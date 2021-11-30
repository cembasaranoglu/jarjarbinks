package builders

import (
	"github.com/gofiber/fiber/v2"
	"jarjarbinks/pkg/infrastructure/ports/http/contracts"
	"jarjarbinks/pkg/infrastructure/ports/http/interfaces"
	"net/http"
)

type okResponseBuilder struct{
}

func (b okResponseBuilder) Build(c *fiber.Ctx, payload *interface{}, err error) error {
	c.Status(http.StatusOK)
	return c.JSON(&contracts.ApiResponseContract{
		Result: payload,
	})
}

func NewOkResponseBuilder() interfaces.ResponseBuilder{
	return &okResponseBuilder{
	}
}

