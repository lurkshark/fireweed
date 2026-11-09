# Rebuild progress

## Current state (2026-11-09)

WF1 is complete. Fireweed applied and archived both changes after separate
audits with no findings and passing implementation verification:

- [0195-run-argv-contract](../../changes/archive/0195-run-argv-contract.md)
  (`eb4a7f2`) documents `gitstore.Request.Args[0]` as the executable and
  adds regression coverage. All 925 cases passed.
- [0196-drill-self-checked-reconstruction](../../changes/archive/0196-drill-self-checked-reconstruction.md)
  (`084fbd4`) gives the drill worker a buildable input workspace with
  dependency implementations and its own eval hook. All 485 cases passed.

The queue is empty and both claims have been released.

WF2 is complete. `go test ./...`, `go vet ./...` and `fireweed ci` passed
on the controller source commit below. The first Go test run alongside CI
failed a timed app eval, which passed on the standalone retry. CI ran with
Doctor enabled in a temporary local clone over `65bf933..6cc2629`.

WF3 is complete. With owner approval, the recorded controller ran
`fireweed drill workflow` once at `4754026` with the configured Codex worker.
The drill passed after its configured retry. All 666 cases passed (181
workflow, 323 app, 162 runner) and the worker reported no contradicted
Decisions. Fireweed cleaned up the scratch worktree and retained reconstruction
`38aa8fc` under `refs/fireweed/drills/workflow/20261109T224519.080152000`.
The [native report](../../.fireweed/drills/workflow.md) is committed in `c62b5be`.

WF4 awaits the owner's disposition of the three reported gaps:

- **FLOW-110 / APP-126.** The CLI's P01 for an unavailable base is already
  specified and covered by APP-126. The drill retry passed it. Close this
  observation unless the owner wants the workflow API contract clarified.
- **FLOW-052 / checks.** A successful boundary hook's raw output is absent
  from the API result. A local public-API probe confirmed that the API returns
  diagnostics only. Decide whether raw boundary output belongs in run reports.
- **FLOW-027 / checks.** A required boundary hook that exits 3 and prints only
  an unattributed error currently passes. A local public-API probe reproduced
  the false pass and confirmed X01 still reports attributed violations. The
  design requires the boundary hook to pass. This is a candidate checks change.

## Controller

The installed `.fireweed/bin/fireweed` was rebuilt from clean main for WF2.

Controller source: `6cc26290333cb47c1a314b16b8b50679afdd07f4`

## Next step

Review the three gap dispositions with the owner (ACCEPTANCE WF4). Queue
agreed changes and record the disposition of the other observations.

Deadline set by the owner: **2026-11-10 22:51:05 UTC**.
