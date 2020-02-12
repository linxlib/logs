// +build darwin freebsd linux netbsd openbsd


package logs

import (
	"io"
	"os"
	"syscall"
	"unsafe"
)

func isTerminal(fd uintptr) bool {
	var termios syscall.Termios
	_, _, err := syscall.Syscall6(syscall.SYS_IOCTL, fd, ioctlReadTermios, uintptr(unsafe.Pointer(&termios)), 0, 0, 0)
	return err == 0
}

func checkIfTerminal(w io.Writer) bool  {
	switch v := w.(type) {
	case *os.File:
		return isTerminal(v.Fd())
	default:
		return false
	}
}

func stderr() io.Writer {
	return os.Stderr
}
