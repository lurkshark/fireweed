---
module: workflow
phase: audit
run: 20261112T041015-0e20df6cfb856bcd117e370d
---

## Motivation

The design's Blocking section says Fireweed records a worker-declared stop
with its reason and Questions. The 0214 implementation worker reported a
blocker and exited successfully without applying the CLI migration. The
controller saw passing evals and archived the change. Workers need a
documented signal that prevents a phase from advancing.

## Interface

The held change document carries the worker's `blocked` reason and
`## Questions`. Done uses the existing blocked outcome and P06 diagnostic.

## Added

### FLOW-137
For a held evals or impl step with a nonempty `blocked` reason and nonempty `## Questions`, Done MUST save the attempt and block the change before it can advance. The block commit MUST preserve the reason and Questions alongside the approved lifecycle state. Done MUST release the claim and return P06.

### FLOW-138
Context for a held evals or impl step MUST tell the worker how to declare a block using the current change's `blocked` reason and `## Questions`.

## Modified

## Removed

## Decisions

## Notes

Use the existing block and saved-attempt path. Validate the held claim and
approved state before recording a block. Carry the worker's Questions into
the state-only block commit and keep other step edits in the saved attempt.
The controller owns the commit. Worker stdout is informational.

Evals: through public Done, submit a declared block in a held impl step with
an otherwise passing eval and a step edit. Assert the P06 result, preserved
reason and Questions, unchanged phase and run, released claim, saved edit and
absence of an archive entry. Exercise the declaration in evals and the guidance
through public Context for both phases. Run after 0230.

Fail-first evidence from the workflow evals: FLOW-137 advanced the declared
evals block to audit and archived the declared impl block after full Verify
passed in both fixtures before the declaration. FLOW-138 found the declaration
instructions missing from Context in both phases. Done rejected the fixture's
unapproved author edit with P03 before saving an attempt or committing.

The implementation can use `packages/workflow/internal/lifecycle/lifecycle.go`
for Done and Context and `packages/workflow/internal/lifecycle/block.go` for
the shared saved-attempt path.
## Provenance
- evals 2026-11-12T04:21:12.05901Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6.1-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=6a9c2d506363e3539cc66efd37bc6a06701249541dc4e4f01ed3aa6454bbb864 tests=197 failing=6 passing=191
