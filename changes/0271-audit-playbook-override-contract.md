---
module: runner
already_satisfied: [new:audit-playbook-override]
---

## Motivation

The design permits project playbook overrides (R20), including the audit
playbook supplied to a restricted auditor. RUN-013 names that playbook but
does not specify override selection. Current Next already reads the invoking
project's .fireweed/playbooks/audit.md; give that behavior an owning contract
without widening the restricted audit workspace.

## Interface

No public signature or type changes.

## Added

### new:audit-playbook-override
Next's audit prompt MUST include the complete invoking-project .fireweed/playbooks/audit.md text when present and the embedded audit playbook when absent.

## Modified

## Removed

## Decisions

## Notes

This addition is legitimately already_satisfied. Extend the existing bounded
audit-workspace fixture with distinct override text and absence fallback,
observing public Next results and the fake auditor's input.
Retain all restricted-tree/history checks, audit edit-import rules, separate
command selection, time/usage-stop behavior and cleanup assertions.
Do not copy the override file, config, interface or implementation into the
audit workspace. No implementation change or real-agent run is requested.
RUN-045 already owns workspace isolation; retain its assertions without
duplicating that obligation here. No unreadable-override hardening is added.

Draft only. Workers obtain context before their step and use fireweed verify;
the controller alone advances and commits lifecycle work. This task does not
authorize starting the change.
