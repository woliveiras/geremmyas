package cli

import (
	"strings"
	"testing"
)

func TestGenericAgentArtifactPlanningRemainsAvailable(t *testing.T) {
	pack := Pack{Name: "synthetic-agent", Files: []FileEntry{{
		Kind: ArtifactAgent, Source: "synthetic/agents", Path: ".",
	}}}
	planned, err := planProjectArtifacts([]Pack{pack}, []string{TargetCopilot, TargetCodex})
	if err != nil {
		t.Fatalf("plan synthetic agent: %v", err)
	}
	got := map[string]bool{}
	for _, entry := range planned {
		got[entry.Target] = true
	}
	for _, path := range []string{".agents/roles", ".github/agents"} {
		if !got[path] {
			t.Errorf("generic agent planner missing %q: %v", path, got)
		}
	}
}

func TestGenericNativeAgentRenderersKeepPermissionBoundaries(t *testing.T) {
	tools := parseAgentTools("[read, search]")
	claude := renderClaudeAgent("reviewer", "Review only", "# Contract", tools)
	if !strings.Contains(claude, "tools: Glob, Grep, Read") || strings.Contains(claude, "Bash") {
		t.Fatalf("Claude renderer expanded read-only tools: %s", claude)
	}
	cursor := renderCursorAgent("reviewer", "Review only", "# Contract", tools)
	if !strings.Contains(cursor, "readonly: true") {
		t.Fatalf("Cursor renderer lost read-only boundary: %s", cursor)
	}
	opencode := renderOpenCodeAgent("Review only", "# Contract", tools)
	for _, clause := range []string{"edit: deny", "bash: deny", "task: deny"} {
		if !strings.Contains(opencode, clause) {
			t.Errorf("OpenCode renderer missing %q: %s", clause, opencode)
		}
	}
}
