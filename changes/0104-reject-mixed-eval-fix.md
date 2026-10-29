---
bases: {}
module: project
---



## Motivation

The design defines eval_fix without spec edits or fixes. Checkpoint0099–0103
found NEW accepts mixed metadata, allowing eval_fix's passing-evals branch to
skip fail-first for fixes/spec changes. The owner approved this scoped correction
gate. Reject the invalid combination centrally before Start/Verify/Done, which
already call project.Check. After the immediate C07 explanation correction,
prove actual NEW rejects the mixed change without starting an agent or claim,
then obtain the checkpoint's scoped re-review before continuing to release CLI.

## Interface

No public API/model changes. Extend existing Check diagnostics with C07 for
this invalid change structure. No workflow behavior or facade changes.

## Added

### new:eval-fix-exclusive-kind
Check MUST report error C07 for a pending change containing eval_fix together
with fixes or any Added, Modified or Removed spec edits.

## Modified



## Removed



## Decisions



## Rejected alternatives



## Notes

Add representative public Check regressions for eval_fix mixed with fixes and
with spec edits, while valid separate eval_fix/fix/spec changes stay accepted.
Use otherwise valid normative fixtures and stamp bases when required so the
only rejection is the new kind conflict. Prefer an Added-only spec edit for
that representative branch rather than a base-hash distraction. Preserve all
existing helpers/assertions and pure-kind behavior. No schema/type-validation
matrix, new kind framework, new diagnostic number, interface, workflow/app or
other-module/protected edits, reorganization, instrumentation or drill.
This addresses the recorded bypass only; no unrelated hardening.

Evals/interfaces freeze before impl. Workers never commit or advance workflow.
Coordinator owns native full checks, isolated audit/application, the immediate
C07 explanation follow-up and scoped checkpoint re-review. Own under5m/full
under15m; do not raise bounds. Five prior applications remain counted; these
two correction applications are explicitly owner-approved, not a count reset,
phase exception or ordinary fifth-checkpoint bypass.
