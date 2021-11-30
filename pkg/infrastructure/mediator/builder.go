package mediator

import (
	"context"
)

type Builder interface {
	RegisterHandler(request Message, handler RequestHandler) Builder
	UseBehaviour(PipelineBehaviour) Builder
	Use(fn func(context.Context, Message, NextFunc)  (interface{}, error)) Builder
	Build() (*Mediator, error)
}
