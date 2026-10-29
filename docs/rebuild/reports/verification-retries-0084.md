# Verification retries 0084 — PASS

0084 applied RUN-015/021/022 through native commits ffa53d0 (evals), 59d1a85
(audit), 667e419 (impl). Full phase evidence is 214 cases: five genuine runtime
failures/209 passes before implementation, then all 214 passing. The permanent
public-interface evals cover default two attempts, configured three attempts,
separate phase allowances, retained edits, latest actual verification feedback,
and exhaustion retaining the cause, invocation results and phase. Interfaces
and evals stayed frozen during implementation.

The initial eval worker hit Sol/high capacity. Its owner-approved one-time
same-profile retry succeeded, validated and removed temporary fixture tests,
and finished the evals. The first implementation worker also hit capacity after
focused tests passed; its separately owner-approved one-time same-profile retry
finished implementation and passed controller quick verification. No profile
was changed. Coordinator phase provenance honestly says manual; actual CLI
logs remain under run 20261029T032332-d2f9262403cce62e6c8ed149.

The coordinator initially launched isolated-audit-1 from the wrong working
directory. It read outside the prepared snapshot and was stopped; it supplies
no accepted audit evidence. The launch deviation and complete log remain in
the run directory. Fresh isolated-audit-2 ran from the snapshot and returned
No findings for RUN-015/021/022. Complete command-log inspection confirms only
supplied change/spec/eval/public-interface/rule reads, without implementation,
history or outside access, execution or writes. That snapshot was removed.
No runtime audit counter or phase was manually changed; round0 was retained.

## Actual NEW next proof

NEW was built at clean 667e419fc75c337650d4befaf1954a1826e8d21e with Go 1.27.2,
darwin/arm64, CGO_ENABLED=0 and vcs.modified=false. The optional module-version
stat-cache write was sandbox-denied; the build exited0.

In an ordinary disposable Go repository, actual NEW next exited0 in 6.53s,
stdout Applied changes/0001-one-value.md, empty stderr. One real eval failed
before implementation; the first implementation returned wrong and failed;
the second retained its edits and received the exact recorded verification
error plus complete hook stdout/stderr, then passed. Exactly four bounded fake
agent calls occurred: evals, separate audit, impl, impl. The result includes
SPEC application, archive, three phase provenance entries, claim release and
a clean working tree. The owned fixture was removed after inspection.

Fixture baseline a44ef87eaf5bdb44de0ab8ae360186cce396fb7b, run
20261029T043351-185bce8110c98772446d2bcc; phase commits
7ca577f62f311da289da5ee6fd2a345d17ffe09f,
a357c72126aca51f32b8701f431ff7d94e12ac49,
728b90ecb665d4e7c5841124dc6e9ca0d29db220. Each of the evals and first-impl
reports has one genuine failing Go case; final impl has one passing case.

The external smoke verifier initially overasserted a detailed Go failure line
that the recipe does not print to hook stdout. The command had already exited0.
The verifier was corrected to compare the actual saved error/stdout/stderr with
the retry prompt, then inspect the same completed fixture without rerunning the
change. No product or frozen-eval change was needed for that helper correction.

This closes only verification retry through the runnable Next/Run path. Saved
attempt refs/state-only blocks, independent continuation, from-attempt resume,
nonzero-agent retry, usage/time stops and remaining kinds/bounds remain open.
One change has applied since checkpoint; no new push or drill occurred. Full
design and final ACCEPTANCE gates remain incomplete.
