package cli

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strconv"
	"strings"
)

const configFileName = "geremmyas.yml"

type Config struct {
	Version int
	Packs   []string
	Targets []string
}

func defaultConfig() Config {
	return Config{
		Version: 1,
		Packs:   []string{"core", "sdd"},
		Targets: defaultTargets(),
	}
}

func parseConfig(r io.Reader) (Config, error) {
	cfg := Config{Version: 1}
	section := ""
	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		if strings.HasPrefix(line, "version:") {
			raw := strings.TrimSpace(strings.TrimPrefix(line, "version:"))
			version, err := strconv.Atoi(raw)
			if err != nil {
				return Config{}, fmt.Errorf("invalid version %q", raw)
			}
			cfg.Version = version
			section = ""
			continue
		}

		if line == "packs:" {
			section = "packs"
			continue
		}
		if line == "targets:" {
			section = "targets"
			continue
		}

		if strings.HasPrefix(line, "- ") && (section == "packs" || section == "targets") {
			value := strings.TrimSpace(strings.TrimPrefix(line, "- "))
			value = strings.Trim(value, `"'`)
			if value == "" {
				continue
			}
			if section == "packs" {
				cfg.Packs = append(cfg.Packs, value)
			} else {
				cfg.Targets = append(cfg.Targets, value)
			}
			continue
		}

		return Config{}, fmt.Errorf("unsupported config line: %s", line)
	}

	if err := scanner.Err(); err != nil {
		return Config{}, err
	}
	if cfg.Version != 1 {
		return Config{}, fmt.Errorf("unsupported config version %d", cfg.Version)
	}
	if len(cfg.Packs) == 0 {
		return Config{}, fmt.Errorf("config must include at least one pack")
	}
	if err := validateTargets(cfg.Targets); err != nil {
		return Config{}, err
	}
	cfg.Targets = normalizeTargets(cfg.Targets)
	return cfg, nil
}

func formatConfig(cfg Config) string {
	packs := append([]string(nil), cfg.Packs...)
	packs = uniqueStrings(packs)
	sort.Strings(packs)

	targets := normalizeTargets(cfg.Targets)

	var b strings.Builder
	b.WriteString("version: 1\n")
	b.WriteString("packs:\n")
	for _, pack := range packs {
		b.WriteString("  - ")
		b.WriteString(pack)
		b.WriteByte('\n')
	}
	b.WriteString("targets:\n")
	for _, target := range targets {
		b.WriteString("  - ")
		b.WriteString(target)
		b.WriteByte('\n')
	}
	return b.String()
}

func uniqueStrings(values []string) []string {
	seen := map[string]bool{}
	out := make([]string, 0, len(values))
	for _, value := range values {
		value = strings.TrimSpace(value)
		if value == "" || seen[value] {
			continue
		}
		seen[value] = true
		out = append(out, value)
	}
	return out
}

func effectiveTargets(cfg Config, flagOverride []string) []string {
	if len(flagOverride) > 0 {
		return normalizeTargets(flagOverride)
	}
	return normalizeTargets(cfg.Targets)
}

// applyTargetsFlag sets cfg.Targets from targetsFlag when non-empty, then validates and normalizes.
func applyTargetsFlag(cfg *Config, targetsFlag string) error {
	if flagTargets := splitCSV(targetsFlag); len(flagTargets) > 0 {
		cfg.Targets = flagTargets
	}
	if err := validateTargets(cfg.Targets); err != nil {
		return err
	}
	cfg.Targets = normalizeTargets(cfg.Targets)
	return nil
}
