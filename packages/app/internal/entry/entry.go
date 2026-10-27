package entry

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"github.com/lurkshark/fireweed/packages/assets"
	"github.com/lurkshark/fireweed/packages/project"
)

type diagnostic struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type envelope struct {
	Schema      int          `json:"schema"`
	Data        any          `json:"data"`
	Diagnostics []diagnostic `json:"diagnostics"`
}

var pending = map[string]bool{
	"init": true, "doctor": true,
	"new-module": true, "stamp": true,
	"list": true, "status": true, "show": true, "context": true,
	"check": true, "coverage": true, "check-commits": true, "ci": true, "pr-body": true,
	"start": true, "verify": true, "done": true, "next": true, "run": true,
	"release": true, "commits": true, "abandon": true,
	"drill": true, "gc": true, "migrate": true, "agents-md": true,
}

func Run(args []string, out, errOut io.Writer) int {
	jsonMode := false
	positional := make([]string, 0, len(args))
	for _, arg := range args {
		if arg == "--json" {
			jsonMode = true
		} else {
			positional = append(positional, arg)
		}
	}
	fail := func(code, message string) int {
		if jsonMode {
			_ = json.NewEncoder(out).Encode(envelope{Schema: 1, Diagnostics: []diagnostic{{Code: code, Message: message}}})
		} else {
			fmt.Fprintf(errOut, "%s: %s\n", code, message)
		}
		return 2
	}
	for _, arg := range positional {
		if strings.HasPrefix(arg, "-") {
			return fail("CLI_OPTION", fmt.Sprintf("Unknown option %q. Use help to see available commands and options.", arg))
		}
	}
	if len(positional) == 0 {
		return fail("CLI_USAGE", "A command is required. Use help to see available commands.")
	}
	command := positional[0]
	count := len(positional) - 1
	var body string
	var err error
	switch command {
	case "help":
		if count != 0 {
			return fail("CLI_USAGE", "help takes no arguments. Use help alone.")
		}
		body = assets.Help()
	case "version":
		if count != 0 {
			return fail("CLI_USAGE", "version takes no arguments. Use version alone.")
		}
		body = assets.Version()
	case "explain", "playbook":
		if count != 1 {
			return fail("CLI_USAGE", fmt.Sprintf("%s requires exactly one %s. Use %s <%s>.", command, lookupKind(command), command, lookupKind(command)))
		}
		if command == "explain" {
			body, err = assets.Explain(positional[1])
		} else {
			body, err = assets.Playbook(positional[1])
		}
		if err != nil {
			return fail("CLI_LOOKUP", fmt.Sprintf("%v. Use a valid %s name.", err, lookupKind(command)))
		}
	case "modules":
		if count != 0 {
			return fail("CLI_USAGE", "modules takes no arguments. Use modules alone.")
		}
		snapshot, err := loadValidatedProject()
		if err != nil {
			return fail("C01", fmt.Sprintf("%v. Provide a valid fireweed.yaml and correct module declarations in SPEC.md, then run modules again.", err))
		}
		modules := moduleResults(snapshot.Modules)
		if jsonMode {
			_ = json.NewEncoder(out).Encode(envelope{Schema: 1, Data: modules, Diagnostics: []diagnostic{}})
		} else {
			printModules(out, modules)
		}
		return 0
	case "new-change":
		if count != 2 {
			return fail("CLI_USAGE", "new-change requires a module and slug. Use new-change <module> <slug>.")
		}
		snapshot, err := loadValidatedProject()
		if err != nil {
			return fail("C01", fmt.Sprintf("%v. Provide a valid fireweed.yaml and correct module declarations in SPEC.md, then retry.", err))
		}
		path, err := project.NewChange(snapshot.Root, positional[1], positional[2])
		if err != nil {
			return fail("CLI_AUTHOR", fmt.Sprintf("%v. Use a valid module and an unused lowercase slug.", err))
		}
		if jsonMode {
			_ = json.NewEncoder(out).Encode(envelope{Schema: 1, Data: path, Diagnostics: []diagnostic{}})
		} else {
			fmt.Fprintln(out, path)
		}
		return 0
	default:
		if pending[command] {
			return fail("CLI_UNIMPLEMENTED", fmt.Sprintf("%s is under construction and is not implemented yet. Use help for available information commands.", command))
		}
		return fail("CLI_COMMAND", fmt.Sprintf("Unknown command %q. Use help to see available commands.", command))
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
