---
eval_fix:
    - PRJ-016
module: project
phase: impl
run: 20261111T220010-7d6302a96fd7261e281b415c
---

## Motivation

PRJ-016 requires Check to report C08 for a Modified or Removed ID absent
from the preceding projection. Its eval requires two C08 diagnostics for
each case, one with the generic text `ID ALP-999 absent from preceding
projection` and one beginning `Modified ID` or `Removed ID`. The spec leaves
the number and wording open, and the project drill lists this as a gap.

## Interface

Audited eval_fix only. Edit only the project evals zone and this change.

## Added

## Modified

## Removed

## Decisions

## Notes

In `Test_PRJ_014_PRJ_015_PRJ_016_CheckDiagnostics` (queue_test.go), require
for each PRJ-016 case at least one non-warning C08 diagnostic whose message
names ALP-999. Leave the PRJ-014 and PRJ-015 cases as they are.

The project evals passed with `fireweed verify absent-edit-diagnostic-evals`.
## Provenance
- evals 2026-11-11T22:01:21.740696Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6.1-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=6a9c2d506363e3539cc66efd37bc6a06701249541dc4e4f01ed3aa6454bbb864 tests=135 failing=0 passing=135
- audit 2026-11-11T22:01:46.781763Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-astra -c model_reasoning_effort="medium" -s workspace-write {prompt} fireweed=0.1.0 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0
## Audit

No findings; checked PRJ-016
