---
already_satisfied:
    - FLOW-009
    - FLOW-015
bases:
    FLOW-009: 08a1783d
    FLOW-015: c68f7655
module: workflow
---



## Motivation

The design requires stable diagnostic causes. D006 assigns existing lifecycle
preconditions P01, claims P02, approved state P03, stale bases P05, blocking P06
and audit submissions P09. Current lifecycle errors are mostly uncoded, while
checks now supply their static/eval/zone/boundary causes. Migrate existing errors
as one cohesive workflow change without adding lifecycle checks or features.

## Interface

Public signatures and types stay unchanged. Returned errors carry the assigned
code followed by a colon, retaining the original cause text and wrapping.

## Added

### new:start-precondition-codes
Start MUST prefix dirty-worktree and queue-eligibility rejections with P01.

### new:unsupported-phase-codes
Verify and Done MUST prefix unsupported-phase rejections with P01.

### new:claim-codes
Start and Current MUST prefix existing claim ownership or contention failures with P02.

### new:approved-state-codes
Verify and Done MUST prefix unapproved author, runtime or audit-state alteration errors with P03 before running hooks.

### new:stale-base-codes
Start, Verify and Done MUST prefix stale-base rejections with P05, retaining the underlying cause.

### new:audit-submission-codes
Audit Done MUST prefix malformed or incomplete No findings submission errors with P09.

### new:blocked-completion-codes
Done MUST prefix its existing blocked-change error with P06 while preserving the committed blocking state and claim release.

### new:change-lookup-codes
Current and Context MUST prefix unknown requested change lookup errors with U02.

### new:dependency-codes
Lifecycle operations MUST preserve assigned static, eval, zone and boundary causes when returning their existing failures, except for lifecycle stale-base classification P05.

## Modified

### FLOW-009
Start MUST report P04 as a warning when neither the local base nor the shared base contains the change file.

### FLOW-015
Start MUST reject with P07 when the shared-base unapproved archive count reaches authoring.max_unapproved, whose default is 5.

## Removed



## Decisions



## Notes

Reuse existing lifecycle failure fixtures for exact code assertions, retaining
all behavior, message, status, pre-hook, HEAD, phase, Questions and claim checks.
Add only a representative public-input case for a named existing cause lacking
one. Do not expand error matrices or add helper infrastructure/instrumentation.
FLOW-009 and FLOW-015 are legitimately already satisfied; other changed IDs
must genuinely fail before implementation. Keep all current coverage and
successful paths. Use no bootstrap source or old implementation as an oracle.

Migrate only existing error branches. Missing/foreign/contended or unreadable
claims use P02; existing approved-state mismatches use P03; existing audit
submission errors use P09; the existing blocking return uses P06. Keep P04 a
warning and P07 an error. Preserve check codes C01–C13, C10/C11/E01/P2-01, Z01
and X01/X02 when those causes are returned; retain warnings X03/X04. A stale
base rejected during lifecycle checks uses P05; a merely missing base retains
C08. Boundary errors use their returned hard diagnostic code, or X02 for an
existing boundary failure without one. Preserve original cause text, avoid
duplicate prefixes, and use P01 only for otherwise uncoded existing operational
preconditions. Unknown requested Current/Context change names use U02; Start's
existing queue-head eligibility rejection remains P01.

Do not add checks for unsupported phases where none exist, new blocked-state
checks, approval features, claims/recovery features or speculative errors.
Do not change audit filtering/round bounds, lifecycle steps, process control,
public interfaces, other modules or enforcement. No file-organization refactor
or general error framework. Controller owns spec application, commits and archive.
