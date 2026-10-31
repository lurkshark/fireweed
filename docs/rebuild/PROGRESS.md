# Rebuild progress

Current handoff only. Per-change evidence remains in archived Provenance
(`changes/archive/`); no separate report inventories or hash manifests.

## State — 2026-10-31

**EC1–EC6 passed; the early Go core-loop cutover is complete.** EC6 installed
the separate NEW pin, aligned protected instructions/setup and passed all
required pin/status/check/doctor smokes. Full product completion remains open.
Main was clean at the tested candidate before this author handoff, with an empty
queue, no shared claims and all acceptance controllers/workers returned.
Only the earlier push through `dae9b28` was owner-authorized; later work is local.

0131–0138 are applied with their archived evidence preserved. 0134 supplies
replaceable-stub guidance; 0135 retains JUnit failures; 0136 supplies bounded
own/dependent retry assertions; 0137 separates Decisions from process Notes;
0138 isolates audit inputs, validates Audit-only edits and cleans up workspaces.
0138's owner-authorized recovery exception is recorded in its archived Notes.
Actual NEW invocations confirmed the guidance and bounded retry feedback; the
retry fixture remains inspectable at
`/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-0136-smoke-agyg5ph1`.

Two normal app eval fixes (`4c794a6`, `20a8117`) removed obsolete audit-cwd
assertions without changing public behavior. Acceptance exposed two core
failures, fixed through ordinary fail-first/audit/application lifecycles:
0139 (`f79a7be`) accepts a valid first Audit insertion before Provenance;
0140 (`9822408`) renders nested doctor templates and preserves canary checks.
Their independent audits and full passing results remain in the archives.

## Gate results

- **EC1:** all prerequisites and demonstrated follow-ups applied; clean main,
  empty queue and no shared claims before handoff.
- **EC2:** focused core review and affected fixes complete, with no unresolved
  finding: fail-first, immutable metadata/zones, independent audit inputs/history,
  Audit-only import/cleanup, approval/P07, saved attempts, worker-commit recovery
  and bounded time/usage stops. NEW has no runtime bootstrap dependency.
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
- Workers and auditors: `codex exec --ephemeral --ignore-user-config
  -m gpt-6-sol -c 'model_reasoning_effort="high"' -s workspace-write {prompt}`.
  Audits use the configured controller audit step and temporary input-only
  workspace; no manual isolation or model/effort substitution.

## Standing rules and remaining scope

- Stop on any blocked change. Audit rounds and shared-base P07 still apply.
  Never record the unapproved count or its limit here; read live controller
  `status --json` and `authoring.max_unapproved` before starting work.
- Push only on explicit owner instruction. Protected author edits stay separate
  from controller-owned worker phases, Specs/Decisions and archives.
- Known warnings: bootstrap C01 identifies embedded recipe specs outside the
  module graph; C18 flags repeated generic requirement sentences. B02 reports
  existing APP-103/104 and CHK-055 case counts. Go's module stat-cache write
  warning is nonfatal. NEW's pre-handoff C16 managed-guidance warning is addressed
  by aligning AGENTS in this handoff. Zero warnings is not a gate.
- The earlier contracts drill's native result passed, but the agent searched
  other modules' internal code, so it does not meet independence. Its retained
  ref is `refs/fireweed/drills/contracts/20261028T103005.180922000`; no new real
  drill or regeneration was run for this milestone.
- Full product completion remains open: real drills/regeneration validation,
  gc/migrate and stale/fix conveniences, additional recipes, version/file-format
  migration, remaining warnings, packaging and every later full-completion gate.
  No real drill/regeneration or deferred product work is authorized here.

## Next

This milestone is complete. Use the recorded NEW pin for subsequent authorized
core-loop work. Await owner scope for later product completion; no further
drill/regeneration, deferred feature work or push is authorized by this goal.
