---
audit_rounds: 0
bases:
    AST-026: ded627c6
module: assets
phase: audit
run: 20261028T204148-2d902c629e93cf96369b58aa
---



## Motivation

The design requires stable diagnostics with accurate cause and corrective-action
explanations. The applied doctor backend now emits D007 C14–C16; explain must
cover those real causes before doctor CLI wiring. Existing init also emits P01
for operational setup failures, while its current explanation only names
lifecycle/new-change. Extend that existing explanation without changing code
assignments or adding speculative behavior.

## Interface



## Added

### AST-027
Explain MUST describe the cause and a corrective action for C14 (doctor hook/canary setup, detection or recovery failures), C15 (agent probe failures or usage limits, and the shared-auditor-command warning) and C16 (missing/stale managed guidance or playbook embedded_base metadata).

## Modified

### AST-026
Explain MUST describe the cause and a corrective action for P01 (operational preconditions, including lifecycle, new-change and init), P02 (claims), P03 (unapproved state alteration), P04 (approval warning), P05 (stale bases), P06 (blocked change), P07 (unapproved-work bound), P09 (audit submissions) and Z01 (forbidden phase-zone edits).

## Removed



## Decisions



## Rejected alternatives



## Notes

Use existing public assets.Explain and existing explanation test conventions.
Keep this content change limited to explanations for emitted causes. C14 guidance
should direct repair of the named hook/template/canary and rerunning doctor;
C15 should cover command/model/login failure, usage limits and configuring an
independent audit command for the warning; C16 should point to agents-md --write
or reviewing/updating a playbook override's embedded_base against the current
embedded playbook. P01 must retain its prior lifecycle/new-change causes and
cover existing init conflicts or filesystem/setup preconditions with an action.
Preserve every prior explanation assertion and interface. No new diagnostics,
behavior, helper framework or platform matrices. Each assigned ID needs genuine
runtime fail-first evidence; unknown-code errors qualify for the new doctor ID,
while init-specific P01 wording fails against the current explanation.
Workers never commit; the controller owns phase state, specs and archive.

Audit phase: use only the change, projected spec, evals, retained public interfaces
and governing audit rules. Do not read implementation files or implementation
history. Review only the assigned IDs and report exact missing assertion clauses;
helper/style critiques and unnamed cases are outside scope.

## Provenance

- evals 2026-10-28T20:45:51.8091Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=b26da634732c7f1c7dc600115582a63cb8f82e1fd2bd1d522d608fe837a04ba7 tests=389 failing=6 passing=383
