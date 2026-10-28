---
bases: {}
fixes:
    - CHK-051
    - CHK-055
    - CHK-056
    - CHK-057
module: checks
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
