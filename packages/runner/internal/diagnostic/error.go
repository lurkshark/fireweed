package diagnostic

import (
	"errors"
	"fmt"

	"github.com/lurkshark/fireweed/packages/contracts"
)

// Errorf keeps the existing code prefix until callers render the structured code.
func Errorf(code, subject, format string, args ...any) error {
	return contracts.CodedError{
		Code:    code,
		Subject: subject,
		Cause:   fmt.Errorf(code+": "+format, args...),
	}
}

func Join(causes ...error) error {
	return PreserveCode(errors.Join(causes...), causes...)
}

// PreserveCode exposes the first coded cause on a multi-error, since ErrorCode
// follows only single-error wrappers. The multi-error still unwraps all causes.
func PreserveCode(err error, causes ...error) error {
	for _, cause := range causes {
		if code, subject, ok := contracts.ErrorCode(cause); ok {
			return contracts.CodedError{Code: code, Subject: subject, Cause: err}
		}
	}
	return err
}
