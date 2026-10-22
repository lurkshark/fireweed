---
module: project
phase: audit
run: 9bd006b5aa4148ac8257841b
---

## Motivation

A separate context or lifecycle process must restore approved config/module/queue metadata from a claim's local Git baseline even after a worker corrupts or deletes current files. Add a read-only immutable source for the existing Snapshot contract, without creating a checkout or reading historical implementation blobs. This supplies the approved metadata that BuildContext requires; live text and phase enforcement remain separate workflow responsibilities.

## Interface

Add LoadAt(start, revision string) (Snapshot, error) to api.go as a thin conversion/delegate to disposable inspection.LoadAt. Keep Snapshot and all existing signatures unchanged. A fatal error returns the zero public Snapshot. Implement behavior under project internal/; add only a new stub during evals, without changing existing implementation.

## Added

### PRJ-024
LoadAt MUST discover the canonical calling checkout from start, including nested directories and linked worktrees, and resolve one explicit nonblank local revision to a full commit before reading its tree. It MUST accept full commit IDs, local branches, annotated/lightweight tags and commit expressions supported by installed Git, in SHA-1 and SHA-256 repositories. An absent/invalid/option-like revision, blob/tree object, unborn HEAD, unavailable repository or Git resolution/enumeration failure MUST return a corrective C01 error and zero Snapshot. All subsequent tree/blob reads MUST be pinned to that resolved commit or its object IDs even if the input ref moves. No current file, index, HEAD or newer commit may replace selected metadata.

### PRJ-025
LoadAt MUST parse the selected commit's fireweed.yaml with public ParseConfig and the embedded binary version, validate configured glob syntax, resolve BaseRef from that config and locally known refs, and return canonical Root. It MUST discover module SPEC.md paths from the selected tree with the same configured-glob, ignored .git/.fireweed subtree, metadata/spec grammar, resolved-default/explicit-empty, retained document/Decisions and name/directory sorting contract as Load. It MUST parse immediate Markdown queue/archive files from that tree, using target prefixes, the existing live/archive handle rules and numeric/slug order; preserve all author/lifecycle/provenance sections; and diagnose duplicate slugs across both lists. It MUST continue past malformed selected documents with the same C02/C03/C04/C05/C09 severity, path and corrective categories as Load, while invalid/missing/unreadable configuration or invalid configured globs is fatal C01. Empty Modules, Changes, Archive and Diagnostics MUST be non-null on success, with diagnostics sorted by path/code/severity/message. Graph/zone/hook/static and queue checks remain the existing public validators' responsibility.

### PRJ-026
Metadata file links MUST resolve only within the selected Git tree. Relative symlink targets and absolute targets lexically within the canonical calling root MUST resolve to selected-tree regular blobs, preserving the target bytes and link-relative path semantics without consulting current worktree link targets. Chained contained links and executable regular blobs MUST work; escaping targets, dangling links, loops, NUL/backslash-invalid paths, directories or gitlinks in place of selected metadata MUST be rejected with the corresponding configuration/module/change category, never read outside/current content or silently substitute another source. Module discovery MUST NOT traverse directory symlinks, gitlinks or nested .git/.fireweed subtrees. Out-of-glob SPEC.md contents and unrelated blobs MUST NOT be read. Explicit queue/archive directory links that resolve within the selected tree MUST expose immediate Markdown entries at their logical changes/ or changes/archive/ paths. Directory links in selected metadata targets MUST use the same contained virtual-tree rules; no filesystem traversal is permitted.

### PRJ-027
LoadAt MUST read only necessary local Git metadata and selected config/module/immediate queue/archive blobs and their committed link targets; it MUST NOT read current working-file contents or unrelated historical code/private/playbook/recipe blobs, including when an eventual error would be ignored. Dirty, deleted, untracked, staged, ignored, symlinked or unreadable current metadata MUST NOT affect the selected snapshot. Successful snapshots and all nested data MUST be independent of inputs and existing repeated results, including empty append backing storage; failures MUST return no partial Snapshot. Every call MUST remain local/noninteractive, preserve files/index/HEAD/branches/claim/other refs/common Git/user/outside data and parent cwd/environment/temp storage, and MUST NOT create a temporary checkout/worktree/export, run native Git or fireweed hooks, agents, bootstrap or network commands, or commit/stage/reset/fetch/push/create branches. Literal Git arguments and hook suppression MUST keep configuration strings and unusual filenames inert. Existing Load behavior MUST remain unchanged.

## Decisions

The caller supplies the approved commit (normally Claim.Base). Resolve it once; metadata body reads use immutable blob IDs, not mutable ref:path expressions. Snapshot has no new provenance field: the caller already retains the full claim baseline. BaseRef is still the selected config's local ref name, resolved through the established defaults; it is not the selected source commit itself. No network refresh occurs.

This is selected-tree discovery, not a scratch checkout. Symlinks are interpreted as committed bytes in a virtual slash-separated tree with lexical containment; absolute links are rooted only in the canonical calling checkout, without consulting current filesystem aliases. File links can target contained committed metadata elsewhere, but discovery does not invent module trees through directory aliases. Gitlinks do not import foreign repository contents. Physical permissions on current metadata cannot restrict an available committed blob. Existing malformed-document diagnostic categories and archive-human-prose treatment remain shared with Load.

This API does not trust or reload a worker's current config/SPEC, does not validate claim ownership, replace a selected change with Claim.State, collect edits, construct a worker prompt or advance a phase. Workflow will compose those responsibilities. Static Check/CheckQueue remain explicit, allowing callers to inspect a diagnosed baseline while repairing source state.

## Rejected alternatives

A temporary checkout/export introduces filesystem mutation, setup hooks and cleanup hazards for a read operation. Reading the current config or symlink targets defeats approval restoration. Loading every historical blob exposes disposable implementation to a later regeneration context. Returning partial successful metadata after fatal source errors can grant incomplete or incorrect permissions. Re-resolving a branch per file mixes incompatible approved states.

## Provenance

- evals | 2026-10-22T23:42:38Z | agent: manual | fireweed: 0.4.0 | playbook: 5867363302d4c874074520977b1627d26991d95ee2fbacb9ff63e827661edd0b | tests: 368 | failing: 134 | passing: 234
- audit | 2026-10-22T23:47:09Z | agent: manual | fireweed: 0.4.0 | playbook: a9c982d24b476621316049b527e6c7bbc319f06b12268af150fa7d8c6df92230 | tests: 0 | failing: 0 | passing: 0
- evals | 2026-10-22T23:53:50Z | agent: manual | fireweed: 0.4.0 | playbook: 5867363302d4c874074520977b1627d26991d95ee2fbacb9ff63e827661edd0b | tests: 381 | failing: 147 | passing: 234

## Audit

- PRJ-024: "resolve one explicit nonblank local revision to a full commit before reading its tree" — invalid/unknown-revision cases assert the final fatal result but do not reject preliminary HEAD tree/body reads whose results are discarded. The unknown-revision native case still whitelists ordinary HEAD metadata; add an ordering guard that rejects tree/body access before successful revision resolution.
- PRJ-025: "resolve BaseRef from that config and locally known refs" — every successful selected configuration uses base main with no distinguishing remote/ref setup. No LoadAt assertion distinguishes resolution from the selected config/local refs from a hard-coded main or current-HEAD-derived BaseRef.
- PRJ-027: "it MUST NOT read current working-file contents" — the native content guard checks absolute path prefixes only, and every read control supplies an absolute path. Relative open/openat requests such as fireweed.yaml from the probe's repository cwd bypass the guard, allowing ignored current-content reads; add relative/dirfd controls and effective-path enforcement.
- PRJ-027: "LoadAt MUST read only necessary local Git metadata and selected config/module/immediate queue/archive blobs and their committed link targets" — the Git body guard accepts any revision prefix before an allowed :path, and accepts arbitrary ^{tree}/^{commit} queries. It therefore permits unselected HEAD/newer-ref metadata reads whose results are ignored; restrict accepted revision-qualified reads to the selected resolved commit and prove an unselected-revision body is rejected while the selected body remains readable.
