---
audit_rounds: 0
bases: {}
module: contracts
phase: impl
run: 20261027T010515-ddcff5bdc56f165ee47ed469
---



## Motivation

The NEW CLI now creates the manual slice's queued change. Its Added spec must use
a new: handle before start assigns an ID (design Change files / late IDs).
The existing item parser accepts assigned IDs only, preventing that next step.

## Interface

Keep the existing ParseItems(section string) ([]Item, error) public API unchanged.
Extend the disposable item-heading grammar; do not add a parallel change parser.

## Added

### CON-006
ParseItems MUST accept new:<slug> item headings and preserve the complete handle in Item.ID, where slug contains lowercase letters, digits and hyphens. It MUST reject empty handles and handles containing other characters. Assigned-ID and handle items MUST retain their source order in mixed input, with the existing tag/body preservation and duplicate-ID checks applying to handles too.

## Modified



## Removed



## Decisions



## Provenance

- evals 2026-10-27T01:06:53.314916Z agent=manual fireweed=0.1.0 playbook=b26da634732c7f1c7dc600115582a63cb8f82e1fd2bd1d522d608fe837a04ba7 tests=118 failing=1 passing=117
- audit 2026-10-27T01:09:49.808283Z agent=manual fireweed=0.1.0 playbook=0fb044cfaca1b9a844a25b4c86223177e52f2f35fb149eddf044e83822665c88 tests=0 failing=0 passing=0

## Audit

No findings. Checked CON-006.
