package evals_test

import (
	"fmt"
	"strings"
	"testing"
	"unicode/utf8"

	"github.com/lurkshark/fireweed/packages/project"
)

// The hook puts assertions only in JUnit. Its last stderr line also separates
// the existing verification feedback from the new failed-case detail section.
const retryDetailHookMarker = "HOOK_ALPHA_STDERR_ONLY"

func retryDetailFixture(t *testing.T, alphaFailed, dependentFailed string) nextFixture {
	t.Helper()
	f := newNextFixture(t, true, true, "retry-impl")
	nextWrite(t, f.root, "packages/aardvark/SPEC.md", "---\nname: aardvark\nid_prefix: AAR\ndepends_on: [alpha]\n---\n# Aardvark\n\n## Specs\n\n### AAR-001\nAardvark MUST return a value.\n\n## Decisions\n")
	nextWrite(t, f.root, "packages/aardvark/api.go", "package aardvark\nfunc Value() string { return \"old\" }\n")
	nextWrite(t, f.root, "packages/aardvark/evals/existing_test.go", "package evals\nimport (\"testing\"; \"example.invalid/next-fixture/packages/aardvark\")\nfunc Test_AAR_001_Value(t *testing.T) { if aardvark.Value() == \"\" { t.Fatal(\"empty\") } }\n")
	// The first evals report supplies fail-first evidence. The first impl agent
	// writes "wrong"; the retry writes "new" and both reports then pass.
	hook := fmt.Sprintf(`#!/bin/sh
phase=$(sed -n 's/^phase: //p' changes/0001-next.md | head -1)
if [ "$phase" = evals ]; then
  cat > "$2" <<'XML'
<testsuite><testcase name="ALP-001" classname="base"/><testcase name="ALP-002_eval" classname="added"><failure message="eval fails before implementation"/></testcase></testsuite>
XML
  exit 1
fi
if grep -q 'return "wrong"' packages/alpha/internal/value.go; then
  if [ "$1" = alpha ]; then
    cat > "$2" <<'XML'
%s
XML
    printf '%%s\n' '%s' >&2
  else
    cat > "$2" <<'XML'
%s
XML
  fi
  exit 1
fi
if [ "$1" = alpha ]; then
  printf '%%s' '<testsuite><testcase name="ALP-001"/><testcase name="ALP-002"/></testsuite>' > "$2"
else
  printf '%%s' '<testsuite><testcase name="AAR-001"/></testsuite>' > "$2"
fi
`, alphaFailed, retryDetailHookMarker, dependentFailed)
	nextWrite(t, f.root, ".fireweed/eval.sh", hook)
	nextCommit(t, f.root, "fixture with dependent JUnit failures")
	for _, d := range project.Check(f.root) {
		if !d.Warning {
			t.Fatalf("invalid retry-detail fixture: %#v", d)
		}
	}
	return f
}

func retryDetailSection(t *testing.T, f nextFixture) (string, string) {
	t.Helper()
	first := nextRead(t, f.logDir, "worker-impl-1.prompt")
	retry := nextRead(t, f.logDir, "worker-impl-2.prompt")
	if first == "" || !strings.HasPrefix(retry, first) {
		t.Fatal("impl verification retry must start with the first impl worker's full prompt")
	}
	const marker = retryDetailHookMarker + "\n"
	i := strings.LastIndex(retry, marker)
	if i < len(first) {
		t.Fatalf("retry prompt lost dependent hook stderr marker: %q", retry[len(first):])
	}
	return retry[len(first) : i+len(marker)], retry[i+len(marker):]
}

func Test_RUN_022_RUN_043_RetryRetainsContextAndIncludesJUnitOnlyFailures(t *testing.T) {
	alpha := `<testsuite>
<testcase name="ALP-001_pass_only" classname="pass.class"/>
<testcase name="ALP-002_first" classname="alpha.first"><failure message="RUN043_ALPHA_FIRST_ASSERTION"/></testcase>
<testcase name="ALP-002_skip_only" classname="skip.class"><skipped/></testcase>
<testcase name="ALP-002_second" classname="alpha.second"><error>RUN043_ALPHA_SECOND_ASSERTION</error></testcase>
</testsuite>`
	dependent := `<testsuite>
<testcase name="AAR-001_dependent" classname="aardvark.dependent"><failure message="RUN043_DEPENDENT_ASSERTION"/></testcase>
<testcase name="AAR-001_pass_only" classname="aardvark.pass"/>
</testsuite>`
	f := retryDetailFixture(t, alpha, dependent)
	r, err := f.run(false)
	if err != nil || r.State.Phase != "applied" || len(r.Steps) != 4 {
		t.Fatalf("Next with JUnit-only dependent failure = %#v, %v", r, err)
	}
	feedback, section := retryDetailSection(t, f)
	if !strings.Contains(feedback, "ALP-002_first") || !strings.Contains(feedback, retryDetailHookMarker) {
		t.Error("retry lost existing verification feedback or hook stderr")
	}
	for _, want := range []string{
		"aardvark", "AAR-001_dependent", "aardvark.dependent", "RUN043_DEPENDENT_ASSERTION",
		"alpha", "ALP-002_first", "alpha.first", "RUN043_ALPHA_FIRST_ASSERTION",
		"ALP-002_second", "alpha.second", "RUN043_ALPHA_SECOND_ASSERTION",
	} {
		if !strings.Contains(section, want) {
			t.Errorf("failed-case section lacks %q: %q", want, section)
		}
	}
	for _, absent := range []string{"ALP-001_pass_only", "ALP-002_skip_only", "AAR-001_pass_only"} {
		if strings.Contains(section, absent) {
			t.Errorf("nonfailed case %q included in detail section", absent)
		}
	}
	previous := -1
	for _, name := range []string{"AAR-001_dependent", "ALP-002_first", "ALP-002_second"} {
		i := strings.Index(section, name)
		if i <= previous {
			t.Errorf("failed cases are not in module and report order: %q", section)
			break
		}
		previous = i
	}
}

func Test_RUN_044_RetryBoundsUnicodeFailureExcerptsAndMarksOmissions(t *testing.T) {
	var alpha strings.Builder
	alpha.WriteString(`<testsuite><testcase name="ALP-001"/>`)
	for i := 0; i < 12; i++ {
		fill := strings.Repeat("界", 5000)
		fmt.Fprintf(&alpha, `<testcase name="ALP-002_long_%02d" classname="large.case"><failure>RUN044_HEAD_%02d_%s_RUN044_TAIL_%02d</failure></testcase>`, i, i, fill, i)
	}
	alpha.WriteString(`</testsuite>`)
	dependent := `<testsuite><testcase name="AAR-001_first" classname="first.case"><failure>RUN044_DEPENDENT_FIRST</failure></testcase></testsuite>`
	f := retryDetailFixture(t, alpha.String(), dependent)
	r, err := f.run(false)
	if err != nil || r.State.Phase != "applied" || len(r.Steps) != 4 {
		t.Fatalf("Next with large JUnit failures = %#v, %v", r, err)
	}
	_, section := retryDetailSection(t, f)
	if n := utf8.RuneCountInString(section); n > 16384 {
		t.Errorf("failed-case section has %d Unicode code points, want at most 16384", n)
	}
	if !strings.Contains(section, "RUN044_HEAD_00_") || !strings.Contains(section, strings.Repeat("界", 4000)) {
		t.Error("first failure did not receive a readable Unicode excerpt")
	}
	for i := 0; i < 12; i++ {
		head := fmt.Sprintf("RUN044_HEAD_%02d_", i)
		start := strings.Index(section, head)
		if start < 0 {
			continue // This case was omitted under the section budget.
		}
		fillRunes := 0
		for _, r := range section[start+len(head):] {
			if r != '界' {
				break
			}
			fillRunes++
		}
		if n := utf8.RuneCountInString(head) + fillRunes; n > 4096 {
			t.Errorf("case %d failure excerpt contains at least %d Unicode code points, want at most 4096", i, n)
		}
		if strings.Contains(section, fmt.Sprintf("RUN044_TAIL_%02d", i)) {
			t.Errorf("case %d failure excerpt continued past its first 4096 code points", i)
		}
	}
	if !strings.Contains(section, "RUN044_DEPENDENT_FIRST") || strings.Index(section, "RUN044_DEPENDENT_FIRST") >= strings.Index(section, "RUN044_HEAD_00_") {
		t.Error("lexically first dependent failure was omitted or misordered")
	}
	if strings.Contains(section, "ALP-002_long_11") {
		t.Error("section exceeded its budget instead of omitting later cases")
	}
	lower := strings.ToLower(section)
	if !strings.Contains(lower, "truncat") && !strings.Contains(lower, "shorten") && !strings.Contains(section, "…") {
		t.Error("shortened diagnostic text has no marker")
	}
	if !strings.Contains(lower, "omitt") && !strings.Contains(lower, "remaining") && !strings.Contains(lower, "more cases") && !strings.Contains(lower, "additional cases") {
		t.Error("omitted cases have no marker")
	}
}
