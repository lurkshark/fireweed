package fixture

import (
	"cmp"
	"fmt"
	"strings"

	"github.com/lurkshark/fireweed/packages/assets"
	"github.com/lurkshark/fireweed/packages/project/internal/source"
)

func CheckCompatibility(config map[string]any) error {
	value, present := config["fireweed"]
	if !present {
		return nil
	}
	version := assets.Version()
	pin, ok := value.(string)
	if !ok {
		return source.FileError("C01", configPath, fmt.Errorf("invalid fireweed pin %v for binary version %s: expected a string", value, version))
	}
	running, ok := parseVersion(version)
	if !ok {
		return source.FileError("C01", configPath, fmt.Errorf("cannot check fireweed pin %q against invalid binary version %q", pin, version))
	}

	compatible := true
	for _, clause := range strings.Split(pin, ",") {
		clause = strings.TrimSpace(clause)
		endpoint := strings.TrimLeft(clause, "<=>")
		operator := clause[:len(clause)-len(endpoint)]
		bound, valid := parseVersion(strings.TrimSpace(endpoint))
		matches, validOperator := compareVersion(operator, running.compare(bound))
		if !valid || !validOperator {
			return source.FileError("C01", configPath, fmt.Errorf("invalid syntax in fireweed pin %q for binary version %s: invalid comparison %q", pin, version, clause))
		}
		compatible = compatible && matches
	}
	if !compatible {
		return source.FileError("C01", configPath, fmt.Errorf("fireweed pin %q excludes running binary version %s", pin, version))
	}
	return nil
}

func compareVersion(operator string, order int) (matches, valid bool) {
	switch operator {
	case "=":
		return order == 0, true
	case "<":
		return order < 0, true
	case "<=":
		return order <= 0, true
	case ">":
		return order > 0, true
	case ">=":
		return order >= 0, true
	default:
		return false, false
	}
}

// Each component holds decimal digits without leading zeros. An empty component
// represents zero, including an omitted patch.
type decimalVersion [3]string

func parseVersion(version string) (decimalVersion, bool) {
	var parsed decimalVersion
	parts := strings.Split(version, ".")
	if len(parts) < 2 || len(parts) > len(parsed) {
		return parsed, false
	}
	for i, part := range parts {
		if part == "" || strings.Trim(part, "0123456789") != "" {
			return parsed, false
		}
		parsed[i] = strings.TrimLeft(part, "0")
	}
	return parsed, true
}

func (version decimalVersion) compare(other decimalVersion) int {
	for i, part := range version {
		// Equal-length decimal components sort lexically in numeric order.
		order := cmp.Or(cmp.Compare(len(part), len(other[i])), strings.Compare(part, other[i]))
		if order != 0 {
			return order
		}
	}
	return 0
}
