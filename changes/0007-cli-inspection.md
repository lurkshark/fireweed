---
module: app
phase: audit
run: f1bf1735e590bfb808e4c9b6
---

## Motivation

Turn the compiling placeholder into an actual executable inspection surface before building more layers. Agents need help, version, diagnostic remediation and embedded protocol text from the new independent binary.

## Interface

Keep Run(args []string, out, errOut io.Writer) int and protected main.go forwarding unchanged. Implement disposable CLI parsing/output/dispatch for help, version, explain and embedded playbook, using assets' public interface. Permanent evals build and launch the NEW product executable in temporary directories; no controller is used as the subject.

## Added

### APP-001
The executable MUST print help with the complete designed command inventory, common --json option and meanings for exit statuses 0, 1, 2 and 3. Empty arguments, help, -h and --help MUST succeed. version, -v and --version MUST print the embedded binary version with exit 0. Help and version MUST work outside a Git repository and MUST NOT create project files.

### APP-002
explain <code> MUST return the embedded diagnostic's code and corrective explanation, and playbook <name> MUST return exact embedded text with its byte-exact SHA-256 and name in JSON. Known codes and all five playbooks MUST succeed with exit 0 outside a project, using only the new binary. Unknown names/codes MUST be usage errors listing valid choices. Missing and extra positional inputs MUST be rejected. Embedded reads MUST NOT modify files or refs.

### APP-003
Every command in this slice MUST accept --json before or after the command/operands and emit exactly one JSON value to stdout with schema_version 1, canonical command name, data and a non-null diagnostics array. Successful diagnostics MUST be empty. JSON errors MUST return exit 2, data null and a diagnostic with code U01, severity error and a nonempty corrective message, without plaintext on stdout or stderr. Plaintext usage errors MUST have empty stdout and a U01 diagnostic on stderr. Output write failures MUST return exit 2.

### APP-004
The CLI MUST reject unknown commands, unknown flags, valued boolean flags and excess operands with exit 2 and valid command/options guidance, without prompting or hanging. -- MUST stop option parsing, so subsequent tokens are operands. A literal --json after -- MUST NOT enable JSON output. Global help MUST provide help for a command without requiring its operands; aliases MUST use canonical help/version names in JSON. None of these inspection commands MAY invoke an agent or mutate the working directory.

## Modified

## Removed

## Decisions

Canonical JSON matches the reference envelope; explain data is {code, explanation}, and embedded playbook data is {name, text, hash, overridden:false}. Text-mode explain is pretty JSON for those two fields, consistent with the reference's structured read output; text-mode playbook is exact text ending in one newline. Boolean flags are --json and --help/-h; --version/-v selects version when no command is present. Help accepts no positional operands; command --help bypasses operand requirements after option validation. --json may occur more than once. Unknown options remain errors even with help requested. This change exposes embedded lookup only; project pins and playbook overrides will be composed through later project/app changes. No runtime bootstrap path is read or executed.

Playbook JSON also includes embedded_hash equal to the selected embedded hash, preserving the reference field used for future override freshness.

## Rejected alternatives

Delegating command execution to the controller would produce false end-to-end evidence and an unusable standalone product. Implementing the entire CLI before its first executable eval would postpone critical argument/output contract checks.

## Provenance

- evals | 2026-10-22T07:43:18Z | agent: manual | fireweed: 0.4.0 | playbook: 5867363302d4c874074520977b1627d26991d95ee2fbacb9ff63e827661edd0b | tests: 6 | failing: 5 | passing: 1
