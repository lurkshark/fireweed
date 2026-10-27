# Rebuild progress

Keep this file short: update sections in place and add one line per applied change.

## Status — 2026-10-26

Stage-2 startup and standalone NEW CLI checkpoint passed. Human review 1 approved
continuation at `3311f676854420e509412e4ca6bc113aeefd9864` (user: "I reviewed it;
continue"). Three applied changes since that review. No pending changes or claims. The manual
lifecycle fixture is not built and every acceptance row remains unverified.

## Controller and baseline

- Pinned controller: `.fireweed/bin/fireweed-bootstrap`, copied from the tested
  sibling executable; version `0.1.0`, embedded source
  `7f925190878e38fc38f30d16bb1cf47aca098cf7`, vcs.modified=false.
  Candidate/destination SHA-256:
  `f1516b5686435c23b54b71c5ac554aba6859b5454f072e0c5eb031dfd6f0fd4f`.
- NEW stage-2 setup baseline: `be35552b706aa3000ff8f11306d129ff1d701b85`.
  Use this explicit nonempty baseline..HEAD for history/CI acceptance. Historical
  restart `d991c26472ad990a5d050a07ce7dc873d1900433` is not the new baseline.
- Product: `bin/fireweed`, clean source
  `5f09706c3cbbc3aca57f8f880f81d309dda5729e` (vcs.modified=false); outside-project
  information commands pass. modules/new-change work, including nested-root
  discovery and JSON; remaining commands explicitly fail as unimplemented. No controller delegation.

## Startup evidence

- Initial tree at `0fee737d7bc2c21c6bf5e28faf299301e2542cb1` was clean, with no
  changes/archives/claims. Read all required design/supporting docs and HANDOFF.
  Design SHA-256 matches handoff:
  `52623b89d8184680f535519e82ceda8a49490b3d4dec000347cbd2cde429b21a`.
- Disposable Go init probe: `/tmp/fireweed-stage2-setup.pyaO9N`. Adapted protected
  version range to >=0.1,<0.2, canaries.yaml maps for evals/ layout, {{name}}
  templates and api.go.tmpl. Retained Go JUnit/boundary hook including unfinished
  tests and package failures; setup copies invoking FIREWEED_BIN into scratch.
  Managed AGENTS already matched. Protected paths and bounds not weakened.
- Worker/auditor argv: Codex ephemeral, ignore-user-config, gpt-6-sol, high effort,
  workspace-write. Audit enabled, max_rounds=2; max_unapproved=5. Startup check and
  full doctor passed, including four canaries/removal recovery and BOTH real
  Sol/high probes, no diagnostics. Scratch worktree removed. Evidence is ignored
  under `.fireweed/evidence/startup-doctor.*`.

## Applied

- `standalone-information` / AST-001–AST-003: evals `a4ceb17` (8 genuine failures), independent audit `826b9e9` (no findings), impl `2248b8e` (8/8 pass, direct dependents/boundaries pass); archive `0001-standalone-information.md`.
- `standalone-cli` / APP-001–APP-003: evals `1298acf` (30 genuine failures), controller audit `7de7d30` (independent findings discarded due folded quote line breaks), impl `db77de7` (30/30 pass); archive `0002-standalone-cli.md`. Independent audit did NOT pass; corrected in next eval_fix.
- `json-error-assertions` / APP-003 eval_fix: evals `eb4704d`, independent audit `e184bb6` (no findings), impl `8796d37` (no code edits; all 26 app cases pass); archive `0003-json-error-assertions.md`.
- `lifecycle-primitives` / CON-001–CON-005: evals `11d0180` (18 failures, 9 passes; every changed ID fails), independent audit `6e341a4` (no findings), impl `8e37dd7` (27/27 own, 26 app and 8 assets cases pass; all direct dependents/boundaries pass); archive `0004-lifecycle-primitives.md`.

- `fixture-authoring` / PRJ-001–PRJ-003: initial evals `1bcba65`, findings `5c647a2`, corrected evals `bfee6b8` (24 failures, 1 pass; all changed IDs fail), independent passing re-audit `06f17c6`, impl `91f036c` (25/25 own and 26 app cases pass, direct dependents/boundaries pass); archive `0005-fixture-authoring.md`. One persistent audit round.

- `project-commands` / APP-004–APP-005: evals `da5dba8` (12 failures, 39 passes; both new IDs fail), independent audit `98e5852` (no findings), impl `5f09706` (51/51 app cases and boundaries pass); archive `0006-project-commands.md`. 1 applied since review 1.

- `change-handles` / CON-006: evals `a22fcd4` (1 failure, 33 passes; CON-006 fails), independent audit `a194629` (no findings), impl `7580897` (34 own, 51 app, 8 assets and 25 project cases pass; boundaries pass); archive `0007-change-handles.md`. 2 applied since review 1; shared-base unapproved count 3.

- `queue-authoring` / PRJ-004–PRJ-007: evals `6e3ca5f` (24 failures, 25 passes), independent audit `776b4d0` (no findings), impl `2ef37df` (53 project and 51 app cases pass, dependents/boundaries pass); archive `0008-queue-authoring.md`. Two contradictory eval fixtures corrected before freezing. 3 applied since review 1; shared-base unapproved count 4.

## Process findings and limits

- Real workers/auditors used required Sol/high. Manually coordinated controller
  transitions record agent=manual in Provenance; actual worker transcripts are
  `.fireweed/evidence/<slug>-*`. Auditors ran separately in `/tmp/fireweed-audit-*`
  with only copied change/projected spec/evals, never implementation or SYSTEM.
- Standalone-cli audit quotes folded spec line breaks. Controller discarded them
  and advanced to impl; an attempted eval edit was rejected and restored exactly
  before implementation. A separately audited eval_fix repaired the known gaps.
  Future quotes must be literal contiguous substrings; inspect phase/audit_rounds
  after every transition. No forbidden eval/interface edits committed in impl.
- Evals-only B03 warnings disappeared after implementation. No claim remains.
- Contracts worker hit account usage limit after writing its implementation. On
  resumption, full controller verification passed the actual code; controller
  applied it without another worker. Usage API now permits ordinary use (1% of
  5-hour window used). No model/provider substitution or fake-agent evidence.
- On resumption origin/main was observed at `6e341a4`; this process did not push
  or move the base. Review count remains four; do not silently reset it.

## Checkpoint verification

- Checked source: `91f036cbb33981a5fb72798da3edec9bdc90b077`. Exact nonempty
  controller history/CI range:
  `be35552b706aa3000ff8f11306d129ff1d701b85..91f036cbb33981a5fb72798da3edec9bdc90b077`.
- Build, `go test ./...`, `go test -race ./...`, `go vet ./...`, controller check,
  check-commits and CI all exited 0. CI ran doctor canaries/recovery, boundaries,
  history and all module coverage: app 26, assets 8, contracts 27, project 25
  cases (86 total including parents), all pass; other modules have no specs yet.
  CI reports doctor=true and no diagnostics. Logs/JSON:
  `.fireweed/evidence/checkpoint-*`. Go printed a nonfatal module-stat-cache
  sandbox write warning during build; build exited 0.
- NEW executable smoke outside a project: help, version --json, explain P2-01 and
  playbook audit all exit 0 with empty stderr. Permanent app eval also exercises
  malformed config, JSON and usage errors. Primary worktree only, no claims.
- These are intermediate checks, not complete acceptance. Fresh-checkout gates,
  full NEW lifecycle/init/doctor/run/CI, full surface, packaging, controller/product
  lifecycle comparison, real drills and isolated NEW self-hosting are unrun.

## Next / resume — review 1 approved

The user explicitly approved continuation from the first five-change checkpoint.
The review count restarts at zero; stop again after five applied changes or a
blocked change. This approval does not move base refs or bypass P07: the current
shared-base unapproved count is one, so P07 may stop this batch before five.
The full goal remains incomplete.

Project queue/projection/stamp/check now work at the public API. Wire stamp/check
and projected views through app next, then add only needed Git, hook/coverage and lifecycle
behavior to finish the NEW manual fail-first/audit/impl/archive fixture. Retain
explicit failures for unimplemented commands; do not polish working primitives
before that slice. Full project grammar/zone validation is still incomplete.

If a change is active on resume, inspect its context, tree and claims first.
Protected author edits belong only between changes, in separate commits.

- Manual slice fixture created at `/tmp/fireweed-stage2-manual-slice/project`
  with hand-authored Go/config/SPEC files. NEW modules --json and new-change toy
  answer --json both exited 0 from packages/toy; queued path changes/0001-answer.md.
  No controller used as the fixture subject. Next add new: handle parsing before
  author/stamp/check and lifecycle; no fixture init or lifecycle success claimed.

- Manual fixture author commit: `95c6a35` with explicit local Git identity.
  changes/0001-answer.md declares new:answer and Answer MUST return 42; no IDs
  assigned. The product still has no stamp/check/lifecycle commands at this point.
