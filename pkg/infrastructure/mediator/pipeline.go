package mediator

import (
	"context"
)



type Pipeline struct {
	behaviours []BehaviorFunc
	pipeline   PipelineFunc
	handlers   map[string]RequestHandler
}

func New() *Pipeline {
	return &Pipeline{
		handlers: make(map[string]RequestHandler),
	}
}

func (p *Pipeline) UseBehaviour(behaviour PipelineBehaviour) Builder {
	return p.Use(behaviour.Process)
}

func (p *Pipeline) Use(call func(context.Context, Message, NextFunc)  (interface{}, error)) Builder {
	p.behaviours = append(p.behaviours, call)
	return p
}

func (p *Pipeline) RegisterHandler(req Message, h RequestHandler) Builder {
	key := req.Key()

	p.handlers[key] = h
	return p
}

func (p *Pipeline) Build() (*Mediator, error) {
	m := newMediator(*p)
	reverseApply(p.behaviours, m.pipe)
	return m, nil
}

func reverseApply(behaviours []BehaviorFunc,
	fn func(BehaviorFunc)) {
	for i := len(behaviours) - 1; i >= 0; i-- {
		fn(behaviours[i])
	}
}