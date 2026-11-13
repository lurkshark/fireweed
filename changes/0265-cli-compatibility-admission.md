---
module: app
---

## Motivation

R20 requires project operations to obey the binary/format compatibility
policy, including commands using a held approved snapshot. APP-004/012/160
deliberately bypass broken live inputs, so merely changing project.Load cannot
cover every command. Check the configuration source each command actually
uses, preserving APP-127's historical endpoint authority and standalone
information commands rather than introducing blanket live-config admission.

## Interface

No public signature or type changes.

## Added

### new:project-command-compatibility
For commands using live project configuration, the CLI MUST require project.CheckCompatibility to accept that configuration before command effects or hook/agent execution. A compatibility rejection MUST return exit 2 with C01 and corrective guidance.

### new:approved-command-compatibility
While a change is owned, modules, coverage, next and run MUST check compatibility against the held approved configuration used by their operations. Compatibility admission MUST NOT add live-config parsing to their approved-snapshot paths or replace live configuration required by other commands with a held snapshot. An approved-configuration rejection MUST return exit 2 with C01 and corrective guidance before command effects or hook/agent execution.

### new:compatibility-command-exceptions
Help forms, version, explain and playbook MUST retain their specified standalone behavior. init and migrate MUST retain their setup and format-validation preconditions. pr-body MUST preserve APP-127's endpoint authority without compatibility admission against live configuration or the recorded endpoint, including when live configuration is malformed.

## Modified



## Removed

## Decisions

## Notes

Run after 0263 and 0264 introduce and complete CheckCompatibility. Extend
existing binary fixtures with representative reading, authoring and lifecycle
commands; test held approved configurations using APP-004/012/160's existing
malformed/missing-live-file scenarios and Next/Run's approved-snapshot path.
Retain the required live spec/projection reads for coverage of other targets.
Owning a claim does not switch check, show, authoring or other live consumers
to approved data. Keep existing root discovery and command prerequisites.

show --history reads the current archive, and check-commits retains its live
configuration prerequisite; historical documents they inspect do not supply
the admission pin. LoadAt stays a raw historical reader. Keep the existing
APP-127/128/129 malformed-live pr-body success fixtures and APP-130/131
failure classifications unchanged; do not use eval_fix to weaken them.

Assert rejection before effects through file/ref snapshots and ordinary fake
commands. Preserve plain/JSON diagnostics and empty JSON stderr. Use a small
representative set, not every command crossed with every pin/format error.
Preserve APP-153/155 migrate validation; format 1 needs no migration.

Drafted at the owner's request after the R01–R29 traceability review.
No lifecycle execution, agent invocation, drill or release is authorized by
this drafting request. Workers must obtain context before their step and use
fireweed verify; only the controller may advance or commit lifecycle work.
