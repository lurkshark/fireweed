---
bases: {}
module: app
---



## Motivation

Milestone 2 needs an initialized project to prove its glue through NEW before
bounded next/run work. The design's doctor and AGENTS.md clauses require the
public doctor command and managed-guidance repair. The checks backend and
emitted C14–C16 explanations now pass; connect them promptly through app, using
ordinary temporary projects, local hooks and bounded fake agents.

## Interface



## Added

### new:doctor-root
doctor MUST discover the nearest ancestor fireweed.yaml and invoke checks.Doctor for that project root.

### new:doctor-bin
doctor MUST supply the actual invoking executable as DoctorRequest.Bin so scratch hooks query the running NEW CLI.

### new:doctor-no-agent
doctor --no-agent MUST skip configured worker and audit commands while retaining hook canary and recovery checks.

### new:doctor-output
Plain doctor output MUST identify each reported probe's outcome and retain reported diagnostic codes, warning status and named failure details.

### new:doctor-status
doctor MUST return exit 0 when no error diagnostics remain, including warning-only results, exit 1 for static or probe failures, and exit 2 for usage or configuration-loading errors.

### new:doctor-json
doctor --json MUST emit schema 1 with an object containing a checks array on success or null data on failure, retaining code/message/path/warning diagnostics and leaving stderr empty.

### new:guidance-root
agents-md --write MUST discover the nearest ancestor fireweed.yaml and refresh AGENTS.md at that project root.

### new:guidance-preserve
agents-md --write MUST regenerate the current fireweed managed block while preserving text outside its markers exactly.

### new:guidance-missing
agents-md --write MUST create AGENTS.md when absent or append the managed block when neither marker exists.

### new:guidance-reject
agents-md --write MUST reject unpaired managed-block markers with corrective P01 and exit 2, preserving the existing text.

### new:guidance-status
agents-md --write MUST print confirmation and exit 0 on success, or exit 2 with corrective U01 for usage, C01 for missing/invalid fireweed.yaml or P01 for operational file/setup failures.

### new:guidance-json
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
