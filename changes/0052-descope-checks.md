---
already_satisfied:
    - CHK-004
    - CHK-006
    - CHK-008
    - CHK-011
    - CHK-012
bases:
    CHK-004: d235dbfa
    CHK-006: b11a6779
    CHK-008: 0cbdbfe0
    CHK-011: 4b0e9e51
    CHK-012: ff56655a
module: checks
---

## Motivation

The 2026-10-26 project review found effort going into hardening beyond the design while core commands stayed unbuilt. The revised design (Spec shape; SYSTEM.md is a plain document) and the SYSTEM.md scope guardrails say that locality and safety properties are conventions checked by review, that specs state observable behavior without aliasing or purity requirements, and that the supported environment is macOS and Linux with ordinary full SHA-1 clones. See D048.

## Interface

No public API change.

## Added

## Modified

### CHK-004
ParseJUnit and Analyze MUST return non-null Tests, BySpec and Diagnostics even when empty, set Module to the target name, preserve testcase evidence and hookExit, and sort diagnostics deterministically by Path, Code, Severity and Message. Duplicate failFirst/already entries MUST NOT duplicate policy diagnostics.

### CHK-006
RunHook MUST supply EOF stdin, return the complete combined output and actual exit, and stop the hook at the earlier of the parent context and snapshot.Config.Timeout; cancellation MUST kill the hook's process group and keep output already emitted. Nil context, blank commands, invalid/nonpositive configured timeouts and repository discovery failures MUST fail with Exit 2 and a useful Err before hook startup; a pre-cancelled valid context MUST prevent startup and return Exit -1 with the context error.

### CHK-008
For a valid report and ordinary completed hook, Evaluate MUST feed actual parsed cases, actual hook exit, the target module, every snapshot module prefix and the supplied failFirst/already lists into Analyze, retaining its exact testcase evidence, ID mappings, policy diagnostics and deterministic order. It MUST set Output and HookExit to the real hook result, preserve the selected module name and honor resolved per-module hooks rather than the project default. It MUST NOT run setup or boundaries hooks or leave report files in the repository.

### CHK-011
An ordinary nonzero boundary process exit MUST still fail with corrective C15 containing its actual exit and output when it has no attributed lines or has any nonempty unattributed lines, even if those unattributed lines appear in the baseline. Empty nonzero output MUST fail. A completed nonzero hook whose every output line is attributed and suppressed by exact baseline or opt-out MAY pass. Hook launch/configuration/signal/context failures MUST always yield corrective C15 retaining actual exit and captured output, even if output contains a known baselined or opted-out violation; such incomplete execution MUST NOT become successful boundary evidence. Nil or pre-cancelled parent contexts MUST fail without hook startup.

### CHK-012
Boundaries MUST return a non-null diagnostic slice sorted deterministically by Path, Code, Severity and Message, and deduplicate identical diagnostics without hiding different violations. Attributed messages MUST retain the complete trimmed violation and corrective guidance. It MUST NOT run eval or setup hooks or agents, or modify the baseline.

## Removed

## Decisions

Narrowed specs are listed in already_satisfied because the current implementation already meets the narrower text; nothing here adds behavior. Clauses removed from a spec become unspecified, not forbidden: existing code that still does more is conforming, and the implementation step may simplify it.

Cancellation reaches the hook's process group only. A hook that daemonizes or starts a new session can outlive cancellation; that is a documented limitation (SYSTEM.md, Process control).

## Rejected alternatives

Keeping the clauses and only deleting their evals: coverage requires every enforced spec to stay covered, and the next audit would reinstate the tests. Rewriting every over-long spec to the new shape limits at once: most long specs carry no hardening and rewriting them adds risk without benefit; they are left for later changes.

## Notes

Evals step: remove subtests that assert only removed clauses: input/result aliasing and spare backing storage, repeated-call independence, parent environment and cwd preservation, overlapping-call isolation, and grandchildren or delayed writes surviving cancellation outside the process group. Keep a plain test that a cancelled hook's direct child in the same group is killed. Every Modified ID must keep at least one referencing test. Implementation step: no product change is required.
