package dashboard

import (
	"encoding/json"
	"errors"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

const cacheDir = ".geremmyas-cache"
const cacheFile = "gitdates.json"

type gitDatesCache struct {
	LastCommit string             `json:"last_commit"`
	Specs      map[string]specGit `json:"specs"`
}

type specGit struct {
	CreatedAt     string `json:"created_at,omitempty"`
	ApprovedAt    string `json:"approved_at,omitempty"`
	ImplementedAt string `json:"implemented_at,omitempty"`
}

var statusLine = regexp.MustCompile(`^[+-].*status:\s*(.+)$`)

// GitOptions controls git date extraction.
type GitOptions struct {
	Root    string
	NoGit   bool
	NoCache bool
}

// ExtractGitDates populates data.Dates from git history and cache.
func ExtractGitDates(data *DashboardData, opts GitOptions) error {
	data.Metrics.GitAvailable = false
	if opts.NoGit {
		return nil
	}
	if _, err := exec.LookPath("git"); err != nil {
		return nil
	}
	cachePath := filepath.Join(opts.Root, cacheDir, cacheFile)
	cache, _ := readGitCache(cachePath)
	if opts.NoCache {
		cache = gitDatesCache{Specs: map[string]specGit{}}
	}
	if cache.Specs == nil {
		cache.Specs = map[string]specGit{}
	}

	git := func(args ...string) *exec.Cmd {
		return exec.Command("git", append([]string{"-C", opts.Root}, args...)...)
	}

	out, err := git("rev-parse", "HEAD").Output()
	if err != nil {
		return nil
	}
	head := strings.TrimSpace(string(out))

	pathspecs := specGitPathspecs()
	logArgs := append([]string{"log", "--all", "-p", "--format=COMMIT:%H %aI", "--"}, pathspecs...)
	if cache.LastCommit != "" && !opts.NoCache {
		logArgs = append([]string{"log", cache.LastCommit + "..HEAD", "-p", "--format=COMMIT:%H %aI", "--"}, pathspecs...)
	}
	logOut, err := git(logArgs...).Output()
	if err != nil && cache.LastCommit == "" {
		return nil
	}
	mergeGitLog(string(logOut), cache.Specs)

	createArgs := append([]string{"log", "--all", "--diff-filter=A", "--name-only",
		"--format=COMMIT:%H %aI", "--"}, pathspecs...)
	createOut, _ := git(createArgs...).Output()
	mergeCreationDates(string(createOut), cache.Specs)

	cache.LastCommit = head
	_ = writeGitCache(cachePath, cache)

	for key, sg := range cache.Specs {
		num := specNumFromPath(key)
		if num == 0 {
			continue
		}
		data.Dates[num] = SpecDates{
			Number:        num,
			CreatedAt:     sg.CreatedAt,
			ApprovedAt:    sg.ApprovedAt,
			ImplementedAt: sg.ImplementedAt,
		}
	}
	data.Metrics.GitAvailable = len(data.Dates) > 0
	return nil
}

func mergeGitLog(output string, specs map[string]specGit) {
	var currentFile, currentDate string
	lines := strings.Split(output, "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "COMMIT:") {
			parts := strings.Fields(strings.TrimPrefix(line, "COMMIT:"))
			if len(parts) >= 2 {
				currentDate = parts[1]
			}
			continue
		}
		if strings.HasPrefix(line, "diff --git") {
			if path, ok := extractSpecFileFromDiff(line); ok {
				currentFile = path
			}
			continue
		}
		m := statusLine.FindStringSubmatch(strings.TrimSpace(line))
		if m == nil || currentFile == "" {
			continue
		}
		status := strings.Trim(m[1], `"' `)
		entry := specs[currentFile]
		switch strings.ToLower(status) {
		case "approved":
			if entry.ApprovedAt == "" {
				entry.ApprovedAt = currentDate
			}
		case "implemented":
			if entry.ImplementedAt == "" {
				entry.ImplementedAt = currentDate
			}
		}
		specs[currentFile] = entry
	}
}

func mergeCreationDates(output string, specs map[string]specGit) {
	var currentDate string
	for _, line := range strings.Split(output, "\n") {
		if strings.HasPrefix(line, "COMMIT:") {
			parts := strings.Fields(strings.TrimPrefix(line, "COMMIT:"))
			if len(parts) >= 2 {
				currentDate = parts[1]
			}
			continue
		}
		if isSpecMarkdownGitPath(line) {
			entry := specs[line]
			if entry.CreatedAt == "" {
				entry.CreatedAt = currentDate
			}
			specs[line] = entry
		}
	}
}

func specGitPathspecs() []string {
	return []string{"specs/*/spec.md", "docs/specs/*/spec.md"}
}

func isSpecMarkdownGitPath(p string) bool {
	if !strings.HasSuffix(p, "/spec.md") {
		return false
	}
	return strings.HasPrefix(p, "specs/") || strings.HasPrefix(p, "docs/specs/")
}

func extractSpecFileFromDiff(line string) (string, bool) {
	for _, prefix := range []string{"docs/specs/", "specs/"} {
		i := strings.Index(line, prefix)
		if i < 0 {
			continue
		}
		rest := line[i:]
		end := strings.Index(rest, " ")
		if end > 0 {
			rest = rest[:end]
		}
		rest = strings.TrimPrefix(rest, "b/")
		rest = strings.TrimPrefix(rest, "a/")
		if isSpecMarkdownGitPath(rest) {
			return rest, true
		}
	}
	return "", false
}

func specNumFromPath(path string) int {
	base := filepath.Base(filepath.Dir(path))
	var n int
	for i := 0; i < len(base) && i < 4; i++ {
		if base[i] < '0' || base[i] > '9' {
			break
		}
		n = n*10 + int(base[i]-'0')
	}
	return n
}

func readGitCache(path string) (gitDatesCache, error) {
	raw, err := os.ReadFile(path)
	if err != nil {
		return gitDatesCache{Specs: map[string]specGit{}}, err
	}
	var c gitDatesCache
	if json.Unmarshal(raw, &c) != nil {
		return gitDatesCache{Specs: map[string]specGit{}}, errors.New("corrupt cache")
	}
	return c, nil
}

func writeGitCache(path string, c gitDatesCache) error {
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	raw, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, raw, 0o644)
}

func parseISODate(s string) (time.Time, bool) {
	s = strings.TrimSpace(s)
	if len(s) >= 10 {
		t, err := time.Parse("2006-01-02", s[:10])
		return t, err == nil
	}
	return time.Time{}, false
}
