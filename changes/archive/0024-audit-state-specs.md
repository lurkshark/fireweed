---
already_satisfied:
    - FLOW-036
    - FLOW-037
    - FLOW-038
    - FLOW-039
    - FLOW-040
audit_rounds: 0
bases: {}
module: workflow
phase: applied
run: 20261027T203601-f3e409a5f0eac8feccb7c61b
---



## Motivation

Continue the review-directed FLOW-003 split with independent audit state rules.
Existing evals already assert each behavior. Keep FLOW-003 until its remaining
passing-audit, Questions and claim-release behavior can be retained separately.
No new feature or implementation change belongs here.

## Interface



## Added

### FLOW-036
Audit Done MUST require every enforced ID in the change's Added, Modified or fixes entries to appear in a No findings submission's checked-ID list.

### FLOW-037
Retained testable audit findings within the round limit MUST return the change to evals.

### FLOW-038
An audit return to evals MUST increment and persist audit_rounds.

### FLOW-039
Audit Done MUST block the change when retained findings would exceed audit.max_rounds, whose default is 2.

### FLOW-040
Audit Done MUST block the change when a retained finding marks its quoted clause untestable.

## Modified



## Removed



## Decisions



## Notes

No new Test functions, t.Run cases or fixture instances. Keep 32 reported workflow
cases and every existing assertion. Retarget the existing checked-No-findings
case, the two-round/default and one-round/configured bound case, and the existing
untestable-finding case. Preserve FLOW-003 labels until final narrowing. These
cases already check phase, persistent count, block commits, Questions and claim
release; do not add scenarios or alter fixtures for asserted behavior.
No interface or implementation edits; impl may edit only this change file.

## Provenance

- evals 2026-10-27T20:37:35.667931Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=b26da634732c7f1c7dc600115582a63cb8f82e1fd2bd1d522d608fe837a04ba7 tests=118 failing=0 passing=118
- audit 2026-10-27T20:38:06.612934Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=0fb044cfaca1b9a844a25b4c86223177e52f2f35fb149eddf044e83822665c88 tests=0 failing=0 passing=0
- impl 2026-10-27T20:39:16.174341Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=07554a3249d8fca0ba70df7504135f297491465caa63458734a4b8351734d0e5 tests=118 failing=0 passing=118

## Audit

Checked IDs: FLOW-036, FLOW-037, FLOW-038, FLOW-039, FLOW-040
No findings
