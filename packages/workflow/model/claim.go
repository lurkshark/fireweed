// Package model retains workflow wire data used by disposable implementations.
package model

type Claim struct {
	Worktree string `json:"worktree"`
	Change   string `json:"change"`
	Run      string `json:"run"`
	Base     string `json:"base"`
	State    string `json:"state"`
	Object   string `json:"-"`
}
