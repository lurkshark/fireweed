---
bases:
    PRJ-055: c323db61
    PRJ-056: 1f45a62f
module: project
phase: applied
run: 20261111T043001-941b76c82838a5e34dac4138
---

## Motivation

The design promises a stable code for every problem and output that agents
can read. Project states its codes two ways. PRJ-057, PRJ-058 and PRJ-059
require a code that `contracts.ErrorCode` reports, while PRJ-055 and PRJ-056
require the code only as the first characters of the error text. Most
project errors carry their code only in the text, so the app recovers it
with a regular expression or by running `project.Check` again (APP-037).
The app drill lists this as a gap, and the project drill's retry changed
error text to match it. SYSTEM.md's diagnostic-code convention makes the
code part of the error; state PRJ-055 and PRJ-056 that way.

## Interface

No signature changes. Errors that carry a code become (or wrap)
`contracts.CodedError`.

## Added

## Modified

### PRJ-055
An unknown module or an ID absent from both current module Specs and matching
archive edits MUST return an error whose contracts ErrorCode is U02 and whose
text names the rejected name.
A current ID with no matching archive edits MUST return an empty slice.
An ID in matching archive edits MUST remain queryable when it no longer exists
in current Specs.

### PRJ-056
Missing or malformed fireweed.yaml MUST produce an error whose contracts
ErrorCode is C01.
Malformed archived change documents or malformed Added/Modified item syntax
MUST produce an error whose contracts ErrorCode is C03 and whose text
identifies the archive path.

## Removed

## Decisions

## Rejected alternatives

A project spec item requiring every coded error to carry its code: the same
rule applies to every module, and copying it into each spec is what C18
warns about. It is a SYSTEM.md convention instead.

## Notes

Apply the SYSTEM.md diagnostic-code convention to every project error that
a spec assigns a code (PRJ-004, PRJ-024, PRJ-065, PRJ-066 and the Load and
Stamp failures), not only PRJ-055 and PRJ-056. Keep the existing message
text, including the leading code, until 0220.

Evals: in spec_history_test.go, replace the PRJ-055 and PRJ-056 prefix
checks (`U02: `, `C01: `, `C03: `) with `contracts.ErrorCode` checks, keeping
the checks for the rejected name and archive path.
## Provenance
- evals 2026-11-11T04:31:13.434594Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6.1-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=6a9c2d506363e3539cc66efd37bc6a06701249541dc4e4f01ed3aa6454bbb864 tests=134 failing=4 passing=130
- audit 2026-11-11T04:31:34.414337Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-astra -c model_reasoning_effort="medium" -s workspace-write {prompt} fireweed=0.1.0 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0
- impl 2026-11-11T05:17:32.584069Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6.1-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=5f856d6bcab308b9ff1a8a7c03b9d2547744112aba81dd30dea8f61ec8062b03 tests=933 failing=0 passing=933
## Audit

No findings; checked PRJ-055, PRJ-056
