---
already_satisfied:
    - new:eval-hook-selection
    - new:eval-hook-placeholders
    - new:eval-hook-bin
    - new:eval-hook-output
    - new:eval-hook-report
bases: {}
module: checks
---



## Motivation

Begin the identified CHK-003 split under SYSTEM's one-behavior spec shape.
Retain the design's Hooks command and report contract, separating hook inputs
and returned data from completion/result validation. Existing public hook evals
already assert these rules. Keep CHK-003 until its remaining rules are split.

## Interface



## Added

### new:eval-hook-selection
Evaluate MUST prefer a module's eval hook override over the global eval hook.

### new:eval-hook-placeholders
Evaluate MUST expand module, module_dir and report placeholders once in shell strings or argv, quoting substituted shell values.

### new:eval-hook-bin
Evaluate MUST set the hook's FIREWEED_BIN environment variable to the request Bin.

### new:eval-hook-output
Evaluate MUST retain the hook process's stdout, stderr and exit code in its returned report.

### new:eval-hook-report
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
