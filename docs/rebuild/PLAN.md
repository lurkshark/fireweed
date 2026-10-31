# Rebuild plan

The design remains product authority; SYSTEM.md governs execution.
[PROGRESS.md](PROGRESS.md) holds current state, evidence and the controller pin;
[ACCEPTANCE.md](ACCEPTANCE.md) holds milestone and full-completion gates;
[GOAL.md](GOAL.md) is the restart prompt. Keep execution order here rather than
repeating completed queues or gate criteria in several documents.

## Completed — Early Go core-loop cutover

EC1–EC6 passed, including 0131–0138, demonstrated follow-up fixes, real-worker
self-hosting and the separate NEW controller handoff at `c1de8f4`. Preserve
archived Provenance and reuse the valid evidence recorded in PROGRESS. Bootstrap
is an explicit recovery option only; the early-cutover queue is closed.

## Current milestone — Drill support and one independent contracts drill

Owner-selected on 2026-10-31. Implement the design's drill workflow and prove
one independent contracts rebuild in an isolated clone. Acceptance DR1–DR4
bounds completion. This starts the real-regeneration work in full gates 7/8;
it does not complete either gate for every module or the full product.

1. **Review the slice once.** Trace the design's Regeneration section and R16
   to current runner/app specs and existing public evals. NEW currently has no
   drill API and its CLI marks drill unimplemented. Reuse existing setup,
   hooks, zone maps, process bounds and regeneration guidance. Add a lower-level
   prerequisite only for a concrete missing behavior needed by this slice.
2. **Implement through the pinned controller.** Author coherent changes for
   the runner drill lifecycle and app command, returning to the CLI as soon as
   prerequisites are usable. Keep normal fail-first, restricted independent
   audit, phase-zone checks, approval/P07, attempts and audit rounds. Permanent
   evals use deterministic fixtures/fake agents for report/ref outcomes,
   cleanup, failures and bounded stops; no real model access in Go tests.
3. **Close the CLI slice.** Build a clean candidate and pass DR2's checks.
   Exercise an actual candidate drill invocation with deterministic fixtures,
   including its output/exit behavior. Identify the executable being tested.
   Main's pinned controller continues to govern implementation changes.
4. **Run one real contracts drill in an isolated clone.** Before invoking an
   agent, inspect scratch setup and reconstruction inputs for independence.
   The rebuild worker receives only retained spec/Decisions/interfaces/evals
   and allowed dependency interfaces; contracts has no dependencies. Keep
   dependent implementations available to the verifier outside the rebuilding
   worker's inputs. Exclude old implementation/history, saved attempts, old
   binaries and bootstrap as reconstruction inputs. If that separation is not
   established, stop before spending the real run.
5. **Review and close.** Use the configured model/effort, check own and all
   transitive dependent evals, inspect the native report and retained ref, and
   verify unchanged retained files plus scratch cleanup. Retain the isolated
   clone until its concise handoff is recorded. Stop on a blocked change,
   failed independence/access, or any gap/contradicted Decision needing owner
   agreement. Do not author gap fixes or launch another module's drill.

Scope ends after this contracts proof and its reviewed evidence. No main pin
promotion, main contracts-code replacement, real regenerate change, additional
module drill, deferred command/recipe/version/packaging work or push is included.

## Controller and candidate

Use the unchanged `/Users/joel/Workspace/fireweed/.fireweed/bin/fireweed` pin
from PROGRESS for all main workflow commands. Keep `.fireweed/bin` first on
worker PATH and `GOCACHE=/tmp/fireweed-rebuild-cache`; FIREWEED_BIN names the
actual invoker. Worker builds of mutable `bin/fireweed` never replace the pin.

The new drill command is tested through an explicitly identified candidate
executable in the isolated clone. Record its source/hash, configure clone and
scratch instructions/setup/PATH consistently, and ensure hooks copy that
invoking candidate. This fixture selection does not repin main. Never silently
fall back to bootstrap or resume another controller's live claim.

## Per-change loop

1. Author: read SYSTEM, projected specs, relevant design clauses and the pinned
   controller's author-change playbook. Motivation names the CLI scenario and
   design/SYSTEM basis. Use new: handles, stamp, check and show --projected;
   commit the authored change separately with no active run/claim.
2. Controller start, then context before each worker step.
3. Evals: black-box tests, described interfaces and allowed new disposable stubs.
   Every Added/Modified/fixes ID needs genuine failure unless already_satisfied.
   Iterate with verify --quick; done runs full checks. Use eval_fix for wrong
   frozen evals, preserving coverage and behavioral obligations.
4. Independent audit: only unasserted spec clauses, restricted inputs, re-audits
   limited to prior finding IDs, and at most the configured findings rounds.
5. Implementation: allowed files only. Controller done verifies, applies,
   archives and commits. Workers never commit. Stop on every blocked change;
   never alter refs, bounds or requirements to avoid a checkpoint.
6. Record current state in PROGRESS and per-change evidence in archived
   Provenance. The native drill report is required product evidence; create no
   supplemental reports, testcase inventories or hash manifests.

## Later full completion

The full requirement inventory and all nine gates remain in ACCEPTANCE. Later
owner-selected slices cover the remaining modules' independent drills, real
regenerate validation, gc/migrate and stale/fix conveniences, recipes,
version/file-format migration, warning completeness and delivery. Reuse valid
cutover evidence; rerun checks only when relevant changes or unresolved failures
invalidate it, in addition to required controller phase checks.
