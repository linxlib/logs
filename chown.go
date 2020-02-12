// +build !linux

package logs

import (
	"os"
)

func chown(_ string, _ os.FileInfo) error {
	return nil
}
