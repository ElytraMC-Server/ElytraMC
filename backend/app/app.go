package app

import (
	"errors"
	"fmt"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"elytra.com/backend/config"
	"elytra.com/backend/database"
	"elytra.com/backend/docs"
	"elytra.com/backend/features/user"
)

type App struct {
	conf  config.Config
	Echo  *echo.Echo
	state State
}

type State struct {
	db        *pgxpool.Pool
	validator *validator.Validate
}

func registerOpenAPI(e *echo.Echo) {
	docs.RegisterDocs(e, "./docs/spec.yaml")
	e.File("/elytra.png", "./elytra.png")
}

func (app *App) SetupRoutes() {
	if app.conf.Mode == config.Dev {
		registerOpenAPI(app.Echo)
	}
	user.RegisterUserRoutes(app.Echo, app.state.db)
}

func (app *App) SetupLogger() {
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

func (app *App) ConfigEcho() {
	app.Echo.Validator = &RequestValidator{
		validator: app.state.validator,
	}
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
	validator := validator.New(validator.WithRequiredStructEnabled())

	return State{db, validator}, nil
}
