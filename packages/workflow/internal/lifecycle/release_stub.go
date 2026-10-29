package lifecycle

import (
	"fmt"

	"github.com/lurkshark/fireweed/packages/gitstore"
)

func Release(root, change string) error {
	repo, err := gitstore.Inspect(root)
	if err != nil {
		return err
	}
	current, err := Current(repo.Root, change)
	if err != nil {
		return err
	}
	if err := gitstore.CASRef(repo.Root, ref(current.Change.Slug), current.ClaimOID, ""); err != nil {
		return fmt.Errorf("P02: release claim: %w", err)
	}
	return nil
}
