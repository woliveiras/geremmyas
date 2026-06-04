package dashboard

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

// RenderOptions configures dashboard generation.
type RenderOptions struct {
	Root string
}

type pageBase struct {
	Title       string
	ActiveNav   string
	HasGit      bool
	Summary     globalSummary
	RootRel     string
}

type globalSummary struct {
	Total        int
	Implemented  int
	InProgress   int
	Pending      int
}

type indexPage struct {
	pageBase
	Families []familyCard
}

type familyCard struct {
	Name        string
	Slug        string
	SpecCount   int
	Implemented int
	InProgress  int
	Pending     int
	ProgressPct float64
	ActivePhase int
}

type familyPage struct {
	pageBase
	Family       Family
	Slug         string
	BoardColumns []BoardColumn
	Blocked      []blockedRow
}

type blockedRow struct {
	Spec   SpecSummary
	Reason string
}

type specPage struct {
	pageBase
	Spec       SpecSummary
	BodyHTML   template.HTML
	PlanHTML   template.HTML
	TasksHTML  template.HTML
	Dates      SpecDates
	BlockedBy  []SpecSummary
	Blocks     []SpecSummary
	OriginHref string
	OriginOK   bool
	Deps       []depLink
}

type depLink struct {
	Number int
	Title  string
	Href   string
	Missing bool
}

type metricsPage struct {
	pageBase
	Metrics      Metrics
	VelLabels    []string
	VelValues    []int
	NoGitMessage bool
}

type listPage struct {
	pageBase
	Items []listItem
	Empty string
}

type listItem struct {
	Title string
	Date  string
	Href  string
	Count int
}

type docPage struct {
	pageBase
	Title    string
	Date     string
	BodyHTML template.HTML
	Related  []SpecSummary
}

// RenderDashboard writes the static site to outputDir.
func RenderDashboard(data DashboardData, outputDir string) error {
	if err := os.MkdirAll(outputDir, 0o755); err != nil {
		return err
	}
	tmpl, err := template.New("layout").Funcs(templateFuncs(data)).ParseFS(Assets,
		"dashboard_assets/templates/*.html", "dashboard_assets/templates/partials.html")
	if err != nil {
		return err
	}
	if err := copyStaticAssets(outputDir); err != nil {
		return err
	}

	summary := computeSummary(data)
	base := pageBase{HasGit: data.Metrics.GitAvailable, Summary: summary}

	if err := renderTemplate(tmpl, filepath.Join(outputDir, "index.html"), "index.html", indexPage{
		pageBase: baseWithNav(base, "overview", "Overview"),
		Families: familyCards(data),
	}); err != nil {
		return err
	}

	for _, fam := range data.Families {
		slug := slugify(fam.Name)
		dir := filepath.Join(outputDir, "families", slug)
		if err := os.MkdirAll(dir, 0o755); err != nil {
			return err
		}
		specs := FamilySpecs(fam)
		var blocked []blockedRow
		for _, s := range specs {
			for _, b := range BlockedBy(s, data.Links) {
				blocked = append(blocked, blockedRow{Spec: s, Reason: fmt.Sprintf("blocked by %04d %s", b.Number, b.Title)})
			}
		}
		if err := renderTemplate(tmpl, filepath.Join(dir, "index.html"), "family.html", familyPage{
			pageBase:     baseWithNav(base, "families", fam.Name),
			Family:       fam,
			Slug:         slug,
			BoardColumns: BoardColumns(specs),
			Blocked:      blocked,
		}); err != nil {
			return err
		}
	}

	for n, spec := range data.Links.SpecsByNumber {
		if err := renderSpecPage(tmpl, outputDir, base, data, *spec); err != nil {
			return fmt.Errorf("spec %04d: %w", n, err)
		}
	}

	labels, values := ChartJSON(data.Metrics)
	if err := renderTemplate(tmpl, filepath.Join(outputDir, "metrics.html"), "metrics.html", metricsPage{
		pageBase:     baseWithNav(base, "metrics", "Metrics"),
		Metrics:      data.Metrics,
		VelLabels:    labels,
		VelValues:    values,
		NoGitMessage: !data.Metrics.GitAvailable,
	}); err != nil {
		return err
	}

	if err := renderPRDPages(tmpl, outputDir, base, data); err != nil {
		return err
	}
	if err := renderBugfixPages(tmpl, outputDir, base, data); err != nil {
		return err
	}
	return renderListPages(tmpl, outputDir, base, data)
}

func renderSpecPage(tmpl *template.Template, outputDir string, base pageBase, data DashboardData, spec SpecSummary) error {
	body, _ := renderMarkdown(spec.Body)
	plan, _ := renderMarkdown(spec.PlanBody)
	tasks, _ := renderMarkdown(spec.TasksBody)
	dir := filepath.Join(outputDir, "specs", fmt.Sprintf("%04d", spec.Number))
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return err
	}
	var deps []depLink
	for _, d := range spec.DependsOn {
		link := depLink{Number: d, Href: fmt.Sprintf("/specs/%04d/index.html", d)}
		if dep, ok := data.Links.SpecsByNumber[d]; ok {
			link.Title = dep.Title
		} else {
			link.Missing = true
			link.Title = "(missing)"
		}
		deps = append(deps, link)
	}
	href, ok := PRDLink(spec.Origin, data.Links)
	if spec.HasPlan {
		_ = renderMarkdownToFile(plan, filepath.Join(dir, "plan.html"), tmpl, base, spec.Title+" plan")
	}
	if spec.Tasks.Total > 0 || spec.TasksBody != "" {
		_ = renderMarkdownToFile(tasks, filepath.Join(dir, "tasks.html"), tmpl, base, spec.Title+" tasks")
	}
	return renderTemplate(tmpl, filepath.Join(dir, "index.html"), "spec.html", specPage{
		pageBase:   baseWithNav(base, "specs", spec.Title),
		Spec:       spec,
		BodyHTML:   body,
		PlanHTML:   plan,
		TasksHTML:  tasks,
		Dates:      data.Dates[spec.Number],
		BlockedBy:  BlockedBy(spec, data.Links),
		Blocks:     Blocks(spec.Number, data.Links),
		OriginHref: href,
		OriginOK:   ok,
		Deps:       deps,
	})
}

func renderPRDPages(tmpl *template.Template, outputDir string, base pageBase, data DashboardData) error {
	for _, prd := range data.PRDs {
		body, _ := renderMarkdown(prd.Body)
		dir := filepath.Join(outputDir, "prds", prd.Slug)
		if err := os.MkdirAll(dir, 0o755); err != nil {
			return err
		}
		var related []SpecSummary
		for _, n := range data.Links.OriginToSpecs[prd.Path] {
			if s, ok := data.Links.SpecsByNumber[n]; ok {
				related = append(related, *s)
			}
		}
		if err := renderTemplate(tmpl, filepath.Join(dir, "index.html"), "doc.html", docPage{
			pageBase: baseWithNav(base, "prds", prd.Title),
			Title:    prd.Title,
			Date:     prd.Date,
			BodyHTML: body,
			Related:  related,
		}); err != nil {
			return err
		}
	}
	return nil
}

func renderBugfixPages(tmpl *template.Template, outputDir string, base pageBase, data DashboardData) error {
	for _, bf := range data.Bugfixes {
		body, _ := renderMarkdown(bf.Body)
		dir := filepath.Join(outputDir, "bugfixes", bf.Slug)
		if err := os.MkdirAll(dir, 0o755); err != nil {
			return err
		}
		if err := renderTemplate(tmpl, filepath.Join(dir, "index.html"), "doc.html", docPage{
			pageBase: baseWithNav(base, "bugfixes", bf.Title),
			Title:    bf.Title,
			Date:     bf.Date,
			BodyHTML: body,
		}); err != nil {
			return err
		}
	}
	return nil
}

func renderListPages(tmpl *template.Template, outputDir string, base pageBase, data DashboardData) error {
	var prdItems []listItem
	for _, prd := range data.PRDs {
		cnt := len(data.Links.OriginToSpecs[prd.Path])
		prdItems = append(prdItems, listItem{
			Title: prd.Title,
			Date:  prd.Date,
			Href:  "/prds/" + prd.Slug + "/index.html",
			Count: cnt,
		})
	}
	emptyPRD := ""
	if len(prdItems) == 0 {
		emptyPRD = "No PRDs"
	}
	if err := os.MkdirAll(filepath.Join(outputDir, "prds"), 0o755); err != nil {
		return err
	}
	if err := renderTemplate(tmpl, filepath.Join(outputDir, "prds", "index.html"), "list.html", listPage{
		pageBase: baseWithNav(base, "prds", "PRDs"),
		Items:    prdItems,
		Empty:    emptyPRD,
	}); err != nil {
		return err
	}

	var bfItems []listItem
	for _, bf := range data.Bugfixes {
		bfItems = append(bfItems, listItem{
			Title: bf.Title,
			Date:  bf.Date,
			Href:  "/bugfixes/" + bf.Slug + "/index.html",
		})
	}
	emptyBF := ""
	if len(bfItems) == 0 {
		emptyBF = "No bugfixes"
	}
	if err := os.MkdirAll(filepath.Join(outputDir, "bugfixes"), 0o755); err != nil {
		return err
	}
	return renderTemplate(tmpl, filepath.Join(outputDir, "bugfixes", "index.html"), "list.html", listPage{
		pageBase: baseWithNav(base, "bugfixes", "Bugfixes"),
		Items:    bfItems,
		Empty:    emptyBF,
	})
}

func renderTemplate(tmpl *template.Template, path, name string, data any) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	return tmpl.ExecuteTemplate(f, name, data)
}

func copyStaticAssets(outputDir string) error {
	return fs.WalkDir(Assets, "dashboard_assets", func(path string, d fs.DirEntry, err error) error {
		if err != nil || d.IsDir() {
			return err
		}
		if !strings.HasPrefix(path, "dashboard_assets/css/") && !strings.HasPrefix(path, "dashboard_assets/js/") {
			return nil
		}
		rel := strings.TrimPrefix(path, "dashboard_assets/")
		dest := filepath.Join(outputDir, rel)
		if err := os.MkdirAll(filepath.Dir(dest), 0o755); err != nil {
			return err
		}
		src, err := Assets.Open(path)
		if err != nil {
			return err
		}
		defer src.Close()
		out, err := os.Create(dest)
		if err != nil {
			return err
		}
		_, err = io.Copy(out, src)
		out.Close()
		return err
	})
}

func templateFuncs(data DashboardData) template.FuncMap {
	return template.FuncMap{
		"statusClass": func(s string) string {
			switch strings.ToLower(s) {
			case "implemented":
				return "status-implemented"
			case "approved":
				return "status-approved"
			case "in review":
				return "status-review"
			default:
				return "status-draft"
			}
		},
		"taskPct": func(t TaskStats) int {
			if t.Total == 0 {
				return 0
			}
			return t.Done * 100 / t.Total
		},
		"formatNum": func(n int) string { return fmt.Sprintf("%04d", n) },
		"roundDays": roundDays,
		"json": func(v any) template.JS {
			b, _ := json.Marshal(v)
			return template.JS(b)
		},
	}
}

func computeSummary(data DashboardData) globalSummary {
	var s globalSummary
	for _, fam := range data.Families {
		for _, ph := range fam.Phases {
			for _, spec := range ph.Specs {
				if spec.Deprecated {
					continue
				}
				s.Total++
				switch strings.ToLower(spec.Status) {
				case "implemented":
					s.Implemented++
				case "draft":
					s.Pending++
				default:
					s.InProgress++
				}
			}
		}
	}
	return s
}

func familyCards(data DashboardData) []familyCard {
	var cards []familyCard
	for _, fam := range data.Families {
		c := familyCard{Name: fam.Name, Slug: slugify(fam.Name)}
		for _, ph := range fam.Phases {
			if ph.Number > c.ActivePhase {
				c.ActivePhase = ph.Number
			}
			for _, spec := range ph.Specs {
				if spec.Deprecated {
					continue
				}
				c.SpecCount++
				switch strings.ToLower(spec.Status) {
				case "implemented":
					c.Implemented++
				case "draft":
					c.Pending++
				default:
					c.InProgress++
				}
			}
		}
		if c.SpecCount > 0 {
			c.ProgressPct = float64(c.Implemented) / float64(c.SpecCount) * 100
		}
		cards = append(cards, c)
	}
	return cards
}

func baseWithNav(b pageBase, nav, title string) pageBase {
	b.ActiveNav = nav
	b.Title = title
	return b
}

func renderMarkdownToFile(body template.HTML, path string, tmpl *template.Template, base pageBase, title string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.WriteString("<!DOCTYPE html><html><head><link rel=\"stylesheet\" href=\"/css/style.css\"></head><body><article>")
	if err != nil {
		return err
	}
	_, err = f.Write([]byte(body))
	if err != nil {
		return err
	}
	_, err = f.WriteString("</article></body></html>")
	return err
}

func slugify(s string) string {
	s = strings.ToLower(s)
	s = strings.ReplaceAll(s, " ", "-")
	var b strings.Builder
	for _, r := range s {
		if (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') || r == '-' {
			b.WriteRune(r)
		}
	}
	out := b.String()
	if out == "" {
		return "family"
	}
	return out
}
