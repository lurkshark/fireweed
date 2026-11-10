---
bases:
    FLOW-111: 1547e4e6
module: workflow
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
