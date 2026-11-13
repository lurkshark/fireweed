package lifecycle

import (
	"errors"
	"fmt"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/gitstore"
	"github.com/lurkshark/fireweed/packages/project"
	"github.com/lurkshark/fireweed/packages/workflow/model"
)

func Block(request model.BlockRequest) (model.BlockResult, error) {
	state, held, err := loadHeld(request.Root, request.Change)
	if err != nil {
		return model.BlockResult{State: state}, err
	}
	return block(state, held, request.Reason, "")
}

func block(state model.State, held claim, reason, questions string) (model.BlockResult, error) {
	result := model.BlockResult{State: state}
	if strings.TrimSpace(reason) == "" {
		return result, failure("P01", held.Path, "blocked reason is required")
	}
	if err := supported(state); err != nil {
		return result, err
	}
	if err := verifyApproval(state, held, questions != ""); err != nil {
		return result, err
	}
	root := state.Snapshot.Root
	doc, err := held.approvedDocument()
	if err != nil {
		return result, err
	}
	doc.Frontmatter["blocked"] = reason
	if questions != "" {
		doc.Body = contracts.ReplaceSection(doc.Body, "Questions", questions)
	}
	change, err := parseChange(state.Change, doc)
	if err != nil {
		return result, err
	}
	rendered, err := contracts.RenderDocument(doc)
	if err != nil {
		return result, err
	}
	attempt, err := snapshotWorktree(root, state.PhaseBase, "save attempt: "+numberedName(change))
	if err != nil {
		return result, err
	}
	blockCommit, err := blockCommit(root, state.PhaseBase, change, rendered)
	if err != nil {
		return result, err
	}
	attemptRef, err := saveAttemptRef(root, change.Slug, attempt)
	if err != nil {
		return result, err
	}
	result.AttemptRef, result.AttemptOID = attemptRef, attempt

	// Restore tracked and staged paths and remove new step files. The run
	// directory remains outside both the attempt and the restoration.
	if _, err := git(root, "restore", "--staged", "--worktree", "--source", state.PhaseBase, "--", ".", excludeRuns); err != nil {
		return result, err
	}
	if _, err := git(root, "clean", "-fd", "--", ".", excludeRuns); err != nil {
		return result, err
	}
	if err := writeFile(disk(root, change.Path), []byte(rendered)); err != nil {
		return result, err
	}
	if err := publish(root, state.PhaseBase, blockCommit, claimsPrefix+held.Slug, state.ClaimOID, ""); err != nil {
		return result, err
	}
	result.State.Change = change
	result.State.Head, result.State.PhaseBase, result.State.ClaimOID = blockCommit, blockCommit, ""
	_, err = git(root, "reset", "-q", blockCommit, "--", ".")
	return result, err
}

func blockCommit(root, base string, change project.Change, rendered string) (string, error) {
	blob, err := gitstore.WriteBlob(root, []byte(rendered))
	if err != nil {
		return "", err
	}
	var commit string
	err = tempIndex(root, base, func(env map[string]string) error {
		if _, err := gitInput(root, nil, env, "update-index", "--add", "--cacheinfo", "100644", blob, change.Path); err != nil {
			return err
		}
		var err error
		commit, err = commitIndex(root, base, "block: "+numberedName(change), env)
		return err
	})
	return commit, err
}

func saveAttemptRef(root, slug, oid string) (string, error) {
	prefix := attemptsPrefix + slug + "/"
	all, err := refs(root, prefix)
	if err != nil {
		return "", err
	}
	used := make(map[string]bool)
	for _, r := range all {
		used[r.name] = true
	}
	for number := 1; ; number++ {
		name := prefix + strconv.Itoa(number)
		if used[name] {
			continue
		}
		err := gitstore.CASRef(root, name, "", oid)
		if err == nil {
			return name, nil
		}
		if _, exists, lookupErr := gitstore.Ref(root, name); lookupErr != nil || !exists {
			return "", errors.Join(err, lookupErr)
		}
	}
}

type savedAttempt struct {
	doc   contracts.Document
	patch string
}

func loadAttempt(root string, change project.Change, number int) (savedAttempt, error) {
	var saved savedAttempt
	name := attemptsPrefix + change.Slug + "/" + strconv.Itoa(number)
	oid, exists, err := gitstore.Ref(root, name)
	if err != nil {
		return saved, coded("P01", name, err)
	}
	if !exists {
		return saved, failure("P01", name, "saved attempt not found: %s", name)
	}
	data, err := gitstore.ReadAt(root, oid, change.Path)
	if err != nil {
		return saved, coded("P01", name, err)
	}
	saved.doc, err = contracts.ParseDocument(string(data))
	if err != nil {
		return saved, coded("P01", name, err)
	}
	for _, key := range []string{"phase", "run"} {
		if str(saved.doc.Frontmatter[key]) != str(change.Document.Frontmatter[key]) {
			return saved, failure("P01", name, "saved attempt has a different %s: %s", key, name)
		}
	}
	saved.patch, err = git(root, "diff", "--binary", "--no-renames", "--no-ext-diff", oid+"^", oid,
		"--", ".", ":(exclude,top)"+change.Path, excludeRuns)
	if err != nil {
		return saved, coded("P01", name, err)
	}
	if saved.patch != "" {
		_, err = gitInput(root, []byte(saved.patch), nil, "apply", "--check", "--")
		if err != nil {
			return saved, coded("P01", name, fmt.Errorf("apply saved attempt: %w", err))
		}
	}
	return saved, nil
}

func restoreAttempt(root string, change project.Change, saved savedAttempt) error {
	doc := change.Document
	editable := []string{"Decisions", "Questions", "Notes"}
	if str(doc.Frontmatter["phase"]) == "audit" {
		editable = []string{"Audit"}
	}
	for _, name := range editable {
		if content, found := contracts.Section(saved.doc.Body, name); found {
			doc.Body = contracts.ReplaceSection(doc.Body, name, content)
		}
	}
	if err := writeDocument(root, change.Path, doc); err != nil {
		return err
	}
	if saved.patch != "" {
		_, err := gitInput(root, []byte(saved.patch), nil, "apply", "--")
		return err
	}
	return nil
}

func pathAt(root, revision string, identity project.Change) (project.Change, bool, error) {
	if revision == "" {
		return identity, false, nil
	}
	for _, file := range []string{archivePath(identity), pendingPath(identity)} {
		data, err := gitstore.ReadAt(root, revision, file)
		if err != nil {
			continue
		}
		identity.Path = file
		doc, err := contracts.ParseDocument(string(data))
		if err != nil {
			return identity, false, err
		}
		change, err := parseChange(identity, doc)
		return change, true, err
	}
	return identity, false, nil
}

type changeCommit struct {
	oid    string
	parent string
	before project.Change
	after  project.Change
}

func phaseCommits(root, revision string, identity project.Change) ([]changeCommit, error) {
	out, err := git(root, "log", "--full-history", "--format=%H %P", revision,
		"--", pendingPath(identity), archivePath(identity))
	if err != nil {
		return nil, err
	}
	var result []changeCommit
	for _, line := range strings.Split(out, "\n") {
		fields := strings.Fields(line)
		if len(fields) == 0 {
			continue
		}
		oid, parent := fields[0], ""
		if len(fields) > 1 {
			parent = fields[1]
		}
		after, present, err := pathAt(root, oid, identity)
		if err != nil {
			return nil, err
		}
		phase := str(after.Document.Frontmatter["phase"])
		if !present || phase == "" {
			continue
		}
		before, _, err := pathAt(root, parent, identity)
		if err != nil {
			return nil, err
		}
		if phase != str(before.Document.Frontmatter["phase"]) {
			result = append(result, changeCommit{oid: oid, parent: parent, before: before, after: after})
		}
	}
	return result, nil
}

func selectPendingOrArchived(root, selector string) (project.Change, error) {
	pending, err := changePaths(root, "", pendingDir)
	if err != nil {
		return project.Change{}, err
	}
	archive, err := changePaths(root, "", archiveDir)
	if err != nil {
		return project.Change{}, err
	}
	return selectChange(append(pending, archive...), selector)
}

func Commits(root, selector string) ([]string, error) {
	repo, err := gitstore.Inspect(root)
	if err != nil {
		return nil, err
	}
	identity, err := selectPendingOrArchived(repo.Root, selector)
	if err != nil {
		return nil, err
	}
	commits, err := phaseCommits(repo.Root, repo.Head, identity)
	if err != nil {
		return nil, err
	}
	result := make([]string, 0, len(commits))
	for _, commit := range commits {
		result = append(result, commit.oid)
	}
	return result, nil
}

func Abandon(root, selector string) (string, error) {
	repo, err := gitstore.Inspect(root)
	if err != nil {
		return "", err
	}
	root = repo.Root
	if err := clean(root, "abandon"); err != nil {
		return "", err
	}
	change, err := selectPendingOrArchived(root, selector)
	if err != nil {
		return "", err
	}
	if path.Dir(change.Path) == archiveDir {
		return "", failure("P01", change.Path, "cannot abandon archived change %s", change.Path)
	}
	name := claimsPrefix + change.Slug
	oid, exists, err := gitstore.Ref(root, name)
	if err != nil {
		return "", coded("P02", change.Slug, err)
	}
	if exists {
		c, err := readClaim(root, ref{name, oid})
		if err != nil {
			return "", err
		}
		if c.Worktree != root || c.Path != change.Path {
			return "", failure("P02", change.Slug, "claim belongs to another worktree or change")
		}
	}
	commits, err := phaseCommits(root, repo.Head, change)
	if err != nil {
		return "", err
	}
	var result string
	err = tempIndex(root, repo.Head, func(env map[string]string) error {
		for _, commit := range commits {
			var patch string
			var err error
			paths := []string{"--", ".", ":(exclude,top)" + pendingPath(change), ":(exclude,top)" + archivePath(change)}
			if commit.parent == "" {
				patch, err = git(root, append([]string{"show", "--format=", "--binary", "--no-renames", commit.oid}, paths...)...)
			} else {
				patch, err = git(root, append([]string{"diff", "--binary", "--no-renames", commit.parent, commit.oid}, paths...)...)
			}
			if err != nil {
				return err
			}
			if patch != "" {
				if _, err := gitInput(root, []byte(patch), env, "apply", "--cached", "--3way", "--reverse", "--"); err != nil {
					return fmt.Errorf("reverse phase commit %s: %w", commit.oid, err)
				}
			}
		}
		if _, err := gitInput(root, nil, env, "update-index", "--force-remove", "--", change.Path); err != nil {
			return err
		}
		result, err = commitIndex(root, repo.Head, "abandon: "+numberedName(change), env)
		return err
	})
	if err != nil {
		return "", err
	}
	if _, err := git(root, "read-tree", "-u", "-m", repo.Head, result); err != nil {
		return "", err
	}
	if err := publish(root, repo.Head, result, name, oid, ""); err != nil {
		_, rollback := git(root, "read-tree", "-u", "-m", result, repo.Head)
		return "", errors.Join(err, rollback)
	}
	return result, nil
}

func drillDate(root, module string) (time.Time, string) {
	doc, err := readDocument(root, path.Join(".fireweed/drills", module+".md"))
	if err != nil {
		return time.Time{}, ""
	}
	date, _ := doc.Frontmatter["date"].(time.Time)
	if date.IsZero() {
		date, _ = time.Parse(time.RFC3339Nano, str(doc.Frontmatter["date"]))
	}
	return date, str(doc.Frontmatter["result"])
}
