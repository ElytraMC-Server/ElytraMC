package docs

import "github.com/labstack/echo/v4"

type route struct {
	docsPath string
}

func RegisterDocs(e *echo.Echo, docsPath string) {
	r := route{docsPath}

	e.GET("/docs", r.ui)
	e.File("/spec.yaml", docsPath)
}

func (r *route) ui(ctx echo.Context) error {
	return ctx.HTML(200, `
<!doctype html>
<html>
  <head>
    <title>Scalar API Reference</title>
    <meta charset="utf-8" />
    <meta
      name="viewport"
      content="width=device-width, initial-scale=1" />
  </head>
  <body>
    <!-- Need a Custom Header? Check out this example https://codepen.io/scalarorg/pen/VwOXqam -->
    <script
      id="api-reference"
      data-url="/spec.yaml"></script>
    <script src="https://cdn.jsdelivr.net/npm/@scalar/api-reference"></script>
  </body>
</html>
	`)
}
