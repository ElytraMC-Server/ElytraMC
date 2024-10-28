package user

import (
	"elytra.com/backend/app"
	"elytra.com/backend/features/user/createUser"
	"elytra.com/backend/features/user/deleteUser"
	"elytra.com/backend/features/user/getUser"
	"github.com/labstack/echo/v4"
)

func RegisterUserRoutes(e *echo.Echo, state app.State) {
	getUser.RegisterGetUser(e, state.Db)
	createUser.RegisterPostUser(e, state.Db)
	deleteUser.RegisterDeleteUser(e, state.Db)
}
