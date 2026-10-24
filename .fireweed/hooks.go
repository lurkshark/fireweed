// Protected, project-owned Go hook. No third-party dependencies.
package main

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"go/parser"
	"go/token"
	"io"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

type Case struct {
	Name    string  `xml:"name,attr"`
	Class   string  `xml:"classname,attr"`
	Failure *string `xml:"failure,omitempty"`
	Error   *string `xml:"error,omitempty"`
	Skipped *string `xml:"skipped,omitempty"`
}
type Suite struct {
	XMLName xml.Name `xml:"testsuite"`
	Cases   []Case   `xml:"testcase"`
}

func main() {
	if len(os.Args) < 2 {
		os.Exit(2)
	}
	if os.Args[1] == "eval" {
		runEvals()
		return
	}

	bin := os.Getenv("FIREWEED_BIN")
	if bin == "" {
		bin = "fireweed"
	}
	cmd := exec.Command(bin, "modules", "--json")
	data, err := cmd.Output()
	if err != nil {
		panic(err)
	}
	var envelope struct {
		Data []struct {
			Name, Dir string
			DependsOn []string `json:"depends_on"`
		}
	}
	if err = json.Unmarshal(data, &envelope); err != nil {
		panic(err)
	}
	mod, err := os.ReadFile("go.mod")
	if err != nil {
		panic(err)
	}
	modulePath := ""
	for _, line := range strings.Split(string(mod), "\n") {
		f := strings.Fields(line)
		if len(f) == 2 && f[0] == "module" {
			modulePath = strings.Trim(f[1], "\"")
		}
	}
	violations := 0
	for _, m := range envelope.Data {
		err = filepath.WalkDir(m.Dir, func(file string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}
			if d.IsDir() || !strings.HasSuffix(file, ".go") {
				return nil
			}
			tree, err := parser.ParseFile(token.NewFileSet(), file, nil, parser.ImportsOnly)
			if err != nil {
				return err
			}
			for _, imp := range tree.Imports {
				imported, _ := strconv.Unquote(imp.Path.Value)
				for _, target := range envelope.Data {
					base := modulePath + "/" + target.Dir
					if imported != base && !strings.HasPrefix(imported, base+"/") {
						continue
					}
					allowed := m.Name == target.Name
					for _, dep := range m.DependsOn {
						if dep == target.Name {
							allowed = true
						}
					}
					if !allowed || (strings.Contains(filepath.ToSlash(file), "/evals/") && imported != base) || (m.Name != target.Name && imported != base) {
						fmt.Printf("%s: %s imports %s outside its declared interface/dependencies\n", m.Name, file, imported)
						violations++
					}
				}
			}
			return nil
		})
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(2)
		}
	}
	if violations > 0 {
		os.Exit(1)
	}
}

// Collection failures are never ID-linked fail-first evidence. In particular,
// a testcase that started but never completed must not serialize as passing.
func runEvals() {
	if len(os.Args) != 4 {
		os.Exit(2)
	}
	cmd := exec.Command("go", "test", "-timeout=45m", "-json", "./"+os.Args[2]+"/...")
	var output bytes.Buffer
	cmd.Stdout = &output
	cmd.Stderr = os.Stderr
	commandErr := cmd.Run()
	dec := json.NewDecoder(&output)
	suite := Suite{}
	indexes := map[string]int{}
	logs := map[string]string{}
	pending := map[string]bool{}
	packages := map[string]bool{}
	failedCases := map[string]bool{}
	packageFailures := map[string]bool{}
	collection := []string{}
	seenPackage := false
	for {
		var ev struct{ Action, Package, Test, Output string }
		err := dec.Decode(&ev)
		if err == io.EOF {
			break
		}
		if err != nil {
			collection = append(collection, "Invalid Go test JSON: "+err.Error())
			break
		}
		if ev.Package != "" {
			seenPackage = true
		}
		if ev.Test == "" {
			if ev.Output != "" {
				fmt.Print(ev.Output)
				logs[ev.Package] += ev.Output
			}
			switch ev.Action {
			case "start":
				packages[ev.Package] = true
			case "pass", "skip":
				delete(packages, ev.Package)
			case "fail":
				delete(packages, ev.Package)
				packageFailures[ev.Package] = true
			}
			continue
		}
		key := ev.Package + "/" + ev.Test
		if ev.Action == "run" {
			if _, exists := indexes[key]; exists {
				collection = append(collection, "Duplicate testcase start: "+key)
			}
			indexes[key] = len(suite.Cases)
			suite.Cases = append(suite.Cases, Case{Name: ev.Test, Class: ev.Package})
			pending[key] = true
		}
		logs[key] += ev.Output
		switch ev.Action {
		case "pass", "fail", "skip":
			i, exists := indexes[key]
			if !exists || !pending[key] {
				collection = append(collection, "Testcase terminal event without one active start: "+key)
				continue
			}
			delete(pending, key)
			text := logs[key]
			if ev.Action == "fail" {
				suite.Cases[i].Failure = &text
				failedCases[ev.Package] = true
			}
			if ev.Action == "skip" {
				suite.Cases[i].Skipped = &text
			}
		}
	}
	if !seenPackage {
		collection = append(collection, "Go test produced no package collection events")
	}
	for key := range pending {
		collection = append(collection, "Unfinished testcase: "+key+"\n"+logs[key])
	}
	for pkg := range packages {
		collection = append(collection, "Unfinished package: "+pkg+"\n"+logs[pkg])
	}
	for pkg := range packageFailures {
		if !failedCases[pkg] {
			collection = append(collection, "Package failed without a completed failing testcase: "+pkg+"\n"+logs[pkg])
		}
	}
	if commandErr != nil && len(packageFailures) == 0 && len(collection) == 0 {
		collection = append(collection, "Go test execution failed: "+commandErr.Error()+"\n"+output.String())
	}
	if len(indexes) == 0 && len(packageFailures) == 0 && commandErr != nil {
		collection = append(collection, "No testcase collection completed")
	}
	if len(collection) > 0 {
		sort.Strings(collection)
		text := strings.Join(collection, "\n\n")
		suite.Cases = append(suite.Cases, Case{Name: "fireweed Go collection error", Class: "fireweed.reporter", Error: &text})
	}
	b, err := xml.Marshal(suite)
	if err != nil {
		panic(err)
	}
	if err = os.WriteFile(os.Args[3], b, 0644); err != nil {
		panic(err)
	}
	if commandErr != nil || len(collection) > 0 {
		os.Exit(1)
	}
}
