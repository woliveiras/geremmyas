package cli

import "strings"

type lintViolation struct {
	Code    string
	Message string
}

const (
	lintViolationMissingTrigger     = "missing-use-when"
	lintViolationMissingNegative    = "missing-negative-scope"
	lintViolationDescriptionTooLong = "description-too-long"
	lintViolationDescriptionMarkup  = "description-markup"
	lintViolationNameMismatch       = "name-mismatch"
	lintViolationBodyTooLong        = "body-too-long"
	maxSkillDescriptionLength       = 1024
	maxSkillBodyLines               = 500
)

func lintDescription(description string) []lintViolation {
	violations := []lintViolation{}
	if !hasUseWhenPhrase(description) {
		violations = append(violations, lintViolation{
			Code:    lintViolationMissingTrigger,
			Message: "description must contain a use when trigger phrase",
		})
	}
	if !hasNegativeScopePhrase(description) {
		violations = append(violations, lintViolation{
			Code:    lintViolationMissingNegative,
			Message: "description must contain a negative-scope phrase",
		})
	}
	if len(description) > maxSkillDescriptionLength {
		violations = append(violations, lintViolation{
			Code:    lintViolationDescriptionTooLong,
			Message: "description must be at most 1024 characters",
		})
	}
	if strings.ContainsAny(description, "<>") {
		violations = append(violations, lintViolation{
			Code:    lintViolationDescriptionMarkup,
			Message: "description must not contain angle-bracket markup",
		})
	}
	return violations
}

func lintName(name, directory string) []lintViolation {
	if name == directory {
		return nil
	}
	return []lintViolation{{
		Code:    lintViolationNameMismatch,
		Message: "skill name must match directory name",
	}}
}

func lintBody(body string) []lintViolation {
	if countBodyLines(body) <= maxSkillBodyLines {
		return nil
	}
	return []lintViolation{{
		Code:    lintViolationBodyTooLong,
		Message: "skill body must be at most 500 lines",
	}}
}

func countBodyLines(body string) int {
	if body == "" {
		return 0
	}
	body = strings.TrimRight(body, "\n")
	if body == "" {
		return 0
	}
	return strings.Count(body, "\n") + 1
}

func hasUseWhenPhrase(description string) bool {
	return strings.Contains(strings.ToLower(description), "use when")
}

func hasNegativeScopePhrase(description string) bool {
	lower := strings.ToLower(description)
	phrases := []string{"do not use", "don't use", "not for"}
	for _, phrase := range phrases {
		if strings.Contains(lower, phrase) {
			return true
		}
	}
	return false
}
