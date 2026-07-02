package cli

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"
)

func TestRunGlobalGeneratesCursorRules(t *testing.T) {
	home := t.TempDir()
	t.Setenv("HOME", home)

	var out strings.Builder
	if code := Run([]string{"global", "--targets", "copilot,cursor", "core", "sdd"}, &out, &out); code != 0 {
		t.Fatalf("global exit code = %d, output: %s", code, out.String())
	}

	mustExist(t, filepath.Join(home, ".agents", "skills"))
	mustExist(t, filepath.Join(home, ".copilot", "instructions"))
	mustExist(t, filepath.Join(home, ".cursor", "rules", "testing.mdc"))
	mustExist(t, filepath.Join(home, ".cursor", "hooks.json"))

	data, err := os.ReadFile(filepath.Join(home, ".cursor", "rules", "skill-bugfix-loop.mdc"))
	if err != nil {
		t.Fatalf("ReadFile skill rule: %v", err)
	}
	if !strings.Contains(string(data), "~/.agents/skills/") {
		t.Fatalf("global skill rule should reference ~/.agents/skills/, got:\n%s", data)
	}
}

func TestRunGlobalGeneratesClaudeAndOpenCode(t *testing.T) {
	home := t.TempDir()
	t.Setenv("HOME", home)

	var out strings.Builder
	if code := Run([]string{"global", "--targets", "claude-code,opencode", "core", "sdd"}, &out, &out); code != 0 {
		t.Fatalf("global exit code = %d, output: %s", code, out.String())
	}

	claudePath := filepath.Join(home, ".claude", "CLAUDE.md")
	mustExist(t, claudePath)
	data, err := os.ReadFile(claudePath)
	if err != nil {
		t.Fatalf("ReadFile CLAUDE.md: %v", err)
	}
	if !strings.Contains(string(data), generatedMarker) {
		t.Fatalf("CLAUDE.md missing generated marker")
	}
	if !strings.Contains(string(data), "~/.agents/skills/") {
		t.Fatalf("global CLAUDE.md should reference ~/.agents/skills/")
	}
	mustExist(t, filepath.Join(home, ".config", "opencode", "AGENTS.md"))
}

func TestRunGlobalGeneratedAgentTargetsCopySkills(t *testing.T) {
	for _, target := range []string{TargetClaudeCode, TargetCodex, TargetOpenCode} {
		t.Run(target, func(t *testing.T) {
			home := t.TempDir()
			t.Setenv("HOME", home)

			var out strings.Builder
			if code := Run([]string{"global", "--targets", target, "core", "sdd"}, &out, &out); code != 0 {
				t.Fatalf("global exit code = %d, output: %s", code, out.String())
			}

			mustExist(t, filepath.Join(home, ".agents", "skills", "bugfix-loop", "SKILL.md"))
		})
	}
}

func TestRunGlobalGeneratedDocsOnlyReferenceInstalledSkills(t *testing.T) {
	targetPaths := map[string]string{
		TargetClaudeCode: filepath.Join(".claude", "CLAUDE.md"),
		TargetCodex:      filepath.Join(".codex", "AGENTS.md"),
		TargetOpenCode:   filepath.Join(".config", "opencode", "AGENTS.md"),
	}

	for target, relPath := range targetPaths {
		t.Run(target, func(t *testing.T) {
			home := t.TempDir()
			t.Setenv("HOME", home)

			var out strings.Builder
			if code := Run([]string{"global", "--targets", target, "core", "sdd"}, &out, &out); code != 0 {
				t.Fatalf("global exit code = %d, output: %s", code, out.String())
			}

			doc := string(testMustRead(t, filepath.Join(home, relPath)))
			for _, skill := range referencedGlobalSkills(doc) {
				mustExist(t, filepath.Join(home, ".agents", "skills", skill, "SKILL.md"))
			}
		})
	}
}

func TestRunGlobalOutputMentionsSkillsForSkillConsumingTargets(t *testing.T) {
	for _, target := range []string{TargetCursor, TargetClaudeCode, TargetCodex, TargetOpenCode} {
		t.Run(target, func(t *testing.T) {
			home := t.TempDir()
			t.Setenv("HOME", home)

			var out strings.Builder
			if code := Run([]string{"global", "--targets", target, "sdd"}, &out, &out); code != 0 {
				t.Fatalf("global exit code = %d, output: %s", code, out.String())
			}

			if !strings.Contains(out.String(), ".agents/skills/") {
				t.Fatalf("global output should mention installed skills path for %s:\n%s", target, out.String())
			}
		})
	}
}

func TestRunGlobalEveryPackIndividually(t *testing.T) {
	catalog, err := loadCatalog()
	if err != nil {
		t.Fatalf("loadCatalog returned error: %v", err)
	}

	for _, pack := range catalog.Packs {
		t.Run(pack.Name, func(t *testing.T) {
			home := t.TempDir()
			t.Setenv("HOME", home)

			var out strings.Builder
			if code := Run([]string{"global", "--targets", "codex", pack.Name}, &out, &out); code != 0 {
				t.Fatalf("global %q exit code = %d, output: %s", pack.Name, code, out.String())
			}
			mustExist(t, filepath.Join(home, ".codex", "AGENTS.md"))
		})
	}
}

func TestRunGlobalUnknownPackWritesNothing(t *testing.T) {
	home := t.TempDir()
	t.Setenv("HOME", home)

	var out strings.Builder
	if code := Run([]string{"global", "--targets", "codex", "missing-pack"}, &out, &out); code == 0 {
		t.Fatalf("global missing-pack exit code = 0, output: %s", out.String())
	}

	mustNotExist(t, filepath.Join(home, ".codex", "AGENTS.md"))
	mustNotExist(t, filepath.Join(home, ".agents", "skills"))
}

func TestRunGlobalCursorOnlyCopiesSkills(t *testing.T) {
	home := t.TempDir()
	t.Setenv("HOME", home)

	var out strings.Builder
	if code := Run([]string{"global", "--targets", "cursor", "sdd"}, &out, &out); code != 0 {
		t.Fatalf("global exit code = %d, output: %s", code, out.String())
	}

	mustExist(t, filepath.Join(home, ".agents", "skills"))
	mustExist(t, filepath.Join(home, ".copilot", "instructions"))
	mustExist(t, filepath.Join(home, ".cursor", "rules"))
}

func TestRunGlobalGeneratesCodexDocument(t *testing.T) {
	home := t.TempDir()
	t.Setenv("HOME", home)

	var out strings.Builder
	if code := Run([]string{"global", "--targets", "codex", "core", "sdd"}, &out, &out); code != 0 {
		t.Fatalf("global exit code = %d, output: %s", code, out.String())
	}

	codexPath := filepath.Join(home, ".codex", "AGENTS.md")
	mustExist(t, codexPath)
	data, err := os.ReadFile(codexPath)
	if err != nil {
		t.Fatalf("ReadFile .codex/AGENTS.md: %v", err)
	}
	content := string(data)
	if !strings.Contains(content, generatedMarker) {
		t.Fatalf("Codex AGENTS.md missing generated marker")
	}
	if !strings.Contains(content, "~/.agents/skills/") {
		t.Fatalf("global Codex AGENTS.md should reference ~/.agents/skills/")
	}
}

func TestGlobalCopyFlagsTargetMatrix(t *testing.T) {
	tests := []struct {
		name             string
		targets          []string
		wantSkills       bool
		wantInstructions bool
	}{
		{name: "copilot", targets: []string{TargetCopilot}, wantSkills: true, wantInstructions: true},
		{name: "cursor", targets: []string{TargetCursor}, wantSkills: true, wantInstructions: true},
		{name: "claude-code", targets: []string{TargetClaudeCode}, wantSkills: true, wantInstructions: true},
		{name: "codex", targets: []string{TargetCodex}, wantSkills: true, wantInstructions: true},
		{name: "opencode", targets: []string{TargetOpenCode}, wantSkills: true, wantInstructions: true},
		{name: "copilot and codex", targets: []string{TargetCopilot, TargetCodex}, wantSkills: true, wantInstructions: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotSkills, gotInstructions := globalCopyFlags(tt.targets)
			if gotSkills != tt.wantSkills || gotInstructions != tt.wantInstructions {
				t.Fatalf("globalCopyFlags(%v) = (%v, %v), want (%v, %v)",
					tt.targets, gotSkills, gotInstructions, tt.wantSkills, tt.wantInstructions)
			}
		})
	}
}

func referencedGlobalSkills(doc string) []string {
	return regexpMatches(`~/\.agents/skills/([^/]+)/SKILL\.md`, doc)
}

func regexpMatches(pattern, text string) []string {
	re := regexp.MustCompile(pattern)
	seen := map[string]bool{}
	var matches []string
	for _, match := range re.FindAllStringSubmatch(text, -1) {
		if len(match) < 2 || seen[match[1]] {
			continue
		}
		if strings.ContainsAny(match[1], "<>*") {
			continue
		}
		seen[match[1]] = true
		matches = append(matches, match[1])
	}
	return matches
}

func TestRunGlobalSDDInstallsGuardrailSkills(t *testing.T) {
	home := t.TempDir()
	t.Setenv("HOME", home)

	var out strings.Builder
	if code := Run([]string{"global", "--targets", "codex", "sdd"}, &out, &out); code != 0 {
		t.Fatalf("global exit code = %d, output: %s", code, out.String())
	}

	for _, skill := range []string{
		"approval-gates-before-implementation",
		"verification-checklists",
		"decision-framework",
		"subagent-selection",
		"agent-rationalization-blocking",
		"abort-criteria",
		"regression-testing",
		"code-review-requesting",
	} {
		t.Run(skill, func(t *testing.T) {
			mustExist(t, filepath.Join(home, ".agents", "skills", skill, "SKILL.md"))
		})
	}
}

func TestRunGlobalCodexIndexesInstructions(t *testing.T) {
	home := t.TempDir()
	t.Setenv("HOME", home)

	var out strings.Builder
	if code := Run([]string{"global", "--targets", "codex", "python-api"}, &out, &out); code != 0 {
		t.Fatalf("global exit code = %d, output: %s", code, out.String())
	}

	data, err := os.ReadFile(filepath.Join(home, ".codex", "AGENTS.md"))
	if err != nil {
		t.Fatalf("ReadFile .codex/AGENTS.md: %v", err)
	}
	content := string(data)
	if !strings.Contains(content, "## Instructions (on demand)") {
		t.Fatalf("Codex AGENTS.md missing Instructions section:\n%s", content)
	}
	if !strings.Contains(content, "~/.codex/instructions/fastapi.instructions.md") {
		t.Fatalf("Codex AGENTS.md should point to ~/.codex/instructions/fastapi.instructions.md:\n%s", content)
	}
	if strings.Contains(content, "~/.copilot/instructions") {
		t.Fatalf("Codex AGENTS.md must not reference ~/.copilot/instructions:\n%s", content)
	}
}

func TestRunGlobalCodexMirrorsInstructionFiles(t *testing.T) {
	home := t.TempDir()
	t.Setenv("HOME", home)

	var out strings.Builder
	if code := Run([]string{"global", "--targets", "codex", "python-api"}, &out, &out); code != 0 {
		t.Fatalf("global exit code = %d, output: %s", code, out.String())
	}

	for _, file := range []string{
		"fastapi.instructions.md",
		"pydantic.instructions.md",
		"python.instructions.md",
	} {
		mustExist(t, filepath.Join(home, ".codex", "instructions", file))
	}
	// Instructions are also copied to the Copilot location unconditionally.
	mustExist(t, filepath.Join(home, ".copilot", "instructions", "fastapi.instructions.md"))
}

func TestRunGlobalCodexOmitsInstructionsWhenNone(t *testing.T) {
	home := t.TempDir()
	t.Setenv("HOME", home)

	var out strings.Builder
	if code := Run([]string{"global", "--targets", "codex", "premortem"}, &out, &out); code != 0 {
		t.Fatalf("global exit code = %d, output: %s", code, out.String())
	}

	data, err := os.ReadFile(filepath.Join(home, ".codex", "AGENTS.md"))
	if err != nil {
		t.Fatalf("ReadFile .codex/AGENTS.md: %v", err)
	}
	if strings.Contains(string(data), "## Instructions (on demand)") {
		t.Fatalf("Codex AGENTS.md should omit Instructions section for a pack with no instructions")
	}
	mustNotExist(t, filepath.Join(home, ".codex", "instructions"))
}

func TestInstructionApplyToLabel(t *testing.T) {
	if got := instructionApplyToLabel(""); got != "all files" {
		t.Fatalf("instructionApplyToLabel(\"\") = %q, want %q", got, "all files")
	}
	if got := instructionApplyToLabel("  "); got != "all files" {
		t.Fatalf("instructionApplyToLabel(whitespace) = %q, want %q", got, "all files")
	}
	if got := instructionApplyToLabel("**/*.py"); got != "**/*.py" {
		t.Fatalf("instructionApplyToLabel(glob) = %q, want %q", got, "**/*.py")
	}
}
