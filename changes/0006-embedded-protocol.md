---
module: assets
---

## Motivation

The new executable needs its own embedded version, protocol playbooks and diagnostic explanations. It must remain usable without the bootstrap checkout or any runtime downloads.

## Interface

Expose BuildInfo {Version string; CurrentFormat, PreviousFormat int}, TextAsset {Name, Text, Hash string}, Info() BuildInfo, Playbooks() []string, Playbook(name string) (TextAsset,error), Codes() []string and Explain(code string) (string,error). Content is embedded in disposable assets implementation, with stable interfaces and permanent public-API evals. Recipe assets will follow in a separate change.

## Added

### new:build-info
Build metadata MUST report version 0.4.0, current format 1 and previous format 0 from a compiled binary, independent of working directory or repository files.

### new:playbooks
The embedded playbook inventory MUST be exactly author-change, audit, evals, impl and regenerate in sorted order. Each lookup MUST return its name, nonempty text and the full lowercase SHA-256 of that exact text; unknown names MUST return errors. Text MUST give concrete commands and rules for its step: authoring a single-module change with local handles, complete specs, stamping and human base approval; fail-first evals with genuine testcase failures and allowed zones; independent exact-clause audits and No findings plus checked IDs; implementation with immutable evals/interfaces and passing dependents/boundaries; regeneration from retained contracts without an old-implementation oracle, preserving interfaces/evals and reporting targeted gaps. Returned inventories MUST be independent slices. Lookups MUST perform no runtime file or network reads.

### new:diagnostics
The diagnostic catalog MUST contain U01, C01 through C16, P01 through P06, P1-01 through P1-03, P2-01, D01 through D07 and H01 through H04, sorted and without duplicates. Each explanation MUST identify the problem and a corrective action. H01 MUST cover protected worker files, H02 immutable impl evals/interfaces, H03 append-only changes by anyone including merge resolutions, and H04 immutable archives with deletion for revert allowed plus duplicate recovery by aborting the merge, listing commits, reverting the incoming duplicate and merging again. Unknown codes MUST return errors. Returned code inventories MUST be independent slices.

## Modified

## Removed

## Decisions

Version and format are data, not mutable global process settings. New binaries may override the version at build time for releases, while these default evals verify an ordinary source build. Playbook hashes are over exact UTF-8 bytes including final newline. Diagnostic wording is fresh; codes and designed meanings retain compatibility with the behavioral reference. Recipe copying, project-owned overrides, stale based_on detection and AGENTS refresh are composed later by project/checks/app.

## Rejected alternatives

Runtime reads from the bootstrap or a source directory would break the single-binary and independent-regeneration requirements. Hashing normalized prose would hide meaningful playbook edits.
