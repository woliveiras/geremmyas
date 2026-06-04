package cli

import (
	"fmt"
	"strings"
)

type markdownFrontmatter struct {
	fields map[string]string
}

func parseMarkdownFrontmatter(content string) (markdownFrontmatter, string, error) {
	fm := markdownFrontmatter{fields: map[string]string{}}
	content = strings.TrimPrefix(content, "\ufeff")
	if !strings.HasPrefix(content, "---") {
		return fm, content, nil
	}

	rest := strings.TrimPrefix(content, "---")
	rest = strings.TrimLeft(rest, "\r\n")
	end := strings.Index(rest, "\n---")
	if end < 0 {
		return fm, content, fmt.Errorf("unclosed frontmatter")
	}

	block := rest[:end]
	body := strings.TrimPrefix(rest[end+len("\n---"):], "\n")
	for _, line := range strings.Split(block, "\n") {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		key, value, ok := strings.Cut(line, ":")
		if !ok {
			continue
		}
		key = strings.TrimSpace(key)
		value = strings.TrimSpace(value)
		value = strings.Trim(value, `"'`)
		fm.fields[key] = value
	}
	return fm, body, nil
}

func (f markdownFrontmatter) get(key string) string {
	return f.fields[key]
}
