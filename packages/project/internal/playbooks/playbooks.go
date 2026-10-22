package playbooks

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"strings"

	"github.com/lurkshark/fireweed/packages/assets"
	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/gitstore"
	"github.com/lurkshark/fireweed/packages/project/model"
)

func Resolve(start, name string) (model.Playbook, error) {
	embedded, err := assets.Playbook(name)
	if err != nil {
		return model.Playbook{}, fmt.Errorf("Choose a valid playbook (%s): %w", strings.Join(assets.Playbooks(), ", "), err)
	}
	repo, err := gitstore.Open(start)
	if err != nil {
		return model.Playbook{}, fmt.Errorf("Open an existing local Git repository before resolving project playbooks: %w", err)
	}
	result := model.Playbook{Name: name, Text: embedded.Text, Hash: embedded.Hash, EmbeddedHash: embedded.Hash}
	path := ".fireweed/playbooks/" + name + ".md"
	failure := func() (model.Playbook, error) {
		return model.Playbook{}, fmt.Errorf("D05: Restore %s as a readable contained regular override file, or remove it to use the embedded playbook", path)
	}
	full, err := gitstore.SafePath(repo.Root, path)
	if err != nil {
		return failure()
	}
	info, err := os.Stat(full)
	if os.IsNotExist(err) {
		return result, nil
	}
	if err != nil || !info.Mode().IsRegular() {
		return failure()
	}
	raw, err := os.ReadFile(full)
	if err != nil {
		return failure()
	}
	result.Text = string(raw)
	sum := sha256.Sum256(raw)
	result.Hash = hex.EncodeToString(sum[:])
	result.Overridden = true
	if document, err := contracts.ParseMarkdown(result.Text); err == nil {
		result.BasedOn, _ = document.Metadata["based_on"].(string)
	}
	result.Stale = result.BasedOn != result.EmbeddedHash
	return result, nil
}
