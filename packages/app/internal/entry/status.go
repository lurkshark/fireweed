package entry

import (
	"fmt"
	"io"
	"sort"
	"strings"

	"github.com/lurkshark/fireweed/packages/workflow"
)

type statusQueueEntry struct {
	Slug        string `json:"slug"`
	Path        string `json:"path"`
	Module      string `json:"module"`
	Phase       string `json:"phase"`
	Blocked     string `json:"blocked"`
	AuditRounds int    `json:"audit_rounds"`
	WaitingOn   string `json:"waiting_on"`
	WaitingRule string `json:"waiting_rule"`
}

type statusQuarantine struct {
	Module string `json:"module"`
	ID     string `json:"id"`
	Reason string `json:"reason"`
}

type statusOptOut struct {
	Module string `json:"module"`
	Reason string `json:"reason"`
}

type statusMetric struct {
	Module          string         `json:"module"`
	Seconds         float64        `json:"seconds"`
	CasesPerSpec    map[string]int `json:"cases_per_spec"`
	EvalLines       int            `json:"eval_lines"`
	DisposableLines int            `json:"disposable_lines"`
}

type statusData struct {
	Modules         []moduleResult     `json:"modules"`
	Queue           []statusQueueEntry `json:"queue"`
	Quarantines     []statusQuarantine `json:"quarantines"`
	BoundaryOptOuts []statusOptOut     `json:"boundary_opt_outs"`
	DriftSinceDrill map[string]int     `json:"drift_since_drill"`
	Unapproved      int                `json:"unapproved"`
	PhaseCommits    map[string]int     `json:"phase_commits"`
	Metrics         []statusMetric     `json:"metrics"`
}

func statusCounts(counts map[string]int) map[string]int {
	if counts == nil {
		return map[string]int{}
	}
	return counts
}

func runStatus(jsonMode bool, out, errOut io.Writer) int {
	snapshot, err := loadValidatedProject()
	if err != nil {
		return cliFailure(2, jsonMode, out, errOut, projectFailureCode(err, "C01"), fmt.Sprintf("%v. Provide a valid fireweed.yaml and module declarations, then retry status.", err))
	}
	report, err := workflow.Status(snapshot.Root)
	diagnostics := reported(report.Diagnostics)
	if err != nil {
		diagnostics = lifecycleError(err, diagnostics, "Use the reported cause to correct the project or Git state, then retry status.")
		last := len(diagnostics) - 1
		if diagnostics[last].Code == "P01" {
			diagnostics[last].Code = projectFailureCode(err, "P01")
		}
		return cliResponse(1, jsonMode, out, errOut, nil, "", diagnostics)
	}

	data := statusData{
		Modules:         moduleResults(report.Modules),
		Queue:           make([]statusQueueEntry, 0, len(report.Queue)),
		Quarantines:     make([]statusQuarantine, 0, len(report.Quarantines)),
		BoundaryOptOuts: make([]statusOptOut, 0, len(report.BoundaryOptOuts)),
		DriftSinceDrill: statusCounts(report.DriftSinceDrill),
		Unapproved:      report.Unapproved,
		PhaseCommits:    statusCounts(report.PhaseCommits),
		Metrics:         make([]statusMetric, 0, len(report.Costs)),
	}
	for _, entry := range report.Queue {
		data.Queue = append(data.Queue, statusQueueEntry{
			Slug: entry.Slug, Path: entry.Path, Module: entry.Module, Phase: entry.Phase,
			Blocked: entry.Blocked, AuditRounds: entry.AuditRounds,
			WaitingOn: entry.WaitingOn, WaitingRule: entry.WaitingRule,
		})
	}
	for _, q := range report.Quarantines {
		data.Quarantines = append(data.Quarantines, statusQuarantine{Module: q.Module, ID: q.ID, Reason: q.Reason})
	}
	for _, optOut := range report.BoundaryOptOuts {
		data.BoundaryOptOuts = append(data.BoundaryOptOuts, statusOptOut{Module: optOut.Module, Reason: optOut.Reason})
	}
	for _, cost := range report.Costs {
		data.Metrics = append(data.Metrics, statusMetric{
			Module: cost.Module, Seconds: cost.Seconds, CasesPerSpec: statusCounts(cost.CasesPerSpec),
			EvalLines: cost.EvalLines, DisposableLines: cost.DisposableLines,
		})
	}
	var plain strings.Builder
	plain.WriteString("Modules:\n")
	printModules(&plain, data.Modules)
	plain.WriteString("Queue:\n")
	if len(data.Queue) == 0 {
		plain.WriteString("  empty\n")
	}
	for _, entry := range data.Queue {
		fmt.Fprintf(&plain, "  slug=%s path=%s module=%s phase=%s blocked=%q audit_rounds=%d waiting_on=%q waiting_rule=%q\n",
			entry.Slug, entry.Path, entry.Module, entry.Phase, entry.Blocked, entry.AuditRounds, entry.WaitingOn, entry.WaitingRule)
	}
	plain.WriteString("Quarantines:\n")
	for _, q := range data.Quarantines {
		fmt.Fprintf(&plain, "  module=%s id=%s reason=%q\n", q.Module, q.ID, q.Reason)
	}
	plain.WriteString("Boundary opt-outs:\n")
	for _, optOut := range data.BoundaryOptOuts {
		fmt.Fprintf(&plain, "  module=%s reason=%q\n", optOut.Module, optOut.Reason)
	}
	writeStatusCounts(&plain, "Drift since drill", data.DriftSinceDrill)
	fmt.Fprintf(&plain, "Unapproved: %d\n", data.Unapproved)
	writeStatusCounts(&plain, "Phase commits", data.PhaseCommits)
	plain.WriteString("Costs:\n")
	for _, cost := range data.Metrics {
		fmt.Fprintf(&plain, "Cost %s: seconds=%g cases_per_spec=%s eval_lines=%d disposable_lines=%d\n",
			cost.Module, cost.Seconds, formatStatusCounts(cost.CasesPerSpec), cost.EvalLines, cost.DisposableLines)
	}
	return cliResponse(0, jsonMode, out, errOut, data, plain.String(), diagnostics)
}

func writeStatusCounts(out io.Writer, label string, counts map[string]int) {
	fmt.Fprintf(out, "%s:\n", label)
	keys := make([]string, 0, len(counts))
	for key := range counts {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		fmt.Fprintf(out, "  %s=%d\n", key, counts[key])
	}
}

func formatStatusCounts(counts map[string]int) string {
	keys := make([]string, 0, len(counts))
	for key := range counts {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	parts := make([]string, 0, len(keys))
	for _, key := range keys {
		parts = append(parts, fmt.Sprintf("%s=%d", key, counts[key]))
	}
	return "[" + strings.Join(parts, " ") + "]"
}
