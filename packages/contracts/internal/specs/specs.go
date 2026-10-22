package specs

import (
	"crypto/sha256"
	"fmt"
	"regexp"
	"strings"
	"unicode"
	"unicode/utf8"
)

type Item struct{ ID, Tag, Body string }

var identifier = regexp.MustCompile(`^([A-Za-z][A-Za-z0-9]*)-([0-9]{3,})$`)
var handle = regexp.MustCompile(`^new:[a-z0-9-]+$`)
var normative = regexp.MustCompile(`\b(MUST|SHOULD)\b`)

// A scanner retains fenced examples while exposing only live headings.
type markdown struct {
	fence byte
	width int
}

func (m *markdown) live(line string) bool {
	trimmed := strings.TrimLeft(line, " ")
	if len(line)-len(trimmed) > 3 {
		return m.fence == 0
	}
	if len(trimmed) < 3 || (trimmed[0] != '`' && trimmed[0] != '~') {
		return m.fence == 0
	}
	width := 0
	for width < len(trimmed) && trimmed[width] == trimmed[0] {
		width++
	}
	if width < 3 {
		return m.fence == 0
	}
	if m.fence == 0 {
		m.fence, m.width = trimmed[0], width
		return false
	}
	if trimmed[0] == m.fence && width >= m.width && strings.TrimSpace(trimmed[width:]) == "" {
		m.fence, m.width = 0, 0
	}
	return false
}
func lines(text string) []string { return strings.Split(strings.ReplaceAll(text, "\r\n", "\n"), "\n") }

func ParseDocument(document, prefix string) ([]Item, error) {
	var scanner markdown
	count, inSpecs := 0, false
	var body []string
	for _, line := range lines(document) {
		if scanner.live(line) && strings.HasPrefix(line, "## ") {
			inSpecs = strings.TrimSpace(strings.TrimPrefix(line, "## ")) == "Specs"
			if inSpecs {
				count++
			}
			continue
		}
		if inSpecs {
			body = append(body, line)
		}
	}
	if count != 1 {
		return nil, fmt.Errorf("expected exactly one Specs section, found %d", count)
	}
	return ParseItems(strings.Join(body, "\n"), prefix, false)
}

func ParseItems(section, prefix string, allowHandles bool) ([]Item, error) {
	items := []Item{}
	seen := map[string]bool{}
	var body []string
	var scanner markdown
	finish := func() error {
		if len(items) == 0 {
			return nil
		}
		item := &items[len(items)-1]
		item.Body = strings.TrimSpace(strings.Join(body, "\n"))
		if item.Body == "" || !normative.MatchString(item.Body) {
			return fmt.Errorf("%s requires a normative MUST or SHOULD body", item.ID)
		}
		return nil
	}
	for _, line := range lines(section) {
		if scanner.live(line) && (line == "###" || strings.HasPrefix(line, "### ")) {
			if err := finish(); err != nil {
				return nil, err
			}
			heading := strings.TrimSpace(strings.TrimPrefix(line, "###"))
			fields := strings.SplitN(heading, " ", 2)
			id := fields[0]
			match := identifier.FindStringSubmatch(id)
			validID := len(match) > 0 && match[1] == prefix
			if !validID && !(allowHandles && handle.MatchString(id)) {
				return nil, fmt.Errorf("invalid spec heading %q", heading)
			}
			if seen[id] {
				return nil, fmt.Errorf("duplicate spec %s", id)
			}
			seen[id] = true
			tag := ""
			if len(fields) > 1 {
				tag = strings.TrimSpace(fields[1])
			}
			if err := validateTags(tag); err != nil {
				return nil, fmt.Errorf("%s: %w", id, err)
			}
			items = append(items, Item{ID: id, Tag: tag})
			body = nil
			continue
		}
		if len(items) > 0 {
			body = append(body, line)
		}
	}
	if err := finish(); err != nil {
		return nil, err
	}
	return items, nil
}

func validateTags(tags string) error {
	seen := map[string]bool{}
	remaining := tags
	for remaining != "" {
		if !strings.HasPrefix(remaining, "[") {
			return fmt.Errorf("malformed tags %q", tags)
		}
		end := strings.IndexByte(remaining, ']')
		if end < 0 {
			return fmt.Errorf("unclosed tag")
		}
		value := remaining[1:end]
		kind := value
		if strings.HasPrefix(value, "quarantine:") {
			kind = "quarantine"
			if strings.TrimSpace(strings.TrimPrefix(value, "quarantine:")) == "" {
				return fmt.Errorf("quarantine needs a reason")
			}
		} else if value != "review" {
			return fmt.Errorf("unknown tag %q", value)
		}
		if seen[kind] {
			return fmt.Errorf("repeated tag %s", kind)
		}
		seen[kind] = true
		remaining = strings.TrimSpace(remaining[end+1:])
	}
	return nil
}

func Hash(id, tag, body string) string {
	heading := "### " + id
	if strings.TrimSpace(tag) != "" {
		heading += " " + strings.TrimSpace(tag)
	}
	normalized := lines(body)
	for i := range normalized {
		normalized[i] = strings.TrimRightFunc(normalized[i], unicode.IsSpace)
	}
	start, end := 0, len(normalized)
	for start < end && normalized[start] == "" {
		start++
	}
	for end > start && normalized[end-1] == "" {
		end--
	}
	sum := sha256.Sum256([]byte(heading + "\n" + strings.Join(normalized[start:end], "\n")))
	return fmt.Sprintf("%x", sum)[:8]
}

func Match(name, classname, id string) bool {
	parts := identifier.FindStringSubmatch(id)
	if len(parts) == 0 {
		return false
	}
	pattern := regexp.MustCompile(`(?i)` + regexp.QuoteMeta(parts[1]) + `[-_]?` + regexp.QuoteMeta(parts[2]))
	for _, text := range []string{name, classname} {
		for _, loc := range pattern.FindAllStringIndex(text, -1) {
			if loc[0] > 0 {
				r, _ := utf8.DecodeLastRuneInString(text[:loc[0]])
				if unicode.IsLetter(r) || unicode.IsDigit(r) {
					continue
				}
			}
			if loc[1] < len(text) {
				r, _ := utf8.DecodeRuneInString(text[loc[1]:])
				if unicode.IsDigit(r) {
					continue
				}
			}
			return true
		}
	}
	return false
}
