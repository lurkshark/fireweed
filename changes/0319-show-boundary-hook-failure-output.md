---
module: checks
phase: audit
run: 20261118T203128-28b85973d0f182d1531dbccf
---

## Motivation

An unattributed nonzero Go boundaries exit caused six dependent failures in the
November 18 assets drill. The X02 diagnostic showed the exit status but hid the
hook's explanation. A local hook that wrote an explanation to stderr reproduced this
on main. Users need the hook's reason to repair recipes and custom hooks.
[Drill reports](https://github.com/lurkshark/fireweed/pull/25)

## Interface

Add the hook's diagnostic output to the existing X02 diagnostic message for an
unattributed nonzero exit. Use the existing diagnostic schema and exit behavior.

## Added

### CHK-086
When a boundary hook exits nonzero without a module-attributed violation, its X02
diagnostic MUST identify the global hook or module override and include the end of
its stderr, if present, up to 4,096 bytes of valid UTF-8. The diagnostic MUST mark
any truncation of the excerpt.

## Modified

## Removed

## Decisions

## Notes

Extend the existing checks.Boundaries unattributed-exit fixtures. Use short stderr
for a global failure and long UTF-8 stderr with distinct head and tail markers for
a module override. Check the identified scope, the excerpt byte ceiling, the tail
marker, valid UTF-8 and a visible omission indicator. Retain a silent failure case,
and preserve existing exit status, attribution, baseline, opt-out and launch-error
assertions.

`fireweed verify show-boundary-hook-failure-output` passed the evals gate. The
global case failed because X02 omitted the short stderr reason. The override
case failed because X02 omitted the UTF-8 tail and truncation marker. The other
140 reported cases passed.

The unattributed-exit diagnostic is built in
`packages/checks/internal/engine/boundaries.go`, which is writable in impl.
## Provenance

- evals 2026-11-18T20:41:46.153674Z agent=codex fireweed=0.2.1 playbook=d2fd7216b035a9614c390d254dd9195ac2c65ae1ed18f07ac93c77b9732d4156 tests=142 failing=2 passing=140

