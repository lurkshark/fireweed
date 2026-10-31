package evals_test

import (
	"testing"

	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/project"
)

func Test_PRJ_057_StampInvalidBasesHasCode(t *testing.T) {
	root := fixture(t)
	queued(t, root, "0001-invalid-bases.md", "bases: [ALP-001]\n", "Needed.\n", "", "### ALP-001\nAlpha MUST change.\n", "")
	err := project.Stamp(root, "invalid-bases")
	if err == nil {
		t.Fatal("Stamp accepted bases that are not a mapping")
	}
	if code, _, ok := contracts.ErrorCode(err); !ok || code != "C08" {
		t.Errorf("Stamp error code = (%q, %v); want (C08, true): %v", code, ok, err)
	}
}

func Test_PRJ_058_MissingPendingChangeHasSubject(t *testing.T) {
	root := fixture(t)
	const target = "changes/0099-absent.md"
	for _, operation := range []struct {
		name string
		call func() error
	}{
		{"Stamp", func() error { return project.Stamp(root, target) }},
		{"Project", func() error { _, err := project.Project(root, target); return err }},
	} {
		t.Run(operation.name, func(t *testing.T) {
			err := operation.call()
			if err == nil {
				t.Fatal("missing pending change was accepted")
			}
			if code, subject, ok := contracts.ErrorCode(err); !ok || code != "U02" || subject != target {
				t.Errorf("error code = (%q, %q, %v); want (U02, %q, true): %v", code, subject, ok, target, err)
			}
		})
	}
}

func Test_PRJ_059_LoadPreservesModuleParseCode(t *testing.T) {
	for _, tc := range []struct {
		name, spec, code string
	}{
		{"malformed_frontmatter", "---\nname: [broken\n---\n# Alpha\n", "C03"},
		{"malformed_spec_heading", "---\nname: alpha\nid_prefix: ALP\ndepends_on: []\n---\n# Alpha\n\n## Specs\n\n### ALP-001 [broken\nAlpha MUST work.\n", "C03"},
		{"duplicate_spec_id", "---\nname: alpha\nid_prefix: ALP\ndepends_on: []\n---\n# Alpha\n\n## Specs\n\n### ALP-001\nAlpha MUST work.\n\n### ALP-001\nAlpha MUST also work.\n", "C05"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			root := fixture(t)
			put(t, root, "packages/alpha/SPEC.md", tc.spec)
			_, err := project.Load(root)
			if err == nil {
				t.Fatal("Load accepted malformed module SPEC.md")
			}
			if code, _, ok := contracts.ErrorCode(err); !ok || code != tc.code {
				t.Errorf("Load error code = (%q, %v); want (%s, true): %v", code, ok, tc.code, err)
			}
		})
	}
}
