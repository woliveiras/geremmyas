package cli

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strings"

	geremmyas "github.com/woliveiras/geremmyas"
)

const globalManifestVersion = 1

type globalManifest struct {
	Version int               `json:"version"`
	Packs   []string          `json:"packs"`
	Targets []string          `json:"targets"`
	Files   map[string]string `json:"files"`
}

type globalReconcileSummary struct {
	Removed   int
	Preserved int
}

func globalManifestPath() (string, error) {
	stateHome := strings.TrimSpace(os.Getenv("XDG_STATE_HOME"))
	if stateHome == "" {
		home, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		stateHome = filepath.Join(home, ".local", "state")
	}
	return filepath.Join(stateHome, "geremmyas", "global-manifest.json"), nil
}

func loadGlobalManifest() (globalManifest, bool, error) {
	path, err := globalManifestPath()
	if err != nil {
		return globalManifest{}, false, err
	}
	data, err := os.ReadFile(path)
	if os.IsNotExist(err) {
		return globalManifest{Version: globalManifestVersion, Files: map[string]string{}}, false, nil
	}
	if err != nil {
		return globalManifest{}, false, err
	}
	var manifest globalManifest
	if err := json.Unmarshal(data, &manifest); err != nil {
		return globalManifest{}, true, fmt.Errorf("read global manifest %s: %w", path, err)
	}
	if manifest.Version != globalManifestVersion {
		return globalManifest{}, true, fmt.Errorf("unsupported global manifest version %d", manifest.Version)
	}
	if manifest.Files == nil {
		manifest.Files = map[string]string{}
	}
	return manifest, true, nil
}

func adoptKnownGlobalFiles(manifest globalManifest, catalog Catalog) (globalManifest, error) {
	known := map[string]string{}
	for _, pack := range catalog.Packs {
		for _, entry := range pack.Files {
			baseDir, relPath, ok := globalDestination(entry.Target)
			if !ok {
				continue
			}
			if err := addEmbeddedDestinationHashes(known, baseDir, relPath, entry.Source); err != nil {
				return manifest, err
			}
		}
	}
	for path, expectedHash := range known {
		actualHash, err := fileSHA256(path)
		if os.IsNotExist(err) {
			continue
		}
		if err != nil {
			return manifest, err
		}
		if actualHash == expectedHash {
			manifest.Files[path] = expectedHash
		}
	}
	return manifest, nil
}

func writeGlobalManifest(manifest globalManifest) error {
	path, err := globalManifestPath()
	if err != nil {
		return err
	}
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	data, err := json.MarshalIndent(manifest, "", "  ")
	if err != nil {
		return err
	}
	data = append(data, '\n')
	tmp, err := os.CreateTemp(filepath.Dir(path), ".global-manifest-*")
	if err != nil {
		return err
	}
	tmpPath := tmp.Name()
	defer os.Remove(tmpPath)
	if _, err := tmp.Write(data); err != nil {
		tmp.Close()
		return err
	}
	if err := tmp.Chmod(0o644); err != nil {
		tmp.Close()
		return err
	}
	if err := tmp.Close(); err != nil {
		return err
	}
	return os.Rename(tmpPath, path)
}

func reconcileGlobalManifest(previous globalManifest, desiredPaths []string, packNames, targets []string) (globalReconcileSummary, error) {
	desired := map[string]bool{}
	for _, path := range desiredPaths {
		desired[filepath.Clean(path)] = true
	}
	nextFiles := map[string]string{}
	summary := globalReconcileSummary{}

	for path, installedHash := range previous.Files {
		path = filepath.Clean(path)
		if desired[path] {
			continue
		}
		if !isManagedGlobalPath(path) {
			return summary, fmt.Errorf("manifest path is outside managed roots: %s", path)
		}
		currentHash, err := fileSHA256(path)
		if os.IsNotExist(err) {
			continue
		}
		if err != nil {
			return summary, err
		}
		if currentHash != installedHash {
			nextFiles[path] = installedHash
			summary.Preserved++
			continue
		}
		if err := os.Remove(path); err != nil {
			return summary, err
		}
		removeEmptyManagedParents(filepath.Dir(path))
		summary.Removed++
	}

	for path := range desired {
		hash, err := fileSHA256(path)
		if os.IsNotExist(err) {
			continue
		}
		if err != nil {
			return summary, err
		}
		nextFiles[path] = hash
	}

	sort.Strings(packNames)
	sort.Strings(targets)
	next := globalManifest{
		Version: globalManifestVersion,
		Packs:   uniqueStrings(packNames),
		Targets: uniqueStrings(targets),
		Files:   nextFiles,
	}
	return summary, writeGlobalManifest(next)
}

func fileSHA256(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	sum := sha256.Sum256(data)
	return hex.EncodeToString(sum[:]), nil
}

func bytesSHA256(data []byte) string {
	sum := sha256.Sum256(data)
	return hex.EncodeToString(sum[:])
}

func isManagedGlobalPath(path string) bool {
	home, err := os.UserHomeDir()
	if err != nil {
		return false
	}
	roots := []string{
		filepath.Join(home, ".agents", "skills"),
		filepath.Join(home, ".copilot", "instructions"),
		filepath.Join(home, ".cursor", "rules"),
		filepath.Join(home, ".cursor", "hooks"),
		filepath.Join(home, ".cursor", "hooks.json"),
		filepath.Join(home, ".claude", "CLAUDE.md"),
		filepath.Join(home, ".config", "opencode", "AGENTS.md"),
		filepath.Join(home, ".codex", "AGENTS.md"),
		filepath.Join(home, ".codex", "instructions"),
	}
	path = filepath.Clean(path)
	for _, root := range roots {
		root = filepath.Clean(root)
		if path == root || strings.HasPrefix(path, root+string(filepath.Separator)) {
			return true
		}
	}
	return false
}

func removeEmptyManagedParents(dir string) {
	for isManagedGlobalPath(dir) {
		if err := os.Remove(dir); err != nil {
			return
		}
		dir = filepath.Dir(dir)
	}
}

func globalDesiredPaths(packs []Pack, targets []string) ([]string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}
	paths := map[string]bool{}
	copySkills, copyInstructions := globalCopyFlags(targets)
	for _, pack := range packs {
		for _, entry := range pack.Files {
			baseDir, relPath, ok := globalDestination(entry.Target)
			if !ok || (strings.HasPrefix(entry.Target, ".github/skills/") && !copySkills) ||
				(strings.HasPrefix(entry.Target, ".github/instructions/") && !copyInstructions) {
				continue
			}
			if err := addEmbeddedDestinationPaths(paths, baseDir, relPath, entry.Source); err != nil {
				return nil, err
			}
		}
	}

	artifacts := collectPackArtifacts(packs)
	if hasTarget(targets, TargetCursor) {
		for _, source := range artifacts.instructions {
			base := strings.TrimSuffix(filepath.Base(source), ".instructions.md")
			paths[filepath.Join(home, ".cursor", "rules", base+".mdc")] = true
		}
		for _, source := range artifacts.skills {
			skillMD, err := findSkillMarkdown(source)
			if err != nil {
				return nil, err
			}
			content, _ := readEmbeddedSource(skillMD)
			fm, _, _ := parseMarkdownFrontmatter(content)
			name := fm.get("name")
			if name == "" {
				name = filepath.Base(filepath.Dir(skillMD))
			}
			paths[filepath.Join(home, ".cursor", "rules", "skill-"+name+".mdc")] = true
		}
		for _, source := range artifacts.agents {
			name := strings.TrimSuffix(filepath.Base(source), ".agent.md")
			paths[filepath.Join(home, ".cursor", "rules", "agent-"+name+".mdc")] = true
		}
		if len(artifacts.agents) > 0 {
			paths[filepath.Join(home, ".cursor", "rules", "geremmyas-agents.mdc")] = true
		}
		if artifacts.hasHooks {
			paths[filepath.Join(home, ".cursor", "hooks", "guardrails.sh")] = true
			paths[filepath.Join(home, ".cursor", "hooks", "guardrails-rules.txt")] = true
			paths[filepath.Join(home, ".cursor", "hooks.json")] = true
		}
	}
	if hasTarget(targets, TargetClaudeCode) {
		paths[filepath.Join(home, ".claude", "CLAUDE.md")] = true
	}
	if hasTarget(targets, TargetOpenCode) {
		paths[filepath.Join(home, ".config", "opencode", "AGENTS.md")] = true
	}
	if hasTarget(targets, TargetCodex) {
		paths[filepath.Join(home, ".codex", "AGENTS.md")] = true
		for _, source := range artifacts.instructions {
			paths[filepath.Join(home, ".codex", "instructions", filepath.Base(source))] = true
		}
	}

	out := make([]string, 0, len(paths))
	for path := range paths {
		out = append(out, filepath.Clean(path))
	}
	sort.Strings(out)
	return out, nil
}

func addEmbeddedDestinationPaths(paths map[string]bool, baseDir, relPath, source string) error {
	info, err := fs.Stat(geremmyas.EmbeddedFiles, source)
	if err != nil {
		return err
	}
	if !info.IsDir() {
		paths[filepath.Join(baseDir, relPath)] = true
		return nil
	}
	return fs.WalkDir(geremmyas.EmbeddedFiles, source, func(path string, d fs.DirEntry, err error) error {
		if err != nil || d.IsDir() {
			return err
		}
		rel, err := filepath.Rel(source, path)
		if err != nil {
			return err
		}
		paths[filepath.Join(baseDir, relPath, rel)] = true
		return nil
	})
}

func addEmbeddedDestinationHashes(paths map[string]string, baseDir, relPath, source string) error {
	info, err := fs.Stat(geremmyas.EmbeddedFiles, source)
	if err != nil {
		return err
	}
	if !info.IsDir() {
		data, err := fs.ReadFile(geremmyas.EmbeddedFiles, source)
		if err != nil {
			return err
		}
		paths[filepath.Join(baseDir, relPath)] = bytesSHA256(data)
		return nil
	}
	return fs.WalkDir(geremmyas.EmbeddedFiles, source, func(path string, d fs.DirEntry, err error) error {
		if err != nil || d.IsDir() {
			return err
		}
		rel, err := filepath.Rel(source, path)
		if err != nil {
			return err
		}
		data, err := fs.ReadFile(geremmyas.EmbeddedFiles, path)
		if err != nil {
			return err
		}
		paths[filepath.Join(baseDir, relPath, rel)] = bytesSHA256(data)
		return nil
	})
}
