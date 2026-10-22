package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io"
	"reflect"
	"strings"
)

// Decode shares the strict type validation used by configuration with document metadata.
func Decode(text string, target any) error {
	typ := reflect.TypeOf(target)
	if typ == nil || typ.Kind() != reflect.Pointer {
		return fmt.Errorf("metadata target must be a pointer")
	}
	var document yaml.Node
	decoder := yaml.NewDecoder(strings.NewReader(text))
	if err := decoder.Decode(&document); err != nil {
		return err
	}
	if len(document.Content) != 1 {
		return fmt.Errorf("missing metadata")
	}
	var extra yaml.Node
	if err := decoder.Decode(&extra); err != io.EOF {
		return fmt.Errorf("metadata must contain one YAML document")
	}
	if err := schema(document.Content[0], typ.Elem(), "metadata"); err != nil {
		return err
	}
	return document.Content[0].Decode(target)
}
func fieldsOf(typ reflect.Type) map[string]reflect.Type {
	fields := map[string]reflect.Type{}
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		tags := strings.Split(field.Tag.Get("yaml"), ",")
		if tags[0] == "-" {
			continue
		}
		inline := false
		for _, option := range tags[1:] {
			if option == "inline" {
				inline = true
			}
		}
		if inline {
			for name, child := range fieldsOf(field.Type) {
				fields[name] = child
			}
			continue
		}
		name := tags[0]
		if name == "" {
			name = strings.ToLower(field.Name)
		}
		fields[name] = field.Type
	}
	return fields
}
