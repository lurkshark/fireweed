# Deterministic CI CLI — 0117

APP-109–113 phasesdc743fe/030965f/58232ca/fb30539/2a63214. Native initial and
corrected evals247 cases4 expected unsupported-CI failures/243 passes; final247
passes at2a63214028a43681556999b7487f5f267ef60a55, about242s full implementation
(own/full limits unchanged). Four focused tests pass in19.869s. Only app/internal
CI composition/parser changed; public API/model and all frozen inputs retained.

Independent restricted audit found closest-root selection, observable default
canary results and default canary-error exit/code/guidance gaps. Same run
20261030T052903-ef4ffe0df0609b399b924ab5 and IDs retained; native round1 correction
adds a competing ancestor and one missed-boundary-canary assertion, preserving
all earlier assertions/cases. Scoped APP-109/111/112-only re-audit PASS. Complete
auditor commands inspected: only permitted snapshot public inputs, no implementation,
SYSTEM/history/outside/code/tests/writes. Snapshots removed; no phase exception.

Prefreeze correction uses supported verify_dependents:none instead of false.
Implementation review corrected current graph-configuration C02 from exit1 to2
alongside C01/C12, preserving APP-112; focused run passed again before full Done.
No frozen eval correction or new error path/allocation. Existing APP-10317 and
APP-10419 case warnings remain explained in their report, with bounded runtime.

Clean NEW2a63214, Go1.27.2 darwin/arm64, CGO0/vcs.modified=false; build passed
without warning. Actual NEW proof PASS in a fresh project with real Go ALP-001,
BET-001 and GAM-001 evals, from tracked nested module cwd. Explicit nonempty
rangec2793bc33293da8d96b99de8a2e61402db845344..b959dc4ecd1fda51698ad2a793d05caceaee675e is one author commit.
Default CI passes static/history/all-module coverage/boundaries/doctor canaries.
Deleting canaries.yaml makes default CI fail exit1/C14; --no-doctor passes the
other checks. A real Gamma Value failure returns exit1/E01 despite
verify_dependents:none, proving all-module evaluation. JSON schema1, confirmation
string/null-data diagnostics and empty stderr inspected. No agents invoked.
Restored fixture is clean/unclaimed, only one worktree; fixture/state removed.
Only fixture preparation was reused, not a lifecycle or prior success rerun.

Local /tmp/fireweed-0117-{evals-worker.log,fixture-correction-worker.log,
evals-focused.json,evals-done.log,audit-worker.log,findings-done.log,
evals-correction-worker.log,evals-correction-done.log,reaudit-worker.log,
audit-pass-done.log,impl-worker.log,config-correction-worker.log,
impl-focused.json,impl-done.log}. Actual /tmp/fireweed-ci-cli-{result.json,
commands.json,run.log}. This is application1 since passing/pushed0112–0116 gate;
CI remains local. Native archive clean/unclaimed before evidence commit.

Next app caller would be the third consecutive app change after0116/0117.
SYSTEM.md42 requires an owner checkpoint before it. Stop here; full goal and
ACCEPTANCE remain incomplete. No new drill or extra push authority is inferred.
