---
audit_pending:
    - AST-028
audit_rounds: 1
bases: {}
module: assets
phase: evals
run: 20261029T024648-f67119a09cc3fbd11b3cf937
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

### AST-028
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

## Provenance

- evals 2026-10-29T02:50:02.521195Z agent=manual fireweed=0.1.0 playbook=b26da634732c7f1c7dc600115582a63cb8f82e1fd2bd1d522d608fe837a04ba7 tests=461 failing=1 passing=460
- audit 2026-10-29T02:59:02.911427Z agent=manual fireweed=0.1.0 playbook=0fb044cfaca1b9a844a25b4c86223177e52f2f35fb149eddf044e83822665c88 tests=0 failing=0 passing=0

## Audit

- AST-028: "followed by a list of IDs on separate lines" — the checked-ID list clause is not fully asserted.
