---
module: assets
phase: applied
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
## Audit

No findings; checked AST-071

## Provenance

- evals 2026-11-17T01:02:48.123432Z agent=codex fireweed=0.2.1 playbook=6a9c2d506363e3539cc66efd37bc6a06701249541dc4e4f01ed3aa6454bbb864 tests=152 failing=1 passing=151
- audit 2026-11-17T01:03:11.457375Z agent=codex fireweed=0.2.1 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0
- impl 2026-11-17T01:32:30.991316Z agent=codex fireweed=0.2.1 playbook=5f856d6bcab308b9ff1a8a7c03b9d2547744112aba81dd30dea8f61ec8062b03 tests=1443 failing=0 passing=1443

