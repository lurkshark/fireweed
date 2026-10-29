package lifecycle

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/lurkshark/fireweed/packages/assets"
	"github.com/lurkshark/fireweed/packages/checks"
	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/gitstore"
	"github.com/lurkshark/fireweed/packages/project"
	"github.com/lurkshark/fireweed/packages/workflow/model"
)

type claim struct {
	Worktree  string            `json:"worktree"`
	Change    string            `json:"change"`
	Run       string            `json:"run"`
	Phase     string            `json:"phase"`
	PhaseBase string            `json:"phase_base"`
	Snapshot  project.Snapshot  `json:"snapshot"`
	Author    map[string]any    `json:"author"`
	Sections  map[string]string `json:"sections"`
	Audit     map[string]any    `json:"audit"`
}

func ref(slug string) string                 { return "refs/fireweed/claims/" + slug }
func file(root, path string) string          { return filepath.Join(root, filepath.FromSlash(path)) }
func read(root, path string) ([]byte, error) { return os.ReadFile(file(root, path)) }
func write(root, path string, data []byte) error {
	p := file(root, path)
	if err := os.MkdirAll(filepath.Dir(p), 0755); err != nil {
		return err
	}
	return os.WriteFile(p, data, 0644)
}
func hard(ds []project.Diagnostic) error {
	for _, d := range ds {
		if !d.Warning {
			if d.Code == "C08" && strings.HasPrefix(d.Message, "stale base for ") {
				return fmt.Errorf("P05: C08: %s", d.Message)
			}
			if strings.HasPrefix(d.Message, d.Code+": ") {
				return errors.New(d.Message)
			}
			return fmt.Errorf("%s: %s", d.Code, d.Message)
		}
	}
	return nil
}
func diag(code, msg string, warning bool) project.Diagnostic {
	return project.Diagnostic{Code: code, Message: msg, Warning: warning}
}
func str(v any) string { s, _ := v.(string); return s }
func num(v any, fallback int) int {
	switch x := v.(type) {
	case int:
		return x
	case float64:
		return int(x)
	case string:
		n, e := strconv.Atoi(x)
		if e == nil {
			return n
		}
	}
	return fallback
}
func mapval(v any) map[string]any { m, _ := v.(map[string]any); return m }
func list(v any) []string {
	var out []string
	switch x := v.(type) {
	case []string:
		return x
	case []any:
		for _, a := range x {
			if s, ok := a.(string); ok {
				out = append(out, s)
			}
		}
	}
	return out
}
func cloneMap(m map[string]any) map[string]any {
	o := map[string]any{}
	for k, v := range m {
		o[k] = v
	}
	return o
}

var authorNames = []string{"Motivation", "Interface", "Added", "Modified", "Removed", "Removed decisions"}

func auditState(doc contracts.Document) map[string]any {
	state := map[string]any{}
	for _, key := range []string{"audit_rounds", "audit_pending"} {
		if value, ok := doc.Frontmatter[key]; ok {
			state[key] = value
		}
	}
	return state
}

func author(doc contracts.Document) (map[string]any, map[string]string) {
	m := cloneMap(doc.Frontmatter)
	for _, k := range []string{"run", "phase", "blocked", "audit_rounds", "audit_pending"} {
		delete(m, k)
	}
	s := map[string]string{}
	for _, n := range authorNames {
		s[n], _ = contracts.Section(doc.Body, n)
	}
	return m, s
}
func equal(a, b any) bool {
	aa, _ := json.Marshal(a)
	bb, _ := json.Marshal(b)
	return string(aa) == string(bb)
}
func id() (string, error) {
	b := make([]byte, 12)
	if _, e := rand.Read(b); e != nil {
		return "", e
	}
	return time.Now().UTC().Format("20060102T150405") + "-" + hex.EncodeToString(b), nil
}
func changeBy(root, target string) (project.Change, error) {
	q, e := project.Queue(root)
	if e != nil {
		return project.Change{}, e
	}
	for _, c := range q {
		if target == "" || target == c.Slug || target == c.Path || target == filepath.Base(c.Path) {
			return c, nil
		}
	}
	return project.Change{}, fmt.Errorf("U02: change %q not found", target)
}
func existing(root, slug string) (claim, string, error) {
	oid, ok, e := gitstore.Ref(root, ref(slug))
	if e != nil {
		return claim{}, "", fmt.Errorf("P02: read claim: %w", e)
	}
	if !ok {
		return claim{}, "", errors.New("P02: change has no claim")
	}
	b, e := gitstore.ReadBlob(root, oid)
	if e != nil {
		return claim{}, "", fmt.Errorf("P02: read claim: %w", e)
	}
	var c claim
	if e = json.Unmarshal(b, &c); e != nil {
		return claim{}, "", fmt.Errorf("P02: decode claim: %w", e)
	}
	return c, oid, nil
}
func state(root string, ch project.Change, c claim, oid, head string) model.State {
	return model.State{Change: ch, Snapshot: c.Snapshot, Run: c.Run, Phase: c.Phase, Head: head, PhaseBase: c.PhaseBase, ClaimOID: oid}
}
func storeClaim(root string, c claim, old string) (string, error) {
	b, e := json.Marshal(c)
	if e != nil {
		return "", e
	}
	oid, e := gitstore.WriteBlob(root, b)
	if e != nil {
		return "", e
	}
	if e = gitstore.CASRef(root, ref(c.Change), old, oid); e != nil {
		return "", e
	}
	return oid, nil
}
func baseRefs(root string, config map[string]any) (local, shared string) {
	base := str(config["base"])
	if base == "" {
		base = "main"
	}
	remote := str(config["remote"])
	if remote == "" {
		remote = "origin"
	}
	local = base
	shared = remote + "/" + base
	if _, ok, _ := gitstore.Ref(root, "refs/remotes/"+shared); !ok {
		shared = local
	}
	return
}
func approved(root, rev, path string) bool { _, e := gitstore.ReadAt(root, rev, path); return e == nil }

var handle = regexp.MustCompile(`new:[a-z0-9-]+`)

func assign(root string, ch project.Change, snap project.Snapshot) (project.Change, error) {
	module := project.Module{}
	for _, m := range snap.Modules {
		if m.Name == ch.Module {
			module = m
			break
		}
	}
	used := map[int]bool{}
	scan := func(items []contracts.Item) {
		for _, it := range items {
			var n int
			if _, e := fmt.Sscanf(it.ID, module.IDPrefix+"-%d", &n); e == nil {
				used[n] = true
			}
		}
	}
	baseSnapshot, _ := project.Load(root)
	for _, m := range baseSnapshot.Modules {
		if m.Name == ch.Module {
			s, _ := contracts.Section(m.Document.Body, "Specs")
			items, _ := contracts.ParseItems(s)
			scan(items)
		}
	}
	entries, _ := os.ReadDir(file(root, "changes/archive"))
	for _, ent := range entries {
		if ent.IsDir() || !strings.HasSuffix(ent.Name(), ".md") {
			continue
		}
		b, e := read(root, "changes/archive/"+ent.Name())
		if e != nil {
			continue
		}
		d, e := contracts.ParseDocument(string(b))
		if e != nil || str(d.Frontmatter["module"]) != ch.Module {
			continue
		}
		for _, section := range []string{"Added", "Modified", "Removed"} {
			s, _ := contracts.Section(d.Body, section)
			its, _ := contracts.ParseItems(s)
			scan(its)
		}
	}
	b, e := read(root, ch.Path)
	if e != nil {
		return ch, e
	}
	mapping := map[string]string{}
	for _, it := range ch.Added {
		if !strings.HasPrefix(it.ID, "new:") {
			continue
		}
		if mapping[it.ID] != "" {
			continue
		}
		n := 1
		for used[n] {
			n++
		}
		used[n] = true
		mapping[it.ID] = fmt.Sprintf("%s-%03d", module.IDPrefix, n)
	}
	content := handle.ReplaceAllStringFunc(string(b), func(h string) string {
		if mapped := mapping[h]; mapped != "" {
			return mapped
		}
		return h
	})
	if content != string(b) {
		if e = write(root, ch.Path, []byte(content)); e != nil {
			return ch, e
		}
		return changeBy(root, ch.Slug)
	}
	return ch, nil
}
func Start(root, target string) (model.State, error) {
	repo, e := gitstore.Inspect(root)
	if e != nil {
		return model.State{}, e
	}
	root = repo.Root
	changes, e := gitstore.Changes(root, repo.Head)
	if e != nil {
		return model.State{}, e
	}
	if len(changes) > 0 {
		return model.State{}, errors.New("P01: start requires clean worktree")
	}
	if e = hard(project.Check(root)); e != nil {
		return model.State{}, e
	}
	q, e := Eligible(root)
	if e != nil {
		return model.State{}, e
	}
	if len(q) == 0 {
		return model.State{}, errors.New("P01: no eligible change")
	}
	ch := q[0]
	if target != "" && target != ch.Slug && target != ch.Path && target != filepath.Base(ch.Path) {
		return model.State{}, errors.New("P01: change is not eligible queue head")
	}
	local, shared := baseRefs(root, mustLoad(root).Config)
	ds := []project.Diagnostic{}
	if !approved(root, local, ch.Path) && !approved(root, shared, ch.Path) {
		ds = append(ds, diag("P04", "change is not approved on local or shared base", true))
	}
	limit := num(mapval(mustLoad(root).Config["authoring"])["max_unapproved"], 5)
	if limit > 0 {
		entries, _ := os.ReadDir(file(root, "changes/archive"))
		count := 0
		for _, ent := range entries {
			if ent.IsDir() || !strings.HasSuffix(ent.Name(), ".md") {
				continue
			}
			p := "changes/archive/" + ent.Name()
			if !approved(root, shared, p) && !approved(root, shared, "changes/"+ent.Name()) {
				count++
			}
		}
		if count >= limit {
			ds = append(ds, diag("P07", "shared-base unapproved archive limit reached", false))
			return model.State{Diagnostics: ds}, errors.New("P07: unapproved archive limit")
		}
	}
	snap, e := project.Project(root, ch.Slug)
	if e != nil {
		return model.State{Diagnostics: ds}, e
	}
	run := str(ch.Document.Frontmatter["run"])
	if run == "" {
		run, e = id()
		if e != nil {
			return model.State{}, e
		}
	}
	phase := str(ch.Document.Frontmatter["phase"])
	if phase == "" {
		phase = "evals"
	}
	c := claim{Worktree: root, Change: ch.Slug, Run: run, Phase: phase, PhaseBase: repo.Head, Snapshot: snap}
	c.Author, c.Sections = author(ch.Document)
	c.Audit = auditState(ch.Document)
	oid, e := storeClaim(root, c, "")
	if e != nil {
		return model.State{Diagnostics: ds}, fmt.Errorf("P02: claim already held: %w", e)
	}
	rollback := func(err error) (model.State, error) {
		_ = gitstore.CASRef(root, ref(ch.Slug), oid, "")
		return model.State{Diagnostics: ds}, err
	}
	ch, e = assign(root, ch, snap)
	if e != nil {
		return rollback(e)
	}
	ch.Document.Frontmatter["run"] = run
	ch.Document.Frontmatter["phase"] = phase
	rendered, e := contracts.RenderDocument(ch.Document)
	if e != nil {
		return rollback(e)
	}
	if e = write(root, ch.Path, []byte(rendered)); e != nil {
		return rollback(e)
	}
	snap, e = project.Project(root, ch.Slug)
	if e != nil {
		return rollback(e)
	}
	c.Snapshot = snap
	c.Author, c.Sections = author(ch.Document)
	c.Audit = auditState(ch.Document)
	oid, e = storeClaim(root, c, oid)
	if e != nil {
		return model.State{}, fmt.Errorf("P02: update claim: %w", e)
	}
	s := state(root, ch, c, oid, repo.Head)
	s.Diagnostics = ds
	return s, nil
}
func mustLoad(root string) project.Snapshot { s, _ := project.Load(root); return s }
func Current(root, target string) (model.State, error) {
	repo, e := gitstore.Inspect(root)
	if e != nil {
		return model.State{}, e
	}
	root = repo.Root
	if target == "" {
		q, e := project.Queue(root)
		if e != nil {
			return model.State{}, e
		}
		for _, ch := range q {
			c, oid, e := existing(root, ch.Slug)
			if e == nil && c.Worktree == root && c.Change == ch.Slug {
				return state(root, ch, c, oid, repo.Head), nil
			}
		}
		return model.State{}, errors.New("P02: worktree holds no change")
	}
	ch, e := changeBy(root, target)
	if e != nil {
		return model.State{}, e
	}
	c, oid, e := existing(root, ch.Slug)
	if e != nil {
		return model.State{}, e
	}
	if c.Worktree != root || c.Change != ch.Slug {
		return model.State{}, errors.New("P02: claim belongs to another worktree")
	}
	return state(root, ch, c, oid, repo.Head), nil
}
func validate(s model.State, c claim) error {
	if s.Change.Slug != c.Change || str(s.Change.Document.Frontmatter["run"]) != c.Run || str(s.Change.Document.Frontmatter["phase"]) != c.Phase {
		return errors.New("P03: approved run or phase changed")
	}
	a, sections := author(s.Change.Document)
	if !equal(a, c.Author) || !equal(sections, c.Sections) {
		return errors.New("P03: approved author metadata changed")
	}
	if !equal(auditState(s.Change.Document), c.Audit) {
		return errors.New("P03: approved audit state changed")
	}
	return nil
}
func loadCurrent(root, target string) (model.State, claim, error) {
	s, e := Current(root, target)
	if e != nil {
		return s, claim{}, e
	}
	c, _, e := existing(s.Snapshot.Root, s.Change.Slug)
	return s, c, e
}
func specs(m project.Module) []contracts.Item {
	section, _ := contracts.Section(m.Document.Body, "Specs")
	its, _ := contracts.ParseItems(section)
	return its
}
func module(s project.Snapshot, name string) project.Module {
	for _, m := range s.Modules {
		if m.Name == name {
			return m
		}
	}
	return project.Module{}
}
func selectedIDs(ch project.Change, projected []contracts.Item) []string {
	var ids []string
	for _, it := range append(append([]contracts.Item{}, ch.Added...), ch.Modified...) {
		if !strings.Contains(it.Tag, "review") && !strings.Contains(it.Tag, "quarantine") {
			ids = append(ids, it.ID)
		}
	}
	for _, id := range list(ch.Document.Frontmatter["fixes"]) {
		for _, item := range projected {
			if item.ID == id && !strings.Contains(item.Tag, "review") && !strings.Contains(item.Tag, "quarantine") {
				ids = append(ids, id)
			}
		}
	}
	return ids
}
func dependents(s project.Snapshot, name, mode string) []string {
	if mode == "none" {
		return nil
	}
	seen := map[string]bool{name: true}
	out := []string{}
	for {
		added := false
		for _, m := range s.Modules {
			if seen[m.Name] {
				continue
			}
			for _, dep := range m.DependsOn {
				if seen[dep] && (mode == "all" || dep == name) {
					out = append(out, m.Name)
					seen[m.Name] = true
					added = true
					break
				}
			}
		}
		if mode != "all" || !added {
			break
		}
	}
	sort.Strings(out)
	return out
}
func reportPath(s model.State, phase, name string) string {
	return filepath.Join(s.Snapshot.Root, ".fireweed", "runs", s.Run, phase, name+".json")
}
func saveReport(s model.State, phase, name string, r checks.Report) {
	b, _ := json.MarshalIndent(r, "", "  ")
	p := reportPath(s, phase, name)
	_ = os.MkdirAll(filepath.Dir(p), 0755)
	_ = os.WriteFile(p, b, 0644)
}
func Verify(ctx context.Context, req model.Request) (model.Result, error) {
	s, c, e := loadCurrent(req.Root, req.Change)
	r := model.Result{State: s, Reports: map[string]checks.Report{}}
	if e != nil {
		return r, e
	}
	if e = validate(s, c); e != nil {
		return r, e
	}
	if e = hard(project.Check(s.Snapshot.Root)); e != nil {
		return r, e
	}
	diff, e := gitstore.Changes(s.Snapshot.Root, c.PhaseBase)
	if e != nil {
		return r, e
	}
	r.Diagnostics = checks.Zones(checks.ZoneRequest{Snapshot: c.Snapshot, Module: s.Change.Module, Phase: c.Phase, ChangePath: s.Change.Path, Base: c.PhaseBase, InterfaceDeclared: strings.TrimSpace(c.Sections["Interface"]) != "", Changes: diff})
	if e = hard(r.Diagnostics); e != nil {
		return r, e
	}
	if c.Phase == "audit" {
		return r, nil
	}
	if c.Phase != "evals" && c.Phase != "impl" {
		return r, fmt.Errorf("P01: unsupported phase %s", c.Phase)
	}
	names := []string{s.Change.Module}
	if c.Phase == "impl" && !req.Quick {
		mode := str(c.Snapshot.Config["verify_dependents"])
		if mode == "" {
			mode = "direct"
		}
		names = append(names, dependents(c.Snapshot, s.Change.Module, mode)...)
	}
	var first error
	for _, name := range names {
		rep, err := checks.Evaluate(ctx, checks.HookRequest{Snapshot: c.Snapshot, Module: name, Bin: req.Bin, Timeout: req.Timeout})
		r.Reports[name] = rep
		saveReport(s, c.Phase, name, rep)
		if err != nil {
			if first == nil {
				first = err
			}
			continue
		}
		items := specs(module(c.Snapshot, name))
		if name == s.Change.Module {
			if err = checks.RequireCoverage(rep, items); err != nil && first == nil {
				first = err
			}
			if c.Phase == "evals" {
				if err = checks.RequireFailFirst(rep, selectedIDs(s.Change, items), list(s.Change.Document.Frontmatter["already_satisfied"])); err != nil {
					r.Diagnostics = append(r.Diagnostics, diag("P2-01", err.Error(), false))
					if first == nil {
						first = err
					}
				}
			}
		}
		if c.Phase == "impl" {
			if err = checks.RequirePass(rep, items); err != nil && first == nil {
				first = err
			}
		}
	}
	if c.Phase == "impl" && !req.Quick {
		ds, err := checks.Boundaries(ctx, checks.BoundaryRequest{Snapshot: c.Snapshot, Bin: req.Bin, Timeout: req.Timeout})
		r.Diagnostics = append(r.Diagnostics, ds...)
		if err != nil && first == nil {
			for _, d := range ds {
				if d.Warning || d.Code == "" {
					continue
				}
				if strings.Contains(err.Error(), d.Message) {
					first = fmt.Errorf("%s: %w", d.Code, err)
				} else {
					first = fmt.Errorf("%s: %s: %w", d.Code, d.Message, err)
				}
				break
			}
			if first == nil {
				first = fmt.Errorf("X02: %w", err)
			}
		}
		if err = hard(ds); err != nil && first == nil {
			first = err
		}
	}
	return r, first
}
func provenance(ch project.Change, phase, agent string, reports map[string]checks.Report) (project.Change, error) {
	if agent == "" {
		agent = "manual"
	}
	playbook, e := assets.Playbook(phase)
	if e != nil {
		return ch, e
	}
	hash := sha256.Sum256([]byte(playbook))
	total, passing, failing := 0, 0, 0
	for _, r := range reports {
		for _, c := range r.Cases {
			total++
			if c.Result == "pass" {
				passing++
			} else if c.Result == "fail" || c.Result == "error" {
				failing++
			}
		}
	}
	line := fmt.Sprintf("- %s %s agent=%s fireweed=%s playbook=%x tests=%d failing=%d passing=%d\n", phase, time.Now().UTC().Format(time.RFC3339Nano), agent, assets.Version(), hash, total, failing, passing)
	body, _ := contracts.Section(ch.Document.Body, "Provenance")
	ch.Document.Body = contracts.ReplaceSection(ch.Document.Body, "Provenance", body+line)
	return ch, nil
}
func commitPhase(s model.State, c claim, oid string, ch project.Change, paths []string, subject, next string) (model.State, error) {
	root := s.Snapshot.Root
	ch.Document.Frontmatter["phase"] = next
	rendered, e := contracts.RenderDocument(ch.Document)
	if e != nil {
		return s, e
	}
	if e = write(root, ch.Path, []byte(rendered)); e != nil {
		return s, e
	}
	head, e := gitstore.Commit(root, subject, paths)
	if e != nil {
		return s, e
	}
	c.Phase = next
	c.PhaseBase = head
	c.Audit = auditState(ch.Document)
	oid, e = storeClaim(root, c, oid)
	if e != nil {
		return s, e
	}
	fresh, e := changeBy(root, ch.Slug)
	if e != nil {
		return s, e
	}
	return state(root, fresh, c, oid, head), nil
}
func auditProcess(s model.State, c claim) (project.Change, string, bool, error) {
	ch := s.Change
	body, _ := contracts.Section(ch.Document.Body, "Audit")
	all := map[string]string{}
	for _, it := range append(append([]contracts.Item{}, ch.Added...), ch.Modified...) {
		all[it.ID] = it.Body
	}
	for _, id := range list(ch.Document.Frontmatter["fixes"]) {
		if all[id] == "" {
			for _, it := range specs(module(c.Snapshot, ch.Module)) {
				if it.ID == id {
					all[id] = it.Body
				}
			}
		}
	}
	allowed := map[string]bool{}
	for _, id := range list(c.Audit["audit_pending"]) {
		allowed[id] = true
	}
	round := num(c.Audit["audit_rounds"], 0)
	var kept []string
	ids := map[string]bool{}
	untestable := false
	rx := regexp.MustCompile(`^\s*-\s*(\[untestable\]\s*)?([A-Za-z][A-Za-z0-9-]*-\d+):\s*"([^"]+)"\s*[—-]\s*(.+?)\s*$`)
	for _, line := range strings.Split(body, "\n") {
		m := rx.FindStringSubmatch(line)
		if m == nil {
			continue
		}
		if all[m[2]] == "" || !strings.Contains(all[m[2]], m[3]) {
			continue
		}
		if round > 0 && !allowed[m[2]] {
			continue
		}
		kept = append(kept, strings.TrimSpace(line))
		ids[m[2]] = true
		if m[1] != "" {
			untestable = true
		}
	}
	if len(kept) > 0 {
		sort.Strings(kept)
		ch.Document.Body = contracts.ReplaceSection(ch.Document.Body, "Audit", "\n"+strings.Join(kept, "\n")+"\n")
		ordered := make([]string, 0, len(ids))
		for id := range ids {
			ordered = append(ordered, id)
		}
		sort.Strings(ordered)
		ch.Document.Frontmatter["audit_pending"] = ordered
		if untestable || round >= num(mapval(c.Snapshot.Config["audit"])["max_rounds"], 2) {
			reason := "audit findings need review"
			ch.Document.Frontmatter["blocked"] = reason
			questions, _ := contracts.Section(ch.Document.Body, "Questions")
			ch.Document.Body = contracts.ReplaceSection(ch.Document.Body, "Questions", questions+"\n"+strings.Join(kept, "\n")+"\n")
			return ch, "blocked", true, nil
		}
		ch.Document.Frontmatter["audit_rounds"] = round + 1
		return ch, "evals", false, nil
	}
	checked := regexp.MustCompile(`(?m)^Checked IDs:\s*(.+)$`).FindStringSubmatch(body)
	if checked == nil || !strings.Contains(body, "No findings") {
		return ch, "", false, errors.New("P09: audit requires Checked IDs and No findings")
	}
	found := map[string]bool{}
	for _, id := range regexp.MustCompile(`[A-Za-z][A-Za-z0-9-]*-\d+`).FindAllString(checked[1], -1) {
		found[id] = true
	}
	for _, id := range selectedIDs(ch, specs(module(c.Snapshot, ch.Module))) {
		if !found[id] {
			return ch, "", false, fmt.Errorf("P09: audit did not check %s", id)
		}
	}
	ch.Document.Body = contracts.ReplaceSection(ch.Document.Body, "Audit", "\nNo findings; checked "+strings.Join(sortedKeys(found), ", ")+"\n")
	delete(ch.Document.Frontmatter, "audit_pending")
	return ch, "impl", false, nil
}
func sortedKeys(m map[string]bool) []string {
	out := make([]string, 0, len(m))
	for s := range m {
		out = append(out, s)
	}
	sort.Strings(out)
	return out
}
func rollbackFiles(root string, originals map[string][]byte) {
	for path, b := range originals {
		if b == nil {
			_ = os.Remove(file(root, path))
		} else {
			_ = write(root, path, b)
		}
	}
}
func apply(s model.State, c claim, ch project.Change, oid string) (model.State, error) {
	root := s.Snapshot.Root
	m := module(c.Snapshot, ch.Module)
	specPath := m.Dir + "/SPEC.md"
	original, e := read(root, specPath)
	if e != nil {
		return s, e
	}
	pending, e := read(root, ch.Path)
	if e != nil {
		return s, e
	}
	archive := "changes/archive/" + filepath.Base(ch.Path)
	if _, e = os.Stat(file(root, archive)); e == nil {
		return s, errors.New("archive already exists")
	}
	originals := map[string][]byte{specPath: original, ch.Path: pending, archive: nil}
	oldDoc, e := contracts.ParseDocument(string(original))
	if e != nil {
		return s, e
	}
	decision, _ := contracts.Section(oldDoc.Body, "Decisions")
	removed, _ := contracts.Section(ch.Document.Body, "Removed decisions")
	for _, line := range strings.Split(removed, "\n") {
		line = strings.TrimSpace(line)
		if line != "" {
			decision = strings.ReplaceAll(decision, line+"\n", "")
			decision = strings.ReplaceAll(decision, line, "")
		}
	}
	added, _ := contracts.Section(ch.Document.Body, "Decisions")
	added = strings.TrimSpace(added)
	if added != "" {
		if !strings.HasSuffix(decision, "\n") {
			decision += "\n"
		}
		decision += fmt.Sprintf("\n- %s: %s\n", ch.Slug, added)
	}
	doc := m.Document
	doc.Body = contracts.ReplaceSection(doc.Body, "Decisions", decision)
	rendered, e := contracts.RenderDocument(doc)
	if e != nil {
		return s, e
	}
	if e = write(root, specPath, []byte(rendered)); e != nil {
		return s, e
	}
	ch.Document.Frontmatter["phase"] = "applied"
	archived, e := contracts.RenderDocument(ch.Document)
	if e != nil {
		rollbackFiles(root, originals)
		return s, e
	}
	if e = write(root, archive, []byte(archived)); e != nil {
		rollbackFiles(root, originals)
		return s, e
	}
	if e = os.Remove(file(root, ch.Path)); e != nil {
		rollbackFiles(root, originals)
		return s, e
	}
	if e = hard(project.Check(root)); e != nil {
		rollbackFiles(root, originals)
		return s, e
	}
	diff, e := gitstore.Changes(root, c.PhaseBase)
	if e != nil {
		rollbackFiles(root, originals)
		return s, e
	}
	paths := []string{specPath, ch.Path, archive}
	for _, d := range diff {
		if d.Path != ch.Path && d.Path != specPath && d.Path != archive {
			paths = append(paths, d.Path)
		}
	}
	head, e := gitstore.Commit(root, "feat("+ch.Module+"): apply "+ch.Slug, paths)
	if e != nil {
		rollbackFiles(root, originals)
		return s, e
	}
	if e = gitstore.CASRef(root, ref(ch.Slug), oid, ""); e != nil {
		return s, e
	}
	c.Phase = "applied"
	c.PhaseBase = head
	s = state(root, ch, c, "", head)
	s.Phase = "applied"
	return s, nil
}
func Done(ctx context.Context, req model.Request) (model.Result, error) {
	req.Quick = false
	s, c, e := loadCurrent(req.Root, req.Change)
	result := model.Result{State: s, Reports: map[string]checks.Report{}}
	if e != nil {
		return result, e
	}
	if e = validate(s, c); e != nil {
		return result, e
	}
	result, e = Verify(ctx, req)
	if e != nil {
		return result, e
	}
	s = result.State
	// A hook can write files. Check its aftermath before committing or applying.
	fresh, e := Current(s.Snapshot.Root, s.Change.Slug)
	if e != nil {
		return result, e
	}
	c, oid, e := existing(s.Snapshot.Root, s.Change.Slug)
	if e != nil {
		return result, e
	}
	if e = validate(fresh, c); e != nil {
		return result, e
	}
	if e = hard(project.Check(s.Snapshot.Root)); e != nil {
		return result, e
	}
	diffAfter, e := gitstore.Changes(s.Snapshot.Root, c.PhaseBase)
	if e != nil {
		return result, e
	}
	postZones := checks.Zones(checks.ZoneRequest{Snapshot: c.Snapshot, Module: fresh.Change.Module, Phase: c.Phase, ChangePath: fresh.Change.Path, Base: c.PhaseBase, InterfaceDeclared: strings.TrimSpace(c.Sections["Interface"]) != "", Changes: diffAfter})
	result.Diagnostics = append(result.Diagnostics, postZones...)
	if e = hard(postZones); e != nil {
		return result, e
	}
	s = fresh
	result.State = fresh
	ch := s.Change
	next := ""
	blocked := false
	switch c.Phase {
	case "evals":
		next = "audit"
		if mapval(c.Snapshot.Config["audit"])["enabled"] == false {
			next = "impl"
		}
	case "audit":
		ch, next, blocked, e = auditProcess(s, c)
		if e != nil {
			return result, e
		}
	case "impl":
		next = "applied"
	default:
		return result, fmt.Errorf("P01: unsupported phase %s", c.Phase)
	}
	ch, e = provenance(ch, c.Phase, req.Agent, result.Reports)
	if e != nil {
		return result, e
	}
	if next == "applied" {
		result.State, e = apply(s, c, ch, oid)
		return result, e
	}
	diff, e := gitstore.Changes(s.Snapshot.Root, c.PhaseBase)
	if e != nil {
		return result, e
	}
	paths := []string{ch.Path}
	if c.Phase == "evals" {
		for _, d := range diff {
			if d.Path != ch.Path {
				paths = append(paths, d.Path)
			}
		}
	}
	subject := "test(" + ch.Module + "): evals " + ch.Slug
	if c.Phase == "audit" {
		subject = "docs(" + ch.Module + "): audit " + ch.Slug
	}
	if blocked {
		subject = "chore(" + ch.Module + "): block " + ch.Slug
	}
	if blocked {
		ch.Document.Frontmatter["phase"] = c.Phase
		rendered, e := contracts.RenderDocument(ch.Document)
		if e != nil {
			return result, e
		}
		if e = write(s.Snapshot.Root, ch.Path, []byte(rendered)); e != nil {
			return result, e
		}
		head, e := gitstore.Commit(s.Snapshot.Root, subject, paths)
		if e != nil {
			return result, e
		}
		if e = gitstore.CASRef(s.Snapshot.Root, ref(ch.Slug), oid, ""); e != nil {
			return result, e
		}
		result.State.Head = head
		result.State.PhaseBase = head
		result.State.Change = ch
		return result, errors.New("P06: change blocked")
	}
	result.State, e = commitPhase(s, c, oid, ch, paths, subject, next)
	return result, e
}
func Context(root, target string, full bool) (string, error) {
	s, e := Current(root, target)
	if e != nil {
		return "", e
	}
	var b strings.Builder
	raw, e := read(s.Snapshot.Root, s.Change.Path)
	if e != nil {
		return "", e
	}
	b.WriteString("# Change\n")
	b.Write(raw)
	m := module(s.Snapshot, s.Change.Module)
	spec, e := contracts.RenderDocument(m.Document)
	if e != nil {
		return "", e
	}
	b.WriteString("\n# Projected spec\n")
	b.WriteString(spec)
	b.WriteString("\n# Interface\n")
	for _, pat := range m.Zones.Interface {
		_ = filepath.WalkDir(file(s.Snapshot.Root, m.Dir), func(path string, d os.DirEntry, err error) error {
			if err != nil || d.IsDir() {
				return nil
			}
			rel, _ := filepath.Rel(file(s.Snapshot.Root, m.Dir), path)
			if !contracts.MatchGlob(pat, filepath.ToSlash(rel)) {
				return nil
			}
			data, e := os.ReadFile(path)
			if e == nil {
				b.WriteString("\n## " + filepath.ToSlash(filepath.Join(m.Dir, rel)) + "\n")
				b.Write(data)
			}
			return nil
		})
	}
	sys, e := read(s.Snapshot.Root, "SYSTEM.md")
	if e == nil {
		b.WriteString("\n# SYSTEM.md\n")
		b.Write(sys)
	}
	play, e := assets.Playbook(s.Phase)
	if e == nil {
		b.WriteString("\n# Playbook\n")
		b.WriteString(play)
	}
	b.WriteString("\n# Other relevant files\n")
	_ = filepath.WalkDir(file(s.Snapshot.Root, m.Dir), func(path string, d os.DirEntry, err error) error {
		if err != nil || d.IsDir() {
			return nil
		}
		rel, _ := filepath.Rel(s.Snapshot.Root, path)
		rel = filepath.ToSlash(rel)
		if rel == m.Dir+"/SPEC.md" {
			return nil
		}
		for _, p := range m.Zones.Interface {
			r := strings.TrimPrefix(rel, m.Dir+"/")
			if contracts.MatchGlob(p, r) {
				return nil
			}
		}
		info, e := d.Info()
		if e != nil {
			return nil
		}
		b.WriteString(fmt.Sprintf("- %s (%d bytes)\n", rel, info.Size()))
		if full {
			data, e := os.ReadFile(path)
			if e == nil {
				b.Write(data)
				b.WriteByte('\n')
			}
		}
		return nil
	})
	return b.String(), nil
}
