package cli

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/charmbracelet/huh"
)

func Run(args []string, stdout, stderr io.Writer) int {
	if len(args) == 0 {
		printHelp(stdout)
		return 0
	}

	catalog, err := loadCatalog()
	if err != nil {
		fmt.Fprintf(stderr, "error: %v\n", err)
		return 1
	}

	var runErr error
	switch args[0] {
	case "help", "--help", "-h":
		printHelp(stdout)
	case "list":
		runErr = runList(stdout, catalog)
	case "init":
		runErr = runInit(args[1:], stdout, catalog)
	case "sync":
		runErr = runSync(args[1:], stdout, catalog)
	case "add":
		runErr = runAdd(args[1:], stdout, catalog)
	case "remove":
		runErr = runRemove(args[1:], stdout, catalog)
	case "global":
		runErr = runGlobal(args[1:], stdout, catalog)
	case "doctor":
		runErr = runDoctor(stdout, catalog)
	default:
		runErr = fmt.Errorf("unknown command %q", args[0])
	}

	if runErr != nil {
		fmt.Fprintf(stderr, "error: %v\n", runErr)
		return 1
	}
	return 0
}

func printHelp(w io.Writer) {
	fmt.Fprintln(w, "geremmyas manages repository-local Copilot agent packs.")
	fmt.Fprintln(w)
	fmt.Fprintln(w, "Usage:")
	fmt.Fprintln(w, "  geremmyas list")
	fmt.Fprintln(w, "  geremmyas init [--packs core,sdd,afk] [--force]")
	fmt.Fprintln(w, "  geremmyas sync [--force]")
	fmt.Fprintln(w, "  geremmyas add <pack>...")
	fmt.Fprintln(w, "  geremmyas remove <pack>...")
	fmt.Fprintln(w, "  geremmyas global <pack>...")
	fmt.Fprintln(w, "  geremmyas doctor")
}

func runList(w io.Writer, catalog Catalog) error {
	for _, pack := range catalog.Packs {
		fmt.Fprintf(w, "%-18s %s\n", pack.Name, pack.Description)
	}
	return nil
}

func runInit(args []string, w io.Writer, catalog Catalog) error {
	fs := flag.NewFlagSet("init", flag.ContinueOnError)
	fs.SetOutput(io.Discard)
	packsFlag := fs.String("packs", "", "comma-separated pack names")
	force := fs.Bool("force", false, "overwrite existing geremmyas.yml")
	if err := fs.Parse(args); err != nil {
		return err
	}

	configPath := filepath.Join(".", configFileName)
	if _, err := os.Stat(configPath); err == nil && !*force {
		return fmt.Errorf("%s already exists; use --force to overwrite", configFileName)
	} else if err != nil && !os.IsNotExist(err) {
		return err
	}

	// Interactive mode when no --packs flag and terminal is available
	if *packsFlag == "" && isInteractive() {
		projectPacks, globalPacks, err := runInteractiveInit(w, catalog)
		if err != nil {
			return err
		}

		// Install project-level packs
		if len(projectPacks) > 0 {
			cfg := Config{Version: 1, Packs: projectPacks}
			if _, err := catalog.Resolve(cfg.Packs); err != nil {
				return err
			}
			if err := os.WriteFile(configPath, []byte(formatConfig(cfg)), 0o644); err != nil {
				return err
			}
			fmt.Fprintf(w, "created %s with %d packs\n", configFileName, len(projectPacks))
		}

		// Install global packs
		if len(globalPacks) > 0 {
			packs, err := catalog.Resolve(globalPacks)
			if err != nil {
				return err
			}
			if err := globalInstallPacks(packs); err != nil {
				return err
			}
			userDir, _ := vsCodeUserDir()
			fmt.Fprintf(w, "installed %d packs globally to %s\n", len(globalPacks), userDir)
		}

		if len(projectPacks) == 0 && len(globalPacks) == 0 {
			fmt.Fprintln(w, "no packs selected")
		}
		return nil
	}

	// Non-interactive: use --packs flag or defaults
	packsList := splitCSV(*packsFlag)
	if len(packsList) == 0 {
		packsList = defaultConfig().Packs
	}
	cfg := Config{Version: 1, Packs: packsList}
	if _, err := catalog.Resolve(cfg.Packs); err != nil {
		return err
	}

	if err := os.WriteFile(configPath, []byte(formatConfig(cfg)), 0o644); err != nil {
		return err
	}
	fmt.Fprintf(w, "created %s\n", configFileName)
	return nil
}

func runSync(args []string, w io.Writer, catalog Catalog) error {
	fs := flag.NewFlagSet("sync", flag.ContinueOnError)
	fs.SetOutput(io.Discard)
	force := fs.Bool("force", false, "overwrite customizable files")
	if err := fs.Parse(args); err != nil {
		return err
	}

	cfg, err := readConfigFile(configFileName)
	if err != nil {
		return err
	}
	packs, err := catalog.Resolve(cfg.Packs)
	if err != nil {
		return err
	}

	root, err := os.Getwd()
	if err != nil {
		return err
	}
	summary, err := syncPacks(root, packs, syncOptions{Force: *force})
	if err != nil {
		return err
	}

	fmt.Fprintf(w, "synced %d packs\n", len(packs))
	fmt.Fprintf(w, "installed=%d updated=%d preserved=%d skipped=%d\n",
		summary.Installed, summary.Updated, summary.Preserved, summary.Skipped)

	if _, statErr := os.Stat(filepath.Join(root, "mise.toml")); statErr == nil {
		fmt.Fprintln(w, "\nhint: run 'mise trust' to trust the mise.toml config file")
	}
	return nil
}

func runAdd(args []string, w io.Writer, catalog Catalog) error {
	if len(args) == 0 {
		return errors.New("add requires at least one pack")
	}
	cfg, err := readConfigFile(configFileName)
	if err != nil {
		return err
	}
	cfg.Packs = uniqueStrings(append(cfg.Packs, args...))
	if _, err := catalog.Resolve(cfg.Packs); err != nil {
		return err
	}
	if err := os.WriteFile(configFileName, []byte(formatConfig(cfg)), 0o644); err != nil {
		return err
	}
	fmt.Fprintf(w, "updated %s\n", configFileName)
	return nil
}

func runRemove(args []string, w io.Writer, catalog Catalog) error {
	cfg, err := readConfigFile(configFileName)
	if err != nil {
		return err
	}

	if len(args) == 0 {
		if !isInteractive() {
			return errors.New("remove requires at least one pack (or run in a terminal for interactive mode)")
		}
		selected, err := runInteractiveRemove(cfg.Packs)
		if err != nil {
			return err
		}
		if len(selected) == 0 {
			fmt.Fprintln(w, "nothing to remove")
			return nil
		}
		args = selected
	}

	remove := map[string]bool{}
	for _, arg := range args {
		remove[arg] = true
	}

	next := []string{}
	for _, pack := range cfg.Packs {
		if !remove[pack] {
			next = append(next, pack)
		}
	}
	cfg.Packs = next
	if len(cfg.Packs) == 0 {
		return errors.New("cannot remove every pack")
	}
	if _, err := catalog.Resolve(cfg.Packs); err != nil {
		return err
	}
	if err := os.WriteFile(configFileName, []byte(formatConfig(cfg)), 0o644); err != nil {
		return err
	}
	fmt.Fprintf(w, "updated %s\n", configFileName)
	return nil
}

func runDoctor(w io.Writer, catalog Catalog) error {
	if err := catalog.ValidateSources(); err != nil {
		return err
	}

	if _, err := os.Stat(configFileName); os.IsNotExist(err) {
		fmt.Fprintln(w, "catalog: ok")
		fmt.Fprintf(w, "%s: missing; run geremmyas init\n", configFileName)
		return nil
	} else if err != nil {
		return err
	}

	cfg, err := readConfigFile(configFileName)
	if err != nil {
		return err
	}
	packs, err := catalog.Resolve(cfg.Packs)
	if err != nil {
		return err
	}
	fmt.Fprintln(w, "catalog: ok")
	fmt.Fprintf(w, "%s: ok\n", configFileName)
	fmt.Fprintf(w, "resolved packs: %d\n", len(packs))
	return nil
}

func readConfigFile(path string) (Config, error) {
	file, err := os.Open(path)
	if os.IsNotExist(err) {
		return Config{}, fmt.Errorf("%s not found; run geremmyas init", configFileName)
	}
	if err != nil {
		return Config{}, err
	}
	defer file.Close()
	return parseConfig(file)
}

func splitCSV(value string) []string {
	parts := strings.Split(value, ",")
	out := make([]string, 0, len(parts))
	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part != "" {
			out = append(out, part)
		}
	}
	return uniqueStrings(out)
}

func runGlobal(args []string, w io.Writer, catalog Catalog) error {
	if len(args) == 0 && isInteractive() {
		// Interactive: show multi-select
		var selected []string
		options := make([]huh.Option[string], len(catalog.Packs))
		for i, p := range catalog.Packs {
			options[i] = huh.NewOption(fmt.Sprintf("%-20s %s", p.Name, p.Description), p.Name)
		}

		form := huh.NewForm(
			huh.NewGroup(
				huh.NewMultiSelect[string]().
					Title("Select packs to install globally").
					Description("Installed to VS Code user-level directory").
					Options(options...).
					Value(&selected),
			),
		)

		if err := form.Run(); err != nil {
			return err
		}

		if len(selected) == 0 {
			fmt.Fprintln(w, "no packs selected")
			return nil
		}
		args = selected
	}

	if len(args) == 0 {
		return errors.New("global requires at least one pack name")
	}

	packs, err := catalog.Resolve(args)
	if err != nil {
		return err
	}

	if err := globalInstallPacks(packs); err != nil {
		return err
	}

	userDir, _ := vsCodeUserDir()
	fmt.Fprintf(w, "installed %d packs globally to %q\n", len(packs), userDir)
	return nil
}
