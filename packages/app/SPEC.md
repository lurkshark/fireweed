---
depends_on:
    - contracts
    - gitstore
    - assets
    - project
    - checks
    - workflow
    - runner
id_prefix: APP
name: app
---

# app

## Specs

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

## Decisions
