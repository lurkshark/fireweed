---
audit_pending:
    - FLOW-087
audit_rounds: 1
bases: {}
module: workflow
phase: evals
run: 20261029T211014-c172b18434a8e46b0f036ac7
---



## Motivation

Design Lifecycle CLI requires commits <change> to list a change's commits on
this branch; abandon uses that list to revert unapplied work. The design
identifies worker commits by change-file phase transitions, never messages.
NEW still rejects commits as unimplemented. Supply only this dependency, then
immediately wire its app caller and prove actual NEW lists the checked phase
commits of a disposable Go lifecycle, excluding unrelated and author commits.

## Interface

Add thin public Commits(root, change string) ([]string, error). Behavior belongs
under workflow/internal. Return full object IDs, newest first.

## Added

### FLOW-087
Commits MUST return the full object IDs, newest first, of commits reachable
from HEAD that change the selected change document's phase relative to their
first parent, following its pending and archive paths. Selection MUST use file
contents rather than commit messages.

### FLOW-088
Commits MUST resolve pending or archived changes by repository-relative path,
filename or unambiguous slug. It MUST reject missing or ambiguous selectors
with U02.

## Modified



## Removed



## Decisions



## Notes

Use public APIs and temporary Git fixtures with explicit identities, distinct
author/unrelated/phase commits and one archive move. Exercise representative
pending/archive selectors, same-message unrelated commits, ordinary phase
transitions and no-phase-change exclusion; no history matrix or private
serialization. Existing gitstore.Run/ReadAt provide Git access. No hooks/agents,
new process primitives, public models, unrelated workflow reorganization, other
module/protected edits, commits or phase commands by workers. Preserve existing
assertions. Evals/interface freeze before impl; separate restricted audit and
native full checks. Own under5m/full under15m, no raised bounds. Second applied
change after correction checkpoint; return immediately to app commits caller.

## Provenance

- evals 2026-10-29T21:25:43.919794Z agent=manual fireweed=0.1.0 playbook=b26da634732c7f1c7dc600115582a63cb8f82e1fd2bd1d522d608fe837a04ba7 tests=384 failing=3 passing=381
- audit 2026-10-29T21:50:14.76299Z agent=manual fireweed=0.1.0 playbook=0fb044cfaca1b9a844a25b4c86223177e52f2f35fb149eddf044e83822665c88 tests=0 failing=0 passing=0

## Audit

- FLOW-087: "Commits MUST return the full object IDs, newest first, of commits reachable" — missing assertion that commits outside HEAD are excluded.
