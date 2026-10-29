# Saved-attempt CLI 0091 — actual NEW proof PASS

APP-087/088 expose start --from-attempt N through the verified workflow API,
validate exactly one positive integer and its start-only scope, and keep current
context/diagnostic/JSON behavior. APP-009 only qualifies unchanged no-flag start
and reuses its original assertions as already_satisfied. No facade/model change.

Native phases 79a5849/7d51927/523f908 record 178 cases: ten genuine failures/168
passes before implementation, then all 178 pass. The real Sol/high worker's
single full app Go JSON check passed in 142.764s. Separate isolated Sol/high
assertion audit passed APP-009/087/088; its complete log shows permitted snapshot
reads only, no SYSTEM/implementation/history/outside reads, writes or execution.
Snapshot removed. No post-audit assertion edits, phase exception or drill.
Coordinator provenance is manual; actual logs are under
20261029T094008-0aae0a3be7e6e86471f9d8d1.

NEW rebuilt at clean 523f908868b4460a3179d698a58ab90131e39a67, Go1.27.2
darwin/arm64, CGO disabled, vcs.modified=false. The optional module stat-cache
write was sandbox-denied; build exited 0. Reusing the actual
[independent-run](independent-run-0089.md) fixture and its saved failed work,
NEW start value-alpha --from-attempt 1 --json and manual correction/NEW done
passed in 2.86s:

- Author unblock a6a754b55f113ce644adae8e880ad81fff766fdf precedes the CLI call.
  Start returned schema1 context string with the restored saved.go path,
  empty diagnostics/stderr and exit 0 from a nested directory. HEAD remained
  that unblock commit; run 20261029T085043-183d7ee49a73545a7e7e4478,
  phase impl and assigned ALP-002 stayed unchanged. The owned claim used the
  current unblock HEAD as its phase base.
- Alpha's wrong-2 implementation and ALPHA_BLOCKED_WORK file were restored from
  refs/fireweed/attempts/value-alpha/1 at 11d3ba60eb1cbca44072d6129f02641692452ed9.
  After correcting Value to new, NEW done passed Alpha's two Go cases and Beta's
  dependent case, applied ALP-002 and archived Alpha at
  442943db5e966611bbcfbb0788b26578b7627e1d with the same run.
- Exactly two commits followed the prior Gamma head: the author unblock and
  the successful impl/archive. Beta's queued document and Gamma's archive stayed
  identical; no claim or working-tree edits remained. The attempt ref/OID was
  unchanged through completion. Fake-agent calls stayed at the earlier seven:
  resume, correction and Done invoked no additional agents.

The first harness invocation stopped before launching NEW because the preceding
block had removed its empty nested cwd. The already-written author unblock was
retained; recreating the directory and continuing from that exact commit avoided
repeating the failed run or changing IDs/refs. This was a harness setup failure,
not a product failure. The repaired harness and actual results are local at
/tmp/fireweed-saved-resume-smoke.py and /tmp/fireweed-saved-resume-smoke-result.json.
The continuation/resume fixture and its state tracker were removed after PASS.
Earlier reports' retained-fixture instructions are now superseded by this cleanup.

This closes the bounded saved-attempt resume scenario for partial R06/R12/R17.
It does not establish agent-commit undo, usage/time stops, remaining kinds/bounds
or final acceptance. The governing tree has no pending change/claim; three
applications since the last checkpoint, last reviewed push 82ef71d.
