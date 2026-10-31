---
audit_rounds: 0
module: app
phase: applied
run: 20261031T054650-d36a45fdadc291fff5034ecf
---



## Motivation

Maintenance after 0128–0132: the CLI layer chooses diagnostic codes, exit statuses
and advice by matching error text in entry.go (stamp C08, show C05/U02),
lifecycle.go (coverage C05), ci.go and check_commits.go (SPEC.md parse C03),
pr_body.go (endpoint configuration), next.go and run.go (approval) and
project.go (projectFailureCode). Replace each with contracts.ErrorCode,
runner.IsApprovalRequired and workflow.IsEndpointConfigError. Codes, exit
statuses and messages stay the same, and existing app evals guard them.

## Interface

No interface changes.

## Added



## Modified



## Removed



## Decisions



## Rejected alternatives



## Notes

Show's U02 for a duplicated handle uses the CodedError subject's `new:` prefix
instead of the message. The wrappedCause regex that reads `CODE:` prefixes is
out of scope for this change.

In-step CLI review reproduced a behavior regression in the pending migration:
with `version: 1`, `modules: [packages/*]` and a module SPEC whose `id_prefix`
is numeric, the baseline NEW `doctor --no-agent --json` returns exit 1/C12.
The pending build instead returns exit 2/C01 and configuration advice.
Doctor's direct project.Load path still hands an uncoded Load error to
projectFailureCode, while loadValidatedProject now adds codes through Check.
Preserve the baseline diagnostic and exit on direct project-loading callers
(including doctor and agents-md) using typed diagnostic codes in app internals.
Use the reproduction to verify the correction before completing this step;
frozen evals and public interfaces stay unchanged.

The updated build was rechecked against the baseline for this fixture:
doctor returned exit 1/C12 and agents-md returned exit 2/C12 with identical JSON
diagnostics in both builds. The direct-load regression is resolved.

## Provenance

- impl 2026-10-31T06:07:56.224149Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.1 playbook=07554a3249d8fca0ba70df7504135f297491465caa63458734a4b8351734d0e5 tests=282 failing=0 passing=282
