---
fixes:
    - GIT-002
module: gitstore
phase: impl
run: 3777b3658f1b2362845d933f
---

## Motivation

The unchanged design requires local Git inspection without fetches. Independent public marker-only discovery on installed Git2.30.1 confirms BaseRef attempts transport when an existing preferred tracking ref names an unavailable commit, then incorrectly returns the local fallback. Edits attempts transport before rejecting an unavailable selected commit or root tree. The same discovery proves a simple modification with an unavailable old blob and a comparison with an irrelevant missing parent can succeed without transport. Preserve those available results rather than requiring a complete object graph.

The environment hint GIT_NO_LAZY_FETCH=1 is ineffective on this installed supported Git. Direct experiments also show remote.promisor=false overrides do not suppress its fetch attempts. A useful error after an attempted fetch is still a violation. This change fixes the existing base-selection contract and specifies the local-only behavior of edit inspection without changing its API.

## Interface

No public API/type changes. Keep Open, BaseRef, Clean, Edits, SafePath, Run and Capture signatures and all existing contracts. Add permanent public-API regressions for GIT-002 and GIT-011, plus independent real-Git controls proving marker-only transport and Trace2 fetch observations in allowed and denied protocol modes. Exercise ordinary/linked SHA-1/SHA-256 repositories supported by installed Git, required and irrelevant missing objects, valid single commit/tag/ancestor inputs, revision-set rejection and available built-in diff/index/worktree configuration and rename behavior. Compare successful available results against independently executed real Git rather than an implementation-specific command allowlist. Preserve all existing evals and their genuine success controls.

## Added

### GIT-011
Edits MUST inspect the requested local baseline without attempting fetch, push, branch creation or network commands, including partially populated repositories on supported Git versions. It MUST return a useful non-nil error and an empty edit result when objects required to resolve the selected commit or determine its edits are unavailable; it MUST NOT return partial successful edits. Missing unrelated ancestors, trees or blob bodies that are not needed for the requested comparison MUST NOT prevent otherwise available results. Available ordinary and linked SHA-1/SHA-256 repositories MUST retain the existing GIT-003 staged, unstaged, untracked, filename and rename semantics, including built-in Git configuration affecting index/worktree comparisons. Full commit IDs and valid single commit/tag/ancestor expressions MUST remain usable; revision sets MUST be rejected. Successful and rejected calls MUST remain noninteractive, use literal arguments, preserve HEAD, refs, original index bytes, working/ignored files and parent cwd/environment, and clean owned temporary storage. No bootstrap or authenticated agent MAY be required or invoked.

## Modified

## Removed

## Decisions

BaseRef chooses an existing exact tracking ref by local ref identity rather than requiring that ref's target commit body. An unavailable target is not permission to fetch or silently substitute a different base. Preserve configured/default remote/branch selection, slash-containing names, symbolic remote defaults and absent-tracking fallback under GIT-002. Existing edit ordering, rename source/destination and filename-byte contracts remain authoritative.

Availability is scoped to the requested operation. Do not validate every reachable commit/blob or decode unrelated content as a precondition. Prove the unavailable-object transport observer with real local Git controls, including a denied protocol where a fetch child may still be attempted. Clear controls before API calls and snapshot all state after fixture construction. Native Git output on available fixtures supplies the comparison result; a reference command that fails is not a positive control.

A temporary minimal Git metadata view avoided transport in an exploratory four-case experiment, but does not establish a production design or configuration compatibility. Any chosen implementation must preserve required built-in comparison configuration, linked-worktree/index/object-format and rename behavior, retain useful errors, and contain/clean only its own temporary resources. Neither original repository configuration nor user environment may be rewritten to suppress the bug. Evaluate implementation only after complete fail-first and an independent restricted audit.

The user's full rebuild goal authorizes local authored execution and commits with the controller. Manual start's normal P04 warning does not assert a human merge or waive approval policy. Keep the pinned bootstrap controller separate from the NEW subject; no pushes or publication. Subsequent workflow Verify composition must also fix its independently confirmed approval/exact-HEAD precondition fetches through genuine regressions, not through this module's worker edits.

Pre-freeze restricted review found and corrected three oracle defects before any eval phase commit: live marker controls must explicitly allow lazy fetch on newer Git; command and ref-transaction traces must detect transient branch creation and direct pack transport commands; preservation assertions must follow the promised HEAD/ref/index/file state rather than require unchanged unrelated object-store/cache layout. Native allowed/denied controls, transient branch/update-ref controls and actual pack-command controls now calibrate observers independently. These preliminary reviews are not formal Audit acceptance.

Forty available configuration leaves across ordinary/linked SHA-1/SHA-256 fixtures independently compare successful native Git against manual premises: autocrlf, included config, info/external attributes, info/global excludes, split-index, timestamp-only, literal names with staged/unstaged edits and executable modes. All preliminary configuration and observer cases pass. Full frozen fail-first evidence remains pending; no product implementation has been changed.

The first complete preflight runs202 records/87.056s with36 genuine terminal transport failures (8 preferred-base and28 required-edit-object leaves), all41 legacy and other controls passing, no skips/errors. Controller correctly rejects that report with C14 because the new required-object group accidentally also carried unchanged GIT-003 in its test name. Correct the ID association to GIT-011 only; the passing available configuration controls retain GIT-003. This report is excluded from accepted fail-first proof. No assertion or contract is weakened, and no implementation was edited.

A second restricted pre-freeze review found five additional assertion defects: invalid BaseRef strings were overconstrained beyond GIT-002; valid branch format/sort operands and checkout/switch literal option-looking names were misclassified; verbose remote-update commands escaped classification; configuration rewrite preservation was omitted; and shell-looking baseline interpolation lacked an explicit trap. Corrected all before freezing, retaining actual native success/error controls and adding four literal baseline forms. The intermediate202-record corrected-ID verify passes enforcement/diagnostics but is superseded by these strengthening corrections and is not the final frozen proof.

The220-record strengthened preflight completes normally in73.112s with enforcement diagnostics[],36 terminal fetch failures and all controls passing. Final pre-freeze review then identifies repository-local scratch omitted by the narrowed metadata snapshot. Add new-path cleanup checks under both CommonDir and selected GitDir, retaining unrelated object/native-cache layout freedom and existing shared-index bytes. Live empty-directory controls cover both locations and TMP; an object-store-layout control proves that allowance. No product code has changed; the final accepted report must run this exact source.

A final narrower cleanup review finds the native-object-cache allowance also exempted temporary object/pack files. Keep actual immutable object additions allowed (proved by native hash-object), but include tmp_* and lock artifacts in cleanup observations, with live loose-object, pack and commit-graph-lock controls. An explicit external-command trap liveness control also calibrates the shell-looking baseline checks. Native linked shared-index layout allowance is also proved directly; classify its real hash filenames at metadata depth rather than exempt arbitrary sharedindex-like names. These are pre-freeze corrections, not changes to product behavior or weakened assertions.

Final exact-source verification completes normally in74.328s with221 records,38 failed records/36 genuine terminal fetch leaves (8 GIT-002,28 GIT-011),41 legacy and142 other passing,0 skips/errors/diagnostics. Exact inventory and all seven eval/interface plus six untouched implementation hashes are checked independently. Final restricted pre-freeze review reports No findings for GIT-002/GIT-003/projected GIT-011, including the native shared-index allowance; this remains distinct from the formal Audit phase. Go vet and whitespace checks pass. Controller done must repeat enforcement and freeze these evals before any implementation.

## Rejected alternatives

An environment hint or failed network request after the fact does not prevent fetch. Requiring all repository objects would reject the demonstrated valid comparisons. Returning a local branch after a broken preferred tracking target changes the configured base. Weakening frozen dependent/native evals, losing built-in comparison behavior, editing another module or calling the bootstrap to implement the product would hide rather than repair the problem.

## Provenance

- evals | 2026-10-23T22:42:27Z | agent: manual | fireweed: 0.4.0 | playbook: 5867363302d4c874074520977b1627d26991d95ee2fbacb9ff63e827661edd0b | tests: 221 | failing: 38 | passing: 183
- audit | 2026-10-23T22:44:49Z | agent: manual | fireweed: 0.4.0 | playbook: a9c982d24b476621316049b527e6c7bbc319f06b12268af150fa7d8c6df92230 | tests: 0 | failing: 0 | passing: 0

## Audit

No findings. Checked GIT-002, projected GIT-011 and retained GIT-003 compatibility.

Reviewed preferred tracking-ref identity despite unavailable targets, configured/default remote selection, symbolic defaults, slash-containing branches and absent-ref fallback. Reviewed required-object rejection with useful errors and empty results, independent available comparisons with irrelevant missing objects, older tag/nested-tag/ancestor selection, revision-set rejection, and ordinary/linked SHA-1/SHA-256 fixtures.

Independent native controls calibrate transport markers, denied-protocol Trace2 detection, transient branch/ref transactions, command-option parsing and external-command traps. Successful native/manual premises retain staged, unstaged, untracked, filename, rename, configuration, attributes, excludes, modes and split-index behavior. Assertions preserve HEAD, refs, original index/configuration, working/ignored files and parent state; cleanup controls cover TMP, repository-local scratch, temporary object/pack files and locks while permitting legitimate object/shared-index storage.

This audit read only the permitted change, public contracts/API and eval source. No tests or commands were executed, no implementation/history/bootstrap/evidence was inspected, and no files were edited.
