package sfconfig

import (
	"fmt"
	"os"
	"strings"
)

type Loader interface {
	// Load loads the source into the config defined by struct s
	Load(s interface{}) error
}

type Validator interface {
	// Validate validates the config struct
	Validate(s interface{}) error
}

type SFConfig struct {
	loaders    []Loader
	validators []Validator
}

func (c *SFConfig) AddLoader(l Loader) {
	c.loaders = append(c.loaders, l)
}

func (c *SFConfig) AddValidator(v Validator) {
	c.validators = append(c.validators, v)
}

func (c *SFConfig) MustLoad(conf interface{}) {
	for i, loader := range c.loaders {
		if err := loader.Load(conf); err != nil {
			fmt.Printf("error loading config from loader %d: %v\n", i, err)
			os.Exit(2)
		}
	}
}

func (c *SFConfig) MustValidate(conf interface{}) {
	for i, validator := range c.validators {
		if err := validator.Validate(conf); err != nil {
			fmt.Printf("error validating config with validator %d: %v\n", i, err)
			os.Exit(2)
		}
	}
}

func (c *SFConfig) Load(conf interface{}) {
	c.MustLoad(conf)
	c.MustValidate(conf)
}

func New() *SFConfig {
	d := &SFConfig{}
	d.loaders = make([]Loader, 0)
	d.validators = make([]Validator, 0)
	return d
}

func NewDefaultConfig(configFile string, envPrefix string) *SFConfig {
	d := New()
	d.AddLoader(&DefaultTagLoader{})

	if strings.HasSuffix(configFile, "json") {
		d.AddLoader(&JSONLoader{ConfigFile: configFile})
	}
	if strings.HasSuffix(configFile, "yml") || strings.HasSuffix(configFile, "yaml") {
		d.AddLoader(&YAMLLoader{ConfigFile: configFile})
	}

	d.AddLoader(&EnvironmentLoader{Prefix: envPrefix, CamelCase: true})

	//d.AddLoader(&FlagLoader{})
	d.AddValidator(&RequiredValidator{})
	return d
}
