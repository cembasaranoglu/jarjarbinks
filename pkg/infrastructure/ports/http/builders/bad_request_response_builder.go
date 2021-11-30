package builders

import (
	"github.com/gofiber/fiber/v2"
	contracts2 "jarjarbinks/pkg/domain/contracts"
	"jarjarbinks/pkg/domain/errors"
	"jarjarbinks/pkg/infrastructure/ports/http/contracts"
	"jarjarbinks/pkg/infrastructure/ports/http/interfaces"
	"net/http"
)

type badRequestResponseBuilder struct{
}

func (b badRequestResponseBuilder) Build(c *fiber.Ctx, payload *interface{}, err error) error {
	domainError := err.(*errors.DomainError)
	c.Status(http.StatusBadRequest)
	return c.JSON(&contracts.ApiResponseContract{
		Messages: &[]contracts2.MessageContract{
			*contracts2.NewBadRequestErrorMessage(domainError.Code, domainError.Message),
		},
	})
}

func NewBadRequestResponseBuilder() interfaces.ResponseBuilder{
	return &badRequestResponseBuilder{}
}
