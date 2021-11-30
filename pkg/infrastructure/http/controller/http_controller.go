package controller

import (
	"github.com/gofiber/fiber/v2"
	"jarjarbinks/pkg/infrastructure/http/controller/interfaces"
)

type HttpController interface {
	interfaces.ControllerBase
	Endpoints() *map[string]map[string]fiber.Handler
}
