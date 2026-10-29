---
bases: {}
module: assets
---



## Motivation

The design says embedded playbooks describe and match the checked protocol.
The actual NEW next smoke returned P09 when its fake auditor wrote inline
No findings/Checked IDs text; the accepted separate-line form then passed.
The current audit playbook merely says to list IDs, without showing that form.
Add the passing example needed by the now-runnable next/run path before real
agents use it. This is the observed guidance correction, not a general parser
rewrite, new diagnostic allocation or deferred drill-gap fix.

## Interface

No public interface change. Extend only the embedded audit playbook text through
the existing Playbook API; retain all previous audit-scope and re-audit guidance.

## Added

### new:passing-audit-example
The audit playbook MUST show a passing ## Audit example with No findings and
Checked IDs: followed by a list of IDs on separate lines, instructing the auditor
to use the actual checked IDs.

## Modified



## Removed



## Decisions



## Rejected alternatives



## Notes

Add a small representative public-interface assertion for the accepted example.
Preserve every earlier test, instruction, lookup, version, hash behavior and
public signature. No parser changes, audit broadening, unrelated playbook/
explanation edits or new instrumentation. The existing embedded hash must remain
derived from the actual text. Own Verify stays below five minutes. Fresh
fail-first and separate isolated assertion audit remain required.

Workers never commit or call controller start/done/next/run, edit phase/counters,
protected docs or other modules. Coordinator owns native commits/application.
After this fifth application, stop for the standing subagent checkpoint review;
PASS authorizes push/continuation, an issue requires human review. Do not start
a sixth change or evade P07. No new drill or phase-write exception is authorized.
