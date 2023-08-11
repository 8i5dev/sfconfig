package sfconfig

import (
	"os"
	"testing"
)

func TestYAML(t *testing.T) {
	m := &YAMLLoader{ConfigFile: testYAML}

	s := &Server{}
	if err := m.Load(s); err != nil {
		t.Error(err)
	}

	testStruct(t, s, getDefaultServer(), "file")
}

func TestYAML_Reader(t *testing.T) {
	f, err := os.Open(testYAML)
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		_ = f.Close()
	}()

	m := &YAMLLoader{Reader: f}
	s := &Server{}
	if err := m.Load(s); err != nil {
		t.Error(err)
	}

	testStruct(t, s, getDefaultServer(), "file")
}

func TestJSON(t *testing.T) {
	m := &JSONLoader{ConfigFile: testJSON}

	s := &Server{}
	if err := m.Load(s); err != nil {
		t.Error(err)
	}

	testStruct(t, s, getDefaultServer(), "file")
}

func TestJSON_Reader(t *testing.T) {
	f, err := os.Open(testJSON)
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		_ = f.Close()
	}()

	m := &JSONLoader{Reader: f}
	s := &Server{}
	if err := m.Load(s); err != nil {
		t.Error(err)
	}

	testStruct(t, s, getDefaultServer(), "file")
}
