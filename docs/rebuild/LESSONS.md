# Lessons from the first rebuild run

The first run (2026-10-21 to 2026-10-26) is preserved on the `rebuild-v1` branch. Its DECISIONS.md entries are cited as v1-Dnnn. These are the findings worth keeping. Treat them as bounded historical input for design-derived authoring, not as extra specs. References to "the bootstrap" below describe the retired 0.4.0 controller used by that run; they are not findings against the fresh 0.1.0 candidate. Do not consult the old branch or implementation as an oracle.

## Design findings

- **Handles are per change.** A `new:` handle is scoped to the change that declares it, even when two queued changes use the same handle text. `already_satisfied` may list a local Added handle (v1-D009, v1-D011).
- **Archived prose is history.** Handle-like text in an archived change's human sections is not a live reference; archived edit sections and metadata still are (v1-D010).
- **Spec history order.** `show --history` orders by the implementation provenance timestamp, newest first, with a total-order fallback. The bootstrap gets this wrong (v1-D012).
- **Coverage evidence.** A skipped testcase covers nothing. The hook's exit code must agree with the report (C11). A test that started and never finished is failure evidence, never a pass. The setup hook now enforces that (v1-D013, v1-D040).
- **Placeholders.** Hook placeholders are substituted once into the authored command, quoted, and never rescanned (v1-D014).
- **Boundary suppressions.** Baseline and opt-out suppression apply only to a boundaries hook that finished normally; a cancelled or failed launch is never clean evidence (v1-D015).
- **Explicit versus omitted keys.** An explicitly empty metadata key differs from an omitted one, both in module defaults and in the "worker changed approved metadata" check (v1-D016).
- **Claims.** Create is create-only. Updates and releases compare and swap against the exact claim object, so a newer owner is never overwritten (v1-D018).
- **Approved metadata.** During a held change, project metadata comes from the claim's base commit, not from current files, so a worker edit cannot change which hooks run before it is rejected (v1-D024).
- **Eligibility.** Blocking is decided on each candidate's projected graph: the preceding queue plus that candidate's own `set` (v1-D026).
- **Approval check.** Comparing a started change with its copy on the base branch must ignore runtime fields (run, phase, blocked, Decisions, Questions, Audit, Provenance) and map `new:` handles to their assigned IDs (v1-D027).
- **Allocation.** Assigning IDs must not add edit sections the author left out (v1-0039).
- **Graph during a held change.** The boundaries hook calls `FIREWEED_BIN modules --json`; while a change is held, that output must be the projected graph including the change's `set` (v1-D042).
- **Embedded data.** `help`, `version`, `explain` and `playbook` work outside a project (v1-D007).

## Environment facts

- On macOS, `t.TempDir()` returns `/var/...` while `os.Getwd()` returns `/private/var/...`. Canonicalize with `filepath.EvalSymlinks` before comparing (v1-D017).
- Parse Git's stdout only. Successful Git commands can write to stderr (v1-D023).
- `git status` may rewrite the index; use `git --no-optional-locks status`. Don't copy the index to avoid writes, because a copied index changes Git's racy-timestamp detection (v1-D039).
- Packed refs that differ only in case collide on case-insensitive filesystems (v1-D038).
- `go test` has a default 10-minute timeout per package. Keep suites well under it rather than raising it (v1-D035).
- `FIREWEED_BIN` in a hook must name the binary that ran the hook, never the controller.
- The first run rejected its default Codex model (v1-D003). Current stage-2 selection is Codex `gpt-6-sol` with high reasoning effort per GOAL.md. A later read-only Sol probe passed at low effort; full high-effort doctor probes remain startup work. Do not infer current availability from old observations.
- Public facades that share types with their own `internal/` code need a retained `model/` subpackage to avoid import cycles (v1-D004).

## What went wrong

- **Hardening spiral.** One-line principles ("never touches the network") became per-spec syscall-level proofs: 19 cgo eval files, Git proxies, I/O interposers. Now: conventions in SYSTEM.md, one PATH-shim test.
- **Edge-case chains.** Three gitstore changes chased processes escaping cancellation and added native objects and 12 MB of per-architecture binaries. Now: process groups plus a documented limitation; no native code.
- **Non-converging audits.** Audits critiqued the test instrumentation, and one change looped through evals and audit five times. Now: audit scope limited to unasserted spec clauses, at most two rounds, then block for a human.
- **Specs as essays.** Specs averaged about 110 words and 6 clauses, plus about 13,000 words of Decisions that workers re-read every step. Now: at most 3 clauses and about 80 words per spec.
- **Decisions as a lab notebook.** Workers wrote test-instrumentation and recovery notes into Decisions, which fireweed appends to SPEC.md: about 13,000 words that every later worker re-read and regeneration would follow. Now: Decisions are one or two sentences of product behavior, stale ones are dropped through Removed decisions, and process notes go elsewhere.
- **Evidence ceremony.** Hash manifests, inventory proofs and a 218 KB progress log. Now: one line per applied change.
- **No human checkpoint.** The agent authored and applied its own changes for five days. Now: the goal explicitly authorizes local authoring and execution, with human checkpoints after five applied changes, on a blocked change, and before a third consecutive change to the same module/behavior. Do not invent an extra per-change approval pause.
