//go:generate sqlc generate
package deleteUser

import (
	"net/http"

	"elytra.com/backend/features/user/deleteUser/database"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	pgx_google_uuid "github.com/vgarvardt/pgx-google-uuid/v5"
)

type route struct {
	db *pgxpool.Pool
}

type DeleteUserRequest struct {
	ID string `json:"id"`
}

func RegisterDeleteUser(e *echo.Echo, db *pgxpool.Pool) {
	route := route{db: db}

	e.POST("/users/delete", route.deleteUser)
}

func (r route) deleteUser(ctx echo.Context) error {
	query := database.New(r.db)

	var request DeleteUserRequest
	err := ctx.Bind(&request)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	id, err := uuid.Parse(request.ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	user := query.DeleteUser(ctx.Request().Context(), pgx_google_uuid.UUID(id))
	if user != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal error")
	}

	return ctx.JSON(204, "Deleted")
}
