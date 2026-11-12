# Releasing fireweed

The release workflow publishes GitHub release archives for macOS and Linux on
amd64 and arm64, SHA-256 checksums, a source archive, and a Homebrew formula.
Pushing a stable `vMAJOR.MINOR.PATCH` tag runs the checks and publishes the
release, then updates `lurkshark/homebrew-tap`. Nothing is published by a branch
push, pull request, or local snapshot build.

Release infrastructure lives in the protected author paths listed in
`fireweed.yaml`. It is maintained between lifecycle changes, separately from
worker steps. This implements the design's Distribution and init / Packaging
requirements. Product behavior and version changes still use the normal
spec/change workflow.

## One-time GitHub setup

1. Create a public `lurkshark/homebrew-tap` repository with an initial commit on
   `main` (a README is enough).
2. Create a fine-grained personal access token restricted to that tap, with
   **Contents: read and write**, and save it in `lurkshark/fireweed` as the Actions
   secret `HOMEBREW_TAP_TOKEN`. Allow this token to update the tap's `main` branch.
   The workflow's ordinary `GITHUB_TOKEN` publishes the release in the source
   repository; it cannot write to the separate tap.
3. Enable Actions in the source repository and allow the release job's
   `contents: write` permission. Restrict creation of `v*` tags to maintainers
   using a repository ruleset if desired.

The workflow checks that the tap exists and the token can read its `main`
branch before publishing. No credentials are needed for local packaging or PR
checks. The tap and token are setup requirements, not created by these files.

## Validate without publishing

Install Go, Python 3, and GoReleaser v2 (validated with 2.20.1), then run:

```sh
export PATH="$PWD/.fireweed/bin:$PATH"
export GOCACHE=/tmp/fireweed-rebuild-cache
make check build
fireweed check
make release-snapshot
ruby -c dist/fireweed.rb
```

`make release-snapshot` validates the GoReleaser config, cross-compiles all four
binaries, builds archives and checksums in ignored `dist/`, and generates the
formula after verifying each archive against its checksum. It never publishes
or updates the tap. Snapshot versions name the packages only; the binary still
reports the version required by its assets spec. The source archive contains
the committed tree, so commit author infrastructure before a real release.

GitHub CI runs the race suite, vet, build and project checks on Linux and macOS,
and separately exercises snapshot packaging. It builds the product at
`bin/fireweed`; it does not replace the local rebuild controller.

`fireweed verify` is for a held lifecycle change. With no held change it reports
P02; use it during version-change worker steps and use the checks above for
protected release-infrastructure edits.

## Publish a version

1. Finish any lifecycle changes and merge the release commit to `main`. Check
   CI and review the release diff.
2. Check `bin/fireweed version`. Currently `AST-001` specifies `0.1.0`, so the
   first release tag is `v0.1.0`. For a later version, change the assets spec,
   evals and implementation through Fireweed first. Do not inject a different
   version with linker flags: `go install`, archives, and project version pins
   must agree.
3. From the reviewed commit, create and push the tag:

   ```sh
   git tag -a v0.1.0 -m 'Release v0.1.0'
   git push origin v0.1.0
   ```

4. Watch the **Release** workflow. It tests Linux and macOS, rejects a malformed
   tag or binary/tag version mismatch, builds the archives, verifies the
   generated formula, publishes the GitHub release, and updates the tap. Stable
   releases only are supported by this workflow and the installer.
5. In a clean installation, run:

   ```sh
   brew install lurkshark/tap/fireweed
   brew test lurkshark/tap/fireweed
   fireweed version
   ```

Existing Homebrew users update with `brew update && brew upgrade fireweed`.
Users who need an older version can use the explicit-version installer or
`go install github.com/lurkshark/fireweed@v0.1.0`.

If the Homebrew job fails after GitHub publication, fix the tap permissions or
token and rerun only the failed job. It downloads the formula attached to the
existing release and does not rebuild or republish binaries. Do not move a
published release tag; use a new version for artifact corrections.

## Homebrew implementation

`scripts/homebrew.py` emits a formula with the matching binary URL and SHA-256
for each OS/architecture and a `brew test` that checks `fireweed version`.
Homebrew installs the binary without compiling Go. This small generator avoids
GoReleaser's [deprecated formula publisher](https://goreleaser.com/customization/publish/homebrew_formulas/)
while keeping Linux support. It follows Homebrew's
[OS and architecture blocks](https://docs.brew.sh/Formula-Cookbook).

The standalone installer downloads only the requested version, verifies its
archive checksum before installing, and never resolves a moving `latest` URL.
Checksums detect corrupt or mismatched downloads; they are served by the same
GitHub release as the binaries. Fireweed itself performs no network operations
or self-updates.
