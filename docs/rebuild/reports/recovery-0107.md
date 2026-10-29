# 0107 recovery checkpoint — owner decision required

Completed this turn: owner-approved blocked batch push to dac2d16, corrections
0104/0105 with native checks and restricted audits, actual NEW mixed-kind
rejection, same checkpoint reviewer scoped PASS, fresh build/vet/race/static/
history16 commits. Corrections pushed to origin/main5358b5f.0106 release CLI
then archived39a1d36,203 own/native passes and restricted audit PASS. Actual NEW
Go start/release preserved dirty content/runtime/HEAD, dropped only owned claim,
and repeated release returned P02; fixture removed. These results remain valid.

Main goal incomplete. Applied since passing checkpoint:1 (0106).0107 is pending,
not applied. Governing run20261029T211014-c172b18434a8e46b0f036ac7, FLOW-087/088,
phase impl, audit_rounds0. Claim0107-change-commits remains owned here. Current
HEAD9ec2165374cdef5d2303652ca34cbe1ef80db79f; only uncommitted change is the
12-line FLOW-087 reachability correction in workflow/evals/commits_test.go.
Its exact patch is saved as /tmp/fireweed-0107-reachability-correction.patch.
No product implementation beyond the declared compiling stub has been made.

## What happened

Initial public evals had102 named passes/3 stub failures. Coordinator noticed
archive destination directory absent before freeze, interrupted exact native
Done without a phase commit, and a fixed Sol/high worker added only error-checked
MkdirAll. Corrected native evals0c13958 passed fail-first:384 cases3 failing/
381 passing. Separate restricted auditor read only permitted snapshot inputs
and reported FLOW-087: "commits reachable from HEAD" — missing assertion.
No forbidden inputs/execution/writes; snapshot removed.

Coordinator submitted that fragment then called Done. It returned Advanced, but
actually committed9ec2165 with phase impl/audit_rounds0 and "No findings. Checked
FLOW-087, FLOW-088." This mechanical record does not represent the independent
auditor's finding. Coordinator failed to verify the resulting phase before
sending the real Sol/high correction worker. The worker added12 lines only,
preserving all prior assertions, and got102 passes/3 genuine stub failures.
Its subsequent Done was rejected: P2-01 impl may not edit evals. No additional
commit, phase change, native blocked metadata or implementation followed.
The claim is held; do not label this mechanical audit a quality PASS.

## Proposed scoped recovery (not executed)

1. Preserve the saved12-line patch. Restore only that uncommitted eval file to
   HEAD; no other discard/reset. Natively release this worktree's0107 claim.
2. One-time explicit phase-write exception: git revert only invalid native audit
   commit9ec2165, restoring committed audit phase/provenance before that pass.
   Preserve the governing run and FLOW-087/088 IDs. No branch/ref/bound tricks,
   history rewrite, replacement change, widened edits or third audit round.
3. Natively resume same0107 in audit. Re-submit the finding as FLOW-087's full
   normative first sentence (rather than fragment): "Commits MUST return the
   full object IDs, newest first, of commits reachable from HEAD that change the
   selected change document's phase relative to their first parent, following
   its pending and archive paths." Inspect the resulting context: require
   phase evals, audit_rounds1, pending FLOW-087 before any worker runs. If the
   controller cannot persist it, stop with state retained; no automatic fallback.
4. Fixed Sol/high worker reapplies only the saved reachability assertion while
   evals correction is authorized; native full fail-first, isolated FLOW-087-only
   re-audit, then valid audit/app implementation/full verification/archive.
   Preserve FLOW-088's existing assertions and prior coverage. Continue immediately
   to app commits caller/actual NEW proof only after success. Any new block stops.

Authority required because SYSTEM says only fireweed owns phase transitions and
protected author commits occur between claims. Earlier phase exceptions were
scoped to0085 and consumed; none currently authorizes reversing0107 phase.
No recovery is assumed approved from elapsed time or this report. Existing goal
remains active at this first recovery checkpoint; no overall completion claim.

Evidence: /tmp/fireweed-0107-{evals-done2.log,audit-worker.log,
findings-done.log,round1-context.txt,round1-worker.log,round1-check.json,
round1-done.log,recovery-context.txt}; native commits0c13958/9ec2165.

## Owner-approved execution — 2026-10-29

Direct answer: "Approve scoped 0107 recovery". Saved patch verified identical;
only its uncommitted eval edit restored. Native release removed the owned claim.
Owner-authorized phase exception consumed by revert 41886128290babe59320719d50545d803fd1e429 of
invalid audit9ec2165. Run20261029T211014-c172b18434a8e46b0f036ac7 and FLOW-087/088
retained in committed audit phase. Resume and full-clause submission next; require
phase evals/audit_rounds1/pendingFLOW-087 before any correction worker.
