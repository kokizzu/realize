// +build !windows

package settings

import "syscall"

// Flimit defines the max number of watched files
func (s *Settings) Flimit() error {
	var rLimit syscall.Rlimit
	rLimit.Max = uint64(s.FileLimit)
	rLimit.Cur = uint64(s.FileLimit)

	if err := syscall.Setrlimit(syscall.RLIMIT_NOFILE, &rLimit); err != nil {
		return err
	}
	return nil
}
