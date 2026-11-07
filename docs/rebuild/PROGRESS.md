# Rebuild handoff

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
retained contract/eval defect. No correction or extra build/test was performed.
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

The original deadline remains **2026-11-07 22:51:05 UTC** (14:51:05 PST).
The completed checks budget began 2026-11-04 22:51:05 UTC; it and the earlier
stabilization budget are closed. The selected two-change follow-up is complete.
Preserve the deadline and configured bounds/counters; none was raised or reset.

At handoff, native cleanup removed the input tree and verifier worktree. The
outside inspector, controller and both workers exited. The owned isolated caller
was removed after its failed native ref and evidence were retained. Main was clean
before protected evidence authoring; no live run/shared claim or queued change
remains. The live unapproved count is zero. Archive phase counters, configuration,
controller/recovery executables, all prior refs and the unrelated linked worktree
are preserved. No reconstructed source was applied, no product/spec/eval/hook or
controller change was authored, and no promotion or push occurred. Protected
evidence is committed separately after terminal stop.

**Controller pin** `/Users/joel/Workspace/fireweed/.fireweed/bin/fireweed`,
version 0.1.0, source `193634686aa6a1c4eb32e69df1c85e1d5ce48bd7`, SHA-256
`23e87b00217ab1a6990352393311bd150b7d60d3a4c138b75e285f518fba117f`.
Built with Go 1.27.2 and `-trimpath` from clean source on darwin/arm64. Its
verified owner-authorized handoff is recorded in commit `200711c`.

Recovery executables remain at `.fireweed/bin/fireweed-99a1d6d`,
`.fireweed/bin/fireweed-pre-ck-handoff` and `.fireweed/bin/fireweed-bootstrap`.
Use them only for explicit owner-directed recovery under SYSTEM's controller
rules. Previous one-time push authorizations are historical; publication and
push require a new explicit owner instruction.
