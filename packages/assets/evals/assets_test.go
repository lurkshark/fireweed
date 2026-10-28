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
	t.Run("C04", func(t *testing.T) {
		text, err := assets.Explain("C04")
		if err != nil {
			t.Fatalf("Explain(C04): %v", err)
		}
		requirePattern(t, text, `(?i)(spec|item)[^.\n]*(non.normative|lacks?[^.\n]*(MUST|SHOULD))`, "spec text without normative wording")
		requirePattern(t, text, `(?i)(add|write|revise|rewrite|edit)[^.\n]*(MUST|SHOULD|normative)`, "corrective action to make the spec normative")
	})
	t.Run("C05", func(t *testing.T) {
		text, err := assets.Explain("C05")
		if err != nil {
			t.Fatalf("Explain(C05): %v", err)
		}
		requirePattern(t, text, `(?i)(foreign|wrong module|wrong prefix)[^.\n]*(ID|identifier)|(ID|identifier)[^.\n]*(foreign|wrong module|wrong prefix)`, "foreign spec ID")
		requirePattern(t, text, `(?i)duplicat[^.\n]*(ID|identifier)|(ID|identifier)[^.\n]*duplicat`, "duplicate spec ID")
		requirePattern(t, text, `(?i)duplicat[^.\n]*edit|repeated[^.\n]*edit`, "duplicate edit identity")
		requirePattern(t, text, `(?i)(use|change|replace|correct|assign|rename)[^.\n]*(module|prefix)`, "corrective action for foreign spec IDs")
		requirePattern(t, text, `(?i)(use|change|replace|correct|assign|rename)[^.\n]*(unique|distinct|unused)`, "corrective action for duplicate spec IDs")
		requirePattern(t, text, `(?i)(remove|combine|merge)[^.\n]*(repeated|duplicate)[^.\n]*edit`, "corrective action for duplicate edits")
	})
	t.Run("C06", func(t *testing.T) {
		text, err := assets.Explain("C06")
		if err != nil {
			t.Fatalf("Explain(C06): %v", err)
		}
		requirePattern(t, text, `(?i)(invalid|malformed|unsupported)[^.\n]*(tag|review|quarantine)`, "invalid spec tag")
		requirePattern(t, text, `(?i)(use|replace|change|correct|remove|edit)[^.\n]*(review|quarantine|tag)`, "corrective action for spec tags")
	})
	t.Run("C07", func(t *testing.T) {
		text, err := assets.Explain("C07")
		if err != nil {
			t.Fatalf("Explain(C07): %v", err)
		}
		requirePattern(t, text, `(?i)(missing|empty|absent)[^.\n]*Motivation`, "missing change Motivation")
		requirePattern(t, text, `(?i)(add|write|fill|provide|complete)[^.\n]*Motivation`, "corrective action to provide Motivation")
	})
	t.Run("C08", func(t *testing.T) {
		text, err := assets.Explain("C08")
		if err != nil {
			t.Fatalf("Explain(C08): %v", err)
		}
		requirePattern(t, text, `(?i)missing[^.\n]*base|base[^.\n]*missing`, "missing edit base")
		requirePattern(t, text, `(?i)stale[^.\n]*base|base[^.\n]*stale`, "stale edit base")
		requirePattern(t, text, `(?i)(absent|missing|not found)[^.\n]*(target|ID|spec)[^.\n]*(preceding|projection)|(target|ID|spec)[^.\n]*(absent|missing|not found)[^.\n]*(preceding|projection)`, "target absent from preceding projection")
		requirePattern(t, text, `(?i)(stamp|update|refresh|correct|fix)[^.\n]*base`, "corrective action for edit base")
		requirePattern(t, text, `(?i)(correct|change|remove|drop|rebase)[^.\n]*(target|ID|edit|change)`, "corrective action for absent target")
	})
	t.Run("C09", func(t *testing.T) {
		text, err := assets.Explain("C09")
		if err != nil {
			t.Fatalf("Explain(C09): %v", err)
		}
		requirePattern(t, text, `(?i)(undeclared|undefined|not declared)[^.\n]*handle`, "undeclared handle reference")
		requirePattern(t, text, `(?i)(declare|add)[^.\n]*(handle|Added)|(correct|change|remove)[^.\n]*reference`, "corrective action for undeclared handle")
	})
	t.Run("X01", func(t *testing.T) {
		text, err := assets.Explain("X01")
		if err != nil {
			t.Fatalf("Explain(X01): %v", err)
		}
		requirePattern(t, text, `(?i)boundar[^.\n]*violation|violation[^.\n]*boundar`, "boundary violation")
		requirePattern(t, text, `(?i)(fix|remove|change|correct)[^.\n]*(import|dependenc|violation)|(declare|add)[^.\n]*dependenc`, "corrective action for boundary violation")
	})
	t.Run("X02", func(t *testing.T) {
		text, err := assets.Explain("X02")
		if err != nil {
			t.Fatalf("Explain(X02): %v", err)
		}
		requirePattern(t, text, `(?i)(required|needed)[^.\n]*boundar[^.\n]*hook[^.\n]*(missing|absent)|(missing|absent)[^.\n]*(required|needed)[^.\n]*boundar[^.\n]*hook`, "missing required boundary hook")
		requirePattern(t, text, `(?i)hook[^.\n]*(fail|timeout|timed out|cancel|not complete|did not complete)|(fail|timeout|timed out|cancel|not complete|did not complete)[^.\n]*hook`, "boundary hook did not complete")
		requirePattern(t, text, `(?i)(configure|add|provide)[^.\n]*hook`, "corrective action for missing hook")
		requirePattern(t, text, `(?i)(fix|repair|rerun|check)[^.\n]*hook`, "corrective action for incomplete hook")
	})
	t.Run("X03", func(t *testing.T) {
		text, err := assets.Explain("X03")
		if err != nil {
			t.Fatalf("Explain(X03): %v", err)
		}
		requirePattern(t, text, `(?i)warn`, "boundary opt-out warning")
		requirePattern(t, text, `(?i)boundar[^.\n]*opt.out[^.\n]*reason|reason[^.\n]*boundar[^.\n]*opt.out`, "reasoned boundary opt-out")
		requirePattern(t, text, `(?i)(review|revisit|remove)[^.\n]*(reason|opt.out)|(configure|add)[^.\n]*hook`, "corrective action for boundary opt-out")
	})
	t.Run("X04", func(t *testing.T) {
		text, err := assets.Explain("X04")
		if err != nil {
			t.Fatalf("Explain(X04): %v", err)
		}
		requirePattern(t, text, `(?i)warn`, "unnecessary boundary enforcement warning")
		requirePattern(t, text, `(?i)(no|none)[^.\n]*boundar[^.\n]*enforcement[^.\n]*(required|needed)|(no|none)[^.\n]*module[^.\n]*(requires?|needs?)[^.\n]*boundar`, "no boundary enforcement required")
		requirePattern(t, text, `(?i)(no (action|fix|change)|leave[^.\n]*hook[^.\n]*absent)|(configure|add)[^.\n]*hook[^.\n]*(if|when)[^.\n]*need`, "action appropriate when enforcement is unnecessary")
	})
	t.Run("U01", func(t *testing.T) {
		text, err := assets.Explain("U01")
		if err != nil {
			t.Fatalf("Explain(U01): %v", err)
		}
		requirePattern(t, text, `(?i)(invalid|unexpected|unsupported|unrecognized)[^.\n]*argument|argument[^.\n]*(invalid|unexpected|unsupported|unrecognized)`, "invalid argument")
		requirePattern(t, text, `(?i)(invalid|unexpected|unsupported|unrecognized)[^.\n]*option|option[^.\n]*(invalid|unexpected|unsupported|unrecognized)`, "invalid option")
		requirePattern(t, text, `(?i)(unsupported|unknown|unrecognized)[^.\n]*command|command[^.\n]*(unsupported|unknown|unrecognized)`, "unsupported command")
		requirePattern(t, text, `(?i)(use|check|consult|see|follow)[^.\n]*(supported|valid)[^.\n]*(syntax|usage)|(check|consult|see|run)[^.\n]*(help|--help)`, "corrective action using supported syntax or help")
	})
	t.Run("U02", func(t *testing.T) {
		text, err := assets.Explain("U02")
		if err != nil {
			t.Fatalf("Explain(U02): %v", err)
		}
		requirePattern(t, text, `(?i)(unknown|unrecognized|invalid)[^.\n]*(lookup|name|playbook|explain)|(lookup|name|playbook|explain)[^.\n]*(unknown|unrecognized|invalid)`, "unknown lookup name")
		requirePattern(t, text, `(?i)(choose|use|select|provide)[^.\n]*(valid|known|listed|supported)[^.\n]*(name|code|playbook|lookup)`, "corrective action to select a valid lookup name")
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
