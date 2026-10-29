---
already_satisfied:
    - APP-009
bases:
    APP-009: e5447502
module: app
---



## Motivation

The design's Resuming section requires start --from-attempt <n>. Workflow now
provides the verified StartFromAttempt operation, but NEW rejects the option.
Wire that caller immediately, then use the retained independent-run Go fixture
to load Alpha's failed impl work, keep its run/phase/IDs and HEAD, correct it and
archive through NEW done. This closes saved-attempt resume through the CLI;
agent-commit undo, usage/time limits and other recovery remain separate gaps.

## Interface

No public facade or model change. Parse the start-only option and compose the
existing workflow.StartFromAttempt and workflow.Context under app disposable
code, retaining existing start text/JSON and diagnostic conventions.

## Added

### new:start-from-attempt-command
start <change> --from-attempt N MUST discover the project root and invoke
workflow.StartFromAttempt with the selected change and N. On success it MUST
return workflow.Context in the existing plain or JSON start shape with exit 0.

### new:attempt-option-validation
For start, --from-attempt MUST require one positive integer and reject missing,
noninteger, nonpositive or repeated values with corrective U01 and exit 2 before
mutation. Its value diagnostics MUST describe valid start usage. Other commands
MUST reject --from-attempt with U01 and exit 2.

## Modified

### APP-009
Without --from-attempt, start <change> MUST discover the root, invoke
workflow.Start and print first-step workflow.Context with exit 0, assigning IDs,
run and phase without committing.

## Removed



## Decisions



## Rejected alternatives



## Notes

APP-009 merely qualifies the unchanged ordinary-start path; already_satisfied
reuses existing assertions rather than manufacturing a failure. The new option
has genuine fail-first evals. Preserve all current default start, warnings,
exit/status/cause and schema1 context/null-data behavior, including lifecycle
P01/P02/P07 propagation and all other commands/options. No hooks/agents during
start, no extra commit or runtime reassignment; restoration belongs to workflow.

Use public app API fixtures and the real public lifecycle to create saved refs
and a committed author unblock. Prove N forwarding with distinct saved work,
root discovery from a nested directory, restored work and unchanged HEAD/run/
phase/IDs/owned claim, both context output shapes, and lifecycle rejection with
corrective code/null JSON. Cover the specified invalid forms and start-only
restriction without mutation. Keep cases representative, under 15 per spec;
reuse fixtures and every existing applicable assertion. No observers or failure
matrix. Own Verify below five minutes, full below fifteen.

Evals only in app's eval zone; the existing API already compiles, no interface
or stub change needed. Real parser/delegation edits follow isolated assertion
audit. Workers never commit, edit SPEC/protected docs/other modules, change
counter/claim/runtime fields in this checkout, or run governing workflow
commands. Fixture-local public workflows/Git with explicit identities are
allowed. No bootstrap source/tests, deleted code, saved rebuild implementation
or rebuild-v1 as an oracle. No drill, phase-write exception, README redesign,
status/recovery-library expansion or unrelated cleanup.

After application rebuild NEW and run the prepared saved-resume proof against
the retained Go fixture. Remove that fixture after successful correction/archive.
