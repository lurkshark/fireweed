package evals_test

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"regexp"
	"strings"
	"testing"
	"time"

	"github.com/lurkshark/fireweed/packages/assets"
	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/gitstore"
	"github.com/lurkshark/fireweed/packages/project"
	"github.com/lurkshark/fireweed/packages/workflow"
)

const oldClause = "Alpha MUST return the old value."
const newClause = "Alpha MUST return the revised value."

func write(t *testing.T, root, name, content string) {
	t.Helper()
	path := filepath.Join(root, filepath.FromSlash(name))
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}
}

func git(t *testing.T, root string, args ...string) string {
	t.Helper()
	cmd := exec.Command("git", args...)
	cmd.Dir = root
	cmd.Env = append(os.Environ(), "GIT_CONFIG_NOSYSTEM=1", "GIT_AUTHOR_NAME=Fixture", "GIT_AUTHOR_EMAIL=fixture@example.invalid", "GIT_COMMITTER_NAME=Fixture", "GIT_COMMITTER_EMAIL=fixture@example.invalid")
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("git %v: %v: %s", args, err, out)
	}
	return strings.TrimSpace(string(out))
}

func head(t *testing.T, root string) string { t.Helper(); return git(t, root, "rev-parse", "HEAD") }
func commit(t *testing.T, root, message string) {
	t.Helper()
	git(t, root, "add", "-A")
	git(t, root, "commit", "-m", message)
}

func fixture(t *testing.T, configExtra, changeFront, added, modified, removed string) string {
	t.Helper()
	root := t.TempDir()
	var err error
	root, err = filepath.EvalSymlinks(root)
	if err != nil {
		t.Fatal(err)
	}
	write(t, root, "SYSTEM.md", "# Fixture system\nLocal lifecycle fixture.\n")
	write(t, root, ".gitignore", ".fireweed/runs/\n")
	write(t, root, "fireweed.yaml", `fireweed: ">=0.1, <0.2"
version: 1
base: main
modules: ["packages/*"]
module_defaults:
  interface: [api.go]
  evals: [evals/**]
  disposable: [internal/**]
hooks:
  eval: [sh, .fireweed/eval.sh, "{module}", "{report}"]
  boundaries: [sh, .fireweed/boundaries.sh]
`+configExtra)
	write(t, root, ".fireweed/eval.sh", `#!/bin/sh
printf 'fixture-hook-output\n'
module="$1"
report="$2"
if [ "$module" = alpha ]; then
  if [ -f packages/alpha/internal/ready ]; then result=''; exitcode=0
  else result='<failure message="expected before implementation"/>'; exitcode=1; fi
  if [ -f .fireweed/unrelated ]; then other='<failure message="known unrelated"/>'; exitcode=1
  else other=''; fi
  printf '<testsuite><testcase name="ALP-001" classname="fixture">%s</testcase><testcase name="ALP-002" classname="fixture">%s</testcase></testsuite>' "$result" "$other" > "$report"
  exit "$exitcode"
fi
if [ "$module" = beta ] && [ -f .fireweed/beta-fail ]; then result='<failure/>'; exitcode=1
elif [ "$module" = gamma ] && [ -f .fireweed/gamma-fail ]; then result='<failure/>'; exitcode=1
else result=''; exitcode=0; fi
case "$module" in beta) id=BET-001;; gamma) id=GAM-001;; *) exit 2;; esac
printf '<testsuite><testcase name="%s" classname="fixture">%s</testcase></testsuite>' "$id" "$result" > "$report"
exit "$exitcode"
`)
	write(t, root, ".fireweed/boundaries.sh", `#!/bin/sh
if [ -f .fireweed/boundary-fail ] && [ -f packages/alpha/internal/ready ]; then printf 'alpha: forbidden dependency\n'; exit 1; fi
`)
	write(t, root, "packages/alpha/SPEC.md", "---\nname: alpha\nid_prefix: ALP\ndepends_on: []\n---\n# Alpha\n\n## Specs\n\n### ALP-001\n"+oldClause+"\n\n### ALP-002\nAlpha MUST retain the second behavior.\n\n## Decisions\nRetain old policy.\n")
	write(t, root, "packages/alpha/api.go", "package alpha\n")
	write(t, root, "packages/beta/SPEC.md", "---\nname: beta\nid_prefix: BET\ndepends_on: [alpha]\n---\n# Beta\n\n## Specs\n\n### BET-001\nBeta MUST follow alpha.\n")
	write(t, root, "packages/gamma/SPEC.md", "---\nname: gamma\nid_prefix: GAM\ndepends_on: [beta]\n---\n# Gamma\n\n## Specs\n\n### GAM-001\nGamma MUST follow beta.\n")
	if modified == "" {
		modified = "### ALP-001\n" + newClause + "\n"
	}
	base := contracts.SpecHash(contracts.Item{ID: "ALP-001", Body: oldClause + "\n"})
	if removed != "" {
		base += "\n  ALP-002: " + contracts.SpecHash(contracts.Item{ID: "ALP-002", Body: "Alpha MUST retain the second behavior.\n"})
	}
	write(t, root, "changes/0001-manual.md", "---\nmodule: alpha\nbases:\n  ALP-001: "+base+"\n"+changeFront+"---\n\n## Motivation\nUsers need a revised alpha value.\n\n## Interface\nAlpha public value is revised.\n\n## Added\n"+added+"\n## Modified\n"+modified+"\n## Removed\n"+removed+"\n## Decisions\nUse the revised policy.\n\n## Removed decisions\nRetain old policy.\n")
	git(t, root, "init")
	git(t, root, "checkout", "-b", "main")
	commit(t, root, "approved fixture")
	for _, d := range project.Check(root) {
		if !d.Warning {
			t.Fatalf("invalid lifecycle fixture: %#v", d)
		}
	}
	return root
}

func start(t *testing.T, root string) workflow.State {
	t.Helper()
	s, err := workflow.Start(root, "manual")
	if err != nil {
		t.Fatalf("Start: %v, diagnostics: %#v", err, s.Diagnostics)
	}
	return s
}

func request(root string) workflow.Request {
	return workflow.Request{Root: root, Change: "manual", Bin: "/tmp/fireweed-fixture", Timeout: 2 * time.Second}
}
func verify(root string, quick bool) (workflow.Result, error) {
	r := request(root)
	r.Quick = quick
	return workflow.Verify(context.Background(), r)
}
func done(root string) (workflow.Result, error) {
	return workflow.Done(context.Background(), request(root))
}

func containsCode(ds []project.Diagnostic, code string) bool {
	for _, d := range ds {
		if d.Code == code {
			return true
		}
	}
	return false
}

func Test_FLOW_001_FLOW_006_FLOW_007_FLOW_008_FLOW_012_FLOW_013_FLOW_014_FLOW_017_FLOW_018_FLOW_019_FLOW_053_StartAssignsLateIDsAndRetainsSnapshotWithoutCommit(t *testing.T) {
	root := fixture(t, "", "", "### new:first\nFirst MUST work.\n\n### new:second\nSecond MUST use new:first.\n", "", "")
	before := head(t, root)
	s := start(t, root+"/.")
	if s.Phase != "evals" || s.Run == "" || s.ClaimOID == "" || s.PhaseBase != before || head(t, root) != before {
		t.Errorf("start state = %#v; head before %s after %s", s, before, head(t, root))
	}
	b, _ := os.ReadFile(filepath.Join(root, "changes/0001-manual.md"))
	if strings.Contains(string(b), "new:first") || strings.Contains(string(b), "new:second") || !strings.Contains(string(b), "### ALP-003") || !strings.Contains(string(b), "### ALP-004") || !strings.Contains(string(b), "use ALP-003") {
		t.Errorf("late IDs not assigned throughout change: %s", b)
	}
	if len(s.Change.Added) != 2 || s.Change.Added[0].ID != "ALP-003" || s.Change.Added[1].ID != "ALP-004" {
		t.Errorf("returned Added IDs = %#v", s.Change.Added)
	}
	alphaProjected := ""
	for _, module := range s.Snapshot.Modules {
		if module.Name == "alpha" {
			alphaProjected = module.Document.Body
		}
	}
	if !strings.Contains(alphaProjected, newClause) || !strings.Contains(alphaProjected, "First MUST work.") || !strings.Contains(alphaProjected, "### ALP-003") || !strings.Contains(alphaProjected, "### ALP-004") {
		t.Errorf("returned snapshot does not project the approved change: %q", alphaProjected)
	}
	edited := strings.Replace(string(b), "First MUST work.", "First MUST change after approval.", 1)
	if edited == string(b) {
		t.Fatal("fixture change has no approved Added body to edit")
	}
	write(t, root, "changes/0001-manual.md", edited)
	current, err := workflow.Current(root, "")
	if err != nil || current.Run != s.Run || current.Change.Slug != "manual" || current.Snapshot.Root != root {
		t.Errorf("Current = %#v, %v", current, err)
	}
	currentAlpha := ""
	for _, module := range current.Snapshot.Modules {
		if module.Name == "alpha" {
			currentAlpha = module.Document.Body
		}
	}
	if currentAlpha != alphaProjected {
		t.Errorf("Current snapshot changed after editing the held change: got %q, approved %q", currentAlpha, alphaProjected)
	}
	if s.Snapshot.Root != root || len(s.Snapshot.Modules) != 3 {
		t.Errorf("returned projected snapshot = %#v", s.Snapshot)
	}
	second := fixture(t, "", "", "", "", "")
	if another := start(t, second); another.Run == "" || another.Run == s.Run {
		t.Errorf("separate starts returned runs %q and %q", s.Run, another.Run)
	}
}

func Test_FLOW_010_FLOW_011_RequiresCleanEligibleQueueHead(t *testing.T) {
	clean := fixture(t, "", "", "", "", "")
	if _, err := workflow.Start(clean, "manual"); err != nil {
		t.Errorf("clean eligible queue head rejected: %v", err)
	}
	for _, dirty := range []string{"tracked", "untracked"} {
		t.Run(dirty, func(t *testing.T) {
			root := fixture(t, "", "", "", "", "")
			if dirty == "tracked" {
				write(t, root, "SYSTEM.md", "edited\n")
			} else {
				write(t, root, "stray.txt", "untracked\n")
			}
			if _, err := workflow.Start(root, "manual"); err == nil {
				t.Error("dirty repository claimed")
			}
		})
	}
	root := fixture(t, "", "", "", "", "")
	write(t, root, "changes/0000-first.md", "---\nmodule: alpha\n---\n## Motivation\nFirst request.\n\n## Interface\n\n## Added\n\n## Modified\n\n## Removed\n\n## Decisions\n")
	commit(t, root, "earlier queued change")
	if _, err := workflow.Start(root, "manual"); err == nil {
		t.Error("later queued change claimed")
	}
}

func Test_FLOW_001_FLOW_016_FLOW_017_LinkedWorktreeExcludesCompetingOwner(t *testing.T) {
	root := fixture(t, "", "", "", "", "")
	other := filepath.Join(t.TempDir(), "linked")
	git(t, root, "worktree", "add", "-b", "linked", other, "HEAD")
	type attempt struct {
		root  string
		state workflow.State
		err   error
	}
	gate := make(chan struct{})
	results := make(chan attempt, 2)
	for _, dir := range []string{root, other} {
		go func(dir string) {
			<-gate
			s, err := workflow.Start(dir, "manual")
			results <- attempt{dir, s, err}
		}(dir)
	}
	close(gate)
	timer := time.NewTimer(10 * time.Second)
	defer timer.Stop()
	var attempts [2]attempt
	for i := range attempts {
		select {
		case attempts[i] = <-results:
		case <-timer.C:
			t.Fatal("simultaneous claim attempts did not finish")
		}
	}
	winners := 0
	for _, result := range attempts {
		if result.err == nil {
			winners++
			if got, err := workflow.Current(result.root, "manual"); err != nil || got.ClaimOID != result.state.ClaimOID {
				t.Errorf("winner lost claim: %#v, %v", got, err)
			}
		} else if _, err := workflow.Current(result.root, "manual"); err == nil {
			t.Errorf("losing worktree %s resolved foreign claim", result.root)
		}
	}
	if winners != 1 {
		t.Errorf("simultaneous claims had %d winners: %#v", winners, attempts)
	}
}

func Test_FLOW_009_FLOW_015_ApprovalWarningAndSharedUnapprovedLimit(t *testing.T) {
	approved := fixture(t, "", "", "", "", "")
	if s := start(t, approved); containsCode(s.Diagnostics, "P04") {
		t.Error("local-base approval warned")
	}
	shared := fixture(t, "", "", "", "", "")
	sharedBase := head(t, shared)
	git(t, shared, "checkout", "-b", "feature")
	git(t, shared, "checkout", "main")
	git(t, shared, "rm", "changes/0001-manual.md")
	commit(t, shared, "local base lacks pending change")
	git(t, shared, "update-ref", "refs/remotes/origin/main", sharedBase)
	git(t, shared, "checkout", "feature")
	if s := start(t, shared); containsCode(s.Diagnostics, "P04") {
		t.Error("shared-base approval warned")
	}
	for _, limit := range []int{4, 5} {
		t.Run(fmt.Sprint(limit), func(t *testing.T) {
			root := fixture(t, "", "", "", "", "")
			git(t, root, "checkout", "-b", "feature")
			git(t, root, "checkout", "main")
			git(t, root, "rm", "changes/0001-manual.md")
			commit(t, root, "base without pending change")
			git(t, root, "update-ref", "refs/remotes/origin/main", head(t, root))
			git(t, root, "checkout", "feature")
			for i := 0; i < limit; i++ {
				write(t, root, fmt.Sprintf("changes/archive/%04d-prior.md", i+1), fmt.Sprintf("---\nmodule: alpha\nphase: applied\n---\n## Motivation\nPrior %d.\n", i))
				commit(t, root, "archived prior")
			}
			before := head(t, root)
			s, err := workflow.Start(root, "manual")
			if limit == 5 {
				if err == nil || !containsCode(s.Diagnostics, "P07") || head(t, root) != before {
					t.Errorf("shared max bound = %#v, %v", s, err)
				}
			}
			if limit == 4 {
				if err != nil || !containsCode(s.Diagnostics, "P04") {
					t.Errorf("unapproved below bound = %#v, %v", s, err)
				}
			}
		})
	}
}

func hookMarker(t *testing.T, root string) string {
	t.Helper()
	marker := filepath.Join(t.TempDir(), "hook-ran")
	p := filepath.Join(root, ".fireweed/eval.sh")
	b, err := os.ReadFile(p)
	if err != nil {
		t.Fatal(err)
	}
	write(t, root, ".fireweed/eval.sh", "printf run > "+fmt.Sprintf("%q", marker)+"\n"+string(b))
	commit(t, root, "track hook observer")
	return marker
}

func Test_FLOW_002_FLOW_020_FLOW_021_RejectsAuthorMetadataAndZonesBeforeHooks(t *testing.T) {
	for _, edit := range []string{"author", "eval-zone"} {
		t.Run(edit, func(t *testing.T) {
			root := fixture(t, "", "", "", "", "")
			write(t, root, "packages/alpha/internal/existing.go", "package internal\n")
			commit(t, root, "existing implementation")
			marker := hookMarker(t, root)
			start(t, root)
			if edit == "author" {
				p := filepath.Join(root, "changes/0001-manual.md")
				b, _ := os.ReadFile(p)
				write(t, root, "changes/0001-manual.md", strings.Replace(string(b), "Users need a revised alpha value.", "Unapproved changed purpose.", 1))
			} else {
				write(t, root, "packages/alpha/internal/existing.go", "package internal // changed in evals\n")
			}
			if r, err := verify(root, true); err == nil {
				t.Errorf("%s accepted: %#v", edit, r)
			}
			if _, err := os.Stat(marker); !os.IsNotExist(err) {
				t.Errorf("hook ran before %s rejection: %v", edit, err)
			}
			if edit == "author" {
				p := filepath.Join(root, "changes/0001-manual.md")
				b, _ := os.ReadFile(p)
				write(t, root, "changes/0001-manual.md", strings.Replace(strings.Replace(string(b), "Unapproved changed purpose.", "Users need a revised alpha value.", 1), "phase: evals", "phase: impl", 1))
				if r, err := verify(root, true); err == nil {
					t.Errorf("approved phase state alteration accepted: %#v", r)
				}
				if _, err := os.Stat(marker); !os.IsNotExist(err) {
					t.Errorf("hook ran before phase-state rejection: %v", err)
				}
			}
		})
	}
}

func Test_FLOW_022_FLOW_023_FLOW_031_EvalsRequireCoverageAndPerIDFailFirst(t *testing.T) {
	for _, scenario := range []string{"missing-coverage", "passing-target-unrelated-failure", "already-satisfied", "fixes"} {
		t.Run(scenario, func(t *testing.T) {
			front := ""
			if scenario == "already-satisfied" {
				front = "already_satisfied: [ALP-001]\n"
			}
			if scenario == "fixes" {
				front = "fixes: [ALP-002]\n"
			}
			root := fixture(t, "", front, "", "", "")
			if scenario == "missing-coverage" {
				p := filepath.Join(root, ".fireweed/eval.sh")
				b, _ := os.ReadFile(p)
				write(t, root, ".fireweed/eval.sh", strings.ReplaceAll(string(b), `name="ALP-001"`, `name="ALP-002"`))
				commit(t, root, "hook omits target")
			}
			if scenario == "already-satisfied" {
				p := filepath.Join(root, ".fireweed/eval.sh")
				b, _ := os.ReadFile(p)
				script := strings.Replace(string(b), `  exit "$exitcode"`, `  if [ -f packages/alpha/evals/no-coverage ]; then sed 's/name="ALP-002"/name="ALP-001"/g' "$report" > "$report.tmp"; mv "$report.tmp" "$report"; fi
  exit "$exitcode"`, 1)
				if script == string(b) {
					t.Fatal("fixture coverage switch was not installed")
				}
				write(t, root, ".fireweed/eval.sh", script)
				commit(t, root, "hook can omit unrelated current spec")
			}
			if scenario == "passing-target-unrelated-failure" || scenario == "already-satisfied" {
				write(t, root, ".fireweed/unrelated", "yes\n")
				commit(t, root, "known unrelated failing case")
			}
			start(t, root)
			if scenario == "passing-target-unrelated-failure" || scenario == "already-satisfied" {
				write(t, root, "packages/alpha/internal/ready", "ready\n")
			}
			before := head(t, root)
			r, err := verify(root, true)
			wantPass := scenario == "already-satisfied"
			if (err == nil) != wantPass {
				t.Errorf("Verify = %#v, %v; want pass %v", r, err, wantPass)
			}
			if head(t, root) != before {
				t.Error("Verify committed")
			}
			if wantPass && (len(r.Reports["alpha"].Cases) == 0 || r.State.Snapshot.Root != root) {
				t.Errorf("reports or held snapshot absent: %#v", r)
			}
			if scenario == "already-satisfied" {
				write(t, root, "packages/alpha/evals/no-coverage", "omit ALP-002\n")
				missing, err := verify(root, true)
				if len(missing.Reports["alpha"].Cases) == 0 {
					t.Errorf("target module evals did not run: %#v, %v", missing, err)
				}
				if err == nil {
					t.Errorf("evals accepted report without unrelated current spec ALP-002: %#v", missing)
				}
			}
		})
	}
}

func Test_FLOW_023_AddedIDRequiresMappedFailingCase(t *testing.T) {
	root := fixture(t, "", "set:\n  interface: [api.go, extra.go]\n", "### new:third\nThird MUST return a value.\n", "", "")
	p := filepath.Join(root, ".fireweed/eval.sh")
	b, _ := os.ReadFile(p)
	script := strings.Replace(string(b), `</testsuite>' "$result" "$other"`, `<testcase name="ALP-003" classname="fixture">%s</testcase></testsuite>' "$result" "$other" "$result"`, 1)
	if script == string(b) {
		t.Fatal("fixture report extension was not installed")
	}
	write(t, root, ".fireweed/eval.sh", script)
	commit(t, root, "report projected added ID")
	s := start(t, root)
	if len(s.Change.Added) != 1 || s.Change.Added[0].ID != "ALP-003" {
		t.Fatalf("assigned Added ID = %#v", s.Change.Added)
	}
	write(t, root, "packages/alpha/extra.go", "package alpha\n")
	r, err := verify(root, true)
	if err != nil {
		t.Fatalf("mapped failing Added case rejected: %#v, %v", r, err)
	}
	if len(r.Reports["alpha"].Cases) != 3 {
		t.Errorf("added ID was not reported: %#v", r.Reports["alpha"])
	}
}

func Test_FLOW_002_FLOW_003_FLOW_022_FLOW_025_FLOW_026_FLOW_027_FLOW_029_FLOW_030_FLOW_031_ImplFreezesEvalsAndRunsDependentsAndBoundaries(t *testing.T) {
	for _, scenario := range []string{"frozen-eval", "direct", "all", "none", "quick-boundary"} {
		t.Run(scenario, func(t *testing.T) {
			config := "audit:\n  enabled: false\n"
			if scenario == "all" {
				config += "verify_dependents: all\n"
			}
			if scenario == "none" {
				config += "verify_dependents: none\n"
			}
			root := fixture(t, config, "", "", "", "")
			if scenario == "frozen-eval" {
				p := filepath.Join(root, ".fireweed/eval.sh")
				b, _ := os.ReadFile(p)
				script := strings.Replace(string(b), `  exit "$exitcode"`, `  if [ -f packages/alpha/internal/no-coverage ]; then sed 's/name="ALP-001"/name="ALP-002"/g' "$report" > "$report.tmp"; mv "$report.tmp" "$report"; fi
  exit "$exitcode"`, 1)
				write(t, root, ".fireweed/eval.sh", script)
				commit(t, root, "coverage switch fixture")
			}
			switch scenario {
			case "direct":
				write(t, root, ".fireweed/beta-fail", "yes\n")
			case "none":
				write(t, root, ".fireweed/beta-fail", "yes\n")
			case "all":
				write(t, root, ".fireweed/gamma-fail", "yes\n")
			case "quick-boundary":
				write(t, root, ".fireweed/boundary-fail", "yes\n")
			}
			if scenario != "frozen-eval" {
				commit(t, root, "configured failing hook")
			}
			start(t, root)
			if _, err := done(root); err != nil {
				t.Fatalf("evals Done: %v", err)
			}
			current, err := workflow.Current(root, "manual")
			if err != nil || current.Phase != "impl" {
				t.Fatalf("impl state = %#v, %v", current, err)
			}
			write(t, root, "packages/alpha/internal/ready", "ready\n")
			if scenario == "frozen-eval" {
				write(t, root, "packages/alpha/evals/new_test.go", "package evals\n")
			}
			before := head(t, root)
			quick, quickErr := verify(root, true)
			full, fullErr := verify(root, false)
			if head(t, root) != before {
				t.Error("Verify committed")
			}
			if scenario == "frozen-eval" {
				if quickErr == nil || fullErr == nil {
					t.Errorf("frozen eval edit accepted: %#v %v; %#v %v", quick, quickErr, full, fullErr)
				}
				if err := os.Remove(filepath.Join(root, "packages/alpha/evals/new_test.go")); err != nil {
					t.Fatal(err)
				}
				write(t, root, "packages/alpha/internal/no-coverage", "yes\n")
				if r, err := verify(root, true); err == nil {
					t.Errorf("impl accepted missing current coverage: %#v", r)
				}
			} else {
				if quickErr != nil {
					t.Errorf("quick rejected own passing evals: %v", quickErr)
				}
				wantFullPass := scenario == "none"
				if (fullErr == nil) != wantFullPass {
					t.Errorf("full Verify = %#v, %v; want pass %v", full, fullErr, wantFullPass)
				}
				if scenario == "direct" && len(full.Reports["beta"].Cases) == 0 {
					t.Error("direct dependent report absent")
				}
				if scenario == "direct" && len(full.Reports["gamma"].Cases) != 0 {
					t.Error("default direct verification ran a transitive dependent")
				}
				if scenario == "all" && len(full.Reports["gamma"].Cases) == 0 {
					t.Error("transitive dependent report absent")
				}
				if scenario == "none" && (len(full.Reports["beta"].Cases) != 0 || len(full.Reports["gamma"].Cases) != 0) {
					t.Error("dependent verification ran with verify_dependents: none")
				}
				if scenario == "direct" && len(full.Reports["beta"].Cases) > 0 && full.Reports["beta"].Cases[0].Result != "fail" {
					t.Errorf("selected direct dependent did not report its failing testcase: %#v", full.Reports["beta"])
				}
				if scenario == "all" && len(full.Reports["gamma"].Cases) > 0 && full.Reports["gamma"].Cases[0].Result != "fail" {
					t.Errorf("selected transitive dependent did not report its failing testcase: %#v", full.Reports["gamma"])
				}
				if scenario == "direct" || scenario == "all" || scenario == "quick-boundary" {
					r, err := done(root)
					if err == nil || head(t, root) != before {
						t.Errorf("Done accepted failed full verification or committed: %#v, %v", r, err)
					}
					current, currentErr := workflow.Current(root, "manual")
					if currentErr != nil || current.Phase != "impl" || current.PhaseBase != before {
						t.Errorf("failed full verification advanced phase: %#v, %v", current, currentErr)
					}
				}
			}
		})
	}
}

func audit(t *testing.T, root, body string) {
	t.Helper()
	p := filepath.Join(root, "changes/0001-manual.md")
	b, err := os.ReadFile(p)
	if err != nil {
		t.Fatal(err)
	}
	text := contracts.ReplaceSection(string(b), "Audit", body)
	write(t, root, "changes/0001-manual.md", text)
}

func toAudit(t *testing.T, root string) workflow.State {
	t.Helper()
	start(t, root)
	r, err := done(root)
	if err != nil || r.State.Phase != "audit" {
		t.Fatalf("evals Done = %#v, %v", r, err)
	}
	return r.State
}

func Test_FLOW_032_FLOW_033_FLOW_046_EvalsDoneUsesDefaultAuditOrDisabledImpl(t *testing.T) {
	for _, disabled := range []bool{false, true} {
		t.Run(fmt.Sprint(disabled), func(t *testing.T) {
			config := ""
			want := "audit"
			if disabled {
				config = "audit:\n  enabled: false\n"
				want = "impl"
			}
			root := fixture(t, config, "", "", "", "")
			start(t, root)
			before := head(t, root)
			r, err := done(root)
			if err != nil || r.State.Phase != want || head(t, root) == before || r.State.PhaseBase != head(t, root) {
				t.Errorf("Done = %#v, %v; want %s with commit", r, err, want)
			}
			current, err := workflow.Current(root, "manual")
			if err != nil || current.Phase != want || current.PhaseBase != head(t, root) {
				t.Errorf("phase base not advanced: %#v, %v", current, err)
			}
		})
	}
}

func Test_FLOW_034_FLOW_036_FLOW_041_FLOW_046_AuditAcceptsOnlyExactClauseFindingsAndCheckedNoFindings(t *testing.T) {
	root := fixture(t, "", "", "", "", "")
	toAudit(t, root)
	audit(t, root, "- ALP-999: \"Missing MUST clause.\" — invalid ID\n- ALP-001: \"Alpha MUST do something else.\" — invalid quote\nChecked IDs: ALP-001\nNo findings\n")
	r, err := done(root)
	if err != nil || r.State.Phase != "impl" || r.State.PhaseBase != head(t, root) {
		t.Errorf("invalid findings were retained: %#v, %v", r, err)
	}
	b, _ := os.ReadFile(filepath.Join(root, "changes/0001-manual.md"))
	if strings.Contains(string(b), "ALP-999") || strings.Contains(string(b), "something else") || !strings.Contains(string(b), "No findings") {
		t.Errorf("audit normalization = %s", b)
	}
	missing := fixture(t, "", "", "", "", "")
	toAudit(t, missing)
	audit(t, missing, "No findings\n")
	before := head(t, missing)
	if _, err := done(missing); err == nil || head(t, missing) != before {
		t.Error("No findings without checked IDs accepted or committed")
	}
}

func Test_FLOW_035_FLOW_037_FLOW_038_FLOW_039_FLOW_042_FLOW_043_FLOW_046_AuditFindingsConvergeAndPersistTwoRoundBound(t *testing.T) {
	root := fixture(t, "", "", "### new:extra [review]\nExtra MUST be visible.\n", "", "")
	toAudit(t, root)
	finding := "- ALP-001: \"" + newClause + "\" — missing assertion\n"
	for round := 1; round <= 2; round++ {
		body := finding
		if round == 2 {
			body += "- ALP-003: \"Extra MUST be visible.\" — new finding on re-audit\n"
		}
		audit(t, root, body)
		r, err := done(root)
		if err != nil || r.State.Phase != "evals" || r.State.PhaseBase != head(t, root) {
			t.Fatalf("finding round %d = %#v, %v", round, r, err)
		}
		current, err := workflow.Current(root, "manual")
		if err != nil || current.Change.Document.Frontmatter["audit_rounds"] != round {
			t.Errorf("persistent rounds %d = %#v, %v", round, current, err)
		}
		if round == 2 {
			b, _ := os.ReadFile(filepath.Join(root, "changes/0001-manual.md"))
			if strings.Contains(string(b), "new finding on re-audit") {
				t.Error("new re-audit ID retained")
			}
		}
		if _, err := done(root); err != nil {
			t.Fatalf("revised evals Done round %d: %v", round, err)
		}
	}
	audit(t, root, finding)
	beforeBlock := head(t, root)
	_, _ = done(root)
	b, _ := os.ReadFile(filepath.Join(root, "changes/0001-manual.md"))
	questions, hasQuestions := contracts.Section(string(b), "Questions")
	if !strings.Contains(string(b), "blocked:") || !hasQuestions || !strings.Contains(questions, newClause) {
		t.Errorf("third finding did not block with question: %s", b)
	}
	if head(t, root) == beforeBlock || !strings.Contains(string(b), "audit_rounds: 2") {
		t.Error("two-round bound was not committed with persistent count")
	}
	if _, err := workflow.Current(root, "manual"); err == nil {
		t.Error("blocked claim remained held")
	}
	configured := fixture(t, "audit:\n  max_rounds: 1\n", "", "", "", "")
	toAudit(t, configured)
	audit(t, configured, finding)
	first, err := done(configured)
	if err != nil || first.State.Phase != "evals" || first.State.PhaseBase != head(t, configured) {
		t.Fatalf("configured first finding = %#v, %v", first, err)
	}
	firstCurrent, err := workflow.Current(configured, "manual")
	if err != nil || firstCurrent.Change.Document.Frontmatter["audit_rounds"] != 1 {
		t.Errorf("configured round count = %#v, %v", firstCurrent, err)
	}
	if _, err := done(configured); err != nil {
		t.Fatalf("configured revised evals Done: %v", err)
	}
	audit(t, configured, finding)
	_, _ = done(configured)
	change, err := os.ReadFile(filepath.Join(configured, "changes/0001-manual.md"))
	configuredQuestions, configuredHasQuestions := contracts.Section(string(change), "Questions")
	if err != nil || !strings.Contains(string(change), "blocked:") || !strings.Contains(string(change), "audit_rounds: 1") || !configuredHasQuestions || !strings.Contains(configuredQuestions, newClause) {
		t.Errorf("configured max_rounds did not block after one return: %s, %v", change, err)
	}
	if _, err := workflow.Current(configured, "manual"); err == nil {
		t.Error("configured bound retained claim")
	}
}

func Test_FLOW_040_FLOW_042_FLOW_043_UntestableFindingBlocksWithQuestion(t *testing.T) {
	root := fixture(t, "", "", "", "", "")
	toAudit(t, root)
	audit(t, root, "- [untestable] ALP-001: \""+newClause+"\" — cannot observe the value\n")
	_, _ = done(root)
	b, _ := os.ReadFile(filepath.Join(root, "changes/0001-manual.md"))
	questions, hasQuestions := contracts.Section(string(b), "Questions")
	if !strings.Contains(string(b), "blocked:") || !hasQuestions || !strings.Contains(questions, newClause) {
		t.Errorf("untestable finding did not block: %s", b)
	}
	if _, err := workflow.Current(root, "manual"); err == nil {
		t.Error("blocked claim remained held")
	}
}

func Test_FLOW_004_FLOW_028_FLOW_044_FLOW_045_FLOW_046_FLOW_047_FLOW_050_ImplDoneAppliesSpecsSetAndDecisionsAndArchives(t *testing.T) {
	root := fixture(t, "audit:\n  enabled: false\nverify_dependents: none\n", "set:\n  interface: [api.go, extra.go]\n  hooks:\n    eval: [sh, .fireweed/eval.sh, \"{module}\", \"{report}\"]\n", "### new:extra [review]\nExtra MUST be visible.\n", "", "### ALP-002\n")
	p := filepath.Join(root, ".fireweed/eval.sh")
	b, _ := os.ReadFile(p)
	script := strings.Replace(string(b), `<testcase name="ALP-002" classname="fixture">%s</testcase>`, "", 1)
	script = strings.Replace(script, `"$result" "$other" > "$report"`, `"$result" > "$report"`, 1)
	write(t, root, ".fireweed/eval.sh", script)
	write(t, root, ".fireweed/global-eval.sh", "#!/bin/sh\nexit 1\n")
	configPath := filepath.Join(root, "fireweed.yaml")
	config, err := os.ReadFile(configPath)
	if err != nil {
		t.Fatal(err)
	}
	write(t, root, "fireweed.yaml", strings.Replace(string(config), "eval: [sh, .fireweed/eval.sh,", "eval: [sh, .fireweed/global-eval.sh,", 1))
	commit(t, root, "report only projected spec IDs")
	start(t, root)
	evalsHead := head(t, root)
	r, err := done(root)
	if err != nil || r.State.Phase != "impl" || head(t, root) == evalsHead || r.State.PhaseBase != head(t, root) {
		t.Fatalf("evals transition = %#v, %v", r, err)
	}
	implBase := head(t, root)
	otherClaim := []byte(`{"worktree":"independent","change":"other"}`)
	otherOID, err := gitstore.WriteBlob(root, otherClaim)
	if err != nil {
		t.Fatal(err)
	}
	otherRef := "refs/fireweed/claims/other"
	if err := gitstore.CASRef(root, otherRef, "", otherOID); err != nil {
		t.Fatal(err)
	}
	write(t, root, "packages/alpha/internal/ready", "ready\n")
	r, err = done(root)
	if err != nil || r.State.Phase != "applied" || head(t, root) == implBase || r.State.PhaseBase != head(t, root) {
		t.Fatalf("impl transition = %#v, %v", r, err)
	}
	if _, err := os.Stat(filepath.Join(root, "changes/0001-manual.md")); !os.IsNotExist(err) {
		t.Errorf("pending change remains: %v", err)
	}
	archive, err := os.ReadFile(filepath.Join(root, "changes/archive/0001-manual.md"))
	if err != nil {
		t.Fatal(err)
	}
	spec, err := os.ReadFile(filepath.Join(root, "packages/alpha/SPEC.md"))
	if err != nil {
		t.Fatal(err)
	}
	appliedHead := head(t, root)
	archiveAtHead, err := gitstore.ReadAt(root, appliedHead, "changes/archive/0001-manual.md")
	if err != nil || string(archiveAtHead) != string(archive) {
		t.Errorf("applied HEAD lacks archived change: %q, %v", archiveAtHead, err)
	}
	specAtHead, err := gitstore.ReadAt(root, appliedHead, "packages/alpha/SPEC.md")
	if err != nil || string(specAtHead) != string(spec) {
		t.Errorf("applied HEAD lacks applied spec: %q, %v", specAtHead, err)
	}
	moduleDoc, err := contracts.ParseDocument(string(spec))
	if err != nil {
		t.Fatal(err)
	}
	wantInterface := []any{"api.go", "extra.go"}
	wantHooks := map[string]any{"eval": []any{"sh", ".fireweed/eval.sh", "{module}", "{report}"}}
	if !reflect.DeepEqual(moduleDoc.Frontmatter["interface"], wantInterface) || !reflect.DeepEqual(moduleDoc.Frontmatter["hooks"], wantHooks) {
		t.Errorf("set values missing from module metadata: %#v", moduleDoc.Frontmatter)
	}
	for _, want := range []string{"phase: applied", "## Provenance"} {
		if !strings.Contains(string(archive), want) {
			t.Errorf("archive misses %q", want)
		}
	}
	for _, want := range []string{newClause, "### ALP-003", "extra.go", "Use the revised policy."} {
		if !strings.Contains(string(spec), want) {
			t.Errorf("applied spec misses %q: %s", want, spec)
		}
	}
	if strings.Contains(string(spec), "### ALP-002") || strings.Contains(string(spec), "Retain old policy.") {
		t.Errorf("removed spec/decision remains: %s", spec)
	}
	if !strings.Contains(string(spec), "- manual: Use the revised policy.") {
		t.Errorf("decision attribution missing: %s", spec)
	}
	if _, err := workflow.Current(root, "manual"); err == nil {
		t.Error("applied claim remained held")
	}
	if got, exists, err := gitstore.Ref(root, otherRef); err != nil || !exists || got != otherOID {
		t.Errorf("application changed independent claim ref: %q, %v, %v", got, exists, err)
	}
	if got, err := gitstore.ReadBlob(root, otherOID); err != nil || string(got) != string(otherClaim) {
		t.Errorf("application changed independent claim blob: %q, %v", got, err)
	}
}

func Test_FLOW_003_FLOW_024_FLOW_048_FLOW_049_FailedDoneKeepsPhaseAndHead(t *testing.T) {
	root := fixture(t, "audit:\n  enabled: false\n", "", "", "", "")
	hookPath := filepath.Join(root, ".fireweed/eval.sh")
	hook, err := os.ReadFile(hookPath)
	if err != nil {
		t.Fatal(err)
	}
	inject := `  if [ -f packages/alpha/internal/ready ]; then printf 'invalid spec fixture\n' > packages/alpha/SPEC.md; fi
  exit "$exitcode"`
	patched := strings.Replace(string(hook), `  exit "$exitcode"`, inject, 1)
	if patched == string(hook) {
		t.Fatal("fixture hook injection was not installed")
	}
	write(t, root, ".fireweed/eval.sh", patched)
	commit(t, root, "tracked invalid-spec hook fixture")
	start(t, root)
	if _, err := done(root); err != nil {
		t.Fatalf("evals Done: %v", err)
	}
	base := head(t, root)
	own, ownErr := verify(root, false)
	if ownErr == nil {
		t.Errorf("impl Verify accepted failing current-module testcase: %#v", own)
	}
	if cases := own.Reports["alpha"].Cases; len(cases) == 0 || cases[0].Result != "fail" || len(cases[0].IDs) == 0 || cases[0].IDs[0] != "ALP-001" {
		t.Errorf("failing current-module testcase absent from Verify report: %#v", own.Reports["alpha"])
	}
	r, err := done(root)
	if err == nil || head(t, root) != base {
		t.Errorf("passing current evals required: %#v, %v", r, err)
	}
	current, currentErr := workflow.Current(root, "manual")
	if currentErr != nil || current.Phase != "impl" || current.PhaseBase != base {
		t.Errorf("failed Done advanced phase: %#v, %v", current, currentErr)
	}
	write(t, root, "packages/alpha/internal/ready", "ready\n")
	write(t, root, "packages/alpha/api.go", "package alpha // edited during impl\n")
	r, err = done(root)
	if err == nil || head(t, root) != base {
		t.Errorf("frozen interface edit accepted or committed: %#v, %v", r, err)
	}
	write(t, root, "packages/alpha/api.go", "package alpha\n")
	r, err = done(root)
	if err == nil || head(t, root) != base {
		t.Errorf("failed static validation applied or committed: %#v, %v", r, err)
	}
	invalidSpec, readErr := os.ReadFile(filepath.Join(root, "packages/alpha/SPEC.md"))
	if readErr != nil || string(invalidSpec) != "invalid spec fixture\n" {
		t.Errorf("successful eval hook did not inject invalid spec: %q, %v", invalidSpec, readErr)
	}
	invalid := false
	for _, d := range project.Check(root) {
		invalid = invalid || !d.Warning
	}
	if !invalid {
		t.Error("injected spec passed static validation")
	}
	if _, err := os.Stat(filepath.Join(root, "changes/archive/0001-manual.md")); !os.IsNotExist(err) {
		t.Errorf("failed static validation created archive: %v", err)
	}
	pending, readErr := os.ReadFile(filepath.Join(root, "changes/0001-manual.md"))
	if readErr != nil || !strings.Contains(string(pending), "phase: impl") || strings.Contains(string(pending), "phase: applied") {
		t.Errorf("failed static validation advanced pending phase: %q, %v", pending, readErr)
	}
}

func Test_FLOW_005_FLOW_051_FLOW_052_FLOW_053_ContextAndProvenanceAndRetainedReports(t *testing.T) {
	root := fixture(t, "audit:\n  enabled: false\n", "", "", "", "")
	write(t, root, "packages/alpha/internal/context-note.txt", "Full context content marker.\n")
	commit(t, root, "context file")
	state := start(t, root)
	if state.Snapshot.Root != root || len(state.Snapshot.Modules) != 3 {
		t.Errorf("held snapshot absent: %#v", state)
	}
	compact, err := workflow.Context(root, "manual", false)
	if err != nil {
		t.Fatal(err)
	}
	full, err := workflow.Context(root, "manual", true)
	if err != nil {
		t.Fatal(err)
	}
	for _, want := range []string{"Users need a revised alpha value.", "Alpha MUST retain the second behavior.", "# Fixture system", "# Write evals", "package alpha"} {
		if !strings.Contains(compact, want) {
			t.Errorf("compact context misses %q", want)
		}
	}
	if strings.Count(compact, newClause) < 2 {
		t.Errorf("compact context misses revised clause in change and projected spec")
	}
	if !strings.Contains(compact, "packages/alpha/internal/context-note.txt") || !regexp.MustCompile(`(?m)context-note\.txt[^\n]*\b29\b`).MatchString(compact) {
		t.Errorf("compact context omitted file path/size: %q", compact)
	}
	if !strings.Contains(full, "Full context content marker.") {
		t.Error("full context omitted relevant file contents")
	}
	if _, err := done(root); err != nil {
		t.Fatalf("evals Done: %v", err)
	}
	write(t, root, "packages/alpha/internal/ready", "ready\n")
	implResult, err := done(root)
	if err != nil {
		t.Fatalf("impl Done: %v", err)
	}
	b, err := os.ReadFile(filepath.Join(root, "changes/archive/0001-manual.md"))
	if err != nil {
		t.Fatal(err)
	}
	provenance, ok := contracts.Section(string(b), "Provenance")
	if !ok || !strings.Contains(provenance, "evals") || !strings.Contains(provenance, "impl") || !strings.Contains(provenance, "manual") || !strings.Contains(provenance, "0.1.0") || !strings.Contains(provenance, "202") {
		t.Errorf("provenance incomplete: %q", provenance)
	}
	for _, phase := range []string{"evals", "impl"} {
		line := ""
		for _, candidate := range strings.Split(provenance, "\n") {
			if strings.HasPrefix(strings.TrimSpace(candidate), "- "+phase+" ") {
				line = strings.TrimSpace(candidate)
				break
			}
		}
		playbook, err := assets.Playbook(phase)
		if err != nil {
			t.Fatal(err)
		}
		hash := sha256.Sum256([]byte(playbook))
		if line == "" || !strings.Contains(line, "playbook="+hex.EncodeToString(hash[:])) {
			t.Errorf("%s provenance line misses matching playbook hash: %q", phase, line)
		}
		if phase == "impl" {
			total, passing, failing := 0, 0, 0
			for _, report := range implResult.Reports {
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
			if total == 0 || passing == 0 || !strings.Contains(line, fmt.Sprintf("tests=%d failing=%d passing=%d", total, failing, passing)) {
				t.Errorf("impl provenance line has wrong result counts for reports %#v: %q", implResult.Reports, line)
			}
		}
	}
	runRoot := filepath.Join(root, ".fireweed/runs", state.Run)
	count := 0
	retained := ""
	_ = filepath.Walk(runRoot, func(path string, info os.FileInfo, err error) error {
		if err == nil && info != nil && !info.IsDir() {
			count++
			b, readErr := os.ReadFile(path)
			if readErr == nil {
				retained += string(b)
			}
		}
		return nil
	})
	if count == 0 || !strings.Contains(retained, "ALP-001") || !strings.Contains(retained, "fixture-hook-output") {
		t.Errorf("reports/output missing under %s (%d files)", runRoot, count)
	}
	configured := fixture(t, "audit:\n  enabled: false\n", "", "", "", "")
	start(t, configured)
	req := request(configured)
	req.Agent = "fixture-agent --model local"
	if _, err := workflow.Done(context.Background(), req); err != nil {
		t.Fatalf("configured-agent evals Done: %v", err)
	}
	change, err := os.ReadFile(filepath.Join(configured, "changes/0001-manual.md"))
	if err != nil || !strings.Contains(string(change), "fixture-agent --model local") {
		t.Errorf("configured agent omitted from provenance: %v", err)
	}
}
