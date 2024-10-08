package main

import (
	"errors"
	"fmt"
	"os"

	"elytra.com/backend/config"
	"elytra.com/backend/database"
	"elytra.com/backend/features/user/getUser"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type App struct {
	conf  config.Config
	Echo  *echo.Echo
	state State
}

type State struct {
	db *pgxpool.Pool
}

func (app *App) setupRoutes() {
	if app.conf.Mode == config.Dev {
		registerOpenAPI(app.Echo)
	}
	getUser.RegisterGetUser(app.Echo, app.state.db)
}

func (app *App) setupLogger() {
	logger := zerolog.New(os.Stdout)
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	app.Echo.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
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

func (app *App) Run() {
	if err := app.Echo.Start(fmt.Sprintf(":%v", app.conf.AppPort)); err != nil {
		log.Fatal().Msg(err.Error())
	}
}

func NewApp(conf config.Config, echo *echo.Echo) (App, error) {
	state, err := NewState(conf)

	if err != nil {
		return App{}, errors.New(fmt.Sprintf("Couldn't initialize the app. Error: %v", err.Error()))
	}

	return App{conf: conf, Echo: echo, state: state}, nil
}

func NewState(conf config.Config) (State, error) {
	db, err := database.NewConnection(conf.DbConn)
	if err != nil {
		return State{}, err
	}

	return State{db}, nil
}
