---
bases: {}
eval_fix:
    - APP-074
    - APP-076
    - APP-077
module: app
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
