---
fixes:
    - FLOW-015
    - FLOW-023
module: workflow
phase: impl
run: f9347464cbe4af07a1cd7b49
---

## Motivation

Compose the existing checked primitives into read-only verification for each held worker phase. Verification must restore committed approval before inspecting worker feedback, retain all useful checking evidence, and reject hooks that mutate protected state. Existing approval lookup and moved-HEAD checks also attempt lazy transport for unavailable selected objects; these genuine regressions need permanent failing public evals.

## Interface

Add workflow.Verify(ctx context.Context, start, slug string) (Verification, error). Empty slug selects the current canonical worktree's unique held claim. Add retained model Verification with string Change/Phase, []checks.Diagnostic Diagnostics, []checks.Coverage Coverage, []gitstore.Edit Edits and []Finding Findings; JSON names are change, phase, diagnostics, coverage, edits, findings (findings omitempty). Finding contains string ID, Clause, Text and bool Untestable with JSON id, clause, text, untestable. Expose aliases through api.go. Keep all existing APIs unchanged. New compiling implementation stubs may reside only in new disposable files.

## Added

### FLOW-041
Verify MUST validate a nonempty explicit slug before repository access, rejecting invalid selectors with corrective C05. It MUST discover the canonical checkout from start, select the exact named held claim or the unique claim owned by this checkout when slug is empty, and reject missing, ambiguous, malformed or foreign selected ownership with corrective P02. Explicit selection MUST remain independent of malformed unrelated claim payloads. It MUST restore the selected approved snapshot and exact started State from the claim's committed Base through LoadApproval, never use current worker configuration/SPEC/queue metadata as authority or silently fall back to unclaimed source. Fatal selection/source failures MUST return a zero Verification and useful non-nil error. Valid ordinary and linked SHA-1/SHA-256 worktrees MUST work from nested start paths.

### FLOW-042
For a valid selection Verify MUST use Project before the selected change and ApplyChange for that change exactly once, including preceding queued effects and the selected set while excluding later effects. It MUST validate the resulting effective project and dependency graph before executing hooks. It MUST compose CheckWorktree with that inclusive approved target and preserve its actual Edits and independent diagnostics. Invalid graph, approved worker state, ownership, HEAD, edit zones or path containment MUST prevent all verification hooks; current blocked MUST add corrective P05 and also prevent hooks. Invalid current change text or unsafe/unreadable current paths MUST fail closed through the existing guards. Structured checking failures MUST return initialized Verification data with nil error, preserving Change/Phase from approval, non-null Diagnostics/Coverage/Edits and useful diagnostics; source failure remains governed by FLOW-041.

### FLOW-043
In evals Verify MUST evaluate exactly the selected inclusive target module with the approved effective hooks/configuration through public Evaluate. Fail-first targets MUST include every Added, Modified and fixes ID, with approved already_satisfied passed under public coverage policy. Coverage, hook status/output, testcase data and all coverage diagnostics MUST remain visible even for expected fail-first failures, unrelated failures, missing/malformed reports or hook errors. It MUST NOT evaluate dependent modules or run boundaries during evals, and MUST NOT treat worker-modified hooks or metadata as approved. Removed IDs MUST remain absent from the projected coverage contract. Legitimately passing fail-first checks MUST yield no verification errors.

### FLOW-044
In audit Verify MUST run no hooks or agents and inspect only the current parsed Audit section. Each retained finding MUST follow the documented one-line format with an optional [untestable] marker, a real eligible target spec ID and a nonempty quote that is an exact substring of that inclusive spec's body. It MUST retain the original line in Text and expose the marker in Untestable, preserve finding order, and discard unknown/review/quarantine IDs, empty/inexact clauses and malformed lines. If no eligible findings remain, a passing audit MUST say No findings case-insensitively and list every Added, Modified and fixes ID using precise public ID matching; an absent/empty/incomplete audit MUST produce corrective P1-03. Valid findings MUST be returned for later Done handling without advancing phase, rewriting Questions or blocking the change. A complete passing audit MUST return empty non-null Findings and no verification error.

### FLOW-045
In impl Verify MUST evaluate the selected inclusive target first, then each effective graph dependent exactly once in ascending module-name order according to approved verify_dependents all/direct/none, default all. Transitive diamond dependents MUST NOT be duplicated; unrelated modules MUST NOT be evaluated. Dependent coverage MUST use approved effective module metadata and MUST NOT impose the selected change's fail-first or already_satisfied policy. It MUST run public Boundaries once against the complete inclusive graph after coverage, even when a coverage hook fails. All available own/dependent coverage and boundary diagnostics MUST remain visible, and failing hooks or coverage MUST NOT be reported as a successful verification. Direct and none MUST retain own coverage and boundaries.

### FLOW-046
After verification hooks Verify MUST recheck the exact stored claim, original HEAD and approved worker state/edit policy, returning the resulting updated Edits and retaining independent pre-hook and coverage diagnostics. Forbidden files or metadata written by hooks and ownership/HEAD changes MUST therefore be rejected. It MUST NOT repair or erase hook changes, create commits, move/reset/stage HEAD/index or advance/update/release claims. Caller cancellation MUST stop the owned active hook and prevent later hooks; errors.Is(err, context.Canceled/DeadlineExceeded) MUST remain usable, with already collected Verification evidence retained when available. A context cancelled before selection MUST return zero Verification and its cancellation error without executing hooks.

### FLOW-047
Apart from explicitly configured verification hooks, Verify MUST be local, bounded and noninteractive: no lazy fetch, push, branch/worktree creation, agents, bootstrap or network commands. Successes and rejected calls MUST preserve original files, ignored data, index bytes, refs, common Git metadata, parent cwd/environment and remove owned temporary storage. Hook-created changes are retained under FLOW-046. LoadApproval and CheckWorktree MUST also reject unavailable selected required claim/HEAD objects without transport, while unrelated missing bodies/ancestors MUST NOT obstruct otherwise locally available results. Repeated results MUST have independent writable slices/maps and deterministic diagnostic ordering by Path, Code, Severity, Message with exact duplicate diagnostics removed; all nonzero Verification collection fields MUST be non-null, including Findings before JSON omitempty encoding.

## Decisions

[checked-verify]
Use the existing public guards and hook/evidence engine, retaining their exact diagnostics rather than inventing a weaker parallel implementation. Fatal selection/source failures use error/zero data; normal checking failures use structured data/nil error. Cancellation retains the standard context error channel. This read-only API does not advance phases, apply specs, persist audit findings, implement retries, or prove real regeneration. Subsequent authored changes own Done, CLI composition and runner behavior.

The bootstrap reference uses substring containment for passing audit IDs; precise ID matching follows the design's spec token identity and prevents ALP-0010 from certifying ALP-001. Public APIs do not require reconstructing a whole unrelated claim inventory for an explicit selector. Ordinary reference comparison is permitted here; no regeneration drill is active.

EVALS start assigns FLOW-041–047/runf9347464cbe4af07a1cd7b49 and reports the normal P04 unmerged warning, using the user's local goal authority without claiming a human merge. Only the declared public facade/new retained types, new eval files and a new compiling disposable stub are added; every existing non-API workflow source file remains byte-identical. Initial compilation found a new helper type mismatch and was corrected before behavioral execution; compile failure is not fail-first evidence. The first calibration independently reproduces16 genuine selected-claim/moved-HEAD transport leaves with unrelated controls passing, but finds a missing required Motivation in the new earlier-queue fixture. Correct that fixture before running complete fail-first and formal audit; no runtime product behavior or existing eval is edited.


Pre-freeze calibration corrected only new evals: public coverage accepts a nonzero hook with genuine executed failures; effective graph diagnostics may retain their native corrective location; cancellation and post-hook checks compare full available evidence. Added native controls prove invalid-selector no-I/O, selected-object transport, configured callback, and Git operation observer liveness. Git Trace2 is paired with a same-PID exec recorder for cwd/canonical metadata, permitting owned bare-view setup while detecting caller branch/push/commit/worktree operations, including nested cwd, symlink aliases, literal option operands and send-pack. Restricted separate auditor /root/audit_start_cli reports No findings for FLOW-041–047 and fixes FLOW-015/FLOW-023 after these corrections. This is pre-freeze source review only; formal Audit follows accepted EVALS. Two interrupted earlier full verification runs used obsolete observer source and are excluded. No existing non-API implementation/eval was changed.


Formal independent Audit after EVALS6aea37e found six assertion gaps; controller commitc110851 records them and returns to EVALS. New permanent regressions assert implicit ownership filtering with a valid foreign claim, native static projection rejection before hooks, approved D04 warning retention beside coverage/post-hook failures, nonempty Errors storage, two-sided empty-collection writes and post-startup stdin/prompt behavior on success/rejection. Native Git controls now observe real transient caller-index staging and resets while permitting owned index and metadata setup. Actual cmd_mode distinguishes path-only resets from full HEAD resets, including empty/nonempty pathspec files; both caller/owned storage combinations are live. After correcting one additional reset-scope oracle issue, restricted pre-freeze re-review reports No findings for all seven Added and both fixes IDs. This still requires a fresh complete EVALS freeze and subsequent formal Audit; product implementation remains untouched.

## Rejected alternatives

Loading the dirty current project before approval can redirect hooks or permissions. Running hooks after failed guards can execute unauthorized configuration. Treating a failing-before report as a Go compile failure provides no ID-linked evidence. Repairing hook edits or committing during Verify would violate the manual iteration contract.

## Provenance

- evals | 2026-10-24T06:09:07Z | agent: manual | fireweed: 0.4.0 | playbook: 5867363302d4c874074520977b1627d26991d95ee2fbacb9ff63e827661edd0b | tests: 2115 | failing: 181 | passing: 1934
- audit | 2026-10-24T06:14:58Z | agent: manual | fireweed: 0.4.0 | playbook: a9c982d24b476621316049b527e6c7bbc319f06b12268af150fa7d8c6df92230 | tests: 0 | failing: 0 | passing: 0
- evals | 2026-10-24T15:35:06Z | agent: manual | fireweed: 0.4.0 | playbook: 5867363302d4c874074520977b1627d26991d95ee2fbacb9ff63e827661edd0b | tests: 2130 | failing: 189 | passing: 1941
- audit | 2026-10-24T15:41:13Z | agent: manual | fireweed: 0.4.0 | playbook: a9c982d24b476621316049b527e6c7bbc319f06b12268af150fa7d8c6df92230 | tests: 0 | failing: 0 | passing: 0

## Audit

No findings. Checked FLOW-041, FLOW-042, FLOW-043, FLOW-044, FLOW-045, FLOW-046, FLOW-047, FLOW-015, FLOW-023.

Fresh restricted source-only auditor /root/audit_checked_verify_final checked the strengthened frozen evals against the permitted public contracts and all seven Added clauses plus both fixes. The six earlier assertion gaps are covered, including live caller-index operation controls and post-startup stdin/prompt observation. No implementation/history/bootstrap/evidence inspection, edits, tests, product execution or Git/controller commands by the auditor.
