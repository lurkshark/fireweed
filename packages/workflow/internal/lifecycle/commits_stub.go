package lifecycle

import "errors"

// Commits is implemented during the implementation phase.
func Commits(root, change string) ([]string, error) {
	return nil, errors.New("commits not implemented")
}
