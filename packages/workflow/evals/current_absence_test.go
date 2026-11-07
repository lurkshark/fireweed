package evals_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/lurkshark/fireweed/packages/workflow"
)

func Test_FLOW_133_CurrentEmptySelectorWithoutOwnedClaim(t *testing.T) {
	assertAbsent := func(t *testing.T, root string) {
		t.Helper()
		_, err := workflow.Current(root, "")
		if err == nil || err.Error() != "P02: worktree holds no change" {
			t.Errorf("Current(root, empty) error = %v; want exact P02 absence signal", err)
		}
	}

	t.Run("unclaimed queued change", func(t *testing.T) {
		root := fixture(t, "", "", "", "", "")
		assertAbsent(t, root)
	})

	t.Run("empty valid queue", func(t *testing.T) {
		root := fixture(t, "", "", "", "", "")
		if err := os.Remove(filepath.Join(root, "changes/0001-manual.md")); err != nil {
			t.Fatal(err)
		}
		assertAbsent(t, root)
	})

	t.Run("claim owned by another worktree", func(t *testing.T) {
		root := fixture(t, "", "", "", "", "")
		other := filepath.Join(t.TempDir(), "linked")
		git(t, root, "worktree", "add", "-b", "linked", other, "HEAD")
		start(t, root)
		assertAbsent(t, other)
	})
}
