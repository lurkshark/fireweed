---
audit_rounds: 0
bases:
    PRJ-013: 403da758
    PRJ-014: f84a2ce6
    PRJ-015: 76e385f9
    PRJ-016: 940ee23a
module: project
phase: applied
run: 20261028T074742-dcbae3e58d8dd0936380dd2a
---



## Motivation

The design requires stable diagnostic causes. D006 assigns C05 to duplicate
identities, C09 to handle references and C08 to bases. Check already rejects
these errors using C01. Migrate those existing paths, treating an absent edit
target as a failure to resolve its preceding base, without adding rejections.

## Interface

No public interface change.

## Added



## Modified

### PRJ-013
Check MUST reject a change that uses the same edit identity more than once across Added, Modified and Removed with diagnostic C05.

### PRJ-014
Check MUST reject a pending change that references a handle absent from that change's Added entries with diagnostic C09.

### PRJ-015
Check MUST reject Modified or Removed entries whose bases hash is missing or differs from the preceding projected spec's canonical hash with diagnostic C08.

### PRJ-016
Check MUST reject Modified or Removed IDs absent from the projection preceding the change with diagnostic C08.

## Removed



## Decisions



## Notes

Reuse the existing CheckDiagnostics table and its Removed-base follow-ups. Add
exact assigned-code assertions matched to each existing rejection's reason;
retain all current inputs, valid cases, nonempty code/message checks and reason
assertions, including missing/stale Modified and Removed bases and both absent
target cases. Preserve earlier exact-code assertions. Do not add fixtures,
scenarios or helpers; reporting labels may change as needed for correct ID
attribution. Changed IDs must genuinely fail before implementation.
Migrate only the existing duplicate edit, undeclared handle, missing/stale base
and absent preceding target paths, including existing projection-application
errors for those absent targets. Preserve messages and error behavior, and leave
other diagnostic causes unchanged. Do not implement missing features or create
speculative errors. Controller owns spec application, commits and archive.

## Provenance

- evals 2026-10-28T07:51:45.052296Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=b26da634732c7f1c7dc600115582a63cb8f82e1fd2bd1d522d608fe837a04ba7 tests=200 failing=7 passing=193
- audit 2026-10-28T07:52:33.147461Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=0fb044cfaca1b9a844a25b4c86223177e52f2f35fb149eddf044e83822665c88 tests=0 failing=0 passing=0
- impl 2026-10-28T07:54:35.974578Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=07554a3249d8fca0ba70df7504135f297491465caa63458734a4b8351734d0e5 tests=200 failing=0 passing=200

## Audit

No findings. Checked PRJ-013, PRJ-014, PRJ-015, PRJ-016.
