package mediator

import (
	"context"
)

type PipelineFunc func(context.Context, Message)  (interface{}, error)


func (p PipelineFunc) Empty() bool { return p == nil }

