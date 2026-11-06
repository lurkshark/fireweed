# Checks milestone restart prompt

**Plan and execution approved on 2026-11-04.** PROGRESS records the running
three-day execution budget and the separate, completed one-time planning push
authorization. CK execution excludes pushing.
PROGRESS records the verified controller pin. Keep it fixed during CK work
and retain the original budget deadline and all enforcement bounds.
[ACCEPTANCE](ACCEPTANCE.md#planned-milestone--checks-contracts-and-proof-ck1ck4)
is the single ordered CK1–CK4 plan and retains the full requirement/command
inventories and completion gates. [PROGRESS](PROGRESS.md) holds current state
and the sole controller pin. [SYSTEM](../../SYSTEM.md) governs process/stops;
[DECISIONS](DECISIONS.md) holds retained design clarifications. The unchanged
[design](../../fireweed-design-v1.md) remains product authority.

Execution restart prompt:

```text
/goal Execute CK1–CK4 in ACCEPTANCE in order under the accepted milestone budget: focused checks contract review and necessary corrections; clean identified candidate validation; one independent checks drill; reviewed evidence and clean handoff. Preserve every full-completion requirement and the exact controller pin in PROGRESS. Stop on the first blocked change, unresolved verification under SYSTEM's correction limit, uncovered owner decision, access/independence failure, budget deadline or CK4 completion. Do not push.

Resume:
- Read AGENTS.md, SYSTEM.md, relevant unchanged design clauses and docs/rebuild/{PROGRESS,ACCEPTANCE,DECISIONS}.md. Record the accepted budget's new execution start/deadline in PROGRESS, retaining it on resume. The completed stabilization deadline does not carry forward.
- Inspect live Git status, queue, runs and common claims; preserve unrelated worktrees and existing evidence/proof refs. Read approval state from status --json and bounds from fireweed.yaml without recording their values in planning documents. Preserve all counters/bounds; never replace a change/drill to reset them.
- Use /Users/joel/Workspace/fireweed/.fireweed/bin/fireweed as the sole workflow controller. Confirm its existing pin and stop on missing/mismatched identity; do not repin or fall back to bootstrap. The candidate executable is only the product under test. Never switch controllers during a run/claim or resume another controller's claim. Worker PATH starts with the current tree's .fireweed/bin. Setup copies the invoker: the pin for workflow/drill scratch, the explicitly selected subject for isolated candidate CLI proofs. Use GOCACHE=/tmp/fireweed-rebuild-cache.
- Workers, separate restricted auditors and regeneration workers use exactly: codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c 'model_reasoning_effort="high"' -s workspace-write {prompt}. Stop on access failure; do not substitute model/effort. Use bounded access probes only as SYSTEM requires. Keep the accepted macOS sh fixture diagnosis; do not investigate it again.
- Follow CK1's normal authoring, fail-first/already_satisfied, audited eval_fix, restricted audit and frozen-phase gates. Context precedes each worker step; workers verify and return, never invoke lifecycle mutators, commit or edit run/phase/Provenance. The owning controller alone advances phases, commits and archives. Protected author/setup work is committed separately without active changes/runs/claims. No manual product Spec/Decision edits.
- CK2 uses a fresh separate candidate and actual affected CLI paths. Preserve history baseline be35552b706aa3000ff8f11306d129ff1d701b85 and pass an explicit nonempty resolved range to check-commits and CI, with hook doctor enabled. Keep exact candidate identity/ranges in committed native evidence, never extra binary hashes or manifests in planning.
- Before CK3's one native drill, prove retained-input isolation and a separate full verifier using only the approved dependency public interfaces as reconstruction inputs. Inspect live inputs and every worker command/full script/patch. If the unchanged pin cannot support required isolation/reporting, stop before the real invocation. Never supply old/deleted implementation, history, saved attempts, binaries, bootstrap or dependency internals. Input Git exposes no old history, remotes or alternates. Enforce checks plus workflow/runner/app evals and coverage with retained inputs unchanged.
- Review the native report, gaps and contradicted Decisions at CK4. No automatic gap fixes or fresh drill after exhausted attempts. Commit reviewed native evidence separately, preserve its ref and prior proofs, clean scratch worktrees and verify no live run/claim. Update PROGRESS and report the CK4 results required by ACCEPTANCE, including live unapproved count. Then stop: no later drills, regenerate change, controller promotion, delivery, publication or push.
```
