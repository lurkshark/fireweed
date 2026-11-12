package lifecycle

import "github.com/lurkshark/fireweed/packages/gitstore"

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
		return codedErrorf("P02", current.Change.Slug, "release claim: %w", err)
	}
	return nil
}
