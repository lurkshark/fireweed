---
module: workflow
phase: audit
run: 22da0470ce4466705c23c958
---

## Motivation

Compose the checked manual start operation from the existing local repository, static queue, eligibility, allocation, claim, context and approval contracts. Stage 2 needs a real transaction that leaves an owned started change without committing, including valid resume and default regeneration. The NEW CLI binding, verification/done and saved-attempt restoration remain follow-up changes; this API is the ordinary start operation without a from-attempt argument.

## Interface

Add Start(start, slug string) (StartResult, error) to api.go as a thin disposable implementation forwarder. Retain StartResult in model/start.go and expose its alias: Context WorkerContext with json context, Claim Claim with json claim, and Diagnostics []checks.Diagnostic with json diagnostics. Existing APIs and models retain their contracts. Evals may add only new disposable stub files and the required facade/model declarations; all behavior belongs under internal/.

## Added

### FLOW-032
Start MUST select one exact valid claim slug from the current repository's queued changes, discovering canonical ordinary or linked roots from nested start directories. Invalid selectors MUST fail with corrective C05 before repository access. Missing or ambiguous targets, invalid project/static queue metadata, unavailable source files and malformed selected state MUST fail with useful corrective diagnostics before creating a claim or changing working files. A dirty tracked, staged or untracked working tree MUST fail with corrective P01. Static warnings MUST remain visible on success. Start MUST use the current committed project and MUST NOT substitute a claim snapshot to bypass these preconditions.

### FLOW-033
Start MUST reject a blocked or otherwise ineligible selected change with corrective P03, using Eligible's complete inclusive queue semantics. Before starting, this canonical worktree MUST hold no existing claim, even for a different or no-longer-queued change. A selected claim held by another linked worktree MUST be rejected with corrective P02. A successful start MUST atomically create exactly one selected claim, recording canonical owner, exact slug, run and the full original HEAD commit. Concurrent competing starts MUST leave one winner; a loser MUST NOT alter the winner's claim or files. Malformed or unreadable existing claim metadata MUST fail closed with useful corrective P02.

### FLOW-034
For a fresh selected change, Start MUST use Allocate's late IDs and all-occurrence bindings, create a new nonempty run identity, and write run plus phase without committing. Spec changes and fixes MUST start in evals; maintenance and regenerate changes MUST start in impl. Different fresh runs MUST have different run identities. The saved started change, returned Context.Change and claim State MUST agree under public Markdown normalization, including retained author structure, metadata, typed edits and identity. The claim State MUST contain the exact saved change bytes. Allocation MUST retain removed/archived/previously assigned reservations and MUST NOT renumber unrelated changes.

### FLOW-035
For an eligible committed evals, audit or impl change with an existing run and no blocked field, Start MUST preserve its run, phase and every assigned ID while obtaining a new claim against current HEAD. It MUST NOT allocate new IDs or regress the phase. Removing blocked from a committed change MUST permit the same resume semantics. Incomplete run/phase pairs, applied queued state and unresolved local Added handles in a started change MUST fail before mutation. A claim alone MUST NOT make an uncommitted active worktree eligible for another start.

### FLOW-036
Start MUST return warning diagnostics P04 when HEAD does not contain the locally known base tip and independently when the selected authored content is absent or unapproved on that base. It MUST resolve the chosen base commit locally once, read the exact selected file from that pinned commit without replacement-object or lazy-fetch interpretation, and use ApprovalMatches against the started content so late IDs and canonical layout do not revoke approval. A malformed, different-target or absent base change MUST be unapproved. Unavailable base ancestry/content MUST never imply approval. Stale or unapproved content MUST warn rather than prevent an otherwise valid manual start. Base warnings MUST retain useful path and corrective guidance and MUST NOT change the configured base or any base ref.

### FLOW-037
On success Start MUST return a complete BuildContext result for the exact newly owned started state, including once-only projected zones/specs/dependencies and the selected phase playbook, with a non-null Diagnostics array and the exact stored Claim including its object token. Failure to construct context, including an unsafe or unreadable required retained file or playbook override, MUST fail the whole start. Start MUST NOT certify a worker step, invoke verification hooks, run an agent or auditor, or advance beyond its initial or resumed phase.

### FLOW-038
Any failure after reservation MUST return a zero StartResult and roll back this operation's working-file writes/deletions and release only its own claim token. The selected change's original exact bytes and permission bits, and every deleted tracked file's contents/type/mode, MUST be restored when still owned by this operation. A replacement claim MUST survive rollback; a concurrently replaced change or recreated deleted path MUST NOT be overwritten. Lost ownership or incomplete rollback MUST produce useful corrective recovery details rather than falsely report a clean rollback. On success the selected change MUST retain its original permission bits. A selected change symlink or unsafe path MUST be rejected before mutation; temporary files MUST be cleaned after success and recoverable rejection.

### FLOW-039
For regenerate true without saved-attempt restoration, Start MUST delete only tracked files classified Disposable under the selected inclusive projected module zones. Protected, SPEC, eval, interface, append-only, preserved, other-module, other and ignored/untracked paths MUST survive, including stricter-zone overlap and metadata overrides. It MUST unlink a disposable leaf symlink without following or changing its target, reject unsafe symlink ancestors before deletion, and never recursively remove a directory containing retained files. Resume of a regenerate change MUST apply the same deletion rules. Start MUST return D01 warning when the target has no readable passing .fireweed/drills/<module>.md report whose base_commit is a locally known ancestor of HEAD with no reachable later target-module apply commit. Missing, malformed, failing, unavailable or stale reports MUST warn; a current passing report MUST suppress D01 even when later applies affect only other modules. Report content, dates and drill refs MUST remain unchanged.

### FLOW-040
Start MUST remain local, deterministic in all fields except fresh run identities, and noninteractive. It MUST NOT fetch, push, create branches/worktrees, commit, move HEAD, stage/reset the index, mutate other refs/claims, invoke Git hooks or fsmonitor, execute hooks/agents/bootstrap, use the network, or change parent cwd/environment. Ordinary and linked SHA-1/SHA-256 repositories MUST retain all unrelated tracked/ignored files, file modes, index bytes, refs, common-Git bookkeeping, attempts and drill reports. Failure paths MUST preserve the same unrelated state. The operation MUST use bounded local subprocesses and retain useful failed-command details without parsing stderr as a successful Git stdout protocol.

## Modified

## Removed

## Decisions

This is the default manual transaction; saved-attempt restoration belongs to Stage 3 and will have a separately authored API. Start's normal approval warnings do not grant unattended approval. The current rebuild goal authorizes local manual execution but does not assert a human merge. Use existing public validators and semantic approval comparison rather than a weaker start-only interpretation. The claim pins original HEAD and exact started bytes; it does not advance HEAD. Rollback uses ownership checks and compare-and-swap release, preserving a replacement winner and explaining incomplete recovery. No implementation or eval belongs in protected infrastructure.

Regeneration rollback holds only the exact files deleted by this transaction and does not expose their contents in worker context or persist an old-code oracle. Generated replacement code and required real-agent drills remain future execution gates. A drill report is current only when its recorded passing base covers the module's reachable apply history; another module's later apply does not stale it. Current passing report metadata uses the reference's result: pass and base_commit fields. Histories are interpreted through archived change state and module identity, never commit subjects.

## Rejected alternatives

Starting from dirty work risks committing unrelated work and violates the manual lifecycle. Treating an existing same-owner claim as implicit resume could overwrite active work. Writing state before atomic reservation loses linked-worktree exclusion. Leaving allocated bytes or partial regeneration deletions after a context error breaks recoverability. Releasing by slug without its exact token can delete a concurrent winner. Approval by filename or spec hashes alone loses authored intent; unconditional approval-layout exemptions would hide allocator regressions. Running hooks during start conflates startup with verify. Recursive directory removal and matching disposable globs without zone precedence can destroy retained contracts.

## Provenance

- evals | 2026-10-23T07:49:51Z | agent: manual | fireweed: 0.4.0 | playbook: 5867363302d4c874074520977b1627d26991d95ee2fbacb9ff63e827661edd0b | tests: 1552 | failing: 113 | passing: 1439
