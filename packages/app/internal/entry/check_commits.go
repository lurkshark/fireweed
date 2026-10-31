package entry

import (
	"fmt"
	"io"
	"strings"

	"github.com/lurkshark/fireweed/packages/assets"
	"github.com/lurkshark/fireweed/packages/checks"
)

type checkCommitsSquash struct {
	Commit string `json:"commit"`
	Change string `json:"change"`
}

type checkCommitsData struct {
	Commits  []string             `json:"commits"`
	Squashed []checkCommitsSquash `json:"squashed"`
}

func runCheckCommits(args []string, jsonMode bool, out, errOut io.Writer) int {
	if len(args) > 1 {
		return cliFailure(2, jsonMode, out, errOut, "U01", "Use check-commits [range] with at most one revision range.")
	}
	snapshot, err := loadValidatedProject()
	if err != nil {
		return cliFailure(2, jsonMode, out, errOut, projectFailureCode(err, "C01"), fmt.Sprintf("%v. Provide a valid fireweed.yaml and module declarations, then retry check-commits.", err))
	}
	revisionRange := ""
	if len(args) == 1 {
		revisionRange = args[0]
	}
	report, err := checks.History(snapshot.Root, revisionRange)
	if err != nil {
		code := projectFailureCode(err, "P01")
		return cliFailure(1, jsonMode, out, errOut, code, fmt.Sprintf("%v. Use a valid Git revision range and repair the recorded project documents, then retry check-commits.", err))
	}
	if len(report.Diagnostics) != 0 {
		diagnostics := make([]diagnostic, 0, len(report.Diagnostics))
		for _, item := range report.Diagnostics {
			guidance, _ := assets.Explain(item.Code)
			diagnostics = append(diagnostics, diagnostic{
				Code: item.Code, Path: item.Path,
				Message: fmt.Sprintf("%s: %s. %s Use explain %s for corrective guidance.", item.Path, item.Message, guidance, item.Code),
			})
		}
		return cliResponse(1, jsonMode, out, errOut, nil, "", diagnostics)
	}
	data := checkCommitsData{Commits: nonNil(report.Commits), Squashed: []checkCommitsSquash{}}
	var plain strings.Builder
	fmt.Fprintf(&plain, "Checked %d commits\n", len(report.Commits))
	for _, item := range report.Squashed {
		data.Squashed = append(data.Squashed, checkCommitsSquash{Commit: item.Commit, Change: item.Change})
		fmt.Fprintf(&plain, "Squashed %s %s\n", item.Commit, item.Change)
	}
	return cliResponse(0, jsonMode, out, errOut, data, plain.String(), nil)
}
