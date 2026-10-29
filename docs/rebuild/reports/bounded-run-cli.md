# Bounded NEW run CLI proof — PASS

Runner dependency 0081 applied at 6bdfb43 (evals 4c8c330, audit b788753):
204 full phase cases passed after seven genuine runtime failures. App wiring
0082 applied at 58351b2 (evals 0e17dae, audit 8220c66): 166 full phase cases
passed after seven runtime failures. Separate isolated Sol/high audits returned
No findings for RUN-016–020 and APP-080–086; full command logs show permitted
reads only, and snapshots were removed. Coordinator provenance is honestly
manual; actual CLI logs remain under runs 20261029T020109-5ec3f2b5c9478ce3dff394f1
and 20261029T022255-defa616cfd7f19df2cabf3fe.

The new app diagnostic fixture's exit/report mismatch was caught while evals
were still editable. A temporary public runner.Run probe showed C11 before the
one-line exit correction and C10/ALP-002 afterward. The harness was removed;
all permanent assertions/IDs stayed intact before native fail-first commitment.
No frozen repair or phase-write exception was needed.

## Actual executable and queue

CGO_ENABLED=0 build passed from clean source
58351b2d4f024cd3b8c5aec9be72685430cf4a3f (Go 1.27.2 darwin/arm64,
vcs.modified=false). An optional version stat-cache write was sandbox-denied
without failing the build. The subject was the independent
/Users/joel/Workspace/fireweed/bin/fireweed, never the bootstrap.

A fresh ordinary Go-recipe fixture used explicit Git identity, three independent
modules/changes and separate deterministic fake worker/auditor scripts, with a
configured 30-second agent timeout. Commands ran from a nested module directory:

| Invocation | Exit / duration | Actual result |
| --- | --- | --- |
| run --max 1 --json | 0 / 6.35s | Alpha applied; two changes remained; limited=true, empty=false |
| run --json | 0 / 9.16s | Beta and Gamma applied in order; limited=false, empty=true |
| run --json on empty queue | 0 / <0.01s | Empty applied array; no new calls or commits |

Each response used schema 1 and empty stderr. Alpha's archive stayed unchanged
during continuation. All three assigned prefix-002 cases genuinely failed in
Go before implementation and passed afterward; saved Cases/IDs/Result fields
were inspected. SPEC application, archives, evals/audit/impl provenance, no
remaining claim and a clean final tree were verified. Exactly nine fake-agent
calls used NEW as FIREWEED_BIN and exactly nine phase commits were created.
Commit phases were inspected from documents, not inferred from subjects.

Fixture range:
3f5fe6b1bfadf2d7ad947dda6bcf475e70802ae1..a98f065ddb10be78b323e7659820f28c925619a5.

| Module | Evals commit | Audit commit | Applied commit |
| --- | --- | --- | --- |
| alpha | 18fcba2bdc2be590e263256146aae9a6fdfa05ca | dc2f6d88d443f907d09d10db1dadf4976a3ece0f | 2b72eb2dc773af1dd569fb6926e66d85fea24142 |
| beta | 7d1fe1019d2a49e3456d4c1a8487094046de334c | 33b3bf89b247f7c457535380e91b44c7b1a6aa25 | 3d8c1028cb7ffa540d71a74e1b663615f438c741 |
| gamma | 8172472caeec969afb26101a9b932025a9c197da | 058fa52f480d75f0e6980d73cfcb7066fd716b13 | a98f065ddb10be78b323e7659820f28c925619a5 |

Runs: alpha 20261029T024424-6b120260d7405b1741383f48,
beta 20261029T024430-048fc577457451c4f454f4ef,
gamma 20261029T024435-6ad43886d86061a066e273f3. The temporary repository was
removed after inspection. Script/results remain under /tmp/fireweed-run-cli-*.

## Scope and next work

This closes the first bounded fake-agent run slice. It does not prove retries,
saved/block attempts, projected independent continuation, resume, agent-commit
undo, --for or usage waits, approval skipping, remaining kinds, real regeneration
or self-hosting, other platforms or final acceptance. Those remain in PLAN and
ACCEPTANCE. Native phase checks passed; a fresh whole-repository race/vet/history
gate follows at the standing five-change review. Applied-since-review is 4; no
new push has occurred. The observed passing-audit syntax guidance mismatch is
the small fifth change, followed by checkpoint review.
