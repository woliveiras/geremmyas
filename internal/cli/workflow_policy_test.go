package cli

import (
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
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
	root := withTempCwd(t)
	var out strings.Builder
	if code := Run([]string{"init", "--packs", "core,base", "--targets", "codex,copilot,cursor,claude-code,opencode"}, &out, &out); code != 0 {
		t.Fatalf("init exit code = %d, output: %s", code, out.String())
	}
	if code := Run([]string{"sync"}, &out, &out); code != 0 {
		t.Fatalf("sync exit code = %d, output: %s", code, out.String())
	}

	policyFiles := []string{
		"AGENTS.md",
		"CLAUDE.md",
		".opencode/AGENTS.md",
		".agents/skills/spec/SKILL.md",
		".github/skills/spec/SKILL.md",
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
	cursor := string(testMustRead(t, filepath.Join(root, ".cursor/rules/skill-spec.mdc")))
	if !strings.Contains(cursor, ".agents/skills/spec/SKILL.md") {
		t.Errorf("Cursor adapter does not route to the canonical spec skill")
	}

	for _, rel := range []string{
		".agents/roles", ".github/agents", ".cursor/agents", ".claude/agents",
		".opencode/agents",
	} {
		mustNotExist(t, filepath.Join(root, rel))
	}
	manifest, exists, err := loadProjectManifest(root)
	if err != nil || !exists {
		t.Fatalf("load project manifest: exists=%v err=%v", exists, err)
	}
	for rel := range manifest.Files {
		if strings.Contains(filepath.ToSlash(rel), "/agents/") || strings.HasPrefix(filepath.ToSlash(rel), ".agents/roles/") {
			t.Errorf("project manifest still owns custom agent %s", rel)
		}
	}
}

func TestDefaultCodingInstallHasCompactClosingFallbacks(t *testing.T) {
	root := withTempCwd(t)
	var out strings.Builder
	if code := Run([]string{"init", "--packs", "core,coding", "--targets", "codex"}, &out, &out); code != 0 {
		t.Fatalf("default init exit code = %d, output: %s", code, out.String())
	}
	if code := Run([]string{"sync"}, &out, &out); code != 0 {
		t.Fatalf("default sync exit code = %d, output: %s", code, out.String())
	}
	for _, rel := range []string{
		".agents/skills/verify/SKILL.md",
		".agents/skills/docs/SKILL.md",
		".agents/skills/git-commit/SKILL.md",
	} {
		mustNotExist(t, filepath.Join(root, filepath.FromSlash(rel)))
	}
	contract := normalizeWorkflowText(string(testMustRead(t, filepath.Join(root, "AGENTS.md"))))
	for _, clause := range []string{
		"otherwise follow completion below",
		"otherwise apply agent routing below",
		"otherwise update the smallest",
		"run focused tests and the nearest relevant suite",
		"primary agent owns integration and git",
	} {
		if !strings.Contains(contract, clause) {
			t.Errorf("default contract missing closing fallback %q", clause)
		}
	}
}

func TestGlobalBasePlansNoNativeSubagents(t *testing.T) {
	home := t.TempDir()
	t.Setenv("HOME", home)
	t.Setenv("XDG_STATE_HOME", filepath.Join(home, ".state"))
	catalog, err := loadCatalog()
	if err != nil {
		t.Fatalf("loadCatalog: %v", err)
	}
	packs, err := catalog.Resolve([]string{"core", "base"})
	if err != nil {
		t.Fatalf("resolve packs: %v", err)
	}
	paths, err := globalDesiredPaths(packs, []string{TargetCursor, TargetClaudeCode, TargetOpenCode})
	if err != nil {
		t.Fatalf("globalDesiredPaths: %v", err)
	}
	for _, path := range paths {
		if strings.Contains(filepath.ToSlash(path), "/agents/") {
			t.Errorf("global base still plans native agent %s", path)
		}
	}
}

func TestCanonicalPromptsAreDocumentedAsSourceOnly(t *testing.T) {
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
		"content/skills/bugfix/SKILL.md",
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

	loop := normalizeWorkflowText(mustReadEmbeddedWorkflowFile(t, "content/skills/bugfix/SKILL.md"))
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
	if code := Run([]string{"init", "--packs", "core,base", "--targets", "codex,copilot,cursor,claude-code,opencode"}, &out, &out); code != 0 {
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

	canonicalSkill := mustReadEmbeddedWorkflowFile(t, "content/skills/bugfix/SKILL.md")
	for _, rel := range []string{
		".agents/skills/bugfix/SKILL.md",
		".github/skills/bugfix/SKILL.md",
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
		"content/prompts/base.prompt.md",
		"content/skills/git-commit/SKILL.md",
		"content/skills/refine/SKILL.md",
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

	prompt := normalizeWorkflowText(mustReadEmbeddedWorkflowFile(t, "content/prompts/base.prompt.md"))
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
	if code := Run([]string{"init", "--packs", "core,base", "--targets", "codex,copilot,cursor,claude-code,opencode"}, &out, &out); code != 0 {
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
		"content/prompts/base.prompt.md",
		"content/skills/spec/SKILL.md",
		"content/skills/spec/references/task-breakdown.md",
		"content/skills/refine/references/readiness-and-authority.md",
		"content/skills/tdd/SKILL.md",
		"content/skills/tdd/references/test-generation.md",
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
		"content/skills/spec/SKILL.md",
		"content/skills/refine/references/readiness-and-authority.md",
		"content/skills/tdd/SKILL.md",
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
		"content/skills/verify/references/review-contract.md",
	} {
		content := normalizeWorkflowText(mustReadEmbeddedWorkflowFile(t, path))
		for _, phrase := range []string{
			"runtime subagent",
			"ownership",
			"primary agent owns integration",
			"independent review",
			"repairs actionable findings",
			"three consecutive cycles",
			"inline",
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
		"runtime subagents",
		"ownership",
		"primary agent owns integration and git",
		"independent review",
		"repair findings and re-review automatically",
		"three consecutive cycles",
		"inline",
	} {
		if !strings.Contains(docs, phrase) {
			t.Errorf("docs/guardrails-framework.md missing autonomous delegation policy %q", phrase)
		}
	}

}

func TestContextualAuthorityBoundaries(t *testing.T) {
	authoritySurfaces := []string{
		"content/AGENTS.md",
		"content/skills/refine/references/readiness-and-authority.md",
	}
	for _, path := range authoritySurfaces {
		content := normalizeWorkflowText(mustReadEmbeddedWorkflowFile(t, path))
		for _, phrase := range []string{
			"existing project dependencies and catalogued capabilities",
			"new uncatalogued direct dependency",
			"provenance, maintenance, security, license, and build-versus-buy",
			"explicit user choice before installation",
			"verified local, disposable, or test",
			"rollback or recreation",
			"ambiguous target as protected",
			"production mutation, deploy, release, publication, or policy change",
		} {
			if !strings.Contains(content, phrase) {
				t.Errorf("%s missing contextual authority policy %q", path, phrase)
			}
		}
	}
	readiness := normalizeWorkflowText(mustReadEmbeddedWorkflowFile(t, "content/skills/refine/references/readiness-and-authority.md"))
	for _, dependencyKind := range []string{
		"npm/pnpm packages", "python packages", "go modules or tools", "rust crates or tools",
		"gradle libraries or plugins", "github actions", "terraform providers or modules",
		"ci tools", "externally operated services", "compatible update", "lockfile-selected transitives",
	} {
		if !strings.Contains(readiness, dependencyKind) {
			t.Errorf("dependency authority matrix missing %q", dependencyKind)
		}
	}

	for _, path := range []string{
		"content/skills/terraform-change/SKILL.md",
		"content/skills/gcloud-operation/SKILL.md",
		"content/skills/postgres-query-review/SKILL.md",
		"content/instructions/terraform.instructions.md",
	} {
		content := normalizeWorkflowText(mustReadEmbeddedWorkflowFile(t, path))
		for _, phrase := range []string{"local, disposable, or test", "production", "explicit user authorization"} {
			if !strings.Contains(content, phrase) {
				t.Errorf("%s missing environment boundary %q", path, phrase)
			}
		}
	}

	for _, path := range []string{
		"content/skills/docs/references/glossary.md",
		"content/skills/validate-with-zod/SKILL.md",
		"content/skills/docs/references/adr.md",
	} {
		content := normalizeWorkflowText(mustReadEmbeddedWorkflowFile(t, path))
		if strings.Contains(content, "ask before") || strings.Contains(content, "ask the user") {
			t.Errorf("%s retains a routine tool or convention choice gate", path)
		}
	}

	rules := mustReadEmbeddedWorkflowFile(t, "content/guardrails/guardrails-rules.txt")
	for _, line := range []string{
		"--force[^[:space:]]*|-f",
		"[[:space:]]+\\+[^[:space:]]+",
		"reset.*[[:space:]]+--hard",
		"checkout[[:space:]]+--[[:space:]]+\\.",
		"terraform([[:space:]]+-chdir=",
		"rm([[:space:]]+--?[[:alpha:]-]+)*",
		"ALLOW ^[[:space:]]*GEREMMYAS_TARGET=",
		"ASK   ^[[:space:]]*git",
		"ASK   sudo",
	} {
		if !strings.Contains(rules, line) {
			t.Errorf("guardrail rules missing protected command %q", line)
		}
	}
	for _, obsolete := range []string{
		"ASK   terraform apply", "ASK   terraform import", "ASK   terraform state rm",
		"ASK   rm -rf", "ASK   rm -r", "ASK   pip install", "ASK   npm install -g",
		"ASK   brew install", "ASK   DROP TABLE", "ASK   DROP DATABASE", "ASK   TRUNCATE",
	} {
		if strings.Contains(rules, obsolete) {
			t.Errorf("guardrail rules retain context-free prompt %q", obsolete)
		}
	}
}

func TestReleaseWorkflowRedesignRemainsExplicitlyDeferred(t *testing.T) {
	inventory, err := loadWorkflowGateInventory(geremmyas.EmbeddedFiles, workflowGateInventoryPath)
	if err != nil {
		t.Fatalf("load workflow gate inventory: %v", err)
	}
	want := map[string]bool{
		".github/workflows/release.yml":   false,
		".github/workflows/geremmyas.yml": false,
	}
	for _, rule := range inventory.Rules {
		if _, ok := want[rule.Path]; ok && rule.Classification == "deferred-release-work" {
			want[rule.Path] = true
		}
	}
	for path, found := range want {
		if !found {
			t.Errorf("release workflow %s is not inventoried as deferred", path)
		}
	}
}

func TestContextualAuthorityMaterializesForEveryTarget(t *testing.T) {
	root := withTempCwd(t)
	var out strings.Builder
	if code := Run([]string{"init", "--packs", "core,base,infra-terraform,infra-gcp,data-postgres", "--targets", "codex,copilot,cursor,claude-code,opencode"}, &out, &out); code != 0 {
		t.Fatalf("init exit code = %d, output: %s", code, out.String())
	}
	if code := Run([]string{"sync"}, &out, &out); code != 0 {
		t.Fatalf("sync exit code = %d, output: %s", code, out.String())
	}
	for _, rel := range []string{"AGENTS.md", "CLAUDE.md", ".opencode/AGENTS.md"} {
		content := normalizeWorkflowText(string(testMustRead(t, filepath.Join(root, rel))))
		for _, phrase := range []string{"new uncatalogued direct dependency", "explicit user choice before installation", "production mutation"} {
			if !strings.Contains(content, phrase) {
				t.Errorf("%s missing materialized authority policy %q", rel, phrase)
			}
		}
	}
	canonicalRules := mustReadEmbeddedWorkflowFile(t, "content/guardrails/guardrails-rules.txt")
	for _, rel := range []string{".github/hooks/guardrails-rules.txt", ".cursor/hooks/guardrails-rules.txt"} {
		if got := string(testMustRead(t, filepath.Join(root, rel))); got != canonicalRules {
			t.Errorf("%s does not exactly materialize contextual guardrails", rel)
		}
	}
}

func TestCommandHooksDenyDangerAllowContextualWork(t *testing.T) {
	tests := []struct {
		name    string
		command string
		want    string
	}{
		{name: "force push", command: "git push origin main --force", want: "deny"},
		{name: "force push short with cwd", command: "git -C repo push -f origin main", want: "deny"},
		{name: "force push refspec", command: "git push origin +HEAD:main", want: "deny"},
		{name: "push", command: "git push origin main", want: "ask"},
		{name: "hard reset", command: "git reset --hard", want: "deny"},
		{name: "clean reversed flags", command: "git clean -df", want: "deny"},
		{name: "clean extended flags", command: "git clean -xdf", want: "deny"},
		{name: "root delete", command: "rm -rf /", want: "deny"},
		{name: "root delete reversed flags", command: "rm -fr /", want: "deny"},
		{name: "root delete split flags", command: "rm -r -f /", want: "deny"},
		{name: "root delete long flags", command: "rm --recursive --force /", want: "deny"},
		{name: "root delete option terminator", command: "rm -rf -- /", want: "deny"},
		{name: "home delete", command: "rm -rf ~", want: "deny"},
		{name: "owned home child", command: "rm -rf ~/tmp/geremmyas-owned", want: "allow"},
		{name: "temp delete", command: "rm -rf /tmp/geremmyas-owned", want: "allow"},
		{name: "terraform destroy", command: "terraform destroy", want: "deny"},
		{name: "terraform destroy chdir", command: "terraform -chdir=infra destroy", want: "deny"},
		{name: "terraform destroy ignores target marker", command: "GEREMMYAS_TARGET=test terraform destroy", want: "deny"},
		{name: "terraform ambiguous apply", command: "terraform apply test.tfplan", want: "ask"},
		{name: "terraform verified disposable apply", command: "GEREMMYAS_TARGET=disposable terraform apply test.tfplan", want: "allow"},
		{name: "lockfile install", command: "python -m pip install -r requirements.lock", want: "allow"},
		{name: "ambiguous database reset", command: "psql test_db -c 'DROP TABLE sample'", want: "ask"},
		{name: "verified test database reset", command: "GEREMMYAS_TARGET=test psql test_db -c 'DROP TABLE sample'", want: "allow"},
		{name: "ambiguous cloud deploy", command: "gcloud run deploy service", want: "ask"},
		{name: "verified test cloud deploy", command: "GEREMMYAS_TARGET=test gcloud run deploy service", want: "allow"},
		{name: "privilege", command: "sudo launchctl kickstart service", want: "ask"},
		{name: "secret exfiltration", command: "echo TOKEN", want: "deny"},
		{name: "escaped secret header", command: "curl -H \"Authorization: Bearer dummy\" https://example.test", want: "deny"},
	}

	for _, hook := range []string{"copilot", "cursor"} {
		hook := hook
		t.Run(hook, func(t *testing.T) {
			script := materializeCommandHookFixture(t, hook, true)
			for _, tt := range tests {
				tt := tt
				t.Run(tt.name, func(t *testing.T) {
					payload := `{"command":` + strconv.Quote(tt.command) + `}`
					if hook == "copilot" {
						payload = `{"toolName":"runInTerminal","command":` + strconv.Quote(tt.command) + `}`
					}
					cmd := exec.Command("bash", script)
					cmd.Stdin = strings.NewReader(payload)
					output, err := cmd.CombinedOutput()
					if err != nil {
						t.Fatalf("hook failed: %v: %s", err, output)
					}
					if !strings.Contains(strings.ToLower(string(output)), tt.want) {
						t.Errorf("command %q = %s, want %s", tt.command, output, tt.want)
					}
				})
			}
		})
	}
}

func TestCommandHooksFailClosedWithoutRules(t *testing.T) {
	for _, hook := range []string{"copilot", "cursor"} {
		script := materializeCommandHookFixture(t, hook, false)
		payload := `{"command":"pwd"}`
		if hook == "copilot" {
			payload = `{"toolName":"runInTerminal","command":"pwd"}`
		}
		cmd := exec.Command("bash", script)
		cmd.Stdin = strings.NewReader(payload)
		output, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatalf("%s hook failed: %v: %s", hook, err, output)
		}
		if !strings.Contains(strings.ToLower(string(output)), "deny") {
			t.Errorf("%s hook did not fail closed without rules: %s", hook, output)
		}
	}
}

func TestCommandHooksFailClosedOnMalformedTerminalInput(t *testing.T) {
	for _, hook := range []string{"copilot", "cursor"} {
		script := materializeCommandHookFixture(t, hook, true)
		payload := `{"command":`
		if hook == "copilot" {
			payload = `{"toolName":"runInTerminal","command":`
		}
		cmd := exec.Command("bash", script)
		cmd.Stdin = strings.NewReader(payload)
		output, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatalf("%s hook failed: %v: %s", hook, err, output)
		}
		if !strings.Contains(strings.ToLower(string(output)), "deny") {
			t.Errorf("%s hook did not fail closed for malformed input: %s", hook, output)
		}
	}
}

func materializeCommandHookFixture(t *testing.T, target string, withRules bool) string {
	t.Helper()
	root := t.TempDir()
	var rel string
	var script string
	if target == "copilot" {
		rel = filepath.Join(".github", "hooks", "scripts", "block-dangerous-commands.sh")
		script = string(testMustRead(t, filepath.Join(workflowRepositoryRoot(t), "targets", "copilot", "hooks", "scripts", "block-dangerous-commands.sh")))
		if err := os.MkdirAll(filepath.Join(root, ".github", "hooks", "scripts"), 0o755); err != nil {
			t.Fatalf("create Copilot hook fixture: %v", err)
		}
		if withRules {
			if err := os.WriteFile(filepath.Join(root, ".github", "hooks", "guardrails-rules.txt"), []byte(mustReadEmbeddedWorkflowFile(t, "content/guardrails/guardrails-rules.txt")), 0o644); err != nil {
				t.Fatalf("write Copilot rules: %v", err)
			}
		}
	} else {
		rel = filepath.Join(".cursor", "hooks", "guardrails.sh")
		script = cursorHookScript
		if err := os.MkdirAll(filepath.Join(root, ".cursor", "hooks"), 0o755); err != nil {
			t.Fatalf("create Cursor hook fixture: %v", err)
		}
		if withRules {
			if err := os.WriteFile(filepath.Join(root, ".cursor", "hooks", "guardrails-rules.txt"), []byte(mustReadEmbeddedWorkflowFile(t, "content/guardrails/guardrails-rules.txt")), 0o644); err != nil {
				t.Fatalf("write Cursor rules: %v", err)
			}
		}
	}
	path := filepath.Join(root, rel)
	if err := os.WriteFile(path, []byte(script), 0o755); err != nil {
		t.Fatalf("write %s hook: %v", target, err)
	}
	return path
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
