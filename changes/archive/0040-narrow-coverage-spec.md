---
already_satisfied:
    - CHK-002
    - CHK-025
    - CHK-026
audit_rounds: 0
bases:
    CHK-002: a4cd59f6
module: checks
phase: applied
run: 20261028T052926-160566b09f6c468fcfed55c6
---



## Motivation

Continue the previously identified checks splits under SYSTEM's one-behavior
spec shape. CHK-002 currently bundles three public APIs. Give fail-first and
passing-result checks independent IDs and retain coverage under CHK-002,
preserving the design's Evals and Implementation step requirements.

## Interface



## Added

### CHK-025
RequireFailFirst MUST find a failed or errored referencing case for every requested enforced ID except already_satisfied; passing, skipped, empty or unrelated failures do not qualify.

### CHK-026
RequirePass MUST reject failed or errored cases referencing any non-quarantined spec.

## Modified

### CHK-002
RequireCoverage MUST cover each current spec except review or quarantine with a non-skipped referencing case.

## Removed



## Decisions



## Notes

Retarget existing junit_test.go tests only: coverage tests retain CHK-002,
the two FailFirst tests gain the fail-first ID instead, and PassRejectsNonQuarantinedFailures
uses the passing-results ID. Preserve all assertions and all 28 reported
checks cases, with no new tests, t.Run cases, fixture instances or helpers.
JUnit failure and error elements both become Result fail in ParseJUnit under
CHK-001; these APIs use that normalized report representation. Do not invent
another result-string value or new behavior. No product code or interface
changes. Implementation may edit only this change file.

Implementation review found the existing coverage, fail-first and passing-result
checks already satisfy their projected specs. `verify --quick` passed all 28
checks cases; no product code or interface edits were needed.

## Provenance

- evals 2026-10-28T05:30:04.731563Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=b26da634732c7f1c7dc600115582a63cb8f82e1fd2bd1d522d608fe837a04ba7 tests=146 failing=0 passing=146
- audit 2026-10-28T05:30:35.611208Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=0fb044cfaca1b9a844a25b4c86223177e52f2f35fb149eddf044e83822665c88 tests=0 failing=0 passing=0
- impl 2026-10-28T05:31:21.006806Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=07554a3249d8fca0ba70df7504135f297491465caa63458734a4b8351734d0e5 tests=146 failing=0 passing=146

## Audit

No findings. Checked CHK-002, CHK-025, CHK-026.
