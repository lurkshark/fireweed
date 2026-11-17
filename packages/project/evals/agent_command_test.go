package evals_test

import (
	"reflect"
	"strings"
	"testing"

	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/project"
)

const workerShellTemplate = "worker-tool --message {prompt}"

func namedAgentConfig() map[string]any {
	return map[string]any{
		"agent": map[string]any{
			"default": "standard",
			"commands": map[string]any{
				"standard": workerShellTemplate,
				"reviewer": []any{"review-tool", "--message", "{prompt}"},
			},
		},
	}
}

func requireAgentCommand(t *testing.T, config map[string]any, role, alias string, want any, inherits bool) {
	t.Helper()
	got, inherited, err := project.ResolveAgentCommand(config, role, alias)
	if err != nil || !reflect.DeepEqual(got, want) || inherited != inherits {
		t.Fatalf("ResolveAgentCommand(%q, %q) = (%#v, %t, %v); want (%#v, %t, nil)",
			role, alias, got, inherited, err, want, inherits)
	}
}

func requireAgentCommandError(t *testing.T, config map[string]any, role, alias, wantCode string, settings ...string) {
	t.Helper()
	_, _, err := project.ResolveAgentCommand(config, role, alias)
	code, subject, ok := contracts.ErrorCode(err)
	if !ok || code != wantCode {
		t.Fatalf("ResolveAgentCommand(%q, %q) error = %v; want %s", role, alias, err, wantCode)
	}
	for _, setting := range settings {
		if strings.Contains(err.Error(), setting) || strings.Contains(subject, setting) {
			return
		}
	}
	t.Errorf("ResolveAgentCommand(%q, %q) error %q (subject %q) must identify one of %q", role, alias, err, subject, settings)
}

func Test_PRJ_086_ResolveExplicitWorkerAlias(t *testing.T) {
	t.Run("select_listed_alias", func(t *testing.T) {
		requireAgentCommand(t, namedAgentConfig(), "worker", "reviewer", []any{"review-tool", "--message", "{prompt}"}, false)
	})
	t.Run("unknown_alias", func(t *testing.T) {
		const alias = "unlisted-worker"
		_, _, err := project.ResolveAgentCommand(namedAgentConfig(), "worker", alias)
		if code, subject, ok := contracts.ErrorCode(err); !ok || code != "U02" || subject != alias {
			t.Errorf("ResolveAgentCommand error code = (%q, %q, %t); want (U02, %q, true): %v", code, subject, ok, alias, err)
		}
		if err == nil || !strings.Contains(err.Error(), alias) {
			t.Errorf("ResolveAgentCommand error = %v; want a message naming %q", err, alias)
		}
	})
	t.Run("malformed_mapping", func(t *testing.T) {
		config := namedAgentConfig()
		config["agent"].(map[string]any)["commands"] = []any{workerShellTemplate}
		requireAgentCommandError(t, config, "worker", "standard", "C01", "agent.commands")
	})
	for _, tc := range []struct {
		name    string
		command any
	}{
		{"empty_shell", ""},
		{"blank_shell", " \t"},
		{"empty_argv", []any{}},
		{"empty_executable", []any{"", "{prompt}"}},
		{"nonstring_argument", []any{"worker-tool", 42}},
		{"invalid_command_type", 42},
	} {
		t.Run(tc.name, func(t *testing.T) {
			config := namedAgentConfig()
			commands := config["agent"].(map[string]any)["commands"].(map[string]any)
			commands["broken-worker"] = tc.command
			requireAgentCommandError(t, config, "worker", "broken-worker", "C01", "broken-worker")
		})
	}
}

func Test_PRJ_087_ResolveNamedWorkerDefault(t *testing.T) {
	t.Run("select_named_default", func(t *testing.T) {
		requireAgentCommand(t, namedAgentConfig(), "worker", "", workerShellTemplate, false)
	})
	t.Run("missing_default", func(t *testing.T) {
		config := namedAgentConfig()
		delete(config["agent"].(map[string]any), "default")
		requireAgentCommandError(t, config, "worker", "", "C01", "agent.default")
	})
	for _, tc := range []struct {
		name      string
		reference any
	}{
		{"empty_default", ""},
		{"nonstring_default", 42},
		{"unlisted_default", "absent-worker"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			config := namedAgentConfig()
			config["agent"].(map[string]any)["default"] = tc.reference
			requireAgentCommandError(t, config, "worker", "", "C01", "agent.default")
		})
	}
	t.Run("invalid_default_command", func(t *testing.T) {
		config := namedAgentConfig()
		config["agent"].(map[string]any)["commands"].(map[string]any)["standard"] = []any{}
		requireAgentCommandError(t, config, "worker", "", "C01", "agent.default", "standard")
	})
}

func Test_PRJ_088_ResolveAuditAliasOrInheritWorkerDefault(t *testing.T) {
	t.Run("select_independent_audit_alias", func(t *testing.T) {
		config := namedAgentConfig()
		config["audit"] = map[string]any{"agent": "reviewer"}
		requireAgentCommand(t, config, "audit", "", []any{"review-tool", "--message", "{prompt}"}, false)
	})
	t.Run("explicitly_name_worker_default", func(t *testing.T) {
		config := namedAgentConfig()
		config["audit"] = map[string]any{"agent": "standard"}
		requireAgentCommand(t, config, "audit", "", workerShellTemplate, false)
	})
	t.Run("inherit_without_audit_section", func(t *testing.T) {
		requireAgentCommand(t, namedAgentConfig(), "audit", "", workerShellTemplate, true)
	})
	t.Run("inherit_without_audit_reference", func(t *testing.T) {
		config := namedAgentConfig()
		config["audit"] = map[string]any{"enabled": true}
		requireAgentCommand(t, config, "audit", "", workerShellTemplate, true)
	})
	for _, tc := range []struct {
		name      string
		reference any
	}{
		{"empty_audit_reference", ""},
		{"null_audit_reference", nil},
		{"nonstring_audit_reference", 42},
		{"unlisted_audit_reference", "absent-auditor"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			config := namedAgentConfig()
			config["audit"] = map[string]any{"agent": tc.reference}
			requireAgentCommandError(t, config, "audit", "", "C01", "audit.agent")
		})
	}
	t.Run("invalid_audit_command", func(t *testing.T) {
		config := namedAgentConfig()
		config["audit"] = map[string]any{"agent": "reviewer"}
		config["agent"].(map[string]any)["commands"].(map[string]any)["reviewer"] = []any{"", "{prompt}"}
		requireAgentCommandError(t, config, "audit", "", "C01", "audit.agent", "reviewer")
	})
}
