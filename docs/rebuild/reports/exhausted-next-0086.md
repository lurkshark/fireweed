# Exhausted NEW next 0086 — PASS

RUN-023/024 and narrowed RUN-015 applied through e0b5c8f/3af876d/c797264.
Full phase evidence: 213 cases, two genuine runtime failures/211 passes before
implementation, then all 213 passing. RUN-015's remaining invocation-failure
behavior was already satisfied and remains covered with its held claim/phase
and captured output. Existing bound/retry assertions remain; terminal work is
now asserted on the attempt ref instead of the checkout. The pre-freeze quick
check caught unchanged RUN-021 labels on the new failing outcomes; only those
case labels were corrected, retaining the assertions and RUN-021 coverage.

The separate isolated Sol/high audit returned No findings for RUN-015/023/024.
Complete command-log inspection confirms supplied snapshot reads only, without
implementation/history/outside access, writes or execution; snapshot removed.
The real implementation worker changed ten lines in the existing exhaustion
branch to use public workflow.Block. Focused tests, all 47 own quick cases and
the app's existing C10 wrapped-cause assertion passed. Evals/interfaces stayed
frozen, and native Done passed all 213 cases and released the claim. Coordinator
provenance is honestly manual; actual fixed-profile logs remain under run
20261029T055120-7a331d675361e12d225d4c11. No phase-write exception was used here.

## Actual CLI proof

NEW was built at clean c7972645f5079fd93c1267fbf5aee4eab7f21f5e with Go 1.27.2,
darwin/arm64, CGO_ENABLED=0 and vcs.modified=false. The optional module-version
stat-cache write was sandbox-denied; build exit remained zero.

Actual NEW next --json exited 1 in 5.61s with schema1/null data,
E01 naming Test_APP_001_Value and the saved ref, and empty stderr. The real Go
eval failed before implementation; both bounded implementation attempts failed.
Exactly four fake-agent calls occurred: evals, separate audit, impl, impl.
The final wrong-2 implementation and retained work are accessible on
refs/fireweed/attempts/one-value/1 at 6ccc4d911e126a570d3c15a625fad198948f6f6d. The checkout's implementation returned
to its original value and its untracked step file was removed. The saved commit
and block commit share the held phase base; the block changes only the pending
change file, preserving run/assigned ID/phase. No claim, archive, applied SPEC
or dirty checkout remained. The expected blocked fixture was removed after
inspection; this is not a blocked governing rebuild change.

Fixture baseline bbee54f097f63c9f9a07d2ea50f1c7864d0871a3, run 20261029T061801-b8bacb5655919ac594a7948a. Branch phase commits:
3e371de232e9b455eb17ed4ecf4ba8d27b06dbe0,
47ae8d84755cf3da23b1a4030df9822de449546e,
e342e5e22d4595e726aca9d0dc0c2df85245733b.
Evals and final impl reports each contain one genuine failing Go case. No
implementation success or apply is claimed for this intentionally failing fixture.

This closes the first runnable exhaustion/save/block/release outcome. The
[lifecycle dependency](saved-attempt-blocking-0085.md) preserves its separately
approved recovery history. Projected waiting/independent continuation, saved
attempt resume, agent-commit undo, nonzero-agent retries, usage/time stops,
remaining kinds and persistent bounds remain open. Three changes have applied
since checkpoint; no push or drill occurred. Final acceptance is incomplete.
