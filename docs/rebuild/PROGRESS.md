# Rebuild handoff

The owner-approved workflow follow-ups are complete: the Current absence
contract is archived in [0178](../../changes/archive/0178-current-absence-contract.md),
and the FLOW-002 audit-zone assertion eval_fix in
[0179](../../changes/archive/0179-audit-zone-assertion.md). The focused review
questions are resolved with existing runtime behavior preserved. Independent
audits and normal full phase verification passed; no runtime/interface change
or reconstructed source import was needed.

Updated WF2 preparation identifies candidate
`79cf33a7690082f28c06bf41a72719d282bc00fe` over the exact range
`be35552b706aa3000ff8f11306d129ff1d701b85..79cf33a7690082f28c06bf41a72719d282bc00fe`.
The appended preparation in the [checks report](../../.fireweed/drills/checks.md)
and [workflow report](../../.fireweed/drills/workflow.md) records candidate
validation, the rejected audit submission and bounded independent retry, and
initial timing-test failures with unchanged passing reruns.

The owner approved one native workflow invocation on this exact candidate under
the unchanged pin, deadline, bounds and live isolation gates. The first native
verification failed in workflow and both transitive dependents, runner and app.
After complete first-trace/native-feedback review and repeated live isolation
checks, the second configured worker was released. It disclosed creating and
removing a temporary `go.mod`, contrary to the explicit reconstruction write
restrictions. The inspector stopped the invocation at the independence checkpoint.
The controller exited by interruption before publishing a new native report,
complete case report or proof ref. No passing independent proof is claimed.
Available evidence and its limits are appended to the workflow report; its earlier
native result and exact preparation remain intact. WF4 and full product completion
remain open. Follow the
[ordered plan](ACCEPTANCE.md#planned-milestone--workflow-contracts-and-regeneration-proof-wf1wf4)
and [owner checkpoint](GOAL.md). Stop before another drill or any correction.

The deadline remains **2026-11-07 22:51:05 UTC** (14:51:05 PST).
The completed checks budget began 2026-11-04 22:51:05 UTC. It and the earlier
stabilization budget are closed. Preserve the deadline and enforcement bounds.
The selected two-change follow-up is complete. The single approved native
invocation is stopped; further drills and corrections need separate agreement.

CK1–CK4 remain complete. Their native report and proof ref
`refs/fireweed/drills/checks/20261107T002339.020159000` are preserved. The previous
failed workflow invocation and reconstruction ref remain retained; they do not
satisfy the passing workflow-proof gate. Preserve prior reports/proof/attempt
refs and all earlier archive counters. Existing budget warnings and the scoped
repeated-workflow warning are explained in the preparation evidence.

At handoff, the stopped drill processes exited and its input tree, verifier
worktree and isolated caller clone were removed. Main was clean before protected
evidence authoring. No live run/shared claim or queued change remains. The live
unapproved count is two; normal approval enforcement and its configured bound
are unchanged. Configuration, controller/recovery executables, all prior refs,
archive contents/counters and the unrelated linked worktree were preserved.
No product, spec or eval correction, reconstructed source import, controller
promotion or push occurred. Protected evidence is committed after the stop.

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
