package lifecycle

import (
	"fmt"
	"os"
	"path"
	"sort"
	"strings"
	"time"

	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/gitstore"
	"github.com/lurkshark/fireweed/packages/project"
	"github.com/lurkshark/fireweed/packages/workflow/model"
)

type appliedChange struct {
	name, module string
	date         time.Time
}

func Status(root string) (model.StatusReport, error) {
	repo, err := gitstore.Inspect(root)
	if err != nil {
		return model.StatusReport{}, err
	}
	root = repo.Root
	snapshot, err := project.Load(root)
	if err != nil {
		return model.StatusReport{}, err
	}
	r := model.StatusReport{Modules: snapshot.Modules, DriftSinceDrill: map[string]int{}, PhaseCommits: map[string]int{}}
	for _, m := range snapshot.Modules {
		r.DriftSinceDrill[m.Name] = 0
		section, _ := contracts.Section(m.Document.Body, "Specs")
		items, err := contracts.ParseItems(section)
		if err != nil {
			return model.StatusReport{}, fmt.Errorf("%s/SPEC.md: %w", m.Dir, err)
		}
		for _, item := range items {
			if reason, ok := strings.CutPrefix(item.Tag, "quarantine:"); ok {
				r.Quarantines = append(r.Quarantines, model.Quarantine{Module: m.Name, ID: item.ID, Reason: strings.TrimSpace(reason)})
			}
		}
		if reason := str(m.Document.Frontmatter["boundaries_opt_out"]); reason != "" {
			r.BoundaryOptOuts = append(r.BoundaryOptOuts, model.BoundaryOptOut{Module: m.Name, Reason: reason})
		}
	}
	queue, err := project.Queue(root)
	if err != nil {
		return model.StatusReport{}, err
	}
	var blocked []project.Change
	for _, candidate := range queue {
		e := model.StatusEntry{Slug: candidate.Slug, Path: candidate.Path, Module: candidate.Module,
			Phase: str(candidate.Document.Frontmatter["phase"]), Blocked: str(candidate.Document.Frontmatter["blocked"]),
			AuditRounds: num(candidate.Document.Frontmatter["audit_rounds"], 0)}
		if e.Phase == "" {
			e.Phase = "queued"
		}
		if e.Blocked != "" {
			blocked = append(blocked, candidate)
		} else if len(blocked) > 0 {
			projected, err := project.Project(root, candidate.Slug)
			if err != nil {
				return model.StatusReport{}, err
			}
			graph := map[string][]string{}
			for _, m := range projected.Modules {
				graph[m.Name] = m.DependsOn
			}
			section, _ := contracts.Section(candidate.Document.Body, "Interface")
			for _, earlier := range blocked {
				if rule := blockingRule(graph, candidate.Module, earlier.Module, strings.TrimSpace(section) != ""); rule != "" {
					e.WaitingOn = strings.TrimSuffix(path.Base(earlier.Path), ".md")
					e.WaitingRule = rule
					break
				}
			}
		}
		r.Queue = append(r.Queue, e)
	}
	archives, err := appliedAtHead(root, repo.Head)
	if err != nil {
		return model.StatusReport{}, err
	}
	base, err := gitstore.BaseRef(root, str(snapshot.Config["remote"]), str(snapshot.Config["base"]))
	if err != nil {
		return model.StatusReport{}, err
	}
	basePaths, err := commitChangePaths(root, base)
	if err != nil {
		return model.StatusReport{}, err
	}
	known := map[string]bool{}
	for _, p := range basePaths {
		known[path.Base(p)] = true
	}
	drillDates := map[string]time.Time{}
	hasDrill := map[string]bool{}
	for _, m := range snapshot.Modules {
		data, err := os.ReadFile(file(root, ".fireweed/drills/"+m.Name+".md"))
		if err != nil && !os.IsNotExist(err) {
			return model.StatusReport{}, err
		}
		if err == nil {
			hasDrill[m.Name] = true
			doc, err := contracts.ParseDocument(string(data))
			if err != nil {
				return model.StatusReport{}, fmt.Errorf("drill %s: %w", m.Name, err)
			}
			drillDates[m.Name] = documentDate(doc.Frontmatter["date"])
		}
	}
	for _, change := range archives {
		if !known[change.name] {
			r.Unapproved++
		}
		if !hasDrill[change.module] || change.date.After(drillDates[change.module]) {
			r.DriftSinceDrill[change.module]++
		}
	}
	sort.Slice(archives, func(i, j int) bool {
		if !archives[i].date.Equal(archives[j].date) {
			return archives[i].date.After(archives[j].date)
		}
		return archives[i].name < archives[j].name
	})
	if len(archives) >= 3 {
		count := 1
		for count < len(archives) && archives[count].module == archives[0].module {
			count++
		}
		if count >= 3 {
			r.Diagnostics = append(r.Diagnostics, diag("P08", fmt.Sprintf("%s: %d consecutive applied changes", archives[0].module, count), true))
		}
	}
	r.PhaseCommits, err = phaseCommitCounts(root, repo.Head, archives)
	if err != nil {
		return model.StatusReport{}, err
	}
	if err := statusCosts(root, &r); err != nil {
		return model.StatusReport{}, err
	}
	return r, nil
}

func appliedAtHead(root, head string) ([]appliedChange, error) {
	paths, err := commitChangePaths(root, head)
	if err != nil {
		return nil, err
	}
	var changes []appliedChange
	for _, p := range paths {
		if path.Dir(p) != "changes/archive" {
			continue
		}
		data, err := gitstore.ReadAt(root, head, p)
		if err != nil {
			return nil, err
		}
		doc, err := contracts.ParseDocument(string(data))
		if err != nil {
			return nil, fmt.Errorf("%s: %w", p, err)
		}
		if str(doc.Frontmatter["phase"]) != "applied" {
			continue
		}
		provenance, _ := contracts.Section(doc.Body, "Provenance")
		var date time.Time
		for _, line := range strings.Split(provenance, "\n") {
			fields := strings.Fields(line)
			if len(fields) >= 3 && fields[0] == "-" && fields[1] == "impl" {
				if parsed := documentDate(fields[2]); parsed.After(date) {
					date = parsed
				}
			}
		}
		changes = append(changes, appliedChange{name: path.Base(p), module: str(doc.Frontmatter["module"]), date: date})
	}
	return changes, nil
}
