package cli

import (
	"fmt"
	"path/filepath"
	"sort"
	"strings"
)

func generateNativeAgents(scope installScope, root, target string, artifacts packArtifacts, opts generatorOptions, summary *generatorSummary) error {
	base, ok := nativeAgentDestination(scope, target)
	if !ok {
		return nil
	}
	for _, source := range artifacts.agents {
		content, err := readEmbeddedSource(source)
		if err != nil {
			return err
		}
		fm, body, err := parseMarkdownFrontmatter(content)
		if err != nil {
			return fmt.Errorf("parse %s: %w", source, err)
		}
		name := strings.TrimSuffix(filepath.Base(source), ".agent.md")
		description := fm.get("description")
		if description == "" {
			description = "geremmyas specialist " + name
		}
		description += " Use proactively when this specialist boundary matches."
		tools := parseAgentTools(fm.get("tools"))

		var rendered string
		switch target {
		case TargetClaudeCode:
			rendered = renderClaudeAgent(name, description, body, tools)
		case TargetOpenCode:
			rendered = renderOpenCodeAgent(description, body, tools)
		case TargetCursor:
			rendered = renderCursorAgent(name, description, body, tools)
		default:
			continue
		}
		if err := writeGeneratedFile(root, filepath.Join(base, name+".md"), []byte(rendered), opts, summary); err != nil {
			return err
		}
	}
	return nil
}

func nativeAgentDestination(scope installScope, target string) (string, bool) {
	switch target {
	case TargetClaudeCode:
		return ".claude/agents", true
	case TargetCursor:
		return ".cursor/agents", true
	case TargetOpenCode:
		if scope == scopeGlobal {
			return ".config/opencode/agents", true
		}
		return ".opencode/agents", true
	default:
		return "", false
	}
}

func parseAgentTools(value string) map[string]bool {
	tools := map[string]bool{}
	value = strings.Trim(strings.TrimSpace(value), "[]")
	for _, item := range strings.Split(value, ",") {
		item = strings.TrimSpace(item)
		if item != "" {
			tools[item] = true
		}
	}
	return tools
}

func renderClaudeAgent(name, description, body string, tools map[string]bool) string {
	mapped := map[string]bool{}
	if tools["read"] {
		mapped["Read"], mapped["Glob"] = true, true
	}
	if tools["search"] {
		mapped["Grep"], mapped["Glob"] = true, true
	}
	if tools["edit"] {
		mapped["Edit"], mapped["Write"] = true, true
	}
	if tools["execute"] {
		mapped["Bash"] = true
	}
	if tools["web"] {
		mapped["WebFetch"], mapped["WebSearch"] = true, true
	}
	if tools["agent"] {
		mapped["Agent"] = true
	}
	names := make([]string, 0, len(mapped))
	for tool := range mapped {
		names = append(names, tool)
	}
	sort.Strings(names)
	return fmt.Sprintf("---\nname: %s\ndescription: %q\ntools: %s\n---\n\n<!-- %s:claude-code-agent -->\n\n%s\n", name, description, strings.Join(names, ", "), generatedMarker, strings.TrimSpace(body))
}

func renderOpenCodeAgent(description, body string, tools map[string]bool) string {
	permission := func(enabled bool) string {
		if enabled {
			return "allow"
		}
		return "deny"
	}
	var bashPolicy string
	if tools["execute"] {
		bashPolicy = "\n  bash:\n    \"*\": allow\n    \"git\": deny\n    \"git *\": deny"
	} else {
		bashPolicy = "\n  bash: deny"
	}
	return fmt.Sprintf("---\ndescription: %q\nmode: subagent\npermission:\n  edit: %s%s\n  task: deny\n  external_directory: deny\n  webfetch: %s\n  websearch: %s\n---\n\n<!-- %s:opencode-agent -->\n\n%s\n",
		description, permission(tools["edit"]), bashPolicy, permission(tools["web"]), permission(tools["web"]), generatedMarker, strings.TrimSpace(body))
}

func renderCursorAgent(name, description, body string, tools map[string]bool) string {
	return fmt.Sprintf("---\nname: %s\ndescription: %q\nmodel: inherit\nreadonly: %t\n---\n\n<!-- %s:cursor-agent -->\n\n%s\n",
		name, description, !tools["edit"], generatedMarker, strings.TrimSpace(body))
}
