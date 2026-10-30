package model

import (
	"time"

	"github.com/lurkshark/fireweed/packages/contracts"
)

// Snapshot is the project data read from a repository root.
type Snapshot struct {
	Root    string
	Config  map[string]any
	Modules []Module
}

// Module is a discovered module and its declared document and zones.
type Module struct {
	Name      string
	IDPrefix  string
	Dir       string
	DependsOn []string
	Zones     contracts.Zones
	Document  contracts.Document
}

// Change is one pending change document and its spec edits.
type Change struct {
	Number   int
	Slug     string
	Path     string
	Module   string
	Document contracts.Document
	Added    []contracts.Item
	Modified []contracts.Item
	Removed  []contracts.Item
}

// SpecHistoryEntry describes one archived change affecting selected specs.
type SpecHistoryEntry struct {
	Path               string
	Module             string
	IDs                []string
	Motivation         string
	Decisions          string
	RemovedDecisions   string
	Provenance         string
	ImplementationDate time.Time
}

// Diagnostic reports a static authoring check.
type Diagnostic struct {
	Code    string
	Message string
	Path    string
	Warning bool
}
