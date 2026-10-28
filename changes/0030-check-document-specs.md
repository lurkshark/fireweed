---
already_satisfied:
    - PRJ-008
    - PRJ-009
    - PRJ-010
    - PRJ-011
    - PRJ-012
audit_pending:
    - PRJ-009
    - PRJ-010
audit_rounds: 1
bases: {}
module: project
phase: impl
run: 20261028T033537-8ab523401bed5218f526ffc1
---



## Motivation

SYSTEM's spec-shape rule requires one observable behavior per item. Begin the
approved PRJ-007 split by giving its document validation rules separate IDs,
preserving the design's Specs and Change files obligations. Existing behavior
and assertions already satisfy these rules; keep PRJ-007 until the remaining
rules have IDs in the following change.

## Interface



## Added

### PRJ-008
Check MUST reject a pending change without a nonempty Motivation section.

### PRJ-009
Check MUST reject a module spec or Added/Modified spec whose body contains no normative keyword.

### PRJ-010
Check MUST reject spec tags other than review or quarantine with a nonempty reason.

### PRJ-011
Check MUST reject duplicate spec IDs within a module's Specs section.

### PRJ-012
Check MUST reject module spec IDs and change edit IDs whose prefix belongs to another module.

## Modified



## Removed



## Decisions



## Notes

Retarget the existing CheckValidAndDiagnostics table cases to the corresponding
Added IDs, retaining PRJ-007 and all existing assertions, including its current
C01 assertion until that spec is narrowed in the next change. No new Test
functions, t.Run cases or fixture instances. Keep 53 reported project cases.
The existing missing_motivation, nonnormative_module/nonnormative_change,
invalid_review_tag/invalid_quarantine_tag, duplicate_module_id and
foreign_module_id/foreign_change_id cases provide the assertions. Only adjust
case names/ID labels as needed for retargeting; retain every other test and
behavior assertion. No product code, helper, fixture or interface changes are
needed.
Implementation may edit only this change file.

Audit correction: within the existing nonnormative_change and
invalid_review_tag cases, reuse each fixture after its original assertion to
check a non-normative Modified item and an unsupported tag with a reason.
Keep the original assertions and the 53 reported cases.

## Provenance

- evals 2026-10-28T03:36:55.429843Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=b26da634732c7f1c7dc600115582a63cb8f82e1fd2bd1d522d608fe837a04ba7 tests=199 failing=0 passing=199
- audit 2026-10-28T03:38:01.799771Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=0fb044cfaca1b9a844a25b4c86223177e52f2f35fb149eddf044e83822665c88 tests=0 failing=0 passing=0
- evals 2026-10-28T03:39:20.994902Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=b26da634732c7f1c7dc600115582a63cb8f82e1fd2bd1d522d608fe837a04ba7 tests=199 failing=0 passing=199
- audit 2026-10-28T03:39:46.545384Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=0fb044cfaca1b9a844a25b4c86223177e52f2f35fb149eddf044e83822665c88 tests=0 failing=0 passing=0

## Audit

No findings. Checked PRJ-009 and PRJ-010.
