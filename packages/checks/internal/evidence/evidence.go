package evidence

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"regexp"
	"sort"
	"strings"

	"github.com/lurkshark/fireweed/packages/checks/model"
	"github.com/lurkshark/fireweed/packages/contracts"
)

type xmlMessage struct {
	Message string `xml:"message,attr"`
	Text    string `xml:",chardata"`
}
type xmlCase struct {
	Name      string       `xml:"name,attr"`
	Classname string       `xml:"classname,attr"`
	Failures  []xmlMessage `xml:"failure"`
	Errors    []xmlMessage `xml:"error"`
	Skipped   *xmlMessage  `xml:"skipped"`
}

func messages(in []xmlMessage) []model.Message {
	if in == nil {
		return nil
	}
	out := make([]model.Message, len(in))
	for i, m := range in {
		out[i] = model.Message{Message: m.Message, Text: m.Text}
	}
	return out
}
func suite(dec *xml.Decoder, start xml.StartElement, cases *[]model.TestCase) error {
	for {
		token, err := dec.Token()
		if err != nil {
			return fmt.Errorf("incomplete JUnit suite: %w", err)
		}
		switch v := token.(type) {
		case xml.StartElement:
			switch v.Name.Local {
			case "testsuite", "testsuites":
				if err := suite(dec, v, cases); err != nil {
					return err
				}
			case "testcase":
				var tc xmlCase
				if err := dec.DecodeElement(&tc, &v); err != nil {
					return err
				}
				item := model.TestCase{Name: tc.Name, Classname: tc.Classname, Failures: messages(tc.Failures), Errors: messages(tc.Errors), IDs: []string{}}
				if tc.Skipped != nil {
					item.Skipped = &model.Message{Message: tc.Skipped.Message, Text: tc.Skipped.Text}
				}
				*cases = append(*cases, item)
			default:
				if err := dec.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			if v.Name != start.Name {
				return fmt.Errorf("unexpected JUnit closing element %s", v.Name.Local)
			}
			return nil
		}
	}
}
func Parse(report []byte) ([]model.TestCase, error) {
	decoder := xml.NewDecoder(bytes.NewReader(report))
	cases := []model.TestCase{}
	seen := false
	for {
		token, err := decoder.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		switch v := token.(type) {
		case xml.StartElement:
			if seen || v.Name.Local != "testsuite" && v.Name.Local != "testsuites" {
				return nil, fmt.Errorf("JUnit requires a single testsuite or testsuites root")
			}
			seen = true
			if err := suite(decoder, v, &cases); err != nil {
				return nil, err
			}
		case xml.CharData:
			if strings.TrimSpace(string(v)) != "" {
				return nil, fmt.Errorf("JUnit contains text outside the root")
			}
		case xml.EndElement:
			return nil, fmt.Errorf("unexpected closing element outside JUnit root")
		}
	}
	if !seen {
		return nil, fmt.Errorf("empty JUnit report; provide a testsuite or testsuites root")
	}
	return cases, nil
}

func ownedSlice[T any](in []T) []T {
	if in == nil {
		return nil
	}
	out := make([]T, len(in))
	copy(out, in)
	return out
}
func cloneCase(in model.TestCase) model.TestCase {
	out := in
	out.Failures = ownedSlice(in.Failures)
	out.Errors = ownedSlice(in.Errors)
	// IDs are derived anew, even when an incoming list purports to be authoritative.
	out.IDs = []string{}
	if in.Skipped != nil {
		copy := *in.Skipped
		out.Skipped = &copy
	}
	return out
}
func diagnostic(cov *model.Coverage, module contracts.Module, code, message string) {
	cov.Diagnostics = append(cov.Diagnostics, model.Diagnostic{Code: code, Severity: "error", Message: message, Path: module.Dir})
}
func candidates(prefixes []string, own string) []*regexp.Regexp {
	seen := map[string]bool{}
	result := []*regexp.Regexp{}
	for _, prefix := range append(ownedSlice(prefixes), own) {
		prefix = strings.ToUpper(prefix)
		if prefix == "" || seen[prefix] {
			continue
		}
		seen[prefix] = true
		result = append(result, regexp.MustCompile(`(?i)(`+regexp.QuoteMeta(prefix)+`)[-_]?([0-9]{3,})`))
	}
	return result
}
func references(tc model.TestCase, patterns []*regexp.Regexp) []string {
	ids := map[string]bool{}
	for _, pattern := range patterns {
		for _, text := range []string{tc.Name, tc.Classname} {
			for _, parts := range pattern.FindAllStringSubmatch(text, -1) {
				id := strings.ToUpper(parts[1]) + "-" + parts[2]
				if contracts.MatchesID(tc.Name, tc.Classname, id) {
					ids[id] = true
				}
			}
		}
	}
	out := make([]string, 0, len(ids))
	for id := range ids {
		out = append(out, id)
	}
	sort.Strings(out)
	return out
}
func policySet(ids []string) map[string]bool {
	out := map[string]bool{}
	for _, id := range ids {
		out[id] = true
	}
	return out
}
func suspended(tag string) bool {
	if strings.Contains(tag, "[review]") {
		return true
	}
	for _, part := range strings.Split(tag, "[") {
		if strings.HasPrefix(part, "quarantine:") {
			reason, _, closed := strings.Cut(strings.TrimPrefix(part, "quarantine:"), "]")
			if closed && strings.TrimSpace(reason) != "" {
				return true
			}
		}
	}
	return false
}
func Analyze(module contracts.Module, prefixes []string, tests []model.TestCase, hookExit int, failFirst, already []string) model.Coverage {
	cov := model.Coverage{Module: module.Name, Tests: make([]model.TestCase, len(tests)), BySpec: map[string][]int{}, Diagnostics: []model.Diagnostic{}, HookExit: hookExit}
	current := map[string]bool{}
	for _, spec := range module.Specs {
		current[spec.ID] = true
	}
	patterns := candidates(prefixes, module.IDPrefix)
	failures := 0
	for i, tc := range tests {
		item := cloneCase(tc)
		item.IDs = references(item, patterns)
		cov.Tests[i] = item
		if item.Skipped == nil && len(item.Failures)+len(item.Errors) > 0 {
			failures++
		}
		if len(item.IDs) == 0 {
			diagnostic(&cov, module, "C12", fmt.Sprintf("Test %q (index %d) has no recognized spec ID; name a current spec in its name or classname.", item.Name, i))
		}
		for _, id := range item.IDs {
			cov.BySpec[id] = append(cov.BySpec[id], i)
			if !current[id] {
				diagnostic(&cov, module, "C12", fmt.Sprintf("Test %q (index %d) references unknown or removed spec %s; remove that reference or name a current spec.", item.Name, i, id))
			}
		}
	}
	if hookExit != 0 && failures == 0 {
		diagnostic(&cov, module, "C11", fmt.Sprintf("Hook exited %d without an executed failing/error testcase; repair the reporter or compilation before using fail-first evidence.", hookExit))
	}
	if hookExit == 0 && failures > 0 {
		diagnostic(&cov, module, "C11", "Hook reported success despite executed failing/error testcases; repair the hook exit/report consistency.")
	}
	wanted, exempt := policySet(failFirst), policySet(already)
	invalid := map[string]bool{}
	for _, list := range [][]string{failFirst, already} {
		for _, id := range list {
			if !current[id] {
				invalid[id] = true
			}
		}
	}
	for id := range invalid {
		diagnostic(&cov, module, "C12", fmt.Sprintf("Policy references unknown ID %s; use only current spec IDs in failFirst and already.", id))
	}
	for _, spec := range module.Specs {
		if suspended(spec.Tag) {
			continue
		}
		executed, failed := false, false
		for _, index := range cov.BySpec[spec.ID] {
			tc := cov.Tests[index]
			if tc.Skipped != nil {
				continue
			}
			executed = true
			failed = failed || len(tc.Failures)+len(tc.Errors) > 0
		}
		if !executed {
			diagnostic(&cov, module, "C13", spec.ID+" has no executed eval; add a non-skipped referencing testcase.")
		}
		if wanted[spec.ID] {
			if !failed && !exempt[spec.ID] {
				diagnostic(&cov, module, "P1-02", spec.ID+" must have an executed failing eval before implementation; add a regression that fails against existing behavior.")
			}
		} else if failed {
			diagnostic(&cov, module, "C14", spec.ID+" has failing evals; implement its required behavior without weakening the tests.")
		}
	}
	sort.SliceStable(cov.Diagnostics, func(i, j int) bool {
		a, b := cov.Diagnostics[i], cov.Diagnostics[j]
		if a.Path != b.Path {
			return a.Path < b.Path
		}
		if a.Code != b.Code {
			return a.Code < b.Code
		}
		if a.Severity != b.Severity {
			return a.Severity < b.Severity
		}
		return a.Message < b.Message
	})
	return cov
}
