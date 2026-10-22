---
depends_on:
    - contracts
    - gitstore
    - assets
    - project
id_prefix: CHK
interface:
    - api.go
    - model/*.go
name: checks
---

# checks

## Specs

### CHK-001
ParseJUnit MUST accept well-formed testsuite or testsuites XML, including nested suites, and preserve actual testcase document order, names, classnames, every failure/error message and text, and skipped presence/message/text. It MUST ignore suite aggregate counts and unrelated report metadata when collecting cases, return a non-null empty slice for an empty valid suite, and reject empty input, malformed or truncated XML, a wrong root, multiple roots or non-whitespace outside the single root atomically with no partial cases. Parsed IDs MUST start as a non-null empty slice; analysis, rather than XML-provided IDs, determines references.

### CHK-002
Analyze MUST derive each testcase's canonical unique sorted IDs from both name and classname using the exact public MatchesID semantics, accepting supported case/separator spellings and rejecting misleading letter/digit/Unicode boundaries. It MUST recognize candidates for the target module prefix even when prefixes is empty, and for every supplied project prefix, while ignoring duplicate prefixes. It MUST replace untrusted incoming IDs, map each recognized ID to all matching zero-based testcase indexes in report order (including skips), and emit C12 for a testcase with no recognized ID or any recognized ID absent from the target module's current specs, including removed IDs and another module's IDs. Every emitted diagnostic MUST have severity error, the target module directory as Path and a nonempty corrective message.

### CHK-003
Analyze MUST emit C13 for every current enforced spec without a non-skipped referencing testcase, including when the only matching cases are skipped. An executed testcase with any failure or error element MUST count as failed, even when its message/text is empty; skipped cases MUST NOT count as executed or failed even if they also contain failure/error elements. For each enforced failFirst ID, at least one executed matching failure/error MUST be required (P1-02 otherwise), unless that ID is explicitly in already; already MUST waive only fail-first and MUST NOT waive executed coverage or ordinary passing requirements. Failures of other enforced specs MUST emit C14, so unrelated failures cannot prove fail-first. review and reasoned quarantine tags, alone or combined, MUST suspend C13, P1-02 and C14 for that spec while preserving its cases, IDs and BySpec visibility. A nonzero hook exit without any executed failing/error testcase, or a zero hook exit with one, MUST emit C11 regardless of tags; aggregate XML counts MUST NOT supply such evidence. Unknown failFirst or already IDs MUST emit C12 rather than granting exemptions.

### CHK-004
ParseJUnit and Analyze MUST return independently owned results across calls. Analyze MUST preserve all input module, prefix, testcase/message/ID slices and policy lists, return non-null Tests, BySpec and Diagnostics even when empty, preserve testcase evidence and hookExit, set Module to the target name, and sort diagnostics deterministically by Path, Code, Severity and Message. No returned nested testcase/message/ID slice, skipped message, BySpec index slice or diagnostics slice MAY alias an input or another result, including empty slices backed by spare capacity. Duplicate failFirst/already entries MUST NOT duplicate policy diagnostics. Repeated identical calls MUST produce equal evidence.

### CHK-005
RunHook MUST execute the supplied command string through sh -c at the canonical repository root found from snapshot.Root. It MUST replace each recognized {module}, {module_dir}, {report} and {git_common_dir} placeholder exactly once with one safely shell-quoted literal value, preserving spaces, apostrophes, metacharacters and placeholder-looking substrings inside values; unknown placeholders MUST remain literal. module name/directory MUST come from the supplied module, report MUST retain its supplied value, and git_common_dir MUST be the canonical shared Git directory, including in linked worktrees. A nil module MUST supply empty module/name-directory values. The child MUST inherit the parent environment but override FIREWEED_MODULE, FIREWEED_MODULE_DIR, FIREWEED_REPORT and FIREWEED_BIN with those values and the currently executing binary's actual path, regardless of stale parent values. No bootstrap executable MAY be substituted for that identity.

### CHK-006
RunHook MUST use the local process adapter's complete combined output, actual exit/Err and literal shell argv semantics, supply EOF rather than interactive stdin, and bound execution by both the parent context and snapshot.Config.Timeout, whichever ends first. Context cancellation/deadline MUST terminate the hook process tree and preserve emitted output. Nil context, blank commands, invalid/nonpositive configured timeouts and repository discovery failures MUST fail with Exit 2 and useful Err before hook startup; a pre-cancelled valid context MUST prevent startup and retain the process adapter's Exit -1/context error contract. RunHook MUST preserve all snapshot/module/command/report input data and parent environment/current directory across successful and failed calls, and overlapping calls MUST NOT mix output, identity, reports or cancellation state.

### CHK-007
Evaluate MUST execute only the supplied module's resolved eval hook, create a unique fresh absolute report path in an owned temporary directory, pass it consistently as {report} and FIREWEED_REPORT, and remove that directory on success and every failure path. A missing/blank eval hook MUST yield C07 without executing a command. Missing, unreadable or invalid JUnit, hook launch/configuration/context/signal failures MUST yield C11, with the selected module directory as Path and a nonempty corrective message. Such invalid evidence MUST return non-null empty Tests and BySpec rather than partial coverage, preserve any captured Output and HookExit, and MUST NOT be accepted as successful fail-first even if a cancelled hook had already written a syntactically valid failing report. Ordinary nonzero process exits with valid reports MUST remain available for real failed/errored testcase evidence rather than being treated as launch failures.

### CHK-008
For a valid report and ordinary completed hook, Evaluate MUST feed actual parsed cases, actual hook exit, the target module, every snapshot module prefix and the supplied failFirst/already lists into Analyze, retaining its exact testcase evidence, ID mappings, policy diagnostics and deterministic order. It MUST set Output and HookExit to the real hook result and preserve the selected module name. It MUST honor resolved per-module hooks rather than silently reverting to the project default. Returned results MUST remain independently owned across calls, and Evaluate MUST preserve all caller inputs and parent process state. It MUST NOT independently invoke setup/boundaries hooks, agents, the bootstrap or network commands, create Git commits/refs, or leave report artifacts in the repository; explicitly authored eval-hook actions remain the caller's hook responsibility.

## Decisions

[junit-evidence]
This slice handles in-memory evidence only. Later checked hooks will attach output and convert missing/invalid report and invocation errors to C11 before composing these checks. BySpec records all references for inspection; coverage and fail-first separately exclude skips. Evidence contradictory to the process exit is always a reporter error, even for review/quarantine specs; those tags suspend spec enforcement, not report consistency. The already list only relaxes the failing-before requirement for a changed spec, and cannot excuse a failing unchanged spec or missing executed coverage. Known project prefixes allow foreign-module references to be reported explicitly as unknown to the selected module; an unrecognized spelling still fails C12 for having no recognized ID. Diagnostics use the established C11–C14 and P1-02 catalog meanings.

Missing executed coverage and missing failing-before evidence are distinct violations: an unenforced exemption is never inferred from C13. For an enforced failFirst ID with no executed testcase and no already exemption, return both C13 and P1-02. Ordinary reference inspection of coverage.go shows a C13 early-continue suppressing P1-02; the authored contract deliberately reports both checks rather than hiding the second unmet requirement.

[hook-evaluation]
Use one-pass placeholder expansion so a value containing another placeholder's spelling cannot trigger a second substitution. Shell quoting belongs here; argv commands and later agent substitutions remain literal at their own boundaries. FIREWEED_BIN is os.Executable for the current process; when invoked through the new CLI it names that independent executable, and public-API eval probes name their own executable. This is never the governing bootstrap path inherited by rebuild hooks.

Evaluate is a low-level selected-module operation. App/workflow callers will choose current or projected specs and supply fail-first IDs explicitly. It neither loads/reprojects an entire queue nor infers an exemption. Missing/malformed/fatal process evidence stops before Analyze; otherwise existing CHK-001–004 determine coverage, including reporter exit inconsistencies. A timeout with a valid failure report is still invalid hook evidence and always C11. Reports are fresh per call and cleaned independently of authored hook side effects. Boundary hook execution remains a separate change.

Fixture commands must distinguish template text from values: use ordinary paths for source/marker literals embedded in a command, and inject placeholder-looking paths through placeholder values when testing one-pass expansion. The initial incomplete cycle accidentally embedded {report} in those original template literals; its four phase commits were fully reverted without applying specs or editing frozen evals. Restart with a direct helper positive control and corrected fixtures.
