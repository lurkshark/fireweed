# Rebuild handoff

The owner-requested sequence is complete and stopped. The two scoped fixes are
applied and independently audited in [0181](../../changes/archive/0181-recorded-regenerate-preservation.md)
and [0182](../../changes/archive/0182-recovery-history-assertions.md). No runtime
source, public interface, hook or configuration correction was needed. The old
pin governed both normal changes; the
[verified corrected-source handoff](../../.fireweed/drills/residual-fixes-controller-handoff.md)
records phase verification, clean-source normal/race/vet/check/doctor/history/CI,
compiled CLI regressions, actual pause/rejection history and cost/timing limits.

Corrected subject/controller source is `ccd3d6b7732d613e80a027ba7969b65adc58bc0e`. Exactly one approved native
workflow invocation then used that new pin and prepared outside inspector.
It **failed native verification**. App passed; workflow and runner have remaining
failures. The focused FLOW-081/FLOW-083 and RUN-027/RUN-028 cases passed. There is
no passing independent workflow proof. WF4 and full completion remain open.

The [current workflow report](../../.fireweed/drills/workflow.md) retains the exact
native prefix followed by owner review. Its
[complete evidence and review](../../.fireweed/drills/workflow-corrected-controller-retry-20261107T232508/review.md)
records every case, actual assertion text, coverage, full first feedback and both
traces, worker gap limitations, durable streams/exact forwarding, retained failed
ref and cleanup. Gap claims are not approved corrections. Normal native failure
publication was observed; rejection/interruption/emergency reporting was not used.
Neither attempt reported a compiler error. No extra diagnostic replay followed.

Failed ref: `refs/fireweed/drills/workflow/20261107T232508.801610000`. It was retained locally; reconstructed code was not
applied. The previous complete workflow report remains
[byte-preserved](../../.fireweed/drills/workflow-before-corrected-controller-retry-3853f17.md).
Original `79cf33a7690082f28c06bf41a72719d282bc00fe`, all prior archives/reports/proof/
attempt refs, old phase counters, recoveries and unrelated worktree are preserved.
Final live counter/warning details are in evidence; no limits were raised or reset.
The valid checks proof and WF1/WF2 history prefixes remain valid.

**Stop for owner review. The approval is consumed.** No further correction,
diagnostic replay, drill, source application, controller update or push is
currently authorized. Follow [GOAL](GOAL.md) and [ACCEPTANCE](ACCEPTANCE.md).
The owning native controller, inspector and both workers exited; input/verifier/
caller scratch is cleaned up. The queue/shared claims are empty and no owned
normal/native/validation run remains active. Protected final evidence is committed
separately after terminal cleanup.

The owner added exactly 72 hours to the original 2026-11-07 22:51:05 UTC deadline.
The existing deadline remains **2026-11-10 22:51:05 UTC** (14:51:05 PST), with no
further extension. All configured bounds, independent audits and isolation gates
remain unchanged. This deadline grants no extra work after the completed invocation.

**Controller pin** `/Users/joel/Workspace/fireweed/.fireweed/bin/fireweed`, version
0.1.0, source `ccd3d6b7732d613e80a027ba7969b65adc58bc0e`, SHA-256
`20278504a8ea89fcd5f23d1de9a6ec849389fd55707054992c59967eec65e2b1`. Built with Go 1.27.2 and `-trimpath` from clean source on darwin/arm64.
Its owner-authorized verified handoff is recorded above; it stayed unchanged
throughout the native retry.

Recovery executables remain at `.fireweed/bin/fireweed-12c571c`,
`.fireweed/bin/fireweed-1936346`, `.fireweed/bin/fireweed-99a1d6d`,
`.fireweed/bin/fireweed-pre-ck-handoff` and `.fireweed/bin/fireweed-bootstrap`.
Use them only for explicit owner-directed recovery with no automatic fallback or
cross-controller resume. Previous push approvals are historical; no push is
currently authorized.
