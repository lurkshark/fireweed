package entry

import (
	"fmt"
	"io"
	"strings"

	"github.com/lurkshark/fireweed/packages/project"
)

type listEntry struct {
	Slug        string `json:"slug"`
	Path        string `json:"path"`
	Module      string `json:"module"`
	Phase       string `json:"phase"`
	Blocked     string `json:"blocked"`
	AuditRounds int    `json:"audit_rounds"`
}

func runList(jsonMode bool, out, errOut io.Writer) int {
	snapshot, err := loadValidatedProject()
	if err != nil {
		return cliFailure(2, jsonMode, out, errOut, projectFailureCode(err, "C01"), fmt.Sprintf("%v. Provide a valid fireweed.yaml and module declarations, then retry list.", err))
	}
	changes, err := project.Queue(snapshot.Root)
	if err != nil {
		return cliFailure(1, jsonMode, out, errOut, projectFailureCode(err, "P01"), fmt.Sprintf("%v. Use corrected queued changes, then retry list.", err))
	}
	entries := make([]listEntry, 0, len(changes))
	var plain strings.Builder
	for _, change := range changes {
		phase, _ := change.Document.Frontmatter["phase"].(string)
		if phase == "" {
			phase = "queued"
		}
		blocked, _ := change.Document.Frontmatter["blocked"].(string)
		rounds, _ := change.Document.Frontmatter["audit_rounds"].(int)
		entry := listEntry{Slug: change.Slug, Path: change.Path, Module: change.Module, Phase: phase, Blocked: blocked, AuditRounds: rounds}
		entries = append(entries, entry)
		fmt.Fprintf(&plain, "slug=%s path=%s module=%s phase=%s blocked=%q audit_rounds=%d\n", entry.Slug, entry.Path, entry.Module, entry.Phase, entry.Blocked, entry.AuditRounds)
	}
	return cliResponse(0, jsonMode, out, errOut, entries, plain.String(), nil)
}
