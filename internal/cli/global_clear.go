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

type globalClearEntry struct {
	Path          string   `json:"path"`
	Ownership     string   `json:"ownership"`
	Status        string   `json:"status"`
	Action        string   `json:"action"`
	Reason        string   `json:"reason"`
	Targets       []string `json:"targets,omitempty"`
	Shared        bool     `json:"shared"`
	InstalledHash string   `json:"installed_hash,omitempty"`
	CurrentHash   string   `json:"current_hash,omitempty"`
	Proof         string   `json:"proof,omitempty"`
}

type globalClearSummary struct {
	Remove   int `json:"remove"`
	Forget   int `json:"forget"`
	Preserve int `json:"preserve"`
}

type globalClearReport struct {
	SchemaVersion    int                     `json:"schema_version"`
	Command          string                  `json:"command"`
	DryRun           bool                    `json:"dry_run"`
	Force            bool                    `json:"force"`
	IncludeAdoptable bool                    `json:"include_adoptable"`
	Manifest         globalInventoryManifest `json:"manifest"`
	RequestedTargets []string                `json:"requested_targets,omitempty"`
	RemainingTargets []string                `json:"remaining_targets"`
	Summary          globalClearSummary      `json:"summary"`
	Entries          []globalClearEntry      `json:"entries"`
	nextManifest     globalManifest
	manifestExists   bool
	manifestDigest   string
}

func runGlobalClear(args []string, w io.Writer, catalog Catalog) error {
	fs := flag.NewFlagSet("global clear", flag.ContinueOnError)
	fs.SetOutput(io.Discard)
	dryRun := fs.Bool("dry-run", false, "report the clear plan without writing")
	force := fs.Bool("force", false, "remove modified regular files with proven ownership")
	includeAdoptable := fs.Bool("include-adoptable", false, "remove exact canonical or trusted generated legacy files")
	jsonOutput := fs.Bool("json", false, "write a stable JSON report")
	targetsFlag := fs.String("targets", "", "comma-separated target scope")
	if err := fs.Parse(args); err != nil {
		return err
	}
	if fs.NArg() != 0 {
		return fmt.Errorf("global clear does not accept pack names")
	}
	requestedTargets := splitCSV(*targetsFlag)
	if err := validateTargets(requestedTargets); err != nil {
		return err
	}
	if len(requestedTargets) > 0 {
		requestedTargets = normalizeTargets(requestedTargets)
	}

	var report globalClearReport
	var err error
	if *dryRun {
		report, err = planGlobalClear(catalog, requestedTargets, *includeAdoptable, *force)
		if err != nil {
			return err
		}
	} else {
		release, lockErr := acquireGlobalMutationLock()
		if lockErr != nil {
			return lockErr
		}
		defer release()
		report, err = planGlobalClear(catalog, requestedTargets, *includeAdoptable, *force)
		if err != nil {
			return err
		}
		if err := applyGlobalClear(&report); err != nil {
			return err
		}
	}
	report.DryRun = *dryRun
	if *jsonOutput {
		encoder := json.NewEncoder(w)
		encoder.SetEscapeHTML(false)
		encoder.SetIndent("", "  ")
		return encoder.Encode(report)
	}
	printGlobalClear(w, report)
	return nil
}

func planGlobalClear(catalog Catalog, requestedTargets []string, includeAdoptable, force bool) (globalClearReport, error) {
	manifestPath, err := globalManifestPath()
	if err != nil {
		return globalClearReport{}, err
	}
	manifest, manifestExists, manifestDigest, err := loadGlobalManifestSnapshot()
	if err != nil {
		return globalClearReport{}, err
	}
	inventory, err := collectGlobalInventoryFromManifest(catalog, nil, includeAdoptable, manifestPath, manifest, manifestExists)
	if err != nil {
		return globalClearReport{}, err
	}
	remainingTargets := subtractStrings(manifest.Targets, requestedTargets)
	if len(requestedTargets) == 0 {
		remainingTargets = nil
	}
	next := globalManifest{
		Version: globalManifestVersion,
		Packs:   append([]string(nil), manifest.Packs...),
		Targets: append([]string(nil), remainingTargets...),
		Files:   map[string]string{},
	}
	if len(remainingTargets) == 0 {
		next.Packs = nil
	}
	report := globalClearReport{
		SchemaVersion:    globalInventorySchemaVersion,
		Command:          "global.clear",
		Force:            force,
		IncludeAdoptable: includeAdoptable,
		Manifest:         inventory.Manifest,
		RequestedTargets: append([]string(nil), requestedTargets...),
		RemainingTargets: append([]string(nil), remainingTargets...),
		Entries:          []globalClearEntry{},
		nextManifest:     next,
		manifestExists:   manifestExists,
		manifestDigest:   manifestDigest,
	}

	for _, file := range inventory.Files {
		entry := globalClearEntry{
			Path:          file.Path,
			Ownership:     file.Ownership,
			Status:        file.Status,
			Targets:       append([]string(nil), file.Targets...),
			Shared:        file.Shared,
			InstalledHash: file.InstalledHash,
			CurrentHash:   file.CurrentHash,
			Proof:         file.Proof,
		}
		eligible := clearEntryEligible(file, manifest.Targets, requestedTargets, remainingTargets)
		switch {
		case file.Ownership == "managed" && !eligible:
			entry.Action, entry.Reason = "preserve", "outside requested target scope"
			report.nextManifest.Files[file.Path] = file.InstalledHash
		case file.Ownership == "managed" && file.Status == "missing":
			entry.Action, entry.Reason = "forget", "owned path is already missing"
		case file.Ownership == "managed" && (file.Status == "current" || file.Status == "obsolete"):
			entry.Action, entry.Reason = "remove", "owned regular file matches its installed hash"
		case file.Ownership == "managed" && file.Status == "modified" && force:
			entry.Action, entry.Reason = "remove", "force enabled for a proven owned regular file"
		case file.Ownership == "managed":
			entry.Action, entry.Reason = "preserve", clearPreserveReason(file.Status)
			report.nextManifest.Files[file.Path] = file.InstalledHash
		case file.Ownership == "adoptable" && eligible && file.Proof == "generated-marker" && !force:
			entry.Action, entry.Reason = "preserve", "generated marker alone requires --force to remove customized legacy content"
		case file.Ownership == "adoptable" && eligible:
			entry.Action, entry.Reason = "remove", "explicitly included exact canonical or trusted generated legacy file"
		default:
			entry.Action, entry.Reason = "preserve", "ownership is not proven for this clear scope"
		}
		addGlobalClearSummary(&report.Summary, entry.Action)
		report.Entries = append(report.Entries, entry)
	}
	sort.Slice(report.Entries, func(i, j int) bool { return report.Entries[i].Path < report.Entries[j].Path })
	return report, nil
}

func clearEntryEligible(file globalInventoryFile, installedTargets, requestedTargets, remainingTargets []string) bool {
	if file.Ownership == "adoptable" && file.Shared {
		return len(requestedTargets) == 0
	}
	if file.Shared {
		return len(remainingTargets) == 0 && (len(requestedTargets) == 0 || targetsIntersect(installedTargets, requestedTargets))
	}
	if len(requestedTargets) == 0 {
		return true
	}
	return targetsIntersect(file.Targets, requestedTargets)
}

func applyGlobalClear(report *globalClearReport) error {
	_, manifestExists, manifestDigest, err := loadGlobalManifestSnapshot()
	if err != nil {
		return err
	}
	if manifestExists != report.manifestExists || manifestDigest != report.manifestDigest {
		return fmt.Errorf("global manifest changed after planning; rerun global clear")
	}
	for i := range report.Entries {
		entry := &report.Entries[i]
		if entry.Action != "remove" {
			continue
		}
		blocked, err := globalPathContainsSymlink(entry.Path)
		if err != nil {
			return err
		}
		info, err := os.Lstat(entry.Path)
		if os.IsNotExist(err) {
			entry.Action, entry.Reason = "forget", "path became missing after planning"
			recountGlobalClearSummary(report)
			continue
		}
		if err != nil || blocked || !info.Mode().IsRegular() {
			preserveGlobalClearEntry(report, entry, "path became unsafe or non-regular after planning")
			continue
		}
		currentHash, err := fileSHA256(entry.Path)
		if err != nil {
			preserveGlobalClearEntry(report, entry, "path became unreadable after planning")
			continue
		}
		if currentHash != entry.CurrentHash {
			preserveGlobalClearEntry(report, entry, "content changed after planning")
			continue
		}
		if err := os.Remove(entry.Path); err != nil {
			return err
		}
		removeEmptyManagedParents(filepath.Dir(entry.Path))
	}
	if report.manifestExists {
		report.nextManifest.Packs = sortedUniqueStrings(report.nextManifest.Packs)
		report.nextManifest.Targets = sortedUniqueStrings(report.nextManifest.Targets)
		if err := writeGlobalManifest(report.nextManifest); err != nil {
			return err
		}
	}
	return nil
}

func preserveGlobalClearEntry(report *globalClearReport, entry *globalClearEntry, reason string) {
	entry.Action = "preserve"
	entry.Reason = reason
	if entry.Ownership == "managed" {
		report.nextManifest.Files[entry.Path] = entry.InstalledHash
	}
	recountGlobalClearSummary(report)
}

func clearPreserveReason(status string) string {
	switch status {
	case "modified":
		return "owned file was modified; use --force to remove it"
	case "symlink":
		return "symlinks are never removed"
	case "non_regular":
		return "non-regular paths are never removed"
	case "unreadable":
		return "unreadable paths are preserved"
	default:
		return "path is preserved conservatively"
	}
}

func addGlobalClearSummary(summary *globalClearSummary, action string) {
	switch action {
	case "remove":
		summary.Remove++
	case "forget":
		summary.Forget++
	case "preserve":
		summary.Preserve++
	}
}

func recountGlobalClearSummary(report *globalClearReport) {
	report.Summary = globalClearSummary{}
	for _, entry := range report.Entries {
		addGlobalClearSummary(&report.Summary, entry.Action)
	}
}

func subtractStrings(installed, removed []string) []string {
	if len(removed) == 0 {
		return nil
	}
	removeSet := map[string]bool{}
	for _, value := range removed {
		removeSet[value] = true
	}
	var out []string
	for _, value := range installed {
		if !removeSet[value] {
			out = append(out, value)
		}
	}
	return sortedUniqueStrings(out)
}

func printGlobalClear(w io.Writer, report globalClearReport) {
	mode := "apply"
	if report.DryRun {
		mode = "dry-run"
	}
	fmt.Fprintf(w, "global clear: %s\n", mode)
	if len(report.RequestedTargets) > 0 {
		fmt.Fprintf(w, "requested targets: %s\n", strings.Join(report.RequestedTargets, ", "))
	}
	fmt.Fprintf(w, "remaining targets: %s\n\n", strings.Join(report.RemainingTargets, ", "))
	fmt.Fprintln(w, "ACTION    STATUS       OWNERSHIP  REASON  PATH")
	for _, entry := range report.Entries {
		fmt.Fprintf(w, "%-9s %-12s %-10s %s  %s\n", entry.Action, entry.Status, entry.Ownership, entry.Reason, entry.Path)
	}
	fmt.Fprintf(w, "\nsummary: remove=%d forget=%d preserve=%d\n", report.Summary.Remove, report.Summary.Forget, report.Summary.Preserve)
}
