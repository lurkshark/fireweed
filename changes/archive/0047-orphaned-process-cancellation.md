---
fixes:
    - GIT-006
    - GIT-010
module: gitstore
phase: applied
run: 358434da4d3867892641e21f
---

## Motivation

The archived live-parent fix0046 stops independently grouped descendants while their ancestors remain discoverable. A separate bounded public-API probe proves a remaining case: the outer command has exited, its independently grouped child is a native orphan with parent PID1, and the outer Run/Capture is still waiting for inherited streams. Both adapters return context.Canceled with Exit-1 after cancellation, yet that child writes after an arm marker created only after the outer return. This still violates the existing entire-owned-tree and no-delayed-write contracts.

Keep cancellation ownership across parent exit and reparenting, including independent process groups and native independent sessions where supported. Preserve ordinary results, retained pretermination output, caller inputs/environment/cwd, unrelated work and overlapping calls. Do not narrow GIT-006/GIT-010 to living parents, a single process group, or processes with a matching executable name.

## Interface

## Added

## Modified

## Removed

## Decisions

This is another fixes-only regression for existing GIT-006 and GIT-010. Public signatures/models, module zones, spec bodies and existing evals remain unchanged. Add permanent public-API tests before implementation. Establish real child readiness and native parent termination/reparenting before cancellation; confirm the outer operation is still pending. Exercise Run and Capture, cancellation and genuine deadline, independently grouped native orphans and independent sessions on platforms with that concept. Use bounded helpers and exact owned cleanup. Positive controls must prove orphan creation, retained streams and delayed-write observation; do not treat a premature outer return or failed helper startup as a successful cancellation test.

A delayed-write marker must be armed only after outer return and written before any pipe output, so a closed stream cannot masquerade as process termination. Include a live unrelated same-image orphan and overlapping public work that survive subject cancellation. Preserve exact context error identity and useful pretermination output, and distinguish ordinary success/nonzero exits. Keep host runtime and cross-compilation evidence separate. Deadline fixtures may coordinate parent exit near an already established real deadline, but must not synthesize a fake DeadlineExceeded error or rely on unproved startup timing.

The ordinary discovery artifacts orphaned-group-cancellation-probe.{go,json} and -proof.json are evidence of the bug, not permanent evals or regeneration. An unchanged public probe observes Run1012ms/Capture1029ms, matching cancellation/Exit-1, native independent child and parent PID1 before cancellation, outer pending, and delayedwrite true for both. Keep any chosen implementation strategy separate from observable test requirements and independently audit the exact assertions after complete genuine fail-first.

The rebuild goal authorizes local authored execution and controller-owned phase/archive commits. Use only the pinned bootstrap as controller, with independently built NEW subjects; leave the reference checkout unchanged. Record ordinary local unmerged warnings without claiming a human merge. No push, publication or release tags.

The implementation retains a cancellation watcher for the complete public call, including inherited streams after the direct command has exited. On macOS, native creation/lifetime identities and stable original-parent identities locate reparented owned descendants; cancellation freezes live parents before child discovery and checks identities before signals. On Linux, each call uses a separate statically linked helper with PR_SET_CHILD_SUBREAPER, so adopted descendants remain scoped to that call across parent exits, groups and sessions. Only this helper becomes a subreaper. It reaps its own children, forwards exact streams and reports the requested executable's outcome over private close-on-exec descriptors. The target retains its original argv, environment and working directory. Combined Run streams share one target pipe; Capture keeps two. The parent loads an architecture-specific, source-generated compressed helper through an anonymous memfd. No calling-program initializer, shell, Go compiler, bootstrap, external helper installation or network is involved at runtime. Helper source and reproducible generation instructions live in the same disposable implementation zone; tracked images support all Go Linux architectures. Ordinary uncancelled completion follows root-exit and stream-drain semantics. Windows retains its existing suspended-start job ownership.

## Rejected alternatives

Current-parent enumeration loses children after reparenting. Killing only the original process group misses independently grouped or session-separated work. Broad same-image, caller-session, user or orphan sweeps can kill unrelated processes. Returning before an owned orphan stops permits delayed writes. Weakening the existing contract or editing frozen evals during implementation would conceal the remaining bug.

## Provenance

- evals | 2026-10-25T06:59:54Z | agent: manual | fireweed: 0.4.0 | playbook: 5867363302d4c874074520977b1627d26991d95ee2fbacb9ff63e827661edd0b | tests: 499 | failing: 14 | passing: 485
- audit | 2026-10-25T07:02:58Z | agent: manual | fireweed: 0.4.0 | playbook: a9c982d24b476621316049b527e6c7bbc319f06b12268af150fa7d8c6df92230 | tests: 0 | failing: 0 | passing: 0
- impl | 2026-10-25T09:22:47Z | agent: manual | fireweed: 0.4.0 | playbook: 653912b194316c62181f62d0a69dc213f7668481ea8b1b9926ef889dd48a9e08 | tests: 3993 | failing: 0 | passing: 3993

## Audit

No findings. Checked GIT-006 and GIT-010.
