---
module: assets
---

## Motivation

SYSTEM.md calls for public-interface evals, and the design treats dependent
evals as evidence of missing module contracts. The latest drill review found
claim construction tied to private storage, error checks tied to punctuation,
and a snapshot comparison that lost numeric kinds through JSON. The evals
playbook can guide authors toward tests that retain those signals.

## Interface

## Added

### new:module-boundary-eval-guidance
The evals playbook MUST direct workers to exercise their module's public
behavior that consumers depend on and to build ordinary dependency state
through public operations. It MUST direct workers to assert structured
diagnostic codes and subjects when contracted, and to inspect typed values before encoding
when type affects callers.

## Modified

## Removed

## Decisions

## Notes

Keep the guidance short and compatible with SYSTEM.md's public-interface
eval conventions. Use held claims, diagnostic identity and typed snapshot
configuration as examples if needed. Exact layout or wording remains useful
where the public spec requires it. Test the guidance through Playbook("evals").
Keep the existing audit scope and verification rules.
