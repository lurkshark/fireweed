package drill

import (
	"fmt"
	"path/filepath"
	"strings"
	"time"

	"github.com/lurkshark/fireweed/packages/checks"
	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/project"
	"github.com/lurkshark/fireweed/packages/runner/internal/invoke"
	"github.com/lurkshark/fireweed/packages/runner/internal/workspace"
	"github.com/lurkshark/fireweed/packages/runner/model"
)

func configuredModel(config map[string]any) string {
	if name, _ := invoke.Mapping(config["agent"])["model"].(string); name != "" {
		return name
	}
	command := invoke.Command(config, "worker")
	args, _ := invoke.Strings(command)
	if shell, ok := command.(string); ok {
		args = strings.Fields(shell)
	}
	for i, arg := range args {
		switch {
		case (arg == "-m" || arg == "--model") && i+1 < len(args):
			return args[i+1]
		case strings.HasPrefix(arg, "--model="):
			return strings.TrimPrefix(arg, "--model=")
		}
	}
	return "unknown"
}

func writeReport(root string, now time.Time, result *model.DrillResult, dependents []project.Module, notes string, failures []string) error {
	var body strings.Builder
	body.WriteString("\n## Own evals\n")
	evalEntries(&body, result.Module, result.Own)
	body.WriteString("\n## Dependent evals\n")
	for _, module := range dependents {
		evalEntries(&body, module.Name, result.Dependents[module.Name])
	}
	for _, heading := range []string{"Contradicted Decisions", "Gaps"} {
		section, _ := contracts.Section(notes, heading)
		fmt.Fprintf(&body, "\n## %s\n%s\n", heading, strings.TrimSpace(section))
	}
	body.WriteString("\n## Agent observations\n")
	for i, step := range result.Steps {
		rel, err := filepath.Rel(root, step.LogPath)
		if err != nil {
			rel = ""
		}
		fmt.Fprintf(&body, "\n### Invocation %d\nLog: %s\n\n%s\n", i+1, filepath.ToSlash(rel), step.Stdout)
	}
	if len(failures) != 0 {
		fmt.Fprintf(&body, "\n## Failures\n%s\n", strings.Join(failures, "\n\n"))
	}
	report, err := contracts.RenderDocument(contracts.Document{
		Frontmatter: map[string]any{
			"date": now.Format(time.RFC3339Nano), "base_commit": result.Base,
			"agent": result.Agent, "model": result.Model, "result": result.Result,
		},
		Body: body.String(),
	})
	if err != nil {
		return err
	}
	path := ".fireweed/drills/" + result.Module + ".md"
	if err := workspace.Write(root, path, []byte(report), 0644); err != nil {
		return fmt.Errorf("write drill report: %w", err)
	}
	result.ReportPath = path
	return nil
}

func evalEntries(body *strings.Builder, module string, report checks.Report) {
	for _, c := range report.Cases {
		ids := strings.Join(c.IDs, ", ")
		if ids == "" {
			ids = "no spec"
		}
		fmt.Fprintf(body, "- %s / %s / %s / %s\n", module, c.Name, ids, c.Result)
		if c.Failure != "" {
			fmt.Fprintf(body, "\n%s\n\n", c.Failure)
		}
	}
}
