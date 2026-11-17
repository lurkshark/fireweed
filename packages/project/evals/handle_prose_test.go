package evals_test

import (
	"strings"
	"testing"

	"github.com/lurkshark/fireweed/packages/project"
)

func Test_PRJ_014_CheckHandlesInChangeProse(t *testing.T) {
	const handle = "new:lookup"
	for _, tc := range []struct {
		name, motivation, iface, added string
		wantC09                        bool
	}{
		{"undeclared_in_motivation", "Users need `" + handle + "`.", "", "", true},
		{"undeclared_in_interface", "Users need a lookup.", "Expose `" + handle + "`.", "", true},
		{"declared_in_added", "Users need `" + handle + "`.", "Expose `" + handle + "`.",
			"### " + handle + "\nThe lookup MUST accept a name.\n", false},
	} {
		t.Run(tc.name, func(t *testing.T) {
			root := fixture(t)
			put(t, root, "packages/alpha/SPEC.md", alphaSpecHeader+"### ALP-001\nAlpha MUST work.\n")
			put(t, root, "changes/0001-lookup.md", "---\nmodule: alpha\n---\n\n"+
				"## Motivation\n\n"+tc.motivation+"\n\n"+
				"## Interface\n\n"+tc.iface+"\n\n"+
				"## Added\n\n"+tc.added+"\n## Modified\n\n## Removed\n\n## Decisions\n")

			diagnostics := project.Check(root)
			found := false
			for _, diagnostic := range diagnostics {
				if diagnostic.Warning {
					continue
				}
				if tc.wantC09 && diagnostic.Code == "C09" && strings.Contains(diagnostic.Message, handle) {
					found = true
				} else {
					t.Errorf("unexpected error diagnostic: %#v", diagnostic)
				}
			}
			if tc.wantC09 && !found {
				t.Errorf("Check diagnostics = %#v; want non-warning C09 naming %q", diagnostics, handle)
			}
		})
	}
}
