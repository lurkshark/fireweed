---
audit_rounds: 0
bases:
    APP-038: 037fc33c
fixes:
    - APP-037
module: app
phase: applied
run: 20261028T095449-04ccdcde82574460479716ce
---



## Motivation

D006 assigns graph validation C02 and spec identity validation C05. Post-0065
NEW smoke demonstrated show returning U02 for a duplicate module name, and U02
for a duplicate ALP-001 even when requesting the distinct ALP-002. These are
static failures, not unknown lookups. APP-038's ambiguous-name wording let those
causes conflict with APP-037. Clarify the distinction and correct only the two
misclassifications, preserving legitimate projected change-local handle ambiguity.

## Interface

No public interface change.

## Added



## Modified

### APP-038
CLI lookup rejections for unknown names or ambiguous projected change-local handles MUST use U02. Invalid graphs, module spec identities and other static validation failures MUST retain their D006 static codes.

## Removed



## Decisions



## Notes

Add the two demonstrated regressions using existing app fixtures and appCause:
show alpha with duplicate graph name alpha must return C02; show alpha ALP-002
with duplicated ALP-001 and a unique ALP-002 must return C05. Preserve the plain
and JSON status/message/cause assertions and all existing inputs and coverage.
Label the regression assertions for APP-037 and APP-038; they must genuinely
fail against 0065 before implementation. No matrix expansion or new helpers.

Keep the existing show-ambiguous-item fixture using two separate queued changes
that each declare the same local handle: those handles are valid in their local scopes and
an ambiguous projected handle lookup remains U02. Do not relabel that case C05
or make lookup resolve inputs it previously rejected. Preserve every other
Code, message, status, JSON shape and success/rejection condition. Correct only
existing show classification: graph-validation failures remain C02, and duplicate
fixed module spec identities remain C05 even if another item was requested.
No new checks, unsupported features, speculative errors, interfaces or other
module edits. Controller owns specs, commits and archive.

## Audit

Checked IDs: APP-037, APP-038
No findings

## Provenance

- evals 2026-10-28T09:56:59.671264Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=b26da634732c7f1c7dc600115582a63cb8f82e1fd2bd1d522d608fe837a04ba7 tests=124 failing=3 passing=121
- audit 2026-10-28T09:57:45.097988Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=0fb044cfaca1b9a844a25b4c86223177e52f2f35fb149eddf044e83822665c88 tests=0 failing=0 passing=0
- impl 2026-10-28T09:59:53.958271Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=07554a3249d8fca0ba70df7504135f297491465caa63458734a4b8351734d0e5 tests=124 failing=0 passing=124
