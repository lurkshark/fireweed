---
audit_rounds: 0
bases: {}
module: checks
phase: audit
run: 20261030T025913-b1814486063e4bf2bf4245fa
---



## Motivation

Design History checks and Merge strategies require NEW check-commits to enforce
H01–H04 using file contents and recognize squashes without false H02 findings.
The command is still unimplemented. PRJ-047–050 now provide recorded project
snapshots; consume that dependency rather than reimplement project parsing.
Build only this history checker, then its missing H01–H04 explanations and
immediate app caller/actual NEW explicit-range proof. Count2 since passing review;
this is the second consecutive history-slice dependency, with normal checkpoints
before further consecutive behavior work. No full-slice success until actual NEW.

## Interface

Add thin public History(root, revisionRange string) (HistoryReport, error).
HistoryReport in checks/model exposes Commits []string, Squashed []SquashedChange
and Diagnostics []project.Diagnostic. SquashedChange exposes Commit and Change
strings (full OID and repository-relative change-document path). Public aliases
and facade only; behavior under checks/internal using existing dependencies.
No new process primitive or dependency. API returns rule findings in Diagnostics;
operational failures are errors rather than clean reports.

## Added

### CHK-070
History MUST list full commit IDs selected by the supplied Git revision range in oldest-first topological order.
For an empty range it MUST use the merge base of HEAD and the locally resolved configured base through HEAD.

### CHK-071
History MUST identify worker commits from new phase assignments in change-file contents across pending/archive moves, rather than commit subjects.
Deleting a change document without assigning a phase MUST not by itself classify the commit as a worker commit.

### CHK-072
History MUST report H01 when a worker commit changes a protected path or a project playbook override.

### CHK-073
History MUST report H02 when an impl-to-applied phase commit changes evals or interface zones.

### CHK-074
History MUST report H03 when any commit modifies, deletes or renames an existing append_only path, including merge-resolution changes identified by combined Git diffs.

### CHK-075
History MUST report H04 when a commit modifies an existing file under changes/archive.
Deleting an archived change file MUST remain allowed by H04.

### CHK-076
History MUST recognize each new applied phase whose parents lack an impl phase as squashed, listing its commit and change without a warning.
H01 and H02 MUST not apply to those squashed commits.
H03 and H04 MUST still apply.

### CHK-077
History MUST return each rule finding as a Diagnostic with its H01–H04 code, repository-relative Path and full commit ID in Message, using the recorded before/after zone maps and the worker change's projected set values.

### CHK-078
History MUST return an error when a revision range or required recorded project/change document cannot be read or parsed.

## Modified



## Removed



## Decisions

A merge merely retaining a phase value already held by a parent is not a new
phase assignment; combined merge diffs still enforce append-only files.
Squashed entries identify the change's repository-relative after-path; ordinary
non-merge phase transitions compare the sole parent.

## Notes

Use public checks/project.LoadAt/gitstore/contracts interfaces and ordinary temp
Git fixtures with explicit identities. Cover actual content-based phase and
impl/archive moves, representative protected/override and eval/interface paths,
append-only author modify/delete/rename and a merge-only range whose resolution
is caught, archive modification/deletion, clean controls and multi-change squash
with H03/H04 still enforced. A representative recorded-zone/set fixture prevents
current-config or changed-zone bypass. Default range can reuse public BaseRef;
no remote fetch, fixture agents or temporary product checkouts. No bootstrap
source/tests, rebuild-v1, deleted/abandoned/saved implementation or old worker-log
oracles. Don't imitate controller JSON quirks or add a feature/error matrix.
Only checks interface/models/evals/new compiling stubs in evals, freeze, separate
restricted Sol/high audit, then implementation/full native verification. Keep
shared diagnostic helper assertions without tagging every incidental assertion
as every new ID; each testcase needs its actual focused IDs. Preserve existing
assertions. Workers never commit or advance phases/release claims, alter protected
files/other modules, raise limits or add instrumentation. Own under5m/full under15m;
only representative cases/spec near15. No extra doctor probes, drills, CI or
PR-body implementation in this dependency. Accurate H01–H04 explanations and
app caller remain required before observable CLI completion.

## Provenance

- evals 2026-10-30T03:45:47.55804Z agent=manual fireweed=0.1.0 playbook=b26da634732c7f1c7dc600115582a63cb8f82e1fd2bd1d522d608fe837a04ba7 tests=524 failing=20 passing=504
