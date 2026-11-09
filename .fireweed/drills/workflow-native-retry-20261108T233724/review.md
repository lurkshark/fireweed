# Supervised native workflow retry failed

The owner authorized one native workflow invocation after the stopped controller
handoff. Pinned Fireweed ran `drill workflow --json` once and used both configured
reconstruction attempts. It exited 1 with E01 after the final workflow, app and
runner evals reported failures. The inspector recorded no rejection or emergency
escalation. WF4 and full completion remain open.

## Candidate and supervision

The candidate was `75efa9a03c4eb7a17673a53bdd8058074e5bb7e1`. The clean
no-hardlinks caller used setup child `68234aa118c7424ac497b7f3987adb21ce006909`.
Only the supplied reconstruction SYSTEM and the external `agent.wrap` prefix
differed from the candidate. The history baseline is
`be35552b706aa3000ff8f11306d129ff1d701b85`, 1,261 commits before the candidate.
Preparation reused the verified controller handoff, prior WF1/WF2 checks and
audited 0188–0189 verification. The earlier own-Verify timing exceptions remain
open. [Pin and validation](../controller-pin-75efa9a.md),
[audited correction evidence](../provenance-report-corrections-handoff.md),
[WF2 history and CI](../workflow-contract-diagnostics-handoff.md).

Protected preparation commit `3965218` updated the inspector's expected source
and pin hash. The worker command, worker SYSTEM, inspection behavior and bounds
were unchanged. The host rehearsal passed its six fake-process scenarios and
the command rejection check. [Rehearsal results](rehearsal-results.json).

The first review checked the 31 actual retained inputs against the candidate and
the supplied worker policy. The workflow SPEC, Decisions, facade/model and
public evals were present, with only the declared dependency facades and models.
Both preflight passes for each attempt checked the same input-only Git history
with one commit, no remotes or alternates, and no unreachable objects. The
verifier kept 564 files unchanged. The first preflight checked the deletion of
all 16 original disposable paths. Both attempts checked that all three
controller copies matched the pin. [Input review](retained-input-review.json),
[gate records](stream-and-gate-receipt.json).

The workers used the configured ephemeral Codex command with gpt-6.1-sol,
high effort, ignored user config and workspace-write. Review covered both
workers' emitted commands, full scripts and incremental patches. Observed
reads stayed within supplied inputs and fresh output. Writes stayed within
workflow's disposable directory and the drill notes. The workers formatted
their output and returned it to the native verifier. There were no observed
worker-local builds or tests, worker commits, or changes to retained inputs.

The first worker returned 13 Go files and notes. The native build succeeded,
and the first eval feedback contained 120 failing rows. Most reflected the
staging command's failure on ignored `.fireweed/runs` or the host Git's rejection
of `worktree list --porcelain -z`. The feedback also recorded waiting-queue
failures and a missing B03 warning. Before releasing the second worker, review
covered the first trace and full native feedback, matched fresh output to the
reviewed writes, and confirmed that the second prompt was exactly the first
prompt plus the controller's feedback. The second worker changed three Go
files and the drill notes to address staging, worktree listing and saved budget
warnings. The second changes matched the emitted patches. [First trace review](first-trace-review.json),
[second trace review](second-trace-review.json).

Both workers exited 0 and forwarded every received byte to the corresponding
native pipe. Both spools matched the controller's invocation logs. The current
report preserves the native bytes as its prefix, including all four streams and
the full first feedback in the second stderr. The receipt records exact byte
offsets, lengths and digests. [Native report](../workflow.md),
[stream and gate receipt](stream-and-gate-receipt.json).

## Final native verification

| Module | Rows | Pass | Fail | Enforced IDs referenced | Suite seconds |
| --- | ---: | ---: | ---: | ---: | ---: |
| Workflow | 175 | 172 | 3 | 134/134 | 159.246 |
| App | 323 | 316 | 7 | 155/155 | 248.480 |
| Runner | 159 | 147 | 12 | 67/67 | 247.865 |

The verifier compiled the final reconstruction and accounted for all 657
reported rows, including parent and subtest rows. All enforced IDs were
referenced. The reported case failures leave the proof open.
[Native fields, diagnostics and failing rows](native-result.json).

- Two workflow Abandon cases and app's owned JSON Abandon case hit the host
  Git's rejection of `git apply --cached --3way`. The app failure appears in
  both its parent and subtest rows.
- FLOW-062 returned the X01 boundary count and omitted the underlying violation
  detail expected by the eval.
- Five app rows and twelve runner rows failed on retained report discovery. The
  reconstruction wrote `<phase>/<module>/report.json`. The affected fixtures
  looked for reports such as `<phase>/alpha.json`. RUN-037 is in this group and
  this result does not resolve the earlier timing exception.

The worker reported no contradicted Decisions. Its eight gap entries concern
the claim format, editable sections, audit grammar, run-output staging,
attempt conflicts, saved budget diagnostics, historical metadata and Git
compatibility. These are worker observations pending owner review.
[Native Decisions and gaps](../workflow.md#contradicted-decisions).

## Retention and stop

The native ref is
`refs/fireweed/drills/workflow/20261108T233724.098992000`, OID
`3f42938d7d8ded117cd40aa5e4ea6b5d984212bd`. The retained tree's parent is the
setup child. Its retained files match that child and its reconstructed files
match the reviewed second output. The ref was transferred from the local
caller. The reconstruction was never applied to main. The previous native
report is available at `3965218:.fireweed/drills/workflow.md`.
[Proof review](proof-review.json), [previous review](../workflow-native-retry-20261108T090949/review.md).

The native controller removed the input and verifier. The remaining caller,
outside staging and rehearsal directories were removed after evidence transfer.
The pinned controller's static and preparation-history checks passed. Idle
Context U02 and Verify P02 confirmed the idle state. The local
`origin/main` tracking ref moved from `5cf1ba4` to `3965218` during the drill,
with a Git reflog entry of “update by push.” The live unapproved count went
from two to zero after that move. This session issued no push and left the ref
movement in place. The isolated caller and prior Fireweed refs were unaffected.
[Cleanup and final checks](cleanup.json).

The authorized invocation is consumed. The handoff stops idle at the existing
deadline of **2026-11-10 22:51:05 UTC**. [PROGRESS](../../../docs/rebuild/PROGRESS.md).
