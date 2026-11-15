---
bases:
    CHK-063: d58f1f3a
    CHK-064: 5165fb12
module: checks
---

## Motivation

Projects using named commands need doctor to probe the commands that lifecycle
workers and auditors will actually use.

## Interface

## Added

### new:doctor-agent-alias-error
Doctor MUST report an agent selection error as a failed C15 probe identifying the setting at fault and skip invocation for that unresolved role.

## Modified

### CHK-063
Doctor MUST resolve the worker and audit commands from the project configuration and run each resolved command once in scratch through agent.wrap with a fixed-reply prompt, requiring exit 0 and the requested reply for success.

### CHK-064
When the audit probe inherits the worker default, Doctor MUST probe that command for audit and emit a C15 warning identifying the shared-command blind spot.

## Removed

## Decisions

## Notes

Move the doctor eval fixtures to named commands and retain their existing
probe, wrapper, prompt, timeout, and log assertions.
