---
module: workflow
phase: impl
run: 87cb84b4509b66f7b60d7a0e
---

## Motivation

Compose the approved text/file policies with actual local Git ownership, a stable HEAD and safe current paths. Verification must reject stale or foreign claims and escaping symlinks before any eval, boundary hook or agent can run.

## Interface

Add retained workflow/model WorktreeCheck with Edits []gitstore.Edit and Diagnostics []checks.Diagnostic, JSON keys edits and diagnostics, and a public workflow WorktreeCheck alias. Add CheckWorktree(snapshot project.Snapshot, approvedModule contracts.Module, changePath string, claim Claim) WorktreeCheck, forwarding to a new disposable implementation. snapshot supplies validated global config and module ownership; approvedModule is the selected resolved target with approved projected zones, as for CheckEdits. snapshot.Root locates the live checkout (nested paths are accepted). This operation composes claim storage, CheckState, CheckEdits and public Git path/diff primitives, not full lifecycle coverage/audit/application. Preserve existing APIs.

## Added

### FLOW-012
CheckWorktree MUST discover the canonical checkout from snapshot.Root, require the exact current stored Claim including Object and fields, require that claim's owner to be this checkout, and require HEAD to equal Claim.Base. It MUST validate Claim.State as an unchanged valid started document through CheckState and public change grammar, with parsed Module, Run and Slug agreeing with the supplied module and Claim. changePath MUST name that queued change directly under changes/, never an archive or escaping/invalid path. Unavailable repositories, missing/stale/tampered/foreign claims, moved HEAD or invalid approved state/identity/path MUST return a corrective P02 at changePath, empty non-null Edits, and stop before current-file comparison or edit collection. Linked worktrees and SHA-1/SHA-256 claim records MUST be supported without changing HEAD or ownership.

### FLOW-013
After valid preconditions, CheckWorktree MUST read the exact current change only after checking its path's containment with the public SafePath contract. Missing, unreadable, directory or escaping/dangling-symlink current changes MUST produce corrective P2-01 at changePath; escaping file contents MUST NOT be read even when read errors would be ignored. Safe in-repository links MAY be read, but their actual type-change edits MUST still obey CheckEdits. Valid current text MUST be compared with Claim.State through CheckState, using the approved phase rather than worker-modified metadata. Invalid current grammar, immutable metadata or protected body edits MUST retain those exact public state diagnostics. Allowed worker prose and blocked edits MUST remain permitted by this guard; blocking/Questions policy is separate.

### FLOW-014
With valid preconditions, CheckWorktree MUST collect the public Git Edits since the exact Claim.Base, including staged, unstaged and untracked files and preserving rename source/destination records. Collection failures MUST return corrective P02 at changePath. It MUST check containment of every edit destination and every rename/copy source, recording corrective P2-01 at each unsafe endpoint, including nonexisting descendants of escaping directory symlinks. It MUST compose CheckEdits using the supplied approved projected target and validated snapshot, retaining independent state, safety and phase/ownership/append-only violations instead of accepting edits after one check fails. Normal permitted changes MUST pass; protected/global/spec/foreign/other-change edits, impl eval/interface renames and evals existing disposable edits MUST fail through the established public policy. No coverage, boundary hook or agent MUST run in this operation.

### FLOW-015
CheckWorktree MUST return fresh non-null Edits and Diagnostics with diagnostics deduplicated and sorted by Path, Code, Severity and Message. Errors MUST have severity error, their specified path and nonempty corrective guidance. Repeated results, including empty results, and input snapshots/modules/claims/maps/slices MUST remain independent and unchanged immediately after result mutation, on successes and rejections. The operation MUST remain local and noninteractive, preserve HEAD, branches, claim/other refs, index bytes, tracked/untracked/ignored/user files, private outside files and parent cwd/environment, and MUST NOT commit, reset, stage, fetch, push, create branches, execute hooks/agents/bootstrap/network commands or leave owned temporary storage. It MAY perform necessary local read-only Git/filesystem inspection. Supplied snapshots' hook/agent strings MUST NOT be executed even on success or failure.

## Decisions

The caller owns static project validation and selection/projection; this guard never trusts a worker-edited phase or set for permissions. Protected config/SPEC changes remain rejected by the global file policy even if a caller reloads a current snapshot. Only a passing aggregate guard can lead to later hook execution. CheckState owns normalized intent, CheckEdits owns phase zones, claim storage owns exact owner/token records and Git primitives own edit/path semantics. This change does not start, advance, block or apply a change and never updates a claim.

Precondition failures return one P02 before collecting edits or reading the current change. Once preconditions pass, edit collection and safety/file policy remain useful even when the current change is missing or invalid, so diagnostics compose; a collection failure is P02 and cannot imply a clean tree. Endpoint safety rejects attempted operations on escaping paths without following outside file contents. Deleting a missing source remains governed by the existing SafePath contract and phase policy. Results expose the actual Git edit records without synthesizing a different rename representation. Error messages from public CheckState and CheckEdits remain intact.

## Rejected alternatives

Trusting only caller-supplied Claim fields would allow stale ownership or a moved HEAD to be accepted. Reading current text before path containment could disclose private files through symlinks. Running evals first could execute hooks for a forbidden step. Reconstructing approval from current worker metadata would allow a phase switch or zone change to grant permissions.

## Provenance

- evals | 2026-10-22T20:14:08Z | agent: manual | fireweed: 0.4.0 | playbook: 5867363302d4c874074520977b1627d26991d95ee2fbacb9ff63e827661edd0b | tests: 915 | failing: 93 | passing: 822
- audit | 2026-10-22T20:22:24Z | agent: manual | fireweed: 0.4.0 | playbook: a9c982d24b476621316049b527e6c7bbc319f06b12268af150fa7d8c6df92230 | tests: 0 | failing: 0 | passing: 0
- evals | 2026-10-22T20:27:01Z | agent: manual | fireweed: 0.4.0 | playbook: 5867363302d4c874074520977b1627d26991d95ee2fbacb9ff63e827661edd0b | tests: 924 | failing: 102 | passing: 822
- audit | 2026-10-22T20:28:41Z | agent: manual | fireweed: 0.4.0 | playbook: a9c982d24b476621316049b527e6c7bbc319f06b12268af150fa7d8c6df92230 | tests: 0 | failing: 0 | passing: 0

## Audit

No findings. Checked FLOW-012, FLOW-013, FLOW-014, FLOW-015.
