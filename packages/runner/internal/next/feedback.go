package next

import (
	"fmt"
	"maps"
	"slices"
	"strings"
	"unicode/utf8"

	"github.com/lurkshark/fireweed/packages/checks"
	"github.com/lurkshark/fireweed/packages/workflow"
)

func verificationFeedback(result workflow.Result, err error) string {
	var feedback strings.Builder
	fmt.Fprintf(&feedback, "\n\n## Last verification\n%v\n", err)
	for _, module := range slices.Sorted(maps.Keys(result.Reports)) {
		report := result.Reports[module]
		for _, stream := range []string{report.Stdout, report.Stderr} {
			if stream == "" {
				continue
			}
			feedback.WriteString(stream)
			if !strings.HasSuffix(stream, "\n") {
				feedback.WriteByte('\n')
			}
		}
	}
	feedback.WriteString(failedCases(result.Reports))
	return feedback.String()
}

func failedCases(reports map[string]checks.Report) string {
	const (
		heading = "\n## Failed testcase details\n"
		omitted = "[Additional failed cases omitted]\n"
		limit   = 16384
	)
	var section strings.Builder
	used := utf8.RuneCountInString(heading)
	for _, module := range slices.Sorted(maps.Keys(reports)) {
		for _, c := range reports[module].Cases {
			if c.Result != "fail" && c.Result != "error" {
				continue
			}
			failure := c.Failure
			if runes := []rune(failure); len(runes) > 4096 {
				failure = string(runes[:4096]) + "\n[Failure text truncated]"
			}
			entry := fmt.Sprintf("- %s / %s / %s\n%s\n", module, c.Classname, c.Name, failure)
			size := utf8.RuneCountInString(entry)
			if section.Len() == 0 {
				section.WriteString(heading)
			}
			if used+size+utf8.RuneCountInString(omitted) > limit {
				section.WriteString(omitted)
				return section.String()
			}
			section.WriteString(entry)
			used += size
		}
	}
	return section.String()
}
