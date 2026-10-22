# Restartable rebuild progress

Status: stage 0 setup; no behavioral changes started. Goal remains active.

## Controller pin

- Source: /Users/joel/Workspace/fireweed-bootstrap at 583b3c53740cf2341e0b0a79afee5b7f646ec1d4 (clean before/after build).
- Build: GOCACHE=/tmp/fireweed-rebuild-cache go build -o /Users/joel/Workspace/fireweed/.fireweed/bin/fireweed-bootstrap . (from bootstrap checkout).
- Toolchain: go1.27.2 darwin/arm64. Bootstrap version JSON: 0.4.0, schema 1.
- Binary SHA-256: f9c48ae6b14887ebc549110478c081bedc450847507c6967735ef8486c98a3ed. Ignored binary must be rebuilt and checksum checked if absent.
- Design SHA-256: 1c17d0ab9496b1b15ad5ca18eed4f4ca623d01cff596bc557bf4c97400d00789.
- Initial target: 1b27763, clean main. No remote changes, pushes or publication.
- Bootstrap help/version/author-change read; Go recipe initialized with codex worker and codex auditor. Codex CLI present and login status reports ChatGPT authentication; a real invocation remains unverified.

## Setup and next work

Eight scaffold modules with proposed dependency DAG; protected design/planning/build/enforcement files; updated Go canary imports; compiling app facade/placeholder. Controller check and go test ./... pass (no behavioral tests yet). Setup baseline and doctor results will be recorded after the setup commit. Next: prove all canaries and recovery; author contracts grammar/ID/hash changes through fail-first, independent audit and impl.

All R01–R21 remain unverified. No product acceptance, real regeneration or self-hosting is claimed.

## Restart procedure

Inspect git status, pinned controller checksum/version, controller status/list and refs/fireweed/claims before any edit. Read current change/context if claimed; resume its phase without resetting work. Only update protected evidence between active changes. Final history/CI uses explicit setup-SHA..HEAD, never main default range.
