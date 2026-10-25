//go:build windows

package evals

import "fmt"

func vanishedPinAncestor(pid int, root, label string) (*nestedOwned, error) {
	return nestedPin(pid, root)
}
func vanishedReaped(pid int, owned *nestedOwned) error {
	live, err := owned.alive()
	if err != nil || live {
		return fmt.Errorf("retained ancestor handle %d still active: %t/%v", pid, live, err)
	}
	return nil
}
