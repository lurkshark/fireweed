---
already_satisfied:
    - PRJ-003
    - PRJ-023
    - PRJ-024
    - PRJ-025
    - PRJ-026
audit_rounds: 0
bases:
    PRJ-003: c0a8d930
module: project
phase: audit
run: 20261028T035627-a8c9ede0a5c2baeb3762ba21
---



## Motivation

Continue the owner-authorized project spec split under SYSTEM's spec-shape
rule. Retain the design's Change files ordering and authoring behavior while
separating PRJ-003's queued skeleton from numbering and argument/collision
checks. Every rule is already implemented and asserted by the public evals.

## Interface



## Added

### PRJ-023
NewChange MUST use the next number after the largest pending or archived change number.

### PRJ-024
NewChange MUST reject a slug already present in the pending queue or archive without overwriting existing files.

### PRJ-025
NewChange MUST reject an unknown target module.

### PRJ-026
NewChange MUST reject an empty slug or a slug containing characters other than lowercase letters, digits and hyphens.

## Modified

### PRJ-003
NewChange MUST create a queued Markdown skeleton for an existing module with Motivation, Interface, Added, Modified, Removed and Decisions sections, containing no assigned spec IDs.

## Removed



## Decisions



## Notes

Retarget existing evals only, preserving every assertion and all 53 reported
project cases. NewChangeSkeletonAndNumbering keeps PRJ-003 and gains the
numbering ID. NewChangeRejectsDuplicatesWithoutOverwrite uses the collision
ID. NewChangeInvalidArguments covers module lookup and slug validation with
its existing unknown_module and uppercase/underscore/slash/empty_slug rows;
label these with the corresponding Added IDs, keeping both relevant IDs on
the table parent as necessary for coverage. No new Test functions, t.Run cases,
fixture instances, helpers, product code or interface changes are needed.
Implementation may edit only this change file. Preserve every unrelated ID.

## Provenance

- evals 2026-10-28T03:57:10.060361Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=b26da634732c7f1c7dc600115582a63cb8f82e1fd2bd1d522d608fe837a04ba7 tests=199 failing=0 passing=199
