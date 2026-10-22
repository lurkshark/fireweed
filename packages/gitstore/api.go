// Package gitstore provides local Git and repository path primitives.
package gitstore

import "github.com/lurkshark/fireweed/packages/gitstore/internal/repository"

type Repository struct{ Root, GitDir, CommonDir string }
type Edit struct{ Status, Path, OldPath string }

func Open(start string) (Repository, error) {
	root, git, common, err := repository.Open(start)
	return Repository{root, git, common}, err
}
func BaseRef(root, base, remote string) (string, error) { return repository.Base(root, base, remote) }
func Clean(root string) (bool, error)                   { return repository.Clean(root) }
func Edits(root, base string) ([]Edit, error) {
	entries, err := repository.Edits(root, base)
	result := make([]Edit, len(entries))
	for i, e := range entries {
		result[i] = Edit{e.Status, e.Path, e.OldPath}
	}
	return result, err
}
func SafePath(root, relative string) (string, error) { return repository.SafePath(root, relative) }
