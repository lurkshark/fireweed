package queue

import (
	"errors"

	"github.com/lurkshark/fireweed/packages/contracts"
)

// causeFailure keeps a rejection's static cause without changing its message.
type causeFailure struct {
	code string
	err  error
}

func (e *causeFailure) Error() string { return e.err.Error() }
func (e *causeFailure) Unwrap() error { return e.err }

func withCause(code string, err error) error { return &causeFailure{code: code, err: err} }

func errorCode(err error, fallback string) string {
	var failure *causeFailure
	if errors.As(err, &failure) {
		return failure.code
	}
	return fallback
}

func itemErrorCode(err error) string {
	if code, _, ok := contracts.ErrorCode(err); ok {
		return code
	}
	return "C03"
}
