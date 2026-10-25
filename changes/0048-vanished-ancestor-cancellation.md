---
fixes:
    - GIT-006
    - GIT-010
module: gitstore
phase: impl
run: d554bc098ada714fba6d6935
---

## Motivation

The direct-orphan regression0047 passes, but a separate bounded public probe proves a remaining macOS violation. Both the requested command and its intermediary have exited and been fully reaped before cancellation. Their live independently sessioned leaf has native parent PID1, and the outer Run/Capture remains pending on inherited streams. Both adapters return matching context.Canceled with Exit-1 in990/995ms, then that descendant writes after a marker armed only after return. Retaining only original-parent identities cannot connect through an already vanished, never-observed intermediary.

Preserve cancellation ownership through multiple ancestor exits, independent groups and independent sessions. Keep the entire-owned-tree contract, unrelated work, ordinary exit results, retained bytes and caller state intact. Do not redefine the contract around living ancestry, process names or a shared caller session.

## Interface

No public signature, model, zone or interface change.

## Added

## Modified

## Removed

## Decisions

This fixes-only change covers existing GIT-006 and GIT-010. Add permanent public-API regressions before implementation and preserve every existing eval, interface and spec body. Exercise Run and Capture with cancellation and genuine deadlines, independently grouped and session-separated descendants where supported. Establish root/intermediary/leaf readiness and their original parent links, then prove both ancestors have terminated and been reaped before cancellation while the leaf remains live and the outer call remains pending. On Windows use exact retained process handles for termination evidence because inherited parent IDs are not rewritten there. A supported ownership anchor may become the new parent; do not require PID1 when an isolated Linux reaper intentionally retains it.

Arm the delayed-write observation only after outer return, and write that marker before pipe output so closed streams cannot masquerade as process termination. Calibrate Native/Run/Capture success and nonzero controls against the same multi-ancestor fixture, retain pretermination stdout/stderr and matching context error identity, and prove unrelated same-image orphans and overlapping public work survive. Preserve argv/override slices, environment and cwd. Register exact owned helper cleanup immediately on readiness, wait bounded helper completion before temporary cleanup, and separate native runtime evidence from cross-compilation. Deadline coordination may release ancestors near an already established real deadline but must not fabricate a context error or substitute an unproved startup delay.

The ordinary refined discovery report orphaned-multi-parent-probe.json and its proof record the bug and native cleanup; they are not permanent evals or regeneration. Native capability and responsibility probes are mechanism research only. They do not establish product integration, Go-only fresh build/install, execution-permission compatibility or universal tree coverage. Choose the implementation only after the permanent fail-first and independent assertion audit; record material execution choices and evidence without narrowing these observable requirements.

The rebuild goal authorizes local authoring and controller-owned worker/phase/archive commits. Use only the pinned .fireweed/bin/fireweed-bootstrap as controller, GOCACHE=/tmp/fireweed-rebuild-cache, and separately built NEW subjects. The four-worker protected hook is already committed and validated through real suites and fresh postcommit doctor. No push, publication, release tags or claimed human merge.

## Rejected alternatives

Polling only current parents, or retaining only an original direct-parent edge, loses descendants once intermediary identities vanish. Killing only the initial process group misses independent groups and sessions. Same-image, caller-session, user or whole-application sweeps endanger unrelated work. Returning before an owned descendant stops permits delayed writes. Editing frozen evals during implementation or weakening the existing entire-tree contract would conceal the bug.

## Provenance

- evals | 2026-10-25T14:57:20Z | agent: manual | fireweed: 0.4.0 | playbook: 5867363302d4c874074520977b1627d26991d95ee2fbacb9ff63e827661edd0b | tests: 535 | failing: 14 | passing: 521
- audit | 2026-10-25T15:05:21Z | agent: manual | fireweed: 0.4.0 | playbook: a9c982d24b476621316049b527e6c7bbc319f06b12268af150fa7d8c6df92230 | tests: 0 | failing: 0 | passing: 0

## Audit

No findings. Checked GIT-006 and GIT-010.
