package entry

import (
	"context"
	"fmt"
	"io"
	"strings"

	"github.com/lurkshark/fireweed/packages/checks"
	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/project"
)

func coverage(out, errOut io.Writer, option options, fail func(string) int) int {
	snapshot, err := project.Load(".")
	if err != nil {
		return loadFailure(out, errOut, option, err)
	}
	name := option.operands[0]
	module := contracts.Module{}
	found := false
	choices := make([]string, 0, len(snapshot.Modules))
	for _, m := range snapshot.Modules {
		choices = append(choices, m.Name)
		if m.Name == name {
			module, found = m, true
		}
	}
	if !found {
		return fail(fmt.Sprintf("unknown module %q; choose an exact module name from: %s. Run fireweed coverage <module>", name, strings.Join(choices, ", ")))
	}
	if option.withChangeSet {
		var selected contracts.Change
		matched := false
		slugs := make([]string, 0, len(snapshot.Changes))
		for _, change := range snapshot.Changes {
			slugs = append(slugs, change.Slug)
			if change.Slug == option.withChange {
				selected, matched = change, true
			}
		}
		if !matched {
			return fail(fmt.Sprintf("unknown queued change %q; valid queued choices: %s. Run fireweed coverage %s --with-change <slug>", option.withChange, strings.Join(slugs, ", "), name))
		}
		if selected.Module != name {
			return fail(fmt.Sprintf("change %q targets module %s. Run fireweed coverage %s --with-change %s", selected.Slug, selected.Module, selected.Module, selected.Slug))
		}
		projected, projectionErr := project.Project(snapshot, selected.Slug)
		if projectionErr == nil {
			for _, candidate := range projected.Modules {
				if candidate.Name == name {
					module = candidate
					break
				}
			}
			module, projectionErr = project.ApplyChange(module, selected)
		}
		if projectionErr != nil {
			return emit(out, errOut, option, nil, []diagnostic{{Code: "C05", Severity: "error", Path: selected.Path, Message: "Cannot preview selected change: " + projectionErr.Error() + ". Correct queued change intent before running coverage."}}, 1)
		}
		snapshot = projected
	}
	result := checks.Evaluate(context.Background(), snapshot, module, nil, nil)
	ds := make([]diagnostic, 0, len(result.Diagnostics))
	code := 0
	for _, d := range result.Diagnostics {
		ds = append(ds, diagnostic{Code: d.Code, Severity: d.Severity, Message: d.Message, Path: d.Path})
		if d.Severity == "error" {
			code = 1
		}
	}
	return emit(out, errOut, option, result, ds, code)
}
