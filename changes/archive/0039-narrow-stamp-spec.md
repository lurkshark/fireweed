---
already_satisfied:
    - PRJ-006
    - PRJ-037
    - PRJ-038
    - PRJ-039
audit_rounds: 0
bases:
    PRJ-006: 51f3a69b
module: project
phase: applied
run: 20261028T052053-ca9c681bda62d2755ae6ec17
---



## Motivation

Complete the authorized remaining project split under SYSTEM's one-behavior
spec shape. Separate PRJ-006's write operation from its rejection and content
preservation rules, retaining the design's Change files canonical bases and
late ID assignment. Existing public evals already assert these behaviors.

## Interface



## Added

### PRJ-037
Stamp MUST reject Modified or Removed IDs absent from the projection preceding the target change.

### PRJ-038
Stamp MUST retain other frontmatter values and human section contents.

### PRJ-039
Stamp MUST succeed on an Added-only change without assigning spec IDs.

## Modified

### PRJ-006
Stamp MUST write each Modified or Removed item's canonical SpecHash from the projection preceding the target change into its bases.

## Removed



## Decisions



## Notes

Retarget existing Stamp evals, preserving every assertion and all 53 project
cases. StampPrecedingProjectionAndPreservation retains PRJ-006 and gains the
content-preservation ID. StampAddedOnlyAndAbsentID gains the Added-only and
absent-target IDs instead of PRJ-006, retaining both Modified and Removed
rejection assertions and its successful Added-only assertion. No new Test
functions, t.Run cases, fixture instances, helpers, product code or interface
changes. Keep every unrelated ID. Implementation may edit only this change file.
Existing Stamp implementation satisfies PRJ-006 and PRJ-037–039. Quick verify
passed all 53 project cases; no product code change was needed.

## Provenance

- evals 2026-10-28T05:21:26.54955Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=b26da634732c7f1c7dc600115582a63cb8f82e1fd2bd1d522d608fe837a04ba7 tests=199 failing=0 passing=199
- audit 2026-10-28T05:22:02.263848Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=0fb044cfaca1b9a844a25b4c86223177e52f2f35fb149eddf044e83822665c88 tests=0 failing=0 passing=0
- impl 2026-10-28T05:22:46.061699Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=07554a3249d8fca0ba70df7504135f297491465caa63458734a4b8351734d0e5 tests=199 failing=0 passing=199

## Audit

No findings. Checked PRJ-006, PRJ-037, PRJ-038, PRJ-039.
