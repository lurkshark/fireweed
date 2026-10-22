// Package contracts exposes the language-independent fireweed document contract.
package contracts

import "github.com/lurkshark/fireweed/packages/contracts/internal/specs"

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
