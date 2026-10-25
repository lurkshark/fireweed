// Protected, project-owned Go hook. No third-party dependencies.
package main

import (
	"bytes"
	"context"
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
	"regexp"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
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
	if os.Args[1] == "eval-one" {
		runOneEvals()
		return
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

type goListEvent struct{ Action, Package, Test, Output string }
type rootInventory map[string]map[string]bool

func enumerateRoots(scope string) (rootInventory, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()
	cmd := exec.CommandContext(ctx, "go", "test", "-json", "-list", ".", scope)
	var out, stderr bytes.Buffer
	cmd.Stdout, cmd.Stderr = &out, &stderr
	runErr := cmd.Run()
	inv := rootInventory{}
	active, seen := map[string]bool{}, map[string]bool{}
	dec := json.NewDecoder(&out)
	for {
		var e goListEvent
		err := dec.Decode(&e)
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("invalid Go listing JSON: %w", err)
		}
		if e.Test != "" {
			return nil, fmt.Errorf("listing unexpectedly executed testcase %s", e.Test)
		}
		switch e.Action {
		case "start":
			if seen[e.Package] {
				return nil, fmt.Errorf("duplicate listed package %s", e.Package)
			}
			if e.Package == "" {
				return nil, fmt.Errorf("listing has unnamed package")
			}
			seen[e.Package], active[e.Package] = true, true
			inv[e.Package] = map[string]bool{}
		case "pass", "skip":
			if !active[e.Package] {
				return nil, fmt.Errorf("listed package terminal without start: %s", e.Package)
			}
			delete(active, e.Package)
		case "fail":
			return nil, fmt.Errorf("Go package listing failed: %s\n%s", e.Package, stderr.String())
		}
		if e.Output != "" {
			for _, line := range strings.Split(e.Output, "\n") {
				name := strings.TrimSpace(line)
				if !token.IsIdentifier(name) || !(strings.HasPrefix(name, "Test") || strings.HasPrefix(name, "Example") || strings.HasPrefix(name, "Fuzz")) {
					continue
				}
				if !active[e.Package] {
					return nil, fmt.Errorf("listed root outside active package: %s", name)
				}
				if inv[e.Package][name] {
					return nil, fmt.Errorf("duplicate listed root %s/%s", e.Package, name)
				}
				inv[e.Package][name] = true
			}
		}
	}
	if runErr != nil {
		return nil, fmt.Errorf("Go listing failed: %w\n%s", runErr, stderr.String())
	}
	if len(inv) == 0 || len(active) != 0 {
		return nil, fmt.Errorf("incomplete Go package listing")
	}
	return inv, nil
}

func reportProblem(suite *Suite, problem string) {
	suite.Cases = append(suite.Cases, Case{Name: "fireweed Go collection error", Class: "fireweed.reporter", Error: &problem})
}

func writeSuite(suite Suite, report string, failed bool) {
	raw, err := xml.Marshal(suite)
	if err == nil {
		err = os.WriteFile(report, raw, 0644)
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}
	if failed {
		os.Exit(1)
	}
}

func runEvals() {
	if len(os.Args) != 4 {
		os.Exit(2)
	}
	inv, err := enumerateRoots("./" + os.Args[2] + "/...")
	if err != nil {
		suite := Suite{}
		reportProblem(&suite, err.Error())
		writeSuite(suite, os.Args[3], true)
		return
	}
	unique := map[string]bool{}
	for _, roots := range inv {
		for root := range roots {
			unique[root] = true
		}
	}
	roots := []string{}
	for root := range unique {
		roots = append(roots, root)
	}
	sort.Strings(roots)
	count := 2
	if len(roots) < 2 {
		count = 1
	}
	groups := make([][]string, count)
	bucket := map[string]int{}
	for i, root := range roots {
		group := i % count
		bucket[root] = group
		groups[group] = append(groups[group], regexp.QuoteMeta(root))
	}
	temp, err := os.MkdirTemp("", "fireweed-eval-shards-")
	if err != nil {
		suite := Suite{}
		reportProblem(&suite, err.Error())
		writeSuite(suite, os.Args[3], true)
		return
	}
	exe, err := os.Executable()
	if err != nil {
		os.RemoveAll(temp)
		suite := Suite{}
		reportProblem(&suite, err.Error())
		writeSuite(suite, os.Args[3], true)
		return
	}
	type shard struct {
		suite          Suite
		packages       map[string]bool
		output, stderr string
		err            error
	}
	results := make([]shard, count)
	var wg sync.WaitGroup
	for i := 0; i < count; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			pattern := "^$"
			if len(groups[i]) > 0 {
				pattern = "^(" + strings.Join(groups[i], "|") + ")$"
			}
			report := filepath.Join(temp, strconv.Itoa(i)+".xml")
			ctx, cancel := context.WithTimeout(context.Background(), 85*time.Minute)
			defer cancel()
			cmd := exec.CommandContext(ctx, exe, "eval-one", os.Args[2], report, pattern)
			var out, stderr bytes.Buffer
			cmd.Stdout, cmd.Stderr = &out, &stderr
			result := shard{err: cmd.Run(), output: out.String(), stderr: stderr.String()}
			raw, readErr := os.ReadFile(report)
			if readErr != nil {
				result.err = fmt.Errorf("missing shard report: %w", readErr)
			} else if parseErr := xml.Unmarshal(raw, &result.suite); parseErr != nil {
				result.err = fmt.Errorf("invalid shard report: %w", parseErr)
			}
			raw, readErr = os.ReadFile(report + ".packages")
			if readErr != nil {
				result.err = fmt.Errorf("missing shard package collection: %w", readErr)
			} else if parseErr := json.Unmarshal(raw, &result.packages); parseErr != nil {
				result.err = fmt.Errorf("invalid shard package collection: %w", parseErr)
			}
			results[i] = result
		}(i)
	}
	wg.Wait()
	os.RemoveAll(temp)
	suite := Suite{}
	failed := false
	allCases := map[string]bool{}
	for i, result := range results {
		fmt.Print(result.output)
		fmt.Fprint(os.Stderr, result.stderr)
		problems := []string{}
		if len(result.packages) != len(inv) {
			problems = append(problems, "shard package inventory differs")
		}
		for pkg := range inv {
			if !result.packages[pkg] {
				problems = append(problems, "shard omitted package "+pkg)
			}
		}
		observed := rootInventory{}
		legitimateFailure := false
		for _, c := range result.suite.Cases {
			key := c.Class + "/" + c.Name
			if allCases[key] {
				problems = append(problems, "duplicate testcase across shards: "+key)
			}
			allCases[key] = true
			suite.Cases = append(suite.Cases, c)
			if c.Failure != nil || c.Error != nil {
				failed = true
				legitimateFailure = true
			}
			if c.Class == "fireweed.reporter" {
				continue
			}
			if !strings.Contains(c.Name, "/") {
				if observed[c.Class] == nil {
					observed[c.Class] = map[string]bool{}
				}
				observed[c.Class][c.Name] = true
				if !inv[c.Class][c.Name] || bucket[c.Name] != i {
					problems = append(problems, "unexpected shard root "+key)
				}
			}
		}
		for pkg, expected := range inv {
			for root := range expected {
				if bucket[root] == i && !observed[pkg][root] {
					problems = append(problems, "shard omitted root "+pkg+"/"+root)
				}
			}
		}
		if result.err != nil {
			failed = true
			if !legitimateFailure {
				problems = append(problems, "shard execution failed: "+result.err.Error())
			}
		}
		if len(problems) > 0 {
			sort.Strings(problems)
			reportProblem(&suite, strings.Join(problems, "\n"))
			failed = true
		}
	}
	writeSuite(suite, os.Args[3], failed)
}

func runOneEvals() {
	if len(os.Args) != 5 {
		os.Exit(2)
	}
	cmd := exec.Command("go", "test", "-timeout=45m", "-json", "./"+os.Args[2]+"/...", "-run", os.Args[4])
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
	seenStarts := map[string]bool{}
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
				if seenStarts[ev.Package] {
					collection = append(collection, "Duplicate package start: "+ev.Package)
				}
				seenStarts[ev.Package] = true
				packages[ev.Package] = true
			case "pass", "skip":
				if !packages[ev.Package] {
					collection = append(collection, "Package terminal without active start: "+ev.Package)
				}
				delete(packages, ev.Package)
			case "fail":
				if !packages[ev.Package] {
					collection = append(collection, "Package terminal without active start: "+ev.Package)
				}
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
	for pkg := range seenStarts {
		for _, line := range strings.Split(logs[pkg], "\n") {
			text := strings.TrimSpace(line)
			if strings.HasPrefix(text, "exit status ") {
				code, err := strconv.Atoi(strings.TrimPrefix(text, "exit status "))
				if err == nil && code > 1 {
					collection = append(collection, "Abnormal test process exit: "+pkg+": "+text)
				}
			}
		}
	}
	if len(collection) > 0 {
		sort.Strings(collection)
		text := strings.Join(collection, "\n\n")
		suite.Cases = append(suite.Cases, Case{Name: "fireweed Go collection error", Class: "fireweed.reporter", Error: &text})
	}
	meta, err := json.Marshal(seenStarts)
	if err != nil {
		panic(err)
	}
	if err = os.WriteFile(os.Args[3]+".packages", meta, 0600); err != nil {
		panic(err)
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
