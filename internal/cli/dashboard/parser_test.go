package dashboard

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func TestParseTaskStats(t *testing.T) {
	content := "- [x] done\n- [~] wip\n- [ ] todo\n  - [x] nested\n"
	stats := ParseTaskStats(content)
	want := TaskStats{Total: 4, Done: 2, InProgress: 1, Pending: 1}
	if stats != want {
		t.Fatalf("stats = %+v, want %+v", stats, want)
	}
}

func TestParseDateFromFilename(t *testing.T) {
	date, slug, ok := ParseDateFromFilename("2026-01-15-onboarding.md")
	if !ok || date != "2026-01-15" || slug != "onboarding" {
		t.Fatalf("got %q %q %v", date, slug, ok)
	}
	if _, _, ok := ParseDateFromFilename("bad.md"); ok {
		t.Fatal("expected false for bad.md")
	}
}

func TestScanSpecsIntegration(t *testing.T) {
	root := t.TempDir()
	writeSpec(t, root, "0001-foo", `---
spec: "0001"
title: Alpha
family: onboarding
phase: 0
status: Implemented
depends_on: [1, 3]
origin: "docs/prds/2026-01-15-x.md"
---
Body`, `- [x] a\n- [ ] b\n`)
	writeSpec(t, root, "0002-bar", "no frontmatter\n", "")
	writeSpec(t, root, "0003-baz", `---
title: No plan
family: onboarding
phase: 1
status: Draft
---
`, "")

	data, err := ScanSpecs(root)
	if err != nil {
		t.Fatal(err)
	}
	if len(data.Warnings) == 0 {
		t.Fatal("expected warning for missing frontmatter")
	}
	var foundAlpha bool
	for _, fam := range data.Families {
		for _, ph := range fam.Phases {
			for _, s := range ph.Specs {
				if s.Number == 1 {
					foundAlpha = true
					if s.Tasks.Total != 2 || s.Tasks.Done != 1 {
						t.Fatalf("tasks = %+v", s.Tasks)
					}
					if len(s.DependsOn) != 2 || s.DependsOn[0] != 1 {
						t.Fatalf("depends = %v", s.DependsOn)
					}
				}
				if s.Number == 3 && s.HasPlan {
					t.Fatal("0003 should not have plan")
				}
			}
		}
	}
	if !foundAlpha {
		t.Fatal("spec 0001 not found")
	}
}

func TestScanSpecsDocsSpecs(t *testing.T) {
	root := t.TempDir()
	writeSpecAt(t, root, "docs/specs", "0010-docs-root", `---
title: From docs/specs
family: platform
phase: 1
status: Draft
---
`, "")
	data, err := ScanSpecs(root)
	if err != nil {
		t.Fatal(err)
	}
	var found bool
	for _, fam := range data.Families {
		for _, ph := range fam.Phases {
			for _, s := range ph.Specs {
				if s.Number == 10 {
					found = true
					if s.Title != "From docs/specs" {
						t.Fatalf("title = %q", s.Title)
					}
				}
			}
		}
	}
	if !found {
		t.Fatal("spec 0010 under docs/specs not found")
	}
}

func TestScanSpecsDuplicateAcrossRoots(t *testing.T) {
	root := t.TempDir()
	writeSpecAt(t, root, "specs", "0005-dup", "---\ntitle: Root\nfamily: f\nphase: 0\nstatus: Draft\n---\n", "")
	writeSpecAt(t, root, "docs/specs", "0005-dup", "---\ntitle: Docs\nfamily: f\nphase: 0\nstatus: Draft\n---\n", "")
	data, err := ScanSpecs(root)
	if err != nil {
		t.Fatal(err)
	}
	count := 0
	for _, fam := range data.Families {
		for _, ph := range fam.Phases {
			count += len(ph.Specs)
		}
	}
	if count != 1 {
		t.Fatalf("want 1 spec, got %d", count)
	}
	dupWarn := false
	for _, w := range data.Warnings {
		if strings.Contains(w.Message, "duplicate spec number 0005") {
			dupWarn = true
		}
	}
	if !dupWarn {
		t.Fatal("expected duplicate warning for docs/specs copy")
	}
}

func TestScanPostmortems(t *testing.T) {
	root := t.TempDir()
	dir := filepath.Join(root, "docs", "postmortems")
	if err := os.MkdirAll(dir, 0o755); err != nil {
		t.Fatal(err)
	}
	path := filepath.Join(dir, "2026-04-01-outage.md")
	if err := os.WriteFile(path, []byte("# Outage\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	pms, _, err := ScanPostmortems(root)
	if err != nil || len(pms) != 1 || pms[0].Date != "2026-04-01" {
		t.Fatalf("postmortems = %+v err=%v", pms, err)
	}
}

func TestWriteSpecReadmesBothRoots(t *testing.T) {
	root := t.TempDir()
	for _, rel := range []string{"specs", "docs/specs"} {
		if err := os.MkdirAll(filepath.Join(root, rel), 0o755); err != nil {
			t.Fatal(err)
		}
	}
	data := DashboardData{Families: []Family{{Name: "F", Phases: []Phase{{Specs: []SpecSummary{{
		Number: 1, Title: "T", Status: "Draft",
	}}}}}}}
	written, err := WriteSpecReadmes(root, data)
	if err != nil {
		t.Fatal(err)
	}
	if len(written) != 2 {
		t.Fatalf("written = %v", written)
	}
}

func TestScanPRDs(t *testing.T) {
	root := t.TempDir()
	dir := filepath.Join(root, "docs", "prds")
	if err := os.MkdirAll(dir, 0o755); err != nil {
		t.Fatal(err)
	}
	path := filepath.Join(dir, "2026-03-20-billing.md")
	if err := os.WriteFile(path, []byte("# Billing\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	prds, _, err := ScanPRDs(root)
	if err != nil || len(prds) != 1 || prds[0].Date != "2026-03-20" {
		t.Fatalf("prds = %+v err=%v", prds, err)
	}
}

func TestScanSpecsPerformance(t *testing.T) {
	root := t.TempDir()
	for i := 1; i <= 100; i++ {
		name := filepath.Join(root, "specs", formatSpecDir(i))
		if err := os.MkdirAll(name, 0o755); err != nil {
			t.Fatal(err)
		}
		content := "---\ntitle: T\nfamily: f\nphase: 0\nstatus: Draft\n---\n"
		if err := os.WriteFile(filepath.Join(name, "spec.md"), []byte(content), 0o644); err != nil {
			t.Fatal(err)
		}
	}
	start := time.Now()
	if _, err := ScanSpecs(root); err != nil {
		t.Fatal(err)
	}
	if time.Since(start) > time.Second {
		t.Fatalf("scan took %v, want < 1s", time.Since(start))
	}
}

func writeSpecAt(t *testing.T, root, specRoot, folder, specBody, tasks string) {
	t.Helper()
	dir := filepath.Join(root, specRoot, folder)
	if err := os.MkdirAll(dir, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dir, "spec.md"), []byte(specBody), 0o644); err != nil {
		t.Fatal(err)
	}
	if tasks != "" {
		if err := os.WriteFile(filepath.Join(dir, "tasks.md"), []byte(strings.ReplaceAll(tasks, `\n`, "\n")), 0o644); err != nil {
			t.Fatal(err)
		}
	}
}

func writeSpec(t *testing.T, root, folder, specBody, tasks string) {
	writeSpecAt(t, root, "specs", folder, specBody, tasks)
}

func formatSpecDir(n int) string {
	return fmt.Sprintf("%04d-s", n)
}
