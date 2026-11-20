---
already_satisfied:
    - PRJ-090
module: project
phase: applied
run: 20261119T225940-c0d588746bc6509e821ea05b
---

## Motivation

Users running `fireweed next` or `fireweed run` need to know when their
configuration is malformed, even when they have no pending changes. In the
`project` reconstruction drill, both commands reported a successful empty
queue because `Queue` skipped configuration validation when `changes/` was
absent.

## Interface

## Added

### PRJ-090
If `fireweed.yaml` contains malformed YAML, `Queue` MUST return an error whose
`contracts.ErrorCode` is `C01`, even when the `changes/` directory is absent.

## Modified

## Removed

## Decisions

## Notes

Add an eval through the public `project.Queue` API. Use an otherwise valid
project fixture with no `changes/` directory, write malformed YAML to
`fireweed.yaml`, and assert that `contracts.ErrorCode` reports `C01`. The
current implementation loads the project before reading the queue.
## Provenance

- evals 2026-11-19T23:00:27.774915Z agent=codex fireweed=0.2.4 playbook=5e05bacdae954379b468564c11429c5532016e26a7afbcd418f9bdb9441e151c tests=327 failing=0 passing=327
- audit 2026-11-19T23:00:40.865364Z agent=codex fireweed=0.2.4 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0
- impl 2026-11-19T23:20:57.164874Z agent=codex fireweed=0.2.4 playbook=5f856d6bcab308b9ff1a8a7c03b9d2547744112aba81dd30dea8f61ec8062b03 tests=1322 failing=0 passing=1322

## Audit

No findings; checked PRJ-090

