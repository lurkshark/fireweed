package model

import (
	"github.com/lurkshark/fireweed/packages/checks"
	"github.com/lurkshark/fireweed/packages/gitstore"
)

type Finding struct {
	ID         string `json:"id"`
	Clause     string `json:"clause"`
	Text       string `json:"text"`
	Untestable bool   `json:"untestable"`
}
type Verification struct {
	Change      string              `json:"change"`
	Phase       string              `json:"phase"`
	Diagnostics []checks.Diagnostic `json:"diagnostics"`
	Coverage    []checks.Coverage   `json:"coverage"`
	Edits       []gitstore.Edit     `json:"edits"`
	Findings    []Finding           `json:"findings,omitempty"`
}
