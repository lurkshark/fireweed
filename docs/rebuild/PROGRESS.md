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

## Controller

The installed `.fireweed/bin/fireweed` was rebuilt from clean main for WF2.

Controller source: `6cc26290333cb47c1a314b16b8b50679afdd07f4`

## Next step

Run `fireweed drill workflow` once, commit the report and review its gaps
with the owner (ACCEPTANCE WF3–WF4).

Deadline set by the owner: **2026-11-10 22:51:05 UTC**.
