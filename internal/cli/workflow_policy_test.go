package cli

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"testing/fstest"

	geremmyas "github.com/woliveiras/geremmyas"
)

func TestWorkflowGateInventoryCoversEveryCataloguedSurface(t *testing.T) {
	root := workflowRepositoryRoot(t)
	catalog, err := loadCatalog()
	if err != nil {
		t.Fatalf("loadCatalog returned error: %v", err)
	}
	inventory, err := loadWorkflowGateInventory(os.DirFS(root), workflowGateInventoryPath)
	if err != nil {
		t.Fatalf("loadWorkflowGateInventory returned error: %v", err)
	}
	classifications := map[string]bool{}
	for _, rule := range inventory.Rules {
		classifications[rule.Classification] = true
	}
	for _, classification := range []string{
		"removed", "retained-authority-boundary", "residual-evidence", "deferred-release-work",
	} {
		if !classifications[classification] {
			t.Errorf("workflow gate inventory has no %q classification", classification)
		}
	}
	surfaces, err := collectWorkflowSurfaceFiles(root, catalog)
	if err != nil {
		t.Fatalf("collectWorkflowSurfaceFiles returned error: %v", err)
	}
	matches, violations, err := scanWorkflowGates(os.DirFS(root), surfaces, inventory)
	if err != nil {
		t.Fatalf("scanWorkflowGates returned error: %v", err)
	}
	if len(violations) != 0 {
		for _, violation := range violations {
			t.Errorf("%s:%d: %s", violation.Path, violation.Line, violation.Message)
		}
	}
	if len(matches) == 0 {
		t.Fatal("workflow gate inventory matched no conversational gates")
	}
}

func TestAutonomousFeaturePolicyMaterializesForEveryTarget(t *testing.T) {
	repositoryRoot := workflowRepositoryRoot(t)
	root := withTempCwd(t)
	var out strings.Builder
	if code := Run([]string{"init", "--packs", "core,sdd", "--targets", "codex,copilot,cursor,claude-code,opencode"}, &out, &out); code != 0 {
		t.Fatalf("init exit code = %d, output: %s", code, out.String())
	}
	if code := Run([]string{"sync"}, &out, &out); code != 0 {
		t.Fatalf("sync exit code = %d, output: %s", code, out.String())
	}

	policyFiles := []string{
		"AGENTS.md",
		"CLAUDE.md",
		".opencode/AGENTS.md",
		".agents/skills/generate-spec/SKILL.md",
		".github/skills/generate-spec/SKILL.md",
		"specs/README.md",
	}
	for _, rel := range policyFiles {
		content := strings.ToLower(string(testMustRead(t, filepath.Join(root, rel))))
		if !strings.Contains(content, "ready") {
			t.Errorf("%s does not materialize the Ready lifecycle", rel)
		}
		if strings.Contains(content, "until the user explicitly approves") ||
			strings.Contains(content, "human approved; implementation may start") {
			t.Errorf("%s retains an obsolete human feature gate", rel)
		}
	}

	codex := string(testMustRead(t, filepath.Join(root, ".codex/AGENTS.md")))
	if !strings.Contains(codex, "Follow `AGENTS.md` at the workspace root") {
		t.Errorf("Codex adapter does not route to the autonomous root contract")
	}
	cursor := string(testMustRead(t, filepath.Join(root, ".cursor/rules/skill-generate-spec.mdc")))
	if !strings.Contains(cursor, ".agents/skills/generate-spec/SKILL.md") {
		t.Errorf("Cursor adapter does not route to the canonical generate-spec skill")
	}

	for _, agent := range []string{
		"architect", "auditor", "documentation", "explorer", "implementer",
		"performance-reviewer", "reviewer", "security-reviewer", "spec-writer", "test-engineer",
	} {
		source := testMustRead(t, filepath.Join(repositoryRoot, "content", "agents", agent+".agent.md"))
		for _, rel := range []string{
			filepath.Join(".agents", "roles", agent+".agent.md"),
			filepath.Join(".github", "agents", agent+".agent.md"),
		} {
			installed := testMustRead(t, filepath.Join(root, rel))
			if string(installed) != string(source) {
				t.Errorf("%s does not exactly materialize %s", rel, agent)
			}
		}
		for _, rel := range []string{
			filepath.Join(".cursor", "agents", agent+".md"),
			filepath.Join(".claude", "agents", agent+".md"),
			filepath.Join(".opencode", "agents", agent+".md"),
		} {
			content := strings.ToLower(string(testMustRead(t, filepath.Join(root, rel))))
			needsName := strings.Contains(rel, ".claude") || strings.Contains(rel, ".cursor")
			if needsName && !strings.Contains(content, "name: "+agent) ||
				!strings.Contains(content, "description:") || !strings.Contains(content, "delegation contract") {
				t.Errorf("%s is not a native specialist definition", rel)
			}
		}
		if strings.Contains(strings.ToLower(string(source)), "execute") {
			opencode := strings.ToLower(string(testMustRead(t, filepath.Join(root, ".opencode", "agents", agent+".md"))))
			for _, policy := range []string{"\"git\": deny", "\"git *\": deny"} {
				if !strings.Contains(opencode, policy) {
					t.Errorf("OpenCode native agent %s missing shell policy %q", agent, policy)
				}
			}
		}
	}
	for _, rel := range []string{"CLAUDE.md", ".opencode/AGENTS.md"} {
		content := strings.ToLower(string(testMustRead(t, filepath.Join(root, rel))))
		if !strings.Contains(content, "native subagents") || !strings.Contains(content, "proactively") {
			t.Errorf("%s does not expose proactive native specialists", rel)
		}
	}

	manifest, exists, err := loadProjectManifest(root)
	if err != nil || !exists {
		t.Fatalf("load native-agent project manifest: exists=%v err=%v", exists, err)
	}
	for _, rel := range []string{
		".cursor/agents/implementer.md",
		".claude/agents/implementer.md",
		".opencode/agents/implementer.md",
	} {
		if _, ok := manifest.Files[rel]; !ok {
			t.Errorf("project manifest does not own native agent %s", rel)
		}
	}

	config := Config{Version: 1, Packs: []string{"core", "sdd"}, Targets: []string{TargetCodex}}
	if err := os.WriteFile(filepath.Join(root, configFileName), []byte(formatConfig(config)), 0o644); err != nil {
		t.Fatalf("write target-change config: %v", err)
	}
	out.Reset()
	if code := Run([]string{"sync"}, &out, &out); code != 0 {
		t.Fatalf("sync after native target removal exit code = %d, output: %s", code, out.String())
	}
	for _, rel := range []string{
		".cursor/agents/implementer.md",
		".claude/agents/implementer.md",
		".opencode/agents/implementer.md",
	} {
		mustNotExist(t, filepath.Join(root, rel))
	}
}

func TestNativeSubagentsAreOwnedByGlobalManifest(t *testing.T) {
	home := t.TempDir()
	t.Setenv("HOME", home)
	t.Setenv("XDG_STATE_HOME", filepath.Join(home, ".state"))
	catalog, err := loadCatalog()
	if err != nil {
		t.Fatalf("loadCatalog: %v", err)
	}
	packs, err := catalog.Resolve([]string{"core", "sdd"})
	if err != nil {
		t.Fatalf("resolve packs: %v", err)
	}
	paths, err := globalDesiredPaths(packs, []string{TargetCursor, TargetClaudeCode, TargetOpenCode})
	if err != nil {
		t.Fatalf("globalDesiredPaths: %v", err)
	}
	want := []string{
		filepath.Join(home, ".cursor", "agents", "implementer.md"),
		filepath.Join(home, ".claude", "agents", "implementer.md"),
		filepath.Join(home, ".config", "opencode", "agents", "implementer.md"),
	}
	got := map[string]bool{}
	for _, path := range paths {
		got[filepath.Clean(path)] = true
	}
	for _, path := range want {
		if !got[filepath.Clean(path)] {
			t.Errorf("global manifest paths missing native agent %s", path)
		}
		if !isManagedGlobalPath(path) {
			t.Errorf("native global agent is outside managed roots: %s", path)
		}
	}
}

func TestSDDPromptIsDocumentedAsCanonicalOnly(t *testing.T) {
	root := workflowRepositoryRoot(t)
	readme := normalizeWorkflowText(string(testMustRead(t, filepath.Join(root, "README.md"))))
	for _, phrase := range []string{
		"canonical examples, not catalog artifacts",
		"geremmyas sync` and `geremmyas global` do not install them",
	} {
		if !strings.Contains(readme, normalizeWorkflowText(phrase)) {
			t.Errorf("README missing canonical-only prompt boundary %q", phrase)
		}
	}
}

func TestBugfixWorkflowRunsAutonomousEvidenceLoop(t *testing.T) {
	surfaces := []string{
		"content/AGENTS.md",
		"content/skills/bugfix-loop/SKILL.md",
	}
	forbidden := []string{
		"approved fix proposal",
		"stop for approval",
		"approval gate",
		"until the user explicitly approves",
		"before the user approves",
	}
	for _, path := range surfaces {
		content := normalizeWorkflowText(mustReadEmbeddedWorkflowFile(t, path))
		for _, phrase := range forbidden {
			if strings.Contains(content, phrase) {
				t.Errorf("%s retains obsolete bugfix gate %q", path, phrase)
			}
		}
	}

	loop := normalizeWorkflowText(mustReadEmbeddedWorkflowFile(t, "content/skills/bugfix-loop/SKILL.md"))
	for _, phrase := range []string{
		"docs/bugfixes/",
		"ranked hypotheses",
		"failed before the fix for the expected reason",
		"before changing production code",
		"actual cause",
		"original reproduction",
		"temporary instrumentation",
		"nearest relevant suite",
		"three consecutive",
	} {
		if !strings.Contains(loop, phrase) {
			t.Errorf("bugfix loop missing autonomous evidence %q", phrase)
		}
	}
}

func TestAutonomousBugfixPolicyMaterializesForEveryTarget(t *testing.T) {
	root := withTempCwd(t)
	var out strings.Builder
	if code := Run([]string{"init", "--packs", "core,sdd", "--targets", "codex,copilot,cursor,claude-code,opencode"}, &out, &out); code != 0 {
		t.Fatalf("init exit code = %d, output: %s", code, out.String())
	}
	if code := Run([]string{"sync"}, &out, &out); code != 0 {
		t.Fatalf("sync exit code = %d, output: %s", code, out.String())
	}

	for _, rel := range []string{
		"AGENTS.md",
		"CLAUDE.md",
		".opencode/AGENTS.md",
	} {
		content := normalizeWorkflowText(string(testMustRead(t, filepath.Join(root, rel))))
		if strings.Contains(content, "stop for approval") || strings.Contains(content, "approval gate") {
			t.Errorf("%s retains an obsolete bugfix approval pause", rel)
		}
		for _, phrase := range []string{
			"reproduce before production edits",
			"rank hypotheses",
			"regression test fails",
			"temporary instrumentation",
			"actual cause",
			"three evidence-driven cycles",
		} {
			if !strings.Contains(content, phrase) {
				t.Errorf("%s missing materialized bugfix invariant %q", rel, phrase)
			}
		}
	}

	canonicalSkill := mustReadEmbeddedWorkflowFile(t, "content/skills/bugfix-loop/SKILL.md")
	for _, rel := range []string{
		".agents/skills/bugfix-loop/SKILL.md",
		".github/skills/bugfix-loop/SKILL.md",
	} {
		materialized := string(testMustRead(t, filepath.Join(root, rel)))
		if materialized != canonicalSkill {
			t.Errorf("%s differs from the canonical autonomous bugfix loop", rel)
		}
	}
}

func TestVerifiedSlicesCommitLocallyByDefault(t *testing.T) {
	surfaces := []string{
		"content/AGENTS.md",
		"content/agents/spec-writer.agent.md",
		"content/prompts/sdd.prompt.md",
		"content/skills/git-commit/SKILL.md",
		"content/skills/requirements-interview/SKILL.md",
	}
	forbidden := []string{
		"commit permission (first)",
		"do not skip the commit permission question",
		"ask the user which exact files",
		"files explicitly approved by the user",
		"ask for explicit confirmation before running `git commit`",
		"run `git commit` only after confirmation",
		"only if i granted commit permission",
		"without explicit permission and confirmation",
	}
	for _, path := range surfaces {
		content := normalizeWorkflowText(mustReadEmbeddedWorkflowFile(t, path))
		for _, phrase := range forbidden {
			if strings.Contains(content, phrase) {
				t.Errorf("%s retains obsolete commit gate %q", path, phrase)
			}
		}
	}

	commitSkill := normalizeWorkflowText(mustReadEmbeddedWorkflowFile(t, "content/skills/git-commit/SKILL.md"))
	for _, phrase := range []string{
		"local commits are the default",
		"no-commits",
		"git status --short",
		"task-owned files or hunks",
		"git diff --cached",
		"conventional commits",
		"tests",
		"documentation",
		"do not push",
		"amend",
		"rebase",
		"merge",
		"tag",
		"release",
		"publication",
		"inspect staged and unstaged state separately",
		"pre-existing staged state",
		"mixed-hunk file",
		"git apply --cached",
		"alternate index",
		"must not advance the current branch",
		"original index remains unchanged",
		"never unstage user-owned state",
		"feature may remain `in progress`",
		"changed files",
		"proposed files",
		"do not wait for the whole feature",
	} {
		if !strings.Contains(commitSkill, phrase) {
			t.Errorf("git-commit skill missing default-commit invariant %q", phrase)
		}
	}

	prompt := normalizeWorkflowText(mustReadEmbeddedWorkflowFile(t, "content/prompts/sdd.prompt.md"))
	for _, phrase := range []string{
		"after each task-owned slice",
		"feature remains `in progress` while more tasks remain",
		"local commit per slice",
		"do not wait for the whole feature",
		"report changed or proposed files",
		"final lifecycle reconciliation",
	} {
		if !strings.Contains(prompt, phrase) {
			t.Errorf("SDD prompt missing slice/feature separation %q", phrase)
		}
	}
}

func TestDefaultCommitPolicyMaterializesForEveryTarget(t *testing.T) {
	root := withTempCwd(t)
	var out strings.Builder
	if code := Run([]string{"init", "--packs", "core,sdd", "--targets", "codex,copilot,cursor,claude-code,opencode"}, &out, &out); code != 0 {
		t.Fatalf("init exit code = %d, output: %s", code, out.String())
	}
	if code := Run([]string{"sync"}, &out, &out); code != 0 {
		t.Fatalf("sync exit code = %d, output: %s", code, out.String())
	}

	for _, rel := range []string{"AGENTS.md", "CLAUDE.md", ".opencode/AGENTS.md"} {
		content := normalizeWorkflowText(string(testMustRead(t, filepath.Join(root, rel))))
		for _, phrase := range []string{
			"atomic local conventional commit by default",
			"feature remains `in progress`",
			"no-commits",
			"do not push",
		} {
			if !strings.Contains(content, phrase) {
				t.Errorf("%s missing materialized commit policy %q", rel, phrase)
			}
		}
	}

	canonicalSkill := mustReadEmbeddedWorkflowFile(t, "content/skills/git-commit/SKILL.md")
	for _, rel := range []string{
		".agents/skills/git-commit/SKILL.md",
		".github/skills/git-commit/SKILL.md",
	} {
		if materialized := string(testMustRead(t, filepath.Join(root, rel))); materialized != canonicalSkill {
			t.Errorf("%s differs from canonical git-commit policy", rel)
		}
	}
}

func TestScanWorkflowGatesRejectsUnclassifiedMatch(t *testing.T) {
	testFS := fstest.MapFS{
		"content/AGENTS.md": {Data: []byte("Ask the user for explicit approval before implementation.\n")},
	}
	inventory := workflowGateInventory{Version: 1}

	_, violations, err := scanWorkflowGates(testFS, []string{"content/AGENTS.md"}, inventory)
	if err != nil {
		t.Fatalf("scanWorkflowGates returned error: %v", err)
	}
	if len(violations) != 1 || !strings.Contains(violations[0].Message, "unclassified") {
		t.Fatalf("violations = %#v, want one unclassified conversational gate", violations)
	}
}

func TestScanWorkflowGatesRejectsRemovedGateAndStaleActiveRule(t *testing.T) {
	testFS := fstest.MapFS{
		"content/AGENTS.md": {Data: []byte("Stop at the approval gate.\n")},
	}
	inventory := workflowGateInventory{Version: 1, Rules: []workflowGateRule{
		{ID: "removed-feature-gate", Classification: "removed", Path: "content/AGENTS.md", Pattern: "(?i)approval gate"},
		{ID: "stale-production-boundary", Classification: "retained-authority-boundary", Path: "content/AGENTS.md", Pattern: "(?i)production authorization"},
	}}

	_, violations, err := scanWorkflowGates(testFS, []string{"content/AGENTS.md"}, inventory)
	if err != nil {
		t.Fatalf("scanWorkflowGates returned error: %v", err)
	}
	if len(violations) != 2 {
		t.Fatalf("violations = %#v, want removed-gate and stale-rule violations", violations)
	}
	joined := violations[0].Message + "\n" + violations[1].Message
	for _, phrase := range []string{"reappeared", "matched no surface"} {
		if !strings.Contains(joined, phrase) {
			t.Errorf("violations missing %q: %s", phrase, joined)
		}
	}
}

func TestScanWorkflowGatesDoesNotHideRemovedGateBehindRetainedRule(t *testing.T) {
	testFS := fstest.MapFS{
		"README.md": {Data: []byte("Ask before changing the workflow.\n")},
	}
	inventory := workflowGateInventory{Version: 1, Rules: []workflowGateRule{
		{ID: "broad-retained", Classification: "retained-authority-boundary", Path: "README.md", Pattern: "(?i)ask"},
		{ID: "removed-workflow-pause", Classification: "removed", Path: "README.md", Pattern: "(?i)ask before"},
	}}

	_, violations, err := scanWorkflowGates(testFS, []string{"README.md"}, inventory)
	if err != nil {
		t.Fatalf("scanWorkflowGates returned error: %v", err)
	}
	joined := ""
	for _, violation := range violations {
		joined += violation.Message + "\n"
	}
	for _, phrase := range []string{"reappeared", "ambiguous workflow gate classifications"} {
		if !strings.Contains(joined, phrase) {
			t.Errorf("overlapping rules did not report %q: %s", phrase, joined)
		}
	}
}

func TestFeatureWorkflowUsesMachineReadyLifecycle(t *testing.T) {
	featureSurfaces := []string{
		"content/AGENTS.md",
		"content/agents/spec-writer.agent.md",
		"content/prompts/sdd.prompt.md",
		"content/skills/generate-spec/SKILL.md",
		"content/skills/generate-spec/references/task-breakdown.md",
		"content/skills/requirements-interview/references/approval-gates.md",
		"content/skills/vertical-tdd/SKILL.md",
		"content/skills/vertical-tdd/references/generate-tests-from-spec.md",
		"content/templates/specs/README.md",
	}
	forbidden := []string{
		"explicitly approves the spec",
		"explicitly approved the spec",
		"stop at the approval gate",
		"before spec approval",
		"until the user explicitly approves",
		"human approved; implementation may start",
		"implementation and feature tests start **only after approved**",
	}

	for _, path := range featureSurfaces {
		path := path
		t.Run(path, func(t *testing.T) {
			data, err := fs.ReadFile(geremmyas.EmbeddedFiles, path)
			if err != nil {
				t.Fatalf("ReadFile(%q): %v", path, err)
			}
			lower := strings.ToLower(string(data))
			for _, phrase := range forbidden {
				if strings.Contains(lower, phrase) {
					t.Errorf("%s retains obsolete human feature gate %q", path, phrase)
				}
			}
		})
	}

	agents := normalizeWorkflowText(mustReadEmbeddedWorkflowFile(t, "content/AGENTS.md"))
	for _, phrase := range []string{
		"machine-ready",
		"read-only, plan-only, no-edits, or no-commits",
		"blast radius",
		"rollback",
		"critical harness evidence",
	} {
		if !strings.Contains(agents, phrase) {
			t.Errorf("content/AGENTS.md missing autonomous feature policy %q", phrase)
		}
	}

	template := mustReadEmbeddedWorkflowFile(t, "content/templates/specs/README.md")
	for _, status := range []string{"Draft", "Ready", "In Progress", "Verified"} {
		if !strings.Contains(template, "**"+status+"**") {
			t.Errorf("spec template missing lifecycle status %q", status)
		}
	}
	for _, phrase := range []string{
		"first implementation task is marked `[~]`",
		"failed evidence",
		"unresolved material decision",
		"independent review",
	} {
		if !strings.Contains(strings.ToLower(template), strings.ToLower(phrase)) {
			t.Errorf("spec template missing transition rule %q", phrase)
		}
	}
	for _, legacy := range []string{"In Review", "Approved", "Implemented", "Completed", "Deprecated"} {
		if !strings.Contains(template, "`"+legacy+"`") {
			t.Errorf("spec template does not preserve historical status %q", legacy)
		}
	}
	if !strings.Contains(strings.ToLower(template), "do not bulk") {
		t.Error("spec template does not forbid bulk migration of historical statuses")
	}

	for _, path := range []string{
		"content/skills/generate-spec/SKILL.md",
		"content/skills/requirements-interview/references/approval-gates.md",
		"content/skills/vertical-tdd/SKILL.md",
	} {
		content := normalizeWorkflowText(mustReadEmbeddedWorkflowFile(t, path))
		for _, phrase := range []string{
			"every acceptance criterion",
			"[x]",
			"[~]",
			"docs",
			"independent review",
			"no actionable finding",
		} {
			if !strings.Contains(content, phrase) {
				t.Errorf("%s missing Verified predicate %q", path, phrase)
			}
		}
	}
}

func TestSpecialistSubagentsAreOrchestratedAutonomously(t *testing.T) {
	for _, path := range []string{
		"content/AGENTS.md",
		"content/agents/references/subagent-selection.md",
	} {
		content := normalizeWorkflowText(mustReadEmbeddedWorkflowFile(t, path))
		for _, phrase := range []string{
			"proactively delegate",
			"explicit file, module, or worktree ownership",
			"primary agent owns integration and git",
			"independent review after each slice or verification wave",
			"repair findings and re-review automatically",
			"three consecutive cycles",
			"supported delegation mechanism or inline",
		} {
			if !strings.Contains(content, phrase) {
				t.Errorf("%s missing autonomous delegation policy %q", path, phrase)
			}
		}
	}
	docsData, err := os.ReadFile(filepath.Join(workflowRepositoryRoot(t), "docs/guardrails-framework.md"))
	if err != nil {
		t.Fatalf("read guardrails documentation: %v", err)
	}
	docs := normalizeWorkflowText(string(docsData))
	for _, phrase := range []string{
		"proactively delegate",
		"explicit file, module, or worktree ownership",
		"primary agent owns integration and git",
		"independent review after each slice or verification wave",
		"repair findings and re-review automatically",
		"three consecutive cycles",
		"supported delegation mechanism or inline",
	} {
		if !strings.Contains(docs, phrase) {
			t.Errorf("docs/guardrails-framework.md missing autonomous delegation policy %q", phrase)
		}
	}

	selection := normalizeWorkflowText(mustReadEmbeddedWorkflowFile(t, "content/agents/references/subagent-selection.md"))
	for _, role := range []string{
		"spec-writer", "architect", "implementer", "test-engineer",
		"security-reviewer", "performance-reviewer", "documentation", "reviewer", "auditor",
	} {
		if !strings.Contains(selection, "`"+role+"`") {
			t.Errorf("subagent selection reference missing specialist %q", role)
		}
	}
	for _, obsolete := range []string{
		"delegate independent, read-heavy work only",
		"security-sensitive operations (review with user first)",
		"anything with shared state",
	} {
		if strings.Contains(selection, obsolete) {
			t.Errorf("subagent selection retains obsolete restriction %q", obsolete)
		}
	}
}

func mustReadEmbeddedWorkflowFile(t *testing.T, path string) string {
	t.Helper()
	data, err := fs.ReadFile(geremmyas.EmbeddedFiles, path)
	if err != nil {
		t.Fatalf("ReadFile(%q): %v", path, err)
	}
	return string(data)
}

func normalizeWorkflowText(value string) string {
	return strings.Join(strings.Fields(strings.ToLower(value)), " ")
}

func workflowRepositoryRoot(t *testing.T) string {
	t.Helper()
	dir, err := os.Getwd()
	if err != nil {
		t.Fatalf("Getwd: %v", err)
	}
	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			t.Fatalf("repository root not found from %s", dir)
		}
		dir = parent
	}
}
