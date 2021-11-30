package builders

import (
	"github.com/gofiber/fiber/v2"
	contracts2 "jarjarbinks/pkg/domain/contracts"
	"jarjarbinks/pkg/infrastructure/ports/http/contracts"
	"jarjarbinks/pkg/infrastructure/ports/http/interfaces"
	"net/http"
)

type internalServerErrorResponseBuilder struct{
}

func (i internalServerErrorResponseBuilder) Build(c *fiber.Ctx, payload *interface{}, err error) error {
	c.Status(http.StatusBadRequest)
	return c.JSON(&contracts.ApiResponseContract{
		Messages: &[]contracts2.MessageContract{
			*contracts2.NewInternalServerErrorMessage("INTERNAL"),
		},
	})
}

func NewInternalServerErrorResponseBuilder() interfaces.ResponseBuilder{
	return &internalServerErrorResponseBuilder{
	}
}

