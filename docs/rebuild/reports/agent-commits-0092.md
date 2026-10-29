# Agent-commit recovery 0092 — actual NEW proof PASS

RUN-027/028 implement the design's soft reset of commits made during a phase
agent invocation. Next checks HEAD after invocation, resets it to the held base
before verification or invocation-error return, keeps index/worktree data and
captured AgentResult/Steps, and returns phase-tagged P05 warnings through existing
State.Diagnostics. Warnings survive later phase, block and failure results.
P05 retains its existing D006 HEAD/base mismatch cause; warning severity and the
message identify recovery. No public facade/model or other-module change.

Native phases 0ea029d/8bf710c/96211bd record 231 cases: two genuine failures/229
passes before implementation, then all 231 pass. A first unused test import was
fixed before the genuine fail-first run; compile failure is not counted as
fail-first. The real Sol/high implementation worker's single runner check passed
53 tests in 122.786s. Separate isolated Sol/high assertion audit passed both IDs;
its complete command log contains permitted snapshot reads, no SYSTEM,
implementation/history/outside reads, writes or execution. Snapshot removed.
No frozen test edit, phase exception, bound change or drill. Coordinator phase
provenance is manual; real logs are under
20261029T100929-08538f37f6e35ed6a14eea8a.

NEW rebuilt at clean 96211bd8e2a842933ca0c0157cc7c7f0065847a8, Go1.27.2
darwin/arm64, CGO disabled, vcs.modified=false. Its optional module stat-cache
write was sandbox-denied; build exited 0. Actual NEW next --json in a disposable
Go fixture passed in 6.12s:

- Bounded fake agents made four commits: evals e7f8b818, audit 8f31e62f, impl
  ca29be8d and 716442d9. All four still existed as commit objects but none was
  an ancestor of the final branch HEAD. Only three checked phase commits followed
  the approved fixture base: d37d80e5/600c5697/8dbe21d5.
- Retained Value=new code, the second impl commit's KEPT_COMMIT_WORK file and
  post-commit KEPT_UNTRACKED_WORK file were in the checked result. ALP-002 had
  real Go fail-first/pass evidence, was applied, and Alpha archived. Beta/Gamma
  remained unstarted. Exactly three agent calls ran.
- Exit 0, schema1 applied result, empty stderr, three P05 warning=true diagnostics
  naming evals/audit/impl and retained work. Working tree clean, claim refs absent.

Fixture baseline 8f6ab3c2bc7205d4cb13612bbb091333f70311aa; run
20261029T103150-9689c2fd6f1772c479972602. Local detailed proof:
/tmp/fireweed-agent-commits-smoke-result.json and
/tmp/fireweed-agent-commits-command-result.json. Fixture and state tracker were
removed after PASS. The governing repository was never reset or committed by a
worker; deliberate fixture-agent commits are the supported recovery test input.

This closes agent-commit undo for the bounded NEW scenario under partial
R12/R17. Usage/time stops, remaining kinds/bounds and full acceptance stay open.
Four applications since the passing checkpoint; last reviewed push 82ef71d.
No pending governing change or claim remains. The next application triggers the
standing separate quality/goal-adherence review before further work or push.
