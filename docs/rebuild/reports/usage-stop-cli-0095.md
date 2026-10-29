# Usage stops 0095 — actual NEW proof PASS

APP-076/077/079/083/084/086 applied at
dbda392145c7c23f41dbb6f645deda048b8c4284. Native phases
6978e2a/51924d9/6dfe616/e546a01/dbda392 record186 cases: seven genuine
failures/179 passes in both evals rounds, then all186 passing in impl.

Fixed Sol/high workers added one public CLI eval file and implemented only
next.go/run.go. The first full evals check had178 passes/seven failure events
(six leaves) in169.540s. Before freeze, a mistaken new Next.Empty assertion was
corrected and its focused check failed only for intended stop behavior. The
explicit earlier-block/later-usage case initially selected blocked Alpha in its
fake script; a fixture-local selector correction made its focused check pass
in2.78s. No existing helper/assertion/caller was weakened.

The isolated audit found APP-076's generic next token did not distinguish the
held change from wait/guidance text. Native round1 preserved run/IDs; the scoped
fixture now names held-quota-work and asserts that identifier. Full fail-first
and APP-076-only re-audit passed. Complete command-log inspection confirms only
permitted snapshot reads, no SYSTEM/implementation/history/outside access, code
execution or writes; snapshots removed. Native provenance is manual; actual
Sol/high CLI logs remain under20261029T150558-9c8afdd02cdb5e6f83345c7e.

The implementation worker's single full check had183 passes/three config-test
failure events in177.09s. A corrective-wording-only implementation correction
passed the focused config check; full native Done then passed186. Evals,
interfaces and other modules stayed frozen. No new diagnostic allocation,
phase-write exception, bound change or drill. Existing C01 explanation covers
invalid settings; clean stops use result reason and warnings, without an error
code. Ordinary earlier-block failure remains exit1/null with cause/ref and stop
reason, as directly asserted through NEW in the permanent CLI evals.

NEW built at clean dbda392145c7c23f41dbb6f645deda048b8c4284, Go1.27.2,
darwin/arm64, CGO disabled, vcs.modified=false. An optional module stat-cache
write was sandbox-denied; build exited0. Actual NEW Go fixture passed:

- run --max2 --json from a nested directory returned3 in6.25s. Alpha's first
  eval agent hit a usage limit, waited5ms, retried with attempts1 and applied.
  Beta's first usage exit stopped against the shared completed5ms/next5ms/
  budget5ms limit. Schema1 retained Alpha's applied summary, empty=false,
  limited=false and reason; diagnostics[], stderr empty. Five agent calls.
- At stop Beta held refs/fireweed/claims/value-beta, run
  20261029T154405-841352e6a569047ffdc272c7, evals phase and BET-002, with its
  partial eval work retained. No Beta phase commit, block or release occurred.
- A later next --json returned0 in4.62s and archived Beta under that same run,
  preserving its exact partial work and assigned ID. Alpha's archive and Gamma's
  queued document stayed identical. Total eight bounded fake-agent calls and
  six checked phase commits: Alpha82112cef/3ad3a69a/efb868da and
  Beta62c71a7a/6c4141f3/45d78c9e. ALP-002 and BET-002 both had real Go
  fail-first/pass reports and applied Specs. Final tree clean, claims absent.

Fixture baseline46b78549b4c91c9ef6eed3ff185b7c5493ede1a0; final
45d78c9ea879d9d7780d45fd9e64207ae6021c98. Detailed local proof/commands:
/tmp/fireweed-usage-cli-result.json and /tmp/fireweed-usage-cli-commands.json.
Fixture and state tracker removed after PASS. The fixture used deterministic
fake agents/auditors; no real-agent unattended/drill acceptance claim.

This closes the bounded usage-budget/owned-resume CLI scenario under partial
R13/R17. Time --for, remaining kinds/bounds, full command inventory and final
gates remain open. Two applications since the passing0089–0093 checkpoint;
last successful push28ced38. No pending governing change or claim.
