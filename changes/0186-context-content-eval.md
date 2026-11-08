---
bases: {}
eval_fix:
    - RUN-013
    - RUN-014
module: runner
phase: audit
run: 20261108T075131-cbdedb2ab362b8e130de4c91
---

## Motivation

The owner requested the context-heading eval correction identified in review
of rejected drill d63c16e. RUN-013 requires full workflow context, whose contents
are defined by FLOW-005 and the design's context command. Neither specifies the
literal headings checked by Test_RUN_013_RUN_014_CompleteSpecChange. Assert the
required content instead of arbitrary heading spelling, preserving full worker
context and restricted auditor isolation.

## Interface

Audited eval-only correction against the existing public contracts. Edit only
the target module evals and this change in evals; implementation changes only
this change. Preserve behavioral coverage and require all own evals to pass.

## Added

## Modified

## Removed

## Decisions

## Notes

Limit edits to the existing CompleteSpecChange eval and its fixture in
packages/runner/evals/next_test.go as needed for clear unique content markers.
Replace literal context-heading requirements with assertions for the actual
change, projected spec, interface contents, SYSTEM, matching phase playbook and
other relevant file contents requested by full context. Ensure these categories
remain distinguishable: finding new spec text inside the change alone must not
stand in for checking projected spec content. Use retained existing-spec content
and fixture markers as appropriate. Read assets public Playbook if needed for
a phase-specific expected playbook. Do not assert exact section titles/case,
formatting or serialization; do not simply delete the heading assertions and
leave any required content category unasserted.

Preserve phase roles, invocation cwd/bin/timeout behavior, ordered phase results,
verification and applied state, captured outputs, audit change/projected-spec/
eval/playbook contents and exclusion of implementation/SYSTEM secrets, and all
remaining lifecycle assertions. Reuse bounded fake agents and public interfaces;
no real agent in permanent evals, private imports or production changes.

Workers use the unchanged pinned controller with the current tree's .fireweed/bin
first on PATH and GOCACHE=/tmp/fireweed-rebuild-cache. Read pinned context before
each held step, run pinned verify --quick, wait for its terminal result and return.
Do not commit, invoke lifecycle mutators, edit run/phase/Provenance or protected
planning/tooling. The outside controller owns full phase verification, independent
restricted audit, commits, archive and claim release. The auditor edits Audit
only, checks the listed eval_fix IDs and preserves configured rounds. A passing
submission says No findings and lists the checked IDs.

This is one of exactly three owner-authorized eval_fix changes. No runtime,
interface/model, spec, Decision, configuration, hook, worker-policy or tooling
change is authorized. Keep every unrelated assertion and every covered ID.
Implementation edits only this change file; frozen evals stay unchanged.
Use current public contracts/evals, without bootstrap/rebuild-v1, saved
reconstruction source or other implementation oracles. Preserve prior evidence,
proof/attempt refs, bounds/counters, controller/recoveries and unrelated worktrees.
Verify and stop after all three changes; no drill, reconstruction import,
controller promotion, later work or push is authorized. The existing deadline
2026-11-10 22:51:05 UTC remains unchanged.
## Provenance
- evals 2026-11-08T08:03:53.984567Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6.1-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=6a9c2d506363e3539cc66efd37bc6a06701249541dc4e4f01ed3aa6454bbb864 tests=159 failing=0 passing=159
