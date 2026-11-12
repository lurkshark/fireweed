---
module: checks
eval_fix:
    - CHK-081
---

## Motivation

The owner requested a fix for the timeout eval that required a CI retry.
In GitHub Actions run 43190033480 (first attempt),
Test_CHK_081_DoctorPersistsFailedAndTimedOutProbeOutput/timeout failed its
four-second stopwatch while the same revision passed on retry. The stopwatch
starts before Doctor creates its scratch worktree and runs all hook canaries;
CHK-065 bounds each agent probe, not the entire Doctor invocation. CHK-081
requires preserving output from timed-out probes. Setup and hook-check cost
must not be mistaken for a probe timeout regression.

## Interface

Audited eval_fix only. Edit only the checks evals zone and this change.
No product behavior, public API, spec text, CI retries or timeout increases.

## Added

## Modified

## Removed

## Decisions

## Notes

Repair Test_CHK_081_DoctorPersistsFailedAndTimedOutProbeOutput in
packages/checks/evals/doctor_agent_logs_test.go. Use a fake-agent readiness
signal so the existing four-second watchdog starts at the selected probe,
excluding Doctor's setup and earlier canary work. Preserve the existing
150ms request timeout and bounded sleeping fake agent. Distinguish an actual
timeout from an ordinary nonzero exit or a missing fixed reply, and ensure
natural completion of the sleeping probe would fail the timeout case.
Preserve the nonzero-exit case, full stdout/stderr retention, final-20-line
diagnostic assertions, absolute log paths and caller/scratch cleanup checks.
Ensure asynchronous test work is joined or cancelled before fixture cleanup.
Do not broaden the edit to other tests without evidence of the same failure.

Run the affected test repeatedly under the race detector, plus the module's
normal fireweed verify and independent audit. Five unmodified local race
runs passed, consistent with the load-sensitive CI symptom; this is not a
claim that the original failure reproduces on every machine.

