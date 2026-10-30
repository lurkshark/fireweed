package entry

import (
	"encoding/json"
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"

	"github.com/lurkshark/fireweed/packages/assets"
	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/project"
)

type diagnostic struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Path    string `json:"path,omitempty"`
	Warning bool   `json:"warning"`
}

type envelope struct {
	Schema      int          `json:"schema"`
	Data        any          `json:"data"`
	Diagnostics []diagnostic `json:"diagnostics"`
}

var pending = map[string]bool{
	"list": true, "status": true,
	"ci": true, "pr-body": true,
	"drill": true, "gc": true, "migrate": true,
}

func Run(args []string, out, errOut io.Writer) int {
	if isInitRequest(args) {
		if contains(args, "--for") {
			return cliFailure(2, contains(args, "--json"), out, errOut, "U01", "--for is only available with run. Use run --for 1h.")
		}
		return runInit(args, out, errOut)
	}
	if isNewModuleRequest(args) {
		return runNewModule(args, out, errOut)
	}
	jsonMode := false
	projected := false
	quick := false
	full := false
	noAgent := false
	write := false
	allowUnmerged := false
	max := 0
	maxSet := false
	forDuration := time.Duration(0)
	forSet := false
	fromAttempt := 0
	fromAttemptSet := false
	withChange := ""
	positional := make([]string, 0, len(args))
	for i := 0; i < len(args); i++ {
		arg := args[i]
		if arg == "--json" {
			jsonMode = true
		} else if arg == "--projected" {
			projected = true
		} else if arg == "--quick" {
			quick = true
		} else if arg == "--full" {
			full = true
		} else if arg == "--no-agent" {
			noAgent = true
		} else if arg == "--write" {
			write = true
		} else if arg == "--allow-unmerged" {
			if allowUnmerged {
				return cliFailure(2, jsonMode || contains(args, "--json"), out, errOut, "U01", "--allow-unmerged may be specified once. Use next or run with --allow-unmerged.")
			}
			allowUnmerged = true
		} else if arg == "--max" {
			if maxSet || i+1 >= len(args) {
				return cliFailure(2, jsonMode || contains(args, "--json"), out, errOut, "U01", "--max requires one positive integer. Use run --max N.")
			}
			i++
			value, parseErr := strconv.Atoi(args[i])
			if parseErr != nil || value <= 0 {
				return cliFailure(2, jsonMode || contains(args, "--json"), out, errOut, "U01", "--max requires a positive integer. Use run --max N.")
			}
			max, maxSet = value, true
		} else if arg == "--for" {
			if forSet || i+1 >= len(args) {
				return cliFailure(2, jsonMode || contains(args, "--json"), out, errOut, "U01", "--for requires one positive Go duration. Use run --for 1h.")
			}
			i++
			value, parseErr := time.ParseDuration(args[i])
			if parseErr != nil || value <= 0 {
				return cliFailure(2, jsonMode || contains(args, "--json"), out, errOut, "U01", "--for requires one positive Go duration. Use run --for 1h.")
			}
			forDuration, forSet = value, true
		} else if arg == "--from-attempt" {
			if fromAttemptSet || i+1 >= len(args) {
				return cliFailure(2, jsonMode || contains(args, "--json"), out, errOut, "U01", "--from-attempt requires one positive integer. Use start <change> --from-attempt N.")
			}
			i++
			value, parseErr := strconv.Atoi(args[i])
			if parseErr != nil || value <= 0 {
				return cliFailure(2, jsonMode || contains(args, "--json"), out, errOut, "U01", "--from-attempt requires one positive integer. Use start <change> --from-attempt N.")
			}
			fromAttempt, fromAttemptSet = value, true
		} else if arg == "--with-change" {
			if i+1 >= len(args) || strings.HasPrefix(args[i+1], "-") {
				return cliFailure(2, jsonMode || contains(args, "--json"), out, errOut, "U01", "--with-change requires a change value. Use coverage <module> --with-change <change>.")
			}
			i++
			withChange = args[i]
		} else {
			positional = append(positional, arg)
		}
	}
	failStatus := func(status int, code, message string) int {
		if jsonMode {
			_ = json.NewEncoder(out).Encode(envelope{Schema: 1, Diagnostics: []diagnostic{{Code: code, Message: message}}})
		} else {
			fmt.Fprintf(errOut, "%s: %s\n", code, message)
		}
		return status
	}
	fail := func(code, message string) int {
		return failStatus(2, code, message)
	}
	for _, arg := range positional {
		if strings.HasPrefix(arg, "-") {
			return fail("U01", fmt.Sprintf("Unknown option %q. Use help to see available commands and options.", arg))
		}
	}
	if len(positional) == 0 {
		return fail("U01", "A command is required. Use help to see available commands.")
	}
	command := positional[0]
	if projected && command != "show" {
		return fail("U01", "--projected is only available with show.")
	}
	if quick && command != "verify" || full && command != "context" || withChange != "" && command != "coverage" {
		return fail("U01", "Use --quick only with verify, --full only with context, and --with-change only with coverage.")
	}
	if noAgent && command != "doctor" || write && command != "agents-md" {
		return fail("U01", "Use --no-agent only with doctor and --write only with agents-md.")
	}
	if allowUnmerged && command != "next" && command != "run" {
		return fail("U01", "--allow-unmerged is only available with next or run.")
	}
	if maxSet && command != "run" {
		return fail("U01", "--max is only available with run. Use run --max N.")
	}
	if forSet && command != "run" {
		return fail("U01", "--for is only available with run. Use run --for 1h.")
	}
	if fromAttemptSet && command != "start" {
		return fail("U01", "--from-attempt is only available with start. Use start <change> --from-attempt N.")
	}
	count := len(positional) - 1
	var body string
	var err error
	switch command {
	case "help":
		if count != 0 {
			return fail("U01", "help takes no arguments. Use help alone.")
		}
		body = assets.Help()
	case "version":
		if count != 0 {
			return fail("U01", "version takes no arguments. Use version alone.")
		}
		body = assets.Version()
	case "explain", "playbook":
		if count != 1 {
			return fail("U01", fmt.Sprintf("%s requires exactly one %s. Use %s <%s>.", command, lookupKind(command), command, lookupKind(command)))
		}
		if command == "explain" {
			body, err = assets.Explain(positional[1])
		} else {
			body, err = assets.Playbook(positional[1])
		}
		if err != nil {
			return fail("U02", fmt.Sprintf("%v. Use a valid %s name.", err, lookupKind(command)))
		}
	case "modules":
		if count != 0 {
			return fail("U01", "modules takes no arguments. Use modules alone.")
		}
		snapshot, err := loadModuleProject()
		if err != nil {
			return fail(projectFailureCode(err, "C01"), fmt.Sprintf("%v. Provide a valid fireweed.yaml and correct module declarations in SPEC.md, then run modules again.", err))
		}
		modules := moduleResults(snapshot.Modules)
		if jsonMode {
			_ = json.NewEncoder(out).Encode(envelope{Schema: 1, Data: modules, Diagnostics: []diagnostic{}})
		} else {
			printModules(out, modules)
		}
		return 0
	case "start", "verify", "done", "context", "coverage", "release", "commits", "abandon":
		return runLifecycle(command, positional[1:], jsonMode, quick, full, withChange, fromAttempt, out, errOut)
	case "check-commits":
		return runCheckCommits(positional[1:], jsonMode, out, errOut)
	case "next":
		if count != 0 {
			return fail("U01", "next takes no arguments. Use next [--allow-unmerged] [--json].")
		}
		return runNext(allowUnmerged, jsonMode, out, errOut)
	case "run":
		if count != 0 {
			return fail("U01", "run takes no arguments. Use run [--for 1h] [--max N] [--allow-unmerged] [--json].")
		}
		return runQueue(max, forDuration, allowUnmerged, jsonMode, out, errOut)
	case "doctor":
		if count != 0 {
			return fail("U01", "doctor takes no arguments. Use doctor [--no-agent] [--json].")
		}
		return runDoctor(noAgent, jsonMode, out, errOut)
	case "agents-md":
		if count != 0 || !write {
			return fail("U01", "agents-md requires --write and no arguments. Use agents-md --write [--json].")
		}
		return runAgentsMD(jsonMode, out, errOut)
	case "new-change":
		if count != 2 {
			return fail("U01", "new-change requires a module and slug. Use new-change <module> <slug>.")
		}
		snapshot, err := loadValidatedProject()
		if err != nil {
			return fail(projectFailureCode(err, "C01"), fmt.Sprintf("%v. Provide a valid fireweed.yaml and correct module declarations in SPEC.md, then retry.", err))
		}
		path, err := project.NewChange(snapshot.Root, positional[1], positional[2])
		if err != nil {
			return fail(newChangeFailureCode(err), fmt.Sprintf("%v. Use a valid module and an unused lowercase slug.", err))
		}
		if jsonMode {
			_ = json.NewEncoder(out).Encode(envelope{Schema: 1, Data: path, Diagnostics: []diagnostic{}})
		} else {
			fmt.Fprintln(out, path)
		}
		return 0
	case "stamp":
		if count != 1 {
			return fail("U01", "stamp requires exactly one change. Use stamp <change>.")
		}
		snapshot, err := loadValidatedProject()
		if err != nil {
			return fail(projectFailureCode(err, "C01"), fmt.Sprintf("%v. Provide a valid fireweed.yaml and module declarations, then retry.", err))
		}
		if err := project.Stamp(snapshot.Root, positional[1]); err != nil {
			code := projectFailureCode(err, "P01")
			if strings.Contains(err.Error(), " bases must be a mapping") {
				code = "C08"
			}
			return failStatus(1, code, err.Error())
		}
		body = fmt.Sprintf("Stamped %s\n", positional[1])
	case "check":
		if count != 0 {
			return fail("U01", "check takes no arguments. Use check alone.")
		}
		snapshot, err := loadValidatedProject()
		if err != nil {
			return fail(projectFailureCode(err, "C01"), fmt.Sprintf("%v. Provide a valid fireweed.yaml and module declarations, then retry.", err))
		}
		results := project.Check(snapshot.Root)
		diagnostics := make([]diagnostic, 0, len(results))
		hasErrors := false
		for _, result := range results {
			diagnostics = append(diagnostics, diagnostic{Code: result.Code, Message: result.Message, Path: result.Path, Warning: result.Warning})
			if !result.Warning {
				hasErrors = true
			}
		}
		if jsonMode {
			var data any
			if !hasErrors {
				data = "Check passed\n"
			}
			_ = json.NewEncoder(out).Encode(envelope{Schema: 1, Data: data, Diagnostics: diagnostics})
		} else {
			for _, result := range diagnostics {
				fmt.Fprintf(errOut, "%s: %s\n", result.Code, result.Message)
			}
			if !hasErrors {
				fmt.Fprintln(out, "Check passed")
			}
		}
		if hasErrors {
			return 1
		}
		return 0
	case "show":
		if count < 1 || count > 2 {
			return fail("U01", "show requires a module and optional item ID. Use show <module> [ID] [--projected].")
		}
		snapshot, err := loadValidatedProject()
		if err != nil {
			code := projectFailureCode(err, "C01")
			return fail(code, fmt.Sprintf("%v. Provide a valid fireweed.yaml and module declarations, then retry.", err))
		}
		if projected {
			snapshot, err = project.Project(snapshot.Root, "")
			if err != nil {
				return fail(projectFailureCode(err, "P01"), err.Error())
			}
		}
		var found *project.Module
		for i := range snapshot.Modules {
			if snapshot.Modules[i].Name == positional[1] {
				if found != nil {
					return fail("U02", fmt.Sprintf("ambiguous module %q", positional[1]))
				}
				found = &snapshot.Modules[i]
			}
		}
		if found == nil {
			return fail("U02", fmt.Sprintf("unknown module %q", positional[1]))
		}
		if count == 1 {
			body, err = contracts.RenderDocument(found.Document)
			if err != nil {
				return fail("C03", err.Error())
			}
		} else {
			section, _ := contracts.Section(found.Document.Body, "Specs")
			items, err := contracts.ParseItems(section)
			if err != nil {
				if strings.HasPrefix(err.Error(), "duplicate item ID: ") {
					if strings.HasPrefix(err.Error(), "duplicate item ID: new:") {
						return fail("U02", err.Error())
					}
					return fail("C05", err.Error())
				}
				return fail("C03", err.Error())
			}
			for _, item := range items {
				if item.ID == positional[2] {
					body = contracts.RenderItem(item)
					break
				}
			}
			if body == "" {
				return fail("U02", fmt.Sprintf("unknown item %q in module %q", positional[2], positional[1]))
			}
		}
	default:
		if pending[command] {
			return fail("U01", fmt.Sprintf("%s is under construction and is not implemented yet. Use help for available information commands.", command))
		}
		return fail("U01", fmt.Sprintf("Unknown command %q. Use help to see available commands.", command))
	}
	if !strings.HasSuffix(body, "\n") {
		body += "\n"
	}
	if jsonMode {
		_ = json.NewEncoder(out).Encode(envelope{Schema: 1, Data: &body, Diagnostics: []diagnostic{}})
	} else {
		_, _ = io.WriteString(out, body)
	}
	return 0
}

func lookupKind(command string) string {
	if command == "explain" {
		return "code"
	}
	return "playbook"
}
