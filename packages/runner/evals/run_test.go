package evals_test

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/lurkshark/fireweed/packages/checks"
	"github.com/lurkshark/fireweed/packages/project"
	"github.com/lurkshark/fireweed/packages/runner"
	"github.com/lurkshark/fireweed/packages/workflow"
)

type runFixture struct{ root, agents string }

var runModules = []struct{ name, prefix string }{{"alpha", "ALP"}, {"beta", "BET"}, {"gamma", "GAM"}}

func newRunFixture(t *testing.T, queued int, unmerged, failBeta bool) runFixture {
	t.Helper()
	root, err := filepath.EvalSymlinks(t.TempDir())
	if err != nil {
		t.Fatal(err)
	}
	agents, err := filepath.EvalSymlinks(t.TempDir())
	if err != nil {
		t.Fatal(err)
	}
	mode := "normal"
	if failBeta {
		mode = "fail-beta"
	}
	nextWrite(t, agents, "worker.sh", `#!/bin/sh
held=''
for file in changes/*.md; do
  [ -f "$file" ] || continue
  case "$(sed -n 's/^phase: //p' "$file" | head -1)" in evals|impl) held="$file"; break;; esac
done
[ -n "$held" ] || { echo 'no held change' >&2; exit 30; }
phase=$(sed -n 's/^phase: //p' "$held" | head -1)
module=$(sed -n 's/^module: //p' "$held" | head -1)
case "$module" in alpha) prefix=ALP;; beta) prefix=BET;; gamma) prefix=GAM;; *) exit 31;; esac
printf '%s|%s|%s|%s\n' "$module" "$phase" "$PWD" "$FIREWEED_BIN" >> "$1/invocations"
if [ "$2" = fail-beta ] && [ "$module" = beta ] && [ "$phase" = evals ]; then echo 'held beta failure' >&2; exit 17; fi
sleep 0.12
case "$phase" in
  evals)
    cat > "packages/$module/evals/added_test.go" <<GO
package evals
import (
  "testing"
  "example.invalid/run-fixture/packages/$module"
)
func Test_${prefix}_002_Added(t *testing.T) {
  if $module.Value() != "new" { t.Fatalf("${prefix}-002: got %q", $module.Value()) }
}
GO
    ;;
  impl)
    cat > "packages/$module/internal/value.go" <<'GO'
package internal
func Value() string { return "new" }
GO
    ;;
  *) exit 32;;
esac
printf '%s-%s' "$module" "$phase"
`)
	nextWrite(t, agents, "auditor.sh", `#!/bin/sh
held=''
for file in changes/*.md; do
  [ -f "$file" ] || continue
  if [ "$(sed -n 's/^phase: //p' "$file" | head -1)" = audit ]; then held="$file"; break; fi
done
[ -n "$held" ] || { echo 'no held audit' >&2; exit 33; }
module=$(sed -n 's/^module: //p' "$held" | head -1)
case "$module" in alpha) prefix=ALP;; beta) prefix=BET;; gamma) prefix=GAM;; *) exit 34;; esac
printf '%s|audit|%s|%s\n' "$module" "$PWD" "$FIREWEED_BIN" >> "$1/invocations"
sleep 0.12
printf '\n## Audit\nNo findings\nChecked IDs: %s-002\n' "$prefix" >> "$held"
printf '%s-audit' "$module"
`)
	nextWrite(t, root, ".gitignore", ".fireweed/runs/\n.fireweed/go-test.log\n")
	nextWrite(t, root, "go.mod", "module example.invalid/run-fixture\n\ngo 1.22\n")
	nextWrite(t, root, "SYSTEM.md", "# Fixture\n")
	nextWrite(t, root, "fireweed.yaml", fmt.Sprintf(`fireweed: ">=0.1, <0.2"
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
agent:
  command: [sh, %q, %q, %q, "{prompt}"]
audit:
  command: [sh, %q, %q, "{prompt}"]
timeout: 50ms
`, filepath.Join(agents, "worker.sh"), agents, mode, filepath.Join(agents, "auditor.sh"), agents))
	nextWrite(t, root, ".fireweed/eval.sh", `#!/bin/sh
case "$1" in alpha) prefix=ALP;; beta) prefix=BET;; gamma) prefix=GAM;; *) exit 35;; esac
go test "./packages/$1/evals" > .fireweed/go-test.log 2>&1
code=$?
if [ "$code" -eq 0 ]; then result=''; else result='<failure message="Go eval failed"/>'; fi
printf '<testsuite><testcase name="%s-001" classname="fixture"/><testcase name="%s-002" classname="fixture">%s</testcase></testsuite>' "$prefix" "$prefix" "$result" > "$2"
exit "$code"
`)
	nextWrite(t, root, ".fireweed/boundaries.sh", "#!/bin/sh\nexit 0\n")
	for _, m := range runModules {
		nextWrite(t, root, "packages/"+m.name+"/SPEC.md", fmt.Sprintf("---\nname: %s\nid_prefix: %s\ndepends_on: []\n---\n# %s\n\n## Specs\n\n### %s-001\n%s MUST return a value.\n\n## Decisions\n", m.name, m.prefix, m.name, m.prefix, m.name))
		nextWrite(t, root, "packages/"+m.name+"/api.go", fmt.Sprintf("package %s\nimport \"example.invalid/run-fixture/packages/%s/internal\"\nfunc Value() string { return internal.Value() }\n", m.name, m.name))
		nextWrite(t, root, "packages/"+m.name+"/internal/value.go", "package internal\nfunc Value() string { return \"old\" }\n")
		nextWrite(t, root, "packages/"+m.name+"/evals/existing_test.go", fmt.Sprintf("package evals\nimport (\"testing\"; \"example.invalid/run-fixture/packages/%s\")\nfunc Test_%s_001_Value(t *testing.T) { if %s.Value() == \"\" { t.Fatal(\"empty\") } }\n", m.name, m.prefix, m.name))
	}
	nextGit(t, root, "init", "-q")
	nextGit(t, root, "checkout", "-q", "-b", "main")
	nextGit(t, root, "config", "user.name", "Fixture")
	nextGit(t, root, "config", "user.email", "fixture@example.invalid")
	if unmerged {
		nextCommit(t, root, "base fixture")
		nextGit(t, root, "checkout", "-q", "-b", "feature")
	}
	for _, m := range runModules[:queued] {
		change := fmt.Sprintf("---\nmodule: %s\nbases: {}\n---\n\n## Motivation\nNew value needed.\n\n## Interface\nExisting Value function.\n\n## Added\n\n### new:value\n%s MUST return the new value.\n\n## Modified\n\n## Removed\n\n## Decisions\n", m.name, m.name)
		nextWrite(t, root, fmt.Sprintf("changes/%04d-%s.md", indexRunModule(m.name)+1, m.name), change)
	}
	nextCommit(t, root, "queued fixture")
	for _, d := range project.Check(root) {
		if !d.Warning {
			t.Fatalf("bad fixture: %#v", d)
		}
	}
	return runFixture{root, agents}
}

func indexRunModule(name string) int {
	for i, m := range runModules {
		if m.name == name {
			return i
		}
	}
	return -1
}

func (f runFixture) request(max int, allow bool) runner.RunRequest {
	return runner.RunRequest{Root: f.root, Bin: "/fixture/run-bin", AllowUnmerged: allow, Timeout: 20 * time.Second, Max: max}
}

func runArchive(t *testing.T, f runFixture, i int) string {
	t.Helper()
	return nextRead(t, f.root, fmt.Sprintf("changes/archive/%04d-%s.md", i+1, runModules[i].name))
}

func runInvocations(t *testing.T, f runFixture) []string {
	t.Helper()
	b, err := os.ReadFile(filepath.Join(f.agents, "invocations"))
	if os.IsNotExist(err) {
		return nil
	}
	if err != nil {
		t.Fatal(err)
	}
	return strings.Split(strings.TrimSpace(string(b)), "\n")
}

func assertRunApplied(t *testing.T, f runFixture, applied []runner.NextResult, want ...string) {
	t.Helper()
	if len(applied) != len(want) {
		t.Fatalf("Applied count = %d, want %d: %#v", len(applied), len(want), applied)
	}
	for i, name := range want {
		r := applied[i]
		if r.Empty || r.State.Phase != "applied" || r.State.Change.Slug != name || r.State.Snapshot.Root != f.root || r.State.Run == "" || len(r.Steps) != 3 {
			t.Errorf("Applied[%d] = %#v", i, r)
		}
		archive := runArchive(t, f, indexRunModule(name))
		if !strings.Contains(archive, "phase: applied") || !strings.Contains(archive, "No findings; checked "+runModules[indexRunModule(name)].prefix+"-002") {
			t.Errorf("%s archive = %q", name, archive)
		}
		id := runModules[indexRunModule(name)].prefix + "-002"
		for _, phase := range []struct{ name, result string }{{"evals", "fail"}, {"impl", "pass"}} {
			path := filepath.Join(f.root, ".fireweed", "runs", r.State.Run, phase.name, name+".json")
			data, err := os.ReadFile(path)
			if err != nil {
				t.Fatal(err)
			}
			var report checks.Report
			if err := json.Unmarshal(data, &report); err != nil {
				t.Fatal(err)
			}
			found := false
			for _, c := range report.Cases {
				if c.Result == phase.result {
					for _, got := range c.IDs {
						if got == id {
							found = true
						}
					}
				}
			}
			if !found {
				t.Errorf("%s %s report lacks %s %s: %#v", name, phase.name, id, phase.result, report)
			}
		}
	}
}

func assertRunAgentLog(t *testing.T, f runFixture, want ...string) {
	t.Helper()
	got := runInvocations(t, f)
	if len(got) != len(want)*3 {
		t.Fatalf("agent calls = %#v, want %d changes", got, len(want))
	}
	for i, name := range want {
		for j, phase := range []string{"evals", "audit", "impl"} {
			line := fmt.Sprintf("%s|%s|%s|/fixture/run-bin", name, phase, f.root)
			if got[i*3+j] != line {
				t.Errorf("agent call %d = %q, want %q", i*3+j, got[i*3+j], line)
			}
		}
	}
}

func Test_RUN_016_UnlimitedQueueOrder(t *testing.T) {
	f := newRunFixture(t, 3, false, false)
	r, err := runner.Run(context.Background(), f.request(0, false))
	if err != nil {
		t.Fatalf("Run: %#v, %v", r, err)
	}
	if !r.Empty || r.Limited {
		t.Errorf("flags = Empty %v Limited %v", r.Empty, r.Limited)
	}
	assertRunApplied(t, f, r.Applied, "alpha", "beta", "gamma")
	assertRunAgentLog(t, f, "alpha", "beta", "gamma")
	if q, err := project.Queue(f.root); err != nil || len(q) != 0 {
		t.Errorf("remaining queue = %#v, %v", q, err)
	}
}

func Test_RUN_016_NegativeMaxIsUnlimited(t *testing.T) {
	f := newRunFixture(t, 1, false, false)
	r, err := runner.Run(context.Background(), f.request(-1, false))
	if err != nil {
		t.Fatalf("Run: %#v, %v", r, err)
	}
	if !r.Empty || r.Limited {
		t.Errorf("flags = %#v", r)
	}
	assertRunApplied(t, f, r.Applied, "alpha")
}

func Test_RUN_017_RUN_018_RUN_019_BoundedContinuationAndForwarding(t *testing.T) {
	f := newRunFixture(t, 3, true, false)
	first, err := runner.Run(context.Background(), f.request(1, true))
	if err != nil {
		t.Fatalf("first Run: %#v, %v", first, err)
	}
	if first.Empty || !first.Limited {
		t.Errorf("first flags = %#v", first)
	}
	assertRunApplied(t, f, first.Applied, "alpha")
	firstArchive := runArchive(t, f, 0)
	q, err := project.Queue(f.root)
	if err != nil || len(q) != 2 || q[0].Slug != "beta" || q[1].Slug != "gamma" {
		t.Errorf("remaining queue = %#v, %v", q, err)
	}
	second, err := runner.Run(context.Background(), f.request(0, true))
	if err != nil {
		t.Fatalf("continuation: %#v, %v", second, err)
	}
	if !second.Empty || second.Limited {
		t.Errorf("continuation flags = %#v", second)
	}
	assertRunApplied(t, f, second.Applied, "beta", "gamma")
	assertRunAgentLog(t, f, "alpha", "beta", "gamma")
	if archive := runArchive(t, f, 0); archive != firstArchive {
		t.Error("first archive changed during continuation")
	}
	if q, err := project.Queue(f.root); err != nil || len(q) != 0 {
		t.Errorf("remaining queue = %#v, %v", q, err)
	}
}

func Test_RUN_018_InitialEmptyAndExactLimit(t *testing.T) {
	t.Run("initial_empty", func(t *testing.T) {
		f := newRunFixture(t, 0, false, false)
		before := nextGit(t, f.root, "rev-parse", "HEAD")
		r, err := runner.Run(context.Background(), f.request(2, false))
		if err != nil || !r.Empty || r.Limited || len(r.Applied) != 0 {
			t.Errorf("empty Run = %#v, %v", r, err)
		}
		if nextGit(t, f.root, "rev-parse", "HEAD") != before || len(runInvocations(t, f)) != 0 {
			t.Error("empty Run changed repo or invoked agents")
		}
	})
	t.Run("exact_limit", func(t *testing.T) {
		f := newRunFixture(t, 2, false, false)
		r, err := runner.Run(context.Background(), f.request(2, false))
		if err != nil {
			t.Fatalf("exact Run: %#v, %v", r, err)
		}
		if !r.Empty || r.Limited {
			t.Errorf("exact flags = %#v", r)
		}
		assertRunApplied(t, f, r.Applied, "alpha", "beta")
		assertRunAgentLog(t, f, "alpha", "beta")
	})
}

func Test_RUN_020_HeldFailureRetainsEarlierResults(t *testing.T) {
	f := newRunFixture(t, 3, false, true)
	r, err := runner.Run(context.Background(), f.request(0, false))
	if err == nil || !strings.Contains(err.Error(), "17") {
		t.Errorf("held failure cause = %v", err)
	}
	assertRunApplied(t, f, r.Applied, "alpha")
	state, currentErr := workflow.Current(f.root, "beta")
	if currentErr != nil || state.Phase != "evals" || state.Run == "" {
		t.Errorf("held beta = %#v, %v", state, currentErr)
	}
	if _, statErr := os.Stat(filepath.Join(f.root, "changes", "0002-beta.md")); statErr != nil {
		t.Errorf("beta no longer held: %v", statErr)
	}
	if _, statErr := os.Stat(filepath.Join(f.root, "changes", "0003-gamma.md")); statErr != nil {
		t.Errorf("gamma no longer queued: %v", statErr)
	}
	if _, statErr := os.Stat(filepath.Join(f.root, "changes", "archive", "0003-gamma.md")); !os.IsNotExist(statErr) {
		t.Errorf("gamma archived: %v", statErr)
	}
	if strings.Contains(nextRead(t, f.root, "changes/0003-gamma.md"), "phase:") {
		t.Error("gamma was started")
	}
	got := runInvocations(t, f)
	if len(got) != 4 || got[3] != "beta|evals|"+f.root+"|/fixture/run-bin" {
		t.Errorf("failure invocation order = %#v", got)
	}
}
