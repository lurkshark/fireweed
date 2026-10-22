package allocation

import (
	"fmt"
	"math/big"
	"regexp"
	"strings"

	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/project/internal/changeview"
	"github.com/lurkshark/fireweed/packages/project/internal/inspection"
)

var handle = regexp.MustCompile(`new:[a-z0-9-]+`)

func Allocate(s inspection.Data, slug string) (contracts.Change, map[string]string, error) {
	fail := func(err error) (contracts.Change, map[string]string, error) {
		return contracts.Change{}, nil, err
	}
	selected := -1
	for i, ch := range s.Changes {
		if ch.Slug == slug && slug != "" {
			if selected >= 0 {
				return fail(fmt.Errorf("change selector %q is ambiguous", slug))
			}
			selected = i
		}
	}
	if selected < 0 {
		return fail(fmt.Errorf("change selector %q is unknown", slug))
	}
	ch := s.Changes[selected]
	target := -1
	for i, m := range s.Modules {
		if m.Name == ch.Module {
			if target >= 0 {
				return fail(fmt.Errorf("target module %q is ambiguous", ch.Module))
			}
			target = i
		}
	}
	if target < 0 {
		return fail(fmt.Errorf("target module %q is missing", ch.Module))
	}
	m := s.Modules[target]
	text, err := changeview.Text(ch)
	if err != nil {
		return fail(err)
	}
	checked, err := contracts.ParseChange(text, changeview.Path(ch), m.IDPrefix)
	if err != nil {
		return fail(fmt.Errorf("invalid selected change: %w", err))
	}
	if (ch.Run == "") != (ch.Phase == "") {
		return fail(fmt.Errorf("run and phase must either both be assigned or both be empty"))
	}
	if ch.Phase == "applied" {
		return fail(fmt.Errorf("an applied change cannot remain queued"))
	}
	started := ch.Run != ""
	for _, item := range checked.Added {
		local := strings.HasPrefix(item.ID, "new:")
		if started && local {
			return fail(fmt.Errorf("started change retains handle %s; recover its assigned IDs before resuming", item.ID))
		}
		if !started && !local {
			return fail(fmt.Errorf("author-chosen Added ID %s must be a local handle before start", item.ID))
		}
	}
	bindings := map[string]string{}
	if !started {
		maximum := new(big.Int)
		reserve := func(items []contracts.Spec) {
			for _, item := range items {
				prefix := m.IDPrefix + "-"
				if !strings.HasPrefix(item.ID, prefix) {
					continue
				}
				digits := strings.TrimPrefix(item.ID, prefix)
				if len(digits) < 3 {
					continue
				}
				valid := true
				for _, digit := range digits {
					if digit < '0' || digit > '9' {
						valid = false
						break
					}
				}
				if !valid {
					continue
				}
				value, ok := new(big.Int).SetString(digits, 10)
				if ok && value.Cmp(maximum) > 0 {
					maximum.Set(value)
				}
			}
		}
		reserve(m.Specs)
		for _, archived := range s.Archive {
			if archived.Module == m.Name {
				reserve(archived.Added)
				reserve(archived.Modified)
				reserve(archived.Removed)
			}
		}
		for i, queued := range s.Changes {
			if i != selected && queued.Module == m.Name && queued.Run != "" && queued.Phase != "" {
				reserve(queued.Added)
			}
		}
		one := big.NewInt(1)
		for _, item := range checked.Added {
			maximum.Add(maximum, one)
			digits := maximum.String()
			if len(digits) < 3 {
				digits = strings.Repeat("0", 3-len(digits)) + digits
			}
			bindings[item.ID] = m.IDPrefix + "-" + digits
		}
		// Match the entire handle namespace token rather than replacing shorter
		// names first. The public parser has already proved every token local.
		text = handle.ReplaceAllStringFunc(text, func(token string) string {
			if id, ok := bindings[token]; ok {
				return id
			}
			return token
		})
		checked, err = contracts.ParseChange(text, changeview.Path(ch), m.IDPrefix)
		if err != nil {
			return fail(fmt.Errorf("invalid allocated change: %w", err))
		}
	}
	// Filename validation is canonical, but the typed identity owns queue order.
	checked.Path, checked.Slug, checked.Number = ch.Path, ch.Slug, ch.Number
	return checked, bindings, nil
}
