---
bases: {}
module: app
---



## Motivation

The design's CLI and Distribution sections and milestone 1 require a usable
independent executable before project/lifecycle operations are expanded.

## Interface

Keep Run(args []string, out, errOut io.Writer) int as the public CLI entry point.
Wire assets through disposable app/internal code; root main.go stays unchanged.

## Added

### new:standalone-information
The CLI MUST run help, version, explain <code> and playbook <name> successfully
outside a project, returning the corresponding assets text on stdout with exit
0 and empty stderr. These commands MUST also work in a directory containing an
invalid fireweed.yaml.

### new:information-json
Each information command MUST accept --json and emit one JSON object with schema
1, string data equal to its plain text result, and an empty diagnostics array.
It MUST return exit 0 and leave stderr empty.

### new:explicit-usage-failures
Unknown commands/options, missing or extra arguments, and unknown lookup names
MUST return exit 2 with a nonempty corrective diagnostic. Commands listed by help
but still under construction MUST explicitly report that they are unimplemented
and return nonzero. Failed --json requests MUST use the schema-1 envelope with
null data and a diagnostic containing code and message.

## Modified



## Removed



## Decisions



## Rejected alternatives
