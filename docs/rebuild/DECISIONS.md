# Rebuild decisions

Record clarifications of the design and intentional differences from the bootstrap here, one short entry each. Dated preparation/restart statements below are historical; current readiness, approvals and next work belong in [PROGRESS.md](PROGRESS.md). Decisions from the first run are on the `rebuild-v1` branch; [LESSONS.md](LESSONS.md) summarizes the ones that still matter.

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

## D006 — Diagnostic allocation (human review 3)

Codes identify a stable cause. Their prefix does not determine severity; the
diagnostic carries its warning status. Assignments below apply to the NEW
product. The pinned controller is unchanged.

| Range | Area and fixed assignments |
| --- | --- |
| C01–C99 | Deterministic configuration and static checks. C01 configuration discovery/decoding; C02 module graph; C03 document/item syntax; C04 normative text; C05 ID ownership/uniqueness; C06 spec tags; C07 change structure/Motivation; C08 bases; C09 handles; C10 coverage; C11 eval-hook/JUnit contract; C12 metadata/zone declarations; C13 queue identity. C14–C16 reserved. C17 spec shape, C18 copied conventions, C19 Decision length retain design meanings. |
| P01–P99, P2-01 | Lifecycle and approval. P01 preconditions; P02 claims; P03 approved state; P04 approval warning; P05 stale base; P06 blocked/Questions; P07 unapproved-work limit; P08 repeated-module warning; P09 audit submission. Existing P2-01 remains missing genuine fail-first evidence. |
| Z01–Z99 | Phase-zone permissions. Existing Z01 remains a forbidden path edit. |
| X01–X99 | Boundaries. X01 unsuppressed violation; X02 required hook missing or unable to complete; X03 reasoned opt-out; X04 boundary enforcement unnecessary. B05 is retired. |
| U01–U99 | CLI usage and lookup. U01 invalid arguments/options; U02 unknown lookup name. Replace the legacy CLI_* labels. |
| E01–E99 | Required eval results. E01 an enforced testcase failed where passing results are required. |
| H01–H04 | History, with the design's exact meanings. |
| B01–B04 | Budget/report warnings only, with the design's exact meanings. No B05 allocation. |

Unlisted numbers are reserved. Every emitted code must have an accurate cause
and corrective action in explain. Wrong label assertions are repaired through
eval_fix without losing the behavior assertions or ID coverage. Normal changes
then establish exact assigned-code assertions and migrate implementation. The
spec split removes the old blanket C01 promise before that migration; no existing
Specs or Decisions are edited by hand.

Owner-approved migration scope: change codes only on error paths that already
exist. Do not stub, scaffold or implement missing behavior to attach or exercise
a diagnostic, and do not add speculative error paths. Existing unsupported-command
rejections may have their codes corrected without implementing those commands.
Apply normal eval_fix and implementation changes to existing emitted diagnostics
and behavior assertions. Explain must be accurate for codes actually emitted;
unused allocations and their tests/explanations wait for the corresponding
justified feature change.

## D007 — Doctor diagnostics

The Milestone 2 doctor feature uses the previously reserved C14–C16 slots:
C14 for doctor hook/canary setup, detection and recovery failures; C15 for agent
probe failures or usage limits, and the shared-auditor-command warning; C16 for
stale or missing managed guidance and playbook override base metadata. Severity
remains a property of the diagnostic. Add accurate assets explanations after
the doctor backend emits these causes and before CLI wiring/checkpoint review;
D006's existing assignments remain unchanged.

## D008 — Regenerate warning

The regenerate lifecycle feature assigns P10 to the warning that its module
lacks a passing drill dated after the latest applied implementation Provenance.
Add its explanation after the workflow emits it and before closing that CLI
slice; this allocates no unrelated diagnostic or drill authorization.
