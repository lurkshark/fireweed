# Spec history query — 0126

PRJ-051–056 applied at8cd9c0e510bdab0dc3f3ff9cb49fadecb28f9bcb, run
20261030T214954-ce905353e72b33c260a183b8, audit_rounds0. Native evals commit
1e76b97eba7c99969445ac90ff02d4af810daf06 recorded681 cases:11 genuine stub
failures/670 passes. Independent restricted Sol/high audit PASS was persisted
as e9ed045b80f4a1a5bc2a80e21d230df65fbd1956. Native full implementation
verification recorded681 passes, then applied/archive/released the claim. Own
project checks stayed below5m; full checks below15m. Existing B02 APP-103/17,
APP-104/19 and CHK-055/16 retain their prior explanations.

The thin facade/model exposes archived Added/Modified selection, exact section
metadata, latest valid implementation Provenance dates with lexical fallback,
module-wide unique IDs, current/removed-ID lookup and C01/C03/U02 causes. Behavior
is limited to the new internal/history reader; no dependency change. Focused
public tests and diff check passed. Evals/interfaces stayed frozen during impl.

Fresh fixed Sol/high workers and a separate isolated assertion auditor were used.
The complete auditor command list stayed within supplied change/projected SPEC/
evals/public API/model reads, with no SYSTEM/implementation/history/outside read,
code execution or write. Snapshot removed. Logs: /tmp/fireweed-0126-*.

This is application5 since the passing0117–0121 review. Stop product authoring
for the approved separate Sol/high quality/goal review and fresh checks over
0122–0126. Push only on PASS; any finding/block requires owner stop. Then immediate
app show --history and actual NEW proof, stopping with that caller local at1.
The query alone does not close the CLI slice or whole-goal acceptance. No drill.
