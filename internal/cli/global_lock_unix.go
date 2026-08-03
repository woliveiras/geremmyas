//go:build linux || darwin

package cli

import (
	"fmt"
	"os"
	"path/filepath"
	"syscall"
)

func acquireGlobalMutationLock() (func() error, error) {
	manifestPath, err := globalManifestPath()
	if err != nil {
		return nil, err
	}
	stateDir := filepath.Dir(manifestPath)
	if err := os.MkdirAll(stateDir, 0o755); err != nil {
		return nil, err
	}
	lockPath := filepath.Join(stateDir, "global-manifest.lock")
	file, err := os.OpenFile(lockPath, os.O_CREATE|os.O_RDWR, 0o600)
	if err != nil {
		return nil, err
	}
	if err := syscall.Flock(int(file.Fd()), syscall.LOCK_EX|syscall.LOCK_NB); err != nil {
		_ = file.Close()
		return nil, fmt.Errorf("another global mutation is in progress (lock: %s)", lockPath)
	}
	return func() error {
		unlockErr := syscall.Flock(int(file.Fd()), syscall.LOCK_UN)
		closeErr := file.Close()
		if unlockErr != nil {
			return unlockErr
		}
		return closeErr
	}, nil
}
