---
already_satisfied:
    - FLOW-081
bases:
    FLOW-081: f47668f4
module: workflow
phase: applied
run: 20261107T214650-d187de5ae2adf49791730c83
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
## Provenance
- evals 2026-11-07T21:52:57.972278Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=6a9c2d506363e3539cc66efd37bc6a06701249541dc4e4f01ed3aa6454bbb864 tests=168 failing=0 passing=168
- audit 2026-11-07T21:59:08.180613Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0
- impl 2026-11-07T22:15:18.168964Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=5f856d6bcab308b9ff1a8a7c03b9d2547744112aba81dd30dea8f61ec8062b03 tests=650 failing=0 passing=650
## Audit

No findings; checked FLOW-081
