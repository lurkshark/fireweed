package evals

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"strings"
	"testing"

	c "github.com/lurkshark/fireweed/packages/contracts"
	p "github.com/lurkshark/fireweed/packages/project"
	w "github.com/lurkshark/fireweed/packages/workflow"
)

func eligibleModule(t *testing.T, name string, deps ...string) c.Module {
	t.Helper()
	prefix := strings.ToUpper(name[:3])
	body := "---\nname: " + name + "\nid_prefix: " + prefix + "\ndepends_on: [" + strings.Join(deps, ",") + "]\n---\n\n## Specs\n\n### " + prefix + "-001 [review]\nThe module MUST retain its contract.\n"
	m, err := c.ParseModule(body, "units/"+name, c.Zones{Interface: []string{"api.go"}, Evals: []string{"evals/**"}, Disposable: []string{"internal/**"}}, c.Hooks{Eval: "hook-trap", Boundaries: "hook-trap", Setup: "hook-trap"})
	if err != nil {
		t.Fatal(err)
	}
	return m
}
func eligibleChange(t *testing.T, n int, slug, module, meta, sections string) c.Change {
	t.Helper()
	prefix := strings.ToUpper(module[:3])
	path := fmt.Sprintf("changes/%04d-%s.md", n, slug)
	ch, err := c.ParseChange("---\nmodule: "+module+"\n"+meta+"---\n\n## Motivation\n\nChoose deterministic queued work.\n\n"+sections, path, prefix)
	if err != nil {
		t.Fatal(err)
	}
	return ch
}
func eligibleSource(t *testing.T) p.Snapshot {
	return p.Snapshot{Root: "/missing/eligibility-root", BaseRef: "never-read", Config: c.Config{Protected: []string{"private/**"}, WorkerWritable: []string{"shared/**"}, Agent: c.Agent{Command: []string{"agent-trap"}}, Hooks: c.Hooks{Eval: "hook-trap", Boundaries: "hook-trap", Setup: "hook-trap"}}, Modules: []c.Module{eligibleModule(t, "alpha"), eligibleModule(t, "bravo", "alpha"), eligibleModule(t, "charlie", "bravo"), eligibleModule(t, "delta"), eligibleModule(t, "echo", "charlie")}, Archive: []c.Change{eligibleChange(t, 99, "archived", "alpha", "", "")}, Diagnostics: []p.Diagnostic{{Code: "C07", Severity: "error", Message: "Unrelated static hook problem", Path: "units/delta/SPEC.md"}}}
}
func eligibleCopy(t *testing.T, s p.Snapshot) p.Snapshot {
	t.Helper()
	raw, err := json.Marshal(s)
	if err != nil {
		t.Fatal(err)
	}
	var out p.Snapshot
	if err = json.Unmarshal(raw, &out); err != nil {
		t.Fatal(err)
	}
	return out
}
func eligibleGraph(t *testing.T, s p.Snapshot, slug string) (p.Snapshot, c.Module, error) {
	t.Helper()
	s = eligibleCopy(t, s)
	sort.Slice(s.Changes, func(i, j int) bool {
		a, b := s.Changes[i], s.Changes[j]
		if a.Number != b.Number {
			return a.Number < b.Number
		}
		if a.Slug != b.Slug {
			return a.Slug < b.Slug
		}
		return a.Path < b.Path
	})
	before, err := p.Project(s, slug)
	if err != nil {
		return p.Snapshot{}, c.Module{}, err
	}
	var selected c.Change
	for _, ch := range s.Changes {
		if ch.Slug == slug {
			selected = ch
			break
		}
	}
	for _, m := range before.Modules {
		if m.Name == selected.Module {
			m, err = p.ApplyChange(m, selected)
			return before, m, err
		}
	}
	return before, c.Module{}, fmt.Errorf("missing target fixture")
}
func eligibleAssert(t *testing.T, s p.Snapshot, slug string, want bool) {
	t.Helper()
	before := policyJSON(t, s)
	got, err := w.Eligible(s, slug)
	if err != nil || got != want {
		t.Errorf("Eligible(%q)=%v,%v want%v,nil", slug, got, err, want)
	}
	if policyJSON(t, s) != before {
		t.Error("eligibility mutated source values")
	}
}
func eligibleError(t *testing.T, s p.Snapshot, slug, detail string) {
	t.Helper()
	before := policyJSON(t, s)
	got, err := w.Eligible(s, slug)
	if got || err == nil || !strings.HasPrefix(err.Error(), "C05:") || !policyCorrective(err.Error()) {
		t.Errorf("invalid prefix/selection result %v,%v", got, err)
	}
	if err != nil && detail != "" && !strings.Contains(err.Error(), detail) {
		t.Errorf("lost public projection detail %q: %v", detail, err)
	}
	if policyJSON(t, s) != before {
		t.Error("rejected eligibility mutated source")
	}
}
func Test_FLOW024EligibleExactSelectionAndQueueOrder(t *testing.T) {
	s := eligibleSource(t)
	a := eligibleChange(t, 2, "zeta", "alpha", "", "")
	b := eligibleChange(t, 10, "later", "delta", "", "")
	s.Changes = []c.Change{b, a}
	eligibleAssert(t, s, "zeta", true)
	eligibleAssert(t, s, "later", false)
	s.Changes = []c.Change{eligibleChange(t, 10000, "later", "delta", "", ""), eligibleChange(t, 9999, "zeta", "alpha", "", "")}
	eligibleAssert(t, s, "zeta", true)
	eligibleAssert(t, s, "later", false)
	s.Changes = []c.Change{eligibleChange(t, 2, "zeta", "alpha", "", ""), eligibleChange(t, 2, "alpha", "delta", "", "")}
	s.Changes[0].Path = "changes/00002-zeta.md"
	eligibleAssert(t, s, "alpha", true)
	eligibleAssert(t, s, "zeta", false)
	// Path is the final tie breaker for nonselected duplicate slug intent.
	first := eligibleChange(t, 1, "blocker", "bravo", "blocked: waiting\nset: {depends_on: [alpha]}\n", "")
	first.Path = "changes/00001-blocker.md"
	last := eligibleChange(t, 1, "blocker", "bravo", "blocked: waiting\nset: {depends_on: []}\n", "")
	last.Path = "changes/0001-blocker.md"
	target := eligibleChange(t, 2, "candidate", "alpha", "", "## Interface\n\nPublic contract edits.\n")
	s.Changes = []c.Change{target, last, first}
	beforeGraph, m, err := eligibleGraph(t, s, "candidate")
	for _, peer := range beforeGraph.Modules {
		if peer.Name == "bravo" && len(peer.DependsOn) != 0 {
			t.Fatal("path-ordered set fixture control")
		}
	}
	if err != nil || m.Name != "alpha" {
		t.Fatalf("ordered prefix fixture control %v", err)
	}
	eligibleAssert(t, s, "candidate", true)
	for _, slug := range []string{"", "../bad", "Upper", "prefix/current", "-bad", "z", "archived", "absent"} {
		eligibleError(t, s, slug, "")
	}
	s.Changes = append(s.Changes, eligibleChange(t, 99, "candidate", "delta", "", ""))
	eligibleError(t, s, "candidate", "")
}
func Test_FLOW025EligibleBlockedDependencyMatrix(t *testing.T) {
	for _, row := range []struct {
		name, blocked, target, inter string
		want                         bool
	}{
		{"same", "alpha", "alpha", "", false},
		{"forward-direct", "alpha", "bravo", "", false},
		{"forward-transitive", "alpha", "charlie", "", false},
		{"forward-long", "alpha", "echo", "", false},
		{"independent", "alpha", "delta", "", true},
		{"reverse-implementation", "charlie", "alpha", "", true},
		{"reverse-direct-interface", "bravo", "alpha", "## Interface\n\nRetained signature changes.\n", false},
		{"reverse-transitive-interface", "echo", "alpha", "## Interface\n\nRetained signature changes.\n", false},
		{"reverse-empty-interface", "bravo", "alpha", "## Interface\n\n", true},
		{"reverse-whitespace-interface", "bravo", "alpha", "## Interface\n\n \t \n", true},
		{"other-section", "bravo", "alpha", "## Notes\n\nInterface text is not a section.\n", true},
		{"fenced-interface", "bravo", "alpha", "## Notes\n\n```markdown\n## Interface\n\nInside a fence.\n```\n", true},
		{"wrong-case-interface", "bravo", "alpha", "## interface\n\nWrong case.\n", true},
	} {
		t.Run(row.name, func(t *testing.T) {
			s := eligibleSource(t)
			s.Changes = []c.Change{eligibleChange(t, 1, "blocked", row.blocked, "blocked: needs-an-answer\n", ""), eligibleChange(t, 2, "candidate", row.target, "", row.inter)}
			eligibleAssert(t, s, "blocked", false)
			eligibleAssert(t, s, "candidate", row.want)
		})
	}
}
func Test_FLOW025EligibleSkipsHeldAndKeepsEveryBlocker(t *testing.T) {
	s := eligibleSource(t)
	s.Changes = []c.Change{eligibleChange(t, 1, "first-block", "alpha", "blocked: first\n", ""), eligibleChange(t, 2, "second-block", "delta", "blocked: second\n", ""), eligibleChange(t, 3, "held", "charlie", "", ""), eligibleChange(t, 4, "candidate", "bravo", "", ""), eligibleChange(t, 5, "independent", "delta", "", "")}
	eligibleAssert(t, s, "held", false)
	eligibleAssert(t, s, "candidate", false)
	eligibleAssert(t, s, "independent", false)
	// A held predecessor does not prevent an independent head, and a later blocker
	// cannot hold work before it. Only one independent head can be eligible.
	s.Changes = []c.Change{eligibleChange(t, 1, "blocker", "alpha", "blocked: waiting\n", ""), eligibleChange(t, 2, "held", "charlie", "", ""), eligibleChange(t, 3, "candidate", "delta", "", ""), eligibleChange(t, 4, "second", "delta", "", ""), eligibleChange(t, 5, "later-block", "delta", "blocked: later\n", "")}
	eligibleAssert(t, s, "candidate", true)
	eligibleAssert(t, s, "second", false)
	// Neither an unrelated discovery error nor invalid later intent certifies or
	// blocks an otherwise valid selected prefix.
	s.Changes = append(s.Changes, eligibleChange(t, 99, "bad-later", "missing", "set: {depends_on: [missing]}\n", ""))
	eligibleAssert(t, s, "candidate", true)
}
func Test_FLOW026EligibleProjectedDependencyRelations(t *testing.T) {
	for _, row := range []struct {
		name, blockModule, blockMeta, target, targetMeta, inter string
		want                                                    bool
	}{
		{"blocked-set-removes-forward", "alpha", "set: {depends_on: []}\n", "bravo", "set: {depends_on: []}\n", "", true},
		{"blocked-set-adds-reverse", "delta", "set: {depends_on: [alpha]}\n", "alpha", "", "## Interface\n\nPublic edit.\n", false},
		{"selected-set-adds-forward", "delta", "", "alpha", "set: {depends_on: [delta]}\n", "", false},
		{"selected-set-removes-reverse", "bravo", "set: {depends_on: []}\n", "alpha", "", "## Interface\n\nPublic edit.\n", true},
	} {
		t.Run(row.name, func(t *testing.T) {
			s := eligibleSource(t)
			s.Changes = []c.Change{eligibleChange(t, 1, "blocker", row.blockModule, "blocked: waiting\n"+row.blockMeta, ""), eligibleChange(t, 2, "candidate", row.target, row.targetMeta, row.inter)}
			_, _, err := eligibleGraph(t, s, "candidate")
			if err != nil {
				t.Fatalf("valid projected graph control %v", err)
			}
			eligibleAssert(t, s, "candidate", row.want)
		})
	}
	t.Run("preceding-unblocked-set", func(t *testing.T) {
		s := eligibleSource(t)
		s.Changes = []c.Change{eligibleChange(t, 1, "blocker", "echo", "blocked: waiting\n", ""), eligibleChange(t, 2, "held-meta", "bravo", "set: {depends_on: [alpha,delta]}\n", "## Interface\n\nChanges the dependency contract.\n"), eligibleChange(t, 3, "candidate", "delta", "", "## Interface\n\nReverse relation changes.\n")}
		eligibleAssert(t, s, "candidate", false)
	})
	t.Run("apply-selected-once", func(t *testing.T) {
		s := eligibleSource(t)
		s.Changes = []c.Change{eligibleChange(t, 1, "candidate", "alpha", "set: {depends_on: []}\n", "## Added\n\n### new:future [review]\nThe target MUST keep this handle.\n")}
		_, m, err := eligibleGraph(t, s, "candidate")
		if err != nil || len(m.Specs) != 2 || m.Specs[1].ID != "new:future" {
			t.Fatalf("once projection control %v %+v", err, m)
		}
		eligibleAssert(t, s, "candidate", true)
	})
}
func Test_FLOW026EligibleEssentialGraphAndProjectionFailures(t *testing.T) {
	for _, mode := range []string{"unknown-target", "duplicate-target", "missing-dependency", "self", "cycle", "blocked-cycle", "behind-head-cycle", "prefix-apply", "selected-apply", "disconnected-cycle", "unknown-preceding-target"} {
		t.Run(mode, func(t *testing.T) {
			s := eligibleSource(t)
			s.Changes = []c.Change{eligibleChange(t, 1, "candidate", "alpha", "", "")}
			detail := ""
			switch mode {
			case "disconnected-cycle":
				s.Modules[1] = eligibleModule(t, "bravo", "charlie")
				s.Changes[0] = eligibleChange(t, 1, "candidate", "delta", "", "")
				if _, _, err := eligibleGraph(t, s, "candidate"); err != nil {
					t.Fatalf("disconnected graph shape control %v", err)
				}
			case "unknown-preceding-target":
				s.Changes = []c.Change{eligibleChange(t, 1, "missing-before", "missing", "blocked: waiting\n", ""), eligibleChange(t, 2, "candidate", "delta", "", "")}
				_, _, err := eligibleGraph(t, s, "candidate")
				if err == nil {
					t.Fatal("unknown preceding projection error premise")
				}
				detail = err.Error()
			case "unknown-target":
				s.Changes[0] = eligibleChange(t, 1, "candidate", "missing", "", "")
			case "duplicate-target":
				s.Modules = append(s.Modules, s.Modules[0])
			case "missing-dependency":
				s.Changes[0] = eligibleChange(t, 1, "candidate", "alpha", "set: {depends_on: [missing]}\n", "")
			case "self":
				s.Changes[0] = eligibleChange(t, 1, "candidate", "alpha", "set: {depends_on: [alpha]}\n", "")
			case "cycle", "blocked-cycle", "behind-head-cycle":
				meta := "set: {depends_on: [charlie]}\n"
				if mode == "blocked-cycle" {
					meta += "blocked: waiting\n"
				}
				s.Changes[0] = eligibleChange(t, 2, "candidate", "alpha", meta, "")
				if mode == "behind-head-cycle" {
					s.Changes = append([]c.Change{eligibleChange(t, 1, "earlier", "delta", "", "")}, s.Changes...)
				}
			case "prefix-apply":
				s.Changes = []c.Change{eligibleChange(t, 1, "broken", "alpha", "blocked: waiting\n", "## Added\n\n### ALP-001 [review]\nDuplicate MUST fail.\n"), eligibleChange(t, 2, "candidate", "delta", "", "")}
				_, _, err := eligibleGraph(t, s, "candidate")
				if err == nil {
					t.Fatal("prefix projection error premise")
				}
				detail = err.Error()
			case "selected-apply":
				s.Changes[0] = eligibleChange(t, 1, "candidate", "alpha", "", "## Added\n\n### ALP-001 [review]\nDuplicate MUST fail.\n")
				_, _, err := eligibleGraph(t, s, "candidate")
				if err == nil {
					t.Fatal("selected apply error premise")
				}
				detail = err.Error()
			}
			eligibleError(t, s, "candidate", detail)
		})
	}
}
func Test_FLOW027EligibleInputOwnershipAndAmbientIndependence(t *testing.T) {
	s := eligibleSource(t)
	s.Modules = []c.Module{s.Modules[4], s.Modules[1], s.Modules[0], s.Modules[3], s.Modules[2]}
	s.Changes = []c.Change{eligibleChange(t, 1, "blocker", "alpha", "blocked: waiting\n", ""), eligibleChange(t, 2, "candidate", "delta", "set: {depends_on: []}\n", "## Added\n\n### new:future [review]\nThis MUST remain unallocated.\n")}
	// Empty visible slices retain live sentinel data in their spare storage.
	backing := []string{"spare-one", "spare-two"}
	s.Modules[3].DependsOn = backing[:0]
	s.Modules[3].Document.Metadata["depends_on"] = backing[:0]
	_, projected, err := eligibleGraph(t, s, "candidate")
	if err != nil || len(projected.Specs) != 2 || projected.Specs[1].ID != "new:future" {
		t.Fatalf("ownership source projection control %v %+v", err, projected)
	}
	original := policyJSON(t, s)
	emptyBefore := append([]string(nil), backing...)
	cwd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	env := os.Environ()
	held := eligibleCopy(t, s)
	held.Changes[1].Blocked = "waiting"
	held.Changes[1].Document.Metadata["blocked"] = "waiting"
	if _, _, err := eligibleGraph(t, held, "candidate"); err != nil {
		t.Fatalf("held ambient source control %v", err)
	}
	_, selectorError := w.Eligible(s, "../bad")
	if selectorError == nil {
		t.Fatal("invalid selector unexpectedly accepted")
	}
	bad := eligibleCopy(t, s)
	bad.Changes = []c.Change{eligibleChange(t, 1, "candidate", "delta", "set: {depends_on: [absent]}\n", "")}
	_, graphError := w.Eligible(bad, "candidate")
	if graphError == nil {
		t.Fatal("invalid graph unexpectedly accepted")
	}
	for _, root := range []string{"/missing/source", "", "../not-a-repository", filepath.Join(t.TempDir(), "absent")} {
		s.Root = root
		held.Root = root
		eligibleAssert(t, held, "candidate", false)
		bad.Root = root
		if _, err := w.Eligible(s, "../bad"); err == nil || err.Error() != selectorError.Error() {
			t.Errorf("Root changed selector error: %v/%v", err, selectorError)
		}
		if _, err := w.Eligible(bad, "candidate"); err == nil || err.Error() != graphError.Error() {
			t.Errorf("Root changed graph error: %v/%v", err, graphError)
		}
		before := policyJSON(t, s)
		for i := 0; i < 2; i++ {
			eligibleAssert(t, s, "candidate", true)
			eligibleError(t, s, "../bad", "")
		}
		if before != policyJSON(t, s) {
			t.Error("ambient source changed")
		}
	}
	s.Root = "/missing/eligibility-root"
	if policyJSON(t, s) != original || !reflect.DeepEqual(backing, emptyBefore) {
		t.Fatal("nested or empty backing mutation")
	}
	// Real ambient state must not choose a different answer from the supplied data.
	current, err := os.Getwd()
	if err != nil || current != cwd || !reflect.DeepEqual(env, os.Environ()) {
		t.Fatal("initial in-process cwd/environment mutation")
	}
	directory, err := filepath.EvalSymlinks(t.TempDir())
	if err != nil {
		t.Fatal(err)
	}
	projectFile := filepath.Join(directory, "fireweed.yaml")
	if err := os.WriteFile(projectFile, []byte("malformed: ["), 0644); err != nil {
		t.Fatal(err)
	}
	if err := os.Chdir(directory); err != nil {
		t.Fatal(err)
	}
	defer os.Chdir(cwd)
	t.Setenv("FIREWEED_BIN", "must-not-run")
	t.Setenv("FIREWEED_ELIGIBILITY_DECOY", "changed")
	eligibleAssert(t, s, "candidate", true)
	after, err := os.Getwd()
	if err != nil || after != directory {
		t.Fatal("in-process cwd mutation")
	}
	// Environment snapshots are taken inside the same process as Eligible.
	expectedEnv := os.Environ()
	eligibleAssert(t, held, "candidate", false)
	if _, err := w.Eligible(s, "../bad"); err == nil || err.Error() != selectorError.Error() {
		t.Errorf("cwd/environment changed selector error %v", err)
	}
	bad.Root = "/missing/eligibility-root"
	if _, err := w.Eligible(bad, "candidate"); err == nil || err.Error() != graphError.Error() {
		t.Errorf("cwd/environment changed graph error %v", err)
	}
	eligibleAssert(t, s, "candidate", true)
	if !reflect.DeepEqual(expectedEnv, os.Environ()) {
		t.Fatal("in-process environment mutation")
	}
	raw, err := os.ReadFile(projectFile)
	if err != nil || string(raw) != "malformed: [" {
		t.Fatal("ambient file mutation")
	}
}
func Test_FLOW024_FLOW026EligibilityFixtureControls(t *testing.T) {
	s := eligibleSource(t)
	s.Changes = []c.Change{eligibleChange(t, 1, "blocker", "alpha", "blocked: waiting\n", ""), eligibleChange(t, 2, "candidate", "delta", "set: {depends_on: []}\n", "## Added\n\n### new:future [review]\nThe projected target MUST retain this item.\n")}
	_, m, err := eligibleGraph(t, s, "candidate")
	if err != nil || m.Name != "delta" || len(m.Specs) != 2 {
		t.Fatalf("public prefix control %v %+v", err, m)
	}
}
