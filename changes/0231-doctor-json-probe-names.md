---
already_satisfied:
    - APP-067
bases:
    APP-067: 27d57f83
module: app
phase: impl
run: 20261112T022510-1f7f440d50e135c8222cd0a2
---

## Motivation

The app's APP-067 eval requires short probe names in `doctor --json`, and
the app takes them from the first word of checks' DoctorCheck names. APP-067
does not mention names, so an app rebuild has to infer that rule from the
eval. The app drill lists this as a gap. 0229 states the checks name format;
state the JSON rule here.

## Interface

None.

## Added

## Modified

### APP-067
doctor --json MUST emit schema 1 with an object containing a checks array on success or null data on failure, retaining code/message/path/warning diagnostics and leaving stderr empty. Each canary check MUST use its probe name as name and keep the hook scope in detail.

## Removed

## Decisions

## Notes

The current app already does this, so APP-067 is listed in
already_satisfied. Run after 0229.

The focused APP-067 evals passed against the existing app. The success
eval checks the short canary names and their eval and boundary hook scopes
in JSON detail, including baseline and recovery for both hooks.
## Provenance
- evals 2026-11-12T02:36:45.62141Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6.1-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=6a9c2d506363e3539cc66efd37bc6a06701249541dc4e4f01ed3aa6454bbb864 tests=323 failing=0 passing=323
- audit 2026-11-12T02:37:19.23138Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-astra -c model_reasoning_effort="medium" -s workspace-write {prompt} fireweed=0.1.0 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0
## Audit

No findings; checked APP-067
