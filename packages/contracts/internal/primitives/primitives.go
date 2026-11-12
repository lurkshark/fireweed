package primitives

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/lurkshark/fireweed/packages/contracts/model"
	"gopkg.in/yaml.v3"
)

func ParseDocument(source string) (model.Document, error) {
	first, next := lineAt(source, 0)
	if first != "---" || next == len(source) && !strings.HasSuffix(source, "\n") {
		return model.Document{}, model.CodedError{Code: "C03", Cause: errors.New("missing YAML frontmatter")}
	}
	for pos := next; pos < len(source); {
		line, after := lineAt(source, pos)
		if line == "---" {
			var frontmatter map[string]any
			if err := yaml.Unmarshal([]byte(source[next:pos]), &frontmatter); err != nil {
				return model.Document{}, model.CodedError{Code: "C03", Cause: fmt.Errorf("invalid YAML frontmatter: %w", err)}
			}
			if frontmatter == nil {
				return model.Document{}, model.CodedError{Code: "C03", Cause: errors.New("YAML frontmatter must be a mapping")}
			}
			return model.Document{Frontmatter: frontmatter, Body: source[after:]}, nil
		}
		pos = after
	}
	return model.Document{}, model.CodedError{Code: "C03", Cause: errors.New("unterminated YAML frontmatter")}
}

func RenderDocument(doc model.Document) (string, error) {
	frontmatter := doc.Frontmatter
	if frontmatter == nil {
		return "", errors.New("render YAML frontmatter: nil mapping")
	}
	data, err := yaml.Marshal(frontmatter)
	if err != nil {
		return "", fmt.Errorf("render YAML frontmatter: %w", err)
	}
	return "---\n" + string(data) + "---\n" + doc.Body, nil
}

// lineAt returns a line without its line ending and the offset of the next line.
func lineAt(source string, start int) (string, int) {
	end := strings.IndexByte(source[start:], '\n')
	if end < 0 {
		return strings.TrimSuffix(source[start:], "\r"), len(source)
	}
	end += start
	return strings.TrimSuffix(source[start:end], "\r"), end + 1
}

// sectionBounds returns a negative start for an absent section, along with any
// fence left open at the end of the document.
func sectionBounds(markdown, name string) (int, int, string) {
	start := -1
	var openFence string
	for pos := 0; pos < len(markdown); {
		line, next := lineAt(markdown, pos)
		lineStart := pos
		pos = next

		content := strings.TrimLeft(line, " ")
		if len(line)-len(content) > 3 {
			continue
		}
		marker, rest := fenceMarker(content)
		if openFence != "" {
			if strings.HasPrefix(marker, openFence) && strings.Trim(rest, " \t") == "" {
				openFence = ""
			}
			continue
		}
		if marker != "" && (marker[0] == '~' || !strings.ContainsRune(rest, '`')) {
			openFence = marker
			continue
		}
		if level, heading := sectionHeading(content); level != 0 {
			if start >= 0 {
				return start, lineStart, ""
			}
			if level == 2 && heading == name {
				start = next
			}
		}
	}
	return start, len(markdown), openFence
}

func fenceMarker(line string) (marker, rest string) {
	if line == "" || (line[0] != '`' && line[0] != '~') {
		return "", line
	}
	rest = strings.TrimLeft(line, line[:1])
	length := len(line) - len(rest)
	if length < 3 {
		return "", line
	}
	return line[:length], rest
}

func sectionHeading(line string) (level int, name string) {
	rest := strings.TrimLeft(line, "#")
	level = len(line) - len(rest)
	if level < 1 || level > 2 || (rest != "" && rest[0] != ' ' && rest[0] != '\t') {
		return 0, ""
	}
	name = strings.TrimRight(rest, " \t")
	beforeClosing := strings.TrimRight(name, "#")
	if strings.HasSuffix(beforeClosing, " ") || strings.HasSuffix(beforeClosing, "\t") {
		name = beforeClosing
	}
	return level, strings.Trim(name, " \t")
}

func Section(markdown, name string) (string, bool) {
	start, end, _ := sectionBounds(markdown, name)
	if start < 0 {
		return "", false
	}
	return markdown[start:end], true
}

func ReplaceSection(markdown, name, body string) string {
	start, end, openFence := sectionBounds(markdown, name)
	if start >= 0 {
		return markdown[:start] + body + markdown[end:]
	}
	if markdown != "" && !strings.HasSuffix(markdown, "\n") {
		markdown += "\n"
	}
	if openFence != "" {
		markdown += openFence + "\n"
	}
	return markdown + "## " + name + "\n" + body
}

var itemHeading = regexp.MustCompile("^### ([A-Z0-9]+-[0-9]+|new:[a-z0-9-]+)(?: \\[([^\\[\\]]+)\\])?$")

func ParseItems(section string) ([]model.Item, error) {
	var items []model.Item
	seen := make(map[string]bool)
	bodyStart := 0
	for pos := 0; pos < len(section); {
		line, next := lineAt(section, pos)
		if strings.HasPrefix(line, "###") && !strings.HasPrefix(line, "####") {
			parts := itemHeading.FindStringSubmatch(line)
			if parts == nil {
				return nil, model.CodedError{Code: "C03", Subject: line, Cause: fmt.Errorf("malformed item heading: %q", line)}
			}
			if seen[parts[1]] {
				return nil, model.CodedError{Code: "C05", Subject: parts[1], Cause: fmt.Errorf("duplicate item ID: %s", parts[1])}
			}
			seen[parts[1]] = true
			if len(items) > 0 {
				items[len(items)-1].Body = section[bodyStart:pos]
			}
			items = append(items, model.Item{ID: parts[1], Tag: parts[2]})
			bodyStart = next
		}
		pos = next
	}
	if len(items) > 0 {
		items[len(items)-1].Body = section[bodyStart:]
	}
	return items, nil
}

func RenderItem(item model.Item) string {
	heading := "### " + item.ID
	if item.Tag != "" {
		heading += " [" + item.Tag + "]"
	}
	return heading + "\n" + item.Body
}

func SpecHash(item model.Item) string {
	content := strings.ReplaceAll(RenderItem(item), "\r\n", "\n")
	lines := strings.Split(content, "\n")
	for i := range lines {
		lines[i] = strings.TrimRight(lines[i], " \t")
	}
	canonical := strings.TrimRight(strings.Join(lines, "\n"), "\n") + "\n"
	sum := sha256.Sum256([]byte(canonical))
	return hex.EncodeToString(sum[:4])
}

var specID = regexp.MustCompile("^([A-Za-z][A-Za-z0-9]*)[-_]?([0-9]+)$")

func MatchesID(id, testcaseName, classname string) bool {
	parts := specID.FindStringSubmatch(id)
	if parts == nil {
		return false
	}
	pattern := regexp.MustCompile("(?i)" + regexp.QuoteMeta(parts[1]) + "[-_]?" + regexp.QuoteMeta(parts[2]))
	for _, text := range []string{testcaseName, classname} {
		for _, match := range pattern.FindAllStringIndex(text, -1) {
			if match[0] > 0 && asciiLetterOrDigit(text[match[0]-1]) {
				continue
			}
			if match[1] < len(text) && text[match[1]] >= '0' && text[match[1]] <= '9' {
				continue
			}
			return true
		}
	}
	return false
}

func asciiLetterOrDigit(ch byte) bool {
	return ch >= 'A' && ch <= 'Z' || ch >= 'a' && ch <= 'z' || ch >= '0' && ch <= '9'
}

func MatchGlob(pattern, path string) bool {
	parts := strings.Split(pattern, "/")
	pathParts := strings.Split(path, "/")
	var match func(int, int) bool
	match = func(pi, si int) bool {
		if pi == len(parts) {
			return si == len(pathParts)
		}
		if parts[pi] == "**" {
			return match(pi+1, si) || si < len(pathParts) && match(pi, si+1)
		}
		return si < len(pathParts) && matchComponent(parts[pi], pathParts[si]) && match(pi+1, si+1)
	}
	return match(0, 0)
}

func matchComponent(pattern, name string) bool {
	pi, ni := 0, 0
	star, resume := -1, 0
	for ni < len(name) {
		if pi < len(pattern) && pattern[pi] == name[ni] && pattern[pi] != '*' {
			pi++
			ni++
		} else if pi < len(pattern) && pattern[pi] == '*' {
			star = pi
			pi++
			resume = ni
		} else if star >= 0 {
			pi = star + 1
			resume++
			ni = resume
		} else {
			return false
		}
	}
	for pi < len(pattern) && pattern[pi] == '*' {
		pi++
	}
	return pi == len(pattern)
}

func ResolveZone(path string, zones model.Zones) string {
	if matchesAny(zones.Protected, path) {
		return "protected"
	}
	if path == "SPEC.md" {
		return "spec"
	}
	for _, zone := range []struct {
		name     string
		patterns []string
	}{
		{"evals", zones.Evals},
		{"interface", zones.Interface},
		{"append_only", zones.AppendOnly},
		{"preserve", zones.Preserve},
		{"disposable", zones.Disposable},
	} {
		if matchesAny(zone.patterns, path) {
			return zone.name
		}
	}
	return "other"
}

func matchesAny(patterns []string, path string) bool {
	for _, pattern := range patterns {
		if MatchGlob(pattern, path) {
			return true
		}
	}
	return false
}
