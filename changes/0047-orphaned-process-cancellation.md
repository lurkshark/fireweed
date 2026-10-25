---
fixes:
    - GIT-006
    - GIT-010
module: gitstore
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

## Rejected alternatives

Current-parent enumeration loses children after reparenting. Killing only the original process group misses independently grouped or session-separated work. Broad same-image, caller-session, user or orphan sweeps can kill unrelated processes. Returning before an owned orphan stops permits delayed writes. Weakening the existing contract or editing frozen evals during implementation would conceal the remaining bug.
