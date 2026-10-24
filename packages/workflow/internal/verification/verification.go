// Package verification composes read-only phase checking from approved authority.
package verification

import (
	"context"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"

	"github.com/lurkshark/fireweed/packages/checks"
	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/gitstore"
	"github.com/lurkshark/fireweed/packages/project"
	"github.com/lurkshark/fireweed/packages/workflow/internal/approval"
	"github.com/lurkshark/fireweed/packages/workflow/internal/claims"
	"github.com/lurkshark/fireweed/packages/workflow/internal/worktree"
	"github.com/lurkshark/fireweed/packages/workflow/model"
)

func Verify(ctx context.Context, start, slug string) (model.Verification, error) {
	zero := model.Verification{}
	if ctx == nil {
		return zero, fmt.Errorf("provide a non-nil verification context")
	}
	if err := ctx.Err(); err != nil {
		return zero, err
	}
	if slug != "" && !claims.ValidSlug(slug) {
		return zero, fmt.Errorf("C05: invalid exact selector %q; use a queued change slug", slug)
	}
	repo, err := gitstore.Open(start)
	if err != nil {
		return zero, fmt.Errorf("P02: discover the selected checkout before verification: %w", err)
	}
	if slug == "" {
		held, err := claims.List(repo.Root)
		if err != nil {
			return zero, fmt.Errorf("P02: inspect held claims before verification: %w", err)
		}
		count := 0
		for _, claim := range held {
			if claim.Worktree == repo.Root {
				slug = claim.Change
				count++
			}
		}
		if count != 1 {
			return zero, fmt.Errorf("P02: expected one held claim owned by this checkout, found %d; inspect claims and use fireweed start or release for the intended change", count)
		}
	}
	source, claim, err := approval.LoadHeld(repo.Root, slug)
	if err != nil {
		return zero, err
	}
	var change contracts.Change
	for _, candidate := range source.Changes {
		if candidate.Slug == slug {
			change = candidate
			break
		}
	}
	if err := ctx.Err(); err != nil {
		return zero, err
	}
	v := model.Verification{Change: slug, Phase: change.Phase, Diagnostics: []checks.Diagnostic{}, Coverage: []checks.Coverage{}, Edits: []gitstore.Edit{}, Findings: []model.Finding{}}
	add := func(code, message string) {
		v.Diagnostics = append(v.Diagnostics, checks.Diagnostic{Code: code, Severity: "error", Path: change.Path, Message: message})
	}
	finish := func(err error) (model.Verification, error) { normalize(&v); return v, err }
	effective, err := project.Project(source, slug)
	if err != nil {
		add("C06", "Restore a valid approved queue projection: "+err.Error())
		return finish(nil)
	}
	var target contracts.Module
	for i, module := range effective.Modules {
		if module.Name == change.Module {
			target, err = project.ApplyChange(module, change)
			if err != nil {
				add("C06", "Restore a valid approved change projection: "+err.Error())
				return finish(nil)
			}
			effective.Modules[i] = target
			break
		}
	}
	for _, d := range project.Check(effective) {
		v.Diagnostics = append(v.Diagnostics, checks.Diagnostic{Code: d.Code, Severity: d.Severity, Message: d.Message, Path: d.Path})
	}
	guard := func() {
		result := worktree.Check(effective, target, change.Path, claim)
		v.Edits = result.Edits
		v.Diagnostics = append(v.Diagnostics, result.Diagnostics...)
	}
	guard()
	var current contracts.Change
	if path, safeErr := gitstore.SafePath(repo.Root, change.Path); safeErr == nil {
		if info, statErr := os.Stat(path); statErr == nil && info.Mode().IsRegular() {
			if raw, readErr := os.ReadFile(path); readErr == nil {
				current, err = contracts.ParseChange(string(raw), change.Path, target.IDPrefix)
				if err != nil {
					add("P2-01", "Restore a valid readable current change before verification: "+err.Error())
				}
				if current.Blocked != "" {
					add("P05", "Resolve the blocked change and its Questions before verification: "+current.Blocked)
				}
			}
		}
	}
	if hasErrors(v.Diagnostics) {
		return finish(nil)
	}
	if err := ctx.Err(); err != nil {
		return finish(err)
	}
	if change.Phase == "audit" {
		audit(&v, current, target, change)
		return finish(nil)
	}
	evaluate := func(module contracts.Module, failFirst, already []string) {
		coverage := checks.Evaluate(ctx, effective, module, failFirst, already)
		v.Coverage = append(v.Coverage, coverage)
		v.Diagnostics = append(v.Diagnostics, coverage.Diagnostics...)
	}
	switch change.Phase {
	case "evals":
		evaluate(target, targets(change), change.AlreadySatisfied)
	case "impl":
		evaluate(target, nil, nil)
		for _, module := range dependents(effective, target.Name) {
			if ctx.Err() != nil {
				break
			}
			evaluate(module, nil, nil)
		}
		if ctx.Err() == nil {
			v.Diagnostics = append(v.Diagnostics, checks.Boundaries(ctx, effective)...)
		}
	default:
		add("P02", "Restore a valid approved worker phase before verification")
	}
	guard()
	return finish(ctx.Err())
}

func targets(change contracts.Change) []string {
	ids := append([]string{}, change.Fixes...)
	for _, group := range [][]contracts.Spec{change.Added, change.Modified} {
		for _, spec := range group {
			ids = append(ids, spec.ID)
		}
	}
	return ids
}
func hasErrors(ds []checks.Diagnostic) bool {
	for _, d := range ds {
		if d.Severity == "error" {
			return true
		}
	}
	return false
}
func normalize(v *model.Verification) {
	sort.Slice(v.Diagnostics, func(i, j int) bool {
		a, b := v.Diagnostics[i], v.Diagnostics[j]
		if a.Path != b.Path {
			return a.Path < b.Path
		}
		if a.Code != b.Code {
			return a.Code < b.Code
		}
		if a.Severity != b.Severity {
			return a.Severity < b.Severity
		}
		return a.Message < b.Message
	})
	ds := v.Diagnostics[:0]
	for _, d := range v.Diagnostics {
		if len(ds) == 0 || ds[len(ds)-1] != d {
			ds = append(ds, d)
		}
	}
	v.Diagnostics = ds
}
func dependents(snapshot project.Snapshot, own string) []contracts.Module {
	policy := snapshot.Config.VerifyDependents
	if policy == "none" {
		return nil
	}
	selected := map[string]bool{own: true}
	for {
		changed := false
		for _, module := range snapshot.Modules {
			if selected[module.Name] {
				continue
			}
			for _, dependency := range module.DependsOn {
				if dependency == own || policy != "direct" && selected[dependency] {
					selected[module.Name] = true
					changed = true
					break
				}
			}
		}
		if !changed || policy == "direct" {
			break
		}
	}
	result := []contracts.Module{}
	for _, module := range snapshot.Modules {
		if module.Name != own && selected[module.Name] {
			result = append(result, module)
		}
	}
	sort.Slice(result, func(i, j int) bool { return result[i].Name < result[j].Name })
	return result
}

var findingLine = regexp.MustCompile(`^- (\[untestable\] )?([A-Z][A-Z0-9]*-[0-9]+): "(.+)" — (.+)$`)

func audit(v *model.Verification, current contracts.Change, module contracts.Module, approved contracts.Change) {
	body := ""
	for _, section := range current.Document.Sections {
		if section.Name == "Audit" {
			body = section.Body
			break
		}
	}
	specs := map[string]string{}
	for _, spec := range module.Specs {
		if !strings.Contains(spec.Tag, "[review]") && !strings.Contains(spec.Tag, "[quarantine:") {
			specs[spec.ID] = spec.Body
		}
	}
	for _, line := range strings.Split(body, "\n") {
		match := findingLine.FindStringSubmatch(line)
		if match == nil || strings.TrimSpace(match[3]) == "" || strings.TrimSpace(match[4]) == "" || !strings.Contains(specs[match[2]], match[3]) {
			continue
		}
		v.Findings = append(v.Findings, model.Finding{ID: match[2], Clause: match[3], Text: line, Untestable: match[1] != ""})
	}
	if len(v.Findings) > 0 {
		return
	}
	valid := strings.Contains(strings.ToLower(body), "no findings")
	for _, id := range targets(approved) {
		valid = valid && contracts.MatchesID(body, "", id)
	}
	if !valid {
		v.Diagnostics = append(v.Diagnostics, checks.Diagnostic{Code: "P1-03", Severity: "error", Path: approved.Path, Message: "Write valid clause findings, or No findings and every Added, Modified and fixes ID in the Audit section."})
	}
}
