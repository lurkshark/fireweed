# Current milestone

**Close full acceptance after v0.2.0.** The released `origin/main` at
`f42333c` includes the acceptance fixes, audited release eval updates and
version change. [PROGRESS](PROGRESS.md) records the evidence for this baseline.

[ACCEPTANCE](ACCEPTANCE.md#full-completion-gates) tracks the remaining gates.
Gates **1, 3, 4, 5 and 7** are open. Process and delivery are met for
v0.2.0, and all eight accepted module drill reports remain evidence for gate 6.

1. Review the focused R05 gap in automated agent-command Provenance and
   agree how to check history across the cutover. Before any new lifecycle
   run, build and record a controller compatible with main under
   [SYSTEM](../../SYSTEM.md).
2. Apply the authorized provenance change through Fireweed. Run the missing
   fresh-checkout Go and controller checks against the final candidate with
   the reviewed history ranges.
3. Complete the product binary exercises in fresh repositories. Run the
   separate isolated self-hosting proof with a real spec change and a real
   drill after owner approval. Review only the earlier drill proofs affected
   by relevant changes.
4. Record gate results, exact revisions and remaining limitations in
   ACCEPTANCE and PROGRESS. Reuse the v0.2.0 delivery evidence for checks
   unaffected by any later changes.

Use the applicable [acceptance criteria](ACCEPTANCE.md#full-completion-gates)
and the authorization rules in [AGENTS](../../AGENTS.md) and SYSTEM.
