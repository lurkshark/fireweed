# Status-warning explanations — 0121

AST-038–041 applied at `79f71130cde2777686c0681142fdd2a89e062ca4`, run
`20261030T160730-90595c2d8e254ac14a11cde0`, audit_rounds0. Native evals
`a745fb4ab1412c58b5528b51a51651711a83c583` recorded712 cases:4 genuine unknown-
code failures/708 passes. Independent restricted Sol/high audit PASS was
persisted as `62cf807a38049060f0c4315010ab50f5897957d0`. Full implementation
verification recorded712 passes, then applied/archive/released the claim.
Own evaluation was about12s and full verification below15m.

Explain now describes B01–B03 cause/default/action and the newly emitted P08
status-warning companion. Only four entries were added; no existing explanation,
API, budget, limit or enforcement changed. Fresh fixed Sol/high workers made no
commits. Four focused cases pass. The independent auditor read only snapshot
AGENTS/change/projected SPEC/own evals/helper; no SYSTEM/private code/history/
outside access, execution or writes. Snapshot removed. Pre-freeze assertion
refinement required minimum-three P08 wording and allowed qualified budgets.*
keys; no cases were added or frozen assertions altered.

A clean CGO-disabled Go1.27.2 darwin/arm64 NEW build identifies this source and
vcs.modified=false. Actual NEW explain B01/B02/B03/P08 --json outside a project
all returned exit0, schema1/nonempty string data, empty diagnostics/stderr.
Results: `/tmp/fireweed-status-warning-explain-proof.json`. Build succeeded with
the known nonfatal sandbox denial writing Go's module stat-cache metadata;
binary revision/settings were inspected, so this is not a failed build or a
reason to change source/repeat valid evidence.

Logs/results: `/tmp/fireweed-0121-{evals-worker,evals-refinement-worker,evals-done,audit-worker,audit-done,impl-worker,impl-done}.log`
and matching focused JSON. Existing native B02 APP-103/17, APP-104/19 and
CHK-055/16 warnings retain prior explanations.

This is application5 since the last passing review. All native work is archived
and unclaimed; new product authoring/push is gated on the separate quality/goal
review and fresh appropriate checks. Caller status is deliberately next after
that PASS/push, as approved; no status CLI/full-goal completion claim here.
