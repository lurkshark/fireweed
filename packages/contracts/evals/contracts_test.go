package evals_test

import (
	"crypto/sha256"
	"encoding/hex"
	"reflect"
	"strings"
	"testing"

	"github.com/lurkshark/fireweed/packages/contracts"
)

func Test_CON_001_ParseDocument(t *testing.T) {
	source := "---\nname: contracts\ncount: 2\n---\n# Title\n\nBody  \n"
	doc, err := contracts.ParseDocument(source)
	if err != nil {
		t.Fatalf("ParseDocument: %v", err)
	}
	if !reflect.DeepEqual(doc.Frontmatter, map[string]any{"name": "contracts", "count": 2}) {
		t.Errorf("frontmatter = %#v", doc.Frontmatter)
	}
	if doc.Body != "# Title\n\nBody  \n" {
		t.Errorf("body = %q", doc.Body)
	}
}

func Test_CON_001_RejectInvalidFrontmatter(t *testing.T) {
	for _, tc := range []struct{ name, source string }{
		{"malformed", "---\nname: [broken\n---\nbody\n"},
		{"non_mapping", "---\n- first\n---\nbody\n"},
		{"unterminated", "---\nname: contracts\nbody\n"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if _, err := contracts.ParseDocument(tc.source); err == nil {
				t.Errorf("ParseDocument(%q) accepted invalid frontmatter", tc.source)
			}
		})
	}
}

func Test_CON_001_RenderRoundTrip(t *testing.T) {
	want := contracts.Document{
		Frontmatter: map[string]any{"name": "contracts", "enabled": true},
		Body:        "# Title\n\nBody\n",
	}
	source, err := contracts.RenderDocument(want)
	if err != nil {
		t.Fatalf("RenderDocument: %v", err)
	}
	got, err := contracts.ParseDocument(source)
	if err != nil {
		t.Fatalf("parse rendered document: %v", err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("round trip = %#v; want %#v", got, want)
	}
}

func Test_CON_002_Sections(t *testing.T) {
	source := "intro\n## Alpha\nA  \n\n## Beta\nB\n\n## Gamma\nC\n"
	if got, ok := contracts.Section(source, "Beta"); !ok || got != "B\n\n" {
		t.Errorf("Section(Beta) = %q, %v", got, ok)
	}
	if _, ok := contracts.Section(source, "Missing"); ok {
		t.Error("Section(Missing) was reported present")
	}
	want := "intro\n## Alpha\nA  \n\n## Beta\nnew\n## Gamma\nC\n"
	if got := contracts.ReplaceSection(source, "Beta", "new\n"); got != want {
		t.Errorf("replacement changed section or surrounding bytes:\n got %q\nwant %q", got, want)
	}
}

func Test_CON_002_AppendSection(t *testing.T) {
	source := "# Doc\n\n## Existing\nkeep\n"
	got := contracts.ReplaceSection(source, "New", "added\n")
	if !strings.HasPrefix(got, source) {
		t.Errorf("append changed original bytes: %q", got)
	}
	if body, ok := contracts.Section(got, "New"); !ok || body != "added\n" {
		t.Errorf("appended section = %q, %v", body, ok)
	}
}

func Test_CON_002_ParseItems(t *testing.T) {
	source := "### CON-001 [draft]\nFirst line\n\nSecond line\n### CON-002\nNext\n"
	want := []contracts.Item{
		{ID: "CON-001", Tag: "draft", Body: "First line\n\nSecond line\n"},
		{ID: "CON-002", Body: "Next\n"},
	}
	got, err := contracts.ParseItems(source)
	if err != nil {
		t.Fatalf("ParseItems: %v", err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("items = %#v; want %#v", got, want)
	}
}

func Test_CON_002_RejectInvalidItems(t *testing.T) {
	for _, tc := range []struct{ name, source string }{
		{"duplicate", "### CON-001\nFirst\n### CON-001\nSecond\n"},
		{"malformed_heading", "### CON-001 [draft\nBody\n"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if _, err := contracts.ParseItems(tc.source); err == nil {
				t.Errorf("ParseItems(%q) accepted invalid items", tc.source)
			}
		})
	}
}

func Test_CON_002_RenderItem(t *testing.T) {
	for _, tc := range []struct {
		item contracts.Item
		want string
	}{
		{contracts.Item{ID: "CON-001", Tag: "draft", Body: "Body\n"}, "### CON-001 [draft]\nBody\n"},
		{contracts.Item{ID: "CON-002", Body: "Line one\n\nLine two\n"}, "### CON-002\nLine one\n\nLine two\n"},
	} {
		if got := contracts.RenderItem(tc.item); got != tc.want {
			t.Errorf("RenderItem(%#v) = %q; want %q", tc.item, got, tc.want)
		}
	}
}

func Test_CON_003_CanonicalHash(t *testing.T) {
	item := contracts.Item{ID: "CON-013", Tag: "draft", Body: "alpha  \r\nbeta\t\r\n\r\n\r\n"}
	canonical := "### CON-013 [draft]\nalpha\nbeta\n"
	sum := sha256.Sum256([]byte(canonical))
	want := hex.EncodeToString(sum[:])[:8]
	if got := contracts.SpecHash(item); got != want {
		t.Errorf("SpecHash = %q; want %q for canonical item %q", got, want, canonical)
	}
	if got := contracts.SpecHash(contracts.Item{ID: "CON-013", Tag: "draft", Body: "alpha\nbeta\n"}); got != want {
		t.Errorf("equivalent LF item hash = %q; want %q", got, want)
	}
}

func Test_CON_003_IDAndTagAffectHash(t *testing.T) {
	base := contracts.Item{ID: "CON-013", Tag: "draft", Body: "body\n"}
	baseHash := contracts.SpecHash(base)
	withID := base
	withID.ID = "CON-014"
	withTag := base
	withTag.Tag = "ready"
	if baseHash == contracts.SpecHash(withID) || baseHash == contracts.SpecHash(withTag) {
		t.Errorf("ID and tag must affect hash: base=%q id=%q tag=%q", baseHash, contracts.SpecHash(withID), contracts.SpecHash(withTag))
	}
}

func Test_CON_004_PositiveMatches(t *testing.T) {
	for _, tc := range []struct{ name, class string }{
		{"Test_con-001_checks", "Suite"},
		{"Test_CON_001", "Suite"},
		{"Test_CON001", "Suite"},
		{"unrelated", "checks.con_001"},
	} {
		if !contracts.MatchesID("CON-001", tc.name, tc.class) {
			t.Errorf("MatchesID(CON-001, %q, %q) = false", tc.name, tc.class)
		}
	}
}

func Test_CON_004_Boundaries(t *testing.T) {
	for _, tc := range []struct{ name, class string }{
		{"XCON-001", "Suite"},
		{"7CON_001", "Suite"},
		{"CON0012", "Suite"},
	} {
		if contracts.MatchesID("CON-001", tc.name, tc.class) {
			t.Errorf("MatchesID accepted embedded or extended ID in %q", tc.name)
		}
	}
}

func Test_CON_005_MatchGlob(t *testing.T) {
	for _, tc := range []struct {
		pattern, path string
		want          bool
	}{
		{"evals/*.go", "evals/check_test.go", true},
		{"evals/*.go", "evals/nested/check.go", false},
		{"internal/**/test.go", "internal/test.go", true},
		{"internal/**/test.go", "internal/a/b/test.go", true},
		{"internal/**/*.go", "internal/a/b/test.go", true},
		{"api.go", "sub/api.go", false},
	} {
		if got := contracts.MatchGlob(tc.pattern, tc.path); got != tc.want {
			t.Errorf("MatchGlob(%q, %q) = %v; want %v", tc.pattern, tc.path, got, tc.want)
		}
	}
}

func Test_CON_005_ResolveZone(t *testing.T) {
	for _, tc := range []struct {
		name, path, want string
		zones            contracts.Zones
	}{
		{"protected_first", "api.go", "protected", contracts.Zones{Protected: []string{"api.go"}, Interface: []string{"api.go"}}},
		{"root_spec", "SPEC.md", "spec", contracts.Zones{Evals: []string{"*"}}},
		{"evals_before_interface", "evals/check.go", "evals", contracts.Zones{Evals: []string{"evals/**"}, Interface: []string{"**"}}},
		{"interface_before_append", "api.go", "interface", contracts.Zones{Interface: []string{"api.go"}, AppendOnly: []string{"*.go"}}},
		{"append_before_preserve", "notes.md", "append_only", contracts.Zones{AppendOnly: []string{"*.md"}, Preserve: []string{"*.md"}}},
		{"preserve_before_disposable", "scratch.go", "preserve", contracts.Zones{Preserve: []string{"*.go"}, Disposable: []string{"*.go"}}},
		{"disposable", "internal/logic.go", "disposable", contracts.Zones{Disposable: []string{"internal/**"}}},
		{"other", "sub/SPEC.md", "other", contracts.Zones{}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := contracts.ResolveZone(tc.path, tc.zones); got != tc.want {
				t.Errorf("ResolveZone(%q) = %q; want %q", tc.path, got, tc.want)
			}
		})
	}
}
