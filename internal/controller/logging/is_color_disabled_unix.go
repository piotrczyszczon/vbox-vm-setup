//go:build aixdarwin || dragonfly || freebsd || linux || netbsd || openbsd || solaris

package logging

func isColorDisabled() bool {
	return false
}
