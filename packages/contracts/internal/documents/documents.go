package documents

import (
	"errors"
	"github.com/lurkshark/fireweed/packages/contracts/model"
)

func Parse(text string) (model.Document, error) {
	return model.Document{}, errors.New("not implemented")
}
func Render(document model.Document) (string, error) { return "", errors.New("not implemented") }
func Module(text, dir string, defaults model.Zones, hooks model.Hooks) (model.Module, error) {
	return model.Module{}, errors.New("not implemented")
}
func Change(text, path, prefix string) (model.Change, error) {
	return model.Change{}, errors.New("not implemented")
}
