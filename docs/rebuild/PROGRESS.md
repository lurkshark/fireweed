# Rebuild handoff

The owner has approved exactly one new workflow retry on frozen candidate
`79cf33a7690082f28c06bf41a72719d282bc00fe`, using the fixed controller below and
the prepared outside inspector. The owner also extended the original deadline
by 72 hours to **2026-11-10 22:51:05 UTC** (14:51:05 PST). Keep all other bounds,
counters, inputs and isolation gates unchanged. Preserve prior reports/refs,
review actual retained inputs and every command/script/patch, retain the terminal
native evidence, clean up only owned scratch, commit a protected handoff and stop.
No correction, reconstructed-code import, further promotion or push is authorized.
This approval supersedes the stopped state below for this single retry only.
Follow [GOAL](GOAL.md).

The scoped CHK-034 fix and separately authorized controller handoff are complete.
[0180](../../changes/archive/0180-retain-failed-hook-output.md) passed normal
fail-first, independent restricted audit and full implementation verification.
Captured hook stdout/stderr/exit are retained on hook execution, report-read and
JUnit-parse errors, with strict rejection and evidence requirements unchanged.
The fixed controller passed normal/race Go tests, vet, checks, hook doctor,
explicit new history/CI range and a compiled-controller compiler-error fixture.
Existing valid history-prefix evidence was reused; the initial bounded full-range
probe timed out and is not claimed as a passing new full-range result. See the
[committed handoff evidence](../../.fireweed/drills/chk034-controller-handoff.md).

The primary pin below was replaced atomically after validation, with no live run
or shared claim. Only outside expected-pin metadata changed in the prepared
wrapper/inspector; no worker-input expansion or wrapper behavior change occurred.
The queue is empty, all owned temporary repositories were removed, and the
unrelated linked worktree and all prior evidence/refs remain intact. Existing
archive counters and configured bounds were preserved; the new archived fix has
its normal phase provenance. No reconstructed code, spec/interface/hook or runner
change was applied, no further native drill ran and no push occurred.

**Stop for owner review.** The frozen subject remains
`79cf33a7690082f28c06bf41a72719d282bc00fe`, distinct from the fixed controller source.
WF4 and full product completion remain open. Another native invocation or
correction requires separate approval. Follow [GOAL](GOAL.md).

The approved safer workflow retry is complete and **failed native verification**.
Exactly one native invocation used candidate
`79cf33a7690082f28c06bf41a72719d282bc00fe`, its two configured reconstruction
attempts, the unchanged pinned controller and the prepared external inspector.
Both workers completed with durable raw stdout/stderr and exact forwarding to
the native worker logs. No independence violation was observed in the reviewed
commands/patches. The native controller published a failed report and retained
reconstruction ref, then exited normally with status 1; no rejection or emergency
stop was needed. These observations establish normal failure publication only,
not native reporting after interruption/rejection.

The [current workflow report](../../.fireweed/drills/workflow.md) retains the exact
native result and reviewed outcome. Workflow, runner and app all reported
`PackageFailure`/C11 and missing-coverage C10 diagnostics; the published own and
dependent case sections are empty. There is no passing workflow proof. Static
review found a missing import in the fresh reconstruction, not an established
retained contract/eval defect. During that retry, no correction or extra
build/test was performed.
The failed reconstruction is retained at
`refs/fireweed/drills/workflow/20261107T172136.182088000`.

WF1/WF2 and the owner-approved follow-ups remain complete: the Current absence
contract is archived in [0178](../../changes/archive/0178-current-absence-contract.md),
and the FLOW-002 audit-zone assertion eval_fix in
[0179](../../changes/archive/0179-audit-zone-assertion.md). The candidate and validated
range remain
`be35552b706aa3000ff8f11306d129ff1d701b85..79cf33a7690082f28c06bf41a72719d282bc00fe`.
The complete earlier workflow report, including preparation, first native failure,
stopped-drill limits and fake rehearsals, is preserved byte-for-byte as
[historical workflow evidence](../../.fireweed/drills/workflow-before-retry-a79f34c.md).
The [checks report](../../.fireweed/drills/checks.md) and valid checks proof
`refs/fireweed/drills/checks/20261107T002339.020159000` remain unchanged.
All earlier reports/proof/attempt refs are preserved.

**Stop for owner review. WF4 and full product completion remain open.** The approval
for this single retry is consumed. Further drills, investigations requiring changes,
and product/spec/eval corrections need separate agreement. The external
[wrapper/inspector handoff](workflow-drill/README.md) remains available, but grants
no additional invocation. Follow the [ordered plan](ACCEPTANCE.md#planned-milestone--workflow-contracts-and-regeneration-proof-wf1wf4)
and [owner checkpoint](GOAL.md).

The original deadline was **2026-11-07 22:51:05 UTC**; the owner has added exactly
72 hours, making the current deadline **2026-11-10 22:51:05 UTC** (14:51:05 PST).
The completed checks budget began 2026-11-04 22:51:05 UTC; it and the earlier
stabilization budget are closed. The selected two-change follow-up is complete.
Preserve the deadline and configured bounds/counters; none was raised or reset.

At the previous native retry handoff, cleanup removed its input tree, verifier
worktree and isolated caller; its inspector, controller and workers had exited.
That evidence commit preceded the separately authorized CHK-034 fix and pin update
above. Normal native failure publication remains the only reporting behavior
established by that retry; the later diagnostics fixture establishes no new drill
result or proof.

**Controller pin** `/Users/joel/Workspace/fireweed/.fireweed/bin/fireweed`,
version 0.1.0, source `12c571cc53c7e03b051811477adc83e5518f157d`, SHA-256
`c9167d3878a1344719574d3ee1179d42295596f2f994c31c556ca25df8aaaada`.
Built with Go 1.27.2 and `-trimpath` from clean source on darwin/arm64. Its
verified owner-authorized handoff is recorded in the linked CHK-034 evidence.
The retiring controller is preserved as `.fireweed/bin/fireweed-1936346`.

Recovery executables remain at `.fireweed/bin/fireweed-99a1d6d`,
`.fireweed/bin/fireweed-pre-ck-handoff` and `.fireweed/bin/fireweed-bootstrap`.
Use them only for explicit owner-directed recovery under SYSTEM's controller
rules. Previous one-time push authorizations are historical; publication and
push require a new explicit owner instruction.
