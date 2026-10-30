package entry

import (
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/lurkshark/fireweed/packages/project"
)

type specHistoryRow struct {
	Path               string   `json:"path"`
	Module             string   `json:"module"`
	IDs                []string `json:"ids"`
	Motivation         string   `json:"motivation"`
	Decisions          string   `json:"decisions"`
	RemovedDecisions   string   `json:"removed_decisions"`
	Provenance         string   `json:"provenance"`
	ImplementationDate string   `json:"implementation_date"`
}

func runShowHistory(args []string, jsonMode bool, out, errOut io.Writer) int {
	root, err := findProjectRoot()
	if err != nil {
		return cliFailure(2, jsonMode, out, errOut, "C01", fmt.Sprintf("%v. Use show --history from a project containing fireweed.yaml.", err))
	}
	id := ""
	if len(args) == 2 {
		id = args[1]
	}
	entries, err := project.SpecHistory(root, args[0], id)
	if err != nil {
		action := "Use a valid archive and correct the reported cause, then retry show --history."
		diagnostics := lifecycleError(err, nil, action)
		status := 1
		if diagnostics[0].Code == "U02" {
			status = 2
			action = "Use a module or ID listed in the project, then retry show --history."
		} else if diagnostics[0].Code == "C01" {
			status = 2
			action = "Use a valid fireweed.yaml, then retry show --history."
		}
		diagnostics = lifecycleError(err, nil, action)
		return cliResponse(status, jsonMode, out, errOut, nil, "", diagnostics)
	}

	rows := make([]specHistoryRow, 0, len(entries))
	var plain strings.Builder
	if len(entries) == 0 {
		plain.WriteString("No history found.\n")
	}
	for _, entry := range entries {
		date := ""
		if !entry.ImplementationDate.IsZero() {
			date = entry.ImplementationDate.Format(time.RFC3339Nano)
		}
		rows = append(rows, specHistoryRow{
			Path: entry.Path, Module: entry.Module, IDs: nonNil(entry.IDs),
			Motivation: entry.Motivation, Decisions: entry.Decisions,
			RemovedDecisions: entry.RemovedDecisions, Provenance: entry.Provenance,
			ImplementationDate: date,
		})
		fmt.Fprintf(&plain, "Path: %s\nModule: %s\nIDs: %s\n", entry.Path, entry.Module, strings.Join(entry.IDs, ", "))
		if date == "" {
			plain.WriteString("Implementation date: undated\n")
		} else {
			fmt.Fprintf(&plain, "Implementation date: %s\n", date)
		}
		writeHistorySection(&plain, "Motivation", entry.Motivation)
		writeHistorySection(&plain, "Decisions", entry.Decisions)
		writeHistorySection(&plain, "Removed decisions", entry.RemovedDecisions)
		writeHistorySection(&plain, "Provenance", entry.Provenance)
	}
	return cliResponse(0, jsonMode, out, errOut, rows, plain.String(), nil)
}

func writeHistorySection(out *strings.Builder, label, content string) {
	fmt.Fprintf(out, "%s:\n%s", label, content)
	if !strings.HasSuffix(content, "\n") {
		out.WriteByte('\n')
	}
	out.WriteByte('\n')
}
