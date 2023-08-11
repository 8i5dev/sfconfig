package sfconfig

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"

	"gopkg.in/yaml.v3"
)

type JSONLoader struct {
	ConfigFile string
	Reader     io.Reader
}

// Load loads the source into the config defined by struct s.
// Defaults to using the Reader if provided, otherwise tries to read from the
// file
func (j *JSONLoader) Load(s interface{}) error {
	var r io.Reader

	if j.Reader != nil {
		r = j.Reader
		//if data, err := io.ReadAll(j.Reader);err != nil {
		//	return fmt.Errorf("reader read error: %v", err)
		//}
	} else if j.ConfigFile != "" {
		jsonFile, err := os.Open(j.ConfigFile)
		if err != nil {
			return fmt.Errorf("open file error: %v", err)
		}
		defer func() {
			_ = jsonFile.Close()
		}()

		r = jsonFile
	} else {
		return errors.New("loader is not set")
	}

	data, err := io.ReadAll(r)
	if err != nil {
		return fmt.Errorf("read file error: %v", err)
	}

	return json.Unmarshal(data, s)
}

// YAMLLoader satisfies the loader interface. It loads the configuration from
// the given yaml file.
type YAMLLoader struct {
	ConfigFile string
	Reader     io.Reader
}

// Load loads the source into the config defined by struct s.
// Defaults to using the Reader if provided, otherwise tries to read from the
// file
func (y *YAMLLoader) Load(s interface{}) error {
	var r io.Reader

	if y.Reader != nil {
		r = y.Reader
	} else if y.ConfigFile != "" {
		yamlFile, err := os.Open(y.ConfigFile)
		if err != nil {
			return err
		}
		defer func() {
			_ = yamlFile.Close()
		}()

		r = yamlFile
	} else {
		return errors.New("loader is not set")
	}

	data, err := io.ReadAll(r)
	if err != nil {
		return fmt.Errorf("read file error: %v", err)
	}

	return yaml.Unmarshal(data, s)
}
