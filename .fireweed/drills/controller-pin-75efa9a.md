# Controller pin 75efa9a — verified idle handoff

The owner requested one controller update. Latest committed source at the request,
`75efa9a03c4eb7a17673a53bdd8058074e5bb7e1`, was built from a clean main checkout
with Go 1.27.2, `-trimpath`, darwin/arm64 and `vcs.modified=false`. The candidate
was atomically installed as version 0.1.0 after validation.

- Pin: `/Users/joel/Workspace/fireweed/.fireweed/bin/fireweed`
- SHA-256: `fea9d492473bbb20883a74fc15ea5921b161727e4ea7096b42b3a5a909d60c65`
- Retired source: `a26f7b9909f6c311416a0706d1b16749d16b9146`
- Explicit recovery: `/Users/joel/Workspace/fireweed/.fireweed/bin/fireweed-a26f7b9`
- Recovery SHA-256: `962149ff80fb5f6e396dbafda4f44fb47bb28840d96d37b3c7c9718444c7bf5a`

All 95 compiler-selected repository source, embedded asset and module files
across 30 project packages match the retiring source. Toolchain, dependency and
build settings match apart from the expected main module version and Git revision
and timestamp. The previous [runtime suite proof](workflow-contract-diagnostics-handoff.md)
and [audited 0188–0189 verification](provenance-report-corrections-handoff.md) were
reused. The only Go changes in the delta are the two selected eval files.

Fresh candidate version and static check passed. Agent-free doctor passed all ten
hook checks in a temporary local clone, including detection of the four injected
failures and their recovery checks. Candidate `check-commits a26f7b9..75efa9a`
checked all fourteen commits with no diagnostics or squash entries. Candidate
and retiring controller status JSON matched byte for byte: empty queue, two
unapproved applied changes, and the same three existing B02 warnings. The
configured cap and all bounds were untouched. The existing runner/app timing
exceptions remain unresolved; this handoff does not retest or close them.

Immediately before promotion, HEAD matched the clean candidate, no shared claim
or Fireweed process was present, and the retiring binary matched its recorded
hash. The retiring copy was preserved and verified before the atomic replacement.
The installed hash, version and static check matched the candidate. Installed
List returned an empty queue; idle Context U02 and Verify P02 confirmed that no
change was held. Those idle responses do not constitute product verification.
The owned doctor clone was removed; the mutable product binary, all existing
recovery executables, Fireweed refs, and main and unrelated worktree state were
verified unchanged by the promotion. Recovery remains explicit, with no fallback.

No product correction, workflow execution, drill preparation, native drill or push
occurred. The failed native report and its proof ref remain preserved. WF4 and
full completion remain open. The single pin-update scope is consumed; stop idle
with the existing deadline **2026-11-10 22:51:05 UTC**. See [PROGRESS](../../docs/rebuild/PROGRESS.md).
