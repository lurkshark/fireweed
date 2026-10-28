---
already_satisfied:
    - PRJ-027
    - PRJ-028
    - PRJ-029
    - PRJ-030
    - PRJ-031
audit_rounds: 0
bases: {}
module: project
phase: impl
run: 20261028T050919-18e1395516a502825fd4fe03
---



## Motivation

Continue the authorized PRJ-004 split under SYSTEM's one-observable-behavior
spec shape, retaining the design's Change files queue and document rules.
Give Queue's existing behaviors dedicated IDs before narrowing PRJ-004 to
change-argument resolution in the next change. Existing assertions already
cover these rules, so implementation stays unchanged.

## Interface



## Added

### PRJ-027
Queue MUST return pending NNNN-slug.md changes in numeric-then-slug order, allowing tied numbers and excluding archives.

### PRJ-028
Queue MUST return each pending change's module, frontmatter and parsed Added, Modified and Removed items.

### PRJ-029
Queue MUST reject a malformed pending change document.

### PRJ-030
Queue MUST reject a pending change naming an unknown module.

### PRJ-031
Queue MUST reject pending changes that share a slug.

## Modified



## Removed



## Decisions



## Notes

Retarget existing queue evals while preserving PRJ-004 for this change.
QueueOrderContentAndArchive gains the order and payload IDs;
QueueRejectsInvalidDocuments gains the malformed/module IDs with corresponding
row labels; TargetAliasesAndAmbiguousSlug gains the slug ID on its existing
ambiguous-queue assertion. Retain every assertion and all 53 project cases.
No new Test functions, t.Run cases, fixture instances, helpers, product code
or interface edits. The next change narrows PRJ-004 only after these rules
have IDs. Implementation may edit only this change file.

## Provenance

- evals 2026-10-28T05:10:10.973452Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=b26da634732c7f1c7dc600115582a63cb8f82e1fd2bd1d522d608fe837a04ba7 tests=199 failing=0 passing=199
- audit 2026-10-28T05:10:39.709375Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=0fb044cfaca1b9a844a25b4c86223177e52f2f35fb149eddf044e83822665c88 tests=0 failing=0 passing=0

## Audit

No findings. Checked IDs: PRJ-027, PRJ-028, PRJ-029, PRJ-030, PRJ-031.
