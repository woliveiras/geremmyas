package dashboard

import (
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/fsnotify/fsnotify"
)

// WatchDirs debounces filesystem events and calls onChange.
// Call the returned stop function when watching is no longer needed (e.g. on
// server shutdown); it closes the watcher and ends the background goroutine.
func WatchDirs(dirs []string, debounce time.Duration, onChange func()) (stop func(), err error) {
	w, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, err
	}
	for _, dir := range dirs {
		if err := addWatchRecursive(w, dir); err != nil {
			_ = w.Close()
			return nil, err
		}
	}
	var (
		closeOnce sync.Once
		done      = make(chan struct{})
	)
	stop = func() {
		closeOnce.Do(func() {
			_ = w.Close()
			close(done)
		})
	}
	go func() {
		defer stop()
		var timer *time.Timer
		stopTimer := func() {
			if timer != nil {
				timer.Stop()
				timer = nil
			}
		}
		trigger := func() {
			stopTimer()
			timer = time.AfterFunc(debounce, onChange)
		}
		for {
			select {
			case <-done:
				stopTimer()
				return
			case ev, ok := <-w.Events:
				if !ok {
					return
				}
				if ev.Op&(fsnotify.Write|fsnotify.Create|fsnotify.Remove|fsnotify.Rename) == 0 {
					continue
				}
				if !shouldRebuild(ev.Name) {
					continue
				}
				trigger()
			case _, ok := <-w.Errors:
				if !ok {
					return
				}
			}
		}
	}()
	return stop, nil
}

func addWatchRecursive(w *fsnotify.Watcher, root string) error {
	if _, err := os.Stat(root); os.IsNotExist(err) {
		return nil
	}
	return filepath.WalkDir(root, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return w.Add(path)
		}
		return nil
	})
}

func shouldRebuild(path string) bool {
	base := filepath.Base(path)
	if base == "README.md" {
		return false
	}
	return strings.HasSuffix(path, ".md")
}
