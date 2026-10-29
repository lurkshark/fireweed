package lifecycle

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/gitstore"
	"github.com/lurkshark/fireweed/packages/workflow/model"
)

func blockGit(root string, input []byte, args ...string) (string, error) {
	argv := append([]string{"git", "-C", root, "-c", "core.hooksPath=/dev/null"}, args...)
	result, err := gitstore.Run(context.Background(), gitstore.Request{
		Args: argv, Input: input, Timeout: 30 * time.Second,
		Env: map[string]string{"GIT_NO_LAZY_FETCH": "1", "GIT_TERMINAL_PROMPT": "0"},
	})
	if err != nil {
		return "", fmt.Errorf("git %s: %w: %s", strings.Join(args, " "), err, result.Stderr)
	}
	return strings.TrimSpace(result.Stdout), nil
}

// The phase base is the last committed approval. Start's assigned IDs, run,
// and phase are held in the claim until the next phase commit.
func blockDocument(root string, s model.State, c claim) (contracts.Document, error) {
	base, err := gitstore.ReadAt(root, c.PhaseBase, s.Change.Path)
	if err != nil {
		return contracts.Document{}, err
	}
	content := string(base)
	baseDoc, err := contracts.ParseDocument(content)
	if err != nil {
		return contracts.Document{}, err
	}
	oldAdded, _ := contracts.Section(baseDoc.Body, "Added")
	newAdded := c.Sections["Added"]
	oldItems, oldErr := contracts.ParseItems(oldAdded)
	newItems, newErr := contracts.ParseItems(newAdded)
	if oldErr == nil && newErr == nil && len(oldItems) == len(newItems) {
		assigned := map[string]string{}
		for i, item := range oldItems {
			if strings.HasPrefix(item.ID, "new:") {
				assigned[item.ID] = newItems[i].ID
			}
		}
		content = handle.ReplaceAllStringFunc(content, func(id string) string {
			if replacement := assigned[id]; replacement != "" {
				return replacement
			}
			return id
		})
	}
	doc, err := contracts.ParseDocument(content)
	if err != nil {
		return contracts.Document{}, err
	}
	for _, name := range authorNames {
		if current, exists := contracts.Section(doc.Body, name); exists && current != c.Sections[name] {
			doc.Body = contracts.ReplaceSection(doc.Body, name, c.Sections[name])
		}
	}
	doc.Frontmatter = cloneMap(c.Author)
	for key, value := range c.Audit {
		doc.Frontmatter[key] = value
	}
	doc.Frontmatter["run"] = c.Run
	doc.Frontmatter["phase"] = c.Phase
	return doc, nil
}

func nextAttempt(root, slug, oid string) (string, error) {
	for n := 1; ; n++ {
		name := fmt.Sprintf("refs/fireweed/attempts/%s/%d", slug, n)
		_, exists, err := gitstore.Ref(root, name)
		if err != nil {
			return "", err
		}
		if exists {
			continue
		}
		if err := gitstore.CASRef(root, name, "", oid); err != nil {
			// A second worktree may have taken this number since the lookup.
			if _, taken, lookupErr := gitstore.Ref(root, name); lookupErr == nil && taken {
				continue
			}
			return "", err
		}
		return name, nil
	}
}

func Block(request model.BlockRequest) (model.BlockResult, error) {
	var result model.BlockResult
	s, c, err := loadCurrent(request.Root, request.Change)
	result.State = s
	if err != nil {
		return result, err
	}
	if err := validate(s, c); err != nil {
		return result, err
	}
	if s.Head != c.PhaseBase {
		return result, errors.New("P05: HEAD differs from the held phase base")
	}
	if strings.TrimSpace(request.Reason) == "" {
		return result, errors.New("blocked reason is empty")
	}
	switch c.Phase {
	case "evals", "audit", "impl":
	default:
		return result, fmt.Errorf("P01: unsupported phase %s", c.Phase)
	}
	root := s.Snapshot.Root
	doc, err := blockDocument(root, s, c)
	if err != nil {
		return result, err
	}
	if _, err = blockGit(root, nil, "add", "-A"); err != nil {
		return result, err
	}
	tree, err := blockGit(root, nil, "write-tree")
	if err != nil {
		return result, err
	}
	attemptOID, err := blockGit(root, []byte("attempt: "+s.Change.Slug+"\n"), "commit-tree", tree, "-p", c.PhaseBase)
	if err != nil {
		return result, err
	}
	attemptRef, err := nextAttempt(root, s.Change.Slug, attemptOID)
	if err != nil {
		return result, err
	}
	result.AttemptRef, result.AttemptOID = attemptRef, attemptOID
	if _, err = blockGit(root, nil, "reset", "--hard", c.PhaseBase); err != nil {
		return result, err
	}
	if _, err = blockGit(root, nil, "clean", "-fd"); err != nil {
		return result, err
	}
	doc.Frontmatter["blocked"] = request.Reason
	rendered, err := contracts.RenderDocument(doc)
	if err != nil {
		return result, err
	}
	if err = write(root, s.Change.Path, []byte(rendered)); err != nil {
		return result, err
	}
	head, err := gitstore.Commit(root, "chore("+s.Change.Module+"): block "+s.Change.Slug, []string{s.Change.Path})
	if err != nil {
		return result, err
	}
	if err = gitstore.CASRef(root, ref(s.Change.Slug), s.ClaimOID, ""); err != nil {
		return result, err
	}
	c.PhaseBase = head
	s.Change.Document = doc
	result.State = state(root, s.Change, c, "", head)
	return result, nil
}
