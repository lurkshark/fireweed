---
eval_fix:
    - CHK-081
module: checks
phase: impl
run: 20261112T153551-b4bf3cdbb2a2279abd81a47d
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

The four-second watchdog begins when the audit fixture signals readiness after
writing both streams. The 150 ms request timeout remains in effect. The eval
requires a C15 deadline-exceeded diagnostic, and the fake audit agent emits the
fixed reply if its five-second sleep completes. The test cancels and joins
Doctor before fixture cleanup on early exits.

The affected test passed ten runs under the race detector. Temporary Go overlay
controls confirmed that the timeout case rejects natural completion with a fixed
reply, an ordinary nonzero exit, and a zero exit with a missing reply. The
controller's `fireweed verify doctor-probe-timeout-eval` passed for checks.

The configured independent auditor failed to start in the restricted workspace
with `Operation not permitted` while initializing the Codex app-server client.
Independent audit remains pending.

## Provenance
- evals 2026-11-12T15:46:37.94541Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6.1-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=6a9c2d506363e3539cc66efd37bc6a06701249541dc4e4f01ed3aa6454bbb864 tests=130 failing=0 passing=130
- audit 2026-11-12T15:47:01.151221Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-astra -c model_reasoning_effort="medium" -s workspace-write {prompt} fireweed=0.1.0 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0
## Audit

No findings; checked CHK-081
