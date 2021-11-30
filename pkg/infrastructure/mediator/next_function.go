package mediator

import "context"

type NextFunc func(ctx context.Context)  (interface{}, error)
