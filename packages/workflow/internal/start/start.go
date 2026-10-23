package start

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"github.com/lurkshark/fireweed/packages/checks"
	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/gitstore"
	"github.com/lurkshark/fireweed/packages/project"
	"github.com/lurkshark/fireweed/packages/workflow/internal/claims"
	workercontext "github.com/lurkshark/fireweed/packages/workflow/internal/context"
	"github.com/lurkshark/fireweed/packages/workflow/internal/eligibility"
	"github.com/lurkshark/fireweed/packages/workflow/model"
)

func problem(format string, args ...any) error {
	return fmt.Errorf("%s. Check the queued source and local ownership, restore valid inputs and retry.", fmt.Sprintf(format, args...))
}
func identity() (string, error) {
	var b [16]byte
	if _, e := rand.Read(b[:]); e != nil {
		return "", problem("cannot create fresh run: %v", e)
	}
	return hex.EncodeToString(b[:]), nil
}

func Start(start, slug string) (out model.StartResult, err error) {
	if !claims.ValidSlug(slug) {
		return out, problem("C05: invalid exact queued slug %q", slug)
	}
	repo, err := gitstore.Open(start)
	if err != nil {
		return out, problem("cannot discover repository: %v", err)
	}
	local := localGit{root: repo.Root, common: repo.CommonDir}
	if err = local.prepare(); err != nil {
		return out, err
	}
	head, err := local.resolveCommit("HEAD")
	if err != nil {
		return out, err
	}
	status, err := local.raw("status", "--porcelain=v1", "-z", "--untracked-files=all")
	if err != nil {
		return out, err
	}
	if status != "" {
		return out, problem("P01: working tree or index is dirty; commit or revert unrelated work before starting")
	}
	headCommit, err := local.commit(head)
	if err != nil {
		return out, err
	}
	if err = local.completeTree(headCommit.tree, map[string]bool{}); err != nil {
		return out, err
	}
	snapshot, err := project.LoadAt(repo.Root, head)
	if err != nil {
		return out, problem("cannot load current committed project: %v", err)
	}
	diagnostics := make([]checks.Diagnostic, 0)
	for _, d := range append(project.Check(snapshot), project.CheckQueue(snapshot)...) {
		if d.Severity == "error" {
			return out, problem("%s: %s (%s)", d.Code, d.Message, d.Path)
		}
		diagnostics = append(diagnostics, checks.Diagnostic{Code: d.Code, Severity: d.Severity, Message: d.Message, Path: d.Path})
	}
	index := -1
	var module contracts.Module
	for i, ch := range snapshot.Changes {
		if ch.Slug == slug {
			if index >= 0 {
				return out, problem("C05: queued slug %s is ambiguous", slug)
			}
			index = i
		}
	}
	if index < 0 {
		return out, problem("C05: no queued change matches %s", slug)
	}
	selected := snapshot.Changes[index]
	for _, m := range snapshot.Modules {
		if m.Name == selected.Module {
			module = m
		}
	}
	eligible, err := eligibility.Check(snapshot, slug)
	if err != nil {
		return out, err
	}
	if !eligible {
		return out, problem("P03: %s is blocked or is not an eligible queue head; resolve the earlier queued work", slug)
	}
	inventory, err := claims.List(repo.Root)
	if err != nil {
		return out, err
	}
	for _, cl := range inventory {
		if cl.Worktree == repo.Root || cl.Change == slug {
			return out, problem("P02: existing claim %s belongs to %s; use release after inspecting that work", cl.Change, cl.Worktree)
		}
	}
	original, err := readNode(repo.Root, selected.Path, false)
	if err != nil {
		return out, problem("cannot read selected change %s: %v", selected.Path, err)
	}
	committed, err := local.file(head, selected.Path)
	if err != nil {
		return out, err
	}
	if string(original.bytes) != committed {
		return out, problem("P01: selected change differs from current HEAD")
	}
	allocated, _, err := project.Allocate(snapshot, slug)
	if err != nil {
		return out, problem("cannot allocate or resume %s: %v", selected.Path, err)
	}
	if allocated.Run == "" {
		allocated.Run, err = identity()
		if err != nil {
			return out, err
		}
		allocated.Phase = "evals"
		if allocated.Kind == "maintenance" || allocated.Regenerate {
			allocated.Phase = "impl"
		}
	}
	allocated.Document.Metadata["run"] = allocated.Run
	allocated.Document.Metadata["phase"] = allocated.Phase
	startedText, err := contracts.RenderMarkdown(allocated.Document)
	if err != nil {
		return out, problem("cannot render started change: %v", err)
	}
	started, err := contracts.ParseChange(startedText, selected.Path, module.IDPrefix)
	if err != nil {
		return out, problem("cannot parse started change: %v", err)
	}
	diagnostics = append(diagnostics, local.approvalWarnings(snapshot.BaseRef, head, module, startedText, selected.Path)...)
	if started.Regenerate && !local.currentDrill(module, head) {
		diagnostics = append(diagnostics, checks.Diagnostic{Code: "D01", Severity: "warning", Path: selected.Path, Message: "No current passing drill report covers this module. Check the local report, run a real drill and review its report before relying on regeneration."})
	}
	nonce, err := identity()
	if err != nil {
		return out, err
	}
	tx := transaction{root: repo.Root, path: selected.Path, original: original, written: startedText}
	defer func() {
		if err != nil {
			out = model.StartResult{}
			err = tx.rollback(err)
		}
	}()
	reserved, reserveErr := local.reserve(model.Claim{Worktree: repo.Root, Change: slug, Run: started.Run, Base: head, State: "reservation:" + nonce})
	if reserveErr != nil {
		if actual, e := claims.Load(repo.Root, slug); e == nil && actual == reserved {
			tx.claim = actual
		}
		return out, reserveErr
	}
	tx.claim = reserved
	if err = tx.owned(); err != nil {
		return out, err
	}
	if !sameNode(repo.Root, selected.Path, original) {
		return out, problem("selected source changed concurrently; recovery preserved its replacement")
	}
	if err = writeAtomic(repo.Root, selected.Path, []byte(startedText), original.mode); err != nil {
		return out, problem("cannot write selected change %s: %v", selected.Path, err)
	}
	tx.didWrite = true
	next, updateErr := local.refresh(tx.claim, startedText)
	if updateErr != nil {
		if actual, e := claims.Load(repo.Root, slug); e == nil && actual == next {
			tx.claim = actual
		}
		return out, updateErr
	}
	tx.claim = next
	if err = tx.owned(); err != nil {
		return out, err
	}
	if !tx.ownsSource() {
		return out, problem("selected change changed concurrently; recovery preserved its replacement")
	}
	snapshot.Changes[index] = started
	if started.Regenerate {
		if err = tx.regenerate(local, snapshot, started); err != nil {
			return out, err
		}
	}
	context, err := workercontext.Build(snapshot, slug)
	if err != nil {
		return out, err
	}
	if err = tx.owned(); err != nil {
		return out, err
	}
	if !tx.ownsSource() {
		return out, problem("selected change changed concurrently; recovery preserved its replacement")
	}
	now, err := local.text("rev-parse", "--verify", "HEAD")
	if err != nil {
		return out, err
	}
	if now != head {
		return out, problem("HEAD changed during start; recover the original baseline before retry")
	}
	return model.StartResult{Context: context, Claim: tx.claim, Diagnostics: diagnostics}, nil
}
