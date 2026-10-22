// Package project discovers and validates repository-local fireweed documents.
package project

import (
	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/project/internal/allocation"
	"github.com/lurkshark/fireweed/packages/project/internal/finalization"
	"github.com/lurkshark/fireweed/packages/project/internal/inspection"
	"github.com/lurkshark/fireweed/packages/project/internal/playbooks"
	"github.com/lurkshark/fireweed/packages/project/internal/projection"
	"github.com/lurkshark/fireweed/packages/project/internal/queuecheck"
	"github.com/lurkshark/fireweed/packages/project/model"
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

func ApplyChange(module contracts.Module, change contracts.Change) (contracts.Module, error) {
	return projection.Apply(module, change)
}
func Project(snapshot Snapshot, before string) (Snapshot, error) {
	result, err := projection.Project(internalSnapshot(snapshot), before)
	return Snapshot{Root: result.Root, BaseRef: result.BaseRef, Config: result.Config, Modules: result.Modules, Changes: result.Changes, Archive: result.Archive, Diagnostics: publicDiagnostics(result.Diagnostics)}, err
}
func BaseHashes(snapshot Snapshot, slug string) (map[string]string, error) {
	return projection.Bases(internalSnapshot(snapshot), slug)
}
func internalSnapshot(s Snapshot) inspection.Data {
	d := inspection.Data{Root: s.Root, BaseRef: s.BaseRef, Config: s.Config, Modules: s.Modules, Changes: s.Changes, Archive: s.Archive}
	for _, v := range s.Diagnostics {
		d.Diagnostics = append(d.Diagnostics, inspection.Diagnostic{Code: v.Code, Severity: v.Severity, Message: v.Message, Path: v.Path})
	}
	return d
}

func CheckQueue(snapshot Snapshot) []Diagnostic {
	return publicDiagnostics(queuecheck.Check(internalSnapshot(snapshot)))
}

func Allocate(snapshot Snapshot, slug string) (contracts.Change, map[string]string, error) {
	return allocation.Allocate(internalSnapshot(snapshot), slug)
}

func Finalize(module contracts.Module, change contracts.Change) (contracts.Module, error) {
	return finalization.Finalize(module, change)
}

type Playbook = model.Playbook

func ResolvePlaybook(start, name string) (Playbook, error) { return playbooks.Resolve(start, name) }

func LoadAt(start, revision string) (Snapshot, error) {
	d, err := inspection.LoadAt(start, revision)
	if err != nil {
		return Snapshot{}, err
	}
	return Snapshot{Root: d.Root, BaseRef: d.BaseRef, Config: d.Config, Modules: d.Modules, Changes: d.Changes, Archive: d.Archive, Diagnostics: publicDiagnostics(d.Diagnostics)}, nil
}
