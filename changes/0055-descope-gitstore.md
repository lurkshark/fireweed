---
already_satisfied:
    - GIT-006
    - GIT-007
    - GIT-008
    - GIT-009
    - GIT-010
    - GIT-011
    - GIT-012
bases:
    GIT-006: 5aef224e
    GIT-007: 638d52ae
    GIT-008: 30358cf2
    GIT-009: 59477ab1
    GIT-010: f5a5fb7e
    GIT-011: 35bd5080
    GIT-012: 327d30d8
module: gitstore
---

## Motivation

The 2026-10-26 project review found effort going into hardening beyond the design while core commands stayed unbuilt. The revised design (Spec shape; SYSTEM.md is a plain document) and the SYSTEM.md scope guardrails say that locality and safety properties are conventions checked by review, that specs state observable behavior without aliasing or purity requirements, and that the supported environment is macOS and Linux with ordinary full SHA-1 clones. See D048.

This is the largest simplification: changes 0046 to 0048 added native process ownership (Darwin native objects, per-architecture Linux reaper executables, Windows job objects) to chase processes that leave their process group. The design asks only for a per-attempt timeout.

## Interface

No public API change. Run and Capture keep their signatures.

## Added

## Modified

### GIT-006
A context already cancelled or expired before Run MUST prevent process startup and return Exit -1 with an Err matching the context error via errors.Is. Cancellation or deadline during execution MUST kill the process group Run started, return promptly with Exit -1 and the matching context Err even if an escaped process still holds an output stream, and retain output emitted before termination. Ordinary exits MUST remain distinguishable from context termination.

### GIT-007
Independent Run calls MUST NOT share captured output or cancellation state, even when overlapping. Invalid inputs and completed, failed or cancelled calls MUST NOT affect a later valid call. Process tests MUST use local bounded helpers and no agent, bootstrap or network dependency.

### GIT-008
Capture MUST execute exactly the supplied executable/argument vector in dir without shell expansion and return complete byte-exact stdout and stderr separately in Stdout and Stderr, including empty streams, NUL/non-UTF8 bytes, long output and concurrently emitted streams, without truncation or deadlock. It MUST inherit the parent environment and apply KEY=value overrides with last entry winning, preserving empty values and equals signs. Existing Run combined-output behavior MUST remain unchanged.

### GIT-009
Capture MUST return Exit 0 and nil Err on success, the actual exit and useful non-nil Err for ordinary nonzero termination, and Exit 2/useful non-nil Err for nil context, empty argv/executable, malformed overrides, nonexistent directories or launch failures without starting a fallback process. Invalid and pre-cancelled calls MUST produce empty streams and no process-created marker. Already cancelled/expired contexts MUST prevent startup with Exit -1 and an Err matching the original context error through errors.Is. Invalid, successful, failing and cancelled calls MUST NOT poison later valid calls.

### GIT-010
Cancellation or deadline during Capture MUST kill the process group it started, return promptly with Exit -1 and the matching context Err, and retain each stream's bytes emitted before termination. Overlapping Run and Capture calls MUST keep independent output and cancellation state. Ordinary termination MUST remain distinguishable from context termination.

### GIT-011
Edits MUST accept full commit IDs and single commit, tag or ancestor expressions as its baseline and reject revision sets, retaining GIT-003 semantics. It MUST NOT fetch, push or create branches, and MUST return an error rather than partial edits when the baseline or its edits cannot be determined.

### GIT-012
Clean MUST retain GIT-003 clean/dirty meaning, including for valid unborn repositories and nested checkout paths. Invalid roots and Git failures MUST return false with a useful error. It MUST NOT fetch, push or create branches.

## Removed

## Decisions

Narrowed specs are listed in already_satisfied because the current implementation already meets the narrower text; nothing here adds behavior. Clauses removed from a spec become unspecified, not forbidden: existing code that still does more is conforming, and the implementation step may simplify it.

Process control uses a new process group per call (Setpgid), kills the group on cancellation or deadline, and bounds stream draining (for example exec.Cmd.WaitDelay) so an escaped process holding a pipe cannot delay the return. Processes that call setsid, daemonize or otherwise leave the group may outlive cancellation; that is a documented limitation. Windows and partial or promisor clones are unsupported. Product code contains no native objects, assembly, cgo or embedded executables.

## Rejected alternatives

Keeping the clauses and only deleting their evals: coverage requires every enforced spec to stay covered, and the next audit would reinstate the tests. Rewriting every over-long spec to the new shape limits at once: most long specs carry no hardening and rewriting them adds risk without benefit; they are left for later changes.

## Notes

Evals step: delete nested_process_test.go, nested_group_*_test.go, orphan_*_test.go, vanished_*_test.go, locality_configuration_test.go and locality_fixture_test.go. In locality_test.go, clean_test.go, capture_test.go, process_test.go and revision_metadata_test.go remove subtests that assert only removed clauses (escaped grandchildren, delayed writes after return, race-instrumented helper survival, promisor or missing objects, fsmonitor and content filters, SHA-256 fixtures, index-byte preservation, Windows environment case, slice backing storage), keeping GIT-002 and GIT-003 coverage. Keep plain tests that a cancelled child and grandchild in the same group are killed and that Run returns promptly when a detached process holds the output pipe. Implementation step: replace the native ownership code with process groups; delete reaper_linux_*.gz, image_linux_*.go, generate_reaper.go, generate/, native/ and Windows-specific code, and remove or reduce the per-platform ownership code in tree_*.go and snapshot_*.go that process groups supersede; drop the copied-index and lazy-fetch machinery if a simpler local implementation satisfies the narrowed specs. Run go vet and the race detector before done.
