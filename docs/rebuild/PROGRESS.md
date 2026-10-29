# Rebuild progress

## Current handoff — 2026-10-29

Manual lifecycle, Go/none/local init, doctor/guidance, bounded agent invocation,
checked NEW next/run, verification retry and clean saved blocking work. NEW run
now continues independent work while blocked dependencies wait, retaining exit 1
and the blocking cause/ref. The full design and final ACCEPTANCE remain open.

**Applied since review: 2 (0089–0090).** The separate Sol/high 0084–0088 checkpoint
passed with build/vet/race/static/history at clean 940ef0e; its explicit new
history extension passed 37 commits with zero diagnostics. The owner directly
approved the reviewed push and the attachment's standing subagent-review/push
process in this chat. The push advanced origin/main from 8aff1e4 to 82ef71d.
See [the checkpoint](reports/checkpoint-0084-0088.md); its earlier rejected
invocations never ran. Normal enforcement remains intact.

## Active slice — saved-attempt resume through NEW start

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
StartFromAttempt is ready; FLOW-074/075 were corrected and re-audited in native
round 1, with same run/IDs and no exception. See
[the API report](reports/saved-resume-api-0090.md). Immediately wire
start --from-attempt <n> promptly, then prove actual NEW restores saved Alpha
work without new IDs/run/phase or a start commit, lets it be corrected, and
archives it through the existing lifecycle. Use only the needed workflow
selection/load dependency and app flag; do not expand to status, agent-commit
undo or time/usage controls before closing this CLI scenario.

Reuse the clean existing plain Git fixture
/private/tmp/fireweed-run-cli-ntnv8b5m, tracked by
/tmp/fireweed-independent-run-state.json. Its Alpha run is
20261029T085043-183d7ee49a73545a7e7e4478; saved attempt 1 is
refs/fireweed/attempts/value-alpha/1 at 11d3ba60eb1cbca44072d6129f02641692452ed9.
It has no live process or claim and needs an author unblock commit before
resume. Gamma already applied; Beta still waits. This toy saved work is proof
data, not an implementation oracle. Remove the fixture after the resume proof.
Detailed local prior proof: /tmp/fireweed-independent-run-smoke-result.json.

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
bf27309d379f74a3af85c5841251cc84088308fe, Go1.27.2 darwin/arm64, CGO disabled,
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
