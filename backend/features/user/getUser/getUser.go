package getUser

import "github.com/labstack/echo/v4"

type route struct {
}

func RegisterGetUser(e *echo.Echo) {
	route := route{}

	e.GET("/users", route.getUsers)
}

type UserDTO struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
}

func (r route) getUsers(ctx echo.Context) error {
	listOfUsers := []UserDTO{
		{1, "John", "Doe"},
		{2, "Jane", "Doe"},
	}

	return ctx.JSON(200, listOfUsers)
}
