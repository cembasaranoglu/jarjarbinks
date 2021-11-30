package api

import (
	"github.com/gofiber/fiber/v2"
	"jarjarbinks/pkg/domain/commands"
	"jarjarbinks/pkg/domain/queries"
	"jarjarbinks/pkg/infrastructure/http/controller"
	"jarjarbinks/pkg/infrastructure/mediator"
	"jarjarbinks/pkg/infrastructure/ports/http/interfaces"
)

type cacheController struct {
	mediator         *mediator.Mediator
	httpPort         interfaces.HttpPort
}


func (p cacheController) findCacheEntryByKeyHttpHandler(c *fiber.Ctx) error{
	query := new(queries.FindCacheEntryByKeyQuery)
	query.EntryKey = c.Params("key")
	result, err := p.mediator.Send(c.Context(), query)
	return p.httpPort.Present(c, &result, err)
}


func (p cacheController) createCacheEntryHttpHandler(c *fiber.Ctx) error{
	command := new(commands.CreateCacheEntryCommand)
	if err := c.BodyParser(&command); err != nil{
		return err
	}else {
		result, err := p.mediator.Send(c.Context(), command)
		return p.httpPort.Present(c, &result, err)
	}
}


func (p cacheController) Endpoints() *map[string]map[string]fiber.Handler {
	return &map[string]map[string]fiber.Handler{
		"": {
			"POST":    p.createCacheEntryHttpHandler,
		},
		"/:key": {
			"GET":    p.findCacheEntryByKeyHttpHandler,
		},
	}
}

func (p cacheController) Name() string {
	return "store"
}

func (p cacheController) Prefix() string {
	return "cache"
}

func (p cacheController) Version() string {
	return "v1"
}


func NewCacheController(
	mediator *mediator.Mediator,
	httpPort interfaces.HttpPort) controller.HttpController {
	return &cacheController{
		mediator:         mediator,
		httpPort:         httpPort,
	}
}
