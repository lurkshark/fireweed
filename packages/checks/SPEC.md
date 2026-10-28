---
depends_on:
    - contracts
    - gitstore
    - assets
    - project
id_prefix: CHK
interface:
    - api.go
    - model/**
name: checks
---

# checks

## Specs

### CHK-001
ParseJUnit MUST accept testsuite or testsuites documents, including nested suites, returning testcase name, classname and time.

### CHK-002
RequireCoverage MUST cover each current spec except review or quarantine with a non-skipped referencing case. Missing coverage MUST return a C10-prefixed error.

### CHK-003
Evaluate MUST run the configured eval hook, accepting shell strings or argv.

### CHK-004
During audit, Zones MUST allow edits only to the current change file.

### CHK-005
Boundaries MUST run applicable global and per-module boundary hooks.

### CHK-006
Zones MUST reject edits to protected paths in every phase.

### CHK-007
Zones MUST reject edits to module SPEC.md files in every phase.

### CHK-008
Zones MUST reject edits within another module in every phase.

### CHK-009
Zones MUST reject edits to change files other than the current change in every phase.

### CHK-010
During evals and impl, Zones MUST allow edits outside modules only when the path matches worker_writable, except for the current change file.

### CHK-011
During evals, Zones MUST allow edits in the target module's evals zone.

### CHK-012
During evals, Zones MUST allow edits in the target module's interface zone only when InterfaceDeclared is true.

### CHK-013
During evals, Zones MUST allow edits in the target module's disposable zone only for files absent from the phase base.

### CHK-014
During evals, Zones MUST allow edits in the target module's other zone.

### CHK-015
During evals, Zones MUST reject edits in the target module's preserve zone.

### CHK-016
During impl, Zones MUST reject edits in the target module's evals zone.

### CHK-017
During impl, Zones MUST reject edits in the target module's interface zone.

### CHK-018
During impl, Zones MUST allow edits in the target module's disposable zone.

### CHK-019
During impl, Zones MUST allow edits in the target module's preserve zone.

### CHK-020
During impl, Zones MUST allow edits in the target module's other zone.

### CHK-021
During evals and impl, Zones MUST allow edits to the current change file.

### CHK-022
During impl, Zones MUST allow only newly added files in the target module's append_only zone.

### CHK-023
During evals, Zones MUST reject edits in the target module's append_only zone.

### CHK-024
Zones MUST check both the source and destination of a rename against phase permissions.

### CHK-025
RequireFailFirst MUST find a failed or errored referencing case for every requested enforced ID except already_satisfied; passing, skipped, empty or unrelated failures do not qualify. Missing genuine fail-first evidence MUST return a P2-01-prefixed error.

### CHK-026
RequirePass MUST reject failed or errored cases referencing any non-quarantined spec with an E01-prefixed error.

### CHK-027
ParseJUnit MUST reject malformed XML with a C11-prefixed error.

### CHK-028
ParseJUnit MUST reject testcases with missing or unknown spec IDs with a C11-prefixed error.

### CHK-029
ParseJUnit MUST map testcase name and classname to current spec IDs using contracts.MatchesID.

### CHK-030
ParseJUnit MUST classify failure and error elements as fail, and skipped elements as skip.

### CHK-031
Evaluate MUST prefer a module's eval hook override over the global eval hook.

### CHK-032
Evaluate MUST expand module, module_dir and report placeholders once in shell strings or argv, quoting substituted shell values.

### CHK-033
Evaluate MUST set the hook's FIREWEED_BIN environment variable to the request Bin.

### CHK-034
Evaluate MUST retain the hook process's stdout, stderr and exit code in its returned report.

### CHK-035
Evaluate MUST parse the hook's written JUnit report into the returned report's cases.

### CHK-036
Evaluate MUST fail with a C11-prefixed error when the eval hook cannot launch or exceeds its timeout.

### CHK-037
Evaluate MUST fail with a C11-prefixed error when the eval hook does not write its JUnit report.

### CHK-038
Evaluate MUST reject a zero exit with mapped failures or a nonzero exit without mapped failures with a C11-prefixed error, while accepting a nonzero exit with mapped failures as valid report evidence.

### CHK-039
Boundaries MUST return errors with diagnostic X01 for unsuppressed module-attributed violations.

### CHK-040
Boundaries MUST return a visible X03 warning containing the reason for a reasoned boundaries_opt_out.

### CHK-041
Boundaries MUST suppress only exact baseline lines from normally exited hooks.

### CHK-042
Boundaries MUST suppress only the opted-out module's violation lines through a reasoned opt-out after a hook exits normally.

### CHK-043
Boundaries MUST report hook launch, cancellation or timeout failures as errors with diagnostic X02.

### CHK-044
Boundaries MUST fail with diagnostic X02 for a missing hook required by a module with disposable globs and no reasoned opt-out.

### CHK-045
Boundaries MUST warn with diagnostic X04 without an error for a missing hook when no boundary enforcement is required.

### CHK-046
Evaluate MUST return C11-prefixed errors for existing eval-hook setup failures, including missing or invalid hook commands.

### CHK-047
Evaluate MUST return a U02-prefixed error for an unknown requested module.

### CHK-048
Evaluate MUST return C03-prefixed errors for malformed module item syntax and C05-prefixed errors for duplicate module item identities.

### CHK-049
Doctor MUST retain project.Check diagnostics and prevent hook or agent probes when the current project has a static error.

### CHK-050
Doctor MUST run probes using current project files in a detached scratch worktree below the common Git directory's fireweed/worktrees, leaving the caller's checkout unchanged.

### CHK-051
Doctor MUST remove its scratch worktree and Git registration after successful or failed probes.

### CHK-052
Doctor MUST run the configured setup hook in scratch before probes; a setup failure MUST produce C14 naming the setup hook and prevent later probes.

### CHK-053
Doctor MUST render two uniquely named throwaway modules from project-owned module templates and the version-1 canary baseline, substituting name, prefix, package, go_module, module_dir and other_dir.

### CHK-054
Doctor MUST report C14 naming missing or malformed module templates or required canary maps and skip the affected probes.

### CHK-055
Doctor MUST require passing baseline evals, coverage and applicable boundaries before planting canaries and after removing each planted canary, reporting C14 with the failed hook and canary when recovery fails.

### CHK-056
Doctor MUST plant the supplied undeclared canary and require its boundary hook to report a module-attributed violation; a missed canary MUST produce C14 naming the hook and canary.

### CHK-057
Doctor MUST plant the supplied internal_eval canary and require its boundary hook to report a module-attributed interface violation; a missed canary MUST produce C14 naming the hook and canary.

### CHK-058
Doctor MUST require the supplied failing_eval canary to produce a failed referencing testcase and nonzero eval exit; compile errors, missing reports and unrelated failures MUST NOT count as detection.

### CHK-059
Doctor MUST require the supplied missing_id canary to be rejected for its unmapped testcase ID; unrelated hook or report failures MUST NOT count as detection.

### CHK-060
Doctor MUST test applicable global eval and boundary hooks and every per-module override through their corresponding canary probes.

### CHK-061
Doctor MUST set probe hooks' FIREWEED_BIN to the request Bin and make the scratch canary modules discoverable through that executable's modules output.

### CHK-062
Doctor MUST retain each real module's reasoned X03 boundaries-opt-out warning and applicable X04 warning when no boundary enforcement is required.

### CHK-063
Doctor MUST run the worker and audit commands once each in scratch through the configured agent.wrap with a fixed-reply prompt, requiring exit 0 and the requested reply for success.

### CHK-064
Doctor MUST use agent.command for the audit probe when audit.command is unset and emit a C15 warning identifying the shared-command blind spot.

### CHK-065
Doctor MUST bound each agent probe by the request Timeout or three minutes when unset, and report C15 failures with the command and its final 20 output lines.

### CHK-066
Doctor MUST classify a configured limit_patterns match in an agent probe's final 20 output lines as a C15 usage-limit diagnostic rather than an ordinary command failure.

### CHK-067
Doctor with NoAgent MUST skip worker and audit command execution while retaining hook canary and recovery checks.

### CHK-068
Doctor MUST emit a C16 warning when the AGENTS.md managed block is missing or differs from the current fireweed guidance, while accepting a current block regardless of surrounding text.

### CHK-069
Doctor MUST emit a C16 warning when a project playbook override's embedded_base is missing or differs from the SHA-256 of its current embedded playbook; a matching base MUST avoid that warning.

## Decisions

[0071-doctor-probes] Doctor requests the literal reply FIREWEED_DOCTOR_OK; limit_patterns match
case-insensitive literal substrings. Project playbook overrides use frontmatter
embedded_base containing the hexadecimal SHA-256 of the embedded playbook text.
