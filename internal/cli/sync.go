package cli

import (
	"bytes"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	geremmyas "github.com/woliveiras/geremmyas"
)

type syncOptions struct {
	Force bool
}

type syncSummary struct {
	Installed int
	Updated   int
	Preserved int
	Skipped   int
}

func syncPacks(root string, packs []Pack, opts syncOptions) (syncSummary, error) {
	summary := syncSummary{}
	copiedTargets := map[string]bool{}

	for _, pack := range packs {
		for _, entry := range pack.Files {
			if err := copyEntry(root, entry, opts, &summary, copiedTargets); err != nil {
				return summary, fmt.Errorf("pack %q: %w", pack.Name, err)
			}
		}
	}

	return summary, nil
}

func copyEntry(root string, entry FileEntry, opts syncOptions, summary *syncSummary, copiedTargets map[string]bool) error {
	info, err := fs.Stat(geremmyas.EmbeddedFiles, entry.Source)
	if err != nil {
		return err
	}

	if !info.IsDir() {
		return copyFile(root, entry.Source, entry.Target, opts, summary, copiedTargets)
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
		return copyFile(root, path, target, opts, summary, copiedTargets)
	})
}

func copyFile(root, source, target string, opts syncOptions, summary *syncSummary, copiedTargets map[string]bool) error {
	target = filepath.Clean(target)
	if target == "." || strings.HasPrefix(target, ".."+string(filepath.Separator)) || filepath.IsAbs(target) {
		return fmt.Errorf("unsafe target path %q", target)
	}
	if copiedTargets[target] {
		summary.Skipped++
		return nil
	}
	copiedTargets[target] = true

	data, err := fs.ReadFile(geremmyas.EmbeddedFiles, filepath.ToSlash(source))
	if err != nil {
		return err
	}

	dest := filepath.Join(root, target)
	current, readErr := os.ReadFile(dest)
	if readErr == nil {
		if bytes.Equal(current, data) {
			summary.Skipped++
			return nil
		}
		if isCustomizableTarget(target) && !opts.Force {
			summary.Preserved++
			return nil
		}
		summary.Updated++
	} else if os.IsNotExist(readErr) {
		summary.Installed++
	} else {
		return readErr
	}

	if err := os.MkdirAll(filepath.Dir(dest), 0o755); err != nil {
		return err
	}
	return os.WriteFile(dest, data, 0o644)
}

func isCustomizableTarget(target string) bool {
	target = filepath.ToSlash(filepath.Clean(target))
	switch target {
	case "AGENTS.md",
		"specs/README.md",
		"mise.toml",
		".github/copilot-instructions.md",
		".github/hooks/guardrails-rules.txt":
		return true
	default:
		return false
	}
}
