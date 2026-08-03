package cli

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestRunContextSeparatesManagedAndExternalSkillRoots(t *testing.T) {
	home := t.TempDir()
	state := t.TempDir()
	project := t.TempDir()
	t.Setenv("HOME", home)
	t.Setenv("XDG_STATE_HOME", state)
	chdirForContextTest(t, project)

	var installOut strings.Builder
	if code := Run([]string{"global", "--targets", "codex", "base"}, &installOut, &installOut); code != 0 {
		t.Fatalf("global exit code = %d, output: %s", code, installOut.String())
	}
	writeContextSkill(t, filepath.Join(project, ".github", "skills", "project-skill", "SKILL.md"))
	writeContextSkill(t, filepath.Join(home, ".agents", "skills", "external", "SKILL.md"))
	writeContextSkill(t, filepath.Join(home, ".agents", "skills", "external", "references", "SKILL.md"))
	writeContextSkill(t, filepath.Join(home, ".codex", "skills", ".system", "system-skill", "SKILL.md"))
	writeContextSkill(t, filepath.Join(home, ".codex", "plugins", "cache", "plugin-a", "skills", "plugin-skill", "SKILL.md"))

	var out strings.Builder
	if code := Run([]string{"context"}, &out, &out); code != 0 {
		t.Fatalf("context exit code = %d, output: %s", code, out.String())
	}
	content := out.String()
	for _, label := range []string{"catalog", "project", "global", "codex-system", "codex-plugin-cache"} {
		if !strings.Contains(content, label) {
			t.Errorf("context output missing %q:\n%s", label, content)
		}
	}
	for _, field := range []string{"top-level=", "nested=", "managed=", "unowned=", "approx-tokens="} {
		if !strings.Contains(content, field) {
			t.Errorf("context output missing %q:\n%s", field, content)
		}
	}
	if !strings.Contains(content, "nested=1") {
		t.Fatalf("context output should report nested skill marker:\n%s", content)
	}
}

func TestRunContextJSONDistinguishesNoStateCodingAndBase(t *testing.T) {
	home := t.TempDir()
	t.Setenv("HOME", home)
	t.Setenv("XDG_STATE_HOME", t.TempDir())

	noState := readContextJSON(t, t.TempDir())
	if noState.Project.Exists {
		t.Fatalf("no-state project unexpectedly has a manifest: %+v", noState.Project)
	}

	codingRoot := materializeContextProject(t, "core,coding")
	coding := readContextJSON(t, codingRoot)
	if got := contextSourceByName(t, coding, "project-portable").Stats.TopLevel; got != 4 {
		t.Fatalf("coding top-level skills = %d, want 4", got)
	}

	baseRoot := materializeContextProject(t, "core,base")
	base := readContextJSON(t, baseRoot)
	if got := contextSourceByName(t, base, "project-portable").Stats.TopLevel; got != 7 {
		t.Fatalf("base top-level skills = %d, want 7", got)
	}
	if strings.Join(coding.Project.Packs, ",") == strings.Join(base.Project.Packs, ",") {
		t.Fatalf("coding and base selections are indistinguishable: coding=%v base=%v", coding.Project.Packs, base.Project.Packs)
	}
	if len(base.SkillCosts) != contextSourceByName(t, base, "catalog").Stats.TopLevel || len(base.PackCosts) != 3 {
		t.Fatalf("context JSON omitted per-skill or pack costs: skills=%d packs=%+v", len(base.SkillCosts), base.PackCosts)
	}
	for _, cost := range base.PackCosts {
		if cost.Name == "base" && (cost.Skills != 7 || cost.DiscoveryTokens == 0 || cost.BodyTokens == 0 || cost.SupportTokens == 0) {
			t.Fatalf("invalid base cost: %+v", cost)
		}
	}
}

func readContextJSON(t *testing.T, root string) contextReport {
	t.Helper()
	var out strings.Builder
	if code := Run([]string{"context", "--root", root, "--json"}, &out, &out); code != 0 {
		t.Fatalf("context JSON exit code = %d, output: %s", code, out.String())
	}
	var report contextReport
	if err := json.Unmarshal([]byte(out.String()), &report); err != nil {
		t.Fatalf("decode context JSON: %v\n%s", err, out.String())
	}
	if report.SchemaVersion != 1 || report.Command != "context" {
		t.Fatalf("unexpected context JSON envelope: version=%d command=%q", report.SchemaVersion, report.Command)
	}
	return report
}

func materializeContextProject(t *testing.T, packs string) string {
	t.Helper()
	root := t.TempDir()
	chdirForContextTest(t, root)
	var out strings.Builder
	if code := Run([]string{"init", "--packs", packs, "--targets", "codex"}, &out, &out); code != 0 {
		t.Fatalf("init %s: %s", packs, out.String())
	}
	if code := Run([]string{"sync"}, &out, &out); code != 0 {
		t.Fatalf("sync %s: %s", packs, out.String())
	}
	return root
}

func contextSourceByName(t *testing.T, report contextReport, name string) contextSource {
	t.Helper()
	for _, source := range report.Sources {
		if source.Name == name {
			return source
		}
	}
	t.Fatalf("context source %q missing: %+v", name, report.Sources)
	return contextSource{}
}

func TestRunContextIgnoresMissingRoots(t *testing.T) {
	home := t.TempDir()
	project := t.TempDir()
	t.Setenv("HOME", home)
	t.Setenv("XDG_STATE_HOME", t.TempDir())
	chdirForContextTest(t, project)

	var out strings.Builder
	if code := Run([]string{"context"}, &out, &out); code != 0 {
		t.Fatalf("context exit code = %d, output: %s", code, out.String())
	}
	if !strings.Contains(out.String(), "global") || !strings.Contains(out.String(), "top-level=0") {
		t.Fatalf("missing roots should report zero counts:\n%s", out.String())
	}
}

func writeContextSkill(t *testing.T, path string) {
	t.Helper()
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		t.Fatalf("create skill dir: %v", err)
	}
	content := "---\nname: example\ndescription: Use when testing context. Do not use otherwise.\n---\n\n# Example\n"
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatalf("write skill: %v", err)
	}
}

func chdirForContextTest(t *testing.T, dir string) {
	t.Helper()
	previous, err := os.Getwd()
	if err != nil {
		t.Fatalf("get working directory: %v", err)
	}
	if err := os.Chdir(dir); err != nil {
		t.Fatalf("change working directory: %v", err)
	}
	t.Cleanup(func() {
		if err := os.Chdir(previous); err != nil {
			t.Errorf("restore working directory: %v", err)
		}
	})
}
