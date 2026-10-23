---
fixes:
    - FLOW-011
module: workflow
---

## Motivation

Public marker-only discovery on installed Git2.30.1 reproduces lazy transport despite GIT_NO_LAZY_FETCH=1: LoadClaim/Claims/UpdateClaim fetch an unavailable selected claim payload or stored base, and CreateClaim/UpdateClaim fetch unavailable requested commits, ancestor revisions, annotated tag objects or tag targets. Every rejected call preserves existing files/refs, but the transport attempt itself violates FLOW-011. Exact-token release and independent-slug creation correctly avoid transport and must keep working. This is a genuine specified-behavior regression with fresh failing public-API evals, not already_satisfied.

Full project0041 verification separately exposed older workflow observers that mistake cat-file commit for the mutating commit command and reject OID-only object inventory and pinned root/subtree IDs. Independent public-only review confirms these necessary metadata reads are permitted by FLOW-020/023 and PRJ-024/027. Correct those eval helpers in this owning module's EVALS phase and independently calibrate their premises; do not weaken private contents, mutable source, unrelated claim, actual mutation/network or native isolation restrictions.

## Interface

No public API/type changes. Keep CreateClaim, LoadClaim, Claims, UpdateClaim and ReleaseClaim signatures and all existing claim/CAS/ownership semantics. Add permanent FLOW-011 public regressions for required unavailable objects and valid independent operations across ordinary/linked SHA-1/SHA-256 fixtures. Observe actual fetch attempts independently through Git Trace2 and a local marker-only transport helper, with live allowed/denied-transport controls. Snapshot state after fixture setup, clear controls before API calls, require corrective zero-result P02 and no partial inventory on rejection, and preserve refs/HEAD/index/files/input/cwd/environment/temp. Required available commits with irrelevant missing objects must remain usable; no complete-object-graph or all-claim decoding precondition is introduced.

Correct workflow's generic Git operation observer and approval source observer in evals only. Parse actual operations and global/command option operands, permit necessary object-name inventory, selected commit/no-walk/tree metadata, and prove real-success byte-equal marker-free allowed controls. Retain live denials for actual mutations/transport, private object bodies/batches, mutable source and unrelated claim inventories. Existing public-success native isolation and hook/fsmonitor controls remain mandatory. Implement the actual no-fetch fix only in disposable workflow internals after fail-first and independent audit.

## Added

## Modified

## Removed

## Decisions

Selected object availability must be established without triggering transport, including older Git versions which ignore the environment hint. All operations retain bounded local subprocesses and literal argument handling. Object-name metadata does not authorize reading arbitrary blob bodies or decoding unrelated claims. A claim payload may be unreadable while exact-token ReleaseClaim remains valid; an unrelated bad claim must not block an independent CreateClaim or selected LoadClaim. Claims inventory still fails atomically when a required entry is unavailable. Missing required requested/stored commits are errors; missing unrelated trees/blobs/ancestors must not invent a stronger baseline requirement. No production code may adapt itself to the invalid legacy metadata observers.

This author-ordered0002 fix runs before the unclaimed committed project0041 implementation. Preserve that project's frozen eval/audit and latest three-file recovery manifest. After workflow application resume the project and require complete own and dependent verification before applying it. No controller, graph, zones, coverage, audit or verification policy changes; no worker commits. The user's rebuild goal authorizes local authored execution with ordinary P04 warnings and does not assert a human merge.

## Rejected alternatives

Relying only on GIT_NO_LAZY_FETCH reproduces the demonstrated old-Git bug. Blocking all object inventory or imposing complete repository availability would reject necessary metadata and independent claims. Disabling dependent/native checks, editing project frozen evals, or changing public error/ownership semantics would hide the regression. A live server, credentials or bootstrap delegation is unnecessary for marker-only transport evidence.
