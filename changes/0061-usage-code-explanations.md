---
bases:
    AST-003: 477c5b10
module: assets
---



## Motivation

The design and D006 require useful explanations for codes actually emitted.
Applied 0060 now emits U01 for existing invalid/unsupported requests and U02
for unknown explain/playbook lookups. Explain currently rejects those codes.

## Interface

No public interface change.

## Added



## Modified

### AST-003
Explain MUST describe the cause and a corrective action for C01 (invalid configuration), C04 (non-normative text), C05 (foreign or duplicate spec/edit identities), C06 (invalid tags), C07 (missing Motivation), C08 (missing/stale edit bases or absent targets), C09 (undeclared handles), X01 (boundary violations), X02 (missing/incomplete boundary hooks), X03 (reasoned boundary opt-out), X04 (unnecessary boundary enforcement), U01 (invalid arguments/options or unsupported commands), U02 (unknown lookup names) and P2-01 (missing genuine fail-first evidence). Explain MUST reject unknown codes.

## Removed



## Decisions



## Notes

Preserve all existing explanation and unknown-code assertions. Extend the
existing Explanations test with one case each for U01 and U02, checking their
specific cause and a corrective action using requirePattern. No new helper or
fixture infrastructure. U01 must cover invalid argument/option and unsupported
command errors; its repair should point to supported syntax/help. U02 must
cover an unknown lookup name and selecting a valid name. These two currently
unknown codes must genuinely fail before implementation. Only add explanation
text for these emitted codes; do not implement commands or explain unused
allocations. Controller owns specs, commits and archive.
