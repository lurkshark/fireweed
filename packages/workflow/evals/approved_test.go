package evals_test

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"testing"

	"github.com/lurkshark/fireweed/packages/gitstore"
	"github.com/lurkshark/fireweed/packages/project"
	"github.com/lurkshark/fireweed/packages/workflow"
)

func Test_FLOW_019_FLOW_139_ApprovedSurvivesLiveEdits(t *testing.T) {
	root := fixture(t, "", "", "", "", "")
	alias := filepath.Join(t.TempDir(), "alias")
	if err := os.Symlink(root, alias); err != nil {
		t.Fatal(err)
	}
	started := start(t, root)
	want, err := snapshotJSON(started.Snapshot)
	if err != nil {
		t.Fatal(err)
	}
	assertApproved := func(label string) {
		t.Helper()
		approved, found, err := workflow.Approved(alias)
		if err != nil || !found {
			t.Errorf("Approved with %s inputs = found %t, error %v; want an owned snapshot", label, found, err)
			return
		}
		if approved.Slug != "manual" || approved.Path != pendingCommitChange || approved.Snapshot.Root != root {
			t.Errorf("Approved with %s inputs = slug %q, path %q, root %q; want manual, %s, %s",
				label, approved.Slug, approved.Path, approved.Snapshot.Root, pendingCommitChange, root)
		}
		got, err := snapshotJSON(approved.Snapshot)
		if err != nil || !bytes.Equal(got, want) {
			t.Errorf("Approved with %s inputs differs from Start snapshot: got %s, want %s, error %v", label, got, want, err)
		}
	}

	const config = "fireweed.yaml"
	const spec = "packages/alpha/SPEC.md"
	write(t, root, config, readFile(t, root, config)+"verify_dependents: all\n")
	write(t, root, spec, strings.Replace(readFile(t, root, spec), "retain the second behavior", "change the second behavior", 1))
	write(t, root, pendingCommitChange, strings.Replace(readFile(t, root, pendingCommitChange), newClause, "Alpha MUST return a different value.", 1))
	assertApproved("changed")

	for _, name := range []string{config, spec, pendingCommitChange} {
		write(t, root, name, "---\ninvalid: [\n---\n")
	}
	assertApproved("malformed")

	for _, name := range []string{config, spec, pendingCommitChange} {
		if err := os.Remove(filepath.Join(root, name)); err != nil {
			t.Fatal(err)
		}
	}
	assertApproved("missing")
}

func snapshotJSON(snapshot project.Snapshot) ([]byte, error) {
	// Empty resolved zone lists have the same meaning regardless of slice representation.
	snapshot.Modules = slices.Clone(snapshot.Modules)
	for i := range snapshot.Modules {
		zones := &snapshot.Modules[i].Zones
		for _, patterns := range []*[]string{
			&zones.Protected, &zones.Evals, &zones.Interface,
			&zones.AppendOnly, &zones.Preserve, &zones.Disposable,
		} {
			if len(*patterns) == 0 {
				*patterns = nil
			}
		}
	}
	return json.Marshal(snapshot)
}

func Test_FLOW_140_ApprovedReportsNoOwnedClaim(t *testing.T) {
	root := fixture(t, "", "", "", "", "")
	other := filepath.Join(t.TempDir(), "linked")
	git(t, root, "worktree", "add", "-b", "linked", other, "HEAD")
	if _, found, err := workflow.Approved(root); found || err != nil {
		t.Errorf("Approved without any claim = found %t, error %v; want false, nil", found, err)
	}
	start(t, root)
	write(t, other, pendingCommitChange, "---\nmodule: [\n---\n")
	if _, found, err := workflow.Approved(other); found || err != nil {
		t.Errorf("Approved from worktree with only a foreign claim = found %t, error %v; want false, nil", found, err)
	}
}

func Test_FLOW_140_ApprovedRejectsUnreadableOrMalformedClaim(t *testing.T) {
	root := fixture(t, "", "", "", "", "")
	started := start(t, root)
	malformed, err := gitstore.WriteBlob(root, []byte("{"))
	if err != nil {
		t.Fatal(err)
	}
	const ref = "refs/fireweed/claims/manual"
	previous := started.ClaimOID
	for _, fault := range []struct{ name, oid string }{
		{"unreadable claim blob", started.PhaseBase}, // The object is a commit, not a blob.
		{"malformed claim JSON", malformed},
	} {
		if err := gitstore.CASRef(root, ref, previous, fault.oid); err != nil {
			t.Fatal(err)
		}
		previous = fault.oid
		_, found, err := workflow.Approved(root)
		assertErrorCode(t, "Approved with "+fault.name, err, "P02")
		if found {
			t.Errorf("Approved reported a snapshot with %s", fault.name)
		}
	}
}
