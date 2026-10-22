---
module: workflow
phase: evals
run: 0f9f3212efebe4892f28a2e7
set:
    interface:
        - api.go
        - model/*.go
---

## Motivation

Provide real atomic local claims shared by linked worktrees. Lifecycle start, verification, recovery and release need an approved baseline and a stable ownership token, without changing branches or committing worker work.

## Interface

Add retained workflow/model Claim data and a public workflow Claim alias. Claim has Worktree, Change, Run, Base and State strings serialized under lowercase JSON keys; Object is the Git blob object identifier used as a compare-and-swap token and is excluded from stored JSON. Add CreateClaim(start, slug, run, base, state string) (Claim, error), LoadClaim(start, slug string) (Claim, error), Claims(start string) ([]Claim, error), UpdateClaim(start string, previous Claim, base, state string) (Claim, error), and ReleaseClaim(start, slug, expectedObject string) error. Thin facade functions forward to a new disposable claims implementation. All operations accept nested working directories and use the canonical repository root and shared common Git directory. Preserve existing APIs.

## Added

### FLOW-008
CreateClaim MUST store a JSON blob at refs/fireweed/claims/<slug> using an atomic create-only Git ref update, with the canonical current worktree root, exact slug/run/state and a full resolved local commit ID for Base. It MUST return that stored Claim and its full blob Object token, without committing or moving HEAD. Only nonempty slugs matching lowercase letters/digits followed by lowercase letters/digits/hyphens and a nonblank run are valid; invalid slugs, missing runs, unavailable repositories and unresolvable commit bases MUST fail with P02 and corrective guidance and create no claim. Empty State MUST be allowed as an initial reservation. Competing creates for the same slug, including from linked worktrees, MUST have exactly one winner; a conflict MUST preserve the winner and report its ownership/release guidance. Different slugs MUST remain independent.

### FLOW-009
UpdateClaim MUST require the exact current Claim, including its Object token, held by the calling canonical worktree; it MUST preserve Worktree, Change and Run, replace State exactly, resolve the new Base to a full local commit ID and atomically replace only the expected blob ref. A stale, modified or foreign-worktree previous Claim MUST fail with P02 and corrective guidance without overwriting current ownership. ReleaseClaim MUST atomically delete only the named ref matching the supplied full Object token, and MUST reject missing/stale/invalid tokens or missing claims with P02 and corrective guidance. Explicit release MAY be performed from another linked worktree; it MUST NOT delete a replacement claim. Updates and releases racing on one token MUST NOT both succeed, except that a genuinely unchanged update MAY retain the same content-addressed token.

### FLOW-010
LoadClaim MUST read the exact shared claim ref, return its complete independent stored fields and full Object token, and reject missing refs, non-blob objects, malformed JSON, trailing JSON values, missing or incorrectly typed required fields, invalid slug/run/worktree/base values and stored Change differing from the requested slug with P02 and corrective guidance. Worktree MUST be an absolute clean path, but its directory MAY no longer exist so stale claims remain inspectable. Claims MUST list every shared claim in ascending Change order, excluding unrelated refs; an empty inventory MUST be non-null. It MUST reject malformed claim refs or payloads atomically rather than silently skipping them or returning a partial inventory. Failed Claim-returning calls MUST return a zero Claim, and separate returned inventories/claims MUST remain independent.

### FLOW-011
Every claim operation MUST remain local and noninteractive, use literal Git arguments, and MUST NOT run fireweed hooks, native Git ref-transaction hooks, agents, the bootstrap or network commands. Successes and rejections MUST preserve HEAD, branches, the caller's index, working files, ignored/user files, other refs, input Claim values, parent cwd and environment; only the selected claim ref and necessary Git object storage MAY change. Temporary owned storage MUST be cleaned on success and failure, including duplicate-create and stale-update failures. Reads and failed precondition checks MUST NOT change refs. Claims MUST work across linked worktrees and in SHA-1 and SHA-256 repositories supported by installed Git, without a hard-coded object-ID width. Claim operations MUST NOT fetch, push, create branches or require author identity, authenticated agents or network access.

## Decisions

This is claim storage, not queue eligibility, one-change-per-worktree policy, approval validation or phase transitions. Start will reserve with empty State before assigning IDs and then update its approved text. Update can move Base after a controller-owned phase commit; a later live-worktree guard checks HEAD against that baseline. Base input may name any locally resolvable commit and is normalized to its full object ID. State is opaque text here; CheckState owns started-document grammar and worker permissions.

Stored JSON consists of the five required string fields. Additional JSON fields are tolerated for forward-compatible reference data; unknown fields do not replace required validation. Empty State is valid. Run is preserved byte-for-byte when nonblank. Load validates Base names as full stored commit IDs and checks local existence; it does not require the stale owner's directory to exist. Git's native reference-transaction hooks are disabled for these metadata operations, so claiming cannot run unrelated project hooks. No fireweed enforcement hook is changed or bypassed in this rebuild.

Update compares the supplied previous record with the actual stored record before a compare-and-swap, then Git enforces the final atomic token check. A content-identical update naturally retains its object ID. Explicit ReleaseClaim is an ownership-token operation rather than an owner-only action, matching manual release intent while preventing a stale release from deleting a newer claim. Unreachable blob objects left by a lost ref race are normal Git object storage, not leaked temporary files. Git errors carry P02 plus a corrective action; callers never receive partial successful data.

## Rejected alternatives

A lock file in one worktree would not arbitrate linked worktrees. Unconditional updates/deletions could overwrite or release a newer owner after a race. Storing claims as commits or branch tips would require author configuration or alter implementation history. Copying an entire controller implementation is unnecessary; use the public local Git/process primitives and new code.

## Provenance

- evals | 2026-10-22T19:36:09Z | agent: manual | fireweed: 0.4.0 | playbook: 5867363302d4c874074520977b1627d26991d95ee2fbacb9ff63e827661edd0b | tests: 769 | failing: 94 | passing: 675
- audit | 2026-10-22T19:40:43Z | agent: manual | fireweed: 0.4.0 | playbook: a9c982d24b476621316049b527e6c7bbc319f06b12268af150fa7d8c6df92230 | tests: 0 | failing: 0 | passing: 0
- evals | 2026-10-22T19:45:07Z | agent: manual | fireweed: 0.4.0 | playbook: 5867363302d4c874074520977b1627d26991d95ee2fbacb9ff63e827661edd0b | tests: 796 | failing: 120 | passing: 676
- audit | 2026-10-22T19:48:32Z | agent: manual | fireweed: 0.4.0 | playbook: a9c982d24b476621316049b527e6c7bbc319f06b12268af150fa7d8c6df92230 | tests: 0 | failing: 0 | passing: 0

## Audit

- FLOW-008: "using an atomic create-only Git ref update" — missing rejection control for an unconditional text-stdin update: claimCASProtocol accepts `update <ref> <newOID>\n` as create-only by treating the trailing empty split chunk as an explicit empty old-object field.
- FLOW-009: "ReleaseClaim MUST atomically delete only the named ref matching the supplied full Object token" — missing assertion that releasing a symbolic claim ref cannot delete its unrelated target ref; every existing release fixture uses a direct ref, so default Git dereferencing could violate this restriction unnoticed.
- FLOW-011: "Temporary owned storage MUST be cleaned on success and failure, including duplicate-create and stale-update failures." — missing directory-type assertion for newly added two-hex-character entries directly under objects; the cleanup checker currently accepts a regular file such as objects/ab as if it were a legitimate fanout directory.
