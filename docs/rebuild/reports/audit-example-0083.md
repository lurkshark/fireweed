# Audit example 0083 — first findings round

First native evals commit 2033c79: 461 cases, one genuine AST-028 runtime failure
and 460 passes. The isolated Sol/high auditor reported the exact clause
followed by a list of IDs on separate lines as incompletely asserted. Full
command-log inspection shows permitted supplied-file reads only. The snapshot
was removed. Native findings commit ee2d097 returned to evals and incremented
audit_rounds to 1; native release provided a clean author boundary.

The original author wording was ambiguous about separating individual IDs
versus separating the No findings line from the Checked IDs field/list. The
first eval required an empty label and a bulleted ID list, checking one ID only.
The pending author change now explicitly requires No findings on one line and
Checked IDs: plus at least two illustrative IDs on the next, with replacement
by actual checked IDs. This preserves the intended accepted example and all
other instructions. Module SPEC is untouched; native application remains required.
No runtime field, counter, ref or phase was manually changed, and no new change
was created to reset bounds. Evals will be corrected in the existing findings
round, followed by AST-028-only re-audit.

An auxiliary Go-fixture check of the earlier below-label layout failed test
collection before reaching audit (C11 PackageFailure); its owned fixture was
removed. A smaller direct NEW lifecycle fixture then confirmed both IDs were
recognized from one comma-separated bullet line and advanced audit to impl.
That format flexibility is not needed by the clarified canonical example; no
product parser change or resolution of the auxiliary collection failure is
claimed. The existing actual Go next/run proofs remain separately recorded.
