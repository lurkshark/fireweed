package primitives

import (
	"errors"

	"github.com/lurkshark/fireweed/packages/contracts/model"
)

var errNotImplemented = errors.New("not implemented")

func ParseDocument(string) (model.Document, error)  { return model.Document{}, errNotImplemented }
func RenderDocument(model.Document) (string, error) { return "", errNotImplemented }
func Section(string, string) (string, bool)         { return "", false }
func ReplaceSection(markdown, _, _ string) string   { return markdown }
func ParseItems(string) ([]model.Item, error)       { return nil, errNotImplemented }
func RenderItem(model.Item) string                  { return "" }
func SpecHash(model.Item) string                    { return "" }
func MatchesID(string, string, string) bool         { return false }
func MatchGlob(string, string) bool                 { return false }
func ResolveZone(string, model.Zones) string        { return "other" }
