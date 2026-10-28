# Rebuild progress

## Current status — 2026-10-28

D006's planned existing-path slice is complete through 0068. The single contracts
drill reports 34 own/339 dependent cases passing and retains its ref with scratch
cleanup complete. It reports an error-prefix contract gap and its agent read
dependent internals, so strict independence acceptance is not established; see
[the reviewed report](reports/contracts-drill.md). No gap fix or repeat drill is
authorized. 0069 now implements recipe-based initialization through NEW.
**Applied since review: 0; one further bounded correction approved.** The owner-approved 0070
correction and scoped re-review passed and were pushed at e083c5f. The next
change, 0071 doctor-probes, originally blocked in evals after two audit findings
rounds with CHK-065's default-timeout assertion missing. Its first auditor exceeded
the input boundary and temporarily
edited implementation; the controller rejected that edit and it was restored.
See [the block report](reports/doctor-0071-block.md). The approved saved-attempt
recovery added the real default-timeout eval in d78d618 (255 phase cases: 29
runtime failures, 226 passes). A fresh restricted audit reviewed all doctor IDs
without implementation access and no longer reports CHK-065, but reported eight
other assertion gaps. The controller blocked 0071 in audit at a81b723, retaining
audit_rounds 2 and both saved refs. The owner now approves one bounded correction
of the eight outstanding IDs, preserving the timeout test and prior passes, with
restricted re-review of only those IDs. Use the native saved eval attempt; if it
cannot restore evals without changing bounds, stop before editing frozen tests.
No doctor implementation was applied; the full goal remains incomplete.

The owner's goal update authorizes a subagent review at every five applied
changes: push to origin and continue when the review finds the work on track, and
stop for human review with details when it reports an issue. Record each review. Separate
phase audits, blocked-change stops and P07 remain enforced; this policy does not
alter bounds. The owner explicitly authorizes each push only after its
passing five-change subagent review. Milestone 1 passed;
Milestone 2 is next; the diagnostic slice and single requested contracts drill are recorded.

Owner approval on 2026-10-28 removes the numeric spec-count caps for all change
kinds, superseding both the five-spec default and the 30-item exceptions. Batch
related work within one module, keeping changes coherent and reviewable and
specs atomic. Spec-only splits preserve all obligations/assertions and behavior;
D006 remains limited to existing error paths, eval corrections and explanations
for emitted codes, preserving behavioral assertions and interfaces without
unsupported features or speculative errors. Normal verification, independent
audits, audit bounds, timeouts and recurring checkpoint reviews remain in force.

Controller: `.fireweed/bin/fireweed-bootstrap` 0.1.0, source `7f925190878e38fc38f30d16bb1cf47aca098cf7`, pin SHA-256 `f1516b5686435c23b54b71c5ac554aba6859b5454f072e0c5eb031dfd6f0fd4f`; unchanged.
Stage-2 history baseline: `be35552b706aa3000ff8f11306d129ff1d701b85`; always check an explicit nonempty baseline..HEAD range. Use `GOCACHE=/tmp/fireweed-rebuild-cache`.
NEW binary: `bin/fireweed`, clean source `4b36a84a8afb20806f4738528b3a7627810a8603` (CGO disabled); [committed manual-lifecycle report](reports/manual-lifecycle.md). Unbuilt commands remain explicit errors; the empty runner is expected.

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
- 0020 verify-evidence-specs — FLOW-020–023; evals 2d7912c/15139aa, audits b43448f/7bbf4eb, impl 42dae8e; one findings round, 32 own cases and dependents pass; implementation unchanged.
- 0021 verify-pass-specs — FLOW-024–027; 3067213/0521c85/8d786c6; 32 own cases and dependents pass, implementation unchanged; resumed saved attempt 1 after human-confirmed Sol access recovery.
- 0022 narrow-verify-zone-spec — FLOW-002/028–031; 534fef4/ea82977/17441fd; 32 own cases and dependents pass; FLOW-002 narrowed, implementation unchanged.
- 0023 evals-transition-specs — FLOW-032–035; 5ebcac7/2481621/5fbaf95; 32 own cases and dependents pass; existing tests retargeted, implementation unchanged.
- 0024 audit-state-specs — FLOW-036–040; 1b1c8f4/8aa115e/c64765f; 32 own cases and dependents pass; existing tests retargeted, implementation unchanged.
- 0025 narrow-done-verification-spec — FLOW-003/041–043; fe63ecf/44f5aa4/9042469; 32 own cases and dependents pass; FLOW-003 narrowed, implementation unchanged.
- 0026 zone-prohibition-specs — CHK-006–010; evals 36fffae/69c2898, audits a643dcb/7bfbe66, impl 7b5634c; one findings round, 28 own cases and dependents pass; implementation unchanged.
- 0027 evals-zone-permission-specs — CHK-011–015; 82bf116/09b7274/d384192; 28 own cases and dependents pass; existing eval retargeted, implementation unchanged.
- 0028 impl-zone-permission-specs — CHK-016–020; e4bb1be/1a73eb6/5c54ca1; 28 own cases and dependents pass; existing eval retargeted, implementation unchanged.
- 0029 narrow-audit-zone-spec — CHK-004/021–024; cdfb41c/e38d19b/db03f62; 28 own cases and dependents pass; CHK-004 narrowed, implementation unchanged.

[Review 5 checkpoint](reports/review5-checkpoint.md): race/vet/controller CI passed, including 261 cases, doctor recovery and history; no pending change or claim. Approved in human review 6; the next slice starts with PRJ-007.

- 0030 check-document-specs — PRJ-008–012; evals f8d1bc2/05825f0, audits 2909f64/4be3541, impl fcc385f; one findings round, 53 own cases and 199 total phase cases pass; implementation unchanged.

- 0031 narrow-check-diagnostics-spec — PRJ-007/013–016; evals 6960718/4441d55, audits 04eb482/1e91828, impl ac3779b; one findings round, 53 own cases and 199 total phase cases pass; PRJ-007 narrowed and blanket C01 removed, implementation unchanged.

[Review 6 checkpoint](reports/review6-checkpoint.md): race/vet/controller CI passed,
including 261 cases, doctor and history; no pending change, claim or scratch
worktree. Approved in human review 7, including consecutive project splits.

- 0032 narrow-load-spec — PRJ-001/017–018; 58d1946/b7d1061/f36a557; 53 own cases and 199 total phase cases pass; no findings, implementation unchanged.

- 0033 narrow-graph-spec — PRJ-002/019–022; d170452/8f084e2/5d1bb0f; 53 own cases and 199 total phase cases pass; no findings, implementation unchanged.

- 0034 narrow-new-change-spec — PRJ-003/023–026; 305d5a3/ea7cd72/2519e8a; 53 own cases and 199 total phase cases pass; no findings, implementation unchanged.

[Review 7 checkpoint](reports/review7-checkpoint.md): three applications; P07
was awaiting shared-base synchronization; human review 8 now approves the
checkpoint and authorizes pushing reviewed main. No pending
change, claim or scratch worktree. Prior race/vet/CI/doctor evidence is reused
because this slice changed only spec wording and eval labels/table metadata.

- 0035 queue-document-specs — PRJ-027–031; c22892f/ac152fa/6187c97; 53 own cases and 199 total phase cases pass; no findings, implementation unchanged.

- 0036 narrow-change-resolution-spec — PRJ-004; 49b9f0e/9bdc62d/0fc27f1; 53 own cases and 199 total phase cases pass; no findings, implementation unchanged.

- 0037 projection-rule-specs — PRJ-032–036; 3002d69/1f13327/136ca0c; 53 own cases and 199 total phase cases pass; no findings, implementation unchanged.

- 0038 narrow-projection-edits-spec — PRJ-005; evals dc8847c/a6ecd35, audits ec740ab/e705509, impl 0850274; one findings round, 53 own cases and 199 total phase cases pass; full Modified-body assertion strengthened, implementation unchanged.

- 0039 narrow-stamp-spec — PRJ-006/037–039; f33d236/7166bc5/430e56c; 53 own cases and 199 total phase cases pass; no findings, implementation unchanged.

[0035–0039 checkpoint](reports/checkpoint-0035-0039.md): five applied project
splits. Subagent checkpoint_0035_0039 returned PASS with no issues. The baseline
history check passed over 232 commits; the owner conditionally authorizes the
push and continuation. The applied count restarts after this passing review.

- 0040 narrow-coverage-spec — CHK-002/025–026; 9d54aea/ef8495d/40bfafb; 28 own cases and 146 total phase cases pass; no findings, implementation unchanged.

- 0041 narrow-junit-spec — CHK-001/027–030; fb61726/73f0bc7/dd588b3; 28 own cases and 146 total phase cases pass; no findings, implementation unchanged.

- 0042 eval-hook-data-specs — CHK-031–035; 9ae18d0/53dd0be/5c585d6; 28 own cases and 146 total phase cases pass; no findings, implementation unchanged.

- 0043 narrow-eval-hook-spec — CHK-003/036–038; be6ed03/5536992/07181b1; 28 own cases and 146 total phase cases pass; no findings, implementation unchanged.

- 0044 boundary-suppression-specs — CHK-039–043; 8492280/7cbffac/06699f9; 28 own cases and 146 total phase cases pass; no findings, implementation unchanged.

[0040–0044 checkpoint](reports/checkpoint-0040-0044.md): five applied checks
splits. Subagent checkpoint_0040_0044 returned PASS with no issues. The explicit
baseline history check passed over 258 commits, authorizing the conditional
push and continuation. The applied count restarts after this passing review.

- 0045 narrow-boundary-hook-spec — CHK-005/044–045; 8d23fd6/cb9e161/ab79c6c; 28 own cases and 146 total phase cases pass; no findings, implementation unchanged.

- 0046 application-state-specs — FLOW-044–048; 5a5df43/7ab36b4/e52698f; 32 own cases and 118 total phase cases pass; no findings, existing metadata/Decision/phase-base assertions strengthened, implementation unchanged.

- 0047 narrow-application-spec — FLOW-004/049–050; 80b6992/6598b5a/2a087d7; 32 own cases and 118 total phase cases pass; no findings, archive/SPEC commit assertion strengthened, implementation unchanged.

- 0048 narrow-context-spec — FLOW-005/051–053; evals c14ce8a/a146cbc, audits 63af469/a7969b6, impl 95a404a; one findings round, 32 own cases and 118 total phase cases pass; per-phase provenance assertions strengthened, implementation unchanged.

- 0049 split-contract-primitives — CON-001/002/005/006/007–015; 208fd0a/e2809a5/71ac66e; 13 affected items, 34 own cases and 261 total phase cases pass; no findings, assertions/test data and implementation unchanged.

[0045–0049 checkpoint](reports/checkpoint-0045-0049.md): five applied changes.
Subagent checkpoint_0045_0049 returned PASS with no issues. The explicit
baseline history check passed over 286 commits. The owner conditionally
authorizes the push and continuation; the applied count restarts after review.

- 0050 split-git-primitive-apis — GIT-003/004/006–010; d0179bd/bf6acb6/3519faa; seven affected items, eight own cases and 207 total phase cases pass; no findings, assertions and implementation unchanged.

- 0051 split-asset-contracts — AST-001/002/004–023; c9210cf/0d856fc/40cceb7; 22 affected items, 20 own cases and 219 total phase cases pass; no findings, assertions/fixtures and implementation unchanged.

- 0052 split-authoring-cli-contracts — APP-003–008/013–026; evals ff54a43/4ad8388, audits 56a829c/bcc1400, impl ed94c03; 20 affected items, one findings round, 86 cases pass; existing zone-array assertion strengthened, implementation/diagnostics unchanged.

- 0053 split-lifecycle-cli-contracts — APP-009–012/027–036; evals 863b578/28baed4, audits 62d07a5/ca18597, impl 11cb39a; 14 affected items, one findings round, 86 cases pass; existing omitted-Done/blocked-Verify assertions completed, implementation unchanged.

- 0054 repair-validation-code-assertions — eval_fix PRJ-007; aa756c5/7aa96d6/706ce7e; 53 own cases and 199 total phase cases pass; no findings, stale blanket C01 comparisons replaced with nonempty-code checks while retaining reason assertions; implementation unchanged.

[0050–0054 checkpoint](reports/checkpoint-0050-0054.md): five applied changes.
Subagent checkpoint_0050_0054 returned PASS with no issues. The explicit baseline
history check passed over 316 commits, authorizing push and continuation. The
applied count restarts after this passing review.

- 0055 validation-cause-codes — PRJ-008–012; 0b91a02/8f875e4/405db73; nine reported failures → 54 own/200 total phase cases pass, no findings; existing C04/C05/C06/C07 paths migrated. One existing table gained a separate reporting parent; no scenarios/fixtures added.

- 0056 validation-code-explanations — AST-003; d12c3a2/f8e16b6/c7b4bea; five reported failures → 24 own/224 total phase cases pass, no findings; four emitted-code explanations added, existing explanation and unknown-code checks retained.

- 0057 edit-reference-codes — PRJ-013–016; 7f099ed/80a1760/12cd5bc; seven reported failures → 54 own/200 total phase cases pass, no findings; existing duplicate edit/handle/base/absent-target paths now C05/C09/C08, fixtures and earlier assertions retained.

- 0058 boundary-cause-codes — CHK-039/040/043–045; 69ae618/700c317/0bfd581; three reported failures → 28 own/146 total phase cases pass, no findings; existing violations/hook failures/opt-out/unnecessary warnings now X01–X04, B05 retired here.

- 0059 reference-boundary-explanations — AST-003; 9d109e6/5ec2a2b/c7eabf9; eight reported failures → 30 own/230 total phase cases pass, no findings; C08/C09/X01–X04 explanations added and C05 repeated-edit repair clarified.

[0055–0059 checkpoint](reports/checkpoint-0055-0059.md): five applied changes.
Subagent checkpoint_0055_0059 returned PASS with no issues. Baseline history
passed over 341 commits; build/race/vet and NEW explanation JSON passed for
source c7eabf9. The passing review authorizes push and continuation; the applied
count restarts. No queue or claim remains. Diagnostic migration is incomplete.

- 0060 usage-lookup-codes — APP-003/013/014; ab0f0d6/9b95b13/df1c10d; 18 reported failures → 94 app cases pass, no findings; existing usage/unsupported rejections now U01 and information lookups U02; eight arity table cases added, all prior behavior assertions retained.

- 0061 usage-code-explanations — AST-003; 3e801bf/dd38afc/875aff7; three reported failures → 32 own/240 total phase cases pass, no findings; U01/U02 explanations added with existing causes and unknown-code checks retained.

- 0062 static-diagnostic-causes — PRJ-040–045 plus fixes PRJ-013; 7e2b3f6/d16c471/1e478b8; 17 reported failures → 74 own/228 total phase cases pass, no findings; existing Check graph/syntax/metadata/queue/target paths classified, C01 retained and duplicate-section C05 corrected; 20 focused public cases added with existing fixture helpers.

- 0063 eval-diagnostic-causes — CHK-002/025–028/036–038/046–048; 4db7bf5/4b1e611/f6b9921; 17 reported failures → 37 own/163 total phase cases pass, no findings; existing eval/coverage/result errors now carry C10/C11/E01/P2-01, with static and lookup causes preserved; signatures unchanged.

- 0064 lifecycle-diagnostic-causes — FLOW-009/015/054–062; evals 6479da3/22f0543, audits 23444aa/8f50b18, impl 4066479; one findings round, 12 reported failures → 45 own/139 total phase cases pass; P04-warning/configured-P07/Done-cause assertions completed and scoped re-audit passed; existing lifecycle causes prefixed, behavior/interfaces retained.

[0060–0064 checkpoint](reports/checkpoint-0060-0064.md): five applied changes.
Subagent checkpoint_0060_0064 returned PASS with no issues. Baseline history
passed over 369 commits; current-source build/race/vet and NEW U01/U02 JSON
checks passed. The owner-authorized push and continuation follow this passing
review; the applied count restarts. No queue or claim remains; D006 is incomplete.

- 0065 remaining-diagnostic-causes — APP-018/019/037–040; 2157abf/7697050/7eb9898; 30 reported failures → 121 app cases pass, no findings; legacy CLI labels removed. Post-application NEW smoke exposed show returning U02 for duplicate graph names and duplicate module spec IDs; correct those static causes next while retaining legitimate projected-handle ambiguity.

- 0066 preserve-show-static-causes — APP-038 plus fixes APP-037; afdc02d/edbac55/39ecdb8; three reported failures → 124 app cases pass, no findings; duplicate graph names/spec IDs retain C02/C05 and valid projected local-handle ambiguity retains U02; rebuilt NEW confirms both observed regressions fixed.

- 0067 emitted-cause-explanations — AST-003/024–026; 1360324/5d3e724/72a2860; 22 reported failures → 51 own/331 total phase cases pass, no findings; sixteen explanations added and C07/U02 clarified. All 30 emitted codes resolve through NEW; race/vet pass. Text review found C02 identity causes and P02 missing/unreadable-claim recovery underdescribed; correct those before the drill.

- 0068 complete-graph-claim-explanations — fixes AST-024/026; d6fb6b1/54892c9/79e57f4; five reported failures → 51 own/331 total phase cases pass, no findings; C02 covers module identities and dependencies, P02 covers missing/invalid/foreign claims and appropriate recovery, P01 includes existing authoring preconditions. NEW C02/P02 JSON smoke passes.

- 0069 project-initialization — APP-041–061; evals b3cbc19/734e476, audits 6fe48c5/07fcb6c, impl cdc64ae; one findings round, 17 reported runtime failures → 144 app cases pass; version-range membership and recipe default-protected assertions completed, scoped re-audit passed. Go/none/local init, optional app, managed instructions, preservation, usage/errors and JSON are wired through NEW; its Go init/modules/check fixture passes. Doctor and execution remain pending.

- 0070 recipe-error-explanations — fixes AST-003/024; 39f82e1/95b998c/4b36a84; four reported runtime failures → 51 own/351 total phase cases pass, no findings; seven assertions added to existing cases and only C12/U02 explanation strings extended. Existing causes/guidance, assertions, public interfaces and init behavior are retained. NEW C12/U02 JSON smoke and owner-required checkpoint re-review pass.

[0065–0069 checkpoint](reports/checkpoint-0065-0069.md): five applications since
the passing 0060–0064 review. Subagent checkpoint_0065_0069 initially returned
BLOCK for missing init recipe causes/repairs in explain C12/U02. The owner
approved 0070 and conditional push/continuation; its scoped re-review now returns
PASS. Build/race/vet and NEW Go init/modules/check passed at cdc64ae; the text-only
correction's full phase checks and NEW C12/U02 JSON pass at 4b36a84. History
validation passed over 397 baseline commits plus the seven-commit correction
extension, both with no diagnostics. The count resets after this passing review.
The owner's intervening direct push
approval synchronized committed history through 3897b1a; it did not reset the
five-application review count.

The planned D006 existing-path slice is complete through 0068: no legacy CLI
labels remain, and all 30 emitted codes have NEW JSON explanations. Full phase
checks passed; race/vet at 72a2860 remains valid for 0068's text-only correction.
No new feature, public interface or speculative error was added. Full diagnostic
surface for unimplemented commands is not claimed. The single contracts drill is now recorded below; no gap changes, repeat drill
or other module drills are authorized.

[Contracts drill review](reports/contracts-drill.md): native pass at 7afe9a0,
34 own/339 dependent cases; retained ref and scratch cleanup verified. One
CON-011/012 error-prefix gap is recorded. Agent reads of dependent internal
implementations exceed the retained-input boundary, so strict independence
acceptance remains unverified. No gap fix/rerun is authorized or performed.

## Next

Perform the newly approved bounded correction of [the eight fresh audit findings](reports/doctor-0071-block.md).
Retain prior passes, the real default-timeout test, counters, requirements and
history. Restore an eval step through the controller before editing tests; if
that cannot be done without changing bounds, stop. Re-review only the eight
outstanding IDs from restricted inputs; any unresolved finding stops again.
No replacement change, bound change or push is authorized by this recovery.
The following sequence remains planned after a passing resolution:

0. The owner-approved 0070 explanation correction and checkpoint re-review pass.
   The authorized push reached e083c5f; continue Milestone 2. The original block,
   human approval and resolution remain in the checkpoint report. Next is the
   checks-owned doctor backend, then accurate explanations for its emitted
   diagnostics and app wiring. D007 allocates its reserved C14–C16 slots.
1. The previously identified project/checks/workflow/contracts/gitstore/assets/app splits are complete through 0053. All obligations and existing assertions are retained under focused specs; do not launch another global rewrite. Continue the agreed diagnostic migration.
2. D006's planned existing-error-path migration and explanations are complete through 0068. Preserve assigned causes and keep explain accurate as justified features add real paths. No unsupported-feature scaffolding, speculative errors or unused-code explanations.
3. Workflow file organization is not a Milestone 2 prerequisite. Defer splitting lifecycle/stub.go unless it causes a demonstrated problem during feature work; a later drill/regenerate may organize the implementation afresh. Do not substitute another speculative cleanup task.
4. The single authorized contracts drill is complete and reviewed in reports/contracts-drill.md. Preserve its native report/ref and limitations. Defer the reported gap and any stricter rerun until human agreement; author no gap changes and drill no other module.
5. Milestone 2 now centers init/doctor setup, projected blocking and fake-agent next/run through NEW. No NEW unattended real-agent execution before that fake-agent run completes. Full acceptance remains unverified.
