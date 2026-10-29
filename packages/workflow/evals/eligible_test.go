package evals_test

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"github.com/lurkshark/fireweed/packages/project"
	"github.com/lurkshark/fireweed/packages/workflow"
)

// A real Start/Block leaves the first queued change blocked, committed, and
// unclaimed. The other changes are ordinary approved fixture-local queue files.
func eligibilityBase(t *testing.T, blockedModule string) string {
	t.Helper()
	root := blockFixture(t, "", "", "", "", "")
	write(t, root, "packages/delta/SPEC.md", "---\nname: delta\nid_prefix: DEL\ndepends_on: []\n---\n# Delta\n\n## Specs\n\n### DEL-001\nDelta MUST work.\n")
	write(t, root, "changes/0001-manual.md", fmt.Sprintf("---\nmodule: %s\n---\n\n## Motivation\nThe first change blocks.\n\n## Interface\n\n## Added\n\n## Modified\n\n## Removed\n\n## Decisions\n", blockedModule))
	commit(t, root, "blocked change fixture")
	checkEligibilityFixture(t, root)
	start(t, root)
	block(t, root, "fixture blocked")
	return root
}

func eligibilityChange(t *testing.T, root string, number int, slug, module, set, iface string) {
	t.Helper()
	write(t, root, fmt.Sprintf("changes/%04d-%s.md", number, slug), fmt.Sprintf("---\nmodule: %s\n%s---\n\n## Motivation\nThe queued change is needed.\n\n## Interface\n%s\n## Added\n\n## Modified\n\n## Removed\n\n## Decisions\n", module, set, iface))
}

func checkEligibilityFixture(t *testing.T, root string) {
	t.Helper()
	for _, d := range project.Check(root) {
		if !d.Warning {
			t.Fatalf("invalid eligibility fixture: %s %s: %s", d.Code, d.Path, d.Message)
		}
	}
	queue, err := project.Queue(root)
	if err != nil {
		t.Fatalf("fixture Queue: %v", err)
	}
	for _, ch := range queue {
		if _, err := project.Project(root, ch.Slug); err != nil {
			t.Fatalf("fixture projection through %s: %v", ch.Slug, err)
		}
	}
}

func readyEligibilityFixture(t *testing.T, root string) {
	t.Helper()
	checkEligibilityFixture(t, root)
	commit(t, root, "approved eligibility queue")
}

func eligibleSlugs(t *testing.T, root string) []string {
	t.Helper()
	changes, err := workflow.Eligible(root)
	if err != nil {
		t.Fatalf("Eligible: %v", err)
	}
	slugs := make([]string, 0, len(changes))
	for _, ch := range changes {
		if ch.Path == "" || ch.Module == "" {
			t.Fatalf("Eligible returned incomplete public change: %#v", ch)
		}
		slugs = append(slugs, ch.Slug)
	}
	return slugs
}

func assertEligible(t *testing.T, root string, want ...string) {
	t.Helper()
	want = append([]string{}, want...)
	if got := eligibleSlugs(t, root); !reflect.DeepEqual(got, want) {
		t.Errorf("eligible slugs = %v, want %v", got, want)
	}
}

func Test_FLOW_068_EligibleOrderAndEmptyQueues(t *testing.T) {
	t.Run("skips waiting and retains independent order", func(t *testing.T) {
		root := eligibilityBase(t, "alpha")
		eligibilityChange(t, root, 2, "waiting", "beta", "", "")
		eligibilityChange(t, root, 3, "independent", "delta", "", "")
		eligibilityChange(t, root, 4, "transitive-wait", "gamma", "", "")
		eligibilityChange(t, root, 5, "second-independent", "delta", "", "")
		readyEligibilityFixture(t, root)
		assertEligible(t, root, "independent", "second-independent")
	})
	t.Run("empty queue", func(t *testing.T) {
		root := blockFixture(t, "", "", "", "", "")
		if err := os.Remove(filepath.Join(root, "changes/0001-manual.md")); err != nil {
			t.Fatal(err)
		}
		readyEligibilityFixture(t, root)
		assertEligible(t, root)
	})
	t.Run("all waiting", func(t *testing.T) {
		root := eligibilityBase(t, "alpha")
		eligibilityChange(t, root, 2, "waiting", "beta", "", "")
		readyEligibilityFixture(t, root)
		assertEligible(t, root)
	})
}

func Test_FLOW_069_EligibleWaitsForSameModule(t *testing.T) {
	root := eligibilityBase(t, "alpha")
	eligibilityChange(t, root, 2, "same-module", "alpha", "", "")
	eligibilityChange(t, root, 3, "independent", "delta", "", "")
	readyEligibilityFixture(t, root)
	assertEligible(t, root, "independent")
}

func Test_FLOW_070_EligibleWaitsForDirectAndTransitiveDependents(t *testing.T) {
	root := eligibilityBase(t, "alpha")
	eligibilityChange(t, root, 2, "direct", "beta", "", "")
	eligibilityChange(t, root, 3, "transitive", "gamma", "", "")
	eligibilityChange(t, root, 4, "independent", "delta", "", "")
	readyEligibilityFixture(t, root)
	assertEligible(t, root, "independent")
}

func Test_FLOW_071_EligibleInterfaceChangeWaitsForBlockedDependent(t *testing.T) {
	root := eligibilityBase(t, "gamma")
	eligibilityChange(t, root, 2, "direct-interface", "beta", "", "Public beta API changes.\n")
	eligibilityChange(t, root, 3, "direct-internal", "beta", "", "")
	eligibilityChange(t, root, 4, "transitive-interface", "alpha", "", "Public alpha API changes.\n")
	eligibilityChange(t, root, 5, "transitive-internal", "alpha", "", "")
	readyEligibilityFixture(t, root)
	assertEligible(t, root, "direct-internal", "transitive-internal")
}

func Test_FLOW_072_EligibleUsesPrecedingAndOwnProjectedSet(t *testing.T) {
	t.Run("preceding set removes dependency", func(t *testing.T) {
		root := eligibilityBase(t, "alpha")
		eligibilityChange(t, root, 2, "detach-beta", "beta", "set:\n  depends_on: []\n", "")
		eligibilityChange(t, root, 3, "gamma-after-detach", "gamma", "", "")
		readyEligibilityFixture(t, root)
		assertEligible(t, root, "detach-beta", "gamma-after-detach")
	})
	t.Run("own set adds or removes dependency", func(t *testing.T) {
		root := eligibilityBase(t, "alpha")
		eligibilityChange(t, root, 2, "delta-now-dependent", "delta", "set:\n  depends_on: [alpha]\n", "")
		eligibilityChange(t, root, 3, "gamma-now-independent", "gamma", "set:\n  depends_on: []\n", "")
		readyEligibilityFixture(t, root)
		assertEligible(t, root, "gamma-now-independent")
	})
}

func Test_FLOW_011_StartSelectsFirstEligiblePastWaitingHead(t *testing.T) {
	makeQueue := func(t *testing.T) string {
		t.Helper()
		root := eligibilityBase(t, "alpha")
		eligibilityChange(t, root, 2, "waiting", "beta", "", "")
		eligibilityChange(t, root, 3, "first-eligible", "delta", "", "")
		eligibilityChange(t, root, 4, "later-eligible", "delta", "", "")
		readyEligibilityFixture(t, root)
		return root
	}
	t.Run("empty target selects first eligible", func(t *testing.T) {
		root := makeQueue(t)
		s, err := workflow.Start(root, "")
		if err != nil || s.Change.Slug != "first-eligible" {
			t.Errorf("Start(empty) selected %q: %v", s.Change.Slug, err)
		}
	})
	t.Run("explicit first eligible", func(t *testing.T) {
		root := makeQueue(t)
		s, err := workflow.Start(root, "first-eligible")
		if err != nil || s.Change.Slug != "first-eligible" {
			t.Errorf("Start(first-eligible) selected %q: %v", s.Change.Slug, err)
		}
	})
	for _, target := range []string{"waiting", "later-eligible"} {
		t.Run("reject "+target, func(t *testing.T) {
			root := makeQueue(t)
			if _, err := workflow.Start(root, target); err == nil || !strings.HasPrefix(err.Error(), "P01") {
				t.Errorf("Start(%s) = %v, want P01 rejection", target, err)
			}
		})
	}
	t.Run("no eligible change", func(t *testing.T) {
		root := eligibilityBase(t, "alpha")
		eligibilityChange(t, root, 2, "waiting", "beta", "", "")
		readyEligibilityFixture(t, root)
		if _, err := workflow.Start(root, ""); err == nil || !strings.HasPrefix(err.Error(), "P01") {
			t.Errorf("Start(empty) = %v, want P01 rejection", err)
		}
	})
}
