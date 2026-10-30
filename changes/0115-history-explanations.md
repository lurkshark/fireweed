---
audit_pending:
    - AST-036
    - AST-037
audit_rounds: 2
bases: {}
module: assets
phase: impl
run: 20261030T044031-868ada5b81566b0bb0c4bb59
---



## Motivation

NEW check-commits must explain the H01–H04 diagnostics now emitted by public
checks.History. Design History checks defines these rules; Explain currently
rejects their codes. Add only accurate causes/actions before the immediate app
caller and actual NEW range proof. Owner explicitly approved this consecutive
history continuation and fifth-change review; this becomes application4 after
0112–0114, leaving the caller as5. No history engine or unrelated diagnostic work.

## Interface

No API/model changes. Extend the existing Explain lookup under assets/internal.

## Added

### AST-034
Explain MUST describe H01 as a worker phase commit changing a protected path or project playbook override.
It MUST advise restoring that edit and separating human-owned setup or override work into an author commit outside worker phases.

### AST-035
Explain MUST describe H02 as an impl-to-applied commit changing evals or interface files.
It MUST advise making those edits in the evals step before implementation and reworking the affected phase history before rerunning history checks.

### AST-036
Explain MUST describe H03 as modifying, deleting or renaming an existing append-only path.
It MUST advise restoring the original path/content and adding a new file for new data.

### AST-037
Explain MUST describe H04 as modifying an archived change file, while allowing archived-file deletion for a revert.
It MUST advise restoring the archive and authoring follow-up work as a pending change.

## Modified



## Removed



## Decisions



## Notes

Owner reply: "Approve continuation and fifth review" for the exact two-change
proposal. Keep count3 until this application; normal audit/full verification,
then caller/actual NEW and fifth quality/goal review/checks before push on PASS.
No new phase-write exception or drill authority. Four focused public Explain
cases can assert cause/action relationships, not merely unrelated keywords;
reuse existing pattern helpers and preserve every earlier assertion/message.
No interface/model, recipe/playbook, checker, other-module/protected edits or
new diagnostic allocations. Evals freeze before impl; fixed real Sol/high worker,
separate isolated assertion auditor, controller-owned phases/commits/archive.
Own under5m/full under15m with unchanged bounds; no matrices, instrumentation or
extra confidence probes. No bootstrap source/tests, rebuild-v1, deleted/abandoned/
saved implementation or old worker-log oracles. Return immediately to app caller.

## Provenance

- evals 2026-10-30T04:44:21.63502Z agent=manual fireweed=0.1.0 playbook=b26da634732c7f1c7dc600115582a63cb8f82e1fd2bd1d522d608fe837a04ba7 tests=680 failing=4 passing=676
- audit 2026-10-30T04:46:40.559502Z agent=manual fireweed=0.1.0 playbook=0fb044cfaca1b9a844a25b4c86223177e52f2f35fb149eddf044e83822665c88 tests=0 failing=0 passing=0
- evals 2026-10-30T04:50:03.682258Z agent=manual fireweed=0.1.0 playbook=b26da634732c7f1c7dc600115582a63cb8f82e1fd2bd1d522d608fe837a04ba7 tests=680 failing=4 passing=676
- audit 2026-10-30T04:52:46.385283Z agent=manual fireweed=0.1.0 playbook=0fb044cfaca1b9a844a25b4c86223177e52f2f35fb149eddf044e83822665c88 tests=0 failing=0 passing=0
- evals 2026-10-30T04:55:32.837924Z agent=manual fireweed=0.1.0 playbook=b26da634732c7f1c7dc600115582a63cb8f82e1fd2bd1d522d608fe837a04ba7 tests=680 failing=4 passing=676
- audit 2026-10-30T04:58:39.413317Z agent=manual fireweed=0.1.0 playbook=0fb044cfaca1b9a844a25b4c86223177e52f2f35fb149eddf044e83822665c88 tests=0 failing=0 passing=0

## Audit

No findings
Checked IDs: AST-036, AST-037
