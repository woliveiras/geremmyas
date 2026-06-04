package cli

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	geremmyas "github.com/woliveiras/geremmyas"
)

const cursorHookScript = `#!/usr/bin/env bash
# geremmyas:generated:cursor
# Cursor beforeShellExecution hook — guardrails from geremmyas core pack.

set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
RULES_FILE="${SCRIPT_DIR}/guardrails-rules.txt"

input=$(cat)
command=""
if command -v jq >/dev/null 2>&1; then
  command=$(echo "$input" | jq -r '.command // empty')
fi
if [[ -z "$command" ]]; then
  command=$(echo "$input" | grep -o '"command"[[:space:]]*:[[:space:]]*"[^"]*"' | head -1 | sed 's/.*"command"[[:space:]]*:[[:space:]]*"\([^"]*\)".*/\1/')
fi

if [[ -z "$command" || ! -f "$RULES_FILE" ]]; then
  echo '{"permission":"allow"}'
  exit 0
fi

while IFS= read -r line; do
  [[ "$line" =~ ^[[:space:]]*# ]] && continue
  [[ "$line" =~ ^[[:space:]]*$ ]] && continue
  ACTION="${line%%[[:space:]]*}"
  PATTERN="${line#*[[:space:]]}"
  PATTERN="${PATTERN#"${PATTERN%%[![:space:]]*}"}"
  [[ -z "$ACTION" || -z "$PATTERN" ]] && continue
  if echo "$command" | grep -qiE "$PATTERN"; then
    case "$ACTION" in
      BLOCK)
        echo '{"permission":"deny","user_message":"Blocked by geremmyas guardrail.","agent_message":"Command matched blocked pattern."}'
        exit 0
        ;;
      ASK)
        echo '{"permission":"ask","user_message":"Confirm this command (geremmyas guardrail).","agent_message":"Command matched ask pattern."}'
        exit 0
        ;;
    esac
  fi
done < "$RULES_FILE"

echo '{"permission":"allow"}'
exit 0
`

func generateCursor(root string, artifacts packArtifacts, opts generatorOptions) (generatorSummary, error) {
	summary := generatorSummary{}

	for _, source := range artifacts.instructions {
		content, err := readEmbeddedSource(source)
		if err != nil {
			return summary, err
		}
		fm, body, err := parseMarkdownFrontmatter(content)
		if err != nil {
			return summary, fmt.Errorf("parse %s: %w", source, err)
		}
		base := strings.TrimSuffix(filepath.Base(source), ".instructions.md")
		rulePath := ".cursor/rules/" + base + ".mdc"
		rule := formatCursorRule(fm.get("description"), fm.get("applyTo"), body)
		if err := writeGeneratedFile(root, rulePath, []byte(rule), opts, &summary); err != nil {
			return summary, err
		}
	}

	for _, source := range artifacts.skills {
		if err := generateCursorSkillRule(root, source, opts, &summary); err != nil {
			return summary, err
		}
	}

	for _, source := range artifacts.agents {
		if err := generateCursorAgentRule(root, source, opts, &summary); err != nil {
			return summary, err
		}
	}

	if artifacts.hasHooks {
		if err := writeGeneratedFile(root, ".cursor/hooks/guardrails.sh", []byte(cursorHookScript), opts, &summary); err != nil {
			return summary, err
		}
		if err := copyEmbeddedToRoot(root, "project/.github/hooks/guardrails-rules.txt", ".cursor/hooks/guardrails-rules.txt", opts, &summary); err != nil {
			return summary, err
		}
		hooksJSON := `{
  "version": 1,
  "hooks": {
    "beforeShellExecution": [
      {
        "command": ".cursor/hooks/guardrails.sh",
        "failClosed": false
      }
    ]
  }
}
`
		if err := writeGeneratedFile(root, ".cursor/hooks.json", []byte(hooksJSON), opts, &summary); err != nil {
			return summary, err
		}
		_ = osChmod(filepath.Join(root, ".cursor", "hooks", "guardrails.sh"), 0o755)
	}

	agentsRule := formatCursorAgentsIndexRule(artifacts.agents)
	if agentsRule != "" {
		if err := writeGeneratedFile(root, ".cursor/rules/geremmyas-agents.mdc", []byte(agentsRule), opts, &summary); err != nil {
			return summary, err
		}
	}

	return summary, nil
}

func formatCursorRule(description, applyTo, body string) string {
	if description == "" {
		description = "geremmyas instruction rule"
	}
	description = strings.ReplaceAll(description, `"`, `'`)
	body = strings.TrimSpace(body)
	var b strings.Builder
	b.WriteString("---\n")
	b.WriteString("description: \"")
	b.WriteString(description)
	b.WriteString(" (geremmyas)\"\n")
	if strings.TrimSpace(applyTo) != "" {
		b.WriteString("globs: ")
		b.WriteString(strings.TrimSpace(applyTo))
		b.WriteByte('\n')
	} else {
		b.WriteString("alwaysApply: false\n")
	}
	b.WriteString("---\n\n")
	b.WriteString("<!-- ")
	b.WriteString(generatedMarker)
	b.WriteString(":cursor -->\n\n")
	b.WriteString(body)
	if !strings.HasSuffix(body, "\n") {
		b.WriteByte('\n')
	}
	return b.String()
}

func generateCursorSkillRule(root, source string, opts generatorOptions, summary *generatorSummary) error {
	skillMD, err := findSkillMarkdown(source)
	if err != nil {
		return err
	}
	content, err := readEmbeddedSource(skillMD)
	if err != nil {
		return err
	}
	fm, body, err := parseMarkdownFrontmatter(content)
	if err != nil {
		return err
	}
	name := fm.get("name")
	if name == "" {
		name = filepath.Base(filepath.Dir(skillMD))
	}
	target := strings.TrimPrefix(filepathToSlash(source), "project/.github/skills/")
	projectPath := ".github/skills/" + target + "/SKILL.md"
	description := fm.get("description")
	if description == "" {
		description = "geremmyas skill " + name
	}
	ruleBody := fmt.Sprintf(
		"When this skill applies, read `%s` and follow it.\n\n%s",
		projectPath,
		strings.TrimSpace(truncateLines(body, 12)),
	)
	rule := formatCursorRule(description, "", ruleBody)
	rulePath := ".cursor/rules/skill-" + name + ".mdc"
	return writeGeneratedFile(root, rulePath, []byte(rule), opts, summary)
}

func generateCursorAgentRule(root, source string, opts generatorOptions, summary *generatorSummary) error {
	content, err := readEmbeddedSource(source)
	if err != nil {
		return err
	}
	fm, body, err := parseMarkdownFrontmatter(content)
	if err != nil {
		return err
	}
	name := strings.TrimSuffix(filepath.Base(source), ".agent.md")
	description := fm.get("description")
	if description == "" {
		description = "geremmyas agent " + name
	}
	projectPath := ".github/agents/" + filepath.Base(source)
	ruleBody := fmt.Sprintf(
		"Manual role (no @agent in Cursor). Read `%s` when user asks for this role.\n\n%s",
		projectPath,
		strings.TrimSpace(body),
	)
	rule := formatCursorRule(description, "", ruleBody)
	rulePath := ".cursor/rules/agent-" + name + ".mdc"
	return writeGeneratedFile(root, rulePath, []byte(rule), opts, summary)
}

func formatCursorAgentsIndexRule(agentSources []string) string {
	if len(agentSources) == 0 {
		return ""
	}
	var b strings.Builder
	b.WriteString("---\n")
	b.WriteString("description: \"geremmyas agent roles — read matching agent file on demand\"\n")
	b.WriteString("alwaysApply: false\n")
	b.WriteString("---\n\n")
	b.WriteString("<!-- ")
	b.WriteString(generatedMarker)
	b.WriteString(":cursor -->\n\n")
	b.WriteString("# Agent roles\n\nCursor has no @agent picker. When the user names a role, read the file:\n\n")
	for _, source := range agentSources {
		name := strings.TrimSuffix(filepath.Base(source), ".agent.md")
		b.WriteString("- `")
		b.WriteString(name)
		b.WriteString("` → `.github/agents/")
		b.WriteString(filepath.Base(source))
		b.WriteString("`\n")
	}
	b.WriteByte('\n')
	return b.String()
}

func findSkillMarkdown(source string) (string, error) {
	source = filepathToSlash(source)
	candidate := filepath.Join(source, "SKILL.md")
	if _, err := fs.Stat(geremmyas.EmbeddedFiles, candidate); err == nil {
		return candidate, nil
	}
	return "", fmt.Errorf("skill markdown not found under %s", source)
}

func readEmbeddedSource(source string) (string, error) {
	data, err := fs.ReadFile(geremmyas.EmbeddedFiles, filepathToSlash(source))
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func copyEmbeddedToRoot(root, source, target string, opts generatorOptions, summary *generatorSummary) error {
	data, err := fs.ReadFile(geremmyas.EmbeddedFiles, filepathToSlash(source))
	if err != nil {
		return err
	}
	return writeGeneratedFile(root, target, data, opts, summary)
}

func truncateLines(body string, maxLines int) string {
	lines := strings.Split(strings.TrimSpace(body), "\n")
	if len(lines) <= maxLines {
		return body
	}
	return strings.Join(lines[:maxLines], "\n") + "\n\n...(see SKILL.md for full procedure)"
}

func osChmod(path string, mode fs.FileMode) error {
	return os.Chmod(path, mode)
}
