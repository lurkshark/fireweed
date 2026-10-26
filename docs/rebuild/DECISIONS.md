# Rebuild decisions

Record clarifications of the design and intentional differences from the bootstrap here, one short entry each. Decisions from the first run are on the `rebuild-v1` branch; [LESSONS.md](LESSONS.md) summarizes the ones that still matter.

## D001 — Restart from the setup baseline

On 2026-10-26 the project owner restarted the rebuild. The first run is preserved on the `rebuild-v1` branch. Local `main` was hard-reset to the setup baseline 14cec73, and the revised design, scope guardrails and rebuild documents were reapplied in separate author commits. `origin/main` still points at the first run until the owner force-pushes. Until then the controller compares against the old base, so `start` reports base and approval warnings; manual execution under the goal's authority proceeds through them as before.

The module graph (eight modules, same dependencies) and the Go recipe setup are unchanged. The Go eval hook gained one fix from the first run: unfinished tests and package build failures are failing cases.
