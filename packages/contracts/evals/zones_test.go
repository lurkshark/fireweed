package evals

import (
	c "github.com/lurkshark/fireweed/packages/contracts"
	"testing"
)

func Test_CON005GlobMatching(t *testing.T) {
	for _, tc := range []struct {
		pattern, path string
		want          bool
	}{
		{"api.go", "api.go", true}, {"*.go", "api.go", true}, {"*.go", "src/api.go", false},
		{"evals/**", "evals/test.go", true}, {"evals/**", "evals/deep/test.go", true},
		{"**/*.go", "api.go", true}, {"**/*.go", "a/b/api.go", true},
		{"a/**/b?.go", "a/b1.go", true}, {"a/**/b?.go", "a/x/y/b2.go", true},
		{"a/*/b.go", "a/b.go", false}, {"a/*/b.go", "a/x/y/b.go", false},
		{"[a-c]?.[gh]o", "b2.go", true}, {"[!a-c]*", "zeta", true}, {"[^a-c]*", "beta", false},
		{"*.go", "api.GO", false}, {"literal.go", "literal.go.bak", false},
		{"a**b", "a/x/b", false}, {"café?.go", "café1.go", true},
	} {
		got, err := c.MatchGlob(tc.pattern, tc.path)
		if err != nil || got != tc.want {
			t.Errorf("%q / %q = %v %v, want %v", tc.pattern, tc.path, got, err, tc.want)
		}
	}
	for _, pattern := range []string{"", "/abs", "a/../b", "./a", "a//b", "a/", "a\\b", "[broken", "[z-a]", "[]", "[!]"} {
		if _, err := c.MatchGlob(pattern, "safe"); err == nil {
			t.Errorf("invalid pattern accepted %q", pattern)
		}
	}
	for _, path := range []string{"", "/abs", "../bad", "a/./b", "a//b", "a/", "a\\b"} {
		if _, err := c.MatchGlob("**", path); err == nil {
			t.Errorf("invalid path accepted %q", path)
		}
	}
}

func Test_CON006FutureOverlap(t *testing.T) {
	for _, tc := range []struct {
		a, b string
		want bool
	}{
		{"evals/**", "internal/**", false}, {"**/*.go", "evals/*.go", true},
		{"a/**/b?.go", "a/b1.go", true}, {"a/*/c", "a/x/y/c", false},
		{"[ab]*.go", "[cd]*.go", false}, {"[a-d]?.go", "[c-f]x.go", true},
		{"api.go", "API.go", false}, {"**/evals/**", "packages/*/evals/*.go", true},
		{"[!a-c]*", "[ab]*", false}, {"[!a-c]*", "z*", true},
		{"future/**/x?.go", "future/new/x1.go", true},
		{"a**b", "a/x/b", false}, {"**", "never-created/path.txt", true},
	} {
		for _, p := range [][2]string{{tc.a, tc.b}, {tc.b, tc.a}} {
			got, err := c.GlobsOverlap(p[0], p[1])
			if err != nil || got != tc.want {
				t.Errorf("%q ∩ %q = %v %v, want %v", p[0], p[1], got, err, tc.want)
			}
		}
	}
	if _, err := c.GlobsOverlap("[broken", "safe"); err == nil {
		t.Error("invalid first pattern accepted")
	}
	if _, err := c.GlobsOverlap("safe", "/absolute"); err == nil {
		t.Error("invalid second pattern accepted")
	}
}

func Test_CON007ZonePrecedence(t *testing.T) {
	all := []string{"**"}
	cases := []struct {
		zones c.Zones
		want  c.Zone
	}{
		{c.Zones{Evals: all, Interface: all, AppendOnly: all, Preserve: all, Disposable: all}, c.EvalZone},
		{c.Zones{Interface: all, AppendOnly: all, Preserve: all, Disposable: all}, c.InterfaceZone},
		{c.Zones{AppendOnly: all, Preserve: all, Disposable: all}, c.AppendOnlyZone},
		{c.Zones{Preserve: all, Disposable: all}, c.PreserveZone},
		{c.Zones{Disposable: all}, c.DisposableZone}, {c.Zones{}, c.OtherZone},
	}
	for _, tc := range cases {
		got, err := c.ZoneFor("packages/m/deep/file.go", "packages/m", tc.zones, nil)
		if err != nil || got != tc.want {
			t.Errorf("zones %+v = %s %v, want %s", tc.zones, got, err, tc.want)
		}
	}
	for _, path := range []string{"fireweed.yaml", "SYSTEM.md", ".fireweed/hooks.go", ".fireweed/deep/x", "AGENTS.md", "CLAUDE.md", "packages/m/AGENTS.md", "packages/other/deep/CLAUDE.md", "docs/plan.md"} {
		got, err := c.ZoneFor(path, "packages/m", cases[0].zones, []string{"docs/**"})
		if err != nil || got != c.Protected {
			t.Errorf("protected %q = %s %v", path, got, err)
		}
	}
	if got, err := c.ZoneFor("packages/m/SPEC.md", "packages/m", cases[0].zones, nil); err != nil || got != c.SpecZone {
		t.Errorf("SPEC precedence: %s %v", got, err)
	}
	if got, err := c.ZoneFor("packages/m/api.go", "packages/m", c.Zones{Interface: []string{"api.go"}, Disposable: all}, []string{"packages/**"}); err != nil || got != c.Protected {
		t.Errorf("protected precedence: %s %v", got, err)
	}
	for _, path := range []string{"packages/m", "packages/more/code.go", "other/code.go"} {
		got, err := c.ZoneFor(path, "packages/m", cases[0].zones, nil)
		if err != nil || got != c.OutsideZone {
			t.Errorf("outside %q: %s %v", path, got, err)
		}
	}
	for _, path := range []string{"../evil", "/absolute", "packages/m/../evil", "packages//m/x"} {
		if _, err := c.ZoneFor(path, "packages/m", c.Zones{}, nil); err == nil {
			t.Errorf("bad path %q accepted", path)
		}
	}
	if _, err := c.ZoneFor("fireweed.yaml", "packages/m", c.Zones{Disposable: []string{"[invalid"}}, nil); err == nil {
		t.Error("stricter protected zone hid invalid module glob")
	}
	if _, err := c.ZoneFor("packages/m/a.go", "packages/m", c.Zones{}, []string{"[invalid"}); err == nil {
		t.Error("invalid protected glob accepted")
	}
	if _, err := c.ZoneFor("safe", "../bad", c.Zones{}, nil); err == nil {
		t.Error("invalid module path accepted")
	}
}
