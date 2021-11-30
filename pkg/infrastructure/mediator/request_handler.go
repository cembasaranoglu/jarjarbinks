package mediator

import (
	"context"
)

type RequestHandler interface {
	Handle(context.Context, Message) (interface{}, error)
}
