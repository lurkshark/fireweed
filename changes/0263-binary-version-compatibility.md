---
module: project
phase: audit
run: 20261114T032839-893b49cba44c12bc3d7ede60
---

## Motivation

The design's Distribution and init / Versioning requires the binary to
refuse projects outside their fireweed version range (R20). The 0.1.0 binary
from main 3b3b64d accepts fireweed: ">=9.0, <10.0". Supply one public
compatibility check for live loading and callers using approved snapshots.

## Interface

Add project.CheckCompatibility(config map[string]any) error to the retained
public facade, implemented in disposable code. It examines the supplied
configuration, so app can validate held approved snapshots without reparsing
malformed or missing live files. No Snapshot or Module fields change.

## Added

### PRJ-082
CheckCompatibility MUST accept an omitted fireweed pin or a nonempty comma-separated conjunction of comparisons using =, <, <=, > or >= and two- or three-component nonnegative decimal versions. It MUST compare against assets.Version numerically, treating an omitted patch component as zero.

### PRJ-083
CheckCompatibility MUST return a C01 CodedError when a present fireweed pin is malformed, is not a string, or excludes assets.Version. The error MUST identify the pin and running binary version and distinguish invalid syntax from incompatibility.

### PRJ-084
Load MUST apply CheckCompatibility to the live configuration before returning a Snapshot, and Check MUST expose a compatibility failure as non-warning C01 for fireweed.yaml.

## Modified



## Removed

## Decisions

## Notes

Implement only the comparison grammar stated above; do not introduce an
external semver dependency, network lookup, prerelease policy or undocumented
operators. Accept surrounding whitespace and pin endpoints as documented.
Use assets.Version to form compatible/incompatible fixtures and exercise the
actual design/init range as well as equality, numeric ordering and malformed
input. Missing pins remain unconstrained for existing unpinned projects.
Keep LoadAt as a historical reader: APP-127 and PRJ-047 need recorded data,
not validation against the binary used to inspect that historical snapshot.
The later app change checks each command's required approved/live configuration
source. It preserves APP-127's historical pr-body path without introducing a
live configuration dependency or validating a historical endpoint's binary pin
against today's executable. The format change extends the same helper.

Drafted at the owner's request after the R01–R29 traceability review.
No lifecycle execution, agent invocation, drill or release is authorized by
this drafting request. Workers must obtain context before their step and use
fireweed verify; only the controller may advance or commit lifecycle work.

The targeted Go evals reproduced fail-first for PRJ-082 and PRJ-083 through
the disposable stub and for PRJ-084 through the current live loader and
static check. `fireweed verify binary-version-compatibility` passed in evals
with fail-first evidence for all three IDs. The unpinned and compatible live
cases and the PRJ-047 recorded-pin case passed. Replace the stub in the disposable
`packages/project/internal/fixture/compatibility.go` and wire live validation
in `packages/project/internal/fixture/load.go`. The existing Check path
reports coded load errors through `packages/project/internal/queue/validate.go`.
## Provenance

- evals 2026-11-14T03:32:41.002593Z agent=codex fireweed=0.1.0 playbook=6a9c2d506363e3539cc66efd37bc6a06701249541dc4e4f01ed3aa6454bbb864 tests=262 failing=30 passing=232

