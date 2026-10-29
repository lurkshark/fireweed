# Run time limit 0097 — actual NEW proof PASS

APP-089/090 and modified APP-084 applied at
ab7b3173871fa6552a4aaadb75ef755a7c0c9b60. Native phase commits
79d6ffd/32ede66/0184b69/81cc839/ab7b317 record196 cases:
10 genuine failures/186 passes in both evals rounds, then196 passes in impl.
Fixed Sol/high implementation worker passed196 cases, zero failures/skips,
in204.687s. Only entry.go/run.go changed privately; interfaces/evals remained
frozen. Invalid duration/other-command uses return corrective U01 before
mutation; valid duration reaches RunRequest.For. Clean time stops share the
existing exit3 result path, while prior blocks retain exit1.

The isolated audit found APP-089's valid-usage wording lacked a direct
assertion. Native round1 retained run/IDs; a scoped positive run/--for usage
pattern corrected that assertion. APP-089-only re-audit passed. Complete audit
command logs show permitted snapshot reads, no implementation/SYSTEM/history/
outside content, code execution or writes; snapshots removed. Native provenance
is manual; real Sol/high CLI logs remain under
20261029T165024-3fef7348df22114b38bedf67. Pre-freeze fixture corrections are
recorded in the change Notes; no existing assertion/helper was weakened.

NEW built at clean ab7b317, Go1.27.2 darwin/arm64, CGO disabled,
vcs.modified=false. Actual Go fixture proof:

- Nested-directory run --max1 --for2s --json returned3 in2.09s, interrupting
  the agent's10s sleep. Schema1 retained applied=[], empty=false, limited=false,
  reason="run time limit 2s reached", diagnostics[] and empty stderr.
- Original claim bdcd40af47eaf8774be8664aa30e049f0efebb14, run
  20261029T172859-31e13c6547f6d46de414b5c9, evals phase, ALP-002 and exact
  partial work survived. HEAD stayed at fixture baseline; no block, saved
  attempt, commit or release occurred.
- Later next --json returned0 in5.84s and archived the same run/ID, retaining
  partial work. Actual ALP-002 Go reports failed before implementation and
  passed afterward. Four calls total (interrupted evals/evals/audit/impl),
  three checked commits4d5f3b7/68e007b/32f0461. Beta/Gamma stayed identical.
  Final tree clean and claims absent.

Fixture baseline bde39998b3ecda1883c33dc2a070407a97164073; final
32f04619b8659c42fddc6d5e9b2f027e142f8bac. Details in
/tmp/fireweed-time-cli-result.json and /tmp/fireweed-time-cli-commands.json.
Fixture/state tracker removed after PASS. Deterministic fake agents/auditor
establish orchestration, not real-agent/drill acceptance.

This closes the time-stop/owned-resume CLI slice under partial R13/R17. Four
applications since0089–0093 review; last push28ced38. Remaining kinds/bounds,
full design and final gates stay open. Next: maintenance starts directly in
implementation, proven through NEW before the fifth-change checkpoint.
