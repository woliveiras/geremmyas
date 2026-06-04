package cli

import (
	"os"
	"path/filepath"
)

type installScope int

const (
	scopeProject installScope = iota
	scopeGlobal
)

type globalPaths struct {
	home           string
	cursorRules    string
	cursorHooks    string
	cursorHooksJSON string
	claudeMD       string
	opencodeAgents string
}

func globalInstallPaths() (globalPaths, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return globalPaths{}, err
	}
	return globalPaths{
		home:            home,
		cursorRules:     filepath.Join(home, ".cursor", "rules"),
		cursorHooks:     filepath.Join(home, ".cursor", "hooks"),
		cursorHooksJSON: filepath.Join(home, ".cursor", "hooks.json"),
		claudeMD:        filepath.Join(home, ".claude", "CLAUDE.md"),
		opencodeAgents:  filepath.Join(home, ".config", "opencode", "AGENTS.md"),
	}, nil
}

func globalSkillPath(skillName string) string {
	return filepath.Join("~", ".agents", "skills", skillName, "SKILL.md")
}
