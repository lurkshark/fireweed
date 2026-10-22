---
depends_on:
    - contracts
id_prefix: AST
name: assets
---

# assets

## Specs

### AST-001
Build metadata MUST report version 0.4.0, current format 1 and previous format 0 from a compiled binary, independent of working directory or repository files.

### AST-002
The embedded playbook inventory MUST be exactly author-change, audit, evals, impl and regenerate in sorted order. Each lookup MUST return its name, nonempty text and the full lowercase SHA-256 of that exact text; unknown names MUST return errors. Text MUST give concrete commands and rules for its step: authoring a single-module change with local handles, complete specs, stamping and human base approval; fail-first evals with genuine testcase failures and allowed zones; independent exact-clause audits and No findings plus checked IDs; implementation with immutable evals/interfaces and passing dependents/boundaries; regeneration from retained contracts without an old-implementation oracle, preserving interfaces/evals and reporting targeted gaps. Returned inventories MUST be independent slices. Lookups MUST perform no runtime file or network reads.

### AST-003
The diagnostic catalog MUST contain U01, C01 through C16, P01 through P06, P1-01 through P1-03, P2-01, D01 through D07 and H01 through H04, sorted and without duplicates. Each explanation MUST identify the problem and a corrective action. H01 MUST cover protected worker files, H02 immutable impl evals/interfaces, H03 append-only changes by anyone including merge resolutions, and H04 immutable archives with deletion for revert allowed plus duplicate recovery by aborting the merge, listing commits, reverting the incoming duplicate and merging again. Unknown codes MUST return errors. Returned code inventories MUST be independent slices.

## Decisions

[embedded-protocol]
Version and format are data, not mutable global process settings. New binaries may override the version at build time for releases, while these default evals verify an ordinary source build. Playbook hashes are over exact UTF-8 bytes including final newline. Diagnostic wording is fresh; codes and designed meanings retain compatibility with the behavioral reference. Recipe copying, project-owned overrides, stale based_on detection and AGENTS refresh are composed later by project/checks/app.

Isolation evidence uses a separately compiled public-API probe. On macOS the native policy denies all workspace source reads, an independent private-file control and network access; it permits OS loader/runtime files. Guard controls verify actual denial before asset lookups. A test-only dylib interposes Darwin file-read/metadata and socket entrypoints after loader startup, exiting 42 on attempts, including ignored errors; positive controls prove file and network interception before testing assets. This requires local cc only for the macOS isolation eval; the product remains Go-only. Paths are canonicalized to avoid /var aliases. Other operating systems retain portable content checks but skip this macOS-only isolation probe; no runtime isolation claim is made for them. Installing the profile requires running the test outside the outer agent sandbox, not weakening the project's hooks.
