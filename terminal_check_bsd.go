// +build darwin freebsd netbsd openbsd

package logs

import (
	"syscall"
)

const ioctlReadTermios = syscall.TIOCGETA

