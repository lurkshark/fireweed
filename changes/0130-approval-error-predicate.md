---
audit_rounds: 0
module: runner
phase: impl
run: 20261031T042130-ae9b6df7b92e26d7b3aed83c
---



## Motivation

App next and run choose their recovery advice by matching `requires approval` in
runner errors. The design gives unapproved changes their own handling (run skips
them without --allow-unmerged); callers need to recognise that refusal without
depending on message wording, which a regenerated runner may change.

## Interface

Add `IsApprovalRequired(err error) bool` to the runner public API, beside the
existing `IsUsageStop`.

Keep the public facade delegating to disposable implementation. During evals,
add the implementation stub in a new internal file so the impl step can replace
it without editing the frozen interface (the wiring repair from 0128).

## Added

### RUN-041
Next MUST return an error for which IsApprovalRequired reports true when the selected change file is on neither the local nor the shared base and AllowUnmerged is false.

### RUN-042
IsApprovalRequired MUST report false for Next errors caused by anything other than a missing approval.

## Modified



## Removed



## Decisions



## Rejected alternatives



## Provenance

- evals 2026-10-31T04:36:40.355209Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.1 playbook=b26da634732c7f1c7dc600115582a63cb8f82e1fd2bd1d522d608fe837a04ba7 tests=362 failing=4 passing=358
- audit 2026-10-31T04:37:10.048125Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.1 playbook=5cb9f9dac24b0b6bb1b970edd4fab4406d2b4d0435ddefc3ef59a2cfdd267f18 tests=0 failing=0 passing=0

## Audit

No findings
Checked IDs: RUN-041, RUN-042
