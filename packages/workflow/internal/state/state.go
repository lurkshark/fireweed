package state

import (
	"fmt"
	"reflect"
	"sort"
	"strings"

	"github.com/lurkshark/fireweed/packages/checks"
	"github.com/lurkshark/fireweed/packages/contracts"
)

func problem(code, path, message string) checks.Diagnostic {
	return checks.Diagnostic{Code: code, Severity: "error", Path: path, Message: message}
}
func claimProblem(path string, err error) []checks.Diagnostic {
	return []checks.Diagnostic{problem("P02", path, fmt.Sprintf("Invalid approved claim state: %v. Restore the approved claim baseline before resuming this step.", err))}
}
func fixedSections(document contracts.Document, permitted map[string]bool) []contracts.Section {
	result := make([]contracts.Section, 0, len(document.Sections))
	for _, section := range document.Sections {
		if !permitted[section.Name] {
			result = append(result, section)
		}
	}
	return result
}

func Check(module contracts.Module, approved, current, path string) []checks.Diagnostic {
	before, err := contracts.ParseChange(approved, path, module.IDPrefix)
	if err != nil {
		return claimProblem(path, err)
	}
	if before.Module != module.Name || strings.TrimSpace(before.Run) == "" {
		return claimProblem(path, fmt.Errorf("the target must match %q and the run must be nonempty", module.Name))
	}
	switch before.Phase {
	case "evals", "audit", "impl":
	default:
		return claimProblem(path, fmt.Errorf("phase %q is not an active worker step", before.Phase))
	}
	after, err := contracts.ParseChange(current, path, module.IDPrefix)
	if err != nil {
		return []checks.Diagnostic{problem("P2-01", path, "Worker edits produced invalid change grammar: "+err.Error()+". Restore valid approved intent and edit only permitted worker sections.")}
	}
	// Each parse owns its metadata. blocked is the only mutable frontmatter field.
	delete(before.Document.Metadata, "blocked")
	delete(after.Document.Metadata, "blocked")
	result := make([]checks.Diagnostic, 0, 2)
	if !reflect.DeepEqual(before.Document.Metadata, after.Document.Metadata) {
		result = append(result, problem("P2-01", path, "Worker changed approved metadata. Restore all approved fields and key presence; only blocked may change during a step."))
	}
	permitted := map[string]bool{"Decisions": true, "Questions": true, "Notes": true}
	if before.Phase == "audit" {
		permitted = map[string]bool{"Audit": true, "Questions": true}
	}
	if before.Document.Preamble != after.Document.Preamble || !reflect.DeepEqual(fixedSections(before.Document, permitted), fixedSections(after.Document, permitted)) {
		result = append(result, problem("P2-01", path, "Worker changed approved body sections or preamble. Restore protected text and relative section order; edit only sections permitted by the approved phase."))
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
