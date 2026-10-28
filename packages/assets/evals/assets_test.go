package evals_test

import (
	"regexp"
	"strings"
	"testing"

	"github.com/lurkshark/fireweed/packages/assets"
)

func Test_AST_001_Version(t *testing.T) {
	if got := assets.Version(); got != "0.1.0" {
		t.Errorf("Version() = %q; want 0.1.0", got)
	}
}

func Test_AST_008_Help(t *testing.T) {
	help := assets.Help()
	if !regexp.MustCompile(`(?i)\bfireweed\b`).MatchString(help) {
		t.Errorf("Help() must identify fireweed; got %q", help)
	}
	// Every command in the design's CLI table, plus the explicitly required
	// version command, must be discoverable by its complete command name.
	commands := []string{
		"init", "doctor", "new-module", "new-change", "stamp", "list", "status",
		"show", "modules", "context", "explain", "playbook", "check", "coverage",
		"check-commits", "ci", "pr-body", "start", "verify", "done", "next", "run",
		"release", "commits", "abandon", "drill", "gc", "migrate", "agents-md",
		"version",
	}
	for _, command := range commands {
		pattern := `(^|[^[:alnum:]_-])` + regexp.QuoteMeta(command) + `($|[^[:alnum:]_-])`
		if !regexp.MustCompile(pattern).MatchString(help) {
			t.Errorf("Help() omits command %q", command)
		}
	}
}

func Test_AST_002_Playbooks(t *testing.T) {
	for _, name := range []string{"author-change", "evals", "audit", "impl", "regenerate"} {
		text, err := assets.Playbook(name)
		if err != nil {
			t.Errorf("Playbook(%q) returned error: %v", name, err)
		}
		if strings.TrimSpace(text) == "" {
			t.Errorf("Playbook(%q) returned no instructions", name)
		}
	}
	if _, err := assets.Playbook("not-a-playbook"); err == nil {
		t.Error("Playbook must reject an unknown name")
	}
}

func Test_AST_009_AST_010_AuditRules(t *testing.T) {
	text, err := assets.Playbook("audit")
	if err != nil {
		t.Fatalf("Playbook(audit): %v", err)
	}
	requirePattern(t, text, `(?i)exact[^.\n]*clauses?|clauses?[^.\n]*exact`, "exact spec clauses")
	requirePattern(t, text, `(?i)(no|without|missing|lack)[^.\n]*(test|eval|assert)`, "findings limited to clauses without assertions")
	requirePattern(t, text, `(?i)re.audit[^.\n]*(only|limit)[^.\n]*IDs?[^.\n]*(prior|previous|preceding|earlier)`, "re-audits limited to prior finding IDs")
}

func Test_AST_011_AST_012_AuthorRules(t *testing.T) {
	text, err := assets.Playbook("author-change")
	if err != nil {
		t.Fatalf("Playbook(author-change): %v", err)
	}
	requirePattern(t, text, `(?i)Motivation`, "Motivation requirement")
	requirePattern(t, text, `(?i)user.visible|design|SYSTEM\.md`, "user-visible or design basis for Motivation")
	requirePattern(t, text, `(?i)narrow|remove|scale back`, "recommendation to narrow costly specs")
	requirePattern(t, text, `(?i)cost|expens|worth|value`, "cost or value reason for narrowing")
	requirePattern(t, text, `already_satisfied`, "already_satisfied for narrowed specs")
}

func Test_AST_003_Explanations(t *testing.T) {
	t.Run("C01", func(t *testing.T) {
		text, err := assets.Explain("C01")
		if err != nil {
			t.Fatalf("Explain(C01): %v", err)
		}
		requirePattern(t, text, `(?i)project|config`, "project configuration")
		requirePattern(t, text, `(?i)invalid|malform|misconfig|missing|error`, "cause of invalid configuration")
		requirePattern(t, text, `(?i)(fix|correct|edit|update|repair)[^.\n]*(field|config|file|setting|project)`, "specific corrective action for configuration")
	})
	t.Run("P2-01", func(t *testing.T) {
		text, err := assets.Explain("P2-01")
		if err != nil {
			t.Fatalf("Explain(P2-01): %v", err)
		}
		requirePattern(t, text, `(?i)fail.first|failing (test|eval|assertion)`, "fail-first context")
		requirePattern(t, text, `(?i)missing|absent|lack|no |not |insufficient|invalid`, "cause of missing evidence")
		requirePattern(t, text, `(?i)evidence|report|test|eval|assertion`, "evidence or assertion")
		requirePattern(t, text, `(?i)(write|add|run|record|capture|produce|rerun)[^.\n]*(test|eval|assertion|evidence|report|verify)`, "corrective action to produce genuine fail-first evidence")
	})
	if _, err := assets.Explain("not-a-code"); err == nil {
		t.Error("Explain must reject an unknown code")
	}
}

func requirePattern(t *testing.T, text, pattern, description string) {
	t.Helper()
	if !regexp.MustCompile(pattern).MatchString(text) {
		t.Errorf("missing %s in %q", description, text)
	}
}
