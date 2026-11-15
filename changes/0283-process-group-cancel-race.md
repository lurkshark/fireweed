---
fixes:
    - GIT-005
module: gitstore
phase: impl
run: 20261115T063707-dbdf60a45054adbee79418d6
---

## Motivation

Fireweed's Doctor cancellation eval failed on the first macOS CI attempt for
PR #13 and during the v0.2.0 release. A direct reproduction through `gitstore.Run`
also took about five seconds after cancellation when a Bash fixture signalled
readiness just before spawning `sleep 5`. GIT-005 and SYSTEM require cancellation
to terminate processes in the launched group. A child that survives the
initial group signal can keep capture pipes open until its normal exit.

## Interface

## Added

## Modified

## Removed

## Decisions

## Notes

Add a public-interface regression for cancellation as a shell launches a child
that inherits its output streams. Keep the race window from the demonstrated
fixture by signalling readiness immediately before spawning the child. Repeat
the bounded case enough to obtain fail-first evidence on macOS. The standalone
reproduction saw 9 slow returns in 60 attempts with 1 ms readiness polling.
Check the cancellation error, captured stdout and stderr, and return before
the child could finish naturally. Preserve the existing test that checks for a
surviving child with a completion marker. Use the same behavioral assertions
on supported hosts.

The fix must terminate the launched group and preserve captured output.
Keep the existing argv, process group, timeout, and context error contracts.

On macOS arm64, `fireweed verify process-group-cancel-race` passed the evals
phase with a fail-first result for GIT-005. On attempt 12, `Run` returned
5.008 seconds after cancellation, exceeding the 3-second bound. The error and
output assertions passed, as did the existing completion-marker eval.
The implementation-writable file is `packages/gitstore/internal/local/local.go`.
## Audit

No findings; checked GIT-005

## Provenance

- evals 2026-11-15T06:41:28.878514Z agent=codex fireweed=0.2.0 playbook=6a9c2d506363e3539cc66efd37bc6a06701249541dc4e4f01ed3aa6454bbb864 tests=14 failing=1 passing=13
- audit 2026-11-15T06:41:43.228404Z agent=codex fireweed=0.2.0 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0

