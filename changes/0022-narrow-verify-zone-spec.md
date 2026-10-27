---
already_satisfied:
    - FLOW-002
    - FLOW-028
    - FLOW-029
    - FLOW-030
    - FLOW-031
audit_rounds: 0
bases:
    FLOW-002: a6ff549a
module: workflow
phase: impl
run: 20261027T202735-fe8faf8d8c6010599509c9b3
---



## Motivation

Finish the review-directed FLOW-002 split. Its coverage, fail-first, approved-state
and full-verification gates now have dedicated IDs. Preserve its remaining
independent behaviors and narrow FLOW-002 to its phase-zone check. All behavior
already works; reuse and retarget existing evals without implementation changes.

## Interface



## Added

### FLOW-028
Verify MUST select hooks from the held approved projected project snapshot.

### FLOW-029
Quick Verify MUST skip dependent-module evals.

### FLOW-030
Quick Verify MUST skip boundary verification.

### FLOW-031
Verify MUST leave HEAD unchanged.

## Modified

### FLOW-002
Verify MUST reject edits forbidden by the approved projected phase zones before running hooks, checking differences since the held phase base.

## Removed



## Decisions



## Notes

No new Test functions, t.Run cases or fixture instances. Keep 32 reported workflow
cases and every existing assertion. Retarget FLOW-002 labels on assertions that
now belong to FLOW-020 through FLOW-027 or the Added IDs; retain FLOW-002 on zone
checks. Existing evals and impl quick/full scenarios already exercise both skips
and HEAD stability. The frozen-eval scenario checks differences since an advanced
phase base.

For approved hook selection, reuse the existing impl-apply fixture with its set
metadata: author a set.hooks.eval override pointing to the existing eval script,
and point the global eval hook to a simple failing script before Start. Keep
its other set/application assertions. No new fixture or separate test case.
No interface or implementation edits; impl may edit only this change file.

## Provenance

- evals 2026-10-27T20:29:39.570882Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=b26da634732c7f1c7dc600115582a63cb8f82e1fd2bd1d522d608fe837a04ba7 tests=118 failing=0 passing=118
- audit 2026-10-27T20:30:21.435482Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=0fb044cfaca1b9a844a25b4c86223177e52f2f35fb149eddf044e83822665c88 tests=0 failing=0 passing=0

## Audit

Checked IDs: FLOW-002, FLOW-028, FLOW-029, FLOW-030, FLOW-031
No findings
