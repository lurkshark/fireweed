---
already_satisfied:
    - FLOW-011
    - FLOW-013
    - FLOW-015
    - FLOW-019
    - FLOW-020
    - FLOW-023
    - FLOW-036
    - FLOW-040
    - FLOW-047
bases:
    FLOW-011: 33e0de6a
    FLOW-013: 2b421e6a
    FLOW-015: 9928733a
    FLOW-019: 12446e0a
    FLOW-020: b17faf30
    FLOW-023: ad1f08c8
    FLOW-027: 71ba1b1b
    FLOW-031: d7ff8e3a
    FLOW-036: ee36115c
    FLOW-040: 0dbeb565
    FLOW-047: a1dee8ea
module: workflow
---

## Motivation

The 2026-10-26 project review found effort going into hardening beyond the design while core commands stayed unbuilt. The revised design (Spec shape; SYSTEM.md is a plain document) and the SYSTEM.md scope guardrails say that locality and safety properties are conventions checked by review, that specs state observable behavior without aliasing or purity requirements, and that the supported environment is macOS and Linux with ordinary full SHA-1 clones. See D048.

The in-flight held-approval-source change (FLOW-048 to FLOW-051 as assigned) was reverted before this change; its graph-query need (D042) is met later in app with existing APIs.

## Interface

No public API change. Eligible and ApprovalMatches keep their signatures; only their purity specs are removed.

## Added

## Modified

### FLOW-011
Claim operations MUST work across linked worktrees, MUST change only the selected claim ref, and MUST NOT run fireweed hooks or agents, commit, fetch, push or create branches. Reads and failed precondition checks MUST NOT change refs.

### FLOW-013
After valid preconditions, CheckWorktree MUST read the current change file; a missing, unreadable or directory path, or a path escaping the repository, MUST produce corrective P2-01 at changePath. Valid current text MUST be compared with Claim.State through CheckState, using the approved phase rather than worker-modified metadata. Invalid current grammar, immutable metadata or protected body edits MUST retain those exact public state diagnostics. Allowed worker prose and blocked edits MUST remain permitted by this guard; blocking/Questions policy is separate.

### FLOW-015
CheckWorktree MUST return non-null Edits and Diagnostics with diagnostics deduplicated and sorted by Path, Code, Severity and Message. Errors MUST have severity error, their specified path and nonempty corrective guidance. It MUST NOT commit, stage, reset, or execute hooks or agents.

### FLOW-019
Every selected content read, including SYSTEM.md, SPEC.md, the queued change, playbook and Files, MUST reject paths escaping the repository. Filesystem/Git enumeration errors MUST fail with corrective C01. Any failed BuildContext MUST return the zero WorkerContext with no partial files or metadata. BuildContext MUST NOT run hooks or agents.

### FLOW-020
LoadApproval MUST accept only an exact valid claim slug and discover the canonical repository from start, including nested and linked checkouts. Invalid selectors MUST fail C05 before repository or content access. It MUST inspect only the exact selected shared claim ref, independently of other claims, and distinguish an absent ref from a malformed payload, symbolic alias or inspection failure. With no selected ref it MUST load the current project through public Load and return its metadata with a zero Claim. With a valid selected claim held by this canonical checkout it MUST load metadata from exactly Claim.Base through public LoadAt and return that exact Claim, including Object. Current config, module SPEC, queued/archive document bytes or moving refs MUST NOT replace this committed metadata.

### FLOW-023
LoadApproval MUST NOT write files or refs, commit, stage or reset, execute hooks or agents, or validate current worker edits.

### FLOW-036
Start MUST return warning diagnostics P04 when HEAD does not contain the locally known base tip and independently when the selected authored content is absent or unapproved on that base. It MUST resolve the chosen base commit locally once, read the exact selected file from that commit, and use ApprovalMatches against the started content so late IDs and canonical layout do not revoke approval. A malformed, different-target or absent base change MUST be unapproved. Unavailable base content MUST never imply approval. Stale or unapproved content MUST warn rather than prevent an otherwise valid manual start. Base warnings MUST retain useful path and corrective guidance and MUST NOT change the configured base or any base ref.

### FLOW-040
Start MUST NOT fetch, push, create branches or worktrees, commit, move HEAD, stage or reset the index, change other claims, or execute hooks or agents. Failure paths MUST leave unrelated files and refs unchanged and report useful failed-command details.

### FLOW-047
Apart from the selected step's configured hooks, Verify MUST NOT run agents, commit, fetch, push or create branches or worktrees, and MUST remove its own temporary storage. Diagnostics MUST be deterministically ordered by Path, Code, Severity and Message with exact duplicates removed, and all Verification collection fields MUST be non-null.

## Removed

### FLOW-027

### FLOW-031

## Decisions

Narrowed specs are listed in already_satisfied because the current implementation already meets the narrower text; nothing here adds behavior. Clauses removed from a spec become unspecified, not forbidden: existing code that still does more is conforming, and the implementation step may simplify it.

The approved-snapshot design (LoadApproval, LoadAt, CheckState, ApprovalMatches) stays: Start, context and Verify depend on it and it implements design behavior. Only its locality, partial-clone, replacement-ref, SHA-256 and aliasing clauses are removed. Native supervision of Git subprocesses is no longer described; gitstore 0055 replaces it with process groups.

## Rejected alternatives

Keeping the clauses and only deleting their evals: coverage requires every enforced spec to stay covered, and the next audit would reinstate the tests. Rewriting every over-long spec to the new shape limits at once: most long specs carry no hardening and rewriting them adds risk without benefit; they are left for later changes.

## Notes

Evals step: delete approval_git_controls_test.go, approval_git_proxy_test.go, approval_isolation_test.go, claim_locality_controls_test.go, claim_locality_test.go, content_approval_isolation_test.go, context_isolation_test.go, eligibility_isolation_test.go, policy_isolation_test.go, read_operation_git_test.go, start_faults_test.go, verify_isolation_test.go, verify_locality_test.go, verify_operations_test.go and worktree_isolation_test.go. Move any shared helper still needed into a plain test file. Delete or rename every test referencing FLOW-027 or FLOW-031. In the remaining files remove subtests that assert only removed clauses (native supervision, fsmonitor, promisor or missing objects, replacement refs, SHA-256 fixtures, index-byte, temp-storage, cwd/environment and aliasing checks). If FLOW-038 rollback loses its only coverage with start_faults_test.go, add a plain fixture-based rollback test without cgo. Every remaining Modified ID must keep at least one referencing test. Implementation step: no product change is required; removing code that only served deleted observers is allowed.
