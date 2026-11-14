package queue

import (
	"fmt"
	"strings"

	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/project/model"
)

func (check *staticCheck) zoneOverlap(file string, module model.Module) {
	for _, eval := range module.Zones.Evals {
		for _, disposable := range module.Zones.Disposable {
			if globsOverlap(eval, disposable) {
				check.error("C12", file, fmt.Sprintf("module %q has overlapping evals glob %q and disposable glob %q", module.Name, eval, disposable))
			}
		}
	}
}

func globsOverlap(left, right string) bool {
	a, b := strings.Split(left, "/"), strings.Split(right, "/")
	type state struct {
		i, j     int
		nonempty bool
	}
	seen := make(map[state]bool)
	var visit func(state) bool
	visit = func(s state) bool {
		if seen[s] {
			return false
		}
		seen[s] = true
		if s.i == len(a) && s.j == len(b) {
			return s.nonempty
		}
		// A recursive wildcard can skip itself or consume a component and stay.
		recursiveA := s.i < len(a) && a[s.i] == "**"
		recursiveB := s.j < len(b) && b[s.j] == "**"
		if recursiveA && visit(state{s.i + 1, s.j, s.nonempty}) {
			return true
		}
		if recursiveB && visit(state{s.i, s.j + 1, s.nonempty}) {
			return true
		}
		if s.i == len(a) || s.j == len(b) || !componentsOverlap(a[s.i], b[s.j]) {
			return false
		}
		next := state{s.i, s.j, true}
		if !recursiveA {
			next.i++
		}
		if !recursiveB {
			next.j++
		}
		return visit(next)
	}
	return visit(state{})
}

func componentsOverlap(a, b string) bool {
	if strings.ContainsRune(a, 0) || strings.ContainsRune(b, 0) {
		return false
	}
	prefixA, _, starA := strings.Cut(a, "*")
	prefixB, _, starB := strings.Cut(b, "*")
	if !starA {
		return matchesFileComponent(b, a)
	}
	if !starB {
		return matchesFileComponent(a, b)
	}
	// With a star on both sides, a shared name can use the longer prefix
	// and suffix and fit both patterns' interior literals between them.
	suffixA := a[strings.LastIndexByte(a, '*')+1:]
	suffixB := b[strings.LastIndexByte(b, '*')+1:]
	return (strings.HasPrefix(prefixA, prefixB) || strings.HasPrefix(prefixB, prefixA)) &&
		(strings.HasSuffix(suffixA, suffixB) || strings.HasSuffix(suffixB, suffixA))
}

func matchesFileComponent(pattern, name string) bool {
	return name != "" && name != "." && name != ".." && contracts.MatchGlob(pattern, name)
}
