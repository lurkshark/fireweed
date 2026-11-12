---
audit_pending:
    - RUN-044
audit_rounds: 1
bases: {}
module: runner
phase: applied
run: 20261031T075040-5475d59eb944f0afb625ad13
---



## Motivation

During 0129, workers passed quick own-module checks while repeated full checks
failed in a dependent module. Manual focused runs exposed the actual assertions.
NEW already forwards verification errors and hook streams, but the assertions
can exist only in JUnit. Consume Case.Failure from 0135 so the next attempt sees
the failed module, testcase and assertion without reconstructing them from logs.
Close this slice with NEW next in a temporary fixture whose fake worker saves
its retry prompt, demonstrating the JUnit-only assertion and bounded excerpt.

## Interface

No public API changes. Next's worker retry prompt gains a bounded failed-testcase
detail section using checks.Case.Failure from 0135-junit-failure-details.

## Added

### RUN-043
On a verification retry, Next MUST append a detail section identifying the module, classname, testcase name and Case.Failure for failed cases from every reported module, including dependents, subject to RUN-044.
It MUST select cases by lexical module name and then report order, excluding passing and skipped cases.

### RUN-044
The failed-case detail section MUST fit within 16,384 Unicode code points, including labels and truncation markers.
Each Case.Failure excerpt MUST contain at most its first 4,096 code points.
Shortened diagnostic text or omitted cases MUST be marked.

## Modified



## Removed



## Decisions



## Rejected alternatives

Requiring repeated full-suite runs just to discover the assertion: deliver the
failure already produced by verification to the retry worker.

## Notes

Depends on the checks report field introduced by 0135. Existing workflow.Result
already carries checks reports, so no workflow API or implementation change is
needed. Preserve RUN-022's existing verification feedback and hook streams;
limits apply only to the new failed-case section. Keep attempt counts, retained
work, usage/time stops and audit restrictions unchanged.

Use the existing bounded fake-agent and local JUnit-hook fixtures. Include an
assertion sentinel in JUnit that does not appear in hook stdout/stderr, including
a dependent-module failure. Check readable bounded excerpts and truncation;
headings need not have an exact prescribed wording. No real agent or network.

## Provenance

- evals 2026-10-31T08:03:11.52703Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.1 playbook=b26da634732c7f1c7dc600115582a63cb8f82e1fd2bd1d522d608fe837a04ba7 tests=364 failing=2 passing=362
- audit 2026-10-31T08:04:06.095812Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.1 playbook=5cb9f9dac24b0b6bb1b970edd4fab4406d2b4d0435ddefc3ef59a2cfdd267f18 tests=0 failing=0 passing=0
- evals 2026-10-31T08:12:36.370233Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.1 playbook=b26da634732c7f1c7dc600115582a63cb8f82e1fd2bd1d522d608fe837a04ba7 tests=364 failing=2 passing=362
- audit 2026-10-31T08:13:50.771143Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.1 playbook=5cb9f9dac24b0b6bb1b970edd4fab4406d2b4d0435ddefc3ef59a2cfdd267f18 tests=0 failing=0 passing=0
- impl 2026-10-31T08:28:24.142163Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.1 playbook=07554a3249d8fca0ba70df7504135f297491465caa63458734a4b8351734d0e5 tests=364 failing=0 passing=364

## Audit

No findings
Checked IDs: RUN-044
