---
bases: {}
fixes:
    - RUN-040
kind: eval_fix
module: runner
---

## Motivation

RUN-040 requires an earlier blocking cause and saved attempt ref to survive a
later time stop. Owner-authorized diagnosis reproduced candidate CI's failure:
the fake gamma agent's six-second sleep can finish before the unchanged
eight-second run budget, advancing to audit while the test expects held evals.
Keep that agent waiting for cancellation so the existing case observes the
intended interruption. Preserve every case, assertion, covered ID, runtime
implementation, public interface and configured bound.

## Interface

## Added

## Modified

## Removed

## Decisions

## Notes

The unchanged protected runner hook reproduced 132 cases with only RUN-040
failing at time_limit_test.go's held-phase assertion (audit instead of evals).
All blocking-cause assertions passed. Correct only the fake-agent wait in
Test_RUN_040_EarlierBlockRetainsCauseBeforeTimeStop; do not increase its deadline
or weaken any assertion. Use normal eval_fix verification and restricted audit.

