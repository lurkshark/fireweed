---
depends_on:
    - contracts
    - gitstore
    - assets
    - project
    - checks
    - workflow
    - runner
id_prefix: APP
name: app
---

# app

## Specs

### APP-001
The executable MUST print help with the complete designed command inventory, common --json option and meanings for exit statuses 0, 1, 2 and 3. Empty arguments, help, -h and --help MUST succeed. version, -v and --version MUST print the embedded binary version with exit 0. Help and version MUST work outside a Git repository and MUST NOT create project files.

### APP-002
explain <code> MUST return the embedded diagnostic's code and corrective explanation, and playbook <name> MUST return exact embedded text with its byte-exact SHA-256 and name in JSON. Known codes and all five playbooks MUST succeed with exit 0 outside a project, using only the new binary. Unknown names/codes MUST be usage errors listing valid choices. Missing and extra positional inputs MUST be rejected. Embedded reads MUST NOT modify files or refs.

### APP-003
Every command in this slice MUST accept --json before or after the command/operands and emit exactly one JSON value to stdout with schema_version 1, canonical command name, data and a non-null diagnostics array. Successful diagnostics MUST be empty. JSON errors MUST return exit 2, data null and a diagnostic with code U01, severity error and a nonempty corrective message, without plaintext on stdout or stderr. Plaintext usage errors MUST have empty stdout and a U01 diagnostic on stderr. Output write failures MUST return exit 2.

### APP-004
The CLI MUST reject unknown commands, unknown flags, valued boolean flags and excess operands with exit 2 and valid command/options guidance, without prompting or hanging. -- MUST stop option parsing, so subsequent tokens are operands. A literal --json after -- MUST NOT enable JSON output. Global help MUST provide help for a command without requiring its operands; aliases MUST use canonical help/version names in JSON. None of these inspection commands MAY invoke an agent or mutate the working directory.

### APP-005
modules MUST discover the project from a nested working directory and publish the resolved module array sorted by name then directory, including names, prefixes, directories, dependencies, resolved zone lists, hooks, specs and Decisions. It MUST apply current project defaults and explicit empty overrides and MUST NOT require a second language-specific graph. Empty projects MUST return an empty non-null array. Both plain output and JSON data MUST preserve these values.

### APP-006
check MUST run public project static validation and return data {valid:true} with exit 0 when there are no error diagnostics, including when warnings exist, and {valid:false} with exit 1 when errors exist. modules MUST return its module array with the same validation diagnostics and exit 0 or 1 by severity. JSON MUST use the schema-1 canonical command/data/non-null diagnostics envelope, with each diagnostic's code, severity, message and repository-relative path in deterministic project order. Plain diagnostics MUST go to stderr with their code and path; data MUST go to stdout. Invalid/missing/incompatible configuration or absence of a Git repository MUST return exit 2 with C01 error diagnostics and null data in JSON, without U01 or plaintext on JSON streams.

### APP-007
modules and check MUST accept --json before or after the command, accept --help/-h without loading a project, and reject excess operands or unknown/valued flags with U01 exit 2 before loading a project. They MUST remain noninteractive and MUST NOT run hooks, agents, the bootstrap or network commands, nor change repository files, ignored content or refs. Project validation MUST expose invalid graph/zone/hook/document/reservation diagnostics from the public project validator rather than silently succeeding.

### APP-008
check MUST combine current/discovery diagnostics with public queued-change validation, including stale/missing bases, projected enforcement, handle scope and ID/reference errors. The combined diagnostics MUST sort by path, code, severity and message without repeating unchanged preexisting problems, preserving code, severity, corrective message and repository-relative path. It MUST return exit 1 and {valid:false} for any errors; a valid queue or warnings alone MUST return exit 0 and {valid:true}. Plain data MUST go to stdout with coded/path-attributed diagnostics on stderr; JSON MUST use the existing schema-1 envelope with no plaintext on either stream. From a nested directory, it MUST process numeric/slug ordered intent regardless of lexical filenames and MUST NOT modify files, ignored content or refs or run hooks, agents, the bootstrap or network commands. modules MUST retain its current-static-only behavior when queued projected errors exist.

## Decisions

[cli-inspection]
Canonical JSON matches the reference envelope; explain data is {code, explanation}, and embedded playbook data is {name, text, hash, overridden:false}. Text-mode explain is pretty JSON for those two fields, consistent with the reference's structured read output; text-mode playbook is exact text ending in one newline. Boolean flags are --json and --help/-h; --version/-v selects version when no command is present. Help accepts no positional operands; command --help bypasses operand requirements after option validation. --json may occur more than once. Unknown options remain errors even with help requested. This change exposes embedded lookup only; project pins and playbook overrides will be composed through later project/app changes. No runtime bootstrap path is read or executed.

Playbook JSON also includes embedded_hash equal to the selected embedded hash, preserving the reference field used for future override freshness.

[project-inspection]
The current graph comes exclusively from SPEC.md and fireweed.yaml through project.Load. Both commands call project.Check on the snapshot, so modules can publish a partial readable graph alongside error diagnostics while returning exit 1. Plain data is pretty JSON, matching the existing structured inspection convention. Validation errors retain data; configuration/discovery fatal errors return null data with C01 and path fireweed.yaml. Help is available before repository loading. These commands expose current static state; queued projection, run status, playbook overrides and workflow commands remain separate changes. Module arrays include the public model's retained document field, providing additional author context without dropping designed graph fields.

[queue-check-cli]
This is composition of the existing two pure validation surfaces. CheckQueue excludes existing Check diagnostics; the CLI globally sorts their concatenation. modules remains a current graph read, so a structurally valid queue that introduces errors affects check only. Discovery errors from loading malformed queue files remain visible to both commands through Check. Existing C01 and U01 behaviors and options stay intact.
