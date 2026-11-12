# Releasing fireweed

The release workflow publishes GitHub release archives for macOS and Linux on
amd64 and arm64, SHA-256 checksums, a source archive, and a Homebrew formula.
Pushing a stable `vMAJOR.MINOR.PATCH` tag runs the checks and publishes the
release, then builds and tests Homebrew bottles and updates
`lurkshark/homebrew-tap`. Nothing is published by a branch push, pull request,
or local snapshot build.

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

Install Go, Node.js 24 or newer, pnpm 10.15.1, Python 3.11 or newer, uv,
and GoReleaser v2 (validated with 2.20.1), then run:

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

`make test` and `make check` give fixture subprocesses an explicit Git identity
and the `C` locale so commits and filename sorting do not depend on the host.
Both CI and release checks install the Node/pnpm and Python/uv recipe tools.

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
   generated formula, and publishes the GitHub release. Its reusable
   **Homebrew bottles** workflow then packages and tests all four native
   bottles, uploads them, and updates the tap. Stable releases only are
   supported by these workflows and the installer.
5. In a clean installation, run:

   ```sh
   brew install lurkshark/tap/fireweed
   brew test lurkshark/tap/fireweed
   fireweed version
   ```

Existing Homebrew users update with `brew update && brew upgrade fireweed`.
Users who need an older version can use the explicit-version installer or
`go install github.com/lurkshark/fireweed@v0.1.0`.

If a bottle or tap job fails after GitHub publication, fix the cause and rerun
only the failed jobs. Successful bottle jobs retain their workflow artifacts
for the merge job. The upload step reuses identical existing assets but refuses
to overwrite different published bytes. The tap job downloads the published
`fireweed-bottled.rb` without rebuilding anything. Do not move a published
release tag or replace its original archives.

To add bottles to an existing release that does not have them, run:

```sh
gh workflow run homebrew-bottles.yml --ref main -f tag=v0.1.0
```

This uses the existing release's `fireweed.rb` and checksum-verified binary
archives. It adds bottle assets and a separate `fireweed-bottled.rb`; it does
not replace the original formula asset, archives, checksums, or release tag.

## Homebrew implementation

`scripts/homebrew.py` emits a formula with the matching binary URL and SHA-256
for each OS/architecture and a `brew test` that checks `fireweed version`.
The formula is the input to native `brew install --build-bottle` and
`brew bottle --json` jobs on macOS 15 (Intel and Apple Silicon) and Ubuntu
24.04 (amd64 and arm64). Each job uninstalls the source installation, pours
its bottle, checks the installed receipt, and runs `brew test`. The Intel
macOS runner first installs current dependencies, building them when needed:
Homebrew no longer publishes new Intel macOS core bottles. This does not
remove Homebrew's developer-tool requirements for Intel users.

Homebrew merges the four metadata files into the final formula's `bottle`
block. Bottle archives use a single hyphen in their release URLs even though
Homebrew creates local filenames with a double hyphen; the publishing job
uses the names in Homebrew's JSON metadata. Only after every bottle passes
and all assets are uploaded does the workflow update the tap.

Matching bottles avoid Fireweed's source-install SDK checks. macOS 15 bottles
also work on newer macOS versions. Homebrew's platform requirements and the
Git dependency still apply; unsupported systems or explicit source installs
may still require developer tools.

Homebrew installs the binary without compiling Go. This small generator avoids
GoReleaser's [deprecated formula publisher](https://goreleaser.com/customization/publish/homebrew_formulas/)
while keeping Linux support. It follows Homebrew's
[OS and architecture blocks](https://docs.brew.sh/Formula-Cookbook).

The standalone installer downloads only the requested version, verifies its
archive checksum before installing, and never resolves a moving `latest` URL.
Checksums detect corrupt or mismatched downloads; they are served by the same
GitHub release as the binaries. Fireweed itself performs no network operations
or self-updates.
