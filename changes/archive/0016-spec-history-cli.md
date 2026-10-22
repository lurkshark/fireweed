---
module: app
phase: applied
run: 795dbbe2bbf903abdce7ba7c
---

## Motivation

Expose current/projected specs and archived reasons through the independent executable. Authors need to inspect tags, queued intent and why permanent behavior exists before writing or changing it. History must follow actual recorded implementation chronology rather than accidentally treating queue filename numbers as timestamps.

## Interface

No public Go interface change. Extend the existing executable Run dispatcher with show <module> [ID] [--projected | --history]. Use public project Load, Check, CheckQueue and Project and retained contracts. JSON module/spec data follows the existing public-model convention, including retained documents for complete context.

## Added

### APP-009
show MUST resolve an exact named module from a nested project directory and return its complete current module in JSON when ID is omitted, or the exact case-sensitive spec item when ID is supplied. Plain output MUST render only the selected ordered specs as level-three headings with retained tags and complete bodies, one blank line between items and one final newline; no specs MUST produce empty plain output. It MUST retain current/discovery diagnostics from public Check, globally sorted by path/code/severity/message, with exit 0 for no errors or warnings alone and exit 1 when errors exist while retaining selected data. Unknown modules or exact IDs MUST produce U01 exit 2 with corrective guidance. Fatal project/configuration failures MUST produce C01 exit 2 and null JSON data.

### APP-010
show --projected MUST use public Project over the complete numeric/slug ordered queue and expose resulting specs and resolved metadata without changing current state or replaying archives. It MUST combine current/discovery and CheckQueue diagnostics, preserving change paths and globally sorted order, and return projected data with exit 0/1 by diagnostic severity. A failed projection MUST return exit 1 with C05 error diagnostics and null JSON data rather than a partial projection. Exact projected ID lookup MUST work for real IDs and local handles; multiple matching local handles MUST produce a corrective U01 ambiguity error rather than choosing one. --projected and --history MUST be mutually exclusive.

### APP-011
show --history MUST return a non-null array of archive rows for the selected module, optionally filtered to exact IDs in Added/Modified/Removed, including retags and removed IDs that are no longer current. It MUST exclude pending entries and fix-only references from ID-filtered history; unfiltered module history MUST include all that module's archived changes. Each row MUST contain the complete change plus its retained Motivation, Decisions and Provenance bodies, preserving historical handle examples. Newest history MUST come first: rows with a parseable RFC3339 timestamp in their final parseable impl provenance line MUST sort before undated rows and by that timestamp descending, then numeric Number, Slug and Path descending to break ties; undated rows MUST use the same numeric/slug/path fallback. Missing history, including an unknown historical ID, MUST return an empty array with exit 0 absent other validation errors. Plain history MUST render the same row array as JSON data, and current/discovery diagnostics MUST follow the current show severity and stream contract without projecting pending changes.

### APP-012
All show modes MUST accept --json before or after command/operands and use the schema-1 canonical show/data/non-null diagnostics envelope without plaintext on JSON streams. Plain data MUST go to stdout and coded/path-attributed diagnostics to stderr. Missing/excess operands, unknown or valued flags, mutually exclusive modes and mode flags on other implemented commands MUST be U01 exit 2 before loading a project. -- MUST end option parsing; help MUST work without loading a project. Output write failures MUST return exit 2. Current, projected and historical reads on valid and invalid queues MUST remain noninteractive and MUST NOT run hooks, agents, the bootstrap or network commands or mutate repository files, ignored contents or refs.

## Modified

## Removed

## Decisions

Exact spec lookup is distinct from tolerant JUnit ID matching. Projected duplicate local handles are valid per-change namespaces, so selecting an ambiguous shared spelling requires showing the whole module instead. Current/history reads expose Check diagnostics only; projected reads also expose CheckQueue and never allocate IDs. JSON retains public-model documents/fields consistently with modules; plain spec reads match the reference's concise heading/body format. Historical removed events are included, matching the reference's useful extension to added/modified/retagged history. ID-filtered fix-only changes are excluded, while unfiltered module history includes maintenance and fixes.

History uses the final parseable impl provenance line because the design explicitly requires newest-first history; dated rows precede undated legacy rows, whose ordering falls back to number/slug/path. This defines a total ordering even when provenance is missing. Author decision D012 records a reference chronology mismatch demonstrated in a separate fixture. Preserve complete raw Provenance regardless of parseability. Claims, drift-since-drill status and other workflow reads remain later slices.

## Rejected alternatives

Requiring a spec to be current before showing history would hide why removed behavior disappeared. Sorting only by queue filename numbers would mislabel newer archives when an earlier-numbered authored change applies later. Sorting by timestamps only when both compared rows have dates would not define a transitive order across mixed legacy data.

## Provenance

- evals | 2026-10-22T16:05:32Z | agent: manual | fireweed: 0.4.0 | playbook: 5867363302d4c874074520977b1627d26991d95ee2fbacb9ff63e827661edd0b | tests: 22 | failing: 6 | passing: 16
- audit | 2026-10-22T16:11:33Z | agent: manual | fireweed: 0.4.0 | playbook: a9c982d24b476621316049b527e6c7bbc319f06b12268af150fa7d8c6df92230 | tests: 0 | failing: 0 | passing: 0
- evals | 2026-10-22T16:13:01Z | agent: manual | fireweed: 0.4.0 | playbook: 5867363302d4c874074520977b1627d26991d95ee2fbacb9ff63e827661edd0b | tests: 22 | failing: 6 | passing: 16
- audit | 2026-10-22T16:13:27Z | agent: manual | fireweed: 0.4.0 | playbook: a9c982d24b476621316049b527e6c7bbc319f06b12268af150fa7d8c6df92230 | tests: 0 | failing: 0 | passing: 0
- impl | 2026-10-22T16:15:31Z | agent: manual | fireweed: 0.4.0 | playbook: 653912b194316c62181f62d0a69dc213f7668481ea8b1b9926ef889dd48a9e08 | tests: 22 | failing: 0 | passing: 22

## Audit

No findings. Checked APP-009, APP-010, APP-011 and APP-012. Separate restricted auditor rechecked exact clauses against permanent public/executable assertions, including corrective lookup guidance in both output formats, case-sensitive module lookup and multiline/fenced body preservation. Same-model separate context; no implementation/history/bootstrap access or test execution.
