package sfconfig

import (
	"os"
	"strings"
	"testing"
)

type (
	ServerEnv struct {
		ServerPort int    `env:""`
		UserName   string `env:""`
		IsEnabled  bool   `env:""`
		MySql      MySql
	}

	MySql struct {
		Host              string   `env:"MYSQL_HOST"`
		Enabled           bool     `env:""`
		Port              int      `required:"true" customRequired:"yes"`
		Hosts             []string `required:"true"`
		DBName            string   `default:"configdb"`
		AvailabilityRatio float64
		unexported        string
	}
)

func TestENV(t *testing.T) {
	m := EnvironmentLoader{}
	s := &ServerEnv{}

	// set env variables
	setEnvVars(t, m.Prefix, m.CamelCase)

	if err := m.Load(s); err != nil {
		t.Error(err)
	}

	if s.ServerPort != 6060 {
		t.Errorf("ServerPort value is wrong: %d, want: %d", s.ServerPort, 6060)
	}
	if s.UserName != "username" {
		t.Errorf("UserName value is wrong: %s, want: %s", s.UserName, "username")
	}
	if s.IsEnabled != true {
		t.Errorf("IsEnabled value is wrong: %v, want: %v", s.IsEnabled, true)
	}
	if s.MySql.Host != "localhost" {
		t.Errorf("MySql.Host value is wrong: %s, want: %s", s.MySql.Host, "localhost")
	}
	if s.MySql.Enabled != true {
		t.Errorf("MySql.Enabled value is wrong: %v, want: %v", s.MySql.Enabled, true)
	}
}

func TestCamelCaseEnv(t *testing.T) {
	m := EnvironmentLoader{CamelCase: true}
	s := &ServerEnv{}

	// set env variables
	setEnvVars(t, m.Prefix, m.CamelCase)

	if err := m.Load(s); err != nil {
		t.Error(err)
	}

	if s.ServerPort != 6060 {
		t.Errorf("ServerPort value is wrong: %d, want: %d", s.ServerPort, 6060)
	}
	if s.UserName != "username" {
		t.Errorf("UserName value is wrong: %s, want: %s", s.UserName, "username")
	}
	if s.IsEnabled != true {
		t.Errorf("IsEnabled value is wrong: %v, want: %v", s.IsEnabled, true)
	}
	if s.MySql.Host != "localhost" {
		t.Errorf("MySql.Host value is wrong: %s, want: %s", s.MySql.Host, "localhost")
	}
	if s.MySql.Enabled != true {
		t.Errorf("MySql.Enabled value is wrong: %v, want: %v", s.MySql.Enabled, true)
	}
}

func TestENVWithPrefix(t *testing.T) {
	m := EnvironmentLoader{Prefix: "prefix"}
	s := &ServerEnv{}

	// set env variables
	setEnvVars(t, m.Prefix, m.CamelCase)

	if err := m.Load(s); err != nil {
		t.Error(err)
	}

	if s.ServerPort != 6060 {
		t.Errorf("ServerPort value is wrong: %d, want: %d", s.ServerPort, 6060)
	}
	if s.UserName != "username" {
		t.Errorf("UserName value is wrong: %s, want: %s", s.UserName, "username")
	}
	if s.IsEnabled != true {
		t.Errorf("IsEnabled value is wrong: %v, want: %v", s.IsEnabled, true)
	}
	if s.MySql.Host != "localhost" {
		t.Errorf("MySql.Host value is wrong: %s, want: %s", s.MySql.Host, "localhost")
	}
	if s.MySql.Enabled != true {
		t.Errorf("MySql.Enabled value is wrong: %v, want: %v", s.MySql.Enabled, true)
	}
}

func TestCamelCaseEnvWithPrefix(t *testing.T) {
	m := EnvironmentLoader{Prefix: "prefix", CamelCase: true}
	s := &ServerEnv{}

	// set env variables
	setEnvVars(t, m.Prefix, m.CamelCase)

	if err := m.Load(s); err != nil {
		t.Error(err)
	}

	if s.ServerPort != 6060 {
		t.Errorf("ServerPort value is wrong: %d, want: %d", s.ServerPort, 6060)
	}
	if s.UserName != "username" {
		t.Errorf("UserName value is wrong: %s, want: %s", s.UserName, "username")
	}
	if s.IsEnabled != true {
		t.Errorf("IsEnabled value is wrong: %v, want: %v", s.IsEnabled, true)
	}
	if s.MySql.Host != "localhost" {
		t.Errorf("MySql.Host value is wrong: %s, want: %s", s.MySql.Host, "localhost")
	}
	if s.MySql.Enabled != true {
		t.Errorf("MySql.Enabled value is wrong: %v, want: %v", s.MySql.Enabled, true)
	}
}

func setEnvVars(t *testing.T, prefix string, camelCase bool) {
	var omitKey = map[string]string{"MYSQL_HOST": ""}
	var env map[string]string
	if camelCase {
		env = map[string]string{
			"SERVER_PORT":    "6060",
			"USER_NAME":      "username",
			"IS_ENABLED":     "true",
			"MYSQL_HOST":     "localhost",
			"MY_SQL_ENABLED": "true",
		}
	} else {
		env = map[string]string{
			"SERVERPORT":    "6060",
			"USERNAME":      "username",
			"ISENABLED":     "true",
			"MYSQL_HOST":    "localhost",
			"MYSQL_ENABLED": "true",
		}
	}

	prefix = strings.ToUpper(prefix)
	for key, val := range env {
		var env string
		if _, ok := omitKey[key]; ok {
			env = key
		} else {
			if prefix == "" {
				env = key
			} else {
				env = prefix + "_" + key
			}
		}

		if err := os.Setenv(env, val); err != nil {
			t.Fatal(err)
		}
	}
}
