---
bases: {}
fixes:
    - CHK-053
module: checks
---



## Motivation

EC3's actual NEW doctor --no-agent on the clean Go candidate fails C14:
"module template evals has directory". The project-owned Go templates contain
valid evals/ and internal/ scaffolding; new-module supports that directory tree,
and the design says doctor creates its throwaway modules from those templates.
CHK-053 requires rendering the project-owned module templates and canary
baseline. Doctor currently rejects directories and reads only flat files, so
it cannot validate the real Go hooks needed for the core-cutover gate. Render
the valid template tree without weakening malformed-input or hook checks.

## Interface

No public API or spec changes.

## Added



## Modified



## Removed



## Decisions



## Rejected alternatives

Flattening or deleting valid project-owned template scaffolding to appease
doctor would evade its required template rendering. Disabling doctor in CI or
skipping its failing gate is not permitted.

## Notes

Use the existing public Doctor fixture and bounded local fake hooks. Add one
representative regression with nested module template files, including .tmpl
rendering and the existing substitutions, and require both throwaway modules
to receive the rendered files before probes succeed. Keep current missing or
malformed template/map rejection, setup, four-canary detection/recovery and
scratch cleanup assertions. No new requirements, interfaces, platform probes,
process instrumentation, real agents or drills belong in permanent evals.

The protected Go setup and template tree remain unchanged. Fix only the checks
implementation after genuine fail-first and separate audit. All ordinary
phase zones, dependent verification, approval and retry/round bounds remain
active. Only the controller commits, applies and archives.

The preceding candidate's real NEW lifecycle succeeded at the retained clone
/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-ec5-qgu1qc4l/candidate:
fail-first cafd070, restricted audit a2b9e66 and applied e2c0b25. Its real audit
workspace contained only change/spec/evals with independent Git metadata and
was removed; the clone is clean and unclaimed. Its CI gate remains pending
because of this doctor failure. Return to the affected gates after this fix.
