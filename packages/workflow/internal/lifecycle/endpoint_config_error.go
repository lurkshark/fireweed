package lifecycle

import "errors"

type endpointConfigError struct{ cause error }

func (e *endpointConfigError) Error() string { return e.cause.Error() }
func (e *endpointConfigError) Unwrap() error { return e.cause }

func IsEndpointConfigError(err error) bool {
	var endpointErr *endpointConfigError
	return errors.As(err, &endpointErr)
}
