---
module: project
---

## Motivation

Start must assign permanent IDs after author approval, preserving removed-ID reservations and every local reference. Build the deterministic allocation transform separately from claims and file mutation so the lifecycle can use it without copying grammar or namespace logic.

## Interface

Add Allocate(snapshot Snapshot, slug string) (contracts.Change, map[string]string, error). It selects one queued change, returns its canonical independently owned document/typed values and a handle-to-real-ID mapping. It does not write state, run or phase. The later lifecycle owns those actions. Retain all existing APIs.

## Added

### new:allocation
Allocate MUST select exactly one queued change by slug and resolve exactly one existing target module. For an unstarted change it MUST assign each local Added handle in declaration order, using the target prefix and consecutive numbers above the greatest numeric ID reserved by current target-module specs, every target-module archive Added/Modified/Removed item, and real Added IDs in other target-module queued changes with both run and phase. ID formatting MUST use at least three decimal digits and support existing larger decimal widths without overflow. Other modules, author prose, bases/reference-only mentions and other unstarted handles MUST NOT reserve numbers. Author-chosen real Added IDs, partial run/phase state and applied queued state MUST error. A valid already started change with real Added IDs MUST retain those IDs/run/phase and return an empty mapping; a started change still containing handles MUST error rather than allocate on resume.

### new:rewrite
Allocate MUST rewrite every occurrence of each declared handle in canonical change metadata, preamble and every body section, including Added/Modified text, Interface, Decisions, Notes and fenced examples. Overlapping handle names MUST NOT corrupt one another. already_satisfied and nested set hook references MUST become the same assigned IDs as their Added declarations. Returned typed fields and Document MUST agree; unrelated author text, section order, tags, real IDs, lifecycle state and metadata MUST retain their semantic values under the public Markdown normalization contract. Typed change fields MUST override stale retained metadata/edit sections. Returned Path, Slug and Number MUST retain the selected typed identity even when filename numbering differs from Number.

### new:allocation-purity
Allocate MUST return errors for missing/ambiguous selectors or targets, invalid selected typed change grammar and undeclared/cross-only handle references, without returning a partial change or allocation mapping. It MUST validate only the selected change's grammar, without replaying or requiring validity of excluded queue entries. Successful empty mappings MUST be fresh non-null maps. Successful results and mappings MUST have no mutable aliases to the supplied snapshot, modules, changes, archive, configuration or nested documents/maps/slices; neither success nor error MAY mutate inputs. Repeated calls with the same input MUST be deterministic. Allocate MUST NOT read/write files, execute hooks/agents/Git or use networking.

## Modified

## Removed

## Decisions

The highest reserved number plus one implements next-free allocation without reusing historical gaps. Use arbitrary-size decimal arithmetic rather than wrapping or truncating large IDs accepted by the public grammar. Queued reservations require assigned state; invalid unstarted real-ID additions do not burn the namespace. Prefix uniqueness and overall queue/application checks remain the existing validators' responsibility. Allocation validates the selected change with canonical typed fields and returns a public parser-normalized representation while preserving its typed identity. Run/phase assignment, ownership, atomic file writes and eligibility belong to workflow. Resume keeps assigned values and returns no new bindings. An empty slug is not an all-queue selector here.

## Rejected alternatives

Predicting IDs during authoring would recreate parallel author collisions. Replacing shorter handles first with naive string replacement would corrupt longer handles. Replaying unrelated malformed predecessors would block allocation for independent eligible changes.
