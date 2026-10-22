---
module: workflow
phase: impl
run: 1715897da7ee515af06978e7
---

## Motivation

Enforce the designed phase-specific file permissions before running hooks or accepting a worker step. A rename changes its source and destination; checking only the destination can erase retained tests or interfaces. Keep this pure file-policy slice separate from later live Git/claim/symlink verification.

## Interface

Add CheckEdits(snapshot project.Snapshot, approvedModule contracts.Module, approvedChange contracts.Change, edits []gitstore.Edit) []checks.Diagnostic to the public workflow facade, forwarding to a new disposable implementation. The snapshot supplies validated module ownership and global protected/worker_writable patterns; approvedModule is the selected resolved module with the change's projected set zones already applied, replacing its stale snapshot zones. approvedChange supplies the immutable started phase and exact change path. Inputs are supplied values; this function neither discovers Git edits nor reads files. Preserve the existing CheckState API.

## Added

### FLOW-004
CheckEdits MUST use the supplied approved module's projected zones in preference to the snapshot's stale selected-module zones, and repository-relative ownership from every other snapshot module. It MUST reject protected paths, every basename SPEC.md, other modules' paths and every changes/ path except a modification of the exact current change with P2-01. worker_writable MUST permit only otherwise unowned paths and MUST NOT override protected, spec, other-module or other-change restrictions. Exact current-change modifications MUST be allowed for all worker phases; CheckState separately enforces its content. Paths outside modules and worker_writable MUST be rejected with P2-01. Prefix-similar module/change paths MUST NOT be mistaken for owned or current paths.

### FLOW-005
For an approved evals phase, CheckEdits MUST permit target evals, other module files such as manifests, permitted worker_writable paths, and new disposable files only. Interface edits MUST require a nonempty exact Interface section in the approved change, otherwise returning P1-01. Preserved and append-only additions or edits, and modification, deletion, type change or rename of disposable files MUST return P1-01. Both source and destination of every rename MUST be checked against their own zones and owners, so moving protected, other-module, eval, interface, preserved, append-only or existing disposable files cannot evade their phase restrictions. A copy's source is unchanged: it MUST have a valid path but MUST NOT be treated as an edit; its destination MUST be checked as a new addition.

### FLOW-006
For an approved impl phase, CheckEdits MUST permit target disposable, preserved and other module files and permitted worker_writable paths, but reject every target eval or interface addition, modification, deletion, type change or rename with P2-01. For audit, it MUST reject every edit except modification of the exact current change with P2-01. Independently of phase and zone precedence, every modified, deleted, type-changed or renamed endpoint matching any owning module's append_only glob MUST return H03, including overlaps with stricter zones and edits to other modules. Append-only destinations of additions or copies MAY be accepted in impl when other restrictions permit; evals MUST reject them with P1-01. All independent H03 and phase/ownership violations MUST remain visible at their own paths.

### FLOW-007
CheckEdits MUST reject a missing run, mismatched approved target or approved phase outside evals/audit/impl with P02 at the current change path and corrective guidance, stopping file-policy comparison. It MUST reject malformed edits with P2-01: only A/M/D/T/R/C statuses are accepted; all paths MUST be canonical slash-separated repository-relative file paths without empty, dot, parent, absolute, backslash or NUL components; R/C MUST have a distinct valid OldPath and other statuses MUST NOT have OldPath. It MUST return fresh non-null diagnostics sorted by Path, Code, Severity and Message, deduplicating identical diagnostics, with severity error and nonempty corrective messages. Inputs, nested maps/slices and earlier results MUST remain unchanged, including immediately after result mutation and on rejected calls. Successful and rejected calls MUST remain noninteractive outside a repository and MUST NOT read or mutate files/refs, run Git, hooks, agents, the bootstrap or network commands, or change parent cwd/environment.

## Decisions

This pure policy consumes the approved projected target, rather than deriving zones from worker-edited change metadata. Lifecycle composition will obtain that target through Project/ApplyChange before calling it. Snapshot ownership is already statically validated; stale target zones are replaced, while foreign module ownership/append-only rules remain effective. This API does not validate live claims, HEAD, symlink containment or actual edit collection; those remain a separate checked worktree operation. CheckState owns current-change body/metadata restrictions.

A rename's original R status applies to both endpoints. Evals cannot rename an existing disposable file into or out of a stub path, and implementation cannot rename a retained eval/interface away. Copies leave their source unchanged and check destination additions. Append-only H03 is independently visible even when stricter worker-zone rules reject the same endpoint. Evals keeps append-only read-only, including additions; implementation permits only additions/copy destinations. Every basename SPEC.md is controller-owned, including nested/outside-module paths. Current-change exemption permits only ordinary M, never deletion, type change, copy or rename.

## Rejected alternatives

Checking only rename destinations would permit erasing frozen evals or protected files. Matching worker_writable before module/protected ownership would broaden worker authority. Reading current filesystem contents here would mix policy with later Git/claim and symlink checks and make pure assertions harder to isolate.

## Provenance

- evals | 2026-10-22T19:13:12Z | agent: manual | fireweed: 0.4.0 | playbook: 5867363302d4c874074520977b1627d26991d95ee2fbacb9ff63e827661edd0b | tests: 662 | failing: 434 | passing: 228
- audit | 2026-10-22T19:16:46Z | agent: manual | fireweed: 0.4.0 | playbook: a9c982d24b476621316049b527e6c7bbc319f06b12268af150fa7d8c6df92230 | tests: 0 | failing: 0 | passing: 0
- evals | 2026-10-22T19:18:20Z | agent: manual | fireweed: 0.4.0 | playbook: 5867363302d4c874074520977b1627d26991d95ee2fbacb9ff63e827661edd0b | tests: 670 | failing: 441 | passing: 229
- audit | 2026-10-22T19:19:04Z | agent: manual | fireweed: 0.4.0 | playbook: a9c982d24b476621316049b527e6c7bbc319f06b12268af150fa7d8c6df92230 | tests: 0 | failing: 0 | passing: 0

## Audit

No findings. Checked FLOW-004, FLOW-005, FLOW-006, FLOW-007.
