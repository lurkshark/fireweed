---
module: checks
phase: impl
run: 22b2367a5ab7064e801a1243
set:
    interface:
        - api.go
        - model/*.go
---

## Motivation

The new product needs trustworthy JUnit evidence before it can govern a change. A compiler error, skipped case, unknown ID or contradictory hook exit cannot stand in for an executed regression. Separate report parsing and coverage analysis from later hook execution so their complete inputs and results can be tested directly.

## Interface

Retain model/*.go for plain public data alongside api.go. Expose Message {Message, Text string}, TestCase {Name, Classname string; Failures, Errors []Message; Skipped *Message; IDs []string}, Diagnostic {Code, Severity, Message, Path string}, and Coverage {Module string; Tests []TestCase; BySpec map[string][]int; Diagnostics []Diagnostic; Output string; HookExit int}, with JSON names matching the reference coverage surface (IDs is spec_ids). ParseJUnit(report []byte) ([]TestCase, error) parses bytes. Analyze(module contracts.Module, prefixes []string, tests []TestCase, hookExit int, failFirst, already []string) Coverage computes evidence. Functions forward through thin interfaces into disposable internal code. No hook invocation is introduced in this change.

## Added

### CHK-001
ParseJUnit MUST accept well-formed testsuite or testsuites XML, including nested suites, and preserve actual testcase document order, names, classnames, every failure/error message and text, and skipped presence/message/text. It MUST ignore suite aggregate counts and unrelated report metadata when collecting cases, return a non-null empty slice for an empty valid suite, and reject empty input, malformed or truncated XML, a wrong root, multiple roots or non-whitespace outside the single root atomically with no partial cases. Parsed IDs MUST start as a non-null empty slice; analysis, rather than XML-provided IDs, determines references.

### CHK-002
Analyze MUST derive each testcase's canonical unique sorted IDs from both name and classname using the exact public MatchesID semantics, accepting supported case/separator spellings and rejecting misleading letter/digit/Unicode boundaries. It MUST recognize candidates for the target module prefix even when prefixes is empty, and for every supplied project prefix, while ignoring duplicate prefixes. It MUST replace untrusted incoming IDs, map each recognized ID to all matching zero-based testcase indexes in report order (including skips), and emit C12 for a testcase with no recognized ID or any recognized ID absent from the target module's current specs, including removed IDs and another module's IDs. Every emitted diagnostic MUST have severity error, the target module directory as Path and a nonempty corrective message.

### CHK-003
Analyze MUST emit C13 for every current enforced spec without a non-skipped referencing testcase, including when the only matching cases are skipped. An executed testcase with any failure or error element MUST count as failed, even when its message/text is empty; skipped cases MUST NOT count as executed or failed even if they also contain failure/error elements. For each enforced failFirst ID, at least one executed matching failure/error MUST be required (P1-02 otherwise), unless that ID is explicitly in already; already MUST waive only fail-first and MUST NOT waive executed coverage or ordinary passing requirements. Failures of other enforced specs MUST emit C14, so unrelated failures cannot prove fail-first. review and reasoned quarantine tags, alone or combined, MUST suspend C13, P1-02 and C14 for that spec while preserving its cases, IDs and BySpec visibility. A nonzero hook exit without any executed failing/error testcase, or a zero hook exit with one, MUST emit C11 regardless of tags; aggregate XML counts MUST NOT supply such evidence. Unknown failFirst or already IDs MUST emit C12 rather than granting exemptions.

### CHK-004
ParseJUnit and Analyze MUST return independently owned results across calls. Analyze MUST preserve all input module, prefix, testcase/message/ID slices and policy lists, return non-null Tests, BySpec and Diagnostics even when empty, preserve testcase evidence and hookExit, set Module to the target name, and sort diagnostics deterministically by Path, Code, Severity and Message. No returned nested testcase/message/ID slice, skipped message, BySpec index slice or diagnostics slice MAY alias an input or another result, including empty slices backed by spare capacity. Duplicate failFirst/already entries MUST NOT duplicate policy diagnostics. Repeated identical calls MUST produce equal evidence.

## Modified

## Removed

## Decisions

This slice handles in-memory evidence only. Later checked hooks will attach output and convert missing/invalid report and invocation errors to C11 before composing these checks. BySpec records all references for inspection; coverage and fail-first separately exclude skips. Evidence contradictory to the process exit is always a reporter error, even for review/quarantine specs; those tags suspend spec enforcement, not report consistency. The already list only relaxes the failing-before requirement for a changed spec, and cannot excuse a failing unchanged spec or missing executed coverage. Known project prefixes allow foreign-module references to be reported explicitly as unknown to the selected module; an unrecognized spelling still fails C12 for having no recognized ID. Diagnostics use the established C11–C14 and P1-02 catalog meanings.

Missing executed coverage and missing failing-before evidence are distinct violations: an unenforced exemption is never inferred from C13. For an enforced failFirst ID with no executed testcase and no already exemption, return both C13 and P1-02. Ordinary reference inspection of coverage.go shows a C13 early-continue suppressing P1-02; the authored contract deliberately reports both checks rather than hiding the second unmet requirement.

## Rejected alternatives

Trusting suite counts permits collection failures or fabricated summary counts to masquerade as executed tests. Joining only against existing IDs would discard removed-ID references. Sharing result backing storage would allow consumers to corrupt evidence used in later checks.

## Provenance

- evals | 2026-10-22T16:21:22Z | agent: manual | fireweed: 0.4.0 | playbook: 5867363302d4c874074520977b1627d26991d95ee2fbacb9ff63e827661edd0b | tests: 29 | failing: 24 | passing: 5
- audit | 2026-10-22T16:25:07Z | agent: manual | fireweed: 0.4.0 | playbook: a9c982d24b476621316049b527e6c7bbc319f06b12268af150fa7d8c6df92230 | tests: 0 | failing: 0 | passing: 0
- evals | 2026-10-22T16:26:32Z | agent: manual | fireweed: 0.4.0 | playbook: 5867363302d4c874074520977b1627d26991d95ee2fbacb9ff63e827661edd0b | tests: 32 | failing: 26 | passing: 6
- audit | 2026-10-22T16:27:11Z | agent: manual | fireweed: 0.4.0 | playbook: a9c982d24b476621316049b527e6c7bbc319f06b12268af150fa7d8c6df92230 | tests: 0 | failing: 0 | passing: 0

## Audit

No findings. Checked CHK-001, CHK-002, CHK-003 and CHK-004. Independent restricted same-model auditor confirmed all seven prior findings resolved by concrete public assertions; no implementation/history/bootstrap access or test execution.
