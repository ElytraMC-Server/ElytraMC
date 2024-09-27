//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen --config=config.yaml ../../../docs/openapi.spec.yaml
package getUser

import "github.com/labstack/echo/v4"

type route struct {
}

func RegisterGetUser(e *echo.Echo) {
	route := route{}
	RegisterHandlers(e, route)
}

type UserDTO struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
}

type Response struct {
	Users []UserDTO `json:"users"`
}

func (r route) GetUsers(ctx echo.Context) error {
	listOfUsers := []UserDTO{
		{1, "John", "Doe"},
		{2, "Jane", "Doe"},
	}

	return ctx.JSON(200, Response{
		Users: listOfUsers,
	})
}
