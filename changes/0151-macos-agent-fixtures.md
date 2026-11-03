---
bases: {}
eval_fix:
    - APP-062
    - APP-063
    - APP-064
    - APP-065
    - APP-066
    - APP-067
    - APP-074
    - APP-075
    - APP-076
    - APP-077
    - APP-078
    - APP-079
    - APP-080
    - APP-081
    - APP-082
    - APP-083
    - APP-084
    - APP-085
    - APP-086
    - APP-089
    - APP-090
    - APP-136
    - APP-137
    - APP-138
    - APP-139
    - APP-140
    - APP-141
    - APP-142
module: app
---

## Motivation

ACCEPTANCE C1 and the design's language-independent eval contract require the
existing CLI assertions to run successfully on this Mac. The owner supplied the
diagnosis: macOS sh kills fixture scripts receiving long prompt arguments.
Correct the app fake-agent fixtures to invoke bash (or pass prompts through
stdin), preserving every assertion and all existing spec coverage. Do not
re-investigate the OS failure or change production behavior or timeouts.

## Interface

None. This is an eval_fix; no new spec handles or public interfaces are needed.

## Added

## Modified

## Removed

## Decisions

## Notes

Review every app eval fixture which supplies {prompt} to a script, including
next/run/drill, usage/time-stop variants and doctor fixtures. Use bash for the
fake worker/auditor and any prompt-bearing script wrapper. Keep hook behavior,
all testcase names, assertions, durations and configured bounds unchanged.
Only app eval files may change. All app evals must pass through normal
fireweed verify, followed by an independent restricted audit and normal full
implementation verification. Workers verify and return; only the outside
owning controller may advance phases, commit or archive.
