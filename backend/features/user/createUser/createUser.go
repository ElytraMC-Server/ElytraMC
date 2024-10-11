//go:generate sqlc generate
package createUser

import (
	"fmt"
	"net/http"

	"github.com/rs/zerolog/log"

	"elytra.com/backend/features/user/createUser/database"
	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
)

type route struct {
	db        *pgxpool.Pool
	validator *validator.Validate
}

type CreateUserRequest struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
}

func RegisterPostUser(e *echo.Echo, db *pgxpool.Pool, validator *validator.Validate) {
	route := route{db: db, validator: validator}

	e.POST("/users", route.createUser)
}

func (r route) createUser(ctx echo.Context) error {
	query := database.New(r.db)

	var request CreateUserRequest
	err := ctx.Bind(&request)
	if err != nil {
		return ctx.String(http.StatusBadRequest, "Bad request")
	}

	if !r.validateStruct(request) {
		return ctx.String(http.StatusBadRequest, "Bad request")
	}

	user, err := query.CreateUser(ctx.Request().Context(), mapToParams(request))
	if err != nil {
		return ctx.String(http.StatusInternalServerError, "Internal Error")
	}

	mapped := mapToDto(user)
	return ctx.JSON(201, mapped)
}

func (r route) validateStruct(request CreateUserRequest) bool {

	// returns nil or ValidationErrors ( []FieldError )
	if r.validator == nil {
		log.Fatal().Msg("text")
	}

	err := r.validator.Struct(request)
	if err != nil {

		// this check is only needed when your code could produce
		// an invalid value for validation such as interface with nil
		// value most including myself do not usually have code like this.
		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Println(err)
			return false
		}

		for _, err := range err.(validator.ValidationErrors) {

			fmt.Println(err.Namespace())
			fmt.Println(err.Field())
			fmt.Println(err.StructNamespace())
			fmt.Println(err.StructField())
			fmt.Println(err.Tag())
			fmt.Println(err.ActualTag())
			fmt.Println(err.Kind())
			fmt.Println(err.Type())
			fmt.Println(err.Value())
			fmt.Println(err.Param())
			fmt.Println()
		}

		// from here you can create your own error messages in whatever language you wish
		return false
	}

	return true
}
