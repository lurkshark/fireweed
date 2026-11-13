package lifecycle

import (
	"errors"
	"fmt"
	"maps"
	"path"
	"slices"
	"strconv"
	"strings"

	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/project"
	"github.com/lurkshark/fireweed/packages/workflow/model"
)

type endpointConfigError struct{ cause error }

func (e endpointConfigError) Error() string { return e.cause.Error() }
func (e endpointConfigError) Unwrap() error { return e.cause }

func IsEndpointConfigError(err error) bool {
	var endpoint endpointConfigError
	return errors.As(err, &endpoint)
}

type commitRange struct {
	base, end  string
	expression string
}

func resolveRange(root, expression string) (commitRange, error) {
	var revisions commitRange
	left, right, ok := strings.Cut(expression, "..")
	triple := ok && strings.HasPrefix(right, ".")
	if triple {
		right = strings.TrimPrefix(right, ".")
	}
	if !ok || left == "" || right == "" || strings.Contains(right, "..") {
		return revisions, failure("P01", expression, "invalid commit range %q", expression)
	}
	resolve := func(revision string) (string, error) {
		oid, err := git(root, "rev-parse", "--verify", "--end-of-options", revision+"^{commit}")
		if err != nil {
			return "", coded("P01", expression, fmt.Errorf("resolve %q in range %q: %w", revision, expression, err))
		}
		return strings.TrimSpace(oid), nil
	}
	base, err := resolve(left)
	if err != nil {
		return revisions, err
	}
	end, err := resolve(right)
	if err != nil {
		return revisions, err
	}
	separator := ".."
	if triple {
		separator = "..."
	}
	return commitRange{base: base, end: end, expression: base + separator + end}, nil
}

func PRBody(root, expression string) (model.PRBodyReport, error) {
	var report model.PRBodyReport
	revisions, err := resolveRange(root, expression)
	if err != nil {
		return report, err
	}
	snapshot, err := project.LoadAt(root, revisions.end)
	if err != nil {
		return report, endpointConfigError{err}
	}
	root = snapshot.Root
	if err := project.ValidateGraph(snapshot.Modules); err != nil {
		return report, endpointConfigError{coded("C02", "", err)}
	}
	touched, err := touchedChanges(root, revisions.expression)
	if err != nil {
		return report, err
	}
	var pending, archived []project.Change
	for _, entry := range []struct {
		dir  string
		dest *[]project.Change
	}{{pendingDir, &pending}, {archiveDir, &archived}} {
		paths, err := changePaths(root, revisions.end, entry.dir)
		if err != nil {
			return report, err
		}
		for _, identity := range paths {
			change, err := readChange(root, revisions.end, identity)
			if err != nil {
				return report, err
			}
			*entry.dest = append(*entry.dest, change)
		}
	}

	var body strings.Builder
	fmt.Fprintf(&body, "# Fireweed changes\n\nRange `%s`\n\n## Selected changes\n", revisions.expression)
	selected := 0
	for _, change := range append(pending, archived...) {
		if !touched[path.Base(change.Path)] {
			continue
		}
		if err := writeSelectedChange(&body, snapshot, revisions.end, change); err != nil {
			return report, err
		}
		selected++
	}
	if selected == 0 {
		body.WriteString("\nNo selected changes.\n")
	}
	quarantines, optOuts, err := exceptions(snapshot)
	if err != nil {
		return report, endpointConfigError{err}
	}
	writeExceptions(&body, quarantines, optOuts, pending)
	costs, budgets, err := latestCosts(root, snapshot.Modules, true)
	if err != nil {
		return report, err
	}
	writeCosts(&body, costs)
	artifacts, err := artifactWarnings(snapshot, revisions)
	if err != nil {
		return report, err
	}
	report.Diagnostics = append(budgets, artifacts...)
	if len(report.Diagnostics) > 0 {
		body.WriteString("\n## Warnings\n")
		for _, d := range report.Diagnostics {
			fmt.Fprintf(&body, "- %s %s\n", d.Code, d.Message)
		}
	}
	report.Body = body.String()
	return report, nil
}

func touchedChanges(root, expression string) (map[string]bool, error) {
	out, err := git(root, "rev-list", "--parents", expression)
	if err != nil {
		return nil, err
	}
	touched := make(map[string]bool)
	for _, line := range strings.Split(out, "\n") {
		fields := strings.Fields(line)
		if len(fields) == 0 {
			continue
		}
		parent := ""
		if len(fields) > 1 {
			parent = fields[1]
		}
		files, err := revisionChanges(root, parent, fields[0], "changes/")
		if err != nil {
			return nil, err
		}
		for _, change := range files {
			if (path.Dir(change.file) == pendingDir || path.Dir(change.file) == archiveDir) && changeIdentity(change.file).Slug != "" {
				touched[path.Base(change.file)] = true
			}
		}
	}
	return touched, nil
}

type revisionChange struct{ status, file string }

func revisionChanges(root, parent, oid string, paths ...string) ([]revisionChange, error) {
	args := []string{"diff-tree", "-r", "--no-commit-id", "--no-renames", "--name-status", "-z"}
	if parent == "" {
		args = append(args, "--root", oid)
	} else {
		args = append(args, parent, oid)
	}
	out, err := git(root, append(append(args, "--"), paths...)...)
	if err != nil || out == "" {
		return nil, err
	}
	fields := gitFields(out)
	if len(fields)%2 != 0 {
		return nil, errors.New("invalid git name-status output")
	}
	var result []revisionChange
	for i := 0; i < len(fields); i += 2 {
		result = append(result, revisionChange{fields[i], fields[i+1]})
	}
	return result, nil
}

func writeSelectedChange(body *strings.Builder, snapshot project.Snapshot, endpoint string, change project.Change) error {
	fmt.Fprintf(body, "\n### %s\n\n- Module %s\n- Kind %s\n- audit_rounds %d\n", change.Path, change.Module,
		changeKind(change), integer(change.Document.Frontmatter["audit_rounds"], 0))
	if provenance := strings.TrimSpace(section(change.Document, "Provenance")); provenance != "" {
		fmt.Fprintf(body, "\n**Provenance**\n\n%s\n", provenance)
	}
	for _, spec := range slices.Concat(change.Added, change.Modified) {
		if strings.EqualFold(strings.TrimSpace(spec.Tag), "review") {
			fmt.Fprintf(body, "\n- Review %s: %s\n", spec.ID, strings.TrimSpace(spec.Body))
		}
	}
	for _, name := range []string{"Decisions", "Removed decisions"} {
		if content := strings.TrimSpace(section(change.Document, name)); content != "" {
			fmt.Fprintf(body, "\n**%s**\n\n%s\n", name, content)
		}
	}

	commits, err := phaseCommits(snapshot.Root, endpoint, change)
	if err != nil {
		return err
	}
	categories := make(map[string]map[string]bool)
	findings := make(map[string]bool)
	recordFindings := func(doc contracts.Document) {
		for _, line := range strings.Split(section(doc, "Audit"), "\n") {
			line = strings.TrimSpace(line)
			if findingPattern.MatchString(line) {
				findings[line] = true
			}
		}
	}
	recordFindings(change.Document)
	for _, commit := range commits {
		recordFindings(commit.after.Document)
		files, err := revisionChanges(snapshot.Root, commit.parent, commit.oid)
		if err != nil {
			return err
		}
		phaseSnapshot, err := project.LoadAt(snapshot.Root, commit.oid)
		if err != nil {
			return err
		}
		module := moduleNamed(phaseSnapshot, change.Module)
		if module.Name == "" {
			continue
		}
		module.Zones = projectedZones(module.Zones, mapping(commit.after.Document.Frontmatter["set"]))
		evals := completedEvals(commit)
		for _, file := range files {
			zone := fileZone(phaseSnapshot, module, file.file)
			category := ""
			switch zone {
			case "evals":
				category = "Eval files"
			case "interface":
				category = "Interface files"
			case "disposable":
				if evals && file.status == "A" {
					category = "New disposable stubs"
				}
			case "other", "spec":
				category = "Module manifests and other files"
			case "worker_writable":
				category = "worker_writable files"
			case "preserve":
				category = "Preserved files"
			}
			if category != "" {
				if categories[category] == nil {
					categories[category] = make(map[string]bool)
				}
				categories[category][file.file] = true
			}
		}
	}
	for _, category := range slices.Sorted(maps.Keys(categories)) {
		fmt.Fprintf(body, "\n**%s**\n\n", category)
		for _, file := range slices.Sorted(maps.Keys(categories[category])) {
			fmt.Fprintf(body, "- %s\n", file)
		}
	}
	if len(findings) != 0 {
		fmt.Fprint(body, "\n**Audit findings**\n\n")
		for _, finding := range slices.Sorted(maps.Keys(findings)) {
			fmt.Fprintln(body, finding)
		}
	}
	return nil
}

func completedEvals(commit changeCommit) bool {
	before := str(commit.before.Document.Frontmatter["phase"])
	after := str(commit.after.Document.Frontmatter["phase"])
	switch before {
	case "evals":
		return after == "audit" || after == "impl"
	case "":
		return after == "audit" || (after == "impl" || after == "applied") && initialPhase(commit.after) == "evals"
	default:
		return after == "applied" && before != "impl" && initialPhase(commit.after) == "evals"
	}
}

func projectedZones(zones contracts.Zones, set map[string]any) contracts.Zones {
	for key, raw := range set {
		switch key {
		case "evals":
			zones.Evals = stringsOf(raw)
		case "interface":
			zones.Interface = stringsOf(raw)
		case "append_only":
			zones.AppendOnly = stringsOf(raw)
		case "preserve":
			zones.Preserve = stringsOf(raw)
		case "disposable":
			zones.Disposable = stringsOf(raw)
		}
	}
	return zones
}

func changeKind(change project.Change) string {
	if regenerate, _ := change.Document.Frontmatter["regenerate"].(bool); regenerate {
		return "regenerate"
	}
	if evalFix(change) {
		return "eval_fix"
	}
	if len(stringsOf(change.Document.Frontmatter["fixes"])) > 0 {
		return "fix"
	}
	if len(change.Added)+len(change.Modified)+len(change.Removed) != 0 {
		return "spec"
	}
	return "maintenance"
}

func writeCosts(body *strings.Builder, costs []model.ModuleCost) {
	if len(costs) == 0 {
		return
	}
	body.WriteString("\n## Costs\n\n| Module | Seconds | Cases per spec | Eval lines | Disposable lines |\n")
	body.WriteString("| --- | ---: | --- | ---: | ---: |\n")
	for _, cost := range costs {
		var cases []string
		for _, id := range slices.Sorted(maps.Keys(cost.CasesPerSpec)) {
			cases = append(cases, fmt.Sprintf("%s: %d", id, cost.CasesPerSpec[id]))
		}
		fmt.Fprintf(body, "| %s | %g | %s | %d | %d |\n", cost.Module, cost.Seconds,
			strings.Join(cases, ", "), cost.EvalLines, cost.DisposableLines)
	}
}

func artifactWarnings(snapshot project.Snapshot, revisions commitRange) ([]project.Diagnostic, error) {
	root := snapshot.Root
	out, err := git(root, "diff", "--numstat", "-z", "--no-renames", "--no-ext-diff", "--diff-filter=A",
		revisions.base, revisions.end, "--")
	if err != nil {
		return nil, err
	}
	var warnings []project.Diagnostic
	for _, record := range gitFields(out) {
		fields := strings.SplitN(record, "\t", 3)
		if len(fields) != 3 {
			continue
		}
		file := fields[2]
		if fileZone(snapshot, ownerOf(snapshot, file), file) != "disposable" {
			continue
		}
		raw, err := git(root, "cat-file", "-s", revisions.end+":"+file)
		if err != nil {
			return nil, err
		}
		size, err := strconv.ParseInt(strings.TrimSpace(raw), 10, 64)
		if err != nil {
			return nil, err
		}
		if size > 262144 || fields[0] == "-" && fields[1] == "-" {
			warnings = append(warnings, warning("B04", file, "%s (%d bytes) is a binary or large disposable addition", file, size))
		}
	}
	return warnings, nil
}
