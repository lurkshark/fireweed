# Show history CLI — 0127

APP-132–135 applied at909a1b1a4531787ba778d6dccf3e33fcc8b7030f, run
20261030T231146-e5304d598617e8a0a182fbc5, audit_rounds0. APP-008/025 gained
only non-history qualifiers, with their existing assertions preserved as
already_satisfied. Native phases:

| Phase | Controller commit | Actual result |
| --- | --- | --- |
| Evals | a695ae5fe97af25029ecd45fad81b6f1a3f5278d | 282 cases:18 genuine unsupported-history failures/264 passes |
| Audit | 3c49293189430e409f40dfecd3316e8646e6d214 | Independent isolated Sol/high PASS for APP-132–135 and APP-008/025 |
| Implementation | 909a1b1a4531787ba778d6dccf3e33fcc8b7030f | 282 passes; native application/archive/claim release |

Owner-approved [clause recovery](show-history-0127-checkpoint.md) preserved only
public inputs, restored the exact named runtime document/two eval edits, natively
abandoned12d67bc and re-authored74798e2 with a fresh run/normal allocation. The
one-time exception is consumed. There was no discarded implementation or reuse.
Fresh fixed Sol/high eval/impl workers and a separate isolated assertion auditor
were used. Complete auditor commands stayed within supplied snapshot public
inputs; no SYSTEM/implementation/history/outside read, code execution or write.
The snapshot was removed. Evals/interfaces remained frozen during impl. Only
the existing APP-001 blind recorder loop was extended; no new observer.

The app is a thin project.SpecHistory caller with labelled ordered plain output,
no-history message, lowercase schema1 JSON rows, RFC3339-or-empty dates and
corrective coded failures. Current/projected show behavior is retained. Root
review/focused tests/diff check passed. Both native own/full runs stayed below5m
(and therefore below the full15m limit); existing B02 APP-103/17 and APP-104/19
warnings retain explanations. Corrected author static check has no new C17.

Clean CGO0 Go1.27.2 darwin/arm64 build and app vet passed. Binary metadata names
909a1b1a4531787ba778d6dccf3e33fcc8b7030f with vcs.modified=false. The known
nonfatal Go module stat-cache write warning was retained, build exit0.

Actual NEW proof passed on its first execution with8 commands from a nested
module directory. Authored archive records deliberately opposed filename and Git
creation order to implementation Provenance order. Filtered/module-wide history,
removed-current-spec lookup, complete metadata, empty current-ID/module arrays,
U02 lookup failures and incompatible-flag U01 all passed in plain/schema1 JSON.
HEAD, refs, index and working status were unchanged, no hooks/agents ran and no
claim existed. Fixture/state removed. These archives are public authored inputs,
not executed lifecycle/drill evidence. Logs/results:
/tmp/fireweed-show-history-cli-{result,commands}.json,
/tmp/fireweed-show-history-cli-{run,build,vet}.log, build-metadata.txt;
/tmp/fireweed-0127-replacement-* contains worker/native phase logs.

Progress review: archive-backed history now closes this small CLI slice, adding
partial R05/R17 evidence. The fifth0122–0126 gate and reviewed push5610268 remain
valid historical evidence. This caller is local application1 since that passing
review; no further push occurred. Tree archived/unclaimed; stop as approved.
Remaining drill/gc/migrate, show--stale/check--fix, recipes/versioning/packaging
and unclosed ACCEPTANCE gates remain. No final-source race/fresh-checkout/CI/
doctor/full-original-history, real regeneration, isolated self-hosting, delivery
or whole-goal completion is claimed.
