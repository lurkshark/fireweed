# Completed CHK-034 fix and controller handoff

The owner-authorized scoped fix is applied and archived in
[0180](../../changes/archive/0180-retain-failed-hook-output.md). Normal fail-first,
independent restricted audit and full implementation verification passed. Captured
hook stdout/stderr/exit now survive report-read, JUnit-parse and hook-execution
errors. Strict C11 rejection, fail-first and coverage enforcement remain intact.
No spec/interface/hook, runner or dependency-import input changes were made.

The separately built controller passed validation and replaced the primary pin
with no live run or shared claim. [PROGRESS](PROGRESS.md) records its fixed source
and hash; the retiring executable and all previous recovery copies are preserved.
Only outside pin metadata changed in the prepared wrapper/inspector. The
[handoff evidence](../../.fireweed/drills/chk034-controller-handoff.md) records
validation, its limits and cleanup. The queue is empty. Stop here for owner review.

The workflow subject remains exactly
`79cf33a7690082f28c06bf41a72719d282bc00fe`, distinct from the fixed controller
source. Its validated range remains
`be35552b706aa3000ff8f11306d129ff1d701b85..79cf33a7690082f28c06bf41a72719d282bc00fe`.
Preserve the original **2026-11-07 22:51:05 UTC** deadline, configured bounds,
existing archive counters, reports/proof/attempt refs and unrelated worktree.
No reconstructed code was applied, no further native drill ran and no push occurred.

The previous workflow retry remains failed. Its exact
[native report](../../.fireweed/drills/workflow.md), complete
[prior report](../../.fireweed/drills/workflow-before-retry-a79f34c.md) and failed
ref `refs/fireweed/drills/workflow/20261107T172136.182088000` remain unchanged.
Normal native failure publication was observed there; native publication after
interruption/rejection was not exercised. The CHK-034 diagnostic fixture is not a
workflow proof. WF4 and full product completion remain open.

Another native invocation or correction requires separate owner approval. The
[prepared retry handoff](workflow-drill/README.md) grants no invocation. Recheck
AGENTS, SYSTEM, PROGRESS, ACCEPTANCE, live status, queue, shared claims, processes,
pin and deadline before any newly approved work. Never switch controllers with a
live run/claim or resume another controller's claim; no automatic recovery fallback.
