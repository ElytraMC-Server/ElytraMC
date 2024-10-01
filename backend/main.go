package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"

	"elytra.com/backend/docs"
	"elytra.com/backend/features/user/getUser"
)

func setupRoutes(e *echo.Echo, conf Config) {
	if conf.Mode == "dev" {
		docs.RegisterDocs(e, "./docs/spec.yaml")
		e.File("/elytra.png", "./elytra.png")
	}
	getUser.RegisterGetUser(e)
}

func setupLogger(e *echo.Echo) {
	logger := zerolog.New(os.Stdout)
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogMethod: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			logger.Info().
				Str("method", v.Method).
				Str("URI", v.URI).
				Int("status", v.Status).
				Msg("request")

			return nil
		},
	}))
}

func bootstrap(e *echo.Echo, conf Config) {
	setupLogger(e)
	setupRoutes(e, conf)
}

type Config struct {
	AppPort string `koanf:"appPort"`
	Mode    string `koanf:"appMode"`
}

func getConfig() (Config, error) {
	koanf := koanf.New(".")
	var conf Config

	if err := koanf.Load(file.Provider("config.yaml"), yaml.Parser()); err != nil {
		return conf, errors.New(fmt.Sprintf("Couldn't load the config: %v", err))
	}

	if err := koanf.Unmarshal("", &conf); err != nil {
		return conf, errors.New(fmt.Sprintf("Couldn't marshal the config: %v", err))
	}

	return conf, nil
}

func main() {
	e := echo.New()

	conf, err := getConfig()

	if err != nil {
		log.Fatal().Msg(err.Error())
	}

	bootstrap(e, conf)

	if conf.Mode == "dev" {
		e.Use(redirectToDocs)
	}

	if err := e.Start(fmt.Sprintf(":%v", conf.AppPort)); err != nil {
		log.Fatal().Msg(err.Error())
	}
}

func redirectToDocs(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if c.Request().URL.Path == "/" {
			return c.Redirect(302, "/docs")
		}
		return next(c)
	}
}
