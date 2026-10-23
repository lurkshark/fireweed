---
fixes:
    - PRJ-027
module: project
phase: impl
run: 24c228986ed5d6567a331bcf
---

## Motivation

The unchanged design requires local Git behavior without fetches or network commands. A marker-only local promisor fixture proves LoadAt attempts a lazy fetch for an unavailable selected configuration blob on installed Git2.30.1 even with GIT_NO_LAZY_FETCH=1. The error is corrective but the transport attempt violates PRJ-027. Protect all selected metadata resolution and reads, including missing commit/tag/ancestor/tree objects, without weakening exact committed-source or selective-content rules.

## Interface

No public interface change. Add permanent black-box regressions invoking public LoadAt and independent controls proving the local transport trap is live.

## Added

## Modified

## Removed

## Decisions

Exercise unavailable selected commit, annotated-tag resolution, ancestor expression, root and nested tree, configuration/module/queued/archive/link metadata blobs. Assert useful C01 or existing path-attributed document diagnostics as appropriate, no partial fatal result, no transport attempt and preservation of unrelated refs/index/current files/configuration. Unavailable unrelated historical/private/disposable blobs must not prevent an otherwise valid selected snapshot. Available selected metadata in a partially populated repository must retain ordinary revision, linked-worktree and supported object-format semantics. Positive controls use a locally configured executable marker helper, never a real remote endpoint. No production command may delegate to the bootstrap.

## Rejected alternatives

Trusting GIT_NO_LAZY_FETCH alone fails on the installed supported Git. Denying a transport protocol still attempts fetch and changes error semantics. Rejecting every partial repository or reading all object bodies discards available committed snapshots and violates selective source access. Do not alter global or repository configuration to suppress the bug.

## Audit

No findings. Checked PRJ-027.

## Provenance

- evals | 2026-10-23T12:39:05Z | agent: manual | fireweed: 0.4.0 | playbook: 5867363302d4c874074520977b1627d26991d95ee2fbacb9ff63e827661edd0b | tests: 581 | failing: 66 | passing: 515
- audit | 2026-10-23T12:45:04Z | agent: manual | fireweed: 0.4.0 | playbook: a9c982d24b476621316049b527e6c7bbc319f06b12268af150fa7d8c6df92230 | tests: 0 | failing: 0 | passing: 0
