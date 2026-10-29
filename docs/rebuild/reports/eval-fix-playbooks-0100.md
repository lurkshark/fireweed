# Eval-fix playbooks 0100 — actual NEW guidance PASS

AST-029/030/031 applied at f418f880a9adc2577ffbb8555147c31a64ab94cc.
Native phases b2841bf/6361924/f418f88 record569 cases: three genuine
failures/566 passes, then all569 passing. Full implementation verification
completed in about12 minutes, within15m; only the documented B02 CHK-055/
16-case warning remained. Own checks:56 passes/three intended failures,
zero skips in11.133s; then59 passes, zero failures/skips in10.859s.

Isolated fixed Sol/high assertion audit passed all three IDs without findings.
Complete command log shows only permitted snapshot reads, no implementation/
SYSTEM/history/outside content, execution or writes; snapshot removed. Native
provenance is manual; actual CLI logs remain under
20261029T181904-9eb88dd4f65008635fa9f934. Evals/interfaces stayed frozen.

Only private content.go changed in impl, adding conditional eval_fix paragraphs
to evals/audit/impl. They instruct evals-only correction with listed coverage
and passing results, audit of corrected assertions, and change-file-only impl.
Ordinary guidance, audit scope/convergence/example and no-commit instructions
remain. No API, recipe/hook/version/diagnostic behavior, code organization,
bound change, phase exception, correction or drill.

NEW built from clean f418f88, Go1.27.2 darwin/arm64, CGO disabled,
vcs.modified=false. Actual playbook evals/audit/impl --json commands from
/private/tmp each returned0, schema1 string data with the conditional guidance,
diagnostics[] and empty stderr. Details: /tmp/fireweed-0100-cli-guidance.json
and /tmp/fireweed-0100-build-info.txt. The valid actual Go eval-fix lifecycle
proof from0099 remains valid and was not repeated.

Partial R17/R20/R23 guidance evidence closes this follow-up. Two applications
since passing0094–0098 checkpoint; last push7e0da27. No pending change/claim.
Full acceptance remains open. The two-change eval-fix slice is complete;
next is the separate missing regenerate lifecycle, without running a drill.
