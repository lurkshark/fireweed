# Latest controller pin — verified idle handoff

The owner-requested pin update promoted clean source
`a26f7b9909f6c311416a0706d1b16749d16b9146` as controller version 0.1.0.
It was built with Go 1.27.2, `-trimpath`, darwin/arm64 and
`vcs.modified=false`. Current pin metadata remains in
[PROGRESS](../../docs/rebuild/PROGRESS.md).

- Controller: `/Users/joel/Workspace/fireweed/.fireweed/bin/fireweed`
- SHA-256: `962149ff80fb5f6e396dbafda4f44fb47bb28840d96d37b3c7c9718444c7bf5a`
- Retiring source: `9a206c558e8e5625c0ae55bb57c42ae8130b4f15`
- Recovery: `/Users/joel/Workspace/fireweed/.fireweed/bin/fireweed-9a206c5`
- Recovery SHA-256: `cc22f47c7aa607c96a13273d29d5bf4ade4f912453840f621331d73dc385f112`

The compiler-selected source/assets/module inputs and toolchain/dependency/build
settings matched the retiring build. Validation reused the prior
[runtime suite proof](workflow-contract-diagnostics-handoff.md) and recent audited
[0185–0187 eval verification](eval-contract-corrections-handoff.md). Fresh candidate
version, static check, all ten hook-doctor checks and the 16-commit explicit history
delta `9a206c5..a26f7b9` passed. Candidate main-repository status exactly matched
the retiring controller, including counts and warnings. Existing own-verification
timing exceptions and B02 warnings remain unresolved.

Promotion checked a clean matching HEAD, empty queue/claims and no active pinned
controller. The retiring executable was preserved before the verified candidate
was atomically installed. No fallback or cross-controller resume occurred.
Direct installed-pin version/static checks passed; idle Context U02 and Verify
P02 confirmed no held change. The owned validation clone was cleaned up.

The owner subsequently authorized pruning the six expanded controller-pin
attachments while retaining this receipt, current pin metadata and recovery
executable. Original attachments and the expanded handoff are recoverable at
`fb86abf` through the [retention index](README.md#controller-pin-attachments).
Actual module drill reports and all proof/attempt refs remain preserved.

The pin-only handoff and attachment cleanup invoked no drill or source correction.
WF4 remains open. Recovery requires explicit owner direction. The existing
**2026-11-10 22:51:05 UTC** deadline and all bounds remain unchanged; stop before
any further work without a new owner-selected scope.
