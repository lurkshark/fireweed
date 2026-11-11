---
bases:
    FLOW-054: 7df3ffcd
    FLOW-055: 487dfc2d
    FLOW-056: 2afdfc82
    FLOW-057: 09d6cb13
    FLOW-058: 7360b3b6
    FLOW-059: cf833220
    FLOW-060: 25d56a42
module: workflow
phase: impl
run: 20261111T063722-da5edaeff080d91cf6f4fc27
---

## Motivation

The design promises a stable code for every problem and output that agents
can read. Workflow requires its lifecycle codes only as message prefixes,
and no workflow error carries a code that `contracts.ErrorCode` can read.
The runner passes these errors on, and the app recovers the code from the
text with a regular expression that takes the last code-shaped token, which
can be a wrapped cause or a spec ID rather than the lifecycle code
(APP-078, APP-085). State these codes through `contracts.ErrorCode`,
following SYSTEM.md's diagnostic-code convention.

## Interface

No signature changes. Errors that carry a code become (or wrap)
`contracts.CodedError`. Diagnostics are unchanged.

## Added

## Modified

### FLOW-054
Start MUST reject a dirty worktree or an ineligible queue with an error whose contracts ErrorCode is P01.

### FLOW-055
Verify and Done MUST reject an unsupported phase with an error whose contracts ErrorCode is P01.

### FLOW-056
Start and Current MUST report existing claim ownership or contention failures with an error whose contracts ErrorCode is P02.

### FLOW-057
Verify and Done MUST reject unapproved author, runtime or audit-state alterations with an error whose contracts ErrorCode is P03 before running hooks.

### FLOW-058
Start, Verify and Done MUST reject a stale base with an error whose contracts ErrorCode is P05, retaining the underlying cause in its chain.

### FLOW-059
During audit, Verify MUST apply Done's audit-submission checks. Verify and Done MUST reject malformed or incomplete submissions with an error whose contracts ErrorCode is P09.

### FLOW-060
Done MUST return its existing blocked-change error with contracts ErrorCode P06 while preserving the committed blocking state and claim release.

## Removed

## Decisions

## Rejected alternatives

A workflow spec item requiring every coded error to carry its code: the
same rule applies to every module, and copying it into each spec is what
C18 warns about. It is a SYSTEM.md convention instead.

## Notes

Only the way each code is stated changes; the conditions are the same.
Apply the convention to workflow's other coded errors too (for example
FLOW-011, FLOW-015, FLOW-061, FLOW-090 to FLOW-096, FLOW-125, FLOW-131).
Keep the existing message text, including leading codes and FLOW-133's
exact text, until 0222. For P05, P05 is the outer code and the stale-base
cause stays in the chain.

Evals: assert `contracts.ErrorCode` for each modified item, including one
P05 case whose cause has its own code.
## Provenance
- evals 2026-11-11T06:49:59.354718Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6.1-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=6a9c2d506363e3539cc66efd37bc6a06701249541dc4e4f01ed3aa6454bbb864 tests=182 failing=11 passing=171
- audit 2026-11-11T06:50:21.491022Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-astra -c model_reasoning_effort="medium" -s workspace-write {prompt} fireweed=0.1.0 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0
## Audit

No findings; checked FLOW-054, FLOW-055, FLOW-056, FLOW-057, FLOW-058, FLOW-059, FLOW-060
