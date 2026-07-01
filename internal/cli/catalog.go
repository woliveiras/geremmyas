package cli

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"sort"

	geremmyas "github.com/woliveiras/geremmyas"
)

type Catalog struct {
	Packs  []Pack `json:"packs"`
	byName map[string]Pack
}

type Pack struct {
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Tier        string      `json:"tier"`
	Depends     []string    `json:"depends"`
	Files       []FileEntry `json:"files"`
}

type FileEntry struct {
	Source string `json:"source"`
	Target string `json:"target"`
}

// Valid pack tiers.
const (
	TierCore  = "core"
	TierStack = "stack"
)

func loadCatalog() (Catalog, error) {
	data, err := fs.ReadFile(geremmyas.EmbeddedFiles, "catalog/packs.json")
	if err != nil {
		return Catalog{}, err
	}

	var catalog Catalog
	if err := json.Unmarshal(data, &catalog); err != nil {
		return Catalog{}, err
	}

	catalog.byName = map[string]Pack{}
	for _, pack := range catalog.Packs {
		if pack.Name == "" {
			return Catalog{}, fmt.Errorf("catalog contains a pack without name")
		}
		if _, exists := catalog.byName[pack.Name]; exists {
			return Catalog{}, fmt.Errorf("catalog contains duplicate pack %q", pack.Name)
		}
		catalog.byName[pack.Name] = pack
	}

	sort.Slice(catalog.Packs, func(i, j int) bool {
		return catalog.Packs[i].Name < catalog.Packs[j].Name
	})

	return catalog, nil
}

func (c Catalog) Pack(name string) (Pack, bool) {
	pack, ok := c.byName[name]
	return pack, ok
}

func (c Catalog) Resolve(names []string) ([]Pack, error) {
	visited := map[string]bool{}
	visiting := map[string]bool{}
	resolved := []Pack{}

	var visit func(string) error
	visit = func(name string) error {
		if visited[name] {
			return nil
		}
		if visiting[name] {
			return fmt.Errorf("cyclic pack dependency at %q", name)
		}

		pack, ok := c.Pack(name)
		if !ok {
			return fmt.Errorf("unknown pack %q", name)
		}

		visiting[name] = true
		for _, dep := range pack.Depends {
			if err := visit(dep); err != nil {
				return err
			}
		}
		visiting[name] = false
		visited[name] = true
		resolved = append(resolved, pack)
		return nil
	}

	for _, name := range uniqueStrings(names) {
		if err := visit(name); err != nil {
			return nil, err
		}
	}

	return resolved, nil
}

func (c Catalog) ValidateSources() error {
	for _, pack := range c.Packs {
		for _, entry := range pack.Files {
			if entry.Source == "" || entry.Target == "" {
				return fmt.Errorf("pack %q contains invalid file entry", pack.Name)
			}
			if _, err := fs.Stat(geremmyas.EmbeddedFiles, entry.Source); err != nil {
				return fmt.Errorf("pack %q references missing source %q: %w", pack.Name, entry.Source, err)
			}
		}
	}
	return nil
}

// ValidateTiers checks that every pack declares a valid tier.
func (c Catalog) ValidateTiers() error {
	for _, pack := range c.Packs {
		switch pack.Tier {
		case TierCore, TierStack:
			// valid
		case "":
			return fmt.Errorf("pack %q is missing a tier (want %q or %q)", pack.Name, TierCore, TierStack)
		default:
			return fmt.Errorf("pack %q has invalid tier %q (want %q or %q)", pack.Name, pack.Tier, TierCore, TierStack)
		}
	}
	return nil
}
