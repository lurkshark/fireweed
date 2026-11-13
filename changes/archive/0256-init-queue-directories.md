---
module: app
phase: applied
run: 20261113T224513-c7aa4d0531459aad633ca7ab
---

## Motivation

The design's Repository layout and Distribution and init sections, tracked
by R01, give a new project a pending queue and archive. On main 3b3b64d,
successful init --recipe none --agent codex --no-app creates neither changes/
nor changes/archive/. Complete the scaffold without overwriting user files.

## Interface

No public signature or type changes.

## Added

### APP-161
Successful init MUST create changes/ and changes/archive/ as directories, including with --no-app and every supported recipe selection. Existing queue and archive directories and their contents MUST be retained.

### APP-162
If changes or changes/archive is occupied by a regular file, init MUST reject with P01 and exit 2 before writing setup files. The diagnostic MUST identify the conflicting path.

## Modified



## Removed

## Decisions

## Notes

Extend the existing binary init fixtures with an empty new project, existing
queue/archive contents, and a regular file at each required directory path.
Reuse the existing recipe-selection tests rather than building a matrix of
every recipe and conflict. Preserve APP-057/058 preflight behavior, --no-app,
JSON failure shape, and the absence of init-time hooks or agent execution.
No placeholder files or extra tracked queue metadata are required.

Drafted at the owner's request after the R01–R29 traceability review.
No lifecycle execution, agent invocation, drill or release is authorized by
this drafting request. Workers must obtain context before their step and use
fireweed verify; only the controller may advance or commit lifecycle work.

The binary init evals failed against the current implementation for both IDs.
For APP-161, fresh projects lacked both directories for none, Go, TypeScript,
Python and the local recipe. For APP-162, a regular file at either required path
still allowed init to exit 0 and write setup files. The fixture with existing
pending and archived contents passed.

Both behaviors belong in the disposable `packages/app/internal/entry/setup.go`,
where init assembles the directory list for preflight and setup writes. No stub
is needed.
## Audit

No findings; checked APP-161, APP-162

## Provenance

- evals 2026-11-13T22:54:59.566983Z agent=codex fireweed=0.1.0 playbook=6a9c2d506363e3539cc66efd37bc6a06701249541dc4e4f01ed3aa6454bbb864 tests=375 failing=9 passing=366
- audit 2026-11-13T22:55:20.001254Z agent=codex fireweed=0.1.0 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0
- impl 2026-11-13T23:03:10.195127Z agent=codex fireweed=0.1.0 playbook=5f856d6bcab308b9ff1a8a7c03b9d2547744112aba81dd30dea8f61ec8062b03 tests=375 failing=0 passing=375

