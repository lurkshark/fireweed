---
already_satisfied:
    - RUN-071
module: runner
phase: impl
run: 20261114T082916-3886eb69581ae4e4ae110ccc
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

### RUN-071
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
## Audit

No findings; checked RUN-071

## Provenance

- evals 2026-11-14T08:39:30.218957Z agent=codex fireweed=0.1.0 playbook=6a9c2d506363e3539cc66efd37bc6a06701249541dc4e4f01ed3aa6454bbb864 tests=180 failing=0 passing=180
- audit 2026-11-14T08:40:27.656922Z agent=codex fireweed=0.1.0 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0

