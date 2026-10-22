---
module: gitstore
---

## Motivation

Approved Git tree/blob restoration needs byte-exact stdout independently of warnings on stderr. Existing Run deliberately merges both streams for hook/agent feedback. Add a separate-stream primitive without losing local argument, environment or process-tree guarantees and without duplicating a runner in project.

## Interface

Add CaptureResult { Stdout, Stderr string; Exit int; Err error } and Capture(ctx context.Context, dir string, argv, overrides []string) CaptureResult to api.go. Keep these retained fields and thin delegation only; implement under internal/process. Add a new stub in evals; existing Run and all other signatures remain unchanged.

## Added

### new:streams
Capture MUST execute exactly the supplied executable/argument vector in dir without shell expansion and return complete byte-exact stdout and stderr separately in Stdout and Stderr, including empty streams, NUL/non-UTF8 bytes, long output and concurrently emitted streams, without truncation or deadlock. It MUST inherit the parent environment and apply KEY=value overrides with last entry winning, preserving empty values and equals signs; Windows keys follow OS case-insensitivity. It MUST preserve caller argv/override slices including spare backing storage and parent environment/current directory. Existing Run combined-output behavior MUST remain unchanged.

### new:outcomes
Capture MUST return Exit 0 and nil Err on success, the actual exit and useful non-nil Err for ordinary nonzero termination, and Exit 2/useful non-nil Err for nil context, empty argv/executable, malformed overrides, nonexistent directories or launch failures without starting a fallback process. Invalid and pre-cancelled calls MUST produce empty streams and no process-created marker. Already cancelled/expired contexts MUST prevent startup with Exit -1 and an Err matching the original context error through errors.Is. Invalid, successful, failing and cancelled calls MUST NOT poison later valid calls. All outcome paths MUST retain caller state and inputs.

### new:ownership
Cancellation or deadline during Capture MUST terminate its entire owned process tree on Linux/macOS/Windows, including children/grandchildren holding output streams, return promptly with Exit -1 and matching context Err, and retain each stream's bytes emitted before termination. Descendants MUST NOT survive to perform delayed writes. Unrelated processes and overlapping Run/Capture calls MUST remain alive and retain independent output/cancellation state, including in a race-instrumented calling executable. Ordinary termination MUST remain distinguishable from context termination. Evals MUST use bounded local helpers without agent/bootstrap/network dependencies.

## Decisions

Separate streams are protocol inputs; no inter-stream ordering is promised. Retain Run's combined stream semantics for existing hooks/agents. Both entry points may share a disposable process execution core with independently owned writers. The adapter itself grants no permission to run Git hooks or remote commands: callers remain responsible for their explicit argv and environment. Exit -1 is the existing internal context-termination convention, distinct from CLI exit meanings.

## Rejected alternatives

Parsing Git protocol bytes out of merged warnings is ambiguous. Shell redirection and temporary export files add quoting, mutation and cleanup concerns. A second process lifecycle implementation in project duplicates ownership/cancellation logic in the wrong module.
