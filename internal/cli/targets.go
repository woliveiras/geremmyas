package cli

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"sort"
	"strings"

	geremmyas "github.com/woliveiras/geremmyas"
)

const (
	TargetCopilot    = "copilot"
	TargetCursor     = "cursor"
	TargetClaudeCode = "claude-code"
	TargetOpenCode   = "opencode"
)

var validTargets = map[string]struct{}{
	TargetCopilot:    {},
	TargetCursor:     {},
	TargetClaudeCode: {},
	TargetOpenCode:   {},
}

func defaultTargets() []string {
	return []string{TargetCopilot}
}

func normalizeTargets(values []string) []string {
	if len(values) == 0 {
		return defaultTargets()
	}
	out := make([]string, 0, len(values))
	seen := map[string]bool{}
	for _, value := range values {
		value = strings.ToLower(strings.TrimSpace(value))
		if value == "" || seen[value] {
			continue
		}
		if _, ok := validTargets[value]; !ok {
			continue
		}
		seen[value] = true
		out = append(out, value)
	}
	sort.Strings(out)
	if len(out) == 0 {
		return defaultTargets()
	}
	return out
}

func validateTargets(values []string) error {
	if len(values) == 0 {
		return nil
	}
	for _, value := range values {
		value = strings.ToLower(strings.TrimSpace(value))
		if value == "" {
			continue
		}
		if _, ok := validTargets[value]; !ok {
			return fmt.Errorf("unsupported target %q (valid: copilot, cursor, claude-code, opencode)", value)
		}
	}
	return nil
}

func hasTarget(targets []string, target string) bool {
	for _, item := range targets {
		if item == target {
			return true
		}
	}
	return false
}

type packArtifacts struct {
	instructions []string
	skills       []string
	agents       []string
	hasHooks     bool
	hasAGENTS    bool
}

func collectPackArtifacts(packs []Pack) packArtifacts {
	artifacts := packArtifacts{}
	instrSeen := map[string]bool{}
	skillSeen := map[string]bool{}
	agentSeen := map[string]bool{}

	for _, pack := range packs {
		for _, entry := range pack.Files {
			target := filepathToSlash(entry.Target)
			source := filepathToSlash(entry.Source)

			switch {
			case strings.HasPrefix(target, ".github/instructions/") && strings.HasSuffix(target, ".instructions.md"):
				addUnique(&artifacts.instructions, source, instrSeen)
			case target == ".github/agents" || strings.HasPrefix(target, ".github/agents/"):
				_ = walkEmbeddedMatches(source, func(path string) error {
					if strings.HasSuffix(path, ".agent.md") {
						addUnique(&artifacts.agents, path, agentSeen)
					}
					return nil
				})
			case strings.HasPrefix(target, ".github/skills/"):
				if strings.HasSuffix(source, "/SKILL.md") || strings.HasSuffix(source, "SKILL.md") {
					addUnique(&artifacts.skills, filepath.Dir(source), skillSeen)
				} else {
					_ = walkEmbeddedMatches(source, func(path string) error {
						if strings.HasSuffix(path, "/SKILL.md") || strings.HasSuffix(path, "SKILL.md") {
							addUnique(&artifacts.skills, filepath.Dir(path), skillSeen)
						}
						return nil
					})
				}
			case target == ".github/hooks" || strings.HasPrefix(target, ".github/hooks/"):
				artifacts.hasHooks = true
			case target == "AGENTS.md":
				artifacts.hasAGENTS = true
			}
		}
	}

	sort.Strings(artifacts.instructions)
	sort.Strings(artifacts.skills)
	sort.Strings(artifacts.agents)
	return artifacts
}

func addUnique(list *[]string, value string, seen map[string]bool) {
	value = filepathToSlash(value)
	if value == "" || seen[value] {
		return
	}
	seen[value] = true
	*list = append(*list, value)
}

func walkEmbeddedMatches(root string, fn func(path string) error) error {
	root = filepathToSlash(root)
	info, err := fs.Stat(geremmyas.EmbeddedFiles, root)
	if err != nil {
		return err
	}
	if !info.IsDir() {
		return fn(root)
	}
	return fs.WalkDir(geremmyas.EmbeddedFiles, root, func(path string, d fs.DirEntry, err error) error {
		if err != nil || d.IsDir() {
			return err
		}
		return fn(filepathToSlash(path))
	})
}

func filepathToSlash(path string) string {
	return strings.ReplaceAll(path, "\\", "/")
}

const generatedMarker = "geremmyas:generated"
