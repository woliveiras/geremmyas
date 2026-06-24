package cli

import (
	"strings"
	"testing"
)

func TestRunVersionCommandPrintsDefaultVersion(t *testing.T) {
	oldVersion := Version
	Version = "dev"
	t.Cleanup(func() { Version = oldVersion })

	for _, args := range [][]string{{"version"}, {"--version"}} {
		t.Run(strings.Join(args, " "), func(t *testing.T) {
			var stdout, stderr strings.Builder
			if code := Run(args, &stdout, &stderr); code != 0 {
				t.Fatalf("Run(%v) exit code = %d, stderr: %s", args, code, stderr.String())
			}
			if got, want := stdout.String(), "geremmyas dev\n"; got != want {
				t.Fatalf("Run(%v) stdout = %q, want %q", args, got, want)
			}
			if stderr.Len() != 0 {
				t.Fatalf("Run(%v) stderr = %q, want empty", args, stderr.String())
			}
		})
	}
}

func TestRunVersionCommandPrintsInjectedVersion(t *testing.T) {
	oldVersion := Version
	Version = "v9.9.9"
	t.Cleanup(func() { Version = oldVersion })

	var stdout, stderr strings.Builder
	if code := Run([]string{"version"}, &stdout, &stderr); code != 0 {
		t.Fatalf("version exit code = %d, stderr: %s", code, stderr.String())
	}
	if got, want := stdout.String(), "geremmyas v9.9.9\n"; got != want {
		t.Fatalf("version stdout = %q, want %q", got, want)
	}
	if stderr.Len() != 0 {
		t.Fatalf("version stderr = %q, want empty", stderr.String())
	}
}

func TestRunHelpMentionsVersionCommand(t *testing.T) {
	var stdout, stderr strings.Builder
	if code := Run([]string{"--help"}, &stdout, &stderr); code != 0 {
		t.Fatalf("help exit code = %d, stderr: %s", code, stderr.String())
	}
	if !strings.Contains(stdout.String(), "  geremmyas version") {
		t.Fatalf("help output missing version command:\n%s", stdout.String())
	}
	if stderr.Len() != 0 {
		t.Fatalf("help stderr = %q, want empty", stderr.String())
	}
}
