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

Fireweed released both claims after applying the changes.

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

WF4 is complete. The owner reviewed the report and approved these dispositions:

- **FLOW-110 / APP-126.** Closed. APP-126 specifies P01 for an unavailable
  base and its eval passed on the drill retry.
- **FLOW-052 / checks.** Deferred raw boundary-hook output retention. The
  current boundary API returns diagnostics only.
- **FLOW-027 / checks.** Queued
  [0197-boundary-hook-nonzero-exit](../../changes/0197-boundary-hook-nonzero-exit.md)
  in `be8fa32` to reject a nonzero exit without a module-attributed violation.
  `fireweed stamp` and `fireweed check` passed. The change is unclaimed.

## Controller

The installed `.fireweed/bin/fireweed` was rebuilt from clean main for WF2.

Controller source: `6cc26290333cb47c1a314b16b8b50679afdd07f4`

## Next step

When selected for implementation, run `0197-boundary-hook-nonzero-exit`
through the normal fail-first, separate audit and implementation lifecycle.

Deadline set by the owner: **2026-11-10 22:51:05 UTC**.
