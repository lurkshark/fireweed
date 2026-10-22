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

### APP-009
show MUST resolve an exact named module from a nested project directory and return its complete current module in JSON when ID is omitted, or the exact case-sensitive spec item when ID is supplied. Plain output MUST render only the selected ordered specs as level-three headings with retained tags and complete bodies, one blank line between items and one final newline; no specs MUST produce empty plain output. It MUST retain current/discovery diagnostics from public Check, globally sorted by path/code/severity/message, with exit 0 for no errors or warnings alone and exit 1 when errors exist while retaining selected data. Unknown modules or exact IDs MUST produce U01 exit 2 with corrective guidance. Fatal project/configuration failures MUST produce C01 exit 2 and null JSON data.

### APP-010
show --projected MUST use public Project over the complete numeric/slug ordered queue and expose resulting specs and resolved metadata without changing current state or replaying archives. It MUST combine current/discovery and CheckQueue diagnostics, preserving change paths and globally sorted order, and return projected data with exit 0/1 by diagnostic severity. A failed projection MUST return exit 1 with C05 error diagnostics and null JSON data rather than a partial projection. Exact projected ID lookup MUST work for real IDs and local handles; multiple matching local handles MUST produce a corrective U01 ambiguity error rather than choosing one. --projected and --history MUST be mutually exclusive.

### APP-011
show --history MUST return a non-null array of archive rows for the selected module, optionally filtered to exact IDs in Added/Modified/Removed, including retags and removed IDs that are no longer current. It MUST exclude pending entries and fix-only references from ID-filtered history; unfiltered module history MUST include all that module's archived changes. Each row MUST contain the complete change plus its retained Motivation, Decisions and Provenance bodies, preserving historical handle examples. Newest history MUST come first: rows with a parseable RFC3339 timestamp in their final parseable impl provenance line MUST sort before undated rows and by that timestamp descending, then numeric Number, Slug and Path descending to break ties; undated rows MUST use the same numeric/slug/path fallback. Missing history, including an unknown historical ID, MUST return an empty array with exit 0 absent other validation errors. Plain history MUST render the same row array as JSON data, and current/discovery diagnostics MUST follow the current show severity and stream contract without projecting pending changes.

### APP-012
All show modes MUST accept --json before or after command/operands and use the schema-1 canonical show/data/non-null diagnostics envelope without plaintext on JSON streams. Plain data MUST go to stdout and coded/path-attributed diagnostics to stderr. Missing/excess operands, unknown or valued flags, mutually exclusive modes and mode flags on other implemented commands MUST be U01 exit 2 before loading a project. -- MUST end option parsing; help MUST work without loading a project. Output write failures MUST return exit 2. Current, projected and historical reads on valid and invalid queues MUST remain noninteractive and MUST NOT run hooks, agents, the bootstrap or network commands or mutate repository files, ignored contents or refs.

### APP-013
coverage MUST discover from a nested project directory, resolve the exact case-sensitive named module and run its resolved eval hook through public Evaluate with empty failFirst/already lists. It MUST return the complete Coverage result, including real testcase evidence/IDs/BySpec, module name, hook exit, captured Output and diagnostics, preserving deterministic public diagnostic order. No error diagnostics MUST mean exit 0; any coverage error MUST mean exit 1 while retaining Coverage data. Plain output MUST be pretty JSON of the same data on stdout with coded/path-attributed diagnostics on stderr; JSON MUST use the schema-1 canonical coverage/data/non-null diagnostics envelope with no plaintext on either stream. The envelope diagnostics MUST equal the data diagnostics. Unknown modules MUST be U01 exit 2 with exact valid choices and no hook startup. Fatal Git/configuration loading failures MUST be C01 exit 2 with null JSON data. Current coverage MUST use current specs/hooks and MUST NOT project or validate pending intent as a prerequisite, so unrelated queue errors cannot replace selected-module coverage.

### APP-014
coverage --with-change MUST resolve an exact queued slug targeting the selected module and use public Project before that change followed by ApplyChange for the selected change, exposing the inclusive numeric/slug ordered projected module's specs and resolved hook metadata to Evaluate. It MUST exclude later queue entries and archives, and MUST NOT allocate IDs or edit current state. Unknown or archived-only slugs and a module/target mismatch MUST produce corrective U01 exit 2 without running a hook. Projection/application failure MUST produce C05 exit 1 attributed to the selected change path, null data and corrective guidance without hook startup. This command MUST still require ordinary passing coverage: Added/Modified/fixes and already_satisfied metadata MUST NOT infer failFirst or grant coverage exemptions. Stale bases alone MUST NOT prevent an otherwise valid preview, consistent with public Project; static queue validation remains check's responsibility.

### APP-015
The executable MUST retain public Evaluate's C07/C11/C12/C13/C14 errors, target module paths, corrective messages, actual hook exits and output for missing hooks/reports, unreadable/malformed reports, timeouts, reporter exit inconsistencies, missing/unknown/removed IDs, skips/uncovered specs and ordinary failures. Review and reasoned quarantine suspensions MUST remain visible in testcase mappings without inventing enforcement errors. Malformed or fatally incomplete hook evidence MUST keep non-null empty Tests/BySpec and MUST NOT become successful coverage, including a timeout after a valid failing report was written. Errors MUST return exit 1 rather than a configuration/usage exit; arbitrary hook stdout/stderr MUST remain only inside Coverage.Output, without corrupting either output format.

### APP-016
coverage MUST accept --json before or after command/operands, --with-change <slug> and --with-change=<slug> anywhere before --, and valid --help/-h without loading a project or running hooks. Repeated with-change values MUST use the last value. Missing/excess operands, empty/missing with-change values, unknown/valued boolean flags, show-only mode flags and with-change on other implemented commands MUST produce corrective U01 exit 2 before loading a project. -- MUST stop option parsing, so subsequent flag-looking text is an operand and cannot enable JSON or with-change. Output write failures MUST return exit 2. Current and projected coverage MUST remain noninteractive and MUST NOT independently run setup/boundaries hooks, agents, the bootstrap or network commands, create Git commits/refs/claims, rewrite specs/changes or modify ignored/user files, except explicitly authored eval-hook actions. Owned report directories MUST be cleaned after success and failure, and FIREWEED_BIN observed by the hook MUST identify the new subject executable rather than an inherited controller path.

## Decisions

[cli-inspection]
Canonical JSON matches the reference envelope; explain data is {code, explanation}, and embedded playbook data is {name, text, hash, overridden:false}. Text-mode explain is pretty JSON for those two fields, consistent with the reference's structured read output; text-mode playbook is exact text ending in one newline. Boolean flags are --json and --help/-h; --version/-v selects version when no command is present. Help accepts no positional operands; command --help bypasses operand requirements after option validation. --json may occur more than once. Unknown options remain errors even with help requested. This change exposes embedded lookup only; project pins and playbook overrides will be composed through later project/app changes. No runtime bootstrap path is read or executed.

Playbook JSON also includes embedded_hash equal to the selected embedded hash, preserving the reference field used for future override freshness.

[project-inspection]
The current graph comes exclusively from SPEC.md and fireweed.yaml through project.Load. Both commands call project.Check on the snapshot, so modules can publish a partial readable graph alongside error diagnostics while returning exit 1. Plain data is pretty JSON, matching the existing structured inspection convention. Validation errors retain data; configuration/discovery fatal errors return null data with C01 and path fireweed.yaml. Help is available before repository loading. These commands expose current static state; queued projection, run status, playbook overrides and workflow commands remain separate changes. Module arrays include the public model's retained document field, providing additional author context without dropping designed graph fields.

[queue-check-cli]
This is composition of the existing two pure validation surfaces. CheckQueue excludes existing Check diagnostics; the CLI globally sorts their concatenation. modules remains a current graph read, so a structurally valid queue that introduces errors affects check only. Discovery errors from loading malformed queue files remain visible to both commands through Check. Existing C01 and U01 behaviors and options stay intact.

[spec-history-cli]
Exact spec lookup is distinct from tolerant JUnit ID matching. Projected duplicate local handles are valid per-change namespaces, so selecting an ambiguous shared spelling requires showing the whole module instead. Current/history reads expose Check diagnostics only; projected reads also expose CheckQueue and never allocate IDs. JSON retains public-model documents/fields consistently with modules; plain spec reads match the reference's concise heading/body format. Historical removed events are included, matching the reference's useful extension to added/modified/retagged history. ID-filtered fix-only changes are excluded, while unfiltered module history includes maintenance and fixes.

History uses the final parseable impl provenance line because the design explicitly requires newest-first history; dated rows precede undated legacy rows, whose ordering falls back to number/slug/path. This defines a total ordering even when provenance is missing. Author decision D012 records a reference chronology mismatch demonstrated in a separate fixture. Preserve complete raw Provenance regardless of parseability. Claims, drift-since-drill status and other workflow reads remain later slices.

[coverage-cli]
coverage is selected-module evidence, not ci or check. It emits Evaluate diagnostics only and does not merge unrelated project static/queue validation. Fatal discovery/configuration errors retain the established C01 contract; a readable selected module remains evaluable despite unrelated pending/static problems. --with-change uses the existing full-queue projection before the selected slug, then applies that change once, so preceding cross-module intent participates consistently with the public projection contract. A failed preview is C05 at the selected change path with null data. It never reserves IDs or writes lifecycle fields.

Normal coverage always passes nil failFirst/already to Evaluate, matching the reference command. A change's already_satisfied metadata cannot excuse an ordinary failing or missing testcase. A structurally valid stale-base preview is still available; check reports C10 separately. Values accept both split and equals syntax, with last repeated value winning. Boolean flags keep their existing no-value contract. Existing command validation stays before help and project loading. Module and slug errors list exact choices or the actual target and a usable corrective command. Hook output is a field of structured evidence, not an extra stdout/stderr stream.
