# History CLI — 0116

APP-105–108 phasesb784784/c016d84/bd4fba5. Native evals243 cases6 expected
unsupported-command failures/237 passes; final243 passes, impl about218s.
Restricted isolated audit PASS, audit_rounds0; complete command inspection found
only allowed snapshot public-input reads, no implementation/SYSTEM/history/outside/
code/tests/writes. Snapshot removed. No API/model change or frozen-file edit.
Current-project failures use existing diagnostic helpers; a narrow recorded-
malformed-SPEC mapping preserves C03 when current documents are valid.

Prefreeze worker repairs fixed one compile issue and the recorded-malformed-SPEC
fixture: the historical bad document must be committed then current SPEC restored;
its error string need not itself include C03. Final public History preconditions
and focused caller evals have no setup/compile failures. No frozen wrong assertion
or repeated lifecycle proof. The new eval file adds4 top-level tests plus3 bounded
usage/error subtests, roughly4s focused, within unchanged own5m/full15m limits.

Clean NEW sourcebd4fba5dfffd3634f0e8cd3c7a639eca1e1e630e, Go1.27.2 darwin/arm64,
CGO0, vcs.modified=false. Actual NEW proof PASS from nested tracked module cwd:
baseline000cd69e45bda0dc9736a0cae3c5e6f61333a957, run20261030T052121-474a0b240a47717bb58a073b.
Real Go ALP-002 fails before implementation and passes afterward; three bounded
fake-agent calls include a separate audit and yield three phase commits/archive/
SPEC/provenance/claim release. Explicit nonempty range returns those3 full OIDs
in order, empty squash/diagnostics; plain output reports3 commits.

An external fixture author creates squash objecte713871b7fc18eb1076248822b6784d32523a383
from that actual checked Go tree, without moving branch/index. NEW lists one
squashed archive/path and no H01/H02 warning. Representative archive mutation
commit03fc0ad2dc307dc3c02b0ea05d708b39551d7ef4 returns exit1/H04,
null data, exact full commit/path and corrective guidance; JSON stderr stays empty.
At every history read files/HEAD/refs/agent log stay retained, no hook/agent is run.
Fixture/state removed after PASS, no harness repair or rerun.

Local logs /tmp/fireweed-0116-{evals-worker.log,evals-focused.json,evals-done.log,
audit-worker.log,audit-done.log,impl-worker.log,impl-focused.json,impl-done.log}.
Actual result/commands/run: /tmp/fireweed-check-commits-cli-{result.json,
commands.json,run.log}. Existing B02 APP-10317/APP-10419 remain explained.
This is application5 after0112–0115: mandatory separate quality/goal review and
fresh appropriate checks now gate any push. Full ACCEPTANCE remains incomplete.
