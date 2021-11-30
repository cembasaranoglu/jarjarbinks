package mediator

import (
	"context"
)

type Sender interface {
	Send(context.Context, Message)  (interface{}, error)
}
