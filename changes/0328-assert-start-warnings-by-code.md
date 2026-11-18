---
eval_fix:
    - "APP-027"
module: app
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
