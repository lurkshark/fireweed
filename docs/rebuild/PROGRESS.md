# Rebuild handoff

Workflow proof preparation is the next target. WF1 contract review and WF2
clean candidate validation remain pending. Follow the
[ordered plan](ACCEPTANCE.md#planned-milestone--workflow-contracts-and-regeneration-proof-wf1wf4)
and use the [restart prompt](GOAL.md). Stop for owner agreement before
corrections and for explicit candidate approval before a native workflow drill.

Preparation retains the deadline **2026-11-07 22:51:05 UTC** (14:51:05 PST).
The completed checks budget began 2026-11-04 22:51:05 UTC. It and the earlier
stabilization budget are closed. Preserve the deadline and enforcement bounds.

CK1–CK4 are complete. The [reviewed checks report](../../.fireweed/drills/checks.md)
records the contract review, validation of candidate `db735ba` and the passing
independent proof. The proof ref is
`refs/fireweed/drills/checks/20261107T002339.020159000`. FLOW-055 recovery is
archived as `0177-prioritize-unsupported-phase`. Earlier reports, recoveries and
handoffs remain in committed history. Preserve prior proof/attempt refs and
unrelated worktrees. Full product completion remains open.

At the completed CK4 handoff, main and the candidate were clean, native scratch
cleanup passed and no live run/shared claim remained. Recheck live state before
execution. Existing testcase-budget warnings remain outside scoped cleanup.

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
