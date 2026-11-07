# Authorized CHK-034 fix and controller handoff

The owner authorized exactly one normal audited fix for CHK-034, followed by a
verified controller handoff after the fix completes with no live run/claim.
[0180](../../changes/0180-retain-failed-hook-output.md) repairs loss of captured
hook stdout/stderr/exit on report-read, JUnit-parse and hook-execution error paths.
Retain strict C11, fail-first and coverage enforcement. No spec/interface/hook,
runner, wrapper behavior or dependency-import input changes are authorized.

Use the current pin recorded in [PROGRESS](PROGRESS.md) through evals, independent
restricted audit and impl. Run context before each step and verify normally;
only fireweed owns phase commits, apply and archive. Preserve all configured
bounds/counters and the original **2026-11-07 22:51:05 UTC** deadline. Do not
change controllers with an active run/claim or resume another controller's claim.

After the applied archive and clean terminal state, build a separate controller
from identified clean fixed source; validate applicable normal/race Go tests, vet,
checks, hook doctor, explicit-range history/CI and the demonstrated diagnostics
path. Reuse unaffected valid evidence. Replace the primary pin atomically only
after validation, preserve existing recovery executables and an explicit recovery
copy of the retiring pin, and record its source/hash in PROGRESS. Update outside
pin metadata only as needed for the authorized handoff. No native drill is needed
or authorized by this instruction. Stop after the handoff.

The workflow subject candidate remains exactly
`79cf33a7690082f28c06bf41a72719d282bc00fe`, with validated range
`be35552b706aa3000ff8f11306d129ff1d701b85..79cf33a7690082f28c06bf41a72719d282bc00fe`.
It is distinct from the fixed controller source revision. Preserve the
[current failed workflow report](../../.fireweed/drills/workflow.md), complete
[prior report](../../.fireweed/drills/workflow-before-retry-a79f34c.md), valid checks
proof, every prior proof/attempt ref, archives and unrelated worktree. Do not apply
reconstructed code, promote before validation, change the subject candidate or push.

The previous single workflow retry is consumed and failed with PackageFailure/C11
and missing-coverage C10 diagnostics. Both raw worker streams were fully forwarded;
normal native failed-report/ref publication was observed, while rejection/emergency
native publication was not exercised. Its failed reconstruction ref is
`refs/fireweed/drills/workflow/20261107T172136.182088000`. WF4 and full product
completion remain open. Another native invocation requires separate approval after
this scoped fix/handoff; worker gap notes authorize no additional corrections.
