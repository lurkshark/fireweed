---
already_satisfied:
    - CHK-055
bases:
    CHK-055: d2bd79ec
module: checks
---

## Motivation

The owner approved resolving the failed checks drill's CHK-055 baseline-scope
question. The unchanged design's Language support section describes doctor as
proving hooks with two throwaway modules. Clarify that baseline eval coverage
belongs to those canary modules, rather than requiring every existing module to
have eval coverage. Preserve all hook, canary and recovery obligations.

## Interface

No public signature, implementation, recipe, hook or configuration change.

## Added

## Modified

### CHK-055
Doctor MUST require passing baseline evals and coverage for the throwaway canary modules, and applicable boundaries, before planting canaries and after removing each planted canary. This baseline coverage requirement applies to the canary modules used to prove hooks, rather than every existing module. Recovery failures MUST produce C14 naming the failed hook and canary.

## Removed

## Decisions

## Notes

The existing RealGoRecipe eval supplies a real module without evals and requires
doctor's successful baseline, all four canaries and recovery. EvalEvidenceIsSpecific
rejects missing baseline coverage before canaries are planted. BaselineAndRecoveryFailures
rejects failed baseline hooks and failed recovery evals/coverage; existing required
boundary and override evals retain their assertions. Reuse this evidence with
legitimate already_satisfied; do not add redundant cases or alter current evals.
CHK-058/CHK-029/CON-004 and their assertions remain unchanged. The failed drill's
reconstructed doctor is not imported or repaired on main.

Evals and implementation workers must run pinned fireweed context for this change
before work and verify with fireweed verify --quick, polling until terminal, then
return. Only the outside owning controller advances phases, commits, applies and
archives with full required verification. Never invoke lifecycle mutators
(including start, done, next, run, release or abandon), git commit or git reset;
never edit run/phase fields or Provenance. Implementation makes no product edits.
Workers may record concise existing evidence in this Notes section only.

The separate restricted auditor edits only the existing Audit body. Findings
identify an affected spec ID and exact normalized clause with no asserting test.
Passing audit writes No findings and Checked CHK-055. Re-audit only previously
flagged IDs, preserving counters and all prior clearances.

