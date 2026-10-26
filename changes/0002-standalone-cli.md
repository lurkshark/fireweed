---
audit_rounds: 0
bases: {}
module: app
phase: impl
run: 20261026T202346-fcd8c0b82ee99647b8ddc47b
---



## Motivation

The design's CLI and Distribution sections and milestone 1 require a usable
independent executable before project/lifecycle operations are expanded.

## Interface

Keep Run(args []string, out, errOut io.Writer) int as the public CLI entry point.
Wire assets through disposable app/internal code; root main.go stays unchanged.

## Added

### APP-001
The CLI MUST run help, version, explain <code> and playbook <name> successfully
outside a project, returning the corresponding assets text on stdout with exit
0 and empty stderr. These commands MUST also work in a directory containing an
invalid fireweed.yaml.

### APP-002
Each information command MUST accept --json and emit one JSON object with schema
1, string data equal to its plain text result, and an empty diagnostics array.
It MUST return exit 0 and leave stderr empty.

### APP-003
Unknown commands/options, missing or extra arguments, and unknown lookup names
MUST return exit 2 with a nonempty corrective diagnostic. Commands listed by help
but still under construction MUST explicitly report that they are unimplemented
and return nonzero. Failed --json requests MUST use the schema-1 envelope with
null data and a diagnostic containing code and message.

## Modified



## Removed



## Decisions



## Rejected alternatives



## Provenance

- evals 2026-10-26T20:27:01.636439Z agent=manual fireweed=0.1.0 playbook=b26da634732c7f1c7dc600115582a63cb8f82e1fd2bd1d522d608fe837a04ba7 tests=30 failing=30 passing=0
- audit 2026-10-26T20:28:24.291091Z agent=manual fireweed=0.1.0 playbook=0fb044cfaca1b9a844a25b4c86223177e52f2f35fb149eddf044e83822665c88 tests=0 failing=0 passing=0

## Audit

No findings. Checked APP-001, APP-002, APP-003.
