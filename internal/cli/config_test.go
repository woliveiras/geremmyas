package cli

import (
	"strings"
	"testing"
)

func TestParseConfig(t *testing.T) {
	cfg, err := parseConfig(strings.NewReader(`
version: 1
packs:
  - core
  - sdd
  - python-api
`))
	if err != nil {
		t.Fatalf("parseConfig returned error: %v", err)
	}
	if cfg.Version != 1 {
		t.Fatalf("Version = %d, want 1", cfg.Version)
	}
	want := []string{"core", "sdd", "python-api"}
	if strings.Join(cfg.Packs, ",") != strings.Join(want, ",") {
		t.Fatalf("Packs = %v, want %v", cfg.Packs, want)
	}
}

func TestParseConfigRejectsUnknownLine(t *testing.T) {
	_, err := parseConfig(strings.NewReader(`
version: 1
unknown: true
packs:
  - core
`))
	if err == nil {
		t.Fatal("parseConfig succeeded, want error")
	}
}

func TestFormatConfigSortsAndDeduplicates(t *testing.T) {
	got := formatConfig(Config{Version: 1, Packs: []string{"sdd", "core", "sdd"}})
	want := "version: 1\npacks:\n  - core\n  - sdd\n"
	if got != want {
		t.Fatalf("formatConfig = %q, want %q", got, want)
	}
}
