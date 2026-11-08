# CHK-034 fix and controller handoff

Protected author evidence after the owner-authorized normal audited fix. This is
not a native drill report, regeneration result, proof ref or claim that WF4 passed.
No native drill or push was invoked during this fix/handoff.

## Cause and scoped repair

The protected Go hook emitted compilation diagnostics, then produced an unmapped
PackageFailure testcase. Strict JUnit rejection was correct. Evaluate returned an
empty Report on hook execution, report-read or parse errors, discarding captured
stdout/stderr/exit before existing native feedback/report consumers could use it.
CHK-034 already required those fields; no spec change was needed.

The fix captures process metadata immediately after runHook and returns it on
those errors. Only successfully parsed cases/time are accepted. C11 rejection,
coverage, genuine fail-first, exit/report agreement and all existing enforcement
remain intact. No hook, runner, interface, spec, configuration, dependency-import
manifest or worker input was changed. Providing imports would not repair this
diagnostic-loss path; no import inventory was introduced.

## Audited lifecycle

Author setup: `959682a27d2ca55b08642165063d25e6d704d78e`.
The old controller stayed pinned throughout run
`20261107T185805-016f27a13886e3ddfbcd9f8f`; normal approval limits remained enabled.
Every phase used the unchanged configured command/model/effort and succeeded on
its first attempt, with independent restricted audit enabled. Fireweed authored
all phase commits and archived [0180](../../changes/archive/0180-retain-failed-hook-output.md):

- Evals `c0a23e0dc568433dc4a5c285151ff72f0fbd7e77`: genuine fail-first on
  unmapped PackageFailure, malformed XML, missing report and hook timeout; 119
  cases, 5 failing (including parent), 114 passing.
- Audit `17ce68f80301f227190c1a760cdeb11f8463707a`: No findings; checked CHK-034.
- Implementation/archive `12c571cc53c7e03b051811477adc83e5518f157d`: full required
  verification, 768 passing and zero failing across checks, app, runner and workflow.

Public regressions assert exact stdout/stderr/exit, unchanged diagnostic causes,
zero accepted cases/time from invalid reports, and continued coverage/fail-first
failure after an unmapped case even when a valid case precedes it. Timeout is
bounded and checks retained output/nonzero exit. Existing assertions remain intact.
Raw [phase reports](README.md#historical-evidence) (`d63c16e:.fireweed/drills/chk034-controller-handoff/impl/checks.json`) and archived
Provenance retain the result. The owning process exited zero and released its claim. Existing B02 testcase-count
warnings for CHK-055 (16), APP-103 (17) and APP-104 (19), against limit 15,
remain visible in status. These IDs and their tests are unchanged; no bound was
raised and no unrelated cleanup was undertaken.

## Clean-source controller validation

Built version 0.1.0 with Go 1.27.2, `-trimpath`, darwin/arm64 from a separate,
clean no-hardlinks clone at exact source
`12c571cc53c7e03b051811477adc83e5518f157d`. Build info confirms that revision and
`vcs.modified=false`. This fixed controller source is distinct from frozen subject
candidate `79cf33a7690082f28c06bf41a72719d282bc00fe`.

The [ordered commands, timestamps and exits](chk034-controller-handoff/validation/results.json)
and matching stdout/stderr files record passing `go test ./...`,
`go test -race ./...`, `go vet ./...`, controller check, and doctor with hook
canaries enabled (`--no-agent` avoids unrelated real-agent access probes).
CI also passed with its hook doctor enabled and all current module evals, under
the unchanged 1,500-second bound.

An initial standalone full-range history probe reached its external 300-second
bound without a history diagnostic; this is retained in
[history-timeout.json](chk034-controller-handoff/validation/history-timeout.json)
and is not claimed as a passing new full-range run. Existing committed WF2 history
evidence in [checks.md](checks.md) already covers
`be35552b706aa3000ff8f11306d129ff1d701b85..79cf33a7690082f28c06bf41a72719d282bc00fe`.
History implementation is byte-unchanged. Reusing that valid prefix, the new
eight-commit delta
`79cf33a7690082f28c06bf41a72719d282bc00fe..12c571cc53c7e03b051811477adc83e5518f157d`
passed explicit check-commits under the same 300-second bound; CI used that same
explicit delta. No configured bound, archive counter or proof requirement was reset.

A separate ordinary Go fixture exercised the real protected hook and compiled
controller, with no agent or Drill call: a deliberately missing `strings` import
made quick verify fail with C11. Its saved
[Report](chk034-controller-handoff/validation/diagnostic-report.json) retains
`undefined: strings` in stdout, `exit status 1` in stderr, HookExit 1 and null
Cases. The fixture claim was released. This establishes the compiled diagnostic
path, not native drill retry/report/ref publication or workflow independence.

## Pin handoff, preservation and stop

After validation and with no live run/shared claim, the owner-authorized primary
replacement was atomic. The primary SHA-256 is
`c9167d3878a1344719574d3ee1179d42295596f2f994c31c556ca25df8aaaada`.
The retiring source `193634686aa6a1c4eb32e69df1c85e1d5ce48bd7` is preserved at
`.fireweed/bin/fireweed-1936346` with SHA-256
`23e87b00217ab1a6990352393311bd150b7d60d3a4c138b75e285f518fba117f`.
All prior recovery executables remain byte-identical, with no fallback execution.
[Promotion receipt](chk034-controller-handoff/promotion.json) records the idle
P02 verify result, which is not product verification. Validation complete.json
describes the earlier pre-promotion stage; promotion.json records the later handoff.

Outside control.py and preflight.py changed only their expected pin hashes.
The preparation-only staging check copied the new pin, retained 31 input files,
394 verifier files and 16 disposable deletion paths, and changed only SYSTEM and
agent.wrap in its isolated setup child. No gate/native invocation was started.
[Preparation receipt](README.md#historical-evidence) (`d63c16e:.fireweed/drills/chk034-controller-handoff/preparation.json`) and expected
hashes are retained. Wrapper behavior, exact command/model/effort, worker-system
instructions, candidate, original deadline and every isolation gate are unchanged;
the earlier fake supervisor rehearsals remain valid for unchanged behavior.

[Integrity checks](chk034-controller-handoff/integrity.json) compare the entry
snapshot: old archives/reports/config/specs/interfaces/hooks, prior Fireweed
proof/attempt refs, old phase counters and unrelated linked worktree are preserved.
The new archive adds exactly its three normal phase commits; live unapproved is
one, not reset. The prior failed native workflow report remains byte-identical.
No reconstructed source was applied. All three owned temporary repositories were
removed only after verifying no claims or linked worktrees; [cleanup](chk034-controller-handoff/cleanup.json)
is recorded. The owning run is terminal, main has no queued change/shared claim,
and no validation or native drill is active.

Stop after this protected author evidence commit. The original deadline remains
2026-11-07 22:51:05 UTC. WF4 and full completion remain open; another native drill
requires separate owner approval and fresh live gates. No native rejection or
interruption reporting behavior is established by this handoff.
