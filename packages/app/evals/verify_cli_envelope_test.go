package evals

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
	"testing"
)

// Inspect keys before decoding values: encoding/json's struct decoder accepts
// case-insensitive field names and silently ignores unknown or duplicate keys.
func cliVerifyEnvelopeKeys(raw string) error {
	d := json.NewDecoder(strings.NewReader(raw))
	token, err := d.Token()
	if err != nil || token != json.Delim('{') {
		return fmt.Errorf("envelope must be an object: %v", err)
	}
	seen := map[string]bool{}
	for d.More() {
		token, err := d.Token()
		if err != nil {
			return err
		}
		key, ok := token.(string)
		if !ok || seen[key] {
			return fmt.Errorf("invalid or duplicate envelope key %v", token)
		}
		switch key {
		case "schema_version", "command", "data", "diagnostics":
		default:
			return fmt.Errorf("noncanonical envelope key %q", key)
		}
		seen[key] = true
		var value json.RawMessage
		if err := d.Decode(&value); err != nil {
			return err
		}
	}
	if token, err = d.Token(); err != nil || token != json.Delim('}') {
		return fmt.Errorf("incomplete envelope: %v", err)
	}
	if len(seen) != 4 {
		return fmt.Errorf("incomplete envelope keys: %v", seen)
	}
	var extra any
	if err := d.Decode(&extra); err != io.EOF {
		return fmt.Errorf("extra envelope output: %v", err)
	}
	return nil
}

func cliVerifyDecode(t *testing.T, got result, command string) envelope {
	t.Helper()
	if err := cliVerifyEnvelopeKeys(got.out); err != nil {
		t.Fatalf("noncanonical JSON envelope: %v: %q", err, got.out)
	}
	return decode(t, got, command)
}

func Test_APP026VerifyEnvelopeControls(t *testing.T) {
	// Data may be any command wire value, including null for fatal/usage
	// outcomes; whitespace and key order have no protocol significance.
	for _, data := range []string{`null`, `"help text"`, `{"change":"current","diagnostics":[],"coverage":[],"edits":[],"phase":"audit"}`} {
		valid := `{"schema_version":1,"command":"verify","data":` + data + `,"diagnostics":[]}`
		for _, raw := range []string{valid, " \n" + valid + "\n\t", `{"diagnostics":[],"data":` + data + `,"command":"verify","schema_version":1}`} {
			if err := cliVerifyEnvelopeKeys(raw); err != nil {
				t.Fatal("canonical envelope rejected", err, raw)
			}
		}
		for _, key := range []string{"schema_version", "command", "data", "diagnostics"} {
			for _, spelling := range []string{strings.ToUpper(key), "extra_" + key} {
				// The nested data may also contain diagnostics. Mutate the
				// envelope member, which is the last occurrence in this fixture.
				needle := `"` + key + `":`
				i := strings.LastIndex(valid, needle)
				raw := valid[:i] + `"` + spelling + `":` + valid[i+len(needle):]
				if cliVerifyEnvelopeKeys(raw) == nil {
					t.Fatal("noncanonical envelope key escaped observer", raw)
				}
			}
			for _, value := range []string{`null`, `[]`} {
				raw := strings.TrimSuffix(valid, "}") + `,"` + key + `":` + value + `}`
				if cliVerifyEnvelopeKeys(raw) == nil {
					t.Fatal("duplicate envelope key escaped observer", raw)
				}
			}
		}
		for _, raw := range []string{strings.TrimSuffix(valid, "}") + `,"extra":true}`, valid + valid, valid + " prompt> ", `[]`, `null`, `{}`, `{"schema_version":1,"command":"verify","diagnostics":[]}`} {
			if cliVerifyEnvelopeKeys(raw) == nil {
				t.Fatal("extra/missing envelope content escaped observer", raw)
			}
		}
	}
}
