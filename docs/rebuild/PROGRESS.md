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

## Controller

The installed `.fireweed/bin/fireweed` ran WF1.

Controller source: `65bf9339b3b7491675625f546a6055eab38b488f`

## Next steps

1. Complete WF2. With no active run or claim, rebuild the controller from
   clean main and record its source commit here. Verify that commit with
   `go test ./...`, `go vet ./...` and `fireweed ci`.
2. Run `fireweed drill workflow` once, commit the report and review its gaps
   with the owner (ACCEPTANCE WF3–WF4).

Deadline set by the owner: **2026-11-10 22:51:05 UTC**.
