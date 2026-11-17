---
module: assets
phase: audit
run: 20261117T005842-6fd96ec088ef13e602419066
---

## Motivation

SYSTEM.md calls for public-interface evals, and the design treats dependent
evals as evidence of missing module contracts. The latest drill review found
claim construction tied to private storage, error checks tied to punctuation,
and a snapshot comparison that lost numeric kinds through JSON. The evals
playbook can guide authors toward tests that retain those signals.

## Interface

## Added

### AST-071
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

`Test_AST_071_PublicBehaviorGuidance` fails against the current
`Playbook("evals")` because the consumer, dependency setup, diagnostic identity
and typed-value guidance is missing. The implementation-writable playbook text
is in `packages/assets/internal/content/content.go`.
## Provenance

- evals 2026-11-17T01:02:48.123432Z agent=codex fireweed=0.2.1 playbook=6a9c2d506363e3539cc66efd37bc6a06701249541dc4e4f01ed3aa6454bbb864 tests=152 failing=1 passing=151

