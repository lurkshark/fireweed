---
already_satisfied:
    - AST-002
    - AST-003
bases:
    AST-002: 9a8f1a37
    AST-003: 69d8655a
module: assets
---

## Motivation

The 2026-10-26 project review found effort going into hardening beyond the design while core commands stayed unbuilt. The revised design (Spec shape; SYSTEM.md is a plain document) and the SYSTEM.md scope guardrails say that locality and safety properties are conventions checked by review, that specs state observable behavior without aliasing or purity requirements, and that the supported environment is macOS and Linux with ordinary full SHA-1 clones. See D048.

## Interface

No public API change.

## Added

## Modified

### AST-002
The embedded playbook inventory MUST be exactly author-change, audit, evals, impl and regenerate in sorted order. Each lookup MUST return its name, nonempty text and the full lowercase SHA-256 of that exact text; unknown names MUST return errors. Text MUST give concrete commands and rules for its step: authoring a single-module change with local handles, complete specs, stamping and human base approval; fail-first evals with genuine testcase failures and allowed zones; independent exact-clause audits and No findings plus checked IDs; implementation with immutable evals/interfaces and passing dependents/boundaries; regeneration from retained contracts without an old-implementation oracle, preserving interfaces/evals and reporting targeted gaps.

### AST-003
The diagnostic catalog MUST contain U01, C01 through C16, P01 through P06, P1-01 through P1-03, P2-01, D01 through D07 and H01 through H04, sorted and without duplicates. Each explanation MUST identify the problem and a corrective action. H01 MUST cover protected worker files, H02 immutable impl evals/interfaces, H03 append-only changes by anyone including merge resolutions, and H04 immutable archives with deletion for revert allowed plus duplicate recovery by aborting the merge, listing commits, reverting the incoming duplicate and merging again. D05 MUST explain both unsafe or unreadable playbook overrides and stale based_on metadata. Its corrective guidance MUST tell the user to restore a readable contained override and, for staleness, compare the embedded hash, review instructions and update based_on. Unknown codes MUST return errors.

## Removed

## Decisions

Narrowed specs are listed in already_satisfied because the current implementation already meets the narrower text; nothing here adds behavior. Clauses removed from a spec become unspecified, not forbidden: existing code that still does more is conforming, and the implementation step may simplify it.

The revised design's new codes (C17, C18, P07, P08, B01 to B04) and playbook rules are product work for R22 to R25, authored separately; this change only removes hardening.

## Rejected alternatives

Keeping the clauses and only deleting their evals: coverage requires every enforced spec to stay covered, and the next audit would reinstate the tests. Rewriting every over-long spec to the new shape limits at once: most long specs carry no hardening and rewriting them adds risk without benefit; they are left for later changes.

## Notes

Evals step: delete isolation_test.go (the cgo runtime-isolation probe). Remove subtests that only assert independent returned slices. Every Modified ID must keep at least one referencing test. Implementation step: no product change is required.
