#!/bin/sh
set -eu
# FIREWEED_BIN is supplied by the invoking controller, including in scratch trees.
mkdir -p .fireweed/bin
if [ "$(cd .fireweed/bin && pwd -P)/fireweed-bootstrap" != "$FIREWEED_BIN" ]; then
    cp "$FIREWEED_BIN" .fireweed/bin/fireweed-bootstrap
    chmod 755 .fireweed/bin/fireweed-bootstrap
fi
ln -sf fireweed-bootstrap .fireweed/bin/fireweed
