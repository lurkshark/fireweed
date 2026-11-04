package evals_test

import (
	"context"
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"strings"
	"testing"
	"time"

	"github.com/lurkshark/fireweed/packages/gitstore"
	"github.com/lurkshark/fireweed/packages/workflow"
)

var gcNow = time.Date(2026, 11, 4, 12, 0, 0, 0, time.UTC)

func gcCall(t *testing.T, root string, now time.Time) workflow.GCResult {
	t.Helper()
	r, err := workflow.GC(context.Background(), workflow.GCRequest{Root: root, Now: now})
	if err != nil {
		t.Fatalf("GC: %v", err)
	}
	return r
}

func gcRef(t *testing.T, root, name, oid string) {
	t.Helper()
	if err := gitstore.CASRef(root, name, "", oid); err != nil {
		t.Fatalf("create %s: %v", name, err)
	}
}

func gcClaim(t *testing.T, root, name, worktree string) {
	t.Helper()
	data, err := json.Marshal(map[string]string{"worktree": worktree, "change": name})
	if err != nil {
		t.Fatal(err)
	}
	oid, err := gitstore.WriteBlob(root, data)
	if err != nil {
		t.Fatal(err)
	}
	gcRef(t, root, "refs/fireweed/claims/"+name, oid)
}

func gcPresent(t *testing.T, root, name string, want bool) {
	t.Helper()
	_, found, err := gitstore.Ref(root, name)
	if err != nil || found != want {
		t.Errorf("ref %s present=%t, error=%v; want present=%t", name, found, err, want)
	}
}

func gcNames(t *testing.T, root string) []string {
	t.Helper()
	out := git(t, root, "for-each-ref", "--format=%(refname)", "refs/fireweed")
	if out == "" {
		return nil
	}
	return strings.Split(out, "\n")
}

func gcPruned(t *testing.T, got, want []string) {
	t.Helper()
	sort.Strings(want)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("GC Pruned = %v, want sorted %v", got, want)
	}
}

func Test_FLOW_128_GCPrunesMissingAndUnregisteredClaimsOnly(t *testing.T) {
	root := fixture(t, "", "", "", "", "")
	linked := filepath.Join(t.TempDir(), "linked")
	git(t, root, "worktree", "add", "-b", "linked", linked, "HEAD")
	gcClaim(t, root, "main-live", root)
	gcClaim(t, root, "linked-live", linked)
	missing := filepath.Join(t.TempDir(), "removed")
	gcClaim(t, root, "missing", missing)
	unregistered := filepath.Join(t.TempDir(), "unregistered")
	if err := os.Mkdir(unregistered, 0o755); err != nil {
		t.Fatal(err)
	}
	gcClaim(t, root, "unregistered", unregistered)
	gcRef(t, root, "refs/fireweed/other/keep", head(t, root))

	r := gcCall(t, root, gcNow)
	gcPruned(t, r.Pruned, []string{"refs/fireweed/claims/missing", "refs/fireweed/claims/unregistered"})
	for _, name := range []string{"missing", "unregistered"} {
		gcPresent(t, root, "refs/fireweed/claims/"+name, false)
	}
	for _, name := range []string{"main-live", "linked-live"} {
		gcPresent(t, root, "refs/fireweed/claims/"+name, true)
	}
	gcPresent(t, root, "refs/fireweed/other/keep", true)
}

func Test_FLOW_129_GCPrunesOnlyArchivedChangeAttempts(t *testing.T) {
	root := fixture(t, "", "", "", "", "")
	write(t, root, "changes/archive/0002-finished.md", "---\nmodule: alpha\nphase: applied\n---\n")
	write(t, root, "changes/0003-pending.md", "---\nmodule: alpha\n---\n")
	oid := head(t, root)
	archived := []string{"refs/fireweed/attempts/finished/1", "refs/fireweed/attempts/finished/2"}
	for _, name := range append(archived, "refs/fireweed/attempts/pending/1", "refs/fireweed/attempts/unknown/1", "refs/fireweed/attempts/finished/not-an-attempt") {
		gcRef(t, root, name, oid)
	}
	r := gcCall(t, root, gcNow)
	gcPruned(t, r.Pruned, archived)
	for _, name := range archived {
		gcPresent(t, root, name, false)
	}
	for _, name := range []string{"refs/fireweed/attempts/pending/1", "refs/fireweed/attempts/unknown/1", "refs/fireweed/attempts/finished/not-an-attempt"} {
		gcPresent(t, root, name, true)
	}
}

func Test_FLOW_130_FLOW_131_GCDrillAgeUsesConfiguredDurationAndNativeUTCDate(t *testing.T) {
	for _, tc := range []struct {
		name, config      string
		old, equal, newer string
	}{
		{"configured", "gc:\n  drill_max_age: 24h\n", "20261103T115959.999999999", "20261103T120000.000000000", "20261103T120000.000000001"},
		{"default", "", "20261005T115959.999999999", "20261005T120000.000000000", "20261005T120000.000000001"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			root := fixture(t, tc.config, "", "", "", "")
			oid := head(t, root)
			prefix := "refs/fireweed/drills/alpha/"
			for _, date := range []string{tc.old, tc.equal, tc.newer} {
				gcRef(t, root, prefix+date, oid)
			}
			gcRef(t, root, prefix+"unrecognized", oid)
			gcRef(t, root, "refs/fireweed/drills/alpha/not-a-date/extra", oid)
			r := gcCall(t, root, gcNow)
			gcPruned(t, r.Pruned, []string{prefix + tc.old})
			gcPresent(t, root, prefix+tc.old, false)
			for _, name := range []string{prefix + tc.equal, prefix + tc.newer, prefix + "unrecognized", "refs/fireweed/drills/alpha/not-a-date/extra"} {
				gcPresent(t, root, name, true)
			}
		})
	}
	// With zero Now, a native date far in the past is expired against any
	// current clock, while a future date stays live.
	root := fixture(t, "", "", "", "", "")
	old := "refs/fireweed/drills/alpha/20200101T000000.000000000"
	future := "refs/fireweed/drills/alpha/20990101T000000.000000000"
	gcRef(t, root, old, head(t, root))
	gcRef(t, root, future, head(t, root))
	gcPruned(t, gcCall(t, root, time.Time{}).Pruned, []string{old})
	gcPresent(t, root, future, true)
}

func Test_FLOW_131_GCInvalidDurationFailsBeforeAnyPruning(t *testing.T) {
	for _, value := range []string{"0s", "-1h", "invalid"} {
		t.Run(value, func(t *testing.T) {
			root := fixture(t, "gc:\n  drill_max_age: "+value+"\n", "", "", "", "")
			claim := "refs/fireweed/claims/stale"
			old := "refs/fireweed/drills/alpha/20200101T000000.000000000"
			gcClaim(t, root, "stale", filepath.Join(t.TempDir(), "gone"))
			gcRef(t, root, old, head(t, root))
			before := gcNames(t, root)
			r, err := workflow.GC(context.Background(), workflow.GCRequest{Root: root, Now: gcNow})
			if err == nil || !strings.HasPrefix(err.Error(), "C01:") {
				t.Errorf("invalid drill_max_age %q: result=%#v error=%v; want C01", value, r, err)
			}
			if got := gcNames(t, root); !reflect.DeepEqual(got, before) {
				t.Errorf("invalid setting pruned refs: before=%v after=%v", before, got)
			}
			gcPresent(t, root, claim, true)
		})
	}
}

func Test_FLOW_132_GCReportsOnlyDeletedRefsAndPreservesRepositoryState(t *testing.T) {
	root := fixture(t, "", "", "", "", "")
	write(t, root, "changes/archive/0002-finished.md", "---\nmodule: alpha\nphase: applied\n---\n")
	write(t, root, "working.txt", "untracked work\n")
	write(t, root, "SYSTEM.md", "staged work\n")
	git(t, root, "add", "SYSTEM.md")
	gcClaim(t, root, "z-stale", filepath.Join(t.TempDir(), "missing"))
	gcRef(t, root, "refs/fireweed/attempts/finished/1", head(t, root))
	gcRef(t, root, "refs/fireweed/drills/alpha/20200101T000000.000000000", head(t, root))
	gcRef(t, root, "refs/fireweed/other/keep", head(t, root))
	want := []string{"refs/fireweed/claims/z-stale", "refs/fireweed/attempts/finished/1", "refs/fireweed/drills/alpha/20200101T000000.000000000"}
	beforeHead := head(t, root)
	beforeIndex := git(t, root, "write-tree")
	beforeStatus := git(t, root, "status", "--porcelain=v1", "--untracked-files=all")
	beforeWork := readFile(t, root, "working.txt")
	beforeStaged := readFile(t, root, "SYSTEM.md")
	r := gcCall(t, root, gcNow)
	gcPruned(t, r.Pruned, want)
	for _, name := range want {
		gcPresent(t, root, name, false)
	}
	gcPresent(t, root, "refs/fireweed/other/keep", true)
	if got := head(t, root); got != beforeHead {
		t.Errorf("GC moved HEAD: %s -> %s", beforeHead, got)
	}
	if got := git(t, root, "write-tree"); got != beforeIndex {
		t.Errorf("GC changed index tree: %s -> %s", beforeIndex, got)
	}
	if got := git(t, root, "status", "--porcelain=v1", "--untracked-files=all"); got != beforeStatus {
		t.Errorf("GC changed worktree status: %q -> %q", beforeStatus, got)
	}
	if readFile(t, root, "working.txt") != beforeWork || readFile(t, root, "SYSTEM.md") != beforeStaged {
		t.Error("GC changed working files")
	}
	if second := gcCall(t, root, gcNow); len(second.Pruned) != 0 {
		t.Errorf("second GC reported refs it did not delete: %v", second.Pruned)
	}
}
