---
bases: {}
fixes:
    - RUN-052
    - RUN-053
    - RUN-055
module: runner
phase: applied
run: 20261101T024402-5b20ef3a216ee76e284f30bf
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

Fail-first: focused public Drill tests on the current implementation showed
RUN-052/RUN-053 losing the project-protected file from both the rebuilt ref and
reconstruction input, and RUN-055 accepting a new protected file and retaining
a rebuilt ref. The protected dependency interface was present and read-only.
The implementation-writable correction point is
`packages/runner/internal/drill/drill.go`; no stub is needed.

## Audit

No findings; checked RUN-052, RUN-053, RUN-055
## Provenance
- evals 2026-11-01T03:01:24.623335Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=6a9c2d506363e3539cc66efd37bc6a06701249541dc4e4f01ed3aa6454bbb864 tests=132 failing=3 passing=129
- audit 2026-11-01T03:01:55.360337Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0
- impl 2026-11-01T03:22:41.869683Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=5f856d6bcab308b9ff1a8a7c03b9d2547744112aba81dd30dea8f61ec8062b03 tests=423 failing=0 passing=423
