package dashboard

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"
	"time"
)

func TestShouldRebuild(t *testing.T) {
	tests := []struct {
		path string
		want bool
	}{
		{"specs/0001-foo/spec.md", true},
		{"specs/README.md", false},
		{"specs/0001-foo/tasks.md", true},
		{"docs/prds/x.txt", false},
	}
	for _, tc := range tests {
		if got := shouldRebuild(tc.path); got != tc.want {
			t.Errorf("shouldRebuild(%q) = %v, want %v", tc.path, got, tc.want)
		}
	}
}

func TestWatchDirsStopEndsGoroutine(t *testing.T) {
	tmp := t.TempDir()
	specs := filepath.Join(tmp, "specs")
	if err := os.MkdirAll(specs, 0o755); err != nil {
		t.Fatal(err)
	}
	before := runtime.NumGoroutine()
	stop, err := WatchDirs([]string{specs}, 50*time.Millisecond, func() {})
	if err != nil {
		t.Fatal(err)
	}
	stop()
	stop() // idempotent
	time.Sleep(100 * time.Millisecond)
	after := runtime.NumGoroutine()
	if after > before+3 {
		t.Fatalf("goroutine leak suspected: before=%d after=%d", before, after)
	}
}
