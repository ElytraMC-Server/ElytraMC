//go:generate sqlc generate
package user

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"

	"elytra.com/backend/features/user/auth/createUser"
	"elytra.com/backend/features/user/auth/loginUser"
	"elytra.com/backend/features/user/deleteUser"
	edituser "elytra.com/backend/features/user/editUser"
	"elytra.com/backend/features/user/getUser"
)

func RegisterUserRoutes(e *echo.Echo, db *pgxpool.Pool) {
	getUser.RegisterGetUser(e, db)
	createUser.RegisterPostUser(e, db)
	deleteUser.RegisterDeleteUser(e, db)
	edituser.RegisterEditUser(e, db)
	loginUser.RegisterLoginUser(e, db)
}
