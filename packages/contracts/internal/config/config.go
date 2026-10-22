package config

import (
	"fmt"
	"io"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode"

	"github.com/lurkshark/fireweed/packages/contracts/model"
	"gopkg.in/yaml.v3"
)

func defaults() model.Config {
	return model.Config{
		Fireweed: ">=0.4, <0.5", Version: 1, Remote: "origin", Modules: []string{"packages/*"},
		ModuleDefaults: model.Zones{Interface: []string{"src/index.ts"}, Evals: []string{"evals/**"}, Disposable: []string{"src/**"}},
		Audit:          model.Audit{Enabled: true}, Attempts: 2, Timeout: "90m", VerifyDependents: "all",
		LimitPatterns: []string{"usage limit", "rate limit"}, LimitWait: "30m", LimitBudget: "4h",
	}
}

func Parse(text, binaryVersion string) (model.Config, error) {
	config := defaults()
	decoder := yaml.NewDecoder(strings.NewReader(text))
	var document yaml.Node
	if err := decoder.Decode(&document); err != nil {
		return model.Config{}, fmt.Errorf("configuration YAML: %w", err)
	}
	if len(document.Content) != 1 || document.Content[0].Kind != yaml.MappingNode {
		return model.Config{}, fmt.Errorf("configuration must be a mapping")
	}
	var extra yaml.Node
	if err := decoder.Decode(&extra); err != io.EOF {
		if err != nil {
			return model.Config{}, err
		}
		return model.Config{}, fmt.Errorf("configuration contains multiple YAML documents")
	}
	if err := schema(document.Content[0], reflect.TypeOf(config), "config"); err != nil {
		return model.Config{}, err
	}
	if err := document.Content[0].Decode(&config); err != nil {
		return model.Config{}, err
	}
	if config.Version < 0 || config.Version > 1 {
		return model.Config{}, fmt.Errorf("unsupported format version %d", config.Version)
	}
	if err := Version(binaryVersion, config.Fireweed); err != nil {
		return model.Config{}, err
	}
	if config.Attempts < 1 {
		return model.Config{}, fmt.Errorf("attempts must be positive")
	}
	for _, field := range []struct {
		name, value string
		positive    bool
	}{{"timeout", config.Timeout, true}, {"limit_wait", config.LimitWait, false}, {"limit_budget", config.LimitBudget, false}} {
		duration, err := time.ParseDuration(field.value)
		if err != nil {
			return model.Config{}, fmt.Errorf("%s: %w", field.name, err)
		}
		if duration < 0 || (field.positive && duration == 0) {
			return model.Config{}, fmt.Errorf("invalid %s duration %s", field.name, field.value)
		}
	}
	if config.VerifyDependents != "all" && config.VerifyDependents != "direct" && config.VerifyDependents != "none" {
		return model.Config{}, fmt.Errorf("verify_dependents must be all, direct or none")
	}
	return config, nil
}

// Validate YAML's actual types before decoding: the YAML library otherwise
// coerces some numeric scalars into strings, hiding malformed argv or hooks.
func schema(node *yaml.Node, typ reflect.Type, path string) error {
	if node.Kind == yaml.AliasNode {
		return schema(node.Alias, typ, path)
	}
	switch typ.Kind() {
	case reflect.Struct:
		if node.Kind != yaml.MappingNode {
			return fmt.Errorf("%s must be a mapping", path)
		}
		fields := fieldsOf(typ)
		seen := map[string]bool{}
		for i := 0; i < len(node.Content); i += 2 {
			key := node.Content[i]
			if key.Tag != "!!str" {
				return fmt.Errorf("%s keys must be strings", path)
			}
			field, ok := fields[key.Value]
			if !ok {
				return fmt.Errorf("unknown %s field %s", path, key.Value)
			}
			if seen[key.Value] {
				return fmt.Errorf("duplicate %s field %s", path, key.Value)
			}
			seen[key.Value] = true
			if err := schema(node.Content[i+1], field, path+"."+key.Value); err != nil {
				return err
			}
		}
	case reflect.Map:
		if node.Kind != yaml.MappingNode || typ.Key().Kind() != reflect.String {
			return fmt.Errorf("%s must be a string-keyed mapping", path)
		}
		seen := map[string]bool{}
		for i := 0; i < len(node.Content); i += 2 {
			key := node.Content[i]
			if key.Tag != "!!str" {
				return fmt.Errorf("%s keys must be strings", path)
			}
			if seen[key.Value] {
				return fmt.Errorf("duplicate key %s", key.Value)
			}
			seen[key.Value] = true
			if err := schema(node.Content[i+1], typ.Elem(), path+"."+key.Value); err != nil {
				return err
			}
		}
	case reflect.Interface:
		// Dynamic set values are checked against their allowed typed fields later.
		return nil
	case reflect.Slice:
		if node.Kind != yaml.SequenceNode {
			return fmt.Errorf("%s must be a list", path)
		}
		for i, item := range node.Content {
			if err := schema(item, typ.Elem(), fmt.Sprintf("%s[%d]", path, i)); err != nil {
				return err
			}
		}
	case reflect.String:
		if node.Kind != yaml.ScalarNode || node.Tag != "!!str" {
			return fmt.Errorf("%s must be a string", path)
		}
	case reflect.Bool:
		if node.Kind != yaml.ScalarNode || node.Tag != "!!bool" {
			return fmt.Errorf("%s must be a boolean", path)
		}
	case reflect.Int:
		if node.Kind != yaml.ScalarNode || node.Tag != "!!int" {
			return fmt.Errorf("%s must be an integer", path)
		}
	default:
		return fmt.Errorf("unsupported schema type %s", typ)
	}
	return nil
}

var numericVersion = regexp.MustCompile(`^v?([0-9]+)\.([0-9]+)(?:\.([0-9]+))?$`)
var comparison = regexp.MustCompile(`^([<>]=?|=)?[ \t]*([0-9]+\.[0-9]+(?:\.[0-9]+)?)`)

func versionNumbers(text string) ([3]int, error) {
	result := [3]int{}
	parts := numericVersion.FindStringSubmatch(text)
	if parts == nil {
		return result, fmt.Errorf("malformed numeric version %q", text)
	}
	for i, value := range parts[1:] {
		if value != "" {
			number, err := strconv.Atoi(value)
			if err != nil {
				return result, err
			}
			result[i] = number
		}
	}
	return result, nil
}
func Version(binaryVersion, constraint string) error {
	actual, err := versionNumbers(binaryVersion)
	if err != nil {
		return err
	}
	remaining := strings.TrimSpace(constraint)
	if remaining == "" {
		return fmt.Errorf("empty binary version constraint")
	}
	satisfied := true
	for remaining != "" {
		clause := comparison.FindStringSubmatch(remaining)
		if clause == nil {
			return fmt.Errorf("malformed binary version constraint %q", constraint)
		}
		required, err := versionNumbers(clause[2])
		if err != nil {
			return err
		}
		cmp := 0
		for i := 0; i < 3; i++ {
			if actual[i] < required[i] {
				cmp = -1
				break
			}
			if actual[i] > required[i] {
				cmp = 1
				break
			}
		}
		switch clause[1] {
		case "", "=":
			satisfied = satisfied && cmp == 0
		case ">":
			satisfied = satisfied && cmp > 0
		case ">=":
			satisfied = satisfied && cmp >= 0
		case "<":
			satisfied = satisfied && cmp < 0
		case "<=":
			satisfied = satisfied && cmp <= 0
		}
		tail := remaining[len(clause[0]):]
		if tail == "" {
			break
		}
		trimmed := strings.TrimLeftFunc(tail, unicode.IsSpace)
		if strings.HasPrefix(trimmed, ",") {
			remaining = strings.TrimSpace(trimmed[1:])
			if remaining == "" {
				return fmt.Errorf("empty constraint clause")
			}
		} else {
			if len(tail) == len(trimmed) {
				return fmt.Errorf("missing constraint separator")
			}
			remaining = trimmed
		}
	}
	if !satisfied {
		return fmt.Errorf("binary %s is outside pinned range %q", binaryVersion, constraint)
	}
	return nil
}
