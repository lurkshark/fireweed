# First next loop — 0079 recovery passed

Current result: reauthored 0079 applied at 1ddea1c with 191 full phase passes.
The isolated audit passed at 54f7f5c; fail-first evals committed 7ed35d4.
No claim remains. Runner library evidence is complete; NEW next CLI wiring
and proof remain pending. Historical evidence follows without retroactive edits.

## Historical block

0079-single-change-execution stopped at the required blocked-change checkpoint.
It is not applied. Native block commit is `00edbd3`; the checkout is clean and
no claim remains. Attempt 1 is retained at
`refs/fireweed/attempts/0079-single-change-execution/1` ->
`b7155214f07205488bdf531ba0c868448961bf8e`. Do not read its implementation as an
oracle; only approved native recovery may restore it.

## Completed evidence

The documentation refocus committed separately at 18da36d and authoring at
8eac2a3. All 29 inventory rows, command inventory and nine gates were preserved;
static checks, local links, code fences and whitespace passed. Newer correction
evidence and the full applied ledger were retained in HISTORY.

Evals/interface/stub commit `eb0598d` passed native fail-first requirements:
191 cases, 12 runtime failures and 179 passes. All new RUN-010–015 IDs fail first;
no compile/collection error supplies evidence. The fixed-profile CLI eval worker
reported 38 own cases, 26 passes and 12 failures in quick verification.

The separate fixed-profile Sol/high auditor returned No findings for all six
IDs. Complete command-log inspection found only permitted snapshot reads, no
implementation/history access, execution or writes. The snapshot was removed.
Native audit commit `0c8841a` records honest coordinator agent=manual provenance;
actual CLI logs remain under run `20261028T232748-8a17fd5f3261c7dfe4264076`.

## Observed blocker

The implementation worker's focused new-behavior suite reached applied, then
failed only the frozen assertion at next_test.go:353 requiring the third commit
subject to contain impl. Existing workflow.Done emits `feat(alpha): apply next`.
The design's State fields and Merge strategies identify commits from phase
state, not subjects. The worker preserved frozen evals/interfaces and other
modules; no implementation Done passed and no application is claimed.

The coordinator added the explicit worker block/Questions, then native Done
saved attempt 1, restored step work, committed only the change state/Questions,
and released the claim. No phase, counter, ref, controller bound or archived
history was edited manually. Controller check passes with the unchanged embedded
recipe SPEC discovery warning. Applied-since-review stays 0.

## Concrete proposed correction — not applied

`/tmp/fireweed-0079-phase-assertion.patch` replaces only the unsupported final
commit-subject assertion with an assertion that HEAD contains the archived
change at phase applied. It retains all phase-count, fail-first/pass, SPEC,
archive, provenance and claim-release assertions and all ID coverage. No
workflow behavior or public API change is proposed.

SYSTEM's blocked-change checkpoint requires human review now. The current new
IDs exist only in blocked 0079, so a separate same-module eval_fix cannot simply
run after it as an eligible independent change. A standard controller path is
native abandon of blocked 0079, then reauthor the same bounded scope with the
corrected assertion and rerun normal fail-first/audit/impl. This keeps history
and normal phase enforcement intact; it needs owner approval at this checkpoint.
Do not start a replacement to bypass a bound, alter phase history, reuse the
consumed phase-write exception, or push this blocked batch. The prior correction
batch through 510cc91 remains pushed and passed. NEW next/run remain unproved.

## Owner-approved native recovery

The owner approved the proposed native abandon/re-author and scoped assertion
correction. Native abandon committed e5059ae, reverting the former phase work
and deleting the blocked pending change. Historical phase/block commits and
attempt 1 remain retained. The same 0079 scope is reauthored; new late IDs are
left to the controller. Only retained API/types/evals and authored spec text were
saved for reuse; no deleted or saved implementation is copied or read. Normal
fail-first and isolated audit must run again. This is the first correction of
the same recorded issue, with applied count 0 and no bound/phase-write exception.

## Reauthored phase verification passed

Fresh fail-first 7ed35d4: 191 cases, 12 runtime failures and 179 passes.
Separate isolated Sol/high assertion audit 54f7f5c: No findings for RUN-010–015;
complete command-log inspection confirms permitted reads only, and the snapshot
was removed. Implementation 1ddea1c passes all 191 phase cases. Retained evals
differ from the former attempt only by the approved committed-phase assertion.
No saved/deleted implementation was supplied to reconstruction. Native Done
applied/archive/claim cleanup passed. Actual CLI logs remain under run
20261028T235723-4038c1c2cfedfafa7aa9c5d5, with honest manual coordinator phase
provenance. No push follows; the applied-since-review count is 1. Return to app
wiring and the actual NEW next fixture now.
