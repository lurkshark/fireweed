# Independent continuation 0089 — actual NEW proof PASS

Modified RUN-010/011/016 and added RUN-025/026 now compose canonical workflow
eligibility with Next and continue Run after successful clean saved blocking.
Only applied results count toward Max. A typed internal blocking error marks
completed lifecycle blocking, so ordinary held failures still stop. Terminal
errors retain the blocking verification cause and saved ref. Queued-file flags
retain their original meaning even when no change is eligible.

Native phase commits 423a9f7/5edbc1d/bf27309 record 217 cases: five genuine
failures/212 passes before implementation, then 217 passes. The real Sol/high
implementation worker's complete runner suite passed in 111.278s. Separate
isolated Sol/high assertion audit passed RUN-010/011/016/025/026; full command
log inspection shows supplied snapshot reads only, without SYSTEM, implementation,
history/outside reads, execution or writes. Snapshot removed. Provenance is
manual, with actual logs under run 20261029T082745-ee6d49740ab2b5b8d2e2120f.

The first evals worker preserved an obsolete all-blocked error assertion and
reported its conflict with modified RUN-010. A scoped pre-audit evals correction
moved that scenario to RUN-010, asserted the new normal Empty result and kept
the no-agent/no-claim checks. Other applicable assertions and interfaces stayed
unchanged. This was the intended behavior migration before freezing, not an
eval_fix or a post-audit test edit. No controller phase exception was used.

NEW rebuilt at clean bf27309d379f74a3af85c5841251cc84088308fe, Go1.27.2
darwin/arm64, CGO disabled, vcs.modified=false. The optional module stat-cache
write was sandbox-denied; build exited 0. Actual NEW init/Go hooks/fake agents
and run --json in a disposable fixture produced PASS in 14.46s:

- Alpha's evals genuinely failed, audit passed, and both impl attempts failed.
  Its final wrong-2 implementation and untracked ALPHA_BLOCKED_WORK file are on
  refs/fireweed/attempts/value-alpha/1 at 11d3ba60eb1cbca44072d6129f02641692452ed9.
  The checkout restored old Alpha code, the block committed only its state, and
  Alpha's spec stayed unapplied.
- Beta depends on Alpha and remained unstarted: no assigned IDs,
  run/phase or archive; old implementation/spec retained. Gamma is independent
  and completed evals/audit/impl, applied GAM-002 and archived. Alpha's blocked
  document stayed identical while Gamma ran.
- Seven bounded fake-agent calls and six branch commits were inspected: Alpha
  evals/audit/block dc084f0/d77082c/db74108, then Gamma evals/audit/impl
  2daafc0/6b9cbc1/96557f3. No claims or working-tree edits remained.
- CLI exit 1, schema1/data null, E01 for ALP-002 and saved ref, empty stderr.
  The successful independent work did not turn the invocation into success.

Fixture baseline c7e2336e4155e86c69872fbe866762c0819cc270. Alpha run
20261029T085043-183d7ee49a73545a7e7e4478; Gamma run
20261029T085053-00f77b1b518b6a44884a0d5f. Detailed local proof is
/tmp/fireweed-independent-run-smoke-result.json. The clean plain Git fixture
/private/tmp/fireweed-run-cli-ntnv8b5m is deliberately retained, tracked by
/tmp/fireweed-independent-run-state.json, for the immediately following saved
attempt resume proof; no active process or claim remains. Remove it after that
scenario. Its toy saved code is proof data, never a rebuild implementation oracle.

This closes independent continuation for partial R12/R17. Saved-attempt loading,
agent-commit undo, usage/time stops, remaining kinds/bounds and final acceptance
remain open. No new drill or packaging/final Go gate is claimed. The governing
checkout has no pending change or claim; one application since checkpoint.
