package primitives

// ErrorCode is implemented in the disposable layer during the implementation step.
func ErrorCode(err error) (code, subject string, ok bool) { return "", "", false }
