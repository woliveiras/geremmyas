package dashboard

import (
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/fsnotify/fsnotify"
)

// WatchDirs debounces filesystem events and calls onChange.
func WatchDirs(dirs []string, debounce time.Duration, onChange func()) error {
	w, err := fsnotify.NewWatcher()
	if err != nil {
		return err
	}
	for _, dir := range dirs {
		if err := addWatchRecursive(w, dir); err != nil {
			_ = w.Close()
			return err
		}
	}
	go func() {
		defer w.Close()
		var timer *time.Timer
		trigger := func() {
			if timer != nil {
				timer.Stop()
			}
			timer = time.AfterFunc(debounce, onChange)
		}
		for {
			select {
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
	return nil
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
