# Eval fix 0099 — actual NEW proof PASS

FLOW-078/079 and modified FLOW-036/023 applied at
23ad9de2a67c22be063ac860bd48175c74c85ceb. FLOW-023 qualifies the preserved
ordinary spec/fix fail-first obligation and is already_satisfied. Native phases
35ce303/3e30063d/23ad9de record361 cases: six failure events/355 passes, then
all361 passing in impl. Own pre-freeze80.805s check had75 leaf cases:71 passes/
four genuine failures (83 pass/six fail events). They expose acceptance of
failing evals, forbidden interface/impl edits, and incomplete audit checked IDs.
The corrected-eval lifecycle and missing-coverage rejection already passed.

Isolated Sol/high assertion audit passed FLOW-023/036/078/079 without findings.
Complete command-log inspection confirms only permitted snapshot reads, no
SYSTEM/implementation/history/outside content, execution or writes; snapshot
removed. Native provenance is manual; real fixed-profile CLI logs remain under
20261029T175930-a47eee99acd8dcc2416ccf47.

Only one new public workflow eval file and one private lifecycle implementation
file changed. Impl restricts eval-fix zones before hooks and after verification,
requires target cases pass and listed-ID coverage, and includes eval_fix IDs
in audit findings/checked-list handling. Ordinary fail-first, approved snapshots,
claim/phase/ref validation, dependents/boundaries and audit bounds remain.
The worker's single own full check passed89 events/75 leaf cases, zero failures
or skips, in77.778s. No implementation correction, API/diagnostic allocation,
bound change, phase exception, reorganization or drill. Evals/interfaces frozen.

NEW built from clean23ad9de, Go1.27.2 darwin/arm64, CGO disabled,
vcs.modified=false. Actual Go eval-fix fixture passed:

- The original ALP-001 test genuinely failed because it demanded a value absent
  from the correct spec. Implementation already satisfied that spec.
- Nested-directory next --json returned0 in4.74s, correcting only the existing
  eval, accepting passing evals, separately auditing ALP-001 and completing a
  no-op impl phase. Schema1 data applied, diagnostics[], stderr empty.
- Three fake-agent calls and three checked commits:
  a5538d6/d4630e5/f1bd2d1. Both evals/impl reports contain passing ALP-001.
  SPEC, interface and implementation stayed identical; final tree clean and
  claims absent. Run20261029T181757-55c047086003af29e55cc280.

Fixture baseline202719b9dee8c1b188295272dabc65cb579aa07f; final
f1bd2d1f3c882e9d46f3ca79335ea469196218dd. Details:
/tmp/fireweed-eval-fix-cli-result.json, command.json and before.jsonl.
Fixture/state tracker removed after PASS. Deterministic fake agents prove
orchestration, not real-agent regeneration/self-hosting.

Partial R08/R09/R17/R23 evidence closes the valid eval-fix CLI path. Embedded
worker guidance still needs scoped eval-fix instructions: impl currently directs
workers toward disposable code despite the new change-file-only rule. That is
the next asset slice. One application since passing0094–0098 checkpoint;
last push7e0da27. Remaining kinds/bounds/inventory and full acceptance stay open.
