# NEW next wiring — recovery and CLI proof passed

Current result: 0080 applied at a11d713 with 159 full phase cases passing, and
the clean-built NEW next Go fixture passed. Applied-since-review is 2. Earlier
block/capacity records below remain historical, without retroactive validation.

## Historical scope and phase record

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
with a configured 30-second agent timeout; state is in
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

## Authorized audit retry and implementation — PASS

The owner approved one retry of the same Sol/high audit and continuation on
PASS. That invocation returned No findings for APP-074–079. Full command-log
inspection confirms only supplied isolated snapshot reads; no implementation,
history, execution or writes were used. The snapshot was removed. Native audit
commit fff9dba records honest manual coordinator advancement. Fresh fixed-profile
implementation passed all six focused cases and quick verification; native full
Done applied/archived at a11d713 with 159 passes and released the claim. Evals and
interfaces stayed frozen. The earlier capacity invocation had no result.

## Actual NEW next Go CLI proof — PASS

CGO_ENABLED=0 build passed at clean source
a11d713e13922087f2172ab5665570a80db1ec0c, Go 1.27.2 darwin/arm64,
vcs.modified=false. An optional module-version stat-cache write outside the
sandbox was denied without failing the build. No fresh full-repository race/vet
or final acceptance claim is added by this milestone.

The actual /Users/joel/Workspace/fireweed/bin/fireweed next command exited 0 in
5.73s, stdout Applied changes/0001-one-value.md, empty stderr. It used an ordinary
Go-recipe repository, explicit Git identity, separate deterministic fake worker
and auditor scripts, and a configured 30-second agent timeout. The recipe's real
Go hook reported APP-001 failing before implementation (one case, exit 1), then
passing afterward (one case, exit 0). Saved case IDs/results were inspected.

Fixture baseline 3cff382d1ad58c2b721337726a95d867d23331e7; checked phase commits:
69c28fae0d595556364def3465cb6b91468660e2 (audit phase),
a4a48a5d32cc7456b813d7ad0751472fb03af985 (impl phase),
09f1195df4a14bdaa8153ad0938db468a331fe5a (applied archive). Product run
20261029T015624-3819560ec6e98a3d5b3524ab retains evals/audit/impl provenance,
SPEC application and the assigned APP-001. Exactly three agent calls used the
NEW executable as FIREWEED_BIN. No pending change, claim or dirty file remained.
The owned temporary repository was removed after these inspections.

Two smoke-helper issues were corrected without changing product source: the
fake auditor initially used an inline checked-ID field rejected with P09, so the
passing fixture uses the canonical separate-line form from the public evals;
then the report reader was corrected to the saved report's Cases/IDs/Result
field names. The final fresh smoke captures command exit/output before checking
all artifacts and passes. Logs/results remain in /tmp/fireweed-next-cli-*.json
and the reusable script /tmp/fireweed-next-cli-smoke.py. This proves the first
bounded fake-agent next lifecycle, not general audit-format flexibility, run,
retry/recovery, real-agent self-hosting or final acceptance. Bounded run follows.
