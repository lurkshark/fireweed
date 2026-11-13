package entry

import (
	"fmt"
	"maps"
	"slices"
	"strings"

	"github.com/lurkshark/fireweed/packages/project"
	"github.com/lurkshark/fireweed/packages/workflow"
)

type zoneRow struct {
	Protected  []string `json:"protected"`
	Evals      []string `json:"evals"`
	Interface  []string `json:"interface"`
	AppendOnly []string `json:"append_only"`
	Preserve   []string `json:"preserve"`
	Disposable []string `json:"disposable"`
}

type moduleRow struct {
	Name      string   `json:"name"`
	IDPrefix  string   `json:"id_prefix"`
	Dir       string   `json:"dir"`
	DependsOn []string `json:"depends_on"`
	Zones     zoneRow  `json:"zones"`
}

func moduleRows(modules []project.Module) []moduleRow {
	rows := make([]moduleRow, 0, len(modules))
	for _, m := range modules {
		z := m.Zones
		rows = append(rows, moduleRow{m.Name, m.IDPrefix, m.Dir, array(m.DependsOn), zoneRow{
			Protected: array(z.Protected), Evals: array(z.Evals), Interface: array(z.Interface),
			AppendOnly: array(z.AppendOnly), Preserve: array(z.Preserve), Disposable: array(z.Disposable),
		}})
	}
	return rows
}

func (m moduleRow) text() string {
	z := m.Zones
	return fmt.Sprintf("%s (%s) %s depends_on=%v protected=%v evals=%v interface=%v append_only=%v preserve=%v disposable=%v",
		m.Name, m.IDPrefix, m.Dir, m.DependsOn, z.Protected, z.Evals, z.Interface, z.AppendOnly, z.Preserve, z.Disposable)
}

func modules(root string, approved *workflow.ApprovedSnapshot) response {
	var snapshot project.Snapshot
	var err error
	if approved != nil {
		snapshot = approved.Snapshot
		err = validateGraph(snapshot.Modules)
	} else {
		snapshot, err = loadGraph(root)
	}
	if err != nil {
		return configFailure(err)
	}
	rows := moduleRows(snapshot.Modules)
	var plain strings.Builder
	for _, m := range rows {
		fmt.Fprintln(&plain, m.text())
	}
	return output(plain.String(), rows)
}

type queueRow struct {
	Slug        string `json:"slug"`
	Path        string `json:"path"`
	Module      string `json:"module"`
	Phase       string `json:"phase"`
	Blocked     string `json:"blocked"`
	AuditRounds int    `json:"audit_rounds"`
}

func makeQueueRow(slug, path, module, phase, blocked string, rounds int) queueRow {
	if phase == "" {
		phase = "queued"
	}
	return queueRow{slug, path, module, phase, blocked, rounds}
}

func (q queueRow) text() string {
	return fmt.Sprintf("%s %s module=%s phase=%s blocked=%q audit_rounds=%d",
		q.Slug, q.Path, q.Module, q.Phase, q.Blocked, q.AuditRounds)
}

type statusQueueRow struct {
	queueRow
	WaitingOn   string `json:"waiting_on"`
	WaitingRule string `json:"waiting_rule"`
}

type quarantineRow struct {
	Module string `json:"module"`
	ID     string `json:"id"`
	Reason string `json:"reason"`
}

type optOutRow struct {
	Module string `json:"module"`
	Reason string `json:"reason"`
}

type metricRow struct {
	Module          string         `json:"module"`
	Seconds         float64        `json:"seconds"`
	CasesPerSpec    map[string]int `json:"cases_per_spec"`
	EvalLines       int            `json:"eval_lines"`
	DisposableLines int            `json:"disposable_lines"`
}

type statusData struct {
	Modules         []moduleRow      `json:"modules"`
	Queue           []statusQueueRow `json:"queue"`
	Quarantines     []quarantineRow  `json:"quarantines"`
	BoundaryOptOuts []optOutRow      `json:"boundary_opt_outs"`
	DriftSinceDrill map[string]int   `json:"drift_since_drill"`
	Unapproved      int              `json:"unapproved"`
	PhaseCommits    map[string]int   `json:"phase_commits"`
	Metrics         []metricRow      `json:"metrics"`
}

func counts(values map[string]int) map[string]int {
	if values == nil {
		return map[string]int{}
	}
	return values
}

func status(root string) response {
	if _, err := project.Load(root); err != nil {
		return configFailure(err)
	}
	report, err := workflow.Status(root)
	if err != nil {
		result := configuredOperationFailure(err, "Correct the reported project state and check that the configured base ref is available locally.")
		result.diagnostics = append(result.diagnostics, report.Diagnostics...)
		return result
	}

	data := statusData{
		Modules:         moduleRows(report.Modules),
		Queue:           make([]statusQueueRow, 0, len(report.Queue)),
		Quarantines:     make([]quarantineRow, 0, len(report.Quarantines)),
		BoundaryOptOuts: make([]optOutRow, 0, len(report.BoundaryOptOuts)),
		DriftSinceDrill: counts(report.DriftSinceDrill),
		Unapproved:      report.Unapproved,
		PhaseCommits:    counts(report.PhaseCommits),
		Metrics:         make([]metricRow, 0, len(report.Costs)),
	}
	var plain strings.Builder
	if len(report.Queue) == 0 {
		fmt.Fprintln(&plain, "Queue is empty.")
	}
	for _, entry := range report.Queue {
		row := makeQueueRow(entry.Slug, entry.Path, entry.Module, entry.Phase, entry.Blocked, entry.AuditRounds)
		fmt.Fprintf(&plain, "%s waiting_on=%q waiting_rule=%q\n", row.text(), entry.WaitingOn, entry.WaitingRule)
		data.Queue = append(data.Queue, statusQueueRow{row, entry.WaitingOn, entry.WaitingRule})
	}
	for _, m := range data.Modules {
		fmt.Fprintln(&plain, m.text())
	}
	for _, q := range report.Quarantines {
		data.Quarantines = append(data.Quarantines, quarantineRow(q))
		fmt.Fprintf(&plain, "Quarantine %s %s: %s\n", q.Module, q.ID, q.Reason)
	}
	for _, optOut := range report.BoundaryOptOuts {
		data.BoundaryOptOuts = append(data.BoundaryOptOuts, optOutRow(optOut))
		fmt.Fprintf(&plain, "Boundary opt-out %s: %s\n", optOut.Module, optOut.Reason)
	}
	for _, name := range slices.Sorted(maps.Keys(data.DriftSinceDrill)) {
		fmt.Fprintf(&plain, "Drift since drill %s: %d\n", name, data.DriftSinceDrill[name])
	}
	fmt.Fprintf(&plain, "Unapproved changes: %d\n", data.Unapproved)
	for _, change := range slices.Sorted(maps.Keys(data.PhaseCommits)) {
		fmt.Fprintf(&plain, "Phase commits %s: %d\n", change, data.PhaseCommits[change])
	}
	for _, cost := range report.Costs {
		data.Metrics = append(data.Metrics, metricRow{cost.Module, cost.Seconds, counts(cost.CasesPerSpec), cost.EvalLines, cost.DisposableLines})
		fmt.Fprintf(&plain, "Cost %s: %gs cases_per_spec=%v eval_lines=%d disposable_lines=%d\n", cost.Module, cost.Seconds, cost.CasesPerSpec, cost.EvalLines, cost.DisposableLines)
	}
	result := output(plain.String(), data)
	result.diagnostics = report.Diagnostics
	return result
}
