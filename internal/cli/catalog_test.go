package cli

import (
	"io/fs"
	"sort"
	"strings"
	"testing"

	geremmyas "github.com/woliveiras/geremmyas"
)

func TestLoadCatalogAndResolveDependencies(t *testing.T) {
	catalog, err := loadCatalog()
	if err != nil {
		t.Fatalf("loadCatalog returned error: %v", err)
	}

	packs, err := catalog.Resolve([]string{"sdd"})
	if err != nil {
		t.Fatalf("Resolve returned error: %v", err)
	}

	got := make([]string, 0, len(packs))
	for _, pack := range packs {
		got = append(got, pack.Name)
	}
	want := []string{"core", "sdd"}
	if len(got) != len(want) {
		t.Fatalf("resolved packs = %v, want %v", got, want)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("resolved packs = %v, want %v", got, want)
		}
	}
}

func TestCatalogSourcesExist(t *testing.T) {
	catalog, err := loadCatalog()
	if err != nil {
		t.Fatalf("loadCatalog returned error: %v", err)
	}
	if err := catalog.ValidateSources(); err != nil {
		t.Fatalf("ValidateSources returned error: %v", err)
	}
}

func TestCatalogCoversEveryTopLevelSkill(t *testing.T) {
	catalog, err := loadCatalog()
	if err != nil {
		t.Fatalf("loadCatalog returned error: %v", err)
	}

	sourceSkills, err := topLevelSkillNames()
	if err != nil {
		t.Fatalf("topLevelSkillNames returned error: %v", err)
	}
	catalogSkills := catalogSkillNames(catalog)

	var missing []string
	for skill := range sourceSkills {
		if !catalogSkills[skill] {
			missing = append(missing, skill)
		}
	}
	sort.Strings(missing)
	if len(missing) > 0 {
		t.Fatalf("top-level skills missing from catalog packs: %s", strings.Join(missing, ", "))
	}
}

func TestResolveBragMePack(t *testing.T) {
	catalog, err := loadCatalog()
	if err != nil {
		t.Fatalf("loadCatalog returned error: %v", err)
	}

	packs, err := catalog.Resolve([]string{"brag-me"})
	if err != nil {
		t.Fatalf("Resolve returned error: %v", err)
	}

	if len(packs) != 1 {
		t.Fatalf("resolved packs = %d, want 1", len(packs))
	}
	if packs[0].Name != "brag-me" {
		t.Fatalf("resolved pack = %q, want brag-me", packs[0].Name)
	}
}

func TestResolveRejectsUnknownPack(t *testing.T) {
	catalog, err := loadCatalog()
	if err != nil {
		t.Fatalf("loadCatalog returned error: %v", err)
	}
	if _, err := catalog.Resolve([]string{"missing-pack"}); err == nil {
		t.Fatal("Resolve succeeded, want error")
	}
}

func topLevelSkillNames() (map[string]bool, error) {
	const root = "project/.github/skills"
	const prefix = root + "/"
	skills := map[string]bool{}

	err := fs.WalkDir(geremmyas.EmbeddedFiles, root, func(path string, d fs.DirEntry, err error) error {
		if err != nil || d.IsDir() {
			return err
		}
		rel := strings.TrimPrefix(path, prefix)
		parts := strings.Split(rel, "/")
		if len(parts) == 2 && parts[1] == "SKILL.md" {
			skills[parts[0]] = true
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return skills, nil
}

func catalogSkillNames(catalog Catalog) map[string]bool {
	const prefix = "project/.github/skills/"
	skills := map[string]bool{}

	for _, pack := range catalog.Packs {
		for _, entry := range pack.Files {
			if !strings.HasPrefix(entry.Source, prefix) {
				continue
			}
			rel := strings.TrimPrefix(entry.Source, prefix)
			name := strings.Split(rel, "/")[0]
			if name != "" {
				skills[name] = true
			}
		}
	}
	return skills
}
