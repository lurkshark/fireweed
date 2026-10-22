package repository

import "errors"

type Edit struct{ Status, Path, OldPath string }

func Open(start string) (string, string, string, error) {
	return "", "", "", errors.New("not implemented")
}
func Base(root, base, remote string) (string, error) { return "", errors.New("not implemented") }
func Clean(root string) (bool, error)                { return false, errors.New("not implemented") }
func Edits(root, base string) ([]Edit, error)        { return nil, errors.New("not implemented") }
func SafePath(root, relative string) (string, error) { return "", errors.New("not implemented") }
