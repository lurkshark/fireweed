package model

import "github.com/lurkshark/fireweed/packages/checks"

type StartResult struct {
	Context     WorkerContext       `json:"context"`
	Claim       Claim               `json:"claim"`
	Diagnostics []checks.Diagnostic `json:"diagnostics"`
}
