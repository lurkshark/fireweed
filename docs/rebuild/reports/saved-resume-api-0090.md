# Saved-attempt API 0090 — archived dependency

FLOW-073–076 add the thin StartFromAttempt facade with unchanged existing
signatures/model. Its disposable lifecycle implementation resolves the canonical
slug and requested ref, checks invalid/missing attempts before claiming, uses
ordinary Start preconditions, and restores only saved step paths and permitted
worker body sections. Current author/runtime/provenance and later independent
commits remain intact; impl retains committed audit and audit resume loads saved
Audit body. HEAD, run, phase and assigned IDs are retained.

Native phases: initial evals 8cbcac8, retained findings d6cb397, corrected evals
9affb4c, passing scoped audit 2d2c483, impl/archive da16b68. Initial 295 cases
had eight genuine failures/287 passes; corrected 297 had ten failures/287 passes;
impl passed all 297. The implementation worker's single workflow Go JSON check
passed all 80 tests in 68.760s. No frozen tests/interfaces changed during impl.

Separate isolated Sol/high audit identified missing requested-attempt-2 selection
(FLOW-074) and audit-phase Audit restoration (FLOW-075). Exact physical-source
quotes were natively retained, returning to evals in round 1 with the same run
and IDs. A scoped real worker added those two cases; the isolated re-audit passed
only FLOW-074/075. FLOW-073/076 stayed passed. Both audit command logs show only
supplied snapshot reads, no SYSTEM/implementation/history/outside reads, code
execution or writes. Both snapshots were removed. No phase-write exception,
round-bound change, coverage weakening or drill occurred.

Actual worker/audit logs are under run
20261029T085817-dd0bc520b24864b5d616184a; coordinator provenance is manual.
The source archive is da16b680e9a1506c0587e0e373bcccd0ea4fbe08. No pending change
or claim remained at application. Two changes have applied since checkpoint;
last reviewed push remains 82ef71d. NEW still uses the valid bf27309 build until
the app caller is ready; it was not rebuilt merely to update the revision label.

This API dependency does not claim CLI resume completion. Immediately wire
start --from-attempt and use the retained Go fixture described in
[PROGRESS](../PROGRESS.md) and [the run proof](independent-run-0089.md).
The prepared /tmp/fireweed-saved-resume-smoke.py has only been syntax-checked,
never run: it will unblock Alpha, start from attempt 1, inspect restored work and
retained state, correct the code, finish through NEW done and clean the fixture.
Full design/final acceptance, agent-commit undo and time/usage remain open.
