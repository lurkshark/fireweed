---
already_satisfied:
    - PRJ-004
audit_rounds: 0
bases:
    PRJ-004: 93a0113f
module: project
phase: impl
run: 20261028T051148-62cbc9ce0b98bd1633865d8e
---



## Motivation

Finish PRJ-004's authorized split under SYSTEM's spec-shape rule. PRJ-027–031
now retain Queue's ordering, payload and rejection rules. Narrow PRJ-004 to
the design's CLI change-argument resolution behavior used by Project and
Stamp, retaining all existing assertions and behavior.

## Interface



## Added



## Modified

### PRJ-004
Project and Stamp MUST resolve a change argument from its unique slug, queued basename or repository-relative queued path, rejecting missing or ambiguous targets.

## Removed



## Decisions



## Notes

Retain PRJ-004 on TargetAliasesAndAmbiguousSlug, which already asserts all three
aliases and missing/ambiguous targets through Project and Stamp. Remove
PRJ-004 from QueueOrderContentAndArchive and QueueRejectsInvalidDocuments;
their PRJ-027–030 IDs now own those assertions. Preserve PRJ-031 on the
ambiguous-slug case and every other assertion/ID. Keep all 53 project cases.
Only rename existing tests/labels; no new Test functions, t.Run cases,
fixtures, helpers, product code or interface edits. Implementation may edit
only this change file.

## Provenance

- evals 2026-10-28T05:12:51.311432Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=b26da634732c7f1c7dc600115582a63cb8f82e1fd2bd1d522d608fe837a04ba7 tests=199 failing=0 passing=199
- audit 2026-10-28T05:13:23.263319Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=0fb044cfaca1b9a844a25b4c86223177e52f2f35fb149eddf044e83822665c88 tests=0 failing=0 passing=0

## Audit

No findings. Checked: PRJ-004.
