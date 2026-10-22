package model

import (
	"github.com/lurkshark/fireweed/packages/checks"
	"github.com/lurkshark/fireweed/packages/gitstore"
)

type WorktreeCheck struct {
	Edits       []gitstore.Edit     `json:"edits"`
	Diagnostics []checks.Diagnostic `json:"diagnostics"`
}
