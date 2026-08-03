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

	packs, err := catalog.Resolve([]string{"base"})
	if err != nil {
		t.Fatalf("Resolve returned error: %v", err)
	}

	got := make([]string, 0, len(packs))
	for _, pack := range packs {
		got = append(got, pack.Name)
	}
	want := []string{"core", "coding", "quality", "base"}
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
				Source: "content/skills/missing",
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
	for _, name := range []string{"core", "coding", "quality", "base"} {
		if tiers[name] != TierCore {
			t.Fatalf("pack %q tier = %q, want %q", name, tiers[name], TierCore)
		}
	}
	for name, tier := range tiers {
		if name == "core" || name == "coding" || name == "quality" || name == "base" {
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
		if file.Source == "content/skills/paper-review" {
			found = true
			break
		}
	}
	if !found {
		t.Fatal("research pack does not include paper-review skill")
	}
}

func TestGameDevFocusedPacksAndCompleteMetapack(t *testing.T) {
	catalog, err := loadCatalog()
	if err != nil {
		t.Fatalf("loadCatalog returned error: %v", err)
	}

	wantByPack := map[string][]string{
		"game-core":        {"game-testing-2d", "gameplay-programming-2d"},
		"game-ui":          {"game-feel-2d", "game-ui-accessibility"},
		"game-systems":     {"game-ai-2d", "game-save-n-progress", "procedural-generation-2d"},
		"game-performance": {"game-performance-2d"},
		"game-audio":       {"game-audio-2d"},
		"game-art":         {"game-art-2d"},
		"game-delivery":    {"game-build-and-release"},
	}
	for packName, want := range wantByPack {
		t.Run(packName, func(t *testing.T) {
			pack, ok := catalog.Pack(packName)
			if !ok {
				t.Fatalf("catalog missing %s pack", packName)
			}
			got := packSourceBasenames(pack, "content/skills/", "")
			if strings.Join(sortedKeys(got), ",") != strings.Join(want, ",") {
				t.Fatalf("%s discoverable skills = %v, want %v", packName, sortedKeys(got), want)
			}
		})
	}

	gameDev, ok := catalog.Pack("game-dev")
	if !ok {
		t.Fatal("catalog missing game-dev pack")
	}
	if len(gameDev.Files) != 0 {
		t.Fatalf("game-dev directly owns %d files, want dependency-only metapack", len(gameDev.Files))
	}
	resolved, err := catalog.Resolve([]string{"game-dev"})
	if err != nil {
		t.Fatalf("Resolve(game-dev) returned error: %v", err)
	}
	got := map[string]bool{}
	for _, pack := range resolved {
		for name := range packSourceBasenames(pack, "content/skills/", "") {
			if got[name] {
				t.Errorf("game-dev resolves duplicate skill %q", name)
			}
			got[name] = true
		}
	}
	want := []string{"game-ai-2d", "game-art-2d", "game-audio-2d", "game-build-and-release", "game-feel-2d", "game-performance-2d", "game-save-n-progress", "game-testing-2d", "game-ui-accessibility", "gameplay-programming-2d", "procedural-generation-2d"}
	if strings.Join(sortedKeys(got), ",") != strings.Join(want, ",") {
		t.Fatalf("game-dev discoverable skills = %v, want %v", sortedKeys(got), want)
	}

	legacy, err := catalog.Resolve([]string{"game-art-2d"})
	if err != nil {
		t.Fatalf("Resolve(game-art-2d) returned error: %v", err)
	}
	if got := packNames(legacy); strings.Join(got, ",") != "game-art,game-art-2d" {
		t.Fatalf("Resolve(game-art-2d) = %v, want compatibility alias over game-art", got)
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

	sourceInstructions, err := embeddedBasenames("content/instructions", ".instructions.md")
	if err != nil {
		t.Fatalf("embeddedBasenames returned error: %v", err)
	}
	catalogInstructions := catalogSourceBasenames(catalog, "content/instructions/", ".instructions.md")

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

func TestCatalogDistributesNoBundledAgents(t *testing.T) {
	catalog, err := loadCatalog()
	if err != nil {
		t.Fatalf("loadCatalog returned error: %v", err)
	}
	for _, pack := range catalog.Packs {
		for _, entry := range pack.Files {
			if entry.Kind == ArtifactAgent || strings.HasPrefix(entry.Source, "content/agents") {
				t.Errorf("pack %q still distributes custom agent artifact %+v", pack.Name, entry)
			}
		}
	}
}

func TestGlobalSubcommandNamesAreNotCatalogPacks(t *testing.T) {
	catalog, err := loadCatalog()
	if err != nil {
		t.Fatalf("loadCatalog returned error: %v", err)
	}
	for _, reserved := range []string{"list", "clear"} {
		if _, exists := catalog.Pack(reserved); exists {
			t.Errorf("reserved global subcommand %q is also a pack", reserved)
		}
	}
}

func TestVerifyContainsBoundedRuntimeReviewContract(t *testing.T) {
	content := strings.ToLower(string(mustReadEmbeddedWorkflowFile(t,
		"content/skills/verify/references/review-contract.md")))
	for _, clause := range []string{
		"runtime subagent", "read-only ownership", "state: findings",
		"primary agent owns integration", "never stage, commit, push",
	} {
		if !strings.Contains(content, clause) {
			t.Errorf("review contract missing %q", clause)
		}
	}
}

func TestWorkflowPacksHaveFocusedDiscoverableSkills(t *testing.T) {
	catalog, err := loadCatalog()
	if err != nil {
		t.Fatalf("loadCatalog returned error: %v", err)
	}
	wants := map[string]map[string]bool{
		"coding":  {"refine": true, "spec": true, "tdd": true, "bugfix": true},
		"quality": {"verify": true, "docs": true, "git-commit": true},
	}
	for packName, want := range wants {
		pack, ok := catalog.Pack(packName)
		if !ok {
			t.Fatalf("catalog missing %s pack", packName)
		}
		got := packSourceBasenames(pack, "content/skills/", "")
		if len(got) != len(want) {
			t.Fatalf("%s discoverable skills = %v, want %v", packName, got, want)
		}
		for name := range want {
			if !got[name] {
				t.Errorf("%s missing discoverable skill %q", packName, name)
			}
		}
	}
	base, ok := catalog.Pack("base")
	if !ok || len(base.Files) != 0 || strings.Join(base.Depends, ",") != "coding,quality" {
		t.Fatalf("base pack = %+v, want dependency-only coding+quality", base)
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
			if strings.HasPrefix(entry.Source, "content/skills/") &&
				strings.HasSuffix(entry.Source, "/SKILL.md") &&
				strings.Count(strings.TrimPrefix(entry.Source, "content/skills/"), "/") > 1 {
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
		"game-dev":      {"game-core", "game-ui", "game-systems", "game-performance", "game-audio", "game-art", "game-delivery", "game-dev"},
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

func sortedKeys(items map[string]bool) []string {
	keys := make([]string, 0, len(items))
	for key := range items {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
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
	const root = "content/skills"
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
	const prefix = "content/skills/"
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
