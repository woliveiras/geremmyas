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

func TestSyncBragMePackCopiesSkillAssets(t *testing.T) {
	catalog, err := loadCatalog()
	if err != nil {
		t.Fatalf("loadCatalog returned error: %v", err)
	}
	packs, err := catalog.Resolve([]string{"brag-me"})
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

	mustExist(t, filepath.Join(root, ".github/skills/brag-me/SKILL.md"))
	mustExist(t, filepath.Join(root, ".github/skills/brag-me/assets/brag-template.md"))
	mustExist(t, filepath.Join(root, ".github/skills/brag-me/assets/reveal-template.html"))
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
	if code := Run([]string{"project", "--targets", "copilot,cursor", "sdd"}, &out, &out); code != 0 {
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
	if code := Run([]string{"init", "--packs", "sdd"}, &out, &out); code != 0 {
		t.Fatalf("init exit code = %d, output: %s", code, out.String())
	}
	if code := Run([]string{"project", "sdd"}, &out, &out); code != 0 {
		t.Fatalf("first project exit code = %d, output: %s", code, out.String())
	}
	readmePath := filepath.Join(root, "specs", "README.md")
	if err := os.WriteFile(readmePath, []byte("# Custom specs index\n"), 0o644); err != nil {
		t.Fatalf("WriteFile returned error: %v", err)
	}

	if code := Run([]string{"project", "sdd"}, &out, &out); code != 0 {
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
