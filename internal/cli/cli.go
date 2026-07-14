package cli

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

var Version = "dev"

func Run(args []string, stdout, stderr io.Writer) int {
	if len(args) == 0 {
		printHelp(stdout)
		return 0
	}

	if args[0] == "version" || args[0] == "--version" {
		printVersion(stdout)
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
	case "version":
		printVersion(stdout)
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
	case "project":
		runErr = runProject(args[1:], stdout, catalog)
	case "global":
		runErr = runGlobal(args[1:], stdout, catalog)
	case "context":
		runErr = runContext(stdout)
	case "lint":
		if runErr = catalog.ValidateTiers(); runErr == nil {
			runErr = runLint(stdout, catalog)
		}
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
	fmt.Fprintln(w, "  geremmyas version")
	fmt.Fprintln(w, "  geremmyas list")
	fmt.Fprintln(w, "  geremmyas init [--packs core,sdd] [--targets copilot,cursor,...] [--force]")
	fmt.Fprintln(w, "  geremmyas sync [--force] [--targets copilot,cursor,claude-code,codex,opencode]")
	fmt.Fprintln(w, "  geremmyas add <pack>...")
	fmt.Fprintln(w, "  geremmyas remove <pack>...")
	fmt.Fprintln(w, "  geremmyas project [--force] [--targets ...] <pack>...")
	fmt.Fprintln(w, "  geremmyas global [--targets copilot,cursor,...] [--force] <pack>...")
	fmt.Fprintln(w, "  geremmyas context")
	fmt.Fprintln(w, "  geremmyas lint")
	fmt.Fprintln(w, "  geremmyas doctor")
}

func printVersion(w io.Writer) {
	fmt.Fprintf(w, "geremmyas %s\n", Version)
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
	targetsFlag := fs.String("targets", "", "comma-separated IDE targets (codex,claude-code,copilot,cursor,opencode)")
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
		projectPacks, globalPacks, err := runInteractiveInit(catalog)
		if err != nil {
			return err
		}

		// Install project-level packs
		if len(projectPacks) > 0 {
			cfg := defaultConfig()
			cfg.Packs = projectPacks
			if err := applyTargetsFlag(&cfg, *targetsFlag); err != nil {
				return err
			}
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
			count, err := globalInstallPacks(packs)
			if err != nil {
				return err
			}
			home := globalInstallDir()
			fmt.Fprintf(w, "installed %d files globally:\n", count)
			fmt.Fprintf(w, "  skills       → %s/.agents/skills/\n", home)
			fmt.Fprintf(w, "  instructions → %s/.copilot/instructions/\n", home)
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
	cfg := defaultConfig()
	cfg.Packs = packsList
	if err := applyTargetsFlag(&cfg, *targetsFlag); err != nil {
		return err
	}
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
	force := fs.Bool("force", false, "overwrite customizable and generated files")
	targetsFlag := fs.String("targets", "", "comma-separated IDE targets (codex,claude-code,copilot,cursor,opencode)")
	if err := fs.Parse(args); err != nil {
		return err
	}

	cfg, err := readConfigFile(configFileName)
	if err != nil {
		return err
	}
	targets := effectiveTargets(cfg, splitCSV(*targetsFlag))
	if err := validateTargets(targets); err != nil {
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
	printSyncSummary(w, len(packs), summary)

	genSummaries, err := runTargetGenerators(root, targets, packs, generatorOptions{Force: *force})
	if err != nil {
		return err
	}
	printGeneratorSummaries(w, genSummaries)

	if hasTarget(targets, TargetCopilot) {
		printMiseHint(w, root)
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

func runProject(args []string, w io.Writer, catalog Catalog) error {
	fs := flag.NewFlagSet("project", flag.ContinueOnError)
	fs.SetOutput(io.Discard)
	force := fs.Bool("force", false, "overwrite customizable and generated files")
	targetsFlag := fs.String("targets", "", "comma-separated IDE targets (codex,claude-code,copilot,cursor,opencode)")
	if err := fs.Parse(args); err != nil {
		return err
	}

	packNames := fs.Args()
	if len(packNames) == 0 && isInteractive() {
		selected, err := runInteractivePackSelect(
			catalog,
			"Select packs to install in this project",
			"Updates geremmyas.yml and syncs files into the current project",
		)
		if err != nil {
			return err
		}
		if len(selected) == 0 {
			fmt.Fprintln(w, "no packs selected")
			return nil
		}
		packNames = selected

		if !*force {
			selectedForce, err := runInteractiveProjectForce()
			if err != nil {
				return err
			}
			*force = selectedForce
		}
	}

	if len(packNames) == 0 {
		return errors.New("project requires at least one pack name")
	}

	cfg, err := readConfigFile(configFileName)
	if err != nil {
		return err
	}
	cfg.Packs = uniqueStrings(append(cfg.Packs, packNames...))
	if err := applyTargetsFlag(&cfg, *targetsFlag); err != nil {
		return err
	}
	targets := cfg.Targets

	packs, err := catalog.Resolve(cfg.Packs)
	if err != nil {
		return err
	}
	if err := os.WriteFile(configFileName, []byte(formatConfig(cfg)), 0o644); err != nil {
		return err
	}
	fmt.Fprintf(w, "updated %s\n", configFileName)

	root, err := os.Getwd()
	if err != nil {
		return err
	}

	summary, err := syncPacks(root, packs, syncOptions{Force: *force})
	if err != nil {
		return err
	}
	printSyncSummary(w, len(packs), summary)

	genSummaries, err := runTargetGenerators(root, targets, packs, generatorOptions{Force: *force})
	if err != nil {
		return err
	}
	printGeneratorSummaries(w, genSummaries)

	if hasTarget(targets, TargetCopilot) {
		printMiseHint(w, root)
	}
	return nil
}

func runDoctor(w io.Writer, catalog Catalog) error {
	if err := catalog.ValidateSources(); err != nil {
		return err
	}
	if err := catalog.ValidateTiers(); err != nil {
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
	fmt.Fprintf(w, "targets: %s\n", strings.Join(normalizeTargets(cfg.Targets), ", "))
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

func printSyncSummary(w io.Writer, packCount int, summary syncSummary) {
	fmt.Fprintf(w, "synced %d packs\n", packCount)
	fmt.Fprintf(w, "installed=%d updated=%d preserved=%d skipped=%d\n",
		summary.Installed, summary.Updated, summary.Preserved, summary.Skipped)
}

func printMiseHint(w io.Writer, root string) {
	if _, statErr := os.Stat(filepath.Join(root, "mise.toml")); statErr == nil {
		fmt.Fprintln(w, "\nhint: run 'mise trust' to trust the mise.toml config file")
	}
}

func runGlobal(args []string, w io.Writer, catalog Catalog) error {
	fs := flag.NewFlagSet("global", flag.ContinueOnError)
	fs.SetOutput(io.Discard)
	targetsFlag := fs.String("targets", "", "comma-separated IDE targets (codex,claude-code,copilot,cursor,opencode)")
	force := fs.Bool("force", false, "overwrite customized generated files")
	if err := fs.Parse(args); err != nil {
		return err
	}
	packArgs := fs.Args()

	if len(packArgs) == 0 && isInteractive() {
		selected, err := runInteractivePackSelect(
			catalog,
			"Select packs to install globally",
			"Installed to user-level directories (~/.agents, ~/.cursor, …)",
		)
		if err != nil {
			return err
		}
		if len(selected) == 0 {
			fmt.Fprintln(w, "no packs selected")
			return nil
		}
		packArgs = selected
	}

	if len(packArgs) == 0 {
		return errors.New("global requires at least one pack name")
	}

	targets := normalizeTargets(splitCSV(*targetsFlag))
	if err := validateTargets(targets); err != nil {
		return err
	}

	packs, err := catalog.Resolve(packArgs)
	if err != nil {
		return err
	}
	previousManifest, manifestExists, err := loadGlobalManifest()
	if err != nil {
		return err
	}
	if !manifestExists {
		previousManifest, err = adoptKnownGlobalFiles(previousManifest, catalog)
		if err != nil {
			return err
		}
	}
	desiredPaths, err := globalDesiredPaths(packs, targets)
	if err != nil {
		return err
	}

	copySkills, copyInstructions := globalCopyFlags(targets)
	count := 0
	if copySkills || copyInstructions {
		count, err = globalInstallPacksFiltered(packs, copySkills, copyInstructions)
		if err != nil {
			return err
		}
	}

	genTargets := generatorTargets(targets)
	if len(genTargets) > 0 {
		genSummaries, err := runGlobalTargetGenerators(genTargets, packs, generatorOptions{Force: *force})
		if err != nil {
			return err
		}
		printGeneratorSummaries(w, genSummaries)
	}

	packNames := make([]string, 0, len(packs))
	for _, pack := range packs {
		packNames = append(packNames, pack.Name)
	}
	reconcileSummary, err := reconcileGlobalManifest(previousManifest, desiredPaths, packNames, targets)
	if err != nil {
		return err
	}
	fmt.Fprintf(w, "reconciled global state: removed=%d preserved=%d\n",
		reconcileSummary.Removed, reconcileSummary.Preserved)

	home := globalInstallDir()
	if count > 0 {
		fmt.Fprintf(w, "installed %d files globally:\n", count)
		if copySkills {
			fmt.Fprintf(w, "  skills       → %s/.agents/skills/\n", home)
		}
		if copyInstructions {
			fmt.Fprintf(w, "  instructions → %s/.copilot/instructions/\n", home)
		}
	}
	if hasTarget(targets, TargetCursor) {
		fmt.Fprintf(w, "  cursor rules → %s/.cursor/rules/\n", home)
		fmt.Fprintf(w, "  cursor hooks → %s/.cursor/hooks.json\n", home)
	}
	if hasTarget(targets, TargetClaudeCode) {
		fmt.Fprintf(w, "  claude-code  → %s/.claude/CLAUDE.md\n", home)
	}
	if hasTarget(targets, TargetCodex) {
		fmt.Fprintf(w, "  codex        → %s/.codex/AGENTS.md\n", home)
		fmt.Fprintf(w, "  codex instr. → %s/.codex/instructions/\n", home)
	}
	if hasTarget(targets, TargetOpenCode) {
		fmt.Fprintf(w, "  opencode     → %s/.config/opencode/AGENTS.md\n", home)
	}
	return nil
}
