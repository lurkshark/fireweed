---
module: project
phase: impl
run: 20261114T014126-fdca6ad831ebf4238c154164
---

## Motivation

The design's Change files frontmatter and Evals step require fixes and
eval_fix to identify real specs and already_satisfied to identify Added or
Modified specs (R04/R07/R23). On main 3b3b64d, fixes: [ALP-999] against only
ALP-001 passes check and quick eval verification because the target is dropped.
Reject invalid lists before they can bypass the requested verification.

## Interface

No public signature or type changes.

## Added

### PRJ-079
When fixes, eval_fix or already_satisfied is present, Check MUST require a list of nonempty string identities and reject other values with non-warning C07 naming the change and field.

### PRJ-080
Check MUST reject a fixes or eval_fix identity absent from the target module's projected Specs after that change, with non-warning C07 naming the change, field and identity. IDs belonging to another module MUST be rejected by the same rule.

### PRJ-081
Check MUST reject an already_satisfied identity that does not name an Added or Modified item of that change, using non-warning C07 naming the change and identity. Before Start, a local Added handle MUST be accepted as that item's identity; after assignment, its assigned ID MUST be accepted.

## Modified



## Removed

## Decisions

## Notes

Do not change review/quarantine exemptions, legitimate already_satisfied,
or the existing rejection of mixed eval_fix changes. Use public Check fixtures
for unknown and foreign IDs, wrong field types, valid current IDs, a same-change
Modified target, and a local Added handle. Include the ALP-999 regression.
Existing lifecycle coverage remains the evidence that valid fixes require a
mapped failing testcase and valid eval_fix targets retain passing coverage.
Do not create an unrelated duplicate-list restriction or require nonempty
lists: empty lists have no targets. Keep local-handle cross-change rules.

Drafted at the owner's request after the R01–R29 traceability review.
No lifecycle execution, agent invocation, drill or release is authorized by
this drafting request. Workers must obtain context before their step and use
fireweed verify; only the controller may advance or commit lifecycle work.

The public Check evals reproduced missing C07 diagnostics for PRJ-079,
PRJ-080 (including `fixes: [ALP-999]`), and PRJ-081. The valid controls and
controller verification passed in the evals phase.

`packages/project/internal/queue/validate.go` is implementation-writable for
all three specs. Check already walks the queue in order and has access to the
change edits and the module's projected items after each application.
## Audit

No findings; checked PRJ-079, PRJ-080, PRJ-081

## Provenance

- evals 2026-11-14T01:45:11.394986Z agent=codex fireweed=0.1.0 playbook=6a9c2d506363e3539cc66efd37bc6a06701249541dc4e4f01ed3aa6454bbb864 tests=229 failing=26 passing=203
- audit 2026-11-14T01:45:25.841459Z agent=codex fireweed=0.1.0 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0

