package cli

import (
	"fmt"
	"io"
	"os"

	"github.com/charmbracelet/huh"
)

func runInteractiveInit(w io.Writer, catalog Catalog) (projectPacks []string, globalPacks []string, err error) {
	if !isInteractive() {
		return nil, nil, fmt.Errorf("interactive mode requires a terminal; use --packs flag instead")
	}

	var projectSelected []string
	var globalSelected []string

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewMultiSelect[string]().
				Title("Select packs to install in this project").
				Description("Use space to select, enter to confirm").
				Options(packOptions(catalog)...).
				Value(&projectSelected),
		),
		huh.NewGroup(
			huh.NewMultiSelect[string]().
				Title("Select packs to install globally (VS Code user-level)").
				Description("These will be available in all your projects").
				Options(packOptions(catalog)...).
				Value(&globalSelected),
		),
	)

	if err := form.Run(); err != nil {
		return nil, nil, err
	}

	return projectSelected, globalSelected, nil
}

func runInteractivePackSelect(catalog Catalog, title, description string) ([]string, error) {
	if !isInteractive() {
		return nil, fmt.Errorf("interactive mode requires a terminal")
	}

	var selected []string
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewMultiSelect[string]().
				Title(title).
				Description(description).
				Options(packOptions(catalog)...).
				Value(&selected),
		),
	)

	if err := form.Run(); err != nil {
		return nil, err
	}
	return selected, nil
}

func runInteractiveProjectForce() (bool, error) {
	force := false
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewConfirm().
				Title("Overwrite customizable project files?").
				Description("Choose no to preserve existing AGENTS.md, mise.toml, Copilot instructions, and hooks.").
				Affirmative("Overwrite").
				Negative("Preserve").
				Value(&force),
		),
	)

	if err := form.Run(); err != nil {
		return false, err
	}
	return force, nil
}

func packOptions(catalog Catalog) []huh.Option[string] {
	options := make([]huh.Option[string], len(catalog.Packs))
	for i, p := range catalog.Packs {
		options[i] = huh.NewOption(fmt.Sprintf("%-20s %s", p.Name, p.Description), p.Name)
	}
	return options
}

func isInteractive() bool {
	fi, err := os.Stdin.Stat()
	if err != nil {
		return false
	}
	return fi.Mode()&os.ModeCharDevice != 0
}

func runInteractiveRemove(installed []string) ([]string, error) {
	if len(installed) == 0 {
		return nil, fmt.Errorf("no packs installed")
	}

	options := make([]huh.Option[string], len(installed))
	for i, p := range installed {
		options[i] = huh.NewOption(p, p)
	}

	var selected []string
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewMultiSelect[string]().
				Title("Select packs to remove").
				Description("Use space to select, enter to confirm").
				Options(options...).
				Value(&selected),
		),
	)

	if err := form.Run(); err != nil {
		return nil, err
	}
	return selected, nil
}
