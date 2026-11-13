---
already_satisfied:
    - CHK-032
    - CHK-082
bases:
    CHK-032: 08797d68
    CHK-082: d043c28e
module: checks
phase: applied
run: 20261113T030643-8d27e4716722e0127834f442
---

## Motivation

The owner approved a consistent placeholder contract before regeneration. CHK-032 and CHK-082 promise shell quoting but leave prequoted placeholders ambiguous. Specify the existing supported template form: placeholders in shell strings are unquoted because Fireweed supplies quoting; argv substitution is literal. This does not require a shell parser.

## Interface

No public API changes.

## Added



## Modified

### CHK-032
For shell templates whose {module}, {module_dir} and {report} placeholders occur outside shell quotes, Evaluate MUST replace each placeholder once with a shell-quoted value. For argv templates it MUST replace placeholders once with literal values. Placeholders already inside shell quotes are outside this contract.

### CHK-082
For shell templates whose {prompt}, {prompt_file} and {git_common_dir} placeholders occur outside shell quotes, Doctor MUST replace each placeholder once with a shell-quoted value. For selected command and wrapper argv it MUST substitute literal values once. Placeholders already inside shell quotes are outside this contract.


## Removed

## Decisions

## Notes

Retain the existing literal-space/apostrophe/one-pass shell and argv assertions, including prompt-file and wrapper behavior. YAML quoting around an argv item is not shell quoting and remains supported. Do not add rejection/validation of prequoted shell placeholders, a shell parser or unrelated command semantics. Current expansion already satisfies this supported-input clarification.
Owner authorized these five regeneration prerequisites and their necessary module-local changes. Stop after this queue; no regeneration or drill is authorized. Follow SYSTEM.md and the controller context. Workers must not read docs/rebuild/PROGRESS.md or ACCEPTANCE.md. Preserve unrelated assertions and retained interfaces after evals. Use existing public-interface fixtures and focused assertions; do not redesign helpers or add a test matrix. Evals and implementation workers run fireweed context for this change before each step and fireweed verify for verification. Restricted auditors follow only the supplied audit playbook and inputs. Only the controller advances phases and commits.
## Provenance
- evals 2026-11-13T03:18:08.865022Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6.1-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=6a9c2d506363e3539cc66efd37bc6a06701249541dc4e4f01ed3aa6454bbb864 tests=130 failing=0 passing=130
- audit 2026-11-13T03:18:27.864767Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-astra -c model_reasoning_effort="medium" -s workspace-write {prompt} fireweed=0.1.0 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0
- impl 2026-11-13T03:38:46.514195Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6.1-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=5f856d6bcab308b9ff1a8a7c03b9d2547744112aba81dd30dea8f61ec8062b03 tests=833 failing=0 passing=833
## Audit

No findings; checked CHK-032, CHK-082
