// +build linux
package logs

import "syscall"

const ioctlReadTermios = syscall.TCGETS
