package dashboard

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

var errSkipSpec = errors.New("skip spec")

var (
	specDirPattern = regexp.MustCompile(`^(\d{4})-([a-z0-9-]+)$`)
	taskCheckbox   = regexp.MustCompile(`(?m)^\s*-\s*\[([ x~])\]`)
	dateFilename   = regexp.MustCompile(`^(\d{4}-\d{2}-\d{2})-(.+)\.md$`)
)

// ParseDateFromFilename extracts YYYY-MM-DD from docs filenames.
func ParseDateFromFilename(name string) (date, slug string, ok bool) {
	m := dateFilename.FindStringSubmatch(name)
	if len(m) != 3 {
		return "", "", false
	}
	return m[1], m[2], true
}

// ParseTaskStats counts task checkboxes in markdown.
func ParseTaskStats(content string) TaskStats {
	var stats TaskStats
	for _, m := range taskCheckbox.FindAllStringSubmatch(content, -1) {
		stats.Total++
		switch m[1] {
		case "x":
			stats.Done++
		case "~":
			stats.InProgress++
		default:
			stats.Pending++
		}
	}
	return stats
}

// SpecRoots are relative paths scanned for NNNN-slug spec folders.
var SpecRoots = []string{"specs", "docs/specs"}

// ScanSpecs walks SpecRoots under root and returns grouped dashboard data.
func ScanSpecs(root string) (DashboardData, error) {
	data := DashboardData{
		Dates: make(map[int]SpecDates),
		Links: LinkIndex{
			SpecsByNumber: make(map[int]*SpecSummary),
			PRDByPath:     make(map[string]*PRD),
			OriginToSpecs: make(map[string][]int),
			Dependents:    make(map[int][]int),
		},
	}
	seen := map[int]string{}
	var specs []SpecSummary

	for _, rel := range SpecRoots {
		specsDir := filepath.Join(root, rel)
		entries, err := os.ReadDir(specsDir)
		if os.IsNotExist(err) {
			continue
		}
		if err != nil {
			return DashboardData{}, err
		}
		for _, ent := range entries {
			if !ent.IsDir() {
				continue
			}
			m := specDirPattern.FindStringSubmatch(ent.Name())
			if m == nil {
				data.Warnings = append(data.Warnings, Warning{
					Path:    filepath.Join(specsDir, ent.Name()),
					Message: "folder name does not match NNNN-slug pattern",
				})
				continue
			}
			num, _ := strconv.Atoi(m[1])
			dir := filepath.Join(specsDir, ent.Name())
			if prev, ok := seen[num]; ok {
				data.Warnings = append(data.Warnings, Warning{
					Path: dir,
					Message: fmt.Sprintf("duplicate spec number %04d (already in %s)", num, prev),
				})
				continue
			}
			seen[num] = rel

			spec, warn, err := parseSpecDir(dir, num, m[2])
			if err != nil {
				if warn != nil {
					data.Warnings = append(data.Warnings, *warn)
				}
				if !errors.Is(err, errSkipSpec) {
					data.Warnings = append(data.Warnings, Warning{Path: dir, Message: err.Error()})
				}
				continue
			}
			if warn != nil {
				data.Warnings = append(data.Warnings, *warn)
			}
			specs = append(specs, spec)
		}
	}

	sort.Slice(specs, func(i, j int) bool { return specs[i].Number < specs[j].Number })
	data.Families = groupSpecsByFamily(specs)
	buildLinkIndex(&data, specs)
	return data, nil
}

func parseSpecDir(dir string, num int, slug string) (SpecSummary, *Warning, error) {
	specPath := filepath.Join(dir, "spec.md")
	raw, err := os.ReadFile(specPath)
	if err != nil {
		return SpecSummary{}, nil, fmt.Errorf("read spec.md: %w", err)
	}
	if len(raw) == 0 {
		return SpecSummary{}, &Warning{Path: specPath, Message: "empty spec.md"}, errSkipSpec
	}

	fm, body, err := parseFrontmatter(string(raw))
	if err != nil {
		return SpecSummary{}, nil, err
	}

	spec := SpecSummary{
		Number: num,
		Slug:   slug,
		Body:   body,
		Dir:    dir,
		Status: "Draft",
	}

	if len(fm) == 0 {
		spec.Family = ungroupedFamily
		spec.Title = slug
		return spec, &Warning{Path: specPath, Message: "missing frontmatter"}, nil
	}

	spec.Title = fm["title"]
	if spec.Title == "" {
		spec.Title = slug
	}
	spec.Family = fm["family"]
	if spec.Family == "" {
		spec.Family = ungroupedFamily
	}
	if p, err := strconv.Atoi(strings.TrimSpace(fm["phase"])); err == nil {
		spec.Phase = p
	}
	if s := strings.TrimSpace(fm["status"]); s != "" {
		spec.Status = s
	}
	spec.Owner = fm["owner"]
	spec.Origin = normalizePath(fm["origin"])
	spec.DependsOn = parseDependsOn(fm["depends_on"])
	spec.Deprecated = strings.EqualFold(spec.Status, "Deprecated")

	if planRaw, err := os.ReadFile(filepath.Join(dir, "plan.md")); err == nil {
		_, planBody, _ := parseFrontmatter(string(planRaw))
		spec.HasPlan = true
		spec.PlanBody = planBody
	}

	if tasksRaw, err := os.ReadFile(filepath.Join(dir, "tasks.md")); err == nil {
		spec.TasksBody = string(tasksRaw)
		spec.Tasks = ParseTaskStats(string(tasksRaw))
	}

	return spec, nil, nil
}

func groupSpecsByFamily(specs []SpecSummary) []Family {
	byFamily := map[string]map[int][]SpecSummary{}
	order := []string{}
	for _, spec := range specs {
		if _, ok := byFamily[spec.Family]; !ok {
			byFamily[spec.Family] = map[int][]SpecSummary{}
			order = append(order, spec.Family)
		}
		byFamily[spec.Family][spec.Phase] = append(byFamily[spec.Family][spec.Phase], spec)
	}
	sort.Strings(order)

	var families []Family
	for _, name := range order {
		phasesMap := byFamily[name]
		var phaseNums []int
		for p := range phasesMap {
			phaseNums = append(phaseNums, p)
		}
		sort.Ints(phaseNums)
		var phases []Phase
		for _, p := range phaseNums {
			list := phasesMap[p]
			sort.Slice(list, func(i, j int) bool { return list[i].Number < list[j].Number })
			phases = append(phases, Phase{Number: p, Specs: list})
		}
		families = append(families, Family{Name: name, Phases: phases})
	}
	return families
}

func buildLinkIndex(data *DashboardData, specs []SpecSummary) {
	for i := range specs {
		s := specs[i]
		data.Links.SpecsByNumber[s.Number] = &specs[i]
		if s.Origin != "" {
			data.Links.OriginToSpecs[s.Origin] = append(data.Links.OriginToSpecs[s.Origin], s.Number)
		}
		for _, dep := range s.DependsOn {
			data.Links.Dependents[dep] = append(data.Links.Dependents[dep], s.Number)
		}
	}
}

// ScanPRDs scans docs/prds/*.md.
func ScanPRDs(root string) ([]PRD, []Warning, error) {
	return scanDocs(root, "docs/prds")
}

// ScanBugfixes scans docs/bugfixes/*.md.
func ScanBugfixes(root string) ([]Bugfix, []Warning, error) {
	bugfixes, warns, err := scanDocs(root, "docs/bugfixes")
	if err != nil {
		return nil, warns, err
	}
	out := make([]Bugfix, len(bugfixes))
	for i, p := range bugfixes {
		out[i] = Bugfix(p)
	}
	return out, warns, nil
}

// ScanPostmortems scans docs/postmortems/*.md.
func ScanPostmortems(root string) ([]Postmortem, []Warning, error) {
	docs, warns, err := scanDocs(root, "docs/postmortems")
	if err != nil {
		return nil, warns, err
	}
	out := make([]Postmortem, len(docs))
	for i, p := range docs {
		out[i] = Postmortem(p)
	}
	return out, warns, nil
}

func scanDocs(root, sub string) ([]PRD, []Warning, error) {
	dir := filepath.Join(root, sub)
	entries, err := os.ReadDir(dir)
	if os.IsNotExist(err) {
		return nil, nil, nil
	}
	if err != nil {
		return nil, nil, err
	}

	var prds []PRD
	var warnings []Warning
	for _, ent := range entries {
		if ent.IsDir() || !strings.HasSuffix(ent.Name(), ".md") {
			continue
		}
		path := filepath.Join(dir, ent.Name())
		raw, err := os.ReadFile(path)
		if err != nil {
			warnings = append(warnings, Warning{Path: path, Message: err.Error()})
			continue
		}
		fm, body, _ := parseFrontmatter(string(raw))
		date, slug, ok := ParseDateFromFilename(ent.Name())
		title := fm["title"]
		if title == "" {
			title = titleFromBody(body, slug)
		}
		if d := fm["date"]; d != "" {
			date = d
		} else if !ok {
			warnings = append(warnings, Warning{Path: path, Message: "cannot extract date from filename"})
			continue
		}
		rel := filepath.ToSlash(filepath.Join(sub, ent.Name()))
		prds = append(prds, PRD{
			Slug:  slug,
			Title: title,
			Date:  date,
			Path:  rel,
			Body:  body,
		})
	}
	sort.Slice(prds, func(i, j int) bool { return prds[i].Date > prds[j].Date })
	return prds, warnings, nil
}

func titleFromBody(body, slug string) string {
	for _, line := range strings.Split(body, "\n") {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "# ") {
			return strings.TrimPrefix(line, "# ")
		}
	}
	return strings.ReplaceAll(slug, "-", " ")
}

func normalizePath(p string) string {
	p = strings.TrimSpace(p)
	p = strings.Trim(p, `"'`)
	return filepath.ToSlash(p)
}

// LoadAll parses specs, PRDs, bugfixes, and postmortems under root.
func LoadAll(root string) (DashboardData, error) {
	data, err := ScanSpecs(root)
	if err != nil {
		return data, err
	}
	prds, w1, err := ScanPRDs(root)
	if err != nil {
		return data, err
	}
	data.PRDs = prds
	data.Warnings = append(data.Warnings, w1...)
	for i := range data.PRDs {
		data.Links.PRDByPath[data.PRDs[i].Path] = &data.PRDs[i]
	}
	bugfixes, w2, err := ScanBugfixes(root)
	if err != nil {
		return data, err
	}
	data.Bugfixes = bugfixes
	data.Warnings = append(data.Warnings, w2...)
	postmortems, w3, err := ScanPostmortems(root)
	if err != nil {
		return data, err
	}
	data.Postmortems = postmortems
	data.Warnings = append(data.Warnings, w3...)
	return data, nil
}
