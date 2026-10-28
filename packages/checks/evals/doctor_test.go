package evals_test

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
	"time"

	"github.com/lurkshark/fireweed/packages/assets"
	"github.com/lurkshark/fireweed/packages/checks"
	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/project"
	"gopkg.in/yaml.v3"
)

const doctorGuidance = "<!-- fireweed:start -->\nThis project uses fireweed. Specs and change files are the source of truth.\nRun `fireweed playbook author-change` before authoring changes and\n`fireweed context <change>` before each step.\n<!-- fireweed:end -->\n"

type doctorFixture struct {
	t    *testing.T
	root string
	log  string
	bin  string
}

func (f doctorFixture) put(name, body string) {
	f.t.Helper()
	path := filepath.Join(f.root, filepath.FromSlash(name))
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		f.t.Fatal(err)
	}
	if err := os.WriteFile(path, []byte(body), 0644); err != nil {
		f.t.Fatal(err)
	}
}

func (f doctorFixture) git(args ...string) string {
	f.t.Helper()
	cmd := exec.Command("git", args...)
	cmd.Dir = f.root
	cmd.Env = append(os.Environ(), "GIT_CONFIG_NOSYSTEM=1", "GIT_TERMINAL_PROMPT=0")
	out, err := cmd.CombinedOutput()
	if err != nil {
		f.t.Fatalf("git %v: %v: %s", args, err, out)
	}
	return string(out)
}

func (f doctorFixture) commit() {
	f.t.Helper()
	f.git("init", "-b", "main")
	f.git("config", "user.name", "Doctor Fixture")
	f.git("config", "user.email", "doctor@example.test")
	f.git("add", ".")
	f.git("commit", "-m", "fixture")
}

func (f doctorFixture) request(noAgent bool) checks.DoctorRequest {
	return checks.DoctorRequest{Root: f.root, Bin: f.bin, NoAgent: noAgent, Timeout: 15 * time.Second}
}

func doctorConfig(t *testing.T, recipe assets.Recipe) string {
	t.Helper()
	config := map[string]any{
		"fireweed": ">=0.1, <0.2", "version": 1, "base": "main", "modules": []string{"packages/*"},
		"limit_patterns":  []string{"rate limit"},
		"module_defaults": map[string]any{"interface": recipe.ModuleDefaults.Interface, "evals": recipe.ModuleDefaults.Evals, "disposable": recipe.ModuleDefaults.Disposable},
		"hooks":           recipe.Hooks,
		"agent":           map[string]any{"command": []string{"sh", ".fireweed/agent.sh", "{prompt}", "{git_common_dir}"}, "wrap": []string{"sh", ".fireweed/wrap.sh"}},
		"audit":           map[string]any{"enabled": true, "command": []string{"sh", ".fireweed/audit.sh", "{prompt}"}},
	}
	data, err := yaml.Marshal(config)
	if err != nil {
		t.Fatal(err)
	}
	return string(data)
}

func (f doctorFixture) editConfig(change func(map[string]any)) {
	f.t.Helper()
	var config map[string]any
	if err := yaml.Unmarshal(mustRead(f.t, filepath.Join(f.root, "fireweed.yaml")), &config); err != nil {
		f.t.Fatal(err)
	}
	change(config)
	data, err := yaml.Marshal(config)
	if err != nil {
		f.t.Fatal(err)
	}
	f.put("fireweed.yaml", string(data))
}

func newDoctorFixture(t *testing.T) doctorFixture {
	t.Helper()
	r, err := assets.BuiltinRecipe("go")
	if err != nil {
		t.Fatal(err)
	}
	f := doctorFixture{t: t, root: t.TempDir(), log: filepath.Join(t.TempDir(), "probes.log")}
	f.put("go.mod", "module example.test/doctorfixture\n\ngo 1.23\n")
	f.put("fireweed.yaml", doctorConfig(t, r))
	f.put("AGENTS.md", "Before.\n"+doctorGuidance+"After.\n")
	f.put("packages/alpha/SPEC.md", "---\nname: alpha\nid_prefix: ALP\ndepends_on: []\n---\n# Alpha\n\n## Specs\n\n### ALP-001\nThe fixture MUST pass.\n")
	f.put("packages/alpha/api.go", "package alpha\n")
	for path, body := range r.Files {
		f.put(path, body)
	}
	f.put(".fireweed/wrap.sh", fmt.Sprintf("#!/bin/sh\nprintf 'wrap:%%s\n' \"$PWD\" >> %q\nexec \"$@\"\n", f.log))
	f.put(".fireweed/agent.sh", fmt.Sprintf("#!/bin/sh\nprintf 'worker:%%s:%%s:%%s\n' \"$PWD\" \"$1\" \"$2\" >> %q\nprintf 'FIREWEED_DOCTOR_OK\\n'\n", f.log))
	f.put(".fireweed/audit.sh", fmt.Sprintf("#!/bin/sh\nprintf 'audit:%%s:%%s\n' \"$PWD\" \"$1\" >> %q\nprintf 'FIREWEED_DOCTOR_OK\\n'\n", f.log))
	return f
}

func localDoctorFixture(t *testing.T) doctorFixture {
	t.Helper()
	f := doctorFixture{t: t, root: t.TempDir(), log: filepath.Join(t.TempDir(), "probes.log")}
	f.bin = filepath.Join(t.TempDir(), "fake-fireweed")
	if err := os.WriteFile(f.bin, []byte(`#!/bin/sh
if test "$1" != modules || test "$2" != --json; then exit 2; fi
printf '{"data":['
first=1
find packages -name SPEC.md | while IFS= read -r spec; do
  dir=${spec%/SPEC.md}; name=$(sed -n 's/^name: //p' "$spec" | head -1)
  if test "$first" = 0; then printf ','; fi; first=0
  printf '{"name":"%s","dir":"%s","depends_on":[],"zones":{"interface":["api.txt"],"evals":["evals/**"],"disposable":["internal/**"]}}' "$name" "$dir"
done
printf ']}'
`), 0700); err != nil {
		t.Fatal(err)
	}
	recipe := assets.Recipe{ModuleDefaults: structZones(), Hooks: map[string]any{
		"eval":       "sh .fireweed/eval.sh {module} {module_dir} {report}",
		"boundaries": "sh .fireweed/boundary.sh",
		"setup":      "sh .fireweed/setup.sh",
	}}
	f.put("fireweed.yaml", doctorConfig(t, recipe))
	f.put("AGENTS.md", doctorGuidance)
	f.put("packages/alpha/SPEC.md", "---\nname: alpha\nid_prefix: ALP\ndepends_on: []\n---\n# Alpha\n\n## Specs\n\n### ALP-001\nThe fixture MUST pass.\n")
	f.put("packages/alpha/api.txt", "public\n")
	f.put(".fireweed/templates/module/SPEC.md", "---\nname: {{name}}\nid_prefix: {{prefix}}\ndepends_on: []\n---\n# {{name}}\n\n## Specs\n\n### {{prefix}}-001\nThe canary MUST pass.\n")
	f.put(".fireweed/templates/module/api.txt.tmpl", "{{package}} {{go_module}} {{module_dir}} {{other_dir}}\n")
	f.put(".fireweed/canaries/canaries.yaml", `version: 1
baseline:
  SPEC.md: "---\nname: {{name}}\nid_prefix: {{prefix}}\ndepends_on: []\n---\n# {{name}}\n\n## Specs\n\n### {{prefix}}-001\nThe canary MUST pass.\n"
  api.txt: "{{package}} {{go_module}} {{module_dir}} {{other_dir}}"
undeclared:
  undeclared.txt: "{{other_dir}}"
internal_eval:
  internal/private.txt: "private"
  evals/internal.txt: "{{module_dir}}"
failing_eval:
  evals/failing.txt: "{{prefix}}-001"
missing_id:
  evals/missing.txt: "UNMAPPED-999"
`)
	f.put(".fireweed/setup.sh", fmt.Sprintf("#!/bin/sh\nprintf 'setup:%%s\n' \"$PWD\" >> %q\n", f.log))
	f.put(".fireweed/eval.sh", fmt.Sprintf(`#!/bin/sh
printf 'eval:%%s:%%s:%%s\n' "$PWD" "$1" "$FIREWEED_BIN" >> %q
prefix=$(sed -n 's/^id_prefix: //p' "$2/SPEC.md" | head -1)
if test -f "$2/evals/failing.txt"; then printf '<testsuite><testcase name="%%s-001"><failure/></testcase></testsuite>' "$prefix" > "$3"; exit 1; fi
if test -f "$2/evals/missing.txt"; then printf '<testsuite><testcase name="UNMAPPED-999"/></testsuite>' > "$3"; exit 0; fi
printf '<testsuite><testcase name="%%s-001"/></testsuite>' "$prefix" > "$3"
`, f.log))
	f.put(".fireweed/boundary.sh", fmt.Sprintf(`#!/bin/sh
printf 'boundaries:%%s:%%s\n' "$PWD" "$FIREWEED_BIN" >> %q
find packages -name SPEC.md | while IFS= read -r spec; do
  dir=${spec%%/SPEC.md}; name=$(sed -n 's/^name: //p' "$spec" | head -1)
  if test -f "$dir/undeclared.txt"; then printf '%%s: undeclared import\n' "$name"; fi
  if test -f "$dir/evals/internal.txt"; then printf '%%s: interface violation\n' "$name"; fi
done
`, f.log))
	f.put(".fireweed/wrap.sh", fmt.Sprintf("#!/bin/sh\nprintf 'wrap:%%s\n' \"$PWD\" >> %q\nexec \"$@\"\n", f.log))
	f.put(".fireweed/agent.sh", fmt.Sprintf("#!/bin/sh\nprintf 'worker:%%s:%%s:%%s\n' \"$PWD\" \"$1\" \"$2\" >> %q\nprintf 'FIREWEED_DOCTOR_OK\\n'\n", f.log))
	f.put(".fireweed/audit.sh", fmt.Sprintf("#!/bin/sh\nprintf 'audit:%%s:%%s\n' \"$PWD\" \"$1\" >> %q\nprintf 'FIREWEED_DOCTOR_OK\\n'\n", f.log))
	return f
}

func structZones() (z contracts.Zones) {
	z.Interface = []string{"api.txt"}
	z.Evals = []string{"evals/**"}
	z.Disposable = []string{"internal/**"}
	return
}

func doctorText(report checks.DoctorReport) string {
	return fmt.Sprintf("checks=%+v diagnostics=%+v", report.Checks, report.Diagnostics)
}
func hasDoctorDiagnostic(report checks.DoctorReport, code string, warning bool, terms ...string) bool {
	for _, d := range report.Diagnostics {
		if d.Code != code || d.Warning != warning {
			continue
		}
		ok := true
		for _, s := range terms {
			if !strings.Contains(strings.ToLower(d.Message), strings.ToLower(s)) {
				ok = false
			}
		}
		if ok {
			return true
		}
	}
	return false
}
func hasDoctorCheck(report checks.DoctorReport, name string, passed bool) bool {
	for _, c := range report.Checks {
		if strings.Contains(c.Name, name) && c.Passed == passed {
			return true
		}
	}
	return false
}
func doctorErrors(report checks.DoctorReport) bool {
	for _, d := range report.Diagnostics {
		if !d.Warning {
			return true
		}
	}
	return false
}
func doctorLog(f doctorFixture) string { b, _ := os.ReadFile(f.log); return string(b) }
func worktrees(f doctorFixture) string { return f.git("worktree", "list", "--porcelain") }

func assertScratchRemoved(t *testing.T, f doctorFixture) {
	t.Helper()
	common := strings.TrimSpace(f.git("rev-parse", "--git-common-dir"))
	if !filepath.IsAbs(common) {
		common = filepath.Join(f.root, common)
	}
	common, err := filepath.EvalSymlinks(common)
	if err != nil {
		t.Fatal(err)
	}
	base := filepath.Join(common, "fireweed", "worktrees")
	for _, line := range strings.Split(doctorLog(f), "\n") {
		if !strings.HasPrefix(line, "setup:") {
			continue
		}
		path := strings.TrimPrefix(line, "setup:")
		for parent := path; ; parent = filepath.Dir(parent) {
			resolved, err := filepath.EvalSymlinks(parent)
			if err == nil {
				suffix, err := filepath.Rel(parent, path)
				if err != nil {
					t.Fatal(err)
				}
				path = filepath.Join(resolved, suffix)
				break
			}
			if parent == filepath.Dir(parent) {
				t.Fatal(err)
			}
		}
		rel, err := filepath.Rel(base, path)
		if err != nil || rel == "." || rel == ".." || strings.HasPrefix(rel, ".."+string(filepath.Separator)) {
			t.Errorf("scratch %q is not below common Git directory %q", path, base)
		}
		if _, err := os.Stat(path); !os.IsNotExist(err) {
			t.Errorf("physical scratch directory %q remains or cannot be checked: %v", path, err)
		}
		return
	}
	t.Error("setup did not record a scratch path")
}

func Test_CHK_049_StaticDiagnosticsStopProbes(t *testing.T) {
	f := localDoctorFixture(t)
	f.put("packages/alpha/SPEC.md", "---\nname: alpha\nid_prefix: ALP\ndepends_on: []\n---\n## Specs\n### ALP-001\nA.\n### ALP-001\nB.\n")
	f.commit()
	want := project.Check(f.root)
	if len(want) == 0 || !doctorErrors(checks.DoctorReport{Diagnostics: want}) {
		t.Fatal("fixture has no static diagnostic")
	}
	report := checks.Doctor(context.Background(), f.request(false))
	remaining := map[project.Diagnostic]int{}
	for _, d := range report.Diagnostics {
		remaining[d]++
	}
	for _, d := range want {
		if remaining[d] == 0 {
			t.Errorf("complete project diagnostic %+v lost: %s", d, doctorText(report))
		}
		remaining[d]--
	}
	if strings.Contains(doctorLog(f), "setup:") || strings.Contains(doctorLog(f), "eval:") || strings.Contains(doctorLog(f), "boundaries:") || strings.Contains(doctorLog(f), "worker:") || strings.Contains(doctorLog(f), "audit:") {
		t.Errorf("static error ran probes: %s", doctorLog(f))
	}
}

func Test_CHK_050_CHK_051_CHK_052_CHK_053_CHK_055_CHK_056_CHK_057_CHK_058_CHK_059_CHK_060_CHK_061_CHK_067_RealGoRecipe(t *testing.T) {
	f := newDoctorFixture(t)
	f.commit()
	f.bin = doctorBinary(t)
	f.editConfig(func(config map[string]any) {
		config["modules"] = []string{"packages/alpha"}
		config["hooks"].(map[string]any)["setup"] = fmt.Sprintf("printf 'setup:%%s\\n' \"$PWD\" >> %q; test -z \"$(git symbolic-ref -q HEAD)\" && test -f .fireweed/doctor-untracked && go mod download", f.log)
	})
	f.put(".fireweed/doctor-untracked", "ready\n")
	for _, d := range project.Check(f.root) {
		if !d.Warning {
			t.Fatalf("invalid Go doctor fixture: %+v", d)
		}
	}
	beforeStatus, beforeWorktrees := f.git("status", "--porcelain", "--untracked-files=all"), worktrees(f)
	report := checks.Doctor(context.Background(), f.request(true))
	if doctorErrors(report) {
		t.Errorf("real Go recipe doctor failed: %s", doctorText(report))
	}
	for _, name := range []string{"baseline", "undeclared", "internal_eval", "failing_eval", "missing_id", "recovery"} {
		if !hasDoctorCheck(report, name, true) {
			t.Errorf("missing passing %s outcome: %s", name, doctorText(report))
		}
	}
	recoveries := 0
	for _, check := range report.Checks {
		if strings.Contains(strings.ToLower(check.Name), "recovery") && check.Passed {
			recoveries++
		}
	}
	if recoveries < 4 {
		t.Errorf("clean recovery not reported after each canary: %s", doctorText(report))
	}
	modules := map[string]bool{}
	for _, check := range report.Checks {
		if check.Module != "" && check.Module != "alpha" {
			modules[check.Module] = true
		}
	}
	if len(modules) < 2 {
		t.Errorf("two distinct throwaway modules absent: %s", doctorText(report))
	}
	if !strings.Contains(doctorLog(f), "setup:") || !strings.Contains(doctorLog(f), "/fireweed/worktrees/") {
		t.Errorf("setup did not run in scratch: %s", doctorLog(f))
	}
	if strings.Contains(doctorLog(f), "worker:") || strings.Contains(doctorLog(f), "audit:") {
		t.Errorf("NoAgent ran agents: %s", doctorLog(f))
	}
	if got := f.git("status", "--porcelain", "--untracked-files=all"); got != beforeStatus {
		t.Errorf("caller changed: before %q after %q", beforeStatus, got)
	}
	if got := worktrees(f); got != beforeWorktrees {
		t.Errorf("scratch registration leaked: before %q after %q", beforeWorktrees, got)
	}
	assertScratchRemoved(t, f)
}

func Test_CHK_051_CHK_052_SetupFailureStopsAndCleans(t *testing.T) {
	f := localDoctorFixture(t)
	f.put(".fireweed/setup.sh", fmt.Sprintf("#!/bin/sh\nprintf 'setup:%%s\\n' \"$PWD\" >> %q\necho failed setup >&2\nexit 17\n", f.log))
	f.commit()
	before := worktrees(f)
	report := checks.Doctor(context.Background(), f.request(false))
	if !hasDoctorDiagnostic(report, "C14", false, "setup") {
		t.Errorf("setup failure absent: %s", doctorText(report))
	}
	if strings.Contains(doctorLog(f), "eval:") || strings.Contains(doctorLog(f), "boundaries:") || strings.Contains(doctorLog(f), "worker:") {
		t.Errorf("setup failure ran later probes: %s", doctorLog(f))
	}
	if got := worktrees(f); got != before {
		t.Errorf("scratch worktree leaked: %q", got)
	}
	assertScratchRemoved(t, f)
}

func Test_CHK_051_CancellationCleansScratchAndStopsSetupGroup(t *testing.T) {
	f := localDoctorFixture(t)
	ready := filepath.Join(t.TempDir(), "ready")
	childFinished := filepath.Join(t.TempDir(), "child-finished")
	f.put(".fireweed/setup.sh", fmt.Sprintf(`#!/bin/sh
printf 'setup:%%s\n' "$PWD" >> %q
sh -c 'sleep 2; printf finished > "$1"' sh %q &
printf ready > %q
wait
`, f.log, childFinished, ready))
	f.commit()
	before := worktrees(f)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	reports := make(chan checks.DoctorReport, 1)
	go func() { reports <- checks.Doctor(ctx, f.request(true)) }()
	deadline := time.After(8 * time.Second)
	for {
		if _, err := os.Stat(ready); err == nil {
			break
		} else if !os.IsNotExist(err) {
			t.Fatal(err)
		}
		select {
		case report := <-reports:
			t.Fatalf("setup exited before cancellation: %s", doctorText(report))
		case <-deadline:
			t.Fatal("setup never signalled readiness")
		case <-time.After(10 * time.Millisecond):
		}
	}
	cancel()
	select {
	case report := <-reports:
		if !hasDoctorDiagnostic(report, "C14", false, "setup") {
			t.Errorf("cancelled setup lacked C14: %s", doctorText(report))
		}
	case <-time.After(5 * time.Second):
		t.Fatal("cancelled doctor did not return promptly")
	}
	if got := worktrees(f); got != before {
		t.Errorf("cancelled doctor retained Git registration: before %q after %q", before, got)
	}
	assertScratchRemoved(t, f)
	time.Sleep(2200 * time.Millisecond)
	if _, err := os.Stat(childFinished); !os.IsNotExist(err) {
		t.Errorf("cancelled setup descendant continued: %v", err)
	}
}

func Test_CHK_054_MissingTemplatesAndCanaryMaps(t *testing.T) {
	for _, tc := range []struct{ name, path, body, term, input, affected string }{
		{"missing template", ".fireweed/templates/module/SPEC.md", "", "template", "SPEC.md", "undeclared"},
		{"malformed template", ".fireweed/templates/module/SPEC.md", "---\nname: [broken\nid_prefix: {{prefix}}\n---\n", "template", "SPEC.md", "undeclared"},
		{"missing baseline", ".fireweed/canaries/canaries.yaml", "version: 1\nbaseline: {}\n", "canary", "baseline", "undeclared"},
		{"missing internal_eval", ".fireweed/canaries/canaries.yaml", "missing_id:\n  evals/missing.txt: missing\n", "internal_eval", "internal_eval", "internal_eval"},
		{"wrong version", ".fireweed/canaries/canaries.yaml", "version: 2\nbaseline: {}\n", "canary", "canaries.yaml", "undeclared"},
		{"malformed canary", ".fireweed/canaries/canaries.yaml", "version: [bad\n", "canary", "canaries.yaml", "undeclared"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			f := localDoctorFixture(t)
			if tc.name == "missing baseline" {
				body := string(mustRead(t, filepath.Join(f.root, tc.path)))
				start, end := strings.Index(body, "baseline:\n"), strings.Index(body, "undeclared:\n")
				if start < 0 || end <= start {
					t.Fatal("fixture cannot isolate baseline map")
				}
				tc.body = body[:start] + "baseline: {}\n" + body[end:]
			}
			if tc.name == "missing internal_eval" {
				body := string(mustRead(t, filepath.Join(f.root, tc.path)))
				tc.body = strings.Replace(body, "internal_eval:\n  internal/private.txt: \"private\"\n  evals/internal.txt: \"{{module_dir}}\"\n", "", 1)
				if tc.body == body {
					t.Fatal("fixture did not remove internal_eval map")
				}
			}
			if tc.body == "" {
				if err := os.Remove(filepath.Join(f.root, tc.path)); err != nil {
					t.Fatal(err)
				}
			} else {
				f.put(tc.path, tc.body)
			}
			f.commit()
			report := checks.Doctor(context.Background(), f.request(true))
			if !hasDoctorDiagnostic(report, "C14", false, tc.term) {
				t.Errorf("missing C14: %s", doctorText(report))
			}
			identified := false
			for _, d := range report.Diagnostics {
				if d.Code == "C14" && !d.Warning && strings.Contains(strings.ToLower(d.Message+" "+d.Path), strings.ToLower(tc.input)) {
					identified = true
				}
			}
			if !identified {
				t.Errorf("C14 did not identify %s: %s", tc.input, doctorText(report))
			}
			if hasDoctorCheck(report, tc.affected, true) {
				t.Errorf("invalid recipe passed canary: %s", doctorText(report))
			}
		})
	}
}

func Test_CHK_055_CHK_056_CHK_057_CHK_058_CHK_059_ZeroHookAndRecovery(t *testing.T) {
	f := localDoctorFixture(t)
	f.put(".fireweed/boundary.sh", "#!/bin/sh\nexit 0\n")
	f.put(".fireweed/eval.sh", `#!/bin/sh
printf '<testsuite><testcase name="ALP-001"/></testsuite>' > "$3"
`)
	f.commit()
	before := worktrees(f)
	report := checks.Doctor(context.Background(), f.request(true))
	for _, tc := range []struct{ name, hook string }{{"undeclared", "boundar"}, {"internal_eval", "boundar"}, {"failing_eval", "eval"}, {"missing_id", "eval"}} {
		if !hasDoctorDiagnostic(report, "C14", false, tc.name, tc.hook) {
			t.Errorf("missed %s canary not diagnosed with hook: %s", tc.name, doctorText(report))
		}
	}
	if !doctorErrors(report) {
		t.Errorf("zero hooks accepted: %s", doctorText(report))
	}
	if got := worktrees(f); got != before {
		t.Errorf("failed canary leaked worktree: %s", got)
	}

	g := localDoctorFixture(t)
	g.put(".fireweed/boundary.sh", fmt.Sprintf(`#!/bin/sh
find packages -name SPEC.md | while IFS= read -r spec; do
  dir=${spec%%/SPEC.md}; name=$(sed -n 's/^name: //p' "$spec" | head -1)
  if test -f "$dir/undeclared.txt"; then printf '%%s' "$name" > .fireweed/recovery-seen; printf '%%s: undeclared import\n' "$name"; fi
done
if test -f .fireweed/recovery-seen && ! find packages -name undeclared.txt | grep -q .; then printf '%%s: lingering violation\n' "$(cat .fireweed/recovery-seen)"; fi
`))
	g.commit()
	r := checks.Doctor(context.Background(), g.request(true))
	if !hasDoctorDiagnostic(r, "C14", false, "recovery", "undeclared", "boundar") {
		t.Errorf("missing failed recovery diagnostic: %s", doctorText(r))
	}
}

func Test_CHK_055_CHK_058_CHK_059_EvalEvidenceIsSpecific(t *testing.T) {
	for _, tc := range []struct{ name, script, canary string }{
		{"missing baseline coverage", `#!/bin/sh
printf '<testsuite/>' > "$3"
`, "baseline"},
		{"failing eval compile error without report", `#!/bin/sh
if test -f "$2/evals/failing.txt"; then echo 'compile error' >&2; exit 2; fi
prefix=$(sed -n 's/^id_prefix: //p' "$2/SPEC.md" | head -1)
printf '<testsuite><testcase name="%s-001"/></testsuite>' "$prefix" > "$3"
`, "failing_eval"},
		{"failing eval with unrelated failure", `#!/bin/sh
if test -f "$2/evals/failing.txt"; then printf '<testsuite><testcase name="ALP-001"><failure/></testcase></testsuite>' > "$3"; exit 1; fi
prefix=$(sed -n 's/^id_prefix: //p' "$2/SPEC.md" | head -1)
printf '<testsuite><testcase name="%s-001"/></testsuite>' "$prefix" > "$3"
`, "failing_eval"},
		{"missing id with unrelated failure", `#!/bin/sh
prefix=$(sed -n 's/^id_prefix: //p' "$2/SPEC.md" | head -1)
if test -f "$2/evals/missing.txt"; then printf '<testsuite><testcase name="%s-001"><failure/></testcase></testsuite>' "$prefix" > "$3"; exit 1; fi
printf '<testsuite><testcase name="%s-001"/></testsuite>' "$prefix" > "$3"
`, "missing_id"},
		{"missing id without report", `#!/bin/sh
if test -f "$2/evals/missing.txt"; then echo 'hook failed' >&2; exit 2; fi
prefix=$(sed -n 's/^id_prefix: //p' "$2/SPEC.md" | head -1)
printf '<testsuite><testcase name="%s-001"/></testsuite>' "$prefix" > "$3"
`, "missing_id"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			f := localDoctorFixture(t)
			f.put(".fireweed/eval.sh", tc.script)
			f.commit()
			r := checks.Doctor(context.Background(), f.request(true))
			if !hasDoctorDiagnostic(r, "C14", false, tc.canary, "eval") {
				t.Errorf("unrelated eval outcome accepted: %s", doctorText(r))
			}
		})
	}
}

func Test_CHK_055_BaselineAndRecoveryFailures(t *testing.T) {
	for _, tc := range []struct{ name, path, script, phase, hook string }{
		{"baseline eval failure", ".fireweed/eval.sh", "#!/bin/sh\necho baseline-eval-failed >&2\nexit 9\n", "baseline", "eval"},
		{"baseline boundary failure", ".fireweed/boundary.sh", "#!/bin/sh\nprintf 'alpha: baseline violation\\n'\n", "baseline", "boundar"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			f := localDoctorFixture(t)
			f.put(tc.path, tc.script)
			f.commit()
			r := checks.Doctor(context.Background(), f.request(true))
			if !hasDoctorDiagnostic(r, "C14", false, tc.phase, tc.hook) {
				t.Errorf("failed baseline accepted: %s", doctorText(r))
			}
		})
	}
	for _, tc := range []struct{ name, failure string }{
		{"recovery eval failure", "echo recovery-eval-failed >&2; exit 9"},
		{"recovery coverage failure", "printf '<testsuite/>' > \"$3\"; exit 0"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			f := localDoctorFixture(t)
			original := string(mustRead(t, filepath.Join(f.root, ".fireweed/eval.sh")))
			injection := "if test -f \"$2/evals/failing.txt\"; then printf '%s' \"$2\" > .fireweed/eval-recovery-module; fi\n" +
				"if test -f .fireweed/eval-recovery-module && test \"$(cat .fireweed/eval-recovery-module)\" = \"$2\" && ! test -f \"$2/evals/failing.txt\"; then " + tc.failure + "; fi\n"
			f.put(".fireweed/eval.sh", strings.Replace(original, "#!/bin/sh\n", "#!/bin/sh\n"+injection, 1))
			f.commit()
			r := checks.Doctor(context.Background(), f.request(true))
			if !hasDoctorDiagnostic(r, "C14", false, "recovery", "failing_eval", "eval") {
				t.Errorf("failed %s accepted: %s", tc.name, doctorText(r))
			}
		})
	}
}

func Test_CHK_056_CHK_057_BoundaryViolationMustNameCanaryModule(t *testing.T) {
	f := localDoctorFixture(t)
	f.put(".fireweed/boundary.sh", `#!/bin/sh
if find packages \( -name undeclared.txt -o -name internal.txt \) | grep -q .; then printf 'alpha: unrelated violation\n'; fi
`)
	f.commit()
	r := checks.Doctor(context.Background(), f.request(true))
	for _, name := range []string{"undeclared", "internal_eval"} {
		if !hasDoctorDiagnostic(r, "C14", false, name, "boundar") {
			t.Errorf("unrelated module violation accepted for %s: %s", name, doctorText(r))
		}
	}
}

func Test_CHK_057_InternalEvalRequiresInterfaceViolation(t *testing.T) {
	f := localDoctorFixture(t)
	f.put(".fireweed/boundary.sh", `#!/bin/sh
find packages -name SPEC.md | while IFS= read -r spec; do
  dir=${spec%/SPEC.md}; name=$(sed -n 's/^name: //p' "$spec" | head -1)
  if test -f "$dir/undeclared.txt"; then printf '%s: undeclared import\n' "$name"; fi
  if test -f "$dir/evals/internal.txt"; then printf '%s: unrelated boundary violation\n' "$name"; fi
done
`)
	f.commit()
	r := checks.Doctor(context.Background(), f.request(true))
	if !hasDoctorDiagnostic(r, "C14", false, "internal_eval", "boundar") || hasDoctorCheck(r, "internal_eval", true) {
		t.Errorf("same-module non-interface violation accepted: %s", doctorText(r))
	}
	if !hasDoctorCheck(r, "undeclared", true) {
		t.Errorf("fixture did not detect the independent undeclared canary: %s", doctorText(r))
	}
}

func Test_CHK_055_CHK_056_CHK_057_MissingRequiredBoundaryHook(t *testing.T) {
	f := localDoctorFixture(t)
	f.editConfig(func(config map[string]any) { delete(config["hooks"].(map[string]any), "boundaries") })
	f.put("packages/alpha/internal/live.txt", "disposable implementation\n")
	f.commit()
	report := checks.Doctor(context.Background(), f.request(true))
	if !hasDoctorDiagnostic(report, "X02", false, "alpha", "boundar") {
		t.Errorf("required boundary hook was not diagnosed: %s", doctorText(report))
	}
	if !hasDoctorDiagnostic(report, "C14", false, "baseline", "boundar") {
		t.Errorf("missing boundary hook was accepted as a clean baseline: %s", doctorText(report))
	}
	for _, canary := range []string{"undeclared", "internal_eval"} {
		if hasDoctorCheck(report, canary, true) {
			t.Errorf("missing boundary hook counted as %s detection: %s", canary, doctorText(report))
		}
		if !hasDoctorDiagnostic(report, "C14", false, canary, "boundar") {
			t.Errorf("missing boundary hook lacked named %s canary failure: %s", canary, doctorText(report))
		}
	}
}

func Test_CHK_055_MissingRequiredHookBesideOverride(t *testing.T) {
	f := localDoctorFixture(t)
	f.editConfig(func(config map[string]any) { delete(config["hooks"].(map[string]any), "boundaries") })
	f.put(".fireweed/override-boundary.sh", fmt.Sprintf("#!/bin/sh\nprintf 'override-boundaries:%%s\\n' \"$PWD\" >> %q\nexec sh .fireweed/boundary.sh\n", f.log))
	f.put("packages/alpha/SPEC.md", strings.Replace(string(mustRead(t, filepath.Join(f.root, "packages/alpha/SPEC.md"))), "depends_on: []", "depends_on: []\nhooks:\n  boundaries: sh .fireweed/override-boundary.sh", 1))
	f.put("packages/beta/SPEC.md", "---\nname: beta\nid_prefix: BET\ndepends_on: []\n---\n# Beta\n\n## Specs\n\n### BET-001\nThe fixture MUST pass.\n")
	f.put("packages/beta/internal/live.txt", "disposable implementation\n")
	f.commit()
	report := checks.Doctor(context.Background(), f.request(true))
	if !hasDoctorDiagnostic(report, "X02", false, "beta", "boundar") {
		t.Errorf("module without its own hook escaped required boundary enforcement: %s", doctorText(report))
	}
	if !strings.Contains(doctorLog(f), "override-boundaries:") {
		t.Errorf("applicable override was not run: %s", doctorLog(f))
	}
}

func Test_CHK_055_ApplicableOverrideNeedsNoGlobalHook(t *testing.T) {
	f := localDoctorFixture(t)
	f.editConfig(func(config map[string]any) { delete(config["hooks"].(map[string]any), "boundaries") })
	f.put(".fireweed/override-boundary.sh", "#!/bin/sh\nexec sh .fireweed/boundary.sh\n")
	f.put("packages/alpha/SPEC.md", strings.Replace(string(mustRead(t, filepath.Join(f.root, "packages/alpha/SPEC.md"))), "depends_on: []", "depends_on: []\nhooks:\n  boundaries: sh .fireweed/override-boundary.sh", 1))
	f.commit()
	report := checks.Doctor(context.Background(), f.request(true))
	if doctorErrors(report) || !hasDoctorCheck(report, "undeclared", true) || !hasDoctorCheck(report, "internal_eval", true) {
		t.Errorf("applicable override was rejected without a global hook: %s", doctorText(report))
	}
}

func Test_CHK_060_CHK_062_ModuleOverridesAndBoundaryWarnings(t *testing.T) {
	f := localDoctorFixture(t)
	f.put(".fireweed/override-eval.sh", fmt.Sprintf("#!/bin/sh\nprintf 'override-eval:%%s\\n' \"$PWD\" >> %q\nexec sh .fireweed/eval.sh \"$@\"\n", f.log))
	f.put(".fireweed/override-boundary.sh", fmt.Sprintf("#!/bin/sh\nprintf 'override-boundaries:%%s:%%s\\n' \"$PWD\" \"$1\" >> %q\nexec sh .fireweed/boundary.sh\n", f.log))
	f.put("packages/alpha/SPEC.md", "---\nname: alpha\nid_prefix: ALP\ndepends_on: []\nhooks:\n  eval: sh .fireweed/override-eval.sh {module} {module_dir} {report}\n  boundaries: sh .fireweed/override-boundary.sh {module}\n---\n# Alpha\n\n## Specs\n\n### ALP-001\nThe fixture MUST pass.\n")
	f.commit()
	report := checks.Doctor(context.Background(), f.request(true))
	if doctorErrors(report) || strings.Count(doctorLog(f), "override-eval:") < 2 || strings.Count(doctorLog(f), "override-boundaries:") < 2 {
		t.Errorf("override probes absent: %s", doctorLog(f))
	}
	g := localDoctorFixture(t)
	g.editConfig(func(config map[string]any) { delete(config["hooks"].(map[string]any), "boundaries") })
	g.put("packages/alpha/SPEC.md", strings.Replace(string(mustRead(t, filepath.Join(g.root, "packages/alpha/SPEC.md"))), "depends_on: []", "depends_on: []\nboundaries_opt_out: legacy", 1))
	g.commit()
	r := checks.Doctor(context.Background(), g.request(true))
	if doctorErrors(r) || !hasDoctorDiagnostic(r, "X03", true, "legacy") || !hasDoctorDiagnostic(r, "X04", true, "no boundary enforcement required") {
		t.Errorf("optional boundaries warning lost: %s", doctorText(r))
	}
}

func Test_CHK_060_OverridesMustDetectTheirCanaries(t *testing.T) {
	for _, tc := range []struct {
		name, hook, command, script, canary string
	}{
		{"eval", "eval", "sh .fireweed/override-eval.sh {module} {module_dir} {report}", `#!/bin/sh
prefix=$(sed -n 's/^id_prefix: //p' "$2/SPEC.md" | head -1)
printf '<testsuite><testcase name="%s-001"/></testsuite>' "$prefix" > "$3"
`, "failing_eval"},
		{"boundary", "boundaries", "sh .fireweed/override-boundary.sh {module}", "#!/bin/sh\nexit 0\n", "internal_eval"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			f := localDoctorFixture(t)
			path := ".fireweed/override-" + tc.name + ".sh"
			f.put(path, tc.script)
			f.put("packages/alpha/SPEC.md", strings.Replace(string(mustRead(t, filepath.Join(f.root, "packages/alpha/SPEC.md"))), "depends_on: []", "depends_on: []\nhooks:\n  "+tc.hook+": "+tc.command, 1))
			f.commit()
			r := checks.Doctor(context.Background(), f.request(true))
			if !hasDoctorDiagnostic(r, "C14", false, tc.canary, tc.name) {
				t.Errorf("broken %s override missed its %s canary without C14: %s", tc.name, tc.canary, doctorText(r))
			}
		})
	}
}

func Test_CHK_063_CHK_064_CHK_065_CHK_066_CHK_067_AgentProbes(t *testing.T) {
	f := localDoctorFixture(t)
	f.commit()
	for _, d := range project.Check(f.root) {
		if !d.Warning {
			t.Fatalf("invalid agent fixture: %+v", d)
		}
	}
	r := checks.Doctor(context.Background(), f.request(false))
	log := doctorLog(f)
	if doctorErrors(r) || strings.Count(log, "worker:") != 1 || strings.Count(log, "audit:") != 1 || strings.Count(log, "wrap:") != 2 || !strings.Contains(log, "FIREWEED_DOCTOR_OK") {
		t.Errorf("worker/audit/wrap probes: %s; log=%s", doctorText(r), log)
	}
	auditPrompt := false
	for _, line := range strings.Split(log, "\n") {
		if fields := strings.SplitN(line, ":", 3); len(fields) == 3 && fields[0] == "audit" && strings.Contains(fields[2], "FIREWEED_DOCTOR_OK") {
			auditPrompt = true
		}
	}
	if !auditPrompt {
		t.Errorf("audit did not receive its fixed-reply prompt: %s", log)
	}
	for _, part := range []string{"/fireweed/worktrees/", "worker:", "audit:"} {
		if !strings.Contains(log, part) {
			t.Errorf("scratch/role marker %s absent: %s", part, log)
		}
	}

	g := localDoctorFixture(t)
	g.editConfig(func(config map[string]any) {
		delete(config["audit"].(map[string]any), "command")
		config["agent"].(map[string]any)["command"] = "sh .fireweed/agent.sh {prompt} {git_common_dir}"
	})
	g.commit()
	rr := checks.Doctor(context.Background(), g.request(false))
	if !hasDoctorDiagnostic(rr, "C15", true, "shared") || strings.Count(doctorLog(g), "worker:") != 2 {
		t.Errorf("audit fallback: %s; log=%s", doctorText(rr), doctorLog(g))
	}

	h := localDoctorFixture(t)
	h.put(".fireweed/agent.sh", fmt.Sprintf("#!/bin/sh\nfor i in $(seq 1 25); do printf 'line-%%02d\\n' \"$i\"; done\necho RATE LIMIT reached\nexit 9\n"))
	h.commit()
	bad := checks.Doctor(context.Background(), h.request(false))
	if !hasDoctorDiagnostic(bad, "C15", false, "usage limit", "rate limit", "line-25") {
		t.Errorf("usage-limit tail missing: %s", doctorText(bad))
	}
	if strings.Contains(doctorText(bad), "line-01") {
		t.Errorf("agent output was not trimmed to final 20 lines: %s", doctorText(bad))
	}
	outOfTail := localDoctorFixture(t)
	outOfTail.put(".fireweed/agent.sh", "#!/bin/sh\necho RATE LIMIT early\nfor i in $(seq 1 25); do printf 'later-%02d\\n' \"$i\"; done\nexit 9\n")
	outOfTail.commit()
	ordinary := checks.Doctor(context.Background(), outOfTail.request(false))
	if !hasDoctorDiagnostic(ordinary, "C15", false, "agent.sh", "later-25") {
		t.Errorf("ordinary out-of-tail failure absent: %s", doctorText(ordinary))
	}
	for _, d := range ordinary.Diagnostics {
		if d.Code == "C15" && !d.Warning && strings.Contains(strings.ToLower(d.Message), "usage limit") {
			t.Errorf("out-of-tail pattern classified as usage limit: %s", doctorText(ordinary))
		}
	}
	for _, tc := range []struct{ name, script string }{{"wrong reply", "#!/bin/sh\necho WRONG_REPLY\n"}, {"nonzero with reply", "#!/bin/sh\necho FIREWEED_DOCTOR_OK\nexit 7\n"}, {"ordinary failure", "#!/bin/sh\nfor i in $(seq 1 25); do printf 'ordinary-%02d\\n' \"$i\"; done\nexit 8\n"}} {
		t.Run(tc.name, func(t *testing.T) {
			q := localDoctorFixture(t)
			q.put(".fireweed/agent.sh", tc.script)
			q.commit()
			r := checks.Doctor(context.Background(), q.request(false))
			if !hasDoctorDiagnostic(r, "C15", false, "agent.sh") {
				t.Errorf("agent failure accepted: %s", doctorText(r))
			}
			if tc.name == "ordinary failure" && (!hasDoctorDiagnostic(r, "C15", false, "ordinary-25") || strings.Contains(doctorText(r), "ordinary-01")) {
				t.Errorf("ordinary failure tail incorrect: %s", doctorText(r))
			}
		})
	}
	wrongAudit := localDoctorFixture(t)
	wrongAudit.put(".fireweed/audit.sh", "#!/bin/sh\necho WRONG_REPLY\n")
	wrongAudit.commit()
	auditFailure := checks.Doctor(context.Background(), wrongAudit.request(false))
	if !hasDoctorDiagnostic(auditFailure, "C15", false, "audit.sh") || hasDoctorCheck(auditFailure, "audit", true) {
		t.Errorf("wrong audit reply accepted: %s", doctorText(auditFailure))
	}

	i := localDoctorFixture(t)
	i.put(".fireweed/audit.sh", "#!/bin/sh\nsleep 2\n")
	i.commit()
	req := i.request(false)
	req.Timeout = 100 * time.Millisecond
	start := time.Now()
	timed := checks.Doctor(context.Background(), req)
	if !hasDoctorDiagnostic(timed, "C15", false, "audit") || time.Since(start) > 4*time.Second {
		t.Errorf("audit timeout not bounded: %s; elapsed=%s", doctorText(timed), time.Since(start))
	}

	defaulted := localDoctorFixture(t)
	defaulted.put(".fireweed/agent.sh", "#!/bin/sh\nsleep 0.2\nprintf 'FIREWEED_DOCTOR_OK\\n'\n")
	defaulted.put(".fireweed/audit.sh", "#!/bin/sh\nsleep 0.2\nprintf 'FIREWEED_DOCTOR_OK\\n'\n")
	defaulted.commit()
	defaultRequest := defaulted.request(false)
	defaultRequest.Timeout = 0
	defaultReport := checks.Doctor(context.Background(), defaultRequest)
	if doctorErrors(defaultReport) || !hasDoctorCheck(defaultReport, "worker", true) || !hasDoctorCheck(defaultReport, "audit", true) {
		t.Errorf("unset timeout did not give both agent probes their default allowance: %s", doctorText(defaultReport))
	}

	j := localDoctorFixture(t)
	j.commit()
	skip := checks.Doctor(context.Background(), j.request(true))
	if strings.Contains(doctorLog(j), "worker:") || strings.Contains(doctorLog(j), "audit:") || !hasDoctorCheck(skip, "undeclared", true) || !strings.Contains(doctorLog(j), j.bin) {
		t.Errorf("NoAgent skipped hooks or ran agent: %s; log=%s", doctorText(skip), doctorLog(j))
	}
}

func Test_CHK_065_UnsetTimeoutBoundsAuditProbe(t *testing.T) {
	f := localDoctorFixture(t)
	f.put(".fireweed/audit.sh", fmt.Sprintf("#!/bin/sh\nprintf 'audit:%%s:%%s\n' \"$PWD\" \"$1\" >> %q\nprintf 'audit-timeout-started\n'\nsleep 210\nprintf 'FIREWEED_DOCTOR_OK\n'\n", f.log))
	f.commit()
	req := f.request(false)
	req.Timeout = 0

	start := time.Now()
	report := checks.Doctor(context.Background(), req)
	elapsed := time.Since(start)
	log := doctorLog(f)
	if !hasDoctorCheck(report, "worker", true) || strings.Count(log, "worker:") != 1 || strings.Count(log, "audit:") != 1 {
		t.Errorf("agent probes did not run once with immediate worker reply: %s; log=%s", doctorText(report), log)
	}
	if !hasDoctorCheck(report, "audit", false) || !hasDoctorDiagnostic(report, "C15", false, "audit.sh", "audit-timeout-started") {
		t.Errorf("long audit probe lacked a C15 timeout with its command and output: %s", doctorText(report))
	}
	if elapsed < 165*time.Second || elapsed > 200*time.Second {
		t.Errorf("unset audit timeout took %s, want a bound near three minutes", elapsed)
	}
}

func Test_CHK_068_CHK_069_ManagedGuidanceAndPlaybookBase(t *testing.T) {
	f := localDoctorFixture(t)
	text, err := assets.Playbook("evals")
	if err != nil {
		t.Fatal(err)
	}
	sum := sha256.Sum256([]byte(text))
	base := hex.EncodeToString(sum[:])
	f.put(".fireweed/playbooks/evals.md", "---\nembedded_base: "+base+"\n---\n# Local override\n")
	f.commit()
	r := checks.Doctor(context.Background(), f.request(true))
	if hasDoctorDiagnostic(r, "C16", true) {
		t.Errorf("current guidance/base warned: %s", doctorText(r))
	}
	f.put("AGENTS.md", "Before.\n"+doctorGuidance+"After.\n")
	surrounded := checks.Doctor(context.Background(), f.request(true))
	if hasDoctorDiagnostic(surrounded, "C16", true) {
		t.Errorf("current managed guidance with surrounding text warned: %s", doctorText(surrounded))
	}
	f.put("AGENTS.md", "Before.\n<!-- fireweed:start -->\nold guidance\n<!-- fireweed:end -->\nAfter.\n")
	f.put(".fireweed/playbooks/evals.md", "---\nembedded_base: stale\n---\n# Local override\n")
	stale := checks.Doctor(context.Background(), f.request(true))
	if !hasDoctorDiagnostic(stale, "C16", true, "AGENTS.md") || !hasDoctorDiagnostic(stale, "C16", true, "evals") {
		t.Errorf("stale guidance/base not warned: %s", doctorText(stale))
	}
	if err := os.Remove(filepath.Join(f.root, "AGENTS.md")); err != nil {
		t.Fatal(err)
	}
	f.put(".fireweed/playbooks/evals.md", "# No embedded base\n")
	missing := checks.Doctor(context.Background(), f.request(true))
	if !hasDoctorDiagnostic(missing, "C16", true, "AGENTS.md") || !hasDoctorDiagnostic(missing, "C16", true, "evals") {
		t.Errorf("missing guidance/base not warned: %s", doctorText(missing))
	}
}

func mustRead(t *testing.T, path string) []byte {
	t.Helper()
	b, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	return b
}
func doctorBinary(t *testing.T) string {
	t.Helper()
	_, file, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("caller")
	}
	root := filepath.Clean(filepath.Join(filepath.Dir(file), "..", "..", ".."))
	bin := filepath.Join(t.TempDir(), "fireweed-new")
	cmd := exec.Command("go", "build", "-o", bin, ".")
	cmd.Dir = root
	cmd.Env = append(os.Environ(), "GOCACHE=/tmp/fireweed-rebuild-cache")
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("build product: %v: %s", err, out)
	}
	return bin
}
