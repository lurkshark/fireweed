---
module: gitstore
phase: audit
run: 9bbd51b0dfcfca62133286a5
---

## Motivation

Hooks and agents need one bounded local process adapter. A timed-out shell can leave descendants running and mutating a repository after verification has returned. Establish literal arguments, environment overrides, complete output and cancellation before adding checked hook and agent orchestration.

## Interface

Add CommandResult {Output string; Exit int; Err error} and Run(ctx context.Context, dir string, argv, overrides []string) CommandResult to retained api.go. The facade only converts/forwards into a new disposable internal process package. Overrides use KEY=value entries layered over the calling environment. No shell is implicit; callers that need a shell supply its executable and arguments. Supported platforms for the cancellation contract are Linux, macOS and Windows.

## Added

### GIT-005
Run MUST execute exactly the supplied executable and argument vector in dir, without shell expansion, and collect complete stdout and stderr into Output while returning the actual exit code and a non-nil Err for a nonzero exit. Successful execution MUST return Exit 0 and nil Err. It MUST inherit the parent environment, apply supplied KEY=value overrides with the last entry for each key winning (including empty values and values containing equals signs), preserve all caller inputs, and leave the caller's working directory and environment unchanged. Empty argv, an empty executable, nil context, malformed environment overrides, nonexistent working directories and launch failures MUST return Exit 2 with a useful non-nil Err without launching a fallback executable. Long output MUST be drained without deadlock or truncation.

### GIT-006
A context already cancelled or expired before Run MUST prevent process startup, return Exit -1 and an Err matching the context cancellation/deadline via errors.Is. Cancellation or deadline during execution MUST terminate the entire spawned process tree on supported platforms, including children and grandchildren that retain output streams, return promptly with Exit -1 and the matching context Err, and retain output emitted before termination. Descendants MUST NOT survive to perform delayed writes after Run returns. Cancellation MUST leave unrelated processes alive. Ordinary successful and failing process termination MUST remain distinguishable from context termination.

### GIT-007
Independent Run calls MUST NOT share captured output or cancellation state, even when overlapping. Run MUST NOT mutate caller argv or override slices, including backing storage beyond their lengths, or the parent environment/current directory. Empty or invalid calls and pre-cancelled calls MUST leave no process-created marker. Invalid inputs and completed successful, failed or cancelled calls MUST NOT poison a later valid call. Process tests MUST use local bounded helpers and no agent/bootstrap/network dependency.

## Modified

## Removed

## Decisions

Exit -1 identifies context termination inside this process API, while CLI layers retain their separately specified 0/1/2/3 meanings. The original context error is exposed for callers to classify timeouts and cancellation without parsing process messages. Launch/configuration errors use Exit 2. Explicit shell invocation remains possible but substitutions and quoting belong to the hook/agent caller, not this adapter. This primitive does not create Git branches or invoke Git by itself. Windows environment keys follow the operating system's case-insensitive semantics. Runtime platform evidence must be reported honestly; cross-compilation cannot establish cancellation behavior on an unexecuted platform.

## Rejected alternatives

Killing only a shell leaves children holding pipes or writing stale results. Implicit shell joining changes literal arguments and permits substitution of authored text. Mutating the parent environment to configure a child makes concurrent hook/agent runs interfere.

## Provenance

- evals | 2026-10-22T16:35:44Z | agent: manual | fireweed: 0.4.0 | playbook: 5867363302d4c874074520977b1627d26991d95ee2fbacb9ff63e827661edd0b | tests: 17 | failing: 8 | passing: 9
