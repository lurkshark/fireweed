---
already_satisfied:
    - APP-037
    - APP-078
    - APP-085
bases:
    APP-037: 9d3b7ee6
    APP-078: 094fe1fd
    APP-085: bf182e31
module: app
phase: audit
run: 20261111T083505-a3181a5d03759644a55b7091
---

## Motivation

The design promises a stable code for every problem. The app learns most
lower-module codes by parsing error text: `next` and `run` take the last
code-shaped token in the message (`wrappedCause`), other commands read a
leading token (`diagnosticPrefix`), and `projectFailureCode` runs
`project.Check` again to guess the code of a project error. The pattern
`[A-Z][A-Z0-9-]+:` also matches spec IDs, and taking the last match reports
a wrapped cause instead of the code the lifecycle assigned, such as P05
over its stale-base cause. The app drill report lists this as a gap. After
0209 to 0212, every coded lower-module error carries its code, so the app
can read it directly.

## Interface

No CLI or JSON changes. Diagnostic codes come from `contracts.ErrorCode`.

## Added

## Modified

### APP-037
Existing static failures exposed by the CLI MUST report the D006 code that the failing Diagnostic carries or that contracts.ErrorCode returns for the failing error: C01 configuration, C02 graph, C03 syntax, C04 normative text, C05 identities, C06 tags, C07 change structure, C08 bases, C09 handles, C12 metadata or C13 queue identity.

### APP-078
next failure diagnostics MUST report the code that contracts.ErrorCode
returns for the returned error, or P01 when it returns none, and include a
corrective action.

### APP-085
run failure diagnostics MUST report the code that contracts.ErrorCode
returns for the returned error, or P01 when it returns none, and include a
corrective action.

## Removed

## Decisions

## Rejected alternatives

Restricting the regular expression to the D006 shape as a separate fix
first: no current error text has been shown to produce a wrong code, so a
fix could not fail first, and this change removes the expression anyway.

## Notes

The existing evals for these IDs already pass, so they are listed in
already_satisfied; keep them passing. The implementation step replaces the
text parsing: take each code from `contracts.ErrorCode` with the command's
existing fallback, remove `wrappedCause`, and remove the `project.Check`
re-run from `projectFailureCode`. `diagnosticPrefix` may remain only to
strip a leading `<code>: ` equal to that code, so the code prints once, as
today; it must no longer choose the code.

Run this only after 0209 to 0212 apply. If a lower-module path still
returns a coded failure without `contracts.CodedError`, block the change
and name the path rather than keeping a text fallback.
## Provenance
- evals 2026-11-11T08:48:01.006913Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6.1-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=6a9c2d506363e3539cc66efd37bc6a06701249541dc4e4f01ed3aa6454bbb864 tests=323 failing=0 passing=323
