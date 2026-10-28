---
bases: {}
module: checks
---



## Motivation

Milestone 2 needs doctor to prove the project-owned glue written after init.
The design's Hooks/recipes/doctor and Worktrees clauses require scratch canary
detection and recovery, bounded agent probes, and visible setup guidance.
Implement the checks-owned backend now, then wire it promptly through app after
assets explains its emitted diagnostics. This is feature work, not another drill.

## Interface

Add checks.Doctor(context.Context, DoctorRequest) DoctorReport with thin aliases
and forwarding in api.go, and the public data under model/:

DoctorRequest has Root string, Bin string, NoAgent bool and Timeout time.Duration.
DoctorCheck has Name string, Module string, Passed bool and Detail string.
DoctorReport has Checks []DoctorCheck and Diagnostics []project.Diagnostic.
Checks record probe outcomes; diagnostics carry errors and warnings. No separate
error return is needed: a non-warning diagnostic means doctor failed. Keep all
implementation internal. Existing public signatures remain unchanged.

## Added

### new:doctor-project
Doctor MUST retain project.Check diagnostics and prevent hook or agent probes when the current project has a static error.

### new:doctor-scratch
Doctor MUST run probes using current project files in a detached scratch worktree below the common Git directory's fireweed/worktrees, leaving the caller's checkout unchanged.

### new:doctor-cleanup
Doctor MUST remove its scratch worktree and Git registration after successful or failed probes.

### new:doctor-setup
Doctor MUST run the configured setup hook in scratch before probes; a setup failure MUST produce C14 naming the setup hook and prevent later probes.

### new:doctor-modules
Doctor MUST render two uniquely named throwaway modules from project-owned module templates and the version-1 canary baseline, substituting name, prefix, package, go_module, module_dir and other_dir.

### new:doctor-inputs
Doctor MUST report C14 naming missing or malformed module templates or required canary maps and skip the affected probes.

### new:doctor-baselines
Doctor MUST require passing baseline evals, coverage and applicable boundaries before planting canaries and after removing each planted canary, reporting C14 with the failed hook and canary when recovery fails.

### new:doctor-undeclared
Doctor MUST plant the supplied undeclared canary and require its boundary hook to report a module-attributed violation; a missed canary MUST produce C14 naming the hook and canary.

### new:doctor-internal-eval
Doctor MUST plant the supplied internal_eval canary and require its boundary hook to report a module-attributed interface violation; a missed canary MUST produce C14 naming the hook and canary.

### new:doctor-failing-eval
Doctor MUST require the supplied failing_eval canary to produce a failed referencing testcase and nonzero eval exit; compile errors, missing reports and unrelated failures MUST NOT count as detection.

### new:doctor-missing-id
Doctor MUST require the supplied missing_id canary to be rejected for its unmapped testcase ID; unrelated hook or report failures MUST NOT count as detection.

### new:doctor-hook-overrides
Doctor MUST test applicable global eval and boundary hooks and every per-module override through their corresponding canary probes.

### new:doctor-hooks-bin
Doctor MUST set probe hooks' FIREWEED_BIN to the request Bin and make the scratch canary modules discoverable through that executable's modules output.

### new:doctor-boundary-warnings
Doctor MUST retain each real module's reasoned X03 boundaries-opt-out warning and applicable X04 warning when no boundary enforcement is required.

### new:doctor-agents
Doctor MUST run the worker and audit commands once each in scratch through the configured agent.wrap with a fixed-reply prompt, requiring exit 0 and the requested reply for success.

### new:doctor-agent-fallback
Doctor MUST use agent.command for the audit probe when audit.command is unset and emit a C15 warning identifying the shared-command blind spot.

### new:doctor-agent-failures
Doctor MUST bound each agent probe by the request Timeout or three minutes when unset, and report C15 failures with the command and its final 20 output lines.

### new:doctor-agent-limits
Doctor MUST classify a configured limit_patterns match in an agent probe's final 20 output lines as a C15 usage-limit diagnostic rather than an ordinary command failure.

### new:doctor-no-agent
Doctor with NoAgent MUST skip worker and audit command execution while retaining hook canary and recovery checks.

### new:doctor-agents-md
Doctor MUST emit a C16 warning when the AGENTS.md managed block is missing or differs from the current fireweed guidance, while accepting a current block regardless of surrounding text.

### new:doctor-playbook-base
Doctor MUST emit a C16 warning when a project playbook override's embedded_base is missing or differs from the SHA-256 of its current embedded playbook; a matching base MUST avoid that warning.

## Modified



## Removed



## Decisions

Doctor requests the literal reply FIREWEED_DOCTOR_OK; limit_patterns match
case-insensitive literal substrings. Project playbook overrides use frontmatter
embedded_base containing the hexadecimal SHA-256 of the embedded playbook text.

## Rejected alternatives



## Notes

Read the current design doctor/worktree/AGENTS/playbook clauses, existing checks
interfaces and project-owned Go recipe through assets.BuiltinRecipe. D007 assigns
C14 hook/canary setup/detection/recovery failures, C15 agent failures/usage limits
and fallback warning, and C16 stale managed guidance/playbook-base warnings.
Retain existing project diagnostics and X03/X04 warnings where applicable.
Report named outcomes for the four manifest canaries and their clean baseline/
recovery checks, so app can present useful results. Do not hide missed canaries
behind success because a hook merely exited 0 or returned an unrelated failure.

Reuse the existing Evaluate, coverage/result checks, Boundaries, public graph
and gitstore process/repository primitives. Use project-owned
.fireweed/templates/module/ and .fireweed/canaries/canaries.yaml, with the Go
recipe's version-1 baseline/undeclared/internal_eval/failing_eval/missing_id
shape. Render .tmpl names without that suffix. The scratch graph must include
the generated modules even when original module patterns differ; use the
project's actual module metadata as the graph source. When Go recipe placeholders
need go_module, derive it from project go.mod; other local recipes need no Go
module file. Static configuration errors must
be diagnosed before executing hooks. A normal Git fixture has a committed HEAD;
tests do not need unborn/partial/promisor/unsupported Git cases.

Doctor must see current project glue, including ordinary tracked/untracked
non-ignored setup edits, without mutating the caller. Keep copying and cleanup
ordinary and bounded; no cache snapshot, byte observers, instrumentation or
platform framework. Scratch worktrees live in the common Git directory and run
setup. Use gitstore.Run timeouts/process groups rather than a second process
supervisor. Use request Timeout for bounded probes; when unset use three minutes.
Agent shell templates and argv expand prompt and git_common_dir once, quoting
shell substitutions. agent.wrap is an argv prefix with the same placeholders.
Keep existing hook placeholder behavior and all prior assertions/signatures.

NoAgent must permit deterministic hook-only verification. Agent tests use local
fake commands with short timeouts, fixed replies, wrappers and controlled final
output; they never authenticate or invoke real agents. Real rebuild workers and
auditors keep the configured Sol/high command. Reuse a small number of fixture
builders and representative scenarios, including one real Go recipe success/
recovery fixture, a zero-success hook that misses canaries, a representative
per-module override, failed recovery and bounded fake-agent failures/usage.
Group independent output assertions in those fixtures, not a Cartesian product
or one fixture per spec. Keep own Verify below five minutes and full below
fifteen; no new process-control/instrumentation probes are needed.

The eval step may add the declared interface/model types and a compiling stub;
all new IDs need runtime failure evidence, not compilation or collection errors.
Implement only checks internal code in impl; evals and interfaces freeze after
evals. Do not change app, assets, templates, setup tooling or planning files.
The next authorized slices explain these newly emitted doctor codes in assets
and wire doctor through app. Do not run any drill or fix the deferred contracts
drill gap. Controller owns phase commits/spec application/archive; workers never
commit.
