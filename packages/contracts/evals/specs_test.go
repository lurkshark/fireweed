package evals

import (
	"crypto/sha256"
	"fmt"
	"reflect"
	"strings"
	"testing"

	c "github.com/lurkshark/fireweed/packages/contracts"
)

func Test_CON001Matching(t *testing.T) {
	for _, name := range []string{"AUTH-014", "auth_014", "AUTH014", "test_auth_014", "x/AuTh014Works", "AUTH014Z", "(AUTH-014)", "AUTH0145 AUTH014"} {
		if !c.MatchesID(name, "", "AUTH-014") {
			t.Errorf("name %q did not match", name)
		}
		if !c.MatchesID("unrelated", name, "AUTH-014") {
			t.Errorf("classname %q did not match", name)
		}
	}
	for _, name := range []string{"AUTH0145", "OAUTH014", "1AUTH014", "éAUTH014", "AUTH014２", "AUTH-140", "AUTH_14", "AUTH--014", "AUTH 014", ""} {
		if c.MatchesID(name, name, "AUTH-014") {
			t.Errorf("false match %q", name)
		}
	}
}

func Test_CON002BaseHash(t *testing.T) {
	canonical := "### AUTH-014 [review]\nA login MUST fail.\n  Indented detail"
	expected := fmt.Sprintf("%x", sha256.Sum256([]byte(canonical)))[:8]
	original := c.Spec{ID: "AUTH-014", Tag: "[review]", Body: "A login MUST fail.\n  Indented detail"}
	if got := c.SpecHash(original); got != expected {
		t.Fatalf("hash %q, want %q", got, expected)
	}
	equivalent := original
	equivalent.Body = "\r\nA login MUST fail. \t\r\n  Indented detail\t\r\n \t\r\n"
	if c.SpecHash(equivalent) != expected {
		t.Error("line endings/trailing whitespace/outer blank lines alter hash")
	}
	for _, different := range []c.Spec{
		{ID: "AUTH-014", Tag: "[quarantine: clock]", Body: original.Body},
		{ID: "AUTH-014", Tag: "", Body: original.Body},
		{ID: "AUTH-014", Tag: original.Tag, Body: original.Body + "!"},
		{ID: "AUTH-015", Tag: original.Tag, Body: original.Body},
	} {
		if c.SpecHash(different) == expected {
			t.Errorf("meaningful change did not affect hash: %+v", different)
		}
	}
}

func Test_CON003SpecGrammar(t *testing.T) {
	body := "A login MUST fail.\n\n```md\n## Specs\n### FAKE-999\n```\n#### detail\nThis is retained."
	doc := "# Module\n## Intro\n### IGNORE-999\n## Specs\nProse before items.\n### AUTH-014\n" + body + "\n\n### AUTH-020\nA message SHOULD be clear.\n## Decisions\n### IGNORE-998\n"
	items, err := c.ParseSpecs(doc, "AUTH")
	want := []c.Spec{{ID: "AUTH-014", Body: body}, {ID: "AUTH-020", Body: "A message SHOULD be clear."}}
	if err != nil || !reflect.DeepEqual(items, want) {
		t.Fatalf("parse got %+v %v, want %+v", items, err, want)
	}
	if empty, err := c.ParseSpecs("## Specs\n\n## Decisions\n", "AUTH"); err != nil || len(empty) != 0 {
		t.Errorf("empty Specs: %v %v", empty, err)
	}
	for _, bad := range []string{
		"# No specs", "## Specs\n## Specs\n", "## Specs\n### WRONG-001\nIt MUST fail.",
		"## Specs\n### AUTH-001\nIt MUST fail.\n### AUTH-001\nIt MUST fail.",
		"## Specs\n### AUTH-1\nIt MUST fail.", "## Specs\n### bad heading\nIt MUST fail.",
		"## Specs\n### AUTH-001\n", "## Specs\n### AUTH-001\nIt does something.",
		"## Specs\n### AUTH-001\nIt MUSTard fail.", "## Specs\n### AUTH-001 extra\nIt MUST fail.",
	} {
		if _, err := c.ParseSpecs(bad, "AUTH"); err == nil {
			t.Errorf("accepted invalid document %q", bad)
		}
	}
	tilde := "## Specs\n### AUTH-001\nIt MUST preserve examples.\n~~~markdown\n## Specs\n### FAKE-001\n~~~"
	if got, err := c.ParseSpecs(tilde, "AUTH"); err != nil || len(got) != 1 || !strings.Contains(got[0].Body, "### FAKE-001") {
		t.Errorf("tilde fence lost: %v %v", got, err)
	}
}

func Test_CON004TagsAndHandles(t *testing.T) {
	for _, tag := range []string{"", "[review]", "[quarantine: clock flaky]", "[review] [quarantine: clock flaky]", "[quarantine: reason] [review]"} {
		items, err := c.ParseItems("### AUTH-001 "+tag+"\nIt MUST work.", "AUTH", false)
		if err != nil || len(items) != 1 || items[0].Tag != tag {
			t.Errorf("tag %q: %+v %v", tag, items, err)
		}
	}
	for _, tag := range []string{"[unknown]", "[review] [review]", "[quarantine]", "[quarantine:]", "[quarantine:   ]", "[review", "review]", "[quarantine: a] [quarantine: b]", "[review] trailing"} {
		if _, err := c.ParseItems("### AUTH-001 "+tag+"\nIt MUST work.", "AUTH", false); err == nil {
			t.Errorf("accepted tag %q", tag)
		}
	}
	for _, handle := range []string{"new:lockout", "new:step-2", "new:123"} {
		got, err := c.ParseItems("### "+handle+"\nIt MUST work.", "AUTH", true)
		if err != nil || len(got) != 1 || got[0].ID != handle {
			t.Errorf("handle %q: %+v %v", handle, got, err)
		}
		if _, err := c.ParseItems("### "+handle+"\nIt MUST work.", "AUTH", false); err == nil {
			t.Errorf("handle allowed without flag: %q", handle)
		}
	}
	for _, handle := range []string{"new:", "new:Upper", "new:under_score", "new:with space"} {
		if _, err := c.ParseItems("### "+handle+"\nIt MUST work.", "AUTH", true); err == nil {
			t.Errorf("accepted malformed handle %q", handle)
		}
	}
	if _, err := c.ParseItems("### new:one\nIt MUST work.\n### new:one\nIt MUST work.", "AUTH", true); err == nil {
		t.Error("duplicate handles accepted")
	}
}
