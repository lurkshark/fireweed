---
module: workflow
phase: impl
run: 634e20449818a1f3eee56425
---

## Motivation

Lifecycle verification must reject workers rewriting the approval they are being checked against. Establish one independently testable state guard before composing filesystem zones, claims and verify/done. The claim will retain approved text; workers may supply their permitted Decisions, Questions and audit findings without changing requirements or lifecycle state.

## Interface

Add CheckState(module contracts.Module, approvedText, currentText, path string) []checks.Diagnostic to the public workflow facade. It consumes supplied text and module identity, uses public Markdown/change parsing, and delegates all behavior to a new disposable internal state package. No module model files or existing dependency APIs change.

## Added

### FLOW-001
CheckState MUST parse approved and current text with the public change grammar, using the supplied module's ID prefix and the supplied change path. An invalid approved document, an approved target different from the supplied module name, a missing run, or an approved phase outside evals/audit/impl MUST return a P02 error at that path with corrective claim guidance. Once approval is valid, invalid current text MUST return P2-01 at that path with corrective worker-edit guidance. Invalid state MUST NOT be accepted as an unchanged change. A valid unchanged started change, including maintenance and regenerate changes, MUST return an empty non-null diagnostic slice.

### FLOW-002
CheckState MUST reject semantic changes to every approved metadata field except blocked with P2-01, including module, nested set/hook/zone values, bases, fixes, already_satisfied, regenerate, run and phase, and adding or removing metadata keys. Valid blocked reasons MAY be added, replaced or removed. For an approved evals or impl phase it MUST permit adding, changing, moving or removing Decisions, Questions and Notes sections; for an approved audit phase it MUST permit only Audit and Questions sections. Every other parsed section, its body and relative order, and the parsed preamble MUST remain unchanged. Protected edits MUST yield P2-01, including Motivation, Interface, Added/Modified/Removed, Removed decisions, Rejected alternatives, Provenance, unknown sections and changes hidden after duplicate or fenced headings. Permissions MUST derive from the approved phase even when current metadata attempts to switch phase. Metadata ordering/quoting and line-ending changes with the same public parsed values MUST NOT be rejected. Exact case-sensitive section names and public fenced-Markdown parsing MUST determine permissions rather than substring replacement.

### FLOW-003
CheckState MUST return fresh non-null diagnostics sorted deterministically by Path, Code, Severity and Message, with severity error, the supplied path and nonempty corrective messages for every rejection. Metadata and protected-body violations in the same valid current document MUST both remain visible; invalid approved or current grammar MUST stop that comparison rather than inventing partial acceptance. Results, including writable empty results, MUST remain independent across repeated calls, and supplied module fields/maps/slices and input strings MUST remain unchanged. It MUST work from supplied text outside any repository, remain noninteractive and MUST NOT create or edit files/refs, execute Git, hooks, agents, the bootstrap or network commands on successful or rejected calls.

## Modified

## Removed

## Decisions

This is a text-state guard, not full lifecycle verification: filesystem zones, symlinks, HEAD/claim ownership, questions/block transitions, fail-first, audit completeness and hook execution remain separate composition. Input module identity is the resolved target from the approved snapshot. Public ParseChange owns grammar and normalized document values; the guard owns permission comparison. It compares parsed metadata keys and nested values, so adding an explicit empty immutable key is still a metadata edit. blocked alone is mutable metadata. Approved started state must have a run and a recognized worker phase; applied/queued state is not a live step baseline.

The approved phase selects section permissions, so an auditor cannot grant itself implementation permissions by changing phase. Non-audit workers may edit Notes as well as Decisions/Questions, matching the reference's documented free-prose workflow. All occurrences of an allowed section are mutable only when public grammar accepts them; every retained protected section occurrence remains compared. Unknown sections and the preamble remain protected. Parser-normalized whitespace and YAML syntax are not independent approval fields. A metadata violation and a protected-body violation produce separate diagnostics, in stable order; invalid text produces one parse/baseline diagnostic first.

## Rejected alternatives

Comparing only typed change metadata loses unknown protected body edits and exact explicit-key presence. Trusting current phase permits self-authorized escalation. A line-based section stripper mistakes fenced examples for actual headings. Reparsing the repository during this text comparison would make its claim baseline dependent on worker-controlled files.

## Provenance

- evals | 2026-10-22T18:41:22Z | agent: manual | fireweed: 0.4.0 | playbook: 5867363302d4c874074520977b1627d26991d95ee2fbacb9ff63e827661edd0b | tests: 160 | failing: 159 | passing: 1
- audit | 2026-10-22T18:44:22Z | agent: manual | fireweed: 0.4.0 | playbook: a9c982d24b476621316049b527e6c7bbc319f06b12268af150fa7d8c6df92230 | tests: 0 | failing: 0 | passing: 0
- evals | 2026-10-22T18:46:58Z | agent: manual | fireweed: 0.4.0 | playbook: 5867363302d4c874074520977b1627d26991d95ee2fbacb9ff63e827661edd0b | tests: 219 | failing: 218 | passing: 1
- audit | 2026-10-22T18:49:25Z | agent: manual | fireweed: 0.4.0 | playbook: a9c982d24b476621316049b527e6c7bbc319f06b12268af150fa7d8c6df92230 | tests: 0 | failing: 0 | passing: 0

## Audit

No findings. Checked FLOW-001, FLOW-002, FLOW-003.
