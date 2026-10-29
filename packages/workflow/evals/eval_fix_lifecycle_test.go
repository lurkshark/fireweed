package evals_test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/lurkshark/fireweed/packages/project"
)

// The approved hook reports the assertion in alpha_test.go as a real testcase.
// The initially wrong assertion fails although alpha's implementation is ready.
func evalFixFixture(t *testing.T, interfaceBody string) string {
	t.Helper()
	root := fixture(t, "", "", "", "", "")
	write(t, root, "changes/0001-manual.md", `---
module: alpha
eval_fix: [ALP-001]
---

## Motivation
Correct the wrong ALP-001 eval while retaining the contract and implementation.

## Interface
`+interfaceBody+`
## Added

## Modified

## Removed

## Decisions

`)
	write(t, root, "packages/alpha/internal/ready", "existing implementation\n")
	write(t, root, "packages/alpha/evals/alpha_test.go", "package alpha_test\n\n// wrong assertion\n")
	hook := readFile(t, root, ".fireweed/eval.sh")
	want := "if [ -f packages/alpha/internal/ready ]; then"
	replacement := "if [ -f packages/alpha/internal/ready ] && grep -q 'correct assertion' packages/alpha/evals/alpha_test.go; then"
	if !strings.Contains(hook, want) {
		t.Fatal("fixture hook condition missing")
	}
	write(t, root, ".fireweed/eval.sh", strings.Replace(hook, want, replacement, 1))
	commit(t, root, "approved eval_fix fixture")
	for _, d := range project.Check(root) {
		if !d.Warning {
			t.Fatalf("invalid eval_fix fixture: %#v", d)
		}
	}
	return root
}

func correctEvalFix(t *testing.T, root string) {
	t.Helper()
	write(t, root, "packages/alpha/evals/alpha_test.go", "package alpha_test\n\n// correct assertion\n")
}

func Test_FLOW_078_FLOW_079_FLOW_036_EvalFixCorrectsEvalAndArchivesWithoutContractOrImplementationEdits(t *testing.T) {
	root := evalFixFixture(t, "\n")
	retained := map[string]string{}
	for _, name := range []string{"packages/alpha/SPEC.md", "packages/alpha/api.go", "packages/alpha/internal/ready"} {
		retained[name] = readFile(t, root, name)
	}
	approved := head(t, root)
	s := start(t, root)
	if s.Phase != "evals" || s.PhaseBase != approved || head(t, root) != approved {
		t.Fatalf("eval_fix Start = %#v, HEAD %s", s, head(t, root))
	}
	correctEvalFix(t, root)
	checked, err := verify(root, true)
	if err != nil || len(checked.Reports["alpha"].Cases) != 2 || head(t, root) != approved {
		t.Fatalf("corrected eval Verify = %#v, %v; HEAD %s", checked, err, head(t, root))
	}
	evals, err := done(root)
	if err != nil || evals.State.Phase != "audit" || head(t, root) == approved {
		t.Fatalf("eval_fix evals Done = %#v, %v", evals, err)
	}
	if got := readFile(t, root, "packages/alpha/evals/alpha_test.go"); !strings.Contains(got, "correct assertion") {
		t.Errorf("corrected eval was not committed: %q", got)
	}
	audit(t, root, "Checked IDs: ALP-001\nNo findings\n")
	impl, err := done(root)
	if err != nil || impl.State.Phase != "impl" {
		t.Fatalf("eval_fix audit Done = %#v, %v", impl, err)
	}
	beforeImpl := head(t, root)
	applied, err := done(root)
	if err != nil || applied.State.Phase != "applied" || head(t, root) == beforeImpl {
		t.Fatalf("eval_fix impl Done = %#v, %v", applied, err)
	}
	if _, err := os.Stat(filepath.Join(root, "changes/0001-manual.md")); !os.IsNotExist(err) {
		t.Errorf("pending eval_fix change remains: %v", err)
	}
	archive := readFile(t, root, "changes/archive/0001-manual.md")
	if !strings.Contains(archive, "phase: applied") || !strings.Contains(archive, "No findings; checked ALP-001") {
		t.Errorf("archive lacks applied phase or checked ID: %s", archive)
	}
	for name, want := range retained {
		if got := readFile(t, root, name); got != want {
			t.Errorf("eval_fix changed retained %s: %q", name, got)
		}
	}
}

func Test_FLOW_079_EvalFixRejectsFailingAndMissingListedEvals(t *testing.T) {
	for _, scenario := range []string{"failing", "missing-listed-coverage"} {
		t.Run(scenario, func(t *testing.T) {
			root := evalFixFixture(t, "\n")
			if scenario == "missing-listed-coverage" {
				correctEvalFix(t, root)
				hook := readFile(t, root, ".fireweed/eval.sh")
				// The hook remains approved fixture infrastructure; only its report
				// selection is changed before Start.
				write(t, root, ".fireweed/eval.sh", strings.ReplaceAll(hook, `name="ALP-001"`, `name="ALP-002"`))
				commit(t, root, "approved hook omits listed ID")
			}
			start(t, root)
			before := head(t, root)
			r, err := verify(root, true)
			if err == nil || head(t, root) != before || len(r.Reports["alpha"].Cases) == 0 {
				t.Fatalf("%s Verify: err=%v, target cases=%#v, HEAD=%s want %s", scenario, err, r.Reports["alpha"].Cases, head(t, root), before)
			}
			if scenario == "failing" && r.Reports["alpha"].Cases[0].Result != "fail" {
				t.Errorf("fixture did not produce a failing target eval: %#v", r.Reports["alpha"])
			}
			if _, err := done(root); err == nil || head(t, root) != before {
				t.Errorf("%s Done advanced: %v", scenario, err)
			}
		})
	}
}

func Test_FLOW_078_EvalFixRejectsInterfaceAndImplementationEditsBeforeHooks(t *testing.T) {
	t.Run("evals interface", func(t *testing.T) {
		root := evalFixFixture(t, "\nThe public alpha interface may be revised.\n")
		marker := hookMarker(t, root)
		start(t, root)
		correctEvalFix(t, root)
		write(t, root, "packages/alpha/api.go", "package alpha // revised public interface\n")
		before := head(t, root)
		if _, err := verify(root, true); err == nil || head(t, root) != before {
			t.Errorf("eval_fix interface Verify = %v", err)
		}
		if _, err := done(root); err == nil || head(t, root) != before {
			t.Errorf("eval_fix interface Done = %v", err)
		}
		if _, err := os.Stat(marker); !os.IsNotExist(err) {
			t.Errorf("hook ran before evals interface rejection: %v", err)
		}
	})
	t.Run("impl implementation", func(t *testing.T) {
		root := evalFixFixture(t, "\n")
		marker := hookMarker(t, root)
		start(t, root)
		correctEvalFix(t, root)
		evals, err := done(root)
		if err != nil || evals.State.Phase != "audit" {
			t.Fatalf("evals Done = %#v, %v", evals, err)
		}
		audit(t, root, "Checked IDs: ALP-001\nNo findings\n")
		impl, err := done(root)
		if err != nil || impl.State.Phase != "impl" {
			t.Fatalf("audit Done = %#v, %v", impl, err)
		}
		if err := os.Remove(marker); err != nil {
			t.Fatal(err)
		}
		write(t, root, "packages/alpha/internal/ready", "unapproved implementation change\n")
		before := head(t, root)
		if _, err := verify(root, true); err == nil || head(t, root) != before {
			t.Errorf("eval_fix impl Verify = %v", err)
		}
		if _, err := done(root); err == nil || head(t, root) != before {
			t.Errorf("eval_fix impl Done = %v", err)
		}
		if _, err := os.Stat(marker); !os.IsNotExist(err) {
			t.Errorf("hook ran before impl implementation rejection: %v", err)
		}
	})
}

func Test_FLOW_036_EvalFixAuditRequiresListedIDBeforeCommit(t *testing.T) {
	root := evalFixFixture(t, "\n")
	start(t, root)
	correctEvalFix(t, root)
	evals, err := done(root)
	if err != nil || evals.State.Phase != "audit" {
		t.Fatalf("evals Done = %#v, %v", evals, err)
	}
	audit(t, root, "Checked IDs: ALP-002\nNo findings\n")
	before := head(t, root)
	if _, err := done(root); err == nil || head(t, root) != before {
		t.Errorf("incomplete eval_fix audit advanced: %v", err)
	}
}
