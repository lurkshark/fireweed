---
eval_fix:
    - RUN-048
    - RUN-049
module: runner
---

## Motivation

The macOS race check for PR #1 failed in
Test_RUN_048_RUN_049_PermittedAuditWorkSurvivesFailureAndStops/time_stop
on head ae659f281711021c075c02cff9f6cf098bb0fc4b (Actions run 43303387406,
job 131480172242). The three-second Run budget expired after the evals
commit but before the auditor started: the invocation log contained only
evals and audit.cwd was absent. The assertions then required partial Audit
text that the auditor never had a chance to write. RUN-048 and RUN-049
cover retaining permitted audit work and cleaning its workspace after an
invocation; startup timing must not prevent this eval from exercising them.

## Interface

Audited eval_fix only. Edit the runner evals zone and this change. No product
code, public API, spec text, CI policy or timeout increases.

## Added

## Modified

## Removed

## Decisions

## Notes

Limit the correction to the time-stop fixture in
packages/runner/evals/audit_workspace_test.go and any small local helper it
needs. Prepare the held audit phase through public workflow/runner APIs before
starting the timed runner.Run call. The existing fake worker can supply its
eval files and invocation record, and workflow.Done can verify and commit the
eval phase outside that timed call. Preserve genuine fixture evidence and
normal lifecycle transitions; never hand-edit phase, claim refs or provenance.

Keep For at three seconds, the request timeout at ten seconds, and the bounded
six-second fake-auditor sleep. Preserve the time-stop classification, zero
applied changes, retained partial Audit and other change contents, held audit
phase and claim, only the evals phase commit, both expected invocations, and
audit workspace cleanup. Leave the agent-failure and usage-stop cases intact.
Do not expand the change to other timing tests without evidence.

Run the affected time-stop case repeatedly with the race detector, followed
by normal fireweed verify and independent audit. Use verify --quick during
worker iteration; the owning controller's done performs full verification.
Record observed results without claiming every machine reproduces the original
CI failure. Stay within SYSTEM's verification and audit bounds.

Audit findings use one line each:
`- ID: "exact single-line substring of the spec" — missing assertion`.
A passing audit includes No findings and Checked IDs for RUN-048 and RUN-049.

Before correction, three focused local race runs passed in 11.294 seconds,
consistent with the CI failure depending on setup cost under load.

