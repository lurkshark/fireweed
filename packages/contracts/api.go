// Package contracts exposes the language-independent fireweed document contract.
package contracts

import (
	"github.com/lurkshark/fireweed/packages/contracts/internal/config"
	"github.com/lurkshark/fireweed/packages/contracts/internal/globs"
	"github.com/lurkshark/fireweed/packages/contracts/internal/specs"
	"github.com/lurkshark/fireweed/packages/contracts/model"
)

// Spec is a normative Markdown item. Tag retains bracketed verification/enforcement tags.
type Spec struct{ ID, Tag, Body string }

func ParseSpecs(document, prefix string) ([]Spec, error) {
	items, err := specs.ParseDocument(document, prefix)
	return publicSpecs(items), err
}
func ParseItems(section, prefix string, allowHandles bool) ([]Spec, error) {
	items, err := specs.ParseItems(section, prefix, allowHandles)
	return publicSpecs(items), err
}
func publicSpecs(items []specs.Item) []Spec {
	result := make([]Spec, len(items))
	for i, item := range items {
		result[i] = Spec{item.ID, item.Tag, item.Body}
	}
	return result
}
func SpecHash(item Spec) string                 { return specs.Hash(item.ID, item.Tag, item.Body) }
func MatchesID(name, classname, id string) bool { return specs.Match(name, classname, id) }

// Zones declares module-relative file patterns in strictness order.
type Zones = model.Zones
type Zone string

const (
	Protected      Zone = "protected"
	SpecZone       Zone = "spec"
	EvalZone       Zone = "evals"
	InterfaceZone  Zone = "interface"
	AppendOnlyZone Zone = "append-only"
	PreserveZone   Zone = "preserved"
	DisposableZone Zone = "disposable"
	OtherZone      Zone = "other"
	OutsideZone    Zone = "outside"
)

func MatchGlob(pattern, path string) (bool, error) { return globs.Match(pattern, path) }
func GlobsOverlap(a, b string) (bool, error)       { return globs.Overlap(a, b) }
func ZoneFor(path, moduleDir string, zones Zones, protected []string) (Zone, error) {
	zone, err := globs.Zone(path, moduleDir, [5][]string{zones.Evals, zones.Interface, zones.AppendOnly, zones.Preserve, zones.Disposable}, protected)
	return Zone(zone), err
}

type Config = model.Config
type Hooks = model.Hooks
type Agent = model.Agent
type Audit = model.Audit

func ParseConfig(text, binaryVersion string) (Config, error) {
	return config.Parse(text, binaryVersion)
}
func CheckVersion(binaryVersion, constraint string) error {
	return config.Version(binaryVersion, constraint)
}
