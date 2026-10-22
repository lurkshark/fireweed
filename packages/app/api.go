package app

import (
	"github.com/lurkshark/fireweed/packages/app/internal/entry"
	"io"
)

// Run executes the fireweed CLI and returns its process exit status.
func Run(args []string, out, errOut io.Writer) int { return entry.Run(args, out, errOut) }
