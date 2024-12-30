package deleteUser

import (
	"net/http"

	"elytra.com/backend/features/user/database"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	pgx_google_uuid "github.com/vgarvardt/pgx-google-uuid/v5"
)

type route struct {
	db *pgxpool.Pool
}

type DeleteUserRequest struct {
	ID uuid.UUID `param:"id" validate:"required,uuid4" uri:"id" example:"ddf6aa7a-625b-4c29-93f0-faf617af5a8e"`
}

func RegisterDeleteUser(e *echo.Echo, db *pgxpool.Pool) {
	route := route{db: db}

	e.DELETE("/users/:id", route.deleteUser)
}

// DeleteUser godoc
// @Summary Delete user
// @Description Delete a user by their UUID if they exist
// @Tags users
// @Param id path string true "UUID of the user to be deleted" example:"123e4567-e89b-12d3-a456-426614174000"
// @Success 204 {string} string "User successfully deleted, no content returned"
// @Failure 400 "Invalid request payload or UUID format"
// @Failure 404 "User not found"
// @Failure 500 "Internal server error"
// @Router /users/{id} [delete]
func (r route) deleteUser(ctx echo.Context) error {
	query := database.New(r.db)

	var request DeleteUserRequest
	err := ctx.Bind(&request)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := ctx.Validate(&request); err != nil {
		return err
	}

	rows, errUser := query.DeleteUser(ctx.Request().Context(), pgx_google_uuid.UUID(request.ID))

	if rows == 0 {
		return ctx.NoContent(http.StatusNotFound)
	}

	if errUser != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal error")
	}

	return ctx.JSON(204, nil)
}
