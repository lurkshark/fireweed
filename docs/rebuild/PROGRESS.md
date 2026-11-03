# Rebuild handoff

**2026-11-02 — current milestone: stabilization and cleanup.** A passed:
exact pinned-source reproduction, pinned status, check and doctor --no-agent.
The initial mismatch was source-path sensitivity without -trimpath; the owner
authorized investigation and the clean original build location reproduced the pin.
B documentation cleanup is complete; C1 is next. Queue and shared claims are empty;
doctor scratch cleanup passed. Existing testcase-budget warnings remain outside
this cleanup. CK1–CK4 are deferred; full product completion remains open.

**Controller pin:** `/Users/joel/Workspace/fireweed/.fireweed/bin/fireweed`,
version 0.1.0, source `281ae9498c0fa20b44dc525494427ff123331ab3`, SHA-256
`bf0dbc59b7fa7ce2ff957a6f21f34903655b6b054a10695cfeef5ece67fa3575`.
Keep it separate from mutable `bin/fireweed`; bootstrap is explicit recovery only.
Worker PATH starts with the current tree's `.fireweed/bin`; use the prescribed
rebuild cache and exact agent argv in [GOAL.md](GOAL.md).

**Next:** C1 audited fixture eval_fix changes, one module at a time, preserving
all assertions, then C2–C8 in [ACCEPTANCE's order](ACCEPTANCE.md#current-milestone--stabilization-and-cleanup).
Execution start: **2026-11-02 19:53:07 PST**. Three-day deadline:
**2026-11-05 19:53:07 PST** (2026-11-06 03:53:07 UTC).
Stop for owner review at that deadline or after C8, whichever comes first.
No push or module drill is authorized for execution.

**Open problems:** accept the supplied diagnosis of macOS sh killing long-prompt
fixtures; C1 must restore passing local Go tests. Agent output is not yet persisted;
C2 addresses it. pnpm/uv must be available for C7 or execution stops.
Preserve normal approval/audit/retry bounds.
Committed evidence remains in [archived changes](../../changes/archive/) and
[native drill reports](../../.fireweed/drills/); git history retains older handoffs.
