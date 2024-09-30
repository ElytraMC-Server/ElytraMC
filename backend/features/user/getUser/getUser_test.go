package getUser

import (
	"testing"

	"elytra.com/backend/test"
	"github.com/pb33f/libopenapi-validator/responses"
)

func TestOpenApi(t *testing.T) {
	doc := test.ParseAndBuildSpec()
	val := responses.NewResponseBodyValidator(&doc)
}
