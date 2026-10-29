# Rebuild progress

## Current handoff — 2026-10-29

Manual lifecycle, Go/none/local init, doctor/guidance, bounded agent invocation,
checked NEW next/run, verification retry and clean saved blocking work. NEW run
now continues independent work while blocked dependencies wait, retaining exit 1
and the blocking cause/ref. The full design and final ACCEPTANCE remain open.

**Applied since review: 0; separate 0089–0093 checkpoint PASS.** The separate Sol/high 0084–0088 checkpoint
passed with build/vet/race/static/history at clean 940ef0e; its explicit new
history extension passed 37 commits with zero diagnostics. The owner directly
approved the reviewed push and the attachment's standing subagent-review/push
process in this chat. The push advanced origin/main from 8aff1e4 to 82ef71d.
See [the checkpoint](reports/checkpoint-0084-0088.md); its earlier rejected
invocations never ran. Normal enforcement remains intact.

## Next slice — usage-limit waits and clean budget stops

0089-independent-continuation applied at bf27309: RUN-010/011/016/025/026,
423a9f7/5edbc1d/bf27309, five genuine failures/212 passes → 217 full phase
passes, isolated assertion audit PASS. The pre-audit all-blocked assertion
migration is recorded without hiding the original conflict. Public interfaces
and frozen assertions stayed unchanged during impl. No governing pending
change or claim remains. See [the report](reports/independent-run-0089.md).

The actual NEW Go run proof passed in 14.46s: Alpha failed both impl attempts,
saved wrong-2 and untracked work, restored its phase base and blocked; dependent
Beta stayed unstarted; Gamma applied/archived. Seven fake-agent calls and six
branch commits, exit1/E01/ref, empty stderr and clean unclaimed checkout were
inspected. Alpha's spec remained unapplied. This closes independent continuation
for this bounded scenario, not the whole recovery milestone.

0090-saved-attempt-resume archived at da16b68 with all 297 cases passing.
FLOW-074/075 were corrected and scoped re-audited in native round 1, retaining
run/IDs without exception. See [the API report](reports/saved-resume-api-0090.md).
0091-saved-attempt-command then archived at 523f908 with all 178 cases passing
and isolated APP-009/087/088 audit PASS; ordinary start assertions were reused.

The actual NEW resume proof passed in 2.86s. From the existing author unblock,
start --from-attempt 1 restored Alpha's wrong-2 and saved work without new
IDs/run/phase or a start commit, returned compact schema1 context, then accepted
manual correction and archived through NEW done. Beta's queued document and
Gamma's archive stayed identical. Ref/OID remained unchanged, no new agents ran,
and final tree/claims were clean. The harness's removed-empty-cwd issue was
repaired before any NEW call; the unblock commit was retained. The retained
fixture and its tracker are now removed. See [the CLI proof](reports/saved-resume-cli-0091.md).

0092-agent-commit-recovery archived at 96211bd with 231 full phase passes
and isolated RUN-027/028 audit PASS. Actual NEW next passed in 6.12s: four
fixture-agent commits were soft-reset away, staged/untracked work survived,
three checked phases archived Alpha, three phase-tagged P05 warnings appeared,
and tree/claims were clean. Fixture removed. See
[the proof](reports/agent-commits-0092.md). No governing pending change or claim.

0093-usage-limit-control did not apply. Its bounded implementation check had
67 passes and one frozen assertion conflict: the ordinary applied-result helper
required three invocations but the usage retry correctly retained four. The
owner approved scoped discard/native abandonment, a fixture-only RUN-016
`eval_fix`, then re-authoring with fresh implementation, including consecutive
runner recovery changes. Native abandonment b08c4d0 released the claim and
removed the uncommitted implementation. Only requirements/evals were retained;
old implementation and logs are not implementation inputs. Prior two-round
assertion audit reached PASS; abandonment addresses the frozen helper conflict,
not an outstanding audit finding or a new round-limit exception.

The approved fixture-only 0093-applied-result-count is now archived at 5b0e642:
RUN-016, b8a9260/bd4d862/5b0e642, 231 full phase passes both times, 53 own
passes, isolated assertion audit PASS. Only the exact-count helper refactor
changed; ordinary three-step assertions/callers are preserved. See
[the recovery report](reports/applied-result-count-0093.md).

The separate Sol/high 0089–0093 quality/goal review passed at clean 06a94e9.
Build/vet/race/static and a 40-commit explicit history extension passed with
zero history diagnostics; unchanged warning explanations remain. See
[the checkpoint](reports/checkpoint-0089-0093.md). The authorized push advanced
origin/main from 82ef71d to 28ced38; continuation is authorized. No new final
acceptance, CI/doctor, drill or self-hosting claim follows.

Next: re-author bounded usage handling using retained contracts/evals, recreate
implementation fresh, wire app exit 3 promptly and close the actual NEW wait/
success/stop/resume scenario. Explicit approval covers consecutive runner
recovery changes. No pending change or claim; no phase-write exception or drill.
Time --for, remaining kinds/bounds, the complete design and final ACCEPTANCE
stay open.

## Authority and deferred work

- After five applied changes since review, obtain a separate subagent quality/
  goal-adherence checkpoint. PASS authorizes pushing to origin and continuing;
  any finding requires a human stop. Current direct approval explicitly covers
  this process. Blocked-change stops, third-consecutive module/behavior
  checkpoints, phase audit bounds and shared-base P07 remain enforced. Pushes
  alone do not reset review counts; no ref/bound/count evasion is permitted.
- Owner removed numeric spec-count caps for every change kind. Keep coherent
  single-module changes, atomic specs and normal checks/audits. Completed global
  spec splits and D006 existing-path migration need no rewrite. Add diagnostic
  explanations only for justified emitted paths.
- Do not reorganize workflow stub.go without a demonstrated feature need.
- The only authorized contracts drill has unresolved independence limitations,
  including dependent-internal reads. Strict regeneration acceptance remains
  unverified; no gap fix, repeat or other-module drill is authorized. See
  [the drill report](reports/contracts-drill.md).
- Prior phase/audit deviations remain historical, not retroactively validated.
  Both approved 0085 phase-write recovery exceptions are consumed. No new
  exception exists. See [0085](reports/saved-attempt-blocking-0085.md),
  [the progress assessment](reports/progress-review-after-0087.md), and
  [HISTORY](HISTORY.md) for archived counts, recoveries and earlier authority.

## Fixed execution context

Controller: /Users/joel/Workspace/fireweed/.fireweed/bin/fireweed-bootstrap,
0.1.0 source 7f925190878e38fc38f30d16bb1cf47aca098cf7, SHA-256
f1516b5686435c23b54b71c5ac554aba6859b5454f072e0c5eb031dfd6f0fd4f; keep
unchanged. Original history baseline be35552b706aa3000ff8f11306d129ff1d701b85;
use explicit nonempty ranges and extend the recorded valid portions.
GOCACHE=/tmp/fireweed-rebuild-cache.

NEW: /Users/joel/Workspace/fireweed/bin/fireweed, built at clean
06a94e95c3bc4a82d74c7799031cb738eff5b7d2, Go1.27.2 darwin/arm64, CGO disabled,
vcs.modified=false. Reuse valid earlier lifecycle/setup/CLI/checkpoint evidence;
no fresh final CI/doctor/fresh-checkout/drill/self-hosting claim follows.

Real workers and separate assertion auditors use codex exec --ephemeral
--ignore-user-config -m gpt-6-sol -c 'model_reasoning_effort="high"'
-s workspace-write {prompt}. Coordinate native start/context/done; workers
never commit. Auditors run in isolated snapshots with change, projected specs,
evals and permitted public interfaces/rules, never implementation, SYSTEM,
history or outside files; inspect complete command logs. Record coordinator
native advancement honestly as manual and retain actual CLI logs. Do not use
native automatic audit dispatch after the documented input/phase deviations.
