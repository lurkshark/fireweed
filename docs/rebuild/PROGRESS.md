# Rebuild progress

Current handoff only. Per-change evidence remains in archived Provenance
(`changes/archive/`); no separate report inventories or hash manifests.

## State — 2026-10-31

**EC1–EC6 passed; the early Go core-loop cutover is complete.** EC6 installed
the separate NEW pin, aligned protected instructions/setup and passed all
required pin/status/check/doctor smokes at handoff commit `c1de8f4`.
Full product completion remains open.

The owner selected the next bounded milestone on 2026-10-31: implement drill
support and prove one independent contracts drill in an isolated clone, with
main's NEW pin unchanged. PLAN orders execution, ACCEPTANCE DR1–DR4 defines
completion, and GOAL is the restart prompt. **Stopped at the required owner
checkpoint:** 0141-drill-lifecycle blocked during implementation at `ec5c27e`.
DR1 is incomplete; DR2–DR4 have not started. No candidate CLI fixture, candidate
identity, isolated proof clone or new native drill report/ref exists yet.
Main is clean with no shared claims or live controller/worker processes; the
original NEW controller hash remains unchanged. The unrelated guidance worktree
was preserved.
Only the earlier push through `dae9b28` was owner-authorized; later work is local.

0131–0138, the two app audit-cwd eval fixes, and demonstrated follow-ups 0139
(Audit insertion) and 0140 (nested doctor templates) are applied. Their phase
checks/audits and 0138's owner-authorized recovery exception remain in archived
Provenance/Notes. Preserve these archives; the completed queue is not restart work.

## Drill checkpoint — 0141

The focused Regeneration/R16 review led to the runner API/lifecycle change.
The pinned NEW controller recorded genuine fail-first, one findings round and
a passing restricted independent re-audit, with frozen evals/interfaces. Both
auditors used the configured Codex model/effort. Manual Audit-only formatting
and carry-forward of previously passed Checked IDs explain the audit lines'
`agent=manual` Provenance; the saved attempt's Notes record those completions.

Implementation exhausted the normal attempts bound. The final verification
report has four failing runner cases (including the parent case) and passing
app dependents. The frozen fake agents for `agent_retry`, `verify_retry` and
`usage_once` leave their initial bad file unchanged on the later successful
invocation, while RUN-056/RUN-057 assertions require it to become `new`.
Implementation cannot repair those fixtures or invent replacement worker output.
The cleanup failure case passed after the bounded implementation correction.

The controller saved implementation and step notes at
`refs/fireweed/attempts/drill-lifecycle/1` and restored the checkout/released the
claim. The queued change retains its Provenance; no replacement change, eval_fix,
retry reset or gap fix was authored. Owner direction is required for frozen-eval
recovery before resuming. Review also identified unresolved native-report
integration details in the saved implementation: its compact frontmatter date
is unreadable by existing drift tracking, and model recording reads `agent.model`
rather than the configured command's `-m` value. These are pending review, not
additional authorized changes after this checkpoint.

The owner authorized a scoped eval_fix recovery on 2026-10-31, preserving
coverage, independent audit, the pin and bounds. Recovery preflight found a
further checkpoint before any frozen edit or new claim: normal eval_fix cannot
repair this unapplied spec change under the current pin. RUN-056/RUN-057 exist
only in 0141's projection (public `show runner RUN-056` returns U02), and main
still has the drill stub. An earlier eval_fix lacks those specs and passing
implementation; a later one is held by the blocked same-module change. The pin's
Start selects the eligible queue head and has no exception for this case.
Eval_fix also cannot be mixed into 0141's Added specs or waive its all-own-pass
gate. No source files, bounds, controller, shared claims or attempt refs changed.

The minimal further recovery would need an explicit one-time frozen-eval
exception: add the missing successful `new` write in the three fake-agent
branches, independently audit the corrected assertions in restricted inputs,
commit only the reviewed correction as an author while unclaimed, then resume
0141 from saved attempt 1 with its IDs, run and audit rounds preserved. This is
not the pin's normal eval_fix lifecycle, so the scoped eval_fix authorization
has not been silently broadened to that exception. No such edit was made.

## Completed cutover evidence

- **EC1:** all prerequisites and demonstrated follow-ups applied; clean main,
  empty queue and no shared claims before handoff.
- **EC2:** focused core review and affected fixes complete, with no unresolved
  finding: fail-first, immutable metadata/zones, independent audit inputs/history,
  Audit-only import/cleanup, approval/P07, saved attempts, worker-commit recovery
  and bounded time/usage stops. NEW has no runtime bootstrap dependency.
  Actual retry CLI fixture:
  `/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-0136-smoke-agyg5ph1`.
- **EC3:** clean candidate `281ae9498c0fa20b44dc525494427ff123331ab3` passed
  build, `go test ./...`, `go test -race ./...` and `go vet ./...` with the
  specified GOCACHE. Both bootstrap and NEW passed `check`, `doctor --no-agent`,
  `check-commits` and `ci` with doctor enabled over the explicit nonempty range
  `be35552b706aa3000ff8f11306d129ff1d701b85..281ae9498c0fa20b44dc525494427ff123331ab3`.
  A CHK-005 startup timeout and EC5's RUN-036 budget failure each passed a focused
  unchanged rerun and the unchanged full command; no bounds or requirements changed.
- **EC4:** equivalent bootstrap/NEW lifecycle smoke passed: schema-1 envelopes,
  three phase commits, matching held refs, restricted audit inputs and cleanup.
  Intentional differences are DECISIONS D005/D010/D011. Reused the valid
  `66f0c67` lifecycle proof after the doctor-only fix; fixture:
  `/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-ec4-t1_zlt95`.
- **EC5:** the exact final NEW binary governed a real spec change with the
  configured workers and separate auditor, normal approval and bounds, genuine
  fail-first (`e4635b0`), audit (`8202abf`) and application/archive (`da08353`).
  NEW CI passed over `17439f926aba8d0e5ffa6b26f6b8499bf56dc4ff..da08353`.
  Frozen interfaces/evals and phase commits/Provenance were inspected. The
  observed audit workspace contained only change/spec/evals, lacked held history
  and implementation, and was removed; the clone is clean and unclaimed with
  no bootstrap executable installed. Retained evidence:
  `/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-ec5-f78ffohx/candidate`.
- **EC6:** verified the separate regular NEW pin's hash/source and distinct
  subject file, with bootstrap's recovery hash unchanged. Protected AGENTS,
  SYSTEM, setup/recipe and worker PATH are aligned in this author handoff.
  The promoted pin passed `status --json`, `check` and `doctor --no-agent`
  with no diagnostics; PATH resolved `fireweed` to the recorded NEW pin.
  Queue is empty, shared claims absent and scratch worktrees cleaned up.

## Controller pin and execution

- Current NEW controller:
  `/Users/joel/Workspace/fireweed/.fireweed/bin/fireweed`, version **0.1.0**,
  source `281ae9498c0fa20b44dc525494427ff123331ab3`, SHA-256
  `bf0dbc59b7fa7ce2ff957a6f21f34903655b6b054a10695cfeef5ece67fa3575`.
  Built from a clean candidate (`vcs.modified=false`). It is a separate copy
  from mutable `bin/fireweed`; worker builds cannot replace this pin.
- Explicit recovery only:
  `/Users/joel/Workspace/fireweed/.fireweed/bin/fireweed-bootstrap`, version 0.1.1,
  source `78242c8251ff5c2ed7007502dc759c414dca5d75`, unchanged SHA-256
  `ba14dfa072bfa72756d41e1fa0b13fec0df654a7ee629e654cbeb954df8d4683`.
  No automatic fallback, cross-controller claim resume or repin during a live run.
  Recovery repinning requires explicit owner instruction.
- Worker PATH starts with the current tree's `.fireweed/bin`; on main use
  `export PATH=/Users/joel/Workspace/fireweed/.fireweed/bin:$PATH`.
  Hooks receive the actual invoking controller as FIREWEED_BIN. Protected setup
  copies that executable atomically into scratch `.fireweed/bin/fireweed`.
- Use `GOCACHE=/tmp/fireweed-rebuild-cache`. History baseline remains
  `be35552b706aa3000ff8f11306d129ff1d701b85`; always pass an explicit range.
- GOAL records the exact worker/auditor/drill argv; no model/effort substitution.
  Audits use the configured controller audit step and temporary input-only
  workspace. Reconstruction inputs must separately meet DR3 independence.

## Standing rules and remaining scope

- Stop on any blocked change, failed independence/access or drill gap/contradicted
  Decision needing owner agreement. No automatic gap fixes or extra module drills.
  Audit rounds and shared-base P07 still apply.
  Never record the unapproved count or its limit here; read live controller
  `status --json` and `authoring.max_unapproved` before starting work.
- Push only on explicit owner instruction. Protected author edits stay separate
  from controller-owned worker phases, Specs/Decisions and archives.
- Known warnings: bootstrap C01 identifies embedded recipe specs outside the
  module graph; C18 flags repeated generic requirement sentences. B02 reports
  existing APP-103/104 and CHK-055 case counts. Go's module stat-cache write
  warning is nonfatal. NEW's C16 guidance warning was resolved by EC6.
  Zero warnings is not a gate.
- The earlier contracts drill's native result passed, but the agent searched
  other modules' internal code, so it does not meet independence. Its retained
  ref is `refs/fireweed/drills/contracts/20261028T103005.180922000`. It was a
  bootstrap run and cannot substitute for DR3's independent NEW proof.
- Beyond the current contracts proof, full completion still needs remaining
  modules' independent drills, actual regenerate validation,
  gc/migrate and stale/fix conveniences, additional recipes, version/file-format
  migration, remaining warnings, packaging and every later full-completion gate.
  No additional real module drill, main contracts replacement, real regenerate
  change, main controller promotion or deferred product work is included here.

## Next

Wait for an owner decision on the newly demonstrated eval_fix recovery
constraint: an explicit one-time frozen-eval exception, or separately scoped
controller support. Do not resume, reset bounds, author follow-ups or start a
real drill automatically.
After an authorized recovery, return to the app command and complete DR1/DR2
before the one isolated contracts proof. Keep main's controller pin unchanged.
No push is authorized; all later full-completion gates remain open.
