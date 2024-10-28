package main

import (
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"

	"elytra.com/backend/app"
	"elytra.com/backend/config"
)

func main() {
	conf, err := config.GetConfig()

	if err != nil {
		log.Fatal().Msg(err.Error())
	}

	app, err := app.NewApp(conf, echo.New())

	if err != nil {
		log.Fatal().Msg(err.Error())
	}

	app.SetupLogger()
	app.SetupRoutes()
	app.ConfigEcho()

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
