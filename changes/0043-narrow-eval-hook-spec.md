---
already_satisfied:
    - CHK-003
    - new:eval-hook-completion
    - new:eval-hook-missing-report
    - new:eval-hook-result-agreement
bases:
    CHK-003: c337293b
module: checks
---



## Motivation

Finish the identified CHK-003 split under SYSTEM's spec-shape rule. CHK-031–035
retain hook selection, substitution, environment and returned data. Give hook
completion and report validity independent IDs, leaving command execution in
CHK-003. Preserve the design's Hooks and fail-first evidence requirements;
all behavior and assertions already exist.

## Interface



## Added

### new:eval-hook-completion
Evaluate MUST fail when the eval hook cannot launch or exceeds its timeout.

### new:eval-hook-missing-report
Evaluate MUST fail when the eval hook does not write its JUnit report.

### new:eval-hook-result-agreement
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
