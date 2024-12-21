package editUser

import (
	"net/http"

	"elytra.com/backend/features/user/database"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
)

type route struct {
	db *pgxpool.Pool
}

type EditUserRequest struct {
	ID       uuid.UUID `param:"id" validate:"required,uuid4" uri:"id" example:"ddf6aa7a-625b-4c29-93f0-faf617af5a8e"`
	Username string    `json:"username" validate:"required" example:"test"`
	Email    string    `json:"email" validate:"required,email" example:"test@gmail.com"`
}

func RegisterEditUser(e *echo.Echo, db *pgxpool.Pool) {
	route := route{db: db}

	e.PATCH("/users/:id", route.editUser)
}

// EditUser godoc
// @Summary Edit user
// @Description Edit one user if it exists
// @Tags users
// @Accept json
// @Param id path string true "UUID of the user"
// @Param account body editUser.EditUserRequest true "Edit user model"
// @Success 204 "User successfully updated, no content returned"
// @Failure 404 "User not found"
// @Failure 500
// @Router /users/{id} [patch]
func (r route) editUser(ctx echo.Context) error {
	query := database.New(r.db)

	var request EditUserRequest
	err := ctx.Bind(&request)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	rows, errUser := query.EditUser(ctx.Request().Context(), mapToParams(request))
	if rows == 0 {
		return ctx.NoContent(http.StatusNotFound)
	}

	if errUser != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal error")
	}
	return ctx.JSON(204, nil)

}
