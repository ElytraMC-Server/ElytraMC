package loginUser

import (
	"database/sql"
	"errors"
	"net/http"
	"time"

	"elytra.com/backend/features/user/database"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
)

type route struct {
	db *pgxpool.Pool
}

type LoginUserRequest struct {
	Username string `json:"name" validate:"required" example:"test"`
	Email    string `json:"email" validate:"required,email" example:"test@gmail.com"`
}

func RegisterLoginUser(e *echo.Echo, db *pgxpool.Pool) {
	route := route{db: db}

	e.POST("/auth/login", route.loginUser)
}

// LoginUser godoc
// @Summary Login user
// @Description Login one user
// @Tags users
// @Accept json
// @Produce json
// @Param account body loginUser.LoginUserRequest true "Login user model"
// @Success 200 "User logged in"
// @Failure 400 "Failed to login the user"
// @Failure 500
// @Router /auth/login [post]
func (r route) loginUser(ctx echo.Context) error {
	query := database.New(r.db)

	var request LoginUserRequest
	err := ctx.Bind(&request)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	rows, err := query.LoginUser(ctx.Request().Context(), mapToParams(request))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) { // Если пользователь не найден
			return ctx.JSON(http.StatusUnauthorized, map[string]string{
				"error": "invalid_credentials",
			})
		}
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal error")
	}

	// Генерация токена
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = uuid.UUID(rows.ID)
	claims["exp"] = time.Now().Add(time.Hour).Unix()

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal error")
	}

	// Возврат токена
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"user_id":    claims["user_id"],
		"token":      t,
		"expires_in": 3600,
		"token_type": "Bearer",
	})
}

