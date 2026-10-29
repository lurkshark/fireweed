# Applied-result count 0093 — archived eval_fix

The owner approved the scoped recovery after the usage change's bounded runner
check passed 67 cases and failed one incorrect frozen-helper assertion. Only
that change's uncommitted runner implementation was discarded. Native abandon
b08c4d0 reverted 0093-usage-limit-control and released its claim; its usage specs
never applied. Only requirements/evals were retained for fresh reuse. Its two
bounded RUN-033 audit corrections reached PASS before the different helper
conflict was found; no outstanding finding or audit bound was bypassed. The
abandoned implementation and its logs remain excluded as implementation inputs.

0093-applied-result-count is an eval_fix of RUN-016, archived at
5b0e6428a54ad1485c8cf56bb514bfa972756c41. Phases b8a9260/bd4d862/5b0e642
record 231 full passes in both evals and impl. The real fixed Sol/high worker's
single bounded public runner check passed 53 cases in 121.421s. The separate
isolated Sol/high audit returned No findings for RUN-016; complete command-log
inspection confirms only supplied snapshot reads, without SYSTEM, implementation,
history, outside reads, code execution or writes. Snapshot removed. Native
coordinator phase provenance is manual; CLI logs are retained under run
20261029T141235-0e6c1ce6b922f1a559353659.

Only packages/runner/evals/run_test.go changed: assertRunApplied delegates to
assertRunAppliedSteps with exactly three; the factored body changes its literal
step count to an explicit argument and retains all other assertions and every
existing caller/test/coverage ID. No product, public interface, model or Spec
change, impl worker, phase-write exception, drill or enforced-bound change.
Future re-authored usage tests must explicitly require all four ordered captured
invocations; implementation must be recreated fresh before CLI exit-3 wiring and
actual NEW usage proof. None of that behavior is claimed complete by this fix.

This is the fifth application since the 0084–0088 checkpoint (0089–0093).
The standing separate quality/goal review and required checks precede any push
or further change. No pending governing change or claim remains at this boundary.
