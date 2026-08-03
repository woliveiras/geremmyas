package cli

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

const legacyGlossaryTemplate = `# Glossary

## [Domain Area]

| Term | Definition | Aliases to Avoid |
|------|-----------|-----------------|
| **Term** | One-sentence definition of what it IS | synonym1, synonym2 |

## Relationships

- A **Term1** belongs to exactly one **Term2**
- A **Term2** can have many **Term3**

## Example Dialogue

> **Dev:** "When a **Term1** does [action], does it create a **Term2**?"
>
> **Domain expert:** "No — a **Term2** is only created when [condition]. A single **Term1** can produce multiple **Term2** if [scenario]."

## Flagged Ambiguities

- "word" was used to mean both **TermA** and **TermB** — these are distinct concepts: [explanation]
`

func TestLegacyArtifactCatalogMapsRemovedSkillsAndAgents(t *testing.T) {
	catalog, err := loadLegacyArtifactCatalog()
	if err != nil {
		t.Fatalf("load legacy catalog: %v", err)
	}
	if len(catalog.Files) == 0 {
		t.Fatal("legacy catalog is empty")
	}

	root := t.TempDir()
	project := map[string]string{}
	if err := addLegacyProjectHashes(project, root); err != nil {
		t.Fatalf("add project hashes: %v", err)
	}
	for _, path := range []string{
		filepath.Join(root, ".agents", "skills", "requirements-interview", "SKILL.md"),
		filepath.Join(root, ".github", "skills", "vertical-tdd", "SKILL.md"),
		filepath.Join(root, ".agents", "roles", "reviewer.agent.md"),
		filepath.Join(root, ".github", "agents", "references", "deep-modules.md"),
	} {
		if project[path] == "" {
			t.Errorf("missing legacy project hash for %s", path)
		}
	}
}

func TestProjectSyncWithoutManifestRemovesOnlyExactLegacyContent(t *testing.T) {
	root := withTempCwd(t)
	exact := filepath.Join(root, ".github", "skills", "generate-glossary", "assets", "glossary-template.md")
	modified := filepath.Join(root, ".agents", "skills", "generate-glossary", "assets", "glossary-template.md")
	for _, path := range []string{exact, modified} {
		if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
			t.Fatalf("create legacy parent: %v", err)
		}
		if err := os.WriteFile(path, []byte(legacyGlossaryTemplate), 0o644); err != nil {
			t.Fatalf("write legacy content: %v", err)
		}
	}
	if err := os.WriteFile(modified, []byte(legacyGlossaryTemplate+"\ncustomized\n"), 0o644); err != nil {
		t.Fatalf("modify legacy content: %v", err)
	}
	if got := bytesSHA256([]byte(legacyGlossaryTemplate)); got != "1c555cefff4ba3d5b0d9c97e623b93eb9480dc1bbef39c02e866d7f5a816aa41" {
		t.Fatalf("legacy fixture hash drifted: %s", got)
	}
	config := Config{Version: 1, Packs: []string{"core", "coding"}, Targets: []string{TargetCopilot}}
	if err := os.WriteFile(filepath.Join(root, configFileName), []byte(formatConfig(config)), 0o644); err != nil {
		t.Fatalf("write config: %v", err)
	}
	var out strings.Builder
	if code := Run([]string{"sync"}, &out, &out); code != 0 {
		t.Fatalf("sync exit code = %d, output: %s", code, out.String())
	}
	mustNotExist(t, exact)
	mustExist(t, modified)
	manifest, _, err := loadProjectManifest(root)
	if err != nil {
		t.Fatalf("load project manifest: %v", err)
	}
	modifiedRel, _ := filepath.Rel(root, modified)
	if _, owned := manifest.Files[filepath.ToSlash(modifiedRel)]; owned {
		t.Fatal("modified legacy content was silently adopted")
	}
}

func TestGlobalClearRecognizesExactRenamedLegacyContent(t *testing.T) {
	home := t.TempDir()
	t.Setenv("HOME", home)
	t.Setenv("XDG_STATE_HOME", t.TempDir())
	path := filepath.Join(home, ".agents", "skills", "generate-glossary", "assets", "glossary-template.md")
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		t.Fatalf("create legacy global parent: %v", err)
	}
	if err := os.WriteFile(path, []byte(legacyGlossaryTemplate), 0o644); err != nil {
		t.Fatalf("write legacy global content: %v", err)
	}
	var out strings.Builder
	if code := Run([]string{"global", "clear", "--include-adoptable"}, &out, &out); code != 0 {
		t.Fatalf("global clear exit code = %d, output: %s", code, out.String())
	}
	mustNotExist(t, path)
}

func TestLegacyProjectAdoptionAndReconcileDoNotFollowParentSymlink(t *testing.T) {
	root := t.TempDir()
	external := t.TempDir()
	externalFile := filepath.Join(external, "generate-glossary", "assets", "glossary-template.md")
	if err := os.MkdirAll(filepath.Dir(externalFile), 0o755); err != nil {
		t.Fatalf("create external legacy parent: %v", err)
	}
	if err := os.WriteFile(externalFile, []byte(legacyGlossaryTemplate), 0o644); err != nil {
		t.Fatalf("write external legacy content: %v", err)
	}
	if err := os.MkdirAll(filepath.Join(root, ".github"), 0o755); err != nil {
		t.Fatalf("create project parent: %v", err)
	}
	if err := os.Symlink(external, filepath.Join(root, ".github", "skills")); err != nil {
		t.Fatalf("create project parent symlink: %v", err)
	}
	catalog, err := loadCatalog()
	if err != nil {
		t.Fatalf("load catalog: %v", err)
	}
	empty := projectManifest{Version: projectManifestVersion, Files: map[string]string{}}
	adopted, err := adoptKnownLegacyProjectFiles(root, empty, catalog)
	if err != nil {
		t.Fatalf("adopt legacy project files: %v", err)
	}
	if len(adopted.Files) != 0 {
		t.Fatalf("adopted external legacy file through symlink: %+v", adopted.Files)
	}
	rel := ".github/skills/generate-glossary/assets/glossary-template.md"
	previous := projectManifest{Version: projectManifestVersion, Files: map[string]string{rel: bytesSHA256([]byte(legacyGlossaryTemplate))}}
	summary, err := reconcileProjectManifest(root, previous, map[string]string{}, nil, nil)
	if err != nil {
		t.Fatalf("reconcile symlinked legacy project file: %v", err)
	}
	if summary.Preserved != 1 {
		t.Fatalf("project symlink reconciliation summary = %+v", summary)
	}
	mustExist(t, externalFile)
}

func TestLegacyGlobalAdoptionAndReconcileDoNotFollowParentSymlink(t *testing.T) {
	home := t.TempDir()
	external := t.TempDir()
	t.Setenv("HOME", home)
	t.Setenv("XDG_STATE_HOME", t.TempDir())
	externalFile := filepath.Join(external, "generate-glossary", "assets", "glossary-template.md")
	if err := os.MkdirAll(filepath.Dir(externalFile), 0o755); err != nil {
		t.Fatalf("create external legacy parent: %v", err)
	}
	if err := os.WriteFile(externalFile, []byte(legacyGlossaryTemplate), 0o644); err != nil {
		t.Fatalf("write external legacy content: %v", err)
	}
	if err := os.MkdirAll(filepath.Join(home, ".agents"), 0o755); err != nil {
		t.Fatalf("create global parent: %v", err)
	}
	if err := os.Symlink(external, filepath.Join(home, ".agents", "skills")); err != nil {
		t.Fatalf("create global parent symlink: %v", err)
	}
	catalog, err := loadCatalog()
	if err != nil {
		t.Fatalf("load catalog: %v", err)
	}
	empty := globalManifest{Version: globalManifestVersion, Files: map[string]string{}}
	adopted, err := adoptKnownGlobalFiles(empty, catalog)
	if err != nil {
		t.Fatalf("adopt legacy global files: %v", err)
	}
	if len(adopted.Files) != 0 {
		t.Fatalf("adopted external global file through symlink: %+v", adopted.Files)
	}
	managedPath := filepath.Join(home, ".agents", "skills", "generate-glossary", "assets", "glossary-template.md")
	previous := globalManifest{Version: globalManifestVersion, Files: map[string]string{managedPath: bytesSHA256([]byte(legacyGlossaryTemplate))}}
	summary, err := reconcileGlobalManifest(previous, nil, nil, nil)
	if err != nil {
		t.Fatalf("reconcile symlinked legacy global file: %v", err)
	}
	if summary.Preserved != 1 {
		t.Fatalf("global symlink reconciliation summary = %+v", summary)
	}
	mustExist(t, externalFile)
}

func TestGlobalInstallFailsClosedOnSkillsParentSymlink(t *testing.T) {
	home := t.TempDir()
	external := t.TempDir()
	t.Setenv("HOME", home)
	t.Setenv("XDG_STATE_HOME", t.TempDir())
	if err := os.MkdirAll(filepath.Join(home, ".agents"), 0o755); err != nil {
		t.Fatalf("create global parent: %v", err)
	}
	if err := os.Symlink(external, filepath.Join(home, ".agents", "skills")); err != nil {
		t.Fatalf("create global skills symlink: %v", err)
	}
	var out strings.Builder
	if code := Run([]string{"global", "--targets", "codex", "base"}, &out, &out); code == 0 {
		t.Fatalf("global install followed parent symlink: %s", out.String())
	}
	entries, err := os.ReadDir(external)
	if err != nil {
		t.Fatalf("read external directory: %v", err)
	}
	if len(entries) != 0 {
		t.Fatalf("global install wrote through parent symlink: %v", entries)
	}
}

func TestLegacyArtifactCatalogMapsOnlySkillsGlobally(t *testing.T) {
	home := t.TempDir()
	t.Setenv("HOME", home)
	global := map[string]string{}
	if err := addLegacyGlobalHashes(global); err != nil {
		t.Fatalf("add global hashes: %v", err)
	}
	if global[filepath.Join(home, ".agents", "skills", "bugfix-loop", "SKILL.md")] == "" {
		t.Fatal("missing legacy global skill hash")
	}
	if _, exists := global[filepath.Join(home, ".agents", "roles", "reviewer.agent.md")]; exists {
		t.Fatal("project-only agent was mapped as global content")
	}
}
