---
module: checks
eval_fix:
    - CHK-050
    - CHK-051
    - CHK-052
    - CHK-053
    - CHK-055
    - CHK-056
    - CHK-057
    - CHK-058
    - CHK-059
    - CHK-060
    - CHK-061
    - CHK-067
---

## Motivation

Checks' real Go recipe doctor eval and the declared-test case of its supplied
missing-ID canary eval build the whole Fireweed CLI before testing Doctor.
RUN-053 omits dependent modules from the checks reconstruction workspace.
The fixture's build of the repository root therefore fails at the missing app
import before either eval exercises Doctor. These evals need to exercise the
real recipe and its canaries using only checks and its declared dependencies.

## Interface

## Added

## Modified

## Removed

## Decisions

## Notes

Replace the shared `doctorBinary` fixture used by the two affected evals with
a small executable built from retained dependencies. It should handle the
Go recipe's `modules --json` request using real project discovery in the
scratch working directory, emitting the module names, directories,
dependencies and interface and eval zones the boundary hook consumes.
Exercise the actual recipe hooks with `FIREWEED_BIN` pointing to this
fixture. Preserve the existing canary inputs, assertions, and spec IDs,
including scratch discovery, boundary failures, missing-ID detection,
NoAgent, recovery, and cleanup. The fixture should build with
`packages/app` absent. Confine the correction to the checks eval zone.
