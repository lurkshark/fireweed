package model

import (
	"time"

	"github.com/lurkshark/fireweed/packages/gitstore"
	"github.com/lurkshark/fireweed/packages/project"
)

type Case struct {
	Name      string
	Classname string
	IDs       []string
	Result    string
	Seconds   float64
}

type Report struct {
	Cases    []Case
	Seconds  float64
	HookExit int
	Stdout   string
	Stderr   string
}

type HookRequest struct {
	Snapshot project.Snapshot
	Module   string
	Bin      string
	Timeout  time.Duration
}

type BoundaryRequest struct {
	Snapshot project.Snapshot
	Bin      string
	Timeout  time.Duration
}

type ZoneRequest struct {
	Snapshot          project.Snapshot
	Module            string
	Phase             string
	ChangePath        string
	Base              string
	InterfaceDeclared bool
	Changes           []gitstore.Change
}
