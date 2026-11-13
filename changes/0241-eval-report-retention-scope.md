---
already_satisfied:
    - FLOW-052
bases:
    FLOW-052: 2d9e57ca
module: workflow
phase: impl
run: 20261113T024216-8a00b31e6c6414295a863990
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
## Provenance
- evals 2026-11-13T02:50:39.843472Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6.1-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=6a9c2d506363e3539cc66efd37bc6a06701249541dc4e4f01ed3aa6454bbb864 tests=200 failing=0 passing=200
- audit 2026-11-13T02:51:01.078976Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-astra -c model_reasoning_effort="medium" -s workspace-write {prompt} fireweed=0.1.0 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0
## Audit

No findings; checked FLOW-052
