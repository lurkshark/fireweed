---
audit_rounds: 0
module: contracts
phase: applied
run: 20261031T004124-adafe55aac45dd6c68a386f3
---



## Motivation

The contracts drill reported a gap at CON-011/CON-012: dependents classify
ParseItems failures by the `duplicate item ID: ` and `malformed item heading: `
message prefixes, which no spec states, so a regenerated contracts module can
break C03/C05 diagnostics in project, checks and app. The design requires stable
diagnostic codes; give parse failures a code callers read without parsing text.

## Interface

Add a public coded error to contracts: a `CodedError` type with `Code` and
`Subject` string fields that wraps an optional cause, and a function
`ErrorCode(err error) (code, subject string, ok bool)` that finds the first
CodedError in an error chain, including through `fmt.Errorf("...: %w", err)`.
Existing message text stays as it is. Caller migrations are handled separately
by queued changes 0129, 0132 and 0133.

## Added

### CON-016
ErrorCode MUST return the code and subject of the first CodedError in an error chain, including one wrapped with %w, and MUST report not ok when the chain has none.

### CON-017
ParseItems MUST return a CodedError with code C05 and the duplicated ID or handle as subject when a Specs section repeats an item heading.

### CON-018
ParseItems MUST return a CodedError with code C03 when an item heading is malformed.

### CON-019
ParseDocument MUST return a CodedError with code C03 when frontmatter is missing, unterminated, invalid YAML or not a mapping.

## Modified



## Removed



## Decisions

CodedError exposes an optional Cause error. Error returns the cause's text when present and the code otherwise; Unwrap returns Cause.

## Rejected alternatives

Exported message-prefix constants: they keep text as the contract and still need
string matching at every caller.

## Provenance

- evals 2026-10-31T00:57:11.890544Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.1 playbook=b26da634732c7f1c7dc600115582a63cb8f82e1fd2bd1d522d608fe837a04ba7 tests=831 failing=16 passing=815
- audit 2026-10-31T00:57:42.94052Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.1 playbook=5cb9f9dac24b0b6bb1b970edd4fab4406d2b4d0435ddefc3ef59a2cfdd267f18 tests=0 failing=0 passing=0
- impl 2026-10-31T03:33:13.958651Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.1 playbook=07554a3249d8fca0ba70df7504135f297491465caa63458734a4b8351734d0e5 tests=831 failing=0 passing=831

## Audit

No findings. Checked CON-001, CON-002, CON-003, CON-004, CON-005, CON-006, CON-007, CON-008, CON-009, CON-010, CON-011, CON-012, CON-013, CON-014, CON-015, CON-016, CON-017, CON-018, CON-019.

## Questions

Resolved by the owner's 2026-10-30 instruction to run all queued changes with
explicit approval: the author repaired the facade to delegate to a disposable
stub before resuming implementation. Evals and assertions are unchanged; caller
migrations remain in 0129, 0132 and 0133. Resume the existing run and phase.
