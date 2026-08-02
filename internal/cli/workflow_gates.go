package cli

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

const workflowGateInventoryPath = "catalog/workflow-gates.json"

var conversationalGatePattern = regexp.MustCompile(`(?i)(^ASK\s|approval gate|explicit (human )?(approval|authority)|explicitly approv|user approval|user has approv|\bask (the )?user\b|\bask whether\b|\bagent asks whether\b|\bask before\b|\bask one focused question\b|\bask for (the paper|paper content|the minimum missing input|what's genuinely missing)|stop (and wait|for approval|until)|without (explicit )?(permission|confirmation|approval|authorization)|permission question|requires confirmation|obtain user approval|human-approved|manual approval|human (accessibility review|listening|play|playtesting|review))`)

type workflowGateInventory struct {
	Version int                `json:"version"`
	Rules   []workflowGateRule `json:"rules"`
}

type workflowGateRule struct {
	ID             string `json:"id"`
	Classification string `json:"classification"`
	Path           string `json:"path"`
	Pattern        string `json:"pattern"`
	Note           string `json:"note"`
}

type workflowGateMatch struct {
	Path           string
	Line           int
	Text           string
	RuleID         string
	Classification string
}

type workflowGateViolation struct {
	Path    string
	Line    int
	Message string
}

func loadWorkflowGateInventory(fsys fs.FS, name string) (workflowGateInventory, error) {
	data, err := fs.ReadFile(fsys, name)
	if err != nil {
		return workflowGateInventory{}, err
	}
	var inventory workflowGateInventory
	if err := json.Unmarshal(data, &inventory); err != nil {
		return workflowGateInventory{}, fmt.Errorf("parse workflow gate inventory: %w", err)
	}
	if inventory.Version != 1 {
		return workflowGateInventory{}, fmt.Errorf("unsupported workflow gate inventory version %d", inventory.Version)
	}
	allowed := map[string]bool{
		"removed":                     true,
		"retained-authority-boundary": true,
		"residual-evidence":           true,
		"deferred-release-work":       true,
	}
	seen := map[string]bool{}
	for _, rule := range inventory.Rules {
		if rule.ID == "" || seen[rule.ID] {
			return workflowGateInventory{}, fmt.Errorf("workflow gate inventory has missing or duplicate rule id %q", rule.ID)
		}
		seen[rule.ID] = true
		if !allowed[rule.Classification] {
			return workflowGateInventory{}, fmt.Errorf("workflow gate rule %q has invalid classification %q", rule.ID, rule.Classification)
		}
		if _, err := path.Match(rule.Path, "probe"); err != nil {
			return workflowGateInventory{}, fmt.Errorf("workflow gate rule %q has invalid path pattern: %w", rule.ID, err)
		}
		if _, err := regexp.Compile(rule.Pattern); err != nil {
			return workflowGateInventory{}, fmt.Errorf("workflow gate rule %q has invalid text pattern: %w", rule.ID, err)
		}
	}
	return inventory, nil
}

func scanWorkflowGates(fsys fs.FS, files []string, inventory workflowGateInventory) ([]workflowGateMatch, []workflowGateViolation, error) {
	type compiledRule struct {
		workflowGateRule
		pattern *regexp.Regexp
	}
	rules := make([]compiledRule, 0, len(inventory.Rules))
	for _, rule := range inventory.Rules {
		compiled, err := regexp.Compile(rule.Pattern)
		if err != nil {
			return nil, nil, fmt.Errorf("compile workflow gate rule %q: %w", rule.ID, err)
		}
		rules = append(rules, compiledRule{workflowGateRule: rule, pattern: compiled})
	}

	files = uniqueStrings(files)
	sort.Strings(files)
	matches := []workflowGateMatch{}
	violations := []workflowGateViolation{}
	ruleMatches := map[string]int{}
	for _, name := range files {
		file, err := fsys.Open(name)
		if err != nil {
			return nil, nil, err
		}
		scanner := bufio.NewScanner(file)
		lineNumber := 0
		for scanner.Scan() {
			lineNumber++
			line := strings.TrimSpace(scanner.Text())
			if !conversationalGatePattern.MatchString(line) {
				continue
			}
			matched := false
			matchedClassifications := map[string]bool{}
			for _, rule := range rules {
				pathMatches, _ := path.Match(rule.Path, name)
				if pathMatches && rule.pattern.MatchString(line) {
					match := workflowGateMatch{
						Path: name, Line: lineNumber, Text: line,
						RuleID: rule.ID, Classification: rule.Classification,
					}
					matches = append(matches, match)
					ruleMatches[rule.ID]++
					matchedClassifications[rule.Classification] = true
					if rule.Classification == "removed" {
						violations = append(violations, workflowGateViolation{
							Path: name, Line: lineNumber,
							Message: fmt.Sprintf("removed conversational gate %q reappeared: %s", rule.ID, line),
						})
					}
					matched = true
				}
			}
			if len(matchedClassifications) > 1 {
				classifications := make([]string, 0, len(matchedClassifications))
				for classification := range matchedClassifications {
					classifications = append(classifications, classification)
				}
				sort.Strings(classifications)
				violations = append(violations, workflowGateViolation{
					Path: name, Line: lineNumber,
					Message: "ambiguous workflow gate classifications: " + strings.Join(classifications, ", "),
				})
			}
			if !matched {
				violations = append(violations, workflowGateViolation{
					Path: name, Line: lineNumber,
					Message: "unclassified conversational gate: " + line,
				})
			}
		}
		scanErr := scanner.Err()
		closeErr := file.Close()
		if scanErr != nil {
			return nil, nil, scanErr
		}
		if closeErr != nil {
			return nil, nil, closeErr
		}
	}
	for _, rule := range rules {
		if rule.Classification != "removed" && rule.Classification != "deferred-release-work" && ruleMatches[rule.ID] == 0 {
			violations = append(violations, workflowGateViolation{
				Path:    workflowGateInventoryPath,
				Message: fmt.Sprintf("stale workflow gate rule %q matched no surface", rule.ID),
			})
		}
	}
	return matches, violations, nil
}

func collectWorkflowSurfaceFiles(root string, catalog Catalog) ([]string, error) {
	paths := map[string]bool{}
	addPath := func(rel string) error {
		full := filepath.Join(root, filepath.FromSlash(rel))
		info, err := os.Stat(full)
		if err != nil {
			return err
		}
		if !info.IsDir() {
			paths[filepath.ToSlash(rel)] = true
			return nil
		}
		return filepath.WalkDir(full, func(filePath string, entry os.DirEntry, walkErr error) error {
			if walkErr != nil {
				return walkErr
			}
			if entry.IsDir() {
				return nil
			}
			relPath, err := filepath.Rel(root, filePath)
			if err != nil {
				return err
			}
			paths[filepath.ToSlash(relPath)] = true
			return nil
		})
	}

	for _, pack := range catalog.Packs {
		for _, entry := range pack.Files {
			if err := addPath(entry.Source); err != nil {
				return nil, err
			}
		}
	}
	for _, rel := range []string{"content/prompts", "targets", "README.md", "docs/guardrails-framework.md"} {
		if err := addPath(rel); err != nil {
			return nil, err
		}
	}
	result := make([]string, 0, len(paths))
	for rel := range paths {
		result = append(result, rel)
	}
	sort.Strings(result)
	return result, nil
}
