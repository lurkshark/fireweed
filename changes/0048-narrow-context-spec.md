---
already_satisfied:
    - FLOW-005
    - new:done-provenance
    - new:retained-run-evidence
    - new:held-state-snapshot
bases:
    FLOW-005: 4832466c
module: workflow
---



## Motivation

Complete the previously identified FLOW-005 bundle under SYSTEM's spec-only
batching and atomic-spec rules. Retain the design's Context and Provenance
requirements while separating context output, provenance, run evidence and
State snapshot exposure. All behavior already exists; preserve the public
interfaces, implementation and all existing assertions.

## Interface



## Added

### new:done-provenance
Each successful Done MUST append provenance containing phase, date, manual or configured agent, version, playbook hash and result counts.

### new:retained-run-evidence
Successful Done MUST retain verification reports and process output under the run directory.

### new:held-state-snapshot
State MUST expose the held projected snapshot so composing commands use the approved graph during hooks.

## Modified

### FLOW-005
Context MUST include the change, projected spec, interface contents, SYSTEM and matching playbook. It MUST list other relevant files by path and size, or include their contents when full is requested.

## Removed



## Decisions



## Notes

Retarget ContextAndProvenanceAndRetainedReports to FLOW-005 and the three Added
IDs, preserving every compact/full context, manual/configured agent, playbook,
count and retained-report assertion. StartAssignsLateIDsAndRetainsSnapshotWithoutCommit
also gains the snapshot ID for its existing approved snapshot assertions,
retaining all current IDs. Existing FLOW-006/019/028 assertions still apply.
Keep all 32 workflow cases, with no new Test functions, t.Run cases, fixture
instances or helper changes. No product behavior or interface edits. Keep the
existing behavioral provenance-format Decision; it still governs formatting.
Implementation may edit only this change file.
