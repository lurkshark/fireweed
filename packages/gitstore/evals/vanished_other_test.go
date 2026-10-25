//go:build !linux && !darwin && !windows

package evals

import "fmt"

func vanishedPinAncestor(pid int, root, label string) (*nestedOwned, error) {
	return nil, fmt.Errorf("native ownership observation unsupported")
}
func vanishedReaped(pid int, owned *nestedOwned) error {
	return fmt.Errorf("native ownership observation unsupported")
}
