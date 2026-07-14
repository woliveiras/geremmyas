package cli

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	geremmyas "github.com/woliveiras/geremmyas"
)

type contextStats struct {
	TopLevel         int
	Nested           int
	FrontmatterBytes int
	Managed          int
	Modified         int
	Unowned          int
}

func (s contextStats) total() int { return s.TopLevel + s.Nested }

func approximateTokens(bytes int) int {
	return (bytes + 3) / 4
}

func runContext(w io.Writer) error {
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	manifest, _, err := loadGlobalManifest()
	if err != nil {
		return err
	}
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	catalogStats, err := collectEmbeddedContextStats("project/.github/skills")
	if err != nil {
		return err
	}
	sources := []struct {
		name     string
		path     string
		manifest map[string]string
	}{
		{name: "project", path: filepath.Join(cwd, ".github", "skills")},
		{name: "global", path: filepath.Join(home, ".agents", "skills"), manifest: manifest.Files},
		{name: "codex-system", path: filepath.Join(home, ".codex", "skills", ".system")},
		{name: "codex-plugin-cache", path: filepath.Join(home, ".codex", "plugins", "cache")},
	}

	fmt.Fprintln(w, "Context usage (approximate; tokens = bytes / 4)")
	printContextStats(w, "catalog", catalogStats, "managed by geremmyas source")
	for _, source := range sources {
		stats, err := collectFilesystemContextStats(source.path, source.manifest)
		if err != nil {
			return err
		}
		note := "observed, not managed"
		if source.name == "global" {
			note = "ownership from global manifest"
		} else if source.name == "codex-plugin-cache" {
			note = "observed cache upper bound; host activation may be smaller"
		}
		printContextStats(w, source.name, stats, note)
	}

	fmt.Fprintln(w, "contracts:")
	printContractStats(w, "project", filepath.Join(cwd, "AGENTS.md"))
	printContractStats(w, "codex-global", filepath.Join(home, ".codex", "AGENTS.md"))
	return nil
}

func printContextStats(w io.Writer, name string, stats contextStats, note string) {
	fmt.Fprintf(w,
		"  %s: top-level=%d nested=%d total=%d frontmatter-bytes=%d approx-tokens=%d managed=%d modified=%d unowned=%d (%s)\n",
		name, stats.TopLevel, stats.Nested, stats.total(), stats.FrontmatterBytes,
		approximateTokens(stats.FrontmatterBytes), stats.Managed, stats.Modified,
		stats.Unowned, note)
}

func printContractStats(w io.Writer, name, path string) {
	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Fprintf(w, "  %s: missing\n", name)
			return
		}
		fmt.Fprintf(w, "  %s: unreadable (%v)\n", name, err)
		return
	}
	fmt.Fprintf(w, "  %s: words=%d bytes=%d approx-tokens=%d\n",
		name, len(strings.Fields(string(data))), len(data), approximateTokens(len(data)))
}

func collectFilesystemContextStats(root string, owned map[string]string) (contextStats, error) {
	stats := contextStats{}
	info, err := os.Lstat(root)
	if os.IsNotExist(err) {
		return stats, nil
	}
	if err != nil {
		return stats, err
	}
	if info.Mode()&os.ModeSymlink != 0 {
		return stats, nil
	}
	err = filepath.WalkDir(root, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.Type()&os.ModeSymlink != 0 {
			if d.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}
		if d.IsDir() || d.Name() != "SKILL.md" {
			return nil
		}
		rel, err := filepath.Rel(root, path)
		if err != nil {
			return err
		}
		if len(strings.Split(filepath.ToSlash(rel), "/")) == 2 {
			stats.TopLevel++
		} else {
			stats.Nested++
		}
		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		stats.FrontmatterBytes += frontmatterBytes(data)
		installedHash, isOwned := owned[filepath.Clean(path)]
		if isOwned {
			stats.Managed++
			if bytesSHA256(data) != installedHash {
				stats.Modified++
			}
		} else {
			stats.Unowned++
		}
		return nil
	})
	return stats, err
}

func collectEmbeddedContextStats(root string) (contextStats, error) {
	stats := contextStats{}
	err := fs.WalkDir(geremmyas.EmbeddedFiles, root, func(path string, d fs.DirEntry, err error) error {
		if err != nil || d.IsDir() || d.Name() != "SKILL.md" {
			return err
		}
		rel := strings.TrimPrefix(path, root+"/")
		if len(strings.Split(rel, "/")) == 2 {
			stats.TopLevel++
		} else {
			stats.Nested++
		}
		data, err := fs.ReadFile(geremmyas.EmbeddedFiles, path)
		if err != nil {
			return err
		}
		stats.FrontmatterBytes += frontmatterBytes(data)
		stats.Managed++
		return nil
	})
	return stats, err
}

func frontmatterBytes(data []byte) int {
	text := string(data)
	if !strings.HasPrefix(text, "---\n") {
		return 0
	}
	end := strings.Index(text[4:], "\n---")
	if end < 0 {
		return 0
	}
	return end
}
