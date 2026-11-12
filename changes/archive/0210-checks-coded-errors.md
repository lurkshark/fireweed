---
already_satisfied:
    - CHK-048
bases:
    CHK-002: 204cb146
    CHK-026: b775b51f
    CHK-027: cd149f35
    CHK-028: d008b19c
    CHK-036: 7479fcd2
    CHK-037: a1b5803c
    CHK-038: 7f0ea57a
    CHK-046: 6b701504
    CHK-047: 501d63ae
    CHK-048: 3b0dee08
module: checks
phase: applied
run: 20261111T051736-9640fdfd1590a860b50eb3a9
---

## Motivation

The design promises a stable code for every problem and output that agents
can read. Checks requires most of its error codes only as message prefixes,
so callers must parse text to learn the code. Workflow and the app pass
these errors on, and the app recovers codes with a regular expression that
also matches spec IDs (APP-078, APP-085). State these codes through
`contracts.ErrorCode`, following SYSTEM.md's diagnostic-code convention.

## Interface

No signature changes. Errors that carry a code become (or wrap)
`contracts.CodedError`. Diagnostics are unchanged.

## Added

## Modified

### CHK-002
RequireCoverage MUST cover each current spec except review or quarantine with a non-skipped referencing case. Missing coverage MUST return an error whose contracts ErrorCode is C10.

### CHK-026
RequirePass MUST reject failed or errored cases referencing any non-quarantined spec with an error whose contracts ErrorCode is E01.

### CHK-027
ParseJUnit MUST reject malformed XML with an error whose contracts ErrorCode is C11.

### CHK-028
ParseJUnit MUST reject, with an error whose contracts ErrorCode is C11, a testcase that references no supplied spec ID or that contains a token with a supplied spec's prefix, matched under the ID rules, naming no supplied spec. Text without such a prefix, including numbered labels, MUST NOT count as an ID reference.

### CHK-036
Evaluate MUST fail with an error whose contracts ErrorCode is C11 when the eval hook cannot launch or exceeds its timeout.

### CHK-037
Evaluate MUST fail with an error whose contracts ErrorCode is C11 when the eval hook does not write its JUnit report.

### CHK-038
Evaluate MUST reject a zero exit with mapped failures or a nonzero exit without mapped failures with an error whose contracts ErrorCode is C11, while accepting a nonzero exit with mapped failures as valid report evidence.

### CHK-046
Evaluate MUST return errors whose contracts ErrorCode is C11 for existing eval-hook setup failures, including missing or invalid hook commands.

### CHK-047
Evaluate MUST return an error whose contracts ErrorCode is U02 for an unknown requested module.

### CHK-048
Evaluate MUST return errors whose contracts ErrorCode is C03 for malformed module item syntax and C05 for duplicate module item identities.

## Removed

## Decisions

## Rejected alternatives

A checks spec item requiring every coded error to carry its code: the same
rule applies to every module, and copying it into each spec is what C18
warns about. It is a SYSTEM.md convention instead.

## Notes

Only the way each code is stated changes; the conditions are the same.
CHK-028 here is the text 0204 introduces. Keep the existing message text,
including the leading code, until 0221.

Evals: change `requireDiagnosticError` and similar helpers to check
`contracts.ErrorCode` and the cause text instead of a leading label. CHK-043's
rule that the boundary error carries no leading X02 is unchanged.

Focused Go evals failed on missing structured codes for CHK-002, CHK-026,
CHK-027, CHK-028, CHK-036, CHK-037, CHK-038, CHK-046 and CHK-047. Both
CHK-048 cases passed because Evaluate already preserves the C03 and C05
errors from contracts.ParseItems. The author needs to mark CHK-048
`already_satisfied` before claiming the updated change. The active claim
locks the approved author metadata.

Both CHK-055 recovery scenarios now run in their parent test with the same
assertions. `fireweed verify checks-coded-errors` recorded 14 CHK-055 cases
and cleared B02. Its only blocker was P2-01 for CHK-048. The other nine
modified IDs had genuine fail-first evidence.

The missing JUnit, coverage and pass codes can be implemented in
`packages/checks/internal/engine/junit.go`. The missing Evaluate codes can
be implemented in `packages/checks/internal/engine/hooks.go`.
## Provenance
- evals 2026-11-11T05:55:57.461239Z agent=manual fireweed=0.1.0 playbook=6a9c2d506363e3539cc66efd37bc6a06701249541dc4e4f01ed3aa6454bbb864 tests=128 failing=19 passing=109
- audit 2026-11-11T05:56:20.925122Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-astra -c model_reasoning_effort="medium" -s workspace-write {prompt} fireweed=0.1.0 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0
- impl 2026-11-11T06:32:46.315707Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6.1-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=5f856d6bcab308b9ff1a8a7c03b9d2547744112aba81dd30dea8f61ec8062b03 tests=798 failing=0 passing=798
## Audit

No findings; checked CHK-002, CHK-026, CHK-027, CHK-028, CHK-036, CHK-037, CHK-038, CHK-046, CHK-047, CHK-048
