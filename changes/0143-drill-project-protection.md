---
bases: {}
fixes:
    - RUN-052
    - RUN-053
    - RUN-055
module: runner
---

## Motivation

DR1 requires approved-zone disposable deletion and unchanged retained inputs
for `fireweed drill <module>`. The design's Modules/zone map gives protected
paths precedence over disposable globs, including fireweed.yaml's project-level
protected list. A public Drill fixture on 0141's implementation declared
packages/alpha/internal/project-protected.txt protected: Drill returned pass,
but the retained rebuilt ref had deleted that file. Its reconstruction input
was also lost. The drill code uses module zones without project protection.
Correct that existing zone obligation and its matching import check before
candidate verification; no new behavior or public API is needed.

## Interface

## Added

## Modified

## Removed

## Decisions

## Notes

Author a focused public regression using the existing deterministic Drill
fixtures. A project protected glob overlapping the target's disposable glob
must retain the file in the rebuilt verification ref and provide it unchanged
as a retained reconstruction input. A new file matching that protected glob
must be rejected as an unauthorized reconstruction edit, without import or a
rebuilt ref. Keep declared dependency interface files available read-only even
when project protection also matches them; protection does not erase their
interface role. Reuse existing own/all-transitive-dependent eval, native report
and cleanup assertions rather than duplicating the entire drill matrix.

These are RUN-052/RUN-053/RUN-055 corrections, not a new zone system. Compose
the approved Snapshot, existing public glob/zone and checks APIs consistently
across deletion, input selection and import validation; retain stricter module
zones and existing built-in protection rules. No public interface, spec text,
configuration, bound, recipe or other module implementation may change. No
speculative platform/binary probes, real agents or real module drills in evals.

Each listed ID needs genuine behavioral fail-first and preserved coverage.
Normal separate restricted audit, frozen phase zones and unchanged controller
bounds apply. Workers read context before each step and never commit. Passing
Audit submissions use No findings and all listed Checked IDs; findings quote
an exact clause with a separator and explanation.

For worker iteration, use focused public drill tests for these IDs and affected
existing drill behavior. Automatic next/done still perform the unchanged full
own/dependent/boundary phase checks; do not skip those, call done manually, commit
or alter hooks/bounds. A pending Audit section is preseeded so the restricted
review can replace only its body without inserting surrounding whitespace.

## Audit

Pending independent audit.
