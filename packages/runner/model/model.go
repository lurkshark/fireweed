package model

import (
	"time"

	"github.com/lurkshark/fireweed/packages/project"
)

type AgentRequest struct {
	Snapshot project.Snapshot
	Role     string
	Prompt   string
	Bin      string
	Timeout  time.Duration
}

type AgentResult struct {
	Stdout     string
	Stderr     string
	ExitCode   int
	UsageLimit bool
}
