# Usage-limit runner 0094 — library dependency archived

RUN-029–035 and the RUN-015 non-limit qualification applied at
a720cf6e2ab07c6f6429992672e4e02e8308a9be. Phases c9a5729/03f9476/5ff9b3d/
b3cbdc2/a720cf6 record 246 cases (15 genuine failures/231 passes), then 247
(16 failures/231 passes) after one scoped correction, and all247 passing in impl.
The fresh fixed Sol/high evals worker reused only retained public evals and the
approved exact-four-invocation assertion; its single check compiled and had
53 passes/15 failures in145.4s. No abandoned implementation/log was an input.

The initial isolated audit checked RUN-015/029–035 and found RUN-033's approved
snapshot assertion did not distinguish a fresh load. Native round1 retained the
same run/IDs. One new public case changes uncommitted fixture budget5ms to0s,
then requires the held approved5ms policy, same claim/run/phase/ID/HEAD/work and
untouched raw file. Its focused compile/fail-first check and full phase passed.
The separate RUN-033-only re-audit returned No findings. Both complete command
logs contain permitted snapshot reads only, with no SYSTEM/implementation/
history/outside access, code execution or writes; snapshots removed. Coordinator
native provenance is manual; actual CLI logs are retained under run
20261029T142712-c76b3cdf9fbe34c2581fe5f8.

Fresh implementation changed only disposable runner code: duration policy,
cancellable waits, shared per-invocation wait budget, retained clean-stop state/
steps, owned approved-snapshot resume and wrapped classifier/failure precedence.
Public facade/evals remained frozen. The first own check had62 passes/7 failures
in154.61s from one verification-attempt off-by-one error. One implementation-only
condition correction yielded69 passes, no failures/skips, exit0 in140.83s
(package-reported138.478s). Full native implementation verification passed247.
No phase-write exception, audit bound change, saved-code oracle or drill.

One application since the passing 0089–0093 checkpoint/push28ced38. Static status
has only the existing C01 recipe warning, B02 CHK-055/16 cases and P08 three
consecutive runner applications; the latter is covered by explicit approved
recovery and the next caller change targets app. No pending change/claim.
This is a library dependency only: current installed NEW is still the checkpoint
build06a94e9. App exit3/JSON/plain output and actual NEW usage proof follow
immediately; no usage CLI completion, --for or final acceptance claim yet.
