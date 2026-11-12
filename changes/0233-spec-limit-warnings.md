---
module: project
phase: audit
run: 20261112T032808-90472063d366a5043a1c6d0d
---

## Motivation

The design's Spec shape, SYSTEM.md and Decisions sections have `check` warn
when a spec item has more than `spec_limits.max_clauses` normative keywords
or `max_words` words (C17), when the same normalized sentence appears in
three or more spec items (C18), and when a module's Decisions exceed
`spec_limits.max_decision_words` (C19). ACCEPTANCE R25 lists C17 and C18 as
required. No spec, code or configuration handling exists for any of them,
so oversized items and copied conventions reach evals and audits unflagged.
This repository's current specs produce no C17 or C19 warning, and two C18
warnings in app.

## Interface

`fireweed.yaml` gains an optional `spec_limits` mapping with `max_clauses`,
`max_words` and `max_decision_words`. Check returns the new warnings.

## Added

### PRJ-071
Load MUST read spec_limits.max_clauses, max_words and max_decision_words as positive integers, defaulting to 3, 80 and 300 when unset, and MUST reject other values with C01.

### PRJ-072
Check MUST return a C17 warning naming each current or pending Added or Modified spec item with more than max_clauses normative keywords or more than max_words whitespace-separated words.

### PRJ-073
Check MUST return a C18 warning naming the items when the same sentence, compared ignoring case, punctuation and whitespace, appears in three or more spec items of the projection after the last pending change.

### PRJ-074
Check MUST return a C19 warning naming the module when a module's Decisions section has more than max_decision_words words.

## Modified

## Removed

## Decisions

## Notes

Normative keywords are the uppercase words the design names: MUST, MUST
NOT (counted once) and SHOULD. The C17 item avoids spelling them out
because an item that lists them would count them.

The two current C18 hits are app's "It MUST leave stderr empty." (APP-093,
APP-096, APP-099, APP-108, APP-113) and "Failures MUST emit null data with
code/message diagnostics." (APP-108, APP-113, APP-117). They are warnings;
leave them for the owner rather than rewording app specs here.

Dependent evals may build fixture specs that repeat a sentence across items
(for example "Alpha MUST change.") and assert an empty diagnostics list. If
a dependent eval fails only because of a new C17, C18 or C19 warning, block
and list those evals for an eval_fix rather than special-casing fixtures.

Run after 0232, which adds the explanations.

`fireweed verify spec-limit-warnings` passed the evals phase. The new evals
failed against the current implementation because Load leaves omitted limits
unset and accepts invalid values, and Check emits none of C17, C18 or C19
for fixtures that exceed the corresponding limits.

The implementation can use `packages/project/internal/fixture/fixture.go`
for configuration and `packages/project/internal/queue/check.go` for warnings,
using the existing projection in `packages/project/internal/queue/project.go`.
## Provenance
- evals 2026-11-12T03:32:35.370871Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6.1-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=6a9c2d506363e3539cc66efd37bc6a06701249541dc4e4f01ed3aa6454bbb864 tests=156 failing=20 passing=136
