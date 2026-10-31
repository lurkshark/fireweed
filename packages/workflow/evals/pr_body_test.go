package evals_test

import (
	"encoding/json"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"
	"time"

	"github.com/lurkshark/fireweed/packages/project"
	"github.com/lurkshark/fireweed/packages/workflow"
)

func prBody(t *testing.T, root, revisionRange string) workflow.PRBodyReport {
	t.Helper()
	r, err := workflow.PRBody(root, revisionRange)
	if err != nil {
		t.Fatalf("PRBody(%q): %v", revisionRange, err)
	}
	return r
}

func prRequire(t *testing.T, body string, terms ...string) {
	t.Helper()
	for _, term := range terms {
		if !strings.Contains(body, term) {
			t.Errorf("PR body lacks %q:\n%s", term, body)
		}
	}
}

func prReject(t *testing.T, body string, terms ...string) {
	t.Helper()
	for _, term := range terms {
		if strings.Contains(body, term) {
			t.Errorf("PR body unexpectedly contains %q:\n%s", term, body)
		}
	}
}

func prPhase(t *testing.T, root, phase string) {
	t.Helper()
	path := "changes/0001-manual.md"
	before := readFile(t, root, path)
	after := before
	for _, old := range []string{"evals", "audit", "impl"} {
		if strings.Contains(after, "phase: "+old+"\n") {
			after = strings.Replace(after, "phase: "+old+"\n", "phase: "+phase+"\n", 1)
			break
		}
	}
	if after == before {
		t.Fatalf("fixture phase missing in %s", path)
	}
	write(t, root, path, after)
}

// The earlier phase commits stay reachable when the requested range contains
// only the final implementation commit. The final commit also archives the
// selected identity, as a real Done does.
func prHistoryFixture(t *testing.T) (root, base, finalPhase, endpoint string) {
	t.Helper()
	root = fixture(t, "worker_writable: [workspace.lock]\n", "", "### ALP-003 [review]\nAdded review clause MUST be inspected.\n", "", "")
	alpha := readFile(t, root, "packages/alpha/SPEC.md")
	alpha = strings.Replace(alpha, "depends_on: []\n", "depends_on: []\npreserve: [internal/keep.go]\n", 1)
	alpha = strings.Replace(alpha, "### ALP-002\n", "### ALP-002 [review]\n", 1)
	write(t, root, "packages/alpha/SPEC.md", alpha)
	manual := readFile(t, root, "changes/0001-manual.md")
	manual = strings.Replace(manual, "### ALP-001\n", "### ALP-001 [review]\n", 1)
	manual = strings.Replace(manual, "module: alpha\n", "module: alpha\nphase: evals\nrun: pr-fixture\naudit_rounds: 0\n", 1)
	write(t, root, "changes/0001-manual.md", manual)
	write(t, root, "packages/alpha/internal/keep.go", "package alpha // preserved baseline\n")
	write(t, root, "packages/alpha/internal/existing.go", "package alpha // existing disposable\n")
	write(t, root, "changes/0003-outside.md", "---\nmodule: gamma\n---\n\n## Motivation\nOutside range change.\n\n## Interface\n\n## Added\n\n## Modified\n\n## Removed\n\n## Decisions\n\n")
	write(t, root, "changes/0004-removed.md", "---\nmodule: beta\n---\n\n## Motivation\nRemoved identity.\n\n## Interface\n\n## Added\n\n## Modified\n\n## Removed\n\n## Decisions\n\n")
	commit(t, root, "fixture phase base")
	base = head(t, root)

	// Evals: an interface edit, one new disposable stub, an eval file, a
	// module manifest, an ordinary module file and a shared writable lockfile.
	write(t, root, "packages/alpha/api.go", "package alpha // RANGE INTERFACE EDIT\n")
	write(t, root, "packages/alpha/evals/report_test.go", "package alpha_test // RANGE EVAL EDIT\n")
	write(t, root, "packages/alpha/internal/new_stub.go", "package alpha // RANGE STUB ADDITION\n")
	write(t, root, "packages/alpha/go.mod", "module fixture.invalid/alpha\n")
	write(t, root, "packages/alpha/other.txt", "RANGE OTHER EDIT\n")
	write(t, root, "workspace.lock", "RANGE SHARED EDIT\n")
	write(t, root, "changes/0002-pending.md", "---\nmodule: beta\neval_fix: [BET-001]\nphase: audit\nrun: beta-pr-fixture\naudit_rounds: 0\n---\n\n## Motivation\nSelected pending identity.\n\n## Interface\n\n## Added\n\n## Modified\n\n## Removed\n\n## Decisions\n\n## Provenance\n- evals 2026-10-30T12:00:00Z tests=7 failing=2 passing=5\n")
	manual = readFile(t, root, "changes/0001-manual.md")
	manual = strings.Replace(manual, "phase: evals\n", "phase: audit\n", 1)
	manual += "\n## Provenance\n- evals 2026-10-30T12:00:00Z tests=11 failing=3 passing=8\n"
	write(t, root, "changes/0001-manual.md", manual)
	commit(t, root, "evals phase with selected pending change")

	manual = readFile(t, root, "changes/0001-manual.md")
	manual = strings.Replace(manual, "phase: audit\n", "phase: evals\n", 1)
	manual = strings.Replace(manual, "audit_rounds: 0\n", "audit_rounds: 1\n", 1)
	manual += "\n## Audit\n- ALP-001: \"" + newClause + "\" — PRIOR AUDIT FINDING missing assertion\n"
	write(t, root, "changes/0001-manual.md", manual)
	commit(t, root, "audit finding round one")

	write(t, root, "packages/alpha/evals/report_test.go", "package alpha_test // CORRECTED EVAL EDIT\n")
	prPhase(t, root, "audit")
	commit(t, root, "corrected eval phase")

	manual = readFile(t, root, "changes/0001-manual.md")
	manual = strings.Replace(manual, "phase: audit\n", "phase: impl\n", 1)
	manual = strings.Replace(manual, "\n## Audit\n- ALP-001: \""+newClause+"\" — PRIOR AUDIT FINDING missing assertion\n", "\n## Audit\nNo findings\nChecked IDs: ALP-001\n", 1)
	write(t, root, "changes/0001-manual.md", manual)
	commit(t, root, "clean audit phase")
	finalPhase = head(t, root)

	write(t, root, "packages/alpha/internal/keep.go", "package alpha // PRESERVED IMPLEMENTATION EDIT\n")
	write(t, root, "packages/alpha/internal/existing.go", "package alpha // IMPLEMENTATION EDIT\n")
	manual = readFile(t, root, "changes/0001-manual.md")
	manual = strings.Replace(manual, "phase: impl\n", "phase: applied\n", 1)
	manual += "- impl 2026-10-30T12:30:00Z tests=13 failing=0 passing=13\n"
	write(t, root, "changes/archive/0001-manual.md", manual)
	if err := os.Remove(filepath.Join(root, "changes/0001-manual.md")); err != nil {
		t.Fatal(err)
	}
	if err := os.Remove(filepath.Join(root, "changes/0004-removed.md")); err != nil {
		t.Fatal(err)
	}
	commit(t, root, "implementation and archive")
	endpoint = head(t, root)
	return
}

func Test_FLOW_114_FLOW_115_FLOW_116_FLOW_117_FLOW_118_FLOW_119_FLOW_120_FLOW_121_PRBodySelectedHistoryAndEndpoint(t *testing.T) {
	root, base, finalPhase, endpoint := prHistoryFixture(t)
	// Later committed and dirty edits must not change the requested endpoint.
	write(t, root, "fireweed.yaml", strings.Replace(readFile(t, root, "fireweed.yaml"), "modules: [\"packages/*\"]", "modules: [] # LATER HEAD CONFIG MARKER", 1))
	write(t, root, "packages/alpha/SPEC.md", readFile(t, root, "packages/alpha/SPEC.md")+"\nLATER HEAD SPEC MARKER\n")
	commit(t, root, "later unrelated HEAD")
	write(t, root, "fireweed.yaml", readFile(t, root, "fireweed.yaml")+"\n# DIRTY WORKTREE CONFIG MARKER\n")
	write(t, root, "packages/alpha/SPEC.md", readFile(t, root, "packages/alpha/SPEC.md")+"\nDIRTY WORKTREE SPEC MARKER\n")
	full := base + ".." + endpoint
	partial := finalPhase + ".." + endpoint
	t.Run("FLOW_114_explicit_endpoints", func(t *testing.T) {
		r := prBody(t, root, full)
		prRequire(t, r.Body, base[:12], endpoint[:12])
		prReject(t, r.Body, "LATER HEAD CONFIG MARKER", "LATER HEAD SPEC MARKER", "DIRTY WORKTREE CONFIG MARKER", "DIRTY WORKTREE SPEC MARKER")
	})
	t.Run("FLOW_115_selected_identities", func(t *testing.T) {
		r := prBody(t, root, full)
		prRequire(t, r.Body, "changes/archive/0001-manual.md", "alpha", "changes/0002-pending.md", "beta", "eval_fix")
		prReject(t, r.Body, "0003-outside", "0004-removed")
		p := prBody(t, root, partial)
		prRequire(t, p.Body, "changes/archive/0001-manual.md")
		prReject(t, p.Body, "0002-pending.md")
	})
	t.Run("FLOW_116_provenance_and_evals", func(t *testing.T) {
		fullReport := prBody(t, root, full)
		prRequire(t, fullReport.Body, "tests=7", "failing=2", "passing=5")
		for _, rr := range []string{full, partial} {
			r := prBody(t, root, rr)
			prRequire(t, r.Body, "tests=11", "failing=3", "passing=8", "report_test.go")
		}
	})
	t.Run("FLOW_117_phase_file_categories", func(t *testing.T) {
		for _, rr := range []string{full, partial} {
			r := prBody(t, root, rr)
			prRequire(t, r.Body, "api.go", "new_stub.go", "go.mod", "other.txt", "workspace.lock")
		}
	})
	t.Run("FLOW_118_preserved_edits", func(t *testing.T) {
		r := prBody(t, root, full)
		prRequire(t, r.Body, "keep.go")
	})
	t.Run("FLOW_119_review_specs", func(t *testing.T) {
		for _, rr := range []string{full, partial} {
			r := prBody(t, root, rr)
			prRequire(t, r.Body, "ALP-003", "Added review clause MUST be inspected.", "ALP-001", newClause)
			prReject(t, r.Body, "Alpha MUST retain the second behavior.")
		}
	})
	t.Run("FLOW_120_prior_corrected_audit", func(t *testing.T) {
		for _, rr := range []string{full, partial} {
			r := prBody(t, root, rr)
			prRequire(t, r.Body, "PRIOR AUDIT FINDING")
			if !strings.Contains(r.Body, "audit_rounds") && !strings.Contains(strings.ToLower(r.Body), "audit rounds") {
				t.Errorf("PR body omits audit round count: %s", r.Body)
			}
		}
	})
	t.Run("FLOW_121_decisions", func(t *testing.T) {
		r := prBody(t, root, full)
		prRequire(t, r.Body, "Use the revised policy.", "Retain old policy.")
	})
}

func prSaveCost(t *testing.T, root, run, phase string, snapshot workflow.CostSnapshot) {
	t.Helper()
	b, err := json.Marshal(snapshot)
	if err != nil {
		t.Fatal(err)
	}
	write(t, root, filepath.Join(".fireweed/runs", run, phase, "costs.json"), string(b))
}

func Test_FLOW_122_FLOW_123_PRBodyEndpointExceptionsAndRetainedCosts(t *testing.T) {
	root := fixture(t, "", "", "", "", "")
	base := head(t, root)
	alpha := readFile(t, root, "packages/alpha/SPEC.md")
	alpha = strings.Replace(alpha, "depends_on: []\n", "depends_on: []\nboundaries_opt_out: ENDPOINT OPT OUT REASON\n", 1)
	alpha = strings.Replace(alpha, "### ALP-002\n", "### ALP-002 [quarantine: ENDPOINT QUARANTINE REASON]\n", 1)
	write(t, root, "packages/alpha/SPEC.md", alpha)
	manual := readFile(t, root, "changes/0001-manual.md")
	write(t, root, "changes/0001-manual.md", strings.Replace(manual, "module: alpha\n", "module: alpha\nblocked: ENDPOINT BLOCK REASON\n", 1))
	commit(t, root, "endpoint exceptions")
	endpoint := head(t, root)
	old := workflow.CostSnapshot{Run: "old", Phase: "evals", RecordedAt: time.Date(2026, 10, 29, 0, 0, 0, 0, time.UTC), Modules: []workflow.ModuleCost{{Module: "alpha", Seconds: 99, CasesPerSpec: map[string]int{"ALP-001": 99}, EvalLines: 99, DisposableLines: 1}}, Diagnostics: []project.Diagnostic{{Code: "B01", Message: "OLD B01", Warning: true}}}
	latest := workflow.CostSnapshot{Run: "new", Phase: "impl", RecordedAt: time.Date(2026, 10, 30, 0, 0, 0, 0, time.UTC), Modules: []workflow.ModuleCost{{Module: "alpha", Seconds: 4.25, CasesPerSpec: map[string]int{"ALP-001": 5}, EvalLines: 17, DisposableLines: 3}, {Module: "beta", Seconds: 2.5, CasesPerSpec: map[string]int{"BET-001": 2}, EvalLines: 8, DisposableLines: 4}}, Diagnostics: []project.Diagnostic{{Code: "B01", Message: "SAVED B01 time", Warning: true}, {Code: "B02", Message: "ALP-001: SAVED B02 cases", Warning: true}, {Code: "B03", Message: "alpha: SAVED B03 ratio", Warning: true}}}
	prSaveCost(t, root, "old", "evals", old)
	prSaveCost(t, root, "new", "impl", latest)
	write(t, root, "fireweed.yaml", readFile(t, root, "fireweed.yaml")+"\nbudgets:\n  cases_per_spec: 1000\n  eval_ratio: 1000\n")
	write(t, root, "packages/alpha/evals/later.txt", strings.Repeat("later\n", 50))
	t.Run("FLOW_122_endpoint_exceptions", func(t *testing.T) {
		r := prBody(t, root, base+".."+endpoint)
		prRequire(t, r.Body, "alpha", "ALP-002", "ENDPOINT QUARANTINE REASON", "ENDPOINT OPT OUT REASON", "0001-manual", "ENDPOINT BLOCK REASON")
	})
	t.Run("FLOW_123_latest_saved_costs", func(t *testing.T) {
		r := prBody(t, root, base+".."+endpoint)
		prRequire(t, r.Body, "SAVED B01", "SAVED B02", "SAVED B03")
		prReject(t, r.Body, "OLD B01")
		section := regexp.MustCompile(`(?ms)^#{1,6} (?:Cost|Costs|Metric|Metrics)\s*\n(.*?)(?:^#{1,6} |\z)`).FindStringSubmatch(r.Body)
		if len(section) != 2 {
			t.Errorf("PR body lacks a Cost/Metric section: %s", r.Body)
		} else {
			rows := regexp.MustCompile(`(?m)^\|\s*(?:alpha|beta|gamma)\s*\|[^\n]*$`).FindAllString(section[1], -1)
			want := map[string]string{
				"alpha": "alpha|4.25|ALP-001: 5|17|3",
				"beta":  "beta|2.5|BET-001: 2|8|4",
			}
			seen := map[string]bool{}
			for _, row := range rows {
				cells := strings.Split(strings.Trim(strings.TrimSpace(row), "|"), "|")
				for i := range cells {
					cells[i] = strings.TrimSpace(cells[i])
				}
				module := cells[0]
				if seen[module] || strings.Join(cells, "|") != want[module] {
					t.Errorf("unexpected or stale cost measurement row %q", row)
				}
				seen[module] = true
			}
			for module := range want {
				if !seen[module] {
					t.Errorf("missing saved %s cost measurement row: %s", module, section[1])
				}
			}
		}
		for _, code := range []string{"B01", "B02", "B03"} {
			if !containsCode(r.Diagnostics, code) {
				t.Errorf("PRBody diagnostics lack saved %s: %#v", code, r.Diagnostics)
			}
		}
	})
}

func Test_FLOW_126_PRBodyEndpointConfigError(t *testing.T) {
	root := fixture(t, "", "", "", "", "")
	base := head(t, root)
	write(t, root, "fireweed.yaml", "[invalid endpoint configuration\n")
	commit(t, root, "invalid endpoint configuration")

	_, err := workflow.PRBody(root, base+".."+head(t, root))
	if err == nil {
		t.Fatal("PRBody accepted malformed endpoint configuration")
	}
	if !workflow.IsEndpointConfigError(err) {
		t.Errorf("IsEndpointConfigError(%v) = false, want true", err)
	}
}

func Test_FLOW_127_PRBodyUnresolvableRangeIsNotEndpointConfigError(t *testing.T) {
	root := fixture(t, "", "", "", "", "")
	_, err := workflow.PRBody(root, head(t, root)+"..not-a-commit")
	if err == nil {
		t.Fatal("PRBody accepted an unresolvable range")
	}
	if workflow.IsEndpointConfigError(err) {
		t.Errorf("IsEndpointConfigError(%v) = true, want false", err)
	}
}

func Test_FLOW_124_FLOW_125_PRBodyDisposableArtifactBoundaryAndRanges(t *testing.T) {
	root := fixture(t, "", "", "", "", "")
	base := head(t, root)
	for path, data := range map[string][]byte{
		"packages/alpha/internal/binary.dat": {0, 1, 2, 3},
		"packages/alpha/internal/large.txt":  []byte(strings.Repeat("L", 262145)),
		"packages/alpha/internal/exact.txt":  []byte(strings.Repeat("E", 262144)),
		"packages/alpha/internal/small.txt":  []byte("ordinary text\n"),
	} {
		full := filepath.Join(root, path)
		if err := os.MkdirAll(filepath.Dir(full), 0755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(full, data, 0644); err != nil {
			t.Fatal(err)
		}
	}
	commit(t, root, "new disposable artifacts")
	endpoint := head(t, root)
	t.Run("FLOW_124_binary_and_large_artifacts", func(t *testing.T) {
		r := prBody(t, root, base+".."+endpoint)
		prRequire(t, r.Body, "binary.dat", "large.txt", "262145", "B04")
		if !regexp.MustCompile("(?m)(binary\\.dat[^\\n]*\\b4\\b|\\b4\\b[^\\n]*binary\\.dat)").MatchString(r.Body) {
			t.Errorf("binary artifact lacks its 4-byte size: %s", r.Body)
		}
		count := 0
		for _, d := range r.Diagnostics {
			if d.Code == "B04" {
				count++
				if !d.Warning || !(strings.Contains(d.Message, "binary.dat") || strings.Contains(d.Message, "large.txt")) || strings.Contains(d.Message, "exact.txt") || strings.Contains(d.Message, "small.txt") {
					t.Errorf("unexpected B04 diagnostic: %#v", d)
				}
			}
		}
		if count != 2 {
			t.Errorf("B04 count = %d, want 2: %#v", count, r.Diagnostics)
		}
	})
	t.Run("FLOW_125_empty_and_invalid_ranges", func(t *testing.T) {
		empty := prBody(t, root, endpoint+".."+endpoint)
		if strings.TrimSpace(empty.Body) == "" || !(strings.Contains(strings.ToLower(empty.Body), "none") || strings.Contains(strings.ToLower(empty.Body), "no change")) {
			t.Errorf("empty range did not explicitly report no selected changes: %q", empty.Body)
		}
		if containsCode(empty.Diagnostics, "B04") {
			t.Errorf("empty range yielded B04: %#v", empty.Diagnostics)
		}
		for _, invalid := range []string{"", "not-a-commit.." + endpoint, base + "..not-a-commit", "not-a-range"} {
			_, err := workflow.PRBody(root, invalid)
			if err == nil || !strings.Contains(err.Error(), "P01") || (invalid != "" && !strings.Contains(err.Error(), "not-a-commit") && !strings.Contains(err.Error(), "not-a-range") && !strings.Contains(err.Error(), "resolve")) {
				t.Errorf("PRBody(%q) error = %v, want P01 with rejected range or cause", invalid, err)
			}
		}
	})
}
