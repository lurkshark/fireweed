---
already_satisfied:
    - FLOW-052
bases:
    FLOW-052: 2d9e57ca
module: workflow
---

## Motivation

The owner approved clarifying FLOW-052 before workflow regeneration. checks.Report carries evaluated modules' process output, while checks.Boundaries exposes diagnostics and an error without raw stdout/stderr. State the existing per-module eval-report retention contract precisely; adding boundary output transport is outside this prerequisite.

## Interface

No public API changes.

## Added



## Modified

### FLOW-052
Successful Done MUST retain evaluated modules' reports, including their eval-hook stdout and stderr, under `.fireweed/runs/<run>/`. Each evaluated module's report MUST be readable as a JSON `checks.Report` at `<phase>/<module>.json` within that run directory. For module names equal to `costs` without regard to case, the report MUST use `<phase>/reports/<module>.json` and coexist with the phase's `costs.json`.


## Removed

## Decisions

## Notes

This narrows the ambiguous reference to process output to output present in each evaluated module's checks.Report. Existing report-path, output and costs-name coexistence assertions remain authoritative; reuse them and add no boundary-output API or new logging infrastructure. No runtime change is intended.
Owner authorized these five regeneration prerequisites and their necessary module-local changes. Stop after this queue; no regeneration or drill is authorized. Follow SYSTEM.md and the controller context. Workers must not read docs/rebuild/PROGRESS.md or ACCEPTANCE.md. Preserve unrelated assertions and retained interfaces after evals. Use existing public-interface fixtures and focused assertions; do not redesign helpers or add a test matrix. Evals and implementation workers run fireweed context for this change before each step and fireweed verify for verification. Restricted auditors follow only the supplied audit playbook and inputs. Only the controller advances phases and commits.
