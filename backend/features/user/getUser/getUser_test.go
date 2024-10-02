package getUser

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/pb33f/libopenapi-validator/responses"

	"elytra.com/backend/test"
)

func TestOpenApi(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	route := route{}
	doc := test.ParseAndBuildSpec()
	validator := responses.NewResponseBodyValidator(&doc)

	route.getUsers(c)

	success, errors := validator.ValidateResponseBody(c.Request(), rec.Result())

	if !success {
		t.Fatal(errors)
	}
}
