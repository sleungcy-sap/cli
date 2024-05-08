//go:build !windows
// +build !windows

package helpers

func PluginPlatform() string {
	return "linux64"
}
