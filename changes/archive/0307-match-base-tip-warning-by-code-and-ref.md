---
eval_fix:
    - RUN-072
module: runner
phase: applied
run: 20261117T105320-b1f1eb8ec47052681a31dcfb
---

## Motivation

The RUN-072 eval must identify the warning that Next is required to retain.
In the workflow drill report in commit `f6f5997`, Start returned a P04
approval warning and a P05 base-tip warning, both naming
`refs/remotes/origin/main`. The eval selected the first diagnostic by ref
text alone and stopped before exercising Next.

## Interface

## Added

## Modified

## Removed

## Decisions

## Notes

In `Test_RUN_072_RetainsStartWarning` and its shared matcher, identify the
base-tip diagnostic by P05 and the selected ref together. Keep the warning
flag check and the full diagnostic equality and uniqueness assertions across
the existing return paths. Keep the separate P04 retention and worker
recovery warning assertions. With this matcher, the focused eval and all its
subtests passed on both main's implementation and the retained reconstruction
`16c778e`.
## Audit

No findings; checked RUN-072

## Provenance

- evals 2026-11-17T11:03:07.677894Z agent=codex fireweed=0.2.1 playbook=6a9c2d506363e3539cc66efd37bc6a06701249541dc4e4f01ed3aa6454bbb864 tests=217 failing=0 passing=217
- audit 2026-11-17T11:03:26.223488Z agent=codex fireweed=0.2.1 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0
- impl 2026-11-17T11:15:38.363886Z agent=codex fireweed=0.2.1 playbook=5f856d6bcab308b9ff1a8a7c03b9d2547744112aba81dd30dea8f61ec8062b03 tests=620 failing=0 passing=620

