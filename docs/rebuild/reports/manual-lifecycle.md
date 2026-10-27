# Milestone 1 — NEW manual Go lifecycle

Verified 2026-10-27 with NEW source `89b70289ee562d228e055f047d48470057beb607`,
built clean (`vcs.modified=false`). Fixture:
`/tmp/fireweed-stage2-manual-slice/project`, authored setup `5357850`.
The original checkout stayed under the pinned controller.

- NEW start assigned TOY-001 and run `20261027T144429-53d4980e0361844d890dea9e`; HEAD stayed at the author commit.
- Public `Test_TOY_001_Answer` called `toy.Answer()` and asserted 42. The disposable stub returned 0; the Go testcase failed and NEW verify accepted genuine fail-first evidence without a commit.
- A SYSTEM.md edit was rejected with Z01. Temporarily returning 42 during evals was rejected with P2-01. Both faults were restored before the evals transition.
- Separate Sol/high audit of only the change, projected spec and eval found no missing assertion for TOY-001.
- A Sol/high implementation worker changed only the internal return value to 42. Git diff against the evals commit was empty for the facade and eval.
- NEW done applied the spec and archived the change. NEW coverage reported TOY-001 pass. The fixture finished clean with no claim.

| Native phase | Commit | Recorded results |
| --- | --- | --- |
| Evals | f04e9f71a8d175d50f87c33f68b35cc61a059dbe | 1 failing testcase |
| Audit | bfb84d0cc80b313b589929c3a2215a0876111364 | No findings; checked TOY-001 |
| Applied | 1323a31dd1605bf86393b1fb8b4ef7c5c799dd8e | 1 passing testcase |

The applied SPEC contains `Answer MUST return 42.` The archive retains the same
run and all three provenance lines. These interactive phases recorded
agent=manual; PROGRESS.md explains the actual child command and the subsequent
switch to configured controller execution. This verifies the manual milestone;
init, doctor, unattended execution, regeneration and full acceptance remain open.
