---
bases: {}
module: workflow
---



## Motivation

The design's regenerate change is maintenance starting in impl with disposable
files deleted, checked against retained contracts and normal dependents. It
also warns without a sufficiently recent passing drill. NEW currently starts
regenerate in evals without deleting implementation. Establish the workflow
path and its reconstruction context. Close the slice through actual NEW in a
Go fixture: start shows deletion/retained contracts and freshness warning,
then a bounded fake worker reconstructs through next and checked archive.
Explain the new emitted P10 warning in the immediately following assets slice.
This feature work authorizes no actual drill.

## Interface

No public API/model changes. Existing Start, Context, Verify/Done and CLI
callers compose the path. P10 is the scoped warning allocated in D008; it is
non-blocking. Read the design's report date/result and archived implementation
Provenance for freshness, not commit ancestry. No new report-writing command.

## Added

### new:regenerate-impl-phase
Start MUST set phase to impl for an unstarted change with regenerate true.

### new:regenerate-disposable-deletion
Before returning from successful Start of an unstarted regenerate change,
Fireweed MUST delete only the target module's files classified as disposable
by the approved projected zone map, retaining all other files.

### new:regenerate-drill-warning
For a regenerate change, Start MUST emit warning P10 when the module has no
passing drill dated after its latest applied change's implementation Provenance.
A newer passing drill MUST suppress that warning.

### new:regenerate-context
For a held regenerate change in impl, Context MUST supply the regenerate
playbook in place of the ordinary impl playbook.

## Modified



## Removed



## Decisions



## Rejected alternatives



## Notes

Keep ordinary spec/fix/eval_fix/maintenance and recorded-phase resume behavior,
approval/claim/clean-tree preconditions, ID/run allocation, full verification,
archive/provenance and boundaries intact. Do not delete before acquiring the
claim. Delete only project files in the target disposable zone, honoring global
protected and interface/evals/preserve/append-only precedence; never Git
metadata or another module. Use the approved snapshot, not unapproved metadata.
Existing Git/public zone facilities suffice; no new API or private imports.
Fresh-start deletion must not repeat on an already recorded phase.

Use representative public fixtures for fresh phase/deletion, retained files
and regenerate context, followed by normal checked reconstruction/archive.
Fixtures may author date/result drill reports as inputs, not run drills. Cover
missing, older passing, newer failed and newer passing reports using an actual
archived implementation Provenance date; compare dates rather than ancestry.
Use valid timestamp data from public documents, including normal YAML timestamp
values; no date-format/platform matrix or fabricated real-drill claim. Keep
coverage/hooks deterministic and distinguish new behavioral failures from
fixture/compiler errors before freeze. Canonicalize setup or compare document
semantics where exact formatting is not a requirement. Preserve all existing
helpers/assertions; add focused eval files only, no stubs needed.

No new command/recipe/hook, generic regeneration framework, scope expansion,
reorganization, other-module/protected edits, process instrumentation, real
regeneration drill or phase exception. Workers never commit or advance phases.
Evals/interfaces freeze in impl; coordinator owns native checks/audit/archive,
fresh NEW build and the actual bounded fake-agent CLI proof. Own under5m/full
under15m; do not raise limits. Add P10 explanation immediately after application
before closing this user-facing slice; no unrelated diagnostics.
