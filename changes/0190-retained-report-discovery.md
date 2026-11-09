---
bases:
    FLOW-052: 7072ed88
module: workflow
phase: audit
run: 20261109T014702-e14683ab7f43e1843ca67070
---

## Motivation

App and runner users need to find the phase reports that Done retains locally
under the design's Provenance rule. The supervised workflow drill wrote those
reports at a different path and 17 downstream eval rows could not read them.
Make the lookup explicit in workflow and keep both artifacts when a module
name would collide with the phase cost snapshot.

## Interface

Keep the existing workflow API, checks.Report fields, cost snapshot location
and ordinary flat report names. The observable change is the collision-safe
location for a module name equal to costs without regard to case.

## Added

## Modified

### FLOW-052
Successful Done MUST retain verification reports and process output under
`.fireweed/runs/<run>/`.
Each evaluated module's report MUST be readable as a JSON `checks.Report`
at `<phase>/<module>.json` within that run directory.
For module names equal to `costs` without regard to case, the report MUST use
`<phase>/reports/<module>.json` and coexist with the phase's `costs.json`.

## Removed

## Decisions

## Notes

In the existing public workflow retention eval, keep all context, provenance,
and hook-output assertions. Check the exact ordinary report locations for
evals and impl using the run identity and the evaluated modules returned by
Done. Decode the reports as checks.Report and assert they retain the
phase-specific testcase IDs/results and hook output. Keep the app and runner
assertions intact.

Add a focused public FLOW-052 fail-first case using the existing lifecycle
fixture and a selected evaluated module named costs. Assert that the collision
report decodes at its distinct path with the expected mapped case and hook
output, and that the phase costs.json still decodes as the correct cost
snapshot. Also cover a case variant in a separate temporary project to check
case-insensitive selection without case-only names in the same project.
Confirm the baseline fails on the missing collision report after a successful
Done, rather than on fixture setup. Keep the eval step out of runtime code.
In impl, keep the change to workflow report path selection; reuse existing
report writing and avoid changing persisted case data or cost measurement.

Use the pinned controller and configured worker PATH and GOCACHE. Read context
before each held step and run pinned verify --quick before returning. The
outside controller owns full verification, phase commits and archive. Workers
never commit, call lifecycle mutators, or edit protected guidance or prior
evidence. Keep the separate restricted audit and existing bounds.
## Provenance
- evals 2026-11-09T01:56:12.541816Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6.1-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=6a9c2d506363e3539cc66efd37bc6a06701249541dc4e4f01ed3aa6454bbb864 tests=178 failing=3 passing=175
