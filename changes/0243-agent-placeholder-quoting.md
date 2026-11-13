---
already_satisfied:
    - RUN-003
bases:
    RUN-003: 69d31adb
module: runner
---

## Motivation

The owner approved aligning agent placeholder semantics with hooks before runner regeneration. RUN-003 leaves already-quoted shell placeholders ambiguous. Specify the same existing supported unquoted shell-template form as checks, preserving literal argv values and one-pass substitution.

## Interface

No public API changes.

## Added



## Modified

### RUN-003
InvokeAgent MUST accept shell strings or argv for selected commands and expand {prompt}, {prompt_file} and {git_common_dir} once in command templates and wrapper argv. For shell placeholders outside shell quotes it MUST shell-quote substituted values; argv substitutions MUST preserve literal values. Placeholders already inside shell quotes are outside this contract.


## Removed

## Decisions

## Notes

Reuse existing shell/argv/wrapper and prompt-file assertions, preserving coverage for spaces, apostrophes and placeholder text inside substituted values. YAML quoting of argv entries remains supported. Do not add rejection of prequoted shell placeholders, a shell parser or new transport behavior. Current expansion already satisfies this clarification.
Owner authorized these five regeneration prerequisites and their necessary module-local changes. Stop after this queue; no regeneration or drill is authorized. Follow SYSTEM.md and the controller context. Workers must not read docs/rebuild/PROGRESS.md or ACCEPTANCE.md. Preserve unrelated assertions and retained interfaces after evals. Use existing public-interface fixtures and focused assertions; do not redesign helpers or add a test matrix. Evals and implementation workers run fireweed context for this change before each step and fireweed verify for verification. Restricted auditors follow only the supplied audit playbook and inputs. Only the controller advances phases and commits.
