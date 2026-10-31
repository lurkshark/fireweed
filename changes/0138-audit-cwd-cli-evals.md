---
audit_rounds: 0
bases: {}
eval_fix:
    - APP-074
    - APP-076
    - APP-077
module: app
phase: impl
run: 20261031T152840-e99af1cbd492b35bdb19ba2f
---



## Motivation

Full 0138 implementation verification found that NEW next succeeds through its
archive but Test_APP_074_APP_076_APP_077_NextDiscoversRootAndAppliesChange still
requires the audit invocation's cwd to be the held checkout. That assumption
is absent from APP-074/076/077 and conflicts with the design's restricted Audit
step and 0138's projected runner behavior. Correct this app eval through the
normal eval_fix lifecycle while preserving the CLI behavior assertions.

## Interface



## Added



## Modified



## Removed



## Decisions



## Rejected alternatives

Changing product audit cwd to satisfy the old assertion would defeat 0138's
required isolation. Another author recovery exception is unnecessary because
this app-only eval_fix can run before 0138 using supported tied-number ordering.

## Notes

Change only the obsolete audit-cwd assertion in the named next CLI test. Keep
project discovery, invoking Bin identity, worker cwd, role invocation, output,
exit status, phase commits, archive, spec application, report and claim checks.
Keep the testcase names and listed-ID coverage. The CLI specs do not prescribe
the auditor's cwd; the runner's own isolation evals assert its new workspace.
The corrected app evals must pass both the current implementation and the
pending isolated-audit implementation without implementing product behavior.

This change was created with new-change and then author-ordered before
0138-isolated-audit-workspaces, before either lifecycle start or ID assignment.
The design explicitly allows tied numbers and orders them by slug. Original
0138 retains its identity, run, audit round and saved implementation attempt 2;
no refs, configured bounds or product requirements are reset.

## Audit

No findings
Checked IDs: APP-074, APP-076, APP-077

## Provenance

- evals 2026-10-31T15:38:16.143912Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.1 playbook=b26da634732c7f1c7dc600115582a63cb8f82e1fd2bd1d522d608fe837a04ba7 tests=282 failing=0 passing=282
- audit 2026-10-31T15:39:01.245559Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.1 playbook=5cb9f9dac24b0b6bb1b970edd4fab4406d2b4d0435ddefc3ef59a2cfdd267f18 tests=0 failing=0 passing=0
