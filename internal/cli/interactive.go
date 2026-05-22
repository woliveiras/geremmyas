package cli

import (
	"fmt"
	"io"
	"os"

	"github.com/charmbracelet/huh"
)

type packChoice struct {
	Name        string
	Description string
	Level       string // "project", "global", "skip"
}

func runInteractiveInit(w io.Writer, catalog Catalog) (projectPacks []string, globalPacks []string, err error) {
	if !isInteractive() {
		return nil, nil, fmt.Errorf("interactive mode requires a terminal; use --packs flag instead")
	}

	choices := make([]packChoice, len(catalog.Packs))
	for i, p := range catalog.Packs {
		choices[i] = packChoice{
			Name:        p.Name,
			Description: p.Description,
			Level:       "skip",
		}
	}

	var projectSelected []string
	var globalSelected []string

	projectOptions := make([]huh.Option[string], len(catalog.Packs))
	for i, p := range catalog.Packs {
		projectOptions[i] = huh.NewOption(fmt.Sprintf("%-20s %s", p.Name, p.Description), p.Name)
	}

	globalOptions := make([]huh.Option[string], len(catalog.Packs))
	for i, p := range catalog.Packs {
		globalOptions[i] = huh.NewOption(fmt.Sprintf("%-20s %s", p.Name, p.Description), p.Name)
	}

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewMultiSelect[string]().
				Title("Select packs to install in this project").
				Description("Use space to select, enter to confirm").
				Options(projectOptions...).
				Value(&projectSelected),
		),
		huh.NewGroup(
			huh.NewMultiSelect[string]().
				Title("Select packs to install globally (VS Code user-level)").
				Description("These will be available in all your projects").
				Options(globalOptions...).
				Value(&globalSelected),
		),
	)

	if err := form.Run(); err != nil {
		return nil, nil, err
	}

	return projectSelected, globalSelected, nil
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
