//go:generate sqlc generate
package createUser

import (
	"net/http"

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

func RegisterPostUser(e *echo.Echo, db *pgxpool.Pool) {
	route := route{db: db}

	e.POST("/users", route.createUser)
}

func (r route) createUser(ctx echo.Context) error {
	query := database.New(r.db)

	var request CreateUserRequest
	err := ctx.Bind(&request)
	if err != nil {
		return ctx.String(http.StatusBadRequest, "Bad request")
	}

	if err := ctx.Validate(&request); err != nil {
		// it's already an HTTP error
		return err
	}

	user, err := query.CreateUser(ctx.Request().Context(), mapToParams(request))
	if err != nil {
		return ctx.String(http.StatusInternalServerError, "Internal Error")
	}

	mapped := mapToDto(user)
	return ctx.JSON(201, mapped)
}
