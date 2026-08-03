package cli

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestSyncPacksCopiesOnlySelectedPacks(t *testing.T) {
	catalog, err := loadCatalog()
	if err != nil {
		t.Fatalf("loadCatalog returned error: %v", err)
	}
	packs, err := catalog.Resolve([]string{"python-api"})
	if err != nil {
		t.Fatalf("Resolve returned error: %v", err)
	}

	root := t.TempDir()
	summary, err := syncPacks(root, packs, syncOptions{})
	if err != nil {
		t.Fatalf("syncPacks returned error: %v", err)
	}
	if summary.Installed == 0 {
		t.Fatalf("Installed = %d, want > 0", summary.Installed)
	}

	mustExist(t, filepath.Join(root, ".github/instructions/python.instructions.md"))
	mustExist(t, filepath.Join(root, ".github/instructions/fastapi.instructions.md"))
	mustExist(t, filepath.Join(root, ".github/instructions/pydantic.instructions.md"))
	mustNotExist(t, filepath.Join(root, ".github/instructions/go.instructions.md"))
}

func TestRunSyncCodexOnlyOmitsCopilotArtifacts(t *testing.T) {
	root := withTempCwd(t)

	var out strings.Builder
	if code := Run([]string{"init", "--packs", "core,base", "--targets", "codex"}, &out, &out); code != 0 {
		t.Fatalf("init exit code = %d, output: %s", code, out.String())
	}
	if code := Run([]string{"sync"}, &out, &out); code != 0 {
		t.Fatalf("sync exit code = %d, output: %s", code, out.String())
	}

	mustExist(t, filepath.Join(root, "AGENTS.md"))
	mustExist(t, filepath.Join(root, ".agents", "skills", "tdd", "SKILL.md"))
	mustExist(t, filepath.Join(root, ".codex", "instructions", "testing.instructions.md"))
	mustExist(t, filepath.Join(root, ".codex", "AGENTS.md"))
	mustNotExist(t, filepath.Join(root, ".github", "copilot-instructions.md"))
	mustNotExist(t, filepath.Join(root, ".github", "agents"))
	mustNotExist(t, filepath.Join(root, ".github", "hooks"))
	mustNotExist(t, filepath.Join(root, ".github", "skills"))
	mustNotExist(t, filepath.Join(root, ".github", "instructions"))
}

func TestRunSyncMixedTargetsMaterializesNativeUnion(t *testing.T) {
	root := withTempCwd(t)

	var out strings.Builder
	if code := Run([]string{"init", "--packs", "core,base", "--targets", "codex,copilot"}, &out, &out); code != 0 {
		t.Fatalf("init exit code = %d, output: %s", code, out.String())
	}
	if code := Run([]string{"sync"}, &out, &out); code != 0 {
		t.Fatalf("sync exit code = %d, output: %s", code, out.String())
	}

	mustExist(t, filepath.Join(root, ".agents", "skills", "tdd", "SKILL.md"))
	mustExist(t, filepath.Join(root, ".codex", "instructions", "testing.instructions.md"))
	mustExist(t, filepath.Join(root, ".github", "skills", "tdd", "SKILL.md"))
	mustExist(t, filepath.Join(root, ".github", "instructions", "testing.instructions.md"))
	mustExist(t, filepath.Join(root, ".github", "hooks", "guardrails-rules.txt"))
}

func TestRunSyncRemovingTargetDeletesOnlyUnchangedOwnedFiles(t *testing.T) {
	root := withTempCwd(t)

	var out strings.Builder
	if code := Run([]string{"init", "--packs", "core,base", "--targets", "codex,copilot"}, &out, &out); code != 0 {
		t.Fatalf("init exit code = %d, output: %s", code, out.String())
	}
	if code := Run([]string{"sync"}, &out, &out); code != 0 {
		t.Fatalf("initial sync exit code = %d, output: %s", code, out.String())
	}

	modified := filepath.Join(root, ".github", "copilot-instructions.md")
	if err := os.WriteFile(modified, []byte("custom\n"), 0o644); err != nil {
		t.Fatalf("WriteFile modified Copilot instructions: %v", err)
	}
	config := Config{Version: 1, Packs: []string{"core", "base"}, Targets: []string{TargetCodex}}
	if err := os.WriteFile(filepath.Join(root, configFileName), []byte(formatConfig(config)), 0o644); err != nil {
		t.Fatalf("WriteFile config: %v", err)
	}

	out.Reset()
	if code := Run([]string{"sync"}, &out, &out); code != 0 {
		t.Fatalf("Codex-only sync exit code = %d, output: %s", code, out.String())
	}

	if got := string(testMustRead(t, modified)); got != "custom\n" {
		t.Fatalf("modified Copilot instructions = %q, want preserved custom content", got)
	}
	mustNotExist(t, filepath.Join(root, ".github", "skills", "tdd", "SKILL.md"))
	mustNotExist(t, filepath.Join(root, ".github", "instructions", "testing.instructions.md"))
	if !strings.Contains(out.String(), "project reconcile:") ||
		!strings.Contains(out.String(), "preserved=1") {
		t.Fatalf("sync output missing reconciliation summary:\n%s", out.String())
	}
}

func TestRunSyncDoesNotFollowOrDeleteUnownedSymlink(t *testing.T) {
	root := withTempCwd(t)
	external := filepath.Join(t.TempDir(), "external.txt")
	if err := os.WriteFile(external, []byte("external\n"), 0o644); err != nil {
		t.Fatalf("WriteFile external: %v", err)
	}
	link := filepath.Join(root, ".github", "skills", "external")
	if err := os.MkdirAll(filepath.Dir(link), 0o755); err != nil {
		t.Fatalf("MkdirAll symlink parent: %v", err)
	}
	if err := os.Symlink(external, link); err != nil {
		t.Fatalf("Symlink: %v", err)
	}

	var out strings.Builder
	if code := Run([]string{"init", "--packs", "core,base", "--targets", "codex"}, &out, &out); code != 0 {
		t.Fatalf("init exit code = %d, output: %s", code, out.String())
	}
	if code := Run([]string{"sync"}, &out, &out); code != 0 {
		t.Fatalf("sync exit code = %d, output: %s", code, out.String())
	}

	info, err := os.Lstat(link)
	if err != nil {
		t.Fatalf("Lstat unowned symlink: %v", err)
	}
	if info.Mode()&os.ModeSymlink == 0 {
		t.Fatalf("%s is no longer a symlink", link)
	}
	if got := string(testMustRead(t, external)); got != "external\n" {
		t.Fatalf("external content = %q, want unchanged", got)
	}
}

func TestRunSyncDoesNotWriteThroughSelectedTargetSymlink(t *testing.T) {
	root := withTempCwd(t)
	externalRoot := t.TempDir()
	link := filepath.Join(root, ".github", "skills", "tdd")
	if err := os.MkdirAll(filepath.Dir(link), 0o755); err != nil {
		t.Fatalf("MkdirAll symlink parent: %v", err)
	}
	if err := os.Symlink(externalRoot, link); err != nil {
		t.Fatalf("Symlink: %v", err)
	}

	var out strings.Builder
	if code := Run([]string{"init", "--packs", "core,base", "--targets", "copilot"}, &out, &out); code != 0 {
		t.Fatalf("init exit code = %d, output: %s", code, out.String())
	}
	if code := Run([]string{"sync"}, &out, &out); code != 0 {
		t.Fatalf("sync exit code = %d, output: %s", code, out.String())
	}

	mustNotExist(t, filepath.Join(externalRoot, "SKILL.md"))
	info, err := os.Lstat(link)
	if err != nil {
		t.Fatalf("Lstat selected target symlink: %v", err)
	}
	if info.Mode()&os.ModeSymlink == 0 {
		t.Fatalf("%s is no longer a symlink", link)
	}
	manifest, _, err := loadProjectManifest(root)
	if err != nil {
		t.Fatalf("load project manifest: %v", err)
	}
	if _, owned := manifest.Files[".github/skills/tdd/SKILL.md"]; owned {
		t.Fatal("project manifest claimed a destination behind a symlink")
	}
}

func TestRunSyncMigratesExactLegacyArtifacts(t *testing.T) {
	root := withTempCwd(t)
	catalog, err := loadCatalog()
	if err != nil {
		t.Fatalf("loadCatalog returned error: %v", err)
	}
	packs, err := catalog.Resolve([]string{"core", "base"})
	if err != nil {
		t.Fatalf("Resolve returned error: %v", err)
	}
	if _, err := syncPacks(root, packs, syncOptions{}); err != nil {
		t.Fatalf("legacy syncPacks returned error: %v", err)
	}
	mustExist(t, filepath.Join(root, ".github", "skills", "tdd", "SKILL.md"))

	config := Config{Version: 1, Packs: []string{"core", "base"}, Targets: []string{TargetCodex}}
	if err := os.WriteFile(filepath.Join(root, configFileName), []byte(formatConfig(config)), 0o644); err != nil {
		t.Fatalf("WriteFile config: %v", err)
	}
	var out strings.Builder
	if code := Run([]string{"sync"}, &out, &out); code != 0 {
		t.Fatalf("migrating sync exit code = %d, output: %s", code, out.String())
	}

	mustNotExist(t, filepath.Join(root, ".github", "skills", "tdd", "SKILL.md"))
	mustExist(t, filepath.Join(root, ".agents", "skills", "tdd", "SKILL.md"))
}

func TestRunSyncReconcilesManifestOwnedNativeAgentsConservatively(t *testing.T) {
	root := withTempCwd(t)
	removedRel := ".cursor/agents/reviewer.md"
	preservedRel := ".cursor/agents/architect.md"
	installed := []byte("<!-- geremmyas:generated:cursor-agent -->\nlegacy agent\n")
	for _, rel := range []string{removedRel, preservedRel} {
		path := filepath.Join(root, filepath.FromSlash(rel))
		if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
			t.Fatalf("create legacy agent parent: %v", err)
		}
		if err := os.WriteFile(path, installed, 0o644); err != nil {
			t.Fatalf("write legacy agent: %v", err)
		}
	}
	manifest := projectManifest{
		Version: projectManifestVersion,
		Packs:   []string{"core", "sdd"},
		Targets: []string{TargetCursor},
		Files: map[string]string{
			removedRel:   bytesSHA256(installed),
			preservedRel: bytesSHA256(installed),
		},
	}
	if err := writeProjectManifest(root, manifest); err != nil {
		t.Fatalf("write legacy manifest: %v", err)
	}
	if err := os.WriteFile(filepath.Join(root, filepath.FromSlash(preservedRel)), []byte("customized legacy agent\n"), 0o644); err != nil {
		t.Fatalf("customize legacy agent: %v", err)
	}
	config := Config{Version: 1, Packs: []string{"core", "base"}, Targets: []string{TargetCursor}}
	if err := os.WriteFile(filepath.Join(root, configFileName), []byte(formatConfig(config)), 0o644); err != nil {
		t.Fatalf("write config: %v", err)
	}

	var out strings.Builder
	if code := Run([]string{"sync"}, &out, &out); code != 0 {
		t.Fatalf("migrating sync exit code = %d, output: %s", code, out.String())
	}
	mustNotExist(t, filepath.Join(root, filepath.FromSlash(removedRel)))
	mustExist(t, filepath.Join(root, filepath.FromSlash(preservedRel)))
	updated, _, err := loadProjectManifest(root)
	if err != nil {
		t.Fatalf("load updated manifest: %v", err)
	}
	if _, exists := updated.Files[removedRel]; exists {
		t.Errorf("removed native agent remains in manifest")
	}
	if updated.Files[preservedRel] != bytesSHA256(installed) {
		t.Errorf("customized native agent lost its original ownership hash")
	}
}

func TestRunSyncRejectsCorruptProjectManifestBeforeCopying(t *testing.T) {
	root := withTempCwd(t)
	if err := os.MkdirAll(filepath.Join(root, ".geremmyas"), 0o755); err != nil {
		t.Fatalf("MkdirAll manifest directory: %v", err)
	}
	if err := os.WriteFile(projectManifestPath(root), []byte("{broken"), 0o644); err != nil {
		t.Fatalf("WriteFile corrupt manifest: %v", err)
	}

	var out strings.Builder
	if code := Run([]string{"init", "--packs", "core,base", "--targets", "codex"}, &out, &out); code != 0 {
		t.Fatalf("init exit code = %d, output: %s", code, out.String())
	}
	if code := Run([]string{"sync"}, &out, &out); code == 0 {
		t.Fatalf("sync corrupt manifest exit code = 0, output: %s", out.String())
	}

	mustNotExist(t, filepath.Join(root, "AGENTS.md"))
	mustNotExist(t, filepath.Join(root, ".agents"))
}

func TestSyncCoreOnlyExcludesStackSkills(t *testing.T) {
	catalog, err := loadCatalog()
	if err != nil {
		t.Fatalf("loadCatalog returned error: %v", err)
	}
	packs, err := catalog.Resolve([]string{"core"})
	if err != nil {
		t.Fatalf("Resolve returned error: %v", err)
	}
	for _, pack := range packs {
		if pack.Tier != TierCore {
			t.Fatalf("core-only resolve pulled non-core pack %q (tier %q)", pack.Name, pack.Tier)
		}
	}

	root := t.TempDir()
	if _, err := syncPacks(root, packs, syncOptions{}); err != nil {
		t.Fatalf("syncPacks returned error: %v", err)
	}

	mustNotExist(t, filepath.Join(root, ".github/skills/premortem"))
	mustNotExist(t, filepath.Join(root, ".github/skills/paper-review"))
}

func TestSyncPacksPreservesCustomizableFiles(t *testing.T) {
	catalog, err := loadCatalog()
	if err != nil {
		t.Fatalf("loadCatalog returned error: %v", err)
	}
	packs, err := catalog.Resolve([]string{"core"})
	if err != nil {
		t.Fatalf("Resolve returned error: %v", err)
	}

	root := t.TempDir()
	agentsPath := filepath.Join(root, "AGENTS.md")
	if err := os.WriteFile(agentsPath, []byte("custom\n"), 0o644); err != nil {
		t.Fatalf("WriteFile returned error: %v", err)
	}

	summary, err := syncPacks(root, packs, syncOptions{})
	if err != nil {
		t.Fatalf("syncPacks returned error: %v", err)
	}
	if summary.Preserved != 1 {
		t.Fatalf("Preserved = %d, want 1", summary.Preserved)
	}

	data, err := os.ReadFile(agentsPath)
	if err != nil {
		t.Fatalf("ReadFile returned error: %v", err)
	}
	if string(data) != "custom\n" {
		t.Fatalf("AGENTS.md = %q, want custom content", string(data))
	}
}

func TestSyncEveryPackIndividually(t *testing.T) {
	catalog, err := loadCatalog()
	if err != nil {
		t.Fatalf("loadCatalog returned error: %v", err)
	}

	for _, pack := range catalog.Packs {
		t.Run(pack.Name, func(t *testing.T) {
			packs, err := catalog.Resolve([]string{pack.Name})
			if err != nil {
				t.Fatalf("Resolve(%q) returned error: %v", pack.Name, err)
			}
			root := t.TempDir()
			if _, err := syncPacks(root, packs, syncOptions{}); err != nil {
				t.Fatalf("syncPacks(%q) returned error: %v", pack.Name, err)
			}
		})
	}
}

func TestRunInitAndAdd(t *testing.T) {
	cwd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Getwd returned error: %v", err)
	}
	root := t.TempDir()
	if err := os.Chdir(root); err != nil {
		t.Fatalf("Chdir returned error: %v", err)
	}
	defer func() {
		if err := os.Chdir(cwd); err != nil {
			t.Fatalf("restore Chdir returned error: %v", err)
		}
	}()

	var out strings.Builder
	if code := Run([]string{"init", "--packs", "core"}, &out, &out); code != 0 {
		t.Fatalf("init exit code = %d, output: %s", code, out.String())
	}
	if code := Run([]string{"add", "python-api"}, &out, &out); code != 0 {
		t.Fatalf("add exit code = %d, output: %s", code, out.String())
	}

	data, err := os.ReadFile(filepath.Join(root, configFileName))
	if err != nil {
		t.Fatalf("ReadFile returned error: %v", err)
	}
	text := string(data)
	if !strings.Contains(text, "  - core\n") || !strings.Contains(text, "  - python-api\n") {
		t.Fatalf("config content missing packs:\n%s", text)
	}
}

func TestRunProjectPersistsTargetsFlag(t *testing.T) {
	root := withTempCwd(t)

	var out strings.Builder
	if code := Run([]string{"init", "--packs", "core"}, &out, &out); code != 0 {
		t.Fatalf("init exit code = %d, output: %s", code, out.String())
	}
	if code := Run([]string{"project", "--targets", "copilot,cursor", "base"}, &out, &out); code != 0 {
		t.Fatalf("project exit code = %d, output: %s", code, out.String())
	}

	data, err := os.ReadFile(filepath.Join(root, configFileName))
	if err != nil {
		t.Fatalf("ReadFile returned error: %v", err)
	}
	text := string(data)
	if !strings.Contains(text, "  - copilot\n") || !strings.Contains(text, "  - cursor\n") {
		t.Fatalf("config missing persisted targets:\n%s", text)
	}

	out.Reset()
	if code := Run([]string{"sync"}, &out, &out); code != 0 {
		t.Fatalf("sync exit code = %d, output: %s", code, out.String())
	}
	mustExist(t, filepath.Join(root, ".cursor/rules/testing.mdc"))
}

func TestRunProjectCodexTargetCopiesReferencedSkills(t *testing.T) {
	root := withTempCwd(t)

	var out strings.Builder
	if code := Run([]string{"init", "--packs", "core", "--targets", "codex"}, &out, &out); code != 0 {
		t.Fatalf("init exit code = %d, output: %s", code, out.String())
	}
	if code := Run([]string{"project", "base"}, &out, &out); code != 0 {
		t.Fatalf("project exit code = %d, output: %s", code, out.String())
	}

	mustExist(t, filepath.Join(root, ".codex", "AGENTS.md"))
	mustExist(t, filepath.Join(root, ".agents", "skills", "bugfix", "SKILL.md"))
}

func TestRunProjectAddsPackAndSyncsFiles(t *testing.T) {
	root := withTempCwd(t)

	var out strings.Builder
	if code := Run([]string{"init", "--packs", "core"}, &out, &out); code != 0 {
		t.Fatalf("init exit code = %d, output: %s", code, out.String())
	}
	if code := Run([]string{"project", "python-api"}, &out, &out); code != 0 {
		t.Fatalf("project exit code = %d, output: %s", code, out.String())
	}

	data, err := os.ReadFile(filepath.Join(root, configFileName))
	if err != nil {
		t.Fatalf("ReadFile returned error: %v", err)
	}
	text := string(data)
	if !strings.Contains(text, "  - core\n") || !strings.Contains(text, "  - python-api\n") {
		t.Fatalf("config content missing packs:\n%s", text)
	}
	mustExist(t, filepath.Join(root, ".github/instructions/python.instructions.md"))
}

func TestRunProjectUnknownPackDoesNotRewriteConfigOrSyncFiles(t *testing.T) {
	root := withTempCwd(t)

	var out strings.Builder
	if code := Run([]string{"init", "--packs", "core"}, &out, &out); code != 0 {
		t.Fatalf("init exit code = %d, output: %s", code, out.String())
	}
	before := string(testMustRead(t, filepath.Join(root, configFileName)))

	if code := Run([]string{"project", "missing-pack"}, &out, &out); code == 0 {
		t.Fatalf("project missing-pack exit code = 0, output: %s", out.String())
	}

	after := string(testMustRead(t, filepath.Join(root, configFileName)))
	if after != before {
		t.Fatalf("config changed after unknown pack:\nbefore:\n%s\nafter:\n%s", before, after)
	}
	mustNotExist(t, filepath.Join(root, ".github", "instructions"))
}

func TestRunProjectPreservesCustomizableFilesByDefault(t *testing.T) {
	root := withTempCwd(t)

	var out strings.Builder
	if code := Run([]string{"init", "--packs", "core"}, &out, &out); code != 0 {
		t.Fatalf("init exit code = %d, output: %s", code, out.String())
	}
	agentsPath := filepath.Join(root, "AGENTS.md")
	if err := os.WriteFile(agentsPath, []byte("custom\n"), 0o644); err != nil {
		t.Fatalf("WriteFile returned error: %v", err)
	}

	if code := Run([]string{"project", "core"}, &out, &out); code != 0 {
		t.Fatalf("project exit code = %d, output: %s", code, out.String())
	}

	data, err := os.ReadFile(agentsPath)
	if err != nil {
		t.Fatalf("ReadFile returned error: %v", err)
	}
	if string(data) != "custom\n" {
		t.Fatalf("AGENTS.md = %q, want custom content", string(data))
	}
}

func TestRunProjectPreservesSpecsReadmeByDefault(t *testing.T) {
	root := withTempCwd(t)

	var out strings.Builder
	if code := Run([]string{"init", "--packs", "base"}, &out, &out); code != 0 {
		t.Fatalf("init exit code = %d, output: %s", code, out.String())
	}
	if code := Run([]string{"project", "base"}, &out, &out); code != 0 {
		t.Fatalf("first project exit code = %d, output: %s", code, out.String())
	}
	readmePath := filepath.Join(root, "specs", "README.md")
	if err := os.WriteFile(readmePath, []byte("# Custom specs index\n"), 0o644); err != nil {
		t.Fatalf("WriteFile returned error: %v", err)
	}

	if code := Run([]string{"project", "base"}, &out, &out); code != 0 {
		t.Fatalf("second project exit code = %d, output: %s", code, out.String())
	}

	data, err := os.ReadFile(readmePath)
	if err != nil {
		t.Fatalf("ReadFile returned error: %v", err)
	}
	if string(data) != "# Custom specs index\n" {
		t.Fatalf("specs/README.md = %q, want custom content", string(data))
	}
}

func TestRunProjectForceOverwritesCustomizableFiles(t *testing.T) {
	root := withTempCwd(t)

	var out strings.Builder
	if code := Run([]string{"init", "--packs", "core"}, &out, &out); code != 0 {
		t.Fatalf("init exit code = %d, output: %s", code, out.String())
	}
	agentsPath := filepath.Join(root, "AGENTS.md")
	if err := os.WriteFile(agentsPath, []byte("custom\n"), 0o644); err != nil {
		t.Fatalf("WriteFile returned error: %v", err)
	}

	if code := Run([]string{"project", "--force", "core"}, &out, &out); code != 0 {
		t.Fatalf("project exit code = %d, output: %s", code, out.String())
	}

	data, err := os.ReadFile(agentsPath)
	if err != nil {
		t.Fatalf("ReadFile returned error: %v", err)
	}
	if string(data) == "custom\n" {
		t.Fatal("AGENTS.md was preserved, want overwritten content")
	}
}

func TestWriteGeneratedFilePreservesCustomFileByDefault(t *testing.T) {
	root := t.TempDir()
	path := filepath.Join(root, ".codex", "AGENTS.md")
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		t.Fatalf("MkdirAll returned error: %v", err)
	}
	if err := os.WriteFile(path, []byte("custom\n"), 0o644); err != nil {
		t.Fatalf("WriteFile returned error: %v", err)
	}

	summary := generatorSummary{}
	if err := writeGeneratedFile(root, ".codex/AGENTS.md", []byte("generated\n"), generatorOptions{}, &summary); err != nil {
		t.Fatalf("writeGeneratedFile returned error: %v", err)
	}

	data := testMustRead(t, path)
	if string(data) != "custom\n" {
		t.Fatalf("generated file overwrite without marker = %q, want custom", data)
	}
	if summary.Preserved != 1 {
		t.Fatalf("Preserved = %d, want 1", summary.Preserved)
	}
}

func TestWriteGeneratedFileUpdatesMarkedFile(t *testing.T) {
	root := t.TempDir()
	path := filepath.Join(root, ".codex", "AGENTS.md")
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		t.Fatalf("MkdirAll returned error: %v", err)
	}
	if err := os.WriteFile(path, []byte("<!-- "+generatedMarker+":codex -->\nold\n"), 0o644); err != nil {
		t.Fatalf("WriteFile returned error: %v", err)
	}

	summary := generatorSummary{}
	if err := writeGeneratedFile(root, ".codex/AGENTS.md", []byte("<!-- "+generatedMarker+":codex -->\nnew\n"), generatorOptions{}, &summary); err != nil {
		t.Fatalf("writeGeneratedFile returned error: %v", err)
	}

	data := testMustRead(t, path)
	if !strings.Contains(string(data), "new") {
		t.Fatalf("generated file was not updated:\n%s", data)
	}
	if summary.Updated != 1 {
		t.Fatalf("Updated = %d, want 1", summary.Updated)
	}
}

func TestWriteGeneratedFileForceOverwritesCustomFile(t *testing.T) {
	root := t.TempDir()
	path := filepath.Join(root, ".codex", "AGENTS.md")
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		t.Fatalf("MkdirAll returned error: %v", err)
	}
	if err := os.WriteFile(path, []byte("custom\n"), 0o644); err != nil {
		t.Fatalf("WriteFile returned error: %v", err)
	}

	summary := generatorSummary{}
	if err := writeGeneratedFile(root, ".codex/AGENTS.md", []byte("generated\n"), generatorOptions{Force: true}, &summary); err != nil {
		t.Fatalf("writeGeneratedFile returned error: %v", err)
	}

	data := testMustRead(t, path)
	if string(data) != "generated\n" {
		t.Fatalf("generated file after force = %q, want generated", data)
	}
	if summary.Updated != 1 {
		t.Fatalf("Updated = %d, want 1", summary.Updated)
	}
}

func withTempCwd(t *testing.T) string {
	t.Helper()
	cwd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Getwd returned error: %v", err)
	}
	root := t.TempDir()
	if err := os.Chdir(root); err != nil {
		t.Fatalf("Chdir returned error: %v", err)
	}
	t.Cleanup(func() {
		if err := os.Chdir(cwd); err != nil {
			t.Fatalf("restore Chdir returned error: %v", err)
		}
	})
	return root
}

func mustExist(t *testing.T, path string) {
	t.Helper()
	if _, err := os.Stat(path); err != nil {
		t.Fatalf("expected %s to exist: %v", path, err)
	}
}

func mustNotExist(t *testing.T, path string) {
	t.Helper()
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		t.Fatalf("expected %s to be absent, stat err: %v", path, err)
	}
}
