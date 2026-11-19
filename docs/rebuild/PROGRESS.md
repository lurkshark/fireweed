# Rebuild progress

## Current state (2026-11-15)

The published v0.2.0 commit is [`f42333c`](https://github.com/lurkshark/fireweed/commit/f42333ca824b24a9d0730fa7b3bd4afa2d0ac377).
The binary and repository pin use the 0.2 series. Fireweed applied and
archived all eight changes queued on main at `79ca187` on
`dev/work-through-queued-changes`, through phase commit `764aed3`. Each passed
a separate audit. [PR #13](https://github.com/lurkshark/fireweed/pull/13)
merged the queue and the
[process-group cancellation fix](../../changes/archive/0283-process-group-cancel-race.md)
on main at `21aae46`. The queue is empty and no claims are active.

[0284](../../changes/archive/0284-exclude-reconstruction-paths.md) added
`drill.exclude` for reconstruction inputs. Fireweed applied it at `3b11aa4`
after a separate audit and 620 passing implementation checks. Local main
configures `docs/**` in `drill.exclude` at `3c78321`.

The RUN-062 fixture evals passed with Apple Git 2.30.1 and Git 2.56.1.
The [GC and abandon change](../../changes/archive/0282-git-compatibility-gc-abandon.md)
reproduced both failures and passed 813 implementation checks on Apple Git
2.30.1. Its focused workflow and CLI evals and the full `go test ./...` suite
passed with Git 2.56.1.

[D011](DECISIONS.md#d011--first-project-format) still applies to project
format 1.

[PR #7](https://github.com/lurkshark/fireweed/pull/7) applied the acceptance
changes 0256 and 0258–0272 and the audited 0257 FLOW-098 fixture correction.
[PR #8](https://github.com/lurkshark/fireweed/pull/8) applied the five audited
release fixture changes 0273–0277 and
[0278](../../changes/archive/0278-release-v0-2-0.md). The version change retains
one expected fail-first case, a separate passing audit and 1,349 passing
implementation checks. Earlier maintenance regenerations 0244–0247 and
follow-ups 0248–0255 remain in the [archive](../../changes/archive/).

[ACCEPTANCE](ACCEPTANCE.md#full-completion-gates) records the current gate
status. Gates **1, 3, 4, 5 and 7 remain open**. Gates **2 and 8 are met for
v0.2.0**. Gate **6 retains its eight accepted historical proofs**.

## v0.2.0 release

- Release candidate `5fa703f` and squash commit `f42333c` have identical
  trees. Linux and macOS build, race suite, vet and product check passed in
  [PR CI](https://github.com/lurkshark/fireweed/actions/runs/43634477716)
  and [main CI](https://github.com/lurkshark/fireweed/actions/runs/43635141809).
  The local macOS arm64 candidate passed `make check build` with Git 2.56.1
  and `GOCACHE=/tmp/fireweed-rebuild-cache`.
- The local GoReleaser dry run verified checksums for the four binary
  archives and the source archive, and Ruby accepted the generated
  formula. The installer used local download fixtures on macOS arm64.
  It installed the valid archive, reported version 0.2.0 and rejected a
  corrupted archive before installing it.
- The [Release workflow](https://github.com/lurkshark/fireweed/actions/runs/43635178790)
  succeeded on attempt 2. The first macOS job failed
  `Test_CHK_081_DoctorPersistsCancelledProbeOutput` on timing. The failed-job
  rerun passed on the same tag. The
  [published release](https://github.com/lurkshark/fireweed/releases/tag/v0.2.0)
  includes four binary archives, a source archive, SHA-256 checksums,
  the formula and four native Homebrew bottles. The bottle jobs tested
  installation and version, and the
  [tap formula](https://github.com/lurkshark/homebrew-tap/blob/main/Formula/fireweed.rb)
  points at v0.2.0. Runtime evidence for the Go suite is Linux and macOS;
  four-target bottle evidence covers installation and version.

## Controller and history

The controller for queued changes 0330 and 0331 reports 0.2.4 and was built
with `-trimpath` from clean committed main at
`8ad16e50aa3e2531594f7f76ff86d214c4e069ea` with no active run or claim.

The local main controller reports 0.2.1. It was built with `-trimpath`
from clean committed main at `23d41e6b2b1e543a9d453b3efd86828095c4d3b6`
after PR #24 merged, with no active run or claim. The rebuilt controller
passed `fireweed check` with the existing C18 warnings.

The controller for changes 0312–0317 reports 0.2.1 and was built with
`-trimpath` from clean committed main at
`67067f5aba44e5fbe3e8a677f428efe20e106cb5` with no active run or claim.
Fireweed applied and archived all six changes through `00e0d32` on
`dev/queued-failure-contract-followups`. Each passed a separate audit and
final phase verification. The queue is empty and no claims are active.

The earlier 0.2.1 controller was built with `-trimpath` from clean committed
main at `3c7832135c949a75ced5ef82d04812ceb2480076` with no active run or
claim. The preceding 0.2.0 controller, built from
`21aae46792fe7261acf306a1696b2d39269bae3f`, ran the phases of 0284.

The recorded controller for the acceptance and release phase runs is
`40b33c70f3de176eda058a96b06d2be1437da972`, built with `-trimpath` from a
clean main checkout on 2026-11-13. Its binary reports 0.1.0. The 0.2.0 product
and main’s 0.2-series pin are now in effect. For the release metadata change,
the 0.2.0 controller was built with `-trimpath` from clean committed main at
`089ea0c6830c2754d6c14ce8d43383f849cdd7f5` with no active run or claim.

The controller for the queue on `dev/work-through-queued-changes` was built
with `-trimpath` from clean committed main at
`79ca1873b3c78010e28e82a43419b89d7182545f` with no active run or claim. It
reports 0.2.0.

| Pinned-controller evidence | Result |
| --- | --- |
| `check` on `ee7f2cc` | Passed with five C18 warnings. |
| `check-commits 9f0ca49..ee7f2cc` | 13 commits passed. |
| `check-commits ec8908d..d67f68f` | 58 commits passed on the PR #7 phase branch. |
| `check-commits ee7f2cc..5fa703f` | 20 commits passed on the PR #8 phase branch. |
| `check` on `764aed3` | Passed with five C18 warnings. |
| `check-commits 79ca187..764aed3` | 25 commits passed on the queue branch. |

Main’s new root `9f0ca49` has the same tree as the retained pre-cutover tip
`d99645e`. The required setup baseline `be35552` is an ancestor of
`d99645e`. The older passing range `be35552..e1e1869` covered 1,304
commits and included a passing standalone Doctor and controller CI with
Doctor enabled. Review and verify both sides of the cutover for gate 4.
A proposed split is `be35552..d99645e` and `9f0ca49..HEAD`, retaining
the phase-branch evidence for squashed PRs.

The release phase checks reported B02 for APP-103 (17 cases) and APP-104
(19 cases) and B03 for app’s eval/disposable line ratio (9,046/2,133).
Retain the nonblocking warnings for disposition in the final controller run.

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

Five passing retries used dependent eval feedback under the earlier drill
behavior. [D009](DECISIONS.md#d009--drill-retries-and-the-dependent-signal)
records that limitation, and
[0202](../../changes/archive/0202-drill-independent-signal.md) keeps dependent
failures out of subsequent worker prompts. Review the effect of later changes
on each relevant proof before seeking approval for a repeat.

## Next step

Review the R05 configured-agent Provenance gap and the gate-4 history split.
Proceed with a focused authorized change for the provenance gap, then gather
the final candidate’s plain Go, controller, binary integration and isolated
self-hosting evidence in the order in
[GOAL](GOAL.md). Preserve the existing release and drill evidence and refresh
the portions affected by later changes. Earlier deferred drill observations
are recorded in the [pre-release progress history](https://github.com/lurkshark/fireweed/blob/ec8908d/docs/rebuild/PROGRESS.md#deferred-gaps).
