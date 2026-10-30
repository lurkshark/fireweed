---
bases: {}
module: workflow
---



## Motivation

NEW pr-body <range> is still unimplemented. The design CLI table, Cost is measured,
Tags and boundary opt-out clauses require a review description with range changes,
review-relevant edits, audit/spec/Decision information, visible exceptions and
budget/opaque-artifact warnings. Build only this missing report, then explain
emitted B04 and wire app immediately with an actual NEW proof. Owner approved
three changes starting at count1, stopping local at count4 without push or drill.

## Interface

Add public PRBody(root, revisionRange string) (PRBodyReport, error) through a thin
api.go facade and model.PRBodyReport containing Body string and Diagnostics
[]project.Diagnostic. All collection/rendering behavior belongs under internal/.
Use existing Git/document/zone/project.LoadAt/cost interfaces; no dependency,
existing API or model behavior changes. An evals-only new stub may return a
non-P01 not-implemented error so tests compile and genuinely fail.

## Added

### new:pr-body-range
PRBody MUST render a Markdown description for an explicit Git commit range, identifying its resolved endpoints.
Project configuration, module/spec metadata and change documents MUST come from the range endpoint rather than later HEAD or working-tree edits.

### new:pr-body-changes
PRBody MUST identify each pending or archived change present at the endpoint whose document identity was touched in the range, including its path, module and change kind.
An empty selection MUST be reported explicitly.

### new:pr-body-evals
For every selected change, PRBody MUST include its recorded eval result counts from Provenance and identify its eval-file edits.

### new:pr-body-author-edits
PRBody MUST identify interface edits, new disposable stubs added in evals, module manifest/other-file edits and worker_writable edits associated with each selected change.

### new:pr-body-preserved
PRBody MUST identify preserved-file edits associated with selected changes.

### new:pr-body-review-specs
PRBody MUST list every review-tagged Added or Modified spec of each selected change with its ID and normative text.
It MUST not list an unrelated pre-existing review spec as changed.

### new:pr-body-audits
PRBody MUST include each selected change's audit_rounds and recorded audit findings, including findings in earlier phase commits that were corrected before the endpoint.

### new:pr-body-decisions
PRBody MUST include the selected changes' added and removed Decisions.

### new:pr-body-exceptions
PRBody MUST list all quarantined endpoint specs with module, ID and reason, every endpoint boundary opt-out with module and reason, and every blocked pending endpoint change with its reason.

### new:pr-body-costs
PRBody MUST show the latest retained measurements and associated B01–B03 diagnostics for modules represented at the endpoint, without fabricating missing measurements.
Saved local cost snapshots MUST remain the measurement source rather than later file counts or recalculated budgets.

### new:pr-body-opaque-additions
PRBody MUST list every binary file and every file larger than 256 KB newly added to a disposable zone between the range endpoints, identifying path and byte size.
Each such file MUST produce a B04 warning in Diagnostics and the Markdown body; ordinary smaller text additions MUST not produce B04.

### new:pr-body-range-errors
An empty or unresolvable range MUST return an error with P01 and the rejected range or resolution cause.
A valid empty range MUST return a report rather than an error.

## Modified



## Removed



## Decisions

Selected changes are identities touched by range commits and present at its end;
removed change documents are absent from that endpoint report. Review context for
a selected change includes its phase commits reachable at the endpoint, so a
range covering only its final phase retains earlier eval/stub/audit information.
Binary follows Git's text/binary classification; 256 KB is262144 bytes, with
strictly larger text files warned. File categories use that phase's declared
metadata; endpoint exceptions use the endpoint's declarations.

## Rejected alternatives

No live Status reuse for endpoint project data, second graph/index, unrelated
library completion, PR publication, new diagnostics beyond emitted B04, or
bootstrap implementation input.

## Notes

Fresh Sol/high eval/impl workers and isolated restricted independent auditors;
max2 persisted finding round trips, re-audit only prior finding IDs. Own Verify
under5m/full under15m; representative public fixtures only. No process/convention
observers or calibration tests. All phase transitions/claims/commits belong to
the controller; use write-capable native Done so Git commit persistence succeeds.
Workers never commit or modify protected/other-module files. No phase exception.
Cost snapshots are local records; endpoint source data is committed. Stop on a
native block or required dependency expansion. After this report, only B04
Explain and immediate app caller/NEW proof remain in the approved count4 slice.
