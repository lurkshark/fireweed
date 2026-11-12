package evals_test

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/lurkshark/fireweed/packages/checks"
)

func doctorProbeLog(t *testing.T, f doctorFixture, report checks.DoctorReport, role string, markers ...string) string {
	t.Helper()
	var detail string
	for _, check := range report.Checks {
		if check.Name == role {
			detail = check.Detail
			break
		}
	}
	if detail == "" {
		t.Errorf("executed %s probe has no check: %s", role, doctorText(report))
		return ""
	}
	at := strings.Index(detail, "log: ")
	if at < 0 {
		t.Errorf("%s detail lacks log path: %q", role, detail)
		return ""
	}
	fields := strings.Fields(detail[at+len("log: "):])
	if len(fields) == 0 {
		t.Errorf("%s detail has empty log path: %q", role, detail)
		return ""
	}
	path := fields[0]
	if !filepath.IsAbs(path) {
		t.Errorf("%s log path is not absolute: %q", role, path)
	}
	rel, err := filepath.Rel(filepath.Join(f.root, ".fireweed", "runs"), path)
	if err != nil {
		t.Errorf("%s log path %q: %v", role, path, err)
		return ""
	}
	parts := strings.Split(filepath.ToSlash(rel), "/")
	if len(parts) != 2 || parts[0] == "" || parts[0] == "." || parts[0] == ".." || parts[1] != "doctor-"+role+"-1.log" {
		t.Errorf("%s log path = %q, want .fireweed/runs/<run>/doctor-%s-1.log", role, path, role)
	}
	data, err := os.ReadFile(path)
	if err != nil {
		t.Errorf("read %s log %q: %v", role, path, err)
		return path
	}
	for _, marker := range markers {
		if !strings.Contains(string(data), marker) {
			t.Errorf("%s log lacks %q: %q", role, marker, data)
		}
	}
	return path
}

func doctorLogRun(t *testing.T, f doctorFixture, worker, audit string, report checks.DoctorReport) string {
	t.Helper()
	base := filepath.Join(f.root, ".fireweed", "runs")
	w, ew := filepath.Rel(base, worker)
	a, ea := filepath.Rel(base, audit)
	if ew != nil || ea != nil {
		t.Errorf("log paths outside run root: %q %q: %v %v", worker, audit, ew, ea)
		return ""
	}
	wRun, aRun := filepath.Dir(w), filepath.Dir(a)
	if wRun != aRun || wRun == "." || wRun == ".." {
		t.Errorf("worker and audit have different run identities: %q %q: %s", worker, audit, doctorText(report))
	}
	return wRun
}

func doctorCallerClean(t *testing.T, f doctorFixture, status, registrations string) {
	t.Helper()
	if got := f.git("status", "--porcelain", "--untracked-files=all"); got != status {
		t.Errorf("caller status changed: before %q after %q", status, got)
	}
	if got := worktrees(f); got != registrations {
		t.Errorf("scratch registration changed: before %q after %q", registrations, got)
	}
	assertScratchRemoved(t, f)
}

func Test_CHK_081_DoctorProbeLogsShareFreshRunAndPreserveStreams(t *testing.T) {
	f := localDoctorFixture(t)
	f.put(".fireweed/agent.sh", "#!/bin/bash\nprintf 'worker stdout Ω without newline'; printf 'worker stderr β without newline' >&2; printf 'FIREWEED_DOCTOR_OK\\n'\n")
	f.put(".fireweed/audit.sh", "#!/bin/bash\nprintf 'audit stdout Ω without newline'; printf 'audit stderr β without newline' >&2; printf 'FIREWEED_DOCTOR_OK\\n'\n")
	f.commit()
	status, registrations := f.git("status", "--porcelain", "--untracked-files=all"), worktrees(f)
	var firstRun string
	for i := 0; i < 2; i++ {
		report := checks.Doctor(context.Background(), f.request(false))
		if doctorErrors(report) || !hasDoctorCheck(report, "worker", true) || !hasDoctorCheck(report, "audit", true) {
			t.Errorf("successful probes changed outcome: %s", doctorText(report))
		}
		worker := doctorProbeLog(t, f, report, "worker", "worker stdout Ω without newlineFIREWEED_DOCTOR_OK\n", "worker stderr β without newline")
		audit := doctorProbeLog(t, f, report, "audit", "audit stdout Ω without newlineFIREWEED_DOCTOR_OK\n", "audit stderr β without newline")
		if worker != "" && audit != "" {
			run := doctorLogRun(t, f, worker, audit, report)
			if i == 1 && run == firstRun {
				t.Errorf("consecutive Doctor invocations reused run identity %q", run)
			}
			firstRun = run
		}
		doctorCallerClean(t, f, status, registrations)
	}
}

func Test_CHK_081_DoctorPersistsFailedAndTimedOutProbeOutput(t *testing.T) {
	failureStdout := "early worker stdout Ω\n"
	for i := 1; i <= 25; i++ {
		failureStdout += fmt.Sprintf("later-%02d\n", i)
	}
	for _, tc := range []struct {
		name, role, script, stdout, stderr string
		timeout                            time.Duration
	}{
		{"failure", "worker", "#!/bin/bash\nprintf 'early worker stdout Ω\\n'; for i in $(seq 1 25); do printf 'later-%02d\\n' \"$i\"; done; printf 'failed worker stderr β' >&2; exit 7\n", failureStdout, "failed worker stderr β", 2 * time.Second},
		{"timeout", "audit", "#!/bin/bash\nprintf 'partial audit stdout Ω'; printf 'partial audit stderr β' >&2; sleep 5\n", "partial audit stdout Ω", "partial audit stderr β", 150 * time.Millisecond},
	} {
		t.Run(tc.name, func(t *testing.T) {
			f := localDoctorFixture(t)
			if tc.role == "worker" {
				f.put(".fireweed/agent.sh", tc.script)
			} else {
				f.put(".fireweed/audit.sh", tc.script)
			}
			f.commit()
			status, registrations := f.git("status", "--porcelain", "--untracked-files=all"), worktrees(f)
			req := f.request(false)
			req.Timeout = tc.timeout
			start := time.Now()
			report := checks.Doctor(context.Background(), req)
			if time.Since(start) > 4*time.Second {
				t.Errorf("%s probe did not honor bounded request timeout", tc.name)
			}
			if !hasDoctorCheck(report, tc.role, false) || !hasDoctorDiagnostic(report, "C15", false, tc.role, tc.stderr) {
				t.Errorf("%s probe outcome/diagnostic lost: %s", tc.name, doctorText(report))
			}
			path := doctorProbeLog(t, f, report, tc.role, tc.stdout, tc.stderr)
			if tc.name == "failure" && path != "" {
				if !strings.Contains(string(mustRead(t, path)), "later-25") || strings.Contains(doctorText(report), "early worker stdout Ω") {
					t.Errorf("failed probe log or final-20-line diagnostic changed: %s", doctorText(report))
				}
			}
			doctorCallerClean(t, f, status, registrations)
		})
	}
}

func Test_CHK_081_DoctorPersistsCancelledProbeOutput(t *testing.T) {
	f := localDoctorFixture(t)
	ready := filepath.Join(t.TempDir(), "worker-ready")
	f.put(".fireweed/agent.sh", fmt.Sprintf("#!/bin/bash\nprintf 'cancel worker stdout Ω'; printf 'cancel worker stderr β' >&2; printf ready > %q; sleep 5\n", ready))
	f.commit()
	status, registrations := f.git("status", "--porcelain", "--untracked-files=all"), worktrees(f)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	reports := make(chan checks.DoctorReport, 1)
	go func() { reports <- checks.Doctor(ctx, f.request(false)) }()
	deadline := time.After(8 * time.Second)
	for {
		if _, err := os.Stat(ready); err == nil {
			break
		} else if !os.IsNotExist(err) {
			t.Fatal(err)
		}
		select {
		case report := <-reports:
			t.Fatalf("worker exited before cancellation: %s", doctorText(report))
		case <-deadline:
			t.Fatal("worker never signalled readiness")
		case <-time.After(10 * time.Millisecond):
		}
	}
	cancel()
	select {
	case report := <-reports:
		if !hasDoctorCheck(report, "worker", false) || !hasDoctorDiagnostic(report, "C15", false, "worker") {
			t.Errorf("cancelled probe outcome lost: %s", doctorText(report))
		}
		doctorProbeLog(t, f, report, "worker", "cancel worker stdout Ω", "cancel worker stderr β")
	case <-time.After(4 * time.Second):
		t.Fatal("cancelled Doctor did not return promptly")
	}
	doctorCallerClean(t, f, status, registrations)
}
