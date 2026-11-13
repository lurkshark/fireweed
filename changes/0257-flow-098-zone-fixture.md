---
eval_fix: [FLOW-098]
module: workflow
---

## Motivation

Full implementation verification for 0258 rejected the FLOW-098 fixture
before it could measure costs. The fixture adds `internal/overlap.txt` to the
projected eval zone and also matches the disposable `internal/**` glob. The
design's disjoint eval/disposable zones rule prohibits that overlap, and
FLOW-098 can exercise projected zone selection with a disjoint file path.

## Interface

## Added

## Modified

## Removed

## Decisions

## Notes

The owner approved this audited eval fix as the prerequisite for resuming
0258. In `packages/workflow/evals/costs_test.go`, move the one-line fixture
file to a module-relative path outside both baseline globs. Update the
queued `set.evals` entry to select the new path only in the held projection.
Preserve every existing assertion and expected cost, including the projected
eval line, excluded protected and preserved files and symlink, unterminated
line, testcase counts, timings, and dependent module cost. Edit only the
workflow evals and this change file. The implementation phase archives the
change without a product edit.
