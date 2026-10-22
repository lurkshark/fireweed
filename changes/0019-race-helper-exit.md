---
already_satisfied:
    - GIT-006
bases:
    GIT-006: 9f347585
module: gitstore
phase: audit
run: d89e9a707529696cba0a7a6b
---

## Motivation

The full go test -race ./... checkpoint after command-process failed Test_GIT006CancellationKillsTree/deadline_false: its unrelated helper deliberately runs three seconds, and the race runtime adds its default one-second exit sleep, exceeding the 3.5-second post-cancellation wait. The process implementation and normal frozen evals pass, including bootstrap CI. Fix this test-only instrumentation mismatch while retaining every cancellation, unrelated-survival, output, promptness and delayed-write assertion unchanged.

already_satisfied is explicit because the existing product already implements GIT-006, as shown by the executed standard suite and transitive controller verification. This repair adds no product behavior or exemption from coverage, race detection, cancellation or the final race gate. The failed full-race log remains evidence. Re-run the exact full race command after repair; no prior race success is claimed.

## Interface

No interface or product implementation change. Clarify the existing unrelated-process guarantee for race-instrumented callers. Configure only the helper subprocess environment to set GORACE=atexit_sleep_ms=0. Race detection remains enabled; the parent process environment, detector and process assertions remain unchanged. Existing tests exercise this helper protocol during go test -race without requiring a second recursive race build.

## Added

## Modified

### GIT-006
A context already cancelled or expired before Run MUST prevent process startup, return Exit -1 and an Err matching the context cancellation/deadline via errors.Is. Cancellation or deadline during execution MUST terminate the entire spawned process tree on supported platforms, including children and grandchildren that retain output streams, return promptly with Exit -1 and the matching context Err, and retain output emitted before termination. Descendants MUST NOT survive to perform delayed writes after Run returns. Cancellation MUST leave unrelated processes alive, including when the calling executable is instrumented for race detection. Ordinary successful and failing process termination MUST remain distinguishable from context termination.

## Removed

## Decisions

Use the Go race runtime's exit-sleep setting only for bounded helper subprocesses. Do not increase the tested process termination bound or remove the live unrelated-process predicate. This is a test harness repair for already implemented behavior, governed through evals/audit/impl/archive so permanent eval changes are reviewable and never smuggled into an implementation commit. Implementation phase requires no product code edits. Keep the original race failure in .fireweed/evidence/go-race-process-checkpoint.log and record the corrected full-run result separately.

## Rejected alternatives

Globally changing the parent GORACE configuration would hide that the default full race command must pass. Extending process cancellation bounds would change product evidence to accommodate a reporter shutdown delay. Editing the already archived eval commit would destroy the original evidence trail.

## Provenance

- evals | 2026-10-22T16:48:34Z | agent: manual | fireweed: 0.4.0 | playbook: 5867363302d4c874074520977b1627d26991d95ee2fbacb9ff63e827661edd0b | tests: 17 | failing: 0 | passing: 17
