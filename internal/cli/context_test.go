package cli

import (
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
	if code := Run([]string{"global", "--targets", "codex", "sdd"}, &installOut, &installOut); code != 0 {
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
