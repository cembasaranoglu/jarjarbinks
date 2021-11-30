package interfaces

import (
	"github.com/gofiber/fiber/v2"
)

type HttpPort interface {
	Present(c *fiber.Ctx, payload *interface{}, err error) error
}
