package claims

import "github.com/lurkshark/fireweed/packages/workflow/model"

func Create(start, slug, run, base, state string) (model.Claim, error) { return model.Claim{}, nil }
func Load(start, slug string) (model.Claim, error)                     { return model.Claim{}, nil }
func List(start string) ([]model.Claim, error)                         { return nil, nil }
func Update(start string, previous model.Claim, base, state string) (model.Claim, error) {
	return model.Claim{}, nil
}
func Release(start, slug, expected string) error { return nil }
