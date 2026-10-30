---
module: contracts
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

### new:error-code-chain
ErrorCode MUST return the code and subject of the first CodedError in an error chain, including one wrapped with %w, and MUST report not ok when the chain has none.

### new:item-duplicate-code
ParseItems MUST return a CodedError with code C05 and the duplicated ID or handle as subject when a Specs section repeats an item heading.

### new:item-malformed-code
ParseItems MUST return a CodedError with code C03 when an item heading is malformed.

### new:document-parse-code
ParseDocument MUST return a CodedError with code C03 when frontmatter is missing, unterminated, invalid YAML or not a mapping.

## Modified

## Removed

## Decisions

## Rejected alternatives

Exported message-prefix constants: they keep text as the contract and still need
string matching at every caller.
