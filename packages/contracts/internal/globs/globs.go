// Package globs compiles path globs to epsilon automata. Intersection walks
// their product, so validation covers possible paths rather than tracked files.
package globs

import (
	"fmt"
	"sort"
	"strings"
	"unicode/utf8"
)

type interval struct{ lo, hi rune }
type edge struct {
	to    int
	chars []interval
}
type state struct {
	eps   []int
	edges []edge
}
type machine []state

var anyRune = []interval{{1, utf8.MaxRune}}
var segmentRune = []interval{{1, '/' - 1}, {'/' + 1, utf8.MaxRune}}

func literal(r rune) []interval { return []interval{{r, r}} }

func validatePath(path string) error {
	if path == "" || strings.HasPrefix(path, "/") || strings.ContainsAny(path, "\\\x00") || !utf8.ValidString(path) {
		return fmt.Errorf("invalid relative path %q", path)
	}
	if len(path) > 1 && path[1] == ':' && ((path[0] >= 'A' && path[0] <= 'Z') || (path[0] >= 'a' && path[0] <= 'z')) {
		return fmt.Errorf("absolute drive path %q", path)
	}
	for _, part := range strings.Split(path, "/") {
		if part == "" || part == "." || part == ".." {
			return fmt.Errorf("invalid path component in %q", path)
		}
	}
	return nil
}
func (m *machine) node() int       { *m = append(*m, state{}); return len(*m) - 1 }
func (m machine) epsilon(a, b int) { m[a].eps = append(m[a].eps, b) }
func (m machine) transition(a, b int, ranges []interval) {
	m[a].edges = append(m[a].edges, edge{b, ranges})
}

func compile(pattern string) (machine, error) {
	if err := validatePath(pattern); err != nil {
		return nil, err
	}
	m := machine{{}}
	current := 0
	text := []rune(pattern)
	for i := 0; i < len(text); i++ {
		r := text[i]
		next := m.node()
		switch r {
		case '*':
			end := i + 1
			for end < len(text) && text[end] == '*' {
				end++
			}
			globstar := end-i == 2 && (i == 0 || text[i-1] == '/') && (end == len(text) || text[end] == '/')
			if globstar && end < len(text) {
				loop := m.node()
				m.epsilon(current, next)
				m.epsilon(current, loop)
				m.transition(loop, loop, anyRune)
				m.transition(loop, next, literal('/'))
				i = end // consume the slash, with a zero-directory alternative
			} else {
				chars := segmentRune
				if globstar {
					chars = anyRune
				}
				m.epsilon(current, next)
				m.transition(current, current, chars)
				i = end - 1
			}
		case '?':
			m.transition(current, next, segmentRune)
		case '[':
			chars, last, err := characterClass(text, i)
			if err != nil {
				return nil, err
			}
			m.transition(current, next, chars)
			i = last
		default:
			m.transition(current, next, literal(r))
		}
		current = next
	}
	// Globstar loops may allocate a later node than the accepting state.
	accept := m.node()
	m.epsilon(current, accept)
	return m, nil
}

func characterClass(text []rune, start int) ([]interval, int, error) {
	end := start + 1
	for end < len(text) && text[end] != ']' {
		end++
	}
	if end == len(text) {
		return nil, 0, fmt.Errorf("unclosed character class")
	}
	pos := start + 1
	negated := false
	if pos < end && (text[pos] == '!' || text[pos] == '^') {
		negated = true
		pos++
	}
	if pos == end {
		return nil, 0, fmt.Errorf("empty character class")
	}
	ranges := []interval{}
	for pos < end {
		low, high := text[pos], text[pos]
		pos++
		if pos+1 < end && text[pos] == '-' {
			high = text[pos+1]
			pos += 2
		}
		if high < low {
			return nil, 0, fmt.Errorf("reversed character range")
		}
		ranges = append(ranges, interval{low, high})
	}
	sort.Slice(ranges, func(i, j int) bool { return ranges[i].lo < ranges[j].lo })
	if negated {
		inverse := []interval{}
		low := rune(1)
		for _, part := range ranges {
			if part.lo > low {
				inverse = append(inverse, interval{low, part.lo - 1})
			}
			if part.hi >= low {
				low = part.hi + 1
			}
		}
		if low <= utf8.MaxRune {
			inverse = append(inverse, interval{low, utf8.MaxRune})
		}
		ranges = inverse
	}
	return intersect(ranges, segmentRune), end, nil
}
func intersect(a, b []interval) []interval {
	result := []interval{}
	for _, x := range a {
		for _, y := range b {
			low, high := max(x.lo, y.lo), min(x.hi, y.hi)
			if low <= high {
				result = append(result, interval{low, high})
			}
		}
	}
	return result
}
func contains(chars []interval, r rune) bool {
	for _, part := range chars {
		if r >= part.lo && r <= part.hi {
			return true
		}
	}
	return false
}
func closure(m machine, initial map[int]bool) map[int]bool {
	queue := []int{}
	for i := range initial {
		queue = append(queue, i)
	}
	for len(queue) > 0 {
		i := queue[0]
		queue = queue[1:]
		for _, j := range m[i].eps {
			if !initial[j] {
				initial[j] = true
				queue = append(queue, j)
			}
		}
	}
	return initial
}
func Match(pattern, path string) (bool, error) {
	m, err := compile(pattern)
	if err != nil {
		return false, err
	}
	if err = validatePath(path); err != nil {
		return false, err
	}
	active := closure(m, map[int]bool{0: true})
	for _, r := range path {
		next := map[int]bool{}
		for i := range active {
			for _, e := range m[i].edges {
				if contains(e.chars, r) {
					next[e.to] = true
				}
			}
		}
		active = closure(m, next)
	}
	return active[len(m)-1], nil
}

// The path state rejects empty, dot, dot-dot and repeated-slash components.
// All other runes behave alike for this state; slash and dot are distinct.
func pathNext(s, category int) int {
	if category == 0 {
		if s == 3 {
			return 0
		}
		return -1
	} // slash
	if category == 1 {
		if s == 0 {
			return 1
		}
		if s == 1 {
			return 2
		}
		return 3
	} // dot
	return 3
}

var categories = [][]interval{literal('/'), literal('.'), {{1, '.' - 1}, {'0', '\\' - 1}, {'\\' + 1, utf8.MaxRune}}}

func Overlap(a, b string) (bool, error) {
	left, err := compile(a)
	if err != nil {
		return false, err
	}
	right, err := compile(b)
	if err != nil {
		return false, err
	}
	type pair struct{ x, y, path int }
	queue := []pair{{0, 0, 0}}
	seen := map[pair]bool{{0, 0, 0}: true}
	add := func(p pair) {
		if !seen[p] {
			seen[p] = true
			queue = append(queue, p)
		}
	}
	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]
		if p.x == len(left)-1 && p.y == len(right)-1 && p.path == 3 {
			return true, nil
		}
		for _, i := range left[p.x].eps {
			add(pair{i, p.y, p.path})
		}
		for _, j := range right[p.y].eps {
			add(pair{p.x, j, p.path})
		}
		for _, x := range left[p.x].edges {
			for _, y := range right[p.y].edges {
				common := intersect(x.chars, y.chars)
				for category, chars := range categories {
					next := pathNext(p.path, category)
					if next >= 0 && len(intersect(common, chars)) > 0 {
						add(pair{x.to, y.to, next})
					}
				}
			}
		}
	}
	return false, nil
}

func Zone(path, moduleDir string, zones [5][]string, protected []string) (string, error) {
	if err := validatePath(path); err != nil {
		return "", err
	}
	if err := validatePath(moduleDir); err != nil {
		return "", err
	}
	patterns := append([]string{}, protected...)
	for _, group := range zones {
		patterns = append(patterns, group...)
	}
	for _, pattern := range patterns {
		if _, err := compile(pattern); err != nil {
			return "", err
		}
	}
	base := path[strings.LastIndex(path, "/")+1:]
	if path == "fireweed.yaml" || path == "SYSTEM.md" || strings.HasPrefix(path, ".fireweed/") || base == "AGENTS.md" || base == "CLAUDE.md" {
		return "protected", nil
	}
	for _, pattern := range protected {
		if matched, _ := Match(pattern, path); matched {
			return "protected", nil
		}
	}
	if !strings.HasPrefix(path, moduleDir+"/") {
		return "outside", nil
	}
	relative := strings.TrimPrefix(path, moduleDir+"/")
	if base == "SPEC.md" {
		return "spec", nil
	}
	names := []string{"evals", "interface", "append-only", "preserved", "disposable"}
	for i, group := range zones {
		for _, pattern := range group {
			if matched, _ := Match(pattern, relative); matched {
				return names[i], nil
			}
		}
	}
	return "other", nil
}
