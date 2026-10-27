---
bases:
    APP-004: 053f5e17
module: app
---



## Motivation

The manual slice now works through library interfaces. Expose its lifecycle,
context and coverage through the NEW executable so the existing Go fixture can
complete fail-first/audit/implementation/archive entirely through CLI commands.
The boundaries hook also needs the held approved module graph.

## Interface



## Added

### new:start-command
start <change> MUST discover the root, invoke workflow.Start and print first-step workflow.Context with exit 0, assigning IDs/run/phase without committing. It MUST preserve reported warnings and return exit 1 for lifecycle rejection or 2 for usage/configuration errors, with corrective code/message diagnostics. --json MUST use schema 1 with context string data on success or null data on failure.

### new:verify-done
verify [change] [--quick] and done [change] MUST invoke workflow with the actual invoking executable as Bin, returning 0 for success, 1 for failed checks/blocked state and 2 for usage/configuration errors. Done MUST always use full checks; verify --quick uses quick checks without commits. --json MUST return schema 1 with result data on success or null on failure and reported code/message diagnostics; plain success names the resulting phase.

### new:context-command
context <change> [--full] MUST return the held workflow.Context text, using compact or full behavior as requested, with exit 0 and empty diagnostics. --json MUST return identical text as schema-1 string data. Usage/configuration errors MUST return 2 and held-state lookup failures 1, with corrective code/message diagnostics and null JSON data.

### new:coverage-command
coverage <module> MUST evaluate current specs, or the selected projection with --with-change <change>, using approved metadata while held and setting FIREWEED_BIN to the actual executable. It MUST map results to IDs, require coverage/pass and return exit 0 on success or 1 on failed checks. Usage/configuration/unknown module/change errors MUST return 2 with corrective diagnostics; --json uses schema 1 with report data on success, null on failure and code/message diagnostics.

## Modified

### APP-004
modules MUST discover the nearest ancestor fireweed.yaml and report the validated graph, using the held approved projected snapshot while this worktree owns a change. With --json it MUST return schema 1 with array data containing name, id_prefix, dir, depends_on and resolved zone arrays, plus empty diagnostics. Success and missing/invalid configuration MUST return exits 0 and 2 respectively, with a corrective C01 failure diagnostic.

## Removed



## Decisions
