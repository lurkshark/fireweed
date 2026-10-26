---
module: workflow
---

## Motivation

The NEW modules command must publish the approved inclusive graph while a local change is held, including when a real boundary hook queries FIREWEED_BIN. The existing public approval loader intentionally supports an unclaimed current-source fallback. Calling it after a separate optional claim lookup can therefore fall back to worker metadata if ownership disappears. Provide a strict optional source selection that callers can safely compose without running Verify recursively. D042 records the equivalent-fixture graph mismatch; app integration belongs in a later single-module change.

## Interface

Add LoadHeldApproval(start string) (project.Snapshot, Claim, bool, error) to the retained workflow facade. It uses the existing public Snapshot and Claim types. Preserve all existing signatures and LoadApproval behavior. Add only a new private stub file/package needed for EVALS to compile.

## Added

### new:held-selection
LoadHeldApproval MUST discover the canonical checkout and common Git directory from start, including nested ordinary and linked SHA-1/SHA-256 worktrees, and select the unique claim owned by that checkout. No owned claim MUST return a zero Snapshot, zero Claim, false and nil error without loading current project configuration, SPEC or queue documents. Valid foreign claims MUST NOT count as owned. An unavailable or malformed claim payload that cannot be safely classified, or ambiguous local ownership, MUST return zero values, false and a useful corrective error.

### new:held-source
For a unique valid owned claim LoadHeldApproval MUST return true, the exact stored Claim including its object token, and the same approved Base snapshot with the exact started State inserted that claimed LoadApproval provides. It MUST use only the selected committed approval and state as project authority, retaining the canonical active Root and approved project configuration, modules, changes and archive. Current worker configuration, SPEC and queue edits MUST NOT replace approved metadata. The returned snapshot MUST remain before inclusive projection, so public Project before the selected change followed by ApplyChange for that change applies preceding and selected effects exactly once and excludes later effects.

### new:held-fail-closed
LoadHeldApproval MUST bind approved source loading to the exact selected claim token. After loading and before returning success it MUST revalidate the unique locally owned selection and that exact token; disappearance, replacement or newly ambiguous local ownership observed at that validation MUST return zero values, false and a useful corrective error. It MUST NOT silently return unclaimed absence or fall back to current project source after selection. Unavailable required claim payload, Base or project-document objects, malformed started State or other selected approved content, and unsafe selected source paths MUST return zero values, false and useful corrective errors. Unavailable unrelated non-claim document bodies or unneeded ancestors MUST NOT prevent a locally complete selected result.

### new:held-observer
Successful and rejected LoadHeldApproval calls MUST remain local, bounded, noninteractive and read-only: no lazy fetch, push, branch/worktree creation, hooks, agents, bootstrap or network commands; no file, ignored-content, index, HEAD, ref, claim, attempts or drill-report mutations; no parent cwd or environment changes. Required local Git failures MUST retain useful command details rather than treating stderr as successful source data. Repeated successful results MUST have independent writable nested slices/maps and deterministic approved values.

## Modified

## Removed

## Decisions

This optional source API does not verify a worker step or validate the effective queue/graph. Callers retain public Project/ApplyChange/Check composition and must not recursively invoke Verify merely to read the graph. Existing LoadApproval retains its documented unclaimed behavior. Optional implicit selection cannot safely ignore an unavailable claim payload: its worktree identity is unknown, so it could be a second local claim. This differs from an explicitly named LoadApproval selection, which can ignore unrelated claim bodies. Absence can be established using local Git ownership metadata before current project documents are loaded.

The rebuild goal authorizes local authored execution and controller-owned phase/archive commits. Use the pinned bootstrap as controller and separately built NEW binaries as subjects; no push, publication or tags, no claimed human merge.

Final ownership validation is the successful read operation's validation point; the API cannot prevent a later independent owner update. Immutable-object caching is allowed but cannot skip that validation. Evals must not require a particular Git command/body-read strategy. Race/fault observers need independent live controls and fresh-source premises; a stub that never selects anything does not establish an ownership-disappearance race. Positive locally complete sources with missing unrelated bodies/ancestors cover the fail-closed/locality contract without treating a fixture-premise failure as fail-first. The saved core/race drafts are planning inputs only and must be compiled, calibrated and corrected during EVALS.

## Rejected alternatives

Separate Claims plus fallback-capable LoadApproval admits a disappearance race. Loading current project metadata before checking ownership lets worker edits obstruct approved graph queries. Running Verify to find ownership recursively executes boundary hooks. Returning an already inclusive projected snapshot would risk applying the selected change twice in existing public composition.
