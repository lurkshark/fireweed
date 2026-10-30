# Abandon CLI slice — 0110/0111

NEW at clean source f82019a4f8368a639231d8075f917be398bb7957 proves the design's
abandon path in an actual Go fixture. Go1.27.2 darwin/arm64, CGO disabled,
vcs.modified=false. No bootstrap invocation supplies product behavior.

0110's FLOW-089–096 dependency, fresh-input recovery and402 full passes are
recorded in [the recovery report](recovery-0109.md). 0111 APP-097–099 phases:
8a8997e/10ac6dd/98241c6/0c734ac/f82019a. Both eval rounds217 cases7 failing/
210 passing; final full implementation217 passing,8 focused test events pass.
Restricted audit found nearest-root selection unasserted; one existing success
fixture now nests a distinguishable project inside another, retaining outer
change/HEAD assertions. Native round1 persisted with same run/IDs; isolated
APP-097-only re-audit PASS. Complete auditor command inspection confirms only
allowed public snapshot reads, no implementation/SYSTEM/history/outside/code
execution/writes; snapshots removed. Real workers used the fixed Sol/high CLI.
Only app internal dispatch/output changes implement the caller; public API and
all earlier assertions remain. Tree clean/unclaimed after native archive.

## Actual NEW evidence

Fixture baseline dbb1a6c28f72ec1a479a1df6d0891b90f2728b89; next run
20261030T003457-01882ca346821f1a0bb0dc45 records actual Go ALP-002 fail-first,
separate bounded fake audit and two intentionally wrong implementation attempts.
Four agent calls total; exit1, committed blocked phase impl, clean checkout and
released claim. Real saved ref refs/fireweed/attempts/value-alpha/1 at
 a224092450abd11c3cb413ff4b981871c25805ba remains unchanged after abandon.
Abandon creates exactly one commit7efe6c14df94f591c030c07cb0c128efcc1980d1 over
6718e876e214c9a307b1e26c96a1ef1e1681153e, returns schema1 confirmation string
with that full OID/empty diagnostics/stderr, restores baseline business files,
and deletes only the pending change relative to baseline. No new phase value,
claim or additional agent call. New-change then Start reuses ALP-002; release
removes its claim while expected uncommitted Start state remains. Fixture removed.

Two harness setup corrections are explicit: the wrapper initially assumed an
unset FIREWEED_PHASE and failed before evals. NEW released that fixture's own
claim; only its uncommitted Start draft was restored, the wrapper was changed
to derive phase/module from the pending document like the existing fixture,
and setup committed in the same fixture. The corrected next then succeeded at
the intended exhaustion. Block correctly removed an empty nested working
directory; the harness continued that same blocked state from the tracked
module directory without repeating next or altering product behavior/assertions.
This is fixture repair, not product correction or a phase-write exception.

Local details: /tmp/fireweed-abandon-cli-{result.json,commands.json,run.log,
corrected-run.log,finish.log}; /tmp/fireweed-0111-{build-info.txt,impl-done.log,
evals-worker.log,audit-worker.log,evals-correction-worker.log,reaudit-worker.log,
impl-worker.log}. No new drill/self-hosting/final acceptance claim. Count5 now
requires separate quality/goal review and fresh checks before push/continuation.
