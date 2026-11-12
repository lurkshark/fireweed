---
audit_rounds: 0
bases: {}
eval_fix:
    - APP-076
    - APP-077
    - APP-079
    - APP-080
    - APP-081
    - APP-082
    - APP-083
    - APP-084
    - APP-086
    - APP-089
    - APP-090
module: app
phase: applied
run: 20261031T160522-cc75ead0e978d616ecc2f33c
---



## Motivation

Full 0138 verification and a complete app eval run demonstrate eight remaining
failures in next usage resume, bounded/unlimited run, run approval, time resume
and shared usage resume. Every failure compares audit cwd with the held root;
the CLI lifecycle, output, archive and exit assertions pass. APP-076–090 do
not prescribe audit cwd, and the design's restricted Audit step requires the
pending runner isolation. Correct these remaining app eval assumptions through
normal eval_fix while preserving all CLI requirements and assertions.

## Interface

No public API changes.

## Added



## Modified



## Removed



## Decisions



## Rejected alternatives

Returning audits to the held checkout would defeat the required isolation.
Another author recovery exception is unnecessary: supported tied-number,
number-then-slug ordering places this normal app eval_fix before the original
0138-isolated-audit-workspaces without changing that change's identity or run.

## Notes

Batch the remaining audit-cwd comparisons in runCLIAssertCalls (run_cli_test.go),
runForCalls (run_for_cli_test.go) and the two invocation loops in usage_cli_test.go
for NextJSONUsageStopAndOwnedResume and RunJSONSharedUsageStopAndResume.
Parse the invocation fields and retain exact module, phase, order, count and
invoking Bin checks. Workers must still use the held root. App specs do not
prescribe audit cwd: accept its workspace independently of held worker cwd;
runner isolation evals own the specific restricted workspace assertions.
Preserve every output/envelope, exit, approval, bound, partial-work, resume
identity, commit/archive, report and claim assertion and all testcase names/IDs.
Do not alter unrelated worker-only comparisons, product behavior, specs or API.
The corrected tests must pass current code and pending audit isolation.

This is the second app correction for the demonstrated obsolete cwd issue,
batching every remaining failure from the complete report. If unresolved after
this simpler scoped correction, stop for SYSTEM's human decision. Original
0138 saved implementation attempt 2 is retained; its claim was released and
only the two restored implementation files returned to HEAD after exact byte
comparison. Its IDs, run, audit round and configured bounds stay unchanged.

## Provenance

- evals 2026-10-31T16:15:21.161101Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.1 playbook=b26da634732c7f1c7dc600115582a63cb8f82e1fd2bd1d522d608fe837a04ba7 tests=282 failing=0 passing=282
- audit 2026-10-31T16:16:14.725795Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.1 playbook=5cb9f9dac24b0b6bb1b970edd4fab4406d2b4d0435ddefc3ef59a2cfdd267f18 tests=0 failing=0 passing=0
- impl 2026-10-31T16:26:19.906283Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.1 playbook=07554a3249d8fca0ba70df7504135f297491465caa63458734a4b8351734d0e5 tests=282 failing=0 passing=282

## Audit

No findings
Checked IDs: APP-076, APP-077, APP-079, APP-080, APP-081, APP-082, APP-083, APP-084, APP-086, APP-089, APP-090
