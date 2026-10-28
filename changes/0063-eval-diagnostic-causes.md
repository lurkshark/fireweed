---
audit_rounds: 0
bases:
    CHK-002: 01780e0f
    CHK-025: 1e3a3c61
    CHK-026: 449e8ad8
    CHK-027: c9b01661
    CHK-028: 2ae67c70
    CHK-036: ef057829
    CHK-037: 3681b776
    CHK-038: 888600aa
module: checks
phase: audit
run: 20261028T084945-cd3602bd735f5ca5842b722d
---



## Motivation

The design requires stable diagnostic causes, and D006 allocates C10 to coverage,
C11 to eval-hook/JUnit contracts, E01 to enforced failing results and P2-01 to
missing genuine fail-first evidence. These checks already reject the inputs but
return uncoded errors, so the composing CLI cannot preserve the cause. Complete
these existing checks as one coherent diagnostic migration.

## Interface

Keep public signatures and types unchanged. Error text for the assigned causes
begins with the code followed by a colon, retaining the original cause text.

## Added

### CHK-046
Evaluate MUST return C11-prefixed errors for existing eval-hook setup failures, including missing or invalid hook commands.

### CHK-047
Evaluate MUST return a U02-prefixed error for an unknown requested module.

### CHK-048
Evaluate MUST return C03-prefixed errors for malformed module item syntax and C05-prefixed errors for duplicate module item identities.

## Modified

### CHK-002
RequireCoverage MUST cover each current spec except review or quarantine with a non-skipped referencing case. Missing coverage MUST return a C10-prefixed error.

### CHK-025
RequireFailFirst MUST find a failed or errored referencing case for every requested enforced ID except already_satisfied; passing, skipped, empty or unrelated failures do not qualify. Missing genuine fail-first evidence MUST return a P2-01-prefixed error.

### CHK-026
RequirePass MUST reject failed or errored cases referencing any non-quarantined spec with an E01-prefixed error.

### CHK-027
ParseJUnit MUST reject malformed XML with a C11-prefixed error.

### CHK-028
ParseJUnit MUST reject testcases with missing or unknown spec IDs with a C11-prefixed error.

### CHK-036
Evaluate MUST fail with a C11-prefixed error when the eval hook cannot launch or exceeds its timeout.

### CHK-037
Evaluate MUST fail with a C11-prefixed error when the eval hook does not write its JUnit report.

### CHK-038
Evaluate MUST reject a zero exit with mapped failures or a nonzero exit without mapped failures with a C11-prefixed error, while accepting a nonzero exit with mapped failures as valid report evidence.

## Removed



## Decisions



## Notes

Reuse existing public JUnit, coverage, fail-first, pass and hook scenarios for
exact error-prefix assertions while preserving all existing status, cause,
report, exemption, mapping and timeout assertions and ID coverage. Add only a
representative public-input case where a new named cause lacks one, using the
existing fixture helpers; no matrix expansion or instrumentation. Changed IDs
must genuinely fail before implementation.

Prefix errors only on existing rejection paths; preserve their cause text,
wrapping, report contents and all success/rejection conditions. ParseJUnit's
other existing XML-structure errors also use C11; Evaluate preserves those
codes and uses C11 for existing setup/execution/report failures. Avoid duplicate
prefixes when propagating errors. Module item parsing retains static causes
C03/C05 and unknown module requests use U02. Do not change Zones or Boundaries,
which already emit their assigned diagnostics, or shared hook behavior for
boundaries. No new validations, unsupported features, speculative errors,
public interfaces or general error framework. Controller owns specs and archive.

## Provenance

- evals 2026-10-28T08:53:24.929099Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=b26da634732c7f1c7dc600115582a63cb8f82e1fd2bd1d522d608fe837a04ba7 tests=163 failing=17 passing=146
