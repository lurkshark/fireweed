package claims

import "strings"

// Inventory only local object identities. Older Git versions can ignore
// GIT_NO_LAZY_FETCH, so content reads need an explicit availability check.
// This neither opens unrelated claim payloads nor requires their completeness.
func (s *storage) available(object string) error {
	if s.objects == nil {
		raw, err := s.git("cat-file", "--batch-all-objects", "--batch-check=%(objectname)")
		if err != nil {
			return err
		}
		objects := make(map[string]bool)
		for _, id := range strings.Fields(raw) {
			if !s.oid(id) {
				return failure("local object inventory contains an invalid object ID")
			}
			objects[id] = true
		}
		s.objects = objects
	}
	if !s.objects[object] {
		return failure("required object %s is unavailable locally; restore the local repository objects", object)
	}
	return nil
}

// rev-list accepts sets, while a claim baseline names a single commit. Let Git
// interpret path/search/reflog expressions, but reject its set operators outside
// those expressions before a range with only one result could be accepted.
func revisionSet(value string) bool {
	for i := 0; i < len(value); i++ {
		if value[i] == ':' {
			return false
		}
		if i+1 < len(value) && (value[i] == '^' || value[i] == '@') && value[i+1] == '{' {
			end := strings.IndexByte(value[i+2:], '}')
			if end < 0 {
				return false
			}
			i += end + 2
			continue
		}
		if strings.HasPrefix(value[i:], "..") || strings.HasPrefix(value[i:], "^@") || strings.HasPrefix(value[i:], "^!") || strings.HasPrefix(value[i:], "^-") {
			return true
		}
	}
	return false
}
