---
already_satisfied:
    - RUN-015
audit_rounds: 0
bases:
    RUN-015: b44edba6
module: runner
phase: impl
run: 20261029T055120-7a331d675361e12d225d4c11
---



## Motivation

The design's Retries and Blocking sections require exhausted work to be saved,
the checkout restored, and only block state committed with claim release.
NEW next already retries verification, and 0085 supplied the lifecycle Block
operation; return immediately to this caller so the CLI can finish that
observable outcome. Prove actual NEW next exhaustion in a disposable Go fixture
with a real failing implementation, saved work/ref, state-only block and clean
release. Broader independent continuation, resume and limit handling follow.

## Interface

No public signature or retained type changes. Compose the existing workflow
Block API inside Next's exhausted-verification path. Existing State and Steps
carry the result; the error identifies the saved attempt ref and original
verification cause. Preserve app's wrapped-cause classification and output.

## Added

### RUN-023
For exhausted phase verification with the worktree's approved claim and an
unchanged HEAD, Next MUST save and block the step through the lifecycle,
leaving a clean unclaimed checkout and a state-only block commit.

### RUN-024
After that blocking succeeds, Next MUST return the blocked lifecycle State and
ordered captured invocations without changing the run, assigned IDs or current
phase. Its error MUST retain the last verification cause and identify the saved
attempt ref.

## Modified

### RUN-015
When agent invocation fails, Next MUST return an error and captured invocation
results without advancing the phase or releasing its held claim.

## Removed



## Decisions



## Rejected alternatives



## Notes

RUN-015 narrows to the unchanged invocation-failure behavior, already satisfied
by its existing case. Move the changed terminal verification outcome to the new
items, retaining every old assertion/ID except the explicitly changed checkout/
claim expectations: failed edits now live on the attempt ref rather than in the
checkout, while phase/run/IDs/captured output and cause stay intact. Do not drop
behavioral checks. Reuse existing Go Next fixtures and representative default
exhaustion cases; no new failure matrix. Lifecycle Block's already-audited Git
semantics need no duplicate library-wide test campaign.

Preserve successful retries, approval checks, audit scope/rounds, nonzero-agent
stops, timeouts, phase ordering and existing Run composition. If lifecycle
blocking itself rejects a state, propagate the failure with verification context
and captured results; do not force release or alter metadata/HEAD to evade it.
Unapproved-state repair, agent-commit undo, from-attempt resume, Questions,
independent continuation, usage waiting, new kinds and unrelated cleanup remain
out of this change. No app or other-module edits; existing C10 cause assertions
must keep passing. A reported block is a failure, never a successful apply.

Workers never commit in this checkout, invoke controller lifecycle operations,
change governing runtime/counters/refs or edit SPEC/protected docs/other modules.
Use public APIs, temporary Git identities, bounded hooks/fake agents; no special
instrumentation, bootstrap/deleted-code/saved-attempt oracle or broad refactor.
Own Verify below five minutes, full below fifteen. Separate isolated assertion
audits remain required. No further phase-write exception is authorized.

## Provenance

- evals 2026-10-29T06:03:45.650988Z agent=manual fireweed=0.1.0 playbook=b26da634732c7f1c7dc600115582a63cb8f82e1fd2bd1d522d608fe837a04ba7 tests=213 failing=2 passing=211
- audit 2026-10-29T06:07:06.610377Z agent=manual fireweed=0.1.0 playbook=0fb044cfaca1b9a844a25b4c86223177e52f2f35fb149eddf044e83822665c88 tests=0 failing=0 passing=0

## Audit

No findings
Checked IDs: RUN-015, RUN-023, RUN-024
