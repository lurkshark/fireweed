package lifecycle

import (
	"regexp"
	"slices"
	"strings"

	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/workflow/model"
)

var findingPattern = regexp.MustCompile(`^- (\[untestable\] )?([A-Z0-9]+-[0-9]+): "(.*)"\s+(?:—|–|-)\s+(.+)$`)
var auditID = regexp.MustCompile(`^[A-Z0-9]+-[0-9]+$`)

type auditResult struct {
	text    string
	next    string
	blocked string
	rounds  int
}

func auditSubmission(state model.State, held claim) (auditResult, error) {
	result := auditResult{next: "impl", rounds: integer(state.Change.Document.Frontmatter["audit_rounds"], 0)}
	approved, err := held.approvedDocument()
	if err != nil {
		return result, err
	}
	specs, err := items(moduleNamed(state.Snapshot, state.Change.Module).Document, "Specs")
	if err != nil {
		return result, err
	}

	allowed := make(map[string]bool)
	if result.rounds > 0 {
		for _, line := range strings.Split(section(approved, "Audit"), "\n") {
			if match := findingPattern.FindStringSubmatch(strings.TrimSpace(line)); match != nil {
				allowed[match[2]] = true
			}
		}
	}
	var findings, checked []string
	var noFindings, hasChecked, untestable bool
	for _, line := range strings.Split(section(state.Change.Document, "Audit"), "\n") {
		line = strings.TrimSpace(line)
		switch {
		case line == "":
		case line == "No findings":
			noFindings = true
		case strings.HasPrefix(line, "Checked IDs:"):
			hasChecked = true
			checked = append(checked, parseAuditIDs(strings.TrimPrefix(line, "Checked IDs:"))...)
		case strings.HasPrefix(line, "No findings; checked "):
			noFindings, hasChecked = true, true
			checked = append(checked, parseAuditIDs(strings.TrimPrefix(line, "No findings; checked "))...)
		default:
			match := findingPattern.FindStringSubmatch(line)
			if match == nil || match[3] == "" {
				return result, failure("P09", state.Change.Path, "malformed audit submission: %s", line)
			}
			if result.rounds != 0 && !allowed[match[2]] {
				continue
			}
			for _, spec := range specs {
				if spec.ID == match[2] && strings.Contains(spec.Body, match[3]) {
					findings = append(findings, line)
					untestable = untestable || match[1] != ""
					break
				}
			}
		}
	}
	if len(findings) > 0 {
		result.text = "\n" + strings.Join(findings, "\n") + "\n\n"
		switch {
		case untestable:
			result.blocked = "audit found an untestable clause"
		case result.rounds >= integer(mapping(state.Snapshot.Config["audit"])["max_rounds"], 2):
			result.blocked = "audit round limit reached"
		default:
			result.next = "evals"
			result.rounds++
		}
		if result.blocked != "" {
			result.next = "audit"
		}
		return result, nil
	}
	if !noFindings || !hasChecked {
		return result, failure("P09", state.Change.Path, "audit requires Checked IDs and No findings")
	}
	for _, id := range changedIDs(state.Change, specs) {
		if !slices.Contains(checked, id) {
			return result, failure("P09", state.Change.Path, "audit did not check %s", id)
		}
	}
	result.text = "\nNo findings; checked " + strings.Join(checked, ", ") + "\n\n"
	return result, nil
}

func parseAuditIDs(text string) []string {
	var ids []string
	for _, id := range strings.Fields(strings.ReplaceAll(text, ",", " ")) {
		if auditID.MatchString(id) {
			ids = append(ids, id)
		}
	}
	return ids
}

func auditEnabled(config map[string]any) bool {
	enabled, ok := mapping(config["audit"])["enabled"].(bool)
	return !ok || enabled
}

func appendSection(doc contracts.Document, name, text string) string {
	body := strings.TrimRight(section(doc, name), "\r\n")
	return contracts.ReplaceSection(doc.Body, name, body+"\n"+strings.TrimSpace(text)+"\n\n")
}
