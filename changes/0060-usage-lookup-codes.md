---
audit_rounds: 0
bases:
    APP-003: c93927b8
    APP-013: dd1c50dd
    APP-014: 4a097bdf
module: app
phase: audit
run: 20261028T081839-6abe11b1c9b65f91cb154373
---



## Motivation

The design requires stable diagnostic codes. D006 replaces legacy CLI_* labels
with U01 for invalid arguments/options and U02 for unknown lookup names. Existing
CLI usage, lookup and explicit unsupported-command rejections still emit legacy
labels. Correct those paths without implementing any unsupported command.

## Interface

No public interface change.

## Added



## Modified

### APP-003
Unknown commands or options and missing or extra arguments MUST return exit 2 with a nonempty corrective U01 diagnostic.

### APP-013
Unknown explain codes or playbook names MUST return exit 2 with a nonempty corrective U02 diagnostic.

### APP-014
Commands listed by help but still under construction MUST explicitly report that they are unimplemented with diagnostic U01 and return nonzero.

## Removed



## Decisions



## Notes

Reuse InvalidRequests, UnimplementedCommands and existing usage-error fixtures.
Add exact U01/U02 assertions for these existing plain and JSON cases while
retaining status, cause/argument, corrective-action, envelope and unimplemented
notice checks. No fixture/helper infrastructure or new scenarios are needed.
Reporting labels may change for correct ID attribution; retain all coverage.
Changed IDs must genuinely fail before implementation. Migrate existing
CLI_USAGE, CLI_OPTION and CLI_COMMAND paths to U01, including lifecycle arity
and misplaced-flag paths. Migrate CLI_LOOKUP to U02 and CLI_UNIMPLEMENTED to U01.
Leave CLI_AUTHOR/STAMP/SHOW/COVERAGE/CONFIG/LIFECYCLE and configuration causes for
subsequent scoped changes. Preserve every existing rejection and success path,
message, exit and JSON shape. Do not implement unsupported commands or add
speculative errors. Controller owns specs, commits and archive.

Eval evidence: APP-003, APP-013 and APP-014 fail on the existing CLI_* codes;
unchanged app IDs pass. Quick verification accepted 15 APP-003 cases. Existing
command fixtures retain their status and message checks.

## Provenance

- evals 2026-10-28T08:27:18.37152Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=b26da634732c7f1c7dc600115582a63cb8f82e1fd2bd1d522d608fe837a04ba7 tests=94 failing=18 passing=76
