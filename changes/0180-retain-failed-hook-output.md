---
bases: {}
fixes:
    - CHK-034
module: checks
---

## Motivation

The owner approved the scoped CHK-034 fix after the workflow retry lost its
compiler diagnostics. A bounded public-API reproduction showed that the protected
Go hook emitted the compiler error and exited nonzero, but Evaluate returned
empty stdout/stderr and HookExit zero after rejecting its unmapped PackageFailure
JUnit case. CHK-034 already requires the hook process's stdout, stderr and exit
code in the returned report. Keep the existing contract and restore these fields
when post-launch hook/report processing fails, so native retry feedback and saved
reports can show the existing diagnostics.

## Interface

No signature, model, spec, Decision, hook, runner, wrapper or configuration change.
Only checks public evals and this change may change during evals. Implementation
is limited to checks internals and this change; preserve frozen evals/interfaces.

## Added

## Modified

## Removed

## Decisions

## Notes

This is one scoped normal fix, not eval_fix or already_satisfied. Add genuine
failing public checks.Evaluate assertions for captured stdout, stderr and nonzero
exit status on the demonstrated unmapped PackageFailure JUnit rejection. Include
representative malformed XML and missing-report failures, and one bounded hook
execution-error path after output has been emitted (cancellation or timeout).
Use deterministic local fake hooks with explicit output/exit markers; no agents,
network, real compiler dependency or process/filesystem instrumentation is needed
inside evals. Keep the evals bounded and under the existing cost limits. Preserve
all existing tests, assertions, IDs and immutable interfaces.

Require the existing C11 rejection causes to remain visible. Invalid/unmapped
JUnit must still yield no accepted testcase evidence and must not satisfy
coverage or genuine fail-first. Do not assign spec IDs to PackageFailure, accept
partial invalid reports, swallow hook errors, change error codes or weaken any
verification requirement. Successful parsing and existing exit/report agreement
behavior must remain unchanged. Avoid unrelated parsing/reporting refactors.

During implementation preserve captured process metadata immediately after the
hook returns, including when it returns an error; retain it on later report-read
or parse errors. Populate validated cases/time only when parsing succeeds. The
existing runner feedback/report consumers already render Report stdout/stderr;
this change does not edit them or the protected Go hook.

The outside owner uses the existing pinned controller through evals, independent
restricted audit and impl, with full phase verification and normal fail-first.
Workers use pinned context before each step, run pinned verify --quick, poll to
terminal and return. Never commit, invoke start/done/release/next/run/abandon,
change run/phase/audit counters/Provenance, or edit protected planning/tooling.
Use the prescribed command/model/effort and GOCACHE. Do not consult bootstrap or
rebuild-v1, reconstruct code or launch any native drill.

The independent auditor sees only its restricted context and edits Audit only.
A passing submission is No findings with Checked CHK-034. Findings must quote the
exact normalized CHK-034 clause and identify a missing asserting test. Keep the
existing audit rounds and retry bounds; never replace this change to reset them.

Preserve candidate 79cf33a7690082f28c06bf41a72719d282bc00fe, the original deadline,
configured bounds/counters, previous reports/proof/attempt refs and unrelated
worktrees. The controller pin stays unchanged for the entire active run/claim.
After this fix is applied/archived and all claims/processes exit, the owner has
separately authorized validation and handoff of a controller built from the fixed
clean source. That protected author handoff is outside worker steps. It must not
change the frozen subject candidate, add dependency-import inputs, apply a failed
reconstruction, launch another drill or push. Stop after the fixed controller
handoff with reviewed evidence; workflow proof remains open.
