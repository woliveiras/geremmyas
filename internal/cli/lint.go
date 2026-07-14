package cli

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"unicode/utf8"
)

type lintViolation struct {
	Code    string
	Message string
}

const (
	lintViolationMissingTrigger     = "missing-use-when"
	lintViolationMissingNegative    = "missing-negative-scope"
	lintViolationDescriptionTooLong = "description-too-long"
	lintViolationDescriptionMarkup  = "description-markup"
	lintViolationNameMismatch       = "name-mismatch"
	lintViolationBodyTooLong        = "body-too-long"
	lintViolationMissingSkillFile   = "missing-skill-md"
	lintViolationNestedSkillFile    = "nested-skill-md"
	lintViolationSDDSkillBudget     = "sdd-skill-budget"
	lintViolationAgentsWordBudget   = "agents-word-budget"
	maxSkillDescriptionLength       = 240
	maxSkillBodyLines               = 250
	maxSDDPublicSkills              = 10
	maxAgentsWords                  = 700
)

type lintFinding struct {
	Path       string
	Violations []lintViolation
}

func lintDescription(description string) []lintViolation {
	violations := []lintViolation{}
	if !hasUseWhenPhrase(description) {
		violations = append(violations, lintViolation{
			Code:    lintViolationMissingTrigger,
			Message: "description must contain a use when trigger phrase",
		})
	}
	if !hasNegativeScopePhrase(description) {
		violations = append(violations, lintViolation{
			Code:    lintViolationMissingNegative,
			Message: "description must contain a negative-scope phrase",
		})
	}
	if utf8.RuneCountInString(description) > maxSkillDescriptionLength {
		violations = append(violations, lintViolation{
			Code:    lintViolationDescriptionTooLong,
			Message: fmt.Sprintf("description must be at most %d characters", maxSkillDescriptionLength),
		})
	}
	if strings.ContainsAny(description, "<>") {
		violations = append(violations, lintViolation{
			Code:    lintViolationDescriptionMarkup,
			Message: "description must not contain angle-bracket markup",
		})
	}
	return violations
}

func lintName(name, directory string) []lintViolation {
	if name == directory {
		return nil
	}
	return []lintViolation{{
		Code:    lintViolationNameMismatch,
		Message: "skill name must match directory name",
	}}
}

func lintBody(body string) []lintViolation {
	if countBodyLines(body) <= maxSkillBodyLines {
		return nil
	}
	return []lintViolation{{
		Code:    lintViolationBodyTooLong,
		Message: fmt.Sprintf("skill body must be at most %d lines", maxSkillBodyLines),
	}}
}

func countBodyLines(body string) int {
	if body == "" {
		return 0
	}
	body = strings.TrimRight(body, "\n")
	if body == "" {
		return 0
	}
	return strings.Count(body, "\n") + 1
}

func hasUseWhenPhrase(description string) bool {
	return strings.Contains(strings.ToLower(description), "use when")
}

func hasNegativeScopePhrase(description string) bool {
	lower := strings.ToLower(description)
	phrases := []string{"do not use", "don't use", "not for"}
	for _, phrase := range phrases {
		if strings.Contains(lower, phrase) {
			return true
		}
	}
	return false
}

func runLint(w io.Writer, catalog Catalog) error {
	root, err := os.Getwd()
	if err != nil {
		return err
	}
	findings, checked, err := collectLintFindings(filepath.Join(root, "project/.github/skills"), root)
	if err != nil {
		return err
	}
	budgetFindings, err := collectRepositoryBudgetFindings(catalog, root)
	if err != nil {
		return err
	}
	findings = append(findings, budgetFindings...)
	sort.Slice(findings, func(i, j int) bool {
		return findings[i].Path < findings[j].Path
	})
	if len(findings) == 0 {
		fmt.Fprintf(w, "lint: ok (%d skills checked)\n", checked)
		return nil
	}
	total := countLintViolations(findings)
	fmt.Fprintf(w, "lint: %d violation(s) in %d skill(s) checked\n", total, checked)
	for _, finding := range findings {
		details := make([]string, 0, len(finding.Violations))
		for _, violation := range finding.Violations {
			details = append(details, fmt.Sprintf("%s (%s)", violation.Code, violation.Message))
		}
		fmt.Fprintf(w, "  %s: %s\n", finding.Path, strings.Join(details, "; "))
	}
	return fmt.Errorf("lint found %d violation(s)", total)
}

func collectLintFindings(skillsRoot, root string) ([]lintFinding, int, error) {
	if _, err := os.Stat(skillsRoot); err != nil {
		return nil, 0, err
	}
	entries := []lintFinding{}
	checked := 0
	directories := map[string]struct{}{}
	err := filepath.WalkDir(skillsRoot, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			rel, relErr := filepath.Rel(skillsRoot, path)
			if relErr != nil || rel == "." {
				return nil
			}
			parts := strings.Split(filepath.ToSlash(rel), "/")
			if len(parts) == 1 {
				directories[path] = struct{}{}
			}
			return nil
		}
		if d.Name() != "SKILL.md" {
			return nil
		}
		checked++
		rel, relErr := filepath.Rel(skillsRoot, path)
		if relErr != nil {
			return relErr
		}
		if strings.Count(filepath.ToSlash(rel), "/") > 1 {
			entries = append(entries, lintFinding{
				Path: relativeLintPath(root, path),
				Violations: []lintViolation{{
					Code:    lintViolationNestedSkillFile,
					Message: "support files must not be named SKILL.md",
				}},
			})
			return nil
		}
		finding, err := lintSkillFile(path, root)
		if err != nil {
			return err
		}
		if len(finding.Violations) > 0 {
			entries = append(entries, finding)
		}
		return nil
	})
	if err != nil {
		return nil, 0, err
	}
	for dir := range directories {
		skillPath := filepath.Join(dir, "SKILL.md")
		if _, err := os.Stat(skillPath); os.IsNotExist(err) {
			entries = append(entries, lintFinding{
				Path: relativeLintPath(root, skillPath),
				Violations: []lintViolation{{
					Code:    lintViolationMissingSkillFile,
					Message: "skill directory must contain SKILL.md",
				}},
			})
		}
	}
	sort.Slice(entries, func(i, j int) bool {
		return entries[i].Path < entries[j].Path
	})
	return entries, checked, nil
}

func collectRepositoryBudgetFindings(catalog Catalog, root string) ([]lintFinding, error) {
	findings := []lintFinding{}
	if count := countSDDPublicSkills(catalog); count > maxSDDPublicSkills {
		findings = append(findings, lintFinding{
			Path: "catalog/packs.json",
			Violations: []lintViolation{{
				Code:    lintViolationSDDSkillBudget,
				Message: fmt.Sprintf("sdd pack must expose at most %d skills; found %d", maxSDDPublicSkills, count),
			}},
		})
	}

	agentsPath := filepath.Join(root, "project/AGENTS.md")
	data, err := os.ReadFile(agentsPath)
	if err != nil && !os.IsNotExist(err) {
		return nil, err
	}
	if err == nil {
		if count := len(strings.Fields(string(data))); count > maxAgentsWords {
			findings = append(findings, lintFinding{
				Path: "project/AGENTS.md",
				Violations: []lintViolation{{
					Code:    lintViolationAgentsWordBudget,
					Message: fmt.Sprintf("project/AGENTS.md must contain at most %d words; found %d", maxAgentsWords, count),
				}},
			})
		}
	}
	return findings, nil
}

func countSDDPublicSkills(catalog Catalog) int {
	seen := map[string]struct{}{}
	for _, pack := range catalog.Packs {
		if pack.Name != "sdd" {
			continue
		}
		for _, entry := range pack.Files {
			const prefix = ".github/skills/"
			target := filepath.ToSlash(entry.Target)
			if !strings.HasPrefix(target, prefix) {
				continue
			}
			name := strings.TrimPrefix(target, prefix)
			if name == "" || strings.Contains(name, "/") {
				continue
			}
			seen[name] = struct{}{}
		}
	}
	return len(seen)
}

func lintSkillFile(path, root string) (lintFinding, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return lintFinding{}, err
	}
	fm, body, err := parseMarkdownFrontmatter(string(data))
	if err != nil {
		return lintFinding{}, err
	}
	directory := filepath.Base(filepath.Dir(path))
	finding := lintFinding{Path: relativeLintPath(root, path)}
	finding.Violations = append(finding.Violations, lintDescription(fm.get("description"))...)
	finding.Violations = append(finding.Violations, lintName(fm.get("name"), directory)...)
	finding.Violations = append(finding.Violations, lintBody(body)...)
	return finding, nil
}

func relativeLintPath(root, path string) string {
	rel, err := filepath.Rel(root, path)
	if err != nil {
		return filepath.ToSlash(path)
	}
	return filepath.ToSlash(rel)
}

func countLintViolations(findings []lintFinding) int {
	total := 0
	for _, finding := range findings {
		total += len(finding.Violations)
	}
	return total
}
