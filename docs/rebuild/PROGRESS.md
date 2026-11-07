# Rebuild handoff

**CK1–CK4 complete.** The independent checks drill on validated candidate
`db735ba` passed its second configured attempt. Checks, app, runner and
workflow passed enforced evals and coverage. The
[reviewed native report](../../.fireweed/drills/checks.md) records CK1,
the exact CK2 validation range and candidate, and the live isolation and
trace review. The passing proof is
`refs/fireweed/drills/checks/20261107T002339.020159000`.

The scoped FLOW-055 correction `0177-prioritize-unsupported-phase` is
archived at `db735ba`. The candidate passed the affected CK2 gates,
including direct Verify and Done CLI checks and explicit-range CI with hook
doctor enabled. The validated range includes the pulled README title merge.
Earlier native reports and scoped recoveries remain in committed history,
including the preceding reviewed report at `e540462`. Prior proof and
attempt refs are retained.

Main and the validated candidate are clean, with no active run or shared
claim. Native input and verifier cleanup passed. The pin, configured bounds,
original deadline and unrelated worktree were preserved. Stop at CK4. The
next module in the planned drill order is workflow, pending separate owner
selection. Full product completion remains open.

[ACCEPTANCE](ACCEPTANCE.md#planned-milestone--checks-contracts-and-proof-ck1ck4)
holds the ordered plan; [GOAL](GOAL.md) is its restart prompt.
Approved plan budget: **three days**, starting **2026-11-04 22:51:05 UTC**
(14:51:05 PST), with deadline **2026-11-07 22:51:05 UTC** (14:51:05 PST).
Retain this original start/deadline on resume. The completed stabilization budget
is closed. Full product completion remains open; every full-completion gate is
retained. Prior proof refs and the unrelated worktree are preserved.

**Handoff authorization (2026-11-04):** the previous pin returned U01 because
native drill was not implemented. The owner explicitly authorized building from
clean committed current source, completing all CK2 gates, then replacing the pin
with no active run or shared claim and committing the protected pin instructions
separately. Validation passed and live status showed an empty queue, no shared
claim and no active change. Approval, audit/retry bounds and the original budget
were preserved; no product or eval correction was needed.

**Controller pin:** `/Users/joel/Workspace/fireweed/.fireweed/bin/fireweed`,
version 0.1.0, source `193634686aa6a1c4eb32e69df1c85e1d5ce48bd7`, SHA-256
`23e87b00217ab1a6990352393311bd150b7d60d3a4c138b75e285f518fba117f`.
Built with Go 1.27.2 and `-trimpath` from clean source on darwin/arm64.
The 2026-11-06 handoff passed normal and race Go tests, vet, candidate check,
hook doctor, and the full explicit history check. Candidate CI with hook
doctor passed for the range since the validated FLOW-062 candidate.
The controller stays separate from mutable `bin/fireweed`.

The immediate predecessor is preserved at `.fireweed/bin/fireweed-99a1d6d`.
Earlier recovery executables remain at `.fireweed/bin/fireweed-pre-ck-handoff`
and `.fireweed/bin/fireweed-bootstrap`. Use these only for explicit
owner-directed recovery, with no automatic fallback or cross-controller
claim resume. Worker PATH starts with the current tree's `.fireweed/bin`.
Use the prescribed rebuild cache and exact agent argv in GOAL.

**Push authorization (2026-11-04):** the owner's separate one-time planning push
is completed; it authorizes no future push, publication, tags or bookkeeping-ref
push.

**Push authorization (2026-11-06).** The owner explicitly approved a one-time
push of all 43 commits from `ac4bf36` through `13a2685` to `origin/main`, with
`ac4bf36` as the exclusive base.

Existing testcase-budget warnings remain outside scoped cleanup. Preserve normal
approval/audit/retry bounds; read live approval state rather than treating a
prior handoff as current approval. Evidence remains in archived Notes/Provenance
and native drill reports; no supplemental inventories or manifests.
