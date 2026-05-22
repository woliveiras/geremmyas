package cli

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	geremmyas "github.com/woliveiras/geremmyas"
)

func vsCodeUserDir() (string, error) {
	switch runtime.GOOS {
	case "darwin":
		home, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		return filepath.Join(home, "Library", "Application Support", "Code", "User"), nil
	case "linux":
		if dir := os.Getenv("XDG_CONFIG_HOME"); dir != "" {
			return filepath.Join(dir, "Code", "User"), nil
		}
		home, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		return filepath.Join(home, ".config", "Code", "User"), nil
	case "windows":
		appdata := os.Getenv("APPDATA")
		if appdata == "" {
			return "", fmt.Errorf("APPDATA not set")
		}
		return filepath.Join(appdata, "Code", "User"), nil
	default:
		return "", fmt.Errorf("unsupported platform %q", runtime.GOOS)
	}
}

func globalInstallPacks(packs []Pack) error {
	userDir, err := vsCodeUserDir()
	if err != nil {
		return fmt.Errorf("detecting VS Code user directory: %w", err)
	}

	for _, pack := range packs {
		for _, entry := range pack.Files {
			if !isGlobalTarget(entry.Target) {
				continue
			}
			if err := globalCopyEntry(userDir, entry); err != nil {
				return fmt.Errorf("pack %q: %w", pack.Name, err)
			}
		}
	}

	return nil
}

// isGlobalTarget returns true for files that belong in the VS Code user-level
// directory (inside .github/). Project-root files like AGENTS.md and mise.toml
// are skipped during global install.
func isGlobalTarget(target string) bool {
	return strings.HasPrefix(target, ".github/") || strings.HasPrefix(target, ".github\\")
}

func globalCopyEntry(userDir string, entry FileEntry) error {
	info, err := fs.Stat(geremmyas.EmbeddedFiles, entry.Source)
	if err != nil {
		return err
	}

	if !info.IsDir() {
		return globalCopyFile(userDir, entry.Source, entry.Target)
	}

	return fs.WalkDir(geremmyas.EmbeddedFiles, entry.Source, func(path string, d fs.DirEntry, err error) error {
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
		target := filepath.Join(entry.Target, rel)
		return globalCopyFile(userDir, path, target)
	})
}

func globalCopyFile(userDir, source, target string) error {
	data, err := fs.ReadFile(geremmyas.EmbeddedFiles, filepath.ToSlash(source))
	if err != nil {
		return err
	}

	dest := filepath.Join(userDir, target)
	if err := os.MkdirAll(filepath.Dir(dest), 0o755); err != nil {
		return err
	}
	return os.WriteFile(dest, data, 0o644)
}
