---
module: checks
phase: impl
run: 40404c0dbcfefce003c43eae
---

## Motivation

Coverage evidence must come from a real configured hook, with safe placeholder expansion, explicit product identity and an independent report path. Compose the established process adapter and JUnit analyzer without mixing in lifecycle transitions or boundary checking. Missing/invalid/partial evidence must not become an accidental fail-first pass.

## Interface

Add RunHook(ctx context.Context, snapshot project.Snapshot, module *contracts.Module, command, report string) gitstore.CommandResult and Evaluate(ctx context.Context, snapshot project.Snapshot, module contracts.Module, failFirst, already []string) Coverage to api.go. Thin facades forward into new disposable internal/hooks code; existing retained model types and ParseJUnit/Analyze remain unchanged. Hook command strings run through sh -c, consistent with the reference; a POSIX shell is a configured-hook prerequisite, including on Windows. No implicit shell is added to gitstore.Run itself.

## Added

### CHK-005
RunHook MUST execute the supplied command string through sh -c at the canonical repository root found from snapshot.Root. It MUST replace each recognized {module}, {module_dir}, {report} and {git_common_dir} placeholder exactly once with one safely shell-quoted literal value, preserving spaces, apostrophes, metacharacters and placeholder-looking substrings inside values; unknown placeholders MUST remain literal. module name/directory MUST come from the supplied module, report MUST retain its supplied value, and git_common_dir MUST be the canonical shared Git directory, including in linked worktrees. A nil module MUST supply empty module/name-directory values. The child MUST inherit the parent environment but override FIREWEED_MODULE, FIREWEED_MODULE_DIR, FIREWEED_REPORT and FIREWEED_BIN with those values and the currently executing binary's actual path, regardless of stale parent values. No bootstrap executable MAY be substituted for that identity.

### CHK-006
RunHook MUST use the local process adapter's complete combined output, actual exit/Err and literal shell argv semantics, supply EOF rather than interactive stdin, and bound execution by both the parent context and snapshot.Config.Timeout, whichever ends first. Context cancellation/deadline MUST terminate the hook process tree and preserve emitted output. Nil context, blank commands, invalid/nonpositive configured timeouts and repository discovery failures MUST fail with Exit 2 and useful Err before hook startup; a pre-cancelled valid context MUST prevent startup and retain the process adapter's Exit -1/context error contract. RunHook MUST preserve all snapshot/module/command/report input data and parent environment/current directory across successful and failed calls, and overlapping calls MUST NOT mix output, identity, reports or cancellation state.

### CHK-007
Evaluate MUST execute only the supplied module's resolved eval hook, create a unique fresh absolute report path in an owned temporary directory, pass it consistently as {report} and FIREWEED_REPORT, and remove that directory on success and every failure path. A missing/blank eval hook MUST yield C07 without executing a command. Missing, unreadable or invalid JUnit, hook launch/configuration/context/signal failures MUST yield C11, with the selected module directory as Path and a nonempty corrective message. Such invalid evidence MUST return non-null empty Tests and BySpec rather than partial coverage, preserve any captured Output and HookExit, and MUST NOT be accepted as successful fail-first even if a cancelled hook had already written a syntactically valid failing report. Ordinary nonzero process exits with valid reports MUST remain available for real failed/errored testcase evidence rather than being treated as launch failures.

### CHK-008
For a valid report and ordinary completed hook, Evaluate MUST feed actual parsed cases, actual hook exit, the target module, every snapshot module prefix and the supplied failFirst/already lists into Analyze, retaining its exact testcase evidence, ID mappings, policy diagnostics and deterministic order. It MUST set Output and HookExit to the real hook result and preserve the selected module name. It MUST honor resolved per-module hooks rather than silently reverting to the project default. Returned results MUST remain independently owned across calls, and Evaluate MUST preserve all caller inputs and parent process state. It MUST NOT independently invoke setup/boundaries hooks, agents, the bootstrap or network commands, create Git commits/refs, or leave report artifacts in the repository; explicitly authored eval-hook actions remain the caller's hook responsibility.

## Modified

## Removed

## Decisions

Use one-pass placeholder expansion so a value containing another placeholder's spelling cannot trigger a second substitution. Shell quoting belongs here; argv commands and later agent substitutions remain literal at their own boundaries. FIREWEED_BIN is os.Executable for the current process; when invoked through the new CLI it names that independent executable, and public-API eval probes name their own executable. This is never the governing bootstrap path inherited by rebuild hooks.

Evaluate is a low-level selected-module operation. App/workflow callers will choose current or projected specs and supply fail-first IDs explicitly. It neither loads/reprojects an entire queue nor infers an exemption. Missing/malformed/fatal process evidence stops before Analyze; otherwise existing CHK-001–004 determine coverage, including reporter exit inconsistencies. A timeout with a valid failure report is still invalid hook evidence and always C11. Reports are fresh per call and cleaned independently of authored hook side effects. Boundary hook execution remains a separate change.

Fixture commands must distinguish template text from values: use ordinary paths for source/marker literals embedded in a command, and inject placeholder-looking paths through placeholder values when testing one-pass expansion. The initial incomplete cycle accidentally embedded {report} in those original template literals; its four phase commits were fully reverted without applying specs or editing frozen evals. Restart with a direct helper positive control and corrected fixtures.

## Rejected alternatives

Sequential string replacement can reinterpret placeholder text in paths. A shared report file admits stale success and concurrent corruption. Treating every nonzero exit as a launch failure would reject real fail-first; treating cancelled partial reports as ordinary failure could accept unfinished verification.

## Provenance

- evals | 2026-10-22T17:15:45Z | agent: manual | fireweed: 0.4.0 | playbook: 5867363302d4c874074520977b1627d26991d95ee2fbacb9ff63e827661edd0b | tests: 62 | failing: 28 | passing: 34
- audit | 2026-10-22T17:17:39Z | agent: manual | fireweed: 0.4.0 | playbook: a9c982d24b476621316049b527e6c7bbc319f06b12268af150fa7d8c6df92230 | tests: 0 | failing: 0 | passing: 0
- evals | 2026-10-22T17:18:56Z | agent: manual | fireweed: 0.4.0 | playbook: 5867363302d4c874074520977b1627d26991d95ee2fbacb9ff63e827661edd0b | tests: 62 | failing: 28 | passing: 34
- audit | 2026-10-22T17:19:47Z | agent: manual | fireweed: 0.4.0 | playbook: a9c982d24b476621316049b527e6c7bbc319f06b12268af150fa7d8c6df92230 | tests: 0 | failing: 0 | passing: 0

## Audit

No findings. Checked CHK-005, CHK-006, CHK-007 and CHK-008. Independent restricted same-model auditor rechecked exact public clauses and corrected fixture premises, including the post-write readiness marker and direct helper control. No implementation/history/bootstrap access or test execution.
