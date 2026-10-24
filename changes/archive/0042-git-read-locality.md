---
fixes:
    - GIT-002
module: gitstore
phase: applied
run: 43bb6e51f2e637a7c66afc35
---

## Motivation

The unchanged design requires local Git inspection without fetches. Independent public marker-only discovery on installed Git2.30.1 confirms BaseRef attempts transport when an existing preferred tracking ref names an unavailable commit, then incorrectly returns the local fallback. Edits attempts transport before rejecting an unavailable selected commit or root tree. The same discovery proves a simple modification with an unavailable old blob and a comparison with an irrelevant missing parent can succeed without transport. Preserve those available results rather than requiring a complete object graph.

The environment hint GIT_NO_LAZY_FETCH=1 is ineffective on this installed supported Git. Direct experiments also show remote.promisor=false overrides do not suppress its fetch attempts. A useful error after an attempted fetch is still a violation. This change fixes the existing base-selection contract and specifies the local-only behavior of edit inspection without changing its API.

## Interface

No public API/type changes. Keep Open, BaseRef, Clean, Edits, SafePath, Run and Capture signatures and all existing contracts. Add permanent public-API regressions for GIT-002 and GIT-011, plus independent real-Git controls proving marker-only transport and Trace2 fetch observations in allowed and denied protocol modes. Exercise ordinary/linked SHA-1/SHA-256 repositories supported by installed Git, required and irrelevant missing objects, valid single commit/tag/ancestor inputs, revision-set rejection and available built-in diff/index/worktree configuration and rename behavior. Compare successful available results against independently executed real Git rather than an implementation-specific command allowlist. Preserve all existing evals and their genuine success controls. Include a live configured agent-shaped fsmonitor callback control, and require both successful and rejected Edits calls to avoid that callback while preserving native built-in comparison behavior and original configuration/index state.

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

The initial phase cycle was completely reverted under D037 after a configured native fsmonitor control exposed a missing permanent assertion. Its221-case fail-first and passing formal audit are retained historically but do not freeze this restarted cycle. Restore the three preserved eval sources only after fresh EVALS start, add a calibrated configured-monitor regression using a compiled local agent-shaped marker (no authenticated agent), and repeat complete fail-first and restricted audit. A selected original index metadata read can invoke the monitor even when an owned metadata view disables it later. Success and rejection paths must prevent that callback. Preserve the candidate separately and restore its two implementation files only after a new passing audit and impl context; never restore its old run/phase/Audit.

Fresh EVALS restored only the three SHA-256-verified saved eval files and adds configured monitor available/rejected controls across ordinary/linked SHA-1/SHA-256 repositories. The focused nine-record test completes normally with four genuine available API codex-marker failures and four rejected-selector passes; all native enabled/disabled monitor premises are live. No implementation was restored or edited. The new start retains the normal P04 warning under the user's local authorization, without claiming human merge. Full230-record fail-first, new independent formal audit and full implementation verification remain required.

Restricted pre-freeze review strengthens configured-monitor coverage further: require a nonempty manual/native staged+unstaged+untracked result, and add unavailable selected commit/root-tree rejections with live promisor controls rather than only an invalid selector. These16 monitor leaves supersede the initial eight-leaf monitor draft; all prior locality/configuration/native controls remain. The complete inventory is now238 records. No product implementation is present or edited.

Final complete fresh-cycle verification runs238 records in86.538s with51 failure records/48 genuine terminal leaves:44 attempt fetch and8 invoke the agent-shaped monitor, with4 overlapping. All41 legacy and146 other records pass;0 skips/errors/diagnostics. Exact testcase inventory, seven eval/interface hashes, six unchanged implementation hashes and no new implementation files are independently checked. Restricted pre-freeze review reports No findings after the nonempty-result/required-object monitor strengthening. Explicit setup..b4195b2 history passes395 commits/inner+outer diagnostics[].

Independent native split-index controls show core.fsmonitor=false is interpreted as an external false command on installed Git2.30.1, while an empty value suppresses callbacks and retains the exact native shared-index path. Any implementation choice must actually prevent callbacks rather than rely on modern boolean interpretation. This native experiment changes no eval/product source. EVALS done must repeat full enforcement before fresh formal audit.

After fresh formal audit, restore only the two SHA-256-verified disposable candidate files and suppress the original shared-index metadata callback plus owned-view monitor via the independently verified empty configuration value. The238-case frozen evals/interface remain unchanged. The new17-case configured-monitor matrix passes in10.452s, including nonempty available edits and all rejection paths; vet/whitespace pass. Effective built-in config, original object storage, existing ref identities, contained metadata, selected index/shared-index, attributes/excludes and shallow/graft context supply native comparisons without promisor fetch settings. Original configuration/index/refs/files remain unmodified; owned view storage is removed on every path. Full own/dependent verification and a fresh scoped race run on this corrected source are still pending.

An ordinary public compatibility probe during full verification confirms a candidate bug on the host's case-insensitive filesystem: valid packed tags topic and TOPIC select distinct commits natively, but mirroring them as loose ref files merges their names and returns the older diff for TOPIC. Stop and exclude that incomplete verify; keep the fresh238-case frozen eval/audit unchanged. Preserve existing ref identities in a sorted owned packed-refs representation instead of loose filenames, without creating a new branch identity or modifying original refs. Rerun the independent native/public selection probe and complete exact-source verification before application. This is an implementation correction within the existing valid-single-commit/tag contract, not an eval/contract weakening.

The same public/native compatibility probe confirms reflog selectors HEAD@{1}, main@{1}, symbolic upstream/HEAD@{1}, and ORIG_HEAD require their existing local metadata. Preserve symbolic aliases, selected-worktree pseudorefs and requested reflog context in the owned view. All six native/public comparisons now agree, including case-distinct packed tags. These are private implementation corrections; frozen238-case evals/API and formal Audit remain unchanged. Prior interrupted full verifications are excluded. Fresh exact-source race and complete controller verification are required.

## Rejected alternatives

An environment hint or failed network request after the fact does not prevent fetch. Requiring all repository objects would reject the demonstrated valid comparisons. Returning a local branch after a broken preferred tracking target changes the configured base. Weakening frozen dependent/native evals, losing built-in comparison behavior, editing another module or calling the bootstrap to implement the product would hide rather than repair the problem.

## Provenance

- evals | 2026-10-23T23:12:36Z | agent: manual | fireweed: 0.4.0 | playbook: 5867363302d4c874074520977b1627d26991d95ee2fbacb9ff63e827661edd0b | tests: 238 | failing: 51 | passing: 187
- audit | 2026-10-23T23:14:02Z | agent: manual | fireweed: 0.4.0 | playbook: a9c982d24b476621316049b527e6c7bbc319f06b12268af150fa7d8c6df92230 | tests: 0 | failing: 0 | passing: 0
- impl | 2026-10-24T00:15:08Z | agent: manual | fireweed: 0.4.0 | playbook: 653912b194316c62181f62d0a69dc213f7668481ea8b1b9926ef889dd48a9e08 | tests: 3317 | failing: 0 | passing: 3317

## Audit

No findings. Checked GIT-002, projected GIT-011 and retained GIT-003 compatibility.

Reviewed preferred tracking-ref identity despite unavailable targets, remote/default selection, symbolic defaults, slash-containing branches and absent-ref fallback. Required-object cases assert useful errors, empty results and no transport; independent native controls preserve available comparisons with irrelevant missing objects, older tags/ancestors, renames and built-in configuration across ordinary/linked SHA-1/SHA-256 fixtures.

The configured-monitor matrix proves a nonempty manual/native staged, unstaged and untracked result. A live agent-shaped callback and paired native override calibrate marker observation before reset. Public success, invalid-selector rejection, unavailable-commit rejection and unavailable-tree rejection independently require no agent/transport marker and preserved state.

Reviewed literal-baseline traps, revision-set rejection, actual-operation parsing, transient branch/ref observation, configuration/index/file/parent preservation, and TMP/repository-local temporary-storage cleanup. Native object and shared-index controls retain legitimate storage-layout freedom without exempting observed temporary files and locks.

This audit read only the permitted change, public SPEC/API and eval source. No tests or commands were executed, no implementation/history/bootstrap/evidence was inspected, and no files were edited.
