package evals

import (
	"fmt"
	checks "github.com/lurkshark/fireweed/packages/checks"
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"strings"
	"testing"
	"unicode"
)

func cliVerifyOutsideInvoke(t *testing.T, dir string, args ...string) result {
	t.Helper()
	scratch, e := filepath.EvalSymlinks(t.TempDir())
	if e != nil {
		t.Fatal(e)
	}
	for _, name := range []string{"TMPDIR", "TMP", "TEMP"} {
		t.Setenv(name, scratch)
	}
	files, temp, env := queueFiles(t, dir), queueFiles(t, scratch), os.Environ()
	cwd, e := os.Getwd()
	if e != nil {
		t.Fatal(e)
	}
	got := invoke(t, dir, args...)
	after, e := os.Getwd()
	if e != nil || after != cwd || !reflect.DeepEqual(env, os.Environ()) || !reflect.DeepEqual(files, queueFiles(t, dir)) || !reflect.DeepEqual(temp, queueFiles(t, scratch)) {
		t.Error("outside-project verify changed files/temp/parent state")
	}
	return got
}

// Retained mutations cannot use a no-change observer. Record the real public
// effect, then require the CLI to add exactly that effect to its own baseline.
type cliVerifyEffectWatch struct {
	fixture cliVerifyCase
	common  string
	temp    string
}

func cliVerifyWatchEffects(t *testing.T, f cliVerifyCase) cliVerifyEffectWatch {
	t.Helper()
	common := strings.TrimSpace(gitFixtureCommand(t, f.Root, "rev-parse", "--git-common-dir"))
	if !filepath.IsAbs(common) {
		common = filepath.Join(f.Root, common)
	}
	scratch, e := filepath.EvalSymlinks(t.TempDir())
	if e != nil {
		t.Fatal(e)
	}
	for _, name := range []string{"TMPDIR", "TMP", "TEMP"} {
		t.Setenv(name, scratch)
	}
	// The configured native empty commit must be identical in both public
	// calibrations and the CLI. This does not normalize an arbitrary new object.
	t.Setenv("GIT_AUTHOR_DATE", "2001-01-01T00:00:00Z")
	t.Setenv("GIT_COMMITTER_DATE", "2001-01-01T00:00:00Z")
	return cliVerifyEffectWatch{f, common, scratch}
}

func (w cliVerifyEffectWatch) capture(t *testing.T) map[string]map[string]string {
	t.Helper()
	plan := queueFiles(t, filepath.Dir(w.fixture.Plan))
	key := filepath.Base(w.fixture.Plan) + "/calls"
	if v, ok := plan[key]; ok {
		mode, _, regular := strings.Cut(v, " bytes:")
		if !regular {
			t.Fatal("hook call log is not a regular file")
		}
		// cliVerifyObserve separately checks every call, identity and report.
		// Only this exact observer file has volatile report/bin contents.
		plan[key] = mode + " bytes:<checked hook-call observer>"
	}
	return map[string]map[string]string{
		"files":  queueFiles(t, w.fixture.Root),
		"common": queueFiles(t, w.common),
		"plan":   plan,
		"temp":   queueFiles(t, w.temp),
	}
}

func cliVerifyReplayEffects(t *testing.T, before, after, current map[string]map[string]string) map[string]map[string]string {
	t.Helper()
	out := map[string]map[string]string{}
	for group, base := range current {
		out[group] = map[string]string{}
		for p, v := range base {
			out[group][p] = v
		}
		for p, v := range before[group] {
			if _, ok := after[group][p]; !ok {
				delete(out[group], p)
			} else if after[group][p] == v {
				continue
			}
		}
		for p, v := range after[group] {
			old, existed := before[group][p]
			if existed && old == v {
				continue
			}
			// Git journals append one known native-hook entry. Restoration
			// between runs legitimately adds its own entries to the prefix.
			if group == "common" && strings.HasPrefix(p, "logs/") {
				prefix, tail, regular := strings.Cut(v, " bytes:")
				_, oldTail, oldRegular := strings.Cut(old, " bytes:")
				curPrefix, curTail, curRegular := strings.Cut(base[p], " bytes:")
				if regular && curRegular && (!existed || oldRegular) {
					if !strings.HasPrefix(tail, oldTail) || !strings.HasPrefix(curTail, oldTail) || curPrefix != prefix {
						t.Fatal("native Git journal replay premise", p)
					}
					out[group][p] = prefix + " bytes:" + curTail + strings.TrimPrefix(tail, oldTail)
					continue
				}
			}
			out[group][p] = v
		}
	}
	return out
}

func cliVerifyAssertEffects(t *testing.T, want, got map[string]map[string]string) {
	t.Helper()
	if reflect.DeepEqual(want, got) {
		return
	}
	for group := range want {
		for p, v := range want[group] {
			if got[group][p] != v {
				t.Errorf("verification effect differs from public oracle: %s/%s", group, p)
			}
		}
		for p := range got[group] {
			if _, ok := want[group][p]; !ok {
				t.Errorf("verification added an effect beyond public oracle: %s/%s", group, p)
			}
		}
	}
}

// Consume the complete public diagnostics while allowing presentation punctuation.
// Any remaining words reveal an extra diagnostic, prompt, path or hook output.
func cliVerifyPlainDiagnosticText(stderr string, ds []checks.Diagnostic, actualReports, reports []string) error {
	text := cliVerifyNormalize(stderr, actualReports).(string)
	messages, paths, codes, optionalCodes, severities := []string{}, []string{}, []string{}, []string{}, []string{}
	for _, d := range ds {
		messages = append(messages, cliVerifyNormalize(d.Message, reports).(string))
		if d.Path != "" {
			paths = append(paths, cliVerifyNormalize(d.Path, reports).(string))
		}
		if strings.Contains(d.Message, d.Code) {
			optionalCodes = append(optionalCodes, d.Code)
		} else {
			codes = append(codes, d.Code)
		}
		severities = append(severities, d.Severity)
	}
	for _, tokens := range [][]string{messages, paths, codes} {
		sort.SliceStable(tokens, func(i, j int) bool { return len(tokens[i]) > len(tokens[j]) })
		for _, token := range tokens {
			if token == "" {
				continue
			}
			at := strings.Index(text, token)
			if at < 0 {
				return fmt.Errorf("plain diagnostic lost %q in %q", token, stderr)
			}
			text = text[:at] + text[at+len(token):]
		}
	}
	for _, code := range optionalCodes {
		if code != "" {
			text = strings.Replace(text, code, "", 1)
		}
	}
	for _, severity := range severities {
		if severity != "" {
			text = strings.Replace(text, severity, "", 1)
		}
	}
	for _, r := range text {
		if !unicode.IsSpace(r) && !strings.ContainsRune(":[]()", r) {
			return fmt.Errorf("plain diagnostic added text %q", text)
		}
	}
	return nil
}
func cliVerifyFatalPlain(t *testing.T, got result, code, message string, exit int) {
	t.Helper()
	if got.code != exit || got.out != "" {
		t.Error("fatal plain exit/stdout", got)
	}
	if e := cliVerifyPlainDiagnosticText(got.err, []checks.Diagnostic{{Code: code, Severity: "error", Message: message}}, nil, nil); e != nil {
		t.Error("fatal discovery plain classification/streams", e, got)
	}
}

// Scope observer settings to one observation, rather than the enclosing test.
// A previous environment-based recorder must never become the next proxy's target.
func cliVerifyRestoreObserverEnv() func() {
	type value struct {
		key, value string
		exists     bool
	}
	old := []value{}
	for _, key := range []string{"TMPDIR", "TMP", "TEMP", "PATH", "FIREWEED_BIN", "GIT_TRACE2_EVENT", "GIT_TRACE2_ENV_VARS", "FW_VERIFY_REAL_GIT", "FW_VERIFY_GIT_LOCATIONS"} {
		v, ok := os.LookupEnv(key)
		old = append(old, value{key, v, ok})
	}
	return func() {
		for _, v := range old {
			if v.exists {
				os.Setenv(v.key, v.value)
			} else {
				os.Unsetenv(v.key)
			}
		}
	}
}
