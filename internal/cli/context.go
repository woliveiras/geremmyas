package cli

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strings"

	geremmyas "github.com/woliveiras/geremmyas"
)

type contextStats struct {
	TopLevel         int `json:"top_level"`
	Nested           int `json:"nested"`
	FrontmatterBytes int `json:"frontmatter_bytes"`
	Managed          int `json:"managed"`
	Modified         int `json:"modified"`
	Unowned          int `json:"unowned"`
}

type contextSelection struct {
	Exists  bool     `json:"exists"`
	Packs   []string `json:"packs,omitempty"`
	Targets []string `json:"targets,omitempty"`
}

type contextSource struct {
	Name         string       `json:"name"`
	Path         string       `json:"path"`
	Stats        contextStats `json:"stats"`
	ApproxTokens int          `json:"approx_tokens"`
	Note         string       `json:"note"`
}

type contextContract struct {
	Name         string `json:"name"`
	Path         string `json:"path"`
	Exists       bool   `json:"exists"`
	Words        int    `json:"words,omitempty"`
	Bytes        int    `json:"bytes,omitempty"`
	ApproxTokens int    `json:"approx_tokens,omitempty"`
	Error        string `json:"error,omitempty"`
}

type contextSkillCost struct {
	Name             string   `json:"name"`
	SelectedBy       []string `json:"selected_by"`
	FrontmatterBytes int      `json:"frontmatter_bytes"`
	BodyBytes        int      `json:"body_bytes"`
	SupportBytes     int      `json:"support_bytes"`
	DiscoveryTokens  int      `json:"discovery_tokens"`
	BodyTokens       int      `json:"body_tokens"`
	SupportTokens    int      `json:"support_tokens"`
}

type contextPackCost struct {
	Name            string `json:"name"`
	Skills          int    `json:"skills"`
	DiscoveryTokens int    `json:"discovery_tokens"`
	BodyTokens      int    `json:"body_tokens"`
	SupportTokens   int    `json:"support_tokens"`
}

type contextReport struct {
	SchemaVersion int                `json:"schema_version"`
	Command       string             `json:"command"`
	Root          string             `json:"root"`
	Project       contextSelection   `json:"project"`
	Global        contextSelection   `json:"global"`
	Sources       []contextSource    `json:"sources"`
	Contracts     []contextContract  `json:"contracts"`
	SkillCosts    []contextSkillCost `json:"skill_costs"`
	PackCosts     []contextPackCost  `json:"pack_costs"`
}

func (s contextStats) total() int { return s.TopLevel + s.Nested }

func approximateTokens(bytes int) int {
	return (bytes + 3) / 4
}

func runContext(args []string, w io.Writer, catalog Catalog) error {
	flags := flag.NewFlagSet("context", flag.ContinueOnError)
	flags.SetOutput(io.Discard)
	rootFlag := flags.String("root", "", "project root to inspect")
	jsonOutput := flags.Bool("json", false, "print machine-readable JSON")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return fmt.Errorf("context does not accept positional arguments")
	}
	report, err := collectContextReport(*rootFlag, catalog)
	if err != nil {
		return err
	}
	if *jsonOutput {
		encoder := json.NewEncoder(w)
		encoder.SetIndent("", "  ")
		return encoder.Encode(report)
	}
	printContextReport(w, report)
	return nil
}

func collectContextReport(rootFlag string, catalog Catalog) (contextReport, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return contextReport{}, err
	}
	globalManifest, globalExists, err := loadGlobalManifest()
	if err != nil {
		return contextReport{}, err
	}
	root := rootFlag
	if root == "" {
		root, err = os.Getwd()
		if err != nil {
			return contextReport{}, err
		}
	}
	root, err = filepath.Abs(root)
	if err != nil {
		return contextReport{}, err
	}
	projectManifest, projectExists, err := loadProjectManifest(root)
	if err != nil {
		return contextReport{}, err
	}
	report := contextReport{
		SchemaVersion: 1,
		Command:       "context",
		Root:          root,
		Project:       contextSelection{Exists: projectExists, Packs: projectManifest.Packs, Targets: projectManifest.Targets},
		Global:        contextSelection{Exists: globalExists, Packs: globalManifest.Packs, Targets: globalManifest.Targets},
	}

	catalogStats, err := collectEmbeddedContextStats("content/skills")
	if err != nil {
		return contextReport{}, err
	}
	sources := []struct {
		name     string
		path     string
		manifest map[string]string
		note     string
	}{
		{name: "catalog", path: "content/skills", note: "managed by geremmyas source"},
		{name: "project-copilot", path: filepath.Join(root, ".github", "skills"), manifest: projectAbsoluteHashes(root, projectManifest.Files), note: "ownership from project manifest"},
		{name: "project-portable", path: filepath.Join(root, ".agents", "skills"), manifest: projectAbsoluteHashes(root, projectManifest.Files), note: "ownership from project manifest"},
		{name: "global-portable", path: filepath.Join(home, ".agents", "skills"), manifest: globalManifest.Files, note: "ownership from global manifest"},
		{name: "codex-system", path: filepath.Join(home, ".codex", "skills", ".system"), note: "observed, not managed"},
		{name: "codex-plugin-cache", path: filepath.Join(home, ".codex", "plugins", "cache"), note: "observed cache upper bound; host activation may be smaller"},
	}
	for index, source := range sources {
		stats := catalogStats
		if index != 0 {
			stats, err = collectFilesystemContextStats(source.path, source.manifest)
		}
		if err != nil {
			return contextReport{}, err
		}
		report.Sources = append(report.Sources, contextSource{Name: source.name, Path: source.path, Stats: stats, ApproxTokens: approximateTokens(stats.FrontmatterBytes), Note: source.note})
	}
	report.Contracts = []contextContract{
		inspectContract("project", filepath.Join(root, "AGENTS.md")),
		inspectContract("codex-global", filepath.Join(home, ".codex", "AGENTS.md")),
	}
	report.SkillCosts, report.PackCosts, err = collectEmbeddedSkillCosts(catalog)
	if err != nil {
		return contextReport{}, err
	}
	return report, nil
}

func projectAbsoluteHashes(root string, files map[string]string) map[string]string {
	out := make(map[string]string, len(files))
	for rel, hash := range files {
		out[filepath.Join(root, filepath.FromSlash(rel))] = hash
	}
	return out
}

func printContextReport(w io.Writer, report contextReport) {
	fmt.Fprintln(w, "Context usage (approximate; tokens = bytes / 4)")
	fmt.Fprintf(w, "state: project=%s global=%s\n", formatContextSelection(report.Project), formatContextSelection(report.Global))
	for _, source := range report.Sources {
		printContextStats(w, source.Name, source.Stats, source.Note)
	}
	fmt.Fprintln(w, "workflow pack upper bounds (body = all selected top-level skill bodies; support = loaded only on demand):")
	for _, cost := range report.PackCosts {
		fmt.Fprintf(w, "  %s: skills=%d discovery-tokens=%d body-tokens=%d support-tokens=%d\n", cost.Name, cost.Skills, cost.DiscoveryTokens, cost.BodyTokens, cost.SupportTokens)
	}
	fmt.Fprintln(w, "contracts:")
	for _, contract := range report.Contracts {
		printContract(w, contract)
	}
}

func collectEmbeddedSkillCosts(catalog Catalog) ([]contextSkillCost, []contextPackCost, error) {
	selectedBy := map[string]map[string]bool{}
	for _, selector := range catalog.Packs {
		resolved, err := catalog.Resolve([]string{selector.Name})
		if err != nil {
			return nil, nil, err
		}
		for _, pack := range resolved {
			for _, entry := range pack.Files {
				if entry.Kind != ArtifactSkill {
					continue
				}
				name := filepath.Base(filepath.Clean(entry.Path))
				if selectedBy[name] == nil {
					selectedBy[name] = map[string]bool{}
				}
				selectedBy[name][selector.Name] = true
			}
		}
	}

	var costs []contextSkillCost
	for name, selectors := range selectedBy {
		root := filepath.ToSlash(filepath.Join("content/skills", name))
		skillData, err := fs.ReadFile(geremmyas.EmbeddedFiles, root+"/SKILL.md")
		if err != nil {
			return nil, nil, err
		}
		frontmatter := frontmatterBytes(skillData)
		body := markdownBodyBytes(skillData)
		support := 0
		err = fs.WalkDir(geremmyas.EmbeddedFiles, root, func(path string, d fs.DirEntry, walkErr error) error {
			if walkErr != nil || d.IsDir() || path == root+"/SKILL.md" {
				return walkErr
			}
			data, readErr := fs.ReadFile(geremmyas.EmbeddedFiles, path)
			if readErr != nil {
				return readErr
			}
			support += len(data)
			return nil
		})
		if err != nil {
			return nil, nil, err
		}
		selectorNames := make([]string, 0, len(selectors))
		for selector := range selectors {
			selectorNames = append(selectorNames, selector)
		}
		sort.Strings(selectorNames)
		costs = append(costs, contextSkillCost{
			Name: name, SelectedBy: selectorNames, FrontmatterBytes: frontmatter,
			BodyBytes: body, SupportBytes: support,
			DiscoveryTokens: approximateTokens(frontmatter), BodyTokens: approximateTokens(body), SupportTokens: approximateTokens(support),
		})
	}
	sort.Slice(costs, func(i, j int) bool { return costs[i].Name < costs[j].Name })

	var packs []contextPackCost
	for _, name := range []string{"coding", "quality", "base"} {
		cost := contextPackCost{Name: name}
		for _, skill := range costs {
			if !containsString(skill.SelectedBy, name) {
				continue
			}
			cost.Skills++
			cost.DiscoveryTokens += skill.DiscoveryTokens
			cost.BodyTokens += skill.BodyTokens
			cost.SupportTokens += skill.SupportTokens
		}
		packs = append(packs, cost)
	}
	return costs, packs, nil
}

func containsString(values []string, want string) bool {
	for _, value := range values {
		if value == want {
			return true
		}
	}
	return false
}

func formatContextSelection(selection contextSelection) string {
	if !selection.Exists {
		return "none"
	}
	return fmt.Sprintf("packs=%s targets=%s", strings.Join(selection.Packs, ","), strings.Join(selection.Targets, ","))
}

func printContextStats(w io.Writer, name string, stats contextStats, note string) {
	fmt.Fprintf(w,
		"  %s: top-level=%d nested=%d total=%d frontmatter-bytes=%d approx-tokens=%d managed=%d modified=%d unowned=%d (%s)\n",
		name, stats.TopLevel, stats.Nested, stats.total(), stats.FrontmatterBytes,
		approximateTokens(stats.FrontmatterBytes), stats.Managed, stats.Modified,
		stats.Unowned, note)
}

func inspectContract(name, path string) contextContract {
	contract := contextContract{Name: name, Path: path}
	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return contract
		}
		contract.Error = err.Error()
		return contract
	}
	contract.Exists = true
	contract.Words = len(strings.Fields(string(data)))
	contract.Bytes = len(data)
	contract.ApproxTokens = approximateTokens(len(data))
	return contract
}

func printContract(w io.Writer, contract contextContract) {
	if !contract.Exists {
		if contract.Error != "" {
			fmt.Fprintf(w, "  %s: unreadable (%s)\n", contract.Name, contract.Error)
		} else {
			fmt.Fprintf(w, "  %s: missing\n", contract.Name)
		}
		return
	}
	fmt.Fprintf(w, "  %s: words=%d bytes=%d approx-tokens=%d\n", contract.Name, contract.Words, contract.Bytes, contract.ApproxTokens)
}

func collectFilesystemContextStats(root string, owned map[string]string) (contextStats, error) {
	stats := contextStats{}
	info, err := os.Lstat(root)
	if os.IsNotExist(err) {
		return stats, nil
	}
	if err != nil {
		return stats, err
	}
	if info.Mode()&os.ModeSymlink != 0 {
		return stats, nil
	}
	err = filepath.WalkDir(root, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.Type()&os.ModeSymlink != 0 {
			if d.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}
		if d.IsDir() || d.Name() != "SKILL.md" {
			return nil
		}
		rel, err := filepath.Rel(root, path)
		if err != nil {
			return err
		}
		if len(strings.Split(filepath.ToSlash(rel), "/")) == 2 {
			stats.TopLevel++
		} else {
			stats.Nested++
		}
		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		stats.FrontmatterBytes += frontmatterBytes(data)
		installedHash, isOwned := owned[filepath.Clean(path)]
		if isOwned {
			stats.Managed++
			if bytesSHA256(data) != installedHash {
				stats.Modified++
			}
		} else {
			stats.Unowned++
		}
		return nil
	})
	return stats, err
}

func collectEmbeddedContextStats(root string) (contextStats, error) {
	stats := contextStats{}
	err := fs.WalkDir(geremmyas.EmbeddedFiles, root, func(path string, d fs.DirEntry, err error) error {
		if err != nil || d.IsDir() || d.Name() != "SKILL.md" {
			return err
		}
		rel := strings.TrimPrefix(path, root+"/")
		if len(strings.Split(rel, "/")) == 2 {
			stats.TopLevel++
		} else {
			stats.Nested++
		}
		data, err := fs.ReadFile(geremmyas.EmbeddedFiles, path)
		if err != nil {
			return err
		}
		stats.FrontmatterBytes += frontmatterBytes(data)
		stats.Managed++
		return nil
	})
	return stats, err
}

func frontmatterBytes(data []byte) int {
	text := string(data)
	if !strings.HasPrefix(text, "---\n") {
		return 0
	}
	end := strings.Index(text[4:], "\n---")
	if end < 0 {
		return 0
	}
	return end
}

func markdownBodyBytes(data []byte) int {
	text := string(data)
	if !strings.HasPrefix(text, "---\n") {
		return len(data)
	}
	end := strings.Index(text[4:], "\n---")
	if end < 0 {
		return len(data)
	}
	start := 4 + end + len("\n---")
	for start < len(text) && (text[start] == '\n' || text[start] == '\r') {
		start++
	}
	return len(text) - start
}
