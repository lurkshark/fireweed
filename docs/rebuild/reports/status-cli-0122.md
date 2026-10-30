# Status CLI — 0122

APP-118–126 applied at `74e281838830e4a33e57ff919988ee39cca50595`.
Run `20261030T171150-8d03756934b1f245e4cc8ad6`, audit_rounds1. Native
implementation verification recorded253 passes, applied/archive/released the
claim, leaving a clean tree with no pending changes. This is application1 since
the passing 0117–0121 review; the caller stays local under the owner's stop.

| Phase | Controller commit | Actual result |
| --- | --- | --- |
| Evals | e466e3ad4ade32ea87f91862548c8c9c259ccd4d | 253 cases:3 genuine unsupported-status failures/250 passes |
| Audit | ed2e28f1d9420c4522fad66b9a9bf369fceed559 | Independent APP-121 plain unmeasured-cost assertion finding, persisted round1 |
| Evals correction | e1a9aa4de8dd5ed91fb935d1aba0bcddf9938a11 | 253 cases:3 genuine failures/250 passes; existing case gained only scoped negative assertion |
| Scoped re-audit | be3249111c6d48d48daa94d53fafe22d5791218e | Independent APP-121-only PASS; other IDs remained passed |
| Implementation | 74e281838830e4a33e57ff919988ee39cca50595 | 253 passes; normal application/archive/release |

Fresh fixed Sol/high workers restored the preserved public assertions and
implemented only CLI dispatch plus internal status rendering. All frozen evals,
interfaces and other modules stayed unchanged during implementation. The worker's
quick verification passed in239.192s; required native own checks remained below5m
and full implementation verification below15m. Existing B02 APP-103/17 and
APP-104/19 warnings retain their prior explanations.

Both independent auditors read only isolated change/projected SPEC/own public
evals and public interfaces/model files. Complete command logs show no SYSTEM,
implementation, history/outside reads, code/test execution or writes. Both
snapshots were removed. The correction re-audit stayed restricted to APP-121.
No new convention observer was added: only the existing app PATH-recorder test
was extended with a blind Git pass-through and existing hook/agent sentinels.
Root implementation review found a thin workflow.Status caller with nonnull
JSON containers, retained diagnostics and deterministic plain count ordering.

A clean CGO-disabled Go1.27.2 darwin/arm64 build identifies source74e2818 and
vcs.modified=false; app vet passes. Build exit0 retained the known nonfatal Go
module stat-cache metadata permission warning. Binary metadata was inspected;
no source fix, permission expansion or repeated build was needed.

## Actual NEW proof

`/Users/joel/Workspace/fireweed/bin/fireweed` passed the disposable status fixture
on its first execution. Native NEW maintenance start/verify --quick/done recorded
real cost snapshots, archived in f1fc702e1c67408b0a74bbaf51be8b5be10ff404 and
released the claim (run20261030T173744-3d09b0569e2e74f615c4947a). Deterministic
fixture hooks supplied two passing ALP-001 JUnit cases with time2 each: status
reported4 JUnit seconds,2 cases,3 eval lines and1 disposable line for alpha only.
Other modules received no fabricated measurement. Custom budgets produced
B01/B02/B03; three alpha archive records produced P08. Warnings retained exit0;
JSON stderr stayed empty and plain diagnostics appeared on stderr.

Plain and schema1 JSON showed four modules/resolved zones, five ordered queued
rows, authored blocked phase/audit rounds, all three waiting rules
(same_module, depends_on_blocked, interface_dependency), quarantine reason,
boundary opt-out reason, unapproved2, phase counts and provenance-date drift2.
Two historical archive records and the failed drill report were authored public
metadata inputs, not executed-change or real-drill evidence. Only the cost
maintenance was actually executed. No hook or agent ran during either read;
HEAD, refs, index, tracked status, pending bytes and hook/agent log were retained.
The unclaimed disposable fixture and its state file were removed.

One actual NEW status --json on this checkout also passed:8 modules, empty queue,
unapproved1,121 archive phase-count entries, no diagnostics and no product cost
snapshots (controller verification is not fabricated as NEW metrics). HEAD,
refs and working-tree status remained unchanged.

Results/logs: `/tmp/fireweed-status-cli-{result,commands,main-result}.json`,
`/tmp/fireweed-status-cli-{run,build,vet}.log`, build-metadata.txt and
`/tmp/fireweed-0122-replacement-{evals,correction,impl}-focused.json` plus native
phase and worker logs with the same prefix. These are partial R17/R24/R25/R27/R29
evidence, not final fresh-checkout/history/CI/race/doctor/regeneration/self-hosting
or whole-goal acceptance.

## Recovery and stop

The earlier uncommitted run and its sandbox commit denial/native resume refusal
remain in the [checkpoint history](status-cli-0122-checkpoint.md). Both explicitly
approved one-time document exceptions are consumed. Native cancellation preserved
public inputs; the replacement used a fresh normal run and allocation, without
implementation reuse, manufactured claims, altered refs/bounds or audit resets.
There was no prior implementation to discard. No push or drill followed this
caller. Stop local at count1; the complete design and ACCEPTANCE remain open.
