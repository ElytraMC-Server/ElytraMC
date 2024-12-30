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
	Username string    `json:"name" validate:"required" example:"test"`
	Email    string    `json:"email" validate:"required,email" example:"test@gmail.com"`
}

func RegisterEditUser(e *echo.Echo, db *pgxpool.Pool) {
	route := route{db: db}

	e.PUT("/users/:id", route.editUser)
}

// EditUser godoc
// @Summary Edit user
// @Description Update user details for an existing user by UUID
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "UUID of the user" example:"123e4567-e89b-12d3-a456-426614174000"
// @Param account body editUser.EditUserRequest true "Edit user model"
// @Success 204 {string} string "User successfully updated, no content returned"
// @Failure 400 "Invalid input or payload"
// @Failure 404 "User not found"
// @Failure 500 "Internal server error"
// @Router /users/{id} [put]
func (r route) editUser(ctx echo.Context) error {
	query := database.New(r.db)

	var request EditUserRequest
	err := ctx.Bind(&request)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := ctx.Validate(&request); err != nil {
		// it's already an HTTP error
		return err
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
