---
audit_rounds: 0
bases: {}
eval_fix:
    - APP-003
module: app
phase: impl
run: 20261026T203404-273992a569da374db2370375
---



## Motivation

The independent standalone-cli audit found that APP-003 tests do not assert exit
2 for JSON invalid inputs or envelope fields for all named input categories.
The controller discarded its line-folded quotes and advanced the prior change.
Correct the frozen evals through the design's eval_fix protocol, retaining APP-003
coverage and requiring all app evals to pass, without changing product behavior.

## Interface



## Added



## Modified



## Removed



## Decisions



## Notes

The auditor's original report is in the isolated input directory
/tmp/fireweed-audit-standalone-cli/audit.md and ignored worker transcripts.
The prior rejected correction was preserved at
/tmp/fireweed-standalone-cli-corrected-evals.go for the eval_fix worker to review.

## Provenance

- evals 2026-10-26T20:35:16.09413Z agent=manual fireweed=0.1.0 playbook=b26da634732c7f1c7dc600115582a63cb8f82e1fd2bd1d522d608fe837a04ba7 tests=26 failing=0 passing=26
- audit 2026-10-26T20:36:36.313905Z agent=manual fireweed=0.1.0 playbook=0fb044cfaca1b9a844a25b4c86223177e52f2f35fb149eddf044e83822665c88 tests=0 failing=0 passing=0

## Audit

No findings and checked APP-003
