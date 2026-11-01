package evals_test

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"slices"
	"sort"
	"strings"
	"testing"
	"time"

	"github.com/lurkshark/fireweed/packages/runner"
)

// All scripts and observations live outside the fixture repository. The worker
// only sees the workspace supplied by Drill; the hook sees the full scratch tree.
type drillFixture struct{ root, observations, base string }

func newDrillFixture(t *testing.T) drillFixture {
	t.Helper()
	root, err := filepath.EvalSymlinks(t.TempDir())
	if err != nil {
		t.Fatal(err)
	}
	observations, err := filepath.EvalSymlinks(t.TempDir())
	if err != nil {
		t.Fatal(err)
	}
	setup := filepath.Join(observations, "setup.sh")
	agent := filepath.Join(observations, "agent.sh")
	wrapper := filepath.Join(observations, "wrapper.sh")
	eval := filepath.Join(observations, "eval.sh")
	nextWrite(t, observations, "setup.sh", `#!/bin/sh
printf '%s\n%s\n' "$PWD" "$FIREWEED_BIN" > "$1/setup"
git symbolic-ref -q HEAD > "$1/setup.branch" || :
git rev-parse HEAD > "$1/setup.head"
mkdir -p .fireweed/bin
printf 'SETUP_BINARY_SECRET' > .fireweed/bin/fireweed
if [ -f "$1/setup-fail" ]; then echo SETUP_FAILURE >&2; exit 17; fi
if [ -f "$1/cleanup-fail" ]; then chmod 0500 "$(dirname "$PWD")"; fi
`)
	nextWrite(t, observations, "wrapper.sh", `#!/bin/sh
printf '%s\n' "$FIREWEED_BIN" >> "$1/wrapper.bin"
shift
exec "$@"
`)
	nextWrite(t, observations, "agent.sh", `#!/bin/sh
count=0
[ ! -f "$1/count" ] || count=$(cat "$1/count")
count=$((count + 1))
printf '%s' "$count" > "$1/count"
: > "$1/invoked-$count"
printf '%s' "$PWD" > "$1/agent.cwd"
printf '%s' "$FIREWEED_BIN" > "$1/agent.bin"
printf '%s' "$2" > "$1/prompt-$count"
if [ "$count" -eq 1 ]; then
  find . -path './.git' -prune -o -type f -print | sort > "$1/input.files"
  git for-each-ref --format='%(refname)' > "$1/input.refs"
  git rev-list --count HEAD > "$1/input.history"
  if git cat-file -e "$(cat "$1/base")^{commit}" 2>/dev/null; then echo present > "$1/input.base"; else echo absent > "$1/input.base"; fi
  if [ -e .git/objects/info/alternates ]; then echo present > "$1/input.alternates"; else echo absent > "$1/input.alternates"; fi
fi
mode=success
[ ! -f "$1/mode" ] || mode=$(cat "$1/mode")
mkdir -p packages/alpha/internal
case "$mode" in
  timeout) printf 'PARTIAL_WORK' > packages/alpha/internal/value.go; sleep 10;;
  usage_always) printf 'PARTIAL_WORK' > packages/alpha/internal/value.go; echo 'RATE LIMIT' >&2; exit 9;;
  usage_once) if [ "$count" -eq 1 ]; then printf 'PARTIAL_WORK' > packages/alpha/internal/value.go; echo 'RATE LIMIT' >&2; exit 9; fi; [ "$(cat packages/alpha/internal/value.go 2>/dev/null)" = PARTIAL_WORK ] || { echo 'prior retry work missing' >&2; exit 8; }; printf 'new' > packages/alpha/internal/value.go;;
  agent_retry|mixed_exhaust) if [ "$count" -eq 1 ]; then printf 'PARTIAL_WORK' > packages/alpha/internal/value.go; echo 'WORKER_FAILURE' >&2; exit 7; fi; if [ "$mode" = agent_retry ]; then [ "$(cat packages/alpha/internal/value.go 2>/dev/null)" = PARTIAL_WORK ] || { echo 'prior retry work missing' >&2; exit 8; }; printf 'new' > packages/alpha/internal/value.go; fi;;
  agent_fail_always) printf 'PARTIAL_WORK' > packages/alpha/internal/value.go; echo 'WORKER_FAILURE' >&2; exit 7;;
  verify_retry) if [ "$count" -eq 1 ]; then printf 'wrong' > packages/alpha/internal/value.go; exit 0; fi; [ "$(cat packages/alpha/internal/value.go 2>/dev/null)" = wrong ] || { echo 'prior retry work missing' >&2; exit 8; }; printf 'new' > packages/alpha/internal/value.go;;
  wrong|mixed_exhaust) printf 'wrong' > packages/alpha/internal/value.go;;
  *) printf 'new' > packages/alpha/internal/value.go;;
esac
case "$mode" in
  change_api) printf 'corrupted' > packages/alpha/api.go;;
  delete_eval) rm packages/alpha/evals/alpha_test.go;;
  add_foreign) mkdir -p packages/core/internal; printf 'bad' > packages/core/internal/foreign.go;;
  add_other) printf 'bad' > packages/alpha/extra.txt;;
  change_preserved) printf 'bad' > packages/alpha/internal/retained.go;;
  add_project_protected) printf 'bad' > packages/alpha/internal/project-new.txt;;
  change_protected_dependency_interface) printf 'bad' > packages/core/api.go;;
esac
cat > DRILL_NOTES.md <<'NOTES'
## Contradicted Decisions
- ALP-001: ordered preference may be contradicted.

## Gaps
- ALP-001, module alpha: ordering needs a spec.
- no spec, module aardvark: caller expectation may need a spec.
NOTES
printf 'worker observation'
`)
	nextWrite(t, observations, "eval.sh", `#!/bin/sh
printf '%s\n' "$1" >> "$3/eval.order"
printf '%s' "$PWD" > "$3/eval.cwd"
if [ ! -f packages/core/internal/secret.txt ] || [ ! -f packages/aardvark/internal/secret.txt ] || [ ! -f packages/zeta/internal/secret.txt ]; then echo dependent internals missing >&2; exit 25; fi
case "$1" in alpha) id=ALP-001;; aardvark) id=AAR-001;; zeta) id=ZET-001;; *) exit 26;; esac
if [ "$1" = alpha ] && [ -f "$3/omit-own-coverage" ]; then printf '<testsuite><testcase name="ALP-001_case"/></testsuite>' > "$2"; exit 0; fi
if [ "$1" = aardvark ] && [ -f "$3/omit-dependent-coverage" ]; then printf '<testsuite><testcase name="AAR-001_case"/></testsuite>' > "$2"; exit 0; fi
fail=''
if [ "$1" = alpha ] && [ "$(cat packages/alpha/internal/value.go 2>/dev/null)" != new ]; then fail='own verification failure'; fi
if [ "$1" != alpha ] && [ -f "$3/fail-dependents" ]; then fail="dependent $1 failure"; fi
if [ -z "$fail" ]; then printf '<testsuite><testcase name="%s_case" classname="fixture"/></testsuite>' "$id" > "$2"; exit 0; fi
printf '<testsuite><testcase name="%s_case" classname="fixture"><failure message="%s"/></testcase></testsuite>' "$id" "$fail" > "$2"
printf '%s\n' "$fail" >&2
exit 1
`)
	nextWrite(t, root, ".gitignore", ".fireweed/drills/\n")
	nextWrite(t, root, "SYSTEM.md", "# Fixture guidance\nSYSTEM_RETAINED\n")
	nextWrite(t, root, "AGENTS.md", "AGENT_INSTRUCTION_SECRET\n")
	nextWrite(t, root, "go.mod", "module example.invalid/drill-fixture\n\ngo 1.22\n")
	nextWrite(t, root, "fireweed.yaml", fmt.Sprintf(`fireweed: ">=0.1, <0.2"
version: 1
base: main
modules: ["packages/*"]
module_defaults:
  interface: [api.go]
  evals: [evals/**]
  preserve: [internal/retained.go]
  disposable: [internal/**]
hooks:
  setup: [sh, %q, %q]
  eval: [sh, %q, "{module}", "{report}", %q]
agent:
  command: [sh, %q, %q, "{prompt}"]
  wrap: [sh, %q, %q]
  model: fixture-model
attempts: 2
timeout: 2s
limit_patterns: [RATE LIMIT]
limit_wait: 2ms
limit_budget: 4ms
verify_dependents: false
`, setup, observations, eval, observations, agent, observations, wrapper, observations))
	for _, module := range []struct{ name, prefix, deps string }{
		{"foundation", "FOU", "[]"}, {"core", "COR", "[foundation]"}, {"alpha", "ALP", "[core]"},
		{"aardvark", "AAR", "[alpha]"}, {"zeta", "ZET", "[aardvark]"},
	} {
		nextWrite(t, root, "packages/"+module.name+"/SPEC.md", fmt.Sprintf("---\nname: %s\nid_prefix: %s\ndepends_on: %s\n---\n# %s\n\n## Specs\n\n### %s-001\n%s MUST preserve its value.\n\n## Decisions\n\n[ordering] Prefer ordered values.\n", module.name, module.prefix, module.deps, module.name, module.prefix, module.name))
		nextWrite(t, root, "packages/"+module.name+"/api.go", "package "+module.name+"\n// INTERFACE_"+module.name+"\n")
		nextWrite(t, root, "packages/"+module.name+"/evals/"+module.name+"_test.go", "package evals\n// EVAL_"+module.name+"\n")
		nextWrite(t, root, "packages/"+module.name+"/internal/secret.txt", "INTERNAL_SECRET_"+module.name+"\n")
	}
	nextWrite(t, root, "packages/alpha/internal/value.go", "OLD_IMPLEMENTATION_SECRET")
	nextWrite(t, root, "packages/alpha/internal/retained.go", "RETAINED_OVERLAP")
	nextWrite(t, root, "packages/alpha/other.txt", "RETAINED_OTHER")
	nextGit(t, root, "init", "-q")
	nextGit(t, root, "checkout", "-q", "-b", "main")
	nextGit(t, root, "config", "user.name", "Fixture")
	nextGit(t, root, "config", "user.email", "fixture@example.invalid")
	nextCommit(t, root, "drill fixture")
	base := nextGit(t, root, "rev-parse", "HEAD")
	nextWrite(t, observations, "base", base)
	return drillFixture{root, observations, base}
}

func (f drillFixture) run(ctx context.Context, timeout time.Duration) (runner.DrillResult, error) {
	return runner.Drill(ctx, runner.DrillRequest{Root: f.root, Module: "alpha", Bin: "/fixture/new-controller", Timeout: timeout})
}

func (f drillFixture) mode(t *testing.T, mode string) {
	t.Helper()
	nextWrite(t, f.observations, "mode", mode)
}

func drillProjectProtectedFixture(t *testing.T) drillFixture {
	t.Helper()
	f := newDrillFixture(t)
	config := nextRead(t, f.root, "fireweed.yaml")
	config = strings.Replace(config, "module_defaults:\n", "protected: [\"packages/alpha/internal/project-*.txt\", \"packages/core/api.go\"]\nmodule_defaults:\n", 1)
	nextWrite(t, f.root, "fireweed.yaml", config)
	nextWrite(t, f.root, "packages/alpha/internal/project-protected.txt", "PROJECT_PROTECTED_INPUT\n")
	nextCommit(t, f.root, "project protected drill inputs")
	f.base = nextGit(t, f.root, "rev-parse", "HEAD")
	nextWrite(t, f.observations, "base", f.base)
	return f
}

func drillObserved(t *testing.T, f drillFixture, path string) string {
	t.Helper()
	return nextRead(t, f.observations, path)
}

func drillNoSetupOrAgent(t *testing.T, f drillFixture) {
	t.Helper()
	for _, path := range []string{"setup", "count"} {
		if _, err := os.Stat(filepath.Join(f.observations, path)); !os.IsNotExist(err) {
			t.Errorf("%s exists: %v", path, err)
		}
	}
}

func drillCleaned(t *testing.T, f drillFixture) {
	t.Helper()
	for _, path := range []string{"setup", "agent.cwd"} {
		b, err := os.ReadFile(filepath.Join(f.observations, path))
		if os.IsNotExist(err) {
			continue
		}
		if err != nil {
			t.Fatal(err)
		}
		cwd := strings.TrimSpace(strings.SplitN(string(b), "\n", 2)[0])
		if cwd == f.root || cwd == "" {
			t.Errorf("%s cwd = %q", path, cwd)
			continue
		}
		if _, err := os.Stat(cwd); !os.IsNotExist(err) {
			t.Errorf("%s workspace remains: %v", cwd, err)
		}
	}
	if out := nextGit(t, f.root, "worktree", "list", "--porcelain"); strings.Count(out, "worktree ") != 1 {
		t.Errorf("scratch worktree remains: %q", out)
	}
}

func Test_RUN_050_DrillPreflightRejectsWithoutEffects(t *testing.T) {
	for _, mode := range []string{"unknown", "dirty", "claim"} {
		t.Run(mode, func(t *testing.T) {
			f := newDrillFixture(t)
			request := runner.DrillRequest{Root: f.root, Module: "alpha", Bin: "/fixture/new-controller"}
			switch mode {
			case "unknown":
				request.Module = "missing"
			case "dirty":
				nextWrite(t, f.root, "stray.txt", "dirty")
			case "claim":
				nextGit(t, f.root, "update-ref", "refs/fireweed/claims/live", f.base)
			}
			before := nextGit(t, f.root, "for-each-ref", "--format=%(refname) %(objectname)", "refs/fireweed")
			_, err := runner.Drill(context.Background(), request)
			if err == nil {
				t.Fatal("preflight accepted " + mode)
			}
			want := map[string]string{"unknown": "missing", "dirty": "dirty", "claim": "claim"}[mode]
			if !strings.Contains(strings.ToLower(err.Error()), want) {
				t.Errorf("%s preflight returned unrelated error: %v", mode, err)
			}
			if nextGit(t, f.root, "rev-parse", "HEAD") != f.base || nextGit(t, f.root, "for-each-ref", "--format=%(refname) %(objectname)", "refs/fireweed") != before {
				t.Error("preflight changed invoking refs")
			}
			drillNoSetupOrAgent(t, f)
			drillCleaned(t, f)
		})
	}
}

func Test_RUN_051_SetupRunsDetachedAtHeadAndFailureStops(t *testing.T) {
	f := newDrillFixture(t)
	nextWrite(t, f.observations, "setup-fail", "yes")
	r, err := f.run(context.Background(), 0)
	if err == nil || !strings.Contains(err.Error(), "SETUP_FAILURE") || !strings.Contains(err.Error(), "17") {
		t.Errorf("setup failure = %#v, %v", r, err)
	}
	drillNoAgent(t, f)
	setupRoot := strings.SplitN(drillObserved(t, f, "setup"), "\n", 2)[0]
	common := nextGit(t, f.root, "rev-parse", "--git-common-dir")
	if !filepath.IsAbs(common) {
		common = filepath.Join(f.root, common)
	}
	if !strings.HasPrefix(setupRoot, filepath.Join(common, "fireweed", "worktrees")+string(os.PathSeparator)) || strings.TrimSpace(drillObserved(t, f, "setup.branch")) != "" || strings.TrimSpace(drillObserved(t, f, "setup.head")) != f.base || !strings.Contains(drillObserved(t, f, "setup"), "/fixture/new-controller") {
		t.Errorf("setup did not run in detached scratch with requested binary: %q", drillObserved(t, f, "setup"))
	}
	drillCleaned(t, f)
}

func drillNoAgent(t *testing.T, f drillFixture) {
	t.Helper()
	if _, err := os.Stat(filepath.Join(f.observations, "count")); !os.IsNotExist(err) {
		t.Errorf("agent invoked: %v", err)
	}
}

func Test_RUN_052_RUN_053_RUN_054_RestrictedInputsAndPrompt(t *testing.T) {
	f := newDrillFixture(t)
	r, err := f.run(context.Background(), 0)
	if err != nil || r.Result != "pass" {
		t.Fatalf("drill = %#v, %v", r, err)
	}
	got := strings.Fields(drillObserved(t, f, "input.files"))
	want := []string{"./SYSTEM.md", "./packages/alpha/SPEC.md", "./packages/alpha/api.go", "./packages/alpha/evals/alpha_test.go", "./packages/alpha/internal/retained.go", "./packages/alpha/other.txt", "./packages/core/api.go", "./packages/foundation/api.go"}
	sort.Strings(want)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("reconstruction inputs = %q, want %q", got, want)
	}
	if strings.TrimSpace(drillObserved(t, f, "input.base")) != "absent" || strings.TrimSpace(drillObserved(t, f, "input.alternates")) != "absent" || strings.TrimSpace(drillObserved(t, f, "input.history")) != "1" || strings.Contains(drillObserved(t, f, "input.refs"), "fireweed/") {
		t.Error("input Git contains source history, alternates or refs")
	}
	prompt := drillObserved(t, f, "prompt-1")
	for _, want := range []string{"ALP-001", "Prefer ordered values", "packages/alpha/api.go", "packages/alpha/evals/alpha_test.go", "DRILL_NOTES.md", "Contradicted Decisions", "Gaps", "no spec", "module", "disposable"} {
		if !strings.Contains(prompt, want) {
			t.Errorf("prompt lacks %q", want)
		}
	}
	for _, forbidden := range []string{f.root, strings.TrimSpace(drillObserved(t, f, "setup")), "OLD_IMPLEMENTATION_SECRET", "INTERNAL_SECRET_", "SETUP_BINARY_SECRET"} {
		if forbidden != "" && strings.Contains(prompt, forbidden) {
			t.Errorf("prompt exposes %q", forbidden)
		}
	}
	for _, term := range []string{"outside", "history", "binar", "dependent", "internal"} {
		if !strings.Contains(strings.ToLower(prompt), term) {
			t.Errorf("prompt does not prohibit external/old inputs or dependent internals: %q", term)
		}
	}
	if got := nextGit(t, f.root, "show", r.Ref+":packages/alpha/internal/retained.go"); got != "RETAINED_OVERLAP" {
		t.Errorf("preserve precedence lost: %q", got)
	}
	if out := nextGit(t, f.root, "ls-tree", "-r", "--name-only", r.Ref, "packages/alpha/internal/secret.txt"); out != "" {
		t.Errorf("target disposable secret survived deletion: %q", out)
	}
	for _, module := range []string{"foundation", "core", "aardvark", "zeta"} {
		if got := nextGit(t, f.root, "show", r.Ref+":packages/"+module+"/internal/secret.txt"); got != "INTERNAL_SECRET_"+module {
			t.Errorf("%s verifier internal changed: %q", module, got)
		}
	}
	drillCleaned(t, f)
}

func Test_RUN_052_RUN_053_ProjectProtectedDisposableIsRetainedInputAndVerifierFile(t *testing.T) {
	f := drillProjectProtectedFixture(t)
	script := nextRead(t, f.observations, "agent.sh")
	anchor := "mode=success\n"
	if !strings.Contains(script, anchor) {
		t.Fatal("fixture agent insertion point missing")
	}
	observe := `if [ -f packages/alpha/internal/project-protected.txt ]; then
  cp packages/alpha/internal/project-protected.txt "$1/protected.input"
fi
if [ -f packages/core/api.go ]; then
  cp packages/core/api.go "$1/protected-interface.input"
fi
`
	nextWrite(t, f.observations, "agent.sh", strings.Replace(script, anchor, observe+anchor, 1))
	r, err := f.run(context.Background(), 0)
	if err != nil || r.Result != "pass" || r.Ref == "" {
		t.Fatalf("protected-path drill = %#v, %v", r, err)
	}
	const protected = "packages/alpha/internal/project-protected.txt"
	if paths := nextGit(t, f.root, "ls-tree", "-r", "--name-only", r.Ref, protected); paths != protected {
		t.Errorf("rebuilt verifier lost project-protected file: %q", paths)
	} else if got := nextGit(t, f.root, "show", r.Ref+":"+protected); got != "PROJECT_PROTECTED_INPUT" {
		t.Errorf("rebuilt verifier changed project-protected file: %q", got)
	}
	if got, readErr := os.ReadFile(filepath.Join(f.observations, "protected.input")); readErr != nil || string(got) != "PROJECT_PROTECTED_INPUT\n" {
		t.Errorf("retained reconstruction input = %q, %v", got, readErr)
	}
	if got, readErr := os.ReadFile(filepath.Join(f.observations, "protected-interface.input")); readErr != nil || string(got) != nextRead(t, f.root, "packages/core/api.go") {
		t.Errorf("protected dependency interface = %q, %v", got, readErr)
	}
	files := strings.Fields(drillObserved(t, f, "input.files"))
	for _, want := range []string{"./" + protected, "./packages/core/api.go"} {
		if !slices.Contains(files, want) {
			t.Errorf("reconstruction input omitted %s: %q", want, files)
		}
	}
	drillCleaned(t, f)
}

func Test_RUN_055_ProjectProtectedAdditionAndDependencyInterfaceEditAreRejected(t *testing.T) {
	for _, mode := range []string{"add_project_protected", "change_protected_dependency_interface"} {
		t.Run(mode, func(t *testing.T) {
			f := drillProjectProtectedFixture(t)
			f.mode(t, mode)
			r, err := f.run(context.Background(), 0)
			if err == nil || r.Result != "fail" || len(r.Steps) != 1 || r.Ref != "" {
				t.Errorf("unauthorized %s reconstruction = %#v, %v", mode, r, err)
			}
			if got := nextGit(t, f.root, "rev-parse", "HEAD"); got != f.base {
				t.Errorf("invalid reconstruction moved invoking HEAD to %s", got)
			}
			if refs := nextGit(t, f.root, "for-each-ref", "--format=%(refname)", "refs/fireweed/drills"); refs != "" {
				t.Errorf("invalid reconstruction retained rebuilt ref: %q", refs)
			}
			if _, statErr := os.Stat(filepath.Join(f.observations, "eval.order")); !os.IsNotExist(statErr) {
				t.Errorf("invalid reconstruction reached evaluation: %v", statErr)
			}
			drillCleaned(t, f)
		})
	}
}

func Test_RUN_055_ReconstructionSubmissionAllowsOnlyDisposableAndNotes(t *testing.T) {
	for _, mode := range []string{"change_api", "delete_eval", "add_foreign", "add_other", "change_preserved"} {
		t.Run(mode, func(t *testing.T) {
			f := newDrillFixture(t)
			f.mode(t, mode)
			r, err := f.run(context.Background(), 0)
			if err == nil || r.Result != "fail" || len(r.Steps) == 0 {
				t.Errorf("invalid %s = %#v, %v", mode, r, err)
			}
			if nextGit(t, f.root, "rev-parse", "HEAD") != f.base || nextGit(t, f.root, "show", "HEAD:packages/alpha/internal/value.go") != "OLD_IMPLEMENTATION_SECRET" {
				t.Error("invalid submission changed invoking implementation")
			}
			if nextGit(t, f.root, "for-each-ref", "--format=%(refname)", "refs/fireweed/drills") != "" {
				t.Error("invalid submission retained a rebuilt ref")
			}
			drillCleaned(t, f)
		})
	}
}

func Test_RUN_056_DrillRetriesAgentAndVerificationWithinOneBound(t *testing.T) {
	for _, mode := range []string{"agent_retry", "verify_retry"} {
		t.Run(mode, func(t *testing.T) {
			f := newDrillFixture(t)
			f.mode(t, mode)
			r, err := f.run(context.Background(), 0)
			if err != nil || r.Result != "pass" || len(r.Steps) != 2 {
				t.Fatalf("retry = %#v, %v", r, err)
			}
			if !strings.Contains(drillObserved(t, f, "prompt-2"), map[string]string{"agent_retry": "WORKER_FAILURE", "verify_retry": "own verification failure"}[mode]) {
				t.Errorf("retry prompt lacks prior failure: %q", drillObserved(t, f, "prompt-2"))
			}
			if got := nextGit(t, f.root, "show", r.Ref+":packages/alpha/internal/value.go"); got != "new" {
				t.Errorf("retry did not retain rebuilt work: %q", got)
			}
			drillCleaned(t, f)
		})
	}
	mixed := newDrillFixture(t)
	mixed.mode(t, "mixed_exhaust")
	mixedResult, mixedErr := mixed.run(context.Background(), 0)
	if mixedErr == nil || mixedResult.Result != "fail" || len(mixedResult.Steps) != 2 || !strings.Contains(drillObserved(t, mixed, "prompt-2"), "WORKER_FAILURE") {
		t.Errorf("agent and verification failures did not share attempts: %#v, %v", mixedResult, mixedErr)
	}
	drillCleaned(t, mixed)
	f := newDrillFixture(t)
	f.mode(t, "usage_always")
	nextWrite(t, f.root, "fireweed.yaml", strings.Replace(nextRead(t, f.root, "fireweed.yaml"), "limit_budget: 4ms", "limit_budget: 0s", 1))
	nextCommit(t, f.root, "no usage wait")
	r, err := f.run(context.Background(), 0)
	if err == nil || !runner.IsUsageStop(err) || r.Result != "stopped" || len(r.Steps) != 1 {
		t.Errorf("usage bounded attempts = %#v, %v", r, err)
	}
	drillCleaned(t, f)
}

func Test_RUN_056_ConfiguredTimeoutBoundsDrillWorker(t *testing.T) {
	f := newDrillFixture(t)
	f.mode(t, "timeout")
	config := strings.Replace(nextRead(t, f.root, "fireweed.yaml"), "timeout: 2s", "timeout: 200ms", 1)
	nextWrite(t, f.root, "fireweed.yaml", config)
	nextCommit(t, f.root, "short configured worker timeout")
	start := time.Now()
	r, err := f.run(context.Background(), 0)
	if err == nil || r.Result == "pass" || len(r.Steps) != 1 || time.Since(start) > 4*time.Second {
		t.Errorf("configured drill timeout = %#v, %v after %s", r, err, time.Since(start))
	}
	drillCleaned(t, f)
}

func Test_RUN_057_UsageDoesNotSpendAttemptAndStopsWhenBudgetExhausted(t *testing.T) {
	f := newDrillFixture(t)
	f.mode(t, "usage_once")
	config := strings.Replace(nextRead(t, f.root, "fireweed.yaml"), "attempts: 2", "attempts: 1", 1)
	config = strings.Replace(config, "limit_wait: 2ms\nlimit_budget: 4ms", "limit_wait: 800ms\nlimit_budget: 800ms", 1)
	nextWrite(t, f.root, "fireweed.yaml", config)
	nextCommit(t, f.root, "one attempt")
	r, err := f.run(context.Background(), 0)
	if err != nil || r.Result != "pass" || len(r.Steps) != 2 || !r.Steps[0].UsageLimit {
		t.Errorf("usage retry = %#v, %v", r, err)
	}
	first, firstErr := os.Stat(filepath.Join(f.observations, "invoked-1"))
	second, secondErr := os.Stat(filepath.Join(f.observations, "invoked-2"))
	if firstErr != nil || secondErr != nil {
		t.Errorf("usage invocation timestamps: first=%v second=%v", firstErr, secondErr)
	} else if gap := second.ModTime().Sub(first.ModTime()); gap < 700*time.Millisecond {
		t.Errorf("configured 800ms usage wait lasted only %s", gap)
	}
	drillCleaned(t, f)

	stopped := newDrillFixture(t)
	stopped.mode(t, "usage_always")
	r, err = stopped.run(context.Background(), 0)
	if err == nil || !runner.IsUsageStop(err) || r.Result != "stopped" || len(r.Steps) != 3 || !r.Steps[2].UsageLimit {
		t.Errorf("usage budget = %#v, %v", r, err)
	}
	drillCleaned(t, stopped)

	cancelled := newDrillFixture(t)
	cancelled.mode(t, "timeout")
	ctx, cancel := context.WithTimeout(context.Background(), 250*time.Millisecond)
	defer cancel()
	r, err = cancelled.run(ctx, 5*time.Second)
	if !errors.Is(err, context.DeadlineExceeded) || r.Result == "pass" {
		t.Errorf("cancelled drill = %#v, %v", r, err)
	}
	drillCleaned(t, cancelled)

	timed := newDrillFixture(t)
	timed.mode(t, "timeout")
	r, err = timed.run(context.Background(), 250*time.Millisecond)
	if err == nil || r.Result == "pass" || len(r.Steps) != 1 {
		t.Errorf("timed drill = %#v, %v", r, err)
	}
	drillCleaned(t, timed)
}

func Test_RUN_058_RUN_059_DrillReturnsAllReportsEvenWhenEvalsFail(t *testing.T) {
	f := newDrillFixture(t)
	f.mode(t, "wrong")
	nextWrite(t, f.observations, "fail-dependents", "yes")
	r, err := f.run(context.Background(), 0)
	if err == nil || r.Result != "fail" || len(r.Steps) != 2 {
		t.Errorf("failed verification = %#v, %v", r, err)
	}
	if len(r.Own.Cases) != 1 || r.Own.Cases[0].Result != "fail" || !strings.Contains(r.Own.Cases[0].Failure, "own verification failure") {
		t.Errorf("own failure details lost: %#v", r.Own)
	}
	if len(r.Dependents) != 2 || r.Dependents["aardvark"].Cases[0].Result != "fail" || r.Dependents["zeta"].Cases[0].Result != "fail" {
		t.Errorf("dependent reports lost: %#v", r.Dependents)
	}
	if got := strings.Fields(drillObserved(t, f, "eval.order")); !reflect.DeepEqual(got, []string{"alpha", "aardvark", "zeta", "alpha", "aardvark", "zeta"}) {
		t.Errorf("own and transitive dependent order = %q", got)
	}
	drillCleaned(t, f)

	dependent := newDrillFixture(t)
	nextWrite(t, dependent.observations, "fail-dependents", "yes")
	depResult, depErr := dependent.run(context.Background(), 0)
	if depErr == nil || depResult.Result != "fail" || depResult.Own.Cases[0].Result != "pass" || depResult.Dependents["aardvark"].Cases[0].Result != "fail" || depResult.Dependents["zeta"].Cases[0].Result != "fail" {
		t.Errorf("dependent-only failures passed or lost reports: %#v, %v", depResult, depErr)
	}
	drillCleaned(t, dependent)
}

func Test_RUN_058_OwnCoverageRequiredAfterPassingCases(t *testing.T) {
	f := newDrillFixture(t)
	spec := nextRead(t, f.root, "packages/alpha/SPEC.md")
	nextWrite(t, f.root, "packages/alpha/SPEC.md", strings.Replace(spec, "## Decisions", "### ALP-002\nAlpha MUST also preserve a second value.\n\n## Decisions", 1))
	nextCommit(t, f.root, "second own coverage obligation")
	nextWrite(t, f.observations, "omit-own-coverage", "yes")
	r, err := f.run(context.Background(), 0)
	if err == nil || r.Result != "fail" || !strings.Contains(err.Error(), "ALP-002") || len(r.Own.Cases) != 1 || r.Own.Cases[0].Result != "pass" || len(r.Dependents) != 2 {
		t.Errorf("missing own coverage = %#v, %v", r, err)
	}
	if got := strings.Fields(drillObserved(t, f, "eval.order")); !reflect.DeepEqual(got, []string{"alpha", "aardvark", "zeta", "alpha", "aardvark", "zeta"}) {
		t.Errorf("coverage failure skipped other evals: %q", got)
	}
	drillCleaned(t, f)
}

func Test_RUN_059_DependentCoverageRequiredAfterPassingCases(t *testing.T) {
	f := newDrillFixture(t)
	spec := nextRead(t, f.root, "packages/aardvark/SPEC.md")
	nextWrite(t, f.root, "packages/aardvark/SPEC.md", strings.Replace(spec, "## Decisions", "### AAR-002\nAardvark MUST also preserve a second value.\n\n## Decisions", 1))
	nextCommit(t, f.root, "second dependent coverage obligation")
	nextWrite(t, f.observations, "omit-dependent-coverage", "yes")
	r, err := f.run(context.Background(), 0)
	if err == nil || r.Result != "fail" || !strings.Contains(err.Error(), "AAR-002") || len(r.Own.Cases) != 1 || r.Own.Cases[0].Result != "pass" || len(r.Dependents) != 2 || len(r.Dependents["aardvark"].Cases) != 1 || r.Dependents["aardvark"].Cases[0].Result != "pass" || len(r.Dependents["zeta"].Cases) != 1 || r.Dependents["zeta"].Cases[0].Result != "pass" {
		t.Errorf("missing dependent coverage = %#v, %v", r, err)
	}
	drillCleaned(t, f)
}

func drillReportSection(report, heading string) string {
	start := strings.Index(report, "## "+heading+"\n")
	if start < 0 {
		return ""
	}
	section := report[start+len("## "+heading+"\n"):]
	if end := strings.Index(section, "\n## "); end >= 0 {
		section = section[:end]
	}
	return section
}

func drillReportLine(section string, tokens ...string) bool {
	for _, line := range strings.Split(section, "\n") {
		all := true
		for _, token := range tokens {
			if !strings.Contains(line, token) {
				all = false
				break
			}
		}
		if all {
			return true
		}
	}
	return false
}

func drillEvalEntry(section, module, testcase, id, result string) bool {
	for _, line := range strings.Split(section, "\n") {
		if strings.Contains(line, testcase) && strings.Contains(line, module) && strings.Contains(line, result) && strings.Contains(strings.ReplaceAll(line, testcase, ""), id) {
			return true
		}
	}
	return false
}

func Test_RUN_060_RUN_061_DrillReportAndRetainedRef(t *testing.T) {
	f := newDrillFixture(t)
	r, err := f.run(context.Background(), 0)
	if err != nil || r.Result != "pass" || r.Module != "alpha" || r.Base != f.base || r.Date == "" || r.Agent == "" || r.Model == "" || r.ReportPath != ".fireweed/drills/alpha.md" || r.Ref != "refs/fireweed/drills/alpha/"+r.Date || len(r.Steps) != 1 {
		t.Fatalf("passing result = %#v, %v", r, err)
	}
	if nextGit(t, f.root, "rev-parse", "HEAD") != f.base || nextGit(t, f.root, "for-each-ref", "--format=%(refname)", r.Ref) != r.Ref || nextGit(t, f.root, "show", r.Ref+":packages/alpha/internal/value.go") != "new" || nextGit(t, f.root, "show", "HEAD:packages/alpha/internal/value.go") != "OLD_IMPLEMENTATION_SECRET" {
		t.Error("retained rebuilt ref or invoking checkout wrong")
	}
	report := nextRead(t, f.root, r.ReportPath)
	for _, want := range []string{"date:", "base_commit: " + f.base, "agent:", "model:", "result: pass", "## Own evals", "## Dependent evals", "## Contradicted Decisions", "## Gaps", "ALP-001_case", "AAR-001_case", "ZET-001_case", "ALP-001", "AAR-001", "ZET-001", "pass", "ordered preference may be contradicted", "module alpha", "no spec", "module aardvark"} {
		if !strings.Contains(report, want) {
			t.Errorf("report lacks %q: %q", want, report)
		}
	}
	for _, entry := range []struct{ section, module, testcase, id, result string }{
		{"Own evals", "alpha", "ALP-001_case", "ALP-001", "pass"},
		{"Dependent evals", "aardvark", "AAR-001_case", "AAR-001", "pass"},
		{"Dependent evals", "zeta", "ZET-001_case", "ZET-001", "pass"},
	} {
		if !drillEvalEntry(drillReportSection(report, entry.section), entry.module, entry.testcase, entry.id, entry.result) {
			t.Errorf("%s lacks a complete %s eval entry: %q", entry.section, entry.module, report)
		}
	}
	if !drillReportLine(drillReportSection(report, "Gaps"), "ALP-001", "alpha", "ordering needs a spec") || !drillReportLine(drillReportSection(report, "Gaps"), "no spec", "aardvark", "caller expectation") {
		t.Errorf("gap entries do not each identify nearest spec and target module: %q", report)
	}
	if !strings.Contains(report, "worker observation") {
		t.Errorf("successful worker observation absent from report: %q", report)
	}
	if len(r.Own.Cases) != 1 || len(r.Dependents) != 2 {
		t.Errorf("passing reports = %#v, %#v", r.Own, r.Dependents)
	}
	if drillObserved(t, f, "agent.bin") != "/fixture/new-controller" || strings.TrimSpace(drillObserved(t, f, "wrapper.bin")) != "/fixture/new-controller" {
		t.Error("worker or wrapper did not receive configured binary")
	}
	drillCleaned(t, f)
}

func Test_RUN_060_RUN_061_FailedDrillReportsFailuresAndRetainsCode(t *testing.T) {
	f := newDrillFixture(t)
	f.mode(t, "wrong")
	r, err := f.run(context.Background(), 0)
	if err == nil || r.Result != "fail" || r.Ref == "" || r.ReportPath != ".fireweed/drills/alpha.md" {
		t.Fatalf("failed drill = %#v, %v", r, err)
	}
	if nextGit(t, f.root, "show", r.Ref+":packages/alpha/internal/value.go") != "wrong" {
		t.Error("failed rebuilt code was not retained")
	}
	report := nextRead(t, f.root, r.ReportPath)
	for _, want := range []string{"result: fail", "own verification failure", "ALP-001_case", "ALP-001", "fail", "## Contradicted Decisions", "## Gaps"} {
		if !strings.Contains(report, want) {
			t.Errorf("failure report lacks %q", want)
		}
	}
	if !drillEvalEntry(drillReportSection(report, "Own evals"), "alpha", "ALP-001_case", "ALP-001", "fail") {
		t.Errorf("failed own eval entry lacks module, testcase, spec ID or result: %q", report)
	}
	drillCleaned(t, f)

	agentFailure := newDrillFixture(t)
	agentFailure.mode(t, "agent_fail_always")
	failed, agentErr := agentFailure.run(context.Background(), 0)
	if agentErr == nil || failed.Result != "fail" || len(failed.Steps) != 2 || failed.ReportPath != ".fireweed/drills/alpha.md" {
		t.Errorf("agent exhaustion report = %#v, %v", failed, agentErr)
	}
	if report := nextRead(t, agentFailure.root, ".fireweed/drills/alpha.md"); !strings.Contains(report, "WORKER_FAILURE") || !strings.Contains(report, "result: fail") {
		t.Errorf("agent failure absent from report: %q", report)
	}
	drillCleaned(t, agentFailure)
}

func Test_RUN_062_CleanupOnSuccessAndFailures(t *testing.T) {
	for _, mode := range []string{"success", "agent_retry", "change_api", "usage_always", "timeout"} {
		t.Run(mode, func(t *testing.T) {
			f := newDrillFixture(t)
			f.mode(t, mode)
			if mode == "timeout" {
				_, _ = f.run(context.Background(), 150*time.Millisecond)
			} else {
				_, _ = f.run(context.Background(), 0)
			}
			for _, path := range []string{"setup", "agent.cwd"} {
				if _, err := os.Stat(filepath.Join(f.observations, path)); err != nil {
					t.Errorf("%s scenario did not reach %s: %v", mode, path, err)
				}
			}
			drillCleaned(t, f)
		})
	}
}

func Test_RUN_062_CleanupFailureCannotPass(t *testing.T) {
	f := newDrillFixture(t)
	nextWrite(t, f.observations, "cleanup-fail", "yes")
	common := nextGit(t, f.root, "rev-parse", "--git-common-dir")
	if !filepath.IsAbs(common) {
		common = filepath.Join(f.root, common)
	}
	parent := filepath.Join(common, "fireweed", "worktrees")
	defer func() {
		_ = os.Chmod(parent, 0755)
		setup, err := os.ReadFile(filepath.Join(f.observations, "setup"))
		if err != nil {
			return
		}
		scratch := strings.TrimSpace(strings.SplitN(string(setup), "\n", 2)[0])
		if _, err := os.Stat(scratch); err == nil {
			nextGit(t, f.root, "worktree", "remove", "--force", scratch)
		}
	}()
	r, err := f.run(context.Background(), 0)
	if err == nil || r.Result == "pass" || len(r.Steps) != 1 {
		t.Errorf("cleanup failure reported a passing drill: %#v, %v", r, err)
	}
	setup := strings.SplitN(drillObserved(t, f, "setup"), "\n", 2)[0]
	if _, statErr := os.Stat(setup); statErr != nil {
		t.Errorf("fixture did not make scratch removal fail: %v", statErr)
	}
}
