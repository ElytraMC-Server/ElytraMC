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
}

func GetConfig() (Config, error) {
	koanf := koanf.New(".")
	var conf Config

	if err := koanf.Load(file.Provider("config/config.yaml"), yaml.Parser()); err != nil {
		return conf, errors.New(fmt.Sprintf("Couldn't load the config: %v", err))
	}

	if err := koanf.Unmarshal("", &conf); err != nil {
		return conf, errors.New(fmt.Sprintf("Couldn't marshal the config: %v", err))
	}

	return conf, nil
}
