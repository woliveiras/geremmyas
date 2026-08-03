//go:build windows

package cli

import (
	"fmt"
	"os"
	"path/filepath"
	"syscall"
	"unsafe"
)

const (
	lockfileFailImmediately = 0x00000001
	lockfileExclusiveLock   = 0x00000002
)

var (
	kernel32Lock     = syscall.NewLazyDLL("kernel32.dll")
	procLockFileEx   = kernel32Lock.NewProc("LockFileEx")
	procUnlockFileEx = kernel32Lock.NewProc("UnlockFileEx")
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
	overlapped := new(syscall.Overlapped)
	result, _, _ := procLockFileEx.Call(
		file.Fd(),
		lockfileFailImmediately|lockfileExclusiveLock,
		0,
		1,
		0,
		uintptr(unsafe.Pointer(overlapped)),
	)
	if result == 0 {
		_ = file.Close()
		return nil, fmt.Errorf("another global mutation is in progress (lock: %s)", lockPath)
	}
	return func() error {
		result, _, callErr := procUnlockFileEx.Call(
			file.Fd(), 0, 1, 0, uintptr(unsafe.Pointer(overlapped)),
		)
		closeErr := file.Close()
		if result == 0 {
			return callErr
		}
		return closeErr
	}, nil
}
