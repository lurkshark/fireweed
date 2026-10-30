package lifecycle

import (
	"fmt"
	"path"
	"sort"
	"strconv"
	"strings"

	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/gitstore"
	"github.com/lurkshark/fireweed/packages/project"
	"github.com/lurkshark/fireweed/packages/workflow/model"
)

type prFile struct{ path, status string }
type prCommit struct {
	oid    string
	parent string
	files  []prFile
}
type prChange struct {
	path, name, module, kind string
	doc                      contracts.Document
	commits                  []prCommit
}

func PRBody(root, revisionRange string) (model.PRBodyReport, error) {
	parts := strings.Split(revisionRange, "..")
	if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
		return model.PRBodyReport{}, fmt.Errorf("P01: invalid range %q", revisionRange)
	}
	var resolved [2]string
	for i, part := range parts {
		out, err := commitsGit(root, "rev-parse", "--verify", "--end-of-options", part+"^{commit}")
		if err != nil {
			return model.PRBodyReport{}, fmt.Errorf("P01: resolve range %q: %w", revisionRange, err)
		}
		resolved[i] = strings.TrimSpace(out)
	}
	endpoint, err := project.LoadAt(root, resolved[1])
	if err != nil {
		return model.PRBodyReport{}, fmt.Errorf("load endpoint project at %s: %w", resolved[1], err)
	}
	root = endpoint.Root
	out, err := commitsGit(root, "rev-list", resolved[0]+".."+resolved[1])
	if err != nil {
		return model.PRBodyReport{}, err
	}
	inRange := map[string]bool{}
	for _, oid := range strings.Fields(out) {
		inRange[oid] = true
	}
	paths, err := commitChangePaths(root, resolved[1])
	if err != nil {
		return model.PRBodyReport{}, err
	}
	byName := map[string]*prChange{}
	for _, p := range paths {
		data, err := gitstore.ReadAt(root, resolved[1], p)
		if err != nil {
			return model.PRBodyReport{}, fmt.Errorf("read endpoint %s: %w", p, err)
		}
		doc, err := contracts.ParseDocument(string(data))
		if err != nil {
			return model.PRBodyReport{}, fmt.Errorf("parse endpoint %s: %w", p, err)
		}
		name := path.Base(p)
		byName[name] = &prChange{path: p, name: name, module: str(doc.Frontmatter["module"]), kind: prKind(doc), doc: doc}
	}
	selected := map[string]bool{}
	out = ""
	if len(inRange) > 0 && len(byName) > 0 {
		// One ancestry traversal supplies retained phase context for all identities.
		out, err = commitsGit(root, "rev-list", "--topo-order", "--parents", resolved[1])
		if err != nil {
			return model.PRBodyReport{}, err
		}
	}
	for _, line := range strings.Split(strings.TrimSpace(out), "\n") {
		fields := strings.Fields(line)
		if len(fields) == 0 {
			continue
		}
		parent := ""
		if len(fields) > 1 {
			parent = fields[1]
		}
		files, err := prCommitFiles(root, fields[0], parent)
		if err != nil {
			return model.PRBodyReport{}, err
		}
		touched := map[string]bool{}
		for _, f := range files {
			if (path.Dir(f.path) == "changes" || path.Dir(f.path) == "changes/archive") && commitChangeName.MatchString(path.Base(f.path)) {
				touched[path.Base(f.path)] = true
			}
		}
		for name := range touched {
			if change := byName[name]; change != nil {
				if inRange[fields[0]] {
					selected[name] = true
				}
				if parent != "" || inRange[fields[0]] {
					change.commits = append(change.commits, prCommit{fields[0], parent, files})
				}
			}
		}
	}
	var changes []*prChange
	for name := range selected {
		changes = append(changes, byName[name])
	}
	sort.Slice(changes, func(i, j int) bool { return changes[i].path < changes[j].path })
	var body strings.Builder
	fmt.Fprintf(&body, "# PR review: %s..%s\n\n## Changes\n\n", resolved[0], resolved[1])
	if len(changes) == 0 {
		body.WriteString("None selected.\n\n")
	}
	cache := map[string]project.Snapshot{resolved[1]: endpoint}
	for _, change := range changes {
		if err := prRenderChange(&body, root, change, cache); err != nil {
			return model.PRBodyReport{}, err
		}
	}
	if err := prExceptions(&body, endpoint, paths, byName); err != nil {
		return model.PRBodyReport{}, err
	}
	result := model.PRBodyReport{}
	if err := prCosts(&body, root, endpoint, &result); err != nil {
		return model.PRBodyReport{}, err
	}
	if err := prArtifacts(&body, root, resolved[0], resolved[1], endpoint, &result); err != nil {
		return model.PRBodyReport{}, err
	}
	result.Body = body.String()
	return result, nil
}

func prCommitFiles(root, oid, parent string) ([]prFile, error) {
	args := []string{"diff-tree", "--no-commit-id", "-r", "-z", "--name-status", "--no-renames"}
	if parent == "" {
		args = append(args, "--root", oid)
	} else {
		args = append(args, parent, oid)
	}
	out, err := commitsGit(root, args...)
	if err != nil {
		return nil, err
	}
	var files []prFile
	records := strings.Split(strings.TrimSuffix(out, "\x00"), "\x00")
	if len(records) == 1 && records[0] == "" {
		return files, nil
	}
	if len(records)%2 != 0 {
		return nil, fmt.Errorf("malformed diff-tree entries at %s", oid)
	}
	for i := 0; i < len(records); i += 2 {
		files = append(files, prFile{records[i+1], records[i]})
	}
	return files, nil
}

func prKind(doc contracts.Document) string {
	if doc.Frontmatter["regenerate"] == true {
		return "regenerate"
	}
	if _, ok := doc.Frontmatter["eval_fix"]; ok {
		return "eval_fix"
	}
	if len(list(doc.Frontmatter["fixes"])) > 0 {
		return "fix"
	}
	for _, section := range []string{"Added", "Modified", "Removed"} {
		content, _ := contracts.Section(doc.Body, section)
		items, _ := contracts.ParseItems(content)
		if len(items) > 0 {
			return "spec"
		}
	}
	return "maintenance"
}

func prPhase(root string, commit prCommit, change *prChange, doc contracts.Document) (string, error) {
	phase := str(doc.Frontmatter["phase"])
	if commit.parent == "" {
		return phase, nil
	}
	previous, err := commitPhaseAt(root, commit.parent, change.name)
	if err != nil {
		return "", err
	}
	if previous != "" {
		return previous, nil
	}
	// Start writes runtime metadata without committing it. The first Done
	// therefore has an authored parent with no phase recorded in Git.
	if phase == "audit" || phase == "impl" {
		switch prKind(doc) {
		case "spec", "fix", "eval_fix":
			return "evals", nil
		}
	}
	return phase, nil
}

func prZones(module project.Module, doc contracts.Document) contracts.Zones {
	zones := module.Zones
	set := mapval(doc.Frontmatter["set"])
	for key, raw := range set {
		values := list(raw)
		switch key {
		case "evals":
			zones.Evals = values
		case "interface":
			zones.Interface = values
		case "append_only":
			zones.AppendOnly = values
		case "preserve":
			zones.Preserve = values
		case "disposable":
			zones.Disposable = values
		}
	}
	return zones
}

func prRootProtected(snapshot project.Snapshot, file string) bool {
	if path.Base(file) == "AGENTS.md" || path.Base(file) == "CLAUDE.md" {
		return true
	}
	for _, pattern := range append([]string{".git", ".git/**", ".fireweed", ".fireweed/**", "SYSTEM.md", "fireweed.yaml"}, list(snapshot.Config["protected"])...) {
		if contracts.MatchGlob(pattern, file) {
			return true
		}
	}
	return false
}

func prRenderChange(body *strings.Builder, root string, change *prChange, snapshots map[string]project.Snapshot) error {
	fmt.Fprintf(body, "### %s\n\nModule: %s; kind: %s\n\n", change.path, change.module, change.kind)
	var provenance, findings []string
	seenProvenance, seenFinding := map[string]bool{}, map[string]bool{}
	rounds := num(change.doc.Frontmatter["audit_rounds"], 0)
	categories := map[string]map[string]bool{}
	for i := len(change.commits) - 1; i >= 0; i-- {
		commit := change.commits[i]
		p := "changes/" + change.name
		data, err := gitstore.ReadAt(root, commit.oid, p)
		if err != nil {
			p = "changes/archive/" + change.name
			data, err = gitstore.ReadAt(root, commit.oid, p)
			if err != nil {
				continue
			}
		}
		doc, err := contracts.ParseDocument(string(data))
		if err != nil {
			return fmt.Errorf("parse %s at %s: %w", p, commit.oid, err)
		}
		if n := num(doc.Frontmatter["audit_rounds"], 0); n > rounds {
			rounds = n
		}
		for _, section := range []struct {
			name  string
			lines *[]string
			seen  map[string]bool
		}{
			{"Provenance", &provenance, seenProvenance}, {"Audit", &findings, seenFinding},
		} {
			content, _ := contracts.Section(doc.Body, section.name)
			for _, line := range strings.Split(content, "\n") {
				line = strings.TrimSpace(line)
				if line == "" || section.name == "Audit" && (line == "No findings" || strings.HasPrefix(line, "Checked IDs:")) || section.seen[line] {
					continue
				}
				section.seen[line] = true
				*section.lines = append(*section.lines, line)
			}
		}
		snapshot, ok := snapshots[commit.oid]
		if !ok {
			snapshot, err = project.LoadAt(root, commit.oid)
			if err != nil {
				return fmt.Errorf("load phase project at %s: %w", commit.oid, err)
			}
			snapshots[commit.oid] = snapshot
		}
		var module project.Module
		for _, candidate := range snapshot.Modules {
			if candidate.Name == change.module {
				module = candidate
				break
			}
		}
		phase, err := prPhase(root, commit, change, doc)
		if err != nil {
			return fmt.Errorf("read prior phase for %s: %w", change.name, err)
		}
		zones := prZones(module, doc)
		for _, f := range commit.files {
			category := ""
			if prRootProtected(snapshot, f.path) {
				continue
			}
			if module.Dir != "" && strings.HasPrefix(f.path, module.Dir+"/") {
				rel := strings.TrimPrefix(f.path, module.Dir+"/")
				switch contracts.ResolveZone(rel, zones) {
				case "evals", "interface", "preserve":
					category = contracts.ResolveZone(rel, zones)
				case "disposable":
					if phase == "evals" && f.status == "A" {
						category = "new eval-step stub"
					} else {
						category = "disposable"
					}
				case "other":
					category = "manifest/other"
				}
			} else {
				for _, pattern := range list(snapshot.Config["worker_writable"]) {
					if contracts.MatchGlob(pattern, f.path) {
						category = "shared worker_writable"
						break
					}
				}
			}
			if category != "" {
				if categories[category] == nil {
					categories[category] = map[string]bool{}
				}
				categories[category][f.path] = true
			}
		}
	}
	fmt.Fprintf(body, "Audit rounds: %d\n\n", rounds)
	for _, lines := range []struct {
		header string
		values []string
	}{{"Provenance", provenance}, {"Audit findings", findings}} {
		if len(lines.values) == 0 {
			continue
		}
		body.WriteString(lines.header + ":\n")
		for _, line := range lines.values {
			body.WriteString(line + "\n")
		}
		body.WriteString("\n")
	}
	for _, category := range []string{"evals", "interface", "new eval-step stub", "disposable", "manifest/other", "shared worker_writable", "preserve"} {
		if len(categories[category]) == 0 {
			continue
		}
		body.WriteString(strings.ToUpper(category[:1]) + category[1:] + " edits:\n")
		var names []string
		for name := range categories[category] {
			names = append(names, name)
		}
		sort.Strings(names)
		for _, name := range names {
			body.WriteString("- " + name + "\n")
		}
		body.WriteString("\n")
	}
	for _, name := range []string{"Added", "Modified"} {
		section, _ := contracts.Section(change.doc.Body, name)
		items, err := contracts.ParseItems(section)
		if err != nil {
			return fmt.Errorf("%s %s: %w", change.path, name, err)
		}
		for _, item := range items {
			if item.Tag == "review" {
				fmt.Fprintf(body, "Review %s %s: %s\n\n", name, item.ID, strings.TrimSpace(item.Body))
			}
		}
	}
	for _, name := range []string{"Decisions", "Removed decisions"} {
		section, _ := contracts.Section(change.doc.Body, name)
		if strings.TrimSpace(section) != "" {
			fmt.Fprintf(body, "%s:\n%s\n\n", name, strings.TrimSpace(section))
		}
	}
	return nil
}

func prExceptions(body *strings.Builder, endpoint project.Snapshot, paths []string, changes map[string]*prChange) error {
	body.WriteString("## Endpoint exceptions\n\n")
	for _, module := range endpoint.Modules {
		section, _ := contracts.Section(module.Document.Body, "Specs")
		items, err := contracts.ParseItems(section)
		if err != nil {
			return fmt.Errorf("%s/SPEC.md: %w", module.Dir, err)
		}
		for _, item := range items {
			if reason, ok := strings.CutPrefix(item.Tag, "quarantine:"); ok {
				fmt.Fprintf(body, "- Quarantine %s %s: %s\n", module.Name, item.ID, strings.TrimSpace(reason))
			}
		}
		if reason := str(module.Document.Frontmatter["boundaries_opt_out"]); reason != "" {
			fmt.Fprintf(body, "- Boundary opt-out %s: %s\n", module.Name, reason)
		}
	}
	for _, p := range paths {
		if path.Dir(p) != "changes" {
			continue
		}
		change := changes[path.Base(p)]
		if change != nil {
			if reason := str(change.doc.Frontmatter["blocked"]); reason != "" {
				fmt.Fprintf(body, "- Blocked %s (%s): %s\n", p, change.module, reason)
			}
		}
	}
	body.WriteString("\n")
	return nil
}

func prCosts(body *strings.Builder, root string, endpoint project.Snapshot, result *model.PRBodyReport) error {
	snapshots, err := Costs(root)
	if err != nil {
		return err
	}
	modules := map[string]bool{}
	for _, module := range endpoint.Modules {
		modules[module.Name] = true
	}
	selected := map[string]int{}
	for i, snapshot := range snapshots {
		for _, cost := range snapshot.Modules {
			if modules[cost.Module] {
				selected[cost.Module] = i
			}
		}
	}
	body.WriteString("## Costs\n\n| Module | Seconds | Cases per spec | Eval lines | Disposable lines |\n| --- | ---: | --- | ---: | ---: |\n")
	var names []string
	for name := range selected {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		for _, cost := range snapshots[selected[name]].Modules {
			if cost.Module != name {
				continue
			}
			var ids, counts []string
			for id := range cost.CasesPerSpec {
				ids = append(ids, id)
			}
			sort.Strings(ids)
			for _, id := range ids {
				counts = append(counts, id+": "+strconv.Itoa(cost.CasesPerSpec[id]))
			}
			fmt.Fprintf(body, "| %s | %g | %s | %d | %d |\n", name, cost.Seconds, strings.Join(counts, ", "), cost.EvalLines, cost.DisposableLines)
			break
		}
	}
	body.WriteString("\n")
	seen := map[project.Diagnostic]bool{}
	for i, snapshot := range snapshots {
		for _, warning := range snapshot.Diagnostics {
			if selectedCostWarning(i, snapshot, selected, warning) && !seen[warning] {
				seen[warning] = true
				result.Diagnostics = append(result.Diagnostics, warning)
				fmt.Fprintf(body, "- %s: %s\n", warning.Code, warning.Message)
			}
		}
	}
	body.WriteString("\n")
	return nil
}

func prArtifacts(body *strings.Builder, root, base, endpoint string, snapshot project.Snapshot, result *model.PRBodyReport) error {
	out, err := commitsGit(root, "diff", "--numstat", "-z", "--diff-filter=A", "--no-renames", base, endpoint)
	if err != nil {
		return err
	}
	var warnings []string
	for _, record := range strings.Split(out, "\x00") {
		if record == "" {
			continue
		}
		fields := strings.SplitN(record, "\t", 3)
		if len(fields) != 3 {
			return fmt.Errorf("malformed diff numstat for %s..%s", base, endpoint)
		}
		p := fields[2]
		if prRootProtected(snapshot, p) {
			continue
		}
		var module project.Module
		var rel string
		for _, candidate := range snapshot.Modules {
			if strings.HasPrefix(p, candidate.Dir+"/") {
				module = candidate
				rel = strings.TrimPrefix(p, candidate.Dir+"/")
				break
			}
		}
		if module.Dir == "" || contracts.ResolveZone(rel, module.Zones) != "disposable" {
			continue
		}
		data, err := gitstore.ReadAt(root, endpoint, p)
		if err != nil {
			return fmt.Errorf("read added artifact %s: %w", p, err)
		}
		if fields[0] != "-" && len(data) <= 262144 {
			continue
		}
		message := fmt.Sprintf("%s: new disposable %s, %d bytes", module.Name, p, len(data))
		result.Diagnostics = append(result.Diagnostics, project.Diagnostic{Code: "B04", Message: message, Path: p, Warning: true})
		warnings = append(warnings, "- B04: "+message)
	}
	body.WriteString("## Disposable artifacts\n\n")
	if len(warnings) == 0 {
		body.WriteString("None.\n\n")
	} else {
		sort.Strings(warnings)
		body.WriteString(strings.Join(warnings, "\n") + "\n\n")
	}
	return nil
}
