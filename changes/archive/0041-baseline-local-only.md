---
fixes:
    - PRJ-027
module: project
phase: applied
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

Implementation resolves a single commit with Git's no-walk missing-object policy before tree access, retaining the resolved object across mutable-ref movement. Local object-name inventory gates one-level tree traversal and selected blob reads; unavailable unrelated blobs remain uninspected. The selected root tree is read from the known commit header: a direct ^{tree} peel was proven to attempt lazy retrieval when that root tree is missing. Revision sets remain rejected rather than selecting an arbitrary member. Current Load behavior and public interfaces are unchanged.

The committed impl run was suspended twice without changing frozen evals or audit because dependent observers rejected conforming metadata reads. App0001 and workflow0002 corrected those observers through their own governed EVALS/audit/implementation cycles. Resumed this original run and restored exactly its three hash-verified saved files. Final full verification passes project581 records (203.384s), app476 (565.300s), checks108 (39.523s), workflow1914 (742.472s), runner0 and boundaries, with0 failures/errors/skips/diagnostics. Frozen project evals/API still match c99774b and both disposable source hashes match the saved candidate. The host wall clock advanced substantially during the run; every suite completed normally with the full expected inventory, and raw evidence retains the observed timestamps. Earlier focused80.588s, source-identical project race202.908s and fresh vet/whitespace results remain separately scoped. Final done repeats required enforcement before application.

## Rejected alternatives

Trusting GIT_NO_LAZY_FETCH alone fails on the installed supported Git. Denying a transport protocol still attempts fetch and changes error semantics. Rejecting every partial repository or reading all object bodies discards available committed snapshots and violates selective source access. Do not alter global or repository configuration to suppress the bug.

## Audit

No findings. Checked PRJ-027.

## Provenance

- evals | 2026-10-23T12:39:05Z | agent: manual | fireweed: 0.4.0 | playbook: 5867363302d4c874074520977b1627d26991d95ee2fbacb9ff63e827661edd0b | tests: 581 | failing: 66 | passing: 515
- audit | 2026-10-23T12:45:04Z | agent: manual | fireweed: 0.4.0 | playbook: a9c982d24b476621316049b527e6c7bbc319f06b12268af150fa7d8c6df92230 | tests: 0 | failing: 0 | passing: 0
- impl | 2026-10-23T21:58:47Z | agent: manual | fireweed: 0.4.0 | playbook: 653912b194316c62181f62d0a69dc213f7668481ea8b1b9926ef889dd48a9e08 | tests: 3079 | failing: 0 | passing: 3079
