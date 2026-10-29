# Regenerate and ordinary fixes — actual NEW proofs PASS

AST-032 applied at31c28d1710be2b59f320bed619797d222dcf8adc. Native
phasesdc6da04/59499c3/31c28d1 record577 cases: one genuine failure/576
passes, then577 impl passes. Own checks:59 pass/one intended failure in11.149s,
then60 passes, zero failures/skips in10.96s. Isolated Sol/high AST-032 audit
passed; complete command log shows permitted snapshot reads only, no
implementation/SYSTEM/history/outside/report-gap content, execution or writes.
Snapshot removed. Native provenance manual; actual fixed-profile CLI logs under
20261029T191501-d2e6014c653fb0ae6a09611b. Impl adds only the P10 map entry,
preserving frozen evals/interfaces and other explanations. No actual drill.

NEW built from clean31c28d1, Go1.27.2 darwin/arm64, CGO disabled,
vcs.modified=false. Two distinct required-kind Go CLI fixtures passed.

Regenerate (closes0101's caller slice):

- Nested-directory start value-alpha --json returned0 in0.24s, with P10
  warning and regenerate playbook in its string context. Disposable value.go
  was absent, contracts retained, HEAD unchanged and no agent had run.
- Owned next --json returned0 in2.59s under the same run
  20261029T192747-0ddbce08d185cfd2cef7b1d3. One impl fake worker confirmed
  the old file was absent before writing reconstructed code. One checked
  archive commit3d8457abd820313a05063ecdc60cbe48f6cbe38d; real ALP-001 Go
  report passed. SPEC/interface/evals identical; final tree clean, no claims.
- Actual explain P10 --json from outside the project returned0 with the
  non-blocking cause/action text. Schema1, expected data and empty stderr
  throughout. Fixture baselineb4c70e002f7ad72262d2afe76280a204c8e442a2;
  start claim a7ad3cc31e277f31c567371987b7205e3af837ec.

Ordinary fixes (existing implementation, no new main change):

- A genuinely empty Value implementation failed the unchanged ALP-001 Go eval.
  With fixes:[ALP-001], nested-directory next --json returned0 in5.33s,
  driving evals/audit/impl and real fail-first/pass reports. Three fake-agent
  calls and three checked commitsb2f322f/6f9b6e6/b522592. SPEC/interface/
  evals stayed identical, implementation corrected, tree clean and claims absent.
- Run20261029T192920-fa3bb9988bc710ab444a2789; baseline
  6cde809d285625797370082b55c601a32e94fb44, final
  b522592efa3ab98f237a905d6148db5ff3d48d21. This proof adds no applied-change
  count and repeats no previously completed ordinary-fixes proof.

Detailed local proof/command/before logs use /tmp/fireweed-regenerate-cli-*
and /tmp/fireweed-fix-cli-*. Both fixtures/state trackers were removed after
PASS. Deterministic fake agents prove kind composition, not real-agent drill,
regeneration sufficiency or self-hosting. Existing0101 date assertions use
authored report inputs, not real-drill claims.

Partial R06/R08/R09/R16/R17/R23/R29 evidence closes the scoped valid regenerate
and ordinary-fixes paths. Four applications since passing0094–0098 review;
last push7e0da27. No pending main change/claim. Remaining management commands,
persistent bounds, design inventory and final gates stay open. Next: minimal
release API dependency, fifth-change checkpoint, then its CLI caller.
