---
audit_rounds: 0
bases: {}
module: workflow
phase: applied
run: 20261029T184119-8942b3694b847f629b8cc91b
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

### FLOW-080
Start MUST set phase to impl for an unstarted change with regenerate true.

### FLOW-081
Before returning from successful Start of an unstarted regenerate change,
Fireweed MUST delete only the target module's files classified as disposable
by the approved projected zone map, retaining all other files.

### FLOW-082
For a regenerate change, Start MUST emit warning P10 when the module has no
passing drill dated after its latest applied change's implementation Provenance.
A newer passing drill MUST suppress that warning.

### FLOW-083
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

## Provenance

- evals 2026-10-29T18:50:45.672515Z agent=manual fireweed=0.1.0 playbook=b26da634732c7f1c7dc600115582a63cb8f82e1fd2bd1d522d608fe837a04ba7 tests=368 failing=6 passing=362
- audit 2026-10-29T18:53:37.27062Z agent=manual fireweed=0.1.0 playbook=0fb044cfaca1b9a844a25b4c86223177e52f2f35fb149eddf044e83822665c88 tests=0 failing=0 passing=0
- impl 2026-10-29T19:11:35.237994Z agent=manual fireweed=0.1.0 playbook=07554a3249d8fca0ba70df7504135f297491465caa63458734a4b8351734d0e5 tests=368 failing=0 passing=368

## Audit

No findings
Checked IDs: FLOW-080, FLOW-081, FLOW-082, FLOW-083
