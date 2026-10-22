package specs

import "errors"

type Item struct{ ID, Tag, Body string }

func ParseDocument(document, prefix string) ([]Item, error) {
	return nil, errors.New("not implemented")
}
func ParseItems(section, prefix string, allowHandles bool) ([]Item, error) {
	return nil, errors.New("not implemented")
}
func Hash(id, tag, body string) string      { return "" }
func Match(name, classname, id string) bool { return false }
