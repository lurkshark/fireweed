// Package project discovers and validates repository-local fireweed documents.
package project

import (
	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/project/internal/inspection"
)

type Diagnostic struct{ Code, Severity, Message, Path string }
type Snapshot struct {
	Root, BaseRef    string
	Config           contracts.Config
	Modules          []contracts.Module
	Changes, Archive []contracts.Change
	Diagnostics      []Diagnostic
}

func Load(start string) (Snapshot, error) {
	d, err := inspection.Load(start)
	return Snapshot{Root: d.Root, BaseRef: d.BaseRef, Config: d.Config, Modules: d.Modules, Changes: d.Changes, Archive: d.Archive, Diagnostics: publicDiagnostics(d.Diagnostics)}, err
}
func Check(s Snapshot) []Diagnostic {
	d := inspection.Data{Root: s.Root, BaseRef: s.BaseRef, Config: s.Config, Modules: s.Modules, Changes: s.Changes, Archive: s.Archive}
	for _, v := range s.Diagnostics {
		d.Diagnostics = append(d.Diagnostics, inspection.Diagnostic{Code: v.Code, Severity: v.Severity, Message: v.Message, Path: v.Path})
	}
	return publicDiagnostics(inspection.Check(d))
}
func publicDiagnostics(ds []inspection.Diagnostic) []Diagnostic {
	result := make([]Diagnostic, len(ds))
	for i, v := range ds {
		result[i] = Diagnostic{Code: v.Code, Severity: v.Severity, Message: v.Message, Path: v.Path}
	}
	return result
}
