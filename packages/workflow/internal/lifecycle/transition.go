package lifecycle

import (
	"context"
	"crypto/sha256"
	"errors"
	"fmt"
	"maps"
	"os"
	"path"
	"slices"
	"strings"
	"time"

	"github.com/lurkshark/fireweed/packages/assets"
	"github.com/lurkshark/fireweed/packages/checks"
	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/gitstore"
	"github.com/lurkshark/fireweed/packages/project"
	"github.com/lurkshark/fireweed/packages/workflow/model"
)

func Done(ctx context.Context, request model.Request) (model.Result, error) {
	state, held, err := loadHeld(request.Root, request.Change)
	result := model.Result{State: state}
	if err != nil {
		return result, err
	}
	if err := supported(state); err != nil {
		return result, err
	}
	reason := strings.TrimSpace(str(state.Change.Document.Frontmatter["blocked"]))
	questions := section(state.Change.Document, "Questions")
	if state.Phase != "audit" && reason != "" && strings.TrimSpace(questions) != "" {
		blocked, err := block(state, held, reason, questions)
		result.State = blocked.State
		if err != nil {
			return result, err
		}
		return result, blockedError(state, reason)
	}

	request.Quick = false
	result, err = verify(ctx, request, state, held)
	if err != nil {
		return result, verificationFailure{err}
	}

	// Hooks run in the checkout. Recheck ownership, approved state and zones
	// before staging any phase changes.
	current, currentClaim, err := loadHeld(request.Root, state.Change.Path)
	if err != nil {
		return result, err
	}
	if current.ClaimOID != state.ClaimOID {
		return result, failure("P02", state.Change.Path, "claim changed during verification")
	}
	if err := verifyApproval(current, currentClaim, false); err != nil {
		return result, err
	}
	result.State = current
	if _, err := checkZones(&result); err != nil {
		return result, err
	}

	doc := current.Change.Document
	doc.Frontmatter = maps.Clone(doc.Frontmatter)
	next, blockReason := "applied", ""
	switch current.Phase {
	case "evals":
		next = "audit"
		if !auditEnabled(current.Snapshot.Config) {
			next = "impl"
		}
	case "audit":
		audit, err := auditSubmission(current, held)
		if err != nil {
			return result, err
		}
		next, blockReason = audit.next, audit.blocked
		doc.Body = contracts.ReplaceSection(doc.Body, "Audit", audit.text)
		if next == "evals" {
			doc.Frontmatter["audit_rounds"] = audit.rounds
		}
		if blockReason != "" {
			doc.Frontmatter["blocked"] = blockReason
			doc.Body = appendSection(doc, "Questions", audit.text)
		}
	}
	doc.Frontmatter["phase"] = next
	line, err := provenance(current, request.Agent, result.Reports)
	if err != nil {
		return result, err
	}
	doc.Body = appendSection(doc, "Provenance", line)
	result, err = finishPhase(result, held, doc, next == "applied" || blockReason != "")
	if err == nil && blockReason != "" {
		err = blockedError(current, blockReason)
	}
	return result, err
}

func blockedError(state model.State, reason string) error {
	return failure("P06", state.Change.Path, "change blocked: %s", reason)
}

func playbook(root, name string) (string, error) {
	data, err := os.ReadFile(disk(root, path.Join(".fireweed/playbooks", name+".md")))
	if err == nil {
		return string(data), nil
	}
	if !os.IsNotExist(err) {
		return "", err
	}
	return assets.Playbook(name)
}

func playbookName(state model.State) string {
	if regenerate, _ := state.Change.Document.Frontmatter["regenerate"].(bool); regenerate && state.Phase == "impl" {
		return "regenerate"
	}
	return state.Phase
}

func provenance(state model.State, agent string, reports map[string]checks.Report) (string, error) {
	if agent == "" {
		agent = "manual"
	}
	book, err := playbook(state.Snapshot.Root, playbookName(state))
	if err != nil {
		return "", err
	}
	total, failing, passing := 0, 0, 0
	for _, report := range reports {
		for _, c := range report.Cases {
			total++
			switch c.Result {
			case "pass":
				passing++
			case "fail", "error":
				failing++
			}
		}
	}
	return fmt.Sprintf("- %s %s agent=%s fireweed=%s playbook=%x tests=%d failing=%d passing=%d",
		state.Phase, time.Now().UTC().Format(time.RFC3339Nano), agent, assets.Version(),
		sha256.Sum256([]byte(book)), total, failing, passing), nil
}

func finishPhase(result model.Result, held claim, doc contracts.Document, release bool) (model.Result, error) {
	state := result.State
	root, pending := state.Snapshot.Root, state.Change.Path
	phase := str(doc.Frontmatter["phase"])
	identity := state.Change
	var files []fileEdit
	if phase == "applied" {
		for _, d := range project.Check(root) {
			if !d.Warning {
				result.Diagnostics = append(result.Diagnostics, d)
				return result, staticFailure([]project.Diagnostic{d})
			}
		}
		spec, err := appliedSpec(state)
		if err != nil {
			return result, err
		}
		if spec.name != "" {
			files = append(files, spec)
		}
		identity.Path = archivePath(state.Change)
		if _, err := os.Lstat(disk(root, identity.Path)); err == nil {
			return result, failure("P01", identity.Path, "archive already exists: %s", identity.Path)
		} else if !os.IsNotExist(err) {
			return result, err
		}
		files = append(files, fileEdit{name: pending, remove: true})
	}
	text, err := contracts.RenderDocument(doc)
	if err != nil {
		return result, err
	}
	files = append(files, fileEdit{name: identity.Path, contents: []byte(text)})
	change, err := parseChange(identity, doc)
	if err != nil {
		return result, err
	}
	rollback, err := editFiles(root, files)
	if err != nil {
		return result, err
	}
	if phase == "applied" {
		diagnostics := project.Check(root)
		if err := staticFailure(diagnostics); err != nil {
			result.Diagnostics = append(result.Diagnostics, diagnostics...)
			return result, errors.Join(err, rollback())
		}
	}

	commit, err := snapshotWorktree(root, state.PhaseBase, fmt.Sprintf("%s: %s", state.Phase, numberedName(change)))
	if err != nil {
		return result, errors.Join(err, rollback())
	}
	var claimOID string
	if !release {
		held.Base, held.Approved = commit, text
		claimOID, err = encodeClaim(root, held)
		if err != nil {
			return result, errors.Join(err, rollback())
		}
	}
	if err := publish(root, state.PhaseBase, commit, claimsPrefix+held.Slug, state.ClaimOID, claimOID); err != nil {
		return result, errors.Join(err, rollback())
	}

	result.State.Change, result.State.Phase = change, phase
	result.State.Head, result.State.PhaseBase, result.State.ClaimOID = commit, commit, claimOID
	_, err = git(root, "reset", "-q", commit, "--", ".")
	return result, err
}

func appliedSpec(state model.State) (fileEdit, error) {
	change := state.Change
	addedDecisions := strings.TrimSpace(section(change.Document, "Decisions"))
	removedDecisions := strings.TrimSpace(section(change.Document, "Removed decisions"))
	_, hasSet := change.Document.Frontmatter["set"]
	if len(change.Added)+len(change.Modified)+len(change.Removed) == 0 && !hasSet && addedDecisions == "" && removedDecisions == "" {
		return fileEdit{}, nil
	}
	// Project uses the shared application rules for set values and spec edits.
	snapshot, err := project.Project(state.Snapshot.Root, change.Path)
	if err != nil {
		return fileEdit{}, err
	}
	module := moduleNamed(snapshot, change.Module)
	doc := module.Document
	if addedDecisions != "" || removedDecisions != "" {
		doc.Body = applyDecisions(doc, change.Slug, addedDecisions, removedDecisions)
	}
	text, err := contracts.RenderDocument(doc)
	return fileEdit{name: path.Join(module.Dir, "SPEC.md"), contents: []byte(text)}, err
}

func decisionEntries(body string) []string {
	var entries []string
	for _, paragraph := range strings.Split(strings.TrimSpace(body), "\n\n") {
		var entry string
		for _, line := range strings.Split(paragraph, "\n") {
			line = strings.TrimSpace(line)
			if strings.HasPrefix(line, "- ") && entry != "" {
				entries = append(entries, entry)
				entry = ""
			}
			if entry == "" {
				entry = line
			} else if line != "" {
				entry += "\n" + line
			}
		}
		if entry != "" {
			entries = append(entries, entry)
		}
	}
	return entries
}

func applyDecisions(doc contracts.Document, slug, added, removed string) string {
	entries := decisionEntries(section(doc, "Decisions"))
	remove := decisionEntries(removed)
	entries = slices.DeleteFunc(entries, func(entry string) bool {
		for _, quoted := range remove {
			quote := strings.Trim(strings.TrimPrefix(strings.TrimPrefix(quoted, "- "), "> "), "\"")
			_, body, attributed := strings.Cut(strings.TrimPrefix(entry, "- "), ": ")
			if quote == entry || quote == strings.TrimPrefix(entry, "- ") || attributed && quote == body {
				return true
			}
		}
		return false
	})
	for _, decision := range decisionEntries(added) {
		entries = append(entries, "- "+slug+": "+strings.TrimPrefix(decision, "- "))
	}
	return contracts.ReplaceSection(doc.Body, "Decisions", "\n"+strings.Join(entries, "\n\n")+"\n")
}

type fileEdit struct {
	name     string
	contents []byte
	remove   bool
}

func (edit fileEdit) apply(root string) error {
	if edit.remove {
		err := os.Remove(disk(root, edit.name))
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	return writeFile(disk(root, edit.name), edit.contents)
}

func editFiles(root string, edits []fileEdit) (func() error, error) {
	var previous []fileEdit
	rollback := func() error {
		var err error
		for i := len(previous) - 1; i >= 0; i-- {
			if restoreErr := previous[i].apply(root); restoreErr != nil {
				err = errors.Join(err, restoreErr)
			}
		}
		return err
	}
	for _, edit := range edits {
		data, err := os.ReadFile(disk(root, edit.name))
		if err != nil && !os.IsNotExist(err) {
			return rollback, errors.Join(err, rollback())
		}
		previous = append(previous, fileEdit{name: edit.name, contents: data, remove: os.IsNotExist(err)})
		if err := edit.apply(root); err != nil {
			return rollback, errors.Join(err, rollback())
		}
	}
	return rollback, nil
}

func snapshotWorktree(root, base, message string) (string, error) {
	var commit string
	err := tempIndex(root, base, func(env map[string]string) error {
		if err := stageWorktree(root, env); err != nil {
			return err
		}
		var err error
		commit, err = commitIndex(root, base, message, env)
		return err
	})
	return commit, err
}

func stageWorktree(root string, index map[string]string) error {
	tracked, err := gitInput(root, nil, index, "diff", "--name-only", "--no-renames", "-z", "--")
	if err != nil {
		return err
	}
	untracked, err := git(root, "ls-files", "--others", "--exclude-standard", "-z")
	if err != nil {
		return err
	}
	staged, err := git(root, "diff", "--cached", "--diff-filter=A", "--no-renames", "--name-only", "-z", "--")
	if err != nil {
		return err
	}
	// A staged addition is tracked even if it matches an ignore rule.
	for _, name := range gitFields(staged) {
		if name != "" && !isRunArtifact(name) {
			if _, err := os.Lstat(disk(root, name)); err == nil {
				untracked += name + "\x00"
			} else if !os.IsNotExist(err) {
				return err
			}
		}
	}
	var paths strings.Builder
	for _, name := range gitFields(tracked + untracked) {
		if name != "" && !isRunArtifact(name) {
			paths.WriteString(name)
			paths.WriteByte(0)
		}
	}
	if paths.Len() == 0 {
		return nil
	}
	_, err = gitInput(root, []byte(paths.String()), index, "--literal-pathspecs", "add", "-A", "-f",
		"--pathspec-from-file=-", "--pathspec-file-nul")
	return err
}

func commitIndex(root, parent, message string, env map[string]string) (string, error) {
	tree, err := gitInput(root, nil, env, "write-tree")
	if err != nil {
		return "", err
	}
	commit, err := git(root, "commit-tree", strings.TrimSpace(tree), "-p", parent, "-m", message)
	return strings.TrimSpace(commit), err
}

func publish(root, oldHead, newHead, name, oldClaim, newClaim string) error {
	if name == "" {
		return gitstore.CASRef(root, "HEAD", oldHead, newHead)
	}
	commands := fmt.Sprintf("update HEAD %s %s\n", newHead, oldHead)
	switch {
	case newClaim != "":
		commands += fmt.Sprintf("update %s %s %s\n", name, newClaim, oldClaim)
	case oldClaim != "":
		commands += fmt.Sprintf("delete %s %s\n", name, oldClaim)
	default:
		commands += fmt.Sprintf("verify %s %040d\n", name, 0)
	}
	_, err := gitInput(root, []byte(commands), nil, "update-ref", "--stdin")
	return err
}
