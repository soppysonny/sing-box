//go:build !(darwin || linux)

package CometBox

import "os"

func getTunnelName(fd int32) (string, error) {
	return "", os.ErrInvalid
}
