---
already_satisfied:
    - FLOW-081
bases:
    FLOW-081: f47668f4
module: workflow
---

## Motivation

The owner approved the scoped residual-failure proposals. The retained failed
workflow drill's FLOW-083 testcase aborted while reading recorded implementation
that Start had deleted, before checking Context. A bounded public-API probe
confirmed correct regenerate playbook selection despite that deletion. FLOW-081
names unstarted regeneration as the deletion trigger, but does not explicitly
state the recorded-phase preservation counterpart. Make that boundary explicit
and give the two behaviors distinct public assertions. The unchanged current
candidate passed the original recorded-phase test, so this is a legitimate
already_satisfied clarification, not a manufactured implementation defect.

## Interface

No public signature/model/interface, runtime source, Decision, hook, configuration
or other-module change. Evals may edit only workflow public evals and this change.
Implementation preserves frozen evals/interfaces and leaves runtime source intact;
only normal controller application of this specification clarification is needed.

## Added

## Modified

### FLOW-081
Before returning from successful Start of an unstarted regenerate change,
Fireweed MUST delete only the target module's files classified as disposable
by the approved projected zone map, retaining all other files. When claiming a
regenerate change with an already recorded phase, Start MUST retain its existing
disposable files.

## Removed

## Decisions

## Notes

Preserve every existing assertion and ID obligation. Keep the existing
unstarted-regeneration deletion/retained-zone assertions. Isolate recorded-phase
file preservation under FLOW-081: a clean approved regenerate change already has
phase impl, a stable recorded run and representative existing disposable content;
ordinary Start must retain that content and the recorded phase/run. Keep a
separate focused FLOW-083 testcase that actually reaches Context and asserts the
exact regenerate playbook is present and ordinary impl playbook absent for the
recorded impl phase. Reuse existing temporary-repository/public-API fixtures and
helpers; avoid a new matrix, duplicate infrastructure or implementation tests.
Moving the existing preservation assertion must not lose it or weaken the Context
assertions. All workflow evals must pass and enforced IDs remain covered.

FLOW-081 already_satisfied is supported by the unchanged candidate's passing
recorded-phase baseline. Do not introduce a fake fail-first, modify runtime code,
use a shadow implementation or change already_satisfied to evade verification.
No source from saved reconstructions is implementation input. Do not inspect
bootstrap/rebuild-v1 or old reconstruction implementation; ordinary current
retained public contracts/evals suffice for this scoped change.

Workers use the current pinned controller on PATH and GOCACHE from SYSTEM. Run
pinned context before each held step, run pinned verify --quick, poll to terminal
and return. Never commit, invoke lifecycle mutators, edit run/phase/Provenance,
change counters or edit protected infrastructure/planning. The outside pinned
controller owns independent restricted audit, full phase verification, all phase
commits, apply/archive and claim release. The auditor edits Audit only, checks the
exact normalized FLOW-081 clauses and preserves the configured rounds; a passing
submission is No findings with Checked FLOW-081.

The controller pin remains unchanged throughout both authorized normal changes.
Preserve the original candidate object, prior reports/proof/attempt refs, existing
archive counters, configured bounds, extended 2026-11-10 22:51:05 UTC deadline,
recovery copies and unrelated worktrees. After this clarification and the separate
runner eval_fix are archived and all runs/claims exit, the owner has authorized
clean corrected-subject validation, a verified controller handoff, then exactly
one workflow drill with that new controller. No worker launches a drill, imports
reconstructed code, promotes a controller or pushes.
