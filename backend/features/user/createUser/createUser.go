package createUser

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

type CreateUserRequest struct {
	Username string `json:"name" validate:"required" example:"test"`
	Email    string `json:"email" validate:"required,email" example:"test@gmail.com"`
}

func RegisterPostUser(e *echo.Echo, db *pgxpool.Pool) {
	route := route{db: db}

	e.POST("/users", route.createUser)
}

// CreateUser godoc
// @Summary Create user
// @Description Create one user
// @Tags users
// @Accept json
// @Produce json
// @Param account body createUser.CreateUserRequest true "Create user model"
// @Success 200 {object} contracts.UserDTO "User creation"
// @Failure 400 "Failed to create the user"
// @Failure 500
// @Router /users [post]
func (r route) createUser(ctx echo.Context) error {
	query := database.New(r.db)

	var request CreateUserRequest
	err := ctx.Bind(&request)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := ctx.Validate(&request); err != nil {
		// it's already an HTTP error
		return err
	}

	user, err := query.CreateUser(ctx.Request().Context(), mapToParams(request))
	if err != nil {
		if utils.IsUniqueViolation(err) {
			return echo.NewHTTPError(http.StatusBadRequest, "Email already in use")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal error")
	}

	mapped := mapToDto(user)
	return ctx.JSON(201, mapped)
}
