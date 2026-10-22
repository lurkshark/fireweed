package entry

import (
	"fmt"
	"io"
	"sort"
	"strings"
	"time"

	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/project"
)

func checkedDiagnostics(snapshot project.Snapshot, queued bool) ([]diagnostic, int) {
	checked := project.Check(snapshot)
	if queued {
		checked = append(checked, project.CheckQueue(snapshot)...)
	}
	sort.SliceStable(checked, func(i, j int) bool {
		a, b := checked[i], checked[j]
		if a.Path != b.Path {
			return a.Path < b.Path
		}
		if a.Code != b.Code {
			return a.Code < b.Code
		}
		if a.Severity != b.Severity {
			return a.Severity < b.Severity
		}
		return a.Message < b.Message
	})
	result := make([]diagnostic, 0, len(checked))
	code := 0
	for _, d := range checked {
		result = append(result, diagnostic{Code: d.Code, Severity: d.Severity, Message: d.Message, Path: d.Path})
		if d.Severity == "error" {
			code = 1
		}
	}
	return result, code
}

func loadFailure(out, errOut io.Writer, option options, err error) int {
	return emit(out, errOut, option, nil, []diagnostic{{Code: "C01", Severity: "error", Path: "fireweed.yaml", Message: "Cannot load project: " + err.Error() + ". Run from a Git project and correct fireweed.yaml."}}, 2)
}

func show(out, errOut io.Writer, option options, fail func(string) int) int {
	snapshot, err := project.Load(".")
	if err != nil {
		return loadFailure(out, errOut, option, err)
	}
	name := option.operands[0]
	var module contracts.Module
	found := false
	choices := make([]string, 0, len(snapshot.Modules))
	for _, m := range snapshot.Modules {
		choices = append(choices, m.Name)
		if m.Name == name {
			module, found = m, true
		}
	}
	if !found {
		return fail(fmt.Sprintf("unknown module %q; choose an exact module name from: %s", name, strings.Join(choices, ", ")))
	}
	diagnostics, code := checkedDiagnostics(snapshot, option.projected)
	id := ""
	if len(option.operands) == 2 {
		id = option.operands[1]
	}
	if option.history {
		return emit(out, errOut, option, archiveHistory(snapshot, name, id), diagnostics, code)
	}
	if option.projected {
		projected, err := project.Project(snapshot, "")
		if err != nil {
			hasC05 := false
			for _, d := range diagnostics {
				hasC05 = hasC05 || d.Code == "C05" && d.Severity == "error"
			}
			if !hasC05 {
				diagnostics = append(diagnostics, diagnostic{Code: "C05", Severity: "error", Message: "Cannot project queue: " + err.Error() + ". Correct the queued change intent before reading projected specs."})
				sort.SliceStable(diagnostics, func(i, j int) bool {
					a, b := diagnostics[i], diagnostics[j]
					if a.Path != b.Path {
						return a.Path < b.Path
					}
					if a.Code != b.Code {
						return a.Code < b.Code
					}
					if a.Severity != b.Severity {
						return a.Severity < b.Severity
					}
					return a.Message < b.Message
				})
			}
			return emit(out, errOut, option, nil, diagnostics, 1)
		}
		for _, m := range projected.Modules {
			if m.Name == name {
				module = m
				break
			}
		}
	}
	items := module.Specs
	var data any = module
	if id != "" {
		selected := []contracts.Spec{}
		for _, item := range items {
			if item.ID == id {
				selected = append(selected, item)
			}
		}
		correction := "Run fireweed show " + name
		if option.projected {
			correction += " --projected"
		}
		correction += " without an ID to list available specs"
		if len(selected) == 0 {
			return fail(fmt.Sprintf("unknown exact spec ID %q. %s", id, correction))
		}
		if len(selected) > 1 {
			return fail(fmt.Sprintf("ambiguous local handle %q. %s", id, correction))
		}
		items, data = selected, selected[0]
	}
	if !option.json {
		data = plainSpecs(items)
	}
	return emit(out, errOut, option, data, diagnostics, code)
}

func plainSpecs(items []contracts.Spec) string {
	parts := make([]string, 0, len(items))
	for _, item := range items {
		heading := "### " + item.ID
		if item.Tag != "" {
			heading += " " + item.Tag
		}
		parts = append(parts, heading+"\n"+item.Body)
	}
	if len(parts) == 0 {
		return ""
	}
	return strings.Join(parts, "\n\n") + "\n"
}

type historyRow struct {
	Change     contracts.Change `json:"change"`
	Motivation string           `json:"motivation"`
	Decisions  string           `json:"decisions"`
	Provenance string           `json:"provenance"`
}

func sectionBody(doc contracts.Document, name string) string {
	for _, section := range doc.Sections {
		if section.Name == name {
			return section.Body
		}
	}
	return ""
}
func implementationTime(provenance string) (time.Time, bool) {
	var date time.Time
	known := false
	for _, line := range strings.Split(provenance, "\n") {
		fields := strings.Split(line, "|")
		if len(fields) < 2 || strings.TrimSpace(strings.TrimPrefix(strings.TrimSpace(fields[0]), "-")) != "impl" {
			continue
		}
		if parsed, err := time.Parse(time.RFC3339Nano, strings.TrimSpace(fields[1])); err == nil {
			date, known = parsed, true
		}
	}
	return date, known
}
func archiveHistory(snapshot project.Snapshot, module, id string) []historyRow {
	rows := []historyRow{}
	for _, change := range snapshot.Archive {
		if change.Module != module {
			continue
		}
		found := id == ""
		for _, group := range [][]contracts.Spec{change.Added, change.Modified, change.Removed} {
			for _, spec := range group {
				found = found || spec.ID == id
			}
		}
		if !found {
			continue
		}
		rows = append(rows, historyRow{change, sectionBody(change.Document, "Motivation"), sectionBody(change.Document, "Decisions"), sectionBody(change.Document, "Provenance")})
	}
	sort.SliceStable(rows, func(i, j int) bool {
		a, b := rows[i], rows[j]
		at, ak := implementationTime(a.Provenance)
		bt, bk := implementationTime(b.Provenance)
		if ak != bk {
			return ak
		}
		if ak && !at.Equal(bt) {
			return at.After(bt)
		}
		if a.Change.Number != b.Change.Number {
			return a.Change.Number > b.Change.Number
		}
		if a.Change.Slug != b.Change.Slug {
			return a.Change.Slug > b.Change.Slug
		}
		return a.Change.Path > b.Change.Path
	})
	return rows
}
