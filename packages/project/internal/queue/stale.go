package queue

import "github.com/lurkshark/fireweed/packages/project/model"

// StaleSpecs is implemented after the evals phase.
func StaleSpecs(root, change string) ([]model.StaleSpec, error) { return nil, nil }

// StampMissing is implemented after the evals phase.
func StampMissing(root, change string) error { return nil }
