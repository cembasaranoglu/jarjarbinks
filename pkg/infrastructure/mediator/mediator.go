package mediator

import (
	"context"
	"errors"
)

type Mediator struct {
	context Pipeline
}

func newMediator(ctx Pipeline) *Mediator {
	return &Mediator{
		context: ctx,
	}
}

func (m *Mediator) Send(ctx context.Context, req Message)  (interface{}, error) {
	if m.context.pipeline.Empty() {
		return m.send(ctx, req)
	}
	return m.context.pipeline(ctx, req)
}

func (m *Mediator) send(ctx context.Context, req Message)  (interface{}, error) {
	key := req.Key()
	handler, ok := m.context.handlers[key]
	if !ok {
		return  nil, errors.New("handler could not be found")
	}
	return handler.Handle(ctx, req)
}

func (m *Mediator) pipe(call BehaviorFunc) {
	if m.context.pipeline.Empty() {
		m.context.pipeline = m.send
	}
	seed := m.context.pipeline

	m.context.pipeline = func(ctx context.Context, msg Message)  (interface{}, error) {
		return call(ctx, msg, func(context.Context)  (interface{}, error) { return seed(ctx, msg) })
	}
}