package cli

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestLintDescriptionMissingTrigger(t *testing.T) {
	violations := lintDescription("Do not use for temporary experiments.")
	assertLintViolationCodes(t, violations, lintViolationMissingTrigger)
}

func TestLintDescriptionMissingNegativeScope(t *testing.T) {
	violations := lintDescription("Use when validating a skill description.")
	assertLintViolationCodes(t, violations, lintViolationMissingNegative)
}

func TestLintDescriptionLengthBoundary(t *testing.T) {
	valid := paddedLintDescription(1024)
	if violations := lintDescription(valid); len(violations) != 0 {
		t.Fatalf("lintDescription(%d chars) returned violations: %+v", len(valid), violations)
	}

	tooLong := valid + "a"
	violations := lintDescription(tooLong)
	assertLintViolationCodes(t, violations, lintViolationDescriptionTooLong)
}

func TestLintDescriptionRejectsMarkup(t *testing.T) {
	violations := lintDescription("Use when reviewing <template> content. Do not use for production.")
	assertLintViolationCodes(t, violations, lintViolationDescriptionMarkup)
}

func TestLintNameMatchesDirectory(t *testing.T) {
	if violations := lintName("skill-authoring", "skill-authoring"); len(violations) != 0 {
		t.Fatalf("lintName returned violations for matching name: %+v", violations)
	}

	violations := lintName("skill-authoring", "skill-authoring-v2")
	assertLintViolationCodes(t, violations, lintViolationNameMismatch)

	violations = lintName("", "skill-authoring")
	assertLintViolationCodes(t, violations, lintViolationNameMismatch)
}

func TestLintBodyLengthBoundary(t *testing.T) {
	if violations := lintBody(bodyWithLines(500)); len(violations) != 0 {
		t.Fatalf("lintBody returned violations for 500 lines: %+v", violations)
	}

	violations := lintBody(bodyWithLines(501))
	assertLintViolationCodes(t, violations, lintViolationBodyTooLong)
}

func TestRunLintReportsViolations(t *testing.T) {
	root := withTempCwd(t)
	writeSkillFixture(t, root, "good-skill", `---
name: good-skill
description: "Use when working on documentation. Do not use for production."
---

# Good skill
`)
	writeSkillFixture(t, root, "bad-skill", `---
name: bad-skill
description: "Missing trigger and negative scope"
---

# Bad skill
`)

	var out strings.Builder
	if code := Run([]string{"lint"}, &out, &out); code == 0 {
		t.Fatalf("lint exit code = %d, want non-zero. output: %s", code, out.String())
	}
	output := out.String()
	if !strings.Contains(output, "lint: ") || !strings.Contains(output, "project/.github/skills/test/bad-skill/SKILL.md") {
		t.Fatalf("lint output missing report:\n%s", output)
	}
	if !strings.Contains(output, lintViolationMissingTrigger) || !strings.Contains(output, lintViolationMissingNegative) {
		t.Fatalf("lint output missing violation codes:\n%s", output)
	}
}

func TestRunLintSucceedsForValidTree(t *testing.T) {
	root := withTempCwd(t)
	writeSkillFixture(t, root, "good-skill", `---
name: good-skill
description: "Use when working on documentation. Do not use for production."
---

# Good skill
`)

	var out strings.Builder
	if code := Run([]string{"lint"}, &out, &out); code != 0 {
		t.Fatalf("lint exit code = %d, output: %s", code, out.String())
	}
	if !strings.Contains(out.String(), "lint: ok (1 skills checked)") {
		t.Fatalf("lint success output unexpected:\n%s", out.String())
	}
}

func TestRunLintReportsMissingSkillAndEmptyDescription(t *testing.T) {
	root := withTempCwd(t)
	writeSkillFixture(t, root, "empty-description", `---
name: empty-description
description: ""
---

# Empty description
`)
	missingDir := filepath.Join(root, "project/.github/skills/test/missing-skill")
	if err := os.MkdirAll(missingDir, 0o755); err != nil {
		t.Fatalf("MkdirAll returned error: %v", err)
	}

	var out strings.Builder
	if code := Run([]string{"lint"}, &out, &out); code == 0 {
		t.Fatalf("lint exit code = %d, want non-zero. output: %s", code, out.String())
	}
	output := out.String()
	if !strings.Contains(output, lintViolationMissingSkillFile) {
		t.Fatalf("lint output missing missing-file violation:\n%s", output)
	}
	if !strings.Contains(output, lintViolationMissingTrigger) || !strings.Contains(output, lintViolationMissingNegative) {
		t.Fatalf("lint output missing empty-description violations:\n%s", output)
	}
}

func assertLintViolationCodes(t *testing.T, violations []lintViolation, want ...string) {
	t.Helper()
	if len(violations) != len(want) {
		t.Fatalf("violations = %+v, want %d items", violations, len(want))
	}
	got := make([]string, 0, len(violations))
	for _, violation := range violations {
		got = append(got, violation.Code)
	}
	if strings.Join(got, ",") != strings.Join(want, ",") {
		t.Fatalf("violation codes = %v, want %v", got, want)
	}
}

func paddedLintDescription(total int) string {
	base := "Use when working on Codex target. Do not use for production. "
	if total < len(base) {
		panic("total too short for base description")
	}
	return base + strings.Repeat("a", total-len(base))
}

func bodyWithLines(lines int) string {
	if lines <= 0 {
		return ""
	}
	items := make([]string, lines)
	for i := range items {
		items[i] = "line"
	}
	return strings.Join(items, "\n")
}

func writeSkillFixture(t *testing.T, root, skillName, content string) string {
	t.Helper()
	path := filepath.Join(root, "project/.github/skills/test", skillName, "SKILL.md")
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		t.Fatalf("MkdirAll returned error: %v", err)
	}
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatalf("WriteFile returned error: %v", err)
	}
	return path
}
