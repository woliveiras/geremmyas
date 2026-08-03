package cli

import (
	"encoding/json"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"testing"
)

func TestRunGlobalListJSONClassifiesManagedFilesWithoutWriting(t *testing.T) {
	home := t.TempDir()
	state := t.TempDir()
	t.Setenv("HOME", home)
	t.Setenv("XDG_STATE_HOME", state)

	var installOut strings.Builder
	if code := Run([]string{"global", "--targets", "codex", "sdd"}, &installOut, &installOut); code != 0 {
		t.Fatalf("install global exit code = %d, output: %s", code, installOut.String())
	}

	manifest, _, err := loadGlobalManifest()
	if err != nil {
		t.Fatalf("loadGlobalManifest: %v", err)
	}
	paths := make([]string, 0, len(manifest.Files))
	for path := range manifest.Files {
		paths = append(paths, path)
	}
	sort.Strings(paths)
	if len(paths) < 3 {
		t.Fatalf("global install produced %d managed files, want at least 3", len(paths))
	}
	modifiedPath := paths[0]
	missingPath := paths[1]
	if err := os.WriteFile(modifiedPath, []byte("user modified\n"), 0o644); err != nil {
		t.Fatalf("modify managed file: %v", err)
	}
	if err := os.Remove(missingPath); err != nil {
		t.Fatalf("remove managed file: %v", err)
	}

	manifestPath := filepath.Join(state, "geremmyas", "global-manifest.json")
	before, err := os.ReadFile(manifestPath)
	if err != nil {
		t.Fatalf("read manifest before list: %v", err)
	}

	var out strings.Builder
	if code := Run([]string{"global", "list", "--json"}, &out, &out); code != 0 {
		t.Fatalf("global list exit code = %d, output: %s", code, out.String())
	}

	var report struct {
		SchemaVersion int    `json:"schema_version"`
		Command       string `json:"command"`
		Manifest      struct {
			Path    string   `json:"path"`
			Exists  bool     `json:"exists"`
			Version int      `json:"version"`
			Packs   []string `json:"packs"`
			Targets []string `json:"targets"`
		} `json:"manifest"`
		Summary map[string]int `json:"summary"`
		Files   []struct {
			Path          string `json:"path"`
			Ownership     string `json:"ownership"`
			Status        string `json:"status"`
			InstalledHash string `json:"installed_hash"`
			CurrentHash   string `json:"current_hash,omitempty"`
		} `json:"files"`
	}
	if err := json.Unmarshal([]byte(out.String()), &report); err != nil {
		t.Fatalf("global list output is not JSON: %v\n%s", err, out.String())
	}
	if report.SchemaVersion != 1 || report.Command != "global.list" {
		t.Fatalf("report identity = version %d command %q", report.SchemaVersion, report.Command)
	}
	if !report.Manifest.Exists || report.Manifest.Version != globalManifestVersion {
		t.Fatalf("manifest metadata = %+v", report.Manifest)
	}
	if report.Manifest.Path != manifestPath {
		t.Fatalf("manifest path = %q, want %q", report.Manifest.Path, manifestPath)
	}

	states := map[string]string{}
	for _, file := range report.Files {
		if file.Ownership != "managed" {
			t.Fatalf("file %s ownership = %q, want managed", file.Path, file.Ownership)
		}
		states[file.Path] = file.Status
	}
	if states[modifiedPath] != "modified" {
		t.Fatalf("modified path status = %q", states[modifiedPath])
	}
	if states[missingPath] != "missing" {
		t.Fatalf("missing path status = %q", states[missingPath])
	}
	if report.Summary["modified"] != 1 || report.Summary["missing"] != 1 {
		t.Fatalf("summary = %#v", report.Summary)
	}
	if report.Summary["current"] == 0 {
		t.Fatalf("summary missing current files: %#v", report.Summary)
	}

	after, err := os.ReadFile(manifestPath)
	if err != nil {
		t.Fatalf("read manifest after list: %v", err)
	}
	if string(after) != string(before) {
		t.Fatalf("global list modified manifest")
	}
}

func TestRunGlobalListClassifiesIntactManagedPathOutsideDesiredStateAsObsolete(t *testing.T) {
	home := t.TempDir()
	state := t.TempDir()
	t.Setenv("HOME", home)
	t.Setenv("XDG_STATE_HOME", state)

	var installOut strings.Builder
	if code := Run([]string{"global", "--targets", "codex", "sdd"}, &installOut, &installOut); code != 0 {
		t.Fatalf("install global exit code = %d, output: %s", code, installOut.String())
	}
	manifest, _, err := loadGlobalManifest()
	if err != nil {
		t.Fatalf("loadGlobalManifest: %v", err)
	}
	obsolete := filepath.Join(home, ".codex", "instructions", "obsolete.instructions.md")
	if err := os.MkdirAll(filepath.Dir(obsolete), 0o755); err != nil {
		t.Fatalf("create obsolete parent: %v", err)
	}
	if err := os.WriteFile(obsolete, []byte("obsolete\n"), 0o644); err != nil {
		t.Fatalf("write obsolete file: %v", err)
	}
	hash, err := fileSHA256(obsolete)
	if err != nil {
		t.Fatalf("hash obsolete file: %v", err)
	}
	manifest.Files[obsolete] = hash
	if err := writeGlobalManifest(manifest); err != nil {
		t.Fatalf("write manifest: %v", err)
	}

	var out strings.Builder
	if code := Run([]string{"global", "list", "--json"}, &out, &out); code != 0 {
		t.Fatalf("global list exit code = %d, output: %s", code, out.String())
	}
	var report globalInventoryReport
	if err := json.Unmarshal([]byte(out.String()), &report); err != nil {
		t.Fatalf("decode global list: %v", err)
	}
	for _, file := range report.Files {
		if file.Path == obsolete {
			if file.Status != "obsolete" {
				t.Fatalf("obsolete file status = %q", file.Status)
			}
			if report.Summary.Obsolete != 1 {
				t.Fatalf("obsolete summary = %+v", report.Summary)
			}
			return
		}
	}
	t.Fatalf("obsolete file missing from report")
}

func TestRunGlobalListMissingManifestIsEmptyAndReadOnly(t *testing.T) {
	home := t.TempDir()
	state := t.TempDir()
	t.Setenv("HOME", home)
	t.Setenv("XDG_STATE_HOME", state)

	var out strings.Builder
	if code := Run([]string{"global", "list", "--json"}, &out, &out); code != 0 {
		t.Fatalf("global list exit code = %d, output: %s", code, out.String())
	}
	var report globalInventoryReport
	if err := json.Unmarshal([]byte(out.String()), &report); err != nil {
		t.Fatalf("decode global list: %v", err)
	}
	if report.Manifest.Exists || len(report.Files) != 0 || report.Summary.Managed != 0 {
		t.Fatalf("missing manifest report = %+v", report)
	}
	mustNotExist(t, filepath.Join(state, "geremmyas", "global-manifest.json"))
	entries, err := os.ReadDir(home)
	if err != nil {
		t.Fatalf("read home: %v", err)
	}
	if len(entries) != 0 {
		t.Fatalf("global list created home entries: %v", entries)
	}
}

func TestRunGlobalListIncludesOnlyExactCanonicalAdoptableFiles(t *testing.T) {
	home := t.TempDir()
	state := t.TempDir()
	t.Setenv("HOME", home)
	t.Setenv("XDG_STATE_HOME", state)

	catalog, err := loadCatalog()
	if err != nil {
		t.Fatalf("loadCatalog: %v", err)
	}
	packs, err := catalog.Resolve([]string{"sdd"})
	if err != nil {
		t.Fatalf("resolve sdd: %v", err)
	}
	if _, err := globalInstallPacksFiltered(packs, true, false); err != nil {
		t.Fatalf("seed legacy skills: %v", err)
	}
	canonical := filepath.Join(home, ".agents", "skills", "bugfix-loop", "SKILL.md")
	modified := filepath.Join(home, ".agents", "skills", "vertical-tdd", "SKILL.md")
	if err := os.WriteFile(modified, []byte("modified legacy\n"), 0o644); err != nil {
		t.Fatalf("modify legacy skill: %v", err)
	}

	var out strings.Builder
	if code := Run([]string{"global", "list", "--include-adoptable", "--json"}, &out, &out); code != 0 {
		t.Fatalf("global list exit code = %d, output: %s", code, out.String())
	}
	var report globalInventoryReport
	if err := json.Unmarshal([]byte(out.String()), &report); err != nil {
		t.Fatalf("decode global list: %v", err)
	}
	states := map[string]string{}
	for _, file := range report.Files {
		states[file.Path] = file.Status
	}
	if states[canonical] != "adoptable" {
		t.Fatalf("canonical legacy status = %q", states[canonical])
	}
	if states[modified] != "unowned" {
		t.Fatalf("modified legacy status = %q, want unowned", states[modified])
	}
	if report.Summary.Adoptable == 0 {
		t.Fatalf("adoptable summary = %+v", report.Summary)
	}
	mustNotExist(t, filepath.Join(state, "geremmyas", "global-manifest.json"))
}

func TestRunGlobalListDoesNotFollowManagedSymlink(t *testing.T) {
	home := t.TempDir()
	state := t.TempDir()
	t.Setenv("HOME", home)
	t.Setenv("XDG_STATE_HOME", state)

	outside := filepath.Join(t.TempDir(), "outside.md")
	if err := os.WriteFile(outside, []byte("outside\n"), 0o644); err != nil {
		t.Fatalf("write outside file: %v", err)
	}
	path := filepath.Join(home, ".codex", "AGENTS.md")
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		t.Fatalf("create codex dir: %v", err)
	}
	if err := os.Symlink(outside, path); err != nil {
		t.Fatalf("create symlink: %v", err)
	}
	manifest := globalManifest{
		Version: globalManifestVersion,
		Packs:   []string{"core"},
		Targets: []string{TargetCodex},
		Files:   map[string]string{path: bytesSHA256([]byte("outside\n"))},
	}
	if err := writeGlobalManifest(manifest); err != nil {
		t.Fatalf("write manifest: %v", err)
	}

	var out strings.Builder
	if code := Run([]string{"global", "list", "--json"}, &out, &out); code != 0 {
		t.Fatalf("global list exit code = %d, output: %s", code, out.String())
	}
	var report globalInventoryReport
	if err := json.Unmarshal([]byte(out.String()), &report); err != nil {
		t.Fatalf("decode global list: %v", err)
	}
	if len(report.Files) != 1 || report.Files[0].Status != "symlink" {
		t.Fatalf("symlink report = %+v", report.Files)
	}
	if report.Files[0].CurrentHash != "" {
		t.Fatalf("symlink target was hashed: %+v", report.Files[0])
	}
}

func TestRunGlobalListFiltersTargetSpecificPathsAndKeepsSharedSkills(t *testing.T) {
	home := t.TempDir()
	state := t.TempDir()
	t.Setenv("HOME", home)
	t.Setenv("XDG_STATE_HOME", state)

	var installOut strings.Builder
	if code := Run([]string{"global", "--targets", "codex,cursor", "sdd"}, &installOut, &installOut); code != 0 {
		t.Fatalf("install global exit code = %d, output: %s", code, installOut.String())
	}
	var out strings.Builder
	if code := Run([]string{"global", "list", "--targets", "codex", "--json"}, &out, &out); code != 0 {
		t.Fatalf("global list exit code = %d, output: %s", code, out.String())
	}
	var report globalInventoryReport
	if err := json.Unmarshal([]byte(out.String()), &report); err != nil {
		t.Fatalf("decode global list: %v", err)
	}
	foundShared := false
	for _, file := range report.Files {
		if strings.Contains(filepath.ToSlash(file.Path), "/.cursor/") {
			t.Fatalf("codex filter included cursor path: %s", file.Path)
		}
		if file.Shared {
			foundShared = true
		}
	}
	if !foundShared {
		t.Fatalf("codex filter omitted shared skills")
	}

	out.Reset()
	if code := Run([]string{"global", "list", "--targets", "not-a-target", "--json"}, &out, &out); code == 0 {
		t.Fatalf("global list accepted unsupported target: %s", out.String())
	}
}

func TestRunGlobalListHumanOutputIsStableAndScoped(t *testing.T) {
	home := t.TempDir()
	state := t.TempDir()
	t.Setenv("HOME", home)
	t.Setenv("XDG_STATE_HOME", state)
	obsolete := filepath.Join(home, ".codex", "instructions", "legacy.md")
	if err := os.MkdirAll(filepath.Dir(obsolete), 0o755); err != nil {
		t.Fatalf("mkdir obsolete parent: %v", err)
	}
	content := []byte("legacy\n")
	if err := os.WriteFile(obsolete, content, 0o644); err != nil {
		t.Fatalf("write obsolete file: %v", err)
	}
	hash := bytesSHA256(content)
	if err := writeGlobalManifest(globalManifest{
		Version: globalManifestVersion,
		Packs:   []string{"core"},
		Targets: []string{TargetCodex},
		Files:   map[string]string{obsolete: hash},
	}); err != nil {
		t.Fatalf("write manifest: %v", err)
	}

	var out strings.Builder
	if code := Run([]string{"global", "list"}, &out, &out); code != 0 {
		t.Fatalf("global list exit code = %d, output: %s", code, out.String())
	}
	for _, fragment := range []string{
		"global manifest:", "packs:", "targets:", "STATUS", "OBSOLETE", "INSTALLED HASH",
		"CURRENT HASH", "summary:", "obsolete", "true", hash, obsolete,
	} {
		if !strings.Contains(out.String(), fragment) {
			t.Fatalf("human output missing %q:\n%s", fragment, out.String())
		}
	}
}

func TestRunGlobalListReportsUnknownUnownedFiles(t *testing.T) {
	home := t.TempDir()
	state := t.TempDir()
	t.Setenv("HOME", home)
	t.Setenv("XDG_STATE_HOME", state)
	unknown := filepath.Join(home, ".agents", "skills", "external", "SKILL.md")
	if err := os.MkdirAll(filepath.Dir(unknown), 0o755); err != nil {
		t.Fatalf("create unknown parent: %v", err)
	}
	if err := os.WriteFile(unknown, []byte("external\n"), 0o644); err != nil {
		t.Fatalf("write unknown file: %v", err)
	}

	var out strings.Builder
	if code := Run([]string{"global", "list", "--json"}, &out, &out); code != 0 {
		t.Fatalf("global list exit code = %d, output: %s", code, out.String())
	}
	var report globalInventoryReport
	if err := json.Unmarshal([]byte(out.String()), &report); err != nil {
		t.Fatalf("decode global list: %v", err)
	}
	if len(report.Files) != 1 || report.Files[0].Path != unknown || report.Files[0].Ownership != "unowned" || report.Files[0].Status != "unowned" {
		t.Fatalf("unowned report = %+v", report.Files)
	}
	if report.Summary.Unowned != 1 {
		t.Fatalf("unowned summary = %+v", report.Summary)
	}
}

func TestRunGlobalListRejectsCanonicalPathCollision(t *testing.T) {
	home := t.TempDir()
	state := t.TempDir()
	t.Setenv("HOME", home)
	t.Setenv("XDG_STATE_HOME", state)
	path := filepath.Join(home, ".codex", "AGENTS.md")
	manifest := globalManifest{
		Version: globalManifestVersion,
		Packs:   []string{"core"},
		Targets: []string{TargetCodex},
		Files: map[string]string{
			path: bytesSHA256(nil),
			filepath.Join(home, ".codex", ".", "AGENTS.md"): bytesSHA256([]byte("different")),
		},
	}
	// filepath.Join cleans paths, so preserve a distinct textual key explicitly.
	manifest.Files[filepath.Dir(path)+string(filepath.Separator)+"."+string(filepath.Separator)+filepath.Base(path)] = bytesSHA256([]byte("different"))
	if err := writeGlobalManifest(manifest); err != nil {
		t.Fatalf("write manifest: %v", err)
	}

	var out strings.Builder
	if code := Run([]string{"global", "list", "--json"}, &out, &out); code == 0 {
		t.Fatalf("global list accepted colliding paths: %s", out.String())
	}
	if !strings.Contains(out.String(), "collide after normalization") {
		t.Fatalf("collision error missing: %s", out.String())
	}
}

func TestRunGlobalListNormalizesOwnedPathBeforeAdoptableComparison(t *testing.T) {
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
		t.Fatalf("resolve sdd: %v", err)
	}
	if _, err := globalInstallPacksFiltered(packs, true, false); err != nil {
		t.Fatalf("materialize canonical skill: %v", err)
	}
	path := filepath.Join(home, ".agents", "skills", "bugfix-loop", "SKILL.md")
	hash, err := fileSHA256(path)
	if err != nil {
		t.Fatalf("hash skill: %v", err)
	}
	rawPath := filepath.Dir(path) + string(filepath.Separator) + "." + string(filepath.Separator) + filepath.Base(path)
	if err := writeGlobalManifest(globalManifest{
		Version: globalManifestVersion,
		Packs:   []string{"sdd"},
		Targets: []string{TargetCodex},
		Files:   map[string]string{rawPath: hash},
	}); err != nil {
		t.Fatalf("write manifest: %v", err)
	}

	var out strings.Builder
	if code := Run([]string{"global", "list", "--include-adoptable", "--json"}, &out, &out); code != 0 {
		t.Fatalf("global list exit code = %d, output: %s", code, out.String())
	}
	var report globalInventoryReport
	if err := json.Unmarshal([]byte(out.String()), &report); err != nil {
		t.Fatalf("decode report: %v", err)
	}
	count := 0
	for _, entry := range report.Files {
		if entry.Path == path {
			count++
			if entry.Ownership != "managed" {
				t.Fatalf("normalized path ownership = %q", entry.Ownership)
			}
		}
	}
	if count != 1 {
		t.Fatalf("normalized path count = %d, want 1", count)
	}
}

func TestRunGlobalListTreatsTrustedGeneratedMarkerAsAdoptable(t *testing.T) {
	home := t.TempDir()
	t.Setenv("HOME", home)
	path := filepath.Join(home, ".codex", "AGENTS.md")
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		t.Fatalf("mkdir codex: %v", err)
	}
	if err := os.WriteFile(path, []byte(generatedMarker+"\nlegacy generated\n"), 0o644); err != nil {
		t.Fatalf("write generated file: %v", err)
	}

	var out strings.Builder
	if code := Run([]string{"global", "list", "--include-adoptable", "--json"}, &out, &out); code != 0 {
		t.Fatalf("global list exit code = %d, output: %s", code, out.String())
	}
	var report globalInventoryReport
	if err := json.Unmarshal([]byte(out.String()), &report); err != nil {
		t.Fatalf("decode report: %v", err)
	}
	for _, entry := range report.Files {
		if entry.Path == path && entry.Status == "adoptable" && entry.Proof == "generated-marker" {
			return
		}
	}
	t.Fatalf("trusted generated file not adoptable: %+v", report.Files)
}

func TestRunGlobalListTreatsUnknownHistoricalPackAsUnresolvedNotObsolete(t *testing.T) {
	home := t.TempDir()
	state := t.TempDir()
	t.Setenv("HOME", home)
	t.Setenv("XDG_STATE_HOME", state)
	path := filepath.Join(home, ".codex", "AGENTS.md")
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		t.Fatalf("create codex dir: %v", err)
	}
	if err := os.WriteFile(path, []byte("legacy\n"), 0o644); err != nil {
		t.Fatalf("write legacy file: %v", err)
	}
	hash, _ := fileSHA256(path)
	if err := writeGlobalManifest(globalManifest{
		Version: globalManifestVersion,
		Packs:   []string{"removed-pack"},
		Targets: []string{TargetCodex},
		Files:   map[string]string{path: hash},
	}); err != nil {
		t.Fatalf("write manifest: %v", err)
	}

	var out strings.Builder
	if code := Run([]string{"global", "list", "--json"}, &out, &out); code != 0 {
		t.Fatalf("global list exit code = %d, output: %s", code, out.String())
	}
	var report globalInventoryReport
	if err := json.Unmarshal([]byte(out.String()), &report); err != nil {
		t.Fatalf("decode global list: %v", err)
	}
	if report.CatalogResolved || report.CatalogError == "" {
		t.Fatalf("catalog resolution = resolved %v error %q", report.CatalogResolved, report.CatalogError)
	}
	if len(report.Files) != 1 || report.Files[0].Status != "current" || report.Files[0].Obsolete {
		t.Fatalf("historical pack file = %+v", report.Files)
	}
}

func TestRunGlobalListRejectsUnknownManifestTarget(t *testing.T) {
	home := t.TempDir()
	state := t.TempDir()
	t.Setenv("HOME", home)
	t.Setenv("XDG_STATE_HOME", state)
	if err := writeGlobalManifest(globalManifest{
		Version: globalManifestVersion,
		Packs:   []string{"core"},
		Targets: []string{"future-target"},
		Files:   map[string]string{},
	}); err != nil {
		t.Fatalf("write manifest: %v", err)
	}
	var out strings.Builder
	if code := Run([]string{"global", "list", "--json"}, &out, &out); code == 0 {
		t.Fatalf("global list accepted incompatible target: %s", out.String())
	}
}

func TestRunGlobalListDetectsHomeRootSymlink(t *testing.T) {
	realHome := t.TempDir()
	parent := t.TempDir()
	linkedHome := filepath.Join(parent, "linked-home")
	if err := os.Symlink(realHome, linkedHome); err != nil {
		t.Fatalf("link home: %v", err)
	}
	state := t.TempDir()
	t.Setenv("HOME", linkedHome)
	t.Setenv("XDG_STATE_HOME", state)
	path := filepath.Join(linkedHome, ".codex", "AGENTS.md")
	if err := os.MkdirAll(filepath.Join(realHome, ".codex"), 0o755); err != nil {
		t.Fatalf("create real codex dir: %v", err)
	}
	if err := os.WriteFile(filepath.Join(realHome, ".codex", "AGENTS.md"), []byte("linked\n"), 0o644); err != nil {
		t.Fatalf("write linked file: %v", err)
	}
	if err := writeGlobalManifest(globalManifest{
		Version: globalManifestVersion,
		Packs:   []string{"core"},
		Targets: []string{TargetCodex},
		Files:   map[string]string{path: bytesSHA256([]byte("linked\n"))},
	}); err != nil {
		t.Fatalf("write manifest: %v", err)
	}

	var out strings.Builder
	if code := Run([]string{"global", "list", "--json"}, &out, &out); code != 0 {
		t.Fatalf("global list exit code = %d, output: %s", code, out.String())
	}
	var report globalInventoryReport
	if err := json.Unmarshal([]byte(out.String()), &report); err != nil {
		t.Fatalf("decode global list: %v", err)
	}
	if len(report.Files) != 1 || report.Files[0].Status != "symlink" || report.Files[0].CurrentHash != "" {
		t.Fatalf("home symlink report = %+v", report.Files)
	}
}

func TestRunGlobalListTargetFilterAppliesToObservedRootsAndSharedAdoptables(t *testing.T) {
	home := t.TempDir()
	state := t.TempDir()
	t.Setenv("HOME", home)
	t.Setenv("XDG_STATE_HOME", state)
	if err := os.MkdirAll(filepath.Join(home, ".codex", "plugins", "cache"), 0o755); err != nil {
		t.Fatalf("create plugin cache: %v", err)
	}
	catalog, err := loadCatalog()
	if err != nil {
		t.Fatalf("loadCatalog: %v", err)
	}
	packs, err := catalog.Resolve([]string{"sdd"})
	if err != nil {
		t.Fatalf("resolve sdd: %v", err)
	}
	if _, err := globalInstallPacksFiltered(packs, true, false); err != nil {
		t.Fatalf("seed legacy skills: %v", err)
	}

	var out strings.Builder
	if code := Run([]string{"global", "list", "--targets", "cursor", "--include-adoptable", "--json"}, &out, &out); code != 0 {
		t.Fatalf("cursor inventory exit code = %d, output: %s", code, out.String())
	}
	var cursorReport globalInventoryReport
	if err := json.Unmarshal([]byte(out.String()), &cursorReport); err != nil {
		t.Fatalf("decode cursor report: %v", err)
	}
	if len(cursorReport.Observed) != 0 || cursorReport.Summary.Adoptable == 0 {
		t.Fatalf("cursor report observed=%+v summary=%+v", cursorReport.Observed, cursorReport.Summary)
	}

	out.Reset()
	if code := Run([]string{"global", "list", "--targets", "codex", "--json"}, &out, &out); code != 0 {
		t.Fatalf("codex inventory exit code = %d, output: %s", code, out.String())
	}
	var codexReport globalInventoryReport
	if err := json.Unmarshal([]byte(out.String()), &codexReport); err != nil {
		t.Fatalf("decode codex report: %v", err)
	}
	if len(codexReport.Observed) != 1 || codexReport.Observed[0].Kind != "plugin-cache" {
		t.Fatalf("codex observed roots = %+v", codexReport.Observed)
	}
}
