# Commits CLI slice — PASS

Clean source6143176f793d965f0670f15a9f19afb0c066705f implements APP-094/095/096
using existing workflow.Commits. Phasescbea8fb/d8401ff/1ee7853/890c3e2/6143176:
both native eval rounds209 cases5 failing/204 passing, then209 full passes.
The single APP-095 finding was a missing operational cause-code assertion;
one-line correction expects existing P01 for the non-Git operational fixture.
Provided U02 selector-code preservation and every prior assertion retained.
Scoped APP-095 re-audit PASS; both independent auditor command lists contain
permitted snapshot reads only, no implementation/SYSTEM/history/outside reads,
code execution or writes. Snapshots removed. Focused implementation check6
named passes; full native app regression check209 passes under15m. No frozen
API/eval changes in implementation, only two app/internal dispatch/output files.

Actual NEW Go fixture baselinedd948bde094d94d3bc1ac9160ccedf13679c5e23 completed next
with real ALP-002 fail-first/pass, separate deterministic audit, three agents
and three checked phase commits. Run20261029T224128-e6958f1a1d8d8b0fe93eec5a, HEAD
4522efe8a6ac3d954cc1326c0c21111b51c33828. Then nested commits slug JSON and archive-path plain
both return exactly the following full phase IDs in newest-first order:

- 4522efe8a6ac3d954cc1326c0c21111b51c33828
- a35c8206277cd03e8282746303558d9fc147f17b
- 5ea2c05ca9a82b9675774524f1cb33885051ea30

Untouched pending Beta returns0 with schema1 data[], diagnostics[]; unknown
selector returns2 with null data/corrective U02. Every JSON response has schema1
and empty stderr. Read calls leave HEAD/archive/call log unchanged; tree clean,
claims none and agent count still3. Fixture removed after PASS. Details:
/tmp/fireweed-commits-cli-{result.json,commands.json,run.log}; clean CGO0
Go1.27.2 darwin/arm64 build info /tmp/fireweed-0108-build-info.txt has
vcs.modified=false and exact source6143176. Existing C01/C18 explanations apply.

This closes the0107 dependency's immediate caller, not full acceptance.
Applied since passing/pushed5358b5f gate:3 (0106–0108). Next: abandon workflow
and immediate app caller/NEW proof, then required fifth-application checkpoint.
No final CI/doctor/fresh-checkout/new real drill/self-hosting claim.
