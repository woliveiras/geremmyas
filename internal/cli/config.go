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
}

func defaultConfig() Config {
	return Config{
		Version: 1,
		Packs:   []string{"core", "sdd", "afk"},
	}
}

func parseConfig(r io.Reader) (Config, error) {
	cfg := Config{Version: 1}
	inPacks := false
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
			inPacks = false
			continue
		}

		if line == "packs:" {
			inPacks = true
			continue
		}

		if inPacks && strings.HasPrefix(line, "- ") {
			pack := strings.TrimSpace(strings.TrimPrefix(line, "- "))
			pack = strings.Trim(pack, `"'`)
			if pack != "" {
				cfg.Packs = append(cfg.Packs, pack)
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
	return cfg, nil
}

func formatConfig(cfg Config) string {
	packs := append([]string(nil), cfg.Packs...)
	packs = uniqueStrings(packs)
	sort.Strings(packs)

	var b strings.Builder
	b.WriteString("version: 1\n")
	b.WriteString("packs:\n")
	for _, pack := range packs {
		b.WriteString("  - ")
		b.WriteString(pack)
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
