---
module: checks
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

### new:boundary-hook-failure-output
When a boundary hook exits nonzero without a module-attributed violation, its X02
diagnostic MUST identify the hook's scope and include a bounded tail of its stderr
when present. If the tail omits earlier output, the diagnostic MUST indicate that it
was truncated.

## Modified

## Removed

## Decisions

## Notes

Extend the existing checks.Boundaries unattributed-exit fixtures with stderr
explaining the failure. Exercise a global hook and a module override, and exercise
long UTF-8 stderr in one case to check a visible tail, truncation indication and
valid text. Preserve the exit status and X02 classification, including when stderr
is empty. Keep existing module-attributed violation, baseline, opt-out and
launch-error behavior.
