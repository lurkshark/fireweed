package queue

import (
	"maps"
	"path"
	"slices"
	"strings"
	"unicode"

	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/project/internal/source"
	"github.com/lurkshark/fireweed/packages/project/model"
)

func (check *staticCheck) outsideSpecs(snapshot model.Snapshot) {
	files, err := source.Git(snapshot.Root, "ls-files", "-z", "--full-name")
	if err != nil {
		return
	}
	seen := make(map[string]bool)
	for _, file := range strings.Split(files, "\x00") {
		if seen[file] || path.Base(file) != "SPEC.md" || strings.HasPrefix(file, ".fireweed/") {
			continue
		}
		seen[file] = true
		if !slices.ContainsFunc(snapshot.Modules, func(module model.Module) bool {
			return module.Dir == "." || strings.HasPrefix(file, module.Dir+"/")
		}) {
			check.warn("C20", file, file+" is outside the configured module directories")
		}
	}
}

func (check *staticCheck) repeatedSentences(modules [][]contracts.Item) {
	matches := make(map[string][]string)
	for _, items := range modules {
		for _, item := range items {
			seen := make(map[string]bool)
			for _, sentence := range strings.FieldsFunc(item.Body, func(r rune) bool {
				return r == '.' || r == '?' || r == '!'
			}) {
				key := strings.Map(func(r rune) rune {
					if unicode.IsPunct(r) || unicode.IsSpace(r) {
						return -1
					}
					return unicode.ToLower(r)
				}, sentence)
				if key != "" && !seen[key] {
					matches[key] = append(matches[key], item.ID)
					seen[key] = true
				}
			}
		}
	}
	for _, key := range slices.Sorted(maps.Keys(matches)) {
		if ids := matches[key]; len(ids) >= 3 {
			slices.Sort(ids)
			check.warn("C18", "", "repeated sentence in "+strings.Join(ids, ", "))
		}
	}
}
