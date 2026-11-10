# Rebuild progress

## Current state (2026-11-10)

WF1 is complete. Fireweed applied and archived both changes after separate
audits with no findings and passing implementation verification:

- [0195-run-argv-contract](../../changes/archive/0195-run-argv-contract.md)
  (`eb4a7f2`) documents `gitstore.Request.Args[0]` as the executable and
  adds regression coverage. All 925 cases passed.
- [0196-drill-self-checked-reconstruction](../../changes/archive/0196-drill-self-checked-reconstruction.md)
  (`084fbd4`) gives the drill worker a buildable input workspace with
  dependency implementations and its own eval hook. All 485 cases passed.

Fireweed released both claims after applying the changes.

WF2 is complete. `go test ./...`, `go vet ./...` and `fireweed ci` passed
on the WF2 controller source commit `6cc2629`. The first Go test run alongside CI
failed a timed app eval, which passed on the standalone retry. CI ran with
Doctor enabled in a temporary local clone over `65bf933..6cc2629`.

WF3 is complete. With owner approval, the recorded controller ran
`fireweed drill workflow` once at `4754026` with the configured Codex worker.
The drill passed after its configured retry. All 666 cases passed (181
workflow, 323 app, 162 runner) and the worker reported no contradicted
Decisions. Fireweed cleaned up the scratch worktree and retained reconstruction
`38aa8fc` under `refs/fireweed/drills/workflow/20261109T224519.080152000`.
The [native report](../../.fireweed/drills/workflow.md) is committed in `c62b5be`.

WF4 is complete. The owner reviewed the report and approved these dispositions:

- **FLOW-110 / APP-126.** Closed. APP-126 specifies P01 for an unavailable
  base and its eval passed on the drill retry.
- **FLOW-052 / checks.** Deferred raw boundary-hook output retention. The
  current boundary API returns diagnostics only.
- **FLOW-027 / checks.** Fireweed applied and archived
  [0197-boundary-hook-nonzero-exit](../../changes/archive/0197-boundary-hook-nonzero-exit.md)
  in `d45652a`. CHK-043 now rejects a nonzero exit without a module-attributed
  violation. Both regressions failed first, the separate audit found no
  gaps, and all 787 implementation cases passed. `fireweed check` and
  `fireweed check-commits a0e25a2..d45652a` passed. The queue is empty and
  Fireweed released the claim.

The owner approved the `0197` worker and restricted audit contexts for OpenAI
Codex and granted standing authorization for the configured change lifecycle.
[SYSTEM.md](../../SYSTEM.md#agent-configuration) records the scope and destination.

## Controller and candidate

The installed `.fireweed/bin/fireweed` was rebuilt with `-trimpath` from clean
main with an empty queue and no claims. Go build metadata confirms the
revision and a clean source tree.

Controller source: `e1e1869131833778f1f8f1a23cf62b0d3b3869b3`

With `GOCACHE=/tmp/fireweed-rebuild-cache`, the fresh macOS arm64 candidate
clone passed the product build, `go test ./...`,
`go test -race ./...`, `go vet ./...`, `fireweed check`, and standalone
`fireweed doctor` including canaries, recovery and both agent probes.
`fireweed check-commits` passed 1,304 commits over
`be35552b706aa3000ff8f11306d129ff1d701b85..e1e1869131833778f1f8f1a23cf62b0d3b3869b3`.
CI with Doctor enabled over the same range passed on the full retry. CI
verified passing evals and coverage for app, assets, checks, contracts,
gitstore, project, runner and workflow. The first CI run failed on the
deadline-sensitive `Test_RUN_036_RUN_037_OneBudgetAcrossChangesRetainsApplied`
eval, which passed in isolation before the full retry.

## Drill proof inventory

**Contracts proof accepted.** The verified controller ran
`fireweed drill contracts` at `e7f1554`. All 1,119 cases passed on the first
attempt, including 64 contracts cases and 1,055 transitive dependent cases.
The [native report](../../.fireweed/drills/contracts.md) is committed in
`24f9d94`. Fireweed cleaned up both scratch locations and retained
`30c7cdfd356d7068f841179c9c4a1fb1321be6c4` at
`refs/fireweed/drills/contracts/20261110T034341.899515000`.
The proof changes only disposable contracts paths.

The report lists no contradicted Decisions. The owner deferred all three gaps:

- `CON-002` / `CON-009` leave fenced headings and level-one section boundaries
  implicit. The reconstruction handles these differently from the current code.
- `CON-008` leaves equality for YAML NaN and offset timestamp keys ambiguous.
- `CON-017` through `CON-019` specify codes and, for duplicates, a subject.
  The evals also pin error wording or prefixes.

**Renewed gitstore proof accepted.** The verified controller
ran `fireweed drill gitstore` at `c3863e5`. All 933 cases passed on the first
attempt, including 13 gitstore cases and 920 transitive dependent cases. Both
GIT-010 expected-absence forms passed. The
[native report](../../.fireweed/drills/gitstore.md) is committed in `c687766`.
Fireweed cleaned up both scratch locations and retained
`d3230d81b89f327ec5a4dfabfcd6381951d7b235` at
`refs/fireweed/drills/gitstore/20261110T060020.420365000`. The proof changes
only disposable gitstore paths.

The owner-approved GIT-010 changes
[0198](../../changes/archive/0198-casref-empty-expected-delete.md) (`feabbc7`)
and [0199](../../changes/archive/0199-casref-zero-expected-delete.md) (`3fa1918`)
passed fail-first, separate audits, and full verification. The earlier report
at `0b34909` and proof `3382ac8` remain in history and at the earlier drill ref.
That reconstruction failed the explicit-zero eval added after its drill.

The renewed report lists no contradicted Decisions. The owner deferred its
remaining gaps:

- `GIT-001` and `GIT-003` leave unborn HEAD and precedence for overlapping
  tracked edits implicit.
- `GIT-005` leaves environment inheritance and executable PATH lookup implicit.
- `GIT-007` leaves literal path syntax and index state after a failed commit
  implicit. Lifecycle callers use the `:(literal)` spelling.
- `GIT-008` through `GIT-010` leave abbreviated blob IDs, ref naming, and
  symbolic-ref lookup and mutation implicit.

**Renewed assets proof accepted.** The verified controller ran
`fireweed drill assets` at `34f9f54`. All 1,049 cases passed on the configured
retry, including 125 assets cases and 924 transitive dependent cases. All four
Doctor cases that failed in the earlier drill passed. The
[native report](../../.fireweed/drills/assets.md) is committed in `db34f44`.
Fireweed cleaned up both scratch locations and retained
`fe1a898a1944436fa7fc066c4cf3781f26ef43cb` at
`refs/fireweed/drills/assets/20261110T080441.699697000`. Its changes are
confined to disposable assets paths.

The earlier failed report remains in `84f7138` and reconstruction `69b7d7c`
remains at `refs/fireweed/drills/assets/20261110T063418.808553000`. It passed
125 assets cases and failed four dependent Doctor cases on `TestMissingIDCanary`.
The owner deferred its other gaps on baseline coverage, mixed Go package zones,
expected-failure JUnit status, and static import resolution.

Fireweed applied [0200](../../changes/archive/0200-doctor-go-missing-id-canary.md)
in `f854288`. CHK-059 now matches the supplied Go test declaration and its
subtests, including names without an underscore after `Test`. The change
passed fail-first, a separate audit with no findings, and 791 implementation
cases. All four previously failing Doctor evals passed with the retained
assets reconstruction `69b7d7c` and the current checks code in a temporary clone.

The owner accepted the proof and approved proceeding to project. The renewed
report lists no contradicted Decisions. The retry restored the
`# Write evals` heading expected by workflow and app, which the current product
already supplies. The report identifies the heading as implicit in AST-002
and the missing-ID handoff to Doctor as implicit in AST-058. The remaining
observations revisit the deferred mixed-zone, JUnit and import-resolution gaps.

**Project proof renewal awaits the PRJ-069 audit.** The verified controller ran
`fireweed drill project` at `d4e8d95`. All 133 project cases and 790 dependent
cases passed after the configured retry. One app case failed. The
[native report](../../.fireweed/drills/project.md) is committed in `dfbb207`.
Fireweed cleaned up both scratch locations and retained the reconstruction
`917fb02dba3b227f5c57ed422455e7e5881735c7` at
`refs/fireweed/drills/project/20261110T144729.916008000`. Its changes are
confined to disposable project paths.

`Test_APP_116_APP_117_ListUsageConfigurationAndQueueErrors` requires a C03
error from `project.Check` when `changes` is a regular file. The reconstruction
classified the queue read failure as C13. The same focused app eval passed
against the current product. The project queue and Check specs leave this
diagnostic implicit.

The owner approved making C03 explicit. Fireweed started
[0201-queue-path-diagnostic](../../changes/0201-queue-path-diagnostic.md)
and assigned `PRJ-069`. The project eval passed with the other 133 cases.
Audit found that the change also promises C03 on the public `Queue` error,
whose classification is private today. The eval asserts the path on `Queue`
and the C03 diagnostic on `Check`. App derives C03 from that diagnostic.

Both auditor submissions omitted the required dash before the finding's
explanation. The controller rejected them with P09 and blocked `0201` in
`d5ea3be`. It saved the audit work at
`refs/fireweed/attempts/queue-path-diagnostic/1` (`dc223bd`) and released the
claim. Audit rounds remain at zero.

The owner approved narrowing `PRJ-069` to the behavior the existing eval
asserts. `Queue` rejects a regular-file `changes` path and identifies it in
the error. `Check` reports a non-warning C03 diagnostic identifying the same
path. Author commit `fde499d` made the revision and removed `blocked` using
the documented resume procedure, preserving the assigned ID, run and audit
phase. `fireweed check` and the focused `PRJ-069` and `APP-116`/`APP-117` evals
passed. Automatic approval review rejected the lifecycle command before it
ran because it required specific authorization to send repository context to
OpenAI Codex. No claim is active. The saved attempt and failed drill ref remain
available.

The report lists no contradicted Decisions. Its other observations cover
diagnostic codes and wording relied on by app, duplicate C08 messages,
change numbering past 9999, intermediate graph validation, and stale-view
and history behavior. It also identifies the design's warning for specs
outside configured globs as an unassigned contract.

**Reuse workflow.** The retained ref and commit `38aa8fc` match the committed
passing report and change only workflow's disposable files. Since that report,
`0197` changed the checks dependency. The retained workflow reconstruction
passed the `FLOW-027` and `FLOW-062` evals and the app CI boundary-failure eval
against the current candidate in a temporary clone.

**Two historical passing refs remain missing.** The prior passing reports
remain in Git history.
The following refs under `refs/fireweed/drills/<module>/` are absent from this
checkout and GitHub origin advertises none. Their exact commits were unavailable
across 119 local object stores, including the historical candidate directories.

| Module | Ref suffix | Expected commit |
| --- | --- | --- |
| project | `20261102T031647.104747000` | `481bc00c1da360d3e8bf76cd5ba04053d2dccaeb` |
| checks | `20261107T002339.020159000` | `c0ca0397e3781410e795f27b52de2b2105350559` |

Reassess reuse if an exact historical proof is recovered from a backup.
Checks needs current proof of the `0197` CHK-043 change. `runner` and `app`
need first native reports and retained refs.

## Next step

Obtain explicit authorization for the configured Codex audit, implementation
and project reconstruction contexts. Resume `0201` at audit with the verified
controller, then renew `project` after Fireweed applies it. Commit the native
report, retain the proof ref and stop for owner review before `checks`. The
remaining drill order is `project`, `checks`, `runner`, then `app` in the
[regeneration goal](GOAL.md).

Deadline set by the owner: **2026-11-10 22:51:05 UTC**.
