# Checkpoint corrections 0076–0078 — PASS

Current result: the authorized scoped subagent re-review and fresh validation
passed. Product head is 0f6f7b2; the clean verification/build source is a34a32a.
The original audit deviation remains historical below. The owner-authorized
correction push may proceed; the full rebuild goal remains incomplete.

## Historical correction and review record

The owner-approved initial push synchronized origin/main at ce165ab with the
0071–0075 findings recorded. Normal checks correction 0076 applied at 995eebb:
evals e2b5b3b reported 298 cases, three runtime failures and 295 passes; audit
37fc2e9 mechanically advanced; implementation 995eebb passed all 298 phase cases.
Evals and public interfaces remained frozen through implementation. No claim or
pending change remains. The runner eval_fix has not started, and no correction
push has occurred.

## Audit input/phase deviation

Native audit-1 read checks and gitstore implementation and edited doctor.go.
The controller rejected that disposable edit with P2-01. Audit-2 read its diff,
restored doctor.go, and committed No findings for CHK-051/055/056/057. Restoration
removed the forbidden edit, but neither invocation is valid restricted audit
input evidence. Their actual logs remain under run
20261028T213607-f7bde2247258cdc46a0653c9; do not consult the provisional code as an
implementation oracle. No archived change or phase history was rewritten.

One supplemental fixed-profile Sol/high assertion reviewer received an isolated
snapshot containing the pre-implementation frozen change/evals, checks specs,
public interfaces and governing rules only. Its log shows only supplied-file
reads, without implementation/history access, execution or writes. The snapshot
was removed; supplemental-assertion-audit-1.log remains in the run directory.
This was a post-application review and does not retroactively validate the native
phase audit.

## Remaining assertion finding

The supplemental reviewer checked CHK-051/055/056/057 and found one gap in
CHK-055: the clause “Doctor MUST require passing baseline evals, coverage and
applicable boundaries before planting canaries” lacks an assertion that failed
baseline checks prevent planting. Existing assertions check C14 for baseline
failure; a program reporting that error and still planting could satisfy them.
The other three reviewed IDs had no reported assertion finding. No further
correction was attempted after this unresolved result.

## Proposed bounded continuation

Explicit approval is needed to add one normal checks eval_fix for CHK-055,
retaining all prior assertions/coverage and adding the missing stop-before-planting
assertion. Product specs and interfaces stay unchanged. Then perform the already
approved runner eval_fix for RUN-007/RUN-009 and obtain the same checkpoint
subagent's scoped re-review before any correction push. Preserve counters,
history, bounds and the existing three-minute default test. To enforce input
isolation, coordinate native start/context/done with real fixed-profile CLI
workers and isolated auditors; controller commits remain authoritative and
coordinator advancement provenance must honestly say manual.

Any further unresolved finding or block stops. No new phase-write exception,
bound change, drill or deferred contracts-gap fix is proposed. The separately
received documentation-refocus proposal remains unapplied in /tmp and must be
reconciled with this newer state at a future safe author boundary.

## Additional eval_fix explicitly approved — 2026-10-28

The owner approved one CHK-055 eval_fix, completion of the existing runner
eval_fix and scoped subagent re-review, using isolated auditors and pushing only
on PASS. Use native controller phase advancement with fixed-profile CLI workers
and auditors; audit snapshots contain only permitted inputs. Record coordinator
advancement honestly. No product-spec/interface, history, counter or bound change
is authorized. The prior audit deviation remains part of the record.

## CHK-055 eval_fix passed

0077 committed evals 064d9ab, isolated audit 0fd7393 and application 435b176.
Both full verification phases passed 298 cases. The real Sol/high eval worker
strengthened five existing baseline cases; source/interfaces/specs stayed frozen.
The separate Sol/high auditor read only an isolated snapshot and returned No
findings, Checked CHK-055. Log inspection confirms no implementation/history
access, execution or writes. The temporary snapshot was removed. Coordinator
phase advancement honestly records agent=manual; actual CLI logs remain under
run 20261028T230031-7c37d81c38953bf7362a2c42. This closes the missing assertion
gap while retaining the earlier native audit deviation as historical fact.
B02 reports 16 cases for CHK-055, one above the aim; all cases/assertions are
retained and this correction adds no cases. Runner correction and final scoped
checkpoint re-review remain pending; no push after ce165ab has occurred.

## Runner eval_fix passed phase checks

0078 committed evals 88119c5, isolated audit 7b8a843 and application 0f6f7b2.
Both full phase checks passed 178 cases. The one existing test now uses a
3-second fake timeout, 5-second descendant delay, 4-second bounded startup/return
windows and a 5.2-second marker observation; stdout/stderr precede readiness.
Both modes, all assertions and IDs remain. The edited parent passed with -race.
The isolated Sol/high auditor returned No findings for RUN-007/RUN-009; complete
log inspection shows only permitted snapshot reads and no writes/execution or
implementation/history access. The snapshot was removed. No product source,
interface/spec or controller limit changed. Coordinator provenance is manual,
with actual fixed-profile CLI logs retained in run
20261028T231208-cfb868c63baddc1bfdad0dda. Fresh full race and scoped checkpoint
re-review remain pending; no push after ce165ab has occurred.

## Final scoped review and verification — PASS

The same checkpoint_0071_0075 subagent returned PASS at 0f6f7b2. Required boundary
coverage is checked before probe overrides can mask absence; override-only and
opt-out behavior remain covered. Cancellation cleanup uses an independent
bounded 30-second context and tests both physical and Git-registration cleanup.
CHK-055 now asserts hooks ran and probes stopped after baseline failures. The
runner fixture retains both modes, readiness, output, prompt return, descendant
completion assertions and all IDs. Specs/interfaces, runner product code and
controller bounds are unchanged. The reviewer explicitly does not retroactively
validate the original 0076 native audit; the approved isolated repairs are on track.

Fresh validation at clean a34a32ac97c3c1a3ff571641117466963f338ca3:

- CGO_ENABLED=0 build, go vet ./..., controller check and go test -race ./...
  passed. Checks race package took 237.297s; runner took 20.576s. Other packages
  passed or reused valid cache. No data-race report or failed case remained.
- The build metadata confirms Go 1.27.2, darwin/arm64 and vcs.modified=false.
  An optional module-version stat-cache write was sandbox-denied; build exit
  remained 0. The unchanged embedded-recipe SPEC.md discovery warning is retained.
- Explicit history a3c044a..a34a32a passed all 19 commits with no diagnostics,
  extending the previously valid 404-commit baseline portion and 52-commit
  extension. Documentation-only report commits follow this verification point.
- The NEW missing-boundary fixture now exits 1 with X02 and C14 diagnostics,
  naming the missing hook and skipped undeclared/internal_eval canaries.
- The fresh healthy Go CLI smoke passed in 22.08s: 12 doctor checks with exactly
  two fake-agent calls, 10 no-agent checks with no extra calls, guidance preserved
  and scratch checkout/registration cleaned. Fixtures were removed.

B02 for 16 CHK-055 cases remains informational; the eval_fix added no cases or
weakened assertions. This passing scoped review closes the correction checkpoint
and resets the five-application review count. No pending change or claim remains.
The next observable outcome is one complete eligible lifecycle through NEW next,
then bounded NEW run --max/continuation/empty-queue scenarios, as ordered by the
updated goal. Recovery and remaining kinds follow that runnable path. No new
drill or deferred contracts-gap fix is authorized.
