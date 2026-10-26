---
bases: {}
eval_fix:
    - APP-003
module: app
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
