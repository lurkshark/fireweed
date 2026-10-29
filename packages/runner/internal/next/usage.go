package next

import (
	"context"
	"errors"
	"fmt"
	"time"
)

type usagePolicy struct {
	wait   time.Duration
	budget time.Duration
}

func readUsagePolicy(config map[string]any) (usagePolicy, error) {
	wait, err := usageDuration(config["limit_wait"], "limit_wait", 30*time.Minute)
	if err != nil {
		return usagePolicy{}, err
	}
	budget, err := usageDuration(config["limit_budget"], "limit_budget", 4*time.Hour)
	if err != nil {
		return usagePolicy{}, err
	}
	if wait <= 0 {
		return usagePolicy{}, fmt.Errorf("C01: limit_wait must be positive")
	}
	if budget < 0 {
		return usagePolicy{}, fmt.Errorf("C01: limit_budget must not be negative")
	}
	return usagePolicy{wait: wait, budget: budget}, nil
}

func usageDuration(value any, name string, fallback time.Duration) (time.Duration, error) {
	if value == nil {
		return fallback, nil
	}
	text, ok := value.(string)
	if !ok {
		return 0, fmt.Errorf("C01: %s must be a duration string, got %T", name, value)
	}
	if text == "" {
		return fallback, nil
	}
	duration, err := time.ParseDuration(text)
	if err != nil {
		return 0, fmt.Errorf("C01: invalid %s %q: %w", name, text, err)
	}
	return duration, nil
}

// UsageBudget records only completed waits, excluding agent and verification time.
type UsageBudget struct {
	waited   time.Duration
	deadline time.Time
	stop     *TimeStopError
}

type TimeStopError struct{ Limit time.Duration }

func (e *TimeStopError) Error() string {
	return fmt.Sprintf("run time limit %s reached", e.Limit)
}

func (b *UsageBudget) SetTimeLimit(deadline time.Time, stop *TimeStopError) {
	b.deadline, b.stop = deadline, stop
}

func (b *UsageBudget) TimeStop(ctx context.Context) error {
	if b.stop != nil && context.Cause(ctx) == b.stop {
		return b.stop
	}
	return nil
}

func (b *UsageBudget) Wait(ctx context.Context, policy usagePolicy) error {
	if err := ctx.Err(); err != nil {
		if stop := b.TimeStop(ctx); stop != nil {
			return stop
		}
		return err
	}
	if b.waited > policy.budget || policy.wait > policy.budget-b.waited {
		return &UsageStopError{Completed: b.waited, Next: policy.wait, Budget: policy.budget}
	}
	if b.stop != nil && time.Until(b.deadline) <= policy.wait {
		return b.stop
	}
	timer := time.NewTimer(policy.wait)
	defer timer.Stop()
	select {
	case <-ctx.Done():
		if stop := b.TimeStop(ctx); stop != nil {
			return stop
		}
		return ctx.Err()
	case <-timer.C:
		b.waited += policy.wait
		return nil
	}
}

type UsageStopError struct {
	Completed time.Duration
	Next      time.Duration
	Budget    time.Duration
}

func (e *UsageStopError) Error() string {
	return fmt.Sprintf("usage limit: completed wait %s; next wait %s exceeds budget %s", e.Completed, e.Next, e.Budget)
}

// CombinedUsageFailure preserves an earlier block as the primary failure.
type CombinedUsageFailure struct {
	Blocked error
	Stop    error
}

func (e *CombinedUsageFailure) Error() string   { return fmt.Sprintf("%v; %v", e.Blocked, e.Stop) }
func (e *CombinedUsageFailure) Unwrap() []error { return []error{e.Blocked, e.Stop} }

func IsUsageStop(err error) bool {
	var combined *CombinedUsageFailure
	if errors.As(err, &combined) {
		return false
	}
	var stop *UsageStopError
	return errors.As(err, &stop)
}
