package primitives

import "github.com/lurkshark/fireweed/packages/contracts/model"

// ErrorCode returns the first coded error encountered while unwrapping err.
func ErrorCode(err error) (code, subject string, ok bool) {
	for err != nil {
		switch coded := err.(type) {
		case model.CodedError:
			return coded.Code, coded.Subject, true
		case *model.CodedError:
			if coded != nil {
				return coded.Code, coded.Subject, true
			}
		}
		wrapped, ok := err.(interface{ Unwrap() error })
		if !ok {
			break
		}
		err = wrapped.Unwrap()
	}
	return "", "", false
}
