---
bases:
    APP-004: e6774676
    APP-012: f47e7f91
module: app
phase: impl
run: 20261113T021756-6c666d67bfecadc626e28d3a
---

## Motivation

The owner approved resolving APP-004/APP-012 before app regeneration. A held approved snapshot must remain accessible even when live configuration or a queued document is malformed. Specify how coverage combines chosen spec text with approved execution metadata, including a projection for a different change, using workflow.Approved.

## Interface

No app API change. Use the new public workflow.Approved query.

## Added

### APP-157
coverage --with-change selecting the held change MUST accept its canonical slug, queued basename or repository-relative path and work despite malformed or missing live configuration, module specs or change documents. Default coverage and coverage selecting another change MUST report the existing cause when their required live spec text or projection cannot be read, without running an eval hook.


## Modified

### APP-004
modules MUST report the validated graph from the held approved snapshot while this worktree owns a change, even when live configuration, module specs or queued documents are malformed or missing. Without an owned claim it MUST report the validated current graph.

### APP-012
coverage MUST evaluate current module spec text by default, the held approved projected spec text when --with-change selects the held change, or the requested live projection's spec text for another queued change. While held, execution MUST use the approved configuration, module graph, zones and hooks, including for another selected projection. It MUST set FIREWEED_BIN to the actual executable.


## Removed

## Decisions

## Notes

Discover the root and query workflow.Approved before parsing live project metadata. For default coverage read the selected module's current spec text at its approved directory without requiring the whole live queue or configuration to parse. For another change, obtain the live projected text through project.Project, then use approved execution metadata; do not execute that other projection's hooks/set values. A requested module absent from the approved graph is an unknown-module U02, not a reason to adopt unapproved execution metadata. Keep normal no-claim behavior, output shapes and diagnostics. Preserve rejection of malformed live selected spec text. Test representative malformed-live modules/held coverage and another change whose hook differs, using local marker hooks. Do not change lifecycle validation or reconstruct claims inside app.
Owner authorized these five regeneration prerequisites and their necessary module-local changes. Stop after this queue; no regeneration or drill is authorized. Follow SYSTEM.md and the controller context. Workers must not read docs/rebuild/PROGRESS.md or ACCEPTANCE.md. Preserve unrelated assertions and retained interfaces after evals. Use existing public-interface fixtures and focused assertions; do not redesign helpers or add a test matrix. Evals and implementation workers run fireweed context for this change before each step and fireweed verify for verification. Restricted auditors follow only the supplied audit playbook and inputs. Only the controller advances phases and commits.

Evals verification passed with fail-first cases for APP-004, APP-012 and APP-157. The current CLI rejects held queries with malformed or missing live inputs and runs the other projection's unapproved hook. The implementation is writable in `packages/app/internal/entry/lifecycle.go` and `packages/app/internal/entry/project.go`.
## Provenance
- evals 2026-11-13T02:29:17.937238Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6.1-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=6a9c2d506363e3539cc66efd37bc6a06701249541dc4e4f01ed3aa6454bbb864 tests=333 failing=2 passing=331
- audit 2026-11-13T02:29:48.855258Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-astra -c model_reasoning_effort="medium" -s workspace-write {prompt} fireweed=0.1.0 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0
## Audit

No findings; checked APP-004, APP-012, APP-157
