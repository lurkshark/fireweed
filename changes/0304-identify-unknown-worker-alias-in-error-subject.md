---
already_satisfied:
    - PRJ-086
bases:
    PRJ-086: 4cba5d9b
module: project
phase: impl
run: 20261117T095147-00f0776ec3ee6d95dc5d3b16
---

## Motivation

The drill CLI uses the error subject to identify an unknown worker alias and
return usage exit status 2 as required by APP-167. The project drill report in
commit `422411b` shows the CLI returning exit status 1 because the reconstructed
resolver named the alias in the error message and left the subject empty.

## Interface

Callers can retrieve the unlisted worker alias from contracts.ErrorCode on
the error returned by project.ResolveAgentCommand.

## Added

## Modified

### PRJ-086
For a worker with a nonempty alias, ResolveAgentCommand MUST select the matching valid entry from the agent.commands mapping. For an unlisted alias it MUST return a U02 CodedError with Subject equal to the supplied alias and a message naming the alias. It MUST return C01 identifying a malformed mapping or invalid selected command.

## Removed

## Decisions

## Notes

In the public unknown-worker-alias eval, assert that contracts.ErrorCode
returns U02 and a subject exactly equal to the supplied alias. Also assert
that the error message names the alias. Keep the existing C01 cases.
Current ResolveAgentCommand already supplies both the subject and message.
## Audit

No findings; checked PRJ-086

## Provenance

- evals 2026-11-17T09:52:44.879972Z agent=codex fireweed=0.2.1 playbook=6a9c2d506363e3539cc66efd37bc6a06701249541dc4e4f01ed3aa6454bbb864 tests=322 failing=0 passing=322
- audit 2026-11-17T09:53:06.465651Z agent=codex fireweed=0.2.1 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0

