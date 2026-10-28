---
audit_rounds: 0
bases: {}
eval_fix:
    - CHK-050
    - CHK-051
module: checks
phase: audit
run: 20261028T203322-c8cf4243d0d0d9b57c2d399a
---



## Motivation

SYSTEM supports ordinary macOS and Linux environments. During 0071, the frozen
scratch-cleanup helper's git rev-parse --path-format=absolute option failed on
this host's Apple Git 2.30.1 even though the same assertions passed with Git 2.53.
The owner explicitly approved a normal eval_fix after 0071 to remove this
helper-only portability defect while preserving the scratch-location and
physical-cleanup requirements and assertions.

## Interface



## Added



## Modified



## Removed



## Decisions



## Rejected alternatives



## Notes

Edit only assertScratchRemoved in packages/checks/evals/doctor_test.go. Obtain
the common Git directory with git rev-parse --git-common-dir; if relative,
resolve it against the fixture root before the existing EvalSymlinks call.
Preserve canonicalization and every assertion of scratch placement, physical
directory removal and Git registration cleanup. Retain all existing cases,
fixtures and ID links, including the real default-timeout test. Do not add tests,
helpers, platform branches, instrumentation, shims or runtime toolchain choices.
Product implementation, interfaces, specs and other modules stay unchanged.
All 71 own cases must pass using ordinary PATH/Apple Git, with no newer-Git
prefix. A focused cleanup run may precede required full controller phase checks;
do not run another whole-suite quick check merely to duplicate those checks.
The controller alone advances/commits. Implementation phase needs no source edit.

## Provenance

- evals 2026-10-28T20:39:26.095182Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=b26da634732c7f1c7dc600115582a63cb8f82e1fd2bd1d522d608fe837a04ba7 tests=260 failing=0 passing=260
