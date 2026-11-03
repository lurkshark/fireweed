# Rebuild handoff

**2026-11-02 — current milestone: stabilization and cleanup.** Planning updated;
execution A–C remains pending. Queue and claims are empty. The pinned executable
is present and matches the pin below, contrary to the supplied earlier missing-file
snapshot; no restoration build or Part A validation has been completed here.
Checks contracts/proof (CK1–CK4) are deferred. Full product completion remains open.

**Controller pin:** `/Users/joel/Workspace/fireweed/.fireweed/bin/fireweed`,
version 0.1.0, source `281ae9498c0fa20b44dc525494427ff123331ab3`, SHA-256
`bf0dbc59b7fa7ce2ff957a6f21f34903655b6b054a10695cfeef5ece67fa3575`.
Keep it separate from mutable `bin/fireweed`; bootstrap is explicit recovery only.
Worker PATH starts with the current tree's `.fireweed/bin`; use the prescribed
rebuild cache and exact agent argv in [GOAL.md](GOAL.md).

**Next:** execute [ACCEPTANCE's stabilization plan](ACCEPTANCE.md#current-milestone--stabilization-and-cleanup):
A restores/verifies the identical controller, B commits the protected process edits,
then C1–C8 fix script fixtures, save agent logs, add prompt files, batch small fixes,
implement gc/migrate, add both recipes and run one real gitstore regenerate change.
Budget: three days from execution start; record start/deadline here when beginning.
Stop for owner review at that deadline or after C8, whichever comes first.
The owner authorized committing/pushing this planning update before A (2026-11-02).
No product-work push or module drill is selected.

**Open problems:** accept the supplied diagnosis of macOS sh killing long-prompt
fixtures; C1 must restore passing local Go tests. Agent output is not yet persisted;
C2 addresses it. pnpm/uv must be available for C7 or execution stops.
Preserve normal approval/audit/retry bounds.
Committed evidence remains in [archived changes](../../changes/archive/) and
[native drill reports](../../.fireweed/drills/); git history retains older handoffs.
