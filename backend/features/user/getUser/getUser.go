package getUser

import (
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"

	"elytra.com/backend/features/user/database"
	"elytra.com/backend/utils"
)

type route struct {
	db *pgxpool.Pool
}

func RegisterGetUser(e *echo.Echo, db *pgxpool.Pool) {
	route := route{db: db}

	e.GET("/users", route.getUsers)
}

func (r route) getUsers(ctx echo.Context) error {
	query := database.New(r.db)

	users, err := query.GetUsers(ctx.Request().Context())

	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	mapped := utils.Map(users, mapToDto)

	return ctx.JSON(200, mapped)
}
