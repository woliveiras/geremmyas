package cli

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestRunDashboardGeneratesSite(t *testing.T) {
	root := withTempCwd(t)
	specDir := filepath.Join(root, "specs", "0001-demo")
	if err := os.MkdirAll(specDir, 0o755); err != nil {
		t.Fatal(err)
	}
	specMD := `---
title: Demo spec
family: platform
phase: 0
status: Draft
---
# Demo
`
	if err := os.WriteFile(filepath.Join(specDir, "spec.md"), []byte(specMD), 0o644); err != nil {
		t.Fatal(err)
	}

	var out strings.Builder
	if code := Run([]string{"dashboard", "--no-git"}, &out, &out); code != 0 {
		t.Fatalf("exit %d: %s", code, out.String())
	}
	index := filepath.Join(root, ".geremmyas", "dashboard", "index.html")
	if _, err := os.Stat(index); err != nil {
		t.Fatalf("missing index.html: %v", err)
	}
	readme, err := os.ReadFile(filepath.Join(root, "specs", "README.md"))
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(readme), "geremmyas:generated") {
		t.Fatalf("readme missing generated marker: %s", readme)
	}
}
