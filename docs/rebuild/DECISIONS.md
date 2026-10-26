# Rebuild decisions

Record clarifications of the design and intentional differences from the bootstrap here, one short entry each. Decisions from the first run are on the `rebuild-v1` branch; [LESSONS.md](LESSONS.md) summarizes the ones that still matter.

## D001 — Restart from the setup baseline

On 2026-10-26 the project owner restarted the rebuild. The first run is preserved on the `rebuild-v1` branch. Local `main` was hard-reset to the setup baseline 14cec73, and the revised design, scope guardrails and rebuild documents were reapplied in separate author commits. The owner force-pushed `main` to `origin` the same day, so the controller's base matches the restart.

The module graph (eight modules, same dependencies) and the Go recipe setup are unchanged. The Go eval hook gained one fix from the first run: unfinished tests and package build failures are failing cases.

## D002 — Usability revision of the design

On 2026-10-26 the owner accepted eight usability changes to the design: approval from the local base branch (with the unapproved-work bound still counted against the remote), `abandon`, `verify --quick`, a doctor agent check, a smaller default `context`, waiting reasons in `status`, `show --stale` and `check --fix`, and a C19 warning for long module Decisions. Parallel independent changes were deferred. ACCEPTANCE R26 to R28 track them; PLAN.md records the new design SHA-256. The pinned bootstrap implements none of these.

## D003 — Any merge strategy

On 2026-10-26 the owner chose to keep per-step commits for review while letting users merge however they like. The design now checks steps before merge (`ci` on the PR branch), recognizes squashed change commits in `check-commits` instead of reporting false H02 violations, and counts drill drift by Provenance dates instead of commit ancestry. ACCEPTANCE R29 tracks it; PLAN.md records the new design SHA-256.

## D004 — Fresh bootstrap and vertical-slice stage 2

On 2026-10-26 the fresh bootstrap completed independently at runtime source
`7f925190878e38fc38f30d16bb1cf47aca098cf7`, version 0.1.0; HANDOFF.md records its
verification. The owner requested a new stage-2 goal that proves a runnable
manual lifecycle through the NEW CLI before completing the remaining surface.
This supersedes the old author-only first session and old controller/reference
instructions. D001–D003 and PLAN's old inspection record remain historical.

Use the tested executable as the pinned controller, without consulting bootstrap
source/tests or rebuild-v1 as an implementation oracle. Keep all design-derived
acceptance gates and human checkpoints. Use Codex `gpt-6-sol` with high reasoning
effort for workers, separate restricted auditor invocations and regeneration.
A read-only Sol/low-effort probe passed; high-effort doctor checks are still required.

Preparation changes only the instructions. Controller installation, reconciling
fireweed.yaml's old version pin and canary/hook setup, doctor and the new history
baseline belong to the goal's startup. No stage-2 worker, init, drill or rebuild was
run during this review. The unchanged design remains product authority; its vendor
examples describe a language/agent-agnostic product and do not select this build's
agent provider.

## D005 — Initial product diagnostic allocation

The design requires stable codes and explanations but does not assign a meaning
to P2-01. Newly authored AST-003 assigns it to missing genuine fail-first evidence.
The controller's operational `explain P2-01` instead describes a forbidden-zone
edit. Keep the authored product meaning and record this intentional difference
in the later controller/product smoke; controller code assignments are not
product authority. Other code mappings remain to be authored with their behavior.
