package evals_test

import (
	"regexp"
	"strings"
	"testing"

	"github.com/lurkshark/fireweed/packages/assets"
)

const helpArg = `(?:<[^>\n]+>|[A-Z][A-Z_-]*)`
const exampleArg = `[[:alnum:]_./:~-]+`

var commandHelpCases = []struct {
	name, purpose, syntax, example string
}{
	{"help", `overview|command list`, "", ""},
	{"version", `(?:installed|binary|fireweed|current)[^\n]*version|version[^\n]*(?:installed|binary|fireweed|current)`, "", ""},
	{"init", `initializ|set up|bootstrap`, `.*--recipe\s+` + helpArg + `.*--agent\s+` + helpArg, `\s+--recipe\s+` + exampleArg + `\s+--agent\s+` + exampleArg},
	{"doctor", `canar`, "", ""},
	{"new-module", `creat\w*[^\n]*module|scaffold`, `\s+` + helpArg + `\s+--prefix\s+` + helpArg, `\s+` + exampleArg + `\s+--prefix\s+` + exampleArg},
	{"new-change", `creat\w*[^\n]*change`, `\s+` + helpArg + `\s+` + helpArg, `\s+` + exampleArg + `\s+` + exampleArg},
	{"stamp", `base[s]?\b[^\n]*hash|hash[^\n]*\bbase`, `\s+` + helpArg, `\s+` + exampleArg},
	{"list", `queue|pending changes`, "", ""},
	{"status", `blocked|quarantine|drift|budget|audit round`, "", ""},
	{"show", `spec`, `\s+` + helpArg + `\s+\[` + helpArg + `\]`, `\s+(?:--stale\s+)?` + exampleArg},
	{"modules", `dependency graph|module graph|zone map`, "", ""},
	{"context", `current step|current phase|worker`, `\s+` + helpArg, `\s+` + exampleArg},
	{"explain", `diagnostic`, `\s+` + helpArg, `\s+` + exampleArg},
	{"playbook", `instructions|guidance`, `\s+` + helpArg, `\s+` + exampleArg},
	{"check", `static|validat`, "", ""},
	{"coverage", `(?:map|match|link)[^\n]*\b(?:ID|spec)|\b(?:ID|spec)[^\n]*(?:map|match|link)`, `\s+` + helpArg, `\s+` + exampleArg},
	{"check-commits", `history|H01`, `\s+` + helpArg, `\s+` + exampleArg},
	{"ci", `deterministic|continuous integration|history[^\n]*check`, `\s+\[` + helpArg + `\]`, ""},
	{"pr-body", `pull request|PR description`, `\s+` + helpArg, `\s+` + exampleArg},
	{"start", `claim|begin`, `\s+` + helpArg, `\s+` + exampleArg},
	{"verify", `check\w*[^\n]*\b(?:step|phase|change)|\b(?:step|phase|change)[^\n]*check`, `\s+\[` + helpArg + `\]`, ""},
	{"done", `advance|commit|complete`, `\s+\[` + helpArg + `\]`, ""},
	{"next", `agent`, "", ""},
	{"run", `queue`, "", ""},
	{"release", `claim`, `\s+` + helpArg, `\s+` + exampleArg},
	{"commits", `branch|commit IDs?|commit hashes`, `\s+` + helpArg, `\s+` + exampleArg},
	{"abandon", `revert|undo|remove`, `\s+` + helpArg, `\s+` + exampleArg},
	{"drill", `rebuild|reconstruct|regenerat`, `\s+` + helpArg, `\s+` + exampleArg},
	{"gc", `prune|stale claim`, "", ""},
	{"migrate", `format`, "", ""},
	{"agents-md", `managed (?:guidance|block|section)`, `\s+--write`, `\s+--write`},
}

func Test_AST_008_Help(t *testing.T) {
	overview := assets.Help()
	requirePattern(t, overview, `(?i)\bfireweed\b`, "Fireweed name")
	commandExamples := `\bfireweed\s+(?:help\s+(?:` + helpCommandNames() + `)(?:$|[^[:alnum:]_-])|(?:` + helpCommandNames() + `)\s+--help\b)`
	listing := regexp.MustCompile(commandExamples+`|\bhelp\s+<command>`).ReplaceAllString(overview, "")
	for _, command := range commandHelpCases {
		requirePattern(t, listing, commandWord(command.name), "overview entry for "+command.name)
	}
	for _, group := range []struct{ label, commands string }{
		{`setup|getting started`, `init|doctor`},
		{`author\w*|creat\w*`, `new-module|new-change|stamp`},
		{`read\w*|inspect\w*`, `list|status|show|modules|context`},
		{`check\w*|validat\w*`, `check|coverage|check-commits|ci`},
		{`lifecycle|work\s?flow|execution`, `start|verify|done|next|run`},
	} {
		requirePattern(t, overview, `(?im)^[ \t#*]*\b(?:`+group.label+`)[^\n]*(?:\n[ \t]*){0,2}\b(?:`+group.commands+`)\b`, "commands grouped under "+group.label)
	}
	requirePattern(t, overview, `\bhelp\s+<command>`, "help <command> navigation")
	requirePattern(t, overview, `<command>\s+--help\b`, "<command> --help navigation")
	requirePattern(t, overview, commandExamples, "concrete command-help example")
	if regexp.MustCompile(`gc\.drill_max_age|720h`).MatchString(overview) {
		t.Error("Help() includes gc configuration details that belong on the gc page")
	}
}

func Test_AST_061_CommandPages(t *testing.T) {
	names := helpCommandNames()
	relatedHelp := regexp.MustCompile(`\bfireweed\s+(?:help\s+(` + names + `)(?:$|[^[:alnum:]_-])|(` + names + `)\s+(?:--help|-h|help)\b)`)
	for _, command := range commandHelpCases {
		page, err := assets.CommandHelp(command.name)
		if err != nil {
			t.Errorf("CommandHelp(%q): %v", command.name, err)
			continue
		}
		invocation := `fireweed[ \t]+` + regexp.QuoteMeta(command.name) + `\b`
		requirePattern(t, page, `(?i)`+command.purpose, command.name+" purpose")
		requirePattern(t, page, `\b`+invocation+command.syntax+`(?:$|[^[:alnum:]_-])`, command.name+" usage syntax")
		requirePattern(t, page, `(?im)^[ \t]*[-*>]?[ \t]*(?:(?:for )?example\b[:,]?[ \t]*|try[ \t]+)?[`+"\x60"+`$ \t]*`+invocation+command.example+`(?:[ \t`+"\x60"+`]|[.!]?$)`, command.name+" invocation example")
		requirePattern(t, page, `(^|[^-\w])--help\b`, command.name+" --help option")
		requirePattern(t, page, `(^|[^-\w])-h\b`, command.name+" -h option")
		requirePattern(t, page, `(?i)(?:show|display|print|read|view|return|open)[^\n]*(?:help|(?:this|command) page)`, command.name+" help option effect")
		requirePattern(t, page, `(?i)--json\b[^\n]*(?:output|format|print|emit|envelope|return|machine)|(?:output|format|print|emit|envelope|return|machine)[^\n]*--json\b`, command.name+" JSON option and its effect")
		related := false
		for _, match := range relatedHelp.FindAllStringSubmatch(page, -1) {
			if match[1]+match[2] != command.name {
				related = true
				break
			}
		}
		if !related {
			t.Errorf("CommandHelp(%q) lacks a related command-help invocation", command.name)
		}
	}
}

func Test_AST_061_CommandOptions(t *testing.T) {
	type helpOption struct{ flag, effect string }
	for _, command := range []struct {
		name    string
		options []helpOption
	}{
		{"init", []helpOption{
			{"--recipe", `recipe|template`}, {"--agent", `worker|codex|claude`},
			{"--audit-agent", `audit(?:or)? (?:agent|command)|separate|independent`}, {"--no-app", `skip|omit|without`},
		}},
		{"doctor", []helpOption{{"--no-agent", `skip|omit|disable|without`}}},
		{"new-module", []helpOption{{"--prefix", `(?:spec|ID)[ -]+prefix|prefix[^\n]*\bIDs?`}}},
		{"show", []helpOption{
			{"--stale", `base|stale change|drift|diverg`}, {"--history", `archiv|provenance|history of`},
			{"--projected", `pending|queue`},
		}},
		{"context", []helpOption{{"--full", `contents?|full text`}}},
		{"check", []helpOption{{"--fix", `(?:stamp|fill|add)[^\n]*missing[^\n]*bases`}}},
		{"coverage", []helpOption{{"--with-change", `project(?:ed|ion)`}}},
		{"ci", []helpOption{{"--no-doctor", `skip|omit|disable|without`}}},
		{"start", []helpOption{{"--from-attempt", `restor|resum|recover`}}},
		{"verify", []helpOption{{"--quick", `dependents?|boundar`}}},
		{"next", []helpOption{{"--allow-unmerged", `approv|base branch|unmerged changes`}}},
		{"run", []helpOption{
			{"--allow-unmerged", `approv|base branch|unmerged changes`}, {"--max", `changes|positive integer`},
			{"--for", `duration|time limit|time budget`},
		}},
		{"agents-md", []helpOption{{"--write", `writ\w*|updat\w*|refresh\w*`}}},
	} {
		page, err := assets.CommandHelp(command.name)
		if err != nil {
			t.Errorf("CommandHelp(%q): %v", command.name, err)
			continue
		}
		for _, option := range command.options {
			requirePattern(t, page, `(?i)`+option.flag+`\b[^\n]*\s+(?:`+option.effect+`)|(?:`+option.effect+`)[^\n]*`+option.flag+`\b`, command.name+" description of "+option.flag)
		}
	}
}

func Test_AST_061_CommandInputDetails(t *testing.T) {
	init := mustCommandHelp(t, "init")
	for _, term := range []string{"go", "typescript-pnpm", "python-uv", "none", "codex", "claude"} {
		requirePattern(t, init, commandWord(term), "init recipe or agent value "+term)
	}
	requirePattern(t, init, `(?i)local[^\n]*(?:path|directory)|(?:path|directory)[^\n]*local`, "local recipe path")
	requirePattern(t, init, `\{prompt\}`, "inline prompt placeholder")
	requirePattern(t, init, `\{prompt_file\}`, "prompt file placeholder")

	show := mustCommandHelp(t, "show")
	requirePattern(t, show, `\bshow\s+--stale\s+`+helpArg, "separate stale-change synopsis")
	requirePattern(t, show, `(?i)mutually exclusive|only one|(?:cannot|can't|do not|don't|never)\s+(?:be\s+)?combin|\[\s*--(?:projected|history)\s*\|\s*--(?:projected|history)\s*\]`, "exclusive show modes")
	requirePattern(t, show, `(?i)(?:spec|item)[^\n]*\bID\b|\bID\b[^\n]*(?:spec|item)`, "optional spec item selector")

	for _, input := range []struct{ command, meaning, description string }{
		{"new-module", `lowercase|lower-case`, "module name format"},
		{"new-change", `(?:lowercase|lower-case|unused|unique)[^\n]*slug|slug[^\n]*(?:lowercase|lower-case|unused|unique)`, "change slug format"},
		{"stamp", `pending change|queued change|change (?:file|selector)`, "change to stamp"},
		{"context", `claim|held change|active change`, "held change"},
		{"coverage", `eval hook|evals? for (?:the |a )?module|module[^\n]*evals?`, "module to evaluate"},
		{"check-commits", `(?:Git|revision|commit)[^\n]*range|range[^\n]*(?:Git|revision|commit)`, "Git revision range"},
		{"ci", `(?:Git|revision|commit)[^\n]*range|range[^\n]*(?:Git|revision|commit)`, "optional Git revision range"},
		{"pr-body", `(?:Git|revision|commit)[^\n]*range|range[^\n]*(?:Git|revision|commit)`, "Git revision range"},
		{"verify", `held change|current (?:claim|change)|(?:the |a )?change held`, "default to the held change"},
		{"done", `held change|current (?:claim|change)|(?:the |a )?change held`, "default to the held change"},
		{"abandon", `unapplied|pending`, "unapplied change selector"},
	} {
		page := mustCommandHelp(t, input.command)
		requirePattern(t, page, `(?i)`+input.meaning, input.command+" "+input.description)
	}
}

func Test_AST_061_UnknownCommandHelp(t *testing.T) {
	const unknown = "no-such-fireweed-command"
	if _, err := assets.CommandHelp(unknown); err == nil || !strings.Contains(err.Error(), unknown) {
		t.Errorf("CommandHelp(%q) must reject and name the command, got %v", unknown, err)
	}
}

func Test_AST_062_HelpNavigation(t *testing.T) {
	page := mustCommandHelp(t, "help")
	for _, form := range []string{"help <command>", "<command> help", "<command> --help", "<command> -h"} {
		requirePattern(t, page, strings.ReplaceAll(regexp.QuoteMeta(form), " ", `\s+`), form)
	}
	requirePattern(t, page, `(?i)overview|command list|lists?[^\n]*commands`, "overview behavior")
	requirePattern(t, page, `(?i)\bschema(?:\s+version)?[[:space:]"':=-]*1\b`, "schema version 1")
	for _, field := range []string{"schema", "data", "diagnostics"} {
		requirePattern(t, page, commandWord(field), "JSON envelope field "+field)
	}
	for _, status := range []struct{ code, meaning string }{
		{"0", `success[^\n]*empty queue|empty queue[^\n]*success`},
		{"1", `(?:fail\w*[^\n]*check\w*|check\w*[^\n]*fail)[^\n]*block|block[^\n]*(?:fail\w*[^\n]*check\w*|check\w*[^\n]*fail)`},
		{"2", `usage[^\n]*configur|configur[^\n]*usage`},
		{"3", `clean[^\n]*stop`},
	} {
		requirePattern(t, page, `(?i)\b`+status.code+`\b[^\n]*(?:`+status.meaning+`)|(?:`+status.meaning+`)[^\n]*\b`+status.code+`\b`, "exit status "+status.code+" meaning")
	}
}

func Test_AST_063_PlaybookAndDiagnosticDiscovery(t *testing.T) {
	playbook := mustCommandHelp(t, "playbook")
	for _, entry := range []struct{ name, purpose string }{
		{"author-change", `author|draft|writ\w*|creat\w*|propos\w*`},
		{"evals", `test\w*|assert\w*|fail.first`},
		{"audit", `review\w*|check\w*|inspect\w*`},
		{"impl", `implement\w*|disposable code`},
		{"regenerate", `rebuild\w*|reconstruct\w*`},
	} {
		requirePattern(t, playbook, `(?i)`+commandWord(entry.name)+`[^\n]*\s+(?:`+entry.purpose+`)\b`, entry.name+" playbook purpose")
	}
	requirePattern(t, playbook, `\bfireweed\s+playbook\s+(?:author-change|evals|audit|impl|regenerate)\b`, "example reading a named playbook")

	explain := mustCommandHelp(t, "explain")
	requirePattern(t, explain, `(?i)\bdiagnostic[ -]+code\b`, "diagnostic-code lookup")
	requirePattern(t, explain, `\bfireweed\s+explain\s+(?:[BCEHPUXZ][0-9]{2}|P2-[0-9]{2})\b`, "concrete diagnostic-code example")
	requirePattern(t, explain, `\bfireweed\s+(?:help\s+playbook|playbook\s+(?:--help|-h|help))\b`, "playbook help for workflow guidance")
	requirePattern(t, explain, `(?i)workflow|(?:step|phase)[ -]+(?:guidance|instructions)`, "workflow guidance through playbooks")
}

func mustCommandHelp(t *testing.T, name string) string {
	t.Helper()
	page, err := assets.CommandHelp(name)
	if err != nil {
		t.Fatalf("CommandHelp(%q): %v", name, err)
	}
	return page
}

func commandWord(name string) string {
	return `(^|[^[:alnum:]_-])` + regexp.QuoteMeta(name) + `($|[^[:alnum:]_-])`
}

func helpCommandNames() string {
	names := make([]string, 0, len(commandHelpCases))
	for _, command := range commandHelpCases {
		names = append(names, regexp.QuoteMeta(command.name))
	}
	return strings.Join(names, "|")
}
