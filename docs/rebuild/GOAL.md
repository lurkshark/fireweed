# Drill milestone goal prompt

Restart prompt for the implementing agent. Editing this document does not start
implementation, invoke a drill or switch controllers. PROGRESS.md records the
pin/state; PLAN.md orders the slice; ACCEPTANCE.md defines DR1–DR4 and the later
full-completion gates. EC1–EC6 are complete and their evidence remains valid
unless relevant changes invalidate it.

```text
/goal Implement NEW drill support in /Users/joel/Workspace/fireweed and prove one independent contracts drill in an isolated clone. Complete ACCEPTANCE DR1–DR4 under the unchanged main NEW controller pin. Stop on any blocked change or gap needing the owner's decision. Full product completion remains open.

Resume:
- Read AGENTS.md, SYSTEM.md and docs/rebuild/{PROGRESS,PLAN,ACCEPTANCE,DECISIONS}.md. The unchanged design is product authority and SYSTEM's guardrails govern execution.
- Inspect live Git status, queue, controller processes and shared claims. Finish existing work only with its owning controller; never reset another worktree's work or release its claim.
- Use /Users/joel/Workspace/fireweed/.fireweed/bin/fireweed for main workflow commands. Verify the source/hash recorded in PROGRESS; keep that pin unchanged throughout this milestone. Never use mutable bin/fireweed as main's controller or silently fall back to bootstrap.
- Put the current tree's .fireweed/bin first on worker PATH. Use GOCACHE=/tmp/fireweed-rebuild-cache. Preserve completed cutover archives, evidence and unrelated worktrees.

Implementation:
- Follow PLAN's per-change loop. Review the drill slice once against the design's Regeneration section and R16, then review only affected fixes. Implement runner drill orchestration and app CLI support; add prerequisites only for concrete unmet behavior needed by this slice.
- Keep normal fail-first, separate restricted audits, immutable metadata/evals/interfaces, phase zones, approval/P07 and configured retry/audit/time/usage bounds. Read live status and authoring limits before starting work; never record their values in planning documents or change them to evade enforcement.
- Cover setup, disposable-only deletion, reconstruction inputs, own/transitive dependent verification, report/ref outcomes and cleanup through public deterministic evals. Keep Go checks independent of real agents/bootstrap.
- Never commit during a worker step. The pinned controller owns phases, Specs/Decisions, application and archives. Protected author/setup/evidence edits belong in separate commits with no active run or shared claim.
- Close implementation with an actual candidate CLI drill fixture and DR2's clean-candidate checks. Reuse valid evidence rather than repeat the completed cutover campaign.

One real proof:
- Use an isolated clone of the tested candidate and identify its executable/source/hash. Configure clone and scratch setup/instructions/PATH for that candidate; hooks copy the actual invoker. This product fixture does not change main's pin or take over a live claim.
- Inspect independence before the real invocation. Rebuild contracts from its retained spec, Decisions, interfaces and evals; it has no dependency interfaces. Do not expose old implementation/history, other modules' internals, saved attempts, old binaries or bootstrap as reconstruction inputs. The verifier may use dependent implementations outside the rebuilding worker's inputs.
- Workers, auditors and the drill worker use exactly: codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c 'model_reasoning_effort="high"' -s workspace-write {prompt}. On access failure, stop and report it; never substitute model or effort.
- Run only the contracts real drill, within normal configured bounds. Require own and all transitive dependent evals to pass, retained contracts/evals unchanged, the design-required native report and retained rebuilt-code ref, and complete scratch cleanup.
- Review the report and independence evidence. Stop on any blocked change, independence/access failure or gap/contradicted Decision needing owner agreement. Do not automatically author gap fixes, reset retries or expand to another module.

Completion and scope:
- Record concise DR1–DR4 results, candidate identity, report/ref and inspectable clone location in PROGRESS. Per-change evidence stays in archived Provenance; commit the reviewed native drill report as an author with no active run/claim. No supplemental report inventories or hash manifests.
- Finish with a clean checkout, no live runs/shared claims and main's original controller pin verified unchanged. Push only on explicit owner instruction.
- No other real module drill, main contracts implementation replacement, real regenerate change, controller promotion, deferred command/recipe/version/packaging work or release is included.
- This goal completes only when DR1–DR4 pass. A failed drill or unresolved gap is a checkpoint, not a waiver. All later full-completion gates remain required.
```
