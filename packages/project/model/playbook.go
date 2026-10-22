// Package model retains public project data used by disposable implementations.
package model

type Playbook struct {
	Name         string `json:"name"`
	Text         string `json:"text"`
	Hash         string `json:"hash"`
	EmbeddedHash string `json:"embedded_hash"`
	BasedOn      string `json:"based_on"`
	Overridden   bool   `json:"overridden"`
	Stale        bool   `json:"stale"`
}
