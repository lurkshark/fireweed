# Rebuild progress

## Current state (2026-11-09)

Main was reset to `515516a` to drop the drill-specific changes 0192
(drill compile checkpoint) and 0193 (focused drill retries) and the
supervised drill harness. Change 0194 (keep run artifacts local) is a
general fix, so its original queue, evals, audit and apply commits were
replayed on top. The previous head, `f35fc0f`, is kept on the local branch
`backup/pre-drill-simplify`. The rewritten main was force-pushed to origin
on 2026-11-09. The bootstrap binary was later removed and
`.fireweed/drills/` pruned to the per-module reports.

The workflow drill was failing because the reconstruction worker could not
compile or test its code, not because of spec gaps. The fix is to change
the product's drill rather than supervise the worker. Two changes are
queued and have not been run:

- [0195-run-argv-contract](../../changes/0195-run-argv-contract.md) states
  that `gitstore.Request.Args[0]` is the executable. This was the gap the
  last workflow drill exposed.
- [0196-drill-self-checked-reconstruction](../../changes/0196-drill-self-checked-reconstruction.md)
  gives the drill worker a buildable input workspace with dependency
  implementations and its own eval hook, while leaving out the old
  implementation, history and dependents.

## Controller

No `.fireweed/bin/fireweed` is installed in this checkout. Before running
the queue, build one from clean main as SYSTEM.md describes and record the
commit here.

Controller source: _not yet built_

## Next steps

1. Run 0195 and 0196 through the normal lifecycle.
2. Rebuild the controller from the result so the drill uses the new input
   workspace.
3. Run `fireweed drill workflow` once, commit the report and review its gaps
   with the owner (ACCEPTANCE WF3–WF4).

Deadline set by the owner: **2026-11-10 22:51:05 UTC**.
