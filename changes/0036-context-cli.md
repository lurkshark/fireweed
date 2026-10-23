---
module: app
phase: evals
run: 8baf00c555aa2a3caf96dc64
---

## Motivation

Expose complete approved worker context through the independently built executable. Separate invocations must restore claim-owned metadata before reading current worker feedback, so editing a config, SPEC or change cannot redirect the context's permission authority. Existing public LoadApproval and BuildContext provide that behavior; this change composes them without adding lifecycle advancement.

## Interface

Keep app.Run(args []string, out, errOut io.Writer) int and all retained public models unchanged. Add context command dispatch and composition only under app/internal. No new public API or worker-writable zone is needed.

## Added

### APP-019
context <change> MUST discover from the current working directory, including nested and linked checkouts, and compose public LoadApproval with BuildContext for the exact selected slug. With no selected claim it MUST use current discovered source; with a valid owned claim it MUST restore Claim.Base metadata and the complete approved State before assembling current feedback. It MUST return the complete public WorkerContext without dropping or adding fields, including the approved Change, inclusive projected target and preceding projected direct dependencies, exact current ChangeText and SpecText, optional System, phase-selected Files, exact selected Playbook and PlaybookHash, and canonical GitCommonDir. Empty existing change text MUST remain empty, while a missing current change MUST use the builder's matching owned-claim recovery. Claimed context MUST remain available with malformed/deleted current config or queue metadata and worker-edited SPEC metadata when its required feedback files remain readable. It MUST NOT reload current project metadata first, project the selected change twice, allocate IDs, certify static validity, require clean files or HEAD equality, or validate worker edits as an advancement precondition. Successful calls MUST return exit 0 with empty diagnostics even when source metadata contains unrelated static/discovery problems; check and verify remain separate operations.

### APP-020
context MUST preserve useful full public LoadApproval/BuildContext error messages and their leading C01, C05, P02 or D05 code as a single severity error diagnostic, with null data and no invented path field. C05 selection/projection failures and P02 claim/ownership/state failures MUST return exit 1; C01 discovery/required-input failures and D05 override failures MUST return exit 2. Missing or ambiguous exact queued selectors, archived-only selectors, missing targets, malformed/foreign/reserved claims, unavailable baselines and invalid approved state MUST NOT fall back to different metadata or expose partial WorkerContext. Unsafe or unreadable required inputs MUST fail atomically; a missing optional SYSTEM.md remains empty. Invalid slug selectors MUST retain C05 before repository/content access. JSON MUST use exactly one schema_version 1 canonical context/data/non-null diagnostics envelope on stdout and empty stderr. Plain success MUST be pretty JSON of the same complete WorkerContext on stdout; plain failure MUST have empty stdout and the coded corrective diagnostic on stderr. Output write failures MUST return exit 2 on success, operational-error and usage output.

### APP-021
context MUST accept --json before or after the command/operand and valid --help/-h without loading a project. Missing/excess operands, unknown or valued boolean flags, show-only modes and coverage-only --with-change MUST be corrective U01 exit 2 before repository/content access. -- MUST stop option parsing, so later flag-looking text is an operand and cannot enable JSON. All calls MUST remain local and noninteractive and preserve files, ignored/outside contents, index, HEAD, branches, claim/other refs, common Git, temporary storage and parent cwd/environment on success and rejection. The command MUST NOT independently run eval/setup/boundaries hooks, agents, bootstrap or network commands, commit, stage, reset, write files/refs or create claims. It MUST inspect only the exact selected claim rather than require every claim to decode, and MUST NOT open disposable code, unrelated private content, transitive dependency contracts or inputs excluded by the public source/context selection contracts. Escaping paths MUST reject before outside content opens, including ignored-error attempts.

## Modified

## Removed

## Decisions

This is a read/repair command. LoadApproval determines authority; BuildContext owns projection and safe current feedback selection. An unclaimed author/resume view is legitimate, and a claimed view tolerates moved HEAD or invalid worker text without permitting advancement. Native/current static diagnostics are not treated as permission to omit context or as successful validation. Errors preserve the public prefix and entire detail without guessing a path from prose; required paths remain visible in their messages. Selector validity is a C05 source error, while arity/options are U01 usage errors. Unknown exact slugs therefore return C05 rather than U01.

Both plain and JSON success expose the public WorkerContext shape; JSON wraps it in the existing envelope. The builder selects author-change for an unstarted change, evals/audit/impl for a started phase, and regenerate for regeneration. Project override bytes and hashes retain the existing resolver contract. This command does not introduce approval, clean-tree, hook, HEAD or current edit validation, which are start/verify/done responsibilities. Existing public loader/builder safety and phase selection remain required; app tests exercise their real composition through the new executable.

## Rejected alternatives

Loading current project metadata before claim ownership would let worker config choose its own authority. Running check or verify before this read would hide the repair context needed to resolve their failures. Flattening WorkerContext into a new CLI-specific model would drop retained contract data. Treating all operational failures as U01 would lose stable protocol meaning and useful public guidance.

## Provenance

- evals | 2026-10-23T02:31:59Z | agent: manual | fireweed: 0.4.0 | playbook: 5867363302d4c874074520977b1627d26991d95ee2fbacb9ff63e827661edd0b | tests: 216 | failing: 47 | passing: 169
- audit | 2026-10-23T02:42:13Z | agent: manual | fireweed: 0.4.0 | playbook: a9c982d24b476621316049b527e6c7bbc319f06b12268af150fa7d8c6df92230 | tests: 0 | failing: 0 | passing: 0

## Audit

- APP-019: "allocate IDs" — missing assertion: every context fixture uses already allocated Added IDs. No handle-bearing context establishes that unallocated local handles remain in the complete structured context and exact source bytes, with specs, changes and refs unchanged. Allocating only when handles exist would pass the current assertions.
- APP-021: "parent cwd/environment on success and rejection" — missing assertion: the preservation case compares the test parent's cwd/environment around executable subprocesses, which cannot expose changes inside app.Run's own process. The native probe calls app.Run but does not compare its process cwd/environment before and after successful and rejected calls.
- APP-021: "It MUST inspect only the exact selected claim rather than require every claim to decode, and MUST NOT open disposable code, unrelated private content, transitive dependency contracts or inputs excluded by the public source/context selection contracts." — missing assertion: the worktree interposer cannot observe reads inside Git child processes, and no Git inventory/tree/body read guard limits them to the selected source. Successful context with a corrupt unrelated claim does not reject ignored broad claim reads, and ignored Git reads of excluded historical/private bodies also remain unobserved. Independently controlled allowed and forbidden selected-ref/tree/body reads are needed at the CLI boundary.
