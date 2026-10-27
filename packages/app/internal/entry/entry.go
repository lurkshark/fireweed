package entry

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"github.com/lurkshark/fireweed/packages/assets"
	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/project"
)

type diagnostic struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Path    string `json:"path,omitempty"`
	Warning bool   `json:"warning,omitempty"`
}

type envelope struct {
	Schema      int          `json:"schema"`
	Data        any          `json:"data"`
	Diagnostics []diagnostic `json:"diagnostics"`
}

var pending = map[string]bool{
	"init": true, "doctor": true,
	"new-module": true,
	"list":       true, "status": true, "context": true,
	"coverage": true, "check-commits": true, "ci": true, "pr-body": true,
	"start": true, "verify": true, "done": true, "next": true, "run": true,
	"release": true, "commits": true, "abandon": true,
	"drill": true, "gc": true, "migrate": true, "agents-md": true,
}

func Run(args []string, out, errOut io.Writer) int {
	jsonMode := false
	projected := false
	positional := make([]string, 0, len(args))
	for _, arg := range args {
		if arg == "--json" {
			jsonMode = true
		} else if arg == "--projected" {
			projected = true
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
			return fail("CLI_OPTION", fmt.Sprintf("Unknown option %q. Use help to see available commands and options.", arg))
		}
	}
	if len(positional) == 0 {
		return fail("CLI_USAGE", "A command is required. Use help to see available commands.")
	}
	command := positional[0]
	if projected && command != "show" {
		return fail("CLI_OPTION", "--projected is only available with show.")
	}
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
	case "stamp":
		if count != 1 {
			return fail("CLI_USAGE", "stamp requires exactly one change. Use stamp <change>.")
		}
		snapshot, err := loadValidatedProject()
		if err != nil {
			return fail("C01", fmt.Sprintf("%v. Provide a valid fireweed.yaml and module declarations, then retry.", err))
		}
		if err := project.Stamp(snapshot.Root, positional[1]); err != nil {
			return failStatus(1, "CLI_STAMP", err.Error())
		}
		body = fmt.Sprintf("Stamped %s\n", positional[1])
	case "check":
		if count != 0 {
			return fail("CLI_USAGE", "check takes no arguments. Use check alone.")
		}
		snapshot, err := loadValidatedProject()
		if err != nil {
			return fail("C01", fmt.Sprintf("%v. Provide a valid fireweed.yaml and module declarations, then retry.", err))
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
			return fail("CLI_USAGE", "show requires a module and optional item ID. Use show <module> [ID] [--projected].")
		}
		snapshot, err := loadValidatedProject()
		if err != nil {
			return fail("C01", fmt.Sprintf("%v. Provide a valid fireweed.yaml and module declarations, then retry.", err))
		}
		if projected {
			snapshot, err = project.Project(snapshot.Root, "")
			if err != nil {
				return fail("CLI_SHOW", err.Error())
			}
		}
		var found *project.Module
		for i := range snapshot.Modules {
			if snapshot.Modules[i].Name == positional[1] {
				if found != nil {
					return fail("CLI_SHOW", fmt.Sprintf("ambiguous module %q", positional[1]))
				}
				found = &snapshot.Modules[i]
			}
		}
		if found == nil {
			return fail("CLI_SHOW", fmt.Sprintf("unknown module %q", positional[1]))
		}
		if count == 1 {
			body, err = contracts.RenderDocument(found.Document)
			if err != nil {
				return fail("CLI_SHOW", err.Error())
			}
		} else {
			section, _ := contracts.Section(found.Document.Body, "Specs")
			items, err := contracts.ParseItems(section)
			if err != nil {
				return fail("CLI_SHOW", err.Error())
			}
			for _, item := range items {
				if item.ID == positional[2] {
					body = contracts.RenderItem(item)
					break
				}
			}
			if body == "" {
				return fail("CLI_SHOW", fmt.Sprintf("unknown item %q in module %q", positional[2], positional[1]))
			}
		}
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
