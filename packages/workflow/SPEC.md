---
depends_on:
    - contracts
    - gitstore
    - assets
    - project
    - checks
id_prefix: FLOW
name: workflow
---

# workflow

## Specs

### FLOW-001
CheckState MUST parse approved and current text with the public change grammar, using the supplied module's ID prefix and the supplied change path. An invalid approved document, an approved target different from the supplied module name, a missing run, or an approved phase outside evals/audit/impl MUST return a P02 error at that path with corrective claim guidance. Once approval is valid, invalid current text MUST return P2-01 at that path with corrective worker-edit guidance. Invalid state MUST NOT be accepted as an unchanged change. A valid unchanged started change, including maintenance and regenerate changes, MUST return an empty non-null diagnostic slice.

### FLOW-002
CheckState MUST reject semantic changes to every approved metadata field except blocked with P2-01, including module, nested set/hook/zone values, bases, fixes, already_satisfied, regenerate, run and phase, and adding or removing metadata keys. Valid blocked reasons MAY be added, replaced or removed. For an approved evals or impl phase it MUST permit adding, changing, moving or removing Decisions, Questions and Notes sections; for an approved audit phase it MUST permit only Audit and Questions sections. Every other parsed section, its body and relative order, and the parsed preamble MUST remain unchanged. Protected edits MUST yield P2-01, including Motivation, Interface, Added/Modified/Removed, Removed decisions, Rejected alternatives, Provenance, unknown sections and changes hidden after duplicate or fenced headings. Permissions MUST derive from the approved phase even when current metadata attempts to switch phase. Metadata ordering/quoting and line-ending changes with the same public parsed values MUST NOT be rejected. Exact case-sensitive section names and public fenced-Markdown parsing MUST determine permissions rather than substring replacement.

### FLOW-003
CheckState MUST return fresh non-null diagnostics sorted deterministically by Path, Code, Severity and Message, with severity error, the supplied path and nonempty corrective messages for every rejection. Metadata and protected-body violations in the same valid current document MUST both remain visible; invalid approved or current grammar MUST stop that comparison rather than inventing partial acceptance. Results, including writable empty results, MUST remain independent across repeated calls, and supplied module fields/maps/slices and input strings MUST remain unchanged. It MUST work from supplied text outside any repository, remain noninteractive and MUST NOT create or edit files/refs, execute Git, hooks, agents, the bootstrap or network commands on successful or rejected calls.

### FLOW-004
CheckEdits MUST use the supplied approved module's projected zones in preference to the snapshot's stale selected-module zones, and repository-relative ownership from every other snapshot module. It MUST reject protected paths, every basename SPEC.md, other modules' paths and every changes/ path except a modification of the exact current change with P2-01. worker_writable MUST permit only otherwise unowned paths and MUST NOT override protected, spec, other-module or other-change restrictions. Exact current-change modifications MUST be allowed for all worker phases; CheckState separately enforces its content. Paths outside modules and worker_writable MUST be rejected with P2-01. Prefix-similar module/change paths MUST NOT be mistaken for owned or current paths.

### FLOW-005
For an approved evals phase, CheckEdits MUST permit target evals, other module files such as manifests, permitted worker_writable paths, and new disposable files only. Interface edits MUST require a nonempty exact Interface section in the approved change, otherwise returning P1-01. Preserved and append-only additions or edits, and modification, deletion, type change or rename of disposable files MUST return P1-01. Both source and destination of every rename MUST be checked against their own zones and owners, so moving protected, other-module, eval, interface, preserved, append-only or existing disposable files cannot evade their phase restrictions. A copy's source is unchanged: it MUST have a valid path but MUST NOT be treated as an edit; its destination MUST be checked as a new addition.

### FLOW-006
For an approved impl phase, CheckEdits MUST permit target disposable, preserved and other module files and permitted worker_writable paths, but reject every target eval or interface addition, modification, deletion, type change or rename with P2-01. For audit, it MUST reject every edit except modification of the exact current change with P2-01. Independently of phase and zone precedence, every modified, deleted, type-changed or renamed endpoint matching any owning module's append_only glob MUST return H03, including overlaps with stricter zones and edits to other modules. Append-only destinations of additions or copies MAY be accepted in impl when other restrictions permit; evals MUST reject them with P1-01. All independent H03 and phase/ownership violations MUST remain visible at their own paths.

### FLOW-007
CheckEdits MUST reject a missing run, mismatched approved target or approved phase outside evals/audit/impl with P02 at the current change path and corrective guidance, stopping file-policy comparison. It MUST reject malformed edits with P2-01: only A/M/D/T/R/C statuses are accepted; all paths MUST be canonical slash-separated repository-relative file paths without empty, dot, parent, absolute, backslash or NUL components; R/C MUST have a distinct valid OldPath and other statuses MUST NOT have OldPath. It MUST return fresh non-null diagnostics sorted by Path, Code, Severity and Message, deduplicating identical diagnostics, with severity error and nonempty corrective messages. Inputs, nested maps/slices and earlier results MUST remain unchanged, including immediately after result mutation and on rejected calls. Successful and rejected calls MUST remain noninteractive outside a repository and MUST NOT read or mutate files/refs, run Git, hooks, agents, the bootstrap or network commands, or change parent cwd/environment.

## Decisions

[approved-state]
This is a text-state guard, not full lifecycle verification: filesystem zones, symlinks, HEAD/claim ownership, questions/block transitions, fail-first, audit completeness and hook execution remain separate composition. Input module identity is the resolved target from the approved snapshot. Public ParseChange owns grammar and normalized document values; the guard owns permission comparison. It compares parsed metadata keys and nested values, so adding an explicit empty immutable key is still a metadata edit. blocked alone is mutable metadata. Approved started state must have a run and a recognized worker phase; applied/queued state is not a live step baseline.

The approved phase selects section permissions, so an auditor cannot grant itself implementation permissions by changing phase. Non-audit workers may edit Notes as well as Decisions/Questions, matching the reference's documented free-prose workflow. All occurrences of an allowed section are mutable only when public grammar accepts them; every retained protected section occurrence remains compared. Unknown sections and the preamble remain protected. Parser-normalized whitespace and YAML syntax are not independent approval fields. A metadata violation and a protected-body violation produce separate diagnostics, in stable order; invalid text produces one parse/baseline diagnostic first.

The initial incomplete cycle was fully reverted before this restart (D017). Temporary cwd fixtures must resolve filesystem aliases before comparing with os.Getwd; establish that premise independently. The public guard must preserve actual parent state, rather than manipulating cwd or environment to satisfy a raw alias string. Duplicate-heading invalid-input cases use raw text because the document renderer correctly rejects duplicate sections before the state guard.

[phase-permissions]
This pure policy consumes the approved projected target, rather than deriving zones from worker-edited change metadata. Lifecycle composition will obtain that target through Project/ApplyChange before calling it. Snapshot ownership is already statically validated; stale target zones are replaced, while foreign module ownership/append-only rules remain effective. This API does not validate live claims, HEAD, symlink containment or actual edit collection; those remain a separate checked worktree operation. CheckState owns current-change body/metadata restrictions.

A rename's original R status applies to both endpoints. Evals cannot rename an existing disposable file into or out of a stub path, and implementation cannot rename a retained eval/interface away. Copies leave their source unchanged and check destination additions. Append-only H03 is independently visible even when stricter worker-zone rules reject the same endpoint. Evals keeps append-only read-only, including additions; implementation permits only additions/copy destinations. Every basename SPEC.md is controller-owned, including nested/outside-module paths. Current-change exemption permits only ordinary M, never deletion, type change, copy or rename.
