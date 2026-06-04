package cli

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestParseMarkdownFrontmatter(t *testing.T) {
	fm, body, err := parseMarkdownFrontmatter(`---
description: "TypeScript rules"
applyTo: "**/*.ts"
---

# Title
`)
	if err != nil {
		t.Fatalf("parseMarkdownFrontmatter returned error: %v", err)
	}
	if fm.get("description") != "TypeScript rules" {
		t.Fatalf("description = %q", fm.get("description"))
	}
	if fm.get("applyTo") != "**/*.ts" {
		t.Fatalf("applyTo = %q", fm.get("applyTo"))
	}
	if !strings.Contains(body, "# Title") {
		t.Fatalf("body = %q", body)
	}
}

func TestParseConfigWithTargets(t *testing.T) {
	cfg, err := parseConfig(strings.NewReader(`
version: 1
packs:
  - core
  - sdd
targets:
  - copilot
  - cursor
  - claude-code
`))
	if err != nil {
		t.Fatalf("parseConfig returned error: %v", err)
	}
	wantTargets := []string{"claude-code", "copilot", "cursor"}
	if strings.Join(cfg.Targets, ",") != strings.Join(wantTargets, ",") {
		t.Fatalf("Targets = %v, want %v", cfg.Targets, wantTargets)
	}
}

func TestParseConfigDefaultsTargetsToCopilot(t *testing.T) {
	cfg, err := parseConfig(strings.NewReader(`
version: 1
packs:
  - core
`))
	if err != nil {
		t.Fatalf("parseConfig returned error: %v", err)
	}
	if len(cfg.Targets) != 1 || cfg.Targets[0] != TargetCopilot {
		t.Fatalf("Targets = %v, want [copilot]", cfg.Targets)
	}
}

func TestValidateTargetsRejectsUnknown(t *testing.T) {
	if err := validateTargets([]string{"vscode"}); err == nil {
		t.Fatal("validateTargets succeeded, want error")
	}
}

func TestFormatConfigIncludesTargets(t *testing.T) {
	got := formatConfig(Config{
		Version: 1,
		Packs:   []string{"sdd", "core"},
		Targets: []string{"cursor", "copilot"},
	})
	if !strings.Contains(got, "targets:\n") {
		t.Fatalf("formatConfig missing targets section:\n%s", got)
	}
	if !strings.Contains(got, "  - copilot\n") || !strings.Contains(got, "  - cursor\n") {
		t.Fatalf("formatConfig targets wrong:\n%s", got)
	}
}

func TestFormatCursorRule(t *testing.T) {
	rule := formatCursorRule("TS rules", "**/*.ts", "# Body\n")
	if !strings.Contains(rule, "globs: **/*.ts") {
		t.Fatalf("rule missing globs:\n%s", rule)
	}
	if !strings.Contains(rule, generatedMarker) {
		t.Fatalf("rule missing generated marker:\n%s", rule)
	}
}

func TestRunSyncGeneratesCursorRules(t *testing.T) {
	root := withTempCwd(t)

	var out strings.Builder
	if code := Run([]string{"init", "--packs", "core,sdd", "--targets", "copilot,cursor"}, &out, &out); code != 0 {
		t.Fatalf("init exit code = %d, output: %s", code, out.String())
	}
	if code := Run([]string{"sync"}, &out, &out); code != 0 {
		t.Fatalf("sync exit code = %d, output: %s", code, out.String())
	}

	mustExist(t, filepath.Join(root, ".cursor/rules/testing.mdc"))
	mustExist(t, filepath.Join(root, ".cursor/hooks.json"))
	mustExist(t, filepath.Join(root, ".github/instructions/testing.instructions.md"))
}

func TestRunSyncGeneratesClaudeAndOpenCode(t *testing.T) {
	root := withTempCwd(t)

	var out strings.Builder
	if code := Run([]string{"init", "--packs", "core,sdd", "--targets", "copilot,claude-code,opencode"}, &out, &out); code != 0 {
		t.Fatalf("init exit code = %d, output: %s", code, out.String())
	}
	if code := Run([]string{"sync"}, &out, &out); code != 0 {
		t.Fatalf("sync exit code = %d, output: %s", code, out.String())
	}

	claudePath := filepath.Join(root, "CLAUDE.md")
	mustExist(t, claudePath)
	data, err := os.ReadFile(claudePath)
	if err != nil {
		t.Fatalf("ReadFile returned error: %v", err)
	}
	if !strings.Contains(string(data), generatedMarker) {
		t.Fatalf("CLAUDE.md missing generated marker")
	}
	mustExist(t, filepath.Join(root, ".opencode/AGENTS.md"))
}
