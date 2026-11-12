---
bases:
    FLOW-111: 1547e4e6
module: workflow
phase: applied
run: 20261111T094613-0a1a131293b1167588705e94
---

## Motivation

The design adds P08 to catch "a regression chain where each change fixes
the edge case the last one exposed". FLOW-111 approximates that with three
consecutive applied changes on one module. Replaying this repository's 200
archived changes, that rule warns on ten runs, six of them planned
module-by-module bootstrap work (for example 0015 to 0025 and 0030 to 0039),
and it says nothing when an unrelated change sits inside a real chain.
Linking changes by the spec IDs they rework, within the last ten applied
changes, warns instead on two clusters: 0056, 0059, 0061, 0067, 0068, 0070
and 0073 kept reworking the same explanation specs, and 0079, 0084, 0086,
0094 and 0096 kept reworking runner execution. Those are the chains the
design describes. Grouping by module also encourages authors to interleave
modules to avoid the warning, which hides the signal.

## Interface

None. The P08 trigger and message change. The diagnostic stays a Status
warning.

## Added

## Modified

### FLOW-111
Status MUST warn with P08 when the most recently applied change ends a chain of three or more applied changes in which each modifies, removes, fixes or eval-fixes a spec ID that its predecessor added or modified. A change's predecessor is the latest of the ten applied changes before it that added or modified such an ID. The warning MUST name the chain's change files and shared IDs.

## Removed

## Decisions

## Rejected alternatives

Requiring the chain's changes to be adjacent: it misses chains with an
unrelated change in between, which is how the replay's real chains looked.
An unbounded window: it links rework months apart and warned on 49 of the
200 replayed changes. Counting only unapproved changes: P07 already bounds
that work to five changes, and a chain across a merge still deserves review.

## Notes

"Added" means an ID in the change's Added section after apply (archived
changes hold assigned IDs), and "modified" an ID in its Modified section.
"Fixes" and "eval-fixes" mean the `fixes` and `eval_fix` lists. Use the
existing applied order Status already computes.

Evals: replace the module-run cases in
`Test_FLOW_109_FLOW_111_StatusProvenanceDriftAndRecentModuleWarning` with
fixtures for a three-change chain (warns, naming files and IDs), three
changes on one module with no shared IDs (no warning), a chain with an
unrelated change between links (warns), and a link to a change more than
ten applied changes back (no warning). Run after 0215 and 0216, so app's
status eval already triggers P08 under this rule.

`go test ./packages/workflow/evals -run '^Test_FLOW_(109_StatusProvenanceDrift|111_StatusReworkChainWarning)$' -count=1 -v`
confirmed that FLOW-109 still passes. The current P08 message omits the chain
files and IDs, misses the chain across unrelated changes, and warns for both
independent same-module changes and a chain broken by a newer predecessor.
`fireweed verify rework-chain-warning` accepted the FLOW-111 fail-first
evidence. The report recorded 183 passing cases and five expected failures.

The P08 trigger lives in the implementation-writable file
`packages/workflow/internal/lifecycle/status.go`.
## Provenance
- evals 2026-11-11T09:56:48.253732Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6.1-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=6a9c2d506363e3539cc66efd37bc6a06701249541dc4e4f01ed3aa6454bbb864 tests=188 failing=5 passing=183
- audit 2026-11-11T09:57:07.670211Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-astra -c model_reasoning_effort="medium" -s workspace-write {prompt} fireweed=0.1.0 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0
- impl 2026-11-11T10:24:10.413494Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6.1-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=5f856d6bcab308b9ff1a8a7c03b9d2547744112aba81dd30dea8f61ec8062b03 tests=678 failing=0 passing=678
## Audit

No findings; checked FLOW-111
