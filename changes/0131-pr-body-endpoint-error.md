---
audit_rounds: 0
module: workflow
phase: impl
run: 20261031T044908-d2a4d72ed9377bd1b5108881
---



## Motivation

App pr-body decides between exit 2 with configuration advice and exit 1 with
range advice by matching `load endpoint project at ` in the PRBody error. The
design's distinct exit codes (2 for configuration errors) then depend on
unspecified message text. Make the endpoint configuration failure identifiable.

## Interface

Add `IsEndpointConfigError(err error) bool` to the workflow public API.

Keep the public facade delegating to disposable implementation. During evals,
add the implementation stub in a new internal file so the impl step can replace
it without editing the frozen interface (the wiring repair from 0128).

## Added

### FLOW-126
PRBody MUST return an error for which IsEndpointConfigError reports true when the project at the range endpoint fails to load.

### FLOW-127
IsEndpointConfigError MUST report false when PRBody fails because the range cannot be resolved.

## Modified



## Removed



## Decisions



## Rejected alternatives



## Provenance

- evals 2026-10-31T05:02:54.54355Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.1 playbook=b26da634732c7f1c7dc600115582a63cb8f82e1fd2bd1d522d608fe837a04ba7 tests=510 failing=2 passing=508
- audit 2026-10-31T05:03:26.541278Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.1 playbook=5cb9f9dac24b0b6bb1b970edd4fab4406d2b4d0435ddefc3ef59a2cfdd267f18 tests=0 failing=0 passing=0

## Audit

Checked IDs: FLOW-126, FLOW-127
No findings
