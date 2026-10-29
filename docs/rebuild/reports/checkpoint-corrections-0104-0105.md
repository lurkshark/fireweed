# Checkpoint0099–0103 scoped corrections — PASS

At clean source4cf21a1014122eb1f4c7b35f3f7d81a480eba4bd, the same separate
Sol/high checkpoint_0099_0103 reviewer returned PASS for the recorded mixed-kind
fail-first bypass. Review was read-only: no tests/product/controller commands,
mutations or drills. It confirmed hard C07 rejection before Start claims,
Verify hooks, Done commits and next agents, plus public regressions and accurate
cause/action guidance. No remaining scoped actionable finding.

## Applied corrections

0104 PRJ-046 adds six lines to central project.Check rejecting presence of
eval_fix with fixes or any Added/Modified/Removed edits. Public pure-kind controls
remain. Phases49c71a3/9aff0a7/117689f/8139419/16f61a9: initial529 cases
3 failing/526 passing; audit named Modified/Removed clauses unasserted; corrected
531 cases5 failing/526 passing, restricted re-audit PASS, implementation531
passing in11m33s. Own82 named events passed. Same run/ID, audit_rounds1.

0105 AST-033 extends only C07's explanation string, retaining Motivation/module
causes/actions. Phases2178977/ace9f64/5ac1f3c/9a59aa1/4cf21a1:592 cases
1 failing/591 passing before and after one native assertion-correction round;
restricted re-audit PASS; final592 passes in3m38s. The correction adds explicit
relationships to both cause/action assertions while preserving earlier checks.
Own61 named events pass. All isolated auditor command lists inspected; only
permitted snapshot reads, no implementation/SYSTEM/history/outside reads or code
execution/writes. Snapshots removed. Workers used the fixed real Sol/high CLI;
initial sandbox initialization refusal occurred before model dispatch and was
resolved using authorized process permissions, not a model/access retry.

## Actual NEW and fresh checks

Clean CGO-disabled NEW records Go1.27.2 darwin/arm64, source4cf21a1,
vcs.modified=false. Build and vet pass. Full race passes with no race reports:
app222.839s, assets19.169s, checks244.662s, project1.955s, runner216.978s,
workflow119.853s; unchanged contracts/gitstore cached. Static check passes with
existing explained C01 embedded-recipe warning; B02 CHK-055/16-case warning
remains explained. Explicit75f05f8fe58ca40b78a0d1ce4b86f4bdfad7f90b..
4cf21a1014122eb1f4c7b35f3f7d81a480eba4bd history passes16 commits with zero
diagnostics/squashed entries, extending recorded valid ranges.

The unchanged actual Go mixed eval_fix:[ALP-001]/fixes:[ALP-002] fixture at
baseline295f9411dab6d4aa35fa2832e016c097b39bd9c2 was incorrectly accepted by
prior NEW check. Rebuilt NEW now returns check exit1 and next exit1 with hard
C07, schema1/null data/empty stderr. Standalone explain C07 returns0 with the
new cause/action. Exact original document/HEAD retained, clean tree, no run/phase,
claims, agents or archives. The harness initially expected next exit2; corrected
to APP-077's ordinary lifecycle-failure exit1 and reran the same unchanged fixture.
No product change or assertion weakening. Fixture removed after PASS.

Local details: /tmp/fireweed-corrections-0104-0105-{build.log,build-info.txt,
vet.log,race.log,static.json,history.json}; /tmp/fireweed-mixed-kind-cli-
{before.json,result.json,commands.json,harness-exit-mismatch.log}.

Owner approved the blocked batch push and exactly this two-change correction
gate. Both corrections, actual NEW proof, native/independent audits and scoped
checkpoint/fresh checks pass. Seven applications are now reviewed; count resets
to0 only at this passing gate. Push is authorized, then return immediately to
app release wiring and actual NEW proof. No final CI/doctor/fresh-checkout,
additional real drill, self-hosting or overall acceptance claim; goal incomplete.
