# Checkpoint 0071–0075 — BLOCK

Five changes have applied since the passing 0065–0069/0070 re-review. Separate
subagent checkpoint_0071_0075 reviewed current code at product commit 4b2fe8e
(checkpoint author head a3c044a) and returned BLOCK. No push or further
implementation occurred. The tree has no pending change or claim. The full
rebuild goal remains incomplete.

## Applied evidence

| Change | IDs | Final eval / audit / impl | Phase results |
| --- | --- | --- | --- |
| 0071 doctor-probes | CHK-049–069 | ebf127f / 8cafea6 / e50061d | 34 failures/226 passes → 260 passes; approved recovery details in doctor-0071-block.md |
| 0072 doctor-helper-portability | eval_fix CHK-050/051 | dc2b1e7 / 426e259 / 762052f | 260 passes; helper only, ordinary Apple Git |
| 0073 doctor-diagnostic-explanations | AST-026/027 | b364af7 / 53cae57 / 8d27f34 | six failures/383 passes → 389 passes |
| 0074 doctor-setup-commands | APP-062–073 | eb05e41 / 3769100 / 19d2b31 | six failures/147 passes → 153 passes; one scoped correction |
| 0075 bounded-agent-invocation | RUN-001–009 | 2e40588 / 6f898ea / 4b2fe8e | 25 failures/153 passes → 178 passes; one scoped correction |

Passing re-audits retained prior passes and reviewed only pending IDs. Coordinator
inspection of phase-auditor logs found no implementation/history reads. Workers
used the configured Sol/high CLI. 0071's earlier audit deviation and explicit
human recoveries remain in its linked report; no counters or saved refs were
reset. The 0075 worker's prose says 24 own cases, but actual quick JSON and phase
results show 25 own cases; the durable count above uses those results.

## Review findings

1. **P1 — required boundary hooks can be absent without failing doctor.**
   packages/checks/internal/engine/doctor.go:374 records whether enforcement is
   required, but emits no missing-hook error. probeHooks at line 542 gets an
   empty boundary variant list, so crossClean and boundary canaries are skipped.
   This violates the design's required-boundary rule and CHK-055–057. A disposable
   module without an opt-out can therefore receive success without enforcement.

   A focused ordinary NEW Go fixture confirmed this at a3c044a: remove the
   boundaries hook, retain the default disposable app module, commit with local
   identity, then run doctor --no-agent --json. It returned exit 0, no diagnostics
   and only baseline/failing_eval/missing_id/recovery checks; undeclared and
   internal_eval were absent. The temporary fixture was removed.

2. **P2 — cancellation leaves scratch Git registration.** Cleanup at
   doctor.go:112 invokes d.git with the cancelled probe context (line 150).
   gitstore.Run rejects an already-cancelled context before launching Git.
   os.RemoveAll deletes only the checkout, leaving its Git registration behind,
   contrary to CHK-051. This finding is confirmed by the current control flow;
   no additional cancellation runtime repro was claimed.

The reviewer found no separate lost assertions, interface/zone or authorization
issue in this slice. 0072 retains its cleanup assertions and explanations retain
previous causes. The passing healthy CLI smoke in doctor-setup-cli.md remains
valid for its tested path; it does not close these two missing behaviors.

## Checkpoint verification

- Clean-source build passed with CGO_ENABLED=0. Go build metadata identifies
  a3c044a7a0f20c02789210fa9be1d4ef0c40949d, Go 1.27.2, darwin/arm64 and
  vcs.modified=false. The sandbox again denied an optional module-version
  stat-cache write; build exit was 0 and the executable is usable.
- go vet ./... passed. Controller check passed with the unchanged C01 warning
  for the embedded recipe's SPEC.md outside module discovery.
- Previously verified baseline be35552..4b36a84 history is reused. Explicit
  extensions 4b36a84..4174402 (46 commits) and 4174402..a3c044a (six commits)
  both passed with no diagnostics; no force/ref/bound changes were made.
- Full go test -race ./... failed in exactly the runner descendant-timeout case.
  Other test packages passed, including checks. No data-race report appeared.
  The timeout subcase exited before readiness with empty output, exit -1 and
  context deadline exceeded. A focused race rerun of that existing parent test
  passed in 4.483 seconds; this does not turn the full race result into a pass.

The same checkpoint subagent separately triaged that observed race failure as a
**timing-dependent frozen eval**, not evidence of a runner product defect. The
500 ms fixture deadline can expire before readiness under load. Its script also
publishes readiness before stdout/stderr, allowing cancellation to precede those
writes. A normal eval_fix for RUN-007/RUN-009 should give the fake command a
reasonable bounded startup/completion window and emit output before readiness,
while preserving readiness, timeout/cancellation, output-retention, descendant
completion assertions and ID coverage. Product code, interfaces, specs and
controller bounds stay unchanged. No eval edit or blanket race rerun was made.

## Human decision and bounded correction proposal

Controller status reports unapproved=5 and P07. origin/main remains
 e083c5f2976575f5ec114a8117b3cc4800ba339a. The standing push authorization applies
only after PASS, so this BLOCK does not authorize a push. The pinned controller
also refuses a sixth unapproved change even with --allow-unmerged. Approval
metadata alone cannot reset this shared-base count; normal synchronization is
needed. Do not alter refs, limits, archives or counts to bypass it.

Proposed human-reviewed sequence, requiring explicit approval:

1. Push the current five-change batch with these findings recorded, synchronizing
   the shared base despite the blocked automated checkpoint review.
2. Author one normal checks fixes change for CHK-051/055/056/057: reject missing
   required boundary-hook coverage using existing rules/diagnostics, preserve
   opt-out/X04 behavior, and use an independent bounded cleanup context so
   cancellation removes both scratch checkout and Git registration. Add focused
   public-interface regressions for the reported behaviors; no API/spec narrowing
   or new diagnostic allocation.
3. Author one normal runner eval_fix for RUN-007/RUN-009 limited to the existing
   observed timing-dependent test, as scoped above. Keep all assertions/coverage
   and verification bounds; change no product implementation.
4. Pass required phase checks and full race verification; have the same subagent
   re-review only these reported issues. Push the correction and continue the
   main goal only if that re-review passes. Any block or unresolved issue stops.

No correction, push, new drill or deferred contracts-gap work is authorized or
performed by this report. No phase-write exception or bound change is proposed.
