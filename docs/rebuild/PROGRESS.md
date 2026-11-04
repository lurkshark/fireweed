# Rebuild handoff

**State:** stabilization A → B → C1–C8 is complete. Both additional recipes passed
fresh default-init doctor proofs; the real gitstore regenerate passed full
verification. Final local `go test ./...` passed at that handoff. The queue and
claims were empty, with prior proof refs and the unrelated worktree preserved.

**Active milestone: CK1–CK4, execution authorized on 2026-11-04.** CK1's focused
review is underway; no correction, candidate checks or real drill has started.
[ACCEPTANCE](ACCEPTANCE.md#planned-milestone--checks-contracts-and-proof-ck1ck4)
holds the ordered plan; [GOAL](GOAL.md) is its restart prompt.
Approved plan budget: **three days**, starting **2026-11-04 22:51:05 UTC**
(14:51:05 PST), with deadline **2026-11-07 22:51:05 UTC** (14:51:05 PST).
Retain this start/deadline on resume. The completed stabilization budget is closed.
Full product completion remains open; every full-completion gate is retained.

**Push authorization (2026-11-04):** the owner authorized a separate planning
commit and one push of all pending `main` commits to origin. This authorizes no
CK execution, future push, publication, tags or bookkeeping-ref push.

**Controller pin:** `/Users/joel/Workspace/fireweed/.fireweed/bin/fireweed`,
version 0.1.0, source `281ae9498c0fa20b44dc525494427ff123331ab3`, SHA-256
`bf0dbc59b7fa7ce2ff957a6f21f34903655b6b054a10695cfeef5ece67fa3575`.
Keep it unchanged and separate from mutable `bin/fireweed`; bootstrap is explicit
owner-directed recovery only. This non-trimpath pin reproduces at its embedded
clean source location. Worker PATH starts with the current tree's `.fireweed/bin`;
use the prescribed rebuild cache and exact agent argv in GOAL.

**Open checkpoints:** CK3 input/verifier isolation under this unchanged pin
must be established before the real
agent. No implementation need is presumed. Existing testcase-budget warnings
remain outside scoped cleanup. Preserve normal approval/audit/retry bounds;
read live approval state at execution rather than treating the prior handoff as
current approval. Evidence remains in archived Notes/Provenance and native drill
reports; no supplemental inventories or manifests.
