package start

import (
	"context"
	"github.com/lurkshark/fireweed/packages/checks"
	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/gitstore"
	"github.com/lurkshark/fireweed/packages/workflow/internal/contentapproval"
	"os"
	"strings"
	"time"
)

type localGit struct {
	root, common string
	objects      map[string]bool
	commits      map[string]commitData
	trees        map[string][]treeEntry
	bodies       map[string]string
}

func (g localGit) capture(args ...string) gitstore.CaptureResult {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	return gitstore.Capture(ctx, g.root, append([]string{"git", "-c", "core.hooksPath=" + os.DevNull, "-c", "core.fsmonitor=false"}, args...), []string{"GIT_TERMINAL_PROMPT=0", "GIT_OPTIONAL_LOCKS=0", "GIT_NO_REPLACE_OBJECTS=1", "GIT_NO_LAZY_FETCH=1", "GIT_COMMITTER_NAME=Fireweed", "GIT_COMMITTER_EMAIL=fireweed@localhost"})
}
func (g localGit) raw(args ...string) (string, error) {
	r := g.capture(args...)
	if r.Err != nil || r.Exit != 0 {
		return "", problem("local Git %s failed (exit %d): %s %v", strings.Join(args, " "), r.Exit, strings.TrimSpace(r.Stdout+r.Stderr), r.Err)
	}
	return r.Stdout, nil
}
func (g localGit) text(args ...string) (string, error) {
	s, e := g.raw(args...)
	return strings.TrimSpace(s), e
}
func (g localGit) approvalWarnings(base, head string, m contracts.Module, current, path string) []checks.Diagnostic {
	ds := make([]checks.Diagnostic, 0)
	warn := func(s string) {
		ds = append(ds, checks.Diagnostic{Code: "P04", Severity: "warning", Path: path, Message: s + ". Check the local base and merge approved authored content before unattended execution."})
	}
	pinned, e := g.resolveCommit(base)
	if e != nil {
		warn("Base tip ancestry is unavailable: " + e.Error())
		warn("Selected authored content is unapproved because base content is unavailable")
		return ds
	}
	if !g.ancestor(pinned, head) {
		warn("HEAD does not contain the locally known base tip " + pinned)
	}
	baseText, e := g.file(pinned, path)
	if e != nil {
		warn("Selected authored content is unapproved because it is absent or unavailable on the pinned base: " + e.Error())
		return ds
	}
	approved, e := contentapproval.Matches(m, baseText, current, path)
	if e != nil || !approved {
		warn("Selected authored content is unapproved on the pinned base")
	}
	return ds
}

// A report covers reachable apply events, not commit messages or mere edits to
// an already applied archive. Inspect parent states to identify each new apply.
func (g localGit) currentDrill(module contracts.Module, head string) bool {
	path := ".fireweed/drills/" + module.Name + ".md"
	n, e := readNode(g.root, path, false)
	if e != nil {
		return false
	}
	doc, e := contracts.ParseMarkdown(string(n.bytes))
	if e != nil {
		return false
	}
	result, _ := doc.Metadata["result"].(string)
	base, _ := doc.Metadata["base_commit"].(string)
	if result != "pass" || base == "" {
		return false
	}
	pinned, e := g.resolveCommit(base)
	if e != nil || !g.ancestor(pinned, head) {
		return false
	}
	later, e := g.later(pinned, head)
	if e != nil {
		return false
	}
	for id, commit := range later {
		applied, e := g.applies(id, module.Name)
		if e != nil {
			return false
		}
		inherited := map[string]bool{}
		for _, parent := range commit.parents {
			previous, e := g.applies(parent, module.Name)
			if e != nil {
				return false
			}
			for run := range previous {
				inherited[run] = true
			}
		}
		for run := range applied {
			if !inherited[run] {
				return false
			}
		}
	}
	return true
}
