package dashboard

import (
	"bytes"
	"html/template"

	"github.com/yuin/goldmark"
)

func renderMarkdown(src string) (template.HTML, error) {
	var buf bytes.Buffer
	if err := goldmark.Convert([]byte(src), &buf); err != nil {
		return "", err
	}
	return template.HTML(buf.String()), nil
}
