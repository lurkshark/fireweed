package globs

import "errors"

func Match(pattern, path string) (bool, error) { return false, errors.New("not implemented") }
func Overlap(a, b string) (bool, error)        { return false, errors.New("not implemented") }
func Zone(path, moduleDir string, zones [5][]string, protected []string) (string, error) {
	return "", errors.New("not implemented")
}
