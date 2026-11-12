---
audit_rounds: 0
bases:
    AST-003: 477c5b10
module: assets
phase: applied
run: 20261028T083045-4f27220f190698eb33960992
---



## Motivation

The design and D006 require useful explanations for codes actually emitted.
Applied 0060 now emits U01 for existing invalid/unsupported requests and U02
for unknown explain/playbook lookups. Explain currently rejects those codes.

## Interface

No public interface change.

## Added



## Modified

### AST-003
Explain MUST describe the cause and a corrective action for C01 (invalid configuration), C04 (non-normative text), C05 (foreign or duplicate spec/edit identities), C06 (invalid tags), C07 (missing Motivation), C08 (missing/stale edit bases or absent targets), C09 (undeclared handles), X01 (boundary violations), X02 (missing/incomplete boundary hooks), X03 (reasoned boundary opt-out), X04 (unnecessary boundary enforcement), U01 (invalid arguments/options or unsupported commands), U02 (unknown lookup names) and P2-01 (missing genuine fail-first evidence). Explain MUST reject unknown codes.

## Removed



## Decisions



## Notes

Preserve all existing explanation and unknown-code assertions. Extend the
existing Explanations test with one case each for U01 and U02, checking their
specific cause and a corrective action using requirePattern. No new helper or
fixture infrastructure. U01 must cover invalid argument/option and unsupported
command errors; its repair should point to supported syntax/help. U02 must
cover an unknown lookup name and selecting a valid name. These two currently
unknown codes must genuinely fail before implementation. Only add explanation
text for these emitted codes; do not implement commands or explain unused
allocations. Controller owns specs, commits and archive.

Fail-first evidence: `go test ./packages/assets/evals -run 'Test_AST_003_Explanations/U0[12]' -count=1` failed for both U01 and U02 because Explain returned `unknown diagnostic code`. `verify --quick` reported both new subcases failing and all existing explanation subcases passing.

## Audit

No findings. Checked: AST-003.

## Provenance

- evals 2026-10-28T08:33:15.770359Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=b26da634732c7f1c7dc600115582a63cb8f82e1fd2bd1d522d608fe837a04ba7 tests=240 failing=3 passing=237
- audit 2026-10-28T08:33:53.445577Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=0fb044cfaca1b9a844a25b4c86223177e52f2f35fb149eddf044e83822665c88 tests=0 failing=0 passing=0
- impl 2026-10-28T08:35:26.21446Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=07554a3249d8fca0ba70df7504135f297491465caa63458734a4b8351734d0e5 tests=240 failing=0 passing=240
