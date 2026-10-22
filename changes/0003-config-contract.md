---
module: contracts
set:
  interface: [api.go, 'model/*.go']
---

## Motivation

All product modules need the same typed configuration and defaults, with omissions distinguishable from explicit empty zone lists and version compatibility enforced before executing project commands.

## Interface

Add retained model/*.go data declarations, aliased through api.go: Zones (same existing fields), Hooks {Eval, Boundaries, Setup string}, Agent {Command, Wrap []string}, Audit {Enabled bool; Command []string}, Config matching the design's fireweed.yaml fields with YAML/JSON snake_case tags. Expose ParseConfig(text, binaryVersion string) (Config, error) and CheckVersion(binaryVersion, constraint string) error. Config duration fields remain strings; consumers may parse validated values with Go time.ParseDuration. No runtime or agent execution is added here.

## Added

### new:config-defaults
Configuration parsing MUST supply format version 1, binary range >=0.4 and <0.5, remote origin, module glob packages/*, interface src/index.ts, evals evals/**, disposable src/**, audit enabled, attempts 2, timeout 90m, verify_dependents all, limit patterns usage limit and rate limit, limit_wait 30m and limit_budget 4h when omitted. Explicit empty lists MUST remain empty, including each module-default zone list, protected paths, worker-writable paths and limit patterns. Omitted base MUST remain unresolved for local Git discovery.

### new:config-fields
Configuration parsing MUST retain every design field, including shell hooks and agent argv/wrap values, per-project zone defaults and audit settings, without executing commands or expanding placeholders. It MUST reject malformed YAML, duplicate keys, multiple YAML documents, non-mapping roots and fields of the wrong type. Returned configurations MUST be independent values, so mutating one parse result cannot affect later parses.

### new:config-version
Configuration parsing MUST accept the current and previous file formats (1 and 0) and reject older or future formats. It MUST reject a binary outside the configured comma- or whitespace-separated conjunction of numeric version comparisons, supporting >, >=, <, <=, = and exact major.minor[.patch] versions. It MUST reject malformed binary versions, constraints and empty explicit constraints. A leading v on binary versions MUST be accepted.

### new:config-validation
Configuration parsing MUST reject attempts below one, nonpositive timeout, negative limit wait or budget, invalid Go duration syntax, and verify_dependents values other than all, direct or none. Explicit audit.enabled false MUST disable audit; omission MUST leave it enabled.

## Modified

## Removed

## Decisions

Only nonnegative numeric release versions are supported initially; prerelease/build constraints are rejected rather than compared incorrectly. Omitted version pin defaults to the current 0.4 range, but an explicitly empty pin is invalid. Go duration zero is valid for limit_wait and limit_budget, enabling deterministic bounded protocol tests. Unrecognized configuration fields are rejected as typing errors to surface typos. Empty mapping input is valid; an empty YAML stream is rejected. Strings containing shell metacharacters are opaque data. ParseConfig performs no file, Git, shell or network operations. Module-level metadata inheritance/overrides and file-format migration are separate later changes.

## Rejected alternatives

Using zero values for omission would lose audit's default enabled state and confuse empty zone lists with inheritance. A mutable global default configuration would leak state across independent projects.
