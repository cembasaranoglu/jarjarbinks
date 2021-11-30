package http

import (
	"github.com/gofiber/fiber/v2"
	"jarjarbinks/pkg/infrastructure/ports/http/builders"
	"jarjarbinks/pkg/infrastructure/ports/http/interfaces"
	"reflect"
)

// TODO : refactor it
type httpPort struct{
	responseBuilders map[string]interfaces.ResponseBuilder
}

func (h httpPort) Present(c *fiber.Ctx, payload *interface{}, err error) error {
	builderName := "Ok"
	if err != nil{
		builderName = reflect.TypeOf(err).String()
	}
	if builder, exists := h.responseBuilders[builderName]; exists{
		return builder.Build(c, payload, err)
	}else{
		return h.responseBuilders["*errors.DatabaseError"].Build(c, payload, err)
	}
}
func NewHttpPort() interfaces.HttpPort {
	return &httpPort{
		responseBuilders: map[string]interfaces.ResponseBuilder{
			"*errors.DomainError":              builders.NewBadRequestResponseBuilder(),
			"*errors.HttpError":                builders.NewBadRequestResponseBuilder(),
			"*errors.DatabaseError":            builders.NewInternalServerErrorResponseBuilder(),
			"*errors.DoesNotExistsDomainError": builders.NewNoContentResponseBuilder(),
			"Ok":                               builders.NewOkResponseBuilder(),
		},
	}
}

