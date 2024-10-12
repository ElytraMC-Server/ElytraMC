package main

import (
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"

	"elytra.com/backend/config"
	"elytra.com/backend/docs"
)

func registerOpenAPI(e *echo.Echo) {
	docs.RegisterDocs(e, "./docs/spec.yaml")
	e.File("/elytra.png", "./elytra.png")
}

func main() {
	conf, err := config.GetConfig()

	if err != nil {
		log.Fatal().Msg(err.Error())
	}

	app, err := NewApp(conf, echo.New())

	if err != nil {
		log.Fatal().Msg(err.Error())
	}

	app.setupLogger()
	app.setupRoutes()
	app.configEcho()

	if conf.Mode == config.Dev {
		app.Echo.Use(redirectToDocs)
	}

	app.Run()
}

func redirectToDocs(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if c.Request().URL.Path == "/" {
			return c.Redirect(302, "/docs")
		}
		return next(c)
	}
}
