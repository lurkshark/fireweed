---
already_satisfied:
    - PRJ-001
    - PRJ-009
    - PRJ-010
    - PRJ-014
    - PRJ-017
    - PRJ-020
    - PRJ-023
    - PRJ-024
    - PRJ-026
    - PRJ-027
bases:
    PRJ-001: b5acc033
    PRJ-009: ca6497c1
    PRJ-010: 4911c01c
    PRJ-014: 7f48193a
    PRJ-017: 5445ca47
    PRJ-020: 9c4f12e8
    PRJ-023: fe275028
    PRJ-024: ab223e2b
    PRJ-026: c0e478ad
    PRJ-027: 4d3af2c5
module: project
---

## Motivation

The 2026-10-26 project review found effort going into hardening beyond the design while core commands stayed unbuilt. The revised design (Spec shape; SYSTEM.md is a plain document) and the SYSTEM.md scope guardrails say that locality and safety properties are conventions checked by review, that specs state observable behavior without aliasing or purity requirements, and that the supported environment is macOS and Linux with ordinary full SHA-1 clones. See D048.

## Interface

No public API change.

## Added

## Modified

### PRJ-001
Load MUST discover the canonical Git root from a nested directory, parse fireweed.yaml using the running embedded version, resolve the local base ref, and discover SPEC.md only in module directories matching configured modules globs. It MUST retain resolved module defaults, explicit empty overrides, dependencies, hooks, specs, Decisions and original documents. Modules MUST be sorted by name then directory. It MUST NOT follow directory symlinks or read files outside the repository. Missing/invalid configuration or invalid configured globs MUST return an error mentioning C01; a nonrepository MUST return an error.

### PRJ-009
Project MUST apply changes in numeric Number then Slug order across their target modules, independent of input slice order. A nonempty before selector MUST return the state after preceding changes only; empty before MUST apply the entire queue. Same-number changes MUST follow slug order. Unknown/ambiguous selectors, unknown target modules and invalid application MUST return errors. Current archive entries MUST NOT be replayed as pending changes.

### PRJ-010
BaseHashes MUST return a non-null map containing exactly the Modified and Removed IDs of the selected queued change, hashing the target specs as projected immediately before that change with the public normalized eight-character SpecHash, including tags; existing bases MUST be replaced in the returned map. Added-only or maintenance changes MUST return an empty map. Missing/ambiguous change, missing target module or missing edited spec MUST return an error.

### PRJ-014
CheckQueue MUST return a non-null diagnostic slice sorted by path, code, severity and message, with nonempty corrective messages and error/warning severity. An empty queue MUST return an empty slice regardless of unrelated input discovery/static diagnostics. A valid queue MUST return an empty slice.

### PRJ-017
Allocate MUST return errors for missing/ambiguous selectors or targets, invalid selected typed change grammar and undeclared/cross-only handle references, without returning a partial change or allocation mapping. It MUST validate only the selected change's grammar, without replaying or requiring validity of excluded queue entries. Successful empty mappings MUST be non-null maps, and repeated calls with the same input MUST be deterministic.

### PRJ-020
Finalize MUST return no partial module on error. ApplyChange MUST remain a preview that retains existing Decisions without performing final decision additions or removals.

### PRJ-023
ResolvePlaybook MUST reject a selected override path that escapes the repository, or is unreadable or a directory, with a corrective error mentioning D05 and the selected repository-relative override path, never silently falling back. A missing selected override MUST fall back to embedded text. Contained regular-file symlink targets MAY be used. Resolution MUST NOT write files or refs or run hooks or agents.

### PRJ-024
LoadAt MUST discover the canonical calling checkout from start, including nested directories and linked worktrees, and resolve one explicit nonblank local revision to a full commit before reading its tree. It MUST accept full commit IDs, local branches, annotated/lightweight tags and commit expressions supported by installed Git. An absent/invalid/option-like revision, blob/tree object, unborn HEAD, unavailable repository or Git resolution/enumeration failure MUST return a corrective C01 error and zero Snapshot. All subsequent tree/blob reads MUST be pinned to that resolved commit or its object IDs even if the input ref moves. No current file, index, HEAD or newer commit may replace selected metadata.

### PRJ-026
LoadAt MUST NOT resolve committed metadata symlinks to content outside the selected tree or to current working files. Module discovery MUST NOT traverse directory symlinks, gitlinks or nested .git/.fireweed subtrees.

### PRJ-027
LoadAt MUST read the selected commit's objects, never current working files, so dirty, deleted, untracked or staged current metadata MUST NOT affect the selected snapshot. Failures MUST return no partial Snapshot. It MUST NOT write files or refs, create a temporary checkout or worktree, or run hooks or agents.

## Removed

## Decisions

Narrowed specs are listed in already_satisfied because the current implementation already meets the narrower text; nothing here adds behavior. Clauses removed from a spec become unspecified, not forbidden: existing code that still does more is conforming, and the implementation step may simplify it.

Symlink support inside committed metadata is no longer specified beyond containment. The implementation may keep its current virtual-tree resolution or replace it with rejection of metadata symlinks.

## Rejected alternatives

Keeping the clauses and only deleting their evals: coverage requires every enforced spec to stay covered, and the next audit would reinstate the tests. Rewriting every over-long spec to the new shape limits at once: most long specs carry no hardening and rewriting them adds risk without benefit; they are left for later changes.

## Notes

Evals step: delete baseline_isolation_test.go, baseline_promisor_test.go, isolation_test.go, playbooks_isolation_test.go and git_operations_test.go. Move any shared helper still needed into a plain test file. In the remaining files remove subtests that assert only removed clauses: aliasing and spare backing storage, purity and no-I/O checks, Git proxy or operation classification, promisor or missing objects, SHA-256 fixtures, chained or contained committed symlinks, index-byte, temp-storage and cwd/environment preservation. Every Modified ID must keep at least one referencing test. Implementation step: no product change is required; simplifying committed-symlink handling within PRJ-026 is allowed.
