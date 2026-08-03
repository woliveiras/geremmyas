package cli

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"

	geremmyas "github.com/woliveiras/geremmyas"
)

type legacyArtifactFile struct {
	Kind   string `json:"kind"`
	Path   string `json:"path"`
	SHA256 string `json:"sha256"`
}

type legacyArtifactCatalog struct {
	Files []legacyArtifactFile `json:"files"`
}

var sha256Pattern = regexp.MustCompile(`^[0-9a-f]{64}$`)

func loadLegacyArtifactCatalog() (legacyArtifactCatalog, error) {
	data, err := fs.ReadFile(geremmyas.EmbeddedFiles, "catalog/legacy-artifacts.json")
	if err != nil {
		return legacyArtifactCatalog{}, err
	}
	var catalog legacyArtifactCatalog
	if err := json.Unmarshal(data, &catalog); err != nil {
		return legacyArtifactCatalog{}, err
	}
	seen := map[string]bool{}
	for _, file := range catalog.Files {
		if file.Kind != ArtifactSkill && file.Kind != ArtifactAgent {
			return legacyArtifactCatalog{}, fmt.Errorf("invalid legacy artifact kind %q", file.Kind)
		}
		if err := validateLogicalPath(file.Path); err != nil {
			return legacyArtifactCatalog{}, fmt.Errorf("invalid legacy artifact path %q: %w", file.Path, err)
		}
		key := file.Kind + ":" + filepath.ToSlash(filepath.Clean(file.Path))
		if seen[key] {
			return legacyArtifactCatalog{}, fmt.Errorf("duplicate legacy artifact %q", key)
		}
		seen[key] = true
		if !sha256Pattern.MatchString(file.SHA256) {
			return legacyArtifactCatalog{}, fmt.Errorf("invalid legacy artifact hash for %q", file.Path)
		}
	}
	return catalog, nil
}

func addLegacyProjectHashes(known map[string]string, root string) error {
	catalog, err := loadLegacyArtifactCatalog()
	if err != nil {
		return err
	}
	for _, file := range catalog.Files {
		var bases []string
		switch file.Kind {
		case ArtifactSkill:
			bases = []string{".agents/skills", ".github/skills"}
		case ArtifactAgent:
			bases = []string{".agents/roles", ".github/agents"}
		}
		for _, base := range bases {
			known[filepath.Join(root, filepath.FromSlash(base), filepath.FromSlash(file.Path))] = file.SHA256
		}
	}
	return nil
}

func addLegacyGlobalHashes(known map[string]string) error {
	catalog, err := loadLegacyArtifactCatalog()
	if err != nil {
		return err
	}
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	for _, file := range catalog.Files {
		if file.Kind == ArtifactSkill {
			known[filepath.Join(home, ".agents", "skills", filepath.FromSlash(file.Path))] = file.SHA256
		}
	}
	return nil
}
