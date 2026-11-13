---
eval_fix:
    - FLOW-098
module: workflow
phase: impl
run: 20261113T232716-79a22fb9aca2aea6c7a17675
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

The fixture now puts the one-line file at `packages/alpha/projected-eval.txt`.
The queued `set.evals` selects it by its exact module-relative path.
## Audit

No findings; checked FLOW-098

## Provenance

- evals 2026-11-13T23:31:20.953704Z agent=codex fireweed=0.1.0 playbook=6a9c2d506363e3539cc66efd37bc6a06701249541dc4e4f01ed3aa6454bbb864 tests=200 failing=0 passing=200
- audit 2026-11-13T23:31:44.418291Z agent=codex fireweed=0.1.0 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0

