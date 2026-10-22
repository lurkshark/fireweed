---
module: assets
phase: audit
run: 4294b7757eb823403f72c0b8
---

## Motivation

The new executable needs its own embedded version, protocol playbooks and diagnostic explanations. It must remain usable without the bootstrap checkout or any runtime downloads.

## Interface

Expose BuildInfo {Version string; CurrentFormat, PreviousFormat int}, TextAsset {Name, Text, Hash string}, Info() BuildInfo, Playbooks() []string, Playbook(name string) (TextAsset,error), Codes() []string and Explain(code string) (string,error). Content is embedded in disposable assets implementation, with stable interfaces and permanent public-API evals. Recipe assets will follow in a separate change.

## Added

### AST-001
Build metadata MUST report version 0.4.0, current format 1 and previous format 0 from a compiled binary, independent of working directory or repository files.

### AST-002
The embedded playbook inventory MUST be exactly author-change, audit, evals, impl and regenerate in sorted order. Each lookup MUST return its name, nonempty text and the full lowercase SHA-256 of that exact text; unknown names MUST return errors. Text MUST give concrete commands and rules for its step: authoring a single-module change with local handles, complete specs, stamping and human base approval; fail-first evals with genuine testcase failures and allowed zones; independent exact-clause audits and No findings plus checked IDs; implementation with immutable evals/interfaces and passing dependents/boundaries; regeneration from retained contracts without an old-implementation oracle, preserving interfaces/evals and reporting targeted gaps. Returned inventories MUST be independent slices. Lookups MUST perform no runtime file or network reads.

### AST-003
The diagnostic catalog MUST contain U01, C01 through C16, P01 through P06, P1-01 through P1-03, P2-01, D01 through D07 and H01 through H04, sorted and without duplicates. Each explanation MUST identify the problem and a corrective action. H01 MUST cover protected worker files, H02 immutable impl evals/interfaces, H03 append-only changes by anyone including merge resolutions, and H04 immutable archives with deletion for revert allowed plus duplicate recovery by aborting the merge, listing commits, reverting the incoming duplicate and merging again. Unknown codes MUST return errors. Returned code inventories MUST be independent slices.

## Modified

## Removed

## Decisions

Version and format are data, not mutable global process settings. New binaries may override the version at build time for releases, while these default evals verify an ordinary source build. Playbook hashes are over exact UTF-8 bytes including final newline. Diagnostic wording is fresh; codes and designed meanings retain compatibility with the behavioral reference. Recipe copying, project-owned overrides, stale based_on detection and AGENTS refresh are composed later by project/checks/app.

Isolation evidence uses a separately compiled public-API probe. On macOS the native policy denies all workspace source reads, an independent private-file control and network access; it permits OS loader/runtime files. Guard controls verify actual denial before asset lookups. Paths are canonicalized to avoid /var aliases. Other operating systems retain portable content checks but skip this macOS-only isolation probe; no runtime isolation claim is made for them. Installing the profile requires running the test outside the outer agent sandbox, not weakening the project's hooks.

## Rejected alternatives

Runtime reads from the bootstrap or a source directory would break the single-binary and independent-regeneration requirements. Hashing normalized prose would hide meaningful playbook edits.

## Provenance

- evals | 2026-10-22T07:11:21Z | agent: manual | fireweed: 0.4.0 | playbook: 5867363302d4c874074520977b1627d26991d95ee2fbacb9ff63e827661edd0b | tests: 3 | failing: 3 | passing: 0
- audit | 2026-10-22T07:14:45Z | agent: manual | fireweed: 0.4.0 | playbook: a9c982d24b476621316049b527e6c7bbc319f06b12268af150fa7d8c6df92230 | tests: 0 | failing: 0 | passing: 0
- evals | 2026-10-22T07:23:20Z | agent: manual | fireweed: 0.4.0 | playbook: 5867363302d4c874074520977b1627d26991d95ee2fbacb9ff63e827661edd0b | tests: 6 | failing: 6 | passing: 0

## Audit

- AST-001: "independent of working directory or repository files" — CWD isolation does not establish independence from absolute repository-file reads.
- AST-002: "Text MUST give concrete commands and rules for its step" — detached keywords do not establish the specified instructions, complete specs, human approval, allowed zones, genuine failed testcase evidence or passing dependent/boundary checks.
- AST-002: "Lookups MUST perform no runtime file or network reads." — no assertion detects runtime file/network reads.
- AST-003: "Each explanation MUST identify the problem and a corrective action." — a word count does not assert problems or corrective actions.
- AST-003: "H01 MUST cover protected worker files, H02 immutable impl evals/interfaces, H03 append-only changes by anyone including merge resolutions, and H04 immutable archives with deletion for revert allowed plus duplicate recovery by aborting the merge, listing commits, reverting the incoming duplicate and merging again." — keywords do not establish rules, immutability or the required incoming-duplicate recovery sequence.
