---
bases: {}
eval_fix:
    - PRJ-007
module: project
---



## Motivation

D006 and the owner's existing-error-path migration constraint require repairing
wrong diagnostic-label assertions before implementation migration. PRJ-007 now
requires nonempty failure code/message fields, but its existing validation table
still demands C01 for unrelated validation causes. Correct only that stale
label expectation while preserving every behavior assertion and ID link.
Existing implementation must remain unchanged and all own evals must pass.

## Interface



## Added



## Modified



## Removed



## Decisions



## Notes

Change only evals/queue_test.go. In CheckValidAndDiagnostics, replace blanket
C01 equality checks with nonempty-code checks and update misleading failure
messages. Preserve the existing per-diagnostic nonempty code/message assertions,
all valid/invalid inputs, specific reason-message assertions, Added/Modified/
Removed cases and every ID label. Keep all 53 project cases; no new tests,
t.Run cases, fixture instances, helpers or test data. All module evals pass
rather than fail first. Do not assign D006 codes here; normal changes will add
exact cause-specific assertions and migrate existing implementation paths.
Do not edit product code, interfaces, Specs or other modules. During evals,
only eval-zone edits are authorized; the controller owns phase state/commits.
Implementation may edit only this change file, then archive.
