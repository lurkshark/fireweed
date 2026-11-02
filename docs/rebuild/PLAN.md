# Rebuild plan

The design remains product authority; SYSTEM.md governs execution.
[PROGRESS.md](PROGRESS.md) records state, evidence and the controller pin;
[ACCEPTANCE.md](ACCEPTANCE.md) defines gates; [GOAL.md](GOAL.md) is the restart
prompt. Keep execution order here and per-change evidence in archived Provenance.

## Completed milestones

EC1–EC6 completed the early Go core-loop cutover and separate NEW pin handoff.
DR1–DR4 completed native drill support and one independent contracts proof,
including owner-approved contracts clarifications in 0144. Preserve the recorded
checks, native report/ref, inspectable clones and immutable archives. Completed
implementation procedures are not a queue to repeat.

## Completed milestone — Stale-change authoring support and independent project drill

Owner-selected and completed on 2026-11-01; **SP1–SP4 passed**. This
succeeded FD1–FD4 and closed the stale-authoring portion of design R28
with a passing independent project proof. PROGRESS records reviewed evidence;
the following execution steps and GOAL prompt retain the completed scope.
The unchanged design remains authority. The scope is project behavior and its
app CLI integration, then **one real
project drill**. Project depends only on contracts, gitstore and assets, whose
reviewed independent proofs remain preserved and reusable.

1. **Define the scoped changes.** Read the design's Stale changes/Fixing a stale
   change clauses, current project/app specs and retained interfaces/evals.
   Author coherent module changes through the pinned NEW controller. Cover
   `show --stale <change>` and `check --fix`; add only the interfaces, CLI
   guidance or diagnostics required by those existing design behaviors. Keep
   spec items atomic and preserve every existing case, assertion and ID.
2. **Implement stale-change authoring support.** For each stale spec, show the
   diff from the stamped text to the current projected text. Recover old text
   by projecting at the commit that last changed the change's bases and verify
   its normalized hash; when matching history cannot be recovered, show current
   text and recorded hash. `check --fix` stamps only missing base entries and
   preserves all existing entries, including stale ones. Ordinary check remains
   read-only and still diagnoses stale bases; the fix flag never resolves a
   conflict silently. Use normal fail-first, restricted independent audit,
   immutable phase zones, approval and configured bounds for project/app work.
3. **Verify the actual NEW CLI and candidate.** Exercise both options in fresh
   local Git fixtures, including historical recovery, fallback, mixed missing/
   valid/stale bases, usage/lookup failures and applicable JSON envelopes/exits.
   Confirm show is read-only and fix changes only permitted missing bases.
   Finish and archive all scoped changes before the drill. Build an explicitly
   identified clean NEW candidate; pass Go tests/race/vet and candidate check,
   doctor --no-agent, explicit-baseline check-commits and CI with hook doctor
   enabled. Reuse valid unaffected evidence; required phase checks still run.
   Main workflow commands keep the original pin; new product options are tested
   against the explicitly selected candidate, not assumed present in that pin.
4. **Prove project independently.** After the feature changes and checks close,
   align the isolated candidate's instructions, setup and worker PATH; hooks
   copy the actual invoker. Inspect the reconstruction inputs before spending
   the real invocation: retained project spec/Decisions/interfaces/evals and
   approved retained files, plus only declared interfaces from contracts,
   gitstore and assets. Keep dependent implementations and build/hook setup in
   a separate verifier. Exclude old target disposable code, embedded data,
   history, attempts, product binaries, bootstrap and dependency internals.
   Never query invoker content as an oracle. Require own and every transitive
   dependent module's evals to pass with retained inputs unchanged, then review
   the complete worker commands, native report and rebuilt-code ref.
5. **Close SP1–SP4.** Commit reviewed `.fireweed/drills/project.md` as protected
   author evidence with no live run/shared claim; retain its
   `refs/fireweed/drills/project/<date>` ref and inspectable clone location.
   Record concise results in PROGRESS, with per-change evidence in archived
   Provenance. Verify cleanup, clean checkouts, unchanged pin/bounds and preserved
   completed proofs/unrelated worktrees. Reassess only evidence invalidated by
   relevant changes; an additional module's real drill needs owner selection.

Use exact configured workers/auditors/drill workers from GOAL. Stop on a blocked
change, independence/access failure or gap/contradicted Decision needing owner
agreement. Do not automatically author drill-gap fixes, reset retries/rounds or
expand the scope. Owner-approved corrections use the normal loop below.

This milestone includes ordinary project/app implementation for the two options;
it includes no import of the independently rebuilt project implementation into
main, real regenerate change, controller promotion, other-module drill, gc,
migrate, additional recipes, file-format migration, unrelated warning cleanup,
deferred product work or push. Full product completion remains open.

## Completed foundation milestone — Independent gitstore and assets drills

Owner-selected and completed on 2026-11-01: **gitstore first, then assets, one
at a time**. FD1–FD4 passed; PROGRESS records the reviewed native evidence.
The sequence below records the completed milestone. Each depends only on
contracts, whose independent proof is complete. ACCEPTANCE FD1–FD4 bounds this
milestone and advanced full regeneration gate 7. Full product completion remains
open; the next planned milestone is defined above.

1. **Prepare the candidate and inputs.** Inspect current state and verify main's
   unchanged controller pin. Select a clean isolated NEW candidate and identify
   its executable/source/hash. Align clone/scratch instructions, setup and PATH
   with that candidate; setup copies the actual invoking controller. Reuse
   unaffected DR2/0144 checks and contracts evidence; run checks invalidated by
   relevant changes. Inspect input membership and independent Git metadata
   before spending either real run, without consulting old target implementation.
2. **Prove gitstore.** Apply approved disposable zones and supply retained target
   inputs plus contracts' declared interfaces only. Keep full-project verification
   outside the rebuilding worker's inputs. Run within existing bounds, review
   own/all-transitive-dependent evals, every worker tool command, gaps and
   contradicted Decisions. Retain its native report/ref and verify cleanup.
   Stop at required owner checkpoints. Finish review and commit evidence with
   no active run/claim before proceeding.
3. **Prove assets.** Repeat the same process only after gitstore's proof closes.
   Delete all target disposable content, including embedded bundle files. The
   candidate's embedded recipes, playbooks, help and explanations are old target
   implementation and cannot supply reconstruction inputs. Gitstore internals
   are not an assets dependency interface. Never overlap the two real drills.
4. **Close FD1–FD4.** Record both candidate identities, native reports/refs and
   inspectable clones concisely in PROGRESS. Preserve retained contracts/evals,
   verify cleanup and unchanged main pin/bounds, and stop once the gates pass.

Use the exact agent argv in GOAL and the execution rules in SYSTEM. Gap fixes
require owner agreement before authoring; they are not automatically authorized
by this plan. The milestone excludes other modules' real drills, main
implementation import/replacement, real regenerate changes, controller promotion,
deferred product work and push. Full product completion remains open.

## Per-change loop for owner-approved corrections

1. Author: read SYSTEM, projected specs, the affected design clauses and the
   pinned controller's author-change playbook. Motivation names the concrete
   failure or agreed gap. Use new: handles, stamp, check and show --projected;
   commit the authored change separately with no active run/claim.
2. Controller start, then context before each worker step. Read live approval
   state and authoring bounds without recording their values in planning docs.
3. Evals: public black-box tests with genuine fail-first per affected ID unless
   legitimately already_satisfied. Preserve existing cases, assertions and ID
   coverage. Iterate with verify --quick; done always runs full checks. Use the
   controller's eval_fix path for wrong frozen evals; do not repeat historical
   recovery exceptions without explicit owner authorization.
4. Independent restricted audit: only unasserted spec clauses, re-audits limited
   to prior finding IDs, with the unchanged configured round/attempt bounds.
5. Implementation: allowed files only, frozen evals/interfaces/metadata. The
   controller verifies, applies, archives and commits; workers never commit.
   Stop whenever a change blocks; never reset counters or alter enforcement.
6. Record current handoff in PROGRESS and per-change evidence in Provenance.
   Reassess evidence affected by the correction; do not repeat unaffected checks
   or expand real-drill scope automatically.

## Later full completion

All nine full-completion gates remain in ACCEPTANCE. After SP1–SP4, the remaining
drill order is checks → workflow → runner → app. Later owner-selected work also
includes real regenerate validation, gc/migrate, additional recipes, version/
file-format migration, warnings and delivery. Coordinate later behavior changes
with affected proofs and repeat only evidence they invalidate. None is started
by this planning update.
