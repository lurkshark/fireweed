# Rebuild progress

Keep this file short: update sections in place and add one line per applied change.

## Status — 2026-10-26

Human review 2 cleared the shared-base gate at `969ab2f`; one change applied
since that review. The NEW CLI can
load modules, author a change, stamp/check it and display projected specs. The
manual start/fail-first/audit/impl/archive lifecycle is still incomplete; every
acceptance row remains unverified.

Shared-base unapproved count is now 1; max_unapproved remains 5. No pending
changes, active claims or scratch worktrees remain. Continue the existing manual
slice and stop again after five applied changes or an actual block.

## Controller, source and reviews

- Controller: `.fireweed/bin/fireweed-bootstrap`, copied tested executable 0.1.0,
  source `7f925190878e38fc38f30d16bb1cf47aca098cf7`, vcs.modified=false;
  pin SHA-256 `f1516b5686435c23b54b71c5ac554aba6859b5454f072e0c5eb031dfd6f0fd4f`.
  It remains unchanged. Never rebuild it from or consult bootstrap source/tests.
- Stage-2 setup baseline: `be35552b706aa3000ff8f11306d129ff1d701b85`.
  Use explicit nonempty baseline..HEAD ranges; never the historical d991c26 range.
- Product: `bin/fireweed`, built at clean source
  `de65319d646fd0e39e9810415a181dd28911762c` (vcs.modified=false).
  help/version/explain/playbook/modules/new-change/stamp/check/show work;
  other commands and unsupported options explicitly fail. No bootstrap delegation.
- Human review 1: user explicitly answered "I reviewed it; continue" for checkpoint
  `3311f676854420e509412e4ca6bc113aeefd9864`. Review count restarted at zero;
  it is now four. That approval did not move shared refs or override P07.
- At the previous gate, the shared base was `origin/main` at `6e341a4e1c85c254968836644586b1120bee705e`.
  Its tree has archives 0001–0003 and queued 0004. Controller reports 5 unapproved
  archives: 0005 fixture-authoring, 0006 project-commands, 0007 change-handles,
  0008 queue-authoring, 0009 authoring-commands. `explain P07` confirms human review
  and merge into shared base before more work. No further start was attempted.
  No bound or base ref was manually changed to evade the gate.
- Human review 2: the user explicitly authorized pushing reviewed main to
  origin/main as an exception to the no-push instruction. The normal push of
  exactly `969ab2f142d66f0aa84e65536a62ad7df322c1c2` exited 0, reporting Everything
  up-to-date. Direct ls-remote and local tracking both confirm that revision;
  controller status now reports unapproved=0, no diagnostics. Evidence
  `.fireweed/evidence/review2-*`. This checkpoint-specific authorization does
  not add product push behavior or authorize automatic pushes of new work.

## Startup and process evidence

- Initial tree `0fee737d7bc2c21c6bf5e28faf299301e2542cb1` was clean/no claims.
  Required design/supporting docs and HANDOFF read. Design SHA-256 unchanged:
  `52623b89d8184680f535519e82ceda8a49490b3d4dec000347cbd2cde429b21a`.
- Disposable init probe `/tmp/fireweed-stage2-setup.pyaO9N`; protected setup adapted
  version >=0.1,<0.2, version-1 YAML canaries for evals/, {{name}} templates and
  api.go.tmpl, scratch controller setup. Existing stdlib Go hooks retained,
  including unfinished tests/package failures. Managed AGENTS matched.
- Startup check/doctor passed four canaries/removal recovery and BOTH real agent
  probes. All real workers/auditors use Codex ephemeral, ignore-user-config,
  gpt-6-sol/high, workspace-write. Audit enabled/2 rounds; max_unapproved=5.
  Manual controller transitions record agent=manual; actual CLI transcripts are
  ignored in `.fireweed/evidence/<slug>-*`. Auditors received only copied change,
  projected specs and evals in isolated `/tmp/fireweed-audit-*` repositories.
- The first standalone-cli auditor's folded quotes were discarded by controller;
  an attempted frozen eval edit was rejected/restored, then repaired through an
  independently audited eval_fix. Quote exact one-line substrings and inspect the
  actual phase after every transition. No forbidden impl eval/interface edits
  were committed. Current batch's audits passed without findings rounds.
- Earlier contracts worker hit usage limit after writing code; full verification
  passed on resume, then controller applied it. No fake agents/model substitution.

## Applied

- `standalone-information` / AST-001–AST-003: evals `a4ceb17` (8 genuine failures), independent audit `826b9e9` (no findings), impl `2248b8e` (8/8 pass, direct dependents/boundaries pass); archive `0001-standalone-information.md`.
- `standalone-cli` / APP-001–APP-003: evals `1298acf` (30 genuine failures), controller audit `7de7d30` (independent findings discarded due folded quote line breaks), impl `db77de7` (30/30 pass); archive `0002-standalone-cli.md`. Independent audit did NOT pass; corrected in next eval_fix.
- `json-error-assertions` / APP-003 eval_fix: evals `eb4704d`, independent audit `e184bb6` (no findings), impl `8796d37` (no code edits; all 26 app cases pass); archive `0003-json-error-assertions.md`.
- `lifecycle-primitives` / CON-001–CON-005: evals `11d0180` (18 failures, 9 passes; every changed ID fails), independent audit `6e341a4` (no findings), impl `8e37dd7` (27/27 own, 26 app and 8 assets cases pass; all direct dependents/boundaries pass); archive `0004-lifecycle-primitives.md`.

- `fixture-authoring` / PRJ-001–PRJ-003: initial evals `1bcba65`, findings `5c647a2`, corrected evals `bfee6b8` (24 failures, 1 pass; all changed IDs fail), independent passing re-audit `06f17c6`, impl `91f036c` (25/25 own and 26 app cases pass, direct dependents/boundaries pass); archive `0005-fixture-authoring.md`. One persistent audit round.

- `project-commands` / APP-004–APP-005: evals `da5dba8` (12 failures, 39 passes; both new IDs fail), independent audit `98e5852` (no findings), impl `5f09706` (51/51 app cases and boundaries pass); archive `0006-project-commands.md`. 1 applied since review 1.

- `change-handles` / CON-006: evals `a22fcd4` (1 failure, 33 passes; CON-006 fails), independent audit `a194629` (no findings), impl `7580897` (34 own, 51 app, 8 assets and 25 project cases pass; boundaries pass); archive `0007-change-handles.md`. 2 applied since review 1; shared-base unapproved count 3.

- `queue-authoring` / PRJ-004–PRJ-007: evals `6e3ca5f` (24 failures, 25 passes), independent audit `776b4d0` (no findings), impl `2ef37df` (53 project and 51 app cases pass, dependents/boundaries pass); archive `0008-queue-authoring.md`. Two contradictory eval fixtures corrected before freezing. 3 applied since review 1; shared-base unapproved count 4.
- `authoring-commands` / APP-006–APP-008: evals `5720dce` (27 failures, 51 passes), independent audit `b2aaa14` (no findings), impl `de65319` (78/78 app cases and boundaries pass); archive `0009-authoring-commands.md`. Ambiguous-module fixture corrected before freezing. 4 applied since review 1.

- `manual-lifecycle-git` / GIT-001–GIT-005: evals `3e82adc` / `93ba9a6` / `5d00cac` (7 genuine failures), findings `98a799f` / `f2ddd0f`, independent passing final audit `5911fab`, impl `92f999a` (8 permanent cases, 2 public smoke probes, 78 app and 53 project cases pass; boundaries pass). Two bounded audit rounds. Archive `0010-manual-lifecycle-git.md`. 1 applied since review 2.

## Latest verification and manual fixture

- Exact checked range:
  `be35552b706aa3000ff8f11306d129ff1d701b85..de65319d646fd0e39e9810415a181dd28911762c`.
  Build, Go test/race/vet, controller check and CI exited 0. CI includes history,
  boundaries and doctor canary recovery; no diagnostics. Coverage: app 78,
  assets 8, contracts 34, project 53 (173 total reported cases including parents).
  Other modules have no specs yet. Logs `.fireweed/evidence/review1-*` and
  `authoring-slice-build.log`; nonfatal module-stat-cache write warning in build.
- Earlier checkpoint at source `91f036c` passed the same current-source checks
  with 86 cases; logs `.fireweed/evidence/checkpoint-*`. Neither checkpoint is
  fresh-checkout evidence or completion of the full design/acceptance gates.
- Hand-authored Go fixture: `/tmp/fireweed-stage2-manual-slice/project` with explicit
  Git identity. NEW modules/new-change ran from packages/toy and created
  changes/0001-answer.md; author commit `95c6a35`. It declares new:answer and
  Answer MUST return 42. NEW stamp answer --json, check --json, show toy --projected
  --json and selected new:answer show all exited 0. Handle remains unassigned;
  author stamping commit `7a34def27b37d5f69aec1eb3abea4ba02ef8fe57`; fixture clean.
- Still unrun: manual lifecycle finish, forbidden-zone/false-fail-first product
  rejections, init/doctor/run and remaining NEW commands, fresh-checkout gates,
  controller/product lifecycle comparison, real drills, isolated NEW self-hosting
  and distribution/delivery. Goal remains incomplete.

## Next / resume — review 2 approved

The gate is cleared by verified shared-base state and explicit human approval.
Review count is one. Keep all bounds and the original stage-2 history baseline.
Continue the existing manual fixture: add only needed Git
and process primitives, hook/JUnit coverage checks, claims/start/verify/done/audit
and apply/archive, wiring each operation through app promptly. Include the two
required product rejections. Do not deepen already-working project/grammar code
before this slice runs end to end. Full project validation/views remain incomplete.

If a change is active on resume, read its context and preserve its work. Protected
author edits belong only between active changes in separate commits.

- Gitstore implementation review exposed selected-new-file staging and index-only
  diff gaps needed by the slice. Both corrected within impl without frozen edits;
  temporary public smoke client `/tmp/fireweed-gitstore-slice-smoke.go` now passes.
  Fixture hooks prepared before start: project-owned Go adapter copied to fixture
  .fireweed/hooks.go, eval/boundaries configured, audit enabled, timeout 30s;
  author commit `5357850929ce87725fc492b1f45b5d1e76b327c9`. No lifecycle success yet.
  Next: checks public JUnit/hook/zone/boundary operations, then workflow and app.
