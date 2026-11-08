# Single approved workflow retry

The owner separately approved exactly one supervised native workflow retry on
validated source `9a206c558e8e5625c0ae55bb57c42ae8130b4f15`, using the unchanged pin below
and prepared outside inspector. Reuse the committed clean-source validation and
unchanged supervisor rehearsals. Use a fresh isolated caller and fresh live gates;
review actual retained inputs before release and every command/script/incremental
patch, full first native feedback and fresh output before any remaining configured
release. Keep exact worker command/model/effort, input policy, bounds/counters and
**2026-11-10 22:51:05 UTC** deadline unchanged. Retain the actual terminal evidence/ref,
review the result, clean only owned scratch, commit protected evidence idle and
stop. Failure/gaps authorize no correction, additional invocation, source import,
controller update, later work or push. This approval supersedes only the previous
native approval stop; all preservation and independence requirements remain.

The owner-selected FLOW-007/FLOW-023 contract clarification and RUN-037 diagnostic
change are applied and independently audited. Clean-source validation and the
idle controller handoff are recorded in the
[preparation evidence](../../../.fireweed/drills/workflow-contract-diagnostics-handoff.md).
PROGRESS records the current pin; `control.py` identifies the validated candidate.
The original 79cf33a, prior corrected ccd3d6b, all preceding native reports/refs,
recovery executables, configured bounds/counters and unrelated worktree are preserved.

**Prepared only. Do not launch a native drill without separate owner approval.**
This sequence exercised inspector preparation and matching primary/caller copies,
then removed its owned idle clones. No live gate, verifier or reconstruction
worker was created. The commands below require a new outside root and fresh live
isolation/input/output review. No further correction, reconstructed-code import,
worker-input expansion or push is authorized. The existing deadline remains
**2026-11-10 22:51:05 UTC**.

The preceding native workflow invocation failed verification; its
[complete evidence](../../../.fireweed/drills/workflow-corrected-controller-retry-20261107T232508/review.md)
remains the current result. WF4 stays open. RUN-037's native timing cause remains
unresolved; richer diagnostics are not a runtime timing fix. Wrapper/inspector
behavior, exact worker command/model/effort, worker SYSTEM and all gates remain
unchanged. Earlier fake rehearsals remain valid for that behavior and establish
no native reporting behavior on rejection/interruption/emergency stop.

`guard.py` is an `agent.wrap` prefix. Its CLI accepts only the exact configured
command/model/effort from SYSTEM, with the original prompt as its last argument.
It retains the 600-second hold and two-attempt cap. It starts no child before
release. Rejection and launch/release share an outside lock; rejection latches
across all remaining attempts, including stale later approval files.

Every received stdout/stderr chunk is written and fsynced to separate
`gate/attempt-N/stdout.bin` and `stderr.bin` files outside the reconstruction tree
before the identical bytes are forwarded to the corresponding native pipe. A
disk spool prevents a blocked reader from blocking the worker drain. No banner,
diagnostic, decoding, line filtering or stream merging is added. Relative order
within each stream is preserved; no cross-stream ordering is promised. Wrapper
diagnostics and byte counts live only in outside gate records. `finished-N.json`
states whether both streams reached EOF and were completely forwarded.

On rejection or inspector SIGINT/SIGTERM, the inspector latches `gate/abort`.
The wrapper signals only its own reconstruction worker process group with TERM,
allows 5 seconds, then KILLs that group if needed. It drains/spools/forwards for
at most another 5 seconds, records completion or truncation, and returns failure.
This includes ordinary descendants after their leader exits. It prevents later
worker releases. The inspector leaves the native controller running for up to
120 seconds after wrapper completion so it can publish failure evidence. If a
wrapper does not finish within 11 seconds, or that reporting window expires,
the inspector escalates to bounded whole-run cleanup.

The original deadline independently starts emergency worker termination, even
without an inspection decision. Emergency cleanup allows the same bounded worker
stop/drain window before signaling the controller group with INT, then KILL after
5 seconds; the final collection window is bounded by another 5 seconds. It never
extends the deadline for reconstruction or releases another worker. Hard kills,
broken pipes, storage errors and escaped/daemonized processes can limit evidence:
preserve received bytes and explicit incomplete records; do not certify unseen
output or invent a native result. As with SYSTEM, processes that escape their
owned group are outside the cleanup guarantee.

The original isolation preflight is retained, with only its outside root supplied
as an argument: exact immutable input identities/bytes, no symlinks, input-only Git
with one commit and no remote/alternates/unreachable objects, verifier separation,
retained verifier bytes, deleted workflow disposable paths, dependency/dependent
implementations and manifests/hooks only in the verifier, exact fresh first
output before attempt two, and all three controller copies matching the pin.
Preflight runs on first readiness and again immediately before each release.
Actual retained-input review precedes release one. Full first commands/scripts,
every incremental patch, native feedback and repeated live isolation review
precede release two. Inspection supplies no reconstruction hints or extra inputs.
The worker SYSTEM policy is unchanged: no local builds/tests/module setup;
missing compilation infrastructure is intentional; format fresh code and return
it to the separate verifier. There are no command shims, interceptors or observers.

## Rehearse without a native drill

From the main repository, run:

```sh
PYTHONDONTWRITEBYTECODE=1 python3 docs/rebuild/workflow-drill/rehearse.py
```

The harness calls the same supervisor functions with explicit fake child argv and
short fixture bounds. It does not put a fake `codex` on PATH, change the production
command check, invoke a Fireweed executable, run a product drill, or certify native
isolation/reporting. It checks normal completion, reviewed second-attempt release,
interruption with termination tails and a descendant, pre-release rejection,
stale approval after rejection, reporting-window exhaustion, stubborn-worker and
whole-controller emergency cleanup, and production command rejection. Raw binary
payloads include NUL/non-UTF-8/CRLF bytes and exceed pipe capacity. Logs and native
forwarded files must match byte for byte, including live prefixes before exit.

## Restart only after new owner authorization

Read AGENTS, SYSTEM, PROGRESS, GOAL and ACCEPTANCE; inspect live pinned status,
queue, shared claims, process metadata and worktrees. Run context before each
applicable step. There is no workflow change during protected preparation, so
context reports U02 and verify reports P02; those are idle-state checks, not
product verification. Stop if the pin, deadline, candidate, bounds or independence
requirements cannot be preserved.

Use a new outside root, never a reconstruction workspace. The following command
only prepares a no-hardlinks isolated caller clone and copies this tooling outside
it. It checks out the exact candidate and commits a local setup child changing
only SYSTEM and `agent.wrap`; product/spec/eval/manifests/hooks remain unchanged.
The main repository is not changed and no drill is invoked. The copied tooling
and expected input/verifier hashes are outside the caller and worker inputs.

```sh
PYTHONDONTWRITEBYTECODE=1 python3 docs/rebuild/workflow-drill/inspector.py prepare /private/tmp/OWNER_SELECTED_NEW_ROOT
```

Inspect that clean setup child and `expected.json`, recheck the live gates and
approval, then launch the staged inspector alongside the single native controller
invocation. **This command requires a new explicit owner approval:**

```sh
PYTHONDONTWRITEBYTECODE=1 python3 /private/tmp/OWNER_SELECTED_NEW_ROOT/tooling/inspector.py launch /private/tmp/OWNER_SELECTED_NEW_ROOT
```

The inspector starts the pinned `drill workflow --json`, holds actual readiness,
and writes outside preflight logs. It releases nothing until an outside reviewer
atomically writes `gate/reviewed-1.json` with `result: "pass"`, the exact `input`
path from `ready-1.json`, the candidate SHA in control.py, and
`actual_retained_inputs_reviewed: true`. This is an attestation after actual
retained-input review, never a substitute for it. The inspector repeats preflight
and requires the ready record to remain unchanged.

While the worker runs, review both durable raw streams, all full commands/scripts
and incremental patches. Before any configured second release, finish that review
and review the full native feedback. Verify fresh output matches the reviewed
trace; the inspector captures its hashes while attempt two is held. Atomically
write `gate/reviewed-2.json` only after review, with the same `result`, `input` and
`candidate` fields, all of the following set to true:

- `full_first_commands_scripts_patches_reviewed`
- `full_native_feedback_reviewed`
- `second_prompt_only_original_plus_native_feedback`
- `fresh_output_matches_reviewed_trace`

Include `first_streams_sha256` with `stdout` and `stderr` hashes of the completed
first raw streams. The inspector checks their identity and complete forwarding,
repeats the original isolation preflight, then releases only the remaining
configured attempt. Readiness/review records and logs never enter worker inputs.

Reject from a second outside terminal as soon as an independence gate fails:

```sh
python3 /private/tmp/OWNER_SELECTED_NEW_ROOT/tooling/inspector.py reject /private/tmp/OWNER_SELECTED_NEW_ROOT 'specific observed reason'
```

Inspector SIGINT/SIGTERM uses the same worker-first stop. Keep the inspector
running through terminal collection in a host environment that can signal its own
child process groups. A preliminary tool-sandbox rehearsal denied a group signal;
the final host rehearsal passed without masking that error or changing worker
sandbox settings. Never send the first stop directly to the
native controller group. Preserve both raw streams and all gate/stop/terminal
records before removing only the owned input tree, verifier worktree and caller
clone. Review any native result/report/ref actually published; if absent, record
that absence. The fake rehearsal proves no native report fields, complete native
case list, proof ref publication or independence result. Commit protected evidence
only with no live run/claim, update the restart handoff, and stop for the owner.
