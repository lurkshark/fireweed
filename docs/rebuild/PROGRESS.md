# Rebuild progress

## Current status — 2026-10-27

Human review 5 approved changes 0020–0024 and the checkpoint report. The reviewed main was pushed at the owner’s explicit request. **Applied since review: 2.** Milestone 1 passed through the NEW binary;
Milestone 2 waits for the spec splits, diagnostic migration and one contracts
drill. Consecutive workflow/checks changes are authorized for the spec split.
Stop after five applied changes, on any block, or on a drill failure that calls
the spec approach into question.

Controller: `.fireweed/bin/fireweed-bootstrap` 0.1.0, source `7f925190878e38fc38f30d16bb1cf47aca098cf7`, pin SHA-256 `f1516b5686435c23b54b71c5ac554aba6859b5454f072e0c5eb031dfd6f0fd4f`; unchanged.
Stage-2 history baseline: `be35552b706aa3000ff8f11306d129ff1d701b85`; always check an explicit nonempty baseline..HEAD range. Use `GOCACHE=/tmp/fireweed-rebuild-cache`.
NEW binary: `bin/fireweed`, clean source `c64765f2bb21c4449cdf6d5b1e1778a71dd1eb4f` (CGO disabled); [committed manual-lifecycle report](reports/manual-lifecycle.md). Unbuilt commands remain explicit errors; the empty runner is expected.

Workers and separate auditors use `codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c 'model_reasoning_effort="high"' -s workspace-write {prompt}`. Interactive work through 0014's initial evals was followed by coordinator `done`, hence legacy `agent=manual` lines. Configured controller execution began at 0014's audit; its provenance records the actual command/model. Continue configured execution. Ignored runtime logs are auxiliary, not reviewable evidence; archives, permanent evals and committed reports carry the evidence.

## Applied changes

- 0001 standalone-information — AST-001–003; phases a4ceb17/826b9e9/2248b8e; 8 fail-first → 8 pass.
- 0002 standalone-cli — APP-001–003; phases 1298acf/7de7d30/db77de7; 30 fail-first → 30 pass; auditor findings were discarded by folded quotes, corrected in 0003.
- 0003 json-error-assertions — APP-003 eval_fix; eb4704d/e184bb6/8796d37; 26 app cases pass, independent audit passed.
- 0004 lifecycle-primitives — CON-001–005; 11d0180/6e341a4/8e37dd7; 18 failures → 27 own cases pass, dependents pass.
- 0005 fixture-authoring — PRJ-001–003; 1bcba65/5c647a2/bfee6b8/06f17c6/91f036c; one findings round, 24 failures → 25 own cases pass.
- 0006 project-commands — APP-004–005; da5dba8/98e5852/5f09706; 12 failures → 51 app cases pass.
- 0007 change-handles — CON-006; a22fcd4/a194629/7580897; 1 failure → 34 own cases pass, dependents pass.
- 0008 queue-authoring — PRJ-004–007; 6e3ca5f/776b4d0/2ef37df; 24 failures → 53 own cases pass, dependents pass.
- 0009 authoring-commands — APP-006–008; 5720dce/b2aaa14/de65319; 27 failures → 78 app cases pass.
- 0010 manual-lifecycle-git — GIT-001–005; evals 3e82adc/93ba9a6/5d00cac, audits 98a799f/f2ddd0f/5911fab, impl 92f999a; two rounds, 7 failures → 8 own cases and 2 public probes pass.
- 0011 lifecycle-checks — CHK-001–005; 69933f7/eff6823/3c5b40b; 25 failures → 28 own cases and 4 public probes pass, dependents pass.
- 0012 manual-lifecycle — FLOW-001–005; evals c5c02de/ca11056/aec7154, audits 0b5aeef/25e4775/8627014, impl 563d2ee; two rounds, 28 failures → 32 own cases and 3 public probes pass.
- 0013 lifecycle-cli — APP-004/009–012; 678f7bb/7db8c47/a9d195c/804cb84/89b7028; one round, 8 failures → 86 app cases, held-coverage probes and real Go lifecycle pass.
- 0014 go-setup-bundle — AST-004–007; 491bffd/7789d62/d24f799/715a1be/03c2375; one round, 6 then 8 failures → 20 own cases and 199 dependent cases pass; configured-agent provenance from audit onward.

- 0015 start-state-specs — FLOW-006–009; ec03162/5223bcb/891d152; already_satisfied, 36 own cases and dependents pass; eval retargeting correction follows in 0016.

- 0016 retarget-start-evals — eval_fix FLOW-001/006–009; 129178c/0b7ea8b/5337b77; restored 32 own cases with every assertion/ID retained, dependents pass.

- 0017 start-precondition-specs — FLOW-010–013; 438072c/b35410d/6b2af96; already_satisfied, 32 own cases and dependents pass; implementation unchanged.

- 0018 start-claim-specs — FLOW-014–017; 9cfb547/31df564/9669555; already_satisfied, 32 own cases and dependents pass; implementation unchanged.

- 0019 narrow-start-claim-spec — FLOW-001/018–019; 0bc1106/dc436e7/bde237f; already_satisfied, 32 own cases and dependents pass; FLOW-001 narrowed, implementation unchanged.

[Checkpoint verification and remaining work](reports/review3-checkpoint.md): race/vet/controller CI passed; 261 reported cases, doctor recovery and history passed. No pending change or claim. Human review 4 cleared this checkpoint.

- 0020 verify-evidence-specs — FLOW-020–023; evals 2d7912c/15139aa, audits b43448f/7bbf4eb, impl 42dae8e; one findings round, 32 own cases and dependents pass; implementation unchanged.

- 0021 verify-pass-specs — FLOW-024–027; 3067213/0521c85/8d786c6; 32 own cases and dependents pass, implementation unchanged; resumed saved attempt 1 after human-confirmed Sol access recovery.

- 0022 narrow-verify-zone-spec — FLOW-002/028–031; 534fef4/ea82977/17441fd; 32 own cases and dependents pass; FLOW-002 narrowed, implementation unchanged.

- 0023 evals-transition-specs — FLOW-032–035; 5ebcac7/2481621/5fbaf95; 32 own cases and dependents pass; existing tests retargeted, implementation unchanged.

- 0024 audit-state-specs — FLOW-036–040; 1b1c8f4/8aa115e/c64765f; 32 own cases and dependents pass; existing tests retargeted, implementation unchanged.

[Review 4 checkpoint](reports/review4-checkpoint.md): race/vet/controller CI passed, including 261 cases, doctor recovery and history; no pending change or claim. Human review 5 cleared this checkpoint; P07 is clear.

- 0025 narrow-done-verification-spec — FLOW-003/041–043; fe63ecf/44f5aa4/9042469; 32 own cases and dependents pass; FLOW-003 narrowed, implementation unchanged.

- 0026 zone-prohibition-specs — CHK-006–010; evals 36fffae/69c2898, audits a643dcb/7bfbe66, impl 7b5634c; one findings round, 28 own cases and dependents pass; implementation unchanged.

## Next

1. Split permanent bundled behaviors through authored changes. Use already_satisfied and retarget existing evals; never hand-edit Specs/Decisions. FLOW-001/FLOW-002/FLOW-003 are split; split CHK-004 and the other bundled items identified in the checkpoint report. Add atomic specs before narrowing an old item when one change cannot hold the full split.
2. Apply D006: remove wrong label assertions through eval_fix, migrate implementation in normal changes, and keep explain accurate. Preserve all behavior assertions and coverage.
3. Before Milestone 2 touches workflow implementation, split lifecycle/stub.go by responsibility as maintenance only. No other cleanup/hardening.
4. After items 1 and 2, run exactly one real controller drill on contracts. Commit its report and list gaps here; author no gap changes and drill no other module.
5. Milestone 2 then centers projected blocking and fake-agent next/run through NEW. No NEW unattended real-agent execution before that fake-agent run completes. Full acceptance remains unverified.
