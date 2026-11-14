package evals_test

import (
	"strings"
	"testing"

	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/project"
)

const changeTargetSlug = "verify-targets"
const changeTargetFile = "0002-" + changeTargetSlug + ".md"
const changeTargetBody = "Alpha MUST work.\n"

func changeTargetFixture(t *testing.T) string {
	t.Helper()
	root := fixture(t)
	put(t, root, "packages/alpha/SPEC.md", alphaSpecHeader+"### ALP-001\n"+changeTargetBody)
	put(t, root, "packages/beta/SPEC.md", betaSpecHeader+"### BET-001\nBeta MUST work.\n")
	return root
}

func checkChangeTargets(t *testing.T, root string, want ...string) {
	t.Helper()
	diagnostics := project.Check(root)
	for _, d := range diagnostics {
		if d.Warning {
			continue
		}
		if len(want) == 0 {
			t.Errorf("valid change targets received an error diagnostic: %#v", d)
			continue
		}
		if d.Code != "C07" || (d.Path != "changes/"+changeTargetFile && !strings.Contains(d.Message, changeTargetSlug)) {
			continue
		}
		matches := true
		for _, part := range want {
			matches = matches && strings.Contains(d.Message, part)
		}
		if matches {
			return
		}
	}
	if len(want) > 0 {
		t.Errorf("Check diagnostics = %#v; want non-warning C07 for %s naming %q", diagnostics, changeTargetFile, want)
	}
}

func Test_PRJ_079_CheckRejectsMalformedChangeTargetLists(t *testing.T) {
	base := contracts.SpecHash(contracts.Item{ID: "ALP-001", Body: changeTargetBody})
	for _, field := range []string{"fixes", "eval_fix", "already_satisfied"} {
		for _, tc := range []struct{ name, value string }{
			{"scalar", "ALP-001"},
			{"null", "null"},
			{"nonstring_member", "[ALP-001, 42]"},
			{"empty_member", `[ALP-001, ""]`},
		} {
			t.Run(field+"/"+tc.name, func(t *testing.T) {
				root := changeTargetFixture(t)
				front := field + ": " + tc.value + "\n"
				modified := ""
				if field == "already_satisfied" {
					front += "bases: {ALP-001: " + base + "}\n"
					modified = "### ALP-001\nAlpha MUST retry.\n"
				}
				queued(t, root, changeTargetFile, front, "Verify the requested behavior.\n", "", modified, "")
				checkChangeTargets(t, root, field)
			})
		}
	}
}

func Test_PRJ_080_CheckFixTargetsAgainstTheProjectionThroughTheirChange(t *testing.T) {
	for _, field := range []string{"fixes", "eval_fix"} {
		for _, tc := range []struct{ name, targets, rejected string }{
			{"unknown", "[ALP-999]", "ALP-999"},
			{"foreign", "[ALP-001, BET-001]", "BET-001"},
			{"added_earlier", "[ALP-002]", ""},
			{"added_later", "[ALP-003]", "ALP-003"},
			{"current", "[ALP-001]", ""},
			{"empty", "[]", ""},
		} {
			t.Run(field+"/"+tc.name, func(t *testing.T) {
				root := changeTargetFixture(t)
				if tc.name == "added_earlier" {
					queued(t, root, "0001-earlier.md", "", "Add an earlier behavior.\n", "### ALP-002\nAlpha MUST index rows.\n", "", "")
				}
				queued(t, root, changeTargetFile, field+": "+tc.targets+"\n", "Verify the requested behavior.\n", "", "", "")
				if tc.name == "added_later" {
					queued(t, root, "0003-later.md", "", "Add a later behavior.\n", "### ALP-003\nAlpha MUST sort rows.\n", "", "")
				}
				if tc.rejected != "" {
					checkChangeTargets(t, root, field, tc.rejected)
				} else {
					checkChangeTargets(t, root)
				}
			})
		}
	}

	t.Run("fixes/removed_in_same_change", func(t *testing.T) {
		root := changeTargetFixture(t)
		base := contracts.SpecHash(contracts.Item{ID: "ALP-001", Body: changeTargetBody})
		queued(t, root, changeTargetFile, "fixes: [ALP-001]\nbases: {ALP-001: "+base+"}\n",
			"Retire the behavior.\n", "", "", "### ALP-001\n")
		checkChangeTargets(t, root, "fixes", "ALP-001")
	})

	t.Run("fixes/repeated_target_added_in_same_change", func(t *testing.T) {
		root := changeTargetFixture(t)
		queued(t, root, changeTargetFile, "fixes: [ALP-002, ALP-002]\n",
			"Add and verify a behavior.\n", "### ALP-002\nAlpha MUST index rows.\n", "", "")
		checkChangeTargets(t, root)
	})
}

func Test_PRJ_081_CheckAlreadySatisfiedTargetsOnlySameChangeAdditionsAndModifications(t *testing.T) {
	const modified = "### ALP-001\nAlpha MUST retry.\n"
	const removed = "### ALP-001\n"
	const local = "### new:local\nAlpha MUST index rows.\n"
	const assigned = "### ALP-002\nAlpha MUST index rows.\n"
	for _, tc := range []struct {
		name, targets, added, modified, removed, rejected string
	}{
		{"unknown", "[ALP-999]", "", "", "", "ALP-999"},
		{"existing_without_edit", "[ALP-001]", "", "", "", "ALP-001"},
		{"removed", "[ALP-001]", "", "", removed, "ALP-001"},
		{"handle_from_earlier_change", "[new:other]", local, "", "", "new:other"},
		{"modified", "[ALP-001]", "", modified, "", ""},
		{"local_handle_before_start", "[new:local]", local, "", "", ""},
		{"assigned_added_id", "[ALP-002]", assigned, "", "", ""},
		{"empty", "[]", "", "", "", ""},
	} {
		t.Run(tc.name, func(t *testing.T) {
			root := changeTargetFixture(t)
			front := "already_satisfied: " + tc.targets + "\n"
			if tc.modified != "" || tc.removed != "" {
				base := contracts.SpecHash(contracts.Item{ID: "ALP-001", Body: changeTargetBody})
				front += "bases: {ALP-001: " + base + "}\n"
			}
			if tc.name == "assigned_added_id" {
				front += "phase: evals\n"
			}
			if tc.name == "handle_from_earlier_change" {
				queued(t, root, "0001-earlier.md", "", "Add an earlier behavior.\n", "### new:other\nAlpha MUST sort rows.\n", "", "")
			}
			queued(t, root, changeTargetFile, front, "Verify the requested behavior.\n", tc.added, tc.modified, tc.removed)
			if tc.rejected != "" {
				checkChangeTargets(t, root, tc.rejected)
			} else {
				checkChangeTargets(t, root)
			}
		})
	}
}
