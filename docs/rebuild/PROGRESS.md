# Rebuild progress

## Current state (2026-11-13)

Main is `3b3b64d`, including merged regenerations and follow-ups 0244–0255.
The owner requested an R01–R29 traceability review and drafts only. The completed
[mapping in ACCEPTANCE](ACCEPTANCE.md#traceability-review--2026-11-13) identifies
existing spec/eval/archive support, confirmed gaps and remaining acceptance
evidence. The revised queue has 16 drafts: 0256 and 0258–0272. The owner
requested removal of the unsupported nested-root restriction (0257), narrowing
of compatibility and override contracts, full Removed-section validation, and
a runner follow-up retaining Start warnings (0272). The remaining drafts cover
initialization, zones, document grammar, targets, guidance and audit attempts.
None has started; no claims, phase commits, agents or drills were created for
this queue.

The review worktree uses a copy of the clean controller built from `40b33c7`,
matching the recorded pin below. D011 records the owner's clarification that
format 1 has no predecessor; future format compatibility/migration remains due.
The historical gate-4 baseline still needs reconciliation with the new main
root, not an unreviewed baseline substitution.

Draft validation: pinned-controller check passes with five nonblocking C18
warnings (four pre-existing and one involving compatibility text). The 16
drafts contain 29 proposed spec items within the shape limits; all R01–R29
mapping rows, local document links and whitespace checks pass. No draft has
run/phase fields or Provenance. Context and verify stop with P02 because no
change is held; no phase verification, agents or acceptance exercise ran.
Do not start a draft just to obtain claim-bound context.

## Completed regeneration follow-ups

Fireweed applied and archived changes 0248–0255 on
`dev/queued-regeneration-followups` at product revision `b220afc`. The queue
is empty and all claims are clear. All required phase checks passed, including
separate audits for 0248–0254. Change 0255 removed the drill cleanup permission
precheck and passed 539 runner and app cases, including the repaired RUN-062
fixture. No standalone drill ran.

Controller `check` passed with C18 repeated-sentence warnings. Phase
verification reported app B02 warnings for APP-103 and APP-104 and a B03
warning for the app eval-to-disposable line ratio (8,690/2,105).

## Core module regeneration

The four owner-approved regenerations are complete on
`dev/regenerate-core-modules` at product revision `466a0ba`, based on clean
main `a74c362` containing prerequisite PR #2 and progressive-help PR #1.
Fireweed removed each module's disposable implementation, then verified,
committed and archived its replacement in dependency order.

| Module | Archived change | Applied commit | Passing cases, including dependents |
| --- | --- | --- | --- |
| project | [0244](../../changes/archive/0244-regenerate-project.md) | `02f7414` | 1,025 |
| workflow | [0245](../../changes/archive/0245-regenerate-workflow.md) | `6bb22fa` | 739 |
| runner | [0246](../../changes/archive/0246-regenerate-runner.md) | `9809d17` | 539 |
| app | [0247](../../changes/archive/0247-regenerate-app.md) | `466a0ba` | 369 |

Each archived change retains its implementation Provenance. These maintenance
regenerations followed the normal implementation-only lifecycle with audit
configuration and all enforcement limits unchanged. Specs, retained interfaces,
evals, unrelated modules and pre-existing archives are unchanged, including
APP-156–APP-158 help and APP-159/APP-160 prerequisite contracts. No drill ran.
The controller stayed pinned to `a74c362` throughout those regenerations.

`make check build` passed all race-enabled Go tests, vet and the product build.
Controller static and history checks passed over `a74c362..466a0ba`, and full
controller CI over that range passed, including every module's evals, boundaries
and doctor canaries with agent probes disabled. The initial CI attempt reported
`Test_CHK_003_CHK_031_CHK_032_CHK_033_CHK_034_CHK_035_ModuleOverrideArgvAndOutput`;
the isolated rerun and full CI retry passed without code, test or limit changes.

Nonblocking warnings remain: three C18 repeated-clause warnings, B02 for
APP-103 (17 cases) and APP-104 (19), and B03 for app's eval/disposable line
ratio of 8,647/2,105, slightly above the limit of 4. Earlier drill proofs and
owner-deferred gaps remain below; no new drill evidence is claimed.

## Controller

`.fireweed/bin/fireweed` was rebuilt with `-trimpath` on 2026-11-13 from a
clean checkout of main's current revision before processing changes 0248–0255.
No run or claim was active. Go build metadata confirms the revision and a clean
tree.

Controller source: `40b33c70f3de176eda058a96b06d2be1437da972`

On a fresh macOS arm64 clone of an earlier controller source
(`e1e1869131833778f1f8f1a23cf62b0d3b3869b3`), with
`GOCACHE=/tmp/fireweed-rebuild-cache`, the product build, `go test ./...`,
`go test -race ./...`, `go vet ./...`, `fireweed check` and standalone
`fireweed doctor` passed, including canaries, recovery and both agent probes.
`fireweed check-commits` passed 1,304 commits over
`be35552b706aa3000ff8f11306d129ff1d701b85..e1e1869131833778f1f8f1a23cf62b0d3b3869b3`.
`fireweed ci` with Doctor over the same range passed on a full retry; the first
run failed the timing-sensitive
`Test_RUN_036_RUN_037_OneBudgetAcrossChangesRetainsApplied`, which
[0207](../../changes/archive/0207-time-budget-eval-margin.md) addresses.

## Drill proofs

Each proof ran `fireweed drill <module>` with the recorded controller and the
configured Codex worker, cleaned up both scratch locations, and changed only
the module's disposable paths. No report lists a contradicted Decision.

| Module | Drill base | Passed on | Cases (own + dependent) | Report commit | Retained ref under `refs/fireweed/drills/` |
| --- | --- | --- | --- | --- | --- |
| contracts | `e7f1554` | attempt 1 | 64 + 1,055 | `24f9d94` | `contracts/20261110T034341.899515000` (`30c7cdf`) |
| gitstore | `c3863e5` | attempt 1 | 13 + 920 | `c687766` | `gitstore/20261110T060020.420365000` (`d3230d8`) |
| assets | `34f9f54` | retry | 125 + 924 | `db34f44` | `assets/20261110T080441.699697000` (`fe1a898`) |
| project | `73e269d` | retry | 134 + 791 | `acc1c26` | `project/20261110T164427.736325000` (`0b51fc1`) |
| checks | `9e81519` | retry | 125 + 666 | `8369532` | `checks/20261110T174227.798503000` (`3f0279d`) |
| workflow | `4754026` | retry | 181 + 485 | `c62b5be` | `workflow/20261109T224519.080152000` (`38aa8fc`) |
| runner | `5b59f35` | retry | 162 + 323 | `9390435` | `runner/20261110T192409.383763000` (`521e048`) |
| app | `2ab6c98` | attempt 1 | 323 + 0 | `dd7d94b` | `app/20261110T201533.930973000` (`b6ace52`) |

The workflow proof was reused after `0197` changed checks: its retained
reconstruction passed the `FLOW-027`, `FLOW-062` and app CI boundary-failure
evals against the later candidate in a temporary clone.

Drill refs are local and never pushed, so they exist only in the clone where
the drills ran. Two earlier passing refs (project `20261102T031647.104747000`,
checks `20261107T002339.020159000`) were lost this way; the renewed proofs
above supersede them. Drills can be repeated, so the committed reports are
the evidence kept; no copy of the reconstructions is stored in the
repository.

**Retries used dependent feedback.** Every retry above received the
dependents' failing eval output and changed code to match it, so those five
passes show the spec plus dependents' evals were enough, not the spec alone.
[DECISIONS D009](DECISIONS.md#d009--drill-retries-and-the-dependent-signal)
records the clarification, and
[0202](../../changes/archive/0202-drill-independent-signal.md) changes the drill so
that dependent failures are reported, not fed back.

## Deferred gaps

The owner deferred these gaps or addressed them through the linked changes.

- **contracts:** fenced headings and level-one section boundaries (CON-002,
  CON-009; [0226](../../changes/archive/0226-fenced-section-boundaries.md)); YAML NaN and offset-timestamp key equality (CON-008); evals that
  pin error wording beyond the specified codes (CON-017 to CON-019;
  [0227](../../changes/archive/0227-item-error-wording-evals.md)).
- **gitstore:** unborn HEAD and overlapping tracked edits (GIT-001, GIT-003);
  environment and PATH lookup (GIT-005); literal path syntax and index state
  after a failed commit (GIT-007); abbreviated blob IDs, ref naming and
  symbolic refs (GIT-008 to GIT-010).
- **assets:** the `# Write evals` heading other modules' evals expect
  (AST-002; [0206](../../changes/archive/0206-context-playbook-eval.md),
  [0208](../../changes/archive/0208-static-cause-evals.md)); mixed-zone Go packages,
  expected-failure JUnit status and static import resolution (AST-018,
  AST-054, AST-056, AST-057).
- **project:** C13 for duplicate slugs (PRJ-024;
  [0203](../../changes/archive/0203-duplicate-slug-code.md)); app evals that pin
  project wording (PRJ-021, PRJ-031, PRJ-041, PRJ-044;
  [0208](../../changes/archive/0208-static-cause-evals.md)); C01 and U02 carried both
  as metadata and as message prefixes (PRJ-004, PRJ-040, PRJ-056, PRJ-058,
  PRJ-066; [0209](../../changes/archive/0209-project-coded-errors.md)); duplicate C08 diagnostics ([0228](../../changes/archive/0228-absent-edit-diagnostic-evals.md)), partial projections, numbering past
  9999, intermediate graph validation, stale-view history, and the design's
  warning for specs outside configured globs
  ([0224](../../changes/archive/0224-spec-outside-globs-explanation.md),
  [0225](../../changes/archive/0225-spec-outside-globs-warning.md)).
- **checks:** which testcase-name tokens count as IDs (CHK-028, CHK-080;
  [0204](../../changes/archive/0204-junit-label-tokens.md)); shell quoting around hook
  placeholders (CHK-032, CHK-082;
  [0242](../../changes/archive/0242-hook-placeholder-quoting.md)); boundary rule attribution (CHK-057); canary
  identity (CHK-059); probe output ordering (CHK-065); Doctor fixtures that
  cannot build the root binary in the reconstruction workspace (CHK-061).
- **workflow:** report-retention scope (FLOW-052;
  [0241](../../changes/archive/0241-eval-report-retention-scope.md) narrows it to
  module reports and eval-hook output); P08 flags module runs rather than rework chains (FLOW-111;
  [0217](../../changes/archive/0217-rework-chain-warning.md), see
  [D010](DECISIONS.md#d010--p08-warns-on-rework-chains)).
- **runner:** P01 for approval refusals (RUN-041;
  [0205](../../changes/archive/0205-approval-refusal-code.md)); shell quoting for
  already-quoted placeholders (RUN-003;
  [0243](../../changes/archive/0243-agent-placeholder-quoting.md)); repeated verification on successful
  phases (RUN-014; [0230](../../changes/archive/0230-single-phase-verification.md)).
- **app:** approved-snapshot access when the live queue cannot be parsed
  (APP-004, APP-012;
  [0239](../../changes/archive/0239-approved-snapshot-query.md),
  [0240](../../changes/archive/0240-approved-cli-metadata.md)); recovering lower-module codes by matching messages
  (APP-037, APP-078, APP-085; 0209–0212 give lower-module errors their
  codes and [0214](../../changes/archive/0214-app-reads-error-codes.md) reads them); Doctor probe names in JSON (APP-067;
  [0229](../../changes/archive/0229-doctor-probe-names.md),
  [0231](../../changes/archive/0231-doctor-json-probe-names.md)); the
  module name and prefix grammar (APP-103, APP-159;
  [0238](../../changes/archive/0238-module-identity-grammar.md)).

## Next step

Owner review of the revised 16 drafts (0256, 0258–0272), then separate
direction for execution if desired. The current request explicitly forbids
running them. Sequence 0263 → 0264 → 0265 and 0262 → 0272; coordinate 0268
with its diagnostic-producing changes before closing their acceptance gaps.
ACCEPTANCE retains the binary next/run warning check and other open evidence.
Do not run drills or the final completion gates as part of drafting.

## Suggested follow-ups

Not yet approved; each needs owner agreement before a change is written.

- **Acceptance closure.** The R01–R29 mapping is complete. Review the drafted
  fixes and the explicitly open runtime/history/delivery evidence in ACCEPTANCE;
  the mapping alone does not close those gates and is not worker input.
- **GIT-007 failed-commit state.** The gitstore drill found that a failed
  commit can leave intent-to-add entries for selected new paths. Probe the
  current code: if a failed lifecycle commit leaves the index changed, the
  next step fails with P01 for a dirty worktree until someone resets it.
  Write a fix only if the probe reproduces it.
- **Flaky timed app eval.** One CI run had a second timing failure besides
  RUN-036/RUN-037 that was never identified. Run the app evals repeatedly
  (`go test -count=20 ./packages/app/...`) to find it. If the work simply
  takes longer than the timer allows, ease the bound, as 0207 does, so that
  timers only catch extreme outliers.
- **App C18 warnings.** The product's `check` warns C18 for two app
  sentences repeated across eight and three items ("It MUST leave stderr
  empty." and the null-data failure rule). Decide whether they become
  SYSTEM.md conventions or stay as specs.
