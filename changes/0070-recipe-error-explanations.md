---
bases: {}
fixes:
    - AST-003
    - AST-024
module: assets
---



## Motivation

The 0065–0069 checkpoint found that init's emitted C12 and U02 recipe errors
lack accurate explain guidance. D006 and the design require every emitted code
to describe its cause and corrective action. The owner approved this scoped
correction and requires checkpoint subagent re-review before push/continuation.
Correct the two existing explanation texts without changing init behavior.

## Interface

No public interface change.

## Added



## Modified



## Removed



## Decisions



## Rejected alternatives



## Notes

C12 already covers invalid metadata/zones under AST-024. Its explanation must
also identify invalid local recipe metadata and direct repair to recipe.yaml,
while retaining module metadata/zone causes and their repair guidance. U02
already covers unknown requested names under AST-003. Its explanation must also
identify an unknown recipe name and advise choosing a supported recipe name or
a readable local recipe directory, while retaining existing lookup causes,
ambiguity guidance and repairs.

In the eval step, extend the existing C12 and U02 explanation assertions through
the public assets.Explain API; preserve every existing assertion, unknown-code
check, fixture and ID. These missing assertions must fail at runtime against
the current text. Reuse existing cases/helpers; no test matrix or new feature is
needed. In implementation, change only the C12/U02 explanation strings.

Do not change init, emitted codes, error conditions, public interfaces, other
explanations or permanent Specs/Decisions. Do not author drill-gap fixes, run
drills, or consult bootstrap/deleted implementations. The controller owns
phase commits/application/archive; workers never commit. The pending five-change
checkpoint count is retained until the approved correction passes re-review.
