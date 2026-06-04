package cli

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/woliveiras/geremmyas/internal/cli/dashboard"
)

const defaultDashboardOut = ".geremmyas/dashboard"

func runDashboard(args []string, w io.Writer) error {
	fs := flag.NewFlagSet("dashboard", flag.ContinueOnError)
	fs.SetOutput(io.Discard)
	output := fs.String("output", defaultDashboardOut, "output directory")
	serve := fs.Bool("serve", false, "start local HTTP server")
	port := fs.Int("port", 8080, "server port")
	watch := fs.Bool("watch", false, "rebuild on file changes (implies --serve)")
	noGit := fs.Bool("no-git", false, "skip git history extraction")
	noCache := fs.Bool("no-cache", false, "force full git rescan")
	if err := fs.Parse(args); err != nil {
		return err
	}
	if *watch {
		*serve = true
	}

	root, err := os.Getwd()
	if err != nil {
		return err
	}
	outDir := filepath.Join(root, *output)
	if filepath.IsAbs(*output) {
		outDir = *output
	}

	build := func() error {
		_, err := dashboard.Build(dashboard.BuildOptions{
			Root:      root,
			OutputDir: outDir,
			NoGit:     *noGit,
			NoCache:   *noCache,
			Quiet:     w,
		})
		return err
	}
	if err := build(); err != nil {
		return err
	}
	if !*serve {
		return nil
	}

	srv := dashboard.NewServer(outDir, *port)
	if *watch {
		dirs := []string{
			filepath.Join(root, "specs"),
			filepath.Join(root, "docs"),
		}
		stopWatch, err := dashboard.WatchDirs(dirs, 300*time.Millisecond, func() {
			if err := build(); err != nil {
				fmt.Fprintf(w, "rebuild error: %v\n", err)
				return
			}
			srv.NotifyReload()
		})
		if err != nil {
			return err
		}
		defer stopWatch()
	}
	return srv.Run()
}
