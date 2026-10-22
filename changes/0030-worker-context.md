---
module: workflow
phase: impl
run: 6ec380df90af73c41671e42a
---

## Motivation

Workers need one complete context assembled from the approved projected contract, exact source documents and only the files their phase may inspect. Compose the existing projection, safe path, Git and playbook APIs before adding the executable context command and lifecycle transitions. Preserve instruction bytes and reject missing or unsafe selected inputs without leaking implementation or silently delivering an incomplete context.

## Interface

Add BuildContext(snapshot project.Snapshot, slug string) (WorkerContext, error) to api.go as a thin disposable delegate. Add retained model types FileContent {Path, Content string} and WorkerContext {Change contracts.Change; ChangeText string; Module contracts.Module; SpecText, System string; Dependencies []contracts.Module; Files []FileContent; Playbook, PlaybookHash, GitCommonDir string}, exposed as public aliases, with snake_case JSON fields matching their names. No other public API changes. Put behavior under internal/context/.

## Added

### FLOW-016
BuildContext MUST select exactly one queued change by its exact Slug from the supplied approved snapshot, reject missing/ambiguous selectors, missing target modules, invalid queued paths or phases outside empty/evals/audit/impl with corrective C05, and return that independent Change. The queued path MUST be directly under changes/ and have the selected slug under the public change filename grammar. It MUST obtain the target Module by projecting all preceding queued changes through public Project and applying the selected change through public ApplyChange, including current specs, metadata, zones and Decisions behavior of those APIs. Later changes MUST NOT affect it. Dependencies MUST contain only that projected target's direct dependencies, in declared order, using modules from the preceding projection; missing dependencies MUST fail C05. Empty Dependencies MUST be non-null. Selection/phase/path/projection errors MUST stop before repository/file access. The caller supplies an approved snapshot; this operation does not reload worker-modified config or module metadata.

### FLOW-017
BuildContext MUST discover the canonical repository from snapshot.Root and return its canonical GitCommonDir, supporting nested roots and linked checkouts. It MUST return exact current queued ChangeText, exact current target SPEC.md as SpecText and exact SYSTEM.md as System, preserving frontmatter, CRLF, empty contents and missing final newlines. A missing SYSTEM.md MUST produce empty System; other system/spec/current-change read failures MUST fail C01 with a corrective path-attributed error. A missing current change MAY use only the exact State from a valid stored claim owned by this canonical checkout whose Change, Run, parsed target and phase match the selected started change; otherwise it MUST fail C01. Unsafe paths MUST fail before content opens and MUST NOT fall back through a claim. It MUST select author-change for an empty phase, otherwise the selected phase, with regenerate taking precedence when Regenerate is true; public ResolvePlaybook MUST provide exact Playbook and PlaybookHash, including overrides, empty text and stale/malformed metadata without new warnings. Resolver errors MUST retain their public corrective details.

### FLOW-018
BuildContext MUST enumerate current tracked and untracked nonignored repository files through local Git, deduplicate and sort Files by repository-relative slash Path and preserve exact content bytes. Tracked ignored files MUST remain eligible; untracked ignored files MUST be excluded. In every phase it MUST include only the target's eval and interface files; outside audit it MUST additionally include target preserved/append-only files and direct dependency interface files. Selection MUST use the projected target zones, preceding projected dependency zones and public ZoneFor precedence with configured protected paths. Audit MUST exclude preserved/append-only and dependency interface contents. Disposable, other/manifest, unrelated module, transitive-only dependency, worker_writable-only, protected and SPEC.md files MUST NOT appear in Files or have their contents read by this enumeration. Missing tracked selected files, unreadable selected files, directories and unsafe/dangling selected paths MUST fail C01 rather than being silently skipped. Files MUST be non-null when empty.

### FLOW-019
Every selected content read, including SYSTEM.md, SPEC.md, the queued change, playbook and Files, MUST be contained through public SafePath before opening content; contained regular symlinks MAY succeed, while escapes, dangling links and missing descendants of escaping directories MUST fail without outside reads even if a later error would be ignored. Filesystem/Git enumeration errors MUST fail with corrective C01. Any failed BuildContext MUST return the zero WorkerContext with no partial files or metadata. Successful results and their nested maps/slices MUST be independent of inputs and repeated results, including immediately after mutation and on empty collections. Successes and rejections MUST remain local/noninteractive, preserve files/index/HEAD/branches/claim/other refs/common Git/user/ignored/outside files and parent cwd/environment/temp storage, and MUST NOT run hooks, agents, bootstrap or network commands, commit, stage, reset, fetch, push or create branches. Necessary read-only local Git inspection is allowed. Hook/agent commands from snapshots MUST NOT execute.

## Decisions

This API consumes an approved captured snapshot, like CheckWorktree; later lifecycle/CLI composition owns capture/approval and current-edit rejection. It does not execute hooks or validate a live step for advancement. Change remains the approved typed record while ChangeText exposes current exact bytes for worker feedback. Missing-change recovery uses the owned matching claim only; other read failures are actionable errors. SYSTEM.md absence matches standalone/minimal project context, while target SPEC.md and selected retained files are required. Atomic zero results prevent partial context being mistaken for success.

Projected context includes preceding queued effects and the selected change, but no later effects. Direct dependency metadata/zones come from the preceding projection and are not recursively expanded. Files contain only phase-selected retained text; manifests and worker_writable files are writable permissions rather than automatic prompt content. Git tracked-plus-untracked enumeration includes new stubs/retained contracts but excludes ignored scratch and binaries. All explicit source reads use contained regular files. These safeguards extend the ordinary reference context behavior, which silently ignores several selected read errors; no real regeneration is claimed by constructing a context.

## Rejected alternatives

Reading entire module trees exposes disposable implementations to reconstruction and unrelated private data to audits. Serving partial success after selected read failures hides missing contracts. Applying every queued change makes context depend on later unapproved intent. Reloading all worker-modified metadata can grant prompt access through an edited zone map. Adding executable/lifecycle transitions in this module change would mix independently testable ownership boundaries.

## Provenance

- evals | 2026-10-22T21:59:47Z | agent: manual | fireweed: 0.4.0 | playbook: 5867363302d4c874074520977b1627d26991d95ee2fbacb9ff63e827661edd0b | tests: 1067 | failing: 123 | passing: 944
- audit | 2026-10-22T22:05:09Z | agent: manual | fireweed: 0.4.0 | playbook: a9c982d24b476621316049b527e6c7bbc319f06b12268af150fa7d8c6df92230 | tests: 0 | failing: 0 | passing: 0
- evals | 2026-10-22T22:16:29Z | agent: manual | fireweed: 0.4.0 | playbook: 5867363302d4c874074520977b1627d26991d95ee2fbacb9ff63e827661edd0b | tests: 1078 | failing: 133 | passing: 945
- audit | 2026-10-22T22:17:52Z | agent: manual | fireweed: 0.4.0 | playbook: a9c982d24b476621316049b527e6c7bbc319f06b12268af150fa7d8c6df92230 | tests: 0 | failing: 0 | passing: 0

## Audit

No findings. Checked FLOW-016, FLOW-017, FLOW-018, FLOW-019.
