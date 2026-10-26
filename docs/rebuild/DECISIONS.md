# Rebuild decisions

Record clarifications of the design and intentional differences from the bootstrap here, one short entry each. Decisions from the first run are on the `rebuild-v1` branch; [LESSONS.md](LESSONS.md) summarizes the ones that still matter.

## D001 — Restart from the setup baseline

On 2026-10-26 the project owner restarted the rebuild. The first run is preserved on the `rebuild-v1` branch. Local `main` was hard-reset to the setup baseline 14cec73, and the revised design, scope guardrails and rebuild documents were reapplied in separate author commits. The owner force-pushed `main` to `origin` the same day, so the controller's base matches the restart.

The module graph (eight modules, same dependencies) and the Go recipe setup are unchanged. The Go eval hook gained one fix from the first run: unfinished tests and package build failures are failing cases.

## D002 — Usability revision of the design

On 2026-10-26 the owner accepted eight usability changes to the design: approval from the local base branch (with the unapproved-work bound still counted against the remote), `abandon`, `verify --quick`, a doctor agent check, a smaller default `context`, waiting reasons in `status`, `show --stale` and `check --fix`, and a C19 warning for long module Decisions. Parallel independent changes were deferred. ACCEPTANCE R26 to R28 track them; PLAN.md records the new design SHA-256. The pinned bootstrap implements none of these.

## D003 — Any merge strategy

On 2026-10-26 the owner chose to keep per-step commits for review while letting users merge however they like. The design now checks steps before merge (`ci` on the PR branch), recognizes squashed change commits in `check-commits` instead of reporting false H02 violations, and counts drill drift by Provenance dates instead of commit ancestry. ACCEPTANCE R29 tracks it; PLAN.md records the new design SHA-256.
