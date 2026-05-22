package cli

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	geremmyas "github.com/woliveiras/geremmyas"
)

// globalDestination resolves the user-level path for a given project target.
// Returns empty string if the target has no user-level equivalent.
func globalDestination(target string) (baseDir string, relPath string, ok bool) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", "", false
	}

	switch {
	case strings.HasPrefix(target, ".github/skills/"):
		rel := strings.TrimPrefix(target, ".github/skills/")
		return filepath.Join(home, ".agents", "skills"), rel, true
	case strings.HasPrefix(target, ".github/instructions/"):
		rel := strings.TrimPrefix(target, ".github/instructions/")
		return filepath.Join(home, ".copilot", "instructions"), rel, true
	default:
		// agents, hooks, copilot-instructions.md, AGENTS.md, mise.toml
		// are project-level only
		return "", "", false
	}
}

func globalInstallDir() string {
	home, _ := os.UserHomeDir()
	return home
}

func globalInstallPacks(packs []Pack) (int, error) {
	count := 0
	for _, pack := range packs {
		for _, entry := range pack.Files {
			baseDir, relPath, ok := globalDestination(entry.Target)
			if !ok {
				continue
			}
			copied, err := globalCopyEntry(baseDir, relPath, entry)
			if err != nil {
				return count, fmt.Errorf("pack %q: %w", pack.Name, err)
			}
			count += copied
		}
	}
	return count, nil
}

func globalCopyEntry(baseDir, relPath string, entry FileEntry) (int, error) {
	info, err := fs.Stat(geremmyas.EmbeddedFiles, entry.Source)
	if err != nil {
		return 0, err
	}

	if !info.IsDir() {
		dest := filepath.Join(baseDir, relPath)
		if err := globalWriteFile(dest, entry.Source); err != nil {
			return 0, err
		}
		return 1, nil
	}

	count := 0
	err = fs.WalkDir(geremmyas.EmbeddedFiles, entry.Source, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		rel, err := filepath.Rel(entry.Source, path)
		if err != nil {
			return err
		}
		dest := filepath.Join(baseDir, relPath, rel)
		if err := globalWriteFile(dest, path); err != nil {
			return err
		}
		count++
		return nil
	})
	if err != nil {
		return count, err
	}
	return count, nil
}

func globalWriteFile(dest, source string) error {
	data, err := fs.ReadFile(geremmyas.EmbeddedFiles, filepath.ToSlash(source))
	if err != nil {
		return err
	}
	if err := os.MkdirAll(filepath.Dir(dest), 0o755); err != nil {
		return err
	}
	return os.WriteFile(dest, data, 0o644)
}
