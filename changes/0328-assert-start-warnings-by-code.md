---
eval_fix:
    - APP-027
module: app
phase: impl
run: 20261119T000630-61184a38a85232bcad090520
---

## Motivation

The November 18 workflow drill returned both P05 and P04. The CLI eval failed
because it expected P04 at index zero, although APP-027 only requires start to
preserve the reported warnings. Earlier drill fixes removed comparable dependence on
YAML presentation.
[Drill reports](https://github.com/lurkshark/fireweed/pull/25)

## Interface

## Added

## Modified

## Removed

## Decisions

## Notes

In Test_APP_009_APP_027_APP_028_APP_040_StartPreservesApprovalWarning, find the P04
warning by code in the JSON diagnostics and assert its existing warning flag and
nonempty message. Assert that the divergent-base fixture also delivers its P05
warning. Keep the success, context, plain output and empty JSON stderr assertions.
Accept either warning order.
## Provenance

- evals 2026-11-19T00:14:49.168893Z agent=codex fireweed=0.2.1 playbook=d2fd7216b035a9614c390d254dd9195ac2c65ae1ed18f07ac93c77b9732d4156 tests=415 failing=0 passing=415
- audit 2026-11-19T00:15:27.213368Z agent=manual fireweed=0.2.1 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0

## Audit

No findings; checked APP-027

