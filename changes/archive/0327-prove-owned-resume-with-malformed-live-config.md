---
eval_fix:
    - RUN-033
module: runner
phase: applied
run: 20261118T234320-c291a8f5270a885fd70169fc
---

## Motivation

RUN-033 requires resume to use the approved snapshot. In the November 18 runner
drill, Next read live configuration before checking for a claim. The owning budget
eval changed the live value to another valid one and passed, while APP-164 detected
both malformed and incompatible live inputs.
[Drill reports](https://github.com/lurkshark/fireweed/pull/25)

## Interface

## Added

## Modified

## Removed

## Decisions

## Notes

Extend the existing owned-resume fixture at the runner.Next interface. After a clean
usage stop, make the live YAML malformed and resume with the already-approved
bounded fake agent. Assert the usage-stop classification, approved configuration,
unchanged run, phase, assigned IDs and saved work, and the approved agent
invocations. Assert that the malformed live file is preserved. Keep the existing
valid-but-different budget scenario and normal approval checks for fresh work.
## Audit

No findings; checked RUN-033

## Provenance

- evals 2026-11-18T23:53:33.099091Z agent=codex fireweed=0.2.1 playbook=d2fd7216b035a9614c390d254dd9195ac2c65ae1ed18f07ac93c77b9732d4156 tests=220 failing=0 passing=220
- audit 2026-11-18T23:54:01.011169Z agent=codex fireweed=0.2.1 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0
- impl 2026-11-19T00:06:28.280593Z agent=codex fireweed=0.2.1 playbook=5f856d6bcab308b9ff1a8a7c03b9d2547744112aba81dd30dea8f61ec8062b03 tests=635 failing=0 passing=635

