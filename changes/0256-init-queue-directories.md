---
module: app
---

## Motivation

The design's Repository layout and Distribution and init sections, tracked
by R01, give a new project a pending queue and archive. On main 3b3b64d,
successful init --recipe none --agent codex --no-app creates neither changes/
nor changes/archive/. Complete the scaffold without overwriting user files.

## Interface

No public signature or type changes.

## Added

### new:queue-directories
Successful init MUST create changes/ and changes/archive/ as directories, including with --no-app and every supported recipe selection. Existing queue and archive directories and their contents MUST be retained.

### new:queue-directory-conflicts
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
