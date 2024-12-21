package getUser

import (
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"

	"elytra.com/backend/features/user/contracts"
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

// GetUsers godoc
// @Summary Get users
// @Description get all users from the database
// @Tags users
// @Produce json
// @Success 200 {array} contracts.UserDTO
// @Failure 500
// @Router /users [get]
func (r route) getUsers(ctx echo.Context) error {
	query := database.New(r.db)

	users, err := query.GetUsers(ctx.Request().Context())

	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	mapped := utils.Map(users, contracts.MapToDto)

	return ctx.JSON(200, mapped)
}
