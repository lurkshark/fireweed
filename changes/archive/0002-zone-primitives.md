---
module: contracts
phase: applied
run: b12d36d6d8aae9082a4bbeb2
---

## Motivation

Zone enforcement must classify future paths as well as existing files. Pattern intersection is required to reject eval/disposable overlaps before a worker can exploit them.

## Interface

Expose Zones with Interface, Evals, Disposable, Preserve and AppendOnly []string; Zone string constants Protected, SpecZone, EvalZone, InterfaceZone, AppendOnlyZone, PreserveZone, DisposableZone, OtherZone, OutsideZone; MatchGlob(pattern, path string) (bool, error); GlobsOverlap(a, b string) (bool, error); and ZoneFor(path, moduleDir string, zones Zones, protected []string) (Zone, error). All paths/globs use repository-relative forward slashes. Module ownership checks will be implemented in project.

## Added

### CON-005
Glob matching MUST support literals, * and ? within a path segment, bracket classes/ranges/negation, and ** across segments, including zero directories for **/. It MUST match the complete relative path, preserve case and reject malformed classes, absolute paths, traversal components and backslashes. Empty paths and empty patterns MUST be rejected.

### CON-006
Glob intersection MUST determine whether two valid patterns can match a common path, without requiring that path to exist; it MUST account for wildcards, classes and zero-directory globstars, and return an error for either invalid pattern.

### CON-007
Zone classification MUST apply strict precedence: protected, SPEC.md, evals, interface, append-only, preserved, disposable, other module files. Protected paths MUST include root fireweed.yaml and SYSTEM.md, all .fireweed descendants, every AGENTS.md and CLAUDE.md, and configured protected globs, even outside the module. Files outside the module MUST otherwise be outside; module path containment MUST use segment boundaries. Invalid repository paths or globs MUST fail classification.

## Modified

## Removed

## Decisions

Globs are case-sensitive with bracket negation `!` or `^`. Path components equal to . or .., repeated slashes and trailing slashes are invalid. A globstar matches slash only when it occupies a complete segment; adjacent ordinary stars remain segment wildcards. Zone fields are module-relative, configured protected globs repository-relative. Literal module directory itself is outside (zones classify files below it). All configured patterns are validated even when a stricter zone wins. The zone resolver classifies only; preserved interface and append-only regeneration behavior follows from strict precedence.

## Rejected alternatives

Checking only tracked files would miss future eval/disposable collisions. Prefix string containment would incorrectly claim sibling directories.

## Provenance

- evals | 2026-10-22T06:28:56Z | agent: manual | fireweed: 0.4.0 | playbook: 5867363302d4c874074520977b1627d26991d95ee2fbacb9ff63e827661edd0b | tests: 7 | failing: 3 | passing: 4
- audit | 2026-10-22T06:30:54Z | agent: manual | fireweed: 0.4.0 | playbook: a9c982d24b476621316049b527e6c7bbc319f06b12268af150fa7d8c6df92230 | tests: 0 | failing: 0 | passing: 0
- evals | 2026-10-22T06:31:19Z | agent: manual | fireweed: 0.4.0 | playbook: 5867363302d4c874074520977b1627d26991d95ee2fbacb9ff63e827661edd0b | tests: 8 | failing: 4 | passing: 4
- audit | 2026-10-22T06:32:06Z | agent: manual | fireweed: 0.4.0 | playbook: a9c982d24b476621316049b527e6c7bbc319f06b12268af150fa7d8c6df92230 | tests: 0 | failing: 0 | passing: 0
- impl | 2026-10-22T06:33:28Z | agent: manual | fireweed: 0.4.0 | playbook: 653912b194316c62181f62d0a69dc213f7668481ea8b1b9926ef889dd48a9e08 | tests: 8 | failing: 0 | passing: 8

## Audit

No findings. Checked CON-005, CON-006 and CON-007 in separate auditor context /root/audit_zone_primitives. All four prior findings are resolved by Test_CON007AuditRegressions. Prior findings and the evals retry remain in the controller phase commits.
