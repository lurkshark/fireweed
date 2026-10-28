---
already_satisfied:
    - CHK-003
    - CHK-036
    - CHK-037
    - CHK-038
audit_rounds: 0
bases:
    CHK-003: c337293b
module: checks
phase: impl
run: 20261028T053805-c45a8c17f97c27326d9ef467
---



## Motivation

Finish the identified CHK-003 split under SYSTEM's spec-shape rule. CHK-031–035
retain hook selection, substitution, environment and returned data. Give hook
completion and report validity independent IDs, leaving command execution in
CHK-003. Preserve the design's Hooks and fail-first evidence requirements;
all behavior and assertions already exist.

## Interface



## Added

### CHK-036
Evaluate MUST fail when the eval hook cannot launch or exceeds its timeout.

### CHK-037
Evaluate MUST fail when the eval hook does not write its JUnit report.

### CHK-038
Evaluate MUST reject a zero exit with mapped failures or a nonzero exit without mapped failures, while accepting a nonzero exit with mapped failures as valid report evidence.

## Modified

### CHK-003
Evaluate MUST run the configured eval hook, accepting shell strings or argv.

## Removed



## Decisions



## Notes

Retarget existing hooks_test.go tests only, preserving every assertion and all
28 checks cases. Keep CHK-003 on successful shell/argv execution tests.
ExitAndReportAgreement gains the result-agreement ID instead of CHK-003 and
retains CHK-034. MissingReportAndLaunchFailure uses the completion/missing-report
IDs instead of CHK-003; TimeoutFails uses completion instead of CHK-003.
Preserve CHK-031–035 everywhere else. No new Test functions, t.Run cases,
fixture instances, helpers, product code or interface changes. Leave
hookSnapshot and its embedded fixture spec unchanged. Implementation may
edit only this change file.

## Provenance

- evals 2026-10-28T05:38:45.790568Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=b26da634732c7f1c7dc600115582a63cb8f82e1fd2bd1d522d608fe837a04ba7 tests=146 failing=0 passing=146
- audit 2026-10-28T05:39:23.211209Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=0fb044cfaca1b9a844a25b4c86223177e52f2f35fb149eddf044e83822665c88 tests=0 failing=0 passing=0

## Audit

No findings.

Checked: CHK-003, CHK-031, CHK-032, CHK-033, CHK-034, CHK-035, CHK-036, CHK-037, CHK-038.
