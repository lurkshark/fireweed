---
bases: {}
eval_fix:
    - FLOW-115
    - FLOW-124
    - FLOW-125
module: workflow
phase: applied
run: 20261108T072748-cb8f72c885753d9412fa663a
---

## Motivation

The owner requested correction of the three eval/contract mismatches exposed by
review of rejected drill d63c16e. FLOW-115 requires explicit empty-selection
reporting and FLOW-125 permits a valid empty range. The current empty-range
assertion rejects the compliant sentence "No selected changes." because it
recognizes only "none" or "no change". Correct that assertion without prescribing
new output wording or relaxing explicit-empty-selection behavior.

## Interface

Audited eval-only correction against the existing public contracts. Edit only
the target module evals and this change in evals; implementation changes only
this change. Preserve behavioral coverage and require all own evals to pass.

## Added

## Modified

## Removed

## Decisions

## Notes

Limit edits to the existing PRBody empty/artifact/range eval in
packages/workflow/evals/pr_body_test.go and a small local helper only if useful.
Accept the observed compliant "No selected changes." wording as well as the
previously accepted explicit-empty forms. Blank output and a report that does
not explicitly identify an empty selection must still fail. Do not reduce this
to merely nonempty output. Attribute the explicit-empty assertion to FLOW-115
while preserving the existing FLOW-124/FLOW-125 mappings and assertions.

Keep binary and strictly-over-256KB artifact path/size assertions, exact/small
text exclusions, B04 diagnostic/body/warning checks and counts, no B04 for an
empty range, successful valid-empty reporting and invalid-range P01/cause checks.
Reuse existing public PRBody fixtures. Do not introduce output-wording specs,
observer-calibration matrices, fake runtime behavior or implementation edits.

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
- evals 2026-11-08T07:34:49.588527Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6.1-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=6a9c2d506363e3539cc66efd37bc6a06701249541dc4e4f01ed3aa6454bbb864 tests=170 failing=0 passing=170
- audit 2026-11-08T07:35:22.332268Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-astra -c model_reasoning_effort="medium" -s workspace-write {prompt} fireweed=0.1.0 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0
- impl 2026-11-08T07:51:12.216261Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6.1-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=5f856d6bcab308b9ff1a8a7c03b9d2547744112aba81dd30dea8f61ec8062b03 tests=652 failing=0 passing=652
## Audit

No findings; checked FLOW-115, FLOW-124, FLOW-125
