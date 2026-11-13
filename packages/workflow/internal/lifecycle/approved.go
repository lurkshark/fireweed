package lifecycle

import "github.com/lurkshark/fireweed/packages/workflow/model"

func Approved(root string) (model.ApprovedSnapshot, bool, error) {
	return model.ApprovedSnapshot{}, false, nil
}
