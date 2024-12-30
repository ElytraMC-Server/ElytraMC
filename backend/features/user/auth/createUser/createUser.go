package createUser

import (
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"

	"elytra.com/backend/features/user/contracts"
	"elytra.com/backend/features/user/database"
	"elytra.com/backend/utils"
)

type route struct {
	db *pgxpool.Pool
}

type CreateUserRequest struct {
	Username string `json:"name" validate:"required" example:"test"`
	Email    string `json:"email" validate:"required,email" example:"test@gmail.com"`
	Password string `json:"password" validate:"required,min=8" example:"securePassword123"`
}

func RegisterPostUser(e *echo.Echo, db *pgxpool.Pool) {
	route := route{db: db}

	e.POST("/auth/register", route.createUser)
}

// CreateUser godoc
// @Summary Create user
// @Description Register a new user and return the created user information
// @Tags users
// @Accept json
// @Produce json
// @Param account body createUser.CreateUserRequest true "Create user model"
// @Success 201 {object} contracts.UserDTO "User successfully created"
// @Failure 400 "Invalid request payload or validation error"
// @Failure 500 "Internal server error"
// @Router /auth/register [post]
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

	password_hash, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal error")
	}

	request.Password = string(password_hash)

	user, err := query.CreateUser(ctx.Request().Context(), mapToParams(request))
	if err != nil {
		if utils.IsUniqueViolation(err) {
			return echo.NewHTTPError(http.StatusBadRequest, "Email already in use")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal error")
	}

	mapped := contracts.MapToDto(user)
	return ctx.JSON(201, mapped)
}
