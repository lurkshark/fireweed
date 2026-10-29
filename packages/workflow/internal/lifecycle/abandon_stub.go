package lifecycle

import "fmt"

// Abandon is implemented during the change's implementation phase.
func Abandon(root, change string) (string, error) {
	return "", fmt.Errorf("abandon is not implemented")
}
