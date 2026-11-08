# Latest controller pin — verified idle handoff

The owner requested updating the controller pin to the latest source. Exactly one
idle handoff is complete. Source is `a26f7b9909f6c311416a0706d1b16749d16b9146`, the
latest committed source when requested. Subsequent protected handoff documentation
records the build; no product source was changed and no drill was invoked.

## Build and reused evidence

A fresh local no-hardlinks clone at the exact source was clean before and after
building. `go build -trimpath` used Go 1.27.2 on darwin/arm64. The
[build metadata](controller-pin-a26f7b9/build-info.txt) records that revision,
`vcs.modified=false` and the expected dependency and compiler settings.

`go list -deps -json .` identified 95 project compiler inputs, including Go files,
embedded recipe assets and go.mod/go.sum. Their Git blob identities are unchanged
from retiring source `9a206c558e8e5625c0ae55bb57c42ae8130b4f15`; see the
[input comparison](controller-pin-a26f7b9/runtime-inputs.json). Toolchain,
dependency versions/checksums and build flags also match the retiring build;
main-module/VCS provenance now identifies the requested latest source.

Validation therefore reuses the prior
[normal/race/vet/CI runtime validation](workflow-contract-diagnostics-handoff.md)
for those unchanged compiled inputs, plus the recent normal audited
[0185–0187 eval verification](eval-contract-corrections-handoff.md) for all changed
evals. Those archives retain independent restricted audits and required phase
checks, with 652 passing case rows covering 355 current IDs. This handoff did not
rerun the complete Go/race/CI suites or claim a new independent regeneration proof.
Existing own-verification timing exceptions and B02 warnings remain unresolved;
no bound was raised or performance correction attempted.

## Fresh candidate checks

The candidate's setup hook copied its exact bytes into the isolated tree's
`.fireweed/bin/fireweed`; scratch PATH used that directory first, with
GOCACHE=/tmp/fireweed-rebuild-cache. The original main pin stayed unchanged until
all candidate checks terminated successfully.

[Validation receipts](controller-pin-a26f7b9/validation.json) retain commands,
outputs, exit codes and elapsed times:

- `version --json`: exit 0, version 0.1.0, no diagnostics.
- `check --json`: exit 0, no diagnostics.
- `check-commits 9a206c558e8e5625c0ae55bb57c42ae8130b4f15..a26f7b9909f6c311416a0706d1b16749d16b9146 --json`: exit 0, all 16 delta commits checked, no diagnostics.
- `doctor --no-agent --json`: exit 0, all ten baseline/canary/recovery checks passed, no diagnostics or stderr. The project Go hooks, randomized fixture modules and candidate setup copy were used; no real agent was invoked.
- Candidate `status --json` in the main repository: exit 0, exactly equal to the retiring controller's parsed output, including counts, queue and warnings.

The build and all read/hook checks were terminal; the main tree and isolated
source were clean, without claims, before promotion. No lifecycle run, worker,
audit or drill was launched for this pin-only request.

## Atomic promotion and preservation

At promotion, HEAD still matched the requested source. Queue and shared claims
were empty; process checks found no active pinned controller. The retiring pin
matched its recorded hash, and the verified candidate matched its build hash.

The old bytes were preserved as `.fireweed/bin/fireweed-9a206c5`, SHA-256
`cc22f47c7aa607c96a13273d29d5bf4ade4f912453840f621331d73dc385f112`.
This is explicit owner-directed recovery only; no automatic fallback.

The verified candidate was written to a same-directory temporary executable,
flushed/fsynced and atomically renamed over
`/Users/joel/Workspace/fireweed/.fireweed/bin/fireweed`. New SHA-256:
`962149ff80fb5f6e396dbafda4f44fb47bb28840d96d37b3c7c9718444c7bf5a`.
The [promotion receipt](controller-pin-a26f7b9/promotion.json) records the idle
guards, exact source/hash, recovery copy and successful replacement.

Prior recoveries, reports, archives, proof/attempt refs, bounds/counters and the
unrelated worktree are preserved. The current failed workflow report/ref is
unchanged. GOAL and PROGRESS record the new pin and stopped state.

After installation, direct main-pin version/static checks passed. Idle Context
returned U02 and idle Verify returned P02 (“worktree holds no change”); these
[post-install receipts](controller-pin-a26f7b9/post-install.json) establish the
expected idle state, not an additional product-verification claim. The
[owned validation root was removed](controller-pin-a26f7b9/cleanup.json) only
after checking its clean source, no claim/additional worktree, terminal checks
and preserved new pin/recovery. No shared cache or unrelated worktree was removed.

## Stop

The single requested pin update is complete. No drill, reconstructed-source
import, source correction, inspector preparation, policy/tooling change or push.
No passing workflow proof or WF4/full-completion claim follows. Any future drill
must recheck its candidate/expected-pin metadata and configured commands under
its separately owner-selected scope. Existing deadline remains
**2026-11-10 22:51:05 UTC**, without extension. No further work is authorized.
