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

func TestDoctorRejectsMissingCatalogSource(t *testing.T) {
	badCatalog := Catalog{
		Packs: []Pack{{
			Name: "broken",
			Files: []FileEntry{{
				Source: "project/.github/skills/missing",
				Target: ".github/skills/missing",
			}},
		}},
	}

	var out strings.Builder
	if err := runDoctor(&out, badCatalog); err == nil {
		t.Fatal("runDoctor succeeded, want missing source error")
	}
}

func TestCatalogTiersValid(t *testing.T) {
	catalog, err := loadCatalog()
	if err != nil {
		t.Fatalf("loadCatalog returned error: %v", err)
	}
	if err := catalog.ValidateTiers(); err != nil {
		t.Fatalf("ValidateTiers returned error: %v", err)
	}

	tiers := map[string]string{}
	for _, pack := range catalog.Packs {
		tiers[pack.Name] = pack.Tier
	}
	for _, name := range []string{"core", "sdd"} {
		if tiers[name] != TierCore {
			t.Fatalf("pack %q tier = %q, want %q", name, tiers[name], TierCore)
		}
	}
	for name, tier := range tiers {
		if name == "core" || name == "sdd" {
			continue
		}
		if tier != TierStack {
			t.Fatalf("pack %q tier = %q, want %q", name, tier, TierStack)
		}
	}
}

func TestValidateTiersRejectsMissingTier(t *testing.T) {
	badCatalog := Catalog{Packs: []Pack{{Name: "broken"}}}
	err := badCatalog.ValidateTiers()
	if err == nil {
		t.Fatal("ValidateTiers succeeded, want missing tier error")
	}
	if !strings.Contains(err.Error(), "broken") {
		t.Fatalf("error %q does not name the offending pack", err)
	}
}

func TestValidateTiersRejectsInvalidTier(t *testing.T) {
	badCatalog := Catalog{Packs: []Pack{{Name: "broken", Tier: "personal"}}}
	err := badCatalog.ValidateTiers()
	if err == nil {
		t.Fatal("ValidateTiers succeeded, want invalid tier error")
	}
	if !strings.Contains(err.Error(), "personal") {
		t.Fatalf("error %q does not name the invalid tier", err)
	}
}

func TestResearchPackIncludesPaperReview(t *testing.T) {
	catalog, err := loadCatalog()
	if err != nil {
		t.Fatalf("loadCatalog returned error: %v", err)
	}
	var research *Pack
	for i := range catalog.Packs {
		if catalog.Packs[i].Name == "research" {
			research = &catalog.Packs[i]
			break
		}
	}
	if research == nil {
		t.Fatal("research pack not found in catalog")
	}
	found := false
	for _, file := range research.Files {
		if file.Source == "project/.github/skills/paper-review" {
			found = true
			break
		}
	}
	if !found {
		t.Fatal("research pack does not include paper-review skill")
	}
}

func TestDoctorWithoutConfigReportsInitHint(t *testing.T) {
	withTempCwd(t)
	catalog, err := loadCatalog()
	if err != nil {
		t.Fatalf("loadCatalog returned error: %v", err)
	}

	var out strings.Builder
	if err := runDoctor(&out, catalog); err != nil {
		t.Fatalf("runDoctor returned error: %v", err)
	}
	if !strings.Contains(out.String(), "geremmyas.yml: missing; run geremmyas init") {
		t.Fatalf("doctor output missing init hint:\n%s", out.String())
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

func TestCatalogCoversEveryInstruction(t *testing.T) {
	catalog, err := loadCatalog()
	if err != nil {
		t.Fatalf("loadCatalog returned error: %v", err)
	}

	sourceInstructions, err := embeddedBasenames("project/.github/instructions", ".instructions.md")
	if err != nil {
		t.Fatalf("embeddedBasenames returned error: %v", err)
	}
	catalogInstructions := catalogSourceBasenames(catalog, "project/.github/instructions/", ".instructions.md")

	var missing []string
	for instruction := range sourceInstructions {
		if !catalogInstructions[instruction] {
			missing = append(missing, instruction)
		}
	}
	sort.Strings(missing)
	if len(missing) > 0 {
		t.Fatalf("instructions missing from catalog packs: %s", strings.Join(missing, ", "))
	}
}

func TestCatalogSDDCoversEveryAgent(t *testing.T) {
	catalog, err := loadCatalog()
	if err != nil {
		t.Fatalf("loadCatalog returned error: %v", err)
	}

	sourceAgents, err := embeddedBasenames("project/.github/agents", ".agent.md")
	if err != nil {
		t.Fatalf("embeddedBasenames returned error: %v", err)
	}
	sdd, ok := catalog.Pack("sdd")
	if !ok {
		t.Fatal("catalog missing sdd pack")
	}
	catalogAgents := packSourceBasenames(sdd, "project/.github/agents/", ".agent.md")

	var missing []string
	for agent := range sourceAgents {
		if !catalogAgents[agent] {
			missing = append(missing, agent)
		}
	}
	sort.Strings(missing)
	if len(missing) > 0 {
		t.Fatalf("agents missing from sdd pack: %s", strings.Join(missing, ", "))
	}
}

func TestCatalogDoesNotReferenceNestedSkillMarkdownAsTopLevelSkill(t *testing.T) {
	catalog, err := loadCatalog()
	if err != nil {
		t.Fatalf("loadCatalog returned error: %v", err)
	}

	var nested []string
	for _, pack := range catalog.Packs {
		for _, entry := range pack.Files {
			if strings.HasPrefix(entry.Source, "project/.github/skills/") &&
				strings.HasSuffix(entry.Source, "/SKILL.md") &&
				strings.Count(strings.TrimPrefix(entry.Source, "project/.github/skills/"), "/") > 1 {
				nested = append(nested, pack.Name+":"+entry.Source)
			}
		}
	}
	sort.Strings(nested)
	if len(nested) > 0 {
		t.Fatalf("catalog references nested SKILL.md files directly: %s", strings.Join(nested, ", "))
	}
}

func TestCatalogDependenciesResolveForEveryPack(t *testing.T) {
	catalog, err := loadCatalog()
	if err != nil {
		t.Fatalf("loadCatalog returned error: %v", err)
	}

	for _, pack := range catalog.Packs {
		t.Run(pack.Name, func(t *testing.T) {
			if _, err := catalog.Resolve([]string{pack.Name}); err != nil {
				t.Fatalf("Resolve(%q) returned error: %v", pack.Name, err)
			}
		})
	}
}

func TestCatalogCompositePackDependencyClosure(t *testing.T) {
	catalog, err := loadCatalog()
	if err != nil {
		t.Fatalf("loadCatalog returned error: %v", err)
	}

	tests := map[string][]string{
		"go-ci":         {"go-base", "infra-ci", "go-ci"},
		"python-ci":     {"python-base", "infra-ci", "python-ci"},
		"react-data":    {"typescript-base", "react-web", "react-data"},
		"typescript-ci": {"typescript-base", "typescript-ci"},
	}
	for packName, want := range tests {
		t.Run(packName, func(t *testing.T) {
			packs, err := catalog.Resolve([]string{packName})
			if err != nil {
				t.Fatalf("Resolve(%q) returned error: %v", packName, err)
			}
			got := packNames(packs)
			if strings.Join(got, ",") != strings.Join(want, ",") {
				t.Fatalf("Resolve(%q) = %v, want %v", packName, got, want)
			}
		})
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

func embeddedBasenames(root, suffix string) (map[string]bool, error) {
	items := map[string]bool{}
	err := fs.WalkDir(geremmyas.EmbeddedFiles, root, func(path string, d fs.DirEntry, err error) error {
		if err != nil || d.IsDir() {
			return err
		}
		if strings.HasSuffix(path, suffix) {
			items[strings.TrimSuffix(path[strings.LastIndex(path, "/")+1:], suffix)] = true
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return items, nil
}

func catalogSourceBasenames(catalog Catalog, prefix, suffix string) map[string]bool {
	items := map[string]bool{}
	for _, pack := range catalog.Packs {
		for name := range packSourceBasenames(pack, prefix, suffix) {
			items[name] = true
		}
	}
	return items
}

func packSourceBasenames(pack Pack, prefix, suffix string) map[string]bool {
	items := map[string]bool{}
	for _, entry := range pack.Files {
		if entry.Source == strings.TrimSuffix(prefix, "/") || strings.HasPrefix(entry.Source, prefix) {
			if strings.HasSuffix(entry.Source, suffix) {
				items[strings.TrimSuffix(entry.Source[strings.LastIndex(entry.Source, "/")+1:], suffix)] = true
				continue
			}
			_ = fs.WalkDir(geremmyas.EmbeddedFiles, entry.Source, func(path string, d fs.DirEntry, err error) error {
				if err != nil || d.IsDir() {
					return err
				}
				if strings.HasSuffix(path, suffix) {
					items[strings.TrimSuffix(path[strings.LastIndex(path, "/")+1:], suffix)] = true
				}
				return nil
			})
		}
	}
	return items
}

func packNames(packs []Pack) []string {
	names := make([]string, 0, len(packs))
	for _, pack := range packs {
		names = append(names, pack.Name)
	}
	return names
}
