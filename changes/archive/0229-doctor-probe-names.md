---
module: checks
phase: applied
run: 20261111T222805-e028ae6f99a11bae90055070
---

## Motivation

`fireweed doctor --json` lists each probe by a short name such as
`baseline` or `recovery`, and the app's eval for APP-067 requires those
names. Checks' DoctorCheck names include the hook scope after the probe
name, and no checks spec says what a name contains, so the app takes the
first word of a string whose format is unstated. The app drill lists this
as a gap. State the format in the module that produces it.

## Interface

None. DoctorCheck.Name's format is specified.

## Added

### CHK-084
Each canary DoctorCheck Name MUST begin with its probe name (baseline, undeclared, internal_eval, failing_eval, missing_id or recovery), followed by a space and the hook scope when one applies.

## Modified

## Removed

## Decisions

## Notes

The current names already have this shape for the canary probes, so this
probably passes first; if it does not, the evals step reports the
mismatch. Other DoctorChecks, such as agent probes, are not affected.
Run before 0231.

The CHK-084 eval passes for successful canary checks across global hooks and
module overrides. It fails for an eval override's failed recovery, which emits
`recovery failing_eval alpha eval override` with the canary name before the
hook scope. The name is produced in
`packages/checks/internal/engine/doctor.go`, which is writable during impl.
## Provenance
- evals 2026-11-11T22:43:34.498859Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6.1-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=6a9c2d506363e3539cc66efd37bc6a06701249541dc4e4f01ed3aa6454bbb864 tests=130 failing=1 passing=129
- audit 2026-11-11T22:43:49.795748Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-astra -c model_reasoning_effort="medium" -s workspace-write {prompt} fireweed=0.1.0 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0
- impl 2026-11-11T23:29:02.950457Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6.1-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=5f856d6bcab308b9ff1a8a7c03b9d2547744112aba81dd30dea8f61ec8062b03 tests=808 failing=0 passing=808
## Audit

No findings; checked CHK-084
