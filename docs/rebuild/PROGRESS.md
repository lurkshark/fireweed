# Rebuild progress

Keep this file short: update sections in place and add one line per applied change.

## Status — 2026-10-27

The NEW binary completed the real manual Go lifecycle at clean source
`89b70289ee562d228e055f047d48470057beb607`: fail-first, both required
rejections, separate audit, frozen implementation, apply/archive and coverage.
Four changes have applied since human review 2. The full goal remains incomplete.

Current local `origin/main` is `c5c02de31bec8eb22569269f25d0ef0e5a69e63a`;
controller status reports 1 unapproved archive. This ref observation does not
restart the independent human-review count. No active rebuild claim remains.
Next: Go setup assets for init, then the five-change human checkpoint.

## Controller, source and reviews

- Controller: `.fireweed/bin/fireweed-bootstrap`, copied tested executable 0.1.0,
  source `7f925190878e38fc38f30d16bb1cf47aca098cf7`, vcs.modified=false;
  pin SHA-256 `f1516b5686435c23b54b71c5ac554aba6859b5454f072e0c5eb031dfd6f0fd4f`.
  It remains unchanged. Never rebuild it from or consult bootstrap source/tests.
- Stage-2 setup baseline: `be35552b706aa3000ff8f11306d129ff1d701b85`.
  Use explicit nonempty baseline..HEAD ranges; never the historical d991c26 range.
- Product: `bin/fireweed`, built at clean source
  `89b70289ee562d228e055f047d48470057beb607` (vcs.modified=false).
  help/version/explain/playbook/modules/new-change/stamp/check/show, start/verify/done,
  held context and coverage work. Remaining commands/options explicitly fail.
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

- `lifecycle-checks` / CHK-001–CHK-005: evals `69933f7` (25 failures, 3 passes), independent audit `eff6823` (no findings), impl `3c5b40b` (28 own, 78 app and 4 public smoke probes pass; dependents/boundaries pass). Two contradictory fixtures corrected before freezing; unknown-ID/protected-path gaps corrected within impl. Archive `0011-lifecycle-checks.md`. 2 applied since review 2.

- `manual-lifecycle` / FLOW-001–FLOW-005: evals `c5c02de` / `ca11056` / `aec7154` (28 failures, 2 passes), findings `0b5aeef` / `25e4775`, independent passing final audit `8627014`, impl `563d2ee` (32 own, 78 app and 3 public probes pass; dependents/boundaries pass). Two bounded audit rounds. Sol usage limit interrupted the first auditor; retry after its reported reset completed independently. Approved-archive counting, mutable Decisions and frozen audit-state gaps corrected within impl. Archive `0012-manual-lifecycle.md`. 3 applied since review 2.

- `lifecycle-cli` / APP-004, APP-009–APP-012: evals `678f7bb` / `a9d195c` (8 failures, 78 passes), findings `7db8c47`, independent passing re-audit `804cb84`, impl `89b7028` (86 app cases, 2 held-coverage public probes and boundaries pass). One audit round; fixture ID setup corrected before freezing and current-coverage metadata corrected within impl. Archive `0013-lifecycle-cli.md`. 4 applied since review 2.

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
- Manual slice verified through NEW only in `/tmp/fireweed-stage2-manual-slice/project`.
  Start kept HEAD `5357850`, assigned TOY-001 and run
  `20261027T144429-53d4980e0361844d890dea9e`. Real Go `Test_TOY_001_Answer`
  failed at 0, passed at 42. SYSTEM edit rejected Z01; passing-before impl rejected
  P2-01; verify never committed. Separate Sol audit passed. Native phase commits:
  evals `f04e9f71a8d175d50f87c33f68b35cc61a059dbe`, audit
  `bfb84d0cc80b313b589929c3a2215a0876111364`, apply
  `1323a31dd1605bf86393b1fb8b4ef7c5c799dd8e`. Eval/facade Git diff against
  evals commit is empty. SPEC/archive/provenance inspected, coverage TOY-001 pass,
  fixture clean/no claim. Logs `.fireweed/evidence/manual-fixture-*`.
- Still unrun: init/doctor/run and remaining NEW surface, fresh-checkout gates,
  controller/product comparison, real drills, isolated NEW self-hosting and
  distribution/delivery. All complete acceptance rows remain unverified.

## Next / resume — review 2 count four

Milestone 1 is verified. Build the Go setup bundle needed by init next, preserving
all bounds. Stop after that fifth applied change for human review even if the
shared-base counter is lower. Init/doctor and bounded unattended execution are
next; do not deepen already-working manual behavior.

If a change is active on resume, read its context and preserve its work. Protected
author edits belong only between active changes in separate commits.
