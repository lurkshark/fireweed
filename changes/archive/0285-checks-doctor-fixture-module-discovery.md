---
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
module: checks
phase: applied
run: 20261116T160641-fa53751da6f82777fa70ba07
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

Both affected Doctor evals passed in a temporary copy containing only checks,
its declared dependencies, and the Go module files. Full `fireweed verify`
passed in the implementation phase.

## Audit

No findings; checked CHK-050, CHK-051, CHK-052, CHK-053, CHK-055, CHK-056, CHK-057, CHK-058, CHK-059, CHK-060, CHK-061, CHK-067

## Provenance

- evals 2026-11-16T16:16:41.860076Z agent=codex fireweed=0.2.1 playbook=6a9c2d506363e3539cc66efd37bc6a06701249541dc4e4f01ed3aa6454bbb864 tests=135 failing=0 passing=135
- audit 2026-11-16T16:17:08.732644Z agent=codex fireweed=0.2.1 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0
- impl 2026-11-16T16:44:04.108408Z agent=codex fireweed=0.2.1 playbook=5f856d6bcab308b9ff1a8a7c03b9d2547744112aba81dd30dea8f61ec8062b03 tests=966 failing=0 passing=966

