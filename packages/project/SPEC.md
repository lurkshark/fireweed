---
depends_on:
    - contracts
    - gitstore
    - assets
id_prefix: PRJ
interface:
    - api.go
    - model/**
name: project
---

# project

## Specs

### PRJ-001
Load MUST read fireweed.yaml and module SPEC.md files only in configured modules globs, returning module names, prefixes, dependencies, documents and zones.

### PRJ-002
ValidateGraph MUST return nil for a valid graph and an explanatory error identifying the rejection cause for an invalid graph.

### PRJ-003
NewChange MUST create a queued Markdown skeleton for an existing module with Motivation, Interface, Added, Modified, Removed and Decisions sections, containing no assigned spec IDs.

### PRJ-004
Project and Stamp MUST resolve a change argument from its unique slug, queued basename or repository-relative queued path, rejecting missing or ambiguous targets. An unknown target MUST return U02 naming the rejected selector and listing valid pending change slugs.

### PRJ-005
Project MUST preserve complete Added and Modified item text and remove Removed items when applying pending spec edits.

### PRJ-006
Stamp MUST write each Modified or Removed item's canonical SpecHash from the projection preceding the target change into its bases.

### PRJ-007
Check MUST return no error diagnostics for valid module specs and changes. Each failure diagnostic MUST have a nonempty code and message.

### PRJ-008
Check MUST reject a pending change without a nonempty Motivation section with diagnostic C07.

### PRJ-009
Check MUST reject a module spec or Added/Modified spec whose body contains no normative keyword with diagnostic C04.

### PRJ-010
Check MUST reject spec tags other than review or quarantine with a nonempty reason with diagnostic C06.

### PRJ-011
Check MUST reject duplicate spec IDs within a module's Specs section with diagnostic C05.

### PRJ-012
Check MUST reject module spec IDs and change edit IDs whose prefix belongs to another module with diagnostic C05.

### PRJ-013
Check MUST reject a change that uses the same edit identity more than once across Added, Modified and Removed with diagnostic C05.

### PRJ-014
Check MUST reject a pending change that references a handle absent from that change's Added entries with diagnostic C09.

### PRJ-015
Check MUST reject Modified or Removed entries whose bases hash is missing or
differs from the preceding projected spec's canonical hash with diagnostic C08.
An absent, null or empty-string entry MUST report `missing base for <ID>`;
a nonempty differing hash MUST report `stale base for <ID>`.

### PRJ-016
Check MUST reject Modified or Removed IDs absent from the projection preceding the change with diagnostic C08.

### PRJ-017
Load MUST inherit an omitted module zone field from module_defaults while retaining an explicitly empty module zone list as empty.

### PRJ-018
Load MUST return an error for missing or malformed config or module documents.

### PRJ-019
ValidateGraph MUST reject module names that are empty or duplicated in the graph.

### PRJ-020
ValidateGraph MUST reject module prefixes that are empty or duplicated in the graph.

### PRJ-021
ValidateGraph MUST reject a dependency that does not name a different existing module.

### PRJ-022
ValidateGraph MUST reject a graph containing a dependency cycle.

### PRJ-023
NewChange MUST use the next number after the largest pending or archived change number.

### PRJ-024
NewChange MUST reject a slug already present in the pending queue or archive with diagnostic C13, without overwriting existing files.

### PRJ-025
NewChange MUST reject an unknown target module with an error beginning
`unknown module ` and naming the rejected module.

### PRJ-026
NewChange MUST reject an empty slug or a slug containing characters other than lowercase letters, digits and hyphens with an error whose contracts.ErrorCode is U01.

### PRJ-027
Queue MUST return pending NNNN-slug.md changes in numeric-then-slug order, allowing tied numbers and excluding archives.

### PRJ-028
Queue MUST return each pending change's module, frontmatter and parsed Added, Modified and Removed items.

### PRJ-029
Queue MUST reject a malformed pending change document.

### PRJ-030
Queue MUST reject a pending change naming an unknown module.

### PRJ-031
Queue MUST reject pending changes that share a slug.

### PRJ-032
Project MUST apply the queue through the named change inclusively, or the whole queue when through is empty.

### PRJ-033
Project MUST apply set values for depends_on, interface, evals, disposable, preserve, append_only, boundaries_opt_out and hooks to the projected module.

### PRJ-034
Project MUST reject identity or unknown fields in set.

### PRJ-035
Project MUST reject an invalid projected module graph.

### PRJ-036
Project MUST keep Added handles local to their declaring change, allowing different changes to reuse handle text.

### PRJ-037
Stamp MUST reject Modified or Removed IDs absent from the projection preceding the target change.

### PRJ-038
Stamp MUST retain other frontmatter values and human section contents.

### PRJ-039
Stamp MUST succeed on an Added-only change without assigning spec IDs.

### PRJ-040
Check MUST report configuration discovery or decoding failures with diagnostic C01.

### PRJ-041
Check MUST report invalid current or projected module graphs with diagnostic C02.

### PRJ-042
Check MUST report unreadable or malformed module/change documents and malformed item syntax with diagnostic C03, retaining C05 for duplicate identities.

### PRJ-043
Check MUST report invalid module metadata or zone field types, and invalid set keys or value types, with diagnostic C12.

### PRJ-044
Check MUST report pending queue identity errors with diagnostic C13.

### PRJ-045
Check MUST report invalid pending change module targets with diagnostic C07.

### PRJ-046
Check MUST report error C07 for a pending change containing eval_fix together
with fixes or any Added, Modified or Removed spec edits.

### PRJ-047
LoadAt MUST return the selected Git revision's project configuration and modules, using its recorded fireweed.yaml module globs and module SPEC.md documents rather than working-tree contents.

### PRJ-048
LoadAt MUST apply the selected revision's module_defaults to omitted zone fields while preserving explicitly empty module zone lists.

### PRJ-049
LoadAt MUST set Snapshot.Root to the canonical repository root and Module.Dir to each module's repository-relative directory.

### PRJ-050
LoadAt MUST return an error when the revision cannot be resolved, its recorded fireweed.yaml is missing or malformed, or a matched recorded module directory lacks a readable SPEC.md.

### PRJ-051
For a nonempty ID, SpecHistory MUST return the target module's archived changes whose Added or Modified items name that ID, including tag-only modifications.
Pending changes and archives belonging to other modules or unrelated spec IDs MUST not appear in that selection.

### PRJ-052
Every history entry MUST identify its archive path and module, the matching IDs, its Motivation, added and removed Decisions, and its complete Provenance section.

### PRJ-053
SpecHistory MUST order entries by their latest valid implementation Provenance timestamp, newest first, with missing or invalid implementation dates after dated entries.
Equal-date or undated entries MUST use archive-path lexical order as a total-order fallback.

### PRJ-054
With an empty ID filter, SpecHistory MUST return each target-module archive containing Added or Modified spec edits once, with its unique affected IDs in lexical order.
A module with no matching archives MUST return an empty slice.

### PRJ-055
An unknown module or an ID absent from both current module Specs and matching
archive edits MUST return an error whose contracts ErrorCode is U02 and whose
text names the rejected name.
A current ID with no matching archive edits MUST return an empty slice.
An ID in matching archive edits MUST remain queryable when it no longer exists
in current Specs.

### PRJ-056
Missing or malformed fireweed.yaml MUST produce an error whose contracts
ErrorCode is C01.
Malformed archived change documents or malformed Added/Modified item syntax
MUST produce an error whose contracts ErrorCode is C03 and whose text
identifies the archive path.

### PRJ-057
Stamp MUST return an error whose contracts ErrorCode is C08 when the change's bases frontmatter is not a mapping.

### PRJ-058
Stamp MUST return an error whose contracts ErrorCode is U02, with the requested change as subject, when no pending change matches it.

### PRJ-059
Load MUST return an error preserving the contracts ErrorCode C03 or C05 when a module SPEC.md fails to parse.

### PRJ-060
StaleSpecs MUST return the target change's Modified and Removed items whose recorded bases differ from the corresponding preceding projected spec hashes, in edit order, with ID, RecordedHash and Current text.
Items without a bases entry MUST be omitted.

### PRJ-061
StaleSpecs MUST recover Before by projecting the spec immediately before the target change at the Git commit that last changed that change file's bases mapping.
It MUST set Recovered true only when the recovered item's normalized SpecHash equals RecordedHash.

### PRJ-062
When matching stamped history cannot be recovered, StaleSpecs MUST return the stale item's Current text and RecordedHash with Recovered false and empty Before.
Unavailable history MUST NOT prevent reporting other stale items.

### PRJ-063
StaleSpecs MUST leave repository files, index and refs unchanged.
It MUST return an empty slice when no recorded bases are stale.

### PRJ-064
StampMissing MUST add canonical preceding-projection hashes only for Modified or Removed IDs absent from bases, preserving every existing entry including stale hashes.
It MUST retain other frontmatter values and human section contents.

### PRJ-065
StampMissing MUST reject a nonmapping bases value or an edited ID absent from the preceding projection with C08 and preserve the target file on rejection.

### PRJ-066
StaleSpecs and StampMissing MUST accept a unique pending slug, queued basename or repository-relative queued path.
An unknown selector MUST return U02 naming the requested change and listing valid pending change slugs; an ambiguous slug MUST return C13.

### PRJ-067
When NewChange cannot create its queued file, it MUST return an error beginning
with `create <queued path>:` and retaining the filesystem cause, where the queued
path is repository-relative.

### PRJ-068
When Check rejects syntax or duplicate identities in a module SPEC.md, the
failure diagnostic MUST identify that document's repository-relative path in
Diagnostic.Path.

### PRJ-069
When `changes` is a regular file, Queue MUST return an error identifying `changes`. Check MUST report a non-warning C03 diagnostic identifying `changes`.

### PRJ-070
Check MUST return a C20 warning naming each tracked file named SPEC.md that is neither a configured module's SPEC.md nor inside a configured module directory or `.fireweed/`.

### PRJ-071
Load MUST read spec_limits.max_clauses, max_words and max_decision_words as positive integers, defaulting to 3, 80 and 300 when unset, and MUST reject other values with C01.

### PRJ-072
Check MUST return a C17 warning naming each current or pending Added or Modified spec item with more than max_clauses normative keywords or more than max_words whitespace-separated words.

### PRJ-073
Check MUST return a C18 warning naming the items when the same sentence, compared ignoring case, punctuation and whitespace, appears in three or more spec items of the projection after the last pending change.

### PRJ-074
Check MUST return a C19 warning naming the module when a module's Decisions section has more than max_decision_words words.

### PRJ-075
Check MUST reject with non-warning C12 when any module's resolved evals and disposable globs match a common possible module-relative file path, whether or not that file exists. The diagnostic MUST identify the module and conflicting globs.

### PRJ-076
Check MUST apply the eval/disposable overlap rule to the current graph and each preceding-through-current queued projection, including inherited defaults, explicit empty lists and set edits.

### PRJ-077
Check MUST reject a module SPEC.md with zero or more than one level-two Specs section using a non-warning C03 diagnostic naming that file. Heading recognition MUST follow CON-002's fenced-code and ATX-heading rules. A single empty Specs section MUST remain valid.

### PRJ-078
Check MUST reject non-whitespace content anywhere in a pending change's Removed section other than valid assigned-ID headings, using non-warning C07 and identifying the change and offending item when present. An empty Removed section or valid headings separated only by whitespace MUST remain valid when their targets and bases are valid.

### PRJ-079
When fixes, eval_fix or already_satisfied is present, Check MUST require a list of nonempty string identities and reject other values with non-warning C07 naming the change and field.

### PRJ-080
Check MUST reject a fixes or eval_fix identity absent from the target module's projected Specs after that change, with non-warning C07 naming the change, field and identity. IDs belonging to another module MUST be rejected by the same rule.

### PRJ-081
Check MUST reject an already_satisfied identity that does not name an Added or Modified item of that change, using non-warning C07 naming the change and identity. Before Start, a local Added handle MUST be accepted as that item's identity; after assignment, its assigned ID MUST be accepted.

### PRJ-082
CheckCompatibility MUST accept an omitted fireweed pin or a nonempty comma-separated conjunction of comparisons using =, <, <=, > or >= and two- or three-component nonnegative decimal versions. It MUST compare against assets.Version numerically, treating an omitted patch component as zero.

### PRJ-083
CheckCompatibility MUST return a C01 CodedError when a present fireweed pin is malformed, is not a string, or excludes assets.Version. The error MUST identify the pin and running binary version and distinguish invalid syntax from incompatibility.

### PRJ-084
Load MUST apply CheckCompatibility to the live configuration before returning a Snapshot, and Check MUST expose a compatibility failure as non-warning C01 for fireweed.yaml.

### PRJ-085
CheckCompatibility MUST accept format version 1 and an omitted version as the initial format, and MUST reject any other present version value with a C01 CodedError identifying the unsupported value and supported format 1. Format acceptance MUST be independent of binary-range acceptance.

## Decisions

[0126-spec-history-query] History reads the project's current archive files, as the archive-backed show
view; it does not infer chronology from commit ancestry or filenames. Entry
metadata are Markdown section contents without their level-two heading; a
nonempty ID filter returns that matching ID in each entry's IDs field.
