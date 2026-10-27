---
bases: {}
module: contracts
---



## Motivation

The NEW CLI now creates the manual slice's queued change. Its Added spec must use
a new: handle before start assigns an ID (design Change files / late IDs).
The existing item parser accepts assigned IDs only, preventing that next step.

## Interface

Keep the existing ParseItems(section string) ([]Item, error) public API unchanged.
Extend the disposable item-heading grammar; do not add a parallel change parser.

## Added

### new:handle-headings
ParseItems MUST accept new:<slug> item headings and preserve the complete handle in Item.ID, where slug contains lowercase letters, digits and hyphens. It MUST reject empty handles and handles containing other characters. Assigned-ID and handle items MUST retain their source order in mixed input, with the existing tag/body preservation and duplicate-ID checks applying to handles too.

## Modified



## Removed



## Decisions
