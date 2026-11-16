package evals_test

import (
	"context"
	"fmt"
	"slices"
	"strings"
	"testing"

	"github.com/lurkshark/fireweed/packages/contracts"
)

func setDrillExclude(t *testing.T, f drillFixture, value string) {
	t.Helper()
	config := nextRead(t, f.root, "fireweed.yaml")
	nextWrite(t, f.root, "fireweed.yaml", config+"drill:\n  exclude: "+value+"\n")
}

func Test_RUN_053_DrillFiltersTrackedInputsAndSeedCommit(t *testing.T) {
	for _, tc := range []struct {
		name, setting string
		filter        bool
	}{
		{"configured", `["docs/**", "MAINTAINERS.md", "examples/*.md"]`, true},
		{"omitted", "", false},
		{"empty", "[]", false},
	} {
		t.Run(tc.name, func(t *testing.T) {
			f := newDrillFixture(t)
			if tc.setting != "" {
				setDrillExclude(t, f, tc.setting)
			}
			files := []struct {
				path        string
				matchesGlob bool
			}{
				{"MAINTAINERS.md", true},
				{"docs/rebuild/history.md", true},
				{"examples/old.md", true},
				{"examples/reference/current.md", false},
			}
			for _, file := range files {
				nextWrite(t, f.root, file.path, file.path+"\n")
			}
			agent := nextRead(t, f.observations, "agent.sh")
			nextWrite(t, f.observations, "agent.sh", strings.Replace(agent, "mode=success\n", `git ls-tree -r --name-only HEAD > "$1/input.seed" || exit 12
mode=success
`, 1))
			eval := nextRead(t, f.root, ".fireweed/eval.sh")
			const verifier = `  printf '%s\n' "$1" >> "$observations/eval.order"`
			nextWrite(t, f.root, ".fireweed/eval.sh", strings.Replace(eval, verifier, `  for doc in MAINTAINERS.md docs/rebuild/history.md examples/old.md examples/reference/current.md; do
    if [ "$(cat "$doc")" != "$doc" ]; then echo "verifier lost $doc" >&2; exit 25; fi
  done
`+verifier, 1))
			nextCommit(t, f.root, "tracked drill documentation and exclusions")
			base := nextGit(t, f.root, "rev-parse", "HEAD")

			r, err := f.run(context.Background(), 0)
			if err != nil || r.Result != "pass" || r.Ref == "" || len(r.Steps) != 1 {
				t.Fatalf("drill result=%q ref=%q steps=%#v error=%v", r.Result, r.Ref, r.Steps, err)
			}
			if order := strings.Fields(drillObserved(t, f, "eval.order")); len(order) == 0 || order[0] != "alpha" {
				t.Errorf("target eval did not run in the verification worktree: %q", order)
			}
			input := strings.Fields(drillObserved(t, f, "input.files"))
			seed := strings.Fields(drillObserved(t, f, "input.seed"))
			for _, file := range files {
				want := !(tc.filter && file.matchesGlob)
				if got := slices.Contains(input, "./"+file.path); got != want {
					t.Errorf("worker workspace contains %s = %t, want %t", file.path, got, want)
				}
				if got := slices.Contains(seed, file.path); got != want {
					t.Errorf("worker seed commit contains %s = %t, want %t", file.path, got, want)
				}
				if got := nextRead(t, f.root, file.path); got != file.path+"\n" {
					t.Errorf("invoking checkout changed %s: %q", file.path, got)
				}
				if got := nextGit(t, f.root, "show", r.Ref+":"+file.path); got != file.path {
					t.Errorf("retained reconstruction changed %s: %q", file.path, got)
				}
			}
			if head := nextGit(t, f.root, "rev-parse", "HEAD"); head != base {
				t.Errorf("invoking HEAD moved from %s to %s", base, head)
			}
			drillCleaned(t, f)
		})
	}
}

func Test_RUN_075_DrillRejectsMalformedExclusionsBeforeSetup(t *testing.T) {
	for _, tc := range []struct{ name, setting string }{
		{"scalar", `"docs/**"`},
		{"mapping", `{docs: true}`},
		{"null list", `null`},
		{"nonstring entry", `["docs/**", 42]`},
		{"empty entry", `[""]`},
		{"absolute path", `["/docs/**"]`},
		{"dot component", `["docs/./history.md"]`},
		{"dot-dot component", `["docs/../history.md"]`},
	} {
		t.Run(tc.name, func(t *testing.T) {
			f := newDrillFixture(t)
			setDrillExclude(t, f, tc.setting)
			nextCommit(t, f.root, "malformed drill exclusion")

			_, err := f.run(context.Background(), 0)
			if code, subject, _ := contracts.ErrorCode(err); code != "C01" || !strings.Contains(subject, "drill.exclude") {
				t.Errorf("malformed drill.exclude returned code=%q subject=%q error=%v, want C01 identifying drill.exclude", code, subject, err)
			}
			drillNoSetupOrAgent(t, f)
		})
	}
}

func Test_RUN_076_DrillRejectsExcludedRequiredInputsBeforeSetup(t *testing.T) {
	for _, tc := range []struct{ name, pattern, path string }{
		{"spec", "packages/alpha/SPEC.md", "packages/alpha/SPEC.md"},
		{"interface", "packages/*/api.go", "packages/alpha/api.go"},
		{"eval", "packages/alpha/evals/**", "packages/alpha/evals/alpha_test.go"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			f := newDrillFixture(t)
			setDrillExclude(t, f, fmt.Sprintf("[%q]", tc.pattern))
			nextCommit(t, f.root, "exclude required drill input")

			_, err := f.run(context.Background(), 0)
			if code, subject, _ := contracts.ErrorCode(err); code != "C01" || !strings.Contains(subject+" "+err.Error(), tc.path) {
				t.Errorf("excluding %q returned code=%q subject=%q error=%v, want C01 identifying %s", tc.pattern, code, subject, err, tc.path)
			}
			drillNoSetupOrAgent(t, f)
		})
	}
}

func Test_RUN_077_DrillRejectsExcludedPathRecreatedBySetup(t *testing.T) {
	f := newDrillFixture(t)
	const excluded = "docs/rebuild/history.md"
	nextWrite(t, f.root, excluded, "tracked operator history\n")
	setDrillExclude(t, f, `["docs/**/*.md"]`)
	setup := nextRead(t, f.root, ".fireweed/setup.sh")
	nextWrite(t, f.root, ".fireweed/setup.sh", strings.Replace(setup, "  exit 0\n", `  mkdir -p docs/rebuild || exit 25
  printf 'recreated by setup\n' > docs/rebuild/history.md || exit 25
  cp docs/rebuild/history.md "$1/recreated-input" || exit 25
  exit 0
`, 1))
	nextCommit(t, f.root, "setup recreates excluded drill input")

	r, err := f.run(context.Background(), 0)
	_, subject, _ := contracts.ErrorCode(err)
	if err == nil || !strings.Contains(subject+" "+err.Error(), excluded) {
		t.Errorf("drill returned %v, want an error identifying %s", err, excluded)
	}
	if got := drillObserved(t, f, "recreated-input"); got != "recreated by setup\n" {
		t.Errorf("input setup did not recreate %s: %q", excluded, got)
	}
	if len(r.Steps) != 0 {
		t.Errorf("drill returned %d steps after setup recreated an excluded path", len(r.Steps))
	}
	drillNoAgent(t, f)
	drillCleaned(t, f)
}
