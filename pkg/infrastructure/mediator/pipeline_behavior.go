package mediator

import (
	"context"
)

type PipelineBehaviour interface {
	Process(context.Context, Message, NextFunc) (interface{}, error)
}

