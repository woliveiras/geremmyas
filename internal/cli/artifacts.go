package cli

import (
	"fmt"
	"path/filepath"
	"sort"
	"strings"
)

const (
	ArtifactAgent               = "agent"
	ArtifactContract            = "contract"
	ArtifactCopilotInstructions = "copilot-instructions"
	ArtifactGuardrail           = "guardrail"
	ArtifactHook                = "hook"
	ArtifactInstruction         = "instruction"
	ArtifactSkill               = "skill"
	ArtifactTemplate            = "template"
	ArtifactTooling             = "tooling"
)

var validArtifactKinds = map[string]bool{
	ArtifactAgent:               true,
	ArtifactContract:            true,
	ArtifactCopilotInstructions: true,
	ArtifactGuardrail:           true,
	ArtifactHook:                true,
	ArtifactInstruction:         true,
	ArtifactSkill:               true,
	ArtifactTemplate:            true,
	ArtifactTooling:             true,
}

func validateLogicalPath(path string) error {
	path = filepath.ToSlash(filepath.Clean(path))
	if path == "" || path == ".." || strings.HasPrefix(path, "../") || filepath.IsAbs(path) {
		return fmt.Errorf("unsafe relative path")
	}
	return nil
}

// planProjectArtifacts maps semantic catalog entries to the files required by
// the selected project targets. The caller remains responsible for copying the
// returned entries and running generators.
func planProjectArtifacts(packs []Pack, targets []string) ([]FileEntry, error) {
	targets = normalizeTargets(targets)
	planned := make([]FileEntry, 0)
	seen := map[string]bool{}

	add := func(entry FileEntry, target string) {
		target = filepath.ToSlash(filepath.Clean(target))
		if seen[target] {
			return
		}
		seen[target] = true
		entry.Target = target
		planned = append(planned, entry)
	}

	for _, pack := range packs {
		for _, entry := range pack.Files {
			if !validArtifactKinds[entry.Kind] {
				return nil, fmt.Errorf("pack %q has invalid artifact kind %q", pack.Name, entry.Kind)
			}
			if err := validateLogicalPath(entry.Path); err != nil {
				return nil, fmt.Errorf("pack %q has invalid artifact path %q: %w", pack.Name, entry.Path, err)
			}

			switch entry.Kind {
			case ArtifactContract, ArtifactTemplate, ArtifactTooling:
				add(entry, entry.Path)
			case ArtifactCopilotInstructions:
				if hasTarget(targets, TargetCopilot) {
					add(entry, ".github/copilot-instructions.md")
				}
			case ArtifactHook:
				if hasTarget(targets, TargetCopilot) {
					add(entry, ".github/hooks")
				}
			case ArtifactGuardrail:
				if hasTarget(targets, TargetCopilot) {
					add(entry, filepath.Join(".github", "hooks", entry.Path))
				}
			case ArtifactAgent:
				if hasAnyNonCopilotTarget(targets) {
					add(entry, filepath.Join(".agents", "roles", entry.Path))
				}
				if hasTarget(targets, TargetCopilot) {
					add(entry, ".github/agents")
				}
			case ArtifactInstruction:
				if hasTarget(targets, TargetCodex) {
					add(entry, filepath.Join(".codex", "instructions", entry.Path))
				}
				if hasTarget(targets, TargetClaudeCode) {
					add(entry, filepath.Join(".claude", "instructions", entry.Path))
				}
				if hasTarget(targets, TargetOpenCode) {
					add(entry, filepath.Join(".opencode", "instructions", entry.Path))
				}
				if hasTarget(targets, TargetCopilot) {
					add(entry, filepath.Join(".github", "instructions", entry.Path))
				}
			case ArtifactSkill:
				if hasAnyNonCopilotTarget(targets) {
					add(entry, filepath.Join(".agents", "skills", entry.Path))
				}
				if hasTarget(targets, TargetCopilot) {
					add(entry, filepath.Join(".github", "skills", entry.Path))
				}
			}
		}
	}

	sort.SliceStable(planned, func(i, j int) bool {
		left, right := planned[i].Target, planned[j].Target
		leftRank, rightRank := projectTargetRank(left), projectTargetRank(right)
		if leftRank != rightRank {
			return leftRank < rightRank
		}
		return left < right
	})
	return planned, nil
}

func hasAnyNonCopilotTarget(targets []string) bool {
	for _, target := range targets {
		if target != TargetCopilot {
			return true
		}
	}
	return false
}

func projectTargetRank(target string) int {
	switch {
	case !strings.HasPrefix(target, "."):
		return 0
	case strings.HasPrefix(target, ".agents/"):
		return 1
	case strings.HasPrefix(target, ".codex/"):
		return 2
	case strings.HasPrefix(target, ".cursor/"):
		return 3
	case strings.HasPrefix(target, ".claude/"):
		return 4
	case strings.HasPrefix(target, ".opencode/"):
		return 5
	case strings.HasPrefix(target, ".github/"):
		return 6
	default:
		return 7
	}
}
