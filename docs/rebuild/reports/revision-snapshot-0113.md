# Revision snapshot dependency — 0113

PRJ-047–050 add public project.LoadAt using the selected commit's Git objects,
recorded configuration/module globs/documents/defaults and canonical paths.
Internal config/module parsing is shared with Load; existing behavior retained.
No temporary checkout/ref, new public model, dependency or process primitive.
This is the minimal prerequisite for recorded-zone history checks, not a CLI
slice completion or a real drill.

Phases0d0af05/114ca12/8e14446/eed7cc4/ea83a82. Both native eval rounds594 cases8
failures/586 passes; final implementation594 passes at
ea83a821e439d26307ac25fe115562e84d394197. Full impl about755s, under15m;8 focused
new test events and6 targeted existing Load events pass. Initial audit found
PRJ-047's retained Config defaults and out-of-glob exclusion unasserted. The
12-line public correction adds an excluded recorded module and direct defaults
assertion without new cases. Native same run/IDs, round1 retained; PRJ-047-only
isolated re-audit PASS. All real workers used fixed Sol/high. Complete auditor
command inspection confirms only allowed public snapshot reads, no implementation,
SYSTEM/history/outside/code execution/writes; snapshots removed. Native archive
released claim; main tree clean. No phase-write exception.

Existing B02 CHK-05516 and APP-10317/APP-10419 warnings remain explained in linked
prior reports. C18's filename fragment "md" is punctuation in SPEC.md rather than
a copied convention. Local details /tmp/fireweed-0113-{evals-done.log,
evals-correction-done.log,audit-worker.log,reaudit-worker.log,impl-done.log,
impl-focused.json,impl-load-regression.json}. Actual NEW remains the clean0112
build until the history caller is available. Count2 since passing/pushed gate;
return now to checks history, justified H01–H04 explanations and immediate app
caller/actual NEW explicit-range proof, retaining normal checkpoint bounds.
