package cli

import (
	"os"
	"path/filepath"
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

func TestRunGlobalCursorOnlyCopiesSkills(t *testing.T) {
	home := t.TempDir()
	t.Setenv("HOME", home)

	var out strings.Builder
	if code := Run([]string{"global", "--targets", "cursor", "sdd"}, &out, &out); code != 0 {
		t.Fatalf("global exit code = %d, output: %s", code, out.String())
	}

	mustExist(t, filepath.Join(home, ".agents", "skills"))
	if _, err := os.Stat(filepath.Join(home, ".copilot", "instructions")); err == nil {
		t.Fatal("cursor-only global should not install ~/.copilot/instructions")
	}
	mustExist(t, filepath.Join(home, ".cursor", "rules"))
}
