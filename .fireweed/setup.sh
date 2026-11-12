#!/bin/sh
set -eu
# FIREWEED_BIN is supplied by the invoking controller, including in scratch trees.
mkdir -p .fireweed/bin
if [ "$(cd .fireweed/bin && pwd -P)/fireweed" != "$FIREWEED_BIN" ]; then
    controller_copy=$(mktemp .fireweed/bin/fireweed.XXXXXX)
    trap 'rm -f "$controller_copy"' EXIT HUP INT TERM
    cp "$FIREWEED_BIN" "$controller_copy"
    chmod 755 "$controller_copy"
    mv -f "$controller_copy" .fireweed/bin/fireweed
    trap - EXIT HUP INT TERM
fi
