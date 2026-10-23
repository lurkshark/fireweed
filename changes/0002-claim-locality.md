---
fixes:
    - FLOW-011
module: workflow
phase: impl
run: 302e5571f4dc2b13329d9ff6
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

Selected object availability must be established without triggering transport, including older Git versions which ignore the environment hint. All operations retain bounded local subprocesses and literal argument handling. Object-name metadata does not authorize reading arbitrary blob bodies or decoding unrelated claims. A claim payload may be unreadable while exact-token ReleaseClaim remains valid; an unrelated bad claim must not block an independent CreateClaim or selected LoadClaim. Claims inventory still fails atomically when a required entry is unavailable. LoadClaim/Claims require available stored commits; UpdateClaim with an available requested replacement MAY either reject unavailable old payload/base with corrective zero/P02 or prove the exact previous content-addressed token and complete a correct no-fetch CAS update. Missing requested replacement commits are errors; missing unrelated trees/blobs/ancestors must not invent a stronger baseline requirement. No production code may adapt itself to the invalid legacy metadata observers.

This author-ordered0002 fix runs before the unclaimed committed project0041 implementation. Preserve that project's frozen eval/audit and latest three-file recovery manifest. After workflow application resume the project and require complete own and dependent verification before applying it. No controller, graph, zones, coverage, audit or verification policy changes; no worker commits. The user's rebuild goal authorizes local authored execution with ordinary P04 warnings and does not assert a human merge.

Pre-freeze restricted review found four real oracle gaps: mandatory rejection of otherwise provable updates, nil-only failed inventory, unclassified new object storage, and unchecked unrelated-claim batch metadata requests. Corrected before freeze: allow exact no-fetch CAS success or corrective rejection when only old content is missing; accept empty inventories; independently decode/hash and bound only selected-operation blobs while preserving existing storage and rejecting garbage; validate explicit batch-check/info requests while permitting OID-only inventory. Initial full matrix produced153 records/60 genuine transport leaf failures and no other failure details, but is superseded by the corrected-source proof. A subsequent pre-freeze review found that returned success tokens also needed validation against preexisting or missing local bodies. Success now independently verifies the selected local blob hash and exact returned fields directly from snapshot bytes; negative controls reject fabricated success with an unchanged dangling token or unrelated existing blob. The first controller preflight stopped before coverage with P2-01 because an explanatory parenthetical had accidentally been added to Interface. Restored the exact approved Interface text and kept clarification only in Decisions; no policy bypass or phase commit. No frozen eval or implementation edits.

The complete-package preflight reached the old Go10m harness timeout at600.279s with1893/1909 expected records; its mechanically acceptable fail-first diagnostics are not accepted complete proof. Nine files were preserved and the own claim released. Separate protected author commit42b7f5f raises only the overall package test budget to20m; fresh doctor canaries all pass, recorded7df0402. This new EVALS run restores identical eight eval files and only allowed Decisions, retaining its new run/claim State. The first restoration script had a syntax error before any writes; a mistakenly launched subsequent preflight was terminated and excluded. Restoration and hash verification were then completed separately before verification. No phase commit, frozen-eval or implementation change was involved in recovery.

The restored complete preflight finishes normally in611.822s with1909 records/61 failed records/60 genuine transport leaves; all1728 legacy and120 new controls pass,0 skips/errors/diagnostics. Legacy-name comparison normalizes full fixture OIDs and only the12-hex fixture token in Test_FLOW009FullReleaseTokenRequired. Before freezing, add four available-base fixture cases (five records including parent) so tag/ancestor unavailability cannot be satisfied by rejecting those forms wholesale: available annotated/nested tags, ancestor updates, range rejection and unrelated-state preservation in ordinary/linked SHA-1/SHA-256 repositories. Final done must rerun the complete1914-record exact source, not reuse the preceding1909-record report.

The available-base draft review caught two further assertion gaps before freeze: tags targeted HEAD, and object storage was only checked at the end. Corrected tags/nested tags now target the distinct older local commit. Every individual call preserves existing object entries and classifies only bounded exact selected states/attempt blobs; unrelated refs, files, temp and parent state remain checked. Corrected valid-base and existing independent CAS/storage/token controls pass10.674s. No product implementation edits.

## Rejected alternatives

Relying only on GIT_NO_LAZY_FETCH reproduces the demonstrated old-Git bug. Blocking all object inventory or imposing complete repository availability would reject necessary metadata and independent claims. Disabling dependent/native checks, editing project frozen evals, or changing public error/ownership semantics would hide the regression. A live server, credentials or bootstrap delegation is unnecessary for marker-only transport evidence.

## Audit

No findings. Checked FLOW-011, related outcome fidelity FLOW-008, FLOW-009 and FLOW-010, and retained observer fidelity FLOW-020 and FLOW-023.

Reviewed the required-unavailable-object matrix and its independent Trace2 and marker-only transport controls, including denied transport, snapshot/reset ordering, ordinary and linked SHA-1/SHA-256 fixtures, and required versus irrelevant missing objects. The assertions retain corrective zero/P02 rejections and atomic empty inventories, permit conforming exact-CAS updates when only old content is unavailable, and require exact-token release. Successful returned tokens are checked against local blob bytes and hashes. Per-call storage checks preserve existing entries, permit bounded selected claim/attempt blobs, and reject unrelated objects and garbage. Available tag, nested-tag and ancestor controls select the distinct older commit; range rejection does not replace those positive outcomes.

Reviewed actual-operation and config-operand parsing, real-success byte-equal marker-free selected commit/OID-only inventory/no-walk/root/subtree controls, and retained private-body/batch, mutable-source, unrelated-claim, mutation and network denials. Native reference-transaction hook and fsmonitor controls, configured hook/agent/bootstrap traps, literal arguments, identity independence, and complete ref/index/HEAD/file/temp/input/parent-state preservation remain asserted. This is a genuine FLOW-011 regression without changed public claim or approval outcomes. The audit used only permitted public contracts and eval source; it did not execute tests or inspect implementation, history, bootstrap or evidence.

## Provenance

- evals | 2026-10-23T19:10:10Z | agent: manual | fireweed: 0.4.0 | playbook: 5867363302d4c874074520977b1627d26991d95ee2fbacb9ff63e827661edd0b | tests: 1914 | failing: 61 | passing: 1853
- audit | 2026-10-23T19:15:07Z | agent: manual | fireweed: 0.4.0 | playbook: a9c982d24b476621316049b527e6c7bbc319f06b12268af150fa7d8c6df92230 | tests: 0 | failing: 0 | passing: 0
