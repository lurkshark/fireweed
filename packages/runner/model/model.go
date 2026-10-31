package model

import (
	"time"

	"github.com/lurkshark/fireweed/packages/checks"
	"github.com/lurkshark/fireweed/packages/project"
	"github.com/lurkshark/fireweed/packages/workflow"
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

type NextRequest struct {
	Root          string
	Bin           string
	AllowUnmerged bool
	Timeout       time.Duration
}

type StepResult struct {
	Phase string
	Role  string
	Agent AgentResult
}

type NextResult struct {
	State workflow.State
	Empty bool
	Steps []StepResult
}

type RunRequest struct {
	Root          string
	Bin           string
	AllowUnmerged bool
	Timeout       time.Duration
	Max           int
	For           time.Duration
}

type RunResult struct {
	Applied []NextResult
	Empty   bool
	Limited bool
}

type DrillRequest struct {
	Root    string
	Module  string
	Bin     string
	Timeout time.Duration
}

type DrillResult struct {
	Module     string
	Base       string
	Date       string
	Agent      string
	Model      string
	Result     string
	ReportPath string
	Ref        string
	Own        checks.Report
	Dependents map[string]checks.Report
	Steps      []AgentResult
}
