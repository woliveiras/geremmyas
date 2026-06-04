package dashboard

import (
	"fmt"
	"io"
	"os"
)

// BuildOptions configures a full dashboard build.
type BuildOptions struct {
	Root       string
	OutputDir  string
	NoGit      bool
	NoCache    bool
	Quiet      io.Writer
}

// Build runs parse → git dates → metrics → render → readme.
func Build(opts BuildOptions) (DashboardData, error) {
	data, err := LoadAll(opts.Root)
	if err != nil {
		return data, err
	}
	for _, msg := range DetectCycles(data.Links) {
		data.Warnings = append(data.Warnings, Warning{Message: msg})
	}
	if err := ExtractGitDates(&data, GitOptions{Root: opts.Root, NoGit: opts.NoGit, NoCache: opts.NoCache}); err != nil {
		return data, err
	}
	ComputeMetrics(&data)
	if err := os.MkdirAll(opts.OutputDir, 0o755); err != nil {
		return data, err
	}
	if err := RenderDashboard(data, opts.OutputDir); err != nil {
		return data, err
	}
	readmePaths, err := WriteSpecReadmes(opts.Root, data)
	if err != nil {
		return data, err
	}
	if opts.Quiet != nil {
		for _, w := range data.Warnings {
			fmt.Fprintf(opts.Quiet, "warning: %s: %s\n", w.Path, w.Message)
		}
		fmt.Fprintf(opts.Quiet, "dashboard: %s\n", opts.OutputDir)
		for _, p := range readmePaths {
			fmt.Fprintf(opts.Quiet, "updated %s\n", p)
		}
	}
	return data, nil
}
