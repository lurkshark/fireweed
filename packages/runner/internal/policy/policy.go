package policy

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/lurkshark/fireweed/packages/contracts"
)

func duration(config map[string]any, field string, fallback time.Duration, allowZero bool) (time.Duration, error) {
	raw, exists := config[field]
	if !exists {
		return fallback, nil
	}
	value, _ := raw.(string)
	parsed, err := time.ParseDuration(value)
	if err != nil || parsed < 0 || parsed == 0 && !allowZero {
		return 0, contracts.CodedError{Code: "C01", Subject: field, Cause: fmt.Errorf("invalid %s duration %v", field, raw)}
	}
	return parsed, nil
}

func Timeout(config map[string]any, requested time.Duration) (time.Duration, error) {
	if requested > 0 {
		return requested, nil
	}
	return duration(config, "timeout", 90*time.Minute, false)
}

func Attempts(config map[string]any) int {
	if attempts, ok := config["attempts"].(int); ok && attempts > 0 {
		return attempts
	}
	return 2
}

type Usage struct {
	wait   time.Duration
	budget time.Duration
}

func UsagePolicy(config map[string]any) (Usage, error) {
	wait, err := duration(config, "limit_wait", 30*time.Minute, false)
	if err != nil {
		return Usage{}, err
	}
	budget, err := duration(config, "limit_budget", 4*time.Hour, true)
	return Usage{wait: wait, budget: budget}, err
}

type usageStop struct {
	completed, next, budget time.Duration
}

func (s usageStop) Error() string {
	return fmt.Sprintf("usage limit stopped execution: completed wait %s, next wait %s, budget %s", s.completed, s.next, s.budget)
}

func IsUsageStop(err error) bool {
	var stop usageStop
	return errors.As(err, &stop) && !isCombined(err)
}

type timeStop struct{ limit time.Duration }

func (s timeStop) Error() string {
	return fmt.Sprintf("run reached its %s time limit", s.limit)
}

func IsTimeStop(err error) bool {
	var stop timeStop
	return errors.As(err, &stop) && !isCombined(err)
}

// Combined preserves both causes while marking the outcome as a failure.
type Combined struct{ Failure, Stop error }

func (e Combined) Error() string   { return fmt.Sprintf("%v\n%v", e.Failure, e.Stop) }
func (e Combined) Unwrap() []error { return []error{e.Failure, e.Stop} }

// WithSecondary keeps the primary diagnostic visible to contracts.ErrorCode
// while preserving both causes for errors.Is and errors.As.
func WithSecondary(primary, secondary error) error {
	if primary == nil {
		return secondary
	}
	if secondary == nil {
		return primary
	}
	var combined error = Combined{Failure: primary, Stop: secondary}
	if code, subject, ok := contracts.ErrorCode(primary); ok {
		combined = contracts.CodedError{Code: code, Subject: subject, Cause: combined}
	}
	return combined
}

func isCombined(err error) bool {
	var combined Combined
	return errors.As(err, &combined)
}

type Budget struct {
	waited   time.Duration
	deadline time.Time
	stop     error
}

func (b *Budget) Bound(ctx context.Context, limit time.Duration) (context.Context, context.CancelFunc) {
	if limit <= 0 {
		return ctx, func() {}
	}
	b.deadline = time.Now().Add(limit)
	b.stop = timeStop{limit}
	return context.WithDeadlineCause(ctx, b.deadline, b.stop)
}

func (b *Budget) Wait(ctx context.Context, usage Usage) error {
	if ctx.Err() != nil {
		return context.Cause(ctx)
	}
	if usage.wait > usage.budget-b.waited {
		return usageStop{b.waited, usage.wait, usage.budget}
	}
	if !b.deadline.IsZero() && time.Now().Add(usage.wait).After(b.deadline) {
		return b.stop
	}
	timer := time.NewTimer(usage.wait)
	defer timer.Stop()
	select {
	case <-ctx.Done():
		return context.Cause(ctx)
	case <-timer.C:
		b.waited += usage.wait
		return nil
	}
}
