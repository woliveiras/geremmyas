package cli

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

type generatorOptions struct {
	Force bool
}

type generatorSummary struct {
	Installed int
	Updated   int
	Preserved int
	Skipped   int
}

func runTargetGenerators(root string, targets []string, packs []Pack, opts generatorOptions) (map[string]generatorSummary, error) {
	summaries := map[string]generatorSummary{}
	artifacts := collectPackArtifacts(packs)

	if hasTarget(targets, TargetCursor) {
		summary, err := generateCursor(root, artifacts, opts)
		if err != nil {
			return summaries, fmt.Errorf("cursor: %w", err)
		}
		summaries[TargetCursor] = summary
	}
	if hasTarget(targets, TargetClaudeCode) {
		summary, err := generateClaudeCode(root, artifacts, opts)
		if err != nil {
			return summaries, fmt.Errorf("claude-code: %w", err)
		}
		summaries[TargetClaudeCode] = summary
	}
	if hasTarget(targets, TargetOpenCode) {
		summary, err := generateOpenCode(root, artifacts, opts)
		if err != nil {
			return summaries, fmt.Errorf("opencode: %w", err)
		}
		summaries[TargetOpenCode] = summary
	}

	return summaries, nil
}

func writeGeneratedFile(root, relPath string, content []byte, opts generatorOptions, summary *generatorSummary) error {
	dest := filepath.Join(root, filepath.FromSlash(relPath))
	current, readErr := os.ReadFile(dest)
	if readErr == nil {
		if string(current) == string(content) {
			summary.Skipped++
			return nil
		}
		if !isGeneratedFile(current) && !opts.Force {
			summary.Preserved++
			return nil
		}
		summary.Updated++
	} else if os.IsNotExist(readErr) {
		summary.Installed++
	} else {
		return readErr
	}

	if err := os.MkdirAll(filepath.Dir(dest), 0o755); err != nil {
		return err
	}
	return os.WriteFile(dest, content, 0o644)
}

func isGeneratedFile(data []byte) bool {
	return strings.Contains(string(data), generatedMarker)
}

func printGeneratorSummaries(w io.Writer, summaries map[string]generatorSummary) {
	if len(summaries) == 0 {
		return
	}
	fmt.Fprintln(w, "generators:")
	for _, target := range []string{TargetCursor, TargetClaudeCode, TargetOpenCode} {
		summary, ok := summaries[target]
		if !ok {
			continue
		}
		fmt.Fprintf(w, "  %s: installed=%d updated=%d preserved=%d skipped=%d\n",
			target, summary.Installed, summary.Updated, summary.Preserved, summary.Skipped)
	}
}
