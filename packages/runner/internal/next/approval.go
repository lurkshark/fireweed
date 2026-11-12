package next

import (
	"errors"
	"fmt"
)

type approvalRequiredError struct {
	path string
}

func (e *approvalRequiredError) Error() string {
	return fmt.Sprintf("change %s requires approval on the local or shared base", e.path)
}

func IsApprovalRequired(err error) bool {
	var approval *approvalRequiredError
	return errors.As(err, &approval)
}
