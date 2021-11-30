package mediator

import (
	"context"
)

type BehaviorFunc func(context.Context, Message, NextFunc)  (interface{}, error)
