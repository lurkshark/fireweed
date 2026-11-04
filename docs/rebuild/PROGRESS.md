# Rebuild handoff

**Current milestone: stabilization and cleanup.** A, B and C1–C7 passed.
The owner-authorized CHK-083/RUN-067 recovery passed audit and full verification.
C6 passed after owner-authorized APP-154 narrowing, retaining every assertion
and its audit counter. C7 includes both additional recipe bundles, passing
fresh default-init doctor proofs and full verification. The owner approved and
pushed prior work. C8 is next; the final local `go test ./...` remains due.
CK1–CK4 are deferred; full product completion remains open. Existing
testcase-budget warnings remain outside this cleanup.

**Controller pin:** `/Users/joel/Workspace/fireweed/.fireweed/bin/fireweed`,
version 0.1.0, source `281ae9498c0fa20b44dc525494427ff123331ab3`, SHA-256
`bf0dbc59b7fa7ce2ff957a6f21f34903655b6b054a10695cfeef5ece67fa3575`.
Keep it separate from mutable `bin/fireweed`; bootstrap is explicit recovery only.
This non-trimpath pin reproduces at its embedded clean source location.
Worker PATH starts with the current tree's `.fireweed/bin`; use the prescribed
rebuild cache and exact agent argv in [GOAL.md](GOAL.md).

**Next:** Execute the single real gitstore regeneration (C8), in
[ACCEPTANCE's order](ACCEPTANCE.md#current-milestone--stabilization-and-cleanup).
Execution start: **2026-11-02 19:53:07 PST**. Three-day deadline:
**2026-11-05 19:53:07 PST** (2026-11-06 03:53:07 UTC).
Stop for owner review at that deadline or after C8, whichever comes first.
No push or module drill is authorized for execution.

**Open problems:** Both stack tools are available;
use the existing pnpm through a project version pin because the unpinned
Corepack default tries the registry. Keep every assertion, logging and
independent-audit control intact.
Preserve normal approval/audit/retry bounds.
Committed evidence remains in [archived changes](../../changes/archive/) and
[native drill reports](../../.fireweed/drills/); git history retains older handoffs.
