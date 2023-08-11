package sfconfig

import (
	"testing"
)

func TestValidators(t *testing.T) {
	v := &RequiredValidator{}
	s := getDefaultServer()
	s.Name = ""

	err := v.Validate(s)
	if err == nil {
		t.Fatal("Name should be required")
	}
}

func TestValidatorsEmbeddedStruct(t *testing.T) {
	v := &RequiredValidator{}
	s := getDefaultServer()
	s.Postgres.Port = 0

	err := v.Validate(s)
	if err == nil {
		t.Fatal("Port should be required")
	}
}
