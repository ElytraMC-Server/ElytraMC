package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"elytra.com/backend/docs"
	"elytra.com/backend/features/user/getUser"
)

func setupRoutes(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	docs.RegisterDocs(e, "./docs/openapi.spec.yaml")
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

func bootstrap(e *echo.Echo) {
	setupLogger(e)
	setupRoutes(e)
}

func main() {
	e := echo.New()

	bootstrap(e)

	if err := e.Start(":9000"); err != nil {
		log.Fatal().Msg(err.Error())
	}
}
