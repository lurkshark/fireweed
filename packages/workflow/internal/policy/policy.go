package policy

import (
	"sort"
	"strings"

	"github.com/lurkshark/fireweed/packages/checks"
	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/gitstore"
	"github.com/lurkshark/fireweed/packages/project"
)

// Paths and patterns use repository slashes, independent of the host platform.
func canonical(path string) bool {
	if path == "" || strings.ContainsAny(path, "\\\x00") || strings.HasPrefix(path, "/") {
		return false
	}
	if len(path) >= 2 && path[1] == ':' && ((path[0] >= 'A' && path[0] <= 'Z') || (path[0] >= 'a' && path[0] <= 'z')) {
		return false
	}
	for _, part := range strings.Split(path, "/") {
		if part == "" || part == "." || part == ".." {
			return false
		}
	}
	return true
}
func matches(patterns []string, path string) (bool, error) {
	found := false
	for _, pattern := range patterns {
		yes, err := contracts.MatchGlob(pattern, path)
		if err != nil {
			return false, err
		}
		found = found || yes
	}
	return found, nil
}
func Check(snapshot project.Snapshot, module contracts.Module, change contracts.Change, edits []gitstore.Edit) []checks.Diagnostic {
	result := make([]checks.Diagnostic, 0)
	seen := map[checks.Diagnostic]bool{}
	add := func(code, path, message string) {
		d := checks.Diagnostic{Code: code, Severity: "error", Path: path, Message: message}
		if !seen[d] {
			seen[d] = true
			result = append(result, d)
		}
	}
	if module.Name != change.Module || strings.TrimSpace(change.Run) == "" || (change.Phase != "evals" && change.Phase != "audit" && change.Phase != "impl") {
		add("P02", change.Path, "Restore the approved claim target, nonempty run and active evals/audit/impl phase before checking worker edits.")
		return result
	}
	// Own this slice: appending the projected target must not alter snapshot data.
	modules := make([]contracts.Module, 0, len(snapshot.Modules)+1)
	for _, candidate := range snapshot.Modules {
		if candidate.Name != module.Name {
			modules = append(modules, candidate)
		}
	}
	modules = append(modules, module)
	described := false
	for _, section := range change.Document.Sections {
		if section.Name == "Interface" && strings.TrimSpace(section.Body) != "" {
			described = true
		}
	}
	malformed := func(path string) {
		add("P2-01", path, "Restore a well-formed edit with canonical relative file paths, an A/M/D/T/R/C status and valid distinct rename/copy endpoints.")
	}
	workerError := func(path string) {
		add("P2-01", path, "Restore this path to its approved state; keep worker edits in the current phase's permitted target-module zones or unowned worker_writable paths.")
	}
	evalError := func(path string) {
		add("P1-01", path, "Restore this path; use only evals, described interfaces, manifests, shared writable files and newly added disposable stubs during evals.")
	}
	endpoint := func(path, status string) {
		var owner *contracts.Module
		for i := range modules {
			if strings.HasPrefix(path, modules[i].Dir+"/") {
				owner = &modules[i]
				break
			}
		}
		appendOnly := false
		if owner != nil {
			var err error
			appendOnly, err = matches(owner.AppendOnly, strings.TrimPrefix(path, owner.Dir+"/"))
			if err != nil {
				workerError(path)
				return
			}
		}
		newFile := status == "A" || status == "C"
		if appendOnly && !newFile {
			add("H03", path, "Restore the existing append-only path; add new files without modifying, deleting, type-changing or renaming existing append-only files.")
		}
		if appendOnly && newFile && change.Phase == "evals" {
			evalError(path)
		}
		if path == change.Path && status == "M" {
			return
		}
		if path == change.Path {
			workerError(path)
			return
		}
		zone, err := contracts.ZoneFor(path, module.Dir, module.Zones, snapshot.Config.Protected)
		if err != nil {
			workerError(path)
			return
		}
		parts := strings.Split(path, "/")
		if zone == contracts.Protected || parts[len(parts)-1] == "SPEC.md" || strings.HasPrefix(path, "changes/") {
			workerError(path)
			return
		}
		if change.Phase == "audit" {
			workerError(path)
			return
		}
		if owner == nil {
			permitted, err := matches(snapshot.Config.WorkerWritable, path)
			if err != nil || !permitted {
				workerError(path)
			}
			return
		}
		if owner.Name != module.Name {
			workerError(path)
			return
		}
		if change.Phase == "impl" {
			if zone == contracts.EvalZone || zone == contracts.InterfaceZone {
				workerError(path)
			}
			return
		}
		switch zone {
		case contracts.InterfaceZone:
			if !described {
				evalError(path)
			}
		case contracts.PreserveZone, contracts.AppendOnlyZone:
			evalError(path)
		case contracts.DisposableZone:
			if !newFile {
				evalError(path)
			}
		}
	}
	for _, edit := range edits {
		valid := true
		if !canonical(edit.Path) {
			malformed(edit.Path)
			valid = false
		}
		switch edit.Status {
		case "R", "C":
			if edit.OldPath == "" || edit.OldPath == edit.Path {
				malformed(edit.Path)
				valid = false
			} else if !canonical(edit.OldPath) {
				malformed(edit.OldPath)
				valid = false
			}
		case "A", "M", "D", "T":
			if edit.OldPath != "" {
				malformed(edit.Path)
				valid = false
			}
		default:
			malformed(edit.Path)
			valid = false
		}
		if !valid {
			continue
		}
		if edit.Status == "R" {
			endpoint(edit.OldPath, edit.Status)
		}
		endpoint(edit.Path, edit.Status)
	}
	sort.Slice(result, func(i, j int) bool {
		a, b := result[i], result[j]
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
	return result
}
