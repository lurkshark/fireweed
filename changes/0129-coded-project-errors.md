---
audit_rounds: 0
blocked: |-
    impl attempts exhausted: P02: failing eval for APP-065: Test_APP_065_APP_066_APP_067_DoctorFailures
    B02: APP-103: 17 cases
    B02: APP-104: 19 cases
    B02: APP-103: 17 cases
    B02: APP-104: 19 cases
module: project
phase: impl
run: 20261031T033337-304bd319d25e3df24a16e023
---



## Motivation

App assigns C08 to Stamp failures by matching ` bases must be a mapping`, U02 to
missing pending changes by matching `pending change ... not found`, and project's
own queue code classifies item errors by message prefix. The design requires
stable diagnostic codes; these depend on unspecified wording, as the contracts
drill showed. Return coded errors using the contracts CodedError from
0128-coded-parse-errors.

## Interface

No new public functions. Stamp and the functions that look up pending changes
return errors that carry a contracts CodedError in their chain.

## Added

### PRJ-057
Stamp MUST return an error whose contracts ErrorCode is C08 when the change's bases frontmatter is not a mapping.

### PRJ-058
Stamp MUST return an error whose contracts ErrorCode is U02, with the requested change as subject, when no pending change matches it.

### PRJ-059
Load MUST return an error preserving the contracts ErrorCode C03 or C05 when a module SPEC.md fails to parse.

## Modified



## Removed



## Decisions



## Rejected alternatives



## Notes

Implementation should also replace the message-prefix check in
internal/queue/cause.go with contracts.ErrorCode.

## Provenance

- evals 2026-10-31T03:35:14.896424Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.1 playbook=b26da634732c7f1c7dc600115582a63cb8f82e1fd2bd1d522d608fe837a04ba7 tests=709 failing=7 passing=702
- audit 2026-10-31T03:35:43.744961Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.1 playbook=5cb9f9dac24b0b6bb1b970edd4fab4406d2b4d0435ddefc3ef59a2cfdd267f18 tests=0 failing=0 passing=0

## Audit

No findings. Checked PRJ-057, PRJ-058, PRJ-059.
