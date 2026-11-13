---
bases:
    APP-004: e6774676
    APP-012: f47e7f91
module: app
---

## Motivation

The owner approved resolving APP-004/APP-012 before app regeneration. A held approved snapshot must remain accessible even when live configuration or a queued document is malformed. Specify how coverage combines chosen spec text with approved execution metadata, including a projection for a different change, using workflow.Approved.

## Interface

No app API change. Use the new public workflow.Approved query.

## Added

### new:held-coverage-inputs
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
