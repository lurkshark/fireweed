from pathlib import Path
import json,subprocess,hashlib
R=Path('/Users/joel/Workspace/fireweed');O=Path(__file__).resolve().parent;m=json.loads((O/'promotion.json').read_text());S=m['source'];H=m['sha256'];assert (O/'validation/complete.json').exists() and (O/'cleanup.json').exists() and (O/'preparation.json').exists();assert not subprocess.check_output(['git','-C',str(R),'for-each-ref','--format=%(refname)','refs/fireweed/claims'],text=True).strip();assert not list((R/'changes').glob('*.md'));assert hashlib.sha256((R/'.fireweed/bin/fireweed').read_bytes()).hexdigest()==H
G='''# Prepared workflow retry — stopped for separate approval

The two owner-selected normal audited changes are complete:
[0183](../../changes/archive/0183-reserved-ids-review-fail-first.md) clarifies
FLOW-007 reserved IDs and FLOW-023 review fail-first exclusion as already_satisfied;
[0184](../../changes/archive/0184-verification-stop-diagnostics.md) improves RUN-037
failure evidence while preserving its two-second budget, four-second hook sleep
and all assertions. No runtime/interface/hook/configuration correction was made.
FLOW-062 and its boundary cause-text test are unchanged. RUN-037's native timing
cause remains unresolved; diagnostics do not establish a timing fix.

Both changes passed normal phase verification and independent restricted audit,
were applied/archived by the old pinned controller and released all claims.
Identified clean source passed normal/race Go suites, vet, check, hook doctor,
explicit history/CI delta with valid prefixes reused, and compiled CLI allocation/
review-enforcement fixtures. The retiring pin is preserved; the verified new pin
is recorded in PROGRESS. The
[committed preparation evidence](../../.fireweed/drills/workflow-contract-diagnostics-handoff.md)
records exact source, commands/results, phase reports, warnings/cost limits,
promotion, expected hashes, preparation-only setup and cleanup.

**No native drill was launched in this sequence. Stop for separate owner approval
of one workflow retry on the identified source and pin.** No active run, queue or
shared claim remains. Owned validation/preparation clones are removed. All prior
native results/reports/proof/attempt refs, original 79cf33a and prior corrected
ccd3d6b objects, bounds/counters, recovery executables and unrelated worktree remain
preserved. The existing deadline is **2026-11-10 22:51:05 UTC**, with no extension.
The preceding native workflow failure remains current; WF4/full completion stay open.

After explicit approval, follow
[the prepared inspector README](workflow-drill/README.md), using a fresh outside
root and the exact candidate/expected pin in control.py. Recheck live queue,
claims, deadline, bounds and pin; run context for the applicable step. Prepare
fresh caller setup, then launch only the single approved invocation. Review actual
retained inputs and all live isolation gates before releasing the first worker;
review all commands/scripts/incremental patches and full native feedback before
any remaining configured release. Repeat preflight and fresh-output identity
checks. Exact worker command/model/effort, worker SYSTEM, input exclusions and all
isolation/output/stop gates remain unchanged. No import inventory, hints,
instrumentation, command interception or expanded inputs are authorized.

Keep received streams durable outside reconstruction, forwarding identical bytes.
On rejection/interruption use the prepared worker-first bounded stop/drain path,
give the native controller its bounded publication opportunity, and retain the
whole-run emergency stop. Earlier fake rehearsals prove wrapper behavior only;
this preparation establishes no new native reporting behavior. Preserve whatever
terminal evidence/ref is actually published, record absence honestly, clean only
owned scratch and commit protected evidence idle. Stop afterward. A failure/gap
authorizes neither another correction nor another drill. No reconstruction import,
further pin update, later work or push is authorized now.
''';(R/'docs/rebuild/GOAL.md').write_text(G)
p=R/'docs/rebuild/PROGRESS.md';s=p.read_text();marker='The owner-requested sequence is complete and stopped.';i=s.index(marker);top='''# Rebuild handoff

The owner-selected contract clarification and diagnostic-only changes are complete,
independently audited and archived:
[0183](../../changes/archive/0183-reserved-ids-review-fail-first.md) makes FLOW-007
ID reservation and FLOW-023 review fail-first exclusion explicit as already_satisfied;
[0184](../../changes/archive/0184-verification-stop-diagnostics.md) improves RUN-037
failure evidence without changing its two-second budget, four-second hook sleep
or retention assertions. FLOW-062 remains unchanged. No runtime/interface/hook/
configuration correction, reconstructed-source import or worker-input expansion.

The old pin governed both normal lifecycles and stayed unchanged until all runs/
claims exited and clean-source validation passed. The
[verified pin and stopped preparation handoff](../../.fireweed/drills/workflow-contract-diagnostics-handoff.md)
records normal/race Go suites, vet, check, hook doctor, explicit history/CI delta
with valid prefixes reused, compiled CLI allocation/review regressions, actual
cost/warning limits, idle promotion, preparation-only setup and cleanup. Original
79cf33a, prior corrected ccd3d6b, all old reports/archives/proof/attempt refs,
counters/bounds, recoveries and unrelated worktree remain preserved.

**Prepared only: no native workflow invocation occurred in this sequence.** No
active run/shared claim or queued change remains; owned validation/preparation
clones are removed. The current pin is below and outside candidate metadata names
the validated source. Stop for separate owner approval of one workflow retry,
using a fresh outside root and fresh live review/isolation gates. No extra
correction, later work or push. Follow GOAL for the restart procedure.

The [remaining-failure investigation](../../.fireweed/drills/workflow-failure-investigation-20261108/findings.md)
confirmed workflow reconstruction defects. RUN-037's native timing cause remains
unresolved; improved diagnostics are not a timing fix. The preceding native
workflow failure remains current and WF4/full completion stay open.

''';s=top+s[i:];old='0.1.0, source `ccd3d6b7732d613e80a027ba7969b65adc58bc0e`, SHA-256\n`20278504a8ea89fcd5f23d1de9a6ec849389fd55707054992c59967eec65e2b1`.';assert old in s;s=s.replace(old,f'0.1.0, source `{S}`, SHA-256\n`{H}`.');old='Its owner-authorized verified handoff is recorded above; it stayed unchanged\nthroughout the native retry.';assert old in s;s=s.replace(old,'Its owner-authorized verified idle handoff is recorded above; no native retry\nhas used this new pin.');s=s.replace('Only the newly approved two-change\nvalidation/preparation sequence above supersedes that stopped state. The next\nnative invocation remains a separate approval; no push is authorized.','The subsequently approved two-change\nvalidation/preparation sequence is now complete and stopped as recorded above.\nThe next native invocation remains a separate approval; no push is authorized.');s=s.replace('Recovery executables remain at `.fireweed/bin/fireweed-12c571c`,','Recovery executables remain at `.fireweed/bin/fireweed-ccd3d6b`,\n`.fireweed/bin/fireweed-12c571c`,');p.write_text(s)
p=R/'docs/rebuild/ACCEPTANCE.md';s=p.read_text();old='''It confirms reconstruction defects for the workflow failures; RUN-037 remains
unreproduced in bounded focused/full-context replay. Proposed follow-ups need
owner selection before normal audited authoring. The native failure and WF4 stop
remain unchanged.

The owner now selected two follow-ups for normal audited execution: FLOW-007/
FLOW-023 already_satisfied clarification and RUN-037 diagnostic-only eval_fix.
Keep FLOW-062 unchanged. After both archive, validate clean source, perform an
idle verified controller update and prepare the stopped retry handoff. The budget
is exactly these two changes plus validation/preparation. Native retry remains a
separate approval; no source import, expanded worker inputs, raised/reset bound
or counter, extra correction or push. Follow current GOAL/PROGRESS; previous
native results and all WF4/full-completion gates remain intact.''';new='''It confirms reconstruction defects for the workflow failures; RUN-037 remains
unreproduced in bounded focused/full-context replay. The native failure and WF4
stop remain unchanged.

The two subsequently owner-selected follow-ups are complete through normal
independent audit and archive: [0183](../../changes/archive/0183-reserved-ids-review-fail-first.md)
clarifies FLOW-007/FLOW-023 as already_satisfied and
[0184](../../changes/archive/0184-verification-stop-diagnostics.md) is a diagnostic-only
RUN-037 eval_fix. FLOW-062, runtime/interfaces/hooks/configuration and existing
bounds remain unchanged. The
[clean-source validation, idle pin update and stopped preparation](../../.fireweed/drills/workflow-contract-diagnostics-handoff.md)
are complete with no native invocation or active run/claim. RUN-037's native
timing cause remains unresolved. This budget is exhausted; stop for separate
owner approval of one retry on the identified source/pin. No source import,
expanded worker inputs, raised/reset bound or counter, extra correction, later
work or push. Follow GOAL/PROGRESS; all prior native results and WF4/full-completion
gates remain intact.''';assert old in s;s=s.replace(old,new);p.write_text(s)
s=(O/'README-prepared.md').read_text();s=s.replace('Only the worker SYSTEM wording changes: no local builds/tests/module setup;','The worker SYSTEM policy is unchanged: no local builds/tests/module setup;');(R/'docs/rebuild/workflow-drill/README.md').write_text(s)
print('Wrote stopped handoff planning and README, no product files.',flush=True)
