//go:build !unix

package CometBox

import (
	"net"
)

func linkFlags(rawFlags uint32) net.Flags {
	panic("stub!")
}
