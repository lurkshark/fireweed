#!/bin/sh
# Install one explicit release; never resolve "latest" or update automatically.
set -eu

fail() { printf '%s\n' "fireweed: $*" >&2; exit 1; }

[ "$#" -ge 1 ] && [ "$#" -le 2 ] || fail "usage: sh install.sh <vMAJOR.MINOR.PATCH> [install-directory]"
tag=$1
printf '%s\n' "$tag" | grep -Eq '^v(0|[1-9][0-9]*)\.(0|[1-9][0-9]*)\.(0|[1-9][0-9]*)$' ||
    fail "supply an explicit stable version, such as v0.1.0"
destination=${2:-"$HOME/.local/bin"}

case "$(uname -s)" in
    Darwin) platform=darwin ;;
    Linux) platform=linux ;;
    *) fail "supported systems are macOS and Linux" ;;
esac
case "$(uname -m)" in
    arm64|aarch64) arch=arm64 ;;
    x86_64|amd64) arch=amd64 ;;
    *) fail "supported architectures are amd64 and arm64" ;;
esac

command -v curl >/dev/null 2>&1 || fail "curl is required"
command -v tar >/dev/null 2>&1 || fail "tar is required"
if command -v sha256sum >/dev/null 2>&1; then
    checksum_tool=sha256sum
elif command -v shasum >/dev/null 2>&1; then
    checksum_tool=shasum
else
    fail "sha256sum or shasum is required"
fi

work=$(mktemp -d)
trap 'rm -rf "$work"' EXIT
trap 'exit 1' HUP INT TERM
archive="fireweed_${tag#v}_${platform}_${arch}.tar.gz"
base="https://github.com/lurkshark/fireweed/releases/download/$tag"
curl --fail --location --silent --show-error --proto '=https' --tlsv1.2 \
    "$base/$archive" -o "$work/$archive"
curl --fail --location --silent --show-error --proto '=https' --tlsv1.2 \
    "$base/fireweed_checksums.txt" -o "$work/checksums.txt"

expected=$(awk -v name="$archive" '$2 == name { print $1 }' "$work/checksums.txt")
[ "${#expected}" -eq 64 ] || fail "missing or ambiguous checksum for $archive"
case "$expected" in *[!0-9a-f]*) fail "invalid checksum for $archive" ;; esac
if [ "$checksum_tool" = sha256sum ]; then
    actual=$(sha256sum "$work/$archive")
else
    actual=$(shasum -a 256 "$work/$archive")
fi
[ "${actual%% *}" = "$expected" ] || fail "checksum mismatch for $archive"

# Extract only the binary, after the archive checksum has passed.
tar -xzf "$work/$archive" -C "$work" fireweed
[ -f "$work/fireweed" ] && [ ! -L "$work/fireweed" ] || fail "archive lacks a regular fireweed binary"
mkdir -p "$destination"
install -m 755 "$work/fireweed" "$destination/fireweed"
printf 'Installed fireweed %s to %s/fireweed\n' "${tag#v}" "$destination"
printf 'Ensure %s is on your PATH.\n' "$destination"
