package lifecycle

import "errors"

type verificationFailure struct{ cause error }

func (e *verificationFailure) Error() string { return e.cause.Error() }
func (e *verificationFailure) Unwrap() error { return e.cause }

func IsVerificationFailure(err error) bool {
	var failure *verificationFailure
	return errors.As(err, &failure)
}
