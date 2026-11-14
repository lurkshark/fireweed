---
module: assets
phase: applied
run: 20261114T062613-7dd1a31e8f07d0d658f87cbf
---

## Motivation

The design's Evals step says recipes show a stub in each language (R07/R19).
AST-044 covers Go, but the TypeScript/pnpm and Python/uv guides contain no
replaceable stub example. Workers need to leave behavior in disposable files
so implementation does not require changing a frozen retained interface.

## Interface

No public signature or type changes.

## Added

### AST-065
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

Evals fail first for AST-065 because both current guides lack the example.
A temporary Go build overlay confirmed that the documented stubs can fail a
greeting assertion and that replacing only the stub makes the assertion pass.
The implementation-writable files are
`packages/assets/internal/recipe/files/typescript-pnpm.md` and
`packages/assets/internal/recipe/files/python-uv.md`. Use `greet` in the retained
`src/index.ts` and `__init__.py` snippets, delegating to a `greet` function in
the new disposable `src/internal/greet.ts` and `internal/greet.py` snippets,
respectively. The fixture adds the retained snippets to the baseline files, so
show them as additions to those files.
## Audit

No findings; checked AST-065

## Provenance

- evals 2026-11-14T06:34:14.738117Z agent=codex fireweed=0.1.0 playbook=6a9c2d506363e3539cc66efd37bc6a06701249541dc4e4f01ed3aa6454bbb864 tests=146 failing=5 passing=141
- audit 2026-11-14T06:34:38.937637Z agent=codex fireweed=0.1.0 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0
- impl 2026-11-14T07:00:22.584259Z agent=codex fireweed=0.1.0 playbook=5f856d6bcab308b9ff1a8a7c03b9d2547744112aba81dd30dea8f61ec8062b03 tests=1330 failing=0 passing=1330

