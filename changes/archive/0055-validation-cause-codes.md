---
audit_rounds: 0
bases:
    PRJ-008: 4efdf265
    PRJ-009: 3ab806fe
    PRJ-010: ac9cb4e7
    PRJ-011: 074dd58a
    PRJ-012: ad7c76da
module: project
phase: applied
run: 20261028T073403-02bb6a871c8d10a594784f1e
---



## Motivation

The design requires stable diagnostic causes for agents and explain. D006 assigns
C07 to missing Motivation, C04 to normative text, C06 to spec tags and C05 to ID
ownership/uniqueness. Existing Check paths still label each C01. Give those
existing rejections their assigned codes without adding errors or features.

## Interface

No public interface change.

## Added



## Modified

### PRJ-008
Check MUST reject a pending change without a nonempty Motivation section with diagnostic C07.

### PRJ-009
Check MUST reject a module spec or Added/Modified spec whose body contains no normative keyword with diagnostic C04.

### PRJ-010
Check MUST reject spec tags other than review or quarantine with a nonempty reason with diagnostic C06.

### PRJ-011
Check MUST reject duplicate spec IDs within a module's Specs section with diagnostic C05.

### PRJ-012
Check MUST reject module spec IDs and change edit IDs whose prefix belongs to another module with diagnostic C05.

## Removed



## Decisions



## Notes

Use existing CheckValidAndDiagnostics table fixtures and its follow-up Modified
normative/tag assertions to assert the exact assigned code for these IDs. Keep
all valid and invalid inputs, reason and nonempty code/message assertions and ID
coverage. Reuse the 53 cases; no new fixture, test case or helper is needed.
Unchanged IDs retain their checks. The exact cause assertion must match the
rejection being tested, not an unrelated diagnostic from the same fixture.
These changed IDs must genuinely fail before implementation; no already_satisfied.
Implementation migrates only the existing error branches that supply these
causes. Keep messages, existing rejection behavior and other cause labels intact.
Do not add unsupported features or speculative error paths. Evals and interfaces
freeze for implementation; controller owns specs, commits and archive.

## Provenance

- evals 2026-10-28T07:37:44.156467Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=b26da634732c7f1c7dc600115582a63cb8f82e1fd2bd1d522d608fe837a04ba7 tests=200 failing=9 passing=191
- audit 2026-10-28T07:38:27.136686Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=0fb044cfaca1b9a844a25b4c86223177e52f2f35fb149eddf044e83822665c88 tests=0 failing=0 passing=0
- impl 2026-10-28T07:40:48.625927Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=07554a3249d8fca0ba70df7504135f297491465caa63458734a4b8351734d0e5 tests=200 failing=0 passing=200

## Audit

No findings. Checked PRJ-008, PRJ-009, PRJ-010, PRJ-011, PRJ-012.
