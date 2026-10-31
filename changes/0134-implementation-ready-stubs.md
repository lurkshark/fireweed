---
bases: {}
module: assets
---



## Motivation

Change 0128's eval worker put ErrorCode's temporary body in api.go, which froze
before implementation and left the next worker unable to replace it. The
design's Evals step permits new disposable stub files and says recipes show
language-specific stub patterns. Make that handoff explicit in the product:
NEW's `bin/fireweed playbook evals` should explain where replaceable stubs go,
and the Go instructions delivered by `init --recipe go` should demonstrate it.
The owner requested this product guidance after reviewing the failure.

## Interface

No public API changes. Expand the instructions returned by Playbook("evals")
and BuiltinRecipe("go").Instructions, including the bundled Go recipe file.

## Added

### new:eval-stub-handoff
The evals playbook MUST instruct workers to keep temporary stub behavior replaceable during implementation without editing frozen interface files.
It MUST require workers to identify the implementation-writable files for missing behavior before completing evals.

### new:go-stub-handoff-example
Go recipe instructions MUST show a retained api.go function delegating to a temporary stub in a new internal/ file.
The example MUST explain that implementation replaces the disposable stub while leaving the retained facade unchanged.

## Modified



## Removed



## Decisions



## Rejected alternatives

Project-only guidance: it helped 0130 and 0131, but does not reach other users.
A new static validator: explicit playbook guidance and a language-specific
example address the demonstrated failure through existing product mechanisms.

## Notes

Keep the general rule language-agnostic in the embedded evals playbook; put the
concrete Go example in the bundled Go recipe. Preserve eval_fix instructions
and existing Go setup/graph guidance. Existing worker context already includes
the matching playbook, so no workflow/context changes or new commands are needed.
Keep this change within assets; it does not change phase-zone enforcement.
