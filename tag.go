package sfconfig

import (
	"reflect"
)

type DefaultTagLoader struct {
	// DefaultTagName is the default tag name for struct fields to define
	// default values for a field.
	DefaultTagName string
}

func (t *DefaultTagLoader) Load(s interface{}) error {
	if t.DefaultTagName == "" {
		t.DefaultTagName = "default"
	}

	for _, field := range StructFields(s) {
		if err := t.processField(field); err != nil {
			return err
		}
	}

	return nil
}

func (t *DefaultTagLoader) processField(field *Field) error {
	switch field.Kind() {
	case reflect.Struct:
		for _, f := range field.Fields() {
			if err := t.processField(f); err != nil {
				return err
			}
		}
	default:
		defaultVal := field.Tag(t.DefaultTagName)
		if defaultVal == "" {
			return nil
		}

		err := FieldSet(field, defaultVal)
		if err != nil {
			return err
		}
	}

	return nil
}
