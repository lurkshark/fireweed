package queue

import "github.com/lurkshark/fireweed/packages/project/model"

func Queue(root string) ([]model.Change, error) { return nil, nil }

func Project(root, through string) (model.Snapshot, error) { return model.Snapshot{}, nil }

func Stamp(root, change string) error { return nil }

func Check(root string) []model.Diagnostic { return nil }
