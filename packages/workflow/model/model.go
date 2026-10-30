package model

import (
	"time"

	"github.com/lurkshark/fireweed/packages/checks"
	"github.com/lurkshark/fireweed/packages/project"
)

type State struct {
	Change      project.Change
	Snapshot    project.Snapshot
	Run         string
	Phase       string
	Head        string
	PhaseBase   string
	ClaimOID    string
	Diagnostics []project.Diagnostic
}

type Request struct {
	Root    string
	Change  string
	Bin     string
	Agent   string
	Quick   bool
	Timeout time.Duration
}

type Result struct {
	State       State
	Reports     map[string]checks.Report
	Costs       []ModuleCost
	Diagnostics []project.Diagnostic
}

type ModuleCost struct {
	Module          string
	Seconds         float64
	CasesPerSpec    map[string]int
	EvalLines       int
	DisposableLines int
}

type CostSnapshot struct {
	Run         string
	Phase       string
	RecordedAt  time.Time
	Modules     []ModuleCost
	Diagnostics []project.Diagnostic
}

type BlockRequest struct {
	Root   string
	Change string
	Reason string
}

type BlockResult struct {
	State      State
	AttemptRef string
	AttemptOID string
}
