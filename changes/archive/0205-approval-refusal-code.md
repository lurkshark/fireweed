---
bases:
    RUN-041: dcfedf6a
module: runner
phase: applied
run: 20261111T022657-c5c29652a3f4de7f915ac1c1
---

## Motivation

The design requires a stable code for every problem. APP-075, APP-077 and
APP-082 require P01 when `next` or `run` refuses an unapproved change, but
RUN-041 only requires `IsApprovalRequired`. The runner drill's first
reconstruction put a different code in the refusal text, the app reported
it, and the drill passed only after the retry showed it the app failures.
Today the app gets P01 only by default, because the runner error carries no
code. Give the refusal its code at the runner boundary.

## Interface

The approval refusal returned by `Next` and `Run` becomes a coded error
readable through `contracts.ErrorCode`. No signature changes.

## Added

## Modified

### RUN-041
Next MUST return an error for which IsApprovalRequired reports true and contracts.ErrorCode returns P01 when the selected change file is on neither the local nor the shared base and AllowUnmerged is false.

## Removed

## Decisions

## Notes

Wrap the existing approval error in a `contracts.CodedError` with code P01
and the change path as subject, keeping `IsApprovalRequired` true through
the wrap. The app evals for APP-075, APP-077 and APP-082 must keep passing
unchanged.
## Provenance
- evals 2026-11-11T02:43:17.673524Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6.1-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=6a9c2d506363e3539cc66efd37bc6a06701249541dc4e4f01ed3aa6454bbb864 tests=166 failing=2 passing=164
- audit 2026-11-11T02:43:41.112518Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-astra -c model_reasoning_effort="medium" -s workspace-write {prompt} fireweed=0.1.0 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0
- impl 2026-11-11T03:05:04.017414Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6.1-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=5f856d6bcab308b9ff1a8a7c03b9d2547744112aba81dd30dea8f61ec8062b03 tests=489 failing=0 passing=489
## Audit

No findings; checked RUN-041
