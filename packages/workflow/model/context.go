package model

import "github.com/lurkshark/fireweed/packages/contracts"

type FileContent struct {
	Path    string `json:"path"`
	Content string `json:"content"`
}

type WorkerContext struct {
	Change       contracts.Change   `json:"change"`
	ChangeText   string             `json:"change_text"`
	Module       contracts.Module   `json:"module"`
	SpecText     string             `json:"spec_text"`
	System       string             `json:"system"`
	Dependencies []contracts.Module `json:"dependencies"`
	Files        []FileContent      `json:"files"`
	Playbook     string             `json:"playbook"`
	PlaybookHash string             `json:"playbook_hash"`
	GitCommonDir string             `json:"git_common_dir"`
}
