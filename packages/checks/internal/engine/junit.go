package engine

import (
	"encoding/xml"
	"fmt"
	"regexp"
	"strings"

	"github.com/lurkshark/fireweed/packages/checks/model"
	"github.com/lurkshark/fireweed/packages/contracts"
)

type junitSuite struct {
	XMLName xml.Name
	Suites  []junitSuite `xml:"testsuite"`
	Cases   []junitCase  `xml:"testcase"`
}

type junitCase struct {
	Name      string    `xml:"name,attr"`
	Classname string    `xml:"classname,attr"`
	Seconds   float64   `xml:"time,attr"`
	Failure   *struct{} `xml:"failure"`
	Error     *struct{} `xml:"error"`
	Skipped   *struct{} `xml:"skipped"`
}

func ParseJUnit(data []byte, specs []contracts.Item) (model.Report, error) {
	var root junitSuite
	if err := xml.Unmarshal(data, &root); err != nil {
		return model.Report{}, fmt.Errorf("C11: parse JUnit: %w", err)
	}
	if root.XMLName.Local != "testsuite" && root.XMLName.Local != "testsuites" {
		return model.Report{}, fmt.Errorf("C11: invalid JUnit root %q", root.XMLName.Local)
	}
	var report model.Report
	idShape := regexp.MustCompile(`^([A-Za-z][A-Za-z0-9]*?)[-_]?[0-9]+$`)
	var prefixes []*regexp.Regexp
	seenPrefixes := map[string]bool{}
	for _, spec := range specs {
		parts := idShape.FindStringSubmatch(spec.ID)
		if parts != nil {
			prefix := strings.ToLower(parts[1])
			if !seenPrefixes[prefix] {
				seenPrefixes[prefix] = true
				prefixes = append(prefixes, regexp.MustCompile(`(?i)`+regexp.QuoteMeta(parts[1])+`[-_]?[0-9]+`))
			}
		}
	}
	var visit func(junitSuite) error
	visit = func(suite junitSuite) error {
		for _, raw := range suite.Cases {
			c := model.Case{Name: raw.Name, Classname: raw.Classname, Seconds: raw.Seconds, Result: "pass"}
			if raw.Failure != nil || raw.Error != nil {
				c.Result = "fail"
			} else if raw.Skipped != nil {
				c.Result = "skip"
			}
			for _, spec := range specs {
				if contracts.MatchesID(spec.ID, c.Name, c.Classname) {
					c.IDs = append(c.IDs, spec.ID)
				}
			}
			for _, value := range []string{c.Name, c.Classname} {
				for _, pattern := range prefixes {
					for _, span := range pattern.FindAllStringIndex(value, -1) {
						if span[0] > 0 && asciiIDChar(value[span[0]-1]) || span[1] < len(value) && value[span[1]] >= '0' && value[span[1]] <= '9' {
							continue
						}
						candidate := value[span[0]:span[1]]
						known := false
						for _, spec := range specs {
							if contracts.MatchesID(spec.ID, candidate, "") {
								known = true
								break
							}
						}
						if !known {
							return fmt.Errorf("C11: testcase %q references unknown spec ID %q", c.Name, candidate)
						}
					}
				}
			}
			if len(c.IDs) == 0 {
				return fmt.Errorf("C11: testcase %q has no known spec ID", c.Name)
			}
			report.Cases = append(report.Cases, c)
			report.Seconds += c.Seconds
		}
		for _, child := range suite.Suites {
			if err := visit(child); err != nil {
				return err
			}
		}
		return nil
	}
	if err := visit(root); err != nil {
		return model.Report{}, err
	}
	return report, nil
}

func asciiIDChar(ch byte) bool {
	return ch >= 'A' && ch <= 'Z' || ch >= 'a' && ch <= 'z' || ch >= '0' && ch <= '9'
}

func exempt(tag, kind string) bool {
	return strings.HasPrefix(strings.ToLower(strings.TrimSpace(tag)), kind)
}

func references(c model.Case, id string) bool {
	for _, got := range c.IDs {
		if got == id {
			return true
		}
	}
	return false
}

func RequireCoverage(report model.Report, specs []contracts.Item) error {
	for _, spec := range specs {
		if exempt(spec.Tag, "review") || exempt(spec.Tag, "quarantine") {
			continue
		}
		covered := false
		for _, c := range report.Cases {
			if c.Result != "skip" && references(c, spec.ID) {
				covered = true
				break
			}
		}
		if !covered {
			return fmt.Errorf("C10: no coverage for %s", spec.ID)
		}
	}
	return nil
}

func RequireFailFirst(report model.Report, ids, already []string) error {
	for _, id := range ids {
		satisfied := false
		for _, done := range already {
			if done == id {
				satisfied = true
				break
			}
		}
		if satisfied {
			continue
		}
		for _, c := range report.Cases {
			if c.Result == "fail" && references(c, id) {
				satisfied = true
				break
			}
		}
		if !satisfied {
			return fmt.Errorf("P2-01: no fail-first case for %s", id)
		}
	}
	return nil
}

func RequirePass(report model.Report, specs []contracts.Item) error {
	for _, spec := range specs {
		if exempt(spec.Tag, "quarantine") {
			continue
		}
		for _, c := range report.Cases {
			if c.Result == "fail" && references(c, spec.ID) {
				return fmt.Errorf("E01: failed case %q for %s", c.Name, spec.ID)
			}
		}
	}
	return nil
}
