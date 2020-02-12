// +build windows

package logs

import (
	"io"
	"os"
	"syscall"
)

func initTerminal(w io.Writer) {
	switch v := w.(type) {
	case *os.File:
		enableVirtualTerminalProcessing(syscall.Handle(v.Fd()), true)
	}
}

func checkIfTerminal(w io.Writer) bool {
	var ret bool
	switch v := w.(type) {
	case *os.File:
		var mode uint32
		err := syscall.GetConsoleMode(syscall.Handle(v.Fd()), &mode)
		ret = (err == nil)
	default:
		ret = false
	}
	if ret {
		initTerminal(w)
	}
	return ret
}

func stderr() io.Writer {
	return NewColorableStderr()
}
