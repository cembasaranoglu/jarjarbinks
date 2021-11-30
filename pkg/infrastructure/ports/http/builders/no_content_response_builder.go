package builders

import (
	"github.com/gofiber/fiber/v2"
	"jarjarbinks/pkg/infrastructure/ports/http/contracts"
	"jarjarbinks/pkg/infrastructure/ports/http/interfaces"
	"net/http"
)

type noContentResponseBuilder struct{
}

func (b noContentResponseBuilder) Build(c *fiber.Ctx, payload *interface{}, err error) error {
	c.Status(http.StatusNoContent)
	return c.JSON(&contracts.ApiResponseContract{})
}

func NewNoContentResponseBuilder() interfaces.ResponseBuilder{
	return &noContentResponseBuilder{
	}
}

