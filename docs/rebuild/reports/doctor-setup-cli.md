# Doctor setup through NEW — 2026-10-28

Result: PASS at clean source 19d2b3172b88be5268bb87998be8b67cb6bc4456.
0074 exposes doctor and agents-md --write through app; this is Milestone 2 setup
evidence, not completion of bounded next/run or final acceptance.

## Phase evidence

APP-062–073 were assigned by the pinned controller. Initial evals 5620604
reported 152 cases, five runtime failures and 147 passes. Audit a02d10d found
three assertion gaps: nearest ancestor selection for doctor and guidance repair,
and complete managed-block creation. Corrected evals eb05e41 reported 153 cases,
six runtime failures and 147 passes. Scoped re-audit 3769100 passed only the three
previously pending IDs (APP-062/068/070); the other nine stayed passed. Both audit
logs show supplied context/eval reads without implementation/history access.
Implementation 19d2b31 passed all 153 app cases, applied/archived the change and
released the claim. Real workers/auditors used the configured Sol/high CLI;
implementation command-log inspection found no prohibited oracle reads.

## Executable and public CLI smoke

Built bin/fireweed with CGO_ENABLED=0 and GOCACHE=/tmp/fireweed-rebuild-cache.
Go build exited 0. The sandbox denied an optional module-version stat-cache
write; go version -m confirms Go 1.27.2, darwin/arm64, CGO_ENABLED=0, the full
source revision above and vcs.modified=false. The resulting executable ran the
smoke successfully.

One ordinary temporary Go repository used explicit local Git identity, NEW init
with Go recipe and separate local worker/auditor scripts, and invocation from a
nested directory. NEW doctor --json passed 12 checks, including all four canaries
and baseline/recovery, and made exactly two fake-agent calls. HEAD and checkout
status were unchanged. NEW agents-md --write --json refreshed a deliberately
stale block while preserving surrounding owner text exactly. NEW doctor
--no-agent --json then passed 10 hook checks without another agent call or a C16
warning. NEW JSON explanations for C14/C15/C16/P01 resolved successfully.

Git registration returned to the single fixture worktree and its physical scratch
root was empty. The fixture was removed. Total smoke time was 16.58 seconds.
These are local fake-agent results; NEW unattended real-agent execution,
remaining recipes, strict regeneration and final acceptance remain unverified.
