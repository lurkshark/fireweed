---
audit_pending:
    - APP-062
    - APP-068
    - APP-070
audit_rounds: 1
bases: {}
module: app
phase: audit
run: 20261028T205002-0d9939dfe0e685f7866c26a4
---



## Motivation

Milestone 2 needs an initialized project to prove its glue through NEW before
bounded next/run work. The design's doctor and AGENTS.md clauses require the
public doctor command and managed-guidance repair. The checks backend and
emitted C14–C16 explanations now pass; connect them promptly through app, using
ordinary temporary projects, local hooks and bounded fake agents.

## Interface



## Added

### APP-062
doctor MUST discover the nearest ancestor fireweed.yaml and invoke checks.Doctor for that project root.

### APP-063
doctor MUST supply the actual invoking executable as DoctorRequest.Bin so scratch hooks query the running NEW CLI.

### APP-064
doctor --no-agent MUST skip configured worker and audit commands while retaining hook canary and recovery checks.

### APP-065
Plain doctor output MUST identify each reported probe's outcome and retain reported diagnostic codes, warning status and named failure details.

### APP-066
doctor MUST return exit 0 when no error diagnostics remain, including warning-only results, exit 1 for static or probe failures, and exit 2 for usage or configuration-loading errors.

### APP-067
doctor --json MUST emit schema 1 with an object containing a checks array on success or null data on failure, retaining code/message/path/warning diagnostics and leaving stderr empty.

### APP-068
agents-md --write MUST discover the nearest ancestor fireweed.yaml and refresh AGENTS.md at that project root.

### APP-069
agents-md --write MUST regenerate the current fireweed managed block while preserving text outside its markers exactly.

### APP-070
agents-md --write MUST create AGENTS.md when absent or append the managed block when neither marker exists.

### APP-071
agents-md --write MUST reject unpaired managed-block markers with corrective P01 and exit 2, preserving the existing text.

### APP-072
agents-md --write MUST print confirmation and exit 0 on success, or exit 2 with corrective U01 for usage, C01 for missing/invalid fireweed.yaml or P01 for operational file/setup failures.

### APP-073
agents-md --write --json MUST emit schema 1 with confirmation string data and empty diagnostics on success, or null data and code/message diagnostics on failure, leaving stderr empty.

## Modified



## Removed



## Decisions

Doctor CLI leaves request Timeout unset, using the backend's existing defaults;
no new timeout flag is introduced. Successful doctor JSON checks use fields name,
module, passed and detail, with diagnostics in the outer envelope.

## Rejected alternatives



## Notes

Reuse checks.Doctor and existing app root discovery, diagnostics/JSON helpers and
managed-block update logic. Keep app's public Run signature and other modules
unchanged; no new interface or stub is needed. Remove only these two commands
from the under-construction dispatcher. Preserve existing generic usage/JSON
contracts: doctor accepts --no-agent, agents-md requires --write, and these flags
are rejected on other commands. Retain assigned static causes for non-configuration
loading failures as APP-037 requires. Missing/extra arguments and unknown options use
corrective U01 before running hooks or writing guidance.

Use public-interface evals and compiled NEW CLI fixtures. Reuse an initialized
Go fixture to establish actual-executable FIREWEED_BIN wiring and canary recovery;
use local fake commands for agent/wrapper behavior and warning-only success.
A representative named failed hook plus retained diagnostics is enough at app;
do not repeat the checks backend's canary/error matrices or three-minute timeout
regression here. Reuse fixtures and assertions across IDs, keep own Verify under
five minutes and all cases bounded. No authenticated agent or bootstrap access.
Exercise guidance replacement with surrounding text, missing-file/absent-block
cases and the existing malformed-marker rejection; keep other project files.

Audit phase uses only this change, projected specs, evals, retained public
interfaces and governing audit rules, without implementation/history reads.
Report only exact missing assertion clauses for assigned IDs. The controller
owns phase transitions, commits, spec application and archive; workers never
commit or change frozen evals/interfaces during implementation.

## Provenance

- evals 2026-10-28T20:56:59.338624Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=b26da634732c7f1c7dc600115582a63cb8f82e1fd2bd1d522d608fe837a04ba7 tests=152 failing=5 passing=147
- audit 2026-10-28T20:58:30.579543Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=0fb044cfaca1b9a844a25b4c86223177e52f2f35fb149eddf044e83822665c88 tests=0 failing=0 passing=0
- evals 2026-10-28T21:02:16.248214Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=b26da634732c7f1c7dc600115582a63cb8f82e1fd2bd1d522d608fe837a04ba7 tests=153 failing=6 passing=147

## Audit

- APP-062: "doctor MUST discover the nearest ancestor fireweed.yaml and invoke checks.Doctor for that project root." — missing assertion that doctor selects the closer project when two ancestor projects have fireweed.yaml.
- APP-068: "agents-md --write MUST discover the nearest ancestor fireweed.yaml and refresh AGENTS.md at that project root." — missing assertion that guidance is refreshed in the closer project when two ancestor projects have fireweed.yaml.
- APP-070: "agents-md --write MUST create AGENTS.md when absent or append the managed block when neither marker exists." — missing assertion that either creation path writes the complete managed block, including its end marker and current guidance.
