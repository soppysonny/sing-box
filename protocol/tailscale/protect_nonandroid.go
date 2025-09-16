//go:build !android

package tailscale

import "github.com/sagernet/sing-box/experimental/CometBox/platform"

func setAndroidProtectFunc(platformInterface platform.Interface) {
}
