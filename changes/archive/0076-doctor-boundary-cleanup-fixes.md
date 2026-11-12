---
audit_rounds: 0
bases: {}
fixes:
    - CHK-051
    - CHK-055
    - CHK-056
    - CHK-057
module: checks
phase: applied
run: 20261028T213607-f7bde2247258cdc46a0653c9
---



## Motivation

The 0071–0075 checkpoint found two doctor violations, and the owner approved
this scoped normal correction. A NEW Go fixture with disposable code and no
boundary hook returned exit 0 with no diagnostics, skipping undeclared and
internal_eval canaries. Cleanup also uses the cancelled probe context, so Git
worktree removal never launches after cancellation and registration survives.
Restore the design's required-boundary rule and CHK-051/055–057 without narrowing
requirements, changing interfaces or adding diagnostic allocations.

## Interface



## Added



## Modified



## Removed



## Decisions



## Rejected alternatives



## Notes

Reference docs/rebuild/reports/checkpoint-0071-0075.md only for the approved
finding scope; current specs and public APIs remain authority. Add focused
public checks.Doctor regressions using existing local doctor fixtures. Cover
missing required boundary-hook coverage, including a real module lacking both
a global hook and its own override while another override exists. Absence must
not be mistaken for successful boundary-canary detection. Preserve existing
X02/C14 assignments and the named hook/canary details required by the fixed IDs;
keep reasoned opt-outs and X04 cases unchanged. Apply the requirement to real
modules; do not invent a global-hook requirement when applicable overrides or
opt-outs cover them.

Exercise cancellation after an ordinary setup/probe starts in scratch, using a
bounded fake command and readiness marker. Assert both physical scratch removal
and Git registration cleanup through the public repository state. Cancellation
must stop the probe process group, then cleanup must still run with its own
bounded context rather than the cancelled probe context. No process observers,
interposers, platform branches, new API or testing-only synchronization hook.

Preserve all prior assertions, interfaces, IDs and the real CHK-065 three-minute
default-timeout regression. Use representative cases, not a new failure matrix.
The new regressions must fail at runtime against the existing implementation.
Use focused new-case runs while developing; required controller full checks
already include the long default-timeout test, so avoid a redundant whole-suite
quick pass solely to duplicate them. Keep own Verify under five minutes and full
implementation Verify under fifteen; change no controller timeouts or bounds.

Implement only the scoped doctor behaviors in this module. Do not refactor
unrelated hooks, change other modules, alter archived changes, fix the deferred
contracts gap or run a drill. Implementation freezes evals/interface. The
controller alone advances, applies, archives and commits; workers never commit.
Audit reads only this change, projected specs, evals, retained interfaces and
governing rules, without implementation/history reads; re-audit only prior
finding IDs and retain normal persisted bounds.

Eval phase: four focused new cases ran against the existing implementation.
Three failed at runtime: cancelled setup retained Git registration, and missing
required boundaries were accepted both alone and beside an applicable override.
The override-only control passed. The cancellation case also checked physical
scratch removal and that the bounded setup descendant did not finish.

Impl: doctor checks required boundary coverage on real modules in scratch before
canary probes, and removes the scratch worktree with an independent bounded
context. The four new focused cases and nearby baseline, override, opt-out and
real Go fixture cases pass; `git diff --check` passes. Full phase checks remain
with the controller.

## Audit

No findings. Checked CHK-051, CHK-055, CHK-056, CHK-057.

## Provenance

- evals 2026-10-28T21:42:27.947682Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=b26da634732c7f1c7dc600115582a63cb8f82e1fd2bd1d522d608fe837a04ba7 tests=298 failing=3 passing=295
- audit 2026-10-28T21:47:04.012113Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=0fb044cfaca1b9a844a25b4c86223177e52f2f35fb149eddf044e83822665c88 tests=0 failing=0 passing=0
- impl 2026-10-28T21:54:54.615748Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=07554a3249d8fca0ba70df7504135f297491465caa63458734a4b8351734d0e5 tests=298 failing=0 passing=298
