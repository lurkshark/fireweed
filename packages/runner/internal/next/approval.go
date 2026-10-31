package next

import "strings"

// IsApprovalRequired is a deliberately incorrect eval-phase placeholder.
func IsApprovalRequired(err error) bool {
	return err != nil && !strings.Contains(err.Error(), "requires approval")
}
