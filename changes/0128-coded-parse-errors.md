---
audit_rounds: 0
module: contracts
phase: audit
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
Existing message text stays as it is; callers must stop depending on it.

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
