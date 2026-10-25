---
fixes:
    - GIT-006
    - GIT-010
module: gitstore
---

## Motivation

Outer public Run and Capture can return context deadline exceeded promptly while a descendant started through another public adapter remains alive in an independent process group. Separate bounded local probes confirm both entry points: the nested child has started before cancellation and still performs a delayed write after the outer call returns. This violates the existing entire-owned-tree and no-delayed-write contracts, and can leave hook or agent work running after its owner stops.

Correct cancellation ownership without changing ordinary outcomes, stream bytes, caller inputs/environment/cwd, overlapping calls or unrelated processes. Preserve the full Linux/macOS/Windows Capture contract; host runtime evidence and cross-compilation evidence must remain distinct.

## Interface

## Added

## Modified

## Removed

## Decisions

This is a fixes-only behavioral regression for existing GIT-006 and GIT-010. Public Run/Capture signatures and models, existing evals and module zones remain unchanged. New permanent tests use bounded local Go helpers, readiness handshakes before cancellation, real independent descendant groups, both cancellation and deadline, retained pre-cancel output/context identity, delayed-write observation and a live unrelated process with the same helper image. Include nested public adapters and independent overlapping calls, including a race-instrumented calling executable. Positive controls must prove helper readiness and the delayed-write observer; do not infer child startup from elapsed time.

The ordinary discovery probes are saved as composed-{run,capture}-cancellation-probe.{go,json}; they are not permanent evals or regeneration evidence. Implement only after genuine complete fail-first and a fresh restricted independent audit. Do not narrow the existing entire-owned-tree contract to the current process-group implementation or change old evals/API during implementation.

The full rebuild goal authorizes local authored execution and commits under the pinned bootstrap controller. Keep its executable separate from the NEW subject and leave the reference checkout unchanged. Record the normal manual unmerged warning without claiming human merge; do not push, publish or tag.

## Rejected alternatives

Killing only the original process group misses independently grouped descendants. Killing processes by executable name or broad caller membership can terminate unrelated or overlapping work. Returning before descendants stop violates the explicit no-delayed-write contract. Relaxing GIT-006/GIT-010 or treating these failures as maintenance would hide a required behavioral fix.
