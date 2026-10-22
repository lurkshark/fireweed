package config

import (
	"errors"
	"github.com/lurkshark/fireweed/packages/contracts/model"
)

func Parse(text, binaryVersion string) (model.Config, error) {
	return model.Config{}, errors.New("not implemented")
}
func Version(binaryVersion, constraint string) error { return errors.New("not implemented") }
