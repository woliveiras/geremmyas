package cli

import (
	"strings"
	"testing"
)

func TestCatalogArtifactKindsAreExplicitAndValid(t *testing.T) {
	catalog, err := loadCatalog()
	if err != nil {
		t.Fatalf("loadCatalog returned error: %v", err)
	}
	if err := catalog.ValidateArtifactKinds(); err != nil {
		t.Fatalf("ValidateArtifactKinds returned error: %v", err)
	}
}

func TestCatalogSourcesUseNeutralOrTargetRoots(t *testing.T) {
	catalog, err := loadCatalog()
	if err != nil {
		t.Fatalf("loadCatalog returned error: %v", err)
	}

	for _, pack := range catalog.Packs {
		for _, entry := range pack.Files {
			if strings.HasPrefix(entry.Source, "content/") ||
				strings.HasPrefix(entry.Source, "targets/") {
				continue
			}
			t.Errorf("pack %q source %q is outside content/ and targets/", pack.Name, entry.Source)
		}
	}
}

func TestValidateArtifactKindsRejectsMissingAndUnknownKinds(t *testing.T) {
	tests := []struct {
		name string
		kind string
	}{
		{name: "missing"},
		{name: "unknown", kind: "assistant-magic"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			catalog := Catalog{Packs: []Pack{{
				Name: "broken",
				Files: []FileEntry{{
					Kind:   tt.kind,
					Source: "content/AGENTS.md",
					Path:   "AGENTS.md",
					Target: "AGENTS.md",
				}},
			}}}
			err := catalog.ValidateArtifactKinds()
			if err == nil {
				t.Fatal("ValidateArtifactKinds succeeded, want error")
			}
			if !strings.Contains(err.Error(), "broken") {
				t.Fatalf("error %q does not name pack", err)
			}
		})
	}
}

func TestDoctorRejectsUnknownArtifactKind(t *testing.T) {
	catalog := Catalog{Packs: []Pack{{
		Name: "broken",
		Tier: TierStack,
		Files: []FileEntry{{
			Kind:   "assistant-magic",
			Source: "content/AGENTS.md",
			Path:   "AGENTS.md",
			Target: "AGENTS.md",
		}},
	}}}

	var out strings.Builder
	err := runDoctor(&out, catalog)
	if err == nil {
		t.Fatal("runDoctor succeeded, want invalid artifact kind error")
	}
	if !strings.Contains(err.Error(), "assistant-magic") {
		t.Fatalf("error %q does not name invalid kind", err)
	}
}

func TestPlanProjectArtifactsUsesKindsInsteadOfGitHubTargets(t *testing.T) {
	packs := []Pack{{
		Name: "test",
		Files: []FileEntry{
			{
				Kind:   ArtifactContract,
				Source: "content/AGENTS.md",
				Path:   "AGENTS.md",
				Target: "AGENTS.md",
			},
			{
				Kind:   ArtifactSkill,
				Source: "content/skills/vertical-tdd",
				Path:   "vertical-tdd",
				Target: ".github/skills/vertical-tdd",
			},
			{
				Kind:   ArtifactInstruction,
				Source: "content/instructions/go.instructions.md",
				Path:   "go.instructions.md",
				Target: ".github/instructions/go.instructions.md",
			},
			{
				Kind:   ArtifactCopilotInstructions,
				Source: "targets/copilot/project-instructions.md",
				Path:   "copilot-instructions.md",
				Target: ".github/copilot-instructions.md",
			},
		},
	}}

	got, err := planProjectArtifacts(packs, []string{TargetCodex, TargetCopilot})
	if err != nil {
		t.Fatalf("planProjectArtifacts returned error: %v", err)
	}
	want := []string{
		"AGENTS.md",
		".agents/skills/vertical-tdd",
		".codex/instructions/go.instructions.md",
		".github/copilot-instructions.md",
		".github/instructions/go.instructions.md",
		".github/skills/vertical-tdd",
	}
	if strings.Join(plannedTargets(got), ",") != strings.Join(want, ",") {
		t.Fatalf("targets = %v, want %v", plannedTargets(got), want)
	}
}

func TestPlanProjectArtifactsIsStableAndDeduplicated(t *testing.T) {
	entry := FileEntry{
		Kind:   ArtifactSkill,
		Source: "content/skills/vertical-tdd",
		Path:   "vertical-tdd",
		Target: ".github/skills/vertical-tdd",
	}
	packs := []Pack{
		{Name: "first", Files: []FileEntry{entry}},
		{Name: "second", Files: []FileEntry{entry}},
	}

	got, err := planProjectArtifacts(packs, []string{TargetOpenCode, TargetCodex, TargetCodex})
	if err != nil {
		t.Fatalf("planProjectArtifacts returned error: %v", err)
	}
	want := []string{".agents/skills/vertical-tdd"}
	if strings.Join(plannedTargets(got), ",") != strings.Join(want, ",") {
		t.Fatalf("targets = %v, want %v", plannedTargets(got), want)
	}
}

func TestPlanProjectArtifactsRejectsUnsafeLogicalPath(t *testing.T) {
	packs := []Pack{{
		Name: "unsafe",
		Files: []FileEntry{{
			Kind:   ArtifactSkill,
			Source: "content/skills/unsafe",
			Path:   "../unsafe",
			Target: ".github/skills/unsafe",
		}},
	}}

	if _, err := planProjectArtifacts(packs, []string{TargetCodex}); err == nil {
		t.Fatal("planProjectArtifacts succeeded, want unsafe path error")
	}
}

func plannedTargets(entries []FileEntry) []string {
	targets := make([]string, 0, len(entries))
	for _, entry := range entries {
		targets = append(targets, entry.Target)
	}
	return targets
}
