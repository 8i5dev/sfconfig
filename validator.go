package sfconfig

import (
	"fmt"
	"reflect"
)

type RequiredValidator struct {
	//  TagName holds the validator tag name. The default is "required"
	TagName string

	// TagValue holds the expected value of the validator. The default is "true"
	TagValue string
}

// Validate validates the given struct against field's zero values. If
// intentionally, the value of a field is `zero-valued`(e.g false, 0, "")
// required tag should not be set for that field.
func (e *RequiredValidator) Validate(s interface{}) error {
	if e.TagName == "" {
		e.TagName = "required"
	}

	if e.TagValue == "" {
		e.TagValue = "true"
	}

	for _, field := range StructFields(s) {
		if err := e.processField("", field); err != nil {
			return err
		}
	}

	return nil
}

func (e *RequiredValidator) processField(fieldName string, field *Field) error {
	fieldName += field.Name()
	switch field.Kind() {
	case reflect.Struct:
		// this is used for error messages below, when we have an error at the
		// child properties add parent properties into the error message as well
		fieldName += "."

		for _, f := range field.Fields() {
			if err := e.processField(fieldName, f); err != nil {
				return err
			}
		}
	default:
		val := field.Tag(e.TagName)
		if val != e.TagValue {
			return nil
		}

		if field.IsZero() {
			return fmt.Errorf("sfconfig: field '%s' is required", fieldName)
		}
	}

	return nil
}
