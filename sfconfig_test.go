package sfconfig

import (
	"encoding/json"
	"errors"
	"testing"
	"time"
)

type (
	Duration time.Duration

	Server struct {
		Name       string `required:"true"`
		Port       int    `default:"6060"`
		ID         int64
		Labels     []int
		Enabled    bool
		Users      []string
		Postgres   Postgres
		unexported string
		Interval   Duration
	}

	// Postgres holds Postgresql database related configuration
	Postgres struct {
		Enabled           bool
		Port              int      `required:"true" customRequired:"yes"`
		Hosts             []string `required:"true"`
		DBName            string   `default:"configdb"`
		AvailabilityRatio float64
		unexported        string
	}
)

func (d *Duration) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Duration(*d).String())
}

func (d *Duration) UnmarshalJSON(b []byte) error {
	var v interface{}
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	switch value := v.(type) {
	case float64:
		*d = Duration(time.Duration(value))
		return nil
	case string:
		var err error
		tmp, err := time.ParseDuration(value)
		if err != nil {
			return err
		}
		*d = Duration(tmp)
		return nil
	default:
		return errors.New("invalid duration")
	}
}

var (
	testTOML = "testdata/config.toml"
	testJSON = "testdata/config.json"
	testYAML = "testdata/config.yaml"
)

func getDefaultServer() *Server {
	return &Server{
		Name:     "koding",
		Port:     6060,
		Enabled:  true,
		ID:       1234567890,
		Labels:   []int{123, 456},
		Users:    []string{"ankara", "istanbul"},
		Interval: Duration(10 * time.Second),
		Postgres: Postgres{
			Enabled:           true,
			Port:              5432,
			Hosts:             []string{"192.168.2.1", "192.168.2.2", "192.168.2.3"},
			DBName:            "configdb",
			AvailabilityRatio: 8.23,
		},
	}
}

func testStruct(t *testing.T, s *Server, d *Server, kind string) {
	if kind == "file" {
		if s.Name != d.Name {
			t.Errorf("Name value is wrong: %s, want: %s", s.Name, d.Name)
		}

		if s.Enabled != d.Enabled {
			t.Errorf("Enabled value is wrong: %t, want: %t", s.Enabled, d.Enabled)
		}

		if s.Interval != d.Interval {
			t.Errorf("Interval value is wrong: %v, want: %v", s.Interval, d.Interval)
		}

		if s.ID != d.ID {
			t.Errorf("ID value is wrong: %v, want: %v", s.ID, d.ID)
		}

		if len(s.Labels) != len(d.Labels) {
			t.Errorf("Labels value is wrong: %d, want: %d", len(s.Labels), len(d.Labels))
		} else {
			for i, label := range d.Labels {
				if s.Labels[i] != label {
					t.Errorf("Label is wrong for index: %d, label: %d, want: %d", i, s.Labels[i], label)
				}
			}
		}

		if len(s.Users) != len(d.Users) {
			t.Errorf("Users value is wrong: %d, want: %d", len(s.Users), len(d.Users))
		} else {
			for i, user := range d.Users {
				if s.Users[i] != user {
					t.Errorf("User is wrong for index: %d, user: %s, want: %s", i, s.Users[i], user)
				}
			}
		}
	} else if kind == "default" {
		if s.Port != d.Port {
			t.Errorf("Port value is wrong: %d, want: %d", s.Port, d.Port)
		}
	}

	testPostgres(t, s.Postgres, d.Postgres, kind)
}

func testPostgres(t *testing.T, s Postgres, d Postgres, kind string) {
	if kind == "file" {
		if s.Enabled != d.Enabled {
			t.Errorf("Postgres enabled is wrong %t, want: %t", s.Enabled, d.Enabled)
		}

		if s.Port != d.Port {
			t.Errorf("Postgres Port value is wrong: %d, want: %d", s.Port, d.Port)
		}

		if s.AvailabilityRatio != d.AvailabilityRatio {
			t.Errorf("AvailabilityRatio is wrong: %f, want: %f", s.AvailabilityRatio, d.AvailabilityRatio)
		}

		if len(s.Hosts) != len(d.Hosts) {
			// do not continue testing if this fails, because others is depending on this test
			t.Fatalf("Hosts len is wrong: %v, want: %v", s.Hosts, d.Hosts)
		}

		for i, host := range d.Hosts {
			if s.Hosts[i] != host {
				t.Fatalf("Hosts number %d is wrong: %v, want: %v", i, s.Hosts[i], host)
			}
		}
	} else if kind == "default" {
		if s.DBName != d.DBName {
			t.Errorf("DBName is wrong: %s, want: %s", s.DBName, d.DBName)
		}
	}
}
