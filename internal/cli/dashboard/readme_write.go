package dashboard

import (
	"os"
	"path/filepath"
)

// WriteSpecReadmes writes the auto-generated index to each existing SpecRoots directory.
func WriteSpecReadmes(root string, data DashboardData) ([]string, error) {
	content := GenerateReadme(data)
	var written []string
	for _, rel := range SpecRoots {
		dir := filepath.Join(root, rel)
		st, err := os.Stat(dir)
		if err != nil || !st.IsDir() {
			continue
		}
		path := filepath.Join(dir, "README.md")
		if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
			return written, err
		}
		written = append(written, filepath.ToSlash(rel)+"/README.md")
	}
	return written, nil
}
