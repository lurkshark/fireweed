package entry

import (
	"fmt"
	"io"
	"strings"

	"github.com/lurkshark/fireweed/packages/project"
)

func runShowStale(selector string, jsonMode bool, out, errOut io.Writer) int {
	snapshot, err := loadValidatedProject()
	if err != nil {
		return cliFailure(2, jsonMode, out, errOut, errorCode(err, "C01"), fmt.Sprintf("%v. Use a valid fireweed.yaml and module declarations, then retry show --stale <change>.", err))
	}
	entries, err := project.StaleSpecs(snapshot.Root, selector)
	if err != nil {
		code := errorCode(err, "P01")
		status := 1
		if code == "U02" || code == "C13" || code == "C01" || code == "C02" || code == "C12" {
			status = 2
		}
		return cliFailure(status, jsonMode, out, errOut, code, fmt.Sprintf("%v (selector %q). Use a unique pending change selector and correct the reported cause, then retry show --stale <change>.", err, selector))
	}
	var plain strings.Builder
	if len(entries) == 0 {
		plain.WriteString("No stale specs found.\n")
	}
	for _, entry := range entries {
		fmt.Fprintf(&plain, "Spec: %s\n", entry.ID)
		if entry.Recovered {
			writeStaleDiff(&plain, entry.Before, entry.Current)
		} else {
			fmt.Fprintf(&plain, "Stamped text unavailable; recorded hash: %s\nCurrent:\n", entry.RecordedHash)
			plain.WriteString(entry.Current)
			if !strings.HasSuffix(entry.Current, "\n") {
				plain.WriteByte('\n')
			}
		}
	}
	return cliResponse(0, jsonMode, out, errOut, plain.String(), plain.String(), nil)
}

func writeStaleDiff(out *strings.Builder, before, current string) {
	oldLines := staleLines(before)
	newLines := staleLines(current)
	lengths := make([][]int, len(oldLines)+1)
	for i := range lengths {
		lengths[i] = make([]int, len(newLines)+1)
	}
	for i := len(oldLines) - 1; i >= 0; i-- {
		for j := len(newLines) - 1; j >= 0; j-- {
			if oldLines[i] == newLines[j] {
				lengths[i][j] = lengths[i+1][j+1] + 1
			} else if lengths[i+1][j] >= lengths[i][j+1] {
				lengths[i][j] = lengths[i+1][j]
			} else {
				lengths[i][j] = lengths[i][j+1]
			}
		}
	}
	i, j := 0, 0
	for i < len(oldLines) || j < len(newLines) {
		switch {
		case i < len(oldLines) && j < len(newLines) && oldLines[i] == newLines[j]:
			fmt.Fprintf(out, " %s\n", oldLines[i])
			i++
			j++
		case i < len(oldLines) && (j == len(newLines) || lengths[i+1][j] >= lengths[i][j+1]):
			fmt.Fprintf(out, "-%s\n", oldLines[i])
			i++
		default:
			fmt.Fprintf(out, "+%s\n", newLines[j])
			j++
		}
	}
}

func staleLines(text string) []string {
	if text == "" {
		return nil
	}
	return strings.Split(strings.TrimSuffix(text, "\n"), "\n")
}
