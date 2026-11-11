# Rebuild progress

## Current state (2026-11-10)

The regeneration milestone is complete. All eight modules have passing native
drill reports, and the owner accepted each proof and deferred its remaining
gaps. Earlier milestone narratives, failed proofs and recovery notes are in
Git history before this revision.

The owner approved changes 0202–0233. Changes 0202–0220 have applied. In 0210,
Fireweed resumed the saved attempt with both CHK-048 assertions intact after
the owner approved marking that ID `already_satisfied`.

The owner approved [the app prerequisite](../../changes/archive/0220-app-history-coded-errors.md)
after the first 0220 run exposed the remaining text-based diagnostic lookup in
`show --history`. Fireweed applied the prerequisite, restored 0220's saved
implementation, and verified all APP-135 cases against the unlabeled project
errors. Changes 0221–0233 remain queued. Links use their archive paths, which
resolve once each applies.

## Controller

`.fireweed/bin/fireweed` was rebuilt with `-trimpath` from clean main after
0207 applied and its claim was released. Go build metadata confirms the
revision and a clean tree.

Controller source: `c5197f710986119ba4370f7e483e7ff4f0b14f37`

On a fresh macOS arm64 clone of the previous controller source
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

The owner deferred these gaps. The queued change that addresses one is linked.

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
  placeholders (CHK-032); boundary rule attribution (CHK-057); canary
  identity (CHK-059); probe output ordering (CHK-065); Doctor fixtures that
  cannot build the root binary in the reconstruction workspace (CHK-061).
- **workflow:** raw boundary-hook output retention (FLOW-052); P08 flags module runs rather than rework chains (FLOW-111;
  [0217](../../changes/archive/0217-rework-chain-warning.md), see
  [D010](DECISIONS.md#d010--p08-warns-on-rework-chains)).
- **runner:** P01 for approval refusals (RUN-041;
  [0205](../../changes/archive/0205-approval-refusal-code.md)); shell quoting for
  already-quoted placeholders (RUN-003); repeated verification on successful
  phases (RUN-014; [0230](../../changes/archive/0230-single-phase-verification.md)).
- **app:** approved-snapshot access when the live queue cannot be parsed
  (APP-004, APP-012); recovering lower-module codes by matching messages
  (APP-037, APP-078, APP-085; 0209–0212 give lower-module errors their
  codes and [0214](../../changes/archive/0214-app-reads-error-codes.md) reads them); Doctor probe names in JSON (APP-067;
  [0229](../../changes/archive/0229-doctor-probe-names.md),
  [0231](../../changes/archive/0231-doctor-json-probe-names.md)); the
  module name and prefix grammar (APP-103).

## Next step

1. Run the remaining changes in queue order with the recorded controller.
2. After 0230 applies, rebuild the controller from clean main with no run or
   claim active and record the new source here. This lets each successful
   phase run full verification once.
3. Await owner direction on the remaining
   [full-completion gates](ACCEPTANCE.md#full-completion-gates-still-required).
   Gate 2 requires the queue to be empty again.

## Suggested follow-ups

Not yet approved; each needs owner agreement before a change is written.

- **Acceptance audit.** Map every R-row in ACCEPTANCE.md to the spec IDs that
  implement it and list rows with no spec. The C17–C19 warnings (0232, 0233)
  were missing although R25 lists them, so other rows may be too. A
  coordinating session or its subagent can do this; it is not worker input.
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
- **App C18 warnings.** After 0233, `check` will warn C18 for two app
  sentences repeated across five and three items ("It MUST leave stderr
  empty." and the null-data failure rule). Decide whether they become
  SYSTEM.md conventions or stay as specs.
