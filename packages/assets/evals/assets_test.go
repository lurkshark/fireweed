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
		requirePattern(t, text, `(?i)(missing|invalid|unknown|undeclared)[^.\n]*(change|pending)[^.\n]*module|(change|pending)[^.\n]*module[^.\n]*(missing|invalid|unknown|undeclared)`, "invalid pending change module target")
		requirePattern(t, text, `(?i)(set|correct|change|select|specify)[^.\n]*(change|pending)[^.\n]*module|(set|correct|change|select|specify)[^.\n]*module[^.\n]*(change|pending|target|existing|valid)`, "corrective action for change module target")
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
		requirePattern(t, text, `(?i)(unknown|ambiguous)[^.\n]*(requested|lookup)[^.\n]*module|(requested|lookup)[^.\n]*module[^.\n]*(unknown|ambiguous)`, "unknown or ambiguous requested module")
		requirePattern(t, text, `(?i)(unknown|ambiguous)[^.\n]*(requested|lookup)[^.\n]*(item|change)|(requested|lookup)[^.\n]*(item|change)[^.\n]*(unknown|ambiguous)`, "unknown or ambiguous requested item or change")
		requirePattern(t, text, `(?i)ambiguous[^.\n]*(projected|local)[^.\n]*handle|(projected|local)[^.\n]*handle[^.\n]*ambiguous`, "ambiguous projected local handle")
		requirePattern(t, text, `(?i)(select|use|provide|specify)[^.\n]*(existing|valid|unambiguous|qualified|full)[^.\n]*(module|item|change|name|ID|handle)`, "corrective action for requested names")
		requirePattern(t, text, `(?i)(unknown|unrecognized|invalid)[^.\n]*recipe name|recipe name[^.\n]*(unknown|unrecognized|invalid)`, "unknown recipe name")
		requirePattern(t, text, `(?i)(choose|use|select)[^.\n]*supported recipe name`, "choose a supported recipe name")
		requirePattern(t, text, `(?i)(choose|use|select|provide)[^.\n]*readable local recipe directory`, "choose a readable local recipe directory")
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

func Test_AST_024_StaticExplanations(t *testing.T) {
	cases := []struct {
		code, cause, action string
	}{
		{"C02", `(?i)(invalid|cycle|unknown|missing)[^.\n]*module[^.\n]*(graph|dependenc)|(module|dependenc)[^.\n]*(graph|cycle|unknown|missing)`, `(?i)(fix|correct|remove|declare|update)[^.\n]*(module|dependenc|graph)`},
		{"C03", `(?i)(unreadable|malformed|invalid)[^.\n]*(document|item|file)|(document|item|file)[^.\n]*(unreadable|malformed|invalid)`, `(?i)(repair|correct|fix|restore)[^.\n]*(document|item|file|syntax|format)`},
		{"C12", `(?i)(invalid|missing|malformed)[^.\n]*(metadata|zone)|(metadata|zone)[^.\n]*(invalid|missing|malformed)`, `(?i)(correct|fix|define|update)[^.\n]*(metadata|zone|field)`},
		{"C13", `(?i)(duplicate|repeated)[^.\n]*(queue|slug)|(queue|slug)[^.\n]*(duplicate|repeated)`, `(?i)(use|assign|choose|rename)[^.\n]*(unused|unique)[^.\n]*slug`},
	}
	for _, tc := range cases {
		t.Run(tc.code, func(t *testing.T) {
			text, err := assets.Explain(tc.code)
			if err != nil {
				t.Fatalf("Explain(%s): %v", tc.code, err)
			}
			requirePattern(t, text, tc.cause, "cause")
			requirePattern(t, text, tc.action, "corrective action")
			if tc.code == "C12" {
				requirePattern(t, text, `(?i)(invalid|missing|malformed)[^.\n]*local recipe metadata|local recipe metadata[^.\n]*(invalid|missing|malformed)`, "invalid local recipe metadata")
				requirePattern(t, text, `(?i)(correct|fix|repair|edit|update)[^.\n]*recipe\.yaml`, "repair recipe.yaml")
			}
			if tc.code == "C02" {
				requirePattern(t, text, `(?i)(empty|missing)[^.\n]*module name|module name[^.\n]*(empty|missing)`, "empty module name cause")
				requirePattern(t, text, `(?i)duplicat[^.\n]*module name|module name[^.\n]*duplicat`, "duplicate module name cause")
				requirePattern(t, text, `(?i)(empty|missing)[^.\n]*prefix|prefix[^.\n]*(empty|missing)`, "empty module prefix cause")
				requirePattern(t, text, `(?i)duplicat[^.\n]*prefix|prefix[^.\n]*duplicat`, "duplicate module prefix cause")
				requirePattern(t, text, `(?i)(correct|fix|rename|assign|give|set|use)[^.\n]*module name`, "corrective action for module names")
				requirePattern(t, text, `(?i)(correct|fix|rename|assign|give|set|use)[^.\n]*prefix`, "corrective action for module prefixes")
			}
		})
	}
}

func Test_AST_025_EvaluationExplanations(t *testing.T) {
	cases := []struct {
		code, cause, action string
	}{
		{"C10", `(?i)(missing|no|lacks?)[^.\n]*(mapped[^.\n]*non.skipped|non.skipped[^.\n]*mapped)[^.\n]*(testcase|case)|(mapped[^.\n]*non.skipped|non.skipped[^.\n]*mapped)[^.\n]*(testcase|case)[^.\n]*(missing|absent)`, `(?i)(add|map|provide|run)[^.\n]*(mapped|non.skipped)[^.\n]*(testcase|case|eval)`},
		{"C11", `(?i)(JUnit|eval|hook)[^.\n]*(report|mapping|exit)[^.\n]*(invalid|missing|mismatch|disagree|fail)|(invalid|missing|mismatch|disagree|fail)[^.\n]*(JUnit|eval|hook)[^.\n]*(report|mapping|exit)`, `(?i)(fix|correct|repair|rerun)[^.\n]*(hook|report|mapping|exit|JUnit)`},
		{"E01", `(?i)(enforced|required)[^.\n]*testcase[^.\n]*(fail|error)|(fail|error)[^.\n]*(enforced|required)[^.\n]*testcase`, `(?i)(fix|repair|correct)[^.\n]*(test|code|implementation|failure)[^.\n]*(rerun|verify|eval)|(fix|repair|correct)[^.\n]*(test|code|implementation|failure)`},
	}
	for _, tc := range cases {
		t.Run(tc.code, func(t *testing.T) {
			text, err := assets.Explain(tc.code)
			if err != nil {
				t.Fatalf("Explain(%s): %v", tc.code, err)
			}
			requirePattern(t, text, tc.cause, "cause")
			requirePattern(t, text, tc.action, "corrective action")
		})
	}
}

func Test_AST_026_LifecycleExplanations(t *testing.T) {
	cases := []struct {
		code, cause, action string
	}{
		{"P01", `(?i)(precondition|prerequisite)[^.\n]*(missing|unmet|fail)|(missing|unmet|fail)[^.\n]*(precondition|prerequisite)`, `(?i)(complete|satisfy|fix|restore)[^.\n]*(precondition|prerequisite|required|state)`},
		{"P02", `(?i)(claim|claimed)[^.\n]*(missing|held|owned|conflict|invalid)|(missing|held|owned|conflict|invalid)[^.\n]*(claim|claimed)`, `(?i)(release|wait|resolve|acquire)[^.\n]*(claim|owner|holder|change)`},
		{"P03", `(?i)(unapproved|unauthorized|unexpected)[^.\n]*(state|alteration|change)|(state|alteration|change)[^.\n]*(unapproved|unauthorized|unexpected)`, `(?i)(restore|revert|review|approve)[^.\n]*(state|change|alteration|work)`},
		{"P04", `(?i)warn[^.\n]*approv|approv[^.\n]*warn`, `(?i)(review|obtain|record|request)[^.\n]*approv`},
		{"P05", `(?i)(stale|mismatch)[^.\n]*base|base[^.\n]*(stale|mismatch)`, `(?i)(refresh|update|repair|restore)[^.\n]*(base|state)[^.\n]*(approved|workflow)|(refresh|update|repair)[^.\n]*base|restore[^.\n]*state[^.\n]*approved`},
		{"P06", `(?i)(blocked|unresolved)[^.\n]*change|change[^.\n]*(blocked|unresolved)`, `(?i)review[^.\n]*Questions`},
		{"P07", `(?i)(shared[^.\n]*base[^.\n]*(bound|limit)|(bound|limit)[^.\n]*shared[^.\n]*base|unapproved[^.\n]*work[^.\n]*(bound|limit))`, `(?i)(review|synchroniz)[^.\n]*(shared|base|change|work)`},
		{"P09", `(?i)(missing|incomplete|invalid)[^.\n]*(No findings|Checked IDs|audit)|(audit|submission)[^.\n]*(missing|incomplete|invalid)[^.\n]*(No findings|Checked IDs)`, `(?is)(submit|record|include|provide).*No findings.*complete.*Checked IDs|(submit|record|include|provide).*complete.*Checked IDs.*No findings`},
		{"Z01", `(?i)(forbidden|unauthorized|outside)[^.\n]*(phase|zone)[^.\n]*edit|edit[^.\n]*(forbidden|unauthorized|outside)[^.\n]*(phase|zone)`, `(?i)(restore|revert)[^.\n]*edit|(move|make)[^.\n]*work[^.\n]*(authorized|appropriate)[^.\n]*(phase|zone)`},
	}
	for _, tc := range cases {
		t.Run(tc.code, func(t *testing.T) {
			text, err := assets.Explain(tc.code)
			if err != nil {
				t.Fatalf("Explain(%s): %v", tc.code, err)
			}
			requirePattern(t, text, tc.cause, "cause")
			requirePattern(t, text, tc.action, "corrective action")
			if tc.code == "P01" {
				requirePattern(t, text, `(?i)new.change[^.\n]*(precondition|prerequisite)|(precondition|prerequisite)[^.\n]*new.change`, "new-change operational precondition")
				requirePattern(t, text, `(?i)lifecycle|\b(start|verify|done)\b`, "lifecycle operational precondition")
				requirePattern(t, text, `(?i)(init|initializ)[^.\n]*(existing|conflict|already present)|(existing|conflict|already present)[^.\n]*(init|initializ)`, "init conflict precondition")
				requirePattern(t, text, `(?i)(init|initializ)[^.\n]*(filesystem|permission|directory|setup|write)|(filesystem|permission|directory|setup|write)[^.\n]*(init|initializ)`, "init filesystem or setup precondition")
				requirePattern(t, text, `(?i)(resolve|remove|rename|move|clear)[^.\n]*(conflict|existing|path|file|directory)`, "corrective action for init conflict")
				requirePattern(t, text, `(?i)(fix|restore|provide|check|ensure|create)[^.\n]*(permission|filesystem|directory|path|setup|writ)`, "corrective action for init setup failure")
			}
			if tc.code == "P02" {
				requirePattern(t, text, `(?i)(unreadable|undecodable)[^.\n]*claim|claim[^.\n]*(unreadable|undecodable)`, "unreadable or undecodable claim cause")
				requirePattern(t, text, `(?i)invalid[^.\n]*claim[^.\n]*data|claim[^.\n]*data[^.\n]*invalid`, "invalid claim data cause")
				requirePattern(t, text, `(?i)(approved|human)[^.\n]*(review|recover)|(review|recover)[^.\n]*(approved|human)`, "approved review or recovery for invalid claims")
				requirePattern(t, text, `(?i)(no|missing|absent)[^.\n]*claim[^.\n]*(start|begin)[^.\n]*eligible[^.\n]*change|(start|begin)[^.\n]*eligible[^.\n]*change[^.\n]*(no|missing|absent)[^.\n]*claim`, "start an eligible change when no claim exists")
				requirePattern(t, text, `(?i)(coordinate|resolve)[^.\n]*(owner|holder|ownership)`, "coordinate with the claim holder")
			}
		})
	}
}

func Test_AST_027_DoctorExplanations(t *testing.T) {
	for _, tc := range []struct {
		code   string
		checks []struct{ pattern, description string }
	}{
		{"C14", []struct{ pattern, description string }{
			{`(?i)(hook|template|canary)[^.\n]*(missing|invalid|broken|fail)|(missing|invalid|broken|fail)[^.\n]*(hook|template|canary)`, "doctor hook, template or canary setup failure"},
			{`(?i)canary[^.\n]*(detect|miss|expos|report)[^.\n]*(violation|failure|error)|(detect|miss|expos|report)[^.\n]*(violation|failure|error)[^.\n]*canary`, "canary detection failure"},
			{`(?i)recover(y|ing)?[^.\n]*(fail|violation|hook|canary)|(fail|violation|hook|canary)[^.\n]*recover(y|ing)?`, "canary recovery failure"},
			{`(?i)(repair|fix|correct|restore)[^.\n]*(hook|template|canary)`, "repair the affected hook, template or canary"},
			{`(?i)(rerun|run again|repeat)[^.\n]*doctor`, "rerun doctor after repair"},
		}},
		{"C15", []struct{ pattern, description string }{
			{`(?i)(agent|worker|audit)[^.\n]*probe[^.\n]*(fail|error|timeout)|(fail|error|timeout)[^.\n]*(agent|worker|audit)[^.\n]*probe`, "failed agent probe"},
			{`(?i)command[^.\n]*(fail|missing|invalid|unavailable|denied)|(fail|missing|invalid|unavailable|denied)[^.\n]*command`, "agent command failure cause"},
			{`(?i)model[^.\n]*(fail|missing|invalid|unavailable|denied)|(fail|missing|invalid|unavailable|denied)[^.\n]*model`, "agent model failure cause"},
			{`(?i)(log[ -]?in|authenticat)[^.\n]*(fail|missing|invalid|unavailable|denied)|(fail|missing|invalid|unavailable|denied)[^.\n]*(log[ -]?in|authenticat)`, "agent login failure cause"},
			{`(?i)(check|fix|correct|configure|restore)[^.\n]*(command|model|log[ -]?in|authenticat)`, "correct agent command, model or login"},
			{`(?i)(usage|rate)[^.\n]*limit|limit[^.\n]*(usage|rate)`, "usage limit cause"},
			{`(?i)(wait|resolve|restore|increase|check)[^.\n]*(limit|quota|access|usage)`, "corrective action for usage limit"},
			{`(?i)(shared|same)[^.\n]*(audit|auditor)[^.\n]*command|(audit|auditor)[^.\n]*command[^.\n]*(shared|same)`, "shared auditor command warning"},
			{`(?i)(configure|set|provide)[^.\n]*(independent|separate|distinct)[^.\n]*(audit|auditor)[^.\n]*command`, "configure an independent audit command"},
		}},
		{"C16", []struct{ pattern, description string }{
			{`(?i)missing[^.\n]*(managed[^.\n]*)?(guidance|AGENTS\.md)|(guidance|AGENTS\.md)[^.\n]*missing`, "missing managed guidance"},
			{`(?i)(stale|out.of.date)[^.\n]*(managed[^.\n]*)?(guidance|AGENTS\.md)|(guidance|AGENTS\.md)[^.\n]*(stale|out.of.date)`, "stale managed guidance"},
			{`(?i)missing[^.\n]*embedded_base|embedded_base[^.\n]*missing`, "missing playbook embedded_base"},
			{`(?i)(stale|out.of.date)[^.\n]*embedded_base|embedded_base[^.\n]*(stale|out.of.date)`, "stale playbook embedded_base"},
			{`(?i)agents-md[^.\n]*--write`, "regenerate managed guidance with agents-md --write"},
			{`(?i)(review|update|refresh|correct)[^.\n]*(playbook[^.\n]*)?(override|embedded_base)`, "review or update the playbook override"},
			{`(?i)embedded_base[^.\n]*(current|embedded)[^.\n]*playbook|(current|embedded)[^.\n]*playbook[^.\n]*embedded_base`, "compare embedded_base with the current embedded playbook"},
		}},
	} {
		t.Run(tc.code, func(t *testing.T) {
			text, err := assets.Explain(tc.code)
			if err != nil {
				t.Fatalf("Explain(%s): %v", tc.code, err)
			}
			for _, check := range tc.checks {
				requirePattern(t, text, check.pattern, check.description)
			}
		})
	}
}

func requirePattern(t *testing.T, text, pattern, description string) {
	t.Helper()
	if !regexp.MustCompile(pattern).MatchString(text) {
		t.Errorf("missing %s in %q", description, text)
	}
}
