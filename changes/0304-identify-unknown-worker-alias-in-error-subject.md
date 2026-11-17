---
module: project
bases:
    PRJ-086: 4cba5d9b
already_satisfied:
    - PRJ-086
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
