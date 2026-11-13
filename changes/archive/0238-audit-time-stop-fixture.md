---
eval_fix:
    - RUN-048
    - RUN-049
module: runner
phase: applied
run: 20261112T213413-dfc069b359ab66e70c488a43
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

The time-stop fixture now starts the change, invokes the existing fake worker,
and verifies and commits evals with the ordinary fixture hook before the timed
Run resumes the held audit. Three focused macOS arm64 race runs passed after
the correction in 15.227 seconds.

Both `fireweed verify --quick audit-time-stop-fixture` and
`fireweed verify audit-time-stop-fixture` passed in the evals phase.
An extra auditor invocation from inside the eval worker's sandbox could not
initialize (`Operation not permitted`). The owning controller runs the required
restricted audit through its configured lifecycle after accepting evals.

## Provenance
- evals 2026-11-12T21:45:45.731133Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6.1-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=6a9c2d506363e3539cc66efd37bc6a06701249541dc4e4f01ed3aa6454bbb864 tests=170 failing=0 passing=170
- audit 2026-11-12T21:46:06.76796Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-astra -c model_reasoning_effort="medium" -s workspace-write {prompt} fireweed=0.1.0 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0
- impl 2026-11-12T21:58:08.397984Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6.1-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=5f856d6bcab308b9ff1a8a7c03b9d2547744112aba81dd30dea8f61ec8062b03 tests=529 failing=0 passing=529
## Audit

No findings; checked RUN-048, RUN-049
