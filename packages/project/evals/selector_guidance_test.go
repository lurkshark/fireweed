package evals_test

import (
	"errors"
	"strings"
	"testing"

	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/project"
)

func pendingSelectorGuidance(t *testing.T, err error, selector string) {
	t.Helper()
	if err == nil {
		t.Fatal("unknown pending selector was accepted")
	}
	if code, subject, ok := contracts.ErrorCode(err); !ok || code != "U02" || subject != selector {
		t.Errorf("error code = (%q, %q, %v); want (U02, %q, true): %v", code, subject, ok, selector, err)
	}
	var coded contracts.CodedError
	if !errors.As(err, &coded) || coded.Cause == nil {
		t.Errorf("lookup error = %v; want a retained U02 cause", err)
	}
	message := err.Error()
	for _, want := range []string{selector, "first", "second"} {
		if !strings.Contains(message, want) {
			t.Errorf("lookup error %q lacks rejected selector or pending slug %q", message, want)
		}
	}
	for _, unwanted := range []string{"0001-first.md", "0002-second.md", "archived"} {
		if strings.Contains(message, unwanted) {
			t.Errorf("lookup guidance %q includes queued filename or archived slug %q", message, unwanted)
		}
	}
}

func emptyPendingSelectorGuidance(t *testing.T, err error, selector string) {
	t.Helper()
	if code, subject, ok := contracts.ErrorCode(err); !ok || code != "U02" || subject != selector {
		t.Fatalf("empty queue error = (%q, %q, %v); want (U02, %q, true): %v", code, subject, ok, selector, err)
	}
	message := strings.ToLower(err.Error())
	if !strings.Contains(message, selector) {
		t.Errorf("empty queue error %q lacks rejected selector", message)
	}
	if !strings.Contains(message, "no pending") && !strings.Contains(message, "no valid pending") && !strings.Contains(message, "none") && !strings.Contains(message, "empty") && !strings.Contains(message, "[]") {
		t.Errorf("empty queue error %q lacks empty-set guidance", message)
	}
}

func Test_PRJ_004_UnknownProjectAndStampSelectorsListPendingSlugs(t *testing.T) {
	root := fixture(t)
	queued(t, root, "0001-first.md", "", "First.\n", "", "", "")
	queued(t, root, "0002-second.md", "", "Second.\n", "", "", "")
	put(t, root, "changes/archive/0003-archived.md", "archived material")
	for _, operation := range []struct {
		name, selector string
		call           func(string) error
	}{
		{"Project", "missing", func(selector string) error { _, err := project.Project(root, selector); return err }},
		{"Stamp", "0099-absent.md", func(selector string) error { return project.Stamp(root, selector) }},
	} {
		t.Run(operation.name, func(t *testing.T) {
			pendingSelectorGuidance(t, operation.call(operation.selector), operation.selector)
		})
	}

	emptyRoot := fixture(t)
	_, err := project.Project(emptyRoot, "missing")
	emptyPendingSelectorGuidance(t, err, "missing")
}

func Test_PRJ_066_UnknownStaleSelectorsListPendingSlugs(t *testing.T) {
	root := fixture(t)
	queued(t, root, "0001-first.md", "", "First.\n", "", "", "")
	queued(t, root, "0002-second.md", "", "Second.\n", "", "", "")
	put(t, root, "changes/archive/0003-archived.md", "archived material")
	for _, operation := range []struct {
		name, selector string
		call           func(string) error
	}{
		{"StaleSpecs", "missing", func(selector string) error { _, err := project.StaleSpecs(root, selector); return err }},
		{"StampMissing", "changes/0099-absent.md", func(selector string) error { return project.StampMissing(root, selector) }},
	} {
		t.Run(operation.name, func(t *testing.T) {
			pendingSelectorGuidance(t, operation.call(operation.selector), operation.selector)
		})
	}

	emptyRoot := fixture(t)
	_, err := project.StaleSpecs(emptyRoot, "missing")
	emptyPendingSelectorGuidance(t, err, "missing")
}
