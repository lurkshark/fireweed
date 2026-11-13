---
bases:
    FLOW-019: 0b34bdcb
module: workflow
---

## Motivation

The owner approved resolving approved-snapshot access before app and workflow regeneration. FLOW-019 and FLOW-053 expose the approved graph, but Current first parses the live queue, preventing callers from retrieving approved metadata after malformed working edits. Provide a read-only claim query without weakening Current, Verify or Done validation.

## Interface

Add ApprovedSnapshot to the retained model with Snapshot project.Snapshot, Slug string and Path string, and expose its alias in api.go. Add Approved(root string) (ApprovedSnapshot, bool, error): bool reports an owned claim, independently of the live queue. Place any new eval-step stub in a new disposable file; no behavior in the facade.

## Added

### new:approved-query
Approved MUST return the calling worktree's held approved snapshot and canonical change slug and repository-relative queued path, with found true, even when working configuration, module specs or change documents are malformed, changed or missing. It MUST use the canonical worktree root in the returned snapshot.

### new:approved-absence
Approved MUST return found false without an error when the calling worktree owns no claim, including when another worktree owns a claim. It MUST return an error whose contracts.ErrorCode is P02 for unreadable or malformed claim data rather than report an approved snapshot.


## Modified

### FLOW-019
When its queued change can be resolved, Current MUST return the project snapshot approved at Start while the calling worktree holds the claim. Approved MUST expose that same snapshot independently of live project and queue parsing.


## Removed

## Decisions

## Notes

Keep existing Current semantics, author/runtime checks and mutations unchanged. The new query reads claims without hooks, agents or writes. It must work with the existing claim records written by the pinned controller, not only a new serialization format: existing records contain worktree, change slug, phase base and snapshot; the queued path can be recovered from that recorded Git tree without parsing live documents. Avoid a claim-format migration. Use public Start plus malformed/missing live inputs and linked-worktree fixtures. Retain existing FLOW-019 assertions and add new coverage for independent access.
Owner authorized these five regeneration prerequisites and their necessary module-local changes. Stop after this queue; no regeneration or drill is authorized. Follow SYSTEM.md and the controller context. Workers must not read docs/rebuild/PROGRESS.md or ACCEPTANCE.md. Preserve unrelated assertions and retained interfaces after evals. Use existing public-interface fixtures and focused assertions; do not redesign helpers or add a test matrix. Evals and implementation workers run fireweed context for this change before each step and fireweed verify for verification. Restricted auditors follow only the supplied audit playbook and inputs. Only the controller advances phases and commits.
