package sfconfig

import (
	"os"
	"reflect"
	"strings"
)

type EnvironmentLoader struct {
	// Prefix prepends given string to every environment variable
	// FIELDNAME will be {PREFIX}_FIELDNAME
	Prefix string

	// CamelCase adds a separator for field names in camelcase form. A
	// fieldname of "AccessKey" would generate an environment name of
	// "{PREFIX}_ACCESSKEY". If CamelCase is enabled, the environment name
	// will be generated in the form of "{PREFIX}_ACCESS_KEY"
	CamelCase bool

	EnvTagName string
}

// Load loads the source into the config defined by struct s
func (e *EnvironmentLoader) Load(s interface{}) error {
	if e.EnvTagName == "" {
		e.EnvTagName = "env"
	}

	for _, field := range StructFields(s) {
		if err := e.processField(e.Prefix, field); err != nil {
			return err
		}
	}

	return nil
}

// processField gets leading name for the env variable and combines the current
// field's name and generates environment variable names recursively
func (e *EnvironmentLoader) processField(prefix string, field *Field) error {
	switch field.Kind() {
	case reflect.Struct:
		fieldName := field.Name()
		prefix = e.generateFieldName(prefix, fieldName)
		for _, f := range field.Fields() {
			if err := e.processField(prefix, f); err != nil {
				return err
			}
		}
	default:
		envVal, ok := field.TagLookup(e.EnvTagName)
		if !ok {
			return nil
		}
		if envVal == "" {
			envVal = e.defaultEnvName(prefix, field.Name())
		}

		v := os.Getenv(envVal)
		if v == "" {
			return nil
		}

		err := FieldSet(field, v)
		if err != nil {
			return err
		}
	}

	return nil
}

func (e *EnvironmentLoader) defaultEnvName(prefix string, fieldName string) string {
	name := fieldName

	if e.CamelCase {
		name = ToSnakeCase(fieldName)
	}

	if prefix != "" {
		name = prefix + "_" + name
	}

	return strings.ToUpper(name)
}

// generateFieldName generates the field name combined with the prefix and the
// field name of struct
func (e *EnvironmentLoader) generateFieldName(prefix string, name string) string {
	fieldName := strings.ToUpper(name)
	if e.CamelCase {
		fieldName = ToSnakeCase(name)
	}

	if prefix == "" {
		return fieldName
	} else {
		return strings.ToUpper(prefix) + "_" + fieldName
	}
}
