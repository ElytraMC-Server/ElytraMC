package main

import (
	"fmt"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"elytra.com/backend/config"
	"elytra.com/backend/docs"
	"elytra.com/backend/features/user/getUser"
)

func setupRoutes(e *echo.Echo, conf config.Config) {
	if conf.Mode == config.Dev {
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

func bootstrap(e *echo.Echo, conf config.Config) {
	setupLogger(e)
	setupRoutes(e, conf)
}

func main() {
	e := echo.New()

	conf, err := config.GetConfig()

	if err != nil {
		log.Fatal().Msg(err.Error())
	}

	bootstrap(e, conf)

	if conf.Mode == config.Prod {
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
