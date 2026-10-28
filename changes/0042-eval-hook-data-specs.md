---
already_satisfied:
    - CHK-031
    - CHK-032
    - CHK-033
    - CHK-034
    - CHK-035
audit_rounds: 0
bases: {}
module: checks
phase: audit
run: 20261028T053501-8ccd279ce0fce2358ab2b579
---



## Motivation

Begin the identified CHK-003 split under SYSTEM's one-behavior spec shape.
Retain the design's Hooks command and report contract, separating hook inputs
and returned data from completion/result validation. Existing public hook evals
already assert these rules. Keep CHK-003 until its remaining rules are split.

## Interface



## Added

### CHK-031
Evaluate MUST prefer a module's eval hook override over the global eval hook.

### CHK-032
Evaluate MUST expand module, module_dir and report placeholders once in shell strings or argv, quoting substituted shell values.

### CHK-033
Evaluate MUST set the hook's FIREWEED_BIN environment variable to the request Bin.

### CHK-034
Evaluate MUST retain the hook process's stdout, stderr and exit code in its returned report.

### CHK-035
Evaluate MUST parse the hook's written JUnit report into the returned report's cases.

## Modified



## Removed



## Decisions



## Notes

Retarget existing hooks_test.go tests and keep CHK-003 until its next narrowing.
GlobalShellHookQuotesAndOnePassExpansion and ArgvPlaceholdersAreOnePass cover
placeholder rules; ModuleOverrideArgvAndOutput covers override selection.
GlobalShellHookQuotesAndOnePassExpansion and ModuleOverrideArgvAndOutput assert
FIREWEED_BIN and returned data; ExitAndReportAgreement also checks nonzero
exit retention. Add the corresponding IDs without removing any assertion or
other label. Keep 28 reported checks cases. No new Test functions, t.Run cases,
fixture instances, helper changes, product code or interfaces. Leave hookSnapshot
and its embedded fixture spec unchanged. Implementation may edit only this change file.

## Provenance

- evals 2026-10-28T05:35:49.258642Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=b26da634732c7f1c7dc600115582a63cb8f82e1fd2bd1d522d608fe837a04ba7 tests=146 failing=0 passing=146
