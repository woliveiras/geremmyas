package cli

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

const projectManifestVersion = 1

type projectManifest struct {
	Version int               `json:"version"`
	Packs   []string          `json:"packs"`
	Targets []string          `json:"targets"`
	Files   map[string]string `json:"files"`
}

type projectReconcileSummary struct {
	Removed   int
	Preserved int
}

type projectSyncResult struct {
	Sync       syncSummary
	Generators map[string]generatorSummary
	Reconcile  projectReconcileSummary
}

func projectManifestPath(root string) string {
	return filepath.Join(root, ".geremmyas", "project-manifest.json")
}

func loadProjectManifest(root string) (projectManifest, bool, error) {
	path := projectManifestPath(root)
	data, err := os.ReadFile(path)
	if os.IsNotExist(err) {
		return projectManifest{Version: projectManifestVersion, Files: map[string]string{}}, false, nil
	}
	if err != nil {
		return projectManifest{}, false, err
	}
	var manifest projectManifest
	if err := json.Unmarshal(data, &manifest); err != nil {
		return projectManifest{}, true, fmt.Errorf("read project manifest %s: %w", path, err)
	}
	if manifest.Version != projectManifestVersion {
		return projectManifest{}, true, fmt.Errorf("unsupported project manifest version %d", manifest.Version)
	}
	if manifest.Files == nil {
		manifest.Files = map[string]string{}
	}
	return manifest, true, nil
}

func writeProjectManifest(root string, manifest projectManifest) error {
	path := projectManifestPath(root)
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	data, err := json.MarshalIndent(manifest, "", "  ")
	if err != nil {
		return err
	}
	data = append(data, '\n')
	tmp, err := os.CreateTemp(filepath.Dir(path), ".project-manifest-*")
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

func syncProjectState(root string, catalog Catalog, packs []Pack, targets []string, opts syncOptions) (projectSyncResult, error) {
	result := projectSyncResult{}
	previous, exists, err := loadProjectManifest(root)
	if err != nil {
		return result, err
	}
	if !exists {
		previous, err = adoptKnownLegacyProjectFiles(root, previous, catalog)
		if err != nil {
			return result, err
		}
	}

	planned, err := planProjectArtifacts(packs, targets)
	if err != nil {
		return result, err
	}
	result.Sync, err = syncEntries(root, planned, opts)
	if err != nil {
		return result, err
	}
	result.Generators, err = runTargetGenerators(root, targets, packs, generatorOptions{Force: opts.Force})
	if err != nil {
		return result, err
	}

	desired, err := projectDesiredHashes(root, planned, packs, targets)
	if err != nil {
		return result, err
	}
	result.Reconcile, err = reconcileProjectManifest(root, previous, desired, resolvedPackNames(packs), targets)
	return result, err
}

func resolvedPackNames(packs []Pack) []string {
	names := make([]string, 0, len(packs))
	for _, pack := range packs {
		names = append(names, pack.Name)
	}
	return names
}

func projectDesiredHashes(root string, planned []FileEntry, packs []Pack, targets []string) (map[string]string, error) {
	expected := map[string]string{}
	for _, entry := range planned {
		if err := addEmbeddedDestinationHashes(expected, root, entry.Target, entry.Source); err != nil {
			return nil, err
		}
	}

	desired := map[string]string{}
	for path, expectedHash := range expected {
		if ownedProjectFileMatches(path, expectedHash) {
			rel, err := filepath.Rel(root, path)
			if err != nil {
				return nil, err
			}
			desired[filepath.ToSlash(rel)] = expectedHash
		}
	}

	for _, rel := range projectGeneratedPaths(packs, targets) {
		path := filepath.Join(root, filepath.FromSlash(rel))
		info, err := os.Lstat(path)
		if err != nil || !info.Mode().IsRegular() {
			continue
		}
		data, err := os.ReadFile(path)
		if err != nil {
			return nil, err
		}
		if !isGeneratedFile(data) && rel != ".cursor/hooks/guardrails-rules.txt" {
			continue
		}
		desired[rel] = bytesSHA256(data)
	}
	return desired, nil
}

func projectGeneratedPaths(packs []Pack, targets []string) []string {
	artifacts := collectPackArtifacts(packs)
	paths := map[string]bool{}
	if hasTarget(targets, TargetCursor) {
		for _, source := range artifacts.instructions {
			base := strings.TrimSuffix(filepath.Base(source), ".instructions.md")
			paths[".cursor/rules/"+base+".mdc"] = true
		}
		for _, source := range artifacts.skills {
			skillMD, err := findSkillMarkdown(source)
			if err != nil {
				continue
			}
			content, _ := readEmbeddedSource(skillMD)
			fm, _, _ := parseMarkdownFrontmatter(content)
			name := fm.get("name")
			if name == "" {
				name = filepath.Base(filepath.Dir(skillMD))
			}
			paths[".cursor/rules/skill-"+name+".mdc"] = true
		}
		for _, source := range artifacts.agents {
			name := strings.TrimSuffix(filepath.Base(source), ".agent.md")
			paths[".cursor/rules/agent-"+name+".mdc"] = true
		}
		if len(artifacts.agents) > 0 {
			paths[".cursor/rules/geremmyas-agents.mdc"] = true
		}
		if artifacts.hasHooks {
			paths[".cursor/hooks/guardrails.sh"] = true
			paths[".cursor/hooks/guardrails-rules.txt"] = true
			paths[".cursor/hooks.json"] = true
		}
	}
	if hasTarget(targets, TargetClaudeCode) {
		paths["CLAUDE.md"] = true
	}
	if hasTarget(targets, TargetOpenCode) {
		paths[".opencode/AGENTS.md"] = true
	}
	if hasTarget(targets, TargetCodex) {
		paths[".codex/AGENTS.md"] = true
	}

	out := make([]string, 0, len(paths))
	for path := range paths {
		out = append(out, path)
	}
	sort.Strings(out)
	return out
}

func adoptKnownLegacyProjectFiles(root string, manifest projectManifest, catalog Catalog) (projectManifest, error) {
	known := map[string]string{}
	legacy, err := planProjectArtifacts(catalog.Packs, []string{TargetCopilot})
	if err != nil {
		return manifest, err
	}
	for _, entry := range legacy {
		if err := addEmbeddedDestinationHashes(known, root, entry.Target, entry.Source); err != nil {
			return manifest, err
		}
	}
	for path, expectedHash := range known {
		if !ownedProjectFileMatches(path, expectedHash) {
			continue
		}
		rel, err := filepath.Rel(root, path)
		if err != nil {
			return manifest, err
		}
		manifest.Files[filepath.ToSlash(rel)] = expectedHash
	}
	return manifest, nil
}

func ownedProjectFileMatches(path, expectedHash string) bool {
	info, err := os.Lstat(path)
	if err != nil || !info.Mode().IsRegular() {
		return false
	}
	hash, err := fileSHA256(path)
	return err == nil && hash == expectedHash
}

func reconcileProjectManifest(root string, previous projectManifest, desired map[string]string, packs, targets []string) (projectReconcileSummary, error) {
	summary := projectReconcileSummary{}
	nextFiles := map[string]string{}
	for rel, installedHash := range previous.Files {
		rel = filepath.ToSlash(filepath.Clean(rel))
		if hash, ok := desired[rel]; ok {
			nextFiles[rel] = hash
			continue
		}
		if !isManagedProjectPath(rel) {
			return summary, fmt.Errorf("project manifest path is outside managed roots: %s", rel)
		}
		path := filepath.Join(root, filepath.FromSlash(rel))
		info, err := os.Lstat(path)
		if os.IsNotExist(err) {
			continue
		}
		if err != nil {
			return summary, err
		}
		if !info.Mode().IsRegular() {
			nextFiles[rel] = installedHash
			summary.Preserved++
			continue
		}
		hash, err := fileSHA256(path)
		if err != nil {
			return summary, err
		}
		if hash != installedHash {
			nextFiles[rel] = installedHash
			summary.Preserved++
			continue
		}
		if err := os.Remove(path); err != nil {
			return summary, err
		}
		removeEmptyProjectParents(root, filepath.Dir(path))
		summary.Removed++
	}
	for rel, hash := range desired {
		nextFiles[rel] = hash
	}
	sort.Strings(packs)
	sort.Strings(targets)
	next := projectManifest{
		Version: projectManifestVersion,
		Packs:   uniqueStrings(packs),
		Targets: uniqueStrings(targets),
		Files:   nextFiles,
	}
	return summary, writeProjectManifest(root, next)
}

func isManagedProjectPath(rel string) bool {
	rel = filepath.ToSlash(filepath.Clean(rel))
	if rel == "" || rel == "." || rel == ".." || strings.HasPrefix(rel, "../") || filepath.IsAbs(rel) {
		return false
	}
	if rel == "AGENTS.md" || rel == "CLAUDE.md" || rel == "mise.toml" || rel == "specs/README.md" {
		return true
	}
	for _, prefix := range []string{".agents/", ".claude/", ".codex/", ".cursor/", ".github/", ".opencode/"} {
		if strings.HasPrefix(rel, prefix) {
			return true
		}
	}
	return false
}

func removeEmptyProjectParents(root, dir string) {
	root = filepath.Clean(root)
	for dir != root && strings.HasPrefix(dir, root+string(filepath.Separator)) {
		if err := os.Remove(dir); err != nil {
			return
		}
		dir = filepath.Dir(dir)
	}
}
