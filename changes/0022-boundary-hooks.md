---
module: checks
phase: audit
run: ac8e6a308cf709bbfebc0f36
---

## Motivation

Compose project-wide boundary enforcement before lifecycle verification. A baseline may permit an existing violation, but it cannot establish that a cancelled or broken checker completed. Keep boundary selection and output attribution language agnostic and use the already verified hook adapter.

## Interface

Add Boundaries(ctx context.Context, snapshot project.Snapshot) []Diagnostic to the retained checks facade, delegating all behavior to a new disposable internal package. Use the existing Diagnostic wire model and public RunHook/process contracts.

## Added

### CHK-009
Boundaries MUST select resolved nonblank boundaries commands from every module without a boundaries_opt_out reason, deduplicate identical command strings and execute them once each in lexical command order. Optional commands on modules without disposable globs MUST still run. A non-opted-out module with disposable globs and no nonblank boundary command MUST yield C07 at its module directory, with severity error and corrective guidance; any such missing required command MUST prevent all hook startup. Empty projects, no-disposable modules without commands and explicitly opted-out modules MUST NOT require a command. Each selected command MUST run through RunHook with nil module and empty report, preserving its canonical root, empty module/report substitutions and environment identity, EOF, parent context and configured timeout behavior. A stale project default MUST NOT replace resolved overrides.

### CHK-010
Boundaries MUST trim surrounding whitespace from nonempty combined-output lines and attribute them to exact case-sensitive module names before a colon, or the first whitespace-delimited word when there is no colon. It MUST emit C15 severity error at the matching module directory for every unignored attributed violation, even after exit 0. It MUST suppress attributed lines for reasoned opted-out modules and exact trimmed lines in the optional .fireweed/boundaries-baseline, but MUST NOT suppress other violations by prefix, substring, case folding or a different message from the same module. Unattributed output after exit 0 MUST be ignored as informational. The baseline MUST be read from the canonical repository, tolerate absence, and reject unreadable, directory or escaping-symlink baseline paths with a corrective C15 at that baseline path before hook startup.

### CHK-011
An ordinary nonzero boundary process exit MUST still fail with corrective C15 containing its actual exit and output when it has no attributed lines or has any nonempty unattributed lines, even if those unattributed lines appear in the baseline. Empty nonzero output MUST fail. A completed nonzero hook whose every output line is attributed and suppressed by exact baseline or opt-out MAY pass. Hook launch/configuration/signal/context failures MUST always yield corrective C15 retaining actual exit and captured output, even if output contains a known baselined or opted-out violation; such incomplete execution MUST NOT become successful boundary evidence. Nil or pre-cancelled parent contexts MUST fail without hook startup. The adapter MUST terminate timed-out/cancelled process trees so no delayed hook writes survive.

### CHK-012
Boundaries MUST return a non-null diagnostic slice, sort diagnostics deterministically by Path, Code, Severity and Message, and deduplicate identical diagnostics from overlapping checkers without hiding different violations. Attributed messages MUST retain the complete trimmed violation and corrective guidance. Repeated and overlapping calls MUST have independent results, hook state and cancellation, preserve all snapshot/module/config/document inputs, parent environment and current directory, and leave the baseline unchanged. It MUST NOT independently run eval/setup hooks, agents, the bootstrap or network commands, create Git commits/refs/claims, modify repository or ignored/user files, or leave report artifacts; explicitly authored boundary-hook actions remain the caller's responsibility. Success, ordinary violation and fatal execution paths MUST preserve these guarantees.

## Modified

## Removed

## Decisions

This API consumes a loaded snapshot and resolved module hooks. It does not run static graph validation, doctor canaries, evals, or an entire lifecycle. Missing required module hooks fail before any checker; modules without disposable code do not require a hook, but their configured optional command is useful and runs. Explicit opt-outs remove a module's command from selection, while its lines in another global checker are suppressed.

Baseline matching is exact after surrounding whitespace normalization. The hook contract uses module-prefixed lines; successful unprefixed informational text is tolerated, as in the reference. Nonzero unprefixed text cannot be laundered through a baseline. Ordinary nonzero exits are acceptable only when actual attributed lines explain all output and all violations are intentionally suppressed. Fatal process errors always fail independently of line suppression. Read failures of this protected policy file fail closed; only a missing baseline means no suppressions. Resolve its path with the repository path-safety contract.

D015 records the source-observed reference gaps: Boundaries in bootstrap internal/fireweed/coverage.go ignores hook Err and negative exit when at least one attributed line is suppressed, and silently ignores baseline read errors. This change intentionally requires completed evidence and readable local policy, consistent with the design's working boundary hook and exact baseline requirements. Diagnostic sort/dedup follows the existing checks convention.

## Rejected alternatives

Running the project default once would ignore resolved overrides. Treating a killed checker as successful because its first line was baselined would accept incomplete verification. Ignoring a broken baseline would silently change the protected policy. Boundary-specific command execution would duplicate established timeout and quoting behavior.

## Provenance

- evals | 2026-10-22T18:10:12Z | agent: manual | fireweed: 0.4.0 | playbook: 5867363302d4c874074520977b1627d26991d95ee2fbacb9ff63e827661edd0b | tests: 105 | failing: 42 | passing: 63
- audit | 2026-10-22T18:15:55Z | agent: manual | fireweed: 0.4.0 | playbook: a9c982d24b476621316049b527e6c7bbc319f06b12268af150fa7d8c6df92230 | tests: 0 | failing: 0 | passing: 0
- evals | 2026-10-22T18:17:50Z | agent: manual | fireweed: 0.4.0 | playbook: 5867363302d4c874074520977b1627d26991d95ee2fbacb9ff63e827661edd0b | tests: 108 | failing: 44 | passing: 64

## Audit

- CHK-010: "or the first whitespace-delimited word when there is no colon" — missing assertion for a line with a colon whose full prefix is not a module name but whose first word is one, such as `alpha informational: detail`; an implementation that always attributes by first word would pass the current fixtures.
- CHK-010: "The baseline MUST be read from the canonical repository" — missing suppression assertion with a baseline at the repository root and snapshot.Root set to a nested directory; the nested-root selection fixture has no baseline.
- CHK-010: "C15 severity error at the matching module directory for every unignored attributed violation, even after exit 0" — missing assertion for an ordinary nonzero checker with an unignored attributed violation; every attributed line in the current ordinary nonzero fixtures is suppressed, so that path's module attribution is unasserted.
- CHK-011: "containing its actual exit and output" — missing assertion retaining the complete combined output in the mixed-attribution fixture; its predicate accepts only `compiler failed` and can lose the preceding baselined line.
- CHK-012: "Repeated and overlapping calls MUST have independent results" — missing an assertion exposing shared backing storage between empty diagnostic results; appending to the first empty result and serializing the still-empty second result cannot detect that alias.
- CHK-012: "preserve all snapshot/module/config/document inputs" — missing preservation assertion with an unsorted module-name order; snapshots covered by the existing preservation assertions are already sorted by module name, so in-place sorting would pass.
- CHK-012: "parent environment and current directory, and leave the baseline unchanged" — missing parent-state and baseline preservation assertions around the overlapping calls; that fixture compares only snapshot values, while these guarantees are checked only on sequential calls.
- CHK-012: "overlapping calls MUST have independent results, hook state and cancellation" — the overlap fixture can reject normal completion: the brief checker sleeps only 0.4 seconds, but long-checker cancellation is allowed 1.5 seconds; if cancellation takes longer than 0.4 seconds, the already-successful brief result is incorrectly classified as cancellation. Assert its actual outcome or establish a controlled still-running premise before making that timing assertion.
