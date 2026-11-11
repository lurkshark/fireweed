package evals_test

import (
	"slices"
	"strings"
	"testing"
)

const dependentVerificationOutput = "RUN-014 dependent rejected the first implementation attempt"

func Test_RUN_014_NextRunsFullVerificationOncePerAttempt(t *testing.T) {
	for _, tc := range []struct {
		name          string
		failDependent bool
		phases        []string
		dependentRuns []string
	}{
		{"pass", false, []string{"evals", "audit", "impl"}, []string{"1"}},
		{"dependent retry", true, []string{"evals", "audit", "impl", "impl"}, []string{"1", "2"}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			f := singleVerificationFixture(t)
			if tc.failDependent {
				nextWrite(t, f.logDir, "reject-first-impl", "yes\n")
			}
			result, err := f.run(false)
			if err != nil || result.Empty || result.State.Phase != "applied" || result.State.ClaimOID != "" {
				t.Fatalf("Next returned phase %q, empty=%t, claim=%q, error=%v", result.State.Phase, result.Empty, result.State.ClaimOID, err)
			}
			var phases []string
			for _, step := range result.Steps {
				phases = append(phases, step.Phase)
				if step.Agent.ExitCode != 0 {
					t.Errorf("%s agent exited with %d", step.Phase, step.Agent.ExitCode)
				}
			}
			if !slices.Equal(phases, tc.phases) {
				t.Errorf("returned phases = %v, want %v", phases, tc.phases)
			}
			if tc.failDependent {
				first := nextRead(t, f.logDir, "worker-impl-1.prompt")
				retry := nextRead(t, f.logDir, "worker-impl-2.prompt")
				if strings.Contains(first, dependentVerificationOutput) || !strings.Contains(retry, dependentVerificationOutput) {
					t.Error("dependent hook output must appear in the verification retry prompt")
				}
			}
			// The dependent hook runs only during full implementation verification.
			if got := strings.Fields(nextRead(t, f.logDir, "dependent-eval-attempts")); !slices.Equal(got, tc.dependentRuns) {
				t.Errorf("dependent eval hook ran for impl attempts %v, want %v", got, tc.dependentRuns)
			}
		})
	}
}

func singleVerificationFixture(t *testing.T) nextFixture {
	t.Helper()
	f := newNextFixture(t, true, true, "normal")
	nextWrite(t, f.root, "packages/beta/SPEC.md", "---\nname: beta\nid_prefix: BET\ndepends_on: [alpha]\n---\n# Beta\n\n## Specs\n\n### BET-001\nBeta MUST accept Alpha's value.\n\n## Decisions\n")
	nextWrite(t, f.root, "fireweed.yaml", nextRead(t, f.root, "fireweed.yaml")+"verify_dependents: direct\nattempts: 2\n")
	nextWrite(t, f.root, ".fireweed/eval.sh", `#!/bin/sh
case "$1" in
  alpha)
    [ -f packages/alpha/evals/next_test.go ] || exit 30
    if grep -q 'return "old"' packages/alpha/internal/value.go; then
      printf '<testsuite><testcase name="ALP-001"/><testcase name="ALP-002"><failure message="old value"/></testcase></testsuite>' > "$2"
      exit 1
    fi
    grep -q 'return "new"' packages/alpha/internal/value.go || exit 30
    printf '<testsuite><testcase name="ALP-001"/><testcase name="ALP-002"/></testsuite>' > "$2"
    ;;
  beta)
    attempt=$(cat "$3/impl-attempts") || exit 30
    printf '%s\n' "$attempt" >> "$3/dependent-eval-attempts" || exit 30
    if [ "$attempt" = 1 ] && [ -f "$3/reject-first-impl" ]; then
      printf '<testsuite><testcase name="BET-001"><failure message="dependent assertion"/></testcase></testsuite>' > "$2"
      printf '%s\n' '`+dependentVerificationOutput+`' >&2
      exit 1
    fi
    printf '<testsuite><testcase name="BET-001"/></testsuite>' > "$2"
    ;;
  *) exit 30 ;;
esac
`)
	nextCommit(t, f.root, "fixture with counted dependent evals")
	return f
}
