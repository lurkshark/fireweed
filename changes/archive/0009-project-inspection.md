---
module: app
phase: applied
run: bf80e89092f8451c72bbb1e5
---

## Motivation

Publish the discovered SPEC.md graph and static diagnostics through the actual independent executable before building authoring and lifecycle operations.

## Interface

Keep Run and protected main.go unchanged. Add modules and check dispatch using project's public Load/Check API. Extend disposable JSON diagnostic representation with repository-relative path. Evals build and invoke the NEW binary using temporary repositories and public contracts.

## Added

### APP-005
modules MUST discover the project from a nested working directory and publish the resolved module array sorted by name then directory, including names, prefixes, directories, dependencies, resolved zone lists, hooks, specs and Decisions. It MUST apply current project defaults and explicit empty overrides and MUST NOT require a second language-specific graph. Empty projects MUST return an empty non-null array. Both plain output and JSON data MUST preserve these values.

### APP-006
check MUST run public project static validation and return data {valid:true} with exit 0 when there are no error diagnostics, including when warnings exist, and {valid:false} with exit 1 when errors exist. modules MUST return its module array with the same validation diagnostics and exit 0 or 1 by severity. JSON MUST use the schema-1 canonical command/data/non-null diagnostics envelope, with each diagnostic's code, severity, message and repository-relative path in deterministic project order. Plain diagnostics MUST go to stderr with their code and path; data MUST go to stdout. Invalid/missing/incompatible configuration or absence of a Git repository MUST return exit 2 with C01 error diagnostics and null data in JSON, without U01 or plaintext on JSON streams.

### APP-007
modules and check MUST accept --json before or after the command, accept --help/-h without loading a project, and reject excess operands or unknown/valued flags with U01 exit 2 before loading a project. They MUST remain noninteractive and MUST NOT run hooks, agents, the bootstrap or network commands, nor change repository files, ignored content or refs. Project validation MUST expose invalid graph/zone/hook/document/reservation diagnostics from the public project validator rather than silently succeeding.

## Modified

## Removed

## Decisions

The current graph comes exclusively from SPEC.md and fireweed.yaml through project.Load. Both commands call project.Check on the snapshot, so modules can publish a partial readable graph alongside error diagnostics while returning exit 1. Plain data is pretty JSON, matching the existing structured inspection convention. Validation errors retain data; configuration/discovery fatal errors return null data with C01 and path fireweed.yaml. Help is available before repository loading. These commands expose current static state; queued projection, run status, playbook overrides and workflow commands remain separate changes. Module arrays include the public model's retained document field, providing additional author context without dropping designed graph fields.

## Rejected alternatives

Calling the bootstrap would conceal missing product behavior. A separate graph in Go tooling would violate the design's single source of module ownership and dependencies.

## Provenance

- evals | 2026-10-22T08:27:25Z | agent: manual | fireweed: 0.4.0 | playbook: 5867363302d4c874074520977b1627d26991d95ee2fbacb9ff63e827661edd0b | tests: 12 | failing: 4 | passing: 8
- audit | 2026-10-22T08:28:54Z | agent: manual | fireweed: 0.4.0 | playbook: a9c982d24b476621316049b527e6c7bbc319f06b12268af150fa7d8c6df92230 | tests: 0 | failing: 0 | passing: 0
- evals | 2026-10-22T08:29:17Z | agent: manual | fireweed: 0.4.0 | playbook: 5867363302d4c874074520977b1627d26991d95ee2fbacb9ff63e827661edd0b | tests: 12 | failing: 4 | passing: 8
- audit | 2026-10-22T08:30:23Z | agent: manual | fireweed: 0.4.0 | playbook: a9c982d24b476621316049b527e6c7bbc319f06b12268af150fa7d8c6df92230 | tests: 0 | failing: 0 | passing: 0
- impl | 2026-10-22T08:31:31Z | agent: manual | fireweed: 0.4.0 | playbook: 653912b194316c62181f62d0a69dc213f7668481ea8b1b9926ef889dd48a9e08 | tests: 12 | failing: 0 | passing: 12

## Audit

No findings. Independent auditor /root/audit_cli_inspection checked APP-005, APP-006 and APP-007 against public contracts, assertions and valid fixture grammar. Plain invalid/warning/valid data gaps are resolved. No implementation/history/bootstrap inspection or test execution by auditor.
