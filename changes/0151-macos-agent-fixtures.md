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
phase: impl
run: 20261103T040710-7a84b102fdee95e56d35cb99
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
## Provenance
- evals 2026-11-03T04:17:39.505831Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=6a9c2d506363e3539cc66efd37bc6a06701249541dc4e4f01ed3aa6454bbb864 tests=304 failing=0 passing=304
- audit 2026-11-03T04:18:23.164423Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0
## Audit

No findings; checked APP-062, APP-063, APP-064, APP-065, APP-066, APP-067, APP-074, APP-075, APP-076, APP-077, APP-078, APP-079, APP-080, APP-081, APP-082, APP-083, APP-084, APP-085, APP-086, APP-089, APP-090, APP-136, APP-137, APP-138, APP-139, APP-140, APP-141, APP-142
