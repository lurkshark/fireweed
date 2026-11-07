# Residual fixes and verified controller handoff

The owner authorized the two proposed fixes through the normal audited workflow,
then a verified controller update and one final workflow drill. This evidence
closes the two corrections and clean-source handoff; it is not a native drill
result. The corrected subject/controller source is `ccd3d6b7732d613e80a027ba7969b65adc58bc0e`. Original baseline
79cf33a remains unchanged and all prior evidence/refs are preserved.

## Normal audited changes

[0181](../../changes/archive/0181-recorded-regenerate-preservation.md) clarifies
FLOW-081: deletion belongs to unstarted regeneration; an already recorded phase
retains disposable files. The existing preservation assertion moved under that ID,
while a separate FLOW-083 test reaches Context and asserts the exact regenerate
playbook and exclusion of ordinary impl. Initial deletion/zone assertions remain.
It is legitimately already_satisfied on the existing implementation. No runtime
source changed. Phase commits are f774410 (168 passing), a9ec21a (No findings,
FLOW-081), and fd7ae49 (650 passing, applied/archive).

The first audit submission was rejected by P01 for whitespace outside Audit.
That rejection was retained, not accepted or hidden. An empty Audit header was
prepared in the held audit document; quick verification reported expected P09
because it contained no Checked/No findings. A fresh independent auditor then
submitted only the accepted Audit body on the same run/pin. Actual audit agent
invocations: two; no configured counter, findings round or bound was reset.

[0182](../../changes/archive/0182-recovery-history-assertions.md) is an audited
eval_fix for RUN-027/RUN-028. Three phase commits are identified from public
history and parsed document phases, not subject words. Four fixture-agent OIDs
are recorded by the existing bounded fake agents and required to be absent from
final ancestry. Reset, staged/unstaged/untracked work, P05 warnings, ordered steps,
claim/run/archive/provenance, applied spec/code and separate invocation-failure
coverage remain. Runtime/spec/interface/hook/configuration stayed unchanged.
Phase commits are 6286a47 (159 passing), 334e948 (No findings, both IDs), and
ccd3d6b (482 passing, applied/archive).

Before the runner evals freeze, the first owning invocation was stopped and its
owned processes cleaned up while retaining the same run/claim. The editable test
was adjusted to parse phase YAML through public contracts instead of requiring
raw `phase: value` formatting. A bounded focused regression passed. A second
configured eval worker endorsed that correction on the same run; the normal
auditor and final worker passed. No frozen eval was edited, phase/counter reset,
manual phase commit or replacement change was used. Raw receipts and logs retain
these actual invocations and the pause. Runner quick verification took about
326 seconds in evals and 324 seconds in impl, exceeding the five-minute own-verify
target; this is a recorded cost limitation, not a claim that the target passed.
No time or coverage bound was raised and no unrelated cost cleanup was added.

## Clean-source validation and promotion

A fresh no-hardlinks clean clone at the exact corrected source built with Go
1.27.2 and `-trimpath` on darwin/arm64. Build info records that revision and
`vcs.modified=false`. [Ordered results](residual-fixes-controller-handoff/validation/results.json)
retain commands, fixed bounds, timestamps and exits for normal/race Go suites,
vet, check, hook doctor (`--no-agent`), explicit history delta and CI with hook
doctor enabled. The first ordinary run failed unchanged CHK-083 cancel because
Doctor did not return within its four-second testcase wait. The exact focused
replay passed without edits; the single ordinary retry and remaining gates passed.
The initial failed command/output is retained under validation/initial-normal-failure
and results.json records both outcomes. No bound, test or runtime was changed;
this is a non-reproduced timing failure, not a proven root-cause diagnosis.
The history delta is
`12c571cc53c7e03b051811477adc83e5518f157d..ccd3d6b7732d613e80a027ba7969b65adc58bc0e`; valid committed WF2 and CHK-034
prefixes are reused, not misrepresented as a fresh standalone full-range run.
History implementation is unchanged.

Compiled CLI fixtures exercise unstarted deletion, recorded-phase file/run
preservation and exact regenerate Context selection. Both fixture claims were
released. No fixture invoked agents or a native drill. Required phase reports
are retained under runs/. Existing B02 warnings for CHK-055 (16), APP-103 (17)
and APP-104 (19), against limit 15, remain visible; these tests were not changed.
Status records live unapproved 3 and exactly three new phase commits per archive,
with old counters intact. Protected author commits do not assert human merge approval.

Only after both lifecycles and validation exited with no active run/shared claim,
the primary was atomically replaced. Current SHA-256 is `20278504a8ea89fcd5f23d1de9a6ec849389fd55707054992c59967eec65e2b1`; retiring source
12c571c and its original bytes are preserved as `.fireweed/bin/fireweed-12c571c`.
All older recoveries remain unchanged. [Promotion receipt](residual-fixes-controller-handoff/promotion.json)
records the handoff and idle P02 check; idle checks are not product validation.
Owned validation repositories were removed only after fixture claim release;
[cleanup](residual-fixes-controller-handoff/validation-cleanup.json) is retained.

Outside tooling changes only expected pin/candidate identities. Exact worker
command/model/effort, SYSTEM input instructions, wrapper behavior, all isolation
and review gates, stop/drain/report windows, two-attempt cap and the existing
2026-11-10 22:51:05 UTC deadline remain unchanged. Earlier fake rehearsals remain
applicable to unchanged supervisor behavior. No command interception, process or
filesystem instrumentation, import inventory, implementation hints or expanded
worker inputs were added. No saved reconstruction was imported.

## Authorized next action

The owner explicitly authorized one downstream workflow drill on this corrected
subject with the new controller; no further confirmation is needed for this
specific sequence. Use a fresh outside root, review actual retained inputs before
release, full first trace/native feedback and fresh output before any remaining
configured release, then retain terminal native evidence/ref and clean owned
scratch. Stop afterward. A failure/gap does not authorize another correction or
drill. Native failure publication after pre-release rejection was observed only
in the superseded invocation; running-worker rejection/interruption reporting is
not established by that report or the fake rehearsals. No native drill or push
occurred during these two fixes and the handoff.
