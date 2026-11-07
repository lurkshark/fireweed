# Rebuild handoff

The two owner-approved residual fixes are complete in normal audited archives
[0181](../../changes/archive/0181-recorded-regenerate-preservation.md) and
[0182](../../changes/archive/0182-recovery-history-assertions.md). The reviewed
[fix/controller evidence](../../.fireweed/drills/residual-fixes-controller-handoff.md)
records preserved obligations, passing phase verification, actual interruptions/
audit rejection and cost limitations. No runtime source, interface, hook or
configuration correction was needed. The old pin governed both changes.

Corrected subject and newly validated controller source:
`ccd3d6b7732d613e80a027ba7969b65adc58bc0e`. Normal/race Go suites, vet, check, hook doctor, explicit history/CI delta
and compiled regenerate CLI behavior passed in clean source. Reuse the valid
WF2/CHK-034 history prefix; the evidence identifies the exact new delta. Original
`79cf33a7690082f28c06bf41a72719d282bc00fe` remains preserved as baseline.
The controller was replaced atomically only after all normal claims/runs exited.
No validation fixture, normal workflow run or shared claim remains active.

The owner has already authorized exactly one final workflow drill with this
corrected subject and new pin. Follow [GOAL](GOAL.md), the prepared outside
inspector and all existing isolation/review/output/stop gates. Retain the terminal
native report/ref and complete review, clean owned scratch, commit a protected
stopped handoff and stop afterward. No additional approval prompt is required for
this specific sequence. No further correction, extra drill, source application
or push is authorized. WF4/full completion remain open pending actual review.

The current [workflow report](../../.fireweed/drills/workflow.md) is still the
previous failed native retry on original 79cf33a. Its
[complete evidence](../../.fireweed/drills/workflow-fixed-controller-retry-20261107T202952/review.md)
and failed ref remain unchanged; it is not a result for the corrected subject.
The [residual investigation](../../.fireweed/drills/workflow-residual-investigation-20261107/proposal.md)
and [superseded pre-release stop](../../.fireweed/drills/workflow-superseded-repeat-20261107T213809/review.md)
are retained separately. That stop released no worker and establishes no
running-worker interruption/drain behavior. Earlier wrapper fake rehearsals
remain valid for unchanged behavior, not evidence of native reporting.

The owner added exactly 72 hours to the original 2026-11-07 22:51:05 UTC deadline.
The current deadline stays **2026-11-10 22:51:05 UTC** (14:51:05 PST), with no further
extension. All configured bounds/counters and independent audits remain intact.
Prior reports/archives/proof/attempt refs, recovery executables and the unrelated
linked worktree are preserved. The valid checks proof and original WF1/WF2 prefix
remain valid. Existing warning/counter details are in committed handoff evidence.

**Controller pin** `/Users/joel/Workspace/fireweed/.fireweed/bin/fireweed`,
version 0.1.0, source `ccd3d6b7732d613e80a027ba7969b65adc58bc0e`, SHA-256
`20278504a8ea89fcd5f23d1de9a6ec849389fd55707054992c59967eec65e2b1`. Built with Go 1.27.2 and `-trimpath` from clean source on darwin/arm64.
Its owner-authorized verified handoff is recorded above.

Recovery executables remain at `.fireweed/bin/fireweed-12c571c`,
`.fireweed/bin/fireweed-1936346`, `.fireweed/bin/fireweed-99a1d6d`,
`.fireweed/bin/fireweed-pre-ck-handoff` and `.fireweed/bin/fireweed-bootstrap`.
Use them only for explicit owner-directed recovery with no automatic fallback or
cross-controller resume. Previous push approvals are historical; no push is
currently authorized.
