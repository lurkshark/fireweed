package lifecycle

import "strings"

// IsEndpointConfigError is a fail-first evals stub.
func IsEndpointConfigError(err error) bool {
	return err != nil && strings.HasPrefix(err.Error(), "P01:")
}
