package config

import (
	"errors"
	"fmt"

	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
)

type appMode string

const (
	Dev  appMode = "dev"
	Prod appMode = "prod"
)

type Config struct {
	AppPort string  `koanf:"appPort"`
	Mode    appMode `koanf:"appMode"`
	DbConn  string  `koanf:"dbConn"`
}

type ConfigLoadError struct {
	errors []error
}

func (loadErr *ConfigLoadError) Error() string {
	return fmt.Sprintf("Got error(s) while loading configuration: %v", loadErr.errors)
}

func GetConfig() (Config, error) {
	koanf := koanf.New(".")
	var conf Config

	var loadErrs []error

	if err := koanf.Load(file.Provider("config/secrets.yaml"), yaml.Parser()); err != nil {
		loadErrs = append(loadErrs, errors.New(fmt.Sprintf("Couldn't load the secrets: %v", err)))
	}

	if err := koanf.Load(file.Provider("config/launchConfig.yaml"), yaml.Parser()); err != nil {
		loadErrs = append(loadErrs, errors.New(fmt.Sprintf("Couldn't load the launch configuration: %v", err)))
	}

	if len(loadErrs) > 0 {
		return Config{}, &ConfigLoadError{errors: loadErrs}
	}

	if err := koanf.Unmarshal("", &conf); err != nil {
		return Config{}, errors.New(fmt.Sprintf("Couldn't marshal the config: %v", err))
	}

	return conf, nil
}
