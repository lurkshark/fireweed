---
module: assets
---

## Motivation

The design's Evals step says recipes show a stub in each language (R07/R19).
AST-044 covers Go, but the TypeScript/pnpm and Python/uv guides contain no
replaceable stub example. Workers need to leave behavior in disposable files
so implementation does not require changing a frozen retained interface.

## Interface

No public signature or type changes.

## Added

### new:additional-language-stubs
The TypeScript/pnpm and Python/uv recipe instructions MUST each show a retained interface delegating to a temporary stub in a new disposable implementation file. Each example MUST explain that implementation replaces the stub without editing the retained interface.

## Modified



## Removed

## Decisions

## Notes

Use each existing recipe's actual zone layout and language imports. Reuse
the existing recipe fixtures to check the example paths and behavior where
useful; no new runtime or package dependency is needed. Preserve ready-made
hook/setup guidance, package exports/import boundaries, and all canaries.
Keep the example small and clearly temporary; it must fail a behavior assertion
rather than fail compilation or collection. Do not auto-upgrade project-owned
recipe copies or add new spec IDs for shared conventions.

Drafted at the owner's request after the R01–R29 traceability review.
No lifecycle execution, agent invocation, drill or release is authorized by
this drafting request. Workers must obtain context before their step and use
fireweed verify; only the controller may advance or commit lifecycle work.
