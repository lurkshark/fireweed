---
fixes: [GIT-005]
module: gitstore
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
