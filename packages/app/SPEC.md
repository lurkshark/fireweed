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
Without --from-attempt, start <change> MUST discover the root, invoke
workflow.Start and print first-step workflow.Context with exit 0, assigning IDs,
run and phase without committing.

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
Plain next MUST identify the applied change or empty queue after successful
runner completion. On a clean usage stop it MUST name the held change and show
the wait/budget reason with instructions for resuming it.

### APP-077
next MUST return exit0 for an applied change or empty queue, exit1 for ordinary
runner execution or lifecycle failure, exit2 for usage or configuration errors,
and exit3 for a clean usage stop.

### APP-078
next failure diagnostics MUST retain existing cause codes from wrapped
dependency errors and include a corrective action.

### APP-079
next --json MUST emit schema1 and leave stderr empty. It MUST use object data
containing change, run, phase and empty on success or a clean usage stop, with
a nonempty reason reporting completed wait, next wait and budget on a stop.
On failure it MUST use null data with code/message diagnostics.

### APP-080
run MUST discover the nearest ancestor fireweed.yaml and invoke runner.Run for
that project with the actual invoking executable as Bin.

### APP-081
run --max N MUST accept a positive integer and supply it as Max, using zero
when --max is absent. It MUST reject missing, noninteger or nonpositive values,
extra positional arguments and unsupported options with corrective U01 and
exit 2; --max MUST be rejected for unrelated implemented commands.

### APP-082
run MUST accept --allow-unmerged and supply AllowUnmerged true only when the
flag is present, using false by default.

### APP-083
After successful runner completion, plain run MUST identify applied changes
and their count, including whether the queue is empty or Max left queued work.
On a clean usage stop it MUST retain the applied names/count and show the
wait/budget reason with instructions for resuming.

### APP-084
run MUST return exit0 for successful queue completion, Max termination or an
empty queue, exit1 for execution or lifecycle failure including any earlier
blocked change, exit2 for usage or configuration errors, and exit3 for a clean
time or usage stop.

### APP-085
run failure diagnostics MUST retain existing cause codes from wrapped
dependency errors and include a corrective action.

### APP-086
run --json MUST emit schema1 and leave stderr empty. It MUST use object data
containing applied change/run/phase summaries, empty and limited on success or
a clean usage stop, with a nonempty reason reporting completed wait, next wait
and budget on a stop. On failure it MUST use null data with code/message
diagnostics.

### APP-087
start <change> --from-attempt N MUST discover the project root and invoke
workflow.StartFromAttempt with the selected change and N. On success it MUST
return workflow.Context in the existing plain or JSON start shape with exit 0.

### APP-088
For start, --from-attempt MUST require one positive integer and reject missing,
noninteger, nonpositive or repeated values with corrective U01 and exit 2 before
mutation. Its value diagnostics MUST describe valid start usage. Other commands
MUST reject --from-attempt with U01 and exit 2.

### APP-089
run --for D MUST require one positive Go duration and pass it as RunRequest.For,
using zero when absent. It MUST reject missing, malformed, nonpositive or repeated
values and use by other commands with corrective U01 and exit2 before mutation,
showing valid run usage.

### APP-090
On a clean runner time stop, run MUST return exit3 with retained applied summaries
and a nonempty reason naming the time limit. Plain output MUST include the
applied count and resume guidance. JSON MUST use the existing schema1 run object
with reason and empty/limited flags, leaving stderr empty.

### APP-091
release <change> MUST discover the nearest project root and invoke
workflow.Release for the named change. On success it MUST print confirmation
with exit0, preserving work, runtime state and HEAD while dropping only the
owned claim.

### APP-092
release MUST return exit1 with the workflow cause code and corrective action
when claim release fails, including missing or foreign claims. It MUST return
exit2 with a corrective diagnostic for usage or project configuration errors.

### APP-093
release --json MUST emit schema1 with confirmation string data and empty
diagnostics on success, or null data with code/message diagnostics on failure.
It MUST leave stderr empty.

### APP-094
commits <change> MUST discover the nearest project root and invoke
workflow.Commits. Plain success MUST print the returned object IDs one per line
in returned order with exit0, returning empty output when the list is empty.

### APP-095
commits MUST return exit2 with corrective U01 for usage errors, U02 for selector
errors, and the existing cause code for project configuration errors. Other
workflow failures MUST return exit1 retaining their cause code and a corrective
action.

### APP-096
commits --json MUST emit schema1 with an array of returned object IDs, including
an empty array for no commits, and empty diagnostics on success. On failure it
MUST emit null data with code/message diagnostics. It MUST leave stderr empty.

### APP-097
abandon <change> MUST discover the nearest project root and invoke workflow.Abandon for the named change.
Plain success MUST print confirmation naming the change and returned full commit object ID with exit0.

### APP-098
abandon MUST return exit2 with corrective U01 for usage errors, U02 for selector errors, and the existing cause code for project configuration errors.
Other workflow failures MUST return exit1 retaining their cause code and a corrective action.

### APP-099
abandon --json MUST emit schema1 with confirmation string data and empty diagnostics on success.
On failure it MUST emit null data with code/message diagnostics.
It MUST leave stderr empty.

### APP-100
new-module <name> --prefix <P> MUST discover the nearest project root and create the named module from .fireweed/templates/module in a new directory matched by the configured module globs.
It MUST use that project-owned copy rather than replace it with embedded templates.

### APP-101
Module creation MUST render template contents using name, prefix, package, module_dir and go_module substitutions, removing a terminal .tmpl suffix from output filenames.
The resulting SPEC.md MUST declare the requested module name and prefix and be loadable as part of the project graph.

### APP-102
new-module MUST reject an existing module name, existing ID prefix or occupied destination before writing module files, preserving the existing module and destination contents.

### APP-103
new-module MUST return exit2 with corrective U01 for invalid arguments and options, the existing cause code for project errors, C03 for missing or malformed templates, and P01 for conflicting module identity or destination.
It MUST reject requests before writing module files when these errors are present.

### APP-104
Successful new-module MUST return exit0 and print the repository-relative module directory.
With --json it MUST emit schema1 string data equal to that directory and empty diagnostics; failures MUST emit null data and code/message diagnostics with empty stderr.

### APP-105
check-commits [range] MUST discover the nearest project root and call checks.History with the supplied range, or an empty range when omitted to use the configured local base.
Plain success MUST return exit0, report the number of checked commits, and list each squashed change with its full commit object ID and change path.

### APP-106
check-commits MUST return exit1 for history-rule violations and report every returned H01–H04 diagnostic with its affected path, full commit object ID and corrective action.

### APP-107
check-commits MUST return exit2 with corrective U01 for invalid arguments/options and the existing cause code for project configuration errors.
Operational history failures MUST return exit1 with their existing cause code, or P01 when uncoded, and a corrective action.

### APP-108
Successful check-commits --json MUST emit schema1 object data with commits in returned order and squashed entries containing commit and change, using empty arrays when appropriate and empty diagnostics.
Failures MUST emit null data with code/message diagnostics.
It MUST leave stderr empty.

### APP-109
ci [range] MUST discover the nearest project root and run static project checks plus history checks on the supplied range, or the configured local-base range when omitted.
Static or history errors MUST prevent a successful CI result.

### APP-110
ci MUST run evals for every discovered module and require passing coverage of its current specs, regardless of verify_dependents.
The configured boundary checks MUST also pass for CI to succeed.

### APP-111
ci MUST run doctor canaries by default with agent probes disabled.
With --no-doctor it MUST skip the doctor canaries while retaining the other deterministic checks.

### APP-112
Successful ci MUST return exit0 with a confirmation.
Usage and project configuration errors MUST return exit2 with corrective guidance, using U01 for usage and the existing cause code for configuration.
Deterministic check failures MUST return exit1 retaining their cause codes and corrective guidance.

### APP-113
Successful ci --json MUST emit schema1 confirmation-string data and collected diagnostics, including warnings.
Failures MUST emit null data with code/message diagnostics.
It MUST leave stderr empty.

### APP-114
list MUST discover the nearest project root and list project.Queue entries in returned order.
Plain success MUST return exit0 and identify every queued change by slug, path and module, with empty output for an empty queue.

### APP-115
Each listed entry MUST report its phase, using queued when the phase is absent.
It MUST report the blocked reason and audit rounds, using an empty reason and zero rounds when absent.

### APP-116
Usage and project configuration errors MUST return exit2 with corrective guidance, using U01 for usage and the existing cause code for configuration.
Queue failures MUST return exit1 with their existing cause code and corrective guidance.

### APP-117
Successful list --json MUST emit schema1 ordered-array data with slug, path, module, phase, blocked and audit_rounds fields, including an empty array for an empty queue, and empty diagnostics.
Failures MUST emit null data with code/message diagnostics.
It MUST leave stderr empty.

### APP-118
status MUST discover the nearest project root and expose workflow.Status for it.
Successful status MUST return exit0, including empty queues and warning-only reports.

### APP-119
Plain status MUST identify every queued entry in returned order by slug, path and module, with its phase, blocked reason, audit rounds, waiting change and blocking rule.
It MUST report an empty queue explicitly when there are no entries.

### APP-120
Plain status MUST identify discovered module names, prefixes, directories, dependencies and resolved zones.
It MUST report quarantines with module/ID/reason, boundary opt-outs with module/reason, each module's drift, the unapproved count and each archived change's phase-commit count from the returned report.

### APP-121
Plain status MUST show every returned module cost with its JUnit seconds, testcase counts per spec, eval lines and disposable lines.
Modules without a recorded cost MUST not receive a fabricated measurement.

### APP-122
Successful status --json MUST emit schema1 object data containing modules, queue, quarantines, boundary_opt_outs, drift_since_drill, unapproved, phase_commits and metrics.
Empty array/map containers MUST appear as arrays/maps rather than null.
It MUST leave stderr empty.

### APP-123
JSON module rows MUST use the existing name, id_prefix, dir, depends_on and zones shape.
Queue rows MUST use slug, path, module, phase, blocked, audit_rounds, waiting_on and waiting_rule.
Quarantine and opt-out rows MUST use module/id/reason and module/reason fields respectively.

### APP-124
JSON metrics MUST use module, seconds, cases_per_spec, eval_lines and disposable_lines fields.
The cases_per_spec, drift_since_drill and phase_commits maps MUST retain their returned identifier/count associations.

### APP-125
status MUST retain the returned diagnostic codes and messages, including budget and P08 warnings.
JSON diagnostics MUST retain path when present and warning status; plain output MUST show each diagnostic code and message.

### APP-126
Usage and project configuration errors MUST return exit2 with corrective guidance, using U01 for usage and the existing configuration cause code.
Status failures MUST return exit1 with corrective guidance, retaining assigned cause codes and using P01 for an unavailable base ref.
JSON failures MUST use null data with code/message diagnostics and empty stderr.

### APP-127
pr-body MUST discover the nearest project root and expose workflow.PRBody for the supplied explicit range.
Requested endpoint data MUST remain authoritative when later HEAD or working-tree metadata differs.

### APP-128
Plain pr-body MUST print the report Body and show every returned diagnostic code and message.
Successful reports MUST return exit0, including empty selections and warning-only reports.

### APP-129
Successful pr-body --json MUST emit schema1 with string data equal to the report Body and retained diagnostic codes, messages, warning status and path when present.
It MUST leave stderr empty.

### APP-130
Missing or extra ranges and invalid options MUST return exit2 with U01 and corrective guidance.
Project discovery or endpoint configuration-loading failures MUST return exit2 with the existing cause code or C01 and corrective guidance.
Unresolvable ranges MUST return exit1 with P01, the range resolution cause and corrective guidance.

### APP-131
JSON pr-body failures MUST emit null data with cause-code/message diagnostics and empty stderr.

## Decisions

[0069-project-initialization] Init accepts codex and claude presets as argv commands; other selections are shell command templates containing {prompt}, stored as strings for one-pass quoted substitution when execution is implemented. Local recipes use recipe.yaml metadata, recipe.md instructions and a files/ project-relative payload; none uses empty language defaults and a generic module scaffold.

[0074-doctor-setup-commands] Doctor CLI leaves request Timeout unset, using the backend's existing defaults;
no new timeout flag is introduced. Successful doctor JSON checks use fields name,
module, passed and detail, with diagnostics in the outer envelope.

[0112-new-module-command] Creation uses the first configured module glob ending in a single wildcard
component under a literal parent, placing the requested name in that component;
a configuration with no such destination is reported as C01 before writing.
The package substitution replaces name hyphens with underscores; go_module is
read from go.mod when present and is otherwise empty.
