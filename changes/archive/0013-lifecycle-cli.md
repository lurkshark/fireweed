---
audit_pending:
    - APP-004
    - APP-009
    - APP-010
    - APP-011
    - APP-012
audit_rounds: 1
bases:
    APP-004: 053f5e17
module: app
phase: applied
run: 20261027T141303-ad015bbc17ea93d3cee80f10
---



## Motivation

The manual slice now works through library interfaces. Expose its lifecycle,
context and coverage through the NEW executable so the existing Go fixture can
complete fail-first/audit/implementation/archive entirely through CLI commands.
The boundaries hook also needs the held approved module graph.

## Interface



## Added

### APP-009
start <change> MUST discover the root, invoke workflow.Start and print first-step workflow.Context with exit 0, assigning IDs/run/phase without committing. It MUST preserve reported warnings and return exit 1 for lifecycle rejection or 2 for usage/configuration errors, with corrective code/message diagnostics. --json MUST use schema 1 with context string data on success or null data on failure.

### APP-010
verify [change] [--quick] and done [change] MUST invoke workflow with the actual invoking executable as Bin, returning 0 for success, 1 for failed checks/blocked state and 2 for usage/configuration errors. Done MUST always use full checks; verify --quick uses quick checks without commits. --json MUST return schema 1 with result data on success or null on failure and reported code/message diagnostics; plain success names the resulting phase.

### APP-011
context <change> [--full] MUST return the held workflow.Context text, using compact or full behavior as requested, with exit 0 and empty diagnostics. --json MUST return identical text as schema-1 string data. Usage/configuration errors MUST return 2 and held-state lookup failures 1, with corrective code/message diagnostics and null JSON data.

### APP-012
coverage <module> MUST evaluate current specs, or the selected projection with --with-change <change>, using approved metadata while held and setting FIREWEED_BIN to the actual executable. It MUST map results to IDs, require coverage/pass and return exit 0 on success or 1 on failed checks. Usage/configuration/unknown module/change errors MUST return 2 with corrective diagnostics; --json uses schema 1 with report data on success, null on failure and code/message diagnostics.

## Modified

### APP-004
modules MUST discover the nearest ancestor fireweed.yaml and report the validated graph, using the held approved projected snapshot while this worktree owns a change. With --json it MUST return schema 1 with array data containing name, id_prefix, dir, depends_on and resolved zone arrays, plus empty diagnostics. Success and missing/invalid configuration MUST return exits 0 and 2 respectively, with a corrective C01 failure diagnostic.

## Removed



## Decisions



## Notes

Evals use a built NEW executable and temporary Git projects with shell JUnit hooks;
hook toggles live in ignored `.fireweed/runs`. A public workflow fixture probe
reached applied, including audit. Quick controller verification reports 78 passing
and 8 failing app cases; every changed ID fails against the current CLI.

## Provenance

- evals 2026-10-27T14:24:35.348641Z agent=manual fireweed=0.1.0 playbook=b26da634732c7f1c7dc600115582a63cb8f82e1fd2bd1d522d608fe837a04ba7 tests=86 failing=8 passing=78
- audit 2026-10-27T14:28:12.174514Z agent=manual fireweed=0.1.0 playbook=0fb044cfaca1b9a844a25b4c86223177e52f2f35fb149eddf044e83822665c88 tests=0 failing=0 passing=0
- evals 2026-10-27T14:32:11.342017Z agent=manual fireweed=0.1.0 playbook=b26da634732c7f1c7dc600115582a63cb8f82e1fd2bd1d522d608fe837a04ba7 tests=86 failing=8 passing=78
- audit 2026-10-27T14:34:34.525113Z agent=manual fireweed=0.1.0 playbook=0fb044cfaca1b9a844a25b4c86223177e52f2f35fb149eddf044e83822665c88 tests=0 failing=0 passing=0
- impl 2026-10-27T14:43:14.895479Z agent=manual fireweed=0.1.0 playbook=07554a3249d8fca0ba70df7504135f297491465caa63458734a4b8351734d0e5 tests=86 failing=0 passing=86

## Audit

No findings

Checked IDs: APP-004, APP-009, APP-010, APP-011, APP-012
