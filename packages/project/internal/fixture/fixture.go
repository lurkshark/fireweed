package fixture

import "github.com/lurkshark/fireweed/packages/project/model"

func Load(root string) (model.Snapshot, error) { return model.Snapshot{}, nil }

func ValidateGraph(modules []model.Module) error { return nil }

func NewChange(root, module, slug string) (string, error) { return "", nil }
