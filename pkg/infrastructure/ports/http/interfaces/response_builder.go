package interfaces

import (
	"github.com/gofiber/fiber/v2"
)

type ResponseBuilder interface {
	Build(c *fiber.Ctx, payload *interface{}, err error) error
}
