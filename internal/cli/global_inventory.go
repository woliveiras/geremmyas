package cli

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

const globalInventorySchemaVersion = 1

type globalInventoryFile struct {
	Path          string   `json:"path"`
	Ownership     string   `json:"ownership"`
	Status        string   `json:"status"`
	Targets       []string `json:"targets,omitempty"`
	Shared        bool     `json:"shared"`
	InstalledHash string   `json:"installed_hash,omitempty"`
	CurrentHash   string   `json:"current_hash,omitempty"`
	Obsolete      bool     `json:"obsolete,omitempty"`
	Proof         string   `json:"proof,omitempty"`
}

type globalInventoryManifest struct {
	Path    string   `json:"path"`
	Exists  bool     `json:"exists"`
	Version int      `json:"version"`
	Packs   []string `json:"packs"`
	Targets []string `json:"targets"`
}

type globalInventorySummary struct {
	Managed    int `json:"managed"`
	Unowned    int `json:"unowned"`
	Current    int `json:"current"`
	Modified   int `json:"modified"`
	Missing    int `json:"missing"`
	Symlink    int `json:"symlink"`
	NonRegular int `json:"non_regular"`
	Unreadable int `json:"unreadable"`
	Adoptable  int `json:"adoptable"`
	Obsolete   int `json:"obsolete"`
}

type globalObservedRoot struct {
	Path      string `json:"path"`
	Ownership string `json:"ownership"`
	Kind      string `json:"kind"`
}

type globalInventoryReport struct {
	SchemaVersion    int                     `json:"schema_version"`
	Command          string                  `json:"command"`
	Manifest         globalInventoryManifest `json:"manifest"`
	RequestedTargets []string                `json:"requested_targets,omitempty"`
	CatalogResolved  bool                    `json:"catalog_resolved"`
	CatalogError     string                  `json:"catalog_error,omitempty"`
	Summary          globalInventorySummary  `json:"summary"`
	Files            []globalInventoryFile   `json:"files"`
	Observed         []globalObservedRoot    `json:"observed,omitempty"`
}

func runGlobalList(args []string, w io.Writer, catalog Catalog) error {
	fs := flag.NewFlagSet("global list", flag.ContinueOnError)
	fs.SetOutput(io.Discard)
	jsonOutput := fs.Bool("json", false, "write a stable JSON inventory")
	targetsFlag := fs.String("targets", "", "comma-separated target filter")
	includeAdoptable := fs.Bool("include-adoptable", false, "include exact canonical files without ownership")
	if err := fs.Parse(args); err != nil {
		return err
	}
	if fs.NArg() != 0 {
		return fmt.Errorf("global list does not accept pack names")
	}

	requestedTargets := splitCSV(*targetsFlag)
	if err := validateTargets(requestedTargets); err != nil {
		return err
	}
	if len(requestedTargets) > 0 {
		requestedTargets = normalizeTargets(requestedTargets)
	}

	report, err := collectGlobalInventory(catalog, requestedTargets, *includeAdoptable)
	if err != nil {
		return err
	}
	if *jsonOutput {
		encoder := json.NewEncoder(w)
		encoder.SetEscapeHTML(false)
		encoder.SetIndent("", "  ")
		return encoder.Encode(report)
	}
	printGlobalInventory(w, report)
	return nil
}

func collectGlobalInventory(catalog Catalog, requestedTargets []string, includeAdoptable bool) (globalInventoryReport, error) {
	manifestPath, err := globalManifestPath()
	if err != nil {
		return globalInventoryReport{}, err
	}
	manifest, exists, _, err := loadGlobalManifestSnapshot()
	if err != nil {
		return globalInventoryReport{}, err
	}
	return collectGlobalInventoryFromManifest(catalog, requestedTargets, includeAdoptable, manifestPath, manifest, exists)
}

func collectGlobalInventoryFromManifest(catalog Catalog, requestedTargets []string, includeAdoptable bool, manifestPath string, manifest globalManifest, exists bool) (globalInventoryReport, error) {
	manifest.Packs = sortedUniqueStrings(manifest.Packs)
	manifest.Targets = sortedUniqueStrings(manifest.Targets)
	if err := validateTargets(manifest.Targets); err != nil {
		return globalInventoryReport{}, fmt.Errorf("incompatible global manifest targets: %w", err)
	}

	report := globalInventoryReport{
		SchemaVersion: globalInventorySchemaVersion,
		Command:       "global.list",
		Manifest: globalInventoryManifest{
			Path:    manifestPath,
			Exists:  exists,
			Version: manifest.Version,
			Packs:   manifest.Packs,
			Targets: manifest.Targets,
		},
		RequestedTargets: append([]string(nil), requestedTargets...),
		Files:            []globalInventoryFile{},
	}

	installedFiles := make(map[string]string, len(manifest.Files))
	paths := make([]string, 0, len(manifest.Files))
	for rawPath, hash := range manifest.Files {
		path := filepath.Clean(rawPath)
		if _, exists := installedFiles[path]; exists {
			return globalInventoryReport{}, fmt.Errorf("global manifest paths collide after normalization: %q", path)
		}
		installedFiles[path] = hash
		paths = append(paths, path)
	}
	sort.Strings(paths)
	desired, desiredKnown, catalogErr, err := currentGlobalDesiredSet(catalog, manifest)
	if err != nil {
		return globalInventoryReport{}, err
	}
	report.CatalogResolved = desiredKnown
	report.CatalogError = catalogErr
	for _, path := range paths {
		targets, shared, err := globalPathTargets(path, manifest.Targets)
		if err != nil {
			return globalInventoryReport{}, err
		}
		if len(requestedTargets) > 0 && !targetsIntersect(targets, requestedTargets) {
			continue
		}
		entry, err := inspectGlobalInventoryFile(path, installedFiles[path], targets, shared)
		if err != nil {
			return globalInventoryReport{}, err
		}
		if desiredKnown && !desired[path] {
			entry.Obsolete = true
			if entry.Status == "current" {
				entry.Status = "obsolete"
			}
		}
		report.Files = append(report.Files, entry)
		addGlobalInventorySummary(&report.Summary, entry)
	}

	unowned, err := collectUnownedGlobalFiles(installedFiles, requestedTargets)
	if err != nil {
		return globalInventoryReport{}, err
	}
	for _, entry := range unowned {
		report.Files = append(report.Files, entry)
		addGlobalInventorySummary(&report.Summary, entry)
	}

	if includeAdoptable {
		adoptable, err := collectAdoptableGlobalFiles(catalog, installedFiles, requestedTargets)
		if err != nil {
			return globalInventoryReport{}, err
		}
		adoptableByPath := map[string]globalInventoryFile{}
		for _, entry := range adoptable {
			adoptableByPath[entry.Path] = entry
		}
		for i, entry := range report.Files {
			adoptableEntry, ok := adoptableByPath[entry.Path]
			if !ok && entry.Ownership == "unowned" && entry.CurrentHash != "" && isTrustedGeneratedGlobalFile(entry.Path) {
				adoptableEntry = entry
				adoptableEntry.Ownership = "adoptable"
				adoptableEntry.Status = "adoptable"
				adoptableEntry.Proof = "generated-marker"
				ok = true
			}
			if !ok || entry.Ownership != "unowned" {
				continue
			}
			report.Files[i] = adoptableEntry
			report.Summary.Unowned--
			report.Summary.Adoptable++
			delete(adoptableByPath, entry.Path)
		}
		for _, entry := range adoptableByPath {
			report.Files = append(report.Files, entry)
			report.Summary.Adoptable++
		}
	}
	sort.Slice(report.Files, func(i, j int) bool { return report.Files[i].Path < report.Files[j].Path })

	report.Observed = collectObservedGlobalRoots(requestedTargets)
	return report, nil
}

func isTrustedGeneratedGlobalFile(path string) bool {
	data, err := os.ReadFile(path)
	if err != nil {
		return false
	}
	return strings.Contains(string(data), generatedMarker)
}

func inspectGlobalInventoryFile(path, installedHash string, targets []string, shared bool) (globalInventoryFile, error) {
	entry := globalInventoryFile{
		Path:          filepath.Clean(path),
		Ownership:     "managed",
		Targets:       append([]string(nil), targets...),
		Shared:        shared,
		InstalledHash: installedHash,
		Proof:         "manifest-hash",
	}
	if !isManagedGlobalPath(path) {
		return entry, fmt.Errorf("manifest path is outside managed roots: %s", path)
	}
	symlink, err := globalPathContainsSymlink(path)
	if err != nil {
		return entry, err
	}
	if symlink {
		entry.Status = "symlink"
		return entry, nil
	}
	info, err := os.Lstat(path)
	if os.IsNotExist(err) {
		entry.Status = "missing"
		return entry, nil
	}
	if err != nil {
		entry.Status = "unreadable"
		return entry, nil
	}
	if !info.Mode().IsRegular() {
		entry.Status = "non_regular"
		return entry, nil
	}
	currentHash, err := fileSHA256(path)
	if err != nil {
		entry.Status = "unreadable"
		return entry, nil
	}
	entry.CurrentHash = currentHash
	if currentHash == installedHash {
		entry.Status = "current"
	} else {
		entry.Status = "modified"
	}
	return entry, nil
}

func currentGlobalDesiredSet(catalog Catalog, manifest globalManifest) (map[string]bool, bool, string, error) {
	desired := map[string]bool{}
	if len(manifest.Packs) == 0 && len(manifest.Targets) == 0 && len(manifest.Files) == 0 {
		return desired, true, "", nil
	}
	if len(manifest.Packs) == 0 || len(manifest.Targets) == 0 {
		return desired, false, "manifest has files or selection state without both packs and targets", nil
	}
	packs, err := catalog.Resolve(manifest.Packs)
	if err != nil {
		return desired, false, err.Error(), nil
	}
	paths, err := globalDesiredPaths(packs, manifest.Targets)
	if err != nil {
		return nil, false, "", err
	}
	for _, path := range paths {
		desired[filepath.Clean(path)] = true
	}
	return desired, true, "", nil
}

func globalPathTargets(path string, installedTargets []string) ([]string, bool, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, false, err
	}
	rel, err := filepath.Rel(home, filepath.Clean(path))
	if err != nil || rel == ".." || strings.HasPrefix(rel, ".."+string(filepath.Separator)) {
		return nil, false, fmt.Errorf("global path is outside home: %s", path)
	}
	rel = filepath.ToSlash(rel)
	switch {
	case rel == ".agents/skills" || strings.HasPrefix(rel, ".agents/skills/"):
		return sortedUniqueStrings(installedTargets), true, nil
	case rel == ".copilot" || strings.HasPrefix(rel, ".copilot/"):
		return []string{TargetCopilot}, false, nil
	case rel == ".cursor" || strings.HasPrefix(rel, ".cursor/"):
		return []string{TargetCursor}, false, nil
	case rel == ".claude" || strings.HasPrefix(rel, ".claude/"):
		return []string{TargetClaudeCode}, false, nil
	case rel == ".config/opencode" || strings.HasPrefix(rel, ".config/opencode/"):
		return []string{TargetOpenCode}, false, nil
	case rel == ".codex" || strings.HasPrefix(rel, ".codex/"):
		return []string{TargetCodex}, false, nil
	default:
		return nil, false, fmt.Errorf("cannot associate managed global path with a target: %s", path)
	}
}

func globalPathContainsSymlink(path string) (bool, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return false, err
	}
	home = filepath.Clean(home)
	homeInfo, err := os.Lstat(home)
	if err != nil {
		return false, err
	}
	if homeInfo.Mode()&os.ModeSymlink != 0 {
		return true, nil
	}
	path = filepath.Clean(path)
	rel, err := filepath.Rel(home, path)
	if err != nil || rel == ".." || strings.HasPrefix(rel, ".."+string(filepath.Separator)) {
		return false, fmt.Errorf("path is outside home: %s", path)
	}
	return pathContainsSymlink(home, rel)
}

func collectAdoptableGlobalFiles(catalog Catalog, owned map[string]string, requestedTargets []string) ([]globalInventoryFile, error) {
	known := map[string]string{}
	for _, pack := range catalog.Packs {
		for _, entry := range pack.Files {
			baseDir, relPath, ok := globalDestination(entry)
			if !ok {
				continue
			}
			if err := addEmbeddedDestinationHashes(known, baseDir, relPath, entry.Source); err != nil {
				return nil, err
			}
		}
	}
	paths := make([]string, 0, len(known))
	for path := range known {
		paths = append(paths, path)
	}
	sort.Strings(paths)
	var out []globalInventoryFile
	for _, path := range paths {
		if _, managed := owned[filepath.Clean(path)]; managed {
			continue
		}
		targets, shared, err := globalPathTargets(path, nil)
		if err != nil {
			return nil, err
		}
		if len(requestedTargets) > 0 && !shared && !targetsIntersect(targets, requestedTargets) {
			continue
		}
		entry, err := inspectGlobalInventoryFile(path, known[path], targets, shared)
		if err != nil {
			return nil, err
		}
		if entry.Status != "current" {
			continue
		}
		entry.Ownership = "adoptable"
		entry.Status = "adoptable"
		entry.Proof = "exact-hash"
		out = append(out, entry)
	}
	return out, nil
}

func collectUnownedGlobalFiles(owned map[string]string, requestedTargets []string) ([]globalInventoryFile, error) {
	roots, err := globalInventoryManagedRoots()
	if err != nil {
		return nil, err
	}
	seen := map[string]bool{}
	var out []globalInventoryFile
	addPath := func(path string) error {
		path = filepath.Clean(path)
		if seen[path] {
			return nil
		}
		seen[path] = true
		if _, managed := owned[path]; managed {
			return nil
		}
		targets, shared, err := globalPathTargets(path, nil)
		if err != nil {
			return err
		}
		if len(requestedTargets) > 0 && !shared && !targetsIntersect(targets, requestedTargets) {
			return nil
		}
		entry := globalInventoryFile{
			Path:      path,
			Ownership: "unowned",
			Targets:   targets,
			Shared:    shared,
		}
		symlink, err := globalPathContainsSymlink(path)
		if err != nil {
			return err
		}
		if symlink {
			entry.Status = "symlink"
			out = append(out, entry)
			return nil
		}
		info, err := os.Lstat(path)
		if os.IsNotExist(err) {
			return nil
		}
		if err != nil {
			entry.Status = "unreadable"
			out = append(out, entry)
			return nil
		}
		if !info.Mode().IsRegular() {
			entry.Status = "non_regular"
			out = append(out, entry)
			return nil
		}
		hash, err := fileSHA256(path)
		if err != nil {
			entry.Status = "unreadable"
		} else {
			entry.Status = "unowned"
			entry.CurrentHash = hash
		}
		out = append(out, entry)
		return nil
	}

	for _, root := range roots {
		info, err := os.Lstat(root)
		if err != nil {
			if os.IsNotExist(err) {
				continue
			}
			return nil, err
		}
		blocked, err := globalPathContainsSymlink(root)
		if err != nil {
			return nil, err
		}
		if blocked {
			if err := addPath(root); err != nil {
				return nil, err
			}
			continue
		}
		if !info.IsDir() {
			if err := addPath(root); err != nil {
				return nil, err
			}
			continue
		}
		err = filepath.WalkDir(root, func(path string, d os.DirEntry, walkErr error) error {
			if walkErr != nil {
				return walkErr
			}
			if path == root || d.IsDir() {
				return nil
			}
			return addPath(path)
		})
		if err != nil {
			return nil, err
		}
	}
	sort.Slice(out, func(i, j int) bool { return out[i].Path < out[j].Path })
	return out, nil
}

func globalInventoryManagedRoots() ([]string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}
	return []string{
		filepath.Join(home, ".agents", "skills"),
		filepath.Join(home, ".copilot", "instructions"),
		filepath.Join(home, ".cursor", "agents"),
		filepath.Join(home, ".cursor", "rules"),
		filepath.Join(home, ".cursor", "hooks"),
		filepath.Join(home, ".cursor", "hooks.json"),
		filepath.Join(home, ".claude", "agents"),
		filepath.Join(home, ".claude", "CLAUDE.md"),
		filepath.Join(home, ".config", "opencode", "agents"),
		filepath.Join(home, ".config", "opencode", "AGENTS.md"),
		filepath.Join(home, ".codex", "AGENTS.md"),
		filepath.Join(home, ".codex", "instructions"),
	}, nil
}

func collectObservedGlobalRoots(requestedTargets []string) []globalObservedRoot {
	if len(requestedTargets) > 0 && !hasTarget(requestedTargets, TargetCodex) {
		return nil
	}
	home, err := os.UserHomeDir()
	if err != nil {
		return nil
	}
	candidates := []globalObservedRoot{
		{Path: filepath.Join(home, ".codex", "plugins", "cache"), Ownership: "external", Kind: "plugin-cache"},
		{Path: filepath.Join(home, ".codex", "skills", ".system"), Ownership: "external", Kind: "runtime-skills"},
	}
	var out []globalObservedRoot
	for _, item := range candidates {
		if _, err := os.Lstat(item.Path); err == nil {
			out = append(out, item)
		}
	}
	return out
}

func addGlobalInventorySummary(summary *globalInventorySummary, entry globalInventoryFile) {
	if entry.Ownership == "managed" {
		summary.Managed++
	} else if entry.Ownership == "unowned" {
		summary.Unowned++
	}
	switch entry.Status {
	case "current":
		summary.Current++
	case "modified":
		summary.Modified++
	case "missing":
		summary.Missing++
	case "symlink":
		summary.Symlink++
	case "non_regular":
		summary.NonRegular++
	case "unreadable":
		summary.Unreadable++
	case "adoptable":
		summary.Adoptable++
	case "obsolete":
		summary.Obsolete++
	}
	if entry.Obsolete && entry.Status != "obsolete" {
		summary.Obsolete++
	}
}

func printGlobalInventory(w io.Writer, report globalInventoryReport) {
	state := "missing"
	if report.Manifest.Exists {
		state = fmt.Sprintf("version=%d", report.Manifest.Version)
	}
	fmt.Fprintf(w, "global manifest: %s (%s)\n", report.Manifest.Path, state)
	fmt.Fprintf(w, "packs: %s\n", strings.Join(report.Manifest.Packs, ", "))
	fmt.Fprintf(w, "targets: %s\n", strings.Join(report.Manifest.Targets, ", "))
	if len(report.RequestedTargets) > 0 {
		fmt.Fprintf(w, "requested targets: %s\n", strings.Join(report.RequestedTargets, ", "))
	}
	if !report.CatalogResolved && report.CatalogError != "" {
		fmt.Fprintf(w, "catalog resolution: unavailable (%s)\n", report.CatalogError)
	}
	fmt.Fprintln(w)
	fmt.Fprintln(w, "STATUS       OWNERSHIP  OBSOLETE  TARGETS              INSTALLED HASH                                                    CURRENT HASH                                                      PATH")
	for _, file := range report.Files {
		targets := strings.Join(file.Targets, ",")
		if file.Shared {
			targets = "shared:" + targets
		}
		fmt.Fprintf(w, "%-12s %-10s %-9t %-20s %-64s %-64s %s\n",
			file.Status, file.Ownership, file.Obsolete, targets,
			file.InstalledHash, file.CurrentHash, file.Path)
	}
	fmt.Fprintf(w,
		"\nsummary: managed=%d unowned=%d current=%d modified=%d missing=%d symlink=%d non-regular=%d unreadable=%d adoptable=%d obsolete=%d\n",
		report.Summary.Managed, report.Summary.Unowned, report.Summary.Current, report.Summary.Modified,
		report.Summary.Missing, report.Summary.Symlink, report.Summary.NonRegular,
		report.Summary.Unreadable, report.Summary.Adoptable, report.Summary.Obsolete)
	for _, observed := range report.Observed {
		fmt.Fprintf(w, "observed external %s: %s\n", observed.Kind, observed.Path)
	}
}

func targetsIntersect(left, right []string) bool {
	for _, a := range left {
		for _, b := range right {
			if a == b {
				return true
			}
		}
	}
	return false
}

func sortedUniqueStrings(values []string) []string {
	out := append([]string(nil), values...)
	sort.Strings(out)
	return uniqueStrings(out)
}
