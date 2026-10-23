package claims

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"time"

	"github.com/lurkshark/fireweed/packages/gitstore"
	"github.com/lurkshark/fireweed/packages/workflow/model"
)

var slugPattern = regexp.MustCompile(`^[a-z0-9][a-z0-9-]*$`)

const prefix = "refs/fireweed/claims/"

type storage struct {
	repo  gitstore.Repository
	ctx   context.Context
	width int
}

func failure(format string, args ...any) error {
	return fmt.Errorf("P02: %s. Use fireweed start or release after checking the local claim and its owner.", fmt.Sprintf(format, args...))
}
func open(start string) (*storage, func(), error) {
	repo, err := gitstore.Open(start)
	if err != nil {
		return nil, func() {}, failure("cannot open the claim repository: %v", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	s := &storage{repo: repo, ctx: ctx}
	format, err := s.git("rev-parse", "--show-object-format")
	if err != nil {
		cancel()
		return nil, func() {}, err
	}
	switch format {
	case "sha1":
		s.width = 40
	case "sha256":
		s.width = 64
	default:
		cancel()
		return nil, func() {}, failure("unsupported local Git object format %q", format)
	}
	return s, cancel, nil
}
func (s *storage) capture(args ...string) gitstore.CaptureResult {
	argv := append([]string{"git", "-c", "core.hooksPath=" + os.DevNull}, args...)
	return gitstore.Capture(s.ctx, s.repo.Root, argv, []string{"GIT_TERMINAL_PROMPT=0", "GIT_OPTIONAL_LOCKS=0", "GIT_NO_REPLACE_OBJECTS=1", "GIT_NO_LAZY_FETCH=1", "GIT_COMMITTER_NAME=Fireweed", "GIT_COMMITTER_EMAIL=fireweed@localhost"})
}
func (s *storage) run(args ...string) gitstore.CommandResult {
	r := s.capture(args...)
	return gitstore.CommandResult{Output: r.Stdout + r.Stderr, Exit: r.Exit, Err: r.Err}
}
func (s *storage) git(args ...string) (string, error) {
	r := s.capture(args...)
	if r.Err != nil || r.Exit != 0 {
		return "", failure("local Git %s failed (exit %d): %s %v", args[0], r.Exit, strings.TrimSpace(r.Stdout+r.Stderr), r.Err)
	}
	return strings.TrimSpace(r.Stdout), nil
}
func (s *storage) oid(value string) bool {
	return len(value) == s.width && strings.Trim(value, "0123456789abcdef") == "" && strings.Trim(value, "0") != ""
}
func validSlug(slug string) error {
	if !slugPattern.MatchString(slug) {
		return failure("invalid claim slug %q", slug)
	}
	return nil
}
func (s *storage) commit(value string) (string, error) {
	if strings.TrimSpace(value) == "" {
		return "", failure("a local commit baseline is required")
	}
	resolved, err := s.git("rev-parse", "--verify", "--end-of-options", value+"^{commit}")
	if err != nil {
		return "", err
	}
	if !s.oid(resolved) {
		return "", failure("baseline %q did not resolve to a full commit ID", value)
	}
	return resolved, nil
}

// Ref updates never dereference a selected name. Reject known aliases as well,
// so ordinary claim inspection cannot mistake another ref for claim ownership.
func (s *storage) direct(slug string) error {
	r := s.run("symbolic-ref", "--quiet", prefix+slug)
	if r.Exit == 0 && r.Err == nil {
		return failure("claim %s is a symbolic alias; inspect it before releasing or replacing it", slug)
	}
	if r.Exit != 1 {
		return failure("cannot inspect claim %s: %s %v", slug, strings.TrimSpace(r.Output), r.Err)
	}
	return nil
}
func (s *storage) object(slug string) (string, error) {
	if err := validSlug(slug); err != nil {
		return "", err
	}
	if err := s.direct(slug); err != nil {
		return "", err
	}
	oid, err := s.git("rev-parse", "--verify", "--end-of-options", prefix+slug)
	if err != nil {
		return "", err
	}
	if !s.oid(oid) {
		return "", failure("claim %s has an invalid object token", slug)
	}
	return oid, nil
}
func (s *storage) load(slug string) (model.Claim, error) {
	zero := model.Claim{}
	object, err := s.object(slug)
	if err != nil {
		return zero, err
	}
	kind, err := s.git("cat-file", "-t", object)
	if err != nil {
		return zero, err
	}
	if kind != "blob" {
		return zero, failure("claim %s does not name a JSON blob", slug)
	}
	raw, err := s.git("cat-file", "blob", object)
	if err != nil {
		return zero, err
	}
	var values map[string]json.RawMessage
	if err = json.Unmarshal([]byte(raw), &values); err != nil {
		return zero, failure("claim %s contains invalid JSON: %v", slug, err)
	}
	fields := make(map[string]string, 5)
	for _, name := range []string{"worktree", "change", "run", "base", "state"} {
		encoded, ok := values[name]
		if !ok {
			return zero, failure("claim %s is missing %s", slug, name)
		}
		var decoded any
		if err = json.Unmarshal(encoded, &decoded); err != nil {
			return zero, failure("claim %s has invalid %s", slug, name)
		}
		text, ok := decoded.(string)
		if !ok {
			return zero, failure("claim %s requires a string %s", slug, name)
		}
		fields[name] = text
	}
	c := model.Claim{Worktree: fields["worktree"], Change: fields["change"], Run: fields["run"], Base: fields["base"], State: fields["state"], Object: object}
	if c.Change != slug || !filepath.IsAbs(c.Worktree) || filepath.Clean(c.Worktree) != c.Worktree || strings.TrimSpace(c.Run) == "" || !s.oid(c.Base) {
		return zero, failure("claim %s has invalid target, owner, run or baseline fields", slug)
	}
	commit, err := s.commit(c.Base)
	if err != nil {
		return zero, err
	}
	if commit != c.Base {
		return zero, failure("claim %s baseline is not a full commit object", slug)
	}
	return c, nil
}
func (s *storage) blob(c model.Claim) (string, error) {
	raw, err := json.Marshal(c)
	if err != nil {
		return "", failure("cannot encode claim: %v", err)
	}
	// Temporary input is private to this invocation and belongs in the common Git
	// directory. It is neither an index nor a working-tree file.
	file, err := os.CreateTemp(s.repo.CommonDir, "fireweed-claim-")
	if err != nil {
		return "", failure("cannot create owned claim storage: %v", err)
	}
	name := file.Name()
	defer os.Remove(name)
	if _, err = file.Write(raw); err != nil {
		file.Close()
		return "", failure("cannot write owned claim storage: %v", err)
	}
	if err = file.Close(); err != nil {
		return "", failure("cannot close owned claim storage: %v", err)
	}
	object, err := s.git("hash-object", "-w", "--", name)
	if err != nil {
		return "", err
	}
	if !s.oid(object) {
		return "", failure("Git did not return a full claim blob ID")
	}
	return object, nil
}
func Create(start, slug, run, base, state string) (model.Claim, error) {
	zero := model.Claim{}
	if err := validSlug(slug); err != nil {
		return zero, err
	}
	if strings.TrimSpace(run) == "" {
		return zero, failure("a nonblank run is required for claim %s", slug)
	}
	s, close, err := open(start)
	if err != nil {
		return zero, err
	}
	defer close()
	resolved, err := s.commit(base)
	if err != nil {
		return zero, err
	}
	if err = s.direct(slug); err != nil {
		return zero, err
	}
	c := model.Claim{Worktree: s.repo.Root, Change: slug, Run: run, Base: resolved, State: state}
	c.Object, err = s.blob(c)
	if err != nil {
		return zero, err
	}
	if _, err = s.git("update-ref", "--no-deref", prefix+slug, c.Object, ""); err != nil {
		owner, readErr := s.load(slug)
		if readErr == nil {
			return zero, failure("claim %s is already held by %s; inspect that owner before release", slug, owner.Worktree)
		}
		return zero, err
	}
	return c, nil
}
func Load(start, slug string) (model.Claim, error) {
	if err := validSlug(slug); err != nil {
		return model.Claim{}, err
	}
	s, close, err := open(start)
	if err != nil {
		return model.Claim{}, err
	}
	defer close()
	return s.load(slug)
}
func List(start string) ([]model.Claim, error) {
	s, close, err := open(start)
	if err != nil {
		return nil, err
	}
	defer close()
	refs, err := s.git("for-each-ref", "--format=%(refname)", prefix)
	if err != nil {
		return nil, err
	}
	result := make([]model.Claim, 0)
	if refs != "" {
		for _, ref := range strings.Split(refs, "\n") {
			if !strings.HasPrefix(ref, prefix) {
				return nil, failure("unexpected claim ref %q", ref)
			}
			c, err := s.load(strings.TrimPrefix(ref, prefix))
			if err != nil {
				return nil, err
			}
			result = append(result, c)
		}
	}
	sort.Slice(result, func(i, j int) bool { return result[i].Change < result[j].Change })
	return result, nil
}
func Update(start string, previous model.Claim, base, state string) (model.Claim, error) {
	zero := model.Claim{}
	if err := validSlug(previous.Change); err != nil {
		return zero, err
	}
	s, close, err := open(start)
	if err != nil {
		return zero, err
	}
	defer close()
	actual, err := s.load(previous.Change)
	if err != nil {
		return zero, err
	}
	if actual != previous || actual.Worktree != s.repo.Root {
		return zero, failure("claim %s is stale, modified or held by another worktree", previous.Change)
	}
	resolved, err := s.commit(base)
	if err != nil {
		return zero, err
	}
	next := actual
	next.Base = resolved
	next.State = state
	next.Object = ""
	next.Object, err = s.blob(next)
	if err != nil {
		return zero, err
	}
	if _, err = s.git("update-ref", "--no-deref", prefix+next.Change, next.Object, previous.Object); err != nil {
		return zero, err
	}
	return next, nil
}
func Release(start, slug, expected string) error {
	if err := validSlug(slug); err != nil {
		return err
	}
	s, close, err := open(start)
	if err != nil {
		return err
	}
	defer close()
	if !s.oid(expected) {
		return failure("a full nonzero object token is required to release claim %s", slug)
	}
	actual, err := s.object(slug)
	if err != nil {
		return err
	}
	if actual != expected {
		return failure("claim %s has a newer ownership token", slug)
	}
	_, err = s.git("update-ref", "--no-deref", "-d", prefix+slug, expected)
	return err
}
