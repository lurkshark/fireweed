# NEW next wiring — 0080 blocked

0079's owner-approved native abandon/re-author succeeded: RUN-010–015 applied at
1ddea1c with all 191 full phase cases passing. The earlier 0079 failure and
recovery remain in [their report](next-0079-block.md). Applied-since-review is 1.
The first CLI slice is still incomplete because 0080 has not applied.

## Phase evidence

Author commit 6b5d461 scoped app wiring to next; evals ff154e5 reported 159 cases,
six genuine runtime failures and 153 passes. No interface/implementation edits
were needed for those fail-first results. The evals use real NEW binaries, Go
fixtures, separate bounded fake worker/auditor scripts and phase-state assertions.

The isolated Sol/high assertion auditor returned No findings for APP-074–079.
Full command-log inspection confirms only permitted snapshot reads, no outside
implementation/history reads, execution or writes. The snapshot was removed.
Native audit commit c79089f records honest manual coordinator advancement;
actual fixed-profile CLI logs remain under run
20261029T000645-7551044ce491f3dfa364881d.

## Concrete frozen-fixture blocker

The implementation worker's focused six-test run passed five cases, including
the full NEW next Go lifecycle with checked commits, SPEC/archive/provenance and
claim release. APP-078 alone fails. Its omit-new hook runs a genuinely failing
Go test, emits a passing JUnit report containing only ALP-001, then exits 1.
Existing checks return C11 for hook-exit/report disagreement before coverage.
The CLI correctly retains that wrapped C11 cause and adds corrective action,
while the frozen assertion requires C10 plus ALP-002. Changing product behavior
to report C10 here would misstate the actual dependency cause.

The coordinator missed this fixture contract error before freezing the evals.
No frozen test, interface or other module was changed during implementation.
Full implementation Verify/Done did not pass and application is not claimed.

Native Done committed only block state/Questions at 246fedf, saved attempt 1 at
refs/fireweed/attempts/0080-next-command/1 ->
df056aa4a36a772e4320e5a056a801fe8370ee13, restored step work and released the
claim. No phase/ref/counter/bound/history was manually rewritten. The tree is
clean; controller check passes with the unchanged embedded-recipe SPEC discovery
warning. No push after 510cc91 has occurred.

## Proposed scoped recovery — awaiting approval

/tmp/fireweed-0080-coverage-fixture.patch adds only exit 0 immediately after the
omit-new branch emits its passing report. That branch then provides an internally
consistent report with genuinely missing ALP-002 coverage, so the expected C10
clause can be tested. The normal fail-first branch, all assertions/IDs and public
interfaces remain unchanged. git apply --check passes; the patch is not applied.

SYSTEM requires human review whenever a change blocks. The proposed native
path is abandon/re-author of 0080, preserving its existing history/attempt and
the same issue/correction count, then restore only retained evals and apply this
one-line correction. Validate the diagnostic fixture through existing public
runner/workflow APIs before freezing it again, then run normal fail-first,
isolated audit and implementation. This is the first correction of this distinct
C11/C10 fixture issue. No new phase-write exception or bound change is proposed.

After application, build NEW and run /tmp/fireweed-next-cli-smoke.py run. Its
prepared, clean ordinary Go repository is /private/tmp/fireweed-next-cli-135aellc
with 30-second hook/agent limits; state is in
/tmp/fireweed-next-cli-smoke-state.json. The external smoke has not run. Clean it
up after the proof. Library success or the uncommitted candidate's focused CLI
passes do not establish final CLI acceptance. Bounded run follows that proof;
all later design/ACCEPTANCE gates and drill deferrals remain open.

## Owner-approved native recovery started

The owner approved native abandon/re-author and the one-line fixture correction.
Native abandon committed c014190 and reverted the previous phase work; all prior
history and attempt 1 remain retained. The same scope and issue are reauthored,
with no counter/bound/phase-write exception. Only retained evals and authored
spec text are supplied for reuse. Before freezing, a temporary public-API check
will validate the fixture returns C10/ALP-002; remove that harness before normal
controller verification. No saved/deleted implementation is an input.

## Corrected fixture and fresh fail-first passed; audit capacity unavailable

The fixed-profile eval worker restored the retained file and applied exactly
the approved exit 0 line. A temporary Go test using the public runner.Next API
and the actual fixture returned verify evals: C10: no coverage for ALP-002,
without C11; the harness was removed before verification. All six CLI specs
compile and fail at runtime against the still-unsupported next command. Native
evals commit 5e6d0bb reports 159 cases: six failures and 153 passes.

The new independent Sol/high auditor invocation terminated with exit 1 after
Selected model is at capacity. It made only a supplied-snapshot pwd/file-list
read and produced no assertion review or No findings result. No substitute model
or manual audit was used. The snapshot was removed; the actual audit/worker
logs remain under run 20261029T012539-cf09ff80f620517ce580ef2f.

The goal explicitly requires stopping on unavailable fixed-profile access.
Native release dropped this worktree's own claim without changing the committed
audit phase, IDs, run or counters. No step work is left uncommitted. This is an
external prerequisite, not a spec finding or a new native block commit. Resume
with native start/context for the existing audit phase and a regenerated isolated
snapshot when the same fixed profile is available/authorized to retry. Do not
repeat the passed eval phase or use saved/deleted implementation as an oracle.
Applied-since-review remains 1; no new push or implementation application occurred.
