// Go project hooks copied by the embedded recipe. Uses the standard library only.
package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"go/parser"
	"go/token"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

type testCase struct {
	Name      string  `xml:"name,attr"`
	Classname string  `xml:"classname,attr"`
	Time      string  `xml:"time,attr"`
	Failure   *string `xml:"failure,omitempty"`
	Skipped   *string `xml:"skipped,omitempty"`
}
type testSuite struct {
	XMLName xml.Name   `xml:"testsuite"`
	Cases   []testCase `xml:"testcase"`
}
type testEvent struct {
	Action  string
	Package string
	Test    string
	Output  string
	Elapsed float64
}
type module struct {
	Name      string   `json:"name"`
	Dir       string   `json:"dir"`
	DependsOn []string `json:"depends_on"`
	Zones     struct {
		Interface []string `json:"interface"`
		Evals     []string `json:"evals"`
	} `json:"zones"`
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "usage: hooks.go eval <module-dir> <report> | boundaries")
		os.Exit(2)
	}
	var err error
	switch os.Args[1] {
	case "eval":
		if len(os.Args) != 4 {
			os.Exit(2)
		}
		err = eval(os.Args[2], os.Args[3])
	case "boundaries":
		err = boundaries()
	default:
		err = fmt.Errorf("unknown hook %q", os.Args[1])
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func eval(dir, report string) error {
	cmd := exec.Command("go", "test", "-json", "./"+filepath.ToSlash(filepath.Clean(dir))+"/...")
	var stdout, stderr bytes.Buffer
	cmd.Stdout, cmd.Stderr = &stdout, &stderr
	runErr := cmd.Run()
	var suite testSuite
	index := map[string]int{}
	logs := map[string]string{}
	finished := map[string]bool{}
	packageFailed := map[string]bool{}
	packageHasFailure := map[string]bool{}
	dec := json.NewDecoder(&stdout)
	for {
		var event testEvent
		if err := dec.Decode(&event); err != nil {
			break
		}
		if event.Test == "" {
			if event.Output != "" {
				fmt.Print(event.Output)
			}
			if event.Action == "fail" {
				packageFailed[event.Package] = true
			}
			continue
		}
		key := event.Package + "/" + event.Test
		if event.Action == "run" {
			index[key] = len(suite.Cases)
			suite.Cases = append(suite.Cases, testCase{Name: event.Test, Classname: event.Package, Time: "0"})
		}
		logs[key] += event.Output
		if i, ok := index[key]; ok {
			switch event.Action {
			case "pass", "fail", "skip":
				finished[key] = true
				suite.Cases[i].Time = strconv.FormatFloat(event.Elapsed, 'f', -1, 64)
				if event.Action == "fail" {
					detail := logs[key]
					suite.Cases[i].Failure = &detail
					packageHasFailure[event.Package] = true
				}
				if event.Action == "skip" {
					detail := logs[key]
					suite.Cases[i].Skipped = &detail
				}
			}
		}
	}
	for key, i := range index {
		if !finished[key] {
			detail := logs[key] + "\nfireweed: test started but did not finish"
			suite.Cases[i].Failure = &detail
			packageHasFailure[suite.Cases[i].Classname] = true
		}
	}
	for pkg := range packageFailed {
		if !packageHasFailure[pkg] {
			detail := "fireweed: package or compile failure\n" + stderr.String()
			suite.Cases = append(suite.Cases, testCase{Name: "PackageFailure", Classname: pkg, Time: "0", Failure: &detail})
		}
	}
	if runErr != nil && !hasFailure(suite.Cases) {
		detail := "fireweed: go test failed\n" + stderr.String() + stdout.String()
		suite.Cases = append(suite.Cases, testCase{Name: "PackageFailure", Classname: dir, Time: "0", Failure: &detail})
	}
	sort.Slice(suite.Cases, func(i, j int) bool {
		if suite.Cases[i].Classname == suite.Cases[j].Classname {
			return suite.Cases[i].Name < suite.Cases[j].Name
		}
		return suite.Cases[i].Classname < suite.Cases[j].Classname
	})
	data, err := xml.Marshal(suite)
	if err != nil {
		return err
	}
	if err := os.WriteFile(report, data, 0644); err != nil {
		return err
	}
	if stderr.Len() > 0 {
		fmt.Fprint(os.Stderr, stderr.String())
	}
	return runErr
}

func hasFailure(cases []testCase) bool {
	for _, c := range cases {
		if c.Failure != nil {
			return true
		}
	}
	return false
}

func boundaries() error {
	bin := os.Getenv("FIREWEED_BIN")
	if bin == "" {
		bin = "fireweed"
	}
	data, err := exec.Command(bin, "modules", "--json").Output()
	if err != nil {
		return fmt.Errorf("query modules: %w", err)
	}
	var graph struct {
		Data []module `json:"data"`
	}
	if err := json.Unmarshal(data, &graph); err != nil {
		return err
	}
	modPath, err := workspaceModulePath()
	if err != nil {
		return err
	}
	var violations []string
	for _, source := range graph.Data {
		err := filepath.WalkDir(source.Dir, func(file string, entry fs.DirEntry, walkErr error) error {
			if walkErr != nil {
				return walkErr
			}
			if entry.IsDir() || !strings.HasSuffix(file, ".go") {
				return nil
			}
			tree, err := parser.ParseFile(token.NewFileSet(), file, nil, parser.ImportsOnly)
			if err != nil {
				return err
			}
			rel, err := filepath.Rel(source.Dir, file)
			if err != nil {
				return err
			}
			fromEval := matchesAny(filepath.ToSlash(rel), source.Zones.Evals)
			for _, imported := range tree.Imports {
				path, err := strconv.Unquote(imported.Path.Value)
				if err != nil {
					return err
				}
				for _, target := range graph.Data {
					base := modPath + "/" + filepath.ToSlash(filepath.Clean(target.Dir))
					if path != base && !strings.HasPrefix(path, base+"/") {
						continue
					}
					declared := source.Name == target.Name
					for _, dependency := range source.DependsOn {
						if dependency == target.Name {
							declared = true
						}
					}
					if !declared || (fromEval && !interfacePackage(target, strings.TrimPrefix(path, base))) {
						violations = append(violations, fmt.Sprintf("%s: %s imports %s (module %s) outside declared dependencies or retained interface", source.Name, filepath.ToSlash(file), path, target.Name))
					}
					break
				}
			}
			return nil
		})
		if err != nil {
			return fmt.Errorf("scan %s: %w", source.Name, err)
		}
	}
	sort.Strings(violations)
	for _, violation := range violations {
		fmt.Println(violation)
	}
	if len(violations) > 0 {
		return fmt.Errorf("%d boundary violation(s)", len(violations))
	}
	return nil
}

func workspaceModulePath() (string, error) {
	file, err := os.Open("go.mod")
	if err != nil {
		return "", err
	}
	defer file.Close()
	scan := bufio.NewScanner(file)
	for scan.Scan() {
		fields := strings.Fields(scan.Text())
		if len(fields) >= 2 && fields[0] == "module" {
			return strings.Trim(fields[1], "\""), nil
		}
	}
	if err := scan.Err(); err != nil {
		return "", err
	}
	return "", fmt.Errorf("go.mod has no module path")
}

func interfacePackage(target module, suffix string) bool {
	packageDir := filepath.Join(target.Dir, filepath.FromSlash(strings.TrimPrefix(suffix, "/")))
	entries, err := os.ReadDir(packageDir)
	if err != nil {
		return false
	}
	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".go") {
			continue
		}
		rel, err := filepath.Rel(target.Dir, filepath.Join(packageDir, entry.Name()))
		if err == nil && matchesAny(filepath.ToSlash(rel), target.Zones.Interface) {
			return true
		}
	}
	return false
}

func matchesAny(file string, patterns []string) bool {
	for _, pattern := range patterns {
		pattern = filepath.ToSlash(pattern)
		if strings.HasSuffix(pattern, "/**") && strings.HasPrefix(file, strings.TrimSuffix(pattern, "/**")+"/") {
			return true
		}
		if ok, _ := filepath.Match(pattern, file); ok {
			return true
		}
	}
	return false
}
