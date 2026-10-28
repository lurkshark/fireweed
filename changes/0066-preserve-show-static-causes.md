---
bases:
    APP-038: 037fc33c
fixes:
    - APP-037
module: app
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
