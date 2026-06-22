package cli

import (
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
