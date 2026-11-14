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
Doctor MUST use the worker default for the audit probe when neither audit.agent nor audit.command is configured and emit a C15 warning identifying the shared-command blind spot.

## Removed

## Decisions
