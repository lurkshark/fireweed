---
audit_rounds: 0
bases:
    AST-003: ecb7d3e9
module: assets
phase: applied
run: 20261028T074132-8b1bb8fdf8bad38c40dc9fbd
---



## Motivation

The design requires every stable diagnostic to be explained. D006 restricts
explanations to codes actually emitted. Applied 0055 now emits C04, C05, C06 and
C07 on existing validation paths. Make explain useful for those current errors.

## Interface

No interface change.

## Added



## Modified

### AST-003
Explain MUST describe the cause and a corrective action for C01 (invalid project configuration), C04 (non-normative spec text), C05 (foreign or duplicate spec IDs), C06 (invalid spec tags), C07 (missing change Motivation) and P2-01 (missing genuine fail-first evidence). Explain MUST reject unknown codes.

## Removed



## Decisions



## Notes

Preserve the C01, P2-01 and unknown-code assertions. Extend the existing
Explanations test with one representative case per newly emitted code, checking
its specific cause and a useful corrective action, through the public interface.
Use the existing requirePattern helper. Do not add new helper infrastructure,
fixtures, unsupported code explanations or feature implementations. The four
new explanations must fail genuinely before implementation. Implementation adds
only the corresponding explanation text; other recipe/playbook assets stay
unchanged. Controller owns spec application, commits and archive.

## Provenance

- evals 2026-10-28T07:43:50.326974Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=b26da634732c7f1c7dc600115582a63cb8f82e1fd2bd1d522d608fe837a04ba7 tests=224 failing=5 passing=219
- audit 2026-10-28T07:44:35.185597Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=0fb044cfaca1b9a844a25b4c86223177e52f2f35fb149eddf044e83822665c88 tests=0 failing=0 passing=0
- impl 2026-10-28T07:46:57.81058Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=07554a3249d8fca0ba70df7504135f297491465caa63458734a4b8351734d0e5 tests=224 failing=0 passing=224

## Audit

No findings. Checked AST-003.
