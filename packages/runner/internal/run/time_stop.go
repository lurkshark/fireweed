package run

import (
	"errors"
	"fmt"

	"github.com/lurkshark/fireweed/packages/runner/internal/next"
)

type combinedTimeFailure struct {
	blocked error
	stop    error
}

func (e *combinedTimeFailure) Error() string   { return fmt.Sprintf("%v; %v", e.blocked, e.stop) }
func (e *combinedTimeFailure) Unwrap() []error { return []error{e.blocked, e.stop} }

func IsTimeStop(err error) bool {
	var combined *combinedTimeFailure
	if errors.As(err, &combined) {
		return false
	}
	var stop *next.TimeStopError
	return errors.As(err, &stop)
}
