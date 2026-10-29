package lifecycle

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/gitstore"
	"github.com/lurkshark/fireweed/packages/workflow/model"
)

type savedFile struct {
	path   string
	status byte
	data   []byte
}

// savedStep reads the difference made by Block's attempt commit, whose parent
// is the held phase base. Later commits on HEAD are not part of that difference.
func savedStep(root, oid string) ([]savedFile, error) {
	result, err := gitstore.Run(context.Background(), gitstore.Request{
		Args:    []string{"git", "-C", root, "-c", "core.hooksPath=/dev/null", "diff", "--no-ext-diff", "--no-textconv", "--name-status", "-z", "--no-renames", oid + "^", oid, "--"},
		Timeout: 30 * time.Second,
		Env:     map[string]string{"GIT_NO_LAZY_FETCH": "1", "GIT_TERMINAL_PROMPT": "0"},
	})
	if err != nil {
		return nil, fmt.Errorf("read saved attempt delta: %w: %s", err, result.Stderr)
	}
	fields := strings.Split(result.Stdout, "\x00")
	if len(fields) == 0 || fields[len(fields)-1] != "" || (len(fields)-1)%2 != 0 {
		return nil, errors.New("malformed saved attempt delta")
	}
	files := make([]savedFile, 0, (len(fields)-1)/2)
	for i := 0; i < len(fields)-1; i += 2 {
		status, path := fields[i], fields[i+1]
		if len(status) != 1 || path == "" {
			return nil, fmt.Errorf("malformed saved attempt entry %q %q", status, path)
		}
		item := savedFile{path: path, status: status[0]}
		switch item.status {
		case 'A', 'M', 'T':
			item.data, err = gitstore.ReadAt(root, oid, path)
			if err != nil {
				return nil, fmt.Errorf("read saved file %s: %w", path, err)
			}
		case 'D':
		default:
			return nil, fmt.Errorf("unsupported saved file status %q for %s", status, path)
		}
		files = append(files, item)
	}
	return files, nil
}

func restoredChange(current, saved contracts.Document, phase string) (string, error) {
	for _, name := range []string{"Decisions", "Questions", "Rejected alternatives", "Notes"} {
		if body, ok := contracts.Section(saved.Body, name); ok {
			current.Body = contracts.ReplaceSection(current.Body, name, body)
		}
	}
	if phase == "audit" {
		if body, ok := contracts.Section(saved.Body, "Audit"); ok {
			current.Body = contracts.ReplaceSection(current.Body, "Audit", body)
		}
	}
	return contracts.RenderDocument(current)
}

func StartFromAttempt(root, change string, attempt int) (model.State, error) {
	if attempt <= 0 {
		return model.State{}, errors.New("P01: attempt number must be positive")
	}
	repo, err := gitstore.Inspect(root)
	if err != nil {
		return model.State{}, err
	}
	root = repo.Root
	ch, err := changeBy(root, change)
	if err != nil {
		return model.State{}, err
	}
	if str(ch.Document.Frontmatter["run"]) == "" || str(ch.Document.Frontmatter["phase"]) == "" {
		return model.State{}, errors.New("P01: saved attempt requires a committed phase")
	}
	name := fmt.Sprintf("refs/fireweed/attempts/%s/%d", ch.Slug, attempt)
	oid, ok, err := gitstore.Ref(root, name)
	if err != nil {
		return model.State{}, fmt.Errorf("read saved attempt %s: %w", name, err)
	}
	if !ok {
		return model.State{}, fmt.Errorf("P01: saved attempt %s does not exist", name)
	}
	files, err := savedStep(root, oid)
	if err != nil {
		return model.State{}, err
	}
	var savedDoc contracts.Document
	for _, item := range files {
		if item.path == ch.Path {
			if item.status == 'D' {
				return model.State{}, fmt.Errorf("saved attempt deleted change %s", ch.Path)
			}
			savedDoc, err = contracts.ParseDocument(string(item.data))
			if err != nil {
				return model.State{}, fmt.Errorf("parse saved change %s: %w", ch.Path, err)
			}
			break
		}
	}

	s, err := Start(root, ch.Slug)
	if err != nil {
		return s, err
	}
	for _, item := range files {
		if item.path == ch.Path {
			continue
		}
		if item.status == 'D' {
			err = os.Remove(file(root, item.path))
			if os.IsNotExist(err) {
				err = nil
			}
		} else {
			_, err = blockGit(root, nil, "restore", "--source="+oid, "--worktree", "--", item.path)
		}
		if err != nil {
			return s, fmt.Errorf("restore saved file %s: %w", item.path, err)
		}
	}
	if savedDoc.Body != "" {
		data, readErr := read(root, ch.Path)
		if readErr != nil {
			return s, fmt.Errorf("read current change %s: %w", ch.Path, readErr)
		}
		current, parseErr := contracts.ParseDocument(string(data))
		if parseErr != nil {
			return s, fmt.Errorf("parse current change %s: %w", ch.Path, parseErr)
		}
		rendered, renderErr := restoredChange(current, savedDoc, s.Phase)
		if renderErr != nil {
			return s, fmt.Errorf("render restored change %s: %w", ch.Path, renderErr)
		}
		if err = write(root, ch.Path, []byte(rendered)); err != nil {
			return s, fmt.Errorf("restore change body %s: %w", ch.Path, err)
		}
	}
	s.Change, err = changeBy(root, ch.Slug)
	if err != nil {
		return s, fmt.Errorf("load restored change %s: %w", ch.Path, err)
	}
	return s, nil
}
