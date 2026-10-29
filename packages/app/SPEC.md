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
The CLI MUST run help, version, explain <code> and playbook <name> successfully
outside a project, returning the corresponding assets text on stdout with exit
0 and empty stderr. These commands MUST also work in a directory containing an
invalid fireweed.yaml.

### APP-002
Each information command MUST accept --json and emit one JSON object with schema
1, string data equal to its plain text result, and an empty diagnostics array.
It MUST return exit 0 and leave stderr empty.

### APP-003
Unknown commands or options and missing or extra arguments MUST return exit 2 with a nonempty corrective U01 diagnostic.

### APP-004
modules MUST report the validated graph, using the held approved projected snapshot while this worktree owns a change.

### APP-005
new-change <module> <slug> MUST use the discovered project root to create the project API's queued skeleton and print its repository-relative path with exit 0.

### APP-006
stamp <change> MUST discover the root, invoke project.Stamp and print confirmation with exit 0, accepting the existing change aliases and leaving Added handles unassigned.

### APP-007
check MUST invoke project.Check and exit 0 when no error diagnostics remain, or exit 1 with all diagnostics when static checks fail.

### APP-008
show <module> [ID] MUST return the current module document or selected item, using the whole-queue project projection with --projected.

### APP-009
start <change> MUST discover the root, invoke workflow.Start and print first-step workflow.Context with exit 0, assigning IDs, run and phase without committing.

### APP-010
verify [change] MUST invoke workflow with the actual invoking executable as Bin, returning 0 for success, 1 for failed checks or blocked state and 2 for usage or configuration errors; plain success names the resulting phase.

### APP-011
context <change> [--full] MUST return the held workflow.Context text, using compact or full behavior as requested, with exit 0 and empty diagnostics.

### APP-012
coverage <module> MUST evaluate current specs, or the selected projection with --with-change <change>, using approved metadata while held and setting FIREWEED_BIN to the actual executable.

### APP-013
Unknown explain codes or playbook names MUST return exit 2 with a nonempty corrective U02 diagnostic.

### APP-014
Commands listed by help but still under construction MUST explicitly report that they are unimplemented with diagnostic U01 and return nonzero.

### APP-015
Failed --json requests MUST use the schema-1 envelope with null data and a diagnostic containing code and message.

### APP-016
modules MUST discover the nearest ancestor fireweed.yaml.

### APP-017
modules --json MUST return schema 1 with array data containing name, id_prefix, dir, depends_on and resolved zone arrays, plus empty diagnostics on success.

### APP-018
modules MUST return exit 0 on success or exit 2 for missing or invalid project inputs, with a corrective failure diagnostic: C01 for configuration, C02 for graph, C03 for document syntax or C12 for metadata errors.

### APP-019
new-change MUST return exit 2 with corrective diagnostics U01 for invalid arguments or slug, U02 for an unknown module and C13 for duplicate slugs, preserving existing files.

### APP-020
new-change --json MUST use schema 1 with string path data on success, or null data and code/message diagnostics on failure.

### APP-021
stamp MUST return exit 2 for usage or configuration errors and exit 1 for stamping rejections, with code/message diagnostics.

### APP-022
stamp --json MUST return schema 1 with string confirmation data on success, or null data and diagnostics on failure.

### APP-023
check MUST return exit 2 for usage or configuration errors with code/message diagnostics.

### APP-024
check --json MUST return schema 1 with string summary data on success, null data on failure and the reported diagnostics.

### APP-025
show --json MUST return identical text as schema-1 string data with empty diagnostics and exit 0 on success.

### APP-026
show MUST return exit 2 with code/message diagnostics for unknown or ambiguous module or item, invalid arguments or configuration errors.

### APP-027
start MUST preserve reported warnings and return exit 1 for lifecycle rejection or exit 2 for usage or configuration errors, with corrective code/message diagnostics.

### APP-028
start --json MUST return schema 1 with context string data on success or null data on failure.

### APP-029
done [change] MUST invoke full workflow.Done with the actual invoking executable as Bin, returning 0 for success, 1 for failed checks or blocked state and 2 for usage or configuration errors; plain success names the resulting phase.

### APP-030
verify --quick MUST use quick checks without committing.

### APP-031
verify and done --json MUST return schema 1 with result data on success or null on failure and reported code/message diagnostics.

### APP-032
context --json MUST return identical workflow.Context text as schema-1 string data.

### APP-033
context MUST return exit 2 for usage or configuration errors and exit 1 for held-state lookup failures, with corrective code/message diagnostics and null JSON data.

### APP-034
coverage MUST map results to IDs and require coverage and pass, returning exit 0 on success or exit 1 on failed checks.

### APP-035
coverage MUST return exit 2 with corrective diagnostics for usage, configuration or unknown module/change errors.

### APP-036
coverage --json MUST return schema 1 with report data on success, null on failure and code/message diagnostics.

### APP-037
Existing static failures exposed by the CLI MUST retain their D006 cause: C01 configuration, C02 graph, C03 syntax, C04 normative text, C05 identities, C06 tags, C07 change structure, C08 bases, C09 handles, C12 metadata or C13 queue identity.

### APP-038
CLI lookup rejections for unknown names or ambiguous projected change-local handles MUST use U02. Invalid graphs, module spec identities and other static validation failures MUST retain their D006 static codes.

### APP-039
Existing operational precondition failures without a more specific D006 cause MUST report P01.

### APP-040
For workflow or checks failures carrying an assigned code, the CLI MUST preserve that code and reported warnings.

### APP-041
init MUST create fireweed.yaml in the current directory with format version 1, a binary range containing assets.Version, base main, remote origin and modules [packages/*].

### APP-042
init MUST write the selected recipe's module_defaults, hooks, protected and worker_writable values into the generated configuration.

### APP-043
init MUST create an editable SYSTEM.md project scaffold when absent and retain an existing SYSTEM.md unchanged.

### APP-044
init MUST create or update the fireweed managed AGENTS.md block, pointing to author-change and context, while preserving text outside the markers.

### APP-045
init --recipe go MUST copy assets.BuiltinRecipe(go)'s Files into their project-relative destinations without changing their contents.

### APP-046
Successful init MUST print the selected recipe's instructions on stdout and return exit 0 with empty stderr.

### APP-047
init --recipe none MUST write minimal setup without language hooks, templates or canaries and print the eval, boundaries and setup hook contract.

### APP-048
init --recipe ./directory MUST load recipe.yaml module_defaults, hooks, protected and worker_writable metadata and recipe.md instructions from that directory.

### APP-049
init with a local recipe MUST copy the directory's files/ contents into matching project-relative destinations and copy its instructions to .fireweed/recipes/local.md.

### APP-050
By default init MUST create a normal packages/app module named app with prefix APP and no dependencies, rendering available module templates with name app, prefix APP and package app.

### APP-051
init --no-app MUST skip creation of the app module.

### APP-052
When the selected recipe supplies no module templates, init MUST create a minimal app SPEC.md with empty Specs and Decisions sections.

### APP-053
init MUST configure codex or claude command presets, or preserve a custom agent command template containing {prompt}, in agent.command.

### APP-054
init MUST enable audit by default and write a separately selected --audit-agent command when supplied, otherwise leaving audit.command unset for agent fallback.

### APP-055
init MUST reject missing --recipe or --agent values, malformed options or extra positional arguments with corrective U01 and exit 2 before creating files; missing-input messages MUST list valid recipe and agent choices.

### APP-056
init MUST return U02 and exit 2 for an unknown recipe name, C03 for unreadable or malformed local recipe documents, and C12 for invalid local recipe metadata types, without creating setup files.

### APP-057
init MUST refuse an existing fireweed.yaml with P01 and exit 2, preserving existing project files.

### APP-058
init MUST refuse conflicting generated recipe or module files with P01 and exit 2 rather than overwrite them.

### APP-059
init MUST retain existing .gitignore entries and ensure .fireweed/runs/ is ignored so lifecycle reports do not become forbidden worktree edits.

### APP-060
init MUST copy setup and print instructions without executing configured agents or hooks.

### APP-061
init --json MUST return schema 1 with instruction text as string data and empty diagnostics on success, or null data and code/message diagnostics on failure, with empty stderr.

### APP-062
doctor MUST discover the nearest ancestor fireweed.yaml and invoke checks.Doctor for that project root.

### APP-063
doctor MUST supply the actual invoking executable as DoctorRequest.Bin so scratch hooks query the running NEW CLI.

### APP-064
doctor --no-agent MUST skip configured worker and audit commands while retaining hook canary and recovery checks.

### APP-065
Plain doctor output MUST identify each reported probe's outcome and retain reported diagnostic codes, warning status and named failure details.

### APP-066
doctor MUST return exit 0 when no error diagnostics remain, including warning-only results, exit 1 for static or probe failures, and exit 2 for usage or configuration-loading errors.

### APP-067
doctor --json MUST emit schema 1 with an object containing a checks array on success or null data on failure, retaining code/message/path/warning diagnostics and leaving stderr empty.

### APP-068
agents-md --write MUST discover the nearest ancestor fireweed.yaml and refresh AGENTS.md at that project root.

### APP-069
agents-md --write MUST regenerate the current fireweed managed block while preserving text outside its markers exactly.

### APP-070
agents-md --write MUST create AGENTS.md when absent or append the managed block when neither marker exists.

### APP-071
agents-md --write MUST reject unpaired managed-block markers with corrective P01 and exit 2, preserving the existing text.

### APP-072
agents-md --write MUST print confirmation and exit 0 on success, or exit 2 with corrective U01 for usage, C01 for missing/invalid fireweed.yaml or P01 for operational file/setup failures.

### APP-073
agents-md --write --json MUST emit schema 1 with confirmation string data and empty diagnostics on success, or null data and code/message diagnostics on failure, leaving stderr empty.

### APP-074
next MUST discover the nearest ancestor fireweed.yaml and invoke runner.Next
for that project, supplying the actual invoking executable as Bin.

### APP-075
next MUST accept --allow-unmerged, passing true only when present. It MUST
reject extra positional arguments and unsupported options with corrective U01
and exit 2; --allow-unmerged MUST be rejected for unrelated implemented commands.

### APP-076
Plain next MUST identify the applied change or empty queue on stdout after
successful runner completion.

### APP-077
next MUST return exit 0 for an applied change or empty queue, exit 1 for runner
execution or lifecycle failure, and exit 2 for usage or configuration-loading
failure.

### APP-078
next failure diagnostics MUST retain existing cause codes from wrapped
dependency errors and include a corrective action.

### APP-079
next --json MUST emit schema 1 with object data containing change, run, phase
and empty on success, or null data with code/message diagnostics on failure.
It MUST leave stderr empty.

## Decisions

[0069-project-initialization] Init accepts codex and claude presets as argv commands; other selections are shell command templates containing {prompt}, stored as strings for one-pass quoted substitution when execution is implemented. Local recipes use recipe.yaml metadata, recipe.md instructions and a files/ project-relative payload; none uses empty language defaults and a generic module scaffold.

[0074-doctor-setup-commands] Doctor CLI leaves request Timeout unset, using the backend's existing defaults;
no new timeout flag is introduced. Successful doctor JSON checks use fields name,
module, passed and detail, with diagnostics in the outer envelope.
