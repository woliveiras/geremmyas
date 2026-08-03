package cli

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestRunGlobalClearDryRunDoesNotMutateFilesOrManifest(t *testing.T) {
	home, state := setupGlobalClearInstall(t, []string{TargetCodex, TargetCursor})
	manifestPath := filepath.Join(state, "geremmyas", "global-manifest.json")
	beforeManifest := testMustRead(t, manifestPath)
	managed := filepath.Join(home, ".codex", "AGENTS.md")
	beforeManaged := testMustRead(t, managed)

	var out strings.Builder
	if code := Run([]string{"global", "clear", "--dry-run", "--json"}, &out, &out); code != 0 {
		t.Fatalf("global clear dry-run exit code = %d, output: %s", code, out.String())
	}
	var report globalClearReport
	if err := json.Unmarshal([]byte(out.String()), &report); err != nil {
		t.Fatalf("decode clear report: %v\n%s", err, out.String())
	}
	if !report.DryRun || report.Summary.Remove == 0 {
		t.Fatalf("dry-run report = %+v", report)
	}
	if got := testMustRead(t, manifestPath); string(got) != string(beforeManifest) {
		t.Fatalf("dry-run changed manifest")
	}
	if got := testMustRead(t, managed); string(got) != string(beforeManaged) {
		t.Fatalf("dry-run changed managed file")
	}
}

func TestRunGlobalClearTargetKeepsSharedSkillsForRemainingTarget(t *testing.T) {
	home, _ := setupGlobalClearInstall(t, []string{TargetCodex, TargetCursor})
	shared := filepath.Join(home, ".agents", "skills", "bugfix-loop", "SKILL.md")
	codex := filepath.Join(home, ".codex", "AGENTS.md")
	cursor := filepath.Join(home, ".cursor", "rules", "skill-bugfix-loop.mdc")

	var out strings.Builder
	if code := Run([]string{"global", "clear", "--targets", "codex", "--json"}, &out, &out); code != 0 {
		t.Fatalf("target clear exit code = %d, output: %s", code, out.String())
	}
	mustNotExist(t, codex)
	mustExist(t, cursor)
	mustExist(t, shared)
	manifest, exists, err := loadGlobalManifest()
	if err != nil || !exists {
		t.Fatalf("load manifest: exists=%t err=%v", exists, err)
	}
	if strings.Join(manifest.Targets, ",") != TargetCursor {
		t.Fatalf("remaining targets = %v", manifest.Targets)
	}
	if _, ok := manifest.Files[shared]; !ok {
		t.Fatalf("shared skill lost ownership")
	}
}

func TestRunGlobalClearPreservesModifiedUnlessForced(t *testing.T) {
	home, _ := setupGlobalClearInstall(t, []string{TargetCodex})
	modified := filepath.Join(home, ".agents", "skills", "bugfix-loop", "SKILL.md")
	if err := os.WriteFile(modified, []byte("custom\n"), 0o644); err != nil {
		t.Fatalf("modify skill: %v", err)
	}

	var out strings.Builder
	if code := Run([]string{"global", "clear", "--json"}, &out, &out); code != 0 {
		t.Fatalf("clear exit code = %d, output: %s", code, out.String())
	}
	mustExist(t, modified)
	manifest, _, err := loadGlobalManifest()
	if err != nil {
		t.Fatalf("load preserved manifest: %v", err)
	}
	if _, ok := manifest.Files[modified]; !ok {
		t.Fatalf("modified file ownership was discarded")
	}

	out.Reset()
	if code := Run([]string{"global", "clear", "--force", "--json"}, &out, &out); code != 0 {
		t.Fatalf("forced clear exit code = %d, output: %s", code, out.String())
	}
	mustNotExist(t, modified)
}

func TestRunGlobalClearNeverRemovesManagedSymlink(t *testing.T) {
	home, _ := setupGlobalClearInstall(t, []string{TargetCodex})
	managed := filepath.Join(home, ".codex", "AGENTS.md")
	external := filepath.Join(t.TempDir(), "external.md")
	if err := os.WriteFile(external, []byte("external\n"), 0o644); err != nil {
		t.Fatalf("write external: %v", err)
	}
	if err := os.Remove(managed); err != nil {
		t.Fatalf("remove managed: %v", err)
	}
	if err := os.Symlink(external, managed); err != nil {
		t.Fatalf("create symlink: %v", err)
	}

	var out strings.Builder
	if code := Run([]string{"global", "clear", "--force", "--json"}, &out, &out); code != 0 {
		t.Fatalf("clear symlink exit code = %d, output: %s", code, out.String())
	}
	info, err := os.Lstat(managed)
	if err != nil || info.Mode()&os.ModeSymlink == 0 {
		t.Fatalf("managed symlink was not preserved: info=%v err=%v", info, err)
	}
	mustExist(t, external)
}

func TestRunGlobalClearIncludeAdoptableRequiresTotalScopeForSharedFiles(t *testing.T) {
	home := t.TempDir()
	state := t.TempDir()
	t.Setenv("HOME", home)
	t.Setenv("XDG_STATE_HOME", state)
	catalog, err := loadCatalog()
	if err != nil {
		t.Fatalf("load catalog: %v", err)
	}
	packs, err := catalog.Resolve([]string{"sdd"})
	if err != nil {
		t.Fatalf("resolve pack: %v", err)
	}
	if _, err := globalInstallPacksFiltered(packs, true, false); err != nil {
		t.Fatalf("seed adoptable: %v", err)
	}
	adoptable := filepath.Join(home, ".agents", "skills", "bugfix-loop", "SKILL.md")

	var out strings.Builder
	if code := Run([]string{"global", "clear", "--targets", "codex", "--include-adoptable"}, &out, &out); code != 0 {
		t.Fatalf("scoped adoptable clear exit code = %d, output: %s", code, out.String())
	}
	mustExist(t, adoptable)

	out.Reset()
	if code := Run([]string{"global", "clear", "--include-adoptable"}, &out, &out); code != 0 {
		t.Fatalf("total adoptable clear exit code = %d, output: %s", code, out.String())
	}
	mustNotExist(t, adoptable)
	mustNotExist(t, filepath.Join(state, "geremmyas", "global-manifest.json"))
}

func TestRunGlobalClearGeneratedMarkerRequiresForce(t *testing.T) {
	home := t.TempDir()
	state := t.TempDir()
	t.Setenv("HOME", home)
	t.Setenv("XDG_STATE_HOME", state)
	path := filepath.Join(home, ".codex", "AGENTS.md")
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		t.Fatalf("mkdir codex: %v", err)
	}
	if err := os.WriteFile(path, []byte(generatedMarker+"\nCUSTOMIZED USER CONTENT\n"), 0o644); err != nil {
		t.Fatalf("write marker file: %v", err)
	}

	var out strings.Builder
	if code := Run([]string{"global", "clear", "--include-adoptable", "--json"}, &out, &out); code != 0 {
		t.Fatalf("clear marker exit code = %d, output: %s", code, out.String())
	}
	mustExist(t, path)
	var report globalClearReport
	if err := json.Unmarshal([]byte(out.String()), &report); err != nil {
		t.Fatalf("decode report: %v", err)
	}
	if len(report.Entries) != 1 || report.Entries[0].Action != "preserve" || report.Entries[0].Proof != "generated-marker" {
		t.Fatalf("marker plan = %+v", report.Entries)
	}

	out.Reset()
	if code := Run([]string{"global", "clear", "--include-adoptable", "--force"}, &out, &out); code != 0 {
		t.Fatalf("forced marker clear exit code = %d, output: %s", code, out.String())
	}
	mustNotExist(t, path)
}

func TestRunGlobalClearFailsClosedOnCorruptManifest(t *testing.T) {
	home := t.TempDir()
	state := t.TempDir()
	t.Setenv("HOME", home)
	t.Setenv("XDG_STATE_HOME", state)
	manifestPath := filepath.Join(state, "geremmyas", "global-manifest.json")
	if err := os.MkdirAll(filepath.Dir(manifestPath), 0o755); err != nil {
		t.Fatalf("mkdir manifest: %v", err)
	}
	if err := os.WriteFile(manifestPath, []byte("not-json\n"), 0o644); err != nil {
		t.Fatalf("write corrupt manifest: %v", err)
	}

	var out strings.Builder
	if code := Run([]string{"global", "clear"}, &out, &out); code == 0 {
		t.Fatalf("clear accepted corrupt manifest: %s", out.String())
	}
	if got := string(testMustRead(t, manifestPath)); got != "not-json\n" {
		t.Fatalf("clear changed corrupt manifest: %q", got)
	}
}

func TestRunGlobalClearFailsClosedOnIncompatibleManifestVersion(t *testing.T) {
	home := t.TempDir()
	state := t.TempDir()
	t.Setenv("HOME", home)
	t.Setenv("XDG_STATE_HOME", state)
	manifestPath := filepath.Join(state, "geremmyas", "global-manifest.json")
	if err := os.MkdirAll(filepath.Dir(manifestPath), 0o755); err != nil {
		t.Fatalf("mkdir manifest: %v", err)
	}
	data := []byte("{\"version\":999,\"packs\":[],\"targets\":[],\"files\":{}}\n")
	if err := os.WriteFile(manifestPath, data, 0o644); err != nil {
		t.Fatalf("write incompatible manifest: %v", err)
	}

	var out strings.Builder
	if code := Run([]string{"global", "clear", "--force"}, &out, &out); code == 0 {
		t.Fatalf("clear accepted incompatible manifest: %s", out.String())
	}
	if got := testMustRead(t, manifestPath); string(got) != string(data) {
		t.Fatalf("clear changed incompatible manifest")
	}
}

func TestRunGlobalClearForgetsMissingAndPreservesUnowned(t *testing.T) {
	home, _ := setupGlobalClearInstall(t, []string{TargetCodex})
	missing := filepath.Join(home, ".codex", "AGENTS.md")
	if err := os.Remove(missing); err != nil {
		t.Fatalf("remove managed file: %v", err)
	}
	unowned := filepath.Join(home, ".agents", "skills", "external", "SKILL.md")
	if err := os.MkdirAll(filepath.Dir(unowned), 0o755); err != nil {
		t.Fatalf("mkdir unowned: %v", err)
	}
	if err := os.WriteFile(unowned, []byte("external\n"), 0o644); err != nil {
		t.Fatalf("write unowned: %v", err)
	}

	var out strings.Builder
	if code := Run([]string{"global", "clear", "--force", "--json"}, &out, &out); code != 0 {
		t.Fatalf("clear exit code = %d, output: %s", code, out.String())
	}
	mustExist(t, unowned)
	var report globalClearReport
	if err := json.Unmarshal([]byte(out.String()), &report); err != nil {
		t.Fatalf("decode clear report: %v", err)
	}
	if report.Summary.Forget == 0 || report.Summary.Preserve == 0 {
		t.Fatalf("clear summary = %+v", report.Summary)
	}
}

func TestApplyGlobalClearPreservesFileChangedAfterPlanning(t *testing.T) {
	home, _ := setupGlobalClearInstall(t, []string{TargetCodex})
	catalog, err := loadCatalog()
	if err != nil {
		t.Fatalf("load catalog: %v", err)
	}
	report, err := planGlobalClear(catalog, nil, false, false)
	if err != nil {
		t.Fatalf("plan clear: %v", err)
	}
	path := filepath.Join(home, ".codex", "AGENTS.md")
	if err := os.WriteFile(path, []byte("changed after plan\n"), 0o644); err != nil {
		t.Fatalf("change after plan: %v", err)
	}
	if err := applyGlobalClear(&report); err != nil {
		t.Fatalf("apply clear: %v", err)
	}
	mustExist(t, path)
	manifest, _, err := loadGlobalManifest()
	if err != nil {
		t.Fatalf("load manifest: %v", err)
	}
	if _, ok := manifest.Files[path]; !ok {
		t.Fatalf("changed-after-plan file lost ownership")
	}
}

func TestApplyGlobalClearRejectsManifestChangedAfterPlanning(t *testing.T) {
	_, _ = setupGlobalClearInstall(t, []string{TargetCodex})
	catalog, err := loadCatalog()
	if err != nil {
		t.Fatalf("load catalog: %v", err)
	}
	report, err := planGlobalClear(catalog, nil, false, false)
	if err != nil {
		t.Fatalf("plan clear: %v", err)
	}
	manifest, _, err := loadGlobalManifest()
	if err != nil {
		t.Fatalf("load manifest: %v", err)
	}
	manifest.Packs = append(manifest.Packs, "concurrent-change")
	if err := writeGlobalManifest(manifest); err != nil {
		t.Fatalf("write concurrent manifest: %v", err)
	}
	if err := applyGlobalClear(&report); err == nil || !strings.Contains(err.Error(), "changed after planning") {
		t.Fatalf("apply error = %v", err)
	}
	for _, entry := range report.Entries {
		if entry.Action == "remove" {
			mustExist(t, entry.Path)
		}
	}
}

func TestRunGlobalClearAbsentTargetAndMissingManifestAreNoOps(t *testing.T) {
	home, state := setupGlobalClearInstall(t, []string{TargetCursor})
	cursor := filepath.Join(home, ".cursor", "hooks.json")
	before := testMustRead(t, filepath.Join(state, "geremmyas", "global-manifest.json"))

	var out strings.Builder
	if code := Run([]string{"global", "clear", "--targets", "codex", "--json"}, &out, &out); code != 0 {
		t.Fatalf("absent target clear exit code = %d, output: %s", code, out.String())
	}
	mustExist(t, cursor)
	after := testMustRead(t, filepath.Join(state, "geremmyas", "global-manifest.json"))
	if string(after) != string(before) {
		t.Fatalf("absent target clear changed manifest")
	}

	emptyHome := t.TempDir()
	emptyState := t.TempDir()
	t.Setenv("HOME", emptyHome)
	t.Setenv("XDG_STATE_HOME", emptyState)
	out.Reset()
	if code := Run([]string{"global", "clear", "--json"}, &out, &out); code != 0 {
		t.Fatalf("missing manifest clear exit code = %d, output: %s", code, out.String())
	}
	mustNotExist(t, filepath.Join(emptyState, "geremmyas", "global-manifest.json"))
}

func TestRunGlobalClearRespectsGlobalMutationLock(t *testing.T) {
	home, _ := setupGlobalClearInstall(t, []string{TargetCodex})
	managed := filepath.Join(home, ".codex", "AGENTS.md")
	release, err := acquireGlobalMutationLock()
	if err != nil {
		t.Fatalf("acquire lock: %v", err)
	}
	defer release()

	var out strings.Builder
	if code := Run([]string{"global", "clear", "--force"}, &out, &out); code == 0 {
		t.Fatalf("clear ignored mutation lock: %s", out.String())
	}
	if !strings.Contains(out.String(), "another global mutation is in progress") {
		t.Fatalf("lock error missing: %s", out.String())
	}
	mustExist(t, managed)
}

func setupGlobalClearInstall(t *testing.T, targets []string) (string, string) {
	t.Helper()
	home := t.TempDir()
	state := t.TempDir()
	t.Setenv("HOME", home)
	t.Setenv("XDG_STATE_HOME", state)
	var out strings.Builder
	args := []string{"global", "--targets", strings.Join(targets, ","), "core", "sdd"}
	if code := Run(args, &out, &out); code != 0 {
		t.Fatalf("seed global install exit code = %d, output: %s", code, out.String())
	}
	return home, state
}
