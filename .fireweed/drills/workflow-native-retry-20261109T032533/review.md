# Supervised native workflow drill failed

The owner authorized one controller promotion and one native invocation. The
installed controller came from the newest committed source at approval,
a54f5da8d4faf9758f1a426cc7559c0726be8607. The drill ran on the separately approved
candidate 7d17a76186fd94d603a018713bdcc1005c89e020. The pin and candidate had
identical compiler inputs. The preparation commit was 4b9cb0e and the isolated
setup child was 4cb2bc783f2ed97e5d407edb2015560a88007975, with only the supplied
worker SYSTEM and external agent.wrap prefix differing from the candidate.
[Controller handoff](../controller-pin-a54f5da.md),
[candidate validation](../workflow-report-boundary-corrections.md).

## Independence and supervision

The host rehearsal passed all fake-process scenarios and command rejection.
Preflight twice checked each attempt. The 31 immutable inputs matched the
candidate and the worker policy, including the workflow SPEC, Decisions, public
evals, and declared dependency facades and models. The input Git repository had
one commit with no old objects, remotes, or alternates. The separate verifier
retained 577 files and removed all 16 original disposable paths. All three
controller copies matched the pinned hash.
[Rehearsal](rehearsal-results.json), [input review](retained-input-review.json),
[gate records](stream-and-gate-receipt.json).

Each worker ran the configured ephemeral Codex command with gpt-6.1-sol, high
effort and workspace-write. The review covered 30 shell commands in each
attempt and all incremental patches. Observed reads stayed within supplied
inputs and fresh output. Observed writes stayed within the disposable workflow
directory and DRILL_NOTES.md. The workers formatted their source and used
read-only source inspection. The outside verifier owned build and eval
execution. No supervisor rejection or emergency escalation occurred.

The first reconstruction compiled and ran the three suites. The native feedback
reported 113 failing rows, counting parent and subtest rows. Workflow had 49,
app had 20, and runner had 44. The main error was Git staging encountering
ignored .fireweed/runs files. The feedback also showed a missing-base code
mismatch, status accepting an unavailable base, and incorrect waiting-queue
selection. Before release two, the review matched all nine fresh files to the
first trace and confirmed the second prompt was the original prompt plus the
exact 240,268-byte native feedback.
[First trace review](first-trace-review.json),
[second trace review](second-trace-review.json).

Both workers exited 0 and forwarded every captured byte. The raw spools matched
both native invocation logs. The report preserves all four raw streams in its
exact native prefix, including the full first feedback in the second stderr.
[Native report](../workflow.md), [stream receipt](stream-and-gate-receipt.json).

## Final verification

The second reconstruction failed to compile. Its claim.go used regexp without
importing it. The compiler emitted undefined regexp errors on lines 576 and
578 in each suite. No final workflow, app, or runner testcases ran, so the
final report contains no testcase rows or coverage. The CLI returned exit 1
with C11 and C10 diagnostics for the package-failure placeholders. Those
diagnostics follow the compiler failure. The first attempt's eval results
cannot establish whether the second attempt addressed those failures.
[Native result](native-result.json).

The worker reported no contradicted Decisions and seven gap entries covering
ID allocation, decision formatting, boundary process output, patch conflicts,
worker-editable sections, approval counting across file moves, and verifier
setup. These are worker observations. [Decisions and gaps](../workflow.md#contradicted-decisions).

## Retention and stop

The controller published refs/fireweed/drills/workflow/20261109T032533.158401000
at cc457c9d257233455f38400a06d11105f41a70e1. The proof tree's 577 retained
files match the setup child, and its eight reconstructed Go files match the
reviewed writes. The exact ref was transferred from the local caller. The
previous native report is available at 4b9cb0e:.fireweed/drills/workflow.md.
[Proof review](proof-review.json), [previous review](../workflow-native-retry-20261108T233724/review.md).

The controller removed the input and verifier. The caller, staging root, and
rehearsal directory were removed after evidence transfer. The final pinned
static and history checks, idle Context U02, idle Verify P02, clean worktree,
empty queue, and absence of shared claims are recorded in the cleanup receipt.
[Cleanup receipt](cleanup.json).

The authorized invocation is consumed. WF4, full completion, and earlier
timing exceptions remain open. The deadline remains 2026-11-10 22:51:05 UTC.
[Rebuild handoff](../../../docs/rebuild/PROGRESS.md).
